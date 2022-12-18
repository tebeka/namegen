package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/docker/docker/pkg/namesgenerator"
)

// gallant_antonelli -> Gallant Antonelli
func humanize(name string) string {
	i := strings.Index(name, "_")
	if i < 0 {
		return name
	}

	var buf bytes.Buffer
	buf.WriteString(strings.ToUpper(name[:1]))
	buf.WriteString(name[1:i])
	buf.WriteByte(' ')
	buf.WriteString(strings.ToUpper(name[i+1 : i+2]))
	buf.WriteString(name[i+2:])

	return buf.String()
}

func main() {
	rand.Seed(time.Now().Unix())
	name := namesgenerator.GetRandomName(0)
	fmt.Println(humanize(name))
}
