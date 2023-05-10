build:
	@go build -o awsIAM

run: build
	@./awsIAM