package main

import (
	"bytes"
	"fmt"
	chunker "github.com/ipfs/go-ipfs-chunker"
	dagtest "github.com/ipfs/go-merkledag/test"
	"github.com/ipfs/go-unixfs/importer/balanced"
	ihelper "github.com/ipfs/go-unixfs/importer/helpers"
	"io/ioutil"
	"os"
)

func genIpfsCID(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	chunk, err := chunker.FromString(bytes.NewReader(data), "")
	if err != nil {
		panic(err)
	}

	dag := dagtest.Mock()

	params := ihelper.DagBuilderParams{
		Dagserv:   dag,
		RawLeaves: false,
		Maxlinks:  ihelper.DefaultLinksPerBlock,
		NoCopy:    false,
	}

	pp, err := params.New(chunk)
	if err != nil {
		panic(err)
	}

	result, err := balanced.Layout(pp)
	if err != nil {
		panic(err)
	}
	fmt.Println("File: ", path)
	fmt.Println("Generated CID: ", result.String())

}

func main() {
	if len(os.Args) != 2 {
		panic("Wrong argument number.")
	}

	args := os.Args[1:]

	path := args[0]

	genIpfsCID(path)
}
