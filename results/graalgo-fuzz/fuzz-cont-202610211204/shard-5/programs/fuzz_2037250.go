package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 4294967295
    var v3 int16 = -32768
    v4 := make(map[uint32]string)
    v5 := []string{"a", ("foo bar" + "\n\t"), "hello", "hello"}
    var v6 uint32 = 2147483647
    for fi7 := 0; fi7 < 15; fi7++ {
        if (!false) {
            v3 = (-(v3))
            for fi8 := 0; fi8 < 2; fi8++ {
                v4[(^(v6))] = "\n\t"
                switch -1 {
                case -486:
                    v2 -= (v2 >> (((v2 & v2)) & 63))
                default:
                    var v9 int32 = -43
                    v9 -= v9
                    var v10 uint16 = 0
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 float64 = 1e10
                fmt.Println(v3)
                _ = fi11
                _ = v12
            }
            if false {
                var v13 uint = (v2 - (^(v2)))
                var v14 bool = false
                _ = v13
                _ = v14
            }
            v6 += (v6 + (v6 << ((v2) & 63)))
        }
        fmt.Println(v2, v3, fi7)
        v15 := GenS1{F0: v3, F1: 42, F2: ((v2 * v2) <= 685)}
        _ = fi7
        _ = v15
    }
    fmt.Println(v6)
    switch 32768 {
    case 32767:
        var v16 uint16 = 2
        _ = v16
    case -9223372036854775808:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        _ = len(v4)
    }
    v1["hello"] = 2147483648
    var v17 float32 = (-2.25 + ((-2.25 / 0.1) + -2.25))
    var v18 int32 = -335
    v19 := GenS1{F0: v3, F1: v18, F2: (1e10 < (1.5 / 3.14159))}
    fmt.Println(v3, v2)
    if false {
        for fi20 := 0; fi20 < 11; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                v2 += 59
                delete(v4, v6)
                _ = fi21
            }
            _ = fi20
        }
    }
    v3 += v3
    fmt.Println(v17)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
