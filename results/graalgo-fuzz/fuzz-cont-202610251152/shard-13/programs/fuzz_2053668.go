package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 uint, p2 uint64) uint64 {
    p1 += (^(p1))
    var v5 string = "a"
    if false {
        v5 = v5
        v6 := GenS1{F0: (false && false), F1: p2, F2: 42}
        if (928 == 65535) {
            var v7 uint64 = 442
            var v8 uint16 = 787
            _ = v7
            _ = v8
        }
        _ = v6
    } else {
        v9 := GenS2{F0: 65535, F1: 100}
        p2 += 233
        v10 := GenS3{F0: -1}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    v4 := make(map[int]int)
    v3 = append(v3, 97)
    var v11 uint16 = 928
    for fi12 := 0; fi12 < 13; fi12++ {
        if false {
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int16 = 0
                _ = fi13
                _ = v14
            }
        } else {
            if (!false) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    fmt.Println(fi12, fi12)
                    v11 = v11
                    _ = fi15
                }
            }
            var v16 int32 = 95
            var v17 float64 = 1e10
            _ = v16
            _ = v17
        }
        v11 -= v11
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 bool = false
            _ = fi18
            _ = v19
        }
        _ = fi12
    }
    var v20 uint64 = 1
    var v21 float64 = 1.5
    v22 := v2[65535]
    v20 = (^(v20))
    for fi23 := 0; fi23 < 2; fi23++ {
        v22 = (v22 - v22)
        _ = fi23
    }
    v21 -= 1.5
    var v24 uint8 = 2
    fmt.Println(v11)
    fmt.Println(v11, v20, v11)
    var v25 int64 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
