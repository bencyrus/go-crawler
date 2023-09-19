# Default target to run when just executing `make`
.DEFAULT_GOAL := run

# Start both frontend and backend services
run:
	@go run *.go