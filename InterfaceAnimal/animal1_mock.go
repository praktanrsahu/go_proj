package InterfaceAnimal

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockAnimal struct {
	ctrl     *gomock.Controller
	recorder *MockAnimalMockRecorder
}

type MockAnimalMockRecorder struct {
	mock *MockAnimal
}

func NewMockAnimal(ctrl *gomock.Controller) *MockAnimal {
	mock := &MockAnimal{ctrl: ctrl}
	mock.recorder = &MockAnimalMockRecorder{mock}
	return mock
}

func (m *MockAnimal) EXPECT() *MockAnimalMockRecorder {
	return m.recorder
}

func (m *MockAnimal) Breathe(val string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Breathe", val)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAnimalMockRecorder) Breathe(val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Breathe", reflect.TypeOf((*MockAnimal)(nil).Breathe), val)
}