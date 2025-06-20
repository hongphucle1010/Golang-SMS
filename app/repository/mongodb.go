package database

import "fmt"

type mongodb struct{}

func (db *mongodb) Exec(q string) error {
    return fmt.Errorf("mysql: not implemented <%s>", q)
}