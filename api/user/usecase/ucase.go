package usecase

import (
	_user "github.com/siends/siends-api/api/user"
	_models "github.com/siends/siends-api/models"
)

type usecase struct {
	User _user.Repository
}

// Input ...
type Input struct {
	User _user.Repository
}

func NewUsecase(input Input) _user.Usecase {
	return &usecase{
		User: input.User,
	}
}

func (ucase *usecase) Signin(doc *_models.UserInput) _models.Response {
	var response _models.Response

	user, err := ucase.User.FindOne(_models.User{Email: doc.Email})
	if err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signin",
			Title:    "Email not registered",
			Status:   400,
			Detail:   "The e-mail address is not registered",
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	if !comparePassword(comparePasswordInput{
		ID:                user.ID,
		UserPassword:      user.Password,
		PasswordToCompare: doc.Password,
	}) {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signin",
			Title:    "Incorrect password",
			Status:   400,
			Detail:   "The password you entered is incorrect",
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	claims := _models.CustomClaims{UserID: user.ID}
	claims.Token, err = createToken(claims)
	if err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signin",
			Title:    "Error generating token",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	response.Data = claims
	return response
}

func (ucase *usecase) Signup(doc *_models.UserInput) _models.Response {
	var response _models.Response
	user := &_models.User{Email: doc.Email}

	if u, _ := ucase.User.FindOne(*user); u != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signup",
			Title:    "Email already exists",
			Status:   400,
			Detail:   "The e-mail address is already registered",
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	if err := ucase.User.InsertOne(user); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signup",
			Title:    "Error creating user",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	password, err := encryptPassword(encryptPasswordInput{
		ID:       user.ID,
		Password: doc.Password,
	})
	if err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signup",
			Title:    "Password encoding error",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	user.Password = password
	if err := ucase.User.Update(user); err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signup",
			Title:    "Error saving password",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	claims := _models.CustomClaims{UserID: user.ID}
	claims.Token, err = createToken(claims)
	if err != nil {
		response.Errors = append(response.Errors, _models.Error{
			Type:     "/errors/signin",
			Title:    "Error generating token",
			Status:   400,
			Detail:   err.Error(),
			Instance: "/api/user/usecase/ucase.go",
		})

		return response
	}

	response.Data = claims
	return response
}
