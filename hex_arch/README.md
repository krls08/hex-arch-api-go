# MYSQL CONFIG

## Get mysql docker instance

docker pull mysql 
- or - 
podman pull docker.io/library/mysql

podman run --name mysql -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 docker.io/library/mysql:latest

podman exec -it mysql /bin/bash

check websocked connection is stablished
/var/run/mysqld/mysqld.sock exists

login in the db

mysql -u root -p
-> 123456

mysql> show databases;

Create user:
mysql> CREATE USER 'db_user' IDENTIFIED BY 'password';
Also can be used
mysql> CREATE USER 'db_user'@'localhost' IDENTIFIED BY 'password';
mysql> CREATE USER 'db_user'@'%' IDENTIFIED BY 'password';


mysql> SELECT User, Host FROM mysql.user;
+------------------+-----------+
| User             | Host      |
+------------------+-----------+
| db_user          | %         |
| root             | %         |
| mysql.infoschema | localhost |
| mysql.session    | localhost |
| mysql.sys        | localhost |
| root             | localhost |
+------------------+-----------+

When Host == % users will be able to connect from any IP

Edit user password:
mysql> ALTER USER 'db_user'@'%' IDENTIFIED BY 'new_password';

Create database:
mysql> CREATE DATABASE hex_arch_db;

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| hex_arch_db        |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+

Use the created database to create a table
mysql> use hex_arch_db;

Create table:
mysql> CREATE TABLE courses
(
    id       VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

mysql> show tables;
+-----------------------+
| Tables_in_hex_arch_db |
+-----------------------+
| courses               |
+-----------------------+
1 row in set (0.00 sec)

grant permisions to user
mysql> GRANT ALL PRIVILEGES ON *.* TO 'db_user'@'%';
grant permisions to certain table:
mysql> GRANT permission_type ON database.table TO 'username'@'localhost';

mysql> SHOW GRANTS FOR 'db_user';
mysql> SHOW GRANTS FOR 'db_user'@'%'; // also...


logout
mysql> \q

login as new user
$ mysql -u db_user -p
> password

mysql> use hex_arch_db;
mysql> select * from courses;

to view if there is any record in courses database

to create records from commandline
mysql> INSERT INTO table_name (column1, column2, column3, ...)
VALUES (value1, value2, value3, ...);
mysql> insert into courses (id, name, duration) values ('4ca073a2-e01f-11ec-9d64-0242ac120006', 'k', '5h'
);
Query OK, 1 row affected (0.18 sec)



## Running test from commandline example
go test -v internal/platform/server/handler/courses/*




