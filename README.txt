READ AND FOLLOW THESE STEPS!

Steps to get started and get a fully functional webserver on localhost: 127.0.0.1:8080

1. Make sure the init.sh file has permission to execute:
    a. In Linux: ls -al
    b. if file is not executable: .rwxr-xr-x run command in terminal: chmod +x init.sh
2. Make sure the watch_tailwindcss.sh ahs permission to run
    a. In Linux: ls -al
    b. if file is not executable: .rwxr-xr-x run command in terminal: chmod +x watch_tailwindcss.sh
3. Run the script init.sh in terminal: ./init.sh
   The script initializes a go-project and install dependencies (chi-router and templ templating library, daisyUI, tailwindcsss) and moves the output.css file to ./static/css folder. The script also download HTMX 4 and alpine.js and puts it in the ./static/js folder 
4. Run the tailwind watcher script in a separate terminal: ./watch_tailwindcss.sh. (The script watches for changes in source files affecting css and recreates output.css)
5. Add into input.css (so the watcher will update output.css when template files uses css classes). Put these lines close to other @source lines: 
    @source "./templates/*.templ";
    @source "./templates/*.go";
6. (Optional) Add Tailwind css typography plugin in input.css. Can be put in the end of the file:
    @plugin "@tailwindcss/typography";
7. Write your backend code in server.go file.

All HTML templates are in the ./templates folder
Before the server can be run, make sure all HTML templates are compiled and go-files are created.

1. cd templates
2. run the command: go tool templ generate

Run the above command whenever you have changed and saved a .templ file

IF you wan hot reload of code while developing run in project root:

    go tool air
    
    (The command uses .air.toml file to activate hot reload of code)
ELSE run in project root:

    go run .

(If css does not get updated try ctrl+F5 to force reload)

Open browser and enter address:

    http://localhost:8080/