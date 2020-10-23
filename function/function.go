package function

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

func ToMd5(string string) string {
	data := []byte(string)
	b := md5.Sum(data)
	pass := hex.EncodeToString(b[:])
	return pass
}

func ToSHA256(string string) string {
	if string != "" {
		data := []byte(string)
		b := sha256.Sum256(data)
		pass := hex.EncodeToString(b[:])
		return pass
	} else {
		return string
	}
}

func UploadAvatar(file multipart.File, filename string) string {
	dir := "/usr/share/nginx/html/CDN/avatar/ams"
	tempFile, err := ioutil.TempFile(dir, filename+"-*.jpg")
	if err != nil {
		return "No Image"
	}
	defer tempFile.Close()
	name := strings.Replace(tempFile.Name(), "/usr/share/nginx/html/CDN/avatar/ams/", "", 1)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	tempFile.Write(fileBytes)

	return name
}

func Uploadbytype(file multipart.File, filename string, bytype string, group string) string {
	var dir string
	if bytype == "person" {
		if group == "" {
			dir = "/usr/share/nginx/html/CDN/person"
		} else {
			dir = "/usr/share/nginx/html/CDN/person/" + group
		}
	} else if bytype == "organization" {
		if group == "" {
			dir = "/usr/share/nginx/html/CDN/organization"
		} else {
			dir = "/usr/share/nginx/html/CDN/organization/" + group
		}
	} else if bytype == "location" {
		if group == "" {
			dir = "/usr/share/nginx/html/CDN/location"
		} else {
			dir = "/usr/share/nginx/html/CDN/location/" + group
		}
	} else if bytype == "categories" {
		if group == "" {
			dir = "/usr/share/nginx/html/CDN/categories"
		} else {
			dir = "/usr/share/nginx/html/CDN/categories/" + group
		}
	}

	tempFile, err := ioutil.TempFile(dir, filename+"-*.jpg")
	if err != nil {
		return "error"
	}
	defer tempFile.Close()

	replaced_dir := "/usr/share/nginx/html/CDN"
	name := strings.Replace(tempFile.Name(), replaced_dir, "", 1)
	if replaced_dir+name != tempFile.Name() {
		os.Remove(tempFile.Name())
		return "error"
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		os.Remove(tempFile.Name())
		return "error"
	}
	tempFile.Write(fileBytes)

	return name
}

func CreateGroupDir(bytype string, group string) {
	if bytype == "categories" {
		dir := "/usr/share/nginx/html/CDN/categories/" + group
		os.Mkdir(dir, os.ModePerm)
	} else if bytype == "person" {
		dir := "/usr/share/nginx/html/CDN/person/" + group
		os.Mkdir(dir, os.ModePerm)
	} else if bytype == "organization" {
		dir := "/usr/share/nginx/html/CDN/organization/" + group
		os.Mkdir(dir, os.ModePerm)
	} else if bytype == "location" {
		dir := "/usr/share/nginx/html/CDN/location/" + group
		os.Mkdir(dir, os.ModePerm)
	}
}

func RemoveImage(dir string) {
	path := "/usr/share/nginx/html/CDN"
	err := os.Remove(path + dir)

	if err != nil {
		log.Println(err)
	}
}

func RemoveGroupDir(bytype string, group string) {
	if bytype == "categories" {
		dir := "/usr/share/nginx/html/CDN/categories/" + group
		os.RemoveAll(dir)
	} else if bytype == "person" {
		dir := "/usr/share/nginx/html/CDN/person/" + group
		os.RemoveAll(dir)
	} else if bytype == "organization" {
		dir := "/usr/share/nginx/html/CDN/organization/" + group
		os.RemoveAll(dir)
	} else if bytype == "location" {
		dir := "/usr/share/nginx/html/CDN/location/" + group
		os.RemoveAll(dir)
	}
}
