package generate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"proto-service/helper"
)

func GenerateProto(args []string) {
	if len(args) == 0 {
		fmt.Println("args can't empty")
		return
	}

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get the root directory:", err)
		return
	}

	// separate ext from filename
	args = helper.TrimString(args)

	for _, arg := range args {
		cmd := exec.Command(`protoc`, "--proto_path=proto", fmt.Sprintf("proto/%s.proto", arg), "--go_out=.", "--go-grpc_out=.")
		cmd.Dir = rootDir

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(output))
			log.Fatalf("error when checking error  : %s \n", err)
		}

		log.Printf("success generate %s.proto\n", arg)
	}

}
