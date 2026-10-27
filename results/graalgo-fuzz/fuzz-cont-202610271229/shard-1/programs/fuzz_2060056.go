package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() int16 {
    var v6 float32 = 0.1
    v7 := GenS1{F0: -1, F1: 42, F2: 32767}
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int64 = 0
    var v2 int64 = v1
    var v3 uint32 = 0
    var v4 int32 = 2147483647
    v5 := make(map[int]uint32)
    v3 -= v3
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := []int{(^(fi8)), fi8, (-(fi8))}
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    v10 := []int{}
    v1 -= v2
    fmt.Println(v1, v4)
    if (v4 == (-(v4))) {
        var v11 uint64 = 0
        v1 += -9223372036854775808
        switch "" {
        case "":
            v2 += -9223372036854775808
        }
        _ = v11
    }
    if (1.5 < (-2.25 - (1.5 / 1e10))) {
        fmt.Println(v1)
        if false {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            switch 2147483647 {
            case 32767:
                v1 = (-(v1))
            case -1:
                if true {
                    if true {
                        v12 := GenS1{F0: (v1 * (-(v1))), F1: 18446744073709551615, F2: -9223372036854775808}
                        v2 = (v2 << ((2) & 63))
                        _ = v12
                    }
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 float64 = (((-2.25 / 0.1) / float64(-2.25)) / -2.25)
                        v14 -= ((1.5 - float64(float32(v14))) * v14)
                        v15 := v5[42]
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                }
            case 127:
                var v16 string = ("foo bar" + "hello")
                _ = v16
            }
            v17 := GenS2{F0: true, F1: 537}
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
