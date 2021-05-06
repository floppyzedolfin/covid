package server

import (
	"testing"

	"github.com/gofiber/fiber"
	"github.com/stretchr/testify/assert"
)

func TestServer_Load(t *testing.T) {
	tt := map[string]struct {
		file string
		loadedData map[string][]entry
	} {
		"2 entries": {
			file: "file://testdata/data.csv",
			loadedData: map[string][]entry{},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := New()
			gouvFile = tc.file
			s.load(&fiber.Ctx{})
			assert.Equal(t, tc.loadedData, s.data)
		})
	}
}

