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

// Package ggfxinput connects ggfx to ggip: it converts the keys, mouse buttons and modifiers in
// ggfx's events to ggip's, and polls the gamepads for a ggfx app.
//
// It is a module of its own, so that ggip does not depend on ggfx, nor ggfx on ggip.
package ggfxinput

import (
	"github.com/ironpark/ggfx"
	"github.com/ironpark/ggip"
)

// keys maps a ggfx key to ggip's by name, which both take from the UI Events code values.
var keys = func() (t [ggfx.KeyMax + 1]struct {
	key ggip.Key
	ok  bool
}) {
	for k := ggfx.Key(0); k <= ggfx.KeyMax; k++ {
		t[k].key, t[k].ok = ggip.KeyFromName(k.String())
	}
	return t
}()

// Key returns the ggip key for a ggfx key. It reports false for a key ggip does not define.
func Key(k ggfx.Key) (ggip.Key, bool) {
	if k < 0 || k > ggfx.KeyMax {
		return 0, false
	}
	e := keys[k]
	return e.key, e.ok
}

// MouseButton returns the ggip mouse button for a ggfx one. Both number the buttons the same way.
func MouseButton(b ggfx.MouseButton) ggip.MouseButton {
	return ggip.MouseButton(b)
}

// Mods returns the modifiers ggfx reports with an event as ggip's.
func Mods(m ggfx.KeyModifiers) ggip.Mods {
	return ggip.Mods{Shift: m.Shift, Ctrl: m.Control, Alt: m.Alt, Meta: m.Meta}
}
