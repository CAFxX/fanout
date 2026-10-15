package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0() int32 {
    var v7 uint16 = 122
    fmt.Println(v7)
    fmt.Println(v7)
    if true {
        v7 += v7
        switch "a" {
        case "a":
            fmt.Println(v7)
        case "":
            v7 += 8
        }
    }
    v7 += v7
    _ = v7
    return 2147483647
}

func main() {
    v1 := []string{("a" + ("" + "\n\t")), (("\n\t" + "a") + "hello")}
    var v2 int64 = 0
    var v3 float32 = (1.5 - float32(0.0))
    var v4 int64 = v2
    var v5 bool = false
    var v6 int = 2147483647
    var v8 uint8 = 104
    var v9 int32 = -1
    fmt.Println(v6, v8)
    var v10 uint16 = 2
    v11 := GenS1{F0: 0, F1: -79}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 = v9
    if v5 {
        v2 += v2
        var v12 int32 = v9
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
