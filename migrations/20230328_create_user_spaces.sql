-- +migrate Up
CREATE TABLE users (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
);

CREATE TABLE spaces (
  `space_id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`space_id`)
);

CREATE TABLE space_members (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT,
  `space_id` BIGINT NOT NULL,
  PRIMARY KEY (`user_id`, `space_id`),
  FOREIGN KEY (`user_id`) REFERENCES users(`user_id`),
  FOREIGN KEY (`space_id`) REFERENCES spaces(`space_id`)
);

CREATE TABLE user_invitations (
  `user_invitation_id` BIGINT NOT NULL AUTO_INCREMENT,
  `sender_id` BIGINT NOT NULL,
  `receiver_id` BIGINT NOT NULL,
  `space_id` BIGINT NOT NULL,
  status VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_invitation_id`),
  FOREIGN KEY (`sender_id`) REFERENCES users(`user_id`),
  FOREIGN KEY (`receiver_id`) REFERENCES users(`user_id`),
  FOREIGN KEY (`space_id`) REFERENCES spaces(`space_id`)
);

-- +migrate Down
DROP TABLE user_invitations;
DROP TABLE space_members;
DROP TABLE spaces;
DROP TABLE users;
