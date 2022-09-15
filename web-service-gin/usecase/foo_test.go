package usecase

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/repository/mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// TestFooUseCase_Save_NotExistsSuccess 新規保存のテストケース
func TestFooUseCase_Save_NotExistsSuccess(t *testing.T) {
	// 期待値
	testID := int64(123)
	testValue := "hoge"
	expectedSaved := &entity.Foo{
		ID:    testID,
		Value: testValue,
	}

	// モックセット
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fooRepoMock := mock.NewMockFooRepository(ctrl)
	fooRepoMock.EXPECT().FindByID(gomock.Eq(testID)).Return(nil, nil).Times(1)
	fooRepoMock.EXPECT().Insert(gomock.Eq(expectedSaved)).Return(nil)
	fooRepoMock.EXPECT().FindByID(gomock.Eq(testID)).Return(expectedSaved, nil).Times(1)

	// モックを注入して作成
	u := &fooUseCase{fooRepo: fooRepoMock}
	// 実行
	actualSaved, err := u.Save(testID, testValue)
	// 結果確認
	if err != nil {
		t.Errorf("予期せぬエラーが発生 %v", err)
	}
	if !reflect.DeepEqual(actualSaved, expectedSaved) {
		t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", expectedSaved, actualSaved)
	}
}
