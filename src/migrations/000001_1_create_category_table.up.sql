CREATE TABLE category
(
    id        UUID PRIMARY KEY,
    site_id   INTEGER      NOT NULL,
    slug      VARCHAR(150) NOT NULL,
    name      VARCHAR(150) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE
)