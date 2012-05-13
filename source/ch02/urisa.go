package urisa

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
    "time"
    "strconv"
    "encoding/base64"
    "encoding/hex"
    "crypto/md5"
    
    "appengine"
    "appengine/urlfetch"
    "appengine/datastore"
)

func init() {
    http.HandleFunc("/", help)
    http.HandleFunc("/chk", chk)
    http.HandleFunc("/qchk", qchk)
}

func help(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, usageHelp)
}
const usageHelp = `
URIsA ~ KSC 4 GAE powdered by go1
{v12.05.4}
usage:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/chk
or with GAE Datastore quick resp. if ahd checked:
    $ curl -d "uri=http://sina.com" urisago1.appsp0t.com/qchk
`

var APPKEY  = "k-60666"
var SECRET  = "99fc9fdbc6761f7d898ad25762407373"
var APIHOST = "open.pc120.com"
var APITYPE = "/phish/"
var PHISHID = map[int] string {
    -1:   "未知",
    0:    "好站",
    1:    "钓鱼!",
    2:    "也许...",
}
type KSC struct {
    Success int    //`json:"success"`
    Phish   int //`json:"phish"`
    Msg     string //`json:"msg"`
}
func _genKSCuri(url string) string {
    println("url len~\t ", len(url))
    maxEncLen := base64.URLEncoding.EncodedLen(len([]byte(url))) 
    println("maxEncLen~\t ", maxEncLen)
    //dst := make([]byte, 256) //<~ 整来的代码,不理解,就一定会出问题...
    dst := make([]byte, maxEncLen) //<~ 整来的代码,不理解,就一定会出问题...
    base64.URLEncoding.Encode(dst, []byte(url))
    println("base64~\t", string(dst))

    args := "appkey=" + APPKEY
    args += "&q=" + string(dst)
    now := time.Now()
    nano := strconv.FormatInt(now.UnixNano(),10)
    //c.Infof("timestamp ~ %v.%v", nano[0:10],nano[10:13])

    args += "&timestamp=" + nano[0:10] + "." + nano[10:13]
    sign_base_string := APITYPE + "?" + args 
    println("sign_base_string~\t ", sign_base_string)
    //md5 hash 严格参数顺序:: appkey -> q -> timestamp
    h := md5.New()
    io.WriteString(h, sign_base_string + SECRET)
    args += "&sign=" + hex.EncodeToString(h.Sum(nil))
    println("sign~\t ", hex.EncodeToString(h.Sum(nil)))
    println("args~\t ", args)
    api_url := "http://"+ APIHOST + APITYPE + "?" + args 
    println("api_url~ ", api_url)

    return api_url
}

func _asKSC(uri string, r *http.Request) (int, int) {
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get(uri)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
    }
    c.Infof("HTTP GET returned status %v", resp.Status)
    if resp.StatusCode != 200 {
        panic(err)
        c.Infof("couldn't get sale data %v", http.StatusInternalServerError)
        //http.Error(w, "couldn't get sale data", http.StatusInternalServerError)
        //return
    }
    defer resp.Body.Close()
    c.Infof("resp.ContentLength %v", resp.ContentLength)
    var buf []byte
    buf, _ = ioutil.ReadAll(resp.Body)
    c.Infof("resp.Body %v", string(buf))

    result := &KSC{}
    err = json.Unmarshal(buf, result)
    if err != nil {
        panic(err)
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
    }
    
    return result.Success, result.Phish
}

type Chked struct {
    Uri     string
    Phish    int
    Tstamp  time.Time
    Cip     string
}
func qchk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    ukey := datastore.NewKey(c, "Uri", url, 0, nil)
    var e2 Chked
    if err := datastore.Get(c, ukey, &e2); err != nil {
        fmt.Fprint(w, "~ ", err.Error() , "\n")
        //panic(err)
        c.Infof("Get Err.~\n\t !!! %v", err.Error())
        //http.Error(w, err.Error(), http.StatusInternalServerError)
        //return
        api_url := _genKSCuri(url)
        _,p := _asKSC(api_url, r)
        e1 := Chked{
            Uri:    url,
            Phish:  p,
            Tstamp: time.Now(),
            Cip:    r.RemoteAddr,
        }
        //key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "chked", nil), &e1)
        ukey, err = datastore.Put(c, ukey, &e1)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //c.Infof("datastore key=%v", key.String())
        fmt.Fprint(w, "ask KCS API srv.\n")
        fmt.Fprint(w, "/qchk(KCS):\t" + PHISHID[p])
    }else{
        c.Infof("datastore ukey=%v", ukey.String())
        fmt.Fprint(w, "datastore Get OK;-) \n")
        c.Infof("co4 datastore:%v \t Phish:%s", e2.Phish ,PHISHID[e2.Phish])
        fmt.Fprint(w, "/qchk(GAE):\t" + PHISHID[e2.Phish])
    }
}

func chk(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    url := r.FormValue("uri")
    
    api_url := _genKSCuri(url)

    s,p := _asKSC(api_url, r)
    c.Infof("Success:%v \t Phish:%s", s ,PHISHID[p])

    fmt.Fprint(w, "/chk(KCS):\t" + PHISHID[p])
}
