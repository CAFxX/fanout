package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 int64; F1 int8; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 4294967295
    v3 := make(map[uint32]int)
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := GenS1{F0: 0, F1: 9223372036854775807, F2: -155}
        for fi6 := 0; fi6 < 8; fi6++ {
            if false {
                fmt.Println(v2, v2, v2)
            }
            _ = fi6
        }
        v2 = fi4
        _ = fi4
        _ = v5
    }
    v7 := []int{(v2 / ((v2) | 1))}
    switch (^(v2)) {
    case 32767:
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    case 2147483647:
        if (0.1 > 1e10) {
            var v8 bool = (true && true)
            var v9 int16 = 0
            v10 := GenS1{F0: 18446744073709551615, F1: 0, F2: -1}
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 int = (v2 + (v2 ^ (-(v2))))
            v11 = (v2 % ((v2) | 1))
            _ = v11
        }
    }
    v2 += v2
    v12 := GenS2{F0: -1, F1: 0, F2: (42 >= 0)}
    var v13 float32 = 1e10
    v14 := GenS3{F0: int64(v2), F1: 75, F2: -1}
    for fi15 := 0; fi15 < 18; fi15++ {
        v13 = float32(float64(((v13 + v13) - v13)))
        var v16 int = -1
        _ = fi15
        _ = v16
    }
    var v17 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
