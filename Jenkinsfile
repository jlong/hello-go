pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        sh 'make tests'
      }
    }
    stage('Tests') {
      steps {
        sh 'make tests'
      }
    }
  }
}