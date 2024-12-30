/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	delete "github.com/karthikbhandary2/user/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "This command is in charge of deleting users.",
	Long: "This command allows you to create users by providing a username",
	Run: func(cmd *cobra.Command, args []string) {
		username := viper.GetString("username")

		// Validate required flags
		if username == "" {
			fmt.Println("Error: --username flag is required.")
			return
		}

		// Print values for debugging
		fmt.Printf("Deleting user: Username=%s\n", username)

		// Call the function and handle errors
		err := delete.DeleteUser(username)
		if err != nil {
			fmt.Printf("Failed to delete user: %v\n", err)
			return
		}

		fmt.Println("User deleted successfully.")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Define flags
	// deleteCmd.PersistentFlags().StringP("username", "u", "", "Name used to delete the user in the server. Eg: EmpID")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// Bind flags to viper
	// viper.BindPFlag("username", deleteCmd.PersistentFlags().Lookup("username"))

}
