package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint16 }
func genFn0(p0 int64) int16 {
    if (316 <= 595) {
        fmt.Println(p0)
        p0 = p0
    }
    v6 := GenS1{F0: 327}
    _ = p0
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{42}
    var v2 int8 = -128
    v3 := make(map[int]bool)
    v4 := []int{4294967295}
    v5 := []string{"foo bar"}
    var v7 float32 = (((1.5 / 1.0) * -2.25) - float32(1e10))
    v8 := GenS2{F0: (((false && ((1e10 - ((0.1 / (-2.25 * (3.14159 * 1.5))) - float64(3.14159))) >= (float64(0.1) * (0.0 / float64(3.14159))))) || false) || true), F1: 1}
    v2 -= (v2 % (((v2 ^ v2)) | 1))
    _ = len(v5)
    v9 := GenS3{F0: 0, F1: 255}
    var v10 uint64 = 18446744073709551615
    var v11 uint64 = v10
    if false {
        v12 := GenS3{F0: 2, F1: 32767}
        if false {
            var v13 int8 = (v2 & v2)
            var v14 bool = (4294967295 < 255)
            v2 = (v13 - (v13 ^ (v2 / (((v2 + (v2 + (-(v13))))) | 1))))
            _ = v13
            _ = v14
        } else {
            var v15 float64 = 1e10
            _ = len(v3)
            _ = v15
        }
        _ = v12
    }
    var v16 int32 = -216
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
