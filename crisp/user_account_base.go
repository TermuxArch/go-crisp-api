// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserAccountData mapping
type UserAccountData struct {
  Data  *UserAccount  `json:"data,omitempty"`
}

// UserAccount mapping
type UserAccount struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// UserAccountCreate mapping
type UserAccountCreate struct {
  Email      *string  `json:"email,omitempty"`
  Passowrd   *string  `json:"password,omitempty"`
  FirstName  *string  `json:"first_name,omitempty"`
  LastName   *string  `json:"last_name,omitempty"`
}


// GetUserAccount resolves the current user account information.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-base-get
func (service *UserService) GetUserAccount() (*UserAccount, *Response, error) {
  url := "user/account"
  req, _ := service.client.NewRequest("GET", url, nil)

  account := new(UserAccountData)
  resp, err := service.client.Do(req, account)
  if err != nil {
    return nil, resp, err
  }

  return account.Data, resp, err
}


// CreateUserAccount creates a new Crisp user account (operator account).
// Reference: https://docs.crisp.im/api/v1/#user-user-account-base-post
func (service *EmailService) CreateUserAccount(email string, password string, firstName string, lastName string) (*Response, error) {
  url := "user/account"
  req, _ := service.client.NewRequest("POST", url, UserAccountCreate{&email, &password, &firstName, &lastName})

  return service.client.Do(req, nil)
}