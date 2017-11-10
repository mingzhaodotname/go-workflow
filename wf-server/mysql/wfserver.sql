CREATE DATABASE IF NOT EXISTS wfserver;

USE wfserver;

CREATE TABLE IF NOT EXISTS wf_template (
    id        VARCHAR(36)         NOT NULL,
    name      VARCHAR(64)         NOT NULL,
    version    VARCHAR(64),
    description   VARCHAR(1024),
    PRIMARY KEY   (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE UNIQUE INDEX wf_template_idx ON wf_template (name, version);

