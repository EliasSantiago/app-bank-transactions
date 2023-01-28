package transactionrepository

import (
	"context"
)

func (repository repository) Check(id string) (bool, error) {
	ctx := context.Background()
	var exist bool
	err := repository.db.QueryRow(
		ctx,
		`SELECT EXISTS(SELECT id FROM wallets WHERE id = $1) AS exist`,
		id,
	).Scan(
		&exist,
	)
	if err != nil {
		return false, err
	}
	return exist, nil
}
