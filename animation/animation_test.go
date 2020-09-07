package animation_test

import (
	"fmt"
	"testing"

	"github.com/faiface/pixel"
	"github.com/svera/quarter/animation"
)

func TestSetCurrentAnim(t *testing.T) {
	var testValues = []struct {
		testName      string
		ID            string
		expectedError string
	}{
		{"Inexistent Animation ID throws an error", "NonExistentID", fmt.Sprintf(animation.ErrorAnimationDoesNotExist, "NonExistentID")},
	}
	for _, tt := range testValues {
		t.Run(tt.testName, func(t *testing.T) {
			an := animation.NewAnimation(pixel.V(0, 0), 5)
			err := an.SetCurrentAnim(tt.ID)
			if err.Error() != tt.expectedError {
				t.Errorf("Expected error \"%s\", got \"%s\"", tt.expectedError, err)
			}
		})
	}
}
