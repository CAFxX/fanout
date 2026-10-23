package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]string)
    var v3 string = "foo bar"
    var v4 uint64 = 167
    var v5 int = 9223372036854775807
    v6 := []uint32{(v1 ^ (v1 | (v1 * (v1 - v1)))), v1, (v1 + 2), 769}
    v4 = 2
    if (!false) {
        var v7 uint32 = 762
        _ = v7
    }
    switch v3 {
    case "hello":
        for fi8 := 0; fi8 < 3; fi8++ {
            fmt.Println(v1, v5)
            _ = fi8
        }
    default:
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: -347, F1: -744, F2: 66}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi9
            _ = v10
        }
    }
    v4 -= 4294967295
    v5 += 18
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v11 uint16 = 2
    v12 := GenS2{F0: v3, F1: 542, F2: 42}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = -128
        _ = fi13
        _ = v14
    }
    switch (v3 + "\n\t") {
    case "\n\t":
        if false {
            v15 := GenS1{F0: -1, F1: 42, F2: 127}
            var v16 int64 = 778
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 612, F1: 1, F2: 118}
            _ = v17
        }
    }
    v5 += (v5 << ((1) & 63))
    var v18 uint8 = 1
    v3 = "hello"
    if (v3 == v3) {
        delete(v2, v5)
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            _ = len(v6)
            _ = fi19
        }
    }
    v20 := GenS1{F0: 1, F1: 408, F2: v18}
    if (!false) {
        var v21 float32 = -2.25
        if (927 == 571) {
            var v22 float32 = 3.14159
            _ = v22
        } else {
            v5 -= (v5 >> ((446) & 63))
            fmt.Println(v18, v4, v3)
            v23 := GenS2{F0: v3, F1: 2, F2: 2147483647}
            _ = v23
        }
        var v24 float32 = v21
        _ = v21
        _ = v24
    } else {
        var v25 int64 = 0
        v11 -= (v11 & (v11 ^ (v11 / (((v11 + v11)) | 1))))
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
