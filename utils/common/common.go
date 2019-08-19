package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var BaseDir string

func init(){
	var err error
	currentFile :=  CurrentFile()
	BaseDir, err = filepath.Abs(path.Dir(path.Dir(path.Dir(currentFile))))
	if err != nil{
		log.Fatal(err)
	}
}

func Strftime(t time.Time) string{
	return t.String()[:19]
}


func Strptime(t string, format string) (time.Time, error){
	structT, err := time.Parse(format, t)
	return structT, err
}


// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}


// 转换查询参数为int
func QueryInt(query string, defaultNum int) int{
	// var num int
	var queryNum int

	if defaultNum < 1{
		defaultNum = 1
	}

	if query == ""{
		queryNum = defaultNum
	}else{
		queryNum, err := strconv.Atoi(query)
		if err != nil{
			queryNum = defaultNum
		}
		if queryNum <= 0{
			queryNum = defaultNum
		}else{
			return queryNum
		}
	}

	return queryNum
}

// 生成随机字符串
func  GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


// map or struct to json
func MapJson(m interface{}) (string, error){
	bytes, err := json.Marshal(m)
	if err != nil{
		return "", err
	}
	return string(bytes), nil

}

// check path is exists
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return true
}


//check path is file
func IsFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return !f.IsDir()
}

// check path is dir
func IsDir(path string) bool{
	f, err := os.Stat(path)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return f.IsDir()
}

//check permission
func NotPermission(path string) bool{
	_, err := os.Stat(path)
	return os.IsPermission(err)
}


func Mkdir(path string)error{
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil{
		return err
	}
	return nil
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}
