package regexp2

import (
	"reflect"
	"regexp"
	"testing"
)

func TestFindAllString(t *testing.T) {
	re1 := regexp.MustCompile(`a.`)
	re2 := MustCompileStd(`a.`)

	a1 := re1.FindAllString("paranormal", -1)
	a2 := re2.FindAllString("paranormal", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllString("paranormal", 2)
	a2 = re2.FindAllString("paranormal", 2)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllString("graal", -1)
	a2 = re2.FindAllString("graal", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllString("none", -1)
	a2 = re2.FindAllString("none", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}
}

func TestFindAllStringIndex(t *testing.T) {
	re1 := regexp.MustCompile(`ab?`)
	re2 := MustCompileStd(`ab?`)

	a1 := re1.FindAllStringIndex("tablett", -1)
	a2 := re2.FindAllStringIndex("tablett", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringIndex("foo", -1)
	a2 = re2.FindAllStringIndex("foo", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}
}

func TestFindAllStringSubmatch(t *testing.T) {
	re1 := regexp.MustCompile(`a(x*)b`)
	re2 := MustCompileStd(`a(x*)b`)

	a1 := re1.FindAllStringIndex("-ab-", -1)
	a2 := re2.FindAllStringIndex("-ab-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringIndex("-axxb-", -1)
	a2 = re2.FindAllStringIndex("-axxb-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringIndex("-ab-axb-", -1)
	a2 = re2.FindAllStringIndex("-ab-axb-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringIndex("-axxb-ab-", -1)
	a2 = re2.FindAllStringIndex("-axxb-ab-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}
}

func TestFindAllStringSubmatchIndex(t *testing.T) {
	re1 := regexp.MustCompile(`a(x*)b`)
	re2 := MustCompileStd(`a(x*)b`)

	a1 := re1.FindAllStringSubmatchIndex("-ab-", -1)
	a2 := re2.FindAllStringSubmatchIndex("-ab-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringSubmatchIndex("-axxb-", -1)
	a2 = re2.FindAllStringSubmatchIndex("-axxb-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringSubmatchIndex("-ab-axb-", -1)
	a2 = re2.FindAllStringSubmatchIndex("-ab-axb-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringSubmatchIndex("-axxb-ab-", -1)
	a2 = re2.FindAllStringSubmatchIndex("-axxb-ab-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringSubmatchIndex("-foo-", -1)
	a2 = re2.FindAllStringSubmatchIndex("-foo-", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	re1 = regexp.MustCompile(`(?i)(?m)\b(?:(?P<Spelling12>colour)|(?P<Spelling13>color))\b`)
	re2 = MustCompileStd(`(?i)(?m)\b(?:(?<Spelling12>colour)|(?<Spelling13>color))\b`)

	a1 = re1.FindAllStringSubmatchIndex("Red is a color. Blue is also a colour.", -1)
	a2 = re2.FindAllStringSubmatchIndex("Red is a color. Blue is also a colour.", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	a1 = re1.FindAllStringSubmatchIndex("none", -1)
	a2 = re2.FindAllStringSubmatchIndex("none", -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}

	/* FIXME: There's a disconnect between rune/byte offets

	re1 = regexp.MustCompile(`(?m)(?:(?P<Smart3>[”“]))`)
	re2 = MustCompileStd(`(?m)(?:(?P<Smart3>[”“]))`)

	s := `(smart or “curly” quotes).`

	a1 = re1.FindAllStringSubmatchIndex(s, -1)
	a2 = re2.FindAllStringSubmatchIndex(s, -1)

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}*/
}

func TestSubexpNames(t *testing.T) {
	re1 := regexp.MustCompile(`(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)`)
	re2 := MustCompileStd(`(?<first>[a-zA-Z]+) (?<last>[a-zA-Z]+)`)

	a1 := re1.SubexpNames()
	a2 := re2.SubexpNames()

	if !reflect.DeepEqual(a1, a2) {
		t.Fatalf("Failed: %v, %v", a1, a2)
	}
}
