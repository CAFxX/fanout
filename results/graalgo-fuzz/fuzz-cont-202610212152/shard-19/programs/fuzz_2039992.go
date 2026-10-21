package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64; F2 string }
func genFn0(p0 uint64, p1 int32, p2 uint8) int {
    v7 := GenS1{F0: "hello", F1: 923, F2: (^(p2))}
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float32 = 3.14159
        p2 -= p2
        p0 -= p0
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        p2 += 127
        var v11 uint16 = 1
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 14, F1: p1, F2: p2}
    var v13 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v7
    return 1
}
func genFn1(p0 uint32) int16 {
    p0 = p0
    var v14 int8 = -127
    v15 := GenS3{F0: 894, F1: 0, F2: "a"}
    _ = p0
    _ = v14
    _ = v15
    return 42
}

func main() {
    var v1 uint = 65535
    v2 := []int{}
    v3 := make(map[uint32]bool)
    var v4 float64 = ((float64(((0.0 * (3.14159 * 3.14159)) / 0.1)) / 1.0) * float64(1.5))
    var v5 int16 = 1
    v6 := []uint32{479, 42, 346}
    var v16 uint = (v1 & 42)
    v3[1] = false
    v17 := GenS3{F0: 28, F1: 1, F2: ("" + "a")}
    var v18 string = "a"
    v5 -= -364
    fmt.Println(v4, v16, v16)
    v5 = 355
    v16 += 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
