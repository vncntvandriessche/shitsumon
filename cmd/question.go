package cmd

import (
	"fmt"
	"os"

	"github.com/dixonwille/wlog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// questionCmd represents the question command
var questionCmd = &cobra.Command{
	Use:   "question",
	Short: "Start the questioning.",
	Long: `Ask shitsumon to start testing your skills.

By default it will ask 10 questions and give you your results.`,
	Run: func(cmd *cobra.Command, args []string) {
		var ui wlog.UI
		var vocabulary map[string]string

		title := viper.GetString("title")
		amount := viper.GetInt("questions")

		ui = wlog.New(os.Stdin, os.Stdout, os.Stdout)
		ui = wlog.AddConcurrent(ui)
		ui = wlog.AddColor(wlog.Blue, wlog.Red, wlog.Cyan, wlog.None, wlog.None, wlog.None, wlog.None, wlog.Green, wlog.Yellow, ui)
		ui = wlog.AddPrefix("", "[x]", "", "", "", "[-]", "[v]", "[!]", ui)

		ui.Info(title)
		err := viper.UnmarshalKey("vocabulary", &vocabulary)

		if err != nil {
			ui.Error("Unable to decode your \"vocabulary item in config file.\"")
		}

		count := 0
		correct := 0
		for question, valid_answer := range vocabulary {
			count++

			answer, err := ui.Ask(fmt.Sprintf("%v) \"%s\"", count, question))

			if err != nil {
				ui.Error("Something went wrong posing this question.")
			}

			if answer == valid_answer {
				ui.Success("Correct answer")
				correct++
			} else {
				ui.Error(fmt.Sprintf("Wrong answer (correct: \"%s\")", valid_answer))
			}

			if count >= amount {
				break
			}
		}

		wrong_answers := count - correct

		ui.Info("\nSummary...")
		ui.Output(fmt.Sprintf("Vocabulary count: %v", len(vocabulary)))
		ui.Output(fmt.Sprintf("Questions asked: %v", count))
		ui.Success(fmt.Sprintf("Correct answers: %v", correct))
		ui.Error(fmt.Sprintf("Wrong answers: %v", wrong_answers))
	},
}

func init() {
	RootCmd.AddCommand(questionCmd)

	viper.SetDefault("questions", "10")
	viper.SetEnvPrefix("shitsumon")
}
