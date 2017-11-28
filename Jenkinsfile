pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'make hello'
      }
    }
    stage('Tests') {
      steps {
        sh 'make tests'
      }
    }
  }
}