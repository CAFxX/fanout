package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    var v1 int64 = -333
    v2 := []int{0, -274, 127, -991}
    var v3 uint16 = 2
    if (2147483648 != 0) {
        fmt.Println(v3, v3)
    }
    v1 -= (v1 - v1)
    v1 -= v1
    if (18446744073709551615 != 18446744073709551615) {
        v1 = v1
    }
    switch 411 {
    case 2147483647:
        var v4 uint16 = (^(v3))
        _ = v4
    default:
        var v5 uint8 = 1
        _ = v5
    }
    if (0 > -1) {
        v6 := GenS1{F0: 42, F1: -1}
        v7 := GenS1{F0: 1000, F1: -931}
        var v8 uint32 = 4294967295
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 1, F1: 1}
    v1 -= v1
    if len(v2) > 0 {
        v2[0] = 32767
    }
    fmt.Println(v1, v3)
    v10 := GenS1{F0: 42, F1: 4294967295}
    fmt.Println(v1, v1)
    if true {
        if (128 > -787) {
            v3 = (v3 - (^(v3)))
        }
        _ = len(v2)
    }
    var v11 uint = 0
    if (!(!false)) {
        v11 = v11
    }
    var v12 int = 0
    fmt.Println(v11, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
