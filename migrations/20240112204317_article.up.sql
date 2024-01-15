BEGIN;

CREATE TABLE IF NOT EXISTS articles
(
    id          SERIAL         ,
    title       VARCHAR(128) NOT NULL,
    body     TEXT         NOT NULL,
    image       TEXT,
    slug       VARCHAR(128) NOT NULL,
    created_by  VARCHAR(128) NOT NULL,
    updated_by  VARCHAR(128) NOT NULL,
    deleted_by  VARCHAR(128),
    created_at  TIMESTAMPTZ  NOT NULL,
    updated_at  TIMESTAMPTZ  NOT NULL,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (id)
);

COMMIT;