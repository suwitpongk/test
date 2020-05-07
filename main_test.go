package main

import "testing"

// func TestMock(t *testing.T) {
// 	assert := assert.New(t)

// 	testObj := new(MyMockedObject)

// 	testObj.On("concat", "1", "2").Return("123", nil)

// 	test, _ := testObj.concat("1", "2")

// 	assert.Equal(test, "123", "they should be equal")
// }

func Test_concat(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{"a", "b"}, "ab", false},
		{"case2", args{"a", "bb"}, "abb", false},
		{"case3", args{"", "bb"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := concat(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("concat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
