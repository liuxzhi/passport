package util

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Unknwon/goconfig"
)

var (
	ConfigFile *goconfig.ConfigFile
	configPath string
)


func init() {
	osArgs := os.Args
	var err error
	for _, v := range osArgs {
		if find := strings.Contains(v, "--config="); find {
			configPath = v[9 : ]
			break
		}
	}

	if configPath == ""  {
		panic("load include files error")
	}

	if false == fileExist(configPath)  {
		panic("load include files error")
	}

	ConfigFile, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		panic(err)
	}

	if err = loadIncludeFiles(); err != nil {
		panic("load include files error:" + err.Error())
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGUSR1)

		for {
			sig := <-ch
			switch sig {
			case syscall.SIGUSR1:
				ReloadConfigFile()
			}
		}
	}()
}

func ReloadConfigFile() {
	var err error
	ConfigFile, err = goconfig.LoadConfigFile(configPath)
	if err != nil {
		fmt.Println("reload config file, error:", err)
		return
	}

	if err = loadIncludeFiles(); err != nil {
		fmt.Println("reload files include files error:", err)
		return
	}
	fmt.Println("reload config file successfully！")
}

func loadIncludeFiles() error {
	includeFile := ConfigFile.MustValue("include_files", "path", "")
	if includeFile != "" {
		includeFiles := strings.Split(includeFile, ",")

		incFiles := make([]string, len(includeFiles))
		for i, incFile := range includeFiles {
			incFiles[i] = incFile
		}
		return ConfigFile.AppendFiles(incFiles...)
	}

	return nil
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}