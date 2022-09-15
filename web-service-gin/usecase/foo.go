package usecase

import (
	"example/web-service-gin/entity"
	"example/web-service-gin/repository"
)

type fooUseCase struct {
	fooRepo repository.FooRepository
}

func (p *fooUseCase) Save(id int64, value string) (*entity.Foo, error) {
	foo, err := p.fooRepo.FindByID(id)
	if err != nil {
		// 何らかのエラー
		return nil, err
	}
	if foo != nil {
		// 既存
		foo.Value = value
		if err := p.fooRepo.Update(foo); err != nil {
			return nil, err
		}
	} else {
		// 新規作成
		foo = &entity.Foo{ID: id, Value: value}
		if err := p.fooRepo.Insert(foo); err != nil {
			return nil, err
		}
	}
	// 再取得
	return p.fooRepo.FindByID(id)
}
