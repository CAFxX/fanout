package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32) uint {
    p0 = p0
    p0 = (p0 & 699)
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 0
    var v3 int32 = 0
    var v4 string = ("\n\t" + "a")
    var v5 uint64 = 4294967295
    v6 := make(map[uint32]string)
    var v7 int32 = -1
    v4 = ("a" + v4)
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 0, F1: int32(v2), F2: 271}
        var v10 float32 = (1e10 / 1e10)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 1, F1: (^(v5))}
    v2 = -946
    v7 -= v3
    fmt.Println(v5)
    v12 := GenS3{F0: 65535, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
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
}
