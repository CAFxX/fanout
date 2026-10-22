package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 32768
    var v3 string = ""
    var v4 int64 = 0
    v3 = ""
    if false {
        var v5 uint32 = 915
        _ = v5
    }
    v6 := GenS1{F0: 65535}
    fmt.Println(v2, v3)
    v4 = (^(v4))
    v1[""] = false
    v7 := v1[v3]
    if v7 {
        v7 = ((v7 && false) || (false && (-93 <= 127)))
        v8 := GenS2{F0: -128, F1: 0}
        var v9 float64 = 1e10
        _ = v8
        _ = v9
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
