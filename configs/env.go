package configs

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnvConfigs() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading configs: %v", err)
		return
	}

	fmt.Println("Configs Loaded")
}
