package validate

import "strings"

// CheckEnv If proper environment is provided in flag
func CheckEnv(environment string) bool {
	env := strings.ToLower(environment)
	return env == "qa" || env == "dev" || env == "prd"
}

// func CheckFlags(flag string) bool {

// }
