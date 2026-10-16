package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 float32 = float32(float64(1e10))
    var v2 int16 = 1000
    v3 := []string{"", ("\n\t" + "")}
    v4 := make(map[string]string)
    if (782 < 2) {
        v5 := GenS1{F0: 65535, F1: "hello"}
        var v6 uint8 = 23
        _ = v5
        _ = v6
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 int64 = 1
    switch -1 {
    case -294:
        v7 = v7
    }
    fmt.Println(v7, v2)
    var v8 int32 = -195
    delete(v4, "\n\t")
    v9 := GenS2{F0: "", F1: 0}
    var v10 bool = (2147483648 == 768)
    delete(v4, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
