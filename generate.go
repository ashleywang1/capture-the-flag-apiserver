package main

import (
	"github.com/solo-io/skv2/codegen"
	"github.com/solo-io/skv2/codegen/model"
	"github.com/solo-io/skv2/codegen/skv2_anyvendor"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
)

func main() {
	log.Println("starting generate")
	apiRoot := "pkg/api"

	skv2Cmd := codegen.Command{
		AppName:      "awang-apiserver",
		RenderProtos: true,
		AnyVendorConfig: &skv2_anyvendor.Imports{
			Local:    []string{
				"api/*.proto",
			},
			External: map[string][]string{
				"github.com/solo-io/protoc-gen-ext":           []string{
					"extproto/*.proto",
					"external/google/**/*.proto",
				},
			},
		},
		Groups: []model.Group{
			model.Group{
				GroupVersion: schema.GroupVersion{
					Group:   "awang.solo.io",
					Version: "v1",
				},
				Module:           "github.com/ashleywang1/capture-the-flag-apiserver",
				Resources:        nil,
				RenderManifests:  false,
				RenderTypes:      false,
				RenderClients:    false,
				RenderController: false,
				MockgenDirective: false,
				ApiRoot:          apiRoot,
			},
		},
	}
	if err := skv2Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
