pipeline {
    agent any
    tools { go '1.24' }
    stages {
        stage('Test') {
            steps {
                // doubt
                sh 'go test ./...'
            }
        }
    }
}