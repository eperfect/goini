package goConfig

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	cfgs     map[string](map[string]string)
	section  string
	filePath string = "config.ini"
)

const (
	SectionDefault string = "defaut"
	AscSpace       byte   = 32
	AscTab         byte   = 9
	AscEnter       byte   = 13
	AscLeftBracket byte   = 91
	AscComment     byte   = 35
)

func InitConfig(path string) {
	filePath = path
	section = SectionDefault
	cfgs = make(map[string](map[string]string))
	ParseConfigFile()
}

func ParseConfigFile() {
	cfgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err during opening file:", err)
		return
	}
	defer cfgFile.Close()

	var bfRd = bufio.NewReader(cfgFile)
	var line []byte
	for {
		line, err = bfRd.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("DONE!")
			} else {
				fmt.Println("error:", err)
			}
			break
		}
		parseConfig(line)
	}
}

func parseConfig(line []byte) {
	var finalStr string = getFilteredStr(line)
	if finalStr != "" {
		if finalStr[0] == AscLeftBracket {
			section = string(finalStr[1 : len(finalStr)-1])
		} else {
			key, value := getKeyValue(finalStr)
			SetValue(section, key, value)
		}
	}
}

func getKeyValue(kvStr string) (string, string) {
	result := strings.Split(kvStr, "=")
	return result[0], result[1]
}

func getFilteredStr(sourceStr []byte) string {
	var tarBytes = make([]byte, 0, len(sourceStr))
	for _, value := range sourceStr {
		if value == AscTab || value == AscSpace || value == AscComment {
			continue
		}
		if value == AscEnter {
			break
		}
		tarBytes = append(tarBytes, value)
	}
	return string(tarBytes)
}

func SetValue(section string, key string, value string) {
	if section == nil {
		section = SectionDefault
	}
	if cfgs[section] == nil {
		cfgs[section] = make(map[string]string)
	}
	cfgs[section][key] = value
}

func GetValue(section string, key string) string {
	if section == nil {
		section = SectionDefault
	}
	if cfgs[section] != nil {
		return cfgs[section][key]
	}
	return nil
}
