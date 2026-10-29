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

// Package ggfxinput connects ggwd, ggfx's window system, to ggip: it converts the keys, mouse
// buttons and modifiers in ggwd's events to ggip's, and polls the gamepads for a ggwd app.
//
// It is a module of its own, so that ggip does not depend on ggfx, nor ggfx on ggip.
package ggfxinput

import (
	"github.com/ironpark/ggwd"
	"github.com/ironpark/ggip"
)

// keys maps a ggwd key to ggip's by name, which both take from the UI Events code values.
var keys = func() (t [ggwd.KeyMax + 1]struct {
	key ggip.Key
	ok  bool
}) {
	for k := ggwd.Key(0); k <= ggwd.KeyMax; k++ {
		t[k].key, t[k].ok = ggip.KeyFromName(k.String())
	}
	return t
}()

// Key returns the ggip key for a ggwd key. It reports false for a key ggip does not define.
func Key(k ggwd.Key) (ggip.Key, bool) {
	if k < 0 || k > ggwd.KeyMax {
		return 0, false
	}
	e := keys[k]
	return e.key, e.ok
}

// MouseButton returns the ggip mouse button for a ggwd one. Both number the buttons the same way.
func MouseButton(b ggwd.MouseButton) ggip.MouseButton {
	return ggip.MouseButton(b)
}

// Mods returns the modifiers ggwd reports with an event as ggip's.
func Mods(m ggwd.KeyModifiers) ggip.Mods {
	return ggip.Mods{Shift: m.Shift, Ctrl: m.Control, Alt: m.Alt, Meta: m.Meta}
}
