pipeline {
    agent any
    stages {
        stage('Test') {
            agent {
                docker { image 'golang:1.24' }
            }
            steps {
                sh 'go test ./...'
            }
        }
    }
}