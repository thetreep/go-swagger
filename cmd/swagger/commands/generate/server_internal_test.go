package generate

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetreep/go-swagger/generator"
)

func TestDeprecated(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stdout)
	s := Server{
		WithContext: true,
	}
	s.apply(new(generator.GenOpts))
	assert.Contains(t, buf.String(), "warning")
}
