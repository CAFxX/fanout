package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64; F1 string; F2 int64 }
type GenS4 struct { F0 string }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 746
    var v3 int64 = v1
    var v4 uint8 = 85
    v4 = 1
    var v5 uint64 = 543
    if ((float64(3.14159) < 1.5) || true) {
        v6 := GenS1{F0: 255, F1: -2147483648, F2: 42}
        fmt.Println(v5, v3, v2)
        v2 += v2
        _ = v6
    }
    var v7 uint8 = v4
    v8 := []int{-942, -726}
    v7 = v7
    v7 -= v4
    var v9 int16 = 32767
    fmt.Println(v3, v4, v1)
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS2{F0: 0, F1: uint32(v4)}
        var v12 int16 = (v9 / (((v9 % ((v9) | 1))) | 1))
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    _ = len(v8)
    switch "\n\t" {
    case "a":
        fmt.Println(v4, v9, v7)
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    if (32768 <= 127) {
        switch "\n\t" {
        case "foo bar":
            v13 := GenS1{F0: 1, F1: 1, F2: 42}
            _ = v13
        case "a":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "":
            v14 := GenS2{F0: 876, F1: 4294967295}
            _ = v14
        default:
            v7 = 1
        }
        if ((^(v3)) != 1) {
            if true {
                var v15 int = 32768
                v16 := GenS3{F0: (^(v3)), F1: "", F2: v3}
                _ = v15
                _ = v16
            } else {
                v17 := GenS4{F0: "hello"}
                fmt.Println(v3)
                _ = v17
            }
            v9 = v9
            fmt.Println(v5, v1)
        }
        fmt.Println(v3)
    }
    v7 = (^(v7))
    var v18 uint16 = 477
    var v19 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
