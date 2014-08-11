// Copyright (c) 2014 Soichiro Kashima

package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	xmlFile, err := os.Open("Material.txt")
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer xmlFile.Close()

	b, _ := ioutil.ReadAll(xmlFile)
	for _, s := range strings.Split(string(b), "\n") {
		if s == "" {
			continue
		}
		r, g, b := hexToInt(s)
		fmt.Printf("%d %d %d %s\n", r, g, b, s)
	}
}

func hexToInt(hexString string) (r, g, b int) {
	raw := hexString
	// Remove prefix '#'
	if strings.HasPrefix(raw, "#") {
		braw := []byte(raw)
		raw = string(braw[1:])
	}

	bytes, _ := hex.DecodeString("FF" + raw)
	r = int(bytes[1])
	g = int(bytes[2])
	b = int(bytes[3])
	return
}
