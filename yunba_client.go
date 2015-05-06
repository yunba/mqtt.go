package mqtt
import (
    "net/http"
    "bytes"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

const (
    yunba_REG_URL = "http://reg.yunba.io:8383/device/reg/"
    yunba_TICK_URL = "http://tick.yunba.io:9999/"
)

// 服务器返回信息
type YunbaInfo struct {
    ErrCode int `json:"e,omitempty"`
    Client string `json:"c,omitempty"`
    UserName string `json:"u,omitempty"`
    Password string `json:"p,omitempty"`
    DeviceId string `json:"d,omitempty"`
}

type YunbaClient struct {
    Appkey string
    DeviceId string
}

func (this *YunbaClient)httpPostJson(url, jsonStr string) (string, error){
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}

func (this *YunbaClient)Reg() (*YunbaInfo, error){
    jsonStr := fmt.Sprintf(`{"a":"%s","p":4, "d":"%s"}`, this.Appkey, this.DeviceId)


    resp, err := this.httpPostJson(yunba_REG_URL, jsonStr)
    if err != nil {
        return nil, err
    }

    regInfo := &YunbaInfo{}
    err = json.Unmarshal([]byte(resp), regInfo)
    if err != nil {
        return nil, err
    }

    return regInfo, nil
}

func (this *YunbaClient)GetHost() (*YunbaInfo, error){
    jsonStr := fmt.Sprintf(`{"a":%s,"n":1,"v":"v1.0.0","o":1}`, this.Appkey)

    resp, err := this.httpPostJson(yunba_TICK_URL, jsonStr)
    if err != nil {
        return nil, err
    }

    urlInfo := &YunbaInfo{}
    err = json.Unmarshal([]byte(resp), urlInfo)
    if err != nil {
        return nil, err
    }
    return urlInfo, nil
}
