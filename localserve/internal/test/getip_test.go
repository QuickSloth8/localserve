package internal

import (
	"fmt"
	"localserve/localserve/internal"
	"testing"
)

// Although this function is supposed to perform a real test
// currently, the func only prints the result for the developer
// to decide, as I didn't figure a better way yet
func TestGetIp(t *testing.T) {
	fmt.Println(internal.GetIp())
}
