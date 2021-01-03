// Code generated by MockGen. DO NOT EDIT.
// Source: ops.go

// Package ops is a generated GoMock package.
package ops

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	inventory_client "github.com/openshift/assisted-installer/src/inventory_client"
)

// MockOps is a mock of Ops interface
type MockOps struct {
	ctrl     *gomock.Controller
	recorder *MockOpsMockRecorder
}

// MockOpsMockRecorder is the mock recorder for MockOps
type MockOpsMockRecorder struct {
	mock *MockOps
}

// NewMockOps creates a new mock instance
func NewMockOps(ctrl *gomock.Controller) *MockOps {
	mock := &MockOps{ctrl: ctrl}
	mock.recorder = &MockOpsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOps) EXPECT() *MockOpsMockRecorder {
	return m.recorder
}

// ExecPrivilegeCommand mocks base method
func (m *MockOps) ExecPrivilegeCommand(liveLogger io.Writer, command string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{liveLogger, command}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecPrivilegeCommand", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecPrivilegeCommand indicates an expected call of ExecPrivilegeCommand
func (mr *MockOpsMockRecorder) ExecPrivilegeCommand(liveLogger, command interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{liveLogger, command}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecPrivilegeCommand", reflect.TypeOf((*MockOps)(nil).ExecPrivilegeCommand), varargs...)
}

// ExecCommand mocks base method
func (m *MockOps) ExecCommand(liveLogger io.Writer, command string, args ...string) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{liveLogger, command}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecCommand", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecCommand indicates an expected call of ExecCommand
func (mr *MockOpsMockRecorder) ExecCommand(liveLogger, command interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{liveLogger, command}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecCommand", reflect.TypeOf((*MockOps)(nil).ExecCommand), varargs...)
}

// Mkdir mocks base method
func (m *MockOps) Mkdir(dirName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", dirName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir
func (mr *MockOpsMockRecorder) Mkdir(dirName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockOps)(nil).Mkdir), dirName)
}

// WriteImageToDisk mocks base method
func (m *MockOps) WriteImageToDisk(ignitionPath, device string, progressReporter inventory_client.InventoryClient, extra []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteImageToDisk", ignitionPath, device, progressReporter, extra)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteImageToDisk indicates an expected call of WriteImageToDisk
func (mr *MockOpsMockRecorder) WriteImageToDisk(ignitionPath, device, progressReporter, extra interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteImageToDisk", reflect.TypeOf((*MockOps)(nil).WriteImageToDisk), ignitionPath, device, progressReporter, extra)
}

// Reboot mocks base method
func (m *MockOps) Reboot() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reboot")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reboot indicates an expected call of Reboot
func (mr *MockOpsMockRecorder) Reboot() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reboot", reflect.TypeOf((*MockOps)(nil).Reboot))
}

