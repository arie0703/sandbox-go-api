package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "microservices-cobra-cli",
	Short: "A microservices CLI application with 7 task commands",
	Long:  "A CLI application built with Cobra that provides 7 numbered task commands (task1-task7) for microservices deployment.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Initialize logging configuration
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	},
}

// executeTask executes a task with proper error handling and logging
func executeTask(taskName string) error {
	// Log task execution start
	log.Printf("Starting execution of %s", taskName)
	
	// Simulate task execution - in a real scenario, this would contain actual task logic
	fmt.Printf("this is %s!\n", taskName)
	
	// Log successful completion
	log.Printf("Successfully completed %s", taskName)
	return nil
}

var task1Cmd = &cobra.Command{
	Use:   "task1",
	Short: "Execute task1",
	Long:  "Execute task1 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task1"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task1: %v\n", err)
			return err
		}
		return nil
	},
}

var task2Cmd = &cobra.Command{
	Use:   "task2",
	Short: "Execute task2",
	Long:  "Execute task2 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task2"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task2: %v\n", err)
			return err
		}
		return nil
	},
}

var task3Cmd = &cobra.Command{
	Use:   "task3",
	Short: "Execute task3",
	Long:  "Execute task3 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task3"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task3: %v\n", err)
			return err
		}
		return nil
	},
}

var task4Cmd = &cobra.Command{
	Use:   "task4",
	Short: "Execute task4",
	Long:  "Execute task4 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task4"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task4: %v\n", err)
			return err
		}
		return nil
	},
}

var task5Cmd = &cobra.Command{
	Use:   "task5",
	Short: "Execute task5",
	Long:  "Execute task5 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task5"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task5: %v\n", err)
			return err
		}
		return nil
	},
}

var task6Cmd = &cobra.Command{
	Use:   "task6",
	Short: "Execute task6",
	Long:  "Execute task6 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task6"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task6: %v\n", err)
			return err
		}
		return nil
	},
}

var task7Cmd = &cobra.Command{
	Use:   "task7",
	Short: "Execute task7",
	Long:  "Execute task7 and output a log message",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := executeTask("task7"); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing task7: %v\n", err)
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(task1Cmd)
	rootCmd.AddCommand(task2Cmd)
	rootCmd.AddCommand(task3Cmd)
	rootCmd.AddCommand(task4Cmd)
	rootCmd.AddCommand(task5Cmd)
	rootCmd.AddCommand(task6Cmd)
	rootCmd.AddCommand(task7Cmd)
}

func Execute() {
	// Set up error handling for invalid commands
	rootCmd.SilenceUsage = false
	rootCmd.SilenceErrors = false
	
	if err := rootCmd.Execute(); err != nil {
		// Log error to stderr with proper formatting
		log.SetOutput(os.Stderr)
		log.Printf("Command execution failed: %v", err)
		os.Exit(1)
	}
}