BEGIN;

CREATE TABLE IF NOT EXISTS projects (
    id SERIAL PRIMARY KEY,
    name            TEXT NOT NULL UNIQUE,
    link            TEXT NULL,
    description     TEXT NULL,
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

CREATE INDEX IF NOT EXISTS idx_app_versions_app_id ON application_versions(application_id);
CREATE INDEX IF NOT EXISTS idx_proj_versions_proj_id ON project_versions(project_id);

COMMIT;
