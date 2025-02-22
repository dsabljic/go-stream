.PHONY: prep
prep:
	go test ./... -coverprofile=coverage.out && ./coverage_cleanup.sh exclude_files.txt && go tool cover -html=coverage.out -o coverage.html