package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tzapio/tzap/pkg/config"
	"github.com/tzapio/tzap/pkg/tzap"
	"github.com/tzapio/tzap/pkg/tzapconnect"
	"github.com/tzapio/tzap/pkg/util/stdin"
)

var gitcommitCmd = &cobra.Command{
	Use:   "gitcommit",
	Short: "prompts chatgpt to generate a commit message and commits it to the current git repo",
	Long:  `tbd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitcommit called")

		diff := exec.Command("git", "diff",
			"--staged",
			"--patch-with-raw",
			"--unified=2",
			"--color=never",
			"--no-renames",
			"--ignore-space-change",
			"--ignore-all-space",
			"--ignore-blank-lines",
		)
		out, err := diff.CombinedOutput()

		if err != nil {
			fmt.Println("Could not get diff:", err)
			return
		}
		println(string(out))

		t := tzap.
			NewWithConnector(
				tzapconnect.WithConfig(config.Configuration{SupressLogs: true}),
			).
			SetHeader(`Write a git commit message maximum 30 words.
			
Template:
{brief git commit message}`).
			AddUserMessage(string(out))
		c, err := t.CountTokens(t.Message.Content)
		if err != nil {
			fmt.Println("Could not count tokens:", err)
			return
		}
		if c >= 3900 {
			println("\n\n\n\n")
			fmt.Printf(
				"WARNING: diff is too long. TRUNCATING TO 3900 of %d estimated tokens\n",
				c)
		}
		fmt.Printf(
			"Summarizing %d estimated tokens\n",
			c)
		if !stdin.ConfirmToContinue() {
			return
		}
		t.Message.Content, err = t.OffsetTokens(t.Message.Content, 0, 3900)
		if err != nil {
			fmt.Println("Could not offset tokens:", err)
			return
		}
		content := t.RequestChat().Data["content"].(string)
		println("\n", content)
		if !stdin.ConfirmToContinue() {
			return
		}
		cmd2 := exec.Command("git", "commit", "-m", `"`+content+`"`)
		if err := cmd2.Run(); err != nil {
			println("Could not git commit. content:", content, " err:", cmd2)
		}
	},
}

func init() {
	rootCmd.AddCommand(gitcommitCmd)
}
