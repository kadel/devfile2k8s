package main

import (
	"flag"
	"fmt"

	"github.com/devfile/library/pkg/devfile"
	"github.com/devfile/library/pkg/devfile/generator"
	"github.com/devfile/library/pkg/devfile/parser"
	"github.com/devfile/library/pkg/devfile/parser/data/v2/common"
	"sigs.k8s.io/yaml"
)

func main() {

	var inputFile string
	flag.StringVar(&inputFile, "input", "devfile.yaml", "Input devfile file")

	parserArgs := parser.ParserArgs{
		Path: inputFile,
	}

	devfile, _, err := devfile.ParseDevfileAndValidate(parserArgs)
	if err != nil {
		panic(err)
	}
	containers, err := generator.GetContainers(devfile, common.DevfileOptions{})
	if err != nil {
		panic(err)
	}
	initContainers, err := generator.GetInitContainers(devfile)
	if err != nil {
		panic(err)
	}

	deployParams := generator.DeploymentParams{
		Containers:     containers,
		InitContainers: initContainers,
		ObjectMeta:     generator.GetObjectMeta("test", "test", map[string]string{}, map[string]string{}),
		TypeMeta:       generator.GetTypeMeta("Deployment", "v1"),
	}

	deployment := generator.GetDeployment(deployParams)

	out, err := yaml.Marshal(deployment)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
