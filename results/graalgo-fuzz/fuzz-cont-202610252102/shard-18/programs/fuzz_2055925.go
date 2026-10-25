package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
func genFn0(p0 uint16, p1 int8) uint8 {
    var v4 uint8 = 13
    var v5 uint8 = (v4 + v4)
    p0 += (p0 ^ p0)
    v5 = 1
    v6 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v5
}
func genFn1() uint {
    v7 := GenS2{F0: true, F1: 2, F2: 42}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint64 = 0
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: (-224 < -743), F1: 1, F2: 104}
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: true, F1: 255, F2: 42}
    _ = v12
    _ = v7
    return 590
}

func main() {
    var v1 uint = 39
    var v2 bool = (false && true)
    v3 := []int{-2147483648, 274}
    if (1 > 42) {
        fmt.Println(v1)
        switch 128 {
        case 127:
            if (-1 < 1) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        default:
            v13 := GenS1{F0: -778}
            _ = v13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (!((fi14 ^ fi14) > fi14)) {
                v2 = true
                var v15 uint8 = uint8(v1)
                _ = v15
            } else {
                v16 := GenS2{F0: (42 != 2), F1: 788, F2: 42}
                fmt.Println(fi14, v2)
                for fi17 := 0; fi17 < 6; fi17++ {
                    if v2 {
                        v2 = (v2 || v2)
                        switch "hello" {
                        case "foo bar":
                            v18 := GenS2{F0: (-2.25 >= (0.0 + float64(-2.25))), F1: 565, F2: 255}
                            _ = v18
                        case "\n\t":
                            var v19 bool = (!v2)
                            _ = v19
                        case "a":
                            v2 = v2
                        default:
                            switch "hello" {
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = fi17
                                }
                            default:
                                v20 := GenS1{F0: 1}
                                _ = v20
                            }
                        }
                        v21 := GenS1{F0: -1}
                        _ = v21
                    }
                    _ = fi17
                }
                _ = v16
            }
            var v22 int8 = -1
            _ = fi14
            _ = v22
        }
    } else {
        v2 = v2
        var v23 uint8 = 255
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = 128
    }
    if (!true) {
        v2 = true
    }
    v2 = (genFn0(758, 127) >= 124)
    fmt.Println(v1, v1)
    if v2 {
        v24 := GenS2{F0: v2, F1: 65535, F2: 0}
        _ = v24
    }
    fmt.Println(v1, v2)
    fmt.Println(v1, v2)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = (v1 - (v1 * (^(v1))))
    var v25 uint8 = 128
    fmt.Println(v1, v2)
    switch 42 {
    case 32768:
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint8 = (^(v25))
            _ = fi26
            _ = v27
        }
    case -108:
        v2 = false
    default:
        v1 = uint(v25)
    }
    v28 := GenS1{F0: 524}
    for fi29 := 0; fi29 < 5; fi29++ {
        v3 = append(v3, fi29)
        for fi30 := 0; fi30 < 7; fi30++ {
            var v31 int8 = 1
            var v32 int = (fi30 ^ fi30)
            _ = fi30
            _ = v31
            _ = v32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            fmt.Println(v2)
            var v34 float32 = 1e10
            _ = fi33
            _ = v34
        }
        _ = fi29
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint64 = 42
        for fi37 := 0; fi37 < 6; fi37++ {
            v25 = v25
            var v38 int16 = 1
            _ = fi37
            _ = v38
        }
        var v39 int = (fi35 ^ (fi35 & fi35))
        _ = fi35
        _ = v36
        _ = v39
    }
    v25 += (v25 | genFn0(240, 42))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
}
