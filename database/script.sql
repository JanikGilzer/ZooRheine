CREATE DATABASE ZooDaba;

USE ZooDaba;

CREATE TABLE gruppe (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE mitarbeiter (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    gruppen_id INT,
    FOREIGN KEY (gruppen_id) REFERENCES gruppe(id)
);

CREATE TABLE zeit (
    id int PRIMARY KEY AUTO_INCREMENT,
    uhr_zeit VARCHAR(255) NOT NULL
);

CREATE TABLE ort(
    id int PRIMARY KEY AUTO_INCREMENT,
    stadt VARCHAR(255) NOT NULL,
    plz VARCHAR(255) NOT NULL
);

CREATE TABLE lieferant (
    id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    adresse VARCHAR(255) NOT NULL,
    ort_id int,
    FOREIGN KEY (ort_id) REFERENCES ort(id)
);

CREATE TABLE revier (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    beschreibung VARCHAR(255) NOT NULL
);

CREATE TABLE pfleger (
    id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    telefonnummer VARCHAR(255) NOT NULL,
    adresse VARCHAR(255) NOT NULL,
    ort_id int,
    revier_id int,
    FOREIGN KEY (ort_id) REFERENCES ort(id),
    FOREIGN Key (revier_id) REFERENCES revier(id)
);

CREATE TABLE gebaude (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    revier_id INT,
    FOREIGN KEY (revier_id) REFERENCES revier(id)
);


CREATE TABLE tierart (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);


CREATE TABLE tier (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    geburtstag date NOT NULL,
    gebaude_id INT,
    tierart_ID INT,
    FOREIGN KEY (gebaude_id) REFERENCES gebaude(id),
    FOREIGN KEY (tierart_ID) REFERENCES tierart(id)
);

CREATE TABLE fuetterungszeit (
    id int PRIMARY KEY AUTO_INCREMENT,
    zeit_id int,
    gebaude_id int,
    FOREIGN KEY (zeit_id) REFERENCES zeit(id),
    FOREIGN KEY (gebaude_id) REFERENCES gebaude(id)
);

CREATE TABLE futter (
    id int PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    lieferant_id int,
    FOREIGN KEY (lieferant_id) REFERENCES lieferant(id)
);

CREATE TABLE benoetigtesfutter (
    id int PRIMARY KEY AUTO_INCREMENT,
    tier_id int,
    futter_id int,
    FOREIGN KEY (tier_id) REFERENCES tier(id),
    FOREIGN KEY (futter_id) REFERENCES futter(id)
);


-- Insert example data into Gruppe
INSERT INTO gruppe (name) VALUES ('Admin'), ('User');

-- Insert example data into User
INSERT INTO mitarbeiter (name, password, gruppen_id) VALUES ('John Doe', 'password123', 1), ('Jane Smith', 'password456', 2);

-- Insert example data into Zeit
INSERT INTO zeit (uhr_zeit) VALUES ('08:00'), ('12:00'), ('16:00');

-- Insert example data into Ort
INSERT INTO ort (stadt, plz) VALUES ('Berlin', '10115'), ('Munich', '80331');

-- Insert example data into Lieferant
INSERT INTO lieferant (name, adresse, ort_id) VALUES ('Supplier A', 'Street 1', 1), ('Supplier B', 'Street 2', 2);

-- Insert example data into Revier
INSERT INTO revier (name, beschreibung) VALUES ('Revier 1', ' '), ('Revier 2', ' ' );

-- Insert example data into Pfleger
INSERT INTO pfleger (name, telefonnummer, adresse, ort_id, revier_id) VALUES ('Keeper A', '123456789', 'Street 3', 1, 1), ('Keeper B', '987654321', 'Street 4', 2, 2);

-- Insert example data into Gebaude
INSERT INTO gebaude (name, revier_id) VALUES ('Building 1', 1), ('Building 2', 2);

Insert Into tierart (id, name) VALUES (1, 'Löwe'), (2, 'Elephant');

-- Insert example data into Tier
INSERT INTO tier (name, geburtstag, gebaude_id, tierart_ID) VALUES ('Leo', '2015-05-01', 1, 1), ('Elefanten', '2010-08-15', 2, 2);

-- Insert example data into FuetterungsZeit
INSERT INTO fuetterungszeit (zeit_id, gebaude_id) VALUES (1, 1), (2, 2);

-- Insert example data into Futter
INSERT INTO futter (name, lieferant_id) VALUES ('Meat', 1), ('Vegetables', 2);

-- Insert example data into BenoetigtesFutter
INSERT INTO benoetigtesfutter (tier_id, futter_id) VALUES (1, 1), (2, 2);