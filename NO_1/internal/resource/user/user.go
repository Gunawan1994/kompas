package user

import (
	"os"
	"time"

	"kompas/internal/service/login"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"

	"kompas/pkg/cache"
	"kompas/pkg/encrypt"
	"kompas/pkg/errors"
)

// Resource class
type Resource struct {
	db    *sqlx.DB
	redis *cache.Redis
}

// New will return object of Resource class
func New(db *sqlx.DB, redis *cache.Redis) *Resource {
	return &Resource{
		db:    db,
		redis: redis,
	}
}

//Login user
func (r *Resource) Login(email string, password string) (login.ModelLogin, error) {
	user := login.ModelLogin{}
	sql := `SELECT 
				email,
				password, 
				id,
				role,
				username
			FROM 
				public.user 
			WHERE
				email = $1`

	err := r.db.Get(&user, sql, email)
	if err != nil {
		return user, errors.Wrap(err)
	}

	match := encrypt.CheckPasswordHash(password, user.Password)
	if match != true {
		return user, errors.New("wrong credential")
	}

	tokenData, err := createToken(user.Id, user.Role)
	if err != nil {
		return user, errors.Wrap(err)
	}

	user.TokenInfo = tokenData

	if err != nil {
		return user, errors.Wrap(err)
	}
	return user, errors.Wrap(err)
}

func createToken(nirp uint64, role string) (login.TokenDetails, error) {
	td := login.TokenDetails{}
	td.AtExpires = 60 * 60 * 24     // 1 day(second, minute, hour)
	td.RtExpires = 60 * 60 * 24 * 7 // 1 week (second, minute, hour, day)
	AtExp := time.Now().Add(time.Hour * 24).Unix()
	RtExp := time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["role"] = role
	atClaims["user_id"] = nirp
	atClaims["exp"] = AtExp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return td, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = nirp
	rtClaims["exp"] = RtExp
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

	if err != nil {
		return td, err
	}
	return td, nil
}
