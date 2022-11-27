package user_test

import (
	"errors"
	"fmt"
	"time"

	"nagase/mocks"
	"nagase/user"

	"testing"

	"github.com/golang/mock/gomock"
)

var now = func() time.Time { return time.Date(1974, time.May, 18, 1, 2, 3, 4, time.UTC) }

func TestUse(t *testing.T) {

	fmt.Println("xxxxxxxqqxxxxxxx")
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	mockCalc := mocks.NewMockCalculator(mockCtrl)
	testUser := &user.User{Doer: mockDoer, Calc: mockCalc}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return("a", nil).Times(1)

	mockCalc.EXPECT().Add(12, 34).Return(234, nil).Times(1)

	testUser.Use()
}

func TestUseReturnsErrorFromDo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	mockCalc := mocks.NewMockCalculator(mockCtrl)
	testUser := &user.User{Doer: mockDoer, Calc: mockCalc}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return dummyError from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return("zz", dummyError).Times(1)
	mockCalc.EXPECT().Add(gomock.Any(), gomock.Any()).Return(234, nil).Times(1)

	result, err := testUser.Use()
	fmt.Println("🐮", result)
	fmt.Println("🐮", err)

	if err != dummyError {
		t.Fail()
	}
}

// func TestUseMatchersExample1(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockDoer := mocks.NewMockDoer(mockCtrl)
// 	testUser := &user.User{Doer: mockDoer}

// 	mockDoer.EXPECT().DoSomething(gomock.Any(), "Hello GoMock")

// 	testUser.Use()
// }

// func TestUseMatchersExample2(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockDoer := mocks.NewMockDoer(mockCtrl)
// 	testUser := &user.User{Doer: mockDoer}

// 	mockDoer.EXPECT().
// 		DoSomething(123, match.OfType("string")).
// 		Return("ssss", nil).
// 		Times(1)

// 	testUser.Use()
// }

// func TestUseOrderExample1(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockDoer := mocks.NewMockDoer(mockCtrl)

// 	callFirst := mockDoer.EXPECT().DoSomething(1, "first this")
// 	mockDoer.EXPECT().DoSomething(2, "then this").After(callFirst)
// 	mockDoer.EXPECT().DoSomething(2, "or this").After(callFirst)

// 	mockDoer.DoSomething(1, "first this")
// 	mockDoer.DoSomething(2, "or this")
// 	mockDoer.DoSomething(2, "then this")
// }

// func TestUseOrderExample2(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockDoer := mocks.NewMockDoer(mockCtrl)

// 	gomock.InOrder(
// 		mockDoer.EXPECT().DoSomething(1, "first this"),
// 		mockDoer.EXPECT().DoSomething(2, "then this"),
// 		mockDoer.EXPECT().DoSomething(3, "then this"),
// 		mockDoer.EXPECT().DoSomething(4, "finally this"),
// 	)

// 	mockDoer.DoSomething(1, "first this")
// 	mockDoer.DoSomething(2, "then this")
// 	mockDoer.DoSomething(3, "then this")
// 	mockDoer.DoSomething(4, "finally this")
// }

// func TestUseActionExamples(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockDoer := mocks.NewMockDoer(mockCtrl)

// 	mockDoer.EXPECT().
// 		DoSomething(gomock.Any(), gomock.Any()).
// 		Return("", nil).
// 		Do(func(x int, y string) {
// 			fmt.Println("Called with x =", x, "and y =", y)
// 		})

// 	mockDoer.EXPECT().
// 		DoSomething(gomock.Any(), gomock.Any()).
// 		Return("", nil).
// 		Do(func(x int, y string) {
// 			if x > len(y) {
// 				t.Fail()
// 			}
// 		})

// 	mockDoer.DoSomething(123, "test")
// 	mockDoer.DoSomething(2, "test")

// }
