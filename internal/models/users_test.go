package models

import (
	"testing"

	"snippetbox.divakaivan.net/internal/assert"
)

func TestUserModelExists(t *testing.T) {

	// will skip if -short flag is passed
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{name: "Valid ID", userID: 1, want: true}, // true
		{name: "Zero ID", userID: 0, want: false},
		{name: "Non-existent ID", userID: 2, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := UserModel{db}

			got, err := m.Exists(tt.userID)

			assert.Equal(t, got, tt.want)
			assert.NilError(t, err)
		})
	}
}
