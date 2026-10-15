package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int32, p2 uint32) string {
    v5 := GenS1{F0: 4294967295, F1: "hello"}
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(fi6, p0, p0)
        if true {
            p0 += p0
            p1 -= p1
        }
        _ = fi6
    }
    v7 := GenS2{F0: 470, F1: 788, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return ""
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]string)
    v3 := []uint32{848, 674, 1}
    var v4 uint64 = 2
    v4 += 1
    v4 = (v4 << ((1) & 63))
    v4 = v4
    var v8 int8 = -19
    var v9 uint16 = 351
    v10 := GenS3{F0: 0}
    var v11 uint16 = 991
    v11 += (^(v11))
    v4 -= (v4 << ((616) & 63))
    fmt.Println(v9)
    var v12 float64 = (-2.25 * 0.1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
