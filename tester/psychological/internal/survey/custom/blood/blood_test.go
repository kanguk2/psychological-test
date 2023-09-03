package blood

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBloodType(t *testing.T) {
	bloodType := GetBloodType([]string{"A", "B"})
	assert.Equal(t, AB, bloodType)

	bloodType = GetBloodType([]string{"B", "A"})
	assert.Equal(t, AB, bloodType)

	bloodType = GetBloodType([]string{"A", "O"})
	assert.Equal(t, A, bloodType)

	bloodType = GetBloodType([]string{"O", "A"})
	assert.Equal(t, A, bloodType)

	bloodType = GetBloodType([]string{"O", "O"})
	assert.Equal(t, O, bloodType)

	bloodType = GetBloodType([]string{"G", "A"})
	assert.Equal(t, "", bloodType)
}
