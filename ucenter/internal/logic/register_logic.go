package logic

import (
	"context"
	"errors"
	"grpc-common/ucenter/types/register"
	"mscoin-common/tools"
	"time"
	"ucenter/internal/domain"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const RegisterCacheKey = "register_code_"

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	CaptchaDomain *domain.CaptchaDomain
	MemberDomain  *domain.MemberDomain
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:           ctx,
		svcCtx:        svcCtx,
		Logger:        logx.WithContext(ctx),
		CaptchaDomain: domain.NewCaptchaDomain(),
		MemberDomain:  domain.NewMemberDomain(svcCtx.DB),
	}
}

func (l *RegisterLogic) RegisterByPhone(in *register.RegisterRequest) (*register.RegisterResponse, error) {
	// 人机验证
	isVerify := l.CaptchaDomain.Verify(in.Captcha.Server, in.Captcha.Token, l.svcCtx.Config.Captcha.Vid, l.svcCtx.Config.Captcha.Key, 2, in.Ip)
	if !isVerify {
		return nil, errors.New("人机校验失败")
	}
	logx.Info("人机校验通过")
	// 校验短信验证码：GetCtx 会把 Redis 里的 JSON 反序列化进 val，因此 val 必须是指针
	var cachedCode string
	if err := l.svcCtx.Cache.GetCtx(l.ctx, RegisterCacheKey+in.Phone, &cachedCode); err != nil {
		return nil, errors.New("验证码已失效或错误")
	}
	if cachedCode != in.Code {
		return nil, errors.New("验证码错误")
	}
	// 注册 验证手机号是否注册
	mem, err := l.MemberDomain.FindByPhone(context.Background(), in.Phone)
	if err != nil {
		return nil, errors.New("数据库链接异常")
	}
	if mem != nil {
		return nil, errors.New("手机号已注册")
	}
	mem, err = l.MemberDomain.Register(context.Background(), in.Phone, in.Password, in.Username, in.Country, in.Promotion, in.SuperPartner)
	return &register.RegisterResponse{}, nil
}

func (l *RegisterLogic) SendCode(req *register.CodeRequest) (*register.NoResponse, error) {
	code := tools.GetCode()
	go func() {
		logx.Info("调用短信平台发送成功")
	}()
	logx.Infof("发送验证码成功：%s \n", code)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := l.svcCtx.Cache.SetWithExpireCtx(ctx, RegisterCacheKey+req.Phone, code, 15*time.Minute)
	if err != nil {
		return nil, errors.New("验证码发送失败")
	}
	return nil, nil
}
