package main

import (
	"encoding/binary"
	"face/Db"
	"face/Global"
	"face/Models"
	"face/Route"
	"fmt"
	"github.com/Kagami/go-face"
	"math"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

func main() {
	Global.ImgPath = "./img"
	Global.JWTKey = "jnjhJ9D@2^32jkFgVfOt"
	Global.DB, _ = Db.GetDB()

	//初始化人脸识别器
	FaceRe, err := face.NewRecognizer(modelDir)
	//FaceRe.RecognizeSingle()
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	//////defer rec.Close()

	Global.FaceRe = FaceRe

	db := Global.DB
	var FaceList []Models.UserList
	var samples []face.Descriptor
	db.Find(&FaceList)
	var cats []int32
	for _, faceData := range FaceList {
		//fmt.Println(faceData.Id)
		cats = append(cats, int32(faceData.Id))
		floatData := make([]float32, len(faceData.Face)/4)
		for i := 0; i < len(floatData); i++ {
			bytes := faceData.Face[i*4 : (i+1)*4]
			floatValue := math.Float32frombits(binary.LittleEndian.Uint32(bytes))
			floatData[i] = floatValue
		}
		var descriptor face.Descriptor
		copy(descriptor[:], floatData)
		samples = append(samples, descriptor)
	}
	Global.FaceRe.SetSamples(samples, cats)

	//nayoungFace, err := rec.RecognizeSingleFile(imagesDir + "/wx.jpg")
	//catID := rec.Classify(nayoungFace.Descriptor)
	//fmt.Println(catID)
	//fmt.Println(time.Now())
	////遍历测试
	//dir, err := os.Open("./testdata/images")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer dir.Close()
	//
	//files, err := dir.Readdir(-1)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, file := range files {
	//	fmt.Println("当前处理：" + file.Name())
	//	Name := strings.Split(file.Name(), ".")[0]
	//
	//	//查询数据库
	//	var Count int64
	//	db.Model(&Global.FaceModel).Where("name=?", Name).Count(&Count)
	//	if Count > 0 {
	//		continue
	//	}
	//	faces, err := rec.RecognizeFile(imagesDir + "/" + file.Name())
	//	if err != nil {
	//		log.Fatalf("无法识别: %v", err)
	//		continue
	//	}
	//	// 打印人脸数量
	//	fmt.Println("图片人脸数量: ", len(faces))
	//	if len(faces) > 1 || len(faces) == 0 {
	//		fmt.Println("图片人脸数量大于一个 不对劲")
	//		continue
	//	}
	//	descriptor := faces[0].Descriptor
	//	descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
	//	db.Create(&Models.Face{Name: Name, Data: descriptorBytes})
	//	fmt.Println("生成" + file.Name() + "人脸数据成功")
	//
	//}

	//
	//
	//
	//
	//
	//descriptor := faces[0].Descriptor
	//fmt.Println(descriptor)
	//fmt.Printf("n1的类型是: %T, n1占用的字节数是: %d\n", descriptor, unsafe.Sizeof(descriptor))
	//descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
	//fmt.Println(descriptorBytes)
	Route.InitRouter()
}
