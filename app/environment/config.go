package environment

import (
	"bytes"
	"os"
	"runtime"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

var HOSTNAME, _ = os.Hostname()
var VERSION = "Development"
var CURRENT_DIR, _ = os.Getwd()

func SetUpEnv() {
	SetUpLogging()

}

func GetEnvString(key string, EnvDefault string) string {

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	}
	return EnvDefault
}

func SetEnvString(key string, value string) {
	_ = os.Setenv(key, value)
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
