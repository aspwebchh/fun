package main

import (
	"net/http"
	"fmt"
	"strconv"
	"io/ioutil"
	"strings"
	"errors"
	"crypto/md5"
	"encoding/hex"
	"regexp"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Trim(str string) string {
	regex, _ := regexp.Compile("(^\\s*)|(\\s*$)")
	return regex.ReplaceAllString(str, "")
}

func Round(val float64, len int) float64 {
	tmp := fmt.Sprintf("%." + strconv.Itoa(len) + "f", val)
	result, _ := strconv.ParseFloat(tmp, 64);
	return result;
}

func ParseFloat(str string) float64 {
	val, err := strconv.ParseFloat(str, 64);
	if err != nil {
		return 0;
	} else {
		return val;
	}
}


func Float2BigInt( fVal float64 ) int64 {
	var intString = fmt.Sprintf("%." + strconv.Itoa(0) + "f", fVal)
	var result,err = strconv.ParseInt(intString,10,64)
	if err == nil {
		return result
	} else {
		return 0
	}
}

func FormatSeconds( seconds float64 ) string {
	var iSeconds = Float2BigInt(seconds)
	var hour = iSeconds / 3600
	iSeconds = iSeconds % 3600
	var minute = iSeconds / 60
	iSeconds = iSeconds % 60
	var resultString = ""
	if hour > 0  {
		resultString += strconv.FormatInt(hour,10) + "小时 "
	}
	if minute > 0  {
		resultString += strconv.FormatInt(minute,10) + "分钟 "
	}
	resultString += strconv.FormatInt(iSeconds,10) + "秒"
	return resultString
}

func Interface2Error(data interface{}) error {
	resultError, ok := data.(error)
	if !ok {
		str, ok := data.(string)
		if !ok {
			return errors.New("未知错误")
		}
		return errors.New(str)
	}
	return resultError
}

func HttpPost(url string) (result string, resultErr error) {
	defer func() {
		err := recover()
		if err != nil {
			result = ""
			resultErr = Interface2Error(err)
		}
	}()
	var urlItems = strings.Split(url, "?");
	var params = "";
	if len(urlItems) == 2 {
		params = urlItems[1]
	}
	client := &http.Client{}
	//client.Timeout = time.Second * 10
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err;
	}
	if resp.StatusCode >= 400 || resp.StatusCode < 200 {
		return "", fmt.Errorf("请求出错， 状态码：" + strconv.Itoa(resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil;
}

func HttpGet(url string) (result string, resultErr error) {
	defer func() {
		err := recover()
		if err != nil {
			result = ""
			resultErr = Interface2Error(err)
		}
	}()
	client := &http.Client{};
	//client.Timeout = time.Second * 5
	request, err := http.NewRequest("GET", url, nil);
	if err != nil {
		return "", err;
	}
	response, err := client.Do(request);
	defer response.Body.Close();
	if err != nil {
		return "", err;
	}
	if response.StatusCode >= 400 || response.StatusCode < 200 {
		return "", fmt.Errorf("请求出错， 状态码：" + strconv.Itoa(response.StatusCode))
	}
	content, err := ioutil.ReadAll(response.Body);
	if err != nil {
		return "", err;
	}
	return string(content), nil;
}

func HttpHead(url string) (resultErr error) {
	defer func() {
		err := recover()
		if err != nil {
			resultErr = Interface2Error(err)
		}
	}()
	response, err := http.Head(url)
	if err != nil {
		return err;
	}
	defer response.Body.Close()
	if response.StatusCode >= 400 || response.StatusCode < 200 {
		return fmt.Errorf("请求出错， 状态码：" + strconv.Itoa(response.StatusCode))
	}
	return nil
}