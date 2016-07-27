package testutils

import (
	"github.com/stretchr/testify/mock"

	m "git.mailbox.com/mailbox/models"
)

type MockDB struct {
	mock.Mock
}

func (db *MockDB) GetDealers() ([]*m.Dealer, error) {
	args := db.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*m.Dealer), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetUsersWith(searchParam string) ([]*m.User, error) {
	args := db.Called(searchParam)
	if args.Get(0) != nil {
		return args.Get(0).([]*m.User), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetUserByID(id string) (*m.User, error) {
	args := db.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*m.User), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetDealerByID(id string) (*m.Dealer, error) {
	args := db.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*m.Dealer), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) CreateParcel(dealerID, userID string) (*m.Parcel, error) {
	args := db.Called(dealerID, userID)
	if args.Get(0) != nil {
		return args.Get(0).(*m.Parcel), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetParcelByID(id string) (*m.Parcel, error) {
	args := db.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*m.Parcel), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetCloseParcels() ([]*m.ParcelWithUserAndDealer, error) {
	args := db.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*m.ParcelWithUserAndDealer), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetOpenParcels() ([]*m.ParcelWithUserAndDealer, error) {
	args := db.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*m.ParcelWithUserAndDealer), nil
	}
	return nil, args.Error(1)
}
