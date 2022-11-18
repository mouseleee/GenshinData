package db

type RoleIndex struct {
	Id        int64
	Nickname  string
	DetailUrl string
}

func InsertRoleIndex(role RoleIndex) (int64, error) {
	q := `insert into role_index(nickname, url) values (?, ?)`
	r, err := pool.Exec(q, role.Nickname, role.DetailUrl)
	if err != nil {
		return 0, err
	}
	i, _ := r.RowsAffected()
	return i, nil
}
