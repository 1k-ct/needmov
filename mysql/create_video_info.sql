CREATE TABLE go_sample.video_info (
    ID INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    videoID VARCHAR(30),
    videoName VARCHAR(50),
    videoDescription TEXT,
    thumbnailURL VARCHAR(50),
    viewCount INT,
    commentCount INT,
    likeCount INT,
    dislikeCount INT,
    uploadDate DATETIME,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);