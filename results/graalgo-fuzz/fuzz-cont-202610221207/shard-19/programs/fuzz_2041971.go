package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]string)
    var v3 int8 = -128
    v4 := []int{9223372036854775807, 479, 0, 127}
    v1 -= v1
    switch -344 {
    case 32768:
        v1 -= v1
    case 4294967295:
        var v5 uint32 = 0
        _ = v5
    case -9223372036854775808:
        v3 -= -128
    default:
        for fi6 := 0; fi6 < 17; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi7
            }
            v8 := GenS1{F0: 0, F1: (false || true), F2: 2}
            var v9 int8 = 127
            _ = fi6
            _ = v8
            _ = v9
        }
    }
    v1 = v1
    v1 = v1
    if (false && (!(!false))) {
        var v10 uint8 = 193
        _ = len(v4)
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint = 0
            _ = fi12
            _ = v13
        }
        v2[fi11] = ("foo bar" + "a")
        v2[(fi11 ^ fi11)] = "a"
        _ = fi11
    }
    var v14 uint8 = 1
    var v15 bool = false
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: -667, F1: (^(v1)), F2: 241}
        v4 = append(v4, fi16)
        _ = fi16
        _ = v17
    }
    var v18 uint = 65535
    v19 := GenS1{F0: 0, F1: (255 <= 1), F2: v18}
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 6; fi21++ {
            v15 = false
            v14 = (v14 * 251)
            var v22 uint64 = 511
            var v23 uint = (v18 + (v18 ^ (v18 + (v18 % ((4294967295) | 1)))))
            v15 = v15
            v15 = v15
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        var v25 uint32 = (v1 + 2147483648)
        var v26 bool = v15
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
