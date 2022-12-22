CREATE TABLE `stocks` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL,
  `symbol` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `name` varchar(256) COLLATE utf8mb4_bin,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `stocks_UNIQUE` (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `intervals` (
  `id` char(36) COLLATE utf8mb4_bin NOT NULL,
  `time` varchar(256) COLLATE utf8mb4_bin NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `intervals_UNIQUE` (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `price` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `stock_id` varchar(36) COLLATE utf8mb4_bin NOT NULL,
  `interval_id` varchar(36) COLLATE utf8mb4_bin NOT NULL,
  `open` float(10,3) unsigned NOT NULL COMMENT '始値',
  `high` float(10,3) unsigned NOT NULL COMMENT '高値',
  `low` float(10,3) unsigned NOT NULL COMMENT '安値',
  `close` float(10,3) unsigned NOT NULL COMMENT '終値',
  `volume` int unsigned NOT NULL COMMENT '終値',
  `acquisition_time` datetime NOT NULL COMMENT '株価取得日時',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE KEY `price_UNIQUE` (`stock_id`,`interval_id`,`acquisition_time`),
  CONSTRAINT `stocks_id` FOREIGN KEY (`stock_id`) REFERENCES `stocks` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `intervals_id` FOREIGN KEY (`interval_id`) REFERENCES `intervals` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;