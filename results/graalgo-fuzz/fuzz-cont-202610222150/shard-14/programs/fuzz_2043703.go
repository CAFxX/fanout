package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 int; F2 uint32 }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    v1 := []int{128}
    var v2 uint32 = 415
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    var v5 uint64 = 168
    v6 := GenS1{F0: -41, F1: true, F2: 833}
    var v7 uint16 = 42
    v8 := GenS2{F0: -227, F1: 77}
    var v9 int16 = 42
    v10 := GenS2{F0: 0, F1: 42}
    if (-2147483648 >= 630) {
        _ = len(v1)
        if false {
            var v11 uint64 = 4294967295
            if (false || false) {
                fmt.Println(v7, v11, v5)
                var v12 int32 = 42
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := v3[32767]
                    v15 := v4[(^(fi13))]
                    if (v15 && (!(false || (false && true)))) {
                        if false {
                            v16 := GenS3{F0: 42, F1: fi13, F2: (v14 & v14)}
                            _ = v16
                        }
                        var v17 uint64 = 2
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                v2 += 4294967295
                _ = fi18
            }
            _ = v11
        }
    } else {
        switch 4294967295 {
        case -2147483648:
            var v19 float32 = (3.14159 + (1e10 - (float32(3.14159) - (0.0 / 1.5))))
            _ = v19
        }
        v2 = (v2 >> ((1) & 63))
    }
    v5 = (v5 ^ uint64(v2))
    if true {
        v9 += -194
        if (!true) {
            v2 += v2
            var v20 int16 = v9
            _ = v20
        }
    } else {
        var v21 int32 = -682
        if (42 != (v9 & v9)) {
            if true {
                var v22 float32 = 1e10
                v23 := GenS4{F0: -116, F1: -1}
                _ = v22
                _ = v23
            }
            fmt.Println(v5, v7, v9)
            v1 = append(v1, -9223372036854775808)
        }
        _ = v21
    }
    if true {
        v2 += v2
    } else {
        for fi24 := 0; fi24 < 1; fi24++ {
            v7 = v7
            fmt.Println(v9, v5)
            fmt.Println(v2)
            _ = fi24
        }
        _ = len(v3)
    }
    v7 = (v7 ^ uint16(v2))
    v4[int(v2)] = (!true)
    v2 -= (^(v2))
    v7 = (v7 | (^(v7)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
