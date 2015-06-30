package api

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	errUnidentifiable = errors.New("unable to identify track and problem")
	errNoFiles        = errors.New("no files submitted")
)

// Iteration represents a version of a particular exercise.
// This gets submitted to the API.
type Iteration struct {
	Key      string            `json:"key"`
	Code     string            `json:"code"`
	TestAnalysis 	 string		   `json:"test_analysis"`
	Dir      string            `json:"dir"`
	Language string            `json:"language"`
	Problem  string            `json:"problem"`
	Solution map[string]string `json:"solution"`
}

// NewIteration prepares an iteration of a problem in a track for submission to the API.
// It takes a dir and a list of files which it will read from disk.
// All paths are assumed to be absolute paths with symlinks resolved.
func NewIteration(dir string, filenames []string) (*Iteration, error) {
	if len(filenames) == 0 {
		return nil, errNoFiles
	}

	iter := &Iteration{
		Dir:      dir,
		Solution: map[string]string{},
	}

	// All the files should be within the exercism path.
	for _, filename := range filenames {
		if !iter.isValidFilepath(filename) {
			return nil, errUnidentifiable
		}
	}

	// Identify language track and problem slug.
	path := filenames[0][len(dir):]
	segments := strings.Split(path, string(filepath.Separator))
	if len(segments) < 4 {
		return nil, errUnidentifiable
	}
	iter.Language = segments[1]
	iter.Problem = segments[2]

	// Run tests and write to out file
	runTestCmd := exec.Command("rake")
	testOutfile, _ := os.Create("./test_out.txt")
	defer testOutfile.Close()
	runTestCmd.Stdout = testOutfile
	runTestCmd.Run()

	// add the test out put to iteration
	testOutput, _ := ioutil.ReadFile("./test_out.txt")
	iter.TestAnalysis = string(testOutput)

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		path := filename[len(iter.RelativePath()):]
		iter.Solution[path] = string(b)
	}
	return iter, nil
}

func (iter *Iteration) RelativePath() string {
	return filepath.Join(iter.Dir, iter.Language, iter.Problem) + string(filepath.Separator)
}

func (iter *Iteration) isValidFilepath(path string) bool {
	if iter == nil {
		return false
	}
	return strings.HasPrefix(strings.ToLower(path), strings.ToLower(iter.Dir))
}

func NewGitIteration(dir string, commit_id string) (*Iteration, error) {

	iter := &Iteration{
		Dir:      dir,
		Solution: map[string]string{},
	}

	// Identify language track and problem slug.
  path, _ := os.Getwd()
	segments := strings.Split(path, "/")
	if len(segments) < 4 {
		return nil, errUnidentifiable
	}
	iter.Language = segments[4]
	iter.Problem = segments[5]


	iter.Solution["path"] = "git"
	iter.Solution["code"] = commit_id
	return iter, nil
}



