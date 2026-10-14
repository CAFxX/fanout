package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint32 = 884
    var v2 float32 = 0.0
    var v3 float32 = v2
    v4 := make(map[uint32]string)
    v5 := make(map[int]bool)
    v6 := []string{}
    v7 := GenS1{F0: (v1 - v1), F1: 2}
    v2 = float32(((3.14159 / 3.14159) * 3.14159))
    v1 -= v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v8 uint64 = 267
    _ = len(v4)
    if false {
        v2 = (v2 / 3.14159)
        fmt.Println(v3, v2, v2)
        switch 128 {
        case 32767:
            switch ("\n\t" + "foo bar") {
            case "hello":
                v9 := GenS1{F0: v1, F1: 222}
                _ = v9
            case "foo bar":
                v10 := GenS1{F0: v1, F1: 255}
                _ = v10
            }
        case -2147483648:
            v3 += v2
        case -9223372036854775808:
            v11 := GenS2{F0: 0, F1: true, F2: 42}
            _ = v11
        }
    }
    v1 -= 329
    fmt.Println(v3)
    _ = len(v4)
    var v12 uint16 = 2
    if false {
        v2 -= (float32((0.1 + 0.1)) + v2)
    } else {
        _ = len(v5)
        fmt.Println(v3, v1, v1)
    }
    if false {
        v13 := GenS3{F0: (^(v12)), F1: (v1 & v1)}
        _ = v13
    } else {
        if true {
            v8 = 1
            v12 = v12
        }
    }
    if false {
        switch "hello" {
        case "hello":
            fmt.Println(v12)
        case "a":
            var v14 string = "a"
            _ = v14
        case "\n\t":
            fmt.Println(v1)
        default:
            if (v8 > (^(v8))) {
                v15 := GenS4{F0: -1}
                _ = v15
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
