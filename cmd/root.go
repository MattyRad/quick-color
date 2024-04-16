package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func colorize(str string, foreground string) termenv.Style {
	output := termenv.NewOutput(os.Stdout)
	colored := output.String(str).Foreground(output.Color(foreground))

	bold := viper.GetBool("bold")
	italic := viper.GetBool("italic")
	faint := viper.GetBool("faint")
	strike := viper.GetBool("strike")
	underline := viper.GetBool("underline")

	if bold {
		colored = colored.Bold()
	}

	if italic {
		colored = colored.Italic()
	}

	if faint {
		colored = colored.Faint()
	}

	if strike {
		colored = colored.CrossOut()
	}

	if underline {
		colored = colored.Underline()
	}

	return colored
}

var rootCmd = &cobra.Command{
	Use:   "quick-color",
	Short: "It colors.",
	Long:  `Really, that's about it, it colors.`,
	Run: func(cmd *cobra.Command, args []string) {
		foreground := "#fff"

		if len(args) >= 1 {
			foreground = args[0]
		}

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(colorize(scanner.Text(), foreground))
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("bold", "b", false, "Bold text")
	rootCmd.PersistentFlags().BoolP("italic", "i", false, "Italic text")
	rootCmd.PersistentFlags().BoolP("faint", "f", false, "Faint text")
	rootCmd.PersistentFlags().BoolP("strike", "s", false, "Strikethrough text")
	rootCmd.PersistentFlags().BoolP("underline", "u", false, "Underline text")
	viper.BindPFlag("bold", rootCmd.PersistentFlags().Lookup("bold"))
	viper.BindPFlag("italic", rootCmd.PersistentFlags().Lookup("italic"))
	viper.BindPFlag("faint", rootCmd.PersistentFlags().Lookup("faint"))
	viper.BindPFlag("strike", rootCmd.PersistentFlags().Lookup("strike"))
	viper.BindPFlag("underline", rootCmd.PersistentFlags().Lookup("underline"))
}
