-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema career_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema career_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `career_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin ;
USE `career_db` ;

-- -----------------------------------------------------
-- Table `career_db`.`category`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`category` (
  `category_id` VARCHAR(32) NOT NULL,
  `name` VARCHAR(32) NOT NULL,
  `goal` VARCHAR(256) NOT NULL,
  PRIMARY KEY (`category_id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;


-- -----------------------------------------------------
-- Table `career_db`.`action`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`action` (
  `action_id` VARCHAR(32) NOT NULL,
  `category_id` VARCHAR(32) NOT NULL,
  `title` VARCHAR(128) NOT NULL,
  `content` VARCHAR(512) NOT NULL,
  `standard_time` VARCHAR(16) NOT NULL,
  `action_type` INT NOT NULL,
  `url` VARCHAR(512) NULL,
  `after` VARCHAR(512) NULL,
  PRIMARY KEY (`action_id`),
  INDEX `tag_id_idx` (`category_id` ASC),
  CONSTRAINT `fk_category_action_category`
    FOREIGN KEY (`category_id`)
    REFERENCES `career_db`.`category` (`category_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;


-- -----------------------------------------------------
-- Table `career_db`.`user`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`user` (
  `user_id` VARCHAR(64) NOT NULL,
  `email` VARCHAR(64) NOT NULL,
  `password` VARCHAR(256) NOT NULL,
  `status` INT NULL DEFAULT NULL,
  `last_login_time` DATETIME,
  PRIMARY KEY (`user_id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;


-- -----------------------------------------------------
-- Table `career_db`.`profile`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`profile` (
  `user_id` VARCHAR(64) NOT NULL,
  `name` VARCHAR(64) NOT NULL,
  `university` VARCHAR(64) NOT NULL,
  `major` VARCHAR(64) NOT NULL,
  `graduation_year` INT NOT NULL,
  `aspiring_occupation` VARCHAR(64) NOT NULL,
  `aspiring_field` VARCHAR(64) NOT NULL,
  `sentence` VARCHAR(1024) NULL DEFAULT NULL,
  `image_path` VARCHAR(128) NULL DEFAULT NULL,
  `job_hunting_status` INT NOT NULL DEFAULT '0',
  `deviation_value` DECIMAL(20,10) NOT NULL,
  PRIMARY KEY (`user_id`),
  INDEX `user_id_idx` (`user_id` ASC),
  CONSTRAINT `fk_user_profile_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `career_db`.`user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;


-- -----------------------------------------------------
-- Table `career_db`.`user_action`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`user_action` (
  `user_action_id` VARCHAR(64)  NOT NULL,
  `user_id` VARCHAR(64) NOT NULL,
  `action_id` VARCHAR(32) NOT NULL,
  `status` INT NOT NULL,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `do` VARCHAR(1024) NULL DEFAULT NULL,
  `reflection` VARCHAR(1024) NULL DEFAULT NULL,
  `next_action` VARCHAR(1024) NULL DEFAULT NULL,
  `evaluate_value` INT NULL DEFAULT NULL,
  PRIMARY KEY (`user_action_id`),
  INDEX `action_id_idx` (`action_id` ASC),
  INDEX `user_id_idx` (`user_id` ASC),
  UNIQUE INDEX `user_action_id_UNIQUE` (`user_action_id` ASC),
  CONSTRAINT `fk_action_user_action`
    FOREIGN KEY (`action_id`)
    REFERENCES `career_db`.`action` (`action_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `fk_user_action_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `career_db`.`user` (`user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;


-- -----------------------------------------------------
-- Table `career_db`.`feedback`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `career_db`.`feedback` (
  `feedback_id` VARCHAR(64) NOT NULL,
  `user_action_id` VARCHAR(64)  NOT NULL,
  `comment` VARCHAR(1024)  NOT NULL,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  INDEX `fk_feedback_user_action1_idx` (`user_action_id` ASC),
  PRIMARY KEY (`feedback_id`),
  CONSTRAINT `fk_feedback_user_action1`
    FOREIGN KEY (`user_action_id`)
    REFERENCES `career_db`.`user_action` (`user_action_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_bin;



SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
