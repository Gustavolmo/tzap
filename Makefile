buildCli: gomodtidy
	cd cli && make release


exGithubDoc:
	go run example/githubdoc/main.go
exMadebygpt:
	go run example/madebygpt/main.go
exRefactoring:
	go run example/refactoring/main.go
exTesting:
	go run example/testing/main.go

gomodtidy:
	go mod tidy
	cd pkg/connectors/openaiconnector && go mod tidy
	cd pkg/tzapconnect && go mod tidy
	cd pkg/connectors/googlevoiceconnector && go mod tidy
	cd examples && go mod tidy
	cd cli && go mod tidy
	go work sync

ts-build:
	cd ts && npm run build