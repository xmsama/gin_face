package main

import (
	"fmt"
	"log"
	"time"
	"unsafe"

	"github.com/Kagami/go-face"
)

const dataDir = "testdata"

// testdata 目录下两个对应的文件夹目录
const (
	modelDir  = dataDir + "/models"
	imagesDir = dataDir + "/images"
)

func main() {
	fmt.Println("Face Recognition...")

	// 初始化识别器
	rec, err := face.NewRecognizer(modelDir)
	if err != nil {
		fmt.Println("Cannot INItialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")
	fmt.Println(time.Now())
	// 调用该方法，传入路径。返回面部数量和任何错误

	faces, err := rec.RecognizeFile(imagesDir + "/bona.jpg")
	if err != nil {
		log.Fatalf("无法识别: %v", err)
	}
	// 打印人脸数量
	fmt.Println(time.Now())
	fmt.Println("图片人脸数量: ", len(faces))

	descriptor := faces[0].Descriptor
	fmt.Println(descriptor)
	fmt.Printf("n1的类型是: %T, n1占用的字节数是: %d\n", descriptor, unsafe.Sizeof(descriptor))
	descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
	fmt.Println(descriptorBytes)
}
