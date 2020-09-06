CREATE TABLE go_sample.channel_info (
    ID INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    channelID VARCHAR(30),
    channelName VARCHAR(50),
    viewCount INT,
    subscriberCount INT,
    videoCount INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);