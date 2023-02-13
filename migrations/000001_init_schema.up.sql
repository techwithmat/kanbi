CREATE TABLE "user" (
  id SERIAL PRIMARY KEY,
  email VARCHAR(30) UNIQUE NOT NULL,
  username VARCHAR(50) NOT NULL,
  password TEXT NOT NULL
);

CREATE TABLE board (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  user_id INT NOT NULL REFERENCES "user"(id)
);

CREATE TABLE board_column (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  board_id INT NOT NULL REFERENCES board(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES "user"(id)
);

CREATE TABLE task (
  id SERIAL PRIMARY KEY,
  title VARCHAR(50) NOT NULL,
  description TEXT,
  board_column_id INT NOT NULL REFERENCES board_column(id) ON DELETE CASCADE,
  user_id INT NOT NULL REFERENCES "user"(id)
);