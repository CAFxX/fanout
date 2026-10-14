package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }

func main() {
    v1 := []int{897, 788, 32768}
    v2 := []int{-2147483648, 9223372036854775807, 1, 9223372036854775807}
    var v3 string = "a"
    var v4 uint8 = 127
    v3 = "foo bar"
    v5 := GenS1{F0: 42, F1: 1, F2: 0}
    var v6 uint16 = 2
    v7 := GenS1{F0: -972, F1: 990, F2: 65535}
    for fi8 := 0; fi8 < 17; fi8++ {
        if false {
            for fi9 := 0; fi9 < 13; fi9++ {
                v6 -= (v6 % ((32767) | 1))
                _ = fi9
            }
            _ = len(v2)
        }
        _ = fi8
    }
    v3 = (v3 + "foo bar")
    v3 = v3
    v3 = v3
    v6 += 1
    fmt.Println(v6)
    v10 := GenS1{F0: 0, F1: 65535, F2: 65535}
    v11 := GenS1{F0: -1, F1: 255, F2: 2147483648}
    fmt.Println(v6, v4)
    v6 -= 255
    v12 := GenS2{F0: false, F1: 597, F2: true}
    if false {
        var v13 string = v3
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
