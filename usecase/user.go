package usecase

import (
	user "Blockchain/database/user"
	"Blockchain/settings"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

func GetUserInfoUseCase(userId int) (user.UserInfo, error) {
	db := settings.Db
	var err error
	var userInfo user.UserInfo
	if err != nil {
		return userInfo, err
	}
	userInfo, err = user.GetUserInfo(db, userId)
	if err != nil {
		return userInfo, err
	}
	return userInfo, nil
}

func AddNewUserUseCase(newUserPlain user.UserPlain) (user.User, error) {
	db := settings.Db
	var err error
	log.Printf("Adding new user. Login: %s, Password: %s", newUserPlain.Login, newUserPlain.Password)
	var newUser user.User
	if err != nil {
		return newUser, err
	}
	wallet := uuid.New().String()
	err = user.AddNewUser(db, newUserPlain, wallet)
	if err != nil {
		return newUser, err
	}
	userId, err := user.GetUserId(db, wallet)
	if err != nil {
		return newUser, err
	}
	err = user.AddNewSender(db, userId)
	if err != nil {
		return newUser, err
	}
	err = user.AddNewRecipient(db, userId)
	if err != nil {
		return newUser, err
	}
	newUser = user.User{newUserPlain, userId, wallet, 100}
	return newUser, nil
}

func LoginUserUseCase(userPlain user.UserPlain) (user.User, error) {
	db := settings.Db
	var err error
	log.Printf("Logging in user. Login: %s, Password: %s", userPlain.Login, userPlain.Password)
	var userToLogin user.User
	if err != nil {
		return userToLogin, err
	}
	userToLogin, err = user.CheckUser(db, userPlain)
	if err != nil {
		return userToLogin, err
	}
	return userToLogin, nil
}

func RandomWalletUseCase(userId int) (user.RandomWallet, error) {
	db := settings.Db
	var err error
	var randomWallet user.RandomWallet
	for {
		randomWallet, err = user.GetRandomWallet(db)
		if err != nil {
			return randomWallet, err
		}
		if randomWallet.UserId != userId {
			return randomWallet, nil
		}
	}
}
