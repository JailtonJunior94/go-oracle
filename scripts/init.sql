-- Create schema
ALTER SESSION SET CONTAINER=FREEPDB1;
CREATE USER JJ IDENTIFIED BY jj QUOTA UNLIMITED ON USERS;
GRANT CONNECT, RESOURCE TO JJ;

-- Create users table within the schema
CREATE TABLE JJ.users (
    id NUMBER PRIMARY KEY,
    name VARCHAR2(100),
    email VARCHAR2(100),
    birth_date DATE,
    active CHAR(1)
);