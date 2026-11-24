# ani-track
Your simple and elegant terminal-based tool for tracking anime, manga, and novel progress.

## Building
Requires ``Go 1.22+``
### Run Directly
```bash
go run .
```
### Build Binary
```bash
go build -o ani-track .
```

## Keybindings

| Key / Shortcut | Action |
| --- | --- |
| `Tab` / `Shift+Tab` | Cycle focus between **Watch History**, **Watch List**, and **Action Section** |
| `Up` / `Down` | Navigate rows within focused table or between action buttons |
| `Enter` | Open Modify dialog for highlighted entry (in table) or activate focused button |
| `a` / `A` | Open **Add Entry** dialog |
| `m` / `M` | Open **Modify / Delete Entry** dialog |
| `d` / `D` | Open **Export / Download** dialog |
| `q` / `Q` | Quit application |
| `?` / `F1` | Show in-app help modal |
| `Esc` | Cancel / close open modal form |
