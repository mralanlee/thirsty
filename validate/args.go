package validate

import "strings"

func checkEnv(environment string) bool {
	env := strings.ToLower(environment)
	return env != "qa" && env != "dev"
}
