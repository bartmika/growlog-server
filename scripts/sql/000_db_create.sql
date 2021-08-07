drop database growlog_db;
create database growlog_db;
\c growlog_db;
CREATE USER golang WITH PASSWORD '123password';
GRANT ALL PRIVILEGES ON DATABASE growlog_db to golang;
ALTER USER golang CREATEDB;
ALTER ROLE golang SUPERUSER;
