// Copyright 2026 The ggip Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ggfxinput_test

import (
	"testing"

	"github.com/ironpark/ggwd"
	"github.com/ironpark/ggip"
	"github.com/ironpark/ggip/ggfxinput"
)

// Both define their keys from the same code values. A key that one adds and the other does not
// shows up here.
func TestEveryKeyMaps(t *testing.T) {
	seen := map[ggip.Key]bool{}
	for k := ggwd.Key(0); k <= ggwd.KeyMax; k++ {
		got, ok := ggfxinput.Key(k)
		if !ok {
			t.Errorf("ggfx key %v has no ggip key", k)
			continue
		}
		if got.String() != k.String() {
			t.Errorf("ggfx key %v maps to ggip key %v", k, got)
		}
		seen[got] = true
	}
	for k := ggip.Key(0); k <= ggip.KeyMax; k++ {
		if !seen[k] {
			t.Errorf("no ggfx key maps to ggip key %v", k)
		}
	}
	if _, ok := ggfxinput.Key(ggwd.KeyMax + 1); ok {
		t.Error("an undefined ggfx key maps")
	}
}

func TestMouseButtonsMap(t *testing.T) {
	pairs := map[ggwd.MouseButton]ggip.MouseButton{
		ggwd.MouseButtonLeft:   ggip.MouseButtonLeft,
		ggwd.MouseButtonMiddle: ggip.MouseButtonMiddle,
		ggwd.MouseButtonRight:  ggip.MouseButtonRight,
		ggwd.MouseButton3:      ggip.MouseButton3,
		ggwd.MouseButton4:      ggip.MouseButton4,
	}
	for from, want := range pairs {
		if got := ggfxinput.MouseButton(from); got != want {
			t.Errorf("MouseButton(%d) = %d; want %d", from, got, want)
		}
	}
	if int(ggwd.MouseButtonMax) != int(ggip.MouseButtonMax) {
		t.Errorf("ggfx has %d mouse buttons and ggip %d", ggwd.MouseButtonMax+1, ggip.MouseButtonMax+1)
	}
}

func TestModsMap(t *testing.T) {
	got := ggfxinput.Mods(ggwd.KeyModifiers{Shift: true, Control: true, Meta: true})
	if want := (ggip.Mods{Shift: true, Ctrl: true, Meta: true}); got != want {
		t.Errorf("Mods = %+v; want %+v", got, want)
	}
}
