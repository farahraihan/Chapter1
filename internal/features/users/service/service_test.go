package service_test

// import (
// 	"chapter1/internal/features/users"
// 	"chapter1/internal/features/users/service"
// 	"chapter1/mocks"
// 	"errors"
// 	"mime/multipart"
// 	"net/textproto"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestLogin(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	LoginRequest := users.User{
// 		Email:    "test@gmail.com",
// 		Password: "password",
// 	}

// 	Result := users.User{
// 		Email:    "test@gmail.com",
// 		Password: "$2a$12$abc",
// 	}

// 	token := "token_string"

// 	t.Run("success login", func(t *testing.T) {
// 		mockQuery.On("Login", LoginRequest.Email).Return(Result, nil).Once()
// 		mockPass.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(nil).Once()
// 		mockJwt.On("GenerateJwt", Result.ID).Return(token, nil).Once()

// 		Result, token, err := userServices.Login(LoginRequest.Email, LoginRequest.Password)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, Result, token)
// 	})

// 	t.Run("fail login user not found", func(t *testing.T) {
// 		expectedError := errors.New("login failed, please try again later")
// 		mockQuery.On("Login", LoginRequest.Email).Return(users.User{}, expectedError).Once()

// 		Result, token, err := userServices.Login(LoginRequest.Email, LoginRequest.Password)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, expectedError, err.Error())
// 		assert.Equal(t, users.User{}, Result)
// 		assert.Equal(t, "", token)
// 	})

// 	t.Run("fail login password mismatch", func(t *testing.T) {
// 		expectedError := errors.New("invalid credentials")
// 		mockQuery.On("Login", LoginRequest.Email).Return(Result, nil).Once()
// 		mockPass.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(expectedError).Once()

// 		Result, token, err := userServices.Login(LoginRequest.Email, LoginRequest.Password)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, expectedError, err.Error())
// 		assert.Equal(t, users.User{}, Result)
// 		assert.Equal(t, "", token)
// 	})

// 	t.Run("fail generate JWT", func(t *testing.T) {
// 		expectedError := errors.New("login failed, please try again later")
// 		mockQuery.On("Login", LoginRequest.Email).Return(Result, nil).Once()
// 		mockPass.On("ComparePassword", []byte(Result.Password), []byte(LoginRequest.Password)).Return(nil).Once()
// 		mockJwt.On("GenerateJwt", Result.ID).Return("", expectedError).Once()

// 		Result, token, err := userServices.Login(LoginRequest.Email, LoginRequest.Password)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, expectedError, err.Error())
// 		assert.Equal(t, users.User{}, Result)
// 		assert.Equal(t, "", token)
// 	})

// }

// func TestRegister(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	RegisterRequest := users.User{
// 		Name:     "farah",
// 		Email:    "farah@gmail.com",
// 		Phone:    "081765993766",
// 		Password: "farah123",
// 		Image:    "",
// 	}

// 	h := make(textproto.MIMEHeader)
// 	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
// 	h.Set("Content-Type", "image/png")

// 	file := &multipart.FileHeader{
// 		Filename: "test.png",
// 		Header:   h,
// 		Size:     123,
// 	}

// 	src, _ := file.Open()
// 	defer src.Close()

// 	t.Run("success register", func(t *testing.T) {
// 		mockPass.On("GeneratePassword", RegisterRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
// 		mockQuery.On("Register", RegisterRequest).Return(nil).Once()

// 		err := userServices.Register(RegisterRequest, src, file.Filename)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("fail generate password", func(t *testing.T) {
// 		expectedError := errors.New("registration failed, please try again later")
// 		mockPass.On("GeneratePassword", RegisterRequest.Password).Return([]byte("another password"), expectedError).Once()

// 		err := userServices.Register(RegisterRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("fail upload to cloudinary", func(t *testing.T) {
// 		expectedError := errors.New("failed to upload image, please try again later")
// 		mockPass.On("GeneratePassword", RegisterRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", expectedError).Once()

// 		err := userServices.Register(RegisterRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("fail register", func(t *testing.T) {
// 		expectedError := errors.New("registration failed, please try again later")
// 		mockPass.On("GeneratePassword", RegisterRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
// 		mockQuery.On("Register", RegisterRequest).Return(expectedError).Once()

// 		err := userServices.Register(RegisterRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)

// 	})
// }

// func TestUpdateUser(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	UpdateRequest := users.User{
// 		Name:     "farah",
// 		Email:    "farah@gmail.com",
// 		Phone:    "081765993766",
// 		Password: "farah123",
// 		Image:    "",
// 	}

// 	h := make(textproto.MIMEHeader)
// 	h.Set("Content-Disposition", `form-data; name="file"; filename="test.png"`)
// 	h.Set("Content-Type", "image/png")

// 	file := &multipart.FileHeader{
// 		Filename: "test.png",
// 		Header:   h,
// 		Size:     123,
// 	}

// 	src, _ := file.Open()
// 	defer src.Close()

// 	t.Run("fail update password", func(t *testing.T) {
// 		expectedError := errors.New("update failed, please try again later")
// 		mockPass.On("GeneratePassword", UpdateRequest.Password).Return([]byte("another password"), expectedError).Once()

// 		err := userServices.UpdateUser(UpdateRequest.ID, UpdateRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("fail upload to cloudinary", func(t *testing.T) {
// 		expectedError := errors.New("failed to upload image, please try again later")
// 		mockPass.On("GeneratePassword", UpdateRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", expectedError).Once()

