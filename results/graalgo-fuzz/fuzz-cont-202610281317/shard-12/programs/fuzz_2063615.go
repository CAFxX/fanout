package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    var v1 int8 = -1
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    var v4 int64 = 150
    v5 := GenS1{F0: v1, F1: false}
    v4 += 1
    var v6 uint16 = 281
    v7 := []int{-994}
    var v8 int = 9223372036854775807
    if ((true || false) || false) {
        v9 := GenS1{F0: v1, F1: (4294967295 < 0)}
        _ = v9
    } else {
        var v10 string = "foo bar"
        var v11 uint64 = 2
        _ = v10
        _ = v11
    }
    v2 += v2
    _ = len(v3)
    var v12 float32 = (1.5 * v2)
    v13 := GenS2{F0: (v4 - v4)}
    var v14 int64 = v4
    for fi15 := 0; fi15 < 7; fi15++ {
        if (!true) {
            v14 += v14
        } else {
            v12 += (1e10 - v12)
            if false {
                v1 -= v1
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (88 > 2) {
                        v17 := GenS3{F0: 42, F1: 598, F2: 189}
                        _ = v17
                    }
                    fmt.Println(fi16, fi16, v1)
                    var v18 uint32 = 65535
                    _ = fi16
                    _ = v18
                }
                v19 := GenS1{F0: v1, F1: false}
                _ = v19
            } else {
                v20 := GenS3{F0: 2, F1: uint(v6), F2: -32768}
                v1 = v1
                _ = v20
            }
        }
        fmt.Println(v4, v8, v4)
        if len(v7) > 0 {
            v7[0] = v8
        }
        _ = fi15
    }
    v21 := GenS1{F0: v1, F1: false}
    var v22 uint8 = 1
    v22 = (v22 * 127)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
