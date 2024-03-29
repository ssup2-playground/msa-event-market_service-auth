package service

import (
	"context"
	"testing"

	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/ssup2-playground/msa-event-market_service-auth/internal/domain/entity"
	"github.com/ssup2-playground/msa-event-market_service-auth/internal/domain/repo"
	"github.com/ssup2-playground/msa-event-market_service-auth/internal/domain/repo/mocks"
	"github.com/ssup2-playground/msa-event-market_service-auth/internal/test"
)

func TestUser(t *testing.T) {
	suite.Run(t, new(userSuite))
}

type userSuite struct {
	suite.Suite

	dbTx           mocks.DBTx
	outboxRepo     mocks.OutboxRepo
	userInfoRepo   mocks.UserInfoRepo
	userSecretRepo mocks.UserSecretRepo

	userService UserService
}

func (u *userSuite) SetupTest() {
	// Init transaction, repo
	u.dbTx = mocks.DBTx{}
	u.outboxRepo = mocks.OutboxRepo{}
	u.userInfoRepo = mocks.UserInfoRepo{}
	u.userSecretRepo = mocks.UserSecretRepo{}

	// Set nooptracer
	opentracing.SetGlobalTracer(opentracing.NoopTracer{})

	// Init service
	u.userService = NewUserServiceImp(&u.dbTx, &u.outboxRepo, &u.userInfoRepo, &u.userInfoRepo, &u.userSecretRepo, &u.userSecretRepo)
}

func (u *userSuite) TestListUserSuccess() {
	u.userInfoRepo.On("List", context.Background(), 0, 50).Return([]entity.UserInfo{
		{
			ID:      test.UserIDCorrect,
			LoginID: test.UserLoginIDCorrect,
			Role:    test.UserRoleCorrect,
			Phone:   test.UserPhoneCorrect,
			Email:   test.UserEmailCorrect,
		},
		{
			ID:      test.UserIDCorrect2,
			LoginID: test.UserLoginIDCorrect2,
			Role:    test.UserRoleCorrect,
			Phone:   test.UserPhoneCorrect,
			Email:   test.UserEmailCorrect,
		},
	}, nil)

	userInfos, err := u.userService.ListUser(context.Background(), 0, 0)
	require.NoError(u.T(), err)
	require.Equal(u.T(), test.UserIDCorrect, userInfos[0].ID)
	require.Equal(u.T(), test.UserLoginIDCorrect, userInfos[0].LoginID)
	require.Equal(u.T(), test.UserIDCorrect2, userInfos[1].ID)
	require.Equal(u.T(), test.UserLoginIDCorrect2, userInfos[1].LoginID)
}

func (u *userSuite) TestListUserServerError() {
	u.userInfoRepo.On("List", context.Background(), 0, 50).Return(nil, repo.ErrServerError)

	_, err := u.userService.ListUser(context.Background(), 0, 0)
	require.Equal(u.T(), ErrRepoServerError, err)
}

func (u *userSuite) TestCreateUserSuccess() {
	userInfo := &entity.UserInfo{
		LoginID: test.UserLoginIDCorrect,
		Role:    test.UserRoleCorrect,
		Phone:   test.UserPhoneCorrect,
		Email:   test.UserEmailCorrect,
	}

	u.dbTx.On("Begin").Return(&u.dbTx, nil)
	u.userInfoRepo.On("WithTx", mock.Anything).Return(&u.userInfoRepo)
	u.userInfoRepo.On("Create", context.Background(), userInfo).Return(nil)
	u.userSecretRepo.On("WithTx", mock.Anything).Return(&u.userSecretRepo)
	u.userSecretRepo.On("Create", context.Background(), mock.Anything).Return(nil)
	u.outboxRepo.On("WithTx", mock.Anything).Return(&u.outboxRepo)
	u.outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
	u.dbTx.On("Commit").Return(nil)

	userInfo, err := u.userService.CreateUser(context.Background(), userInfo, test.UserPasswdCorrect)
	require.NoError(u.T(), err)
	require.Equal(u.T(), test.UserLoginIDCorrect, userInfo.LoginID)
	require.Equal(u.T(), test.UserRoleCorrect, userInfo.Role)
	require.Equal(u.T(), test.UserPhoneCorrect, userInfo.Phone)
	require.Equal(u.T(), test.UserEmailCorrect, userInfo.Email)
}

