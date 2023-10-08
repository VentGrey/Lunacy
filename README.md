# Lunacy: Keybinds viewer for LeftWM 🖥️⌨️

Easily view your LeftWM keybindings in a sleek terminal interface. No more fumbling through config files!

## TODO 💡
- Create a Makefile
- Create CI/CD
- Add tests
- Improve TUI
- Add to [Epitaph](https://github.com/VentGrey/Epitaph)

## Introduction 📖

LeftWM is an amazing window manager, but with the transition from TOML to RON configuration, I needed an easier way to view and manage my keybindings. Enter *"Lunacy"*, the Keybinds Viewer: a simple and intuitive terminal-based UI to help you quickly reference your LeftWM keybinds.

## Features ✨

- **Simple UI**: No more scrolling through configuration files.
- **Intuitive Display**: Keybindings are presented in an easy-to-read table format.
- **Grouping**: Similar commands are grouped together for better clarity.

## Installation 📥

```shell
go build -o keybinds-viewer
sudo mv keybinds-viewer /usr/local/bin/
```

## Usage 🚀

After installation, simply run:

```shell
lunacy
```

## Why Go and not Rust? 🤔

While Rust is a fantastic language with many safety features and has a growing ecosystem (and the language used to make LeftWM itself), my decision to use Go was driven by a few factors:

1. Simplicity: Go's straightforward syntax and semantics make it easy to read and write. This was crucial for quickly developing a prototype and making iterations.

2. Standard Library: Go's extensive standard library, especially for tasks related to file I/O and regular expressions, played a significant role in the decision.

3. Concurrency: Although this project doesn't leverage Go's concurrency features, having the Goroutines model available for future extensions was a plus.

4. Popularity in System Tools: Many system-level tools are being written in Go due to its simplicity and the ease of distributing static binaries.

5. All of the above were lies, I haven't written Rust in a long time, I don't want to deal with the `ron` library and it's serde mess, besides, I really don't need to parse the whole RON config, just the keybindings part.

In essence, while Rust might have been a good fit for this project, Go's attributes made it the language of choice for quick development, easy distribution, and potential future extensions.

## License 📜
This project is licensed under the GPL-3 License. See the LICENSE file for details.
