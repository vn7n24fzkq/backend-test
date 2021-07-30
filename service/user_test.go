package service

import (
	"testing"

	"vn7n24fzkq/backend-test/dao"
)

func BenchmarkAuthUser(b *testing.B) {
	//TODO
}

func TestAuthUser(t *testing.T) {
	//TODO
}

func TestCreateUser(t *testing.T) {
	userService := getUserService(t)
	var testUser = dao.User{
		Username:       "testUser",
		PasswordDigest: "password",
		Salt:           "salt",
	}
	var anotherTestUser = dao.User{
		Username:       "anothertestUser",
		PasswordDigest: "password",
		Salt:           "salt",
	}
	var user, err = userService.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Should not get any error when creating user. %s", err)
	} else if user.Username != testUser.Username {
		t.Fatalf("The username of created User is %s, but it should be %s", user.Username, testUser.Username)
	}

	user, err = userService.CreateUser(testUser)
	if err == nil {
		t.Fatalf("Should get an error when creating a duplicated user.")
	}

	user, err = userService.CreateUser(anotherTestUser)
	if err != nil {
		t.Fatalf("Should not get any error when creating a different user.")
	}
}

func TestGetUserByID(t *testing.T) {
	userService := getUserService(t)
	var testUser = dao.User{
		Username:       "testUser",
		PasswordDigest: "password",
		Salt:           "salt",
	}
	var user, err = userService.GetUserByID(1)
	if err == nil {
		t.Fatal("Should get an error when try to get an user which is not exists.")
	}
	user, _ = userService.CreateUser(testUser)
	var createdUserID = user.ID
	user, err = userService.GetUserByID(createdUserID)
	if err != nil {
		t.Fatalf("Should not get any error when getting an exists user. %s", err)
	} else if createdUserID != user.ID {
		t.Fatalf("The user ID is %d, and it should be %d", user.ID, createdUserID)
	}
}

func TestUpdateUserByID(t *testing.T) {
	userService := getUserService(t)
	var testUser = dao.User{
		Username:       "testUser",
		PasswordDigest: "password",
		Salt:           "salt",
	}

	var user, _ = userService.CreateUser(testUser)
	var err = userService.UpdateUserByID(2, testUser)
	if err == nil {
		t.Fatalf("Should get an error when updating a not exists user")
	}
	testUser.Username = "updatedName"
	testUser.PasswordDigest = "updatedPassword"
	testUser.Salt = "updatedSalt"
	err = userService.UpdateUserByID(user.ID, testUser)
	if err != nil {
		t.Fatalf("Should not get any error when updating user. %s", err)
	}

	updatedUser, _ := userService.GetUserByID(user.ID)
	if updatedUser.Username != testUser.Username || updatedUser.PasswordDigest != testUser.PasswordDigest || updatedUser.Salt != testUser.Salt {
		t.Fatalf("Except the ID, user should be same with \n%+v\n,but it is \n%+v", testUser, updatedUser)
	}
}

func TestDeleteUserByID(t *testing.T) {
	userService := getUserService(t)
	var testUser = dao.User{
		Username:       "testUser",
		PasswordDigest: "password",
		Salt:           "salt",
	}
	var user, _ = userService.CreateUser(testUser)
	var err = userService.DeleteUserByID(user.ID)
	if err != nil {
		t.Fatalf("Should not get any error when deleting a exists user. %s", err)
	}
	user, err = userService.GetUserByID(user.ID)
	if err == nil {
		t.Fatalf("Should any error when getting a deleted user. %s", err)
	}
}

func getUserService(t *testing.T) *UserService {
	db := GetTestDB(t)
	// Initialize DAO
	userDAO := dao.NewUserDAO(db)

	return NewUserService(userDAO)
}
