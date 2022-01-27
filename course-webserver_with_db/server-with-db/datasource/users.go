package datasource

func (ds *pgDB) GetUsers() ([]string, error) {
	rows, err := ds.db.Query(`select name from users;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []string{}
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		users = append(users, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
