pipeline {
  agent any
  options {
    skipDefaultCheckout(true)
  }
  stages {
    stage('Clean') {
      steps {
        deleteDir()
      }
    }
    stage('Checkout') {
      steps {
        checkout scm
      }
    }
    // stage ต่อ ๆ ไป
  }
}