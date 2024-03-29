package user

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/manicar2093/health_records/internal/db/entities"
	"github.com/manicar2093/health_records/internal/db/repositories"
	"github.com/manicar2093/health_records/internal/services"
	"github.com/manicar2093/health_records/pkg/aws"
	"github.com/manicar2093/health_records/pkg/logger"
	"github.com/manicar2093/health_records/pkg/validators"
	"github.com/manicar2093/health_records/pkg/validators/nullsql"
)

var (
	emailAttributeName          string = "email"
	emailVerifiedAttributeName  string = "email_verified"
	emailVerifiedAttributeValue string = "true"
)

type UserCreator interface {
	Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error)
}

type UserCreatorImpl struct {
	authProvider aws.CongitoClient
	passGen      services.PassGen
	userRepo     repositories.UserRepository
	uuidGen      services.UUIDGenerator
	validator    validators.ValidatorService
}

func NewUserCreatorImpl(
	authProvider aws.CongitoClient,
	passGen services.PassGen,
	userRepo repositories.UserRepository,
	uuidGen services.UUIDGenerator,
	validator validators.ValidatorService,
) *UserCreatorImpl {
	return &UserCreatorImpl{
		authProvider: authProvider,
		passGen:      passGen,
		userRepo:     userRepo,
		uuidGen:      uuidGen,
		validator:    validator,
	}
}

func (c *UserCreatorImpl) Run(ctx context.Context, user *UserCreatorRequest) (*UserCreatorResponse, error) {
	logger.Info(user)
	if err := c.isValidRequest(user); err != nil {
		logger.Error(err)
		return nil, err
	}

	pass, err := c.passGen.Generate()
	if err != nil {
		logger.Error(err)
		return nil, ErrGenerationPass
	}

	requestData := cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:        &config.CognitoPoolID,
		Username:          &user.Email,
		TemporaryPassword: &pass,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{Name: &emailAttributeName, Value: &user.Email},
			{Name: &emailVerifiedAttributeName, Value: &emailVerifiedAttributeValue},
		},
	}

	userOutput, err := c.authProvider.AdminCreateUser(&requestData)
	if err != nil {
		logger.Error(err)
		return nil, ErrSavingUserAWS
	}

	userEntity := entities.User{
		Name:          user.Name,
		LastName:      user.LastName,
		RoleID:        int32(user.RoleID),
		GenderID:      nullsql.ValidateIntSQLValid(int64(user.GenderID)),
		Email:         user.Email,
		Birthday:      user.Birthday,
		AvatarUrl:     fmt.Sprintf("%s%s.svg", config.AvatarURLSrc, c.uuidGen.New()),
		BiotypeID:     nullsql.ValidateIntSQLValid(int64(user.BiotypeID)),
		BoneDensityID: nullsql.ValidateIntSQLValid(int64(user.BoneDensityID)),
	}

	userEntity.UserUUID = *userOutput.User.Username

	err = c.userRepo.SaveUser(ctx, &userEntity)

	if err != nil {
		logger.Error(err)
		return nil, ErrSavingUserDB
	}

	return &UserCreatorResponse{UserCreated: &userEntity}, nil
}

func (c *UserCreatorImpl) isValidRequest(createUserReq *UserCreatorRequest) error {
	if validation := c.validator.Validate(createUserReq); !validation.IsValid {
		return validation.Err
	}

	if createUserReq.RoleID == CustomerRole {
		if validation := c.validator.Validate(
			createUserReq.GetCustomerCreationValidations(),
		); !validation.IsValid {
			return validation.Err
		}
	}

	return nil

}
