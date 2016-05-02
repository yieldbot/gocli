/*
 * gocli
 * Copyright (c) 2015 Yieldbot, Inc.
 * For the full copyright and license information, please view the LICENSE.txt file.
 */

package gocli

import (
	"testing"
)

func TestStrPadRight(t *testing.T) {
	if strPadRight("hello", " ", 10) != "hello     " {
		t.Error("invalid return by strPadRight")
	}
}
