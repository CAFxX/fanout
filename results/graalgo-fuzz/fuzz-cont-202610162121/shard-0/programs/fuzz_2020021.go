package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 int16) uint16 {
    switch 4294967295 {
    case 127:
        switch 4294967295 {
        case -9223372036854775808:
            fmt.Println(p0)
        default:
            var v5 uint = 18446744073709551615
            _ = v5
        }
    case 9223372036854775807:
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 9223372036854775807, F1: -655}
            _ = fi6
            _ = v7
        }
    }
    var v8 int64 = 42
    _ = p0
    _ = v8
    return 2
}
func genFn1(p0 int8, p1 uint8, p2 bool) int32 {
    p0 = (p0 / (((-(p0))) | 1))
    p0 -= (p0 ^ 127)
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = -1
        _ = fi9
        _ = v10
    }
    if (!(p2 || true)) {
        v11 := []int{0, 396}
        fmt.Println(p0, p0, p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 975
}

func main() {
    var v1 uint8 = 0
    var v2 float64 = (float64(0.0) / 1e10)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]int)
    switch 0 {
    case 128:
        for fi12 := 0; fi12 < 6; fi12++ {
            v1 -= 255
            switch "" {
            case "":
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 int = (fi12 << ((42) & 63))
                    v14 = v14
                    fmt.Println(fi12, v2)
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi12
        }
    }
    switch 128 {
    case -2147483648:
        v2 = (v2 + (((1.5 - float64(((float32(v2) - (0.1 / 1e10)) + 3.14159))) * 1e10) + (3.14159 - v2)))
    case 2147483647:
        var v15 float32 = 1e10
        _ = v15
    case 128:
        for fi16 := 0; fi16 < 15; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                v1 += v1
                _ = fi17
            }
            _ = fi16
        }
    default:
        if (0 < 32767) {
            var v18 uint16 = 0
            _ = v18
        } else {
            v2 += v2
        }
    }
    v2 = v2
    var v19 uint8 = 0
    if false {
        var v20 string = ("foo bar" + "foo bar")
        _ = v20
    }
    v21 := GenS1{F0: -393, F1: 42}
    v22 := GenS1{F0: -1, F1: 9223372036854775807}
    if (v19 != (v1 - v19)) {
        v23 := []int{}
        fmt.Println(v1, v2, v1)
        var v24 float64 = -2.25
        _ = v23
        _ = v24
    } else {
        fmt.Println(v2, v1)
        fmt.Println(v1, v2)
        switch 9223372036854775807 {
        case 369:
            if false {
                fmt.Println(v2, v2, v19)
            } else {
                v25 := []int{-9223372036854775808, 42, 693}
                fmt.Println(v1, v2)
                _ = v25
            }
        case 32768:
            v26 := GenS1{F0: -893, F1: -1}
            _ = v26
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
