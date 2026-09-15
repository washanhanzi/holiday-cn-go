"""Exercise the real update script against an isolated Git submodule."""

import importlib.util
from datetime import datetime, timedelta, timezone
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
        shutil.copy2(ROOT / "holidaycn.go", self.repo / "holidaycn.go")
        shutil.copytree(ROOT / "cmd", self.repo / "cmd")
        (self.repo / "pkg/holiday").mkdir(parents=True)
        self.command(
            self.repo, "go", "run", "./cmd/generator", "holiday-cn", "pkg/holiday",
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
            self.repo, "git", "diff", "--name-only", "HEAD", "--", "pkg/holiday",
        ).stdout

    def assert_cross_year_query(self, off, name):
        (self.repo / "lookup_test.go").write_text(f'''package holidaycn

import (
    "testing"
    "time"
)

func TestUpdatedLookup(t *testing.T) {{
    date := time.Date(2024, 12, 28, 0, 0, 0, 0, time.UTC)
    off, name, err := IsRestDay(date)
    if err != nil || off != {str(off).lower()} || name != {json.dumps(name)} {{
        t.Fatalf("unexpected result: (%v, %q, %v)", off, name, err)
    }}
}}
''')
        self.command(self.repo, "go", "test", ".")

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

    def test_cross_year_lookup_is_added_and_removed(self):
        self.write_year(year=2024)
        year_file = self.upstream / "2025.json"
        data = json.loads(year_file.read_text())
        data["days"].append({
            "date": "2024-12-28", "name": "Makeup day", "isOffDay": False,
        })
        year_file.write_text(json.dumps(data))
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "true")
        self.assertIn("pkg/holiday/year_2024.go", self.generated_diff())
        self.assert_cross_year_query(False, "Makeup day")
        self.commit(self.repo)

        (self.upstream / "2025.json").unlink()
        self.commit(self.upstream)
        _, changed = self.update()
        self.assertEqual(changed, "true")
        self.assertIn("pkg/holiday/year_2024.go", self.generated_diff())
        self.assert_cross_year_query(True, "")

    def test_invalid_date_preserves_generated_files(self):
        year_file = self.upstream / "2025.json"
        data = json.loads(year_file.read_text())
        data["days"][0]["date"] = "2025-02-30"
        year_file.write_text(json.dumps(data))
        self.commit(self.upstream)
        result, changed = self.update(check=False)
        self.assertNotEqual(result.returncode, 0)
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
    now = datetime(2026, 9, 14, 10, 30, 45, tzinfo=timezone.utc)

    def test_v04_release_ignores_legacy_and_prerelease_tags(self):
        self.assertEqual(
            release_tag.next_tag(
                "v0.3.20260101\nv0.4.20260914\nv2.0.0\nv0.4.20260914120000-rc.1",
                "", self.now,
            ),
            "v0.4.20260914103045",
        )

    def test_same_second_collision_advances_to_next_second(self):
        self.assertEqual(
            release_tag.next_tag("v0.4.20260914103045", "", self.now),
            "v0.4.20260914103046",
        )

    def test_clock_skew_preserves_version_order_across_midnight(self):
        self.assertEqual(
            release_tag.next_tag("v0.4.20260914235959", "", self.now),
            "v0.4.20260915000000",
        )

    def test_retry_reuses_tag_for_same_commit(self):
        self.assertEqual(
            release_tag.next_tag(
                "v0.4.20260913120000\nv0.4.20260914103045",
                "v0.4.20260913120000", self.now,
            ),
            "v0.4.20260913120000",
        )

    def test_tags_use_utc(self):
        self.assertEqual(
            release_tag.next_tag(
                "", "",
                datetime(2026, 9, 14, 18, 30, 45, tzinfo=timezone(timedelta(hours=8))),
            ),
            "v0.4.20260914103045",
        )

    def test_module_path_must_match_release_series(self):
        release_tag.validate_module(
            "module github.com/washanhanzi/holiday-cn-go\n\ngo 1.16\n"
        )
        for module in ("github.com/washanhanzi/holiday-cn-go/v2", "example.com/other"):
            with self.subTest(module=module), self.assertRaises(ValueError):
                release_tag.validate_module(f"module {module}\n\ngo 1.16\n")


if __name__ == "__main__":
    unittest.main()
