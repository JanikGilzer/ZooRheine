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
INSERT INTO mitarbeiter (name, password, gruppen_id) VALUES ('admin', '$2a$14$692xvxXH6onXCNoJnJInBO2esb/Ki9xDOFA6yMMRxdx21lv/zR/8C', 1), ('verwaltung', '$2a$12$gtW1QNcT0fPTPY26C0ScGOuwpPRtNikN4v4hcgzASo5us0XKbX0JS', 2), ('pfleger', '1234', 3);

-- Insert example data into Zeit
INSERT INTO zeit (uhr_zeit) VALUES
                                ('08:00'),
                                ('09:00'),
                                ('10:00'),
                                ('11:00'),
                                ('12:00'),
                                ('13:00'),
                                ('14:00'),
                                ('15:00'),
                                ('16:00'),
                                ('17:00');;

-- Insert example data into Ort
INSERT INTO ort (stadt, plz) VALUES   ('Berlin', '10115'),
                                      ('München', '80331'),
                                      ('Hamburg', '20095'),
                                      ('Köln', '50667'),
                                      ('Frankfurt', '60311'),
                                      ('Stuttgart', '70173'),
                                      ('Düsseldorf', '40210'),
                                      ('Leipzig', '04109'),
                                      ('Dresden', '01067'),
                                      ('Bremen', '28195');

-- Insert example data into Lieferant
INSERT INTO lieferant (name, adresse, ort_id) VALUES   ('Futterland GmbH', 'Futterstraße 1', 1),
                                                       ('Tierbedarf AG', 'Tierweg 5', 2),
                                                       ('ZooFutter Inc.', 'Zooallee 10', 3),
                                                       ('Futter Express', 'Expressweg 3', 4),
                                                       ('Nahrung Spezialisten', 'Nahrungstraße 7', 5),
                                                       ('Animal Supplies', 'Tierstraße 15', 6),
                                                       ('Exotische Futter', 'Exotenweg 2', 7),
                                                       ('BioFutter', 'Biostraße 11', 8),
                                                       ('NaturFutter GmbH', 'Naturallee 4', 9),
                                                       ('Premium Zoo Service', 'Premiumweg 9', 10);

-- Insert example data into Revier
INSERT INTO revier (id, name, beschreibung) VALUES
                                               (1,'Bergrevier', 'Gebirgiges Terrain mit seltenen Tieren'),
                                               (2,'Regenwaldrevier', 'Dichter, grüner Bereich mit vielfältiger Flora'),
                                               (3,'Wüstenrevier', 'Trocken und heiß, ideal für Wüstentiere'),
                                               (4,'Savannenrevier', 'Große Weide mit afrikanischen Tieren'),
                                               (5,'Polarrevier', 'Kalte Umgebung für Eis- und Schneetiere'),
                                               (6,'Dschungelrevier', 'Üppiger, tropischer Bereich mit exotischen Vögeln'),
                                               (7,'Reptilienrevier', 'Gebäude für Schlangen, Echsen und Krokodile'),
                                               (8,'Aquatisches Revier', 'Wasserwelt mit Fischen und Amphibien'),
                                               (9,'Schwartz-Wald', ''),
                                               (10,'Vogelrevier', 'Offene Voliere für verschiedene Vogelarten');



-- Insert example data into Pfleger
INSERT INTO pfleger (name, telefonnummer, adresse, ort_id, revier_id) VALUES
                                                                          ('Hans Müller', '030123456', 'Pflegerstraße 10', 1, 1),
                                                                          ('Anna Schmidt', '089123456', 'Zooweg 2', 2, 2),
                                                                          ('Peter Becker', '040123456', 'Tierpfad 3', 3, 3),
                                                                          ('Laura Fischer', '0221123456', 'Pflegeweg 5', 4, 4),
                                                                          ('Markus Weber', '069987654', 'Tierstraße 8', 5, 5),
                                                                          ('Sabrina Hoffmann', '071112233', 'Futterstraße 3', 6, 6),
                                                                          ('Jens Schneider', '021123456', 'Zooallee 7', 7, 7),
                                                                          ('Katrin Bauer', '034123456', 'Tiergehege 4', 8, 8),
                                                                          ('Tim Lehmann', '035112233', 'Wildpfad 2', 9, 9),
                                                                          ('Ute Meyer', '051123456', 'Vogelweg 6', 10, 10);

