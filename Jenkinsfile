pipeline {

  environment {
        INSTALLER = 'quay.io/ocpmetal/assisted-installer'
        CONTROLLER = 'quay.io/ocpmetal/assisted-installer-controller'
  }
  agent {
    node {
      label 'host'
    }
  }
  stages {
    stage('build') {
      steps {
        sh 'skipper make'
      }
    }


  stage('publish images on push to master') {
              when {
                  branch 'osher/test'
              }

              steps {
                  withCredentials([usernamePassword(credentialsId: 'ocpmetal_cred', passwordVariable: 'PASS', usernameVariable: 'USER')]) {
                      sh '''docker login quay.io -u $USER -p $PASS'''
                  }

                  sh '''docker tag  ${INSTALLER} ${INSTALLER}:osher_test'''
                  sh '''docker tag  ${INSTALLER} ${INSTALLER}:${GIT_COMMIT}'''
                  sh '''docker push ${INSTALLER}:osher_test'''
                  sh '''docker push ${INSTALLER}:${GIT_COMMIT}'''

                  sh '''docker tag  ${CONTROLLER} ${CONTROLLER}:osher_test'''
                  sh '''docker tag  ${CONTROLLER} ${CONTROLLER}:${GIT_COMMIT}'''
                  sh '''docker push ${CONTROLLER}:osher_test'''
                  sh '''docker push ${CONTROLLER}:${GIT_COMMIT}'''
              }
   }
}
}