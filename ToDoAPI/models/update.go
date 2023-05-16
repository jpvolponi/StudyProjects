package models

fun Update (id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		retunr 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title=$1, description=$2, done=$3 WHERE id=$4`, todo.Title, todo.Description, todo.Done, id)
}