package database

//func AddNewUser(newUser gateway.UserPlain) error {
//	db, err := sql.Open("mysql", "root:Everything7tays@tcp(127.0.0.1:3306)/blockchain")
//	if err != nil {
//		panic(err.Error())
//		return err
//	}
//	defer db.Close()
//	insert, err := db.Query("INSERT INTO blockchain.user (login, password, balance) " +
//		"VALUES ("+ newUser.Login +", "+ newUser.Password +", 100);")
//	if err != nil {
//		panic(err.Error())
//		return err
//	}
//	defer insert.Close()
//	return nil
//}
