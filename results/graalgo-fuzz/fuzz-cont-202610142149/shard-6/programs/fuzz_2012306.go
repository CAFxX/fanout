package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 string, p1 uint32, p2 string) int16 {
    switch -10 {
    case 32767:
        var v6 uint = 12
        _ = v6
    case -2147483648:
        var v7 int32 = 2147483647
        _ = v7
    }
    fmt.Println(p2)
    fmt.Println(p1, p1, p2)
    p1 -= p1
    v8 := GenS1{F0: -519, F1: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 779
}
func genFn1(p0 int16, p1 uint32, p2 uint) uint8 {
    p1 -= 993
    var v9 uint64 = 878
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 128
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = 0
    v3 := []uint32{2147483647}
    var v4 uint16 = 32767
    var v5 int = -301
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v10 int = v5
    fmt.Println(v5, v2, v5)
    var v11 uint8 = 2
    if (!false) {
        fmt.Println(v5, v1, v4)
        switch -2147483648 {
        case 0:
            if len(v3) > 0 {
                v3[0] = 2
            }
        case -9223372036854775808:
            for fi12 := 0; fi12 < 7; fi12++ {
                if (genFn0("", 42, "a") < v2) {
                    if false {
                        if (1 < 0) {
                            v13 := GenS2{F0: 545, F1: -240}
                            _ = v13
                        }
                        var v14 uint8 = 127
                        _ = v14
                    } else {
                        v15 := GenS3{F0: v11, F1: 0}
                        _ = v15
                    }
                    var v16 float32 = v1
                    _ = v16
                } else {
                    v11 = v11
                }
                v17 := GenS2{F0: 42, F1: int64(v11)}
                _ = fi12
                _ = v17
            }
        default:
            var v18 uint64 = 501
            _ = v18
        }
    }
    var v19 int = (v5 - v10)
    var v20 int = -1
    v5 += (v5 & (v10 * v10))
    v2 = (v2 << ((4294967295) & 63))
    var v21 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
