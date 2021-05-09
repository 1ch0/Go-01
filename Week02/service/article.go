package service

import (
	"context"

	"github.com/1ch0/Go-01/dao"
	"github.com/1ch0/Go-01/model"
)

func GetArticle(ctx context.Context, id int) (*model.Article, error) {
	return dao.GetArticle(ctx, id)
}
