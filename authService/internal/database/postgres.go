package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v4/log/logrusadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type DB struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, conf string) *DB {
	dbConf, err := pgxpool.ParseConfig(conf)
	if err != nil {
		log.Fatalf("can't parse postgres config: %v", err)
	}

	logrusLogger := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.JSONFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}

	dbConf.ConnConfig.Logger = logrusadapter.NewLogger(logrusLogger)

	pool, err := pgxpool.ConnectConfig(ctx, dbConf)
	if err != nil {
		log.Fatalf("error connecting to postgres: %v", err)
	}

	db := &DB{Pool: pool}

	db.Ping(ctx)

	return db
}

func (db *DB) Ping(ctx context.Context) {
	if err := db.Pool.Ping(ctx); err != nil {
		log.Fatalf("can't ping postgres: %v", err)
	}

	log.Println("postgres connected")
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) Migrate() error {
	_, b, _, _ := runtime.Caller(0)

	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(b))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error create the migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error migrate up: %v", err)
	}

	log.Println("migration done")

	return nil
}

func (db *DB) Drop() error {
	_, b, _, _ := runtime.Caller(0)

	migrationPath := fmt.Sprintf("file:///%s/migrations", path.Dir(b))

	m, err := migrate.New(migrationPath, db.conf.Database.URL)
	if err != nil {
		return fmt.Errorf("error create the migrate instance: %v", err)
	}

	if err := m.Drop(); err != nil {
		return fmt.Errorf("error drop: %v", err)
	}

	log.Println("migration drop")

	return nil
}

func (db *DB) Truncate(ctx context.Context) error {
	if _, err := db.Pool.Exec(ctx, `
		DELETE FROM users;
	`); err != nil {
		return fmt.Errorf("error truncate: %v", err)
	}

	return nil
}
