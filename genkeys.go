// Copyright 2015 Hajime Hoshi
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

//go:build ignore

// The key name convention follows the Web standard: https://www.w3.org/TR/uievents-code/#code-value-tables

package main

import (
	"bufio"
	"cmp"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

var keyNameToCode map[string]string

func init() {
	// A key's name is its code without the "Key" prefix of the letters (e.g., 'A' vs 'KeyA').
	keyNameToCode = map[string]string{
		"Comma":          "Comma",
		"Period":         "Period",
		"AltLeft":        "AltLeft",
		"AltRight":       "AltRight",
		"CapsLock":       "CapsLock",
		"ControlLeft":    "ControlLeft",
		"ControlRight":   "ControlRight",
		"ShiftLeft":      "ShiftLeft",
		"ShiftRight":     "ShiftRight",
		"Enter":          "Enter",
		"Space":          "Space",
		"Tab":            "Tab",
		"Delete":         "Delete",
		"End":            "End",
		"Home":           "Home",
		"Insert":         "Insert",
		"PageDown":       "PageDown",
		"PageUp":         "PageUp",
		"ArrowDown":      "ArrowDown",
		"ArrowLeft":      "ArrowLeft",
		"ArrowRight":     "ArrowRight",
		"ArrowUp":        "ArrowUp",
		"Escape":         "Escape",
		"Backspace":      "Backspace",
		"Quote":          "Quote",
		"Minus":          "Minus",
		"Slash":          "Slash",
		"Semicolon":      "Semicolon",
		"Equal":          "Equal",
		"BracketLeft":    "BracketLeft",
		"Backslash":      "Backslash",
		"BracketRight":   "BracketRight",
		"Backquote":      "Backquote",
		"NumLock":        "NumLock",
		"Pause":          "Pause",
		"PrintScreen":    "PrintScreen",
		"ScrollLock":     "ScrollLock",
		"ContextMenu":    "ContextMenu",
		"NumpadAdd":      "NumpadAdd",
		"NumpadDecimal":  "NumpadDecimal",
		"NumpadDivide":   "NumpadDivide",
		"NumpadMultiply": "NumpadMultiply",
		"NumpadSubtract": "NumpadSubtract",
		"NumpadEnter":    "NumpadEnter",
		"NumpadEqual":    "NumpadEqual",
		"MetaLeft":       "MetaLeft",
		"MetaRight":      "MetaRight",
		"IntlBackslash":  "IntlBackslash",
	}

	// ASCII: 0 - 9
	for c := '0'; c <= '9'; c++ {
		name := "Digit" + string(c)
		keyNameToCode[name] = name
	}
	// ASCII: A - Z
	for c := 'A'; c <= 'Z'; c++ {
		keyNameToCode[string(c)] = "Key" + string(c)
	}
	// Function keys
	for i := 1; i <= 24; i++ {
		name := "F" + strconv.Itoa(i)
		keyNameToCode[name] = name
	}
	// Numpad
	// https://www.w3.org/TR/uievents-code/#key-numpad-section
	for c := '0'; c <= '9'; c++ {
		name := "Numpad" + string(c)
		keyNameToCode[name] = name
	}
}

const keysTmpl = `{{.License}}

{{.DoNotEdit}}

package ggip

import (
	"fmt"
	"strings"
)

// A Key represents a keyboard key.
// These keys represent physical keys of US keyboard.
// For example, KeyQ represents Q key on US keyboards and ' (quote) key on Dvorak keyboards.
type Key int

// Keys.
const (
{{range $index, $name := .KeyNamesWithoutMods}}Key{{$name}}{{if eq $index 0}} Key = iota{{end}}
{{end}}	KeyAlt
	KeyControl
	KeyShift
	KeyMeta
	KeyMax = KeyMeta
)

// String returns a string representing the key.
//
// If k is an undefined key, String returns an empty string.
func (k Key) String() string {
	switch k {
	{{range $name := .KeyNames}}case Key{{$name}}:
		return {{$name | printf "%q"}}
	{{end}}}
	return ""
}

func keyNameToKeyCode(name string) (Key, bool) {
	switch strings.ToLower(name) {
	{{range $name := .KeyNames}}case {{$name | printf "%q" | ToLower}}:
		return Key{{$name}}, true
	{{end}}}
	return 0, false
}

// Code returns the key's code value in the UI Events KeyboardEvent code values, the code of a
// browser's KeyboardEvent: https://www.w3.org/TR/uievents-code/.
//
// Code returns an empty string for KeyAlt, KeyControl, KeyShift and KeyMeta, which stand for
// either side's key, and for an undefined key.
func (k Key) Code() string {
	switch k {
	{{range $name, $code := .KeyNameToCode}}case Key{{$name}}:
		return {{$code | printf "%q"}}
	{{end}}}
	return ""
}

// KeyFromCode returns the key whose code value is code. See [Key.Code].
func KeyFromCode(code string) (Key, bool) {
	switch code {
	{{range $name, $code := .KeyNameToCode}}case {{$code | printf "%q"}}:
		return Key{{$name}}, true
	{{end}}}
	return 0, false
}

// KeyFromName returns the key whose String is name, ignoring case.
func KeyFromName(name string) (Key, bool) {
	return keyNameToKeyCode(name)
}

// MarshalText implements encoding.TextMarshaler.
func (k Key) MarshalText() ([]byte, error) {
	return []byte(k.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler
func (k *Key) UnmarshalText(text []byte) error {
	key, ok := keyNameToKeyCode(string(text))
	if !ok {
		return fmt.Errorf("ggip: unexpected key name: %s", string(text))
	}
	*k = key
	return nil
}
`

func digitKey(name string) int {
	if len(name) != 1 {
		return -1
	}
	c := name[0]
	if c < '0' || '9' < c {
		return -1
	}
	return int(c - '0')
}

func alphabetKey(name string) rune {
	if len(name) != 1 {
		return -1
	}
	c := rune(name[0])
	if c < 'A' || 'Z' < c {
		return -1
	}
	return c
}

func functionKey(name string) int {
	if len(name) < 2 {
		return -1
	}
	if name[0] != 'F' {
		return -1
	}
	i, err := strconv.Atoi(name[1:])
	if err != nil {
		return -1
	}
	return i
}

func keyNamesCmp(k0, k1 string) int {
	d0, d1 := digitKey(k0), digitKey(k1)
	a0, a1 := alphabetKey(k0), alphabetKey(k1)
	f0, f1 := functionKey(k0), functionKey(k1)
	if d0 != -1 {
		if d1 != -1 {
			return cmp.Compare(d0, d1)
		}
		return -1
	}
	if a0 != -1 {
		if d1 != -1 {
			return 1
		}
		if a1 != -1 {
			return cmp.Compare(a0, a1)
		}
		return -1
	}
	if d1 != -1 {
		return 1
	}
	if a1 != -1 {
		return 1
	}
	if f0 != -1 && f1 != -1 {
		return cmp.Compare(f0, f1)
	}
	return cmp.Compare(k0, k1)
}

const license = `// Copyright 2013 The Ebitengine Authors
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
`

func main() {
	// Follow the standard comment rule (https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source).
	doNotEdit := "// Code generated by genkeys.go using 'go generate'. DO NOT EDIT."

	var keyNames, keyNamesWithoutMods []string
	for name := range keyNameToCode {
		keyNames = append(keyNames, name)
		keyNamesWithoutMods = append(keyNamesWithoutMods, name)
	}
	// Keys for modifiers
	keyNames = append(keyNames, "Alt", "Control", "Shift", "Meta")

	slices.SortFunc(keyNames, keyNamesCmp)
	slices.SortFunc(keyNamesWithoutMods, keyNamesCmp)

	f, err := os.Create("keys.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	funcs := template.FuncMap{
		"ToLower": strings.ToLower,
	}
	tmpl, err := template.New("keys.go").Funcs(funcs).Parse(keysTmpl)
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: According to godoc, maps are automatically sorted by key.
	w := bufio.NewWriter(f)
	if err := tmpl.Execute(w, struct {
		License             string
		DoNotEdit           string
		KeyNameToCode       map[string]string
		KeyNames            []string
		KeyNamesWithoutMods []string
	}{
		License:             license,
		DoNotEdit:           doNotEdit,
		KeyNameToCode:       keyNameToCode,
		KeyNames:            keyNames,
		KeyNamesWithoutMods: keyNamesWithoutMods,
	}); err != nil {
		log.Fatal(err)
	}

	if err := w.Flush(); err != nil {
		log.Fatal(err)
	}
}
