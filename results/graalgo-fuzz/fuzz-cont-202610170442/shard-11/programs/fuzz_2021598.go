package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = -111
    v3 := []uint32{42, 65535, 0, 2}
    v4 := make(map[string]int)
    var v5 uint64 = 42
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2, v5)
        case "\n\t":
            v7 := GenS1{F0: "", F1: 255, F2: 42}
            _ = v7
        case "hello":
            var v8 int = (fi6 / ((fi6) | 1))
            _ = v8
        default:
            v5 += v5
        }
        _ = fi6
    }
    if false {
        var v9 int64 = -9223372036854775808
        v5 += v5
        var v10 int32 = -236
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: "foo bar", F1: 118, F2: v5}
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: "", F1: 296, F2: 780}
        fmt.Println(v5, v2, fi12)
        var v14 uint16 = 163
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int64 = 1
    var v16 uint64 = (v5 & 18446744073709551615)
    switch 42 {
    case 4294967295:
        if len(v3) > 0 {
            v3[0] = 65535
        }
    default:
        if ((v15 ^ v15) < (v15 >> ((255) & 63))) {
            for fi17 := 0; fi17 < 11; fi17++ {
                if (!(("a" > "\n\t") || false)) {
                    var v18 uint = 693
                    fmt.Println(v16)
                    _ = v18
                } else {
                    var v19 int = (-(fi17))
                    var v20 uint = 18446744073709551615
                    _ = v19
                    _ = v20
                }
                v2 = (v2 >> ((65535) & 63))
                var v21 uint16 = 1
                _ = fi17
                _ = v21
            }
        }
    }
    _ = len(v3)
    if (!false) {
        fmt.Println(v5)
        for fi22 := 0; fi22 < 13; fi22++ {
            var v23 float64 = 0.1
            _ = fi22
            _ = v23
        }
    }
    var v24 int64 = v15
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
