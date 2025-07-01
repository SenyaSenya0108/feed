CREATE TABLE category
(
    id        UUID    DEFAULT gen_random_uuid(),
    site_id   INTEGER      NOT NULL,
    slug      VARCHAR(150) NOT NULL,
    name      VARCHAR(150) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    parent_id UUID         NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES category (id)
)