# Golang Cli Tool for Storing Content to S3


I want to create a tool called data-bin.

- it is written in Golang
- it will run in Aws
- it will be deployed by terraform
- it will have a Makefile for all the tasks to build, test, deploy
- it is immplemented as a Aws Lambda function which accepts the payload as POST request
- the traffic is encrypted https
- each request will have a POST body and this will be stored as an object in a dedicated S3 object
- project structure will be standard with deployment for the aws deploy stuff and there will be go.mod at root
- there will be cmd/data-bin folder with main.go
- there will be internal with sub packages folder
