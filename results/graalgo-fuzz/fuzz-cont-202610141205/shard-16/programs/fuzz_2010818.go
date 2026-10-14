package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16) bool {
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 += (p0 & p0)
        var v8 int64 = 370
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 307 {
    case 2147483647:
        if (1 > 255) {
            if true {
                for fi10 := 0; fi10 < 5; fi10++ {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        v12 := GenS1{F0: 359, F1: 74}
                        var v13 bool = true
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = fi10
                }
                fmt.Println(p0)
            }
            fmt.Println(p0)
            p0 += p0
        }
    case -459:
        v14 := []int{}
        _ = v14
    }
    if false {
        p0 = (p0 - (p0 | 32767))
        v15 := GenS2{F0: 2}
        _ = v15
    }
    p0 -= -169
    var v16 bool = (true || (65535 != 4294967295))
    _ = p0
    _ = v16
    return (1 >= 129)
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    var v3 uint = 585
    var v4 int16 = 1
    v5 := make(map[int]bool)
    var v6 bool = false
    var v17 uint64 = 4294967295
    fmt.Println(v1)
    if v6 {
        v18 := []int{25, 2147483647}
        v5[127] = false
        switch 337 {
        case 42:
            v2[32767] = 0
        case 253:
            var v19 uint16 = 1
            _ = v19
        default:
            v17 = (v17 & v17)
        }
        _ = v18
    }
    v3 = (v3 * 65535)
    fmt.Println(v4)
    if v6 {
        var v20 uint16 = uint16(v17)
        switch 829 {
        case 56:
            if (float32(1e10) == 0.0) {
                v20 = v20
                v20 = 972
            } else {
                switch 2147483647 {
                case -175:
                    v17 = v17
                case 1:
                    v21 := []int{42, -664}
                    _ = v21
                default:
                    v4 = (v4 - v4)
                }
                if v6 {
                    v22 := []int{9223372036854775807, -2147483648, 32767}
                    _ = v22
                }
            }
        case 42:
            v20 += v20
        default:
            var v23 uint64 = uint64(v3)
            _ = v23
        }
        _ = v20
    }
    fmt.Println(v3, v6)
    v6 = (("\n\t" != ("a" + "a")) && v6)
    v4 -= v4
    fmt.Println(v4, v17, v3)
    v24 := v5[32768]
    var v25 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
