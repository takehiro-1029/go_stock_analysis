CREATE TABLE `stock_price` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL,
  `symbol_id` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `open` float(10,3) unsigned NOT NULL COMMENT '始値',
  `high` float(10,3) unsigned NOT NULL COMMENT '高値',
  `low` float(10,3) unsigned NOT NULL COMMENT '安値',
  `close` float(10,3) unsigned NOT NULL COMMENT '終値',
  `volume` int unsigned NOT NULL COMMENT '終値',
  `acquisition_time` datetime NOT NULL COMMENT '株価取得日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stocks_UNIQUE` (`symbol_id`,`acquisition_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;