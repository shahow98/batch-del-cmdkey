package core

import (
	"bytes"
	"io"
	"os/exec"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

const targetT string = "目标: "

func GetAllCmdKeys() ([]string, error) {
	cmdkeyCmd := exec.Command("cmdkey", "/list")

	cmdkeyOut, err := cmdkeyCmd.Output()
	if err != nil {
		return nil, err
	}
	utf8Data, err := GbkToUtf8(cmdkeyOut)
	if err != nil {
		return nil, err
	}

	data := string(utf8Data)
	lines := strings.Split(data, "\n")

	var targets []string
	for _, line := range lines {
		if strings.Contains(line, targetT) {
			target := strings.ReplaceAll(line, targetT, "")
			target = strings.TrimSpace(target)
			targets = append(targets, target)
		}
	}
	return targets, nil
}

func DelCmdkeys(targets []string) error {
	for _, target := range targets {
		cmdkeyCmd := exec.Command("cmdkey", "/delete:"+target)
		_, err := cmdkeyCmd.Output()
		if err != nil {
			return err
		}
	}
	return nil
}

func GbkToUtf8(gbkData []byte) ([]byte, error) {
	decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(gbkData), decoder)
	utf8Data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return utf8Data, nil
}
