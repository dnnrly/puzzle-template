package main

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunDetectsMissingMainFile(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("cmd/puzzle", 0755)

	err := Run(&Config{
		fs: memFs,
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to read main file")
}

func TestRunFailsWhenWritingInvalidFile(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("cmd/puzzle", 0755)
	afero.WriteFile(memFs, "cmd/puzzle/main.go", []byte(`func initPuzzles() {
		// next puzzle
	}`), 0644)

	err := Run(&Config{
		fs: afero.NewReadOnlyFs(memFs),
	})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unable to write main file")
}

func TestRunRegistersNewTest(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("cmd/puzzle", 0755)
	afero.WriteFile(memFs, "cmd/puzzle/main.go", []byte(`func initPuzzles() {
		// next puzzle
	}`), 0644)

	err := Run(&Config{
		fs: memFs,
	})
	assert.NoError(t, err)

	contents, err := afero.ReadFile(memFs, "cmd/puzzle/main.go")
	assert.NoError(t, err)
	assert.Contains(t, string(contents), "append(Solutions, puzzle.Puzzle001)")
	assert.Contains(t, string(contents), "// next puzzle")
}

func TestRunCreatesPuzzleFile(t *testing.T) {
	memFs := afero.NewMemMapFs()
	memFs.MkdirAll("cmd/puzzle", 0755)
	afero.WriteFile(memFs, "cmd/puzzle/main.go", []byte(`func initPuzzles() {
		// next puzzle
	}`), 0644)

	err := Run(&Config{
		fs: memFs,
	})
	assert.NoError(t, err)

	contents, err := afero.ReadFile(memFs, "puzzle001.go")
	assert.NoError(t, err)
	assert.Contains(t, string(contents), "package puzzle")
	assert.Contains(t, string(contents), "func Puzzle001() int")
}
