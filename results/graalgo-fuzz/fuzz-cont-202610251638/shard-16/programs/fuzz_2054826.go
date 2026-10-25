package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 int8, p1 int64, p2 uint8) uint16 {
    if ((p2 & (p2 << ((749) & 63))) < 1) {
        if true {
            fmt.Println(p1)
        }
        for fi5 := 0; fi5 < 7; fi5++ {
            p2 -= p2
            _ = fi5
        }
    } else {
        switch 128 {
        case 42:
            switch "foo bar" {
            case "hello":
                fmt.Println(p0, p0)
            default:
                p2 = (p2 / ((42) | 1))
            }
        case -743:
            for fi6 := 0; fi6 < 3; fi6++ {
                p2 = (p2 | 79)
                _ = fi6
            }
        }
    }
    v7 := GenS1{F0: 32767, F1: 524}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{955, 65535}
    var v3 int8 = -123
    v4 := []int{}
    v1 = (v1 + v1)
    var v8 uint8 = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 -= v3
    v3 = v3
    v8 = 134
    v9 := GenS2{F0: v1, F1: 9223372036854775807}
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int = 0
        _ = fi10
        _ = v11
    }
    _ = len(v2)
    if false {
        v8 = v8
        _ = len(v2)
        var v12 bool = (0 != 239)
        _ = v12
    } else {
        if ("a" < "a") {
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint8 = v8
                var v15 float64 = float64((((0.1 * 3.14159) - -2.25) / (float32((float64(-2.25) * (((-2.25 - 3.14159) / 0.1) * ((1.5 - (1e10 * 0.0)) / (((0.1 - -2.25) * 1e10) / 1.5))))) * -2.25)))
                var v16 int64 = 0
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS2{F0: 1, F1: fi17}
            for fi19 := 0; fi19 < 14; fi19++ {
                v4 = append(v4, (fi19 % ((fi17) | 1)))
                v8 += (^(v8))
                v8 = v8
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    }
    var v20 int64 = -61
    v3 -= (v3 << ((2) & 63))
    if len(v2) > 0 {
        v2[0] = 42
    }
    var v21 int = -9223372036854775808
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
