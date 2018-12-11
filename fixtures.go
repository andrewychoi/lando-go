package main

func getFixtureQueries() []string {
	insertUsers := `
	INSERT INTO users (id, first_name, last_name, email) VALUES
	(1, 'Andrew', 'Choi', 'andrew@choi.com'),
	(2, 'Alice', 'Smith', 'alice@example.com'),
	(3, 'Bob', 'Smith', 'bob@example.com');`
	return []string{insertUsers}
}