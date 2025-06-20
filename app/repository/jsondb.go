package database

type jsondb struct{}

func (db *jsondb) Exec(q string) error { return nil }