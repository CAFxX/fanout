package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
type GenS4 struct { F0 int32; F1 uint; F2 string }
func genFn0(p0 int) string {
    p0 -= (p0 | (p0 ^ (p0 / ((p0) | 1))))
    v6 := []int{-121, 617}
    _ = p0
    _ = v6
    return ""
}

func main() {
    v1 := []string{}
    var v2 uint = 298
    var v3 bool = true
    v4 := make(map[uint32]bool)
    var v5 int16 = 1000
    if v3 {
        if (!(!false)) {
            var v7 float32 = float32(((1e10 * 0.1) - float64(-2.25)))
            _ = v7
        }
        if v3 {
            for fi8 := 0; fi8 < 6; fi8++ {
                fmt.Println(v3, v5)
                _ = fi8
            }
        } else {
            var v9 uint32 = 2147483647
            if true {
                v10 := GenS1{F0: 4294967295, F1: -19, F2: 128}
                var v11 uint = (^(v2))
                var v12 uint8 = 1
                _ = v10
                _ = v11
                _ = v12
            }
            _ = v9
        }
        v4[398] = v3
    }
    switch 523 {
    case 42:
        switch int(v5) {
        case 981:
            var v13 uint8 = 128
            _ = v13
        case -2147483648:
            v14 := GenS1{F0: 507, F1: int8(v5), F2: 237}
            _ = v14
        }
    case 9223372036854775807:
        v15 := GenS2{F0: (42 < 1)}
        _ = v15
    default:
        v16 := GenS3{F0: 32767, F1: 178, F2: 665}
        _ = v16
    }
    v5 = -1
    fmt.Println(v5)
    delete(v4, 4294967295)
    var v17 int8 = 100
    switch 32767 {
    case 315:
        var v18 int32 = 624
        _ = v18
    case 9223372036854775807:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case -789:
        for fi19 := 0; fi19 < 15; fi19++ {
            if v3 {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 bool = v3
                    var v22 uint8 = 42
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 int8 = -8
                    v25 := GenS4{F0: 2147483647, F1: 625, F2: "foo bar"}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                var v26 int32 = -399
                _ = v26
            }
            v2 += 0
            v1 = append(v1, "a")
            _ = fi19
        }
    default:
        v27 := GenS1{F0: 42, F1: 64, F2: 1}
        _ = v27
    }
    v28 := v4[691]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
