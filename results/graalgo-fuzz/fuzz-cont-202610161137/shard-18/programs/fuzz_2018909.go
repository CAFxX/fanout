package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
func genFn0(p0 uint32) uint32 {
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 uint16 = 1
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint8 = 128
    for fi11 := 0; fi11 < 14; fi11++ {
        if true {
            v12 := []int{}
            if len(v12) > 0 {
                v12[0] = 2147483647
            }
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint8 = v10
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(fi13, v10)
            if false {
                if (1.5 == 1e10) {
                    var v17 int64 = -443
                    v17 -= v17
                    _ = v17
                }
            }
            v10 -= v10
            _ = fi16
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            switch "" {
            case "\n\t":
                p0 += 5
            case "hello":
                switch -9223372036854775808 {
                case -2147483648:
                    var v19 uint = 556
                    _ = v19
                }
            default:
                if (!true) {
                    v20 := []int{}
                    _ = v20
                }
            }
            _ = fi18
        }
        _ = fi13
    }
    p0 = p0
    _ = p0
    _ = v10
    return 136
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 127
    v3 := []int{19}
    var v4 float64 = -2.25
    v5 := make(map[int]uint32)
    var v6 float64 = v4
    v21 := GenS1{F0: 42, F1: 127, F2: true}
    var v22 uint = 255
    v22 = (v22 & (v22 << ((1) & 63)))
    var v23 int16 = -32768
    var v24 uint16 = 32767
    var v25 uint = (v22 % ((18446744073709551615) | 1))
    fmt.Println(v25, v25)
    v26 := GenS1{F0: 42, F1: -1, F2: (-1 == -9223372036854775808)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
