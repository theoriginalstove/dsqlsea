CREATE TABLE records (
    id         INTEGER PRIMARY KEY,
    tenant_id  BIGINT NOT NULL,
    name       TEXT NOT NULL,
    age        INT NOT NULL,
    status     TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
