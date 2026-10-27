package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool, p1 uint32) uint8 {
    if p0 {
        var v4 uint = 255
        _ = v4
    } else {
        var v5 uint8 = 0
        _ = v5
    }
    v6 := []int{641, 174, 2147483647}
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            switch "hello" {
            case "hello":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
        } else {
            v8 := GenS1{F0: 2147483647, F1: 18446744073709551615}
            for fi9 := 0; fi9 < 3; fi9++ {
                if (false && p0) {
                    p0 = p0
                    var v10 uint16 = 0
                    _ = v10
                }
                _ = fi9
            }
            _ = v8
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (2 <= 180) {
            p1 = (^(p1))
            p0 = (0.0 <= float32(((0.1 - 3.14159) * 0.0)))
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 201
}

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 149
    var v3 uint64 = 923
    var v11 float32 = (1.5 + 0.1)
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{(fi12 << ((42) & 63)), -9223372036854775808}
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 bool = false
            v3 += v3
            if (!((0 <= (^(v1))) || ((!(!(!v15))) || v15))) {
                _ = len(v13)
            } else {
                fmt.Println(fi14, v1)
                var v16 int16 = 1000
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v17 := []int{-383, 32768}
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 int8 = 100
        _ = fi18
        _ = v19
    }
    v11 = (v11 / v11)
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = len(v17)
    switch 662 {
    case 4294967295:
        if len(v17) > 0 {
            v17[0] = 127
        }
    case 9223372036854775807:
        var v20 int64 = -1
        _ = v20
    default:
        switch 9223372036854775807 {
        case 32768:
            if (100 >= 0) {
                _ = len(v17)
                v17 = append(v17, 42)
                v21 := GenS1{F0: -671, F1: (v3 - (v3 - (v3 | v3)))}
                _ = v21
            } else {
                var v22 int16 = 29
                _ = v22
            }
        default:
            if (false && (127 > -1)) {
                v3 += 1
                switch "a" {
                case "a":
                    v11 -= v11
                default:
                    v23 := GenS2{F0: genFn0((!((^(v2)) > (v2 | v2))), (v1 + (v1 / ((v1) | 1)))), F1: 326}
                    _ = v23
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
}
