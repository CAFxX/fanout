package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint32) string {
    var v7 uint = 858
    var v8 int = 4294967295
    fmt.Println(p0, v8)
    _ = p0
    _ = v7
    _ = v8
    return (("hello" + "\n\t") + (("\n\t" + "a") + "hello"))
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 0
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    v5 := []uint32{0}
    v6 := make(map[int]int)
    v2 = (^(v2))
    v9 := GenS1{F0: v2, F1: 4294967295}
    v3 += (^(v3))
    var v10 uint64 = 2
    v10 = (v10 ^ v10)
    var v11 int32 = 1
    fmt.Println(v3, v11)
    var v12 int = -9223372036854775808
    var v13 uint32 = 4294967295
    v10 = v10
    v14 := GenS2{F0: v11}
    delete(v6, 127)
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int = 4294967295
        v4[(v13 << ((1) & 63))] = (-(v2))
        v17 := GenS3{F0: 255}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        var v19 float64 = (0.1 + 3.14159)
        if len(v5) > 0 {
            v5[0] = 812
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v12, v2, v2)
    v11 -= -349
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
