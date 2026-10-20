package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int16) uint8 {
    p0 += p0
    v7 := GenS1{F0: 0, F1: 1, F2: false}
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 uint16 = 222
    v2 := make(map[string]bool)
    var v3 uint64 = 1
    v4 := make(map[int]uint32)
    var v5 int32 = -1
    var v6 int16 = 68
    fmt.Println(v5)
    v3 += v3
    var v8 int64 = -714
    delete(v4, 32768)
    v9 := GenS2{F0: v5, F1: true}
    v8 += v8
    v10 := GenS3{F0: 128}
    var v11 int = 711
    var v12 uint16 = v1
    for fi13 := 0; fi13 < 15; fi13++ {
        fmt.Println(v11)
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
