package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint; F2 uint64 }
func genFn0(p0 bool, p1 uint8) int64 {
    var v4 int = -789
    if (!((!false) && true)) {
        v4 = v4
        v4 -= (v4 % (((v4 % (((^(v4))) | 1))) | 1))
        for fi5 := 0; fi5 < 13; fi5++ {
            p1 = (p1 - (p1 | p1))
            _ = fi5
        }
    }
    if p0 {
        p1 -= (^(p1))
        fmt.Println(v4, v4, p0)
        var v6 bool = p0
        _ = v6
    } else {
        var v7 int32 = 1
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint = 18446744073709551615
            var v10 int16 = 129
            fmt.Println(v9, v10)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint8, p1 uint, p2 bool) uint8 {
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 6; fi12++ {
        p0 = (p0 - (p0 ^ 0))
        for fi13 := 0; fi13 < 1; fi13++ {
            if (32767 == 1) {
                var v14 bool = (-128 <= 89)
                _ = v14
            }
            switch 32768 {
            case 127:
                var v15 int32 = -638
                _ = v15
            case 0:
                var v16 int16 = -888
                fmt.Println(v11)
                _ = v16
            default:
                switch "foo bar" {
                case "\n\t":
                    v17 := []int{(fi12 / ((fi13) | 1))}
                    _ = v17
                case "foo bar":
                    v18 := GenS1{F0: -404, F1: v11}
                    _ = v18
                case "a":
                    var v19 int16 = -184
                    _ = v19
                default:
                    var v20 uint16 = 0
                    _ = v20
                }
            }
            _ = fi13
        }
        _ = fi12
    }
    var v21 uint32 = v11
    var v22 uint = p1
    v11 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v21
    _ = v22
    return p0
}

func main() {
    var v1 int16 = -32768
    var v2 float64 = 1.5
    v3 := []int{128, -464}
    v1 = v1
    v23 := GenS2{F0: (("foo bar" + "") + ("\n\t" + ("" + "\n\t"))), F1: 2, F2: 4294967295}
    _ = len(v3)
    v24 := GenS1{F0: -1, F1: 229}
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int8 = 119
        if true {
            if (true || true) {
                v2 = v2
                fmt.Println(v26, v1)
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v26, fi25, v2)
                    _ = fi27
                }
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi25
        _ = v26
    }
    v28 := GenS1{F0: 346, F1: 65535}
    v1 += (-(v1))
    v2 += v2
    var v29 uint64 = 2
    var v30 float32 = 1.5
    switch 32767 {
    case 32767:
        v29 = (v29 * v29)
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
