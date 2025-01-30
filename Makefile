.SILENT:

server:
	air

tailwind:
	tailwindcss -i pkg/twthemes/input.css -o pkg/twthemes/output.css --watch

deps:
	go install github.com/air-verse/air@latest && \
		go install github.com/a-h/templ/cmd/templ@latest && \
		go mod tidy

dev:
	make -j2 tailwind server
