// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package types

import (
	"github.com/maraino/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorageNS

type MockStateStorageNS struct {
	mock.Mock
}

func (s *MockStateStorageNS) WriteKeyNS(input *WriteKeyInputNS) (*WriteKeyOutputNS, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*WriteKeyOutputNS), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockStateStorageNS) ReadKeyNS(input *ReadKeyInputNS) (*ReadKeyOutputNS, error) {
	ret := s.Called(input)
	if out := ret.Get(0); out != nil {
		return out.(*ReadKeyOutputNS), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
