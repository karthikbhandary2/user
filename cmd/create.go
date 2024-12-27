package cmd

import (
	"fmt"
	create "github.com/karthikbhandary2/user/pkg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "This command is in charge of creating users.",
	Long:  "This command allows you to create users by providing a username and a full name.",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")
		name := viper.GetString("name")

		// Validate required flags
		if username == "" || name == "" {
			fmt.Println("Error: Both --username and --name flags are required.")
			return
		}

		// Print values for debugging
		fmt.Printf("Creating user: Username=%s, Name=%s\n", username, name)

		// Call the function and handle errors
		err := create.CreateUser(username, name)
		if err != nil {
			fmt.Printf("Failed to create user: %v\n", err)
			return
		}

		fmt.Println("User created successfully.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Define flags
	createCmd.PersistentFlags().StringP("username", "u", "", "Name used to create the user in the server. Eg: EmpID")
	createCmd.PersistentFlags().StringP("name", "n", "", "Actual name of the user. Eg: Dio Brando")
	
	// Bind flags to viper
	viper.BindPFlag("username", createCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("name", createCmd.PersistentFlags().Lookup("name"))
}
