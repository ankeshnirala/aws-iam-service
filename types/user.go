package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const rootRole string = "root"
const role string = "member"

// user_type?: string;
// username?: string;
// email?: string;
// password?: string;
// mobile?: number | string;
// fcm_token?: string;
// referral_code?: string;
// email_verified?: string;
// sign_in_count?: number;
// current_sign_in_ip?: string | string[];
// last_sign_in_ip?: string | string[];
// current_sign_in_at?: Date;
// last_sign_in_at?: Date;
// is_blocked?: string;
// refferedCode?: string;
// device_type?: string;
// status?: string;
// current_latitude?: string | number;
// current_longitude?: string | number;
// soc_id?: string;
// soc_log_type?: string;
// password_reset_token?: string;
// password_reset_expires?: string | number;

type UserPrimaryDetails struct {
	ID        primitive.ObjectID `json:"id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"email,required"`
	Password  string             `json:"password" validate:"required"`
	Username  string             `json:"username"`
	UserType  string             `json:"userType"`
	Phone     string             `json:"mobile"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type UserSecondaryDetails struct {
	ID           primitive.ObjectID `json:"id"`
	UserID       primitive.ObjectID `json:"userID"`
	EmailVerfied bool               `json:"emailVerified"`
	IsBlocked    bool               `json:"isBlocked"`
	SignInCount  int                `json:"signInCount"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserLoginHistory struct {
	ID     primitive.ObjectID `json:"id"`
	UserID primitive.ObjectID `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email" validate:"email,required"`
	Password  string             `json:"password" validate:"required"`
	Role      string             `json:"role" validate:"required"`
	CreatedAt time.Time          `json:"createdAt"`
	CreatedBy primitive.ObjectID `json:"createdBy"`
}

func (u *User) ValidPassword(pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw)) == nil
}

func NewRootUser(firstName, lastName, email, password string) (*User, error) {

	encPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      rootRole,
		Password:  string(encPw),
		CreatedAt: time.Now().UTC(),
	}, nil
}

func NewUser(firstName, lastName, email, password, createdBy string) (*User, error) {

	encPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	objectId, err := primitive.ObjectIDFromHex(createdBy)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Role:      role,
		Password:  string(encPw),
		CreatedAt: time.Now().UTC(),
		CreatedBy: objectId,
	}, nil
}
