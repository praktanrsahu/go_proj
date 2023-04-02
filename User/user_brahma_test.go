package User

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"sample/mocktest/InterfaceAnimal"
	"testing"
)

func TestBrahma_Use(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		testUser func(mock *InterfaceAnimal.MockAnimal)
		want     string
		err      string
		val      string
	}{
		{
			name: "Success",
			testUser: func(mock *InterfaceAnimal.MockAnimal) {
				mock.EXPECT().Breathe(gomock.Any()).Return("15 Breath per min", nil)
			},
			want: "15 Breath per min",
			err:  "",
		},
		{
			name: "Failure",
			testUser: func(mock *InterfaceAnimal.MockAnimal) {
				mock.EXPECT().Breathe(gomock.Any()).Return("", fmt.Errorf("Error while executing method "))
			},
			want: "",
			err:  "Error while executing method ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockUser := InterfaceAnimal.NewMockAnimal(ctrl)
			testuser := Brahma{typebrahma: mockUser}
			if test.testUser != nil {
				test.testUser(mockUser)
			}
			got, err := testuser.Use(test.val)
			assert.Equal(t, test.want, got)
			if err != nil || test.err != "" {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}
