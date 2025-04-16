pipeline {
    agent any
    environment {
        PATH = "$PATH:/opt/homebrew/bin:/usr/local/bin"
    }
    stages {
        stage('Run Tests') {
            tools { go '1.24.2' }
            steps {
                sh 'go test ./...'
            }
        }

        stage('Auth GCP') {
            steps {
                withCredentials([file(credentialsId: 'gcloud-creds', variable: 'GCLOUD_CREDS')]) {
                    sh '''
                        gcloud auth activate-service-account --key-file="$GCLOUD_CREDS"
                        gcloud auth configure-docker asia-northeast3-docker.pkg.dev --quiet
                    '''
                }
            }
        }

        stage('Build Image') {
            steps {
                script {
                    env.IMAGE_TAG = "asia-northeast3-docker.pkg.dev/dataengcamp-427114/snippetbox-app/snippetbox-app:latest"
                }
                sh '''
                    echo "Building image with tag: $IMAGE_TAG"
                    docker build . --file Dockerfile --tag $IMAGE_TAG
                '''
            }
        }

        stage('Push Image') {
            steps {
                sh '''
                    echo "Pushing image with tag: $IMAGE_TAG"
                    docker push $IMAGE_TAG
                '''
            }
        }
    }
}