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
		ID            int
		expectedError string
	}{
		{"Animation ID below 0 does not exist", -1, fmt.Sprintf(animation.ErrorAnimationDoesNotExist, -1)},
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
