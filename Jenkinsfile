pipeline {
  agent any

  environment {
    DOCKERHUB_REPO = "eyekaewalee/loan-api"
    DOCKER_CREDS   = "dockerhub-eyekaewalee"
    IMAGE_TAG      = "${env.GIT_COMMIT}"
  }

  stages {
    stage('Clean') {
      steps { deleteDir() }
    }

    stage('Checkout') {
      steps { checkout scm }
    }

    stage('Docker Build (linux/amd64)') {
      steps {
        sh """
          docker build --platform=linux/amd64 \
            -t ${DOCKERHUB_REPO}:${IMAGE_TAG} \
            -t ${DOCKERHUB_REPO}:latest \
            .
        """
      }
    }

    stage('Docker Login & Push') {
      steps {
        withCredentials([usernamePassword(credentialsId: DOCKER_CREDS, usernameVariable: 'DH_USER', passwordVariable: 'DH_PASS')]) {
          sh """
            echo "$DH_PASS" | docker login -u "$DH_USER" --password-stdin
            docker push ${DOCKERHUB_REPO}:${IMAGE_TAG}
            docker push ${DOCKERHUB_REPO}:latest
            docker logout
          """
        }
      }
    }
  }
}
