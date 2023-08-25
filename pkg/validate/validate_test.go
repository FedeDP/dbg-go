package validate

import (
	"fmt"
	"github.com/falcosecurity/test-infra/images/update-dbg/dbg-go/pkg/autogenerate"
	"github.com/falcosecurity/test-infra/images/update-dbg/dbg-go/pkg/root"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

func generateConfigFile(dkConf autogenerate.DriverkitYaml, confName string) (func(), error) {
	data, err := yaml.Marshal(dkConf)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(confName, data, 0644)
	if err != nil {
		return nil, err
	}
	return func() {
		_ = os.Remove(confName)
	}, nil
}

func TestValidateConfig(t *testing.T) {
	opts := Options{Options: root.Options{
		Architecture:  "x86_64",
		DriverVersion: []string{"1.0.0+driver"},
	}}

	tests := map[string]struct {
		dkConf        autogenerate.DriverkitYaml
		confName      string
		errorExpected bool
	}{
		"correct config": {
			dkConf: autogenerate.DriverkitYaml{
				KernelVersion: "1",
				KernelRelease: "5.10.0",
				Target:        "centos",
				Architecture:  "amd64",
				Output: autogenerate.DriverkitYamlOutputs{
					Module: fmt.Sprintf(autogenerate.OutputPathFmt+".ko",
						opts.DriverVersion,
						opts.Architecture,
						"falco",
						"centos",
						"5.10.0",
						"1"),
					Probe: fmt.Sprintf(autogenerate.OutputPathFmt+".o",
						opts.DriverVersion,
						opts.Architecture,
						"falco",
						"centos",
						"5.10.0",
						"1"),
				},
				KernelUrls:       nil,
				KernelConfigData: "",
			},
			confName:      "centos_5.10.0_1.yaml",
			errorExpected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			cleanup, err := generateConfigFile(test.dkConf, test.confName)
			assert.NoError(t, err)
			t.Cleanup(cleanup)
			err = validateConfig(opts, "1.0.0+driver", test.confName)
			if test.errorExpected {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}