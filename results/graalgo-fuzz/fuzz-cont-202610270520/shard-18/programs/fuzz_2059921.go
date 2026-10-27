package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
func genFn0(p0 int32, p1 int64) int64 {
    p0 -= (p0 & p0)
    p0 = (^(p0))
    p0 = p0
    _ = p0
    _ = p1
    return (p1 - 9223372036854775807)
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 1.5
    v3 := make(map[uint32]int)
    var v4 float32 = 1.5
    v5 := make(map[int]int)
    var v6 int8 = 42
    for fi7 := 0; fi7 < 6; fi7++ {
        if (int(v6) != fi7) {
            var v8 int = fi7
            v9 := []int{v8, (v8 / (((fi7 * (v8 >> ((1) & 63)))) | 1)), -1}
            v8 -= (v8 * v8)
            _ = v8
            _ = v9
        } else {
            switch (fi7 | fi7) {
            case 4294967295:
                v6 -= (v6 & (v6 | (-(v6))))
            case 456:
                for fi10 := 0; fi10 < 1; fi10++ {
                    v11 := GenS1{F0: 588, F1: 0, F2: fi10}
                    fmt.Println(v6, v6)
                    _ = fi10
                    _ = v11
                }
            case 582:
                if (false && (!(false || true))) {
                    v5[fi7] = (fi7 / ((fi7) | 1))
                    if false {
                        var v12 int8 = (^(v6))
                        _ = v12
                    }
                }
            default:
                v6 += (-(v6))
            }
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int64 = 1
                var v15 uint16 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = fi7
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        v4 += 3.14159
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v6)
            _ = fi17
        }
        _ = fi16
    }
    if (((0.0 / v4) + (v4 / v4)) <= (((float32((float64(v4) / -2.25)) / v4) * 1e10) + v4)) {
        v18 := []int{80, 867}
        fmt.Println(v6, v4, v4)
        _ = v18
    }
    v4 = float32(-2.25)
    v4 -= v4
    var v19 uint = 0
    v6 -= 1
    switch ("hello" + "a") {
    case "foo bar":
        v4 -= (v4 - ((v4 / v4) + (v4 + v4)))
    default:
        v20 := GenS1{F0: (v19 | (v19 ^ v19)), F1: (v19 << ((626) & 63)), F2: 916}
        _ = v20
    }
    var v21 uint8 = 0
    v22 := []int{-2147483648}
    if ((v4 > v4) && false) {
        v6 += (v6 + (v6 * 1))
        fmt.Println(v4, v6, v2)
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
    } else {
        var v23 uint64 = 1
        v24 := GenS1{F0: v19, F1: v19, F2: 127}
        if false {
            if true {
                v21 = 1
                v22 = append(v22, 0)
                v1[2147483647] = 4294967295
            }
            for fi25 := 0; fi25 < 19; fi25++ {
                v4 = 0.0
                v4 = ((float32((v2 - (v2 - (((-2.25 - 0.0) / v2) * 3.14159)))) * (v4 + (v4 - (v4 - (0.1 / (v4 - ((v4 - float32(v2)) + v4))))))) / 0.1)
                _ = fi25
            }
            v26 := GenS1{F0: 2, F1: uint(v21), F2: 4294967295}
            _ = v26
        } else {
            v2 += 0.0
            var v27 int8 = v6
            var v28 uint8 = (v21 << (((^(v19))) & 63))
            _ = v27
            _ = v28
        }
        _ = v23
        _ = v24
    }
    switch 128 {
    case 813:
        switch 1 {
        case -548:
            var v29 uint64 = 42
            _ = v29
        case 127:
            v19 = v19
        case 42:
            var v30 uint32 = 2
            _ = v30
        default:
            for fi31 := 0; fi31 < 4; fi31++ {
                fmt.Println(v19)
                _ = fi31
            }
        }
    case -992:
        v19 -= (v19 >> (((v19 % (((v19 >> ((306) & 63))) | 1))) & 63))
    case 1:
        v4 += -2.25
    }
    v2 -= v2
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint = (v19 | 33)
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        _ = fi32
        _ = v33
    }
    delete(v5, 32768)
    var v34 int = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
