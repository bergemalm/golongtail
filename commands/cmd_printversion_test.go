package commands

import (
	"io/ioutil"
	"testing"
)

func TestPrintVersionIndex(t *testing.T) {
	testPath, _ := ioutil.TempDir("", "test")
	fsBlobPathPrefix := "fsblob://" + testPath
	createVersionData(t, fsBlobPathPrefix)
	executeCommandLine("upsync", "--source-path", testPath+"/version/v1", "--target-path", fsBlobPathPrefix+"/index/v1.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")
	executeCommandLine("upsync", "--source-path", testPath+"/version/v2", "--target-path", fsBlobPathPrefix+"/index/v2.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")
	executeCommandLine("upsync", "--source-path", testPath+"/version/v3", "--target-path", fsBlobPathPrefix+"/index/v3.lvi", "--storage-uri", fsBlobPathPrefix+"/storage")

	cmd, err := executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v1.lvi")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}
	cmd, err = executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v2.lvi")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}
	cmd, err = executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v3.lvi")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}

	cmd, err = executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v1.lvi", "--compact")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}
	cmd, err = executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v2.lvi", "--compact")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}
	cmd, err = executeCommandLine("print-version", "--version-index-path", fsBlobPathPrefix+"/index/v3.lvi", "--compact")
	if err != nil {
		t.Errorf("%s: %s", cmd, err)
	}
}