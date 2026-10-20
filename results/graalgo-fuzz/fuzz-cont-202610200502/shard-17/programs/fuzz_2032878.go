package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{"", "hello"}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    _ = len(v1)
    v4 := GenS1{F0: -33}
    var v5 uint = 255
    var v6 uint64 = 940
    var v7 float32 = 3.14159
    var v8 int16 = 32767
    v8 = (v8 - int16(v6))
    delete(v2, 65535)
    v9 := GenS1{F0: -2147483648}
    switch ("\n\t" + "foo bar") {
    case "hello":
        v10 := GenS2{F0: v5, F1: uint8(v5)}
        _ = v10
    case "":
        switch -9223372036854775808 {
        case 937:
            var v11 int16 = v8
            _ = v11
        }
    }
    fmt.Println(v6)
    if true {
        v12 := GenS3{F0: 1}
        _ = v12
    } else {
        fmt.Println(v5)
        v6 -= 411
    }
    if false {
        switch 39 {
        case -416:
            fmt.Println(v7, v6, v5)
        default:
            v3["foo bar"] = -877
        }
    }
    var v13 int16 = v8
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
