#!groovy​

pipeline {
  agent any
  stages {
    stage('build2pub') {
      steps {
        sh "heplify-server/build/make-build-image.sh TRUE"
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