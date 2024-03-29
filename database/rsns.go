package database

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"tectonic-sync/utils"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func UpdateRsns(nc []utils.NameChange, verbose bool) error {
	b := &pgx.Batch{}

	for _, user := range nc {
		query := psql.Update("rsn").Set("rsn", user.NewName).Where(squirrel.Eq{"wom_id": strconv.Itoa(user.PlayerId)})
		sql, args, err := query.ToSql()
		if err != nil {
			return err
		}
		b.Queue(sql, args...)
	}

	results := db.SendBatch(context.Background(), b)
	defer results.Close()

	errs := make([]error, 0, len(nc))
	updated := 0

	for _, user := range nc {
		commandTag, err := results.Exec()
		if err != nil {
			errs = append(errs, fmt.Errorf("Error updating user [%s -> %s](%d): %w", user.OldName, user.NewName, user.PlayerId, err))
			continue
		}

		if verbose {
			fmt.Printf("[%s -> %s](%d)\n", user.OldName, user.NewName, user.PlayerId)
		}

		// Check the commandTag to determine the success of the update
		if commandTag.RowsAffected() == 1 {
			fmt.Printf("User [%s -> %s](%d) successfully updated\n", user.OldName, user.NewName, user.PlayerId)
			updated++
		}
	}

	fmt.Printf("Updated %d/%d fetched users.\n", updated, len(nc))

	errs = append(errs, results.Close())

	return errors.Join(errs...)
}
