package dotenv

import (
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

// MessageOut indicates message output.
// Default is `os.Stdout`.
var MessageOut io.Writer = os.Stdout

// Load apply .env to ENVIRONMENT VARIABLE.
func Load() {
	// just print messages to stdout before logger setup.
	if _, found := os.LookupEnv("ENV_FILE"); !found {
		fmt.Fprintln(MessageOut, "no env file specified. try to load default .env.")
		os.Setenv("ENV_FILE", ".env")
	}

	envfile := os.Getenv("ENV_FILE")
	err := godotenv.Load(envfile)
	if err != nil {
		fmt.Fprintf(MessageOut, "no env file loaded %v\n", err)
	} else {
		fmt.Fprintf(MessageOut, "env file loaded: %v\n", envfile)
	}
}
