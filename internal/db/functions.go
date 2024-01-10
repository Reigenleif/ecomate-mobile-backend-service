package db

import "context"

func Get[T any](ctx context.Context, query string) ([]T, error) {
	var list []T
	rows, _ := GetDB().Query(ctx, query)
	defer rows.Close()

	for rows.Next() {
		var item T
		err := rows.Scan(&item)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}

	return list, nil
}

func GetFirst[T any](ctx context.Context, query string) (T, error) {
	var item T
	rows, _ := GetDB().Query(ctx, query)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&item)
		if err != nil {
			return item, err
		}
	}

	return item, nil
}

// Bentar
func Exec(ctx context.Context, query string) (bool, error) {
	ct, _ := GetDB().Exec(ctx, query)

	if (ct.RowsAffected() < 1) {
		return false, nil
	}


	return true, nil
}