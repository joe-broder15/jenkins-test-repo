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

    stage('Publish to Nexus') {
      steps {
        nexusArtifactUploader(
          nexusVersion:   'nexus3',
          protocol:       'http',
          nexusUrl:       "10.1.1.129:8081/repository/golang-builds",
          credentialsId:  'nexus-creds',           // your Jenkins creds ID
          groupId:        'com.mycompany.go',      // adjust path prefix
          version:        "1.0.${env.BUILD_NUMBER}", 
          repository:     'go-binaries',           // your raw-hosted repo name
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
