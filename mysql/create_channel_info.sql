CREATE TABLE go_sample.channel_infos (
    ID INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    channel_id VARCHAR(30),
    channel_name VARCHAR(50),
    view_count INT,
    subscriber_count INT,
    video_count INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);