// 		err := userServices.UpdateUser(UpdateRequest.ID, UpdateRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("fail update user", func(t *testing.T) {
// 		expectedError := errors.New("update failed, please try again later")
// 		mockPass.On("GeneratePassword", UpdateRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
// 		mockQuery.On("UpdateUser", UpdateRequest.ID, UpdateRequest).Return(expectedError).Once()

// 		err := userServices.UpdateUser(UpdateRequest.ID, UpdateRequest, src, file.Filename)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)

// 	})

// 	t.Run("success update user", func(t *testing.T) {
// 		mockPass.On("GeneratePassword", UpdateRequest.Password).Return([]byte("farah123"), nil).Once()
// 		mockCld.On("UploadToCloudinary", src, file.Filename).Return("", nil).Once()
// 		mockQuery.On("UpdateUser", UpdateRequest.ID, UpdateRequest).Return(nil).Once()

// 		err := userServices.UpdateUser(UpdateRequest.ID, UpdateRequest, src, file.Filename)
// 		assert.Nil(t, err)
// 	})
// }

// func TestDeleteUser(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	var adminID uint = 1
// 	var memberID uint = 2

// 	t.Run("user not admin", func(t *testing.T) {
// 		expectedError := errors.New("access denied")
// 		mockQuery.On("IsAdmin", memberID).Return(false, expectedError).Once()

// 		err := userServices.DeleteUser(memberID, adminID)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("fail delete user", func(t *testing.T) {
// 		expectedError := errors.New("delete failed, please try again later")
// 		mockQuery.On("IsAdmin", adminID).Return(true, nil).Once()
// 		mockQuery.On("DeleteUser", memberID).Return(expectedError).Once()

// 		err := userServices.DeleteUser(adminID, memberID)
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 	})

// 	t.Run("success delete user", func(t *testing.T) {
// 		mockQuery.On("IsAdmin", adminID).Return(true, nil).Once()
// 		mockQuery.On("DeleteUser", memberID).Return(nil).Once()

// 		err := userServices.DeleteUser(adminID, memberID)
// 		assert.Nil(t, err)
// 	})
// }

// func TestGetUserByID(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	GetResponse := users.User{
// 		ID:      1,
// 		Name:    "Farah",
// 		Email:   "farah@gmail.com",
// 		Phone:   "087990988990",
// 		IsAdmin: true,
// 		Image:   "https:\\image.png",
// 	}

// 	t.Run("success get user by ID", func(t *testing.T) {
// 		mockQuery.On("GetUserByID", GetResponse.ID).Return(users.User{}, nil).Once()

// 		result, err := userServices.GetUserByID(GetResponse.ID)
// 		assert.Nil(t, err)
// 		assert.NotNil(t, result)
// 	})

// 	t.Run("fail get user by ID", func(t *testing.T) {
// 		expectedError := errors.New("failed to retrieve user, please try again later")
// 		mockQuery.On("GetUserByID", GetResponse.ID).Return(users.User{}, expectedError).Once()

// 		result, err := userServices.GetUserByID(GetResponse.ID)
// 		assert.NotNil(t, err)
// 		assert.ErrorContains(t, expectedError, err.Error())
// 		assert.Equal(t, users.User{}, result)
// 	})

// }

// func TestGetAllUser(t *testing.T) {
// 	mockQuery := mocks.NewUQuery(t)
// 	mockPass := mocks.NewPassUtilInterface(t)
// 	mockJwt := mocks.NewJwtUtilityInterface(t)
// 	mockCld := mocks.NewCloudinaryUtilityInterface(t)

// 	userServices := service.NewUserServices(mockQuery, mockPass, mockJwt, mockCld)

// 	GetResponse := []users.User{
// 		{
// 			ID:      1,
// 			Name:    "Farah",
// 			Email:   "farah@gmail.com",
// 			Phone:   "087990988990",
// 			IsAdmin: true,
// 			Image:   "https:\\image.png",
// 		},
// 		{
// 			ID:      2,
// 			Name:    "Dion",
// 			Email:   "dion@gmail.com",
// 			Phone:   "087960988990",
// 			IsAdmin: false,
// 			Image:   "https:\\image.png",
// 		},
// 	}

// 	var userID uint = 3

// 	t.Run("user not admin", func(t *testing.T) {
// 		expectedError := errors.New("access denied")
// 		mockQuery.On("IsAdmin", userID).Return(false, expectedError).Once()

// 		users, totalItems, err := userServices.GetAllUsers(userID, 10, 1, "")
// 		assert.NotNil(t, err)
// 		assert.Equal(t, expectedError, err)
// 		assert.Equal(t, 0, totalItems)
// 		assert.Nil(t, users)
// 	})

// 	t.Run("fail get all user data", func(t *testing.T) {
// 		expectedError := errors.New("failed to retrieve users, please try again later")
// 		mockQuery.On("IsAdmin", userID).Return(true, nil).Once()
// 		mockQuery.On("GetAllUsers", 10, 1, "").Return(nil, 0, expectedError).Once()

// 		users, totalItems, err := userServices.GetAllUsers(userID, 10, 1, "")
// 		assert.NotNil(t, err)
// 		assert.Equal(t, 0, totalItems)
// 		assert.Equal(t, expectedError, err)
// 		assert.Nil(t, nil, users)
// 	})

// 	t.Run("success get all user data", func(t *testing.T) {
// 		mockQuery.On("IsAdmin", userID).Return(true, nil).Once()
// 		mockQuery.On("GetAllUsers", 10, 1, "").Return(GetResponse, 2, nil).Once()

// 		users, totalItems, err := userServices.GetAllUsers(userID, 10, 1, "")
// 		assert.Nil(t, err)
// 		assert.Equal(t, 2, totalItems)
// 		assert.Equal(t, GetResponse, users)
// 	})
// }
