package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]int)
    v3 := []uint32{0, 2, 230, 65535}
    v3 = append(v3, 2147483648)
    var v4 uint32 = 169
    for fi5 := 0; fi5 < 10; fi5++ {
        v1 = 1e10
        v4 += 4294967295
        _ = len(v3)
        _ = fi5
    }
    v3 = append(v3, 37)
    var v6 int32 = 1
    var v7 int8 = 1
    var v8 uint64 = 4294967295
    switch "\n\t" {
    case "hello":
        v9 := GenS1{F0: 42, F1: (v4 / (((v4 ^ v4)) | 1)), F2: 42}
        _ = v9
    case "\n\t":
        var v10 int64 = -1
        _ = v10
    }
    switch "hello" {
    case "foo bar":
        if (true && true) {
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 16; fi12++ {
                    if ((fi12 / ((fi12) | 1)) < fi11) {
                        v4 -= (v4 * v4)
                        var v13 int32 = (v6 - v6)
                        fmt.Println(v13, v13, v6)
                        var v14 int8 = (v7 - (v7 + 42))
                        _ = v13
                        _ = v14
                    }
                    v8 -= (v8 % (((^(v8))) | 1))
                    v15 := GenS2{F0: 1, F1: v4}
                    _ = fi12
                    _ = v15
                }
                _ = fi11
            }
        } else {
            _ = len(v2)
            fmt.Println(v1)
            fmt.Println(v4)
        }
    case "":
        v16 := GenS1{F0: 255, F1: 373, F2: 42}
        _ = v16
    case "\n\t":
        var v17 uint8 = 255
        _ = v17
    default:
        fmt.Println(v8, v7)
    }
    if true {
        fmt.Println(v6, v7, v4)
        var v18 float64 = v1
        _ = v18
    }
    switch "hello" {
    case "hello":
        v6 -= -963
    case "a":
        fmt.Println(v1)
    default:
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                v2[v4] = 4294967295
                _ = fi20
            }
            var v21 int32 = (v6 << ((489) & 63))
            _ = len(v3)
            _ = fi19
            _ = v21
        }
    }
    if (("" + "a") >= "foo bar") {
        _ = len(v2)
        _ = len(v2)
    }
    v22 := GenS2{F0: -73, F1: (v4 ^ v4)}
    v4 = v4
    v23 := GenS2{F0: v7, F1: v4}
    v24 := v2[4294967295]
    v25 := GenS1{F0: 0, F1: 2147483648, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
