package entities

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"shape/entities/request"
	"time"
)

var (
	lastAuthUserConfigFilename string = ""
	ErrOnCreateUser                   = errors.New("Create user fail")
)

type (
	User struct {
		UserID    string `json:"user_id" `
		Username  string `json:"username"`
		Password  string `json:"password,omitempty"`
		Fullname  string `json:"fullname"`
		Email     string `json:"email" `
		Disabled  bool   `json:"disabled"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	UserDB struct {
		Users []User `json:"users"`
	}
)

func (c *UserDB) Load(filename string) (*UserDB, error) {
	return readAuthUser(filename)
}

func readAuthUser(filename string) (*UserDB, error) {

	log.Info("Loading user from ", filename)

	confData, jerr := getConfigFromFile(filename)
	if jerr != nil {
		//log.Fatal("Failed to load configuration: ", jerr)
		return nil, jerr
	}
	lastAuthUserConfigFilename = filename

	return decodeAuthUser(confData)
}

func decodeAuthUser(data []byte) (*UserDB, error) {

	config := &UserDB{}

	// try to decode json first and yaml in the next step
	if err := json.Unmarshal(data, &config); err != nil {
		if err = yaml.Unmarshal(data, &config); err != nil {
			return nil, err
		}
		return nil, err
	}

	return config, nil
}

func (c *UserDB) Save() error {
	return c.WriteConfig()
}

func (c *UserDB) WriteConfig() error {
	if len(lastAuthUserConfigFilename) < 3 {
		log.Error("Missing config filename")
		return fmt.Errorf("Filename missing")
	}

	err := json.Unmarshal(c.ConfigJSON(), &c)
	if err != nil {
		return err
	}

	newData, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(lastAuthUserConfigFilename, newData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserDB) ConfigJSON() []byte {
	str, _ := json.Marshal(c)
	return str
}

func (c *UserDB) NewUser(user request.UserRequest) (*User, error) {
	pHashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	if ok := c.CheckUserUnique(user.Username); ok {
		return nil, ErrOnCreateUser
	}
	newuser := &User{
		UserID:    uuid.NewString(),
		Username:  user.Username,
		Password:  string(pHashed),
		CreatedAt: time.Now().String(),
		Fullname:  user.Fullname,
		Email:     user.Email,
	}
	c.Users = append(c.Users, *newuser)
	if err := c.Save(); err != nil {
		return nil, err
	}
	return newuser, nil
}

func (c *UserDB) GetUserByID(id string) (*User, error) {
	cnt := len(c.Users)
	for i := 0; i < cnt; i++ {
		if c.Users[i].UserID == id {
			return &c.Users[i], nil
		}
	}
	return nil, nil
}

func (c *UserDB) GetUserByName(username string) (*User, error) {
	for i := 0; i < len(c.Users); i++ {
		if c.Users[i].Username == username {
			return &c.Users[i], nil
		}
	}
	return nil, nil
}

func (c *UserDB) UpdateUser(id string, input request.UserRequest) error {
	currentUser, err := c.GetUserByID(id)
	if err != nil {
		return err
	}
	if input.Password != "" && input.NewPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(input.Password)); err != nil {
			return err
		}
		pHashed, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		currentUser.Password = string(pHashed)
	}
	if input.Disabled == true {
		currentUser.Disabled = true
	} else {
		currentUser.Disabled = false
	}
	if input.Fullname != "" {
		currentUser.Fullname = input.Fullname
	}
	if input.Email != "" {
		currentUser.Email = input.Email
	}
	currentUser.UpdatedAt = time.Now().String()

	if err := c.Save(); err != nil {
		return err
	}
	return nil
}

func (c *UserDB) DisabledUser(id string) (*User, error) {
	currentUser, err := c.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	currentUser.Disabled = true
	fmt.Println(c)
	if err = c.Save(); err != nil {
		return nil, err
	}
	return currentUser, err
}

func (c *UserDB) GetsPublic() (UserDB, error) {
	PubUser := c
	for i, _ := range PubUser.Users {
		PubUser.Users[i].Password = ""
	}
	return *PubUser, nil
}

func (c *UserDB) GetPublic(username string) (*User, error) {
	PubUser := c
	for i, v := range PubUser.Users {
		if v.Username == username {
			PubUser.Users[i].Password = ""
			return &PubUser.Users[i], nil
		}
	}
	return nil, nil
}

func (c *UserDB) Gets() UserDB {
	return *c
}
func (c *UserDB) CheckUserUnique(username string) bool {
	for _, user := range c.Users {
		if user.Username == username {
			return true
		}
	}
	return false
}
