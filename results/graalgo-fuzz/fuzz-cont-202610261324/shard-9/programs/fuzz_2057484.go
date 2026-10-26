package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }

func main() {
    var v1 int64 = 0
    var v2 uint32 = 65535
    var v3 uint8 = 1
    v2 += 65535
    fmt.Println(v3, v2)
    var v4 int8 = -100
    switch ("hello" + ("hello" + "hello")) {
    case "\n\t":
        if true {
            if true {
                v5 := []int{32768}
                _ = v5
            }
            v6 := GenS1{F0: (v3 ^ 177)}
            _ = v6
        }
    default:
        v7 := []int{128, 32768, 2147483647}
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 uint16 = 2
        v10 := GenS2{F0: 32767, F1: (v4 - v4)}
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = 0
    v11 := []int{}
    switch ("" + "\n\t") {
    case "hello":
        var v12 int8 = v4
        _ = v12
    case "a":
        var v13 bool = (754 < 134)
        _ = v13
    case "":
        var v14 int32 = 1
        _ = v14
    default:
        _ = len(v11)
    }
    var v15 uint = 4294967295
    var v16 uint16 = 255
    var v17 int64 = (v1 & (v1 >> (((v15 * (^(v15)))) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
