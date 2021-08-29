package main

import (
	"log"
	"net"
	"os"

	"template/logger"

	// ${module_name}/app/${domain_name}/controller -> ${domain_name_initial}c
	tc "template/app/template/controller"
	// ${module_name}/app/${domain_name}/gateway -> ${domain_name_initial}g
	tg "template/app/template/gateway"
	// ${module_name}/app/${domain_name}/usecase -> ${domain_name_initial}u
	tu "template/app/template/usecase"
	// protocol buffer ${domain_name} -> pb${domain_name}
	pbTemplate "template/proto/template"

	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "19003"
	}

	listenPort, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	server := grpc.NewServer()

	// 依存性を解決 TODO 別ファイル管理??
	templateGateway := tg.NewTemplateGateway()

	templateUsecase := tu.NewTemplateUsecase(templateGateway)

	templateController := tc.NewTemplateController(templateUsecase)

	// サーバーへ登録
	pbTemplate.RegisterTemplateServiceServer(server, templateController)

	logger.StartLog()

	server.Serve(listenPort)
}
