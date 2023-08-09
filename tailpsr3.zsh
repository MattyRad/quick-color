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
