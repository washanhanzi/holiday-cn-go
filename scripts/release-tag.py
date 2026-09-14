"""Choose a stable v2 tag, reusing an existing tag for this exact commit."""

import re
import subprocess


def versions(tags):
    return [
        tuple(map(int, match.groups()))
        for tag in tags.splitlines()
        if (match := re.fullmatch(r"v2\.([0-9]+)\.([0-9]+)", tag))
    ]


def next_tag(all_tags, head_tags):
    existing = versions(head_tags)
    if existing:
        minor, patch = max(existing)
    elif released := versions(all_tags):
        minor, patch = max(released)
        patch += 1
    else:
        minor, patch = 0, 0
    return f"v2.{minor}.{patch}"


if __name__ == "__main__":
    all_tags = subprocess.check_output(["git", "tag", "--list"], text=True)
    head_tags = subprocess.check_output(
        ["git", "tag", "--points-at", "HEAD"], text=True
    )
    print(next_tag(all_tags, head_tags))
