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
                        export PATH=$PATH:/opt/homebrew/bin:/usr/local/bin
                        
                        IMAGE_TAG="asia-northeast3-docker.pkg.dev/dataengcamp-427114/snippetbox-app/snippetbox-app:latest"
                        
                        gcloud auth activate-service-account --key-file="$GCLOUD_CREDS"
                        gcloud auth configure-docker asia-northeast3-docker.pkg.dev --quiet
                        
                        echo "Building image with tag: $IMAGE_TAG"
                        docker build . --file Dockerfile --tag $IMAGE_TAG
                        
                        docker images | grep snippetbox-app
                        
                        echo "Pushing image with tag: $IMAGE_TAG"
                        docker push $IMAGE_TAG
                    '''
                }
            }
        }
    }
}