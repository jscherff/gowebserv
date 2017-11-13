// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package service

import (
	`crypto/rsa`
	`net/http`
	`time`
	jwt `github.com/dgrijalva/jwt-go`
)

// ModelService is an interface to help manage database models.
type ModelService interface {
	Create(token *jwt.Token) (*http.Cookie, error)
	Extract(req *http.Request) (*http.Cookie, error)
}

// modelService is a service that implements the ModelService interface.
type modelService struct {
	ModelService
	Tables []string
	Columns map[string][]string
}

// NewAuthTokenService returns an object that implements the AuthTokenService interface.
func NewModelService(driver, dsn string) *modelService {
	return &modelService{}
}

// Create generates a new authentication http.Cookie from a jwt.Token.
func (this *modelService) Create(token *jwt.Token, key *rsa.PrivateKey) (*http.Cookie, error) {

	tokenString, err := token.SignedString(key)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name: `Auth`,
		Value: tokenString,
		Expires: time.Now().Add(time.Hour * 1),
		HttpOnly: true,
	}

	return cookie, nil
}

// Extract extracts the 'Auth' http.Cookie from an http.Request.
func (this *modelService) Extract(req *http.Request) (*http.Cookie, error) {

	if cookie, err := req.Cookie(`Auth`); err != nil {
		return nil, err
	} else {
		return cookie, nil
	}
}
