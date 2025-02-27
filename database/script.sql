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
INSERT INTO gruppe (name) VALUES ('Admin'), ('Verwaltung'), ('Pfleger');

-- Insert example data into User
INSERT INTO mitarbeiter (name, password, gruppen_id) VALUES ('admin', '$2a$14$692xvxXH6onXCNoJnJInBO2esb/Ki9xDOFA6yMMRxdx21lv/zR/8C', 1), ('verwaltung', '$2a$14$bAYK588LBSy0CVNJkA0Kn.5.JF/.EEQsjKRv.HQgz7S7YYAk8Kywu', 2), ('pfleger', '1234', 3);

-- Insert example data into Zeit
INSERT INTO zeit (uhr_zeit) VALUES
                                ('06:00'), ('11:00'), ('13:00'), ('15:00'), ('18:00'), ('20:00');

-- Insert example data into Ort
INSERT INTO ort (stadt, plz) VALUES ('Berlin', '10115'), ('Munich', '80331'), ('Dresden', '01067'),
                                    ('Leipzig', '04109'),
                                    ('Düsseldorf', '40213'),
                                    ('Hannover', '30159'),
                                    ('Nürnberg', '90403');

-- Insert example data into Lieferant
INSERT INTO lieferant (name, adresse, ort_id) VALUES ('Fleischlieferung Müller', 'Fleischstraße 22', 3),
                                                    ('Gemüsebau Schmidt', 'Gemüseweg 45', 4),
                                                    ('Fischhandel Fischer', 'Fischkai 12', 5),
                                                    ('Exotisches Futterland', 'Exotenweg 8', 6);

-- Insert example data into Revier
INSERT INTO revier (id, name, beschreibung) VALUES
                                               (1,'Höhle', ' '),
                                               (2,'Tiefsee', ' ' ),
                                               (3,'Fallout', ''),
                                               (4,'Savanne', ''),
                                               (5,'Frankfurter-HBF', ''),
                                               (6,'Jungle', ''),
                                               (7,'Erde', ''),
                                               (8,'Strand', ''),
                                               (9,'Schwartz-Wald', ''),
                                               (10,'Kleintier', '');

UPDATE revier SET beschreibung =
                      CASE
                          WHEN name = 'Höhle' THEN 'Afrikanische Savannentiere wie Löwen, Elefanten und Giraffen'
                          WHEN name = 'Tiefsee' THEN 'Tiefseebewohner wie Riesenkalmare und Anglerfische'
                          WHEN name = 'Fallout' THEN 'Schlangen, Echsen und Schildkröten aus aller Welt'
                          WHEN name = 'Savanne' THEN 'Kleine Säugetiere wie Waschbären und Quokkas'
                          WHEN name = 'Frankfurter-HBF' THEN 'Tiere aus dem Schwarzwald wie Wölfe und Wildschweine'
                          WHEN name = 'Jungle' THEN 'Tiere aus Nord- und Südamerika wie Bisons und Faultiere'
                          WHEN name = 'Erde' THEN 'Tiere aus Nord- und Südamerika wie Bisons und Faultiere'
                          WHEN name = 'Strand' THEN 'Tiere aus Nord- und Südamerika wie Bisons und Faultiere'
                          WHEN name = 'Schwartz-Wald' THEN 'Tiere aus Nord- und Südamerika wie Bisons und Faultiere'
                          WHEN name = 'Kleintier' THEN 'Tiere aus Nord- und Südamerika wie Bisons und Faultiere'
                          ELSE beschreibung
                          END;


-- Insert example data into Pfleger
INSERT INTO pfleger (name, telefonnummer, adresse, ort_id, revier_id) VALUES
                                                                          ('Lisa Großtier', '0151123457', 'Savannenweg 1', 3, 1),
                                                                          ('Max Reptil', '0170987655', 'Schlangenpfad 2', 4, 3),
                                                                          ('Sophie Meer', '0160123457', 'Fischkai 8', 5, 2),
                                                                          ('Tom Kleintier', '0151123458', 'Waldweg 9', 6, 4),
                                                                          ('Julia Wald', '0170987656', 'Schwarzwaldstraße 10', 7, 5);

-- Insert example data into Gebaude
INSERT INTO gebaude (name, revier_id) VALUES
                                          ('Neandertaler', 1),
                                          ('Kraken', 2),
                                          ('Rat Rat Rat', 3),
                                          ('Löwen', 4),
                                          ('Elefanten', 4),
                                          ('Giraffen', 4),
                                          ('Ratten', 5),
                                          ('Faultiere', 6),
                                          ('Waschbären', 6),
                                          ('Delphine', 2),
                                          ('Dornteufel', 4),
                                          ('Wal' , 8),
                                          ('Schnabeltiere', 9),
                                          ('Mantis Shrimp', 2),
                                          ('Kalmar', 2),
                                          ('Nahwal', 8),
                                          ('Quokka', 6),
                                          ('Schildkröten', 6),
                                          ('Wölfe', 7),
                                          ('Bären', 9),
                                          ('Hund', 9);

                                          

Insert Into tierart (id, name) VALUES (1, 'Löwe'), (2, 'Elephant'), (3, 'Ratte'), (4, 'Faultier'), (5, 'Wal');
Insert Into tierart (id, name) VALUES (6, 'Dino'), (7, 'Delphin'), (8, 'Dornteufel'), (9, 'Schnabel tier'), (10, 'Giraffe');
Insert Into tierart (id, name) VALUES (11, 'Mantis Shrimp'), (12, 'Kalmar'), (13, 'Nahwal'), (14, 'Quokka'), (15, 'Schildkröte');
Insert Into tierart (id, name) VALUES (16, 'Waschbär'), (17, 'Wolf');


-- Insert example data into Tier
INSERT INTO tier (name, geburtstag, gebaude_id, tierart_ID) VALUES
                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2),

                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2),

                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2),
                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2),

                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2),

                                                                ('Leo', '2015-05-01', 1, 1),
                                                                ('Elefanten', '2010-08-15', 2, 2);

-- Insert example data into FuetterungsZeit
INSERT INTO fuetterungszeit (zeit_id, gebaude_id) VALUES (1, 1), (2, 2);

-- Insert example data into Futter
INSERT INTO futter (name, lieferant_id) VALUES ('Meat', 1), ('Vegetables', 2);

-- Insert example data into BenoetigtesFutter
INSERT INTO benoetigtesfutter (tier_id, futter_id) VALUES (1, 1), (2, 2);