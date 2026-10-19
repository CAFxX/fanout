package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 string; F1 string }

func main() {
    var v1 uint32 = 91
    v2 := make(map[uint32]bool)
    v3 := []uint32{2147483647, 174}
    var v4 int = 206
    var v5 int16 = -32768
    var v6 int8 = -73
    _ = len(v2)
    var v7 uint8 = 42
    var v8 int16 = (^(v5))
    fmt.Println(v5, v7, v8)
    v5 += (v5 << ((4294967295) & 63))
    switch 4294967295 {
    case -871:
        v9 := GenS1{F0: 2, F1: 9223372036854775807, F2: v4}
        _ = v9
    }
    if (!(0 > -1)) {
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 uint32 = (^(v1))
            var v12 string = "hello"
            for fi13 := 0; fi13 < 10; fi13++ {
                v1 = (v1 >> ((42) & 63))
                if false {
                    var v14 bool = (v12 >= "a")
                    v8 += 0
                    v15 := GenS2{F0: (v12 + "a"), F1: 0, F2: (v7 & (v7 + 255))}
                    _ = v14
                    _ = v15
                } else {
                    var v16 bool = true
                    _ = len(v2)
                    v17 := GenS2{F0: "\n\t", F1: 1, F2: v7}
                    v18 := GenS3{F0: ((v12 + v12) + v12), F1: v12}
                    v19 := v2[(v11 / (((v11 + v11)) | 1))]
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v8 = v8
                v20 := GenS3{F0: "\n\t", F1: v12}
                v4 = fi13
                _ = fi13
                _ = v20
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v1 = v1
    } else {
        _ = len(v2)
    }
    var v21 uint8 = v7
    v22 := GenS3{F0: "hello", F1: "hello"}
    fmt.Println(v8, v7, v6)
    if (((0.1 - 0.1) * 1.5) != 1e10) {
        v6 = v6
        v7 = v7
        var v23 bool = (42 >= 100)
        _ = v23
    } else {
        v24 := v2[v1]
        var v25 uint16 = 170
        fmt.Println(v24, v21, v6)
        _ = v24
        _ = v25
    }
    v5 = (-(v8))
    for fi26 := 0; fi26 < 7; fi26++ {
        v21 += (v21 * (v7 * (v21 - (^(v7)))))
        fmt.Println(v8)
        var v27 float64 = 1.5
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
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
