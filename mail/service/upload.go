package service

import (
	"GOproject/GIT/mail/conf"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadStatic(file multipart.File, userId uint, userName string) (filepath string, err error) {
	bId := strconv.Itoa(int(userId)) //拼接路径
	basepath := "." + conf.AvatarPath + "user" + bId + "/"
	if !DirExistNot(basepath) {
		CreateDir(basepath)
	}
	//avatarpath := basepath + userName + ".jpg"
	//	content, err := io.ReadAll(file)
	//	if err != nil {
	//		return "", err
	//	}
	//	err = os.WriteFile(avatarpath, content, 0777)
	//	if err != nil {
	//		return
	//	}
	//	return "user" + bId + "/" + userName+".jpg", nil
	//}
	img, format, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	avatarpath := basepath + userName + "." + format
	f, err := os.Create(avatarpath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(f, img, nil)
	case "png":
		err = png.Encode(f, img)
	default:
		err = fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + "." + format, nil
}

func UploadProductTolocalStatic(file multipart.File, userId uint, productName string) (filepath string, err error) {
	bId := strconv.Itoa(int(userId)) //拼接路径
	basepath := "." + conf.ProductPath + "boss" + bId + "/"
	if !DirExistNot(basepath) {
		CreateDir(basepath)
	}
	img, format, err := image.Decode(file)
	if err != nil {
		return "", err
	}
	productPath := basepath + productName + "." + format
	f, err := os.Create(productPath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(f, img, nil)
	case "png":
		err = png.Encode(f, img)
	default:
		err = fmt.Errorf("unsupported image format: %s", format)
	}
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + "." + format, nil
}

// 判断是文件路径是否存在
func DirExistNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func CreateDir(dirname string) bool {
	err := os.MkdirAll(dirname, 0777)
	if err != nil {
		return false
	}
	return true
}
