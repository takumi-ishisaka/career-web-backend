package persistence

import (
	"database/sql"
	"net/smtp"

	"github.com/pkg/errors"

	"github.com/nokin-all-of-career/career-web-backend/domain"
	"github.com/nokin-all-of-career/career-web-backend/domain/repository"
	"github.com/nokin-all-of-career/career-web-backend/infra"
)

type userPersistence struct {
	DB *sql.DB
}

// NewUserPersistence : make user persistence
func NewUserPersistence(DB *sql.DB) repository.UserRepository {
	return &userPersistence{
		DB: DB,
	}
}

func (up userPersistence) Insert(userID, email, password string) error {
	stmt, err := up.DB.Prepare("INSERT INTO user(user_id, email, password, status,last_login_time) VALUES(?, ?, ?, ?,NULL)")
	if err != nil {
		return errors.Wrap(err, "userPersistence.Insert()")
	}

	_, err = stmt.Exec(userID, email, password, domain.USER)
	if err != nil {
		return errors.Wrap(err, "userApplication.Insert()")
	}

	return nil
}

func (up userPersistence) SelectByEmail(email string) (*domain.User, error) {
	row := up.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)
	user, err := convertToUser(row)
	if err != nil {
		return nil, errors.Wrap(err, "userPersistence.SelectByEmail()")
	}

	return user, nil
}

func (up userPersistence) SelectByPrimaryKey(userID string) (*domain.User, error) {
	row := up.DB.QueryRow("SELECT * FROM user WHERE user_id = ?", userID)
	user, err := convertToUser(row)
	if err != nil {
		return nil, errors.Wrap(err, "userPersistence.SelectByPrimaryKey()")
	}

	return user, nil
}

func (up userPersistence) Update(userID, email string) error {
	stmt, err := up.DB.Prepare("UPDATE user SET email=? WHERE user_id=?")
	if err != nil {
		return errors.Wrap(err, "userPersistence.Update()")
	}

	_, err = stmt.Exec(email, userID)
	if err != nil {
		return errors.Wrap(err, "userApplication.Update()")
	}

	return nil
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	nullTime := new(sql.NullTime)
	err := row.Scan(&user.UserID, &user.Email, &user.Password, &user.Status, nullTime)
	if nullTime.Valid {
		user.LastLoginTime = nullTime.Time
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, nil
		}
		return nil, err
	}

	return &user, nil
}

func (up userPersistence) SendValidationEmail(from, to string, url string) error {
	msg := []byte("" +
		"From: オルキャリ運営 <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: 【オルキャリ】メールアドレスの確認\r\n" +
		"\r\n" +
		to + "様\r\n" +
		"\r\n" +
		"この度はオルキャリにご登録いただき誠にありがとうございます。\r\n" +
		"本メールはメールアドレス確認用のメールになります。\r\n" +
		"以下のURLをクリックしていただき、メールアドレスの確認を完了してください。\r\n" +
		"\r\n" +
		url + "\r\n" +
		"")
	err := smtp.SendMail("smtp.gmail.com:587", infra.Auth, from, []string{to}, msg)
	if err != nil {
		return errors.Wrap(err, "userPersistence.SendvalidationEmail()")
	}
	return nil
}
