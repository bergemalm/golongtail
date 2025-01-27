package commands

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestValidateVersion(t *testing.T) {
	testPath, _ := os.MkdirTemp("", "test")
	fsBlobPathPrefix := "fsblob://" + testPath
	createVersionData(t, fsBlobPathPrefix)
	executeCommandLine("upsync", "--source-path", testPath+"/version/v1", "--target-path", fsBlobPathPrefix+"/index/v1.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")
	executeCommandLine("upsync", "--source-path", testPath+"/version/v2", "--target-path", fsBlobPathPrefix+"/index/v2.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")
	executeCommandLine("upsync", "--source-path", testPath+"/version/v3", "--target-path", fsBlobPathPrefix+"/index/v3.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")

	cmd, err := executeCommandLine("validate-version", "--storage-uri", fsBlobPathPrefix+"/storage", "--version-index-path", fsBlobPathPrefix+"/index/v1.lvi")
	assert.NoError(t, err, cmd)
	cmd, err = executeCommandLine("validate-version", "--storage-uri", fsBlobPathPrefix+"/storage", "--version-index-path", fsBlobPathPrefix+"/index/v2.lvi")
	assert.NoError(t, err, cmd)
	cmd, err = executeCommandLine("validate-version", "--storage-uri", fsBlobPathPrefix+"/storage", "--version-index-path", fsBlobPathPrefix+"/index/v3.lvi")
	assert.NoError(t, err, cmd)

	os.RemoveAll(path.Join(testPath, "storage"))
	executeCommandLine(fmt.Sprintf("init-remote-store --storage-uri %s", fsBlobPathPrefix+"/storage"))

	cmd, err = executeCommandLine("validate-version", "--storage-uri", fsBlobPathPrefix+"/storage", "--version-index-path", fsBlobPathPrefix+"/index/v3.lvi")
	if err == nil {
		t.Errorf("%s: OK", cmd)
	}
}
