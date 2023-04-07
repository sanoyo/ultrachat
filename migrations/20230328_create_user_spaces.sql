-- +migrate Up
CREATE TABLE users (
  userId BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (userId)
);

CREATE TABLE spaces (
  spaceId BIGINT NOT NULL,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY (spaceId)
);

CREATE TABLE space_members (
  userId BIGINT NOT NULL,
  spaceId BIGINT NOT NULL,
  PRIMARY KEY (userId, spaceId),
  FOREIGN KEY (userId) REFERENCES users(userId),
  FOREIGN KEY (spaceId) REFERENCES spaces(spaceId)
);

CREATE TABLE user_invitations (
  userInvitationId BIGINT NOT NULL,
  senderId BIGINT NOT NULL,
  receiverId BIGINT NOT NULL,
  spaceId BIGINT NOT NULL,
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
