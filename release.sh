#!/usr/bin/env bash

set -euo pipefail

# Optional: ensure you're on the correct branch
RELEASE_BRANCH="main"
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)

if [[ "$CURRENT_BRANCH" != "$RELEASE_BRANCH" ]]; then
  echo "⚠️  Not on '$RELEASE_BRANCH' branch (currently on '$CURRENT_BRANCH')."
  read -p "Continue anyway? (y/N): " confirm
  [[ "$confirm" =~ ^[Yy]$ ]] || exit 1
fi

# Optional: check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
  echo "❌ You have uncommitted changes. Please commit or stash them before releasing."
  exit 1
fi

# Ensure changelog exists or prompt for it
if [[ ! -f CHANGELOG.md ]]; then
  echo "⚠️  CHANGELOG.md not found. Consider adding one before release."
fi

# Tagging (optional: ask user or auto-bump)
echo "📌 The latest tag is:"
git describe --tags --abbrev=0 || echo "(none found)"
read -p "Enter new version tag (e.g. v1.2.3): " VERSION

# Create git tag
git tag "$VERSION"
git push origin "$VERSION"

# Run GoReleaser (you can add --snapshot or --rm-dist for testing)
echo "🚀 Running GoReleaser..."
goreleaser release --clean

echo "✅ Release $VERSION complete!"
