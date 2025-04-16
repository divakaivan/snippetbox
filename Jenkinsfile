pipeline {
    agent any
    tools { go '1.24.2' }
    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        },
        stage('Push to Artifact Registry') {
            steps {
                withCredentials([file(credentialsId: 'gcloud-creds', variable: 'GCLOUD_CREDS')]) {
                    sh '''
                        export PATH=$PATH:/opt/homebrew/bin
                        
                        gcloud auth activate-service-account --key-file="$GCLOUD_CREDS"
                        gcloud auth configure-docker asia-northeast3-docker.pkg.dev --quiet
                    '''
                }
            }
        }
    }
}