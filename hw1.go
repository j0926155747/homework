package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)
type Co struct
{
	Temp_c float64
	Icon_url string
}

type Message struct
{
	Current_observation Co
}

func main() {

	str:="https://goo.gl/mSzFnk"
	
	resp, err := http.Get(str)  
    if err != nil {
		fmt.Println("網址讀取錯誤: ",err.Error());
        return
    }

    input, err1 := ioutil.ReadAll(resp.Body) 
	resp.Body.Close()    
    if err1 != nil {
		fmt.Println("讀成byte檔錯誤: ",err1.Error())
        return
    }
	
	
	var m Message
	err2 := json.Unmarshal(input, &m)
	if err2 != nil {
        fmt.Println("json error: ", err.Error())
        return
    }
	fmt.Println(m.Current_observation.Temp_c)
	fmt.Println(m.Current_observation.Icon_url)
	return
}
