package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", ":8080")
	tmpl := template.Must(template.ParseGlob("templates/*"))
	count := 0
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
			tmpl.ExecuteTemplate(w, "base",count)
	})

	http.HandleFunc("/count-up", func (w http.ResponseWriter, r *http.Request) {
		count = count + 1
		tmpl.ExecuteTemplate(w, "count", count)
	})

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	fmt.Println("running on http://localhost"+os.Getenv(("PORT")))
	http.ListenAndServe("localhost"+os.Getenv("PORT"), nil)
}

// export default async (req, context) => {
//   let todos = context.cookies.get("todos")
//     ? JSON.parse(atob(context.cookies.get("todos")))
//     : [];

//   if (req.method === "GET") {
//     const body = todos.map(makeTodo);
//     return new Response(body, {
//       headers: { "content-type": "text/html" },
//     });
//   }

//   if (req.method === "POST") {
//     const data = await req.formData();

//     if (!data) {
//       return new Response("", {
//         headers: { "content-type": "text/html" },
//       });
//     }
//     const todo = {
//       id: crypto.randomUUID(),
//       title: data.get("new-todo"),
//       done: false,
//     };
//     todos = [...todos, todo];

//     context.cookies.set({
//       name: "todos",
//       value: btoa(JSON.stringify(todos)),
//       httpOnly: true,
//     });

//     const body = makeTodo(todo);
//     return new Response(body, {
//       headers: { "content-type": "text/html" },
//     });
//   }

//   if (req.method === "PATCH") {
//     const data = await req.formData();

//     if (!data) {
//       return new Response("", {
//         headers: { "content-type": "text/html" },
//       });
//     }
//     const todo = todos.find((todo) => todo.id === data.get("id"));

//     todo.done = data.get("done");

//     context.cookies.set({
//       name: "todos",
//       value: btoa(JSON.stringify(todos)),
//       httpOnly: true,
//       sameSite: "strict",
//     });

//     const body = makeTodo(todo);
//     return new Response(body, {
//       headers: { "content-type": "text/html" },
//     });
//   }

//   if (req.method === "DELETE") {
//     const data = await req.formData();

//     if (!data) {
//       return new Response("", {
//         headers: { "content-type": "text/html" },
//       });
//     }

//     todos = todos.filter((todo) => todo.id !== data.get("id"));

//     context.cookies.set({
//       name: "todos",
//       value: btoa(JSON.stringify(todos)),
//       httpOnly: true,
//     });

//     return new Response("", {
//       headers: { "content-type": "text/html" },
//     });
//   }
// };
