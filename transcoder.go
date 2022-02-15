package main

import (
	"os"
	"strings"
	"path/filepath"
	"github.com/modfy/fluent-ffmpeg"
	"io/ioutil"
	"fmt"
	"bytes"
)

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, "."+filepath.Ext(fileName))
}

func main() {
	filename := os.Args[1]
	var extension = filepath.Ext(filename)
	var name = filename[0:len(filename)-len(extension)]
	fmt.Printf(filename)
	fmt.Printf("\n/Users/amirhossiensorouri/Desktop/projects/stageshine/transcoder/videos/"+name+".m3u8")

	buf := &bytes.Buffer{}
	err := fluentffmpeg.NewCommand("").
			InputPath("/Users/amirhossiensorouri/Desktop/projects/stageshine/create-rest-api-in-go-tutorial/videos/"+filename).
			OutputFormat("hls").
			OutputPath("/Users/amirhossiensorouri/Desktop/projects/stageshine/transcoder/videos/"+name+".m3u8").
			Overwrite(true).
			OutputLogs(buf). // provide a io.Writer
			Run()
	
	if err != nil {
		fmt.Printf("%+v", err)
	}

	out, _ := ioutil.ReadAll(buf) // read logs
	fmt.Println(string(out))
	return
}