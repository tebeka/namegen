package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	enCaser = cases.Title(language.English)
)

// gallant_antonelli -> Gallant Antonelli
func humanize(name string) string {
	name = strings.Replace(name, "_", " ", 1)

	return enCaser.String(name)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: namegen\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	rand.Seed(time.Now().Unix())
	name := namesgenerator.GetRandomName(0)
	fmt.Println(humanize(name))
}
