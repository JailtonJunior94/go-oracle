alter session set container=FREEPDB1;
---
create user customer identified by customer;
---
grant all privileges to customer;