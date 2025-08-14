package cmd

import (
	"bytes"
	"fmt"
	"grep/internal"
	"io"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var outputFile string
var after, before int = 0, 0

var rootCmd = &cobra.Command{
	Use:   "grep",
	Short: "grep is a service which searchs for specific text within files.",
	Long:  "grep is a service which searchs for specific text within files.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOut(os.Stdout) // ensure Cobra help/usage prints to STDOUT
		cmd.SetErr(os.Stderr) // ensure errors print to STDERR

		cis, _ := cmd.Flags().GetBool("case_insensitive")
		lineCount, _ := cmd.Flags().GetBool("count")

		searchString := args[0]
		var res []string
		var output string

		if len(args) == 1 {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read from the console.")
				os.Exit(1)
			}

			buffer := bytes.NewReader(data)
			res = internal.SearchFile(searchString, buffer, cis, after, before)
		} else {
			fi, err := os.Stat(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if fi.IsDir() {
				fl, err := os.ReadDir(args[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
				for _, f := range fl {
					fp := path.Join(args[1], f.Name())

					file, err := internal.OpenFile(fp)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}
					sr := internal.SearchFile(searchString, file, cis, after, before)

					for i := range sr {
						sr[i] = fp + ":" + sr[i]
					}
					res = append(res, sr...)
				}
			} else {
				file, err := internal.OpenFile(args[1])
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}

				res = internal.SearchFile(searchString, file, cis, after, before)
			}
		}

		if lineCount {
			output += fmt.Sprintf("%d\n", len(res))
		} else {
			for _, r := range res {
				output += r + "\n"
			}
		}

		if outputFile != "" {
			internal.WriteToFile(output, outputFile)
		} else {
			fmt.Fprint(os.Stdout, output)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("case_insensitive", "i", false, "Search without case-sensitivity")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Write output to a file")
	rootCmd.Flags().IntVarP(&after ,"nlines_after", "A", 0, "Print n lines after the match")
	rootCmd.Flags().IntVarP(&before ,"nlines_before", "B", 0, "Print n lines before the match")
	rootCmd.Flags().BoolP("count", "C", false, "Only print count of matches instead of actual matched lines")
}
