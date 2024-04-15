CREATE TABLE IF NOT EXISTS users(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  name varchar(50) NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS recipes(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  title varchar(50) NOT NULL,
  description text,
  user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS ingredients(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  name varchar(50) NOT NULL,
  amount integer NOT NULL,
  unit varchar(20) NOT NULL,
  recipe_id bigint NOT NULL REFERENCES recipes(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS procedures(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  description text NOT NULL,
  order_number integer NOT NULL,
  recipe_id bigint NOT NULL REFERENCES recipes(id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS likes(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  recipe_id bigint NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS comments(
  id bigint AUTO_INCREMENT NOT NULL primary key,
  comment text NOT NULL,
  user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  recipe_id bigint NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
/* TEST DATA */
INSERT INTO users (name)
VALUES ('John Doe'),
  ('Hoge 太郎'),
  ('Fuga 次郎');
INSERT INTO recipes (title, description, user_id)
VALUES ('味噌ラーメン', '美味しい味噌ラーメンの作り方です', 1),
  ('バナナスムージー', '美味しいバナナスムージーの作り方です', 2),
  ('納豆ごはん', '美味しい納豆ごはんの作り方です', 3);
INSERT INTO ingredients (name, amount, unit, recipe_id)
VALUES ('味噌', '15', 'g', 1),
  ('中華麺(乾)', '100', 'g', 1),
  ('チャーシュー', '30', 'g', 1),
  ('バナナ', '80', 'g', 2),
  ('牛乳', '200', 'cc', 2),
  ('納豆', '1', 'パック', 3),
  ('ごはん', '200', 'g', 3);
INSERT INTO procedures (description, order_number, recipe_id)
VALUES ('中華麺を茹でます', 1, 1),
  ('味噌とお湯を混ぜます', 2, 1),
  ('どんぶりに盛り付けて完成', 3, 1),
  ('バナナを粉砕します', 1, 2),
  ('牛乳と混ぜて完成', 2, 2),
  ('納豆にご飯を乗せるだけ！', 1, 3);
INSERT INTO likes (user_id, recipe_id)
VALUES (2, 1),
  (3, 1),
  (1, 3);
INSERT INTO comments (comment, user_id, recipe_id)
VALUES ('美味しそう！', 2, 1),
  ('作り方がわかりにくい', 3, 1);