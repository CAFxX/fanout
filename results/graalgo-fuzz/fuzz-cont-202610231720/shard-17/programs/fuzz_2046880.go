package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 uint32 = 4294967295
    var v4 uint64 = 42
    var v5 int16 = 0
    if false {
        var v6 int16 = v5
        _ = len(v1)
        _ = v6
    } else {
        switch 32768 {
        case -1:
            var v7 float64 = -2.25
            _ = v7
        case 969:
            var v8 bool = ("" < ("hello" + "\n\t"))
            _ = v8
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int16 = v5
    v5 = (-(v9))
    v10 := GenS1{F0: v9, F1: 141, F2: (^(v3))}
    fmt.Println(v9, v4)
    var v11 float32 = 3.14159
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: (v5 + -1), F1: 42, F2: v3}
            fmt.Println(v11)
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v15 := GenS1{F0: 49, F1: 4294967295, F2: 159}
    switch 763 {
    case 9223372036854775807:
        v16 := GenS1{F0: (v9 ^ -436), F1: 0, F2: v3}
        _ = v16
    case -599:
        var v17 int16 = (v9 | v5)
        _ = v17
    }
    v2 = append(v2, (^(v3)))
    delete(v1, 42)
    switch 810 {
    case -929:
        if true {
            for fi18 := 0; fi18 < 4; fi18++ {
                var v19 uint32 = (v3 * (v3 ^ v3))
                _ = fi18
                _ = v19
            }
            v5 = 0
        } else {
            v3 = v3
        }
    case 127:
        if false {
            v5 = v5
            var v20 string = "foo bar"
            var v21 uint8 = 1
            _ = v20
            _ = v21
        } else {
            var v22 float64 = 3.14159
            fmt.Println(v22)
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            switch fi23 {
            case 1:
                if ((fi23 | 128) != 1) {
                    v9 = v5
                    v3 = (v3 - v3)
                }
            case 9223372036854775807:
                for fi24 := 0; fi24 < 8; fi24++ {
                    _ = len(v1)
                    v9 = v5
                    var v25 int16 = v5
                    _ = len(v2)
                    v11 = -2.25
                    _ = fi24
                    _ = v25
                }
            case -1:
                for fi26 := 0; fi26 < 11; fi26++ {
                    v1[uint32(v4)] = (v3 % ((165) | 1))
                    _ = fi26
                }
            default:
                v9 = v9
            }
            v5 -= 1000
            v27 := GenS1{F0: v9, F1: 42, F2: v3}
            _ = fi23
            _ = v27
        }
    }
    v1[0] = (v3 ^ 222)
    switch "" {
    case "a":
        v28 := v1[(v3 ^ v3)]
        _ = v28
    case "\n\t":
        v1[(v3 ^ v3)] = (^(v3))
    case "foo bar":
        v4 = v4
    }
    fmt.Println(v3, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
