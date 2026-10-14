package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int16) int8 {
    if false {
        v4 := GenS1{F0: 128}
        _ = v4
    }
    if (82 > 1) {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 int = (^(fi5))
            _ = fi5
            _ = v6
        }
    }
    p0 = p0
    p0 -= (-(p0))
    _ = p0
    return 20
}

func main() {
    var v1 bool = (1.5 < 0.0)
    v2 := make(map[string]uint32)
    var v3 string = "hello"
    v7 := GenS2{F0: 774, F1: 0}
    v8 := []int{-1}
    fmt.Println(v1)
    v9 := v2[v3]
    switch 127 {
    case 1:
        delete(v2, "hello")
    default:
        switch 42 {
        case -140:
            switch 587 {
            case 32767:
                if (0 != -86) {
                    var v10 float32 = ((float32(0.1) / 1e10) - (float32(3.14159) + 1.5))
                    v10 = float32((-2.25 + 1.5))
                    _ = v10
                } else {
                    v3 = v3
                    v1 = true
                }
            case -294:
                var v11 float64 = 1.5
                _ = v11
            default:
                v12 := GenS3{F0: -9223372036854775808}
                _ = v12
            }
        case 127:
            for fi13 := 0; fi13 < 13; fi13++ {
                v1 = (!v1)
                var v14 uint = 18446744073709551615
                _ = fi13
                _ = v14
            }
        case 128:
            var v15 int32 = 1
            _ = v15
        }
    }
    var v16 int16 = 1000
    v17 := GenS4{F0: 160}
    fmt.Println(v9, v3)
    fmt.Println(v16, v3)
    v9 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
