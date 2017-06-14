package test

import (
	"testing"
	"github.com/inver/gostone/test/infrastructure"
	"bytes"
	"fmt"
	"os"
	"archive/zip"
	"path/filepath"
	"io"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
	"github.com/inver/gostone/parser"
	"github.com/inver/gostone/evaluator"
	"strconv"
)

func unzip(url, target string) error {
	content, err := loadArchive(url)
	if err != nil {
		return err
	}

	size := len(content)
	reader, err := zip.NewReader(bytes.NewReader(content), int64(size))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func loadArchive(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func loadTests() (map[string][]infrastructure.TestUnit, error) {
	err := unzip("https://github.com/MegafonWebLab/histone-java2/archive/master.zip", "./tmp")
	if err != nil {
		return nil, err
	}

	basedir := "./tmp/histone-java2-master/test/src/main/resources"
	res := make(map[string][]infrastructure.TestUnit)

	filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".json") && !info.IsDir() {
			file, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			var units []infrastructure.TestUnit
			err = json.Unmarshal(file, &units)
			if err != nil {
				//panic(err) do nothing
				return nil
			}
			res[path] = units
		}
		return nil
	})

	return res, nil
}

func TestConcreteSimple(t *testing.T) {
	units, err := loadTests()
	if err != nil {
		fmt.Println(err)
	}

	baseUri := ""
	parser := new(parser.Parser)
	teamplateEvaluator := new(evaluator.Evaluator)
	for _, unitFile := range units {
		for _, testCases := range unitFile {
			//todo add more information to panic
			for index, testCase := range testCases.Cases {
				testName := testCases.Name + "$" + strconv.Itoa(index)
				t.Run(testName, func(t *testing.T) {
					root, err := parser.Process(testCase.Input, baseUri)
					if err != nil {
						panic(err)
					}
					ctx := make(map[string]evaluator.EvalNode)
					res, err := teamplateEvaluator.Process(*root, ctx)

					//todo rework to t.Run()
					if res != testCase.ExpectedResult {
						t.Fail()
					}
				})
			}
		}

	}
}
