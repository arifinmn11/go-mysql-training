CREATE DATABASE rakamin;
USE rakamin;

CREATE TABLE Customers (
  CustomerId  INT NOT NULL AUTO_INCREMENT,
  FirstName   VARCHAR(255),
  LastName    VARCHAR(255),
  Phone       VARCHAR(255),
  City        VARCHAR(255),
  Birthday    DATE,
  PRIMARY KEY (CustomerId)
);

CREATE TABLE Orders (
  OrderId     VARCHAR(255),
  OrderDate   DATE,
  CustomerId  INT,
  PRIMARY KEY (OrderId),
  FOREIGN KEY (CustomerId) REFERENCES Customers(CustomerId) ON DELETE CASCADE
);

CREATE TABLE Items (
  ItemId        INT,
  ProductName   VARCHAR(255),
  Price         INT,
  PRIMARY KEY (ItemId)
);

CREATE TABLE OrderItems (
  OrderId     VARCHAR(255),
  ItemId      INT,
  Quantity    INT,
  CONSTRAINT PkOrderItems PRIMARY KEY (OrderId, ItemId)
);

INSERT INTO Customers (FirstName, LastName, Phone, City, Birthday) VALUES ("Husnul", "Hamidiah", "219-200-1600", "Jakarta", "1990-08-11"), ("James", "Bond", "219-201-7960", "Malang", "1992-01-01"), ("Ethan", "Hunt", "219-203-2642", "Surabaya", "1989-12-09"), ("Harry", "Potter", "219-205-0172", "London", "1989-11-04");