run: 
	go run main.go

mocks:
	@mockery --name="ProductServiceType" --dir="./service" --output="./controllers/mocks"

tests:
	go test ./...