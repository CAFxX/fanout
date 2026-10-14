package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0(p0 uint8) int16 {
    if true {
        if (0 < -1) {
            v6 := []int{9223372036854775807}
            v6 = append(v6, 32767)
            _ = v6
        } else {
            p0 = (p0 ^ p0)
            var v7 int = 32767
            _ = v7
        }
        if false {
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 string = "a"
                    var v11 int16 = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
        } else {
            v12 := []int{-9223372036854775808, 958, 613}
            _ = v12
        }
        switch 128 {
        case 32767:
            fmt.Println(p0)
        case -655:
            for fi13 := 0; fi13 < 1; fi13++ {
                p0 -= p0
                _ = fi13
            }
        case 1:
            p0 -= (p0 ^ p0)
        default:
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = 1
                switch (fi14 + fi14) {
                case 42:
                    v16 := []int{fi14}
                    _ = v16
                }
                _ = fi14
                _ = v15
            }
        }
    } else {
        fmt.Println(p0)
        var v17 int64 = -1
        _ = v17
    }
    var v18 uint = 2
    v19 := []int{32767, 70, 9223372036854775807, 938}
    _ = p0
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]bool)
    var v3 float32 = (0.0 / 1e10)
    var v4 int32 = -1
    v5 := make(map[uint32]uint32)
    v1 += v1
    _ = len(v2)
    for fi20 := 0; fi20 < 13; fi20++ {
        v4 += -1
        var v21 uint8 = 128
        _ = fi20
        _ = v21
    }
    var v22 int16 = -1
    var v23 uint = v1
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := []int{fi24, (-(fi24)), fi24}
        for fi26 := 0; fi26 < 14; fi26++ {
            v23 = v1
            if ((fi24 < fi26) || (!(v22 != 1000))) {
                v27 := v2[int(v23)]
                var v28 float64 = (-2.25 * -2.25)
                fmt.Println(v4)
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    switch "" {
    case "":
        delete(v5, 348)
    case "foo bar":
        for fi29 := 0; fi29 < 11; fi29++ {
            v30 := GenS1{F0: (-(v22)), F1: 127, F2: -1}
            v31 := []int{}
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    v23 = v1
    v1 = v23
    _ = len(v2)
    v4 = v4
    _ = len(v2)
    var v32 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
