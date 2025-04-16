pipeline {
    agent any
    tools { go '1.19' }
    stages {
        stage('Test') {
            steps {
                // doubt
                sh 'go test ./...'
            }
        }
    }
}