package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"os"
)

func getSecureRandInt64() (int64, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(math.MaxInt64)))
	if err != nil {
		return 0, err
	}
	return val.Int64(), nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func printCustomBanner() {

	fmt.Println("")
	fmt.Println("██████╗  ██████╗ ██████╗ ")
	fmt.Println("██╔══██╗██╔════╝██╔════╝ ")
	fmt.Println("██████╔╝██║     ██║  ███╗")
	fmt.Println("██╔══██╗██║     ██║   ██║")
	fmt.Println("██║  ██║╚██████╗╚██████╔╝")
	fmt.Println("╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ")
	fmt.Println("rolling-code-generator")
	fmt.Println("⇨ version 0.1")
	fmt.Println("⇨ service uuid " + uuidServiceKey)
}
