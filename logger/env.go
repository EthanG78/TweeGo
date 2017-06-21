package logger

import "os"

func GetEnv (name string) string{
	env := os.Getenv(name)
	if env == ""{
		panic("Missing environment variable " + name)
	}
	return env
}