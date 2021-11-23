package usecase

import (
	"Blockchain/database"
	user "Blockchain/database/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func GetUserInfoUseCase(userId int) (user.UserInfo, error) {
	db, err := database.Connection()
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
	db, err := database.Connection()
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
	db, err := database.Connection()
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
	db, err := database.Connection()
	var randomWallet user.RandomWallet
	if err != nil {
		return randomWallet, err
	}
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
