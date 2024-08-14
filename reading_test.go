package mplgo

import (
	"image/color"
	"slices"
	"testing"
)

func TestRgbaArrayFromLine(t *testing.T) {
	have, err := rgbaArrayFromLine("[1.01    1.01          0.99 0.99]", 4)
	want := []float64{1.01, 1.01, 0.99, 0.99}

	if !slices.Equal(have, want) || err != nil {
		t.Fatalf("Array from line failed: have %#v, want %#v", have, want)
	}
}

func TestRgbaArrayFromStdout(t *testing.T) {
	have, err := rgbaArrayFromStdout("[[1.02     2.00000000000]\n[0.9992 22]]", 2, 2)
	want := [][]float64{{1.02, 2.0}, {0.9992, 22.0}}

	for i := range have {
		if !slices.Equal(have[i], want[i]) || err != nil {
			t.Fatalf("Array from stdout failed: have %#v, want %#v", have, want)
		}
	}
}

func TestRunPython(t *testing.T) {
	have, dontWant, err := runPython(PY_EXTRACTOR, "viridis", 3)
	want := `[[0.267004 0.004874 0.329415 1.      ]
 [0.127568 0.566949 0.550556 1.      ]
 [0.993248 0.906157 0.143936 1.      ]]
`

	if err != nil {
		t.Fatalf("runPython created a non-nil error")
	}

	if dontWant != "" {
		t.Fatalf("runPython wrote to stdout: %s", dontWant)
	}

	if have != want {
		t.Fatalf("runPython failed, have %s, want %s", have, want)
	}
}

func TestGetCmap(t *testing.T) {
	cmap, err := GetCmapCustom(PY_EXTRACTOR, "viridis", 3, 4)

	if err != nil {
		t.Fatalf("Recieved error from GetCmap")
	}

	want := color.RGBA{68, 1, 84, 255}

	if cmap.name != "viridis" {
		t.Fatalf("Color map name not viridis, got %s", cmap.name)
	}

	if cmap.steps != 3 {
		t.Fatalf("Not correct number of steps in color map, wanted 3 got %d", cmap.steps)
	}

	if cmap.fsteps != 3.0 {
		t.Fatalf("Not correct number of fsteps in color map, wanted 3 got %f", cmap.fsteps)
	}

	if want != cmap.data[0] {
		t.Fatalf("Not correct color extracted from maps: have: %#v, want :%#v", cmap.data[0], want)
	}
}
