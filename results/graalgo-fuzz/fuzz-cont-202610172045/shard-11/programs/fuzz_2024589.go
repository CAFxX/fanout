package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }

func main() {
    var v1 string = "hello"
    var v2 bool = (0 != 557)
    var v3 uint8 = 0
    v4 := []int{-73, -653, -2147483648, -571}
    v5 := make(map[uint32]string)
    var v6 uint64 = uint64(v3)
    var v7 int8 = -1
    for fi8 := 0; fi8 < 8; fi8++ {
        switch 9223372036854775807 {
        case 9223372036854775807:
            v2 = v2
        case 922:
            for fi9 := 0; fi9 < 10; fi9++ {
                switch -422 {
                case 32767:
                    v1 = v1
                }
                _ = fi9
            }
        case 32767:
            fmt.Println(v7)
        }
        var v10 uint32 = 1
        v6 -= (v6 | 979)
        _ = fi8
        _ = v10
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS1{F0: -672, F1: 870, F2: (v7 ^ v7)}
        _ = fi11
        _ = v12
    }
    v3 = (v3 % (((v3 + 0)) | 1))
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 string = v1
        switch fi13 {
        case 32767:
            for fi15 := 0; fi15 < 5; fi15++ {
                v2 = (((v3 - v3) == (^(v3))) || false)
                _ = fi15
            }
        }
        v2 = false
        _ = fi13
        _ = v14
    }
    switch v1 {
    case "\n\t":
        var v16 uint = 18446744073709551615
        _ = v16
    }
    v17 := GenS1{F0: -1, F1: -1, F2: (v7 + v7)}
    var v18 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
