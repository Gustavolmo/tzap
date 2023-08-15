module github.com/tzapio/tzap/pkg/connectors/openaiconnector

go 1.20

replace github.com/tzapio/tzap => ../../../

require (
	github.com/sashabaranov/go-openai v1.14.1
	github.com/tzapio/tokenizer v0.0.4
	github.com/tzapio/tzap v0.9.3
)

require github.com/dlclark/regexp2 v1.10.0 // indirect
