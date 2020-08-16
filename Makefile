mod:
	go mod init
	go get k8s.io/client-go@v0.18.8 
build:
	go build cmd/client-object/client-go-demo.go
run:
	go run cmd/client-object/client-go-demo.go