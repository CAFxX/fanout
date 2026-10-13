package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-16}
    v3 := make(map[string]int)
    var v4 uint8 = 59
    var v5 int = 2147483647
    var v6 int32 = 275
    var v7 int32 = (v6 - (v6 & 1))
    switch (^(v5)) {
    case -624:
        switch (("\n\t" + "\n\t") + "\n\t") {
        case "a":
            var v8 uint8 = 255
            _ = v8
        }
    case -475:
        var v9 int16 = 1
        _ = v9
    default:
        var v10 int8 = -95
        _ = v10
    }
    if false {
        var v11 bool = (v5 == v5)
        switch (v5 + (v5 / ((-1) | 1))) {
        case 4294967295:
            fmt.Println(v7, v11, v6)
        case 32768:
            v5 -= -211
        default:
            _ = len(v2)
        }
        _ = v11
    } else {
        var v12 int32 = 1
        v13 := GenS1{F0: -32768, F1: v5}
        switch "\n\t" {
        case "hello":
            for fi14 := 0; fi14 < 11; fi14++ {
                fmt.Println(v4)
                _ = fi14
            }
        }
        _ = v12
        _ = v13
    }
    if (-32768 > 32767) {
        var v15 uint32 = 187
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint64 = 42
        var v18 uint32 = 42
        v5 -= (fi16 + -9223372036854775808)
        _ = fi16
        _ = v17
        _ = v18
    }
    v5 = v5
    var v19 uint8 = 255
    v6 += (v7 % (((^(v7))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
