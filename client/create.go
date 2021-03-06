package client

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	gflag "github.com/jessevdk/go-flags"
)

func (cli *HyperClient) HyperCmdCreate(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("\"create\" requires a minimum of 1 argument, please provide POD spec file.\n")
	}
	var opts struct {
		Yaml      bool     `short:"y" long:"yaml" default:"false" default-mask:"-" description:"create a pod based on Yaml file"`
	}
	var parser = gflag.NewParser(&opts, gflag.Default)
	parser.Usage = "create [OPTIONS] POD_FILE\n\ncreate a pod, but without running it"
	args, err := parser.Parse()
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			return err
		} else {
			return nil
		}
	}
	jsonFile := args[1]
	if _, err := os.Stat(jsonFile); err != nil {
		return err
	}
	jsonbody, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	if opts.Yaml == true {
		jsonbody, err = cli.ConvertYamlToJson(jsonbody)
		if err != nil {
			return err
		}
	}
	podId, err := cli.CreatePod(string(jsonbody))
	if err != nil {
		return err
	}
	fmt.Printf("Pod ID is %s\n", podId)
	return nil
}
