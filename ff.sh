#!/usr/bin/env bash
set -euo pipefail

TITLE=$(i3-fast-find | dmenu -l 10)
i3-fast-find -focus "$TITLE"
