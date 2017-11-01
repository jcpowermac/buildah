package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

var (
	Var1 = []string{"ONE=1", "TWO=2"}
)

func createTmpFile(content []byte) (string, error) {
	tmpfile, err := ioutil.TempFile(os.TempDir(), "unittest")
	if err != nil {
		return "", err
	}

	if _, err := tmpfile.Write(content); err != nil {
		return "", err

	}
	if err := tmpfile.Close(); err != nil {
		return "", err
	}
	return tmpfile.Name(), nil
}

func TestConvertStringSliceToMap(t *testing.T) {
	strSlice := []string{"BLAU=BLUE", "GELB=YELLOW"}
	result, _ := convertStringSliceToMap(strSlice, "=")
	assert.Equal(t, result["BLAU"], "BLUE")
}

func TestConvertStringSliceToMapBadData(t *testing.T) {
	strSlice := []string{"BLAU=BLUE", "GELB^YELLOW"}
	_, err := convertStringSliceToMap(strSlice, "=")
	assert.Error(t, err)
}

func TestGetAllLabels(t *testing.T) {
	fileLabels := []string{}
	labels, _ := getAllLabels(fileLabels, Var1)
	assert.Equal(t, len(labels), 2)
}

func TestGetAllLabelsBadKeyValue(t *testing.T) {
	inLabels := []string{"ONE1", "TWO=2"}
	fileLabels := []string{}
	_, err := getAllLabels(fileLabels, inLabels)
	assert.Error(t, err, assert.AnError)
}

func TestGetAllLabelsBadLabelFile(t *testing.T) {
	fileLabels := []string{"/foobar5001/be"}
	_, err := getAllLabels(fileLabels, Var1)
	assert.Error(t, err, assert.AnError)
}

func TestGetAllLabelsFile(t *testing.T) {
	content := []byte("THREE=3")
	tFile, err := createTmpFile(content)
	defer os.Remove(tFile)
	assert.NoError(t, err)
	fileLabels := []string{tFile}
	result, _ := getAllLabels(fileLabels, Var1)
	assert.Equal(t, len(result), 3)
}

func TestGetAllEnvironmentVariables(t *testing.T) {
	fileEnvs := []string{}
	result, _ := getAllEnvironmentVariables(fileEnvs, Var1)
	assert.Equal(t, len(result), 2)
}

func TestGetAllEnvironmentVariablesBadKeyValue(t *testing.T) {
	inEnvs := []string{"ONE1", "TWO=2"}
	fileEnvs := []string{}
	_, err := getAllEnvironmentVariables(fileEnvs, inEnvs)
	assert.Error(t, err, assert.AnError)
}

func TestGetAllEnvironmentVariablesBadEnvFile(t *testing.T) {
	fileEnvs := []string{"/foobar5001/be"}
	_, err := getAllEnvironmentVariables(fileEnvs, Var1)
	assert.Error(t, err, assert.AnError)
}

func TestGetAllEnvironmentVariablesFile(t *testing.T) {
	content := []byte("THREE=3")
	tFile, err := createTmpFile(content)
	defer os.Remove(tFile)
	assert.NoError(t, err)
	fileEnvs := []string{tFile}
	result, _ := getAllEnvironmentVariables(fileEnvs, Var1)
	assert.Equal(t, len(result), 3)
}
