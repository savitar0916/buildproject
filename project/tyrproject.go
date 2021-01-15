package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

var (
	//FName - 專案名稱
	FName string
	//App -
	App string = "app"
	//Domain -
	Domain string = "domain"
	//Infra -
	Infra string = "infra"
	//Interface -
	Interface string = "interface"
	//Usecase -
	Usecase string = "usecase"
	//Config -
	Config string = "config"
)

//SetProject -
func SetProject() {
	FName = getFName()
	setapp()
	setconfig()
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

func setapp() {
	os.Mkdir("app", os.ModePerm)
	setdomain()
	setinfra()
	setinterface()
	setusecase()
}

func setconfig() {
	os.Mkdir(Config, os.ModePerm)

	Content := []byte("package config")
	config := fmt.Sprintf("%s/%s", Config, "config.go")
	err := ioutil.WriteFile(config, Content, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
	Data := fmt.Sprintf(`%s%s`, "package config\nimport (\n\"time\"\n)\nvar (\n// ConsulHost -\nConsulHost string\n// APP -\nAPP = \"\"\n// VERSION -\nVERSION string\n// BUILD -\nBUILD string\n// COMMIT -\nCOMMIT string\n// PORT - server port\nPORT string\n", "// ZONE - time zone\nZONE = time.FixedZone(\"\", 8*60*60)\n// ISUSEIP - IP\nISUSEIP bool\n)\nvar (\n// C - \nC = &Config{}\n)\n// Config - 配置\ntype Config struct {\n}\n//api -\nconst (\n)\n")
	envContent := []byte(Data)

	env := fmt.Sprintf("%s/%s", Config, "env.go")
	err = ioutil.WriteFile(env, envContent, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}

func setdomain() {
	//建立資料夾
	Path := fmt.Sprintf("%s/%s", App, Domain)
	os.MkdirAll(Path, os.ModePerm)
	Model := fmt.Sprintf("%s/%s", Path, "model")
	os.MkdirAll(Model, os.ModePerm)
	Repository := fmt.Sprintf("%s/%s", Path, "repository")
	os.MkdirAll(Repository, os.ModePerm)
	Service := fmt.Sprintf("%s/%s", Path, "service")
	os.MkdirAll(Service, os.ModePerm)

	RepoData := []byte("package repository")
	Repo := fmt.Sprintf("%s/%s%s", Repository, FName, "_repository.go")
	err := ioutil.WriteFile(Repo, RepoData, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}

	SerData := []byte("package service")
	SerInter := fmt.Sprintf("%s/%s", Service, "interface.go")
	err = ioutil.WriteFile(SerInter, SerData, os.ModePerm)
	if err != nil {
		fmt.Println("ioutil WriteFile error: ", err)
	}
}

func setinfra() {
	//取得建立的路徑
	Path := fmt.Sprintf("%s/%s", App, Infra)
	os.MkdirAll(Path, os.ModePerm)
	Config := fmt.Sprintf("%s/%s", Path, "config")
	os.MkdirAll(Config, os.ModePerm)
	Logger := fmt.Sprintf("%s/%s", Path, "logger")
	os.MkdirAll(Logger, os.ModePerm)
}

func setinterface() {
	//建立資料夾
	Path := fmt.Sprintf("%s/%s", App, Interface)
	os.MkdirAll(Path, os.ModePerm)
	API := fmt.Sprintf("%s/%s", Path, "api")
	os.MkdirAll(API, os.ModePerm)
	Controller := fmt.Sprintf("%s/%s", Path, "controller")
	os.MkdirAll(Controller, os.ModePerm)
}

func setusecase() {
	Path := fmt.Sprintf("%s/%s", App, Usecase)
	os.MkdirAll(Path, os.ModePerm)
	content := []byte("package " + Usecase)
	Interface := fmt.Sprintf("%s/%s", Path, "interface.go")
	Usecase := fmt.Sprintf("%s/%s%s", Path, FName, "_usecase.go")
	UsecaseTest := fmt.Sprintf("%s/%s%s", Path, FName, "_test.go")
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
