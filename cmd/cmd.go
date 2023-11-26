package cmd

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/router"
	"log"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "bookkeeping",
	}
	srvCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}
	dbCmd := &cobra.Command{
		Use: "db",
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CtreateUserTable()
		},
	}

	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	rootCmd.AddCommand(srvCmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(migrateCmd)
	database.GormConnect()
	defer database.Close()
	rootCmd.Execute()
}

func RunServer() {
	r := router.New()
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("r.Run 的下一行")
}
