-- +migrate Up
CREATE TABLE users (
  userId VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  PRIMARY KEY (userId)
);

CREATE TABLE spaces (
  spaceId VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (spaceId)
);

CREATE TABLE space_members (
  userId VARCHAR(255) NOT NULL,
  spaceId VARCHAR(255) NOT NULL,
  PRIMARY KEY (userId, spaceId),
  FOREIGN KEY (userId) REFERENCES users(userId),
  FOREIGN KEY (spaceId) REFERENCES spaces(spaceId)
);

CREATE TABLE user_invitations (
  userInvitationId VARCHAR(255) NOT NULL,
  senderId VARCHAR(255) NOT NULL,
  receiverId VARCHAR(255) NOT NULL,
  spaceId VARCHAR(255) NOT NULL,
  status VARCHAR(255) NOT NULL,
  PRIMARY KEY (userInvitationId),
  FOREIGN KEY (senderId) REFERENCES users(userId),
  FOREIGN KEY (receiverId) REFERENCES users(userId),
  FOREIGN KEY (spaceId) REFERENCES spaces(spaceId)
);

-- +migrate Down
DROP TABLE user_invitations;
DROP TABLE space_members;
DROP TABLE spaces;
DROP TABLE users;
