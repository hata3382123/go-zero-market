package dao

import (
	"context"
	"mscoin-common/msdb"
	"mscoin-common/msdb/gorms"
	"ucenter/internal/model"

	"gorm.io/gorm"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m *MemberDao) FindByPhone(ctx context.Context, phone string) (mem *model.Member, err error) {
	session := m.conn.Session(ctx)
	err = session.Model(&model.Member{}).Where("mobile_phone=?", phone).Limit(1).Find(&mem).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return mem, err
}

func NewMemberDao(db *msdb.MsDB) *MemberDao {
	return &MemberDao{conn: gorms.New(db.Conn)}
}
