package controllers

import (
	"encoding/json"
	"fmt"
	"go-todo-app/app/dto"
	"go-todo-app/app/models"
	"log"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var signUpRequestDto dto.SignUpRequestDto

	err := json.Unmarshal(body, &signUpRequestDto)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	user := models.User{
		Name:     signUpRequestDto.Name,
		Email:    signUpRequestDto.Email,
		PassWord: signUpRequestDto.Password,
	}

	if err := user.CreateUser(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func signin(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var signInDto dto.SignInRequestDto

	err := json.Unmarshal(body, &signInDto)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	user, err := models.GetUserByEmail(signInDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user.PassWord == models.Encrypt(signInDto.Password) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		signInResponseDto := dto.SignInResponseDto{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		res, err := json.Marshal(signInResponseDto)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		http.SetCookie(w, &cookie)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(res))

		return
	}

	w.WriteHeader(http.StatusUnauthorized)

	return
}

func signout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		err := session.DeleteSessionByUUID()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
	}

	return
}
