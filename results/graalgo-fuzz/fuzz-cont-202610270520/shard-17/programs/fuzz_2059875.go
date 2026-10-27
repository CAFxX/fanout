package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar", "\n\t"}
    var v2 int8 = 127
    var v3 string = "a"
    v4 := make(map[int]int)
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 446, F1: -1}
    var v7 int = 360
    var v8 uint16 = 42
    v2 = -128
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint32 = 827
        var v11 int64 = -1
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if true {
            var v13 uint8 = 197
            v7 += fi12
            v7 -= (-(v7))
            _ = v13
        } else {
            var v14 int32 = 1
            v2 = (v2 & -128)
            _ = v14
        }
        _ = fi12
    }
    switch (v7 << ((18446744073709551615) & 63)) {
    case 4294967295:
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 float64 = (float64(1.5) + (1.5 * 0.0))
            v2 += v2
            if (int32(v7) > 2147483647) {
                v3 = ((("\n\t" + ((v3 + v3) + "foo bar")) + ("" + (v3 + v3))) + (("a" + "") + (v3 + "a")))
                var v17 uint64 = 0
                v7 += 773
                _ = v17
            } else {
                v2 = v2
            }
            _ = fi15
            _ = v16
        }
    default:
        var v18 int64 = -107
        _ = v18
    }
    var v19 uint16 = v8
    for fi20 := 0; fi20 < 5; fi20++ {
        v8 = 32767
        fmt.Println(v19, v3, fi20)
        var v21 uint32 = 2147483647
        _ = fi20
        _ = v21
    }
    v8 += 437
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint32 = 65535
        delete(v4, (v7 / ((v7) | 1)))
        switch "foo bar" {
        case "hello":
            var v24 uint = 480
            _ = v24
        case "":
            var v25 int16 = int16(fi22)
            _ = v25
        default:
            v2 += (^(v2))
        }
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
