package main

import (
	"io"
	"net/http"
	"os/exec"
	"strings"
)

var list = map[string]string{
	"+":     "amixer set Master playback 5%+",
	"-":     "amixer set Master playback 5%-",
	"next":  "foobnix --next",
	"prev":  "foobnix --prev",
	"pause": "foobnix --pause",
	"play":  "foobnix --play",
}

func get_cmd(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `<html>
		<head>
			<style type='text/css'>
			body {
				zoom: 200%;
			}
			</style>
			<title>Remote control</title>
		</head> 
	<body>
		<form method="POST" action="">
			Enter command: <br>
			<input type="text" autofocus name="cmd" placeholder="">
			<button>Send</button>
		</form>
	</body>
	</html>`)

	run(strings.TrimSpace(r.FormValue("cmd")))
}

func run(cmd string) {
	if len(cmd) != 0 {
		if value, ok := list[cmd]; ok {
			execute(value)
			print("Executed: " + value + "\n")
		}
	}
}

func execute(cmd string) {
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]
	out, _ := exec.Command(head, parts...).Output()
	print(string(out))
}

func main() {
	print("RemoteServer started \n")
	http.HandleFunc("/linux", get_cmd)
	http.ListenAndServe(":1337", nil)
}
