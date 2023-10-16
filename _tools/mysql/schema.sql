use playground;

CREATE TABLE IF NOT EXISTS city
(
    id          INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '都市ID',
    country     VARCHAR(3) NOT NULL COMMENT '国',
    name        VARCHAR(255) NOT NULL COMMENT '都市名',
    created_at  DATETIME  default current_timestamp,
    updated_at  DATETIME default current_timestamp on update current_timestamp
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='都市';

CREATE TABLE IF NOT EXISTS city_detail
(
    id         INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '都市詳細ID',
    city_id    INT UNSIGNED NOT NULL COMMENT '都市ID',
    lat        FLOAT NOT NULL COMMENT '緯度',
    lon        FLOAT NOT NULL COMMENT '経度',
    population INT NOT NULL COMMENT '人口',
    timezone   INT NOT NULL COMMENT 'UTCとの時差(秒)',
    created_at DATETIME  default current_timestamp,
    updated_at DATETIME default current_timestamp on update current_timestamp,
    UNIQUE KEY unique_city_population (city_id, population),
    FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE CASCADE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='都市詳細';

CREATE TABLE IF NOT EXISTS city_temperature
(
    id          BIGINT  AUTO_INCREMENT PRIMARY KEY,
    city_id     INT UNSIGNED NOT NULL COMMENT '都市ID',
    date        DATETIME     NOT NULL COMMENT '日時',
    temperature FLOAT        NOT NULL COMMENT '摂氏',
    pressure    INT          NOT NULL COMMENT '気圧',
    humidity    INT          NOT NULL COMMENT '湿度',
    created_at  DATETIME  default current_timestamp,
    updated_at  DATETIME default current_timestamp on update current_timestamp,
    UNIQUE KEY unique_city_date (city_id, date),
    FOREIGN KEY (city_id) REFERENCES city(id) ON DELETE CASCADE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='都市の気温';
