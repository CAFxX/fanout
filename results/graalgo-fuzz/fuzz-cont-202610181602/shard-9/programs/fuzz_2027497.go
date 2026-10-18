package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 float64 = 0.0
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 69, F1: 2, F2: 32767}
    fmt.Println(v2)
    fmt.Println(v1)
    v2 = v2
    switch 32768 {
    case 32767:
        v5 := GenS2{F0: ("a" + ("hello" + "a"))}
        _ = v5
    case 128:
        v1 -= v1
    default:
        v1 = 3.14159
    }
    v6 := GenS2{F0: "a"}
    v1 -= v1
    v2 -= 42
    fmt.Println(v1, v1)
    if true {
        delete(v3, 2147483647)
        if (!false) {
            v1 += (v1 / v1)
        }
    }
    if (!(!(4294967295 <= 2))) {
        var v7 string = "\n\t"
        _ = v7
    } else {
        v1 += (3.14159 + v1)
        if false {
            fmt.Println(v1, v2)
            fmt.Println(v1)
        }
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    v9 := GenS1{F0: 100, F1: 994, F2: 2}
    v2 = (v2 ^ 1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
