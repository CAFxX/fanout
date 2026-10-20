package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 uint16, p1 uint64) bool {
    p0 = (p0 * 515)
    p0 = (^(p0))
    _ = p0
    _ = p1
    return false
}
func genFn1(p0 int32, p1 bool) uint {
    p0 = (p0 + 1)
    p0 += (^(p0))
    var v4 int = 32768
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 uint = 0
    var v3 int16 = 632
    v5 := []int{int(v2)}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 bool = (v3 > (v3 ^ v3))
        _ = fi6
        _ = v7
    }
    var v8 float64 = -2.25
    v1 = -176
    v2 = genFn1((v1 * (v1 % (((v1 << (((v2 + (v2 / ((v2) | 1)))) & 63))) | 1))), (-128 != 127))
    v9 := GenS1{F0: 138, F1: 469, F2: 42}
    if true {
        v2 = genFn1(1, (!(!false)))
        var v10 bool = false
        switch 460 {
        case 32768:
            var v11 int8 = -1
            _ = v11
        }
        _ = v10
    } else {
        fmt.Println(v8, v3, v1)
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v3 = v3
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: "", F1: 65535}
            var v15 bool = (!(4294967295 == 1))
            _ = fi13
            _ = v14
            _ = v15
        }
        v2 += 18446744073709551615
        _ = fi12
    }
    switch "foo bar" {
    case "\n\t":
        var v16 uint = 255
        _ = v16
    case "":
        for fi17 := 0; fi17 < 2; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 uint8 = 2
                v3 = (v3 ^ (v3 >> ((v2) & 63)))
                v3 -= v3
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    default:
        fmt.Println(v1, v8)
    }
    var v20 uint16 = 32767
    var v21 float32 = 1e10
    fmt.Println(v8, v2, v20)
    switch "a" {
    case "":
        var v22 int32 = (^(v1))
        _ = v22
    default:
        v5 = append(v5, 444)
    }
    v2 = uint(v20)
    var v23 uint64 = 840
    for fi24 := 0; fi24 < 13; fi24++ {
        v2 = v2
        v20 -= (v20 >> ((genFn1((v1 >> (((v2 | (^(v2)))) & 63)), true)) & 63))
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
