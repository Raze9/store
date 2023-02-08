package service

import (
	"GOproject/GIT/mail/dao"
	"GOproject/GIT/mail/model"
	"GOproject/GIT/mail/pkg/e"
	"GOproject/GIT/mail/pkg/util"
	"GOproject/GIT/mail/serializer"
	"context"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id"form:"id"`
	Name          string `json:"name"form:"name"`
	CategoryId    int    `json:"category_id"form:"category_id"`
	Title         string `json:"title"form:"title"`
	Info          string `json:"info"form:"info"`
	ImgPath       string `json:"img_path"form:"img_path"`
	Price         string `json:"price"form:"price"`
	DiscountPrice string `json:"discount_price"form:"discount_price"`
	OnSale        bool   `json:"onSale"form:"onSale"`
	Num           int    `json:"num"form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetuserbyId(uId)
	tmp, _ := files[0].Open()
	path, err := UploadProductTolocalStatic(tmp, uId, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("boos creat err", err)
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    uint(service.CategoryId),
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        uId,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)

	err = productDao.CreatProduct(product)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status:  code,
			Message: e.GetMsg(code),
			Error:   err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productimgdao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductTolocalStatic(tmp, uId, service.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
				Error:   err.Error(),
			}
		}
		productimg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productimgdao.CreatProductImg(productimg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status:  code,
				Message: e.GetMsg(code),
				Error:   err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status:  code,
		Message: e.GetMsg(code),
		Data:    serializer.BuildProduct(product),
	}
}
