part := 1
dataset := full

day%:
	@echo "Running day$*..."
	go run main.go $* $(part) $(dataset)