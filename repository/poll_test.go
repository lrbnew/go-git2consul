package repository

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestPollBranches(t *testing.T) {

	r, cleanup := tempGitInitPath(t)
	defer cleanup()

	cfg := loadConfig(t)

	repos, err := loadRepos(cfg)
	if err != nil {
		t.Fatal(err)
	}
	repo := repos[0]

	expected := tempCommitRepo(r, t)

	err = repo.PollBranches()
	if err != nil {
		t.Fatal(err)
	}

	//Verify that the file changed
	filePath := filepath.Join("test-fixtures", "example", "foo")
	actual, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Polling failed to pull files")
	}

	// Cleanup
	defer func() {
		repos[0].CloneCh()
		err = os.RemoveAll(repo.store)
		if err != nil {
			t.Fatal(err)
		}
	}()
}