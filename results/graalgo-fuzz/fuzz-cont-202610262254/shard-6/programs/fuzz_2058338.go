package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    var v1 float32 = 1.5
    var v2 float64 = 3.14159
    var v3 float64 = v2
    v4 := make(map[int]int)
    var v5 string = "hello"
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int16 = 410
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := GenS1{F0: v5}
    v1 = 0.0
    switch -1 {
    case 4294967295:
        var v10 int32 = 567
        _ = v10
    case 326:
        fmt.Println(v5, v3)
    case -9223372036854775808:
        v11 := GenS2{F0: false, F1: 865}
        _ = v11
    default:
        var v12 int32 = 309
        _ = v12
    }
    switch (v5 + "") {
    case "foo bar":
        v1 -= (float32(v3) / v1)
    case "a":
        v3 += (v3 * v2)
    }
    if false {
        var v13 uint32 = 65535
        var v14 string = ("" + v5)
        _ = v13
        _ = v14
    }
    v1 = v1
    var v15 uint64 = 1
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
