pipeline {
  agent any

  environment {
    IMAGE_NAME = "eyekaewalee/loan-api"
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Docker Build (linux/amd64)') {
      steps {
        sh '''
          docker build --platform=linux/amd64 \
            -t $IMAGE_NAME:${GIT_COMMIT} \
            -t $IMAGE_NAME:latest .
        '''
      }
    }

    stage('Docker Login & Push') {
      steps {
        withCredentials([usernamePassword(
          credentialsId: 'dockerhub-eyekaewalee',
          usernameVariable: 'DOCKER_USER',
          passwordVariable: 'DOCKER_PASS'
        )]) {
          sh '''
            echo "$DOCKER_PASS" | docker login -u "$DOCKER_USER" --password-stdin
            docker push $IMAGE_NAME:${GIT_COMMIT}
            docker push $IMAGE_NAME:latest
          '''
        }
      }
    }
  }
}
