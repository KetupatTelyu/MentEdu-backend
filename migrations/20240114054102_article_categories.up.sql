BEGIN;

CREATE TABLE IF NOT EXISTS article_categories
(
    article_id  INT           NOT NULL,
    category_id INT           NOT NULL,
    created_by  VARCHAR(128)  NOT NULL,
    updated_by  VARCHAR(128)  NOT NULL,
    deleted_by  VARCHAR(128),
    created_at  TIMESTAMPTZ   NOT NULL,
    updated_at  TIMESTAMPTZ   NOT NULL,
    deleted_at  TIMESTAMPTZ,
    PRIMARY KEY (article_id, category_id),
    FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE SET NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE SET NULL
    );

COMMIT;
