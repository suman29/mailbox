package mocks

import (
	m "git.mailbox.com/mailbox/models"
	"github.com/stretchr/testify/mock"

	_ "github.com/lib/pq"
)

// MockDB is an autogenerated mock type for the MockDB type
type MockDB struct {
	mock.Mock
}

// CreateParcel provides a mock function with given fields: dealerID, ownerID
func (_m *MockDB) CreateParcel(dealerID string, ownerID string) (*m.Parcel, error) {
	ret := _m.Called(dealerID, ownerID)

	var r0 *m.Parcel
	if rf, ok := ret.Get(0).(func(string, string) *m.Parcel); ok {
		r0 = rf(dealerID, ownerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*m.Parcel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(dealerID, ownerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDealerByID provides a mock function with given fields: id
func (_m *MockDB) GetDealerByID(id string) (*m.Dealer, error) {
	ret := _m.Called(id)

	var r0 *m.Dealer
	if rf, ok := ret.Get(0).(func(string) *m.Dealer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*m.Dealer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDealers provides a mock function with given fields:
func (_m *MockDB) GetDealers() ([]*m.Dealer, error) {
	ret := _m.Called()

	var r0 []*m.Dealer
	if rf, ok := ret.Get(0).(func() []*m.Dealer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*m.Dealer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParcelByID provides a mock function with given fields: id
func (_m *MockDB) GetParcelByID(id string) (*m.Parcel, error) {
	ret := _m.Called(id)

	var r0 *m.Parcel
	if rf, ok := ret.Get(0).(func(string) *m.Parcel); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*m.Parcel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetParcelsWith provides a mock function with given fields: searchParam
func (_m *MockDB) GetParcelsWith(searchParam string) ([]*m.Parcel, error) {
	ret := _m.Called(searchParam)

	var r0 []*m.Parcel
	if rf, ok := ret.Get(0).(func(string) []*m.Parcel); ok {
		r0 = rf(searchParam)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*m.Parcel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(searchParam)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: _a0
func (_m *MockDB) GetUserByID(_a0 string) (*m.User, error) {
	ret := _m.Called(_a0)

	var r0 *m.User
	if rf, ok := ret.Get(0).(func(string) *m.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*m.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsersWith provides a mock function with given fields: _a0
func (_m *MockDB) GetUsersWith(_a0 string) ([]*m.User, error) {
	ret := _m.Called(_a0)

	var r0 []*m.User
	if rf, ok := ret.Get(0).(func(string) []*m.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*m.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (db *MockDB) GetCloseParcels() ([]*m.ParcelUserDetails, error) {
	args := db.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*m.ParcelUserDetails), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB) GetOpenParcels() ([]*m.ParcelUserDetails, error) {
	args := db.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]*m.ParcelUserDetails), nil
	}
	return nil, args.Error(1)
}

func (db *MockDB)UpdateParcelStatusById(parcelId string, status bool) error{
	return nil
}
