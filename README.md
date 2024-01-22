# quick-color

Color terminal output from stdout. That is all.

```zsh
echo "hello world" | quick-color '#0a0'
```

I had a hard time finding a binary that does this. Someone point it out if it exists.

The main goal is to produce colored terminal output for tailing log levels in a way that doesn't use hard-to-understand escape sequences.

## Installation

See [releases](https://github.com/MattyRad/quick-color/releases).

## Examples

```zsh
#!/usr/bin/env zsh

tail ${@:1} > /dev/null | tee \
	>(grep --line-buffered DEBUG | quick-color '#333') \
	>(grep --line-buffered INFO | quick-color '#888') \
	>(grep --line-buffered NOTICE | quick-color '#fff') \
	>(grep --line-buffered WARNING | quick-color '#fccb00') \
	>(grep --line-buffered ERROR | quick-color '#ff6900') \
	>(grep --line-buffered CRITICAL | quick-color --bold '#ff6900') \
	>(grep --line-buffered ALERT | quick-color '#f44336') \
	>(grep --line-buffered EMERGENCY | quick-color --bold '#f44336') \
	> /dev/null
```

```zsh
#!/usr/bin/env bash

tail ${@:1} | tee \
	>(grep --line-buffered DEBUG | quick-color '#333') \
	>(grep --line-buffered INFO | quick-color '#888') \
	>(grep --line-buffered NOTICE | quick-color '#fff') \
	>(grep --line-buffered WARNING | quick-color '#fccb00') \
	>(grep --line-buffered ERROR | quick-color '#ff6900') \
	>(grep --line-buffered CRITICAL | quick-color --bold '#ff6900') \
	>(grep --line-buffered ALERT | quick-color '#f44336') \
	>(grep --line-buffered EMERGENCY | quick-color --bold '#f44336') \
	> /dev/null
```
