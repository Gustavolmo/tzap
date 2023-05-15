package main

import (
	"github.com/tzapio/tzap/cli/cmd/cmdutil"
	"github.com/tzapio/tzap/pkg/config"
	"github.com/tzapio/tzap/pkg/embed"
	"github.com/tzapio/tzap/pkg/tzap"
	"github.com/tzapio/tzap/pkg/tzapconnect"
	tutil "github.com/tzapio/tzap/pkg/util"
)

func main() {
	openai_apikey, err := tzapconnect.LoadOPENAI_APIKEY()
	if err != nil {
		panic(err)
	}
	tzap.
		NewWithConnector(
			tzapconnect.WithConfig(openai_apikey, config.Configuration{MD5Rewrites: true})).
		WorkTzap(func(t *tzap.Tzap) {
			files, err := tutil.ListFilesInDir("./")
			if err != nil {
				panic(err)
			}
			files = cmdutil.GetNonExcludedFiles(files)
			embed.PrepareEmbeddingsFromFiles(t, files)
		})
}
