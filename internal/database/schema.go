package database

var schema = `
CREATE TABLE IF NOT EXISTS users(
	id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
	username TEXT,
	email TEXT UNIQUE,
	password TEXT,
	isAdmin BOOLEAN
);
`
