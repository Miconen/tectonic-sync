package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"tectonic-sync/database"
	"tectonic-sync/utils"
)

func main() {
	groupFlag := flag.String("group", "", "group")
	dbFlag := flag.String("db", "", "db")
	guildFlag := flag.String("guild", "", "guild")
	verboseFlag := flag.Bool("verbose", false, "verbose")

	flag.Parse()

	if *groupFlag != "" {
		os.Setenv("GROUP_ID", *groupFlag)
	}
	if *dbFlag != "" {
		os.Setenv("DATABASE_URL", *dbFlag)
	}
	if *guildFlag != "" {
		os.Setenv("GUILD_ID", *guildFlag)
	}

	guild := os.Getenv("GUILD_ID")
	if guild == "" {
		fmt.Fprintf(os.Stderr, "Error getting environment variable: \"GUILD_ID\"\n")
		os.Exit(1)
	}

	group := os.Getenv("GROUP_ID")
	if group == "" {
		fmt.Fprintf(os.Stderr, "Error getting environment variable: \"GROUP_ID\"\n")
		os.Exit(1)
	}

	conn, err := database.InitDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	users, err := utils.GetNameChanges(group)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching name changes: %v\n", err)
		os.Exit(1)
	}

	err = database.UpdateRsns(guild, users, *verboseFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating users: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
