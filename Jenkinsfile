pipeline {
  agent any

  tools {
    go 'golang' // name must match what you configured in Jenkins tools
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

    stage('Archive Binary') {
      steps {
        archiveArtifacts artifacts: 'bin/**', fingerprint: true
      }
    }
  }
}
