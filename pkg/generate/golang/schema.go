package generate

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"
)

// SchemaVersion is the version of the schema used to generate the lock file
const SchemaVersion = 3

// GoPackage is a struct to hold Nix Go module parameters
type GoPackage struct {
	GoPackagePath string `toml:"-"`
	Version       string `toml:"version"`
	Hash          string `toml:"hash"`
	ReplacedPath  string `toml:"replaced,omitempty"`
}

// Output is the output of the generator
type Output struct {
	SchemaVersion int                  `toml:"schema"`
	Mod           map[string]GoPackage `toml:"mod"`

	// Packages with passed import paths trigger `go install` based on this list
	SubPackages []string `toml:"subPackages,omitempty"`

	// Packages with passed import paths has a "default package" which pname & version is inherit from
	GoPackagePath string `toml:"goPackagePath,omitempty"`
}

// Marshal marshals the output to a byte slice
func Marshal(pkgs []GoPackage, goPackagePath string, subPackages []string) ([]byte, error) {
	out := &Output{
		SchemaVersion: SchemaVersion,
		Mod:           make(map[string]GoPackage),
		SubPackages:   subPackages,
		GoPackagePath: goPackagePath,
	}

	for _, pkg := range pkgs {
		out.Mod[pkg.GoPackagePath] = pkg
	}

	var buf bytes.Buffer
	e := toml.NewEncoder(&buf)
	err := e.Encode(out)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ReadCache reads the cache file and returns a map of GoPackages
func ReadCache(filePath string) map[string]GoPackage {
	ret := make(map[string]GoPackage)

	if filePath == "" {
		return ret
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		return ret
	}

	var output Output
	_, err = toml.Decode(string(b), &output)
	if err != nil {
		return ret
	}

	if output.SchemaVersion != SchemaVersion {
		return ret
	}

	for k, v := range output.Mod {
		v.GoPackagePath = k
		ret[k] = v
	}

	return ret
}
