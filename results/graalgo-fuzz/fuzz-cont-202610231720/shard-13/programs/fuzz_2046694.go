package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -265
    v3 := []uint32{907, 2, 994, 42}
    fmt.Println(v2)
    for fi4 := 0; fi4 < 8; fi4++ {
        fmt.Println(fi4, v2)
        v1[fi4] = (fi4 * fi4)
        for fi5 := 0; fi5 < 12; fi5++ {
            if (620 >= 32767) {
                if (((0.0 * 0.1) - -2.25) == (1.5 / -2.25)) {
                    var v6 string = "hello"
                    fmt.Println(v6, v2)
                    _ = v6
                }
                var v7 string = "a"
                _ = v7
            } else {
                fmt.Println(fi5, fi4, v2)
            }
            v8 := GenS1{F0: 2, F1: 27, F2: 217}
            v2 += -1
            _ = fi5
            _ = v8
        }
        _ = fi4
    }
    var v9 int = 1
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS2{F0: 0}
        var v12 uint = 0
        var v13 int64 = 9223372036854775807
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int32 = (v2 | 0)
    for fi15 := 0; fi15 < 2; fi15++ {
        if ((-(v2)) < v14) {
            v16 := GenS3{F0: 33}
            _ = v16
        }
        _ = fi15
    }
    v9 = v9
    var v17 uint16 = 60
    _ = len(v1)
    var v18 bool = true
    v9 = (v9 + 118)
    v9 = 128
    var v19 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
