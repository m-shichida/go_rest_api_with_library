CREATE TABLE IF NOT EXISTS places (
    id int AUTO_INCREMENT,
    name varchar(100) NOT NULL,
    description varchar(300) NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);