func (u *userSuite) TestGetUserSuccess() {
	u.userInfoRepo.On("Get", context.Background(), test.UserIDCorrect).Return(&entity.UserInfo{
		ID:      test.UserIDCorrect,
		LoginID: test.UserLoginIDCorrect,
		Role:    test.UserRoleCorrect,
		Phone:   test.UserPhoneCorrect,
		Email:   test.UserEmailCorrect,
	}, nil)

	userInfo, err := u.userService.GetUser(context.Background(), test.UserIDCorrect)
	require.NoError(u.T(), err)
	require.Equal(u.T(), test.UserIDCorrect, userInfo.ID)
	require.Equal(u.T(), test.UserLoginIDCorrect, userInfo.LoginID)
	require.Equal(u.T(), test.UserRoleCorrect, userInfo.Role)
	require.Equal(u.T(), test.UserPhoneCorrect, userInfo.Phone)
	require.Equal(u.T(), test.UserEmailCorrect, userInfo.Email)
}

func (u *userSuite) TestGetUserRepoNotFoundError() {
	u.userInfoRepo.On("Get", context.Background(), test.UserIDCorrect).Return(&entity.UserInfo{}, repo.ErrNotFound)

	_, err := u.userService.GetUser(context.Background(), test.UserIDCorrect)
	require.Equal(u.T(), ErrRepoNotFound, err)
}

func (u *userSuite) TestUpdateUserSuccess() {
	userInfo := &entity.UserInfo{
		ID:      test.UserIDCorrect,
		LoginID: test.UserLoginIDCorrect,
		Role:    test.UserRoleCorrect,
		Phone:   test.UserPhoneCorrect,
		Email:   test.UserEmailCorrect,
	}

	u.dbTx.On("Begin").Return(&u.dbTx, nil)
	u.userInfoRepo.On("WithTx", mock.Anything).Return(&u.userInfoRepo)
	u.userInfoRepo.On("Get", context.Background(), test.UserIDCorrect).Return(nil, nil)
	u.userInfoRepo.On("WithTx", mock.Anything).Return(&u.userInfoRepo)
	u.userInfoRepo.On("Update", context.Background(), userInfo).Return(nil)
	u.userSecretRepo.On("WithTx", mock.Anything).Return(&u.userSecretRepo)
	u.userSecretRepo.On("Update", context.Background(), mock.Anything).Return(nil)
	u.dbTx.On("Commit").Return(nil)

	err := u.userService.UpdateUser(context.Background(), userInfo, test.UserPasswdCorrect)
	require.NoError(u.T(), err)
}

func (u *userSuite) TestDeleteUserSuccess() {
	userInfo := &entity.UserInfo{
		ID:      test.UserIDCorrect,
		LoginID: test.UserLoginIDCorrect,
		Role:    test.UserRoleCorrect,
		Phone:   test.UserPhoneCorrect,
		Email:   test.UserEmailCorrect,
	}

	u.dbTx.On("Begin").Return(&u.dbTx, nil)
	u.userInfoRepo.On("WithTx", mock.Anything).Return(&u.userInfoRepo)
	u.userInfoRepo.On("Get", context.Background(), test.UserIDCorrect).Return(userInfo, nil)
	u.userInfoRepo.On("WithTx", mock.Anything).Return(&u.userInfoRepo)
	u.userInfoRepo.On("Delete", context.Background(), test.UserIDCorrect).Return(nil)
	u.userSecretRepo.On("WithTx", mock.Anything).Return(&u.userSecretRepo)
	u.userSecretRepo.On("Delete", context.Background(), mock.Anything).Return(nil)
	u.outboxRepo.On("WithTx", mock.Anything).Return(&u.outboxRepo)
	u.outboxRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
	u.dbTx.On("Commit").Return(nil)

	err := u.userService.DeleteUser(context.Background(), test.UserIDCorrect)
	require.NoError(u.T(), err)
}
