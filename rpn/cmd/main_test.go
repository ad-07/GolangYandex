package main

// import (
// 	"testing"
// )

// func TestNewWorld(t *testing.T) {
// 	type test struct {
// 		height  int
// 		width   int
// 		wantErr bool
// 	}

// 	tests := []test{
// 		{height: 0, width: 4, wantErr: true},
// 		{height: -1, width: 4, wantErr: true},
// 		{height: 5, width: 0, wantErr: true},
// 		{height: 5, width: 6, wantErr: false},
// 	}

// 	for _, tt := range tests {
// 		height := tt.height
// 		width := tt.width
// 		world, err := NewWorld(height, width)
// 		if err != nil {
// 			if tt.wantErr {
// 				continue
// 			}
// 			t.Errorf("Unexpected error: %s", err)
// 		}

// 		if world.Height != height {
// 			t.Errorf("expected height: %d, actual height: %d", height, world.Height)
// 		}
// 		if world.Width != width {
// 			t.Errorf("expected width: %d, actual width: %d", width, world.Width)
// 		}

// 		if len(world.Cells) != height {
// 			t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
// 		}
// 	}
// }

// import (
// 	"testing"
// )

// func TestCalc(t *testing.T) {
// 	testCasesSuccess := []struct {
// 		name           string
// 		expression     string
// 		expectedResult float64
// 	}{
// 		{
// 			name:           "simple",
// 			expression:     "1+1",
// 			expectedResult: 2,
// 		},
// 		{
// 			name:           "priority",
// 			expression:     "(2+2)*2",
// 			expectedResult: 8,
// 		},
// 		{
// 			name:           "priority",
// 			expression:     "2+2*2",
// 			expectedResult: 6,
// 		},
// 		{
// 			name:           "/",
// 			expression:     "1/2",
// 			expectedResult: 0.5,
// 		},
// 		// {
// 		// 	name:       "simple",
// 		// 	expression: "1+1*",
// 		// 	expectedResult: ,
// 		// }
// 	}

// 	for _, testCase := range testCasesSuccess {
// 		t.Run(testCase.name, func(t *testing.T) {
// 			val, err := Calc(testCase.expression)
// 			if err != nil {
// 				t.Fatalf("successful case %s returns error", testCase.expression)
// 			}
// 			if val != testCase.expectedResult {
// 				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
// 			}
// 		})
// 	}
// }