package env

import "github.com/joho/godotenv"

func MustLoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
