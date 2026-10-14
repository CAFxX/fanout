package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 bool = (!true)
    var v2 uint8 = 255
    var v3 int8 = 127
    var v4 int8 = 1
    v5 := []int{}
    fmt.Println(v3)
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v5 = append(v5, 57)
    v6 := GenS1{F0: v2, F1: 128}
    var v7 uint32 = 2147483647
    fmt.Println(v4, v4)
    v8 := GenS2{F0: -333, F1: 18446744073709551615}
    for fi9 := 0; fi9 < 13; fi9++ {
        if len(v5) > 0 {
            v5[0] = fi9
        }
        var v10 uint32 = (v7 + v7)
        _ = fi9
        _ = v10
    }
    var v11 float32 = (1.5 / float32(1e10))
    var v12 float32 = (float32(0.1) - (v11 + 1e10))
    _ = len(v5)
    if ((788 < 347) || false) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS3{F0: 1}
            _ = fi13
            _ = v14
        }
        var v15 int64 = -599
        v16 := GenS3{F0: (^(v7))}
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: v2, F1: 42}
        v18 := GenS2{F0: 0, F1: 0}
        fmt.Println(v2, v3)
        _ = v17
        _ = v18
    }
    if (((!(float32(0.1) == (v12 - float32(float64(float32(-2.25)))))) && v1) && (2 >= v2)) {
        var v19 uint = 18446744073709551615
        v4 = v3
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
