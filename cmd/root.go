package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "quick-color",
	Short: "It colors.",
	Long:  `Really, that's about it, it colors.`,
	Run: func(cmd *cobra.Command, args []string) {
		output := termenv.NewOutput(os.Stdout)

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			var inputLines []string

			line := scanner.Text()

			inputLines = append(inputLines, line)

			allInput := ""
			if len(inputLines) > 0 {
				allInput = inputLines[0]
				for i := 1; i < len(inputLines); i++ {
					allInput += "\n" + inputLines[i]
				}
			}

			s := output.String(allInput)

			if len(args) < 1 {
				fmt.Println(s)
				os.Exit(1)
			}

			arg_color := args[0]

			colored := s.Foreground(output.Color(arg_color))

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

			fmt.Println(colored)

			os.Stdout.Sync()
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
