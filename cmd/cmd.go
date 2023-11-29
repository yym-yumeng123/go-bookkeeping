package cmd

import (
	"bookkeeping/internal/database"
	"bookkeeping/internal/email"
	"bookkeeping/internal/router"
	"log"

	"github.com/spf13/cobra"
)

func Run() {
	// 根命令
	rootCmd := &cobra.Command{
		Use: "bookkeeping",
	}
	// server 命令
	serverCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}
	// 数据库命令
	dbCmd := &cobra.Command{
		Use: "db",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	createCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateTable()
		},
	}

	emailCmd := &cobra.Command{
		Use: "email",
		Run: func(cmd *cobra.Command, args []string) {
			email.Send()
		},
	}

	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	DropColumnCmd := &cobra.Command{
		Use: "drop",
		Run: func(cmd *cobra.Command, args []string) {
			database.DropColumn()
		},
	}

	crudCmd := &cobra.Command{
		Use: "crud",
		Run: func(cmd *cobra.Command, args []string) {
			database.Crud()
		},
	}

	database.GormConnect()
	defer database.Close()

	rootCmd.AddCommand(serverCmd, dbCmd)
	rootCmd.AddCommand(createCmd, migrateCmd, DropColumnCmd, crudCmd)
	rootCmd.AddCommand(emailCmd)
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
