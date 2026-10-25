package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }

func main() {
    var v1 int8 = -100
    var v2 uint64 = 8
    v3 := []uint32{2147483647}
    v4 := make(map[uint32]int)
    var v5 int64 = -119
    v2 += (^(v2))
    var v6 string = ("hello" + "a")
    _ = len(v4)
    if true {
        v7 := GenS1{F0: 4294967295, F1: 1, F2: 1}
        for fi8 := 0; fi8 < 15; fi8++ {
            v3 = append(v3, 4294967295)
            v5 = (v5 | v5)
            _ = fi8
        }
        fmt.Println(v2)
        _ = v7
    } else {
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(v1, v5, v6)
            var v10 int64 = v5
            _ = fi9
            _ = v10
        }
        v6 = ""
    }
    if false {
        for fi11 := 0; fi11 < 20; fi11++ {
            if len(v3) > 0 {
                v3[0] = 42
            }
            var v12 uint8 = 255
            for fi13 := 0; fi13 < 4; fi13++ {
                fmt.Println(fi11, fi11, v2)
                v3 = append(v3, 1)
                v14 := GenS2{F0: v6, F1: 42, F2: 42}
                var v15 int16 = 42
                v6 = (v6 + "foo bar")
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi11
            _ = v12
        }
    } else {
        _ = len(v3)
    }
    v2 = v2
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS3{F0: v2, F1: 42, F2: v1}
        _ = fi16
        _ = v17
    }
    v1 -= -128
    var v18 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
