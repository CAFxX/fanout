package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 42
    v3 := make(map[uint32]bool)
    var v4 float64 = 0.0
    var v5 int32 = 1
    v6 := GenS1{F0: 127, F1: 399, F2: 375}
    for fi7 := 0; fi7 < 11; fi7++ {
        switch (^(fi7)) {
        case -847:
            v5 = v5
        case 689:
            fmt.Println(fi7)
        case 462:
            if false {
                var v8 float32 = -2.25
                var v9 int = 0
                _ = v8
                _ = v9
            }
        default:
            var v10 uint32 = 2147483648
            _ = v10
        }
        _ = fi7
    }
    v11 := []int{}
    for fi12 := 0; fi12 < 14; fi12++ {
        v11 = append(v11, (fi12 - (fi12 >> ((1) & 63))))
        if false {
            v4 = ((v4 * ((v4 + (float64(0.1) - 3.14159)) - (((v4 + 1e10) * float64(((-2.25 / -2.25) - (1e10 * 0.0)))) + (0.0 / v4)))) - (v4 / v4))
            v11 = append(v11, 0)
            for fi13 := 0; fi13 < 18; fi13++ {
                switch fi12 {
                case 1:
                    if (true && ((float64((0.0 - (0.1 * 0.0))) + (0.1 / -2.25)) >= 0.0)) {
                        v14 := GenS1{F0: 74, F1: 875, F2: 186}
                        _ = v14
                    }
                default:
                    v15 := GenS1{F0: 2, F1: uint32(v2), F2: 42}
                    v16 := GenS1{F0: 128, F1: 4294967295, F2: 42}
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
        } else {
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 int = (fi17 >> ((uint(v1)) & 63))
                v19 := v3[4294967295]
                var v20 uint8 = 255
                var v21 int8 = 1
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 8; fi22++ {
                if false {
                    v11 = append(v11, (fi12 << ((2) & 63)))
                    if true {
                        v5 += 42
                        var v23 uint8 = 255
                        var v24 int8 = -1
                        v4 = float64(3.14159)
                        var v25 uint32 = 2
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
                _ = len(v3)
                _ = fi22
            }
            var v26 int = (fi12 ^ fi12)
            _ = v26
        }
        _ = fi12
    }
    var v27 float32 = 0.1
    v11 = append(v11, 1)
    v1 = v1
    if true {
        if (-2147483648 > 32768) {
            v1 -= 1
            v4 -= ((((v4 - v4) - v4) * v4) / (v4 / (v4 - ((v4 - (float64(-2.25) + float64(1e10))) * v4))))
        }
        v4 = v4
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        if (!((^(v5)) > v5)) {
            for fi29 := 0; fi29 < 9; fi29++ {
                var v30 int64 = -1
                if len(v11) > 0 {
                    v11[0] = (fi28 ^ fi28)
                }
                _ = fi29
                _ = v30
            }
            fmt.Println(v5)
            if (false || (1 > -716)) {
                v2 = v2
                var v31 uint = 2
                v1 = v1
                _ = v31
            } else {
                for fi32 := 0; fi32 < 2; fi32++ {
                    v2 = v2
                    fmt.Println(fi32, v2)
                    switch "" {
                    case "\n\t":
                        fmt.Println(v5, v27)
                    case "":
                        v33 := v3[65535]
                        _ = v33
                    default:
                        _ = len(v11)
                    }
                    _ = fi32
                }
                var v34 int32 = v5
                _ = v34
            }
        }
        _ = fi28
    }
    v3[2] = (1 == 831)
    v27 = float32(1.5)
    v11 = append(v11, -806)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
