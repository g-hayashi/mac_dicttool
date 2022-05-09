package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"howett.net/plist"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	var mode, file string
	if len(os.Args) != 3 || (os.Args[1] != "google" && os.Args[1] != "mac") {
		fmt.Print("Program args: <progname> [mac|google](case sensitive) [dict file]\n")
		fmt.Print("mac: convert plist to txt(utf-16)\n")
		fmt.Print("google: convert utf8 to utf16\n")
		return
	} else {
		mode = os.Args[1]
		file = os.Args[2]
	}
	// ソースファイルを開く
	source, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer source.Close()

	// 書き込み先ファイルを用意
	dest, err := os.Create("out_" + file + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	switch mode {
	case "mac":
		var data interface{}
		plist.NewDecoder(source).Decode(&data)
		writer := bufio.NewWriter(
			transform.NewWriter(dest, unicode.UTF16(
				unicode.LittleEndian, unicode.UseBOM).NewEncoder()))

		// plist直下がまずarrayなので、まずは[]を走査する
		for _, i := range data.([]interface{}) {
			// arrayのなかはdict(map)なので、mapのKey,Valueを走査する
			dict := i.(map[string]interface{})
			var phrase, shortcut string
			for j, k := range dict {
				switch j {
				case "phrase":
					phrase = k.(string)
				case "shortcut":
					shortcut = k.(string)
				}
			}
			line := fmt.Sprintf("%s\t%s\t%s\t", phrase, shortcut, "名詞")
			fmt.Fprintln(writer, line)
		}
		writer.Flush()
	case "google":
		scanner := bufio.NewScanner(
			transform.NewReader(source, unicode.UTF8.NewDecoder()))
		writer := bufio.NewWriter(
			transform.NewWriter(dest, unicode.UTF16(
				unicode.LittleEndian, unicode.UseBOM).NewEncoder()))

		// 変換しながら書き込み
		for scanner.Scan() {
			_, err = fmt.Fprintln(writer, scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		writer.Flush()
	}
	log.Println("done")
}
