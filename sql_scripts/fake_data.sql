DROP TABLE IF EXISTS address;
CREATE TABLE address (
  IP           VARCHAR(16) NOT NULL,
  lbname       VARCHAR(100) NOT NULL,
  PRIMARY KEY (IP)
);

INSERT INTO address
  (IP, lbname)
VALUES
  ("234.234.46.3", "mike"),
  ("225.235.44.53", "mike"),
  ("223.237.24.31", "abc"),
  ("219.222.42.123", "mike"),
  ("205.234.34.23", "aj");
