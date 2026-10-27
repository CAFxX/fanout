package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 string }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    var v4 uint = 42
    v1 += 187
    v4 = (v4 & v4)
    v5 := []int{384, 2147483647, -113}
    for fi6 := 0; fi6 < 15; fi6++ {
        _ = len(v5)
        var v7 int = fi6
        _ = fi6
        _ = v7
    }
    switch (("hello" + "a") + "") {
    case "a":
        v4 += v4
    default:
        v8 := GenS1{F0: -1, F1: 409, F2: ("hello" + "\n\t")}
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float64 = 3.14159
        _ = fi9
        _ = v10
    }
    var v11 float64 = (float64(-2.25) * 3.14159)
    v12 := GenS2{F0: 465}
    fmt.Println(v11, v1, v1)
    if (!false) {
        var v13 int = 9223372036854775807
        var v14 int8 = 48
        var v15 int16 = 0
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 int8 = 100
        fmt.Println(v11)
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 2
    v3[-1] = 139
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
