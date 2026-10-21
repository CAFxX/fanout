package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]int)
    var v3 int = 32768
    var v4 int64 = 9223372036854775807
    var v5 uint64 = 2
    for fi6 := 0; fi6 < 14; fi6++ {
        if true {
            var v7 int16 = 0
            fmt.Println(v3, v7)
            _ = v7
        }
        _ = fi6
    }
    v8 := GenS1{F0: "foo bar"}
    v3 -= (v3 | (^(v3)))
    v4 += v4
    if (false || true) {
        if ((704 <= 42) && true) {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint64 = (v1 ^ v1)
                fmt.Println(fi9, fi9)
                _ = len(v2)
                _ = fi9
                _ = v10
            }
        }
        v11 := v2[-704]
        v3 -= v3
        _ = v11
    }
    var v12 uint32 = 701
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := []int{v3, (^(v3))}
        _ = fi13
        _ = v14
    }
    v3 -= v3
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS2{F0: (v5 % ((v5) | 1)), F1: 1000, F2: 641}
        var v17 float32 = -2.25
        v18 := GenS3{F0: -242, F1: 255, F2: 42}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    switch "hello" {
    case "":
        var v19 int64 = v4
        _ = v19
    case "a":
        v20 := []int{(v3 - (v3 * (v3 / (((v3 >> ((255) & 63))) | 1)))), -1, (v3 / ((2147483647) | 1))}
        _ = v20
    default:
        var v21 uint64 = (v1 >> ((611) & 63))
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
