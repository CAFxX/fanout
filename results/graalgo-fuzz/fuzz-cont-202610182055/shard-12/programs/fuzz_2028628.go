package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }

func main() {
    v1 := []uint32{1, 600, 358}
    var v2 uint = 0
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    fmt.Println(v2)
    var v5 int64 = -9223372036854775808
    var v6 int32 = 42
    if (32767 == -32768) {
        var v7 uint16 = 511
        var v8 uint8 = 128
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: 127}
    var v10 int8 = 1
    v5 += 42
    if (v4 <= float32(1e10)) {
        switch "a" {
        case "":
            _ = len(v3)
        case "a":
            var v11 uint = v2
            _ = v11
        case "foo bar":
            var v12 int32 = v6
            _ = v12
        default:
            if true {
                v6 += (v6 * v6)
                v1 = append(v1, 42)
                for fi13 := 0; fi13 < 17; fi13++ {
                    v5 += (v5 >> ((4294967295) & 63))
                    switch "\n\t" {
                    case "a":
                        v10 = (v10 % ((v10) | 1))
                    }
                    _ = fi13
                }
            } else {
                var v14 string = "\n\t"
                var v15 int32 = (v6 ^ v6)
                _ = v14
                _ = v15
            }
        }
        v1 = append(v1, 42)
        if true {
            _ = len(v1)
            for fi16 := 0; fi16 < 13; fi16++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v17 := GenS1{F0: 1, F1: fi16}
                _ = fi16
                _ = v17
            }
            var v18 int8 = 127
            _ = v18
        }
    } else {
        switch "\n\t" {
        case "":
            v2 = v2
        case "a":
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 int64 = (v5 / ((v5) | 1))
                _ = fi19
                _ = v20
            }
        default:
            var v21 uint8 = 120
            _ = v21
        }
        v5 -= (v5 * -816)
        var v22 int16 = 0
        _ = v22
    }
    fmt.Println(v10, v5, v6)
    v23 := GenS2{F0: int64(v6)}
    v24 := GenS3{F0: v6}
    v25 := GenS4{F0: -1, F1: -1, F2: (v10 & (v10 / ((v10) | 1)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
