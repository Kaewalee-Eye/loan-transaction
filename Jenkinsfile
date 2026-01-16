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

    stage('Buildx Init') {
      steps {
        sh '''
          docker buildx create --name multiarch --use >/dev/null 2>&1 || docker buildx use multiarch
          docker buildx inspect --bootstrap
        '''
      }
    }

    stage('Docker Login') {
      steps {
        withCredentials([usernamePassword(credentialsId: DOCKER_CREDS, usernameVariable: 'DH_USER', passwordVariable: 'DH_PASS')]) {
          sh 'echo "$DH_PASS" | docker login -u "$DH_USER" --password-stdin'
        }
      }
    }

    stage('Build & Push (linux/amd64)') {
      steps {
        sh """
          docker buildx build --platform linux/amd64 \
            -t ${DOCKERHUB_REPO}:${IMAGE_TAG} \
            -t ${DOCKERHUB_REPO}:latest \
            --push .
        """
      }
    }

    stage('Logout') {
      steps { sh 'docker logout || true' }
    }
  }
}
