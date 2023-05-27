/*
Copyright © 2023 github.com/qzakwani
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	// "os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cpz",
	Short: "Copy js css files from dir1 to dir2.",
	Long: `
	Copy js and css files from dir1 to dir2 ->
	Delete existing css js files named index from dir2 ->
	Paste the files in dir2 and renaming to be index.js and index.css.


	Intended to update changes from bundled files in a js project like svelte to 
	be components in another.
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		from := viper.GetString("from")
		to := viper.GetString("to")
		if from == "" || !strings.HasSuffix(from, "/") {
			fmt.Println()
			fmt.Println("ERROR: <from> path was not found! or not ending with /")
			fmt.Println()
			return
		}
		if to == "" || !strings.HasSuffix(to, "/") {
			fmt.Println()
			fmt.Println("ERROR: <to> path was not found! or not ending with /")
			fmt.Println()
			return
		}

		if _, err := os.Stat(from); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "\nError: from -> ", err)
			return
		}

		if _, err := os.Stat(to); os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "\nError: to -> ", err)
			return
		}

		var output string
		if temp := viper.GetString("output"); temp == "" {
			output = "index"
		} else {
			output = temp
		}

		files, err := os.ReadDir(from)
		if err != nil {
			fmt.Fprintln(os.Stderr, "\nError: ", err)
			return

		}

		for _, file := range files {
			name := file.Name()
			if ex := filepath.Ext(name); ex == ".css" || ex == ".js" {
				fmt.Fprintln(os.Stderr, "Moving: ", name)
				err := move(from, to, name, output+ex)
				if err != nil {
					fmt.Fprintln(os.Stderr, "\nError: ", err)
					return
				}
			}
		}

		fmt.Fprintln(os.Stderr, "\nDone.")

	},
}

func move(from, to, name, output string) error {
	c := exec.Command("cp", "-f", from+name, to+output)
	err := c.Run()

	return err

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Search config in home directory with name ".cpz" (without extension).
	viper.AddConfigPath(".")
	viper.SetConfigType("json")
	viper.SetConfigName(".cpz")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		fmt.Println()

	} else {
		fmt.Println()
		fmt.Println("ERROR: Config file was not found. <.cpz.json> ")
		fmt.Println()
		os.Exit(1)
	}
}
