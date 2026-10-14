package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0(p0 int, p1 uint16, p2 uint32) string {
    p0 = p0
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 string) int16 {
    var v5 int = -9223372036854775808
    switch v5 {
    case 9223372036854775807:
        v6 := []int{v5, (v5 * 128), (^(v5))}
        _ = v6
    case 771:
        fmt.Println(p0)
    default:
        v7 := GenS1{F0: -1, F1: (v5 << ((18446744073709551615) & 63)), F2: 791}
        _ = v7
    }
    _ = p0
    _ = v5
    return -967
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 2
    var v3 uint16 = 347
    v4 := []uint32{uint32(v2), 1}
    var v8 uint8 = 2
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 bool = true
            v12 := GenS2{F0: 0}
            v8 += (v8 / (((^(v8))) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    switch -610 {
    case 32768:
        v4 = append(v4, 2147483648)
    case 128:
        v3 += v3
    case 0:
        var v13 int64 = -609
        _ = v13
    default:
        switch "\n\t" {
        case "foo bar":
            v2 = uint64(v8)
        default:
            var v14 float64 = float64((float32(1.5) - -2.25))
            _ = v14
        }
    }
    fmt.Println(v2, v2, v3)
    if (true || (v3 > (v3 % ((0) | 1)))) {
        v1 = v1
    }
    v3 = 42
    v15 := GenS3{F0: 2147483647, F1: (^(v8)), F2: true}
    v16 := GenS3{F0: 0, F1: 127, F2: false}
    v17 := GenS2{F0: 42}
    switch 0 {
    case 32768:
        v3 = uint16(v8)
    default:
        var v18 float64 = (1.5 * (-2.25 + 1e10))
        _ = v18
    }
    var v19 int = int(v2)
    v20 := GenS3{F0: 682, F1: 255, F2: (false || (false || ((false || ((100 >= 36) || true)) || true)))}
    v3 = (v3 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
