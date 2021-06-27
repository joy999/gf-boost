package gbbase

import "context"

func (this *Service) SetServiceManage(s *ServiceManage) {
	this.ServiceManage = s
}

func (this *Service) Close() {
	this.ContextManage.Close()
	this.ServiceManage = nil
}

func (this *Service) SetCtx(ctx context.Context) {
	this.ServiceManage.SetCtx(ctx)
	this.TransactionManage.SetCtx(ctx)
}
