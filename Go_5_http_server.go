package main

// 处理http
//func ListenAndServe(addr string, handler Handler) error
//该方法有两个参数:
// 第一个参数 addr    地 ;
// 第二个参数表示 务 处理程序， 通常为空 ，
// 这意味着 务 调用 http.DefaultServeMux 进行处理，
// 而编写的业务逻 辑处理程序 http.Handle() 或 http.HandleFunc()
// 默认 注入 http.DefaultServeMux 中


// 处理https
//net/http 包还提供 http.ListenAndServeTLS() 方法，用于处理 HTTPS 连接 求:
//func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error
//开  SSL    务也很简单，如下列代码所示:
//http.Handle("/foo", fooHandler)
//http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)) })
//log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)) 或者是:

//ss := &http.Server{
//	Addr: ":10443",
//	Handler:myHandler,
//	ReadTimeout:10 * time.Second,
//	WriteTimeout:10 * time.Second,
//	MaxHeaderBytes: 1 << 20,
//}
//log.Fatal(ss.ListenAndServeTLS("cert.pem", "key.pem"))