package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64) int16 {
    switch 42 {
    case -852:
        v6 := GenS1{F0: 731, F1: 4294967295}
        _ = v6
    case -9223372036854775808:
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: int64(fi7), F1: 663}
            var v9 string = "hello"
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        var v10 string = ((("" + "") + "") + "")
        _ = v10
    }
    switch ("" + "a") {
    case "hello":
        for fi11 := 0; fi11 < 5; fi11++ {
            if (!true) {
                v12 := GenS1{F0: 42, F1: 4294967295}
                v13 := []int{(fi11 ^ fi11), 127, (fi11 + (^(fi11)))}
                fmt.Println(p0, p0)
                _ = v12
                _ = v13
            }
            _ = fi11
        }
    case "a":
        var v14 bool = (false || false)
        _ = v14
    default:
        var v15 uint16 = 1
        _ = v15
    }
    fmt.Println(p0)
    switch 128 {
    case 32767:
        switch 9223372036854775807 {
        case 42:
            for fi16 := 0; fi16 < 20; fi16++ {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(p0, p0)
                default:
                    fmt.Println(fi16, p0)
                }
                var v17 int16 = -170
                p0 += (p0 % ((p0) | 1))
                _ = fi16
                _ = v17
            }
        }
    case 415:
        v18 := []int{-841, 769, 1, 1}
        _ = v18
    }
    _ = p0
    return 1000
}

func main() {
    var v1 uint64 = 845
    v2 := []uint32{91, 776}
    v3 := []uint32{42, 0, 859, 4294967295}
    v4 := make(map[string]bool)
    var v5 uint16 = 2
    fmt.Println(v5, v5)
    var v19 uint8 = 127
    fmt.Println(v5, v19)
    fmt.Println(v5, v19, v19)
    for fi20 := 0; fi20 < 9; fi20++ {
        if true {
            var v21 int16 = -32768
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint16 = uint16(v1)
                _ = fi22
                _ = v23
            }
            if false {
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS2{F0: (v5 ^ (v5 * 821))}
                    var v26 float64 = (((1e10 + 0.0) / -2.25) - (float64((float32((0.1 * 0.0)) * 1.5)) - -2.25))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                var v27 int16 = int16(fi20)
                var v28 int32 = 0
                _ = v27
                _ = v28
            }
            _ = v21
        } else {
            v29 := GenS1{F0: 42, F1: 460}
            v2 = append(v2, 191)
            _ = v29
        }
        for fi30 := 0; fi30 < 10; fi30++ {
            v1 += (^(v1))
            var v31 bool = true
            _ = fi30
            _ = v31
        }
        _ = fi20
    }
    v19 += v19
    fmt.Println(v1)
    v19 -= 127
    var v32 uint = 2
    var v33 uint = (v32 % ((2) | 1))
    switch 2147483647 {
    case 2147483647:
        v3 = append(v3, 4294967295)
    }
    switch 839 {
    case -317:
        v1 -= (^(v1))
    case 128:
        if false {
            if (-2.25 < 1e10) {
                var v34 float64 = 1e10
                _ = v34
            }
        }
    default:
        v32 = (v33 & 65535)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
