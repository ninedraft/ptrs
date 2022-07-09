package ptrs_test

import (
	"ptrs"
	"testing"

	"golang.org/x/exp/slices"
)

func TestOf(test *testing.T) {
	var x = 10
	var p = ptrs.Of(x)
	if *p != x {
		test.Errorf("%v is expected, got %v", x, *p)
	}
}

func TestDeref(test *testing.T) {
	var x = 10

	var got, ok = ptrs.Deref(&x)
	if !ok {
		test.Errorf("ok=true is expected, got %v", ok)
	}
	if got != x {
		test.Errorf("%v is expected, got %v", x, got)
	}
}

func TestDeref_Nil(test *testing.T) {
	var _, ok = ptrs.Deref[int](nil)
	if ok {
		test.Errorf("ok=false is expected, got %v", ok)
	}
}

func TestDerefOr(test *testing.T) {
	var x = 10

	var got = ptrs.DerefOr(&x, 100*x+1)
	if got != x {
		test.Errorf("%v is expected, got %v", x, got)
	}
}

func TestDerefOr_Nil(test *testing.T) {
	var x = 10
	var got = ptrs.DerefOr(nil, 10)
	if got != x {
		test.Errorf("%v is expected, got %v", x, got)
	}
}

func TestNew(test *testing.T) {
	var p = ptrs.New[int]()
	if *p != 0 {
		test.Errorf("0 is expected, got %v", *p)
	}
}

func TestFlatten(test *testing.T) {
	var pp = []*int{ptrs.Of(1), ptrs.Of(2), ptrs.Of(3)}

	var got = ptrs.Flatten(pp)

	var expected = []int{1, 2, 3}
	if !slices.Equal(got, expected) {
		test.Errorf("%v is expected, got %v", expected, got)
	}
}

func TestFlatten_Nil(test *testing.T) {
	var pp = []*int{nil, ptrs.Of(1), nil, ptrs.Of(2), nil, ptrs.Of(3), nil}

	var got = ptrs.Flatten(pp)

	var expected = []int{1, 2, 3}
	if !slices.Equal(got, expected) {
		test.Errorf("%v is expected, got %v", expected, got)
	}
}

func TestRef(test *testing.T) {
	var values = []int{1, 2, 3}

	var got = ptrs.Ref(values)
	for i, ptr := range got {
		if *ptr != values[i] {
			test.Errorf("got[%d] expected to be equal to &%d", i, values[i])
		}
	}
}

func TestRef_EmptyOrNil(test *testing.T) {
	if !slices.Equal(ptrs.Ref([]int{}), []*int{}) {
		test.Errorf("Ref([]) expected to return an empty slice")
	}
	if !slices.Equal(ptrs.Ref[int](nil), []*int{}) {
		test.Errorf("Ref(nil) expected to return an empty slice")
	}
}

func TestMake(test *testing.T) {
	var got = ptrs.Make(4, func(i int, ptr *int) {
		*ptr = i
	})

	var expected = []*int{ptrs.Of(0), ptrs.Of(1), ptrs.Of(2), ptrs.Of(3)}
	if !ptrs.EqualSlice(got, expected) {
		test.Errorf("%v is expected, got %v", expected, got)
	}
}

func TestMake_Nil(test *testing.T) {
	var got = ptrs.Make[int](4, nil)

	var zero = 0
	var expected = []*int{&zero, &zero, &zero, &zero}
	if !ptrs.EqualSlice(got, expected) {
		test.Errorf("%v is expected, got %v", expected, got)
	}
}

func TestEqualSlice_SamePtrs(test *testing.T) {
	var x = 0
	var a = []*int{&x, &x, &x}
	var b = []*int{&x, &x, &x}
	if !ptrs.EqualSlice(a, b) {
		test.Errorf("a expected to be equal to b")
	}
}

func TestEqualSlice_SameValues(test *testing.T) {
	var a = []*int{nil, ptrs.Of(1), nil, ptrs.Of(2), nil, ptrs.Of(3), nil}
	var b = []*int{nil, ptrs.Of(1), nil, ptrs.Of(2), nil, ptrs.Of(3), nil}
	if !ptrs.EqualSlice(a, b) {
		test.Errorf("a expected to be equal to b")
	}
}

func TestEqualSlice_DifferentLengths(test *testing.T) {
	var x = 0
	var a = []*int{&x, &x, &x}
	var b = []*int{&x, &x}
	if ptrs.EqualSlice(a, b) {
		test.Errorf("a expected to be not equal to b")
	}
}

func TestEqualSlice_DifferentValues(test *testing.T) {
	var x = 0
	var a = []*int{&x, &x, &x}
	var b = []*int{&x, &x, ptrs.Of(10*x + 1)}
	if ptrs.EqualSlice(a, b) {
		test.Errorf("a expected to be not equal to b")
	}
}

func TestEqual(test *testing.T) {
	var a = ptrs.Of(1)
	var b = ptrs.Of(1)
	if !ptrs.Equal(a, b) {
		test.Errorf("a is expected to be equal to b")
	}
}

func TestEqual_SamePtrs(test *testing.T) {
	var a = ptrs.Of(1)
	if !ptrs.Equal(a, a) {
		test.Errorf("a is expected to be equal to a")
	}
}

func TestEqual_DifferentValues(test *testing.T) {
	var a = ptrs.Of(1)
	var b = ptrs.Of(2)
	if ptrs.Equal(a, b) {
		test.Errorf("a is expected to be not equal to b")
	}
}

func TestEqual_NilAndValue(test *testing.T) {
	var a = ptrs.Of(1)
	if ptrs.Equal(a, nil) {
		test.Errorf("a is expected to be not equal to b")
	}
}

func TestEqual_ValueAndNil(test *testing.T) {
	var b = ptrs.Of(1)
	if ptrs.Equal(nil, b) {
		test.Errorf("a is expected to be not equal to b")
	}
}

func TestEqual_Nil(test *testing.T) {
	if !ptrs.Equal[int](nil, nil) {
		test.Errorf("nil is expected to be equal to an other nil")
	}
}
