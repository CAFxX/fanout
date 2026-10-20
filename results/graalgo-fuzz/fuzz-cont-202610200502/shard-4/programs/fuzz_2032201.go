package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = ((!false) || (!false))
    var v3 uint8 = 42
    v4 := make(map[int]uint32)
    var v5 uint32 = 581
    if ((187 >= 42) && true) {
        v6 := []int{-253, -2147483648}
        var v7 string = "hello"
        _ = v6
        _ = v7
    } else {
        var v8 float32 = float32((0.1 + 1e10))
        var v9 int32 = 1
        _ = v8
        _ = v9
    }
    v10 := []int{-1, 127}
    if v2 {
        var v11 int = -592
        if v2 {
            v12 := GenS1{F0: (65535 > 255)}
            _ = v12
        } else {
            if true {
                v11 += (v11 / ((v11) | 1))
            }
            var v13 uint16 = 964
            for fi14 := 0; fi14 < 13; fi14++ {
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 bool = true
                    v13 += v13
                    _ = fi15
                    _ = v16
                }
                if len(v10) > 0 {
                    v10[0] = (-(fi14))
                }
                _ = fi14
            }
            _ = v13
        }
        var v17 float32 = ((1e10 / 1.0) - 0.0)
        _ = v11
        _ = v17
    } else {
        if len(v10) > 0 {
            v10[0] = 2147483647
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        switch ("hello" + ("a" + (("a" + "") + "foo bar"))) {
        case "hello":
            v5 = (v5 + v5)
        case "":
            switch (fi18 + -335) {
            case 42:
                if (78 > -128) {
                    var v19 uint32 = v5
                    v10 = append(v10, fi18)
                    _ = v19
                }
            case 127:
                if v2 {
                    var v20 uint16 = 553
                    _ = v20
                }
            case -263:
                v21 := GenS1{F0: ((!(0 >= 1)) && v2)}
                _ = v21
            default:
                v22 := GenS2{F0: ((-(fi18)) >= (-(fi18))), F1: 0}
                _ = v22
            }
        case "foo bar":
            var v23 int = 0
            _ = v23
        }
        _ = fi18
    }
    v4[1] = (^(v5))
    v24 := GenS2{F0: false, F1: 4294967295}
    var v25 int = -1
    if false {
        var v26 uint32 = (v5 % ((138) | 1))
        _ = v26
    }
    v25 -= 2147483647
    var v27 int32 = -1
    v28 := v1[(v25 / (((v25 * (v25 * (v25 >> ((255) & 63))))) | 1))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
