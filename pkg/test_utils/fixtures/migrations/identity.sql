
CREATE TABLE IF NOT EXISTS identity
(
    id         UUID PRIMARY KEY                                   NOT NULL,
    username   text                                               NOT NULL,
    pass_hash   text                                               NOT NULL,
    eth_address bytea,
    country    text,
    state      text                                               NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_identity_username ON identity (username);

CREATE UNIQUE INDEX IF NOT EXISTS idx_identity_eth_address ON identity (eth_address) WHERE state = 'SUBMITTED';