package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 uint64, p1 uint, p2 string) uint16 {
    p0 = p0
    fmt.Println(p2)
    var v7 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 bool = (-9223372036854775808 > -9223372036854775808)
    v2 := make(map[uint32]string)
    var v3 uint64 = 42
    var v4 float64 = float64(-2.25)
    var v5 float64 = v4
    var v6 uint16 = 2
    var v8 uint = 2
    v1 = (2147483647 != 4294967295)
    v6 += 1
    v6 += (v6 | (^(v6)))
    var v9 int64 = 42
    var v10 uint64 = v3
    v11 := GenS1{F0: 2, F1: 1, F2: (((((1e10 - (0.0 / 3.14159)) / float32(((1.5 - v5) * v5))) / 0.1) == (1e10 / 1.0)) && true)}
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{(fi12 << ((v8) & 63)), fi12, (fi12 * 32767), 0}
        v4 += v4
        var v14 int64 = v9
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := []int{fi15, fi15}
        var v17 uint8 = 42
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{}
    v10 += (v3 | (v10 << (((^(v8))) & 63)))
    var v19 int64 = -9223372036854775808
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 int = 128
        fmt.Println(v10)
        _ = fi20
        _ = v21
    }
    fmt.Println(v6, v10, v10)
    var v22 uint8 = 76
    for fi23 := 0; fi23 < 11; fi23++ {
        v24 := GenS2{F0: v22}
        _ = fi23
        _ = v24
    }
    if true {
        v9 = -403
        var v25 uint16 = (v6 << ((v8) & 63))
        if v1 {
            for fi26 := 0; fi26 < 2; fi26++ {
                v27 := GenS3{F0: v8, F1: 4294967295, F2: (v9 / ((-1) | 1))}
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
                var v28 float32 = -2.25
                _ = fi26
                _ = v27
                _ = v28
            }
            v29 := GenS4{F0: (v22 ^ 1), F1: int(v9), F2: 2147483647}
            _ = v29
        }
        _ = v25
    }
    for fi30 := 0; fi30 < 17; fi30++ {
        v3 += v3
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
