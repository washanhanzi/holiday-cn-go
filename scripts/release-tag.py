"""Choose a UTC timestamp tag in the v0.4 series for the unsuffixed module."""

from datetime import datetime, timedelta, timezone
from pathlib import Path
import re
import subprocess


MODULE_PATH = "github.com/washanhanzi/holiday-cn-go"
TAG_PATTERN = re.compile(r"v0\.4\.([1-9][0-9]{13})")


def timestamps(tags):
    return [
        datetime.strptime(match.group(1), "%Y%m%d%H%M%S").replace(
            tzinfo=timezone.utc
        )
        for tag in tags.splitlines()
        if (match := TAG_PATTERN.fullmatch(tag))
    ]


def next_tag(all_tags, head_tags, now):
    existing = timestamps(head_tags)
    if existing:
        timestamp = max(existing)
    else:
        timestamp = now.astimezone(timezone.utc).replace(microsecond=0)
        released = timestamps(all_tags)
        if released:
            timestamp = max(timestamp, max(released) + timedelta(seconds=1))
    return "v0.4." + timestamp.strftime("%Y%m%d%H%M%S")


def validate_module(go_mod):
    match = re.search(r"^module\s+(\S+)\s*$", go_mod, re.MULTILINE)
    if not match or match.group(1) != MODULE_PATH:
        raise ValueError(f"v0.4 releases require module {MODULE_PATH}")


if __name__ == "__main__":
    validate_module(Path("go.mod").read_text())
    all_tags = subprocess.check_output(["git", "tag", "--list"], text=True)
    head_tags = subprocess.check_output(
        ["git", "tag", "--points-at", "HEAD"], text=True
    )
    print(next_tag(all_tags, head_tags, datetime.now(timezone.utc)))
