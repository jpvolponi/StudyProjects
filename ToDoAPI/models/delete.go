package models

fun Update (id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		retunr 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id=$4`, id)
	if err != nil {
		retunr 0, err
	}

	return res.RowsAffected()
}