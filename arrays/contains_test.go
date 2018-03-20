package arrays

import "testing"

// TestArraycontains
func TestContains(t *testing.T) {
	type test struct {
		data interface{}
		key  interface{}
		want bool
	}
	var testNil []interface{}
	testees := []test{
		{
			data: []int{1, 2, 3, 4, 5},
			key:  1,
			want: true,
		},
		{
			data: []int{1, 2, 3, 4, 5},
			key:  6,
			want: false,
		},
		{
			data: []int{1, 2, 3, 4, 5},
			key:  "aaa",
			want: false,
		},
		{
			data: []string{"aaa", "bbb", "ccc"},
			key:  "bbb",
			want: true,
		},
		{
			data: []string{"aaa", "bbb", "ccc"},
			key:  "ddd",
			want: false,
		},
		{
			data: []string{"aaa", "bbb", "ccc"},
			key:  123,
			want: false,
		},
		{
			data: []bool{true, false},
			key:  true,
			want: false,
		},
		{
			data: testNil,
			key:  true,
			want: false,
		},
	}
	for _, testee := range testees {
		actual := Contains(testee.data, testee.key)
		if actual != testee.want {
			t.Errorf("Response assert failed : actual %t, expect %t", actual, testee.want)
		}
	}
}
