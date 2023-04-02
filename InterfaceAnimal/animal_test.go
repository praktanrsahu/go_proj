package InterfaceAnimal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCat_Breath(t *testing.T) {
	tests := []struct {
		name  string
		sound string
		val   string
		want  string
		err   string
	}{
		{
			name: "Success",
			val:  "15",
			want: "I breath 15 times per minute",
			err:  "",
		},
		{
			name: "Failure",
			val:  "25",
			want: "",
			err:  "Breath provided- 25 , Please provide appropriate. ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Cat{
				sound: test.sound,
			}
			got, err := c.Breath(test.val)
			assert.Equal(t, test.want, got)
			if err != nil || test.err != "" {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestCat_Sound(t *testing.T) {
	tests := []struct {
		name    string
		sound   string
		want    string
		wantErr string
	}{
		{
			name:    "Failure",
			sound:   "Bhow",
			want:    "",
			wantErr: "sound provided for Cat: Bhow , Please provide appropriate. ",
		},
		{
			name:    "Success",
			sound:   "Meow!!",
			want:    "My sound is Meow!!",
			wantErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := Cat{
				sound: test.sound,
			}
			got, err := c.Sound()
			assert.Equal(t, test.want, got)
			if err != nil || test.wantErr != "" {
				assert.EqualError(t, err, test.wantErr)
			}
		})
	}
}

func TestDog_Breath(t *testing.T) {
	tests := []struct {
		name  string
		sound string
		val   string
		want  string
		err   string
	}{
		{
			name: "Success",
			val:  "25",
			want: "I breath 25 times per minute",
			err:  "",
		},
		{
			name: "Failure",
			val:  "15",
			want: "",
			err:  "Breath provided- 15 , Please provide appropriate. ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := Dog{
				sound: test.sound,
			}
			got, err := d.Breath(test.val)
			assert.Equal(t, test.want, got)
			if err != nil || test.err != "" {
				assert.EqualError(t, err, test.err)
			}
		})
	}
}

func TestDog_Sound(t *testing.T) {
	tests := []struct {
		name    string
		sound   string
		want    string
		wantErr string
	}{
		{
			name:    "Failure",
			sound:   "Meow",
			want:    "",
			wantErr: "sound provided for Dog: Meow , Please provide appropriate. ",
		},
		{
			name:    "Success",
			sound:   "Bhow!!",
			want:    "My sound is Bhow!!",
			wantErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			d := Dog{
				sound: test.sound,
			}
			got, err := d.Sound()
			assert.Equal(t, test.want, got)
			if err != nil || test.wantErr != "" {
				assert.EqualError(t, err, test.wantErr)
			}
		})
	}
}
