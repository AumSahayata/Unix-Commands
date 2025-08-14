package cmd

import (
	"fmt"
	"os"
	"gotree/internal"

	"github.com/spf13/cobra"
)

var levels int

var rootCmd = &cobra.Command{
	Use:"gotree",
	Short: "This command lists down all the files and directories in the given directory.",
	Long: "This command lists down all the files and directories in the given directory.",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string){
		fullPath, _ := cmd.Flags().GetBool("fullPath")
		printDir, _ := cmd.Flags().GetBool("printDir")
		printPerm, _ := cmd.Flags().GetBool("printPerm")
		sort, _ := cmd.Flags().GetBool("sort")

		output := internal.ScanDir(args[0], levels, sort, printPerm, fullPath, printDir)
		fmt.Print(output)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("fullPath", "f", false, "Print full paths of the directories and files.")
	rootCmd.Flags().BoolP("printDir", "d", false, "Print only directories.")
	rootCmd.Flags().BoolP("printPerm", "p", false, "Prints permissions of the files and directories.")
	rootCmd.Flags().BoolP("sort", "t", false, "Sort by modification time.")
	
	rootCmd.Flags().IntVarP(&levels, "levels", "L", -1, "Traverse to specified level.")
}