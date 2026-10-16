package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint }
func genFn0(p0 uint8, p1 uint) int16 {
    var v6 int64 = -1
    v7 := GenS1{F0: 912}
    p1 -= 823
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 uint8, p1 int32) int16 {
    p1 -= (p1 - 463)
    p1 += p1
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 uint16 = 562
    var v2 uint64 = 2
    var v3 uint32 = 201
    v4 := []string{"", "foo bar", "\n\t"}
    var v5 uint = 65535
    v3 = (v3 + 4294967295)
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 bool = (42 < -1)
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 int8 = 42
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    switch "" {
    case "\n\t":
        for fi13 := 0; fi13 < 15; fi13++ {
            v1 = uint16(v5)
            _ = fi13
        }
    default:
        fmt.Println(v1, v1)
    }
    if true {
        var v14 uint64 = v2
        _ = v14
    }
    v15 := GenS1{F0: v5}
    switch 4294967295 {
    case 1:
        for fi16 := 0; fi16 < 6; fi16++ {
            if (464 != v2) {
                fmt.Println(v1, v5)
            } else {
                for fi17 := 0; fi17 < 20; fi17++ {
                    switch fi17 {
                    case 32768:
                        var v18 float64 = (v8 + v8)
                        _ = v18
                    case 1:
                        fmt.Println(v1, v1)
                    }
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    _ = fi17
                }
            }
            _ = fi16
        }
    default:
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 19; fi20++ {
                v4 = append(v4, "foo bar")
                _ = fi20
            }
            _ = fi19
        }
    }
    v21 := GenS1{F0: (v5 | (v5 + v5))}
    v3 -= (v3 & 2147483647)
    fmt.Println(v2, v3, v8)
    if (9223372036854775807 == -1) {
        var v22 uint8 = uint8(v1)
        var v23 int16 = 1000
        _ = v22
        _ = v23
    } else {
        if (true || (486 == 1000)) {
            var v24 float64 = v8
            v25 := GenS1{F0: uint(v3)}
            v8 = (0.1 * (v8 + v8))
            _ = v24
            _ = v25
        } else {
            var v26 uint16 = v1
            v27 := GenS2{F0: 65535, F1: (v5 << ((255) & 63))}
            _ = v26
            _ = v27
        }
        v3 -= uint32(v5)
        if false {
            var v28 uint = v5
            _ = v28
        }
    }
    var v29 float32 = 1.5
    var v30 int = 42
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
