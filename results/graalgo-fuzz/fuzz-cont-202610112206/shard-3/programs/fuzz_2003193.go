package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64 }
func genFn0() uint16 {
    var v6 uint8 = 35
    if false {
        switch -817 {
        case -641:
            v6 = (v6 ^ 2)
        case 32768:
            v6 = (v6 * 1)
        default:
            var v7 uint16 = 42
            _ = v7
        }
        v8 := GenS1{F0: 176, F1: 0, F2: -9223372036854775808}
        v6 = (v6 * (v6 << ((65535) & 63)))
        _ = v8
    }
    v6 += (^(v6))
    fmt.Println(v6)
    var v9 uint32 = 42
    _ = v6
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{804, 530, 4294967295, 545}
    v2 := []int{127, 613}
    var v3 bool = ((false || true) && (32 < 128))
    v4 := []int{444, -966, 42}
    var v5 int16 = 1000
    v5 = (-(v5))
    v3 = (441 == 0)
    var v10 int16 = (v5 / ((42) | 1))
    var v11 float64 = 0.1
    var v12 int64 = 0
    v13 := GenS2{F0: 9223372036854775807}
    if (127 < 128) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v5 = v10
            var v15 int64 = v12
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int8 = -1
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    } else {
        v5 -= (v5 + -32768)
        for fi18 := 0; fi18 < 7; fi18++ {
            if (18446744073709551615 == 42) {
                v19 := GenS1{F0: 143, F1: 42, F2: -1}
                var v20 uint32 = 791
                _ = v19
                _ = v20
            }
            switch 128 {
            case 1:
                if (fi18 != 32768) {
                    v3 = (!true)
                    _ = len(v2)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = len(v4)
                }
            default:
                var v21 int = -2147483648
                _ = v21
            }
            fmt.Println(v12, v10, v12)
            _ = fi18
        }
        v22 := GenS1{F0: 170, F1: 2, F2: v12}
        _ = v22
    }
    if false {
        _ = len(v4)
        v23 := GenS2{F0: (v12 ^ 42)}
        _ = v23
    } else {
        if true {
            var v24 uint64 = 1
            for fi25 := 0; fi25 < 4; fi25++ {
                var v26 uint8 = 255
                fmt.Println(v11, v11)
                _ = fi25
                _ = v26
            }
            _ = v24
        }
        _ = len(v1)
    }
    v5 += v10
    switch -734 {
    case 2147483647:
        fmt.Println(v3)
    case -1:
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS2{F0: (v12 + v12)}
            v10 = 1
            v12 = v12
            _ = fi27
            _ = v28
        }
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v29 int64 = (^(v12))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
