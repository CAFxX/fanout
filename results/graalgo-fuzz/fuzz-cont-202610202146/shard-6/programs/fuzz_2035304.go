package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 float64 = (-2.25 / -2.25)
    var v2 uint64 = 18446744073709551615
    v3 := []string{"a", "", "hello", "hello"}
    v4 := []string{"", "\n\t", ""}
    v5 := []int{955, 6}
    v6 := GenS1{F0: 32768}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int16 = -32768
        switch 127 {
        case 127:
            switch fi7 {
            case 32768:
                for fi9 := 0; fi9 < 19; fi9++ {
                    var v10 int16 = (v8 - (v8 * -103))
                    _ = fi9
                    _ = v10
                }
            default:
                switch "a" {
                case "":
                    fmt.Println(v1, v8)
                default:
                    var v11 uint = 1
                    _ = v11
                }
            }
        default:
            for fi12 := 0; fi12 < 15; fi12++ {
                v8 += (-(v8))
                _ = fi12
            }
        }
        _ = fi7
        _ = v8
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    switch 32767 {
    case -934:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 312:
        switch 32768 {
        case 127:
            v13 := GenS1{F0: -2147483648}
            _ = v13
        case 32767:
            v1 += float64(0.0)
        }
    }
    switch 1 {
    case -9223372036854775808:
        for fi14 := 0; fi14 < 1; fi14++ {
            v2 -= (v2 ^ (v2 >> ((42) & 63)))
            for fi15 := 0; fi15 < 14; fi15++ {
                v16 := GenS2{F0: 9223372036854775807, F1: fi14}
                v2 += (v2 ^ (v2 + (v2 % ((v2) | 1))))
                v2 += (v2 * 4294967295)
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    default:
        var v17 uint = 255
        _ = v17
    }
    if true {
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        v3 = append(v3, "")
    }
    var v18 uint16 = 143
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint8 = 1
        _ = len(v3)
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v1, v2)
    if true {
        switch "foo bar" {
        case "":
            if len(v5) > 0 {
                v5[0] = -634
            }
        default:
            fmt.Println(v18, v18, v18)
        }
        var v21 float32 = (3.14159 + 1e10)
        v21 -= -2.25
        _ = v21
    } else {
        switch 693 {
        case 2147483647:
            v2 -= (v2 << ((2) & 63))
        case 525:
            v22 := GenS2{F0: 420, F1: 310}
            _ = v22
        case -2147483648:
            var v23 uint16 = (v18 | 65535)
            _ = v23
        }
        var v24 int8 = 106
        _ = v24
    }
    if (4294967295 < 1) {
        v25 := GenS3{F0: 42, F1: 174}
        for fi26 := 0; fi26 < 17; fi26++ {
            var v27 int = 1
            v28 := GenS1{F0: (v27 - v27)}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
