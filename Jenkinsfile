pipeline {
    agent any
    tools { go '1.24.2' }
    stages {
        stage('Test') {
            steps {
                // doubt
                sh 'go test ./...'
            }
        }
    }
}