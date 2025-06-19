pipeline {
  agent any

  tools {
    go 'golang'
  }

  environment {
    GOPATH = "${WORKSPACE}/go"
    PATH = "${GOPATH}/bin:${PATH}"
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        sh 'go build -o bin/myapp ./...'
      }
    }

    stage('Test') {
      steps {
        sh 'go test ./...'
      }
    }

    stage('Publish to Nexus') {
      steps {
        nexusArtifactUploader(
          nexusVersion:   'nexus3',
          protocol:       'http',
          nexusUrl:       "10.1.1.129:8081/repository/golang-builds",
          credentialsId:  'nexus-creds',           
          groupId:        'com.mycompany.go',     
          version:        "1.0.${env.BUILD_NUMBER}", 
          repository:     'go-binaries',           
          artifacts: [
            [artifactId: 'myapp',
             classifier: '',
             file:       'bin/myapp',
             type:       'binary']
          ]
        )
      }
    }
  }
}
