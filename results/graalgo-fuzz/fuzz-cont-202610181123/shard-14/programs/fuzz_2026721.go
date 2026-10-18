package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }

func main() {
    var v1 uint16 = 916
    v2 := make(map[int]int)
    v3 := []uint32{0, 256}
    v1 -= (v1 + (v1 | (v1 << ((882) & 63))))
    _ = len(v3)
    var v4 uint16 = v1
    v5 := GenS1{F0: -32768, F1: -665}
    var v6 int = -2147483648
    switch "" {
    case "a":
        v7 := v2[v6]
        _ = v7
    case "\n\t":
        v8 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v8
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v4 -= 1
    _ = len(v3)
    if true {
        var v9 float64 = 1e10
        v10 := GenS1{F0: 392, F1: v6}
        v11 := GenS2{F0: true}
        _ = v10
        _ = v11
        _ = v9
    } else {
        var v12 uint8 = 228
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 int16 = -372
            for fi15 := 0; fi15 < 15; fi15++ {
                _ = len(v3)
                _ = fi15
            }
            switch int(v12) {
            case 128:
                var v16 uint16 = v1
                _ = v16
            default:
                v4 += v4
            }
            _ = fi13
            _ = v14
        }
        fmt.Println(v6)
        _ = v12
    }
    fmt.Println(v6, v6, v6)
    v4 -= (^(v4))
    if ((v4 << ((255) & 63)) < 0) {
        var v17 float64 = -2.25
        v6 = v6
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS3{F0: fi18}
            v6 -= (v6 * v6)
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    switch "foo bar" {
    case "\n\t":
        v1 -= (v4 * v4)
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            for fi21 := 0; fi21 < 3; fi21++ {
                v22 := GenS3{F0: (v6 * v6)}
                v23 := GenS1{F0: 42, F1: (-(fi20))}
                var v24 int32 = -260
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 uint8 = 42
            _ = fi20
            _ = v25
        }
    }
    var v26 uint = 18446744073709551615
    for fi27 := 0; fi27 < 11; fi27++ {
        v6 -= (fi27 + 963)
        delete(v2, (fi27 >> (((^(v26))) & 63)))
        v26 = (v26 ^ (v26 | 65535))
        _ = fi27
    }
    var v28 int16 = -1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
