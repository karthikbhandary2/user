package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"

  homedir "github.com/mitchellh/go-homedir"
  "github.com/spf13/viper"

)


var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "user",
  Short: "Create Users.",
  Long: `Using this you can create users in your Linux servers.`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Here you will define your flags and configuration settings.
  // Cobra supports persistent flags, which, if defined here,
  // will be global for your application.

  rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.user.yaml)")


  // Cobra also supports local flags, which will only run
  // when this action is called directly.
  // rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
  
  //create command flags
  rootCmd.PersistentFlags().StringP("username", "u","", "name used to create the user in the server. Eg: EmpID")
	rootCmd.PersistentFlags().StringP("name", "n","", "actual name of the user. Eg: Dio Brando")
  rootCmd.PersistentFlags().StringP("version", "v", "", "version of the application")

  //bind flags to viper
  viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
  viper.BindPFlag("version", rootCmd.PersistentFlags().Lookup("version"))

  rootCmd.AddCommand(createCmd)
  rootCmd.AddCommand(deleteCmd)
  rootCmd.AddCommand(versionCmd)
}


// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".usercreation" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".user")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil {
    fmt.Println("Using config file:", viper.ConfigFileUsed())
  }
}

