package model_test

import (
	"errors"
	"testing"
	"time"

	"github.com/codefluence-x/altair/provider/plugin/oauth/entity"
	"github.com/codefluence-x/altair/provider/plugin/oauth/model"
	"github.com/codefluence-x/altair/provider/plugin/oauth/query"
	mockdb "github.com/codefluence-x/monorepo/db/mock"
	"github.com/codefluence-x/monorepo/exception"
	"github.com/codefluence-x/monorepo/kontext"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOauthAccessToken(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("one", func(t *testing.T) {
		t.Run("Given oauth access token", func(t *testing.T) {
			t.Run("When database operation complete it will return oauth access token data", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				row := mockdb.NewMockRow(mockCtrl)

				expectedData := entity.OauthAccessToken{
					ID:    1,
					Token: "token",
				}

				sqldb.EXPECT().QueryRowContext(gomock.Any(), "oauth-access-token-one", query.SelectOneOauthAccessToken, expectedData.ID).Return(row)
				row.EXPECT().Scan(gomock.Any()).DoAndReturn(func(dest ...interface{}) exception.Exception {
					val0, _ := dest[0].(*int)
					*val0 = expectedData.ID

					val1, _ := dest[3].(*string)
					*val1 = expectedData.Token
					return nil
				})

				oauthAccessTokenModel := model.NewOauthAccessToken()
				data, err := oauthAccessTokenModel.One(kontext.Fabricate(), expectedData.ID, sqldb)

				assert.Nil(t, err)
				assert.Equal(t, expectedData, data)
			})

		})
	})

	t.Run("OneByToken", func(t *testing.T) {
		t.Run("Given oauth access token", func(t *testing.T) {
			t.Run("When database operation complete it will return oauth access token data", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				row := mockdb.NewMockRow(mockCtrl)

				expectedData := entity.OauthAccessToken{
					ID:    1,
					Token: "token",
				}

				sqldb.EXPECT().QueryRowContext(gomock.Any(), "oauth-access-token-one-by-token", query.SelectOneOauthAccessTokenByToken, expectedData.Token).Return(row)
				row.EXPECT().Scan(gomock.Any()).DoAndReturn(func(dest ...interface{}) exception.Exception {
					val0, _ := dest[0].(*int)
					*val0 = expectedData.ID

					val1, _ := dest[3].(*string)
					*val1 = expectedData.Token
					return nil
				})

				oauthAccessTokenModel := model.NewOauthAccessToken()
				data, err := oauthAccessTokenModel.OneByToken(kontext.Fabricate(), expectedData.Token, sqldb)

				assert.Nil(t, err)
				assert.Equal(t, expectedData, data)
			})

		})
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("Given context and access token insertable", func(t *testing.T) {
			t.Run("When database operation complete it will create oauth access token data", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				result := mockdb.NewMockResult(mockCtrl)

				expectedID := 1
				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-create", query.InsertOauthAccessToken,
					insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn,
				).Return(result, nil)

				result.EXPECT().LastInsertId().Return(int64(expectedID), nil)

				oauthAccessTokenModel := model.NewOauthAccessToken()
				ID, err := oauthAccessTokenModel.Create(kontext.Fabricate(), insertable, sqldb)

				assert.Equal(t, expectedID, ID)
				assert.Nil(t, err)
			})

			t.Run("When database operation failed then it will return error", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-create", query.InsertOauthAccessToken,
					insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn,
				).Return(nil, exception.Throw(errors.New("unexpected")))

				oauthAccessTokenModel := model.NewOauthAccessToken()
				_, err := oauthAccessTokenModel.Create(kontext.Fabricate(), insertable, sqldb)

				assert.Equal(t, exception.Unexpected, err.Type())
				assert.NotNil(t, err)
			})

			t.Run("When database operation is complete but get last inserted id error then it will return error", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				result := mockdb.NewMockResult(mockCtrl)

				insertable := entity.OauthAccessTokenInsertable{
					OauthApplicationID: 1,
					ResourceOwnerID:    1,
					Token:              "token",
					Scopes:             "public users stores",
					ExpiresIn:          time.Now().Add(time.Hour * 24),
				}

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-create", query.InsertOauthAccessToken,
					insertable.OauthApplicationID, insertable.ResourceOwnerID, insertable.Token, insertable.Scopes, insertable.ExpiresIn,
				).Return(result, nil)

				result.EXPECT().LastInsertId().Return(int64(0), exception.Throw(errors.New("unexpected error")))

				oauthAccessTokenModel := model.NewOauthAccessToken()
				ID, err := oauthAccessTokenModel.Create(kontext.Fabricate(), insertable, sqldb)

				assert.Equal(t, 0, ID)
				assert.NotNil(t, err)
				assert.Equal(t, exception.Unexpected, err.Type())
			})
		})
	})

	t.Run("Revoke", func(t *testing.T) {
		t.Run("Given context and token", func(t *testing.T) {
			t.Run("When database operation complete it will return nil", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				result := mockdb.NewMockResult(mockCtrl)

				token := "token"

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-revoke", query.RevokeAccessToken, token).Return(result, nil)
				result.EXPECT().RowsAffected().Return(int64(1), nil)

				oauthAccessTokenModel := model.NewOauthAccessToken()
				err := oauthAccessTokenModel.Revoke(kontext.Fabricate(), token, sqldb)
				assert.Nil(t, err)
			})

			t.Run("When database operation complete but there is no updated rows it will return error", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				result := mockdb.NewMockResult(mockCtrl)

				token := "token"

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-revoke", query.RevokeAccessToken, token).Return(result, nil)
				result.EXPECT().RowsAffected().Return(int64(0), nil)

				oauthAccessTokenModel := model.NewOauthAccessToken()
				err := oauthAccessTokenModel.Revoke(kontext.Fabricate(), token, sqldb)
				assert.NotNil(t, err)
				assert.Equal(t, exception.NotFound, err.Type())
			})

			t.Run("When database operation failed it will return error", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				token := "token"

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-revoke", query.RevokeAccessToken, token).Return(nil, exception.Throw(errors.New("unexpected error")))

				oauthAccessTokenModel := model.NewOauthAccessToken()
				err := oauthAccessTokenModel.Revoke(kontext.Fabricate(), token, sqldb)
				assert.NotNil(t, err)
			})

			t.Run("When database operation complete but get rows affected error it will return error", func(t *testing.T) {
				sqldb := mockdb.NewMockDB(mockCtrl)
				result := mockdb.NewMockResult(mockCtrl)

				token := "token"

				sqldb.EXPECT().ExecContext(gomock.Any(), "oauth-access-token-revoke", query.RevokeAccessToken, token).Return(result, nil)
				result.EXPECT().RowsAffected().Return(int64(0), exception.Throw(errors.New("unexpected error")))

				oauthAccessTokenModel := model.NewOauthAccessToken()
				err := oauthAccessTokenModel.Revoke(kontext.Fabricate(), token, sqldb)
				assert.NotNil(t, err)
			})
		})
	})
}
