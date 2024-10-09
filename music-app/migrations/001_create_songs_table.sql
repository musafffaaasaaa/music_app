CREATE TABLE songs (
  id SERIAL PRIMARY KEY,
  group_name TEXT NOT NULL,
  song_title TEXT NOT NULL,
  release_date DATE,
  lyrics TEXT,
  external_link TEXT
);
 