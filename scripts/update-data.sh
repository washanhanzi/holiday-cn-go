#!/usr/bin/env bash
set -euo pipefail

: "${GITHUB_OUTPUT:?GITHUB_OUTPUT must point to the step output file}"
echo "changed=false" >> "$GITHUB_OUTPUT"

previous=$(git rev-parse HEAD:holiday-cn)
git submodule update --init --remote -- holiday-cn
current=$(git -C holiday-cn rev-parse HEAD)

changed_files=$(git -C holiday-cn diff --name-only "$previous" "$current" -- ':(top,glob)[0-9][0-9][0-9][0-9].json')
if [[ -z "$changed_files" ]]; then
  echo "No holiday JSON changes."
  exit 0
fi

# Generate in a temporary directory so failures leave committed Go data intact.
generated_dir=$(mktemp -d)
trap 'rm -rf "$generated_dir"' EXIT
go run ./cmd/generator holiday-cn "$generated_dir"

# Remove obsolete year files when upstream removes or renames a data file.
for year_file in pkg/holiday/year_[0-9][0-9][0-9][0-9].go; do
  if [[ -f "$year_file" ]]; then
    rm -- "$year_file"
  fi
done
cp "$generated_dir"/*.go pkg/holiday/

# Include newly generated files in the diff without staging their contents.
git add --intent-to-add -- pkg/holiday
if git diff --quiet -- pkg/holiday; then
  echo "Holiday JSON changed, but generated Go data is unchanged."
  exit 0
fi

echo "changed=true" >> "$GITHUB_OUTPUT"
