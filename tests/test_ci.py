"""Exercise the real update script against an isolated Git submodule."""

import importlib.util
import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


ROOT = Path(__file__).resolve().parents[1]
spec = importlib.util.spec_from_file_location(
    "release_tag", ROOT / "scripts/release-tag.py"
)
release_tag = importlib.util.module_from_spec(spec)
spec.loader.exec_module(release_tag)


class UpdateDataTests(unittest.TestCase):
    def command(self, cwd, *args, check=True):
        result = subprocess.run(
            args, cwd=cwd, env=self.env, text=True,
            stdout=subprocess.PIPE, stderr=subprocess.STDOUT,
        )
        if check and result.returncode:
            self.fail(f"{args} failed:\n{result.stdout}")
        return result

    def commit(self, repo):
        self.command(repo, "git", "add", ".")
        self.command(repo, "git", "commit", "-m", "fixture")

    def write_year(self, year=2025, name="New Year"):
        data = {
            "year": year,
            "papers": [],
            "days": [{"date": f"{year}-01-01", "name": name, "isOffDay": True}],
        }
        (self.upstream / f"{year}.json").write_text(json.dumps(data))

    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="holiday-ci-test-")
        self.addCleanup(self.temp.cleanup)
        base = Path(self.temp.name)
        self.repo = base / "project"
        self.upstream = base / "upstream"
        self.output = base / "outputs"
        self.env = os.environ.copy()
        self.env.update({
            "GITHUB_OUTPUT": str(self.output),
            "GIT_ALLOW_PROTOCOL": "file",
            "GIT_AUTHOR_NAME": "CI Test",
            "GIT_AUTHOR_EMAIL": "ci-test@example.com",
            "GIT_COMMITTER_NAME": "CI Test",
            "GIT_COMMITTER_EMAIL": "ci-test@example.com",
            "GIT_CONFIG_GLOBAL": os.devnull,
            "GIT_CONFIG_NOSYSTEM": "1",
            "GOWORK": "off",
        })
        for repo in (self.upstream, self.repo):
            repo.mkdir()
            self.command(repo, "git", "init", "--initial-branch=main")
        self.write_year()
        self.commit(self.upstream)
        self.command(
            self.repo, "git", "submodule", "add", str(self.upstream), "holiday-cn"
        )
        shutil.copy2(ROOT / "go.mod", self.repo / "go.mod")
        shutil.copytree(ROOT / "cmd", self.repo / "cmd")
        (self.repo / "pkg/holiday").mkdir(parents=True)
        self.command(
            self.repo, "go", "run", "./cmd/generator", "holiday-cn", "pkg/holiday"
        )
        self.commit(self.repo)

    def update(self, check=True):
        self.output.write_text("")
        result = self.command(
            self.repo, "bash", str(ROOT / "scripts/update-data.sh"), check=check
        )
        outputs = dict(
            line.split("=", 1) for line in self.output.read_text().splitlines()
        )
        return result, outputs["changed"]

    def generated_diff(self):
        return self.command(
            self.repo, "git", "diff", "--name-only", "HEAD", "--", "pkg/holiday"
        ).stdout

    def test_no_upstream_changes(self):
        _, changed = self.update()
        self.assertEqual(changed, "false")
        self.assertEqual(self.generated_diff(), "")

    def test_non_data_changes_are_ignored(self):
        (self.upstream / "README.md").write_text("New documentation")
        (self.upstream / "schema.json").write_text("{}")
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "false")
        self.assertEqual(self.generated_diff(), "")

    def test_year_data_changes_are_regenerated(self):
        self.write_year(name="Updated holiday")
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "true")
        self.assertIn(
            "Updated holiday",
            (self.repo / "pkg/holiday/year_2025.go").read_text(),
        )

    def test_metadata_only_json_changes_do_not_publish(self):
        year_file = self.upstream / "2025.json"
        data = json.loads(year_file.read_text())
        data["papers"] = ["https://example.com/announcement"]
        year_file.write_text(json.dumps(data))
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "false")
        self.assertEqual(self.generated_diff(), "")

    def test_added_and_removed_years(self):
        (self.upstream / "2025.json").unlink()
        self.write_year(year=2026)
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "true")
        self.assertFalse((self.repo / "pkg/holiday/year_2025.go").exists())
        self.assertTrue((self.repo / "pkg/holiday/year_2026.go").exists())
        self.command(self.repo, "go", "build", "./...")

    def test_invalid_json_fails_without_replacing_go_data(self):
        (self.upstream / "2025.json").write_text("{broken")
        self.commit(self.upstream)
        result, changed = self.update(check=False)
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(changed, "false")
        self.assertEqual(self.generated_diff(), "")


class ReleaseTagTests(unittest.TestCase):
    def test_first_v2_release_ignores_legacy_and_prerelease_tags(self):
        self.assertEqual(
            release_tag.next_tag("v0.3.20260101\nv2.0.0-rc.1", ""), "v2.0.0"
        )

    def test_patch_increment_is_numeric(self):
        self.assertEqual(
            release_tag.next_tag("v2.0.9\nv2.0.10", ""), "v2.0.11"
        )

    def test_release_tracks_latest_minor(self):
        self.assertEqual(
            release_tag.next_tag("v2.0.100\nv2.1.2", ""), "v2.1.3"
        )

    def test_retry_reuses_tag_for_same_commit(self):
        self.assertEqual(
            release_tag.next_tag("v2.0.0\nv2.0.1", "v2.0.0"), "v2.0.0"
        )


if __name__ == "__main__":
    unittest.main()
