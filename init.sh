go mod init goServerSkeleton
go get -u github.com/go-chi/chi/v5
go get -tool github.com/a-h/templ/cmd/templ@latest

#local install of air: so you get code hot reload while developing
go get -tool github.com/air-verse/air@latest

# install daisyui and tailwind. creates a input.css file and moves the output file to 
# ./static/css
curl -sL daisyui.com/fast | bash
mv output.css static/css

#get htmx 4, apline compat and alpine
curl -L https://cdn.jsdelivr.net/npm/htmx.org@4.0.0 -o ./static/js/htmx_4_0_0.js
curl -L https://cdn.jsdelivr.net/npm/htmx.org@4.0.0/dist/ext/hx-alpine-compat.js -o ./static/js/hx-alpine-compat.js
curl -L https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js -o ./static/js/alpine_3_x_x.min.js

