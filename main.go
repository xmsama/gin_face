package main

import (
	"face/Db"
	"face/Global"
	"face/Models"
	"fmt"
	"github.com/Kagami/go-face"
	"log"
	"os"
	"strings"
	"unsafe"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

func main() {
	//初始化人脸识别器
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()

	//干一个数据库
	Global.DB, _ = Db.GetDB()
	db := Global.DB
	//遍历测试
	dir, err := os.Open("./testdata/images")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		fmt.Println("当前处理：" + file.Name())
		Name := strings.Split(file.Name(), ".")[0]

		//查询数据库
		var Count int64
		db.Model(&Global.FaceModel).Where("name=?", Name).Count(&Count)
		if Count > 0 {
			continue
		}
		faces, err := rec.RecognizeFile(imagesDir + "/" + file.Name())
		if err != nil {
			log.Fatalf("无法识别: %v", err)
			continue
		}
		// 打印人脸数量
		fmt.Println("图片人脸数量: ", len(faces))
		if len(faces) > 0 || len(faces) == 0 {
			fmt.Println("图片人脸数量大于一个 不对劲")
			continue
		}
		descriptor := faces[0].Descriptor
		descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
		db.Create(&Models.Face{Name: Name, Data: descriptorBytes})
		fmt.Println("生成" + file.Name() + "人脸数据成功")

	}
	//
	//// 调用该方法，传入路径。返回面部数量和任何错误
	//faces, err := rec.RecognizeFile(imagesDir + "/bona.jpg")
	//if err != nil {
	//	log.Fatalf("无法识别: %v", err)
	//}
	//// 打印人脸数量
	//fmt.Println("图片人脸数量: ", len(faces))
	//if len(faces) > 1 {
	//	fmt.Println("图片人脸数量大于一个 不对劲")
	//}
	//
	//
	//
	//
	//descriptor := faces[0].Descriptor
	//fmt.Println(descriptor)
	//fmt.Printf("n1的类型是: %T, n1占用的字节数是: %d\n", descriptor, unsafe.Sizeof(descriptor))
	//descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
	//fmt.Println(descriptorBytes)

}
