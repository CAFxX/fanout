package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string; F1 int16; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0(p0 uint16, p1 uint8) uint8 {
    var v5 uint8 = p1
    if true {
        if false {
            var v6 string = "hello"
            v7 := []int{4294967295, 242, 1}
            var v8 bool = false
            _ = v6
            _ = v7
            _ = v8
        } else {
            v9 := []int{128, 32768, -302}
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: 1}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return p1
}
func genFn1(p0 string, p1 uint16, p2 uint64) uint8 {
    p1 -= (^(p1))
    var v11 int = -778
    v12 := GenS2{F0: p0, F1: 32767, F2: 42}
    var v13 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return genFn0(120, 128)
}

func main() {
    v1 := []int{128}
    var v2 uint32 = 42
    v3 := make(map[uint32]int)
    var v4 int64 = -1
    v2 = v2
    v4 = v4
    v14 := v3[2147483648]
    v15 := GenS2{F0: "foo bar", F1: -32768, F2: 745}
    for fi16 := 0; fi16 < 19; fi16++ {
        v4 += 9223372036854775807
        _ = fi16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS2{F0: "", F1: 1, F2: v2}
        switch 2147483647 {
        case 9223372036854775807:
            v4 -= 42
        case -1:
            v2 += (v2 / ((v2) | 1))
        default:
            switch v14 {
            case 32767:
                v2 -= v2
            case 42:
                v19 := GenS1{F0: 65535, F1: -593}
                _ = v19
            }
        }
        v20 := GenS3{F0: (!((!false) || false)), F1: 18446744073709551615}
        _ = fi17
        _ = v18
        _ = v20
    }
    v21 := GenS1{F0: 18446744073709551615, F1: -454}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = (v4 >> ((0) & 63))
    v22 := GenS1{F0: 2, F1: 175}
    _ = len(v3)
    v14 = 681
    var v23 float64 = 1e10
    if (!false) {
        v4 -= (v4 | v4)
        v23 += 0.1
    }
    fmt.Println(v23, v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
