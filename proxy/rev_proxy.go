package main
import (
  "log"
  "net/http"
  "net/http/httputil"
  "net/url"
  "github.com/gorilla/mux"
)

var (
  hostProxy = make(map[string]string)
  proxies = make(map[string]*httputil.ReverseProxy)
)

func init(){
  hostProxy["machine0.com"] = "http://192.168.0.135:10080"
  hostProxy["machine1.com"] = "http://192.168.0.135:20080"
  //create a struct and url parse the domain
/* func init()(w http.ResponseWriter, r *http.Request){
  domain["url1.com"=r.URL.Query().Get("url")]
      sheenanigan it
      for k,v := range domain{
      remote,err := url.Parse(v)//log the err
      proxies[k] = httputil.NewSingleHostReverseProxy(remote)
    }
}  */
  for k,v := range hostProxy {
    remote, err := url.Parse(v)
    if err != nil {
      log.Fatal("Unable to pass proxy target")
    }
    proxies[k] = httputil.NewSingleHostReverseProxy(remote)
  }
}

func main(){
  rout := mux.NewRouter()
  for host,proxy := range proxies {
    rout.Host(host).Handler(proxy)
  }
  log.Fatal((http.ListenAndServe(":80",rout)))
}
