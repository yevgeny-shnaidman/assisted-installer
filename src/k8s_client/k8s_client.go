package k8s_client

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/openshift/assisted-installer/src/utils"
	"k8s.io/apimachinery/pkg/labels"

	"k8s.io/api/certificates/v1beta1"

	"github.com/openshift/assisted-installer/src/ops"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"

	operatorv1 "github.com/openshift/client-go/operator/clientset/versioned"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	certificatesv1beta1client "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
)

//go:generate mockgen -source=k8s_client.go -package=k8s_client -destination=mock_k8s_client.go
type K8SClient interface {
	ListMasterNodes() (*v1.NodeList, error)
	PatchEtcd() error
	UnPatchEtcd() error
	ListNodes() (*v1.NodeList, error)
	RunOCctlCommand(args []string, kubeconfigPath string, o ops.Ops) (string, error)
	ApproveCsr(csr *v1beta1.CertificateSigningRequest) error
	ListCsrs() (*v1beta1.CertificateSigningRequestList, error)
	GetConfigMap(namespace string, name string) (*v1.ConfigMap, error)
	GetPodLogs(namespace string, podName string, sinceSeconds int64) (string, error)
	GetPods(namespace string, labelMatch map[string]string) ([]v1.Pod, error)
}

type K8SClientBuilder func(configPath string, logger *logrus.Logger) (K8SClient, error)

type k8sClient struct {
	log      *logrus.Logger
	client   *kubernetes.Clientset
	ocClient *operatorv1.Clientset
	// CertificateSigningRequestInterface is interface
	csrClient certificatesv1beta1client.CertificateSigningRequestInterface
}

func NewK8SClient(configPath string, logger *logrus.Logger) (K8SClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return &k8sClient{}, errors.Wrap(err, "loading kubeconfig")
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &k8sClient{}, errors.Wrap(err, "creating a Kubernetes client")
	}
	ocClient, err := operatorv1.NewForConfig(config)
	if err != nil {
		return &k8sClient{}, errors.Wrap(err, "creating a Kubernetes client")
	}
	csrClient := client.CertificatesV1beta1().CertificateSigningRequests()

	return &k8sClient{logger, client, ocClient, csrClient}, nil
}

func (c *k8sClient) ListMasterNodes() (*v1.NodeList, error) {
	nodes, err := c.client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{LabelSelector: "node-role.kubernetes.io/master"})
	if err != nil {
		return &v1.NodeList{}, err
	}
	return nodes, nil
}

func (c *k8sClient) ListNodes() (*v1.NodeList, error) {
	nodes, err := c.client.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return &v1.NodeList{}, err
	}
	return nodes, nil
}

func (c *k8sClient) PatchEtcd() error {
	c.log.Info("Patching etcd")
	data := []byte(`{"spec": {"unsupportedConfigOverrides": {"useUnsupportedUnsafeNonHANonProductionUnstableEtcd": true}}}`)
	result, err := c.ocClient.OperatorV1().Etcds().Patch(context.Background(), "cluster", types.MergePatchType, data, metav1.PatchOptions{})
	if err != nil {
		return errors.Wrap(err, "Failed to patch etcd")
	}
	c.log.Info(result)
	return nil
}

func (c *k8sClient) UnPatchEtcd() error {
	c.log.Info("UnPatching etcd")
	data := []byte(`{"spec": {"unsupportedConfigOverrides": null}}`)
	result, err := c.ocClient.OperatorV1().Etcds().Patch(context.Background(), "cluster", types.MergePatchType, data, metav1.PatchOptions{})
	if err != nil {
		return errors.Wrap(err, "Failed to unpatch etcd")
	}
	c.log.Info(result)
	return nil
}

func (c *k8sClient) RunOCctlCommand(args []string, kubeconfigPath string, o ops.Ops) (string, error) {
	c.log.Infof("Running oc command with args %v", args)
	args = append([]string{fmt.Sprintf("--kubeconfig=%s", kubeconfigPath)}, args...)
	outPut, err := o.ExecPrivilegeCommand(utils.NewLogWriter(c.log), "oc", args...)
	if err != nil {
		return "", err
	}
	return outPut, nil
}

func (c k8sClient) ListCsrs() (*v1beta1.CertificateSigningRequestList, error) {
	csrs, err := c.csrClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		c.log.Errorf("Failed to get list of csrs. err : %e", err)
		return nil, err
	}
	return csrs, nil
}

func (c k8sClient) ApproveCsr(csr *v1beta1.CertificateSigningRequest) error {

	csr.Status.Conditions = append(csr.Status.Conditions, certificatesv1beta1.CertificateSigningRequestCondition{
		Type:           certificatesv1beta1.CertificateApproved,
		Reason:         "NodeCSRApprove",
		Message:        "This CSR was approved by the assisted-installer-controller",
		LastUpdateTime: metav1.Now(),
	})
	if _, err := c.csrClient.UpdateApproval(context.TODO(), csr, metav1.UpdateOptions{}); err != nil {
		c.log.Errorf("Failed to approve csr %v, err %e", csr, err)
		return err
	}
	return nil
}

func (c *k8sClient) GetConfigMap(namespace string, name string) (*v1.ConfigMap, error) {
	cm, err := c.client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (c *k8sClient) GetPods(namespace string, labelMatch map[string]string) ([]v1.Pod, error) {
	listOptions := metav1.ListOptions{}
	if labelMatch != nil {
		labelSelector := metav1.LabelSelector{MatchLabels: labelMatch}
		listOptions.LabelSelector = labels.Set(labelSelector.MatchLabels).String()
	}

	pod, err := c.client.CoreV1().Pods(namespace).List(context.TODO(), listOptions)
	if err != nil {
		return nil, err
	}

	return pod.Items, nil
}

func (c *k8sClient) GetPodLogs(namespace string, podName string, sinceSeconds int64) (string, error) {
	podLogOpts := v1.PodLogOptions{}
	if sinceSeconds > 0 {
		podLogOpts.SinceSeconds = &sinceSeconds
	}
	req := c.client.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