-- Insert example data into Gebaude
INSERT INTO gebaude (name, revier_id) VALUES
                                          ('Alpenhaus', 1),         -- Für das Bergrevier
                                          ('Regenwaldhaus', 2),      -- Für das Regenwaldrevier
                                          ('Wüstenoase', 3),         -- Für das Wüstenrevier
                                          ('Savannenhaus', 4),       -- Für das Savannenrevier
                                          ('Polargebäude', 5),       -- Für das Polarrevier
                                          ('Dschungelhaus', 6),      -- Für das Dschungelrevier
                                          ('Reptilienhaus', 7),      -- Für das Reptilienrevier
                                          ('Aquarium', 8),           -- Für das Aquatische Revier
                                          ('Schwarzwaldhaus', 9),    -- Für den Schwarz-Wald
                                          ('Voliere', 10);           -- Für das Vogelrevier

                                          

Insert Into tierart (id, name) VALUES (1, 'Löwe'), (2, 'Elephant'), (3, 'Ratte'), (4, 'Faultier'), (5, 'Wal');
Insert Into tierart (id, name) VALUES (6, 'Dino'), (7, 'Delphin'), (8, 'Dornteufel'), (9, 'Schnabel tier'), (10, 'Giraffe');
Insert Into tierart (id, name) VALUES (11, 'Mantis Shrimp'), (12, 'Kalmar'), (13, 'Nahwal'), (14, 'Quokka'), (15, 'Schildkröte');
Insert Into tierart (id, name) VALUES (16, 'Waschbär'), (17, 'Wolf'),   (18, 'Steinbock'),      -- Alpenhaus (Bergrevier)
                                      (19, 'Bergziege'),      -- Alpenhaus (Bergrevier)
                                      (20, 'Orang-Utan'),     -- Regenwaldhaus (Regenwaldrevier)
                                      (21, 'Tukan'),          -- Regenwaldhaus (Regenwaldrevier)
                                      (22, 'Kamel'),          -- Wüstenrevier (Wüstenoase)
                                      (23, 'Fennek'),         -- Wüstenrevier (Wüstenoase)
                                      (24, 'Nashorn'),        -- Savannenrevier (Savannenhaus)
                                      (25, 'Hyäne'),          -- Savannenrevier (Savannenhaus)
                                      (26, 'Eisbär'),         -- Polargebäude (Polarrevier)
                                      (27, 'Pinguin'),        -- Polargebäude (Polarrevier)
                                      (28, 'Tiger'),          -- Dschungelrevier (Dschungelhaus)
                                      (29, 'Affe'),           -- Dschungelrevier (Dschungelhaus)
                                      (30, 'Leguan'),         -- Reptilienrevier (Reptilienhaus)
                                      (31, 'Chamäleon'),      -- Reptilienrevier (Reptilienhaus)
                                      (32, 'Hai'),            -- Aquatisches Revier (Aquarium)
                                      (33, 'Seepferdchen'),   -- Aquatisches Revier (Aquarium)
                                      (34, 'Fuchs'),          -- Schwarzwald (Schwarzwaldhaus)
                                      (35, 'Reh'),            -- Schwarzwald (Schwarzwaldhaus)
                                      (36, 'Wildschwein'),    -- Schwarzwald (Schwarzwaldhaus)
                                      (37, 'Papagei'),        -- Vogelrevier (Voliere)
                                      (38, 'Adler'),          -- Vogelrevier (Voliere)
                                      (39, 'Eule');           -- Vogelrevier (Voliere)


