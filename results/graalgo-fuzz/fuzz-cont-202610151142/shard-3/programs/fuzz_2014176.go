package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 uint16 = 513
    var v5 int64 = -290
    var v6 uint16 = (v4 >> ((366) & 63))
    var v7 int16 = -114
    v8 := GenS1{F0: (v7 + 0), F1: 1, F2: v4}
    for fi9 := 0; fi9 < 11; fi9++ {
        v6 -= (v6 | v6)
        _ = fi9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            v5 = v5
        } else {
            _ = len(v1)
            v3 += (v3 / ((v3) | 1))
        }
        _ = fi10
    }
    var v11 string = ""
    v12 := GenS2{F0: (!(-1 < -128)), F1: v5, F2: uint8(v4)}
    var v13 int = -354
    v13 = (v13 * (v13 * (v13 * (v13 - -2147483648))))
    v14 := GenS1{F0: (v7 ^ v7), F1: v13, F2: v6}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
