package migrations

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

func migratedCmd(resolver func(driver string) *sqlx.DB) *cobra.Command {
	var cmd = new(cobra.Command)
	cmd.Use = "migrated"
	cmd.Short = "show migrated files list"
	cmd.Run = func(cmd *cobra.Command, args []string) {
		var err error
		driver, err := cmd.Flags().GetString("driver")
		if err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		}

		db := resolver(driver)
		if db == nil {
			fmt.Printf("failed: %s database driver not found\n", driver)
			return
		}

		res, err := db.Query("SELECT name FROM migrations WHERE is_seed = FALSE;")
		if err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		}

		for res.Next() {
			var table string
			res.Scan(&table)
			fmt.Println(table)
		}
	}
	return cmd
}
