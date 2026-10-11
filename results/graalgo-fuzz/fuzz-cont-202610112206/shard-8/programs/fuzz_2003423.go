package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0(p0 int32, p1 int16) int16 {
    v7 := []int{-31, 2147483647}
    var v8 float64 = 1e10
    var v9 int = 4294967295
    p0 -= 2147483647
    var v10 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 495
}

func main() {
    var v1 uint16 = 777
    var v2 uint64 = 2
    var v3 int8 = 1
    var v4 float32 = 0.1
    v5 := make(map[string]uint32)
    v6 := make(map[string]int)
    v1 += (v1 - v1)
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{fi11, (fi11 % (((fi11 - 1)) | 1))}
        switch (fi11 & fi11) {
        case 127:
            var v13 int64 = -1
            _ = v13
        case -2147483648:
            v3 = int8(v1)
        default:
            var v14 uint32 = uint32(v2)
            _ = v14
        }
        if false {
            switch "" {
            case "hello":
                v4 -= ((1.5 - v4) - (v4 * float32(-2.25)))
            case "\n\t":
                v15 := GenS1{F0: genFn0(-1, 958), F1: int8(v2)}
                _ = v15
            case "a":
                fmt.Println(v4, v4, fi11)
            }
        } else {
            fmt.Println(v2, fi11, v4)
            var v16 bool = (true || true)
            v17 := GenS2{F0: 43}
            _ = v16
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v2, v2)
    var v18 int32 = -1
    var v19 float32 = (float32(1.5) * v4)
    v20 := []int{0}
    switch "a" {
    case "":
        v21 := GenS3{F0: 1, F1: -944}
        _ = v21
    case "\n\t":
        v22 := GenS4{F0: 1, F1: 0, F2: v1}
        _ = v22
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            fmt.Println(fi23, v2)
            _ = fi23
        }
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
