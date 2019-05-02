BEGIN TRANSACTION;

TRUNCATE TABLE "todos";

INSERT INTO "todos"
VALUES
(1, 'Test Title one', 'Test Body one'),
(2, 'Test Title two', 'Test Body two');

COMMIT;