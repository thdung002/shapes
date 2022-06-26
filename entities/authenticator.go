package entities

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"shape/entities/request"
	"time"
)

var (
	ErrModifyOther       = errors.New("Modifying other user is not allowed.")
	ErrModifyAdmin       = errors.New("Modifying admin user is not allowed.")
	ErrForbidden         = errors.New("Forbidden")
	ErrExpiredToken      = errors.New("Token is expired")
	ErrCredentialMissing = errors.New("Missing credential while generating token")
	ErrOnLogin           = errors.New("Login fail")
)

type (
	JwtClaims struct {
		Username string
		Deleted  bool
		Fullname string
		Email    string
		jwt.StandardClaims
	}

	UserIntf interface {
		GetUserByID(string) (*User, error)
		GetsPublic() (UserDB, error)
		GetUserByName(string) (*User, error)
		Gets() UserDB
		UpdateUser(string, request.UserRequest) error
		NewUser(userRequest request.UserRequest) (*User, error)
		DisabledUser(string) (*User, error)
		GetPublic(username string) (*User, error)
	}
	Auth struct {
		DB           UserIntf
		secret       []byte
		nonce        int
		Username     string
		Email        string
		Fullname     string
		Deleted      bool
		AccessToken  *jwt.Token
		RefreshToken *jwt.Token
	}
	Options struct {
		Secret      string
		HmacPadding string
		Nonce       int
	}
)

func NewAuth(userCfg string, options *Options) *Auth {
	cfg := UserDB{}
	load, err := cfg.Load(userCfg)
	if err != nil {
		return nil
	}
	return &Auth{
		DB:     load,
		secret: []byte(options.Secret),
		nonce:  options.Nonce,
	}
}

func (d *Auth) authenticate(username, password string) (*User, error) {
	user, err := d.DB.GetUserByName(username)
	if user == nil {
		return nil, ErrOnLogin
	}
	if err != nil {
		return nil, ErrOnLogin
	}
	if user.Disabled {
		return nil, ErrOnLogin
	}
	// Compare bcrypt hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrOnLogin
	}
	return user, nil
}

func (d *Auth) Authenticate(username, password string) (bool, error) {
	user, err := d.authenticate(username, password)
	if err != nil {
		return false, err
	}
	return user != nil, nil
}

func (d *Auth) AuthenticateJWT(username, password string) (*map[string]string, error) {
	user, err := d.authenticate(username, password)
	if err != nil {
		return nil, err
	}
	d.Username = username

	d.Fullname = user.Fullname
	d.Email = user.Email
	return d.GenerateJWT()
}

func (d *Auth) GenerateJWT() (*map[string]string, error) {
	if d.Username != "" {
		acc, err := d.generateAccessToken()
		if err != nil {
			return nil, err
		}
		ref, err := d.generateRefreshToken()
		if err != nil {
			return nil, err
		}
		return &map[string]string{"token": acc, "refresh_token": ref}, nil
	}
	return nil, ErrCredentialMissing
}

func (d *Auth) generateRefreshToken() (string, error) {
	if d.Username != "" {
		claims := &JwtClaims{
			d.Username,
			d.Deleted,
			d.Fullname,
			d.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			},
		}
		d.RefreshToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		refreshTokenStr, err := d.RefreshToken.SignedString(d.secret)
		if err != nil {
			return "", err
		}
		return refreshTokenStr, nil
	}
	return "", ErrCredentialMissing
}

func (d *Auth) generateAccessToken() (string, error) {
	if d.Username != "" {
		claims := &JwtClaims{
			d.Username,
			d.Deleted,
			d.Fullname,
			d.Email,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			},
		}
		d.AccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenStr, err := d.AccessToken.SignedString(d.secret)
		if err != nil {
			return "", err
		}
		return accessTokenStr, nil
	}
	return "", ErrCredentialMissing
}

func (d *Auth) ValidateToken(token string) (jwt.MapClaims, error) {
	res, err := d.Parse(token)
	if err != nil {
		return nil, err
	}
	return res.Claims.(jwt.MapClaims), nil
}

func (d *Auth) Parse(token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return d.secret, nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (d *Auth) RenewToken(token string) (string, error) {
	_, err := d.ValidateToken(token)
	if err != nil {
		return "", err
	}
	return d.generateAccessToken()
}
