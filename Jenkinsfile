#!groovy​

pipeline {
  agent any

  parameters {
    choice(choices: "hepsrv\n"+
      "hepipe", description: 'product type', name: 'HEP_TYPE')
  }

  stages {
    stage('build2pub') {
      steps {
        sh "build/make-build-image.sh ${params.HEP_TYPE}"
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