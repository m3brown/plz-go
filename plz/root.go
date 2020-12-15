/*
Copyright © 2020 Mike Brown <brown.3.mike@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plz

import (
	"fmt"
    "path"
	"github.com/spf13/cobra"
    "log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plz-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func getCommand(id string, commands []Command) (Command, error) {
    for _, cmd := range commands {
        if cmd.Id == id {
            return cmd, nil
        }
    }
    return Command{}, fmt.Errorf("Could not find command '%s' in config", id)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    //cfg, err := plzConfig()
    filePath, err := configFilePath()

    if err != nil {
        fmt.Println(err)
        return
    }

    cfg, yamlErr := parseYaml(path.Join(filePath, configFilename))

    if yamlErr != nil {
        fmt.Println("ERROR!")
        fmt.Println(yamlErr)
        return
    }
    fmt.Println("CONFIG!")
    fmt.Println(cfg.Commands[0].Cmd)
    fmt.Println(cfg.Commands[0].Id)
    fmt.Println(cfg.Commands[0].Cmds)

    fmt.Println("HERE")
    cmd, rest := getArgs()
    fmt.Println(cmd)
    fmt.Println(rest)

    command, cmdErr := getCommand(cmd, cfg.Commands)
    if cmdErr != nil {
        fmt.Println(cmdErr)
        return
    }

    execErr := execCommand(command, rest, filePath)
    if execErr != nil {
        fmt.Println(execErr)
    }

    //execCommand("echo hello; sleep 3; echo foo; exit 1; echo bar")
	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.plz-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
        log.Print(fmt.Sprintf("got configFile: %s", cfgFile))
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".plz-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".plz-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
