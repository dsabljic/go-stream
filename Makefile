.PHONY: cover
cover:
	go test ./... -coverprofile=coverage.out && ./coverage_cleanup.sh exclude.txt && go tool cover -html=coverage.out -o coverage.html