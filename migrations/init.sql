CREATE DATABASE IF NOT EXISTS redikru;
USE redikru;

CREATE TABLE IF NOT EXISTS companies (
  id CHAR(36) PRIMARY KEY,
  name TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS jobs (
  id CHAR(36) PRIMARY KEY,
  company_id CHAR(36),
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL
);

CREATE INDEX idx_jobs_created_at ON jobs(created_at);
-- fulltext could be considered, but MySQL fulltext on TEXT requires engine InnoDB + MySQL 5.6+ and other config.
