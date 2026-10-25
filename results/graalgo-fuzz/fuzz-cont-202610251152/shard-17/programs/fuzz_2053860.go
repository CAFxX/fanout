package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
func genFn0(p0 int, p1 bool, p2 int8) bool {
    p2 = p2
    if ((-2.25 * 1.5) != 1e10) {
        fmt.Println(p1, p2, p2)
    }
    fmt.Println(p0, p1)
    v5 := []int{p0, (p0 * p0), (p0 * (p0 - 4294967295)), (p0 | (p0 - (p0 - p0)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (!(!(p1 || (1 < 65535))))
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 131
    v3 := make(map[int]bool)
    v4 := []uint32{65535}
    var v6 float64 = 0.1
    v3[128] = (4294967295 < 1)
    var v7 int32 = 814
    v8 := GenS1{F0: -1, F1: 4294967295}
    v6 += 3.14159
    v1 -= v1
    v7 += 42
    v9 := GenS1{F0: 42, F1: 4294967295}
    var v10 uint64 = v2
    fmt.Println(v1, v10, v10)
    v11 := GenS1{F0: 0, F1: 42}
    if (false || false) {
        fmt.Println(v1, v1)
        if true {
            for fi12 := 0; fi12 < 14; fi12++ {
                switch "" {
                case "foo bar":
                    var v13 int16 = -1
                    _ = v13
                default:
                    v1 += 42
                }
                _ = fi12
            }
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int = int(v10)
                _ = fi14
                _ = v15
            }
        }
    }
    v4 = append(v4, 65535)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int8 = -85
        var v18 string = "\n\t"
        _ = fi16
        _ = v17
        _ = v18
    }
    switch -940 {
    case 871:
        v10 = (v2 / (((^(v10))) | 1))
    case 42:
        var v19 int16 = 1000
        _ = v19
    default:
        if (false && true) {
            fmt.Println(v1, v1)
            v20 := GenS2{F0: 127}
            _ = v20
        } else {
            v3[4294967295] = false
            switch -763 {
            case 0:
                if (!true) {
                    if ("" != ("a" + "a")) {
                        v7 = (-(v7))
                        for fi21 := 0; fi21 < 1; fi21++ {
                            v10 -= 829
                            _ = fi21
                        }
                    } else {
                        var v22 int16 = -32768
                        var v23 uint64 = v10
                        var v24 int = 112
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    switch 461 {
                    case -9223372036854775808:
                        var v25 uint = 4294967295
                        _ = v25
                    }
                    _ = len(v4)
                }
            case -1:
                var v26 int16 = -32768
                _ = v26
            }
        }
    }
    v6 = (((v6 - 1.5) / 1.0) / 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
