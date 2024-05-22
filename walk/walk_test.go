package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	tcs := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{"struct with 1 string field", struct{ Name string }{"Megan"}, []string{"Megan"}},
		{"struct with 2 string fields", struct{ Name, City string }{"Megan", "Orlando"}, []string{"Megan", "Orlando"}},
		{
			"struct with 1 string and 1 int field",
			struct {
				Name string
				Age  uint8
			}{"Megan", 30},
			[]string{"Megan"},
		},
		{
			"nested fields",
			Person{
				"Megan",
				Profile{30, "Orlando"},
			},
			[]string{"Megan", "Orlando"},
		},
		{
			"pointers to things",
			&Person{
				"Megan",
				Profile{30, "Orlando"},
			},
			[]string{"Megan", "Orlando"},
		},
		{
			"slices",
			[]Profile{
				{33, "Orlando"},
				{35, "Tampa"},
			},
			[]string{"Orlando", "Tampa"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Orlando"},
				{35, "Tampa"},
			},
			[]string{"Orlando", "Tampa"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			var got []string
			walk(tc.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, tc.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, tc.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
