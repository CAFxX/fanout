package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }
type GenS4 struct { F0 int16; F1 string; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint8; F2 bool }

func main() {
    v1 := []string{"", "\n\t", "a"}
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    v4 := GenS1{F0: 255}
    v5 := GenS1{F0: 443}
    if false {
        _ = len(v1)
        v6 := GenS2{F0: (!(0 > 0)), F1: 65535, F2: 255}
        var v7 uint = 18446744073709551615
        _ = v6
        _ = v7
    }
    var v8 uint16 = 42
    v8 = (v8 % ((0) | 1))
    v9 := GenS3{F0: v8, F1: 1, F2: 1}
    v10 := GenS4{F0: 32767, F1: "a", F2: v3}
    if true {
        v3 -= (^(v3))
    } else {
        v3 += 2
        if (((true || (2 <= 66)) && (!false)) && false) {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            v11 := GenS5{F0: 1, F1: 0, F2: true}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v11
        }
    }
    if true {
        if (-1 < -990) {
            var v12 int8 = 127
            var v13 uint16 = (v8 * v8)
            _ = v12
            _ = v13
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v14 uint8 = 112
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
