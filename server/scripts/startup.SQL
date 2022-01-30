DROP TABLE IF EXISTS car;
CREATE TABLE car (
  id            INT NOT NULL,
  brand         VARCHAR(64) NOT NULL,
  model         VARCHAR(64) NOT NULL,
  horse_power   INT NOT NULL,
  PRIMARY KEY (`id`)
);