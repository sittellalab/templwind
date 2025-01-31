package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sittellalab/templwind/internal"
	"github.com/sittellalab/templwind/pkg/twlib"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal(usageText)
	}

	switch args[1] {
	case "dev":
		runDevServer(args[2:])

	case "help", "-h", "-help", "--help":
		fmt.Print(usageText)

	case "version", "-v", "-version", "--version":
		fmt.Println("v0.1.0")

	default:
		log.Fatal(usageText)
	}
}

const usageText = `usage: templwind <command> [<args>...]

templwind - Accessible web components for Go developers

See docs at https://github.com/dimmerz92/templwind

commands:
dev -p [port]           Runs the dev server for templwind component development
help, -h, --help        Prints usage information
version, -v, --version  Prints the version
`

func runDevServer(args []string) {
	cmd := flag.NewFlagSet("dev", flag.ExitOnError)
	port := cmd.Int("p", 8000, "Use a port between 1024 and 65535")
	cmd.Parse(args)

	if *port < 1024 || *port > 65535 {
		log.Fatal("PORT ERROR: Use a port between 1024 and 65535")
	}

	// Styles
	http.HandleFunc("GET /output.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/twthemes/output.css")
	})

	// Index page
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		twlib.Render(w, r, http.StatusOK, internal.Index())
	})

	// Test route for uploading a directory (prints file names with dir structure)
	http.HandleFunc("POST /folder", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(200 << 20)
		files, _ := r.MultipartForm.File["folders"]
		for _, file := range files {
			fmt.Println(twlib.GetStructuredFilename(file))
		}
	})

	// Test route for uploading many files
	http.HandleFunc("POST /files", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(200 << 20)
		files, _ := r.MultipartForm.File["files"]
		for _, file := range files {
			fmt.Println(file.Filename)
		}
	})

	// Test route for a single file
	http.HandleFunc("POST /file", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(200 << 20)
		fmt.Println(r.MultipartForm.File["file"][0].Filename)
	})

	fmt.Printf("Starting dev server on port %d", *port)
	if *port != 8000 {
		fmt.Println("\nEnsure the app_port and proxy_port are updated in .air.toml for server reloading")
		fmt.Println("Navigate to localhost:<proxy_port>, where the proxy_port is the port set in .air.toml")
	} else {
		fmt.Println("\nNavigate to localhost:8001 for hot reloading")
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
