CREATE TABLE records (
    id         BIGSERIAL PRIMARY KEY,
    tenant_id  BIGINT NOT NULL,
    name       TEXT NOT NULL,
    age        INT NOT NULL,
    status     TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE profiles (
    id         BIGSERIAL PRIMARY KEY,
    record_id  BIGINT NOT NULL REFERENCES records(id),
    bio        TEXT,
    avatar_url TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE tags (
    id         BIGSERIAL PRIMARY KEY,
    name       TEXT NOT NULL
);

CREATE TABLE record_tags (
    record_id  BIGINT NOT NULL REFERENCES records(id),
    tag_id     BIGINT NOT NULL REFERENCES tags(id),
    PRIMARY KEY (record_id, tag_id)
);
