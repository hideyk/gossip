server:
	go run cmd/server/main.go port=$(port) 

client:
	go run cmd/client/main.go port=$(port)