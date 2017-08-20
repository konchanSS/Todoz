-- +migrate Up

CREATE TABLE todos (
	id bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	body VARCHAR(256) NOT NULL COMMENT '内容',
	is_finished tinyint(1) DEFAULT 0 NOT NULL COMMENT 'ステータス',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
  deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE todos;