-- Insert example data into Tier
INSERT INTO tier (name, geburtstag, gebaude_id, tierart_ID) VALUES
                                                                -- Alpenhaus (Bergrevier) – Tierarten: Steinbock (18) und Bergziege (19)
                                                                ('Billy', '2015-04-01', 1, 18),
                                                                ('Zara', '2016-06-10', 1, 19),

                                                                -- Regenwaldhaus (Regenwaldrevier) – Tierarten: Orang-Utan (20) und Tukan (21)
                                                                ('Koko', '2012-08-15', 2, 20),
                                                                ('Rio', '2014-09-09', 2, 21),

                                                                -- Wüstenoase (Wüstenrevier) – Tierarten: Kamel (22) und Fennek (23)
                                                                ('Cammy', '2010-03-21', 3, 22),
                                                                ('Fenny', '2011-07-30', 3, 23),

                                                                -- Savannenhaus (Savannenrevier) – Tierarten: Löwe (1), Elephant (2), Nashorn (24), Hyäne (25), Giraffe (10) und Quokka (14)
                                                                ('Leo', '2013-05-05', 4, 1),
                                                                ('Ella', '2010-01-12', 4, 2),
                                                                ('Rocky', '2012-11-20', 4, 24),
                                                                ('Hyani', '2015-08-18', 4, 25),

                                                                -- Polargebäude (Polarrevier) – Tierarten: Wal (5), Eisbär (26), Pinguin (27) und Narwal (13)
                                                                ('Wally', '2008-12-12', 5, 5),
                                                                ('Icy', '2011-09-09', 5, 26),
                                                                ('Pingu', '2013-10-10', 5, 27),
                                                                ('Nari', '2014-11-11', 5, 13),

                                                                -- Dschungelhaus (Dschungelrevier) – Tierarten: Faultier (4), Tiger (28) und Affe (29)
                                                                ('Slowy', '2016-06-06', 6, 4),
                                                                ('Tigo', '2015-05-05', 6, 28),
                                                                ('Momo', '2017-07-07', 6, 29),

                                                                -- Reptilienhaus (Reptilienrevier) – Tierarten: Dino (6), Dornteufel (8), Leguan (30) und Chamäleon (31)
                                                                ('Rex', '2005-03-03', 7, 6),
                                                                ('Spike', '2006-04-04', 7, 8),
                                                                ('Lenny', '2007-07-07', 7, 30),
                                                                ('Cammy_R', '2008-08-08', 7, 31),

                                                                -- Aquarium (Aquatisches Revier) – Tierarten: Delphin (7), Hai (32), Seepferdchen (33), Kalmar (12), Mantis Shrimp (11) und Schildkröte (15)
                                                                ('Flipper', '2010-01-01', 8, 7),
                                                                ('Jaws', '2009-02-02', 8, 32),
                                                                ('Seppi', '2011-03-03', 8, 33),
                                                                ('Kalle', '2012-04-04', 8, 12),


                                                                -- Schwarzwaldhaus (Schwarzwald) – Tierarten: Ratte (3), Wolf (17), Fuchs (34), Reh (35), Wildschwein (36) und Waschbär (16)
                                                                ('Ratti', '2018-07-07', 9, 3),
                                                                ('Luna', '2017-08-08', 9, 17),
                                                                ('Foxy', '2016-09-09', 9, 34),
                                                                ('Bambi', '2015-10-10', 9, 35),


                                                                -- Voliere (Vogelrevier) – Tierarten: Schnabel tier (9), Papagei (37), Adler (38) und Eule (39)
                                                                ('Beaky', '2012-01-01', 10, 9),
                                                                ('Polly', '2013-02-02', 10, 37),
                                                                ('Sky', '2014-03-03', 10, 38),
                                                                ('Hooty', '2015-04-04', 10, 39);

-- Insert example data into FuetterungsZeit
INSERT INTO fuetterungszeit (zeit_id, gebaude_id) VALUES   (1, 1),   -- 08:00 im Elefantenhaus
                                                           (2, 2),   -- 09:00 im Löwengehege
                                                           (3, 3),   -- 10:00 im Reptilienhaus
                                                           (4, 4),   -- 11:00 im Aquarium
                                                           (5, 5),   -- 12:00 in der Bärenhütte
                                                           (6, 6),   -- 13:00 in der Wüstenoase
                                                           (7, 7),   -- 14:00 im Pinguinhaus
                                                           (8, 8),   -- 15:00 im Bauernhofgebäude
                                                           (9, 9),   -- 16:00 in der Voliere Nord
                                                           (10, 10); -- 17:00 im Regenwaldhaus

-- Insert example data into Futter
INSERT INTO futter (name, lieferant_id) VALUES   ('Obstmischung', 1),
                                                 ('Fleischstücke', 2),
                                                 ('Fischfutter', 3),
                                                 ('Gemüse-Mix', 4),
                                                 ('Knochenbrühe', 5),
                                                 ('Insekten-Snack', 6),
                                                 ('Nagerfutter', 7),
                                                 ('Samenmischung', 8),
                                                 ('Blattgemüse', 9),
                                                 ('Spezialfutter', 10);


INSERT INTO tierart (id, name) VALUES
                                   (40, 'Gämse'),
                                   (41, 'Murmeltiere'),
                                   (42, 'Bär');

