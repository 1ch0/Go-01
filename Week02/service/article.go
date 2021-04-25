package service

import (
	"context"

	"github.com/1ch0/Go-01/Week02/dao"
	"github.com/1ch0/Go-01/Week02/model"
)

func GetArticle(ctx context.Context, id int) (*model.Article, error) {
	return dao.GetArticle(ctx, id)
}
