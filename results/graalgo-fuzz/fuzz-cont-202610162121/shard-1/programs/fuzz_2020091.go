package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
func genFn0() int16 {
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint16 = 321
        v7 := GenS1{F0: 565, F1: -9223372036854775808, F2: -2147483648}
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "a" {
            case "":
                v4 -= v4
            }
            if false {
                v4 = v4
                v9 := []int{}
                if (-2.25 > 0.0) {
                    var v10 int = int(v4)
                    _ = v10
                }
                var v11 int64 = 0
                if len(v9) > 0 {
                    v9[0] = fi8
                }
                _ = len(v9)
                _ = v11
                _ = v9
            } else {
                fmt.Println(fi5)
                var v12 uint = 42
                _ = v12
            }
            v13 := GenS1{F0: (v4 - 30), F1: 364, F2: 517}
            _ = fi8
            _ = v13
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 127 {
    case 1:
        v4 = v4
    case 128:
        if true {
            var v14 int32 = -2147483648
            _ = v14
        } else {
            var v15 uint32 = 2
            _ = v15
        }
    case 127:
        v4 = v4
    default:
        var v16 uint32 = 815
        _ = v16
    }
    _ = v4
    return -32768
}
func genFn1(p0 int, p1 int16, p2 int16) int32 {
    if (411 >= 806) {
        p1 = (p1 | p1)
        if true {
            var v17 int16 = (^(p2))
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        }
    }
    var v19 uint8 = 1
    if false {
        var v20 float64 = 0.1
        _ = v20
    }
    var v21 int32 = 1
    for fi22 := 0; fi22 < 17; fi22++ {
        if false {
            v21 -= (v21 - (-(v21)))
            for fi23 := 0; fi23 < 6; fi23++ {
                v21 = 42
                _ = fi23
            }
            var v24 uint8 = (v19 >> ((65535) & 63))
            _ = v24
        } else {
            for fi25 := 0; fi25 < 4; fi25++ {
                p2 = -1
                _ = fi25
            }
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v21
    return -223
}

func main() {
    var v1 bool = false
    var v2 float64 = ((float64((((3.14159 - 0.1) * 1e10) * ((1.5 * float32((0.0 - 0.1))) + (((0.1 - 1.5) / 3.14159) - (float32(0.0) * 1e10))))) / 1.5) - (((((((0.1 - 3.14159) * (1.5 / float64((3.14159 * 1.5)))) - float64(1e10)) + 1.5) - 0.0) * 3.14159) + 0.1))
    var v3 uint16 = 42
    v3 += (v3 | v3)
    v1 = v1
    fmt.Println(v3, v2)
    switch 4294967295 {
    case 316:
        var v26 uint16 = (v3 / ((149) | 1))
        _ = v26
    default:
        v1 = false
    }
    fmt.Println(v1, v2)
    var v27 uint8 = 127
    v28 := []int{580}
    if len(v28) > 0 {
        v28[0] = -2147483648
    }
    var v29 int8 = 1
    var v30 int = 128
    fmt.Println(v29, v29, v29)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
