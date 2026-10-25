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

package ggip_test

import (
	"testing"

	"github.com/ironpark/ggip"
)

func TestKeyboardModsFromEitherSide(t *testing.T) {
	var kb ggip.Keyboard
	kb.Set(ggip.KeyShiftRight, true)
	kb.Set(ggip.KeyControlLeft, true)
	kb.Set(ggip.KeyMeta, true)
	if got, want := kb.Mods(), (ggip.Mods{Shift: true, Ctrl: true, Meta: true}); got != want {
		t.Errorf("Mods() = %+v; want %+v", got, want)
	}
	kb.Set(ggip.KeyShiftRight, false)
	if kb.Mods().Shift {
		t.Error("Shift is held after its only key was released")
	}
	if !kb.IsPressed(ggip.KeyControlLeft) || kb.IsPressed(ggip.KeyControlRight) {
		t.Error("IsPressed does not tell the two Control keys apart")
	}
	kb.Reset()
	if kb.Mods() != (ggip.Mods{}) || kb.IsPressed(ggip.KeyControlLeft) {
		t.Error("Reset left a key held")
	}
}

func TestKeyboardIgnoresUndefinedKeys(t *testing.T) {
	var kb ggip.Keyboard
	kb.Set(ggip.KeyMax+1, true)
	kb.Set(-1, true)
	if kb.IsPressed(ggip.KeyMax+1) || kb.IsPressed(-1) {
		t.Error("an undefined key reads as held")
	}
}

func TestModsOr(t *testing.T) {
	got := ggip.Mods{Shift: true}.Or(ggip.Mods{Alt: true})
	if want := (ggip.Mods{Shift: true, Alt: true}); got != want {
		t.Errorf("Or = %+v; want %+v", got, want)
	}
}
