package testutils

import (
	"context"

	"github.com/HotCodeGroup/warscript-utils/models"
	"google.golang.org/grpc"
)

type FakeAuthClient struct {
	Users map[int64]*models.InfoUser

	Failer
}

func (c *FakeAuthClient) GetUserByID(ctx context.Context,
	in *models.UserID, opts ...grpc.CallOption) (*models.InfoUser, error) {
	return nil, nil
}

func (c *FakeAuthClient) GetUserByUsername(ctx context.Context,
	in *models.Username, opts ...grpc.CallOption) (*models.InfoUser, error) {
	return nil, nil
}

func (c *FakeAuthClient) GetSessionInfo(ctx context.Context,
	in *models.SessionToken, opts ...grpc.CallOption) (*models.SessionPayload, error) {
	return nil, nil
}

func (c *FakeAuthClient) GetUsersByIDs(ctx context.Context,
	in *models.UserIDs, opts ...grpc.CallOption) (*models.InfoUsers, error) {
	if err := c.NextFail(); err != nil {
		return nil, err
	}

	users := &models.InfoUsers{
		Users: make([]*models.InfoUser, 0),
	}
	for _, id := range in.IDs {
		for uid, info := range c.Users {
			if id.ID == uid {
				users.Users = append(users.Users, info)
				break
			}
		}
	}

	return users, nil
}
