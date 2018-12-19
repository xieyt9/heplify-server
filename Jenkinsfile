#!groovy​

pipeline {
  agent any
  stages {
    stage('build2pub') {
      steps {
        sh "build/make-build-image.sh TRUE"
      }
    }
  }

  post {
    always {
      echo 'One way or another, I have finished'
      dir("${env.WORKSPACE}/jenkins") {
        deleteDir() /* clean up our workspace */
      }
    }
  }
}