package main

func getCreationQueries() []string {

	return []string{
		`CREATE TABLE IF NOT EXISTS users (
		  	id SERIAL PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS people (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users (id),
			frequency TEXT
		);`, 
		`CREATE TABLE IF NOT EXISTS meetings (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users (id),
			person_id INTEGER REFERENCES people (id),
			datetime TIMESTAMP WITH TIME ZONE NOT NULL,
			note TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			user_id INTEGER REFERENCES users (id),
			person_id INTEGER REFERENCES people (id),
			datetime TIMESTAMP WITH TIME ZONE NOT NULL
		);`,
	}

}

func getDeletionQueries() []string {
	return []string{
		`DROP TABLE IF EXISTS notes;`,
		`DROP TABLE IF EXISTS meetings;`,
		`DROP TABLE IF EXISTS people;`,
		`DROP TABLE IF EXISTS users;`,
	}
}
