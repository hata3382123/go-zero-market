package domain

import (
	"context"
	"errors"
	"mscoin-common/msdb"
	"ucenter/internal/dao"
	"ucenter/internal/model"
	"ucenter/internal/repo"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberDomain struct {
	MemberRepo repo.MemberRepo
}

func NewMemberDomain(db *msdb.MsDB) *MemberDomain {
	return &MemberDomain{
		MemberRepo: dao.NewMemberDao(db),
	}
}

func (d *MemberDomain) FindByPhone(ctx context.Context, phone string) (*model.Member, error) {
	mem, err := d.MemberRepo.FindByPhone(ctx, phone)
	if err != nil {
		logx.Info(err)
		return nil, errors.New("数据库链接异常")
	}
	return mem, nil
}

func (d *MemberDomain) Register(ctx context.Context, phone string, password string, username string, country string, promotion string, partner string) error {
	panic("implement me")
}
