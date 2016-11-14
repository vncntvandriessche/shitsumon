package cmd

import (
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
		ui = wlog.New(os.Stdin, os.Stdout, os.Stdout)
		ui = wlog.AddConcurrent(ui)
		ui = wlog.AddColor(wlog.Blue, wlog.Red, wlog.Cyan, wlog.None, wlog.None, wlog.None, wlog.None, wlog.Green, wlog.Yellow, ui)
		ui = wlog.AddPrefix("-", "[x]", "", "", "", "[-]", "[v]", "[!]", ui)

		ui.Info(viper.GetString("title"))

		var v map[string]string
		err := viper.UnmarshalKey("vocabulary", &v)

		if err != nil {
			ui.Error("Unable to decode your \"vocabulary item in config file.\"")
			ui.Output("")
			ui.Output("Printing trace...")
			panic(err)
		}

		for k, v := range v {
			answer, err := ui.Ask("\"" + k + "\"")

			if err != nil {
				ui.Error("Something went pretty wrong posing this question.")
				ui.Output("")
				ui.Output("Printing trace...")
				panic(err)
			}

			if answer == v {
				ui.Success("Correct answer")
			} else {
				ui.Error("Wrong answer (correct: \"" + v + "\")")
			}
		}

		// ui.Error("Error message")
		// ui.Info("Info message")
		// ui.Success("Success message")
		// ui.Warn("Warning message")

		// ui.Output("|" + q + "|")
	},
}

func init() {
	RootCmd.AddCommand(questionCmd)

	// Here you will define your flags and configuration settings.
	viper.SetDefault("questions", "10")
	viper.SetEnvPrefix("shitsumon")
	viper.ReadInConfig()

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
