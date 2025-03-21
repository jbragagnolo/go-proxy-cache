//go:build all || unit
// +build all unit

package jwt

import (
	"testing"

	"github.com/fabiocicerchia/go-proxy-cache/config"
	"github.com/fabiocicerchia/go-proxy-cache/utils/slice"
)

func TestContains(t *testing.T) {
	v := []string{"a", "b"}
	res := slice.ContainsString(v, "a")
	if !res {
		t.Error("Expected true but got", res)
	}

	res = slice.ContainsString(v, "c")
	if res {
		t.Error("Expected false but got", res)
	}
	res = slice.ContainsString(v, "")
	if res {
		t.Error("Expected false but got", res)
	}

	v = []string{}
	res = slice.ContainsString(v, "a")
	if res {
		t.Error("Expected false but got", res)
	}

	res = slice.ContainsString(v, "")
	if res {
		t.Error("Expected false but got", res)
	}

}

func TestIsExcluded(t *testing.T) {
	jwtConfig := &config.Jwt{ExcludedPaths: []string{"/a"}}
	res := IsExcluded(jwtConfig.ExcludedPaths, "/a")
	if !res {
		t.Error("Expected true but got", res)
	}

	res = IsExcluded(jwtConfig.ExcludedPaths, "/b")
	if res {
		t.Error("Expected false but got", res)
	}

	jwtConfig.ExcludedPaths = []string{`^/c/[0-9]+$`}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c`)
	if res {
		t.Error("Expected false but got", res)
	}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c/1234`)
	if !res {
		t.Error("Expected true but got", res)
	}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c/1234/d`)
	if res {
		t.Error("Expected false but got", res)
	}

	jwtConfig.ExcludedPaths = []string{`^/c\/[0-9]+\/d\?.+$`}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c`)
	if res {
		t.Error("Expected false but got", res)
	}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c/1234/d`)
	if res {
		t.Error("Expected false but got", res)
	}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c/1234/f?key1=val1&key2=val2`)
	if res {
		t.Error("Expected false but got", res)
	}
	res = IsExcluded(jwtConfig.ExcludedPaths, `/c/1234/d?key1=val1&key2=val2`)
	if !res {
		t.Error("Expected true but got", res)
	}
}
