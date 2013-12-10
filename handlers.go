package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

func ReturnString(params map[string]interface{}, body []byte) ([]byte, error, int) {
	html := string(body)
	pdfString, err := runCmdFromStdin(populateStdin(html))
	if err != nil {
		return nil, err, http.StatusBadRequest
	}
	return []byte(pdfString), nil, http.StatusOK
}

func populateStdin(str string) func(io.WriteCloser) {
	return func(stdin io.WriteCloser) {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewBufferString(str))
	}
}

func runCmdFromStdin(populate_stdin_func func(io.WriteCloser)) (string, error) {
	cmd := exec.Command("wkhtmltopdf", "-", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", fmt.Errorf("%s", "request could not be performed")
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("%s", "request could not be performed")
	}
	err = cmd.Start()
	if err != nil {
		return "", fmt.Errorf("%s", "request could not be performed")
	}
	populate_stdin_func(stdin)
	buffer := &bytes.Buffer{}
	go io.Copy(buffer, stdout)
	err = cmd.Wait()
	if err != nil {
		return "", fmt.Errorf("%s", "request could not be performed")
	}
	return buffer.String(), nil
}
