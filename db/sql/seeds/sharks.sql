BEGIN TRANSACTION;

TRUNCATE TABLE "sharks";

INSERT INTO "sharks"
VALUES
(1, 'Basking Shark', 'Cetorhinus maximus', 'Description of basking shark', 'Image of basking shark'),
(2, 'Zebra Bullhead Shark', 'Heterodontus zebra', 'Description of zebra shark', 'Image of zebra shark');

COMMIT;