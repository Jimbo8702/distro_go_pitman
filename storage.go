package main

// Database represents the central database to store fetched data.
type Database struct {
	// conn *pgx.Conn
}

// NewDatabase creates a new Database instance.
func NewDatabase() (*Database, error) {
	// Connect to PostgreSQL
	// conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
	// if err != nil {
	// 	return nil, err
	// }

	return &Database{
		// conn: conn,
	}, nil
}

// SaveData saves fetched data to the database.
func (db *Database) SaveData(data string) error {
	// _, err := db.conn.Exec(context.Background(), "INSERT INTO crawled_data (data) VALUES ($1)", data)
	return nil
}