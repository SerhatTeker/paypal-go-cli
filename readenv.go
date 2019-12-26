package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// loads variables from .env into the system
func loadEnv() {
	envFile := ".env.prod"
	err := godotenv.Load(envFile)

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// reads env variables from os
func readEnv() {
	loadEnv()

	// Get env variables
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	// TODO: Change tests
	fmt.Println("FOO: ", s3Bucket)
	fmt.Println("FOO: ", secretKey)
}
