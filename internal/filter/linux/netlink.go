package linux

import (
	"github.com/stretchr/testify/mock"
	"github.com/vishvananda/netlink"
)

type Netlinker interface {
	LinkAdd(link netlink.Link) error
	LinkDel(link netlink.Link) error
	LinkSetUp(link netlink.Link) error
	LinkSetDown(link netlink.Link) error
	LinkByName(name string) (netlink.Link, error)
}

type RealNetlinker struct{}

func (r *RealNetlinker) LinkAdd(link netlink.Link) error {
	return netlink.LinkAdd(link)
}

func (r *RealNetlinker) LinkDel(link netlink.Link) error {
	return netlink.LinkDel(link)
}

func (r *RealNetlinker) LinkSetUp(link netlink.Link) error {
	return netlink.LinkSetUp(link)
}

func (r *RealNetlinker) LinkSetDown(link netlink.Link) error {
	return netlink.LinkSetDown(link)
}

func (r *RealNetlinker) LinkByName(name string) (netlink.Link, error) {
	return netlink.LinkByName(name)
}

type MockNetlinker struct {
	mock.Mock
}

func (m *MockNetlinker) LinkAdd(link netlink.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockNetlinker) LinkDel(link netlink.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockNetlinker) LinkSetUp(link netlink.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockNetlinker) LinkSetDown(link netlink.Link) error {
	args := m.Called(link)
	return args.Error(0)
}

func (m *MockNetlinker) LinkByName(name string) (netlink.Link, error) {
	args := m.Called(name)
	return args.Get(0).(netlink.Link), args.Error(1)
}
