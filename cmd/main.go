package main

import (
	"database/sql"
	"log"

	migration "github.com/jailtonjunior94/go-oracle/pkg/database/migrate"

	_ "github.com/sijms/go-ora/v2"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "go-oracle",
		Short: "Go Oracle",
	}

	migrate := &cobra.Command{
		Use:   "migrate",
		Short: "Go Oracle Migrations",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := sql.Open("oracle", "oracle://orcl:SuperPassword@2025@localhost:1521/FREEPDB1")
			if err != nil {
				log.Fatalf("error connecting to database: %v", err)
			}

			if err = conn.Ping(); err != nil {
				log.Fatalf("error pinging database: %v", err)
			}

			migrate, err := migration.NewMigrateOracle(conn, "file://../database/migrations", "FREEPDB1")
			if err != nil {
				log.Fatalf("error creating migration: %v", err)
			}

			if err = migrate.Execute(); err != nil {
				log.Fatalf("error executing migration: %v", err)
			}
		},
	}

	root.AddCommand(migrate)
	if err := root.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err)
	}
}
