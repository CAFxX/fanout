package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 uint32, p1 int64) uint16 {
    p1 -= p1
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            p0 -= (^(p0))
            if (!("a" != "\n\t")) {
                v7 := GenS1{F0: p1}
                _ = v7
            }
            _ = fi6
        }
        var v8 int32 = -1
        v9 := GenS1{F0: 1}
        _ = fi5
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(p0, fi10)
        _ = fi10
    }
    p1 -= (p1 - (p1 + 42))
    var v11 int8 = 1
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    var v1 uint = 535
    var v2 int8 = -1
    v3 := make(map[string]uint32)
    var v4 uint8 = 89
    _ = len(v3)
    v4 = v4
    v12 := GenS1{F0: -1}
    var v13 uint8 = 2
    v2 = 122
    v14 := GenS2{F0: 65535, F1: 255}
    v15 := []int{-1, -1}
    delete(v3, "")
    v13 -= v4
    _ = len(v3)
    v16 := GenS3{F0: 276, F1: -1}
    v4 += v13
    _ = len(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
