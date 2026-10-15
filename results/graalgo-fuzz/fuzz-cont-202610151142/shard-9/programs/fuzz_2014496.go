package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    var v1 uint8 = 255
    var v2 int32 = 388
    v3 := make(map[string]uint32)
    v4 := GenS1{F0: v2, F1: (v2 >> ((85) & 63))}
    switch "\n\t" {
    case "\n\t":
        var v5 uint32 = 4294967295
        _ = v5
    case "hello":
        v6 := v3[("hello" + "hello")]
        _ = v6
    default:
        v2 = -2147483648
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if true {
            if (9223372036854775807 > -159) {
                var v8 int = (fi7 << ((659) & 63))
                for fi9 := 0; fi9 < 19; fi9++ {
                    fmt.Println(v1, v2, v1)
                    var v10 float64 = 0.1
                    v11 := GenS1{F0: v2, F1: (v2 & (v2 & v2))}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint16 = 255
            _ = fi12
            _ = v13
        }
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    var v14 float32 = 0.1
    var v15 uint16 = 640
    fmt.Println(v2)
    switch "\n\t" {
    case "foo bar":
        v15 += (v15 | uint16(v1))
    default:
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := v3[(("\n\t" + "hello") + "foo bar")]
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: 2147483647, F1: (v2 << ((2) & 63))}
    v3[("foo bar" + ("a" + ""))] = 4294967295
    for fi19 := 0; fi19 < 11; fi19++ {
        v3[("\n\t" + "hello")] = 2
        for fi20 := 0; fi20 < 18; fi20++ {
            _ = len(v3)
            switch fi19 {
            case 496:
                var v21 uint64 = uint64(v1)
                _ = v21
            case 0:
                _ = len(v3)
            }
            v22 := GenS1{F0: (v2 + v2), F1: -138}
            _ = fi20
            _ = v22
        }
        _ = fi19
    }
    v14 = (v14 * float32((0.1 + float64((((0.0 / v14) * 0.1) + float32(0.0))))))
    var v23 bool = true
    v3["a"] = 1
    for fi24 := 0; fi24 < 14; fi24++ {
        v25 := []int{(fi24 + 2147483647), (fi24 % ((fi24) | 1))}
        v26 := GenS1{F0: v2, F1: v2}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int64 = -863
    v14 = (1.5 / (float32(1e10) - v14))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
}
