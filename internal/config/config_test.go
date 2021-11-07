package config

import (
	"os"
	"testing"
)

func TestMustLoad(t *testing.T) {

	// test: once-ed configuration
	one := MustLoad()
	two := MustLoad()

	if one != two {
		t.Errorf("failed to initialize configuration once and only once")
	}

	// test: ensure default-ness
	expected := "local"

	if one.Environment != expected {
		t.Errorf("failed to default environment variables")
	}

	// test: ensure configurability
	expected = "dev"
	_ = os.Setenv("ENVIRONMENT", expected)
	configured := mustLoad()

	if configured.Environment != expected {
		t.Errorf("failed to yield custom configuraiton via environment variables")
	}

}
