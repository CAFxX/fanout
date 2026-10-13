package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float64 = (1e10 / 3.14159)
    var v2 int = 32768
    var v3 int32 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    v6 := []uint32{65535}
    v2 = (v2 + v2)
    v2 += (v2 | (v2 ^ (v2 & v2)))
    v1 -= -2.25
    if (!false) {
        v7 := GenS1{F0: 128}
        var v8 int32 = int32(v2)
        fmt.Println(v1, v1)
        _ = v7
        _ = v8
    } else {
        var v9 int16 = -1
        _ = v9
    }
    if (122 < -1) {
        if false {
            fmt.Println(v2, v2, v3)
        }
        var v10 bool = (v3 != v3)
        _ = v10
    }
    fmt.Println(v3, v3, v2)
    fmt.Println(v3, v2)
    switch ("" + (("" + "") + ("foo bar" + "foo bar"))) {
    case "a":
        v1 = 3.14159
    case "":
        if false {
            fmt.Println(v3, v1)
        }
    }
    fmt.Println(v1, v1, v2)
    var v11 uint = 248
    fmt.Println(v3, v2, v11)
    fmt.Println(v1, v11, v11)
    if true {
        var v12 uint32 = 1
        fmt.Println(v12, v3)
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
