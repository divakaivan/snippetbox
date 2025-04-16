pipeline {
    agent any
    tools { go '1.24.2' }
    stages {
        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Push to Artifact Registry') {
            steps {
                withCredentials([file(credentialsId: 'gcloud-creds', variable: 'GCLOUD_CREDS')]) {
                    sh '''
                        export PATH=$PATH:/opt/homebrew/bin
                        
                        gcloud auth activate-service-account --key-file="$GCLOUD_CREDS"
                        gcloud auth configure-docker asia-northeast3-docker.pkg.dev --quiet
                        docker build . --file Dockerfile --tag asia-northeast3/snippetbox-app:latest
                        docker push asia-northeast3-docker.pkg.dev/snippetbox/snippetbox-app/snippetbox-app:latest
                    '''
                }
            }
        }
    }
}