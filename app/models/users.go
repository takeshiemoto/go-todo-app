package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (uuid,name,email,password,created_at) VALUES (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(&user.ID, &user.UUID, &user.Email, &user.PassWord, &user.CreatedAt)

	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {
	cmd := `DELETE from users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, email, password, created_at FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(&user.ID, &user.UUID, &user.Email, &user.PassWord, &user.CreatedAt)

	return user, err
}

func (u *User) CreateSession() (session Session, err error) {
	session = Session{}

	cmd := `INSERT INTO sessions (uuid, email, user_id, created_at) VALUES (?, ?, ?, ?)`

	_, err = Db.Exec(cmd, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd = `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? AND email = ?`
	err = Db.QueryRow(cmd, u.ID, u.Email).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)

	return session, err
}

func (s *Session) CheckSession() (valid bool, err error) {
	cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`

	err = Db.QueryRow(cmd, s.UUID).Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	if err != nil {
		valid = false

		return
	}

	if s.ID != 0 {
		valid = true
	}

	return valid, err
}

func (s *Session) DeleteSessionByUUID() (err error) {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err = Db.Exec(cmd, s.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (s *Session) GetUserBySession() (user User, err error) {
	user = User{}

	cmd := `SELECT id, uuid, name, email, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, s.UserID).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt)

	return user, err
}
