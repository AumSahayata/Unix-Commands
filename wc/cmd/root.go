package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"wc/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "wc",
	Short: "wc is a service which counts the lines, words and characters in a file.",
	Long: "wc is a service which counts the lines, words and characters in a file.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.SetOut(os.Stdout) // ensure Cobra help/usage prints to STDOUT
		cmd.SetErr(os.Stderr) // ensure errors print to STDERR
		
		lines, _ := cmd.Flags().GetBool("lines")
		words, _ := cmd.Flags().GetBool("words")
		char, _ := cmd.Flags().GetBool("char")

		if len(args) == 0 {
			data, err := io.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to read from the console.")
				os.Exit(1)
			}
			buffer := bytes.NewReader(data)

			_, outStr, err := wcFunction(buffer, lines, words, char)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to count from the input.")
				os.Exit(1)
			} else {
				fmt.Fprint(os.Stdout, outStr, "\n")
			}
		}

		var total [3]int
		var hadError bool
		
		for _, path := range args {
			f, err := internal.OpenFile(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				hadError = true
				continue
			}
			defer f.Close()

			temp, outStr, err := wcFunction(f, lines, words, char)
			if err != nil {
				hadError = true
			} else {
				total[0] += temp[0]
				total[1] += temp[1]
				total[2] += temp[2]
				fmt.Fprint(os.Stdout, outStr, " ", path, "\n")
			}
		}

		if len(args) > 1 {
			fmt.Fprintf(os.Stdout, "%8d %8d %8d total", total[0], total[1], total[2])
		}

		if hadError {
			os.Exit(1)
		}
	},
}

func wcFunction(file io.Reader, lines, words, char bool) ([3]int, string, error) {
	var output string
	var count [3]int
	var err error

	if lines {
		count[0], err = internal.CountLines(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		output += fmt.Sprintf("%8d", count[0])
	}

	if words {
		count[1], err = internal.CountWords(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		output += fmt.Sprintf("%8d", count[1])
	}

	if char {
		count[2], err = internal.CountCharacters(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		output += fmt.Sprintf("%8d", count[2])
	}

	if !lines && !words && !char {
		count[0], err = internal.CountLines(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		count[1], err = internal.CountWords(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		count[2], err = internal.CountCharacters(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return [3]int{}, "", err
		}
		output = fmt.Sprintf("%8d %8d %8d", count[0], count[1], count[2])
	}

	return count, output, nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("lines", "l", false, "Count the number of lines present in the file.")
	rootCmd.Flags().BoolP("words", "w", false, "Count the number of words present in the file.")
	rootCmd.Flags().BoolP("char", "c", false, "Count the number of characters present in the file.")
}