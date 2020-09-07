CREATE TABLE go_sample.video_infos (
    ID INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    video_id VARCHAR(30),
    video_name VARCHAR(50),
    video_description TEXT,
    thumbnail_url VARCHAR(50),
    view_count INT,
    comment_count INT,
    like_count INT,
    dislike_count INT,
    upload_date DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);