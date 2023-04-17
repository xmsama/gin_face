package Utils

import (
	"encoding/base64"
	"fmt"
	"os"
)

func AddFace(ID int, B64Blob []byte) {
	//B64 = strings.Replace(B64, "data:image/jpeg;base64,", "", 1)
	//B64 = B64[strings.IndexByte(B64, ',')+1:]
	//_, err := base64.StdEncoding.DecodeString(B64)
	//if err != nil {
	//	fmt.Println("B64图片解码失败:", err)
	//	return
	//}
	//db := Global.DB
	//db.Model(&Global.UserListModel).Where("id=  ? ", ID).Updates(map[string]interface{}{"image": Global.ImgPath + "/" + strconv.Itoa(ID) + ".jpg"})
	//faces, err := Global.FaceRe.RecognizeFile(Global.ImgPath + "/" + strconv.Itoa(ID) + ".jpg")
	//if err != nil {
	//	fmt.Println("识别出现错误", err)
	//}
	//descriptor := faces[0].Descriptor
	//descriptorBytes := (*(*[1 << 30]byte)(unsafe.Pointer(&descriptor[0])))[:len(descriptor)*4]
	//db.Model(&Global.UserListModel).Where("id=  ? ", ID).Updates(map[string]interface{}{"face": descriptorBytes})

}

func Blob2B64(Blob []byte) (b64 string) {
	return base64.StdEncoding.EncodeToString(Blob)
}
func ReadImage(path string) (b64 string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 读取文件内容
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将文件内容转换为 Base64 编码
	encoded := base64.StdEncoding.EncodeToString(buffer)
	return "data:image/jpg;base64," + encoded

}
