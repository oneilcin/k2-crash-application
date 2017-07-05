// -*- mode: groovy -*-
// Jenkins pipeline
// See documents at https://jenkins.io/doc/book/pipeline/jenkinsfile/

podTemplate(label: 'k2-crash-application', containers: [
        containerTemplate(name: 'jnlp', image: 'jenkinsci/jnlp-slave:2.62-alpine', args: '${computer.jnlpmac} ${computer.name}'),
        containerTemplate(name: 'golang', image: 'golang:latest', ttyEnabled: true, command: 'cat'),
]) {
    node('k2-crash-application') {
        container('golang') {
            stage('Checkout') {
                echo 'Checking out....'
                checkout scm
                sh 'go version'
            }

            stage('Build') {
                echo 'Building....'
                sh 'go get -v -d -t ./... || true'
                sh 'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v'
            }
            stage('Test') {
                echo 'Building....'
                sh 'TF_ACC=Y go test -v'
            }
            stage('Deploy') {
                echo 'Deploying...'
            }
        }
    }
}
