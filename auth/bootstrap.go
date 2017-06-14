package auth

import (
	"database/sql"
	"fmt"
	"os"
)

// BootstrapAdminFromEnvironment does exactly what it says on the can
func BootstrapAdminFromEnvironment() error {
	bootstrapKey := os.Getenv("BOOTSTRAP_ADMIN_KEY")
	if bootstrapKey == "" {
		fmt.Println("No admin bootstrap key set")
		return nil
	}

	if _, err := GetAPIKeyByKey(bootstrapKey); err != sql.ErrNoRows {
		if err == nil {
			fmt.Println("Admin bootstrap key already defined")
			return nil
		}
		return err
	}

	fmt.Println("Bootstrapping API key from environment...")

	apiKey := APIKey{
		Key:   bootstrapKey,
		Admin: true,
	}
	return AddAPIKey(apiKey)
}
