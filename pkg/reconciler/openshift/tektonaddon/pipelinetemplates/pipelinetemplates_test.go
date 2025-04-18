package tektonaddon

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	mf "github.com/manifestival/manifestival"
	"gotest.tools/v3/golden"
)

func TestGeneratePipelineTemplates(t *testing.T) {
	addonLocation := filepath.Join("testdata")

	manifest := mf.Manifest{}

	err := GeneratePipelineTemplates(addonLocation, &manifest)
	assertNoError(t, err)
	for _, m := range manifest.Resources() {
		jsonPipeline, err := m.MarshalJSON()
		assertNoError(t, err)
		golden.Assert(t, string(jsonPipeline), strings.ReplaceAll(fmt.Sprintf("%s.golden", m.GetName()), "/", "-"))
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("assertion failed; expected no error %v", err)
	}
}
