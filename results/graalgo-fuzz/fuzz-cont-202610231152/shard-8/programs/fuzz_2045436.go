package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint32 = 42
    var v2 uint64 = 368
    var v3 int16 = 683
    v4 := []int{1, -2147483648, -670}
    _ = len(v4)
    v3 -= v3
    switch -498 {
    case 0:
        var v5 int64 = 1
        _ = v5
    case -9223372036854775808:
        switch 32767 {
        case 9223372036854775807:
            if (!(false || false)) {
                var v6 int = 42
                fmt.Println(v1, v3)
                _ = v6
            }
        case 2147483647:
            for fi7 := 0; fi7 < 11; fi7++ {
                for fi8 := 0; fi8 < 17; fi8++ {
                    fmt.Println(v2, fi7, v1)
                    v3 = (v3 - 1)
                    v4 = append(v4, fi7)
                    v3 = v3
                    var v9 int32 = 969
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
        case 596:
            if (477 != 32768) {
                var v10 int32 = 451
                _ = len(v4)
                var v11 int8 = 110
                _ = v10
                _ = v11
            }
        default:
            if (uint8(v1) < 218) {
                if false {
                    var v12 uint64 = v2
                    v12 = v2
                    var v13 uint32 = uint32(v12)
                    _ = v12
                    _ = v13
                }
                var v14 float32 = 3.14159
                _ = v14
            }
        }
    }
    if (!true) {
        _ = len(v4)
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint16 = 812
            var v17 int8 = 80
            _ = fi15
            _ = v16
            _ = v17
        }
        v3 -= (v3 | (-(v3)))
    }
    if (false && true) {
        var v18 float32 = 0.1
        if (true || (true || false)) {
            switch 351 {
            case 32768:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                if false {
                    var v19 uint16 = 255
                    v20 := GenS1{F0: -972, F1: -1, F2: v3}
                    for fi21 := 0; fi21 < 17; fi21++ {
                        if false {
                            v22 := GenS2{F0: (^(v1))}
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v4 = append(v4, fi23)
                                switch (fi23 * fi21) {
                                case 0:
                                    if ((fi23 >= fi23) || true) {
                                        var v24 uint8 = 127
                                        var v25 uint8 = uint8(v1)
                                        _ = v24
                                        _ = v25
                                    }
                                default:
                                    v3 = (-(v3))
                                }
                                _ = fi23
                            }
                            v4 = append(v4, fi21)
                            _ = v22
                        } else {
                            v18 = (1e10 * v18)
                            v18 = (v18 / v18)
                        }
                        _ = fi21
                    }
                    _ = v19
                    _ = v20
                }
            }
            var v26 uint16 = 0
            if len(v4) > 0 {
                v4[0] = 127
            }
            _ = v26
        }
        _ = v18
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        fmt.Println(v3, v3, v3)
        var v28 float32 = (1e10 * (3.14159 + -2.25))
        var v29 uint64 = 1
        _ = fi27
        _ = v28
        _ = v29
    }
    switch ("hello" + "") {
    case "foo bar":
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 uint8 = 42
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = v3
            _ = fi30
            _ = v31
        }
    case "\n\t":
        v4 = append(v4, -933)
    }
    v1 = (^(v1))
    v3 = v3
    v32 := GenS2{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
