package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	m "git.mailbox.com/mailbox/models"
	tu "git.mailbox.com/mailbox/testutils"
	u "git.mailbox.com/mailbox/utils"
)

var dealerParcelSearch = &m.Dealer{
	ID:   u.SPtr("bda1103c-4024-4ea4-b955-58c1c2c702b7"),
	Name: u.SPtr("Flipkart"),
}

var ownerParcelSearch = &m.User{
	ID:      u.SPtr("cda1103c-4024-4ea4-b955-58c1c2c702b7"),
	Email:   strfmt.Email("mello@mello.com"),
	Name:    u.SPtr("Mello"),
	EmpID:   "11113",
	PhoneNo: u.SPtr("9910399900"),
}

var receiverParcelSearch = &m.User{
	ID:      u.SPtr("dda1103c-4024-4ea4-b955-58c1c2c702b7"),
	Email:   strfmt.Email("jello@jello.com"),
	Name:    u.SPtr("Jello"),
	EmpID:   "11114",
	PhoneNo: u.SPtr("9900299900"),
}

var parcels = []*m.Parcel{
	&m.Parcel{
		ID:         "ada1103c-4024-4ea4-b955-58c1c2c702b7",
		DealerID:   u.SPtr("bda1103c-4024-4ea4-b955-58c1c2c702b7"),
		OwnerID:    u.SPtr("cda1103c-4024-4ea4-b955-58c1c2c702b7"),
		RecieverID: u.SPtr("dda1103c-4024-4ea4-b955-58c1c2c702b7"),
		Owner:      ownerParcelSearch,
		Dealer:     dealerParcelSearch,
		Reciever:   receiverParcelSearch,
	},
}

func TestParcelSearchSuccess(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=mello", nil)
	require.NoError(t, err, "failed to create a request: dealers")

	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)
	mockDbObj.On("GetParcelsWith", "mello").Return(parcels, nil)
	mockDbObj.On("GetDealerByID", "bda1103c-4024-4ea4-b955-58c1c2c702b7").Return(dealerParcelSearch, nil)
	mockDbObj.On("GetUserByID", "cda1103c-4024-4ea4-b955-58c1c2c702b7").Return(ownerParcelSearch, nil)
	mockDbObj.On("GetUserByID", "dda1103c-4024-4ea4-b955-58c1c2c702b7").Return(receiverParcelSearch, nil)

	parcelSearchHandler(mockDbObj)(w, r)

	var actualParcels []*m.Parcel
	err = json.Unmarshal(w.Body.Bytes(), &actualParcels)
	require.NoError(t, err, "failed to unmarshal the response")

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Equal(t, parcels[0].ID, actualParcels[0].ID)
	assert.Equal(t, parcels[0].Owner, actualParcels[0].Owner)
	assert.Equal(t, parcels[0].Dealer, actualParcels[0].Dealer)
	assert.Equal(t, parcels[0].Reciever, actualParcels[0].Reciever)

	mockDbObj.AssertExpectations(t)
}

func TestParcelSearchSuccessNoQueryParam(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)

	parcelSearchHandler(mockDbObj)(w, r)

	var actualParcels []*m.Parcel
	err = json.Unmarshal(w.Body.Bytes(), &actualParcels)
	require.NoError(t, err, "failed to unmarshal the response")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 0, len(actualParcels))
}

func TestParcelSearchSuccessWhenLessThan3QueryParam(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=li", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)

	parcelSearchHandler(mockDbObj)(w, r)

	var actualParcels []*m.User
	err = json.Unmarshal(w.Body.Bytes(), &actualParcels)
	require.NoError(t, err, "failed to unmarshal the response")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 0, len(actualParcels))
}

func TestParcelSearchDBFailure(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=mello", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)
	mockDbObj.On("GetParcelsWith", "mello").Return(nil, fmt.Errorf("failure to connect to db"))

	parcelSearchHandler(mockDbObj)(w, r)

	var actualError *m.Error
	err = json.Unmarshal(w.Body.Bytes(), &actualError)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	require.NotNil(t, actualError)

	assert.Equal(t, int32(500), *actualError.Code)
	assert.Equal(t, "Internal server error", *actualError.Message)

	require.Nil(t, actualError.Fields)

	mockDbObj.AssertExpectations(t)
}

func TestParcelSearchDealerGetFailure(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=mello", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)
	mockDbObj.On("GetParcelsWith", "mello").Return(parcels, nil)
	mockDbObj.On("GetDealerByID", "bda1103c-4024-4ea4-b955-58c1c2c702b7").Return(nil, fmt.Errorf("failed to get the owner"))

	parcelSearchHandler(mockDbObj)(w, r)

	var actualError *m.Error
	err = json.Unmarshal(w.Body.Bytes(), &actualError)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	require.NotNil(t, actualError)

	assert.Equal(t, int32(500), *actualError.Code)
	assert.Equal(t, "Internal server error", *actualError.Message)

	require.Nil(t, actualError.Fields)

	mockDbObj.AssertExpectations(t)
}

func TestParcelSearchOwnerGetFailure(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=mello", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)
	mockDbObj.On("GetParcelsWith", "mello").Return(parcels, nil)
	mockDbObj.On("GetDealerByID", "bda1103c-4024-4ea4-b955-58c1c2c702b7").Return(dealerParcelSearch, nil)
	mockDbObj.On("GetUserByID", "cda1103c-4024-4ea4-b955-58c1c2c702b7").Return(nil, fmt.Errorf("failed to get the owner"))

	parcelSearchHandler(mockDbObj)(w, r)

	var actualError *m.Error
	err = json.Unmarshal(w.Body.Bytes(), &actualError)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	require.NotNil(t, actualError)

	assert.Equal(t, int32(500), *actualError.Code)
	assert.Equal(t, "Internal server error", *actualError.Message)

	require.Nil(t, actualError.Fields)

	mockDbObj.AssertExpectations(t)
}

func TestParcelSearchReceiverGetFailure(t *testing.T) {
	r, err := http.NewRequest("GET", "/parcels/search?q=mello", nil)
	require.NoError(t, err, "failed to create a request: dealers")
	w := httptest.NewRecorder()

	mockDbObj := new(tu.MockDB)
	mockDbObj.On("GetParcelsWith", "mello").Return(parcels, nil)
	mockDbObj.On("GetDealerByID", "bda1103c-4024-4ea4-b955-58c1c2c702b7").Return(dealerParcelSearch, nil)
	mockDbObj.On("GetUserByID", "cda1103c-4024-4ea4-b955-58c1c2c702b7").Return(ownerParcelSearch, nil)
	mockDbObj.On("GetUserByID", "dda1103c-4024-4ea4-b955-58c1c2c702b7").Return(nil, fmt.Errorf("failed to get the owner"))

	parcelSearchHandler(mockDbObj)(w, r)

	var actualError *m.Error
	err = json.Unmarshal(w.Body.Bytes(), &actualError)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	require.NotNil(t, actualError)

	assert.Equal(t, int32(500), *actualError.Code)
	assert.Equal(t, "Internal server error", *actualError.Message)

	require.Nil(t, actualError.Fields)

	mockDbObj.AssertExpectations(t)
}
