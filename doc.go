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

// Package ggip handles keyboard and mouse input once a window system has reported it: the keys,
// named after the UI Events code values, the modifier keys, the mouse buttons, which keys are
// held, and shortcuts written as chords such as "cmd+s".
//
// ggip does not read the keyboard or the mouse; a window system does, and the app feeds what it
// reports to ggip. Gamepads are different, as they belong to the process rather than to a
// window: the gamepad package reads them.
package ggip

//go:generate go run genkeys.go
//go:generate gofmt -s -w .
