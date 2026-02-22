package cli

import (
	"fmt"
)

// CliService provides CLI-specific business logic
type CliService struct{}

func (s *CliService) PrintInfo() {
	fmt.Println("Multi-Platform CLI")
	fmt.Println("==================")
	fmt.Println("Available commands:")
	fmt.Println("  info   - Show application information")
	fmt.Println("  help   - Show this help message")
}

func (s *CliService) GetVersion() string {
	return "1.0.0"
}
