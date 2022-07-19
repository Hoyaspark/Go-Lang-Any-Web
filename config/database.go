package config

import (
	"context"
	"database/sql"
	"errors"
)

var ErrNotFoundDatabaseFromContext = errors.New("not found database from context")

func NewDatabase() *sql.DB {
	con := DBProperties["mysql"]

	url := con.Username + ":" + con.Password + "@tcp(" + con.Address + ":" + con.Port + ")/" + con.Name + "?parseTime=true"

	db, err := sql.Open(con.Kind, url)

	checkErrForPanic(err)

	err = db.Ping()

	checkErrForPanic(err)

	// 항상 연결되어있는 커넥션 풀 갯수 설정
	db.SetMaxIdleConns(5)
	// 최대로 연결할 수 있는 커넥션 갯수 설정
	db.SetMaxOpenConns(10)

	return db
}

func ContextWithDatabase(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, DatabaseKey, db)
}

func DatabaseFromContext(ctx context.Context) (*sql.DB, error) {
	if db, ok := ctx.Value(DatabaseKey).(*sql.DB); ok {
		return db, nil
	} else {
		return nil, ErrNotFoundDatabaseFromContext
	}
}
