package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"

	"github.com/mahmoudk1000/relen/internal/database"
)

var (
	instance *database.Queries
	conn     *sql.DB
	once     sync.Once
	initErr  error
	schema   = `
BEGIN;

CREATE TABLE IF NOT EXISTS projects (
    id              SERIAL PRIMARY KEY,
    name            TEXT NOT NULL UNIQUE,
    link            TEXT,
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS applications (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id      INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    name            TEXT NOT NULL,
    description     TEXT,
    repo_url        TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE (project_id, name)
);

CREATE TABLE IF NOT EXISTS project_versions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id      INTEGER NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    version         TEXT NOT NULL,
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE (project_id, version)
);

CREATE TABLE IF NOT EXISTS application_versions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    application_id  UUID NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    version         TEXT NOT NULL,
    git_hash        TEXT,
    description     TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE (application_id, version)
);

CREATE TABLE IF NOT EXISTS project_version_apps (
    project_version_id     UUID NOT NULL REFERENCES project_versions(id) ON DELETE CASCADE,
    application_id         UUID NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    application_version_id UUID NOT NULL REFERENCES application_versions(id),

    PRIMARY KEY (project_version_id, application_id)
);

CREATE INDEX IF NOT EXISTS idx_application_versions_app_id ON application_versions(application_id);
CREATE INDEX IF NOT EXISTS idx_app_versions_app_id ON application_versions(name);
CREATE INDEX IF NOT EXISTS idx_applications_project_id ON applications(project_id);
CREATE INDEX IF NOT EXISTS idx_proj_versions_proj_id ON project_versions(name);

COMMIT;
		`
)

func Init(connectionString string) error {
	once.Do(func() {
		conn, initErr = sql.Open("postgres", connectionString)
		if initErr != nil {
			initErr = fmt.Errorf("failed to open database: %w", initErr)
			return
		}

		if err := conn.Ping(); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}

		instance = database.New(conn)
	})

	return initErr
}

func BuildSchema() error {
	var exists bool

	err := conn.QueryRow(`
		SELECT EXISTS (
		SELECT FROM information_schema.tables
		WHERE table_schema = 'public' AND table_name = 'projects'
		)
		`).Scan(&exists)
	if err != nil {
		return fmt.Errorf("failed to check if schema exists: %w", err)
	}

	if _, err := conn.Exec(schema); err != nil {
		return fmt.Errorf("failed to initialize database schema: %w", err)
	}

	return nil
}

func Get() *database.Queries {
	if instance == nil {
		panic("database not initialized: call db.Init() first")
	}
	return instance
}

func GetConn() *sql.DB {
	if conn == nil {
		panic("database not initialized: call db.Init() first")
	}
	return conn
}

func Close() error {
	if conn != nil {
		return conn.Close()
	}
	return nil
}
