package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 int16 = -32768
    var v3 bool = (!false)
    v4 := []string{}
    var v5 bool = (!(-2147483648 < 9223372036854775807))
    var v6 int16 = (v2 ^ v2)
    switch "" {
    case "a":
        v2 = (v6 / ((v6) | 1))
    }
    fmt.Println(v6)
    fmt.Println(v5, v6, v1)
    var v7 int = -794
    switch "" {
    case "foo bar":
        var v8 string = "hello"
        _ = v8
    case "":
        var v9 uint8 = (^(v1))
        _ = v9
    }
    var v10 float32 = 0.1
    if true {
        v11 := GenS1{F0: (v6 & v6), F1: 42}
        var v12 int64 = -407
        _ = v11
        _ = v12
    } else {
        v6 = v6
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int64 = -9223372036854775808
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