-- 2. Insert two new buildings in the Bergrevier (Alpenrevier).
--    Assuming revier_id = 1 corresponds to the alpine area.
INSERT INTO gebaude (name, revier_id) VALUES
                                          ('Panorama-Haus', 1),
                                          ('AlpenOase', 1);

-- 3. Insert four animals for each new building.
--    For "Alpenpanorama-Haus":
--      - "Alpensteinbock" (using existing tierart id 18: Steinbock)
--      - "Gämse" (newly added, id 40)
--      - "Murmeltiere" (newly added, id 41)
--      - "Adler" (using existing tierart id 38: Adler)
INSERT INTO tier (name, geburtstag, gebaude_id, tierart_ID) VALUES
                                                                ('Alpensteinbock_APH', '2018-04-01', (SELECT id FROM gebaude WHERE name = 'Panorama-Haus'), 18),
                                                                ('Gämse_APH', '2019-05-05', (SELECT id FROM gebaude WHERE name = 'Panorama-Haus'), 40),
                                                                ('Murmeltiere_APH', '2020-06-06', (SELECT id FROM gebaude WHERE name = 'Panorama-Haus'), 41),
                                                                ('Adler_APH', '2017-07-07', (SELECT id FROM gebaude WHERE name = 'Panorama-Haus'), 38);




-- Insert example data into BenoetigtesFutter
INSERT INTO benoetigtesfutter (tier_id, futter_id) VALUES    (1, 4),   -- Billy (Steinbock) erhält Gemüse-Mix
                                                             (2, 4),   -- Zara (Bergziege) erhält Gemüse-Mix
                                                             (3, 1),   -- Koko (Orang-Utan) erhält Obstmischung
                                                             (4, 8),   -- Rio (Tukan) erhält Samenmischung
                                                             (5, 1),   -- Cammy (Kamel) erhält Obstmischung
                                                             (6, 1),   -- Fenny (Fennek) erhält Obstmischung
                                                             (7, 2),   -- Leo (Löwe) erhält Fleischstücke
                                                             (8, 4),   -- Ella (Elephant) erhält Gemüse-Mix
                                                             (9, 4),   -- Rocky (Nashorn) erhält Gemüse-Mix
                                                             (10, 2),  -- Hyani (Hyäne) erhält Fleischstücke
                                                             (11, 4),  -- Gemma (Giraffe) erhält Gemüse-Mix
                                                             (12, 1),  -- Quokki (Quokka) erhält Obstmischung
                                                             (13, 3),  -- Wally (Wal) erhält Fischfutter
                                                             (14, 2),  -- Icy (Eisbär) erhält Fleischstücke
                                                             (15, 3),  -- Pingu (Pinguin) erhält Fischfutter
                                                             (16, 3),  -- Nari (Narwal) erhält Fischfutter
                                                             (17, 1),  -- Slowy (Faultier) erhält Obstmischung
                                                             (18, 2),  -- Tigo (Tiger) erhält Fleischstücke
                                                             (19, 1),  -- Momo (Affe) erhält Obstmischung
                                                             (20, 10), -- Rex (Dino) erhält Spezialfutter
                                                             (21, 6),  -- Spike (Dornteufel) erhält Insekten-Snack
                                                             (22, 4),  -- Lenny (Leguan) erhält Gemüse-Mix
                                                             (23, 6),  -- Cammy_R (Chamäleon) erhält Insekten-Snack
                                                             (24, 3),  -- Flipper (Delphin) erhält Fischfutter
                                                             (25, 3),  -- Jaws (Hai) erhält Fischfutter
                                                             (26, 3),  -- Seppi (Seepferdchen) erhält Fischfutter
                                                             (27, 3),  -- Kalle (Kalmar) erhält Fischfutter
                                                             (28, 3),  -- Mantis (Mantis Shrimp) erhält Fischfutter
                                                             (29, 4),  -- Shelly (Schildkröte) erhält Gemüse-Mix
                                                             (30, 7),  -- Ratti (Ratte) erhält Nagerfutter
                                                             (31, 2),  -- Luna (Wolf) erhält Fleischstücke
                                                             (32, 2),  -- Foxy (Fuchs) erhält Fleischstücke
                                                             (33, 4),  -- Bambi (Reh) erhält Gemüse-Mix
                                                             (34, 4); -- Oink (Wildschwein) erhält Gemüse-Mix