// ExtractFromIgnition mocks base method
func (m *MockOps) ExtractFromIgnition(ignitionPath, fileToExtract string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtractFromIgnition", ignitionPath, fileToExtract)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExtractFromIgnition indicates an expected call of ExtractFromIgnition
func (mr *MockOpsMockRecorder) ExtractFromIgnition(ignitionPath, fileToExtract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtractFromIgnition", reflect.TypeOf((*MockOps)(nil).ExtractFromIgnition), ignitionPath, fileToExtract)
}

// SystemctlAction mocks base method
func (m *MockOps) SystemctlAction(action string, args ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{action}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SystemctlAction", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SystemctlAction indicates an expected call of SystemctlAction
func (mr *MockOpsMockRecorder) SystemctlAction(action interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{action}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SystemctlAction", reflect.TypeOf((*MockOps)(nil).SystemctlAction), varargs...)
}

// PrepareController mocks base method
func (m *MockOps) PrepareController() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareController")
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareController indicates an expected call of PrepareController
func (mr *MockOpsMockRecorder) PrepareController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareController", reflect.TypeOf((*MockOps)(nil).PrepareController))
}

// GetVGByPV mocks base method
func (m *MockOps) GetVGByPV(pvName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVGByPV", pvName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVGByPV indicates an expected call of GetVGByPV
func (mr *MockOpsMockRecorder) GetVGByPV(pvName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVGByPV", reflect.TypeOf((*MockOps)(nil).GetVGByPV), pvName)
}

// RemoveVG mocks base method
func (m *MockOps) RemoveVG(vgName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVG", vgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveVG indicates an expected call of RemoveVG
func (mr *MockOpsMockRecorder) RemoveVG(vgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVG", reflect.TypeOf((*MockOps)(nil).RemoveVG), vgName)
}

// RemoveLV mocks base method
func (m *MockOps) RemoveLV(lvName, vgName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLV", lvName, vgName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLV indicates an expected call of RemoveLV
func (mr *MockOpsMockRecorder) RemoveLV(lvName, vgName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLV", reflect.TypeOf((*MockOps)(nil).RemoveLV), lvName, vgName)
}

// RemovePV mocks base method
func (m *MockOps) RemovePV(pvName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePV", pvName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePV indicates an expected call of RemovePV
func (mr *MockOpsMockRecorder) RemovePV(pvName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePV", reflect.TypeOf((*MockOps)(nil).RemovePV), pvName)
}

// Wipefs mocks base method
func (m *MockOps) Wipefs(device string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wipefs", device)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wipefs indicates an expected call of Wipefs
func (mr *MockOpsMockRecorder) Wipefs(device interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wipefs", reflect.TypeOf((*MockOps)(nil).Wipefs), device)
}

// GetMCSLogs mocks base method
func (m *MockOps) GetMCSLogs() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMCSLogs")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMCSLogs indicates an expected call of GetMCSLogs
func (mr *MockOpsMockRecorder) GetMCSLogs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMCSLogs", reflect.TypeOf((*MockOps)(nil).GetMCSLogs))
}

// UploadInstallationLogs mocks base method
func (m *MockOps) UploadInstallationLogs(isBootstrap bool) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadInstallationLogs", isBootstrap)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadInstallationLogs indicates an expected call of UploadInstallationLogs
func (mr *MockOpsMockRecorder) UploadInstallationLogs(isBootstrap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadInstallationLogs", reflect.TypeOf((*MockOps)(nil).UploadInstallationLogs), isBootstrap)
}

// ReloadHostFile mocks base method
func (m *MockOps) ReloadHostFile(filepath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReloadHostFile", filepath)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReloadHostFile indicates an expected call of ReloadHostFile
func (mr *MockOpsMockRecorder) ReloadHostFile(filepath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReloadHostFile", reflect.TypeOf((*MockOps)(nil).ReloadHostFile), filepath)
}

// CreateOpenshiftSshManifest mocks base method
func (m *MockOps) CreateOpenshiftSshManifest(filePath, template, sshPubKeyPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOpenshiftSshManifest", filePath, template, sshPubKeyPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOpenshiftSshManifest indicates an expected call of CreateOpenshiftSshManifest
func (mr *MockOpsMockRecorder) CreateOpenshiftSshManifest(filePath, template, sshPubKeyPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOpenshiftSshManifest", reflect.TypeOf((*MockOps)(nil).CreateOpenshiftSshManifest), filePath, template, sshPubKeyPath)
}

// GetMustGatherLogs mocks base method
func (m *MockOps) GetMustGatherLogs(workDir, kubeconfigPath string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMustGatherLogs", workDir, kubeconfigPath)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMustGatherLogs indicates an expected call of GetMustGatherLogs
func (mr *MockOpsMockRecorder) GetMustGatherLogs(workDir, kubeconfigPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMustGatherLogs", reflect.TypeOf((*MockOps)(nil).GetMustGatherLogs), workDir, kubeconfigPath)
}

// CreateRandomHostname mocks base method
func (m *MockOps) CreateRandomHostname(hostname string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRandomHostname", hostname)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRandomHostname indicates an expected call of CreateRandomHostname
func (mr *MockOpsMockRecorder) CreateRandomHostname(hostname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRandomHostname", reflect.TypeOf((*MockOps)(nil).CreateRandomHostname), hostname)
}

// GetHostname mocks base method
func (m *MockOps) GetHostname() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostname")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostname indicates an expected call of GetHostname
func (mr *MockOpsMockRecorder) GetHostname() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostname", reflect.TypeOf((*MockOps)(nil).GetHostname))
}
