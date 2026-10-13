package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    v1 := []uint32{852, 829, 65535, 65535}
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    switch 127 {
    case -1:
        var v4 string = "a"
        _ = v4
    case -932:
        if true {
            v5 := GenS1{F0: 2147483647}
            _ = v5
        } else {
            switch ("" + ("hello" + "foo bar")) {
            case "foo bar":
                if true {
                    var v6 int64 = -1
                    _ = v6
                }
            case "a":
                var v7 uint16 = 608
                _ = v7
            default:
                var v8 int16 = -463
                _ = v8
            }
            var v9 int16 = -145
            if (false && false) {
                v10 := GenS2{F0: 1, F1: 65535, F2: 2}
                _ = v10
            } else {
                v11 := v3[4294967295]
                var v12 int = 127
                _ = v11
                _ = v12
            }
            _ = v9
        }
    }
    var v13 uint16 = 32767
    v14 := GenS1{F0: -1}
    var v15 int = 2147483647
    var v16 float32 = ((0.1 / 3.14159) + 1e10)
    v13 = 255
    v13 -= v13
    if true {
        var v17 uint16 = (v13 >> ((18446744073709551615) & 63))
        _ = v17
    } else {
        v15 = (v15 | (v15 << ((255) & 63)))
        fmt.Println(v16)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
