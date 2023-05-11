module github.com/tzapio/tzap/example

go 1.20

replace github.com/tzapio/tzap => ../

replace github.com/tzapio/tzap/pkg/connectors/openaiconnector => ../pkg/connectors/openaiconnector

replace github.com/tzapio/tzap/pkg/connectors/redisembeddbconnector => ../pkg/connectors/redisembeddbconnector

replace github.com/tzapio/tzap/pkg/tzapconnect => ../pkg/tzapconnect

replace github.com/tzapio/tzap/cli => ../cli

require (
	github.com/tzapio/tzap v0.0.0-00010101000000-000000000000
	github.com/tzapio/tzap/cli v0.0.0-00010101000000-000000000000
	github.com/tzapio/tzap/pkg/tzapconnect v0.0.0-00010101000000-000000000000
)

require (
	github.com/dlclark/regexp2 v1.9.0 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/sashabaranov/go-openai v1.9.3 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/tiktoken-go/tokenizer v0.1.0 // indirect
	github.com/tzapio/tzap/pkg/connectors/openaiconnector v0.0.0-00010101000000-000000000000 // indirect
)
