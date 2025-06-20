package database

import "sms/app/environment"

type DB interface {
    Exec(q string) error
}

func NewDB(env environment.Env) DB {
    switch env {
    case environment.PROD:
        return &mongodb{}
    case environment.DEV:
        return &jsondb{}
    default:
        panic("unknown environment")
    }
}