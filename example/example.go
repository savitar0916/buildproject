package example

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

//Data -
type Data struct {
	//FName - 專案名稱
	FName string ""
	//App -
	App string "app"
	//Domain -
	Domain string "domain"
	//Infra -
	Infra string "infra"
	//Interface -
	Interface string "interface"
	//Usecase -
	Usecase string "usecase"
}

func (d *Data) setProject() {
	d.FName = getFName()
	setapp(d)
	//setconfig()
}

func getFName() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, filename := path.Split(pwd)
	return filename
}

func setapp(d *Data) {
	os.Mkdir("app", os.ModePerm)
	setdomain(d)
	setinfra(d)
	setinterface(d)
	setusecase(d)
}

func setconfig() {
	os.Mkdir("config", os.ModePerm)
}

func setdomain(d *Data) {
	//建立資料夾
	Path := fmt.Sprintf("%s/%s", d.App, d.Domain)
	os.MkdirAll(Path, os.ModePerm)
	Model := fmt.Sprintf("%s/%s", Path, "model")
	os.MkdirAll(Model, os.ModePerm)
	Repository := fmt.Sprintf("%s/%s", Path, "repository")
	os.MkdirAll(Repository, os.ModePerm)
	Service := fmt.Sprintf("%s/%s", Path, "service")
	os.MkdirAll(Service, os.ModePerm)
}

func setinfra(d *Data) {
	//建立資料夾
	Path := fmt.Sprintf("%s/%s", d.App, d.Infra)
	os.MkdirAll(Path, os.ModePerm)
	Config := fmt.Sprintf("%s/%s", Path, "config")
	os.MkdirAll(Config, os.ModePerm)
	Logger := fmt.Sprintf("%s/%s", Path, "logger")
	os.MkdirAll(Logger, os.ModePerm)
}

func setinterface(d *Data) {
	//建立資料夾
	Path := fmt.Sprintf("%s/%s", d.App, d.Interface)
	os.MkdirAll(Path, os.ModePerm)
	API := fmt.Sprintf("%s/%s", Path, "api")
	os.MkdirAll(API, os.ModePerm)
	Controller := fmt.Sprintf("%s/%s", Path, "controller")
	os.MkdirAll(Controller, os.ModePerm)
}

func setusecase(d *Data) {
	Path := fmt.Sprintf("%s/%s", d.App, d.Usecase)
	os.MkdirAll(Path, os.ModePerm)
	content := []byte("package " + d.Usecase)
	Interface := fmt.Sprintf("%s/%s", Path, "interface.go")
	Usecase := fmt.Sprintf("%s/%s%s", Path, d.FName, "_usecase.go")
	UsecaseTest := fmt.Sprintf("%s/%s%s", Path, d.FName, "_test.go")
	fmt.Println(d.FName)
	//將指定內容寫入到檔案中
	err := ioutil.WriteFile(Interface, content, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
	err = ioutil.WriteFile(Usecase, content, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
	err = ioutil.WriteFile(UsecaseTest, content, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}
