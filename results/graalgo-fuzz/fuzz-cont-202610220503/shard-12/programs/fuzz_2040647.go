package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
func genFn0(p0 uint, p1 int32, p2 uint) uint16 {
    v5 := GenS1{F0: 150, F1: p1, F2: 0}
    v6 := GenS1{F0: 49, F1: int32(p2), F2: 710}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 93
}
func genFn1(p0 bool, p1 int32, p2 uint) uint64 {
    p2 = (p2 << ((255) & 63))
    p1 = (p1 % (((p1 - 1)) | 1))
    p1 = (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    return 433
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    v4 := []int{0}
    fmt.Println(v1)
    for fi7 := 0; fi7 < 19; fi7++ {
        if (0 < 2147483647) {
            for fi8 := 0; fi8 < 10; fi8++ {
                if (!(42 != 4294967295)) {
                    fmt.Println(fi7, fi8)
                } else {
                    var v9 string = "a"
                    _ = v9
                }
                fmt.Println(fi8, v1)
                _ = fi8
            }
        } else {
            for fi10 := 0; fi10 < 13; fi10++ {
                v4 = append(v4, fi10)
                v1 = (v1 & (v1 ^ v1))
                _ = fi10
            }
            v1 -= v1
        }
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v1, fi11, fi7)
            _ = fi11
        }
        _ = fi7
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float32 = -2.25
        _ = len(v3)
        _ = fi12
        _ = v13
    }
    var v14 float32 = 0.1
    fmt.Println(v14)
    v15 := GenS2{F0: 42}
    v14 = (float32(0.1) * 1.5)
    v1 -= v1
    var v16 uint = 919
    v16 = v16
    var v17 uint64 = uint64(v16)
    var v18 int32 = 0
    var v19 int8 = 42
    delete(v3, 230)
    v20 := GenS3{F0: 42, F1: genFn0(v16, 1, (v16 * v16)), F2: 708}
    var v21 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
