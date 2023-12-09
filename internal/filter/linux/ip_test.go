package linux

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
)

func TestCreateIfb(t *testing.T) {
	mockNetlinker := new(MockNetlinker)
	mockNetlinker.On("LinkAdd", mock.AnythingOfType("*netlink.Ifb")).Return(nil)

	err := CreateIfb(mockNetlinker)

	mockNetlinker.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestSetUpIfb(t *testing.T) {
	mockNetlinker := new(MockNetlinker)
	mockNetlinker.On("LinkByName", ifbName).Return(&netlink.Dummy{}, nil)
	mockNetlinker.On("LinkSetUp", mock.AnythingOfType("*netlink.Dummy")).Return(nil)

	err := SetUpIfb(mockNetlinker)

	mockNetlinker.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestTearDownIfb(t *testing.T) {
	mockNetlinker := new(MockNetlinker)
	mockNetlinker.On("LinkByName", ifbName).Return(&netlink.Dummy{}, nil)
	mockNetlinker.On("LinkSetDown", mock.AnythingOfType("*netlink.Dummy")).Return(nil)

	err := TearDownIfb(mockNetlinker)

	mockNetlinker.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDeleteIfb(t *testing.T) {
	mockNetlinker := new(MockNetlinker)
	mockNetlinker.On("LinkByName", ifbName).Return(&netlink.Dummy{}, nil)
	mockNetlinker.On("LinkDel", mock.AnythingOfType("*netlink.Dummy")).Return(nil)

	err := DeleteIfb(mockNetlinker)

	mockNetlinker.AssertExpectations(t)
	assert.Nil(t, err)
}
