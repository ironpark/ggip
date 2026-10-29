# ggip

> **Merged into [ggwd](https://github.com/ironpark/ggwd).** The keys, `Mods`, `Keyboard` and
> `Chord` are in package `ggwd`, whose events carry them, and `gamepad` and `gamepaddb` are
> `ggwd/gamepad` and `ggwd/gamepaddb`. `ggfxinput.StartGamepads` is `gamepad.Start`, and the
> conversions are gone. This repository is no longer developed.

ggip (go general input) handles input for a GUI: the keyboard and the mouse once a window system
has reported them, and the gamepads, which it reads itself. It was split out of
[ggfx](https://github.com/ironpark/ggfx) and [ggui](https://github.com/ironpark/ggui), and
depends on neither.

- `ggip`: the keys, named after the [UI Events code values](https://www.w3.org/TR/uievents-code/)
  (`Key.Code`, `KeyFromCode`), the modifiers (`Mods`, with `Cmd` for the platform's command key),
  the mouse buttons, `Keyboard`, which tracks the held keys, and `Chord`, a shortcut such as
  `"cmd+shift+z"` that parses, matches a press and labels itself for a menu (`⇧⌘Z`,
  `Ctrl+Shift+Z`).
- `gamepad`: reads the gamepads (IOKit and GameController on macOS, DirectInput and XInput on
  Windows, evdev on Linux, the Gamepad API in the browser), and `Tracker` turns the polls into
  connect, disconnect, button and axis events.
- `gamepaddb`: the standard layout mappings in SDL_GameControllerDB's format.
- `ggfxinput`: a module of its own that connects ggwd, ggfx's window system, to ggip, so that
  neither imports the other.

## Using it with ggwd

`ggfxinput` converts the keys, buttons and modifiers in ggwd's events, and polls the gamepads on
the main thread, waking the loop with `ggwd.Wake` when it found something:

```go
type app struct {
	keyboard ggip.Keyboard
	pads     *ggfxinput.Gamepads
	save     ggip.Chord
}

func (a *app) HandleEvent(ev ggwd.Event) error {
	switch ev := ev.(type) {
	case ggwd.StartEvent:
		w, err := ggwd.NewWindow(&ggwd.WindowOptions{Title: "input"})
		if err != nil {
			return err
		}
		a.save = ggip.MustChord("cmd+s")
		a.pads = ggfxinput.StartGamepads(w)
	case ggwd.KeyEvent:
		key, ok := ggfxinput.Key(ev.Key)
		if !ok {
			break
		}
		a.keyboard.Set(key, ev.Pressed)
		mods := a.keyboard.Mods().Or(ggfxinput.Mods(ev.Modifiers))
		if ev.Pressed && a.save.Matches(key, mods) {
			log.Print("save")
		}
	case ggwd.FocusEvent:
		if !ev.Focused {
			a.keyboard.Reset()
		}
	case ggwd.WakeEvent:
		for _, ev := range a.pads.Take() {
			if ev, ok := ev.(gamepad.StandardButtonEvent); ok {
				log.Printf("gamepad %d: %v pressed=%v", ev.ID, ev.Button, ev.Pressed)
			}
		}
	}
	return nil
}
```

Without ggfxinput, poll the gamepads the way `ggfxinput/gamepads.go` does: call `gamepad.Update`
on the main thread, pass a window handle on Windows, and hand `Tracker.Update`'s events to the
loop. See the `gamepad` package documentation.

The gamepad code comes from [Ebitengine](https://ebitengine.org/) v2.10.2 through ggfx.

## License

Apache-2.0, the same as Ebitengine. See [LICENSE](LICENSE) and [NOTICE.md](NOTICE.md).
