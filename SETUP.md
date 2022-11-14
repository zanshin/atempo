# Install

## Atempo

    go install github.com/zanshin/atempo@latest

## Create Database
Create a MySQL database, add a user account, and grant that account all privileges.

    mysql> CREATE DATABASE atempo;
    mysql> CREATE USER atempo IDENTIFIED BY 'password';
    mysql> GRANT ALL PRIVILEGES ON *.* TO 'atempo'@'localhost';

## Create Database Tables
Run these SQL queries to create the database tables.

### visit table

    DROP TABLE IF EXISTS visit;
    CREATE TABLE visit (id int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, dt timestamp, ipv4 int UNSIGNED, referrer varchar(255), browser_family varchar(255), browser_version varchar(15), platform varchar(255), platform_version varchar(255), resolution varchar(13), country_name varchar(45), region_name varchar(45), city_name varchar(50)) ENGINE=InnoDB;

### href_click table

    DROP TABLE IF EXISTS href_click;
    CREATE TABLE href_click (id int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, dt timestamp, url varchar(255), ip_address varchar(15), href varchar(255), href_rectangle polygon) ENGINE=InnoDB;

### page_view table

    DROP TABLE IF EXISTS page_view;
    CREATE TABLE page_view (id int UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, dt timestamp, url varchar(255), ip_address varchar(15), user_agent varchar(255), screen_height int, screen_width int) ENGINE=InnoDB;

## Configuration
Using the template `config.json` file, fill in the database information.

## Running

    atempo -p config.json -l
