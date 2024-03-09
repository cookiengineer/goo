package goo

import esbuild "github.com/evanw/esbuild/pkg/api"
import "os"
import "os/user"
import "strings"

func Build(layout Layout, theme Theme, output_folder string) bool {

	var result bool = false

	folder := os.Getenv("GOO_FOLDER")

	if folder == "" {

		home, err1 := os.UserHomeDir()

		if err1 == nil {
			folder = home + "/go/pkg/mod/github.com/cookiengineer/goo"
		} else {

			user, err2 := user.Current()

			if err2 == nil {
				folder = "/home/" + user.Username + "/go/pkg/mod/github.com/cookiengineer/goo"
			}

		}

	}

	// TODO: Find a way to use the embed.FS
	// Maybe with a randomized folder in /tmp that is used for esbuild?

	entrypoint_css := ""
	entrypoint_mjs := ""
	external := make([]string, 0)

	if layout == AppLayout {

		if theme == DefaultTheme {

			entrypoint_css = folder+"/assets/app-default.css"
			entrypoint_mjs = folder+"/assets/app-default.mjs"

			external = []string{
				folder+"/assets/themes/default/cantarell-bold-italic.woff",
				folder+"/assets/themes/default/cantarell-bold-italic.woff2",
				folder+"/assets/themes/default/cantarell-bold.woff",
				folder+"/assets/themes/default/cantarell-bold.woff2",
				folder+"/assets/themes/default/cantarell-italic.woff",
				folder+"/assets/themes/default/cantarell-italic.woff2",
				folder+"/assets/themes/default/cantarell-regular.woff",
				folder+"/assets/themes/default/cantarell-regular.woff2",
				folder+"/assets/themes/default/vera-mono.woff",
			}
		}

	} else {
		// TODO
	}

	if entrypoint_css != "" && entrypoint_mjs != "" {

		esbuild_result := esbuild.Build(esbuild.BuildOptions{
			EntryPoints: []string{entrypoint_css, entrypoint_mjs},
			External:    []string{"*.woff", "*.woff2"},
			Outdir:      "/tmp",
			Bundle:      true,
			Write:       false,
			LogLevel:    esbuild.LogLevelInfo,
		})

		if len(esbuild_result.Errors) == 0 {

			result = true

			for f := 0; f < len(esbuild_result.OutputFiles); f++ {

				file := esbuild_result.OutputFiles[f]
				name := file.Path[strings.LastIndex(file.Path, "/")+1:]

				os.WriteFile(output_folder + "/" + name, file.Contents, 0666)

			}

			for e := 0; e < len(external); e++ {

				name := external[e][strings.LastIndex(external[e], "/")+1:]
				buffer, err := os.ReadFile(external[e])

				if err == nil {
					os.WriteFile(output_folder + "/" + name, buffer, 0666)
				}


			}

		}

	}

	return result

}
