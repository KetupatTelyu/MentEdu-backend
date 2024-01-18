BEGIN;

CREATE TABLE IF NOT EXISTS consultation_details
(
    id                UUID          NOT NULL,
    consultation_id  UUID          NOT NULL,
    meeting_url       TEXT  NOT NULL,
    created_by        VARCHAR(128)  NULL,
    updated_by        VARCHAR(128)  NULL,
    deleted_by        VARCHAR(128)  NULL,
    created_at        TIMESTAMPTZ   NOT NULL,
    updated_at        TIMESTAMPTZ   NOT NULL,
    deleted_at        TIMESTAMPTZ,
    PRIMARY KEY (id),
    FOREIGN KEY (consultation_id) REFERENCES consultation_orders (id) ON DELETE SET NULL
    );

COMMIT;
