CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  completed BOOL NOT NULL DEFAULT FALSE,
  collection TEXT NOT NULL,
  UNIQUE (title, collection)
);
