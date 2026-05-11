# LazyVim Cheat Sheet

A focused reference for day-to-day work — keep this open in a split while you build.

## The Four Modes

Everything in Vim depends on understanding these:

- **Normal** — default mode. Where commands live. `Esc` always brings you here.
- **Insert** — typing text. Enter with `i`, `a`, `o`. Exit with `Esc` (or `jk` if you map it).
- **Visual** — selecting text. Enter with `v` (char), `V` (line), `Ctrl+v` (block).
- **Command** — typing `:` commands like `:w`, `:q`.

## Save, Quit, Undo, Redo

| Command | Action |
|---|---|
| `:w` | Save |
| `:q` | Quit |
| `:q!` | Quit without saving |
| `:wq` or `ZZ` | Save and quit |
| `u` | Undo (Normal mode) |
| `Ctrl+r` | Redo |
| `U` | Undo all changes on current line |

## Movement

Resist the arrow keys — your wrists will thank you.

| Key | Action |
|---|---|
| `h j k l` | Left, down, up, right |
| `w` / `b` | Next / previous word |
| `e` | End of word |
| `0` / `$` | Start / end of line |
| `^` | First non-blank character |
| `gg` / `G` | Top / bottom of file |
| `Ctrl+u` / `Ctrl+d` | Half-page up / down |
| `{` / `}` | Previous / next paragraph |
| `%` | Jump to matching bracket |
| `f<char>` | Jump to next `<char>` on line; `;` repeats |
| `*` | Search for word under cursor |

## Editing Basics

| Key | Action |
|---|---|
| `i` / `a` | Insert before / after cursor |
| `I` / `A` | Insert at start / end of line |
| `o` / `O` | New line below / above |
| `x` | Delete character |
| `dd` | Delete line |
| `yy` | Yank (copy) line |
| `p` / `P` | Paste after / before |
| `dw` | Delete word |
| `d$` | Delete to end of line |
| `cw` | Change word |
| `r<char>` | Replace single character |
| `R` | Replace mode |
| `.` | Repeat last change (gold) |
| `>>` / `<<` | Indent / dedent line |

**The grammar:** `<count><action><motion>`. So `d3w` deletes 3 words, `y$` yanks to end of line, `ci"` changes inside quotes. Once this clicks, Vim feels like a language.

## Visual Mode

- `v` then move — select
- `y` copy, `d` delete, `c` change, `>` indent

## Search & Replace

| Command | Action |
|---|---|
| `/foo` | Search forward |
| `?foo` | Search backward |
| `n` / `N` | Next / previous match |
| `:%s/old/new/g` | Replace all in file |
| `:%s/old/new/gc` | Replace with confirmation |

## LazyVim-Specific

The leader key in LazyVim is **Space**. Press `<Space>` and pause — **which-key** pops up showing every available shortcut. Learn it instead of memorizing.

| Shortcut | Action |
|---|---|
| `<Space>` | Open which-key menu |
| `<Space>e` | File explorer (Neo-tree) |
| `<Space><Space>` | Find files (Telescope) |
| `<Space>/` | Grep across project |
| `<Space>fb` | Find buffers |
| `<Space>sk` | Search keymaps (when you forget a binding) |
| `<Space>qq` | Quit all |
| `<Space>L` | Lazy plugin manager |
| `<Space>cm` | Mason (LSP installer) |

## Terminal Inside Nvim

| Shortcut | Action |
|---|---|
| `<Space>ft` | Open floating terminal |
| `Ctrl+/` | Toggle terminal at bottom |
| `:term` | Open a terminal in command mode |

Inside terminal, `Esc` enters Normal mode for the terminal buffer; `i` returns to typing.

For a Go project, just keep `Ctrl+/` open and run `go run .` or `go test ./...` there.

## Splits and Windows

| Shortcut | Action |
|---|---|
| `<Space>|` | Vertical split |
| `<Space>-` | Horizontal split |
| `Ctrl+w` then `h/j/k/l` | Move between splits |
| `Ctrl+w` then `q` | Close current split |

## Buffers and Tabs

| Shortcut | Action |
|---|---|
| `<Space>bd` | Close buffer |
| `Shift+h` / `Shift+l` | Previous / next buffer |
| `<Space>bp` | Pin a buffer |

## Go-Specific Bits

| Shortcut | Action |
|---|---|
| `gd` | Go to definition (LSP) |
| `gr` | References |
| `K` | Hover docs |
| `<Space>ca` | Code actions |
| `<Space>cr` | Rename symbol |
| `<Space>cf` | Format file |

`gofmt`/`goimports` runs on save by default in LazyVim's Go extra. Enable it via `:LazyExtras`, then toggle `lang.go` with `x`.

## Two Pieces of Advice

Run `:Tutor` once — it's a 25-minute interactive walkthrough built into Nvim and will save you a week of fumbling.

Don't try to memorize this list. Pin it somewhere, use which-key (`<Space>`), and let muscle memory build over a few weeks.
