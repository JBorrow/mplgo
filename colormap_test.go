package mplgo

import "testing"

func TestLiveCmapObject(t *testing.T) {
	cmap, err := GetCmap("viridis", 3)

	if err != nil {
		t.Fatalf("Recieved a non-nil error from GetCmap")
	}

	haveMap := cmap.Map(1.0)
	wantMap := cmap.data[len(cmap.data)-1]

	if haveMap != wantMap {
		t.Fatalf("Tried to map 1.0, have %#v and want %#v", haveMap, wantMap)
	}

	haveMap = cmap.Map(0.0)
	wantMap = cmap.data[0]

	if haveMap != wantMap {
		t.Fatalf("Tried to map 0.0, have %#v and want %#v", haveMap, wantMap)
	}

	haveMap = cmap.Map(10.0)
	wantMap = cmap.data[len(cmap.data)-1]

	if haveMap != wantMap {
		t.Fatalf("Tried to map 10.0, have %#v and want %#v", haveMap, wantMap)
	}

	haveMap = cmap.Map(-100000.0)
	wantMap = cmap.data[0]

	if haveMap != wantMap {
		t.Fatalf("Tried to map -100000.0, have %#v and want %#v", haveMap, wantMap)
	}
}
