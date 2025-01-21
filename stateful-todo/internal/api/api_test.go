package api

import (
	"fmt"
	"testing"
)

func TestNewServer(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		t.Parallel()
		fmt.Println("test succeed")
	})
}
