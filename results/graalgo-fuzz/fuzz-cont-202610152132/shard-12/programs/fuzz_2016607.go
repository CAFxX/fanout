package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 uint8 = 255
    var v5 uint32 = 2
    v6 := GenS1{F0: (v4 | (^(v4))), F1: v4}
    var v7 uint = 0
    if true {
        v7 = v7
    }
    v4 += (v4 - 42)
    if (!false) {
        var v8 int16 = 32767
        v4 = v4
        v9 := GenS1{F0: 128, F1: (v4 - v4)}
        _ = v8
        _ = v9
    } else {
        v2 = append(v2, (v5 - (v5 / ((65535) | 1))))
        v10 := GenS1{F0: (v4 / ((128) | 1)), F1: (^(v4))}
        _ = len(v1)
        _ = v10
    }
    if (100 < 42) {
        var v11 uint32 = 800
        var v12 int32 = -1
        fmt.Println(v5)
        _ = v11
        _ = v12
    } else {
        if true {
            if len(v2) > 0 {
                v2[0] = v5
            }
            v4 = (v4 << ((994) & 63))
            fmt.Println(v7, v7, v5)
        } else {
            v1[1] = v5
        }
        if true {
            if (42 >= 0) {
                v5 -= (v5 | (v5 / (((v5 ^ (v5 / ((v5) | 1)))) | 1)))
                switch (("\n\t" + "hello") + "hello") {
                case "foo bar":
                    v4 = (v4 % (((v4 % ((127) | 1))) | 1))
                case "\n\t":
                    var v13 string = ""
                    _ = v13
                case "a":
                    v14 := GenS2{F0: v5, F1: 1, F2: v5}
                    _ = v14
                default:
                    v5 = v5
                }
                fmt.Println(v4, v7)
            }
            v4 = 116
            v15 := GenS2{F0: 1, F1: 978, F2: uint32(v4)}
            _ = v15
        }
    }
    var v16 int64 = 1
    switch 32768 {
    case 9223372036854775807:
        v17 := v1[2147483647]
        _ = v17
    default:
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 uint = v7
            var v20 int8 = -128
            v16 = (v16 - v16)
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    _ = len(v2)
    v4 = v4
    var v21 int64 = -793
    v22 := GenS1{F0: (^(v4)), F1: 179}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
