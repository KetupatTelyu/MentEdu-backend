BEGIN;

CREATE TABLE IF NOT EXISTS consultation_orders
(
    id           UUID          NOT NULL,
    user_id      UUID          NOT NULL,
    purpose      VARCHAR(255)  NOT NULL,
    date_time    TIMESTAMPTZ   NOT NULL,
    status       VARCHAR(255)  NOT NULL,
    consultant_id UUID,
    created_by   VARCHAR(128)  NULL,
    updated_by   VARCHAR(128)  NULL,
    deleted_by   VARCHAR(128)  NULL,
    created_at   TIMESTAMPTZ   NOT NULL,
    updated_at   TIMESTAMPTZ   NOT NULL,
    deleted_at   TIMESTAMPTZ,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    FOREIGN KEY (consultant_id) REFERENCES users (id) ON DELETE SET NULL
    );

COMMIT;
