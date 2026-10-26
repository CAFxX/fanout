package main

import (
    "fmt"
)


func main() {
    var v1 float64 = -2.25
    var v2 bool = (1 < 1)
    var v3 bool = (2 >= 1)
    v4 := make(map[int]string)
    v3 = false
    var v5 int64 = 0
    fmt.Println(v5)
    var v6 string = (("" + "a") + "foo bar")
    v3 = v2
    switch -1 {
    case 128:
        v2 = false
    }
    v5 += (v5 % (((v5 - -1)) | 1))
    var v7 int8 = 42
    v4[4294967295] = ("a" + "a")
    var v8 float32 = -2.25
    var v9 float64 = v1
    fmt.Println(v5, v7, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
