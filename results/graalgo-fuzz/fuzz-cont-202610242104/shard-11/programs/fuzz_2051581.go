package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }

func main() {
    var v1 int64 = 1
    var v2 uint32 = 0
    v3 := make(map[uint32]bool)
    var v4 uint32 = 42
    fmt.Println(v2, v4, v2)
    var v5 int8 = 127
    fmt.Println(v4)
    v6 := GenS1{F0: 60, F1: 0, F2: 1}
    var v7 int64 = (-(v1))
    v4 = 2147483647
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 string = (("foo bar" + "foo bar") + ("hello" + ""))
            var v11 uint8 = uint8(v4)
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 += (v2 * (v2 + v4))
        v7 -= v1
        _ = fi8
    }
    v3[v4] = (-2.25 > (((1e10 / 1.0) / (1e10 / -2.25)) - (0.1 / 1e10)))
    if (((float64(((-2.25 - float32(1e10)) - 3.14159)) + (float64(3.14159) / -2.25)) - 0.1) != 0.1) {
        if (!false) {
            v12 := []int{-9223372036854775808}
            _ = v12
        }
    } else {
        v7 = v7
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v2)
                var v15 int32 = 68
                _ = fi14
                _ = v15
            }
            v16 := GenS1{F0: 86, F1: 146, F2: fi13}
            _ = fi13
            _ = v16
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint16 = 592
            v2 = (v2 | v4)
            switch (("" + ("a" + "\n\t")) + "\n\t") {
            case "a":
                for fi19 := 0; fi19 < 6; fi19++ {
                    v5 -= v5
                    _ = fi19
                }
            }
            _ = fi17
            _ = v18
        }
    }
    switch 387 {
    case 136:
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS1{F0: -746, F1: 102, F2: 32768}
            var v22 bool = true
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        var v23 uint = 42
        _ = v23
    }
    switch "\n\t" {
    case "\n\t":
        if false {
            var v24 int8 = (v5 ^ (v5 * (v5 ^ 127)))
            fmt.Println(v1)
            var v25 int64 = v1
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 int32 = 42
                fmt.Println(v4)
                v5 = (v5 & (v5 - -34))
                _ = fi26
                _ = v27
            }
            if true {
                for fi28 := 0; fi28 < 17; fi28++ {
                    var v29 uint = 373
                    for fi30 := 0; fi30 < 7; fi30++ {
                        v31 := v3[(v2 | 1)]
                        var v32 uint = (v29 >> ((v29) & 63))
                        switch 32767 {
                        case -10:
                            if (fi28 != -782) {
                                v33 := GenS2{F0: (false || v31), F1: 163, F2: v29}
                                _ = v33
                            }
                        case 128:
                            var v34 float32 = 0.1
                            _ = v34
                        case -215:
                            var v35 int16 = 1
                            _ = v35
                        }
                        v29 += v32
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    v36 := []int{(fi28 / ((0) | 1)), fi28}
                    _ = fi28
                    _ = v29
                    _ = v36
                }
            }
        }
    }
    v37 := []int{}
    v38 := GenS2{F0: (65535 >= 65535), F1: 1, F2: 2}
    v1 += (v7 * -9223372036854775808)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
