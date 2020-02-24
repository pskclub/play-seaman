package members

import (
	"bitbucket.org/3dsinteractive/seaman"
	"golang-guideline/config"
	"golang-guideline/models"
	"golang-guideline/products"
)

type MemberServiceOptions struct {
	Context        seaman.IContext
	Config         config.IConfig
	ProductService products.IProductService
}

type IMemberService interface {
	GetPage(pageOptions *models.PageOptions, options *GetMemberPageServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError)
	Find(id string) (interface{}, seaman.IError)
}

func NewMemberService(options *MemberServiceOptions) IMemberService {

	return &memberService{
		ctx:            options.Context,
		cfg:            options.Config,
		productService: options.ProductService,
	}
}

type memberService struct {
	ctx            seaman.IContext
	cfg            config.IConfig
	productService products.IProductService
}

func (s memberService) GetPage(pageOptions *models.PageOptions, options *GetMemberPageServiceOptions) ([]interface{}, *models.PaginationMeta, seaman.IError) {
	s.productService.Delete("4343")
	panic("implement me")
}

func (s memberService) Find(id string) (interface{}, seaman.IError) {
	panic("implement me")
}
