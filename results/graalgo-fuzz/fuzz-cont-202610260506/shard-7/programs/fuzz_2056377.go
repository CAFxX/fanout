package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    var v1 uint64 = 867
    var v2 uint32 = 2
    var v3 uint8 = 128
    fmt.Println(v2, v1, v3)
    v4 := []int{32767, -9223372036854775808, 287}
    v2 -= (^(v2))
    v5 := GenS1{F0: 513}
    switch "" {
    case "hello":
        v2 = (v2 % ((v2) | 1))
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        v3 = 249
        _ = fi6
    }
    _ = len(v4)
    fmt.Println(v2)
    var v7 uint8 = (v3 << ((2) & 63))
    var v8 int8 = 92
    switch 9223372036854775807 {
    case 32768:
        v4 = append(v4, 32767)
    case 1:
        v2 += v2
    case 0:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v9 := GenS2{F0: ("\n\t" + "hello")}
    if true {
        var v10 uint16 = 2
        if (!((false && false) || true)) {
            fmt.Println(v8, v7)
            v11 := GenS3{F0: true, F1: true}
            v12 := GenS2{F0: "hello"}
            _ = v11
            _ = v12
        }
        var v13 uint64 = v1
        _ = v10
        _ = v13
    }
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
