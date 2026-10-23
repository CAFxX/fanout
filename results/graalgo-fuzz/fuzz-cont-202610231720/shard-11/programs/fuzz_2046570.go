package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 610
    var v2 int = -1
    v3 := []uint32{24, 2147483648, 5, 42}
    var v4 int64 = -9223372036854775808
    if false {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 string = ""
            _ = fi5
            _ = v6
        }
        _ = len(v3)
    } else {
        fmt.Println(v4, v1, v4)
        for fi7 := 0; fi7 < 15; fi7++ {
            fmt.Println(v4, v1, v4)
            var v8 uint = 493
            _ = fi7
            _ = v8
        }
        v4 = (v4 * v4)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v1)
    switch "" {
    case "":
        v9 := GenS1{F0: 32767, F1: 18446744073709551615, F2: 0}
        _ = v9
    case "hello":
        switch "foo bar" {
        case "a":
            v10 := GenS1{F0: 1, F1: 4294967295, F2: 18446744073709551615}
            _ = v10
        default:
            if false {
                v2 = v2
                for fi11 := 0; fi11 < 13; fi11++ {
                    v2 = (v2 % ((fi11) | 1))
                    if (1 > -1) {
                        if true {
                            for fi12 := 0; fi12 < 14; fi12++ {
                                switch "hello" {
                                case "":
                                    var v13 int8 = 127
                                    _ = v13
                                }
                                fmt.Println(v4)
                                _ = fi12
                            }
                            for fi14 := 0; fi14 < 8; fi14++ {
                                switch 0 {
                                case 9223372036854775807:
                                    var v15 float32 = 1.5
                                    _ = v15
                                }
                                switch (-(v2)) {
                                case -153:
                                    v1 = (^(v1))
                                default:
                                    var v16 int = -2147483648
                                    _ = v16
                                }
                                _ = fi14
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch (v2 | (v2 + v2)) {
                        case 0:
                            v2 = (^(fi11))
                        case 42:
                            v1 -= v1
                        case 128:
                            var v17 float32 = 0.0
                            _ = v17
                        default:
                            var v18 uint16 = 32767
                            _ = v18
                        }
                        var v19 int = (v2 & v2)
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 14; fi20++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v4, v2)
                            var v21 float64 = (-2.25 * (-2.25 - 0.1))
                            _ = fi20
                            _ = v21
                        }
                        var v22 float32 = (3.14159 / (1e10 / (0.1 + (float32(float64((((float32(3.14159) - 0.0) / 3.14159) - 0.1))) / 1.5))))
                        v4 -= v4
                        _ = v22
                    }
                    _ = fi11
                }
            }
        }
    }
    v2 = (v2 >> ((105) & 63))
    v23 := GenS2{F0: 992}
    switch (v2 / ((9223372036854775807) | 1)) {
    case -2147483648:
        var v24 int = (v2 ^ (v2 * v2))
        _ = v24
    default:
        var v25 int = v2
        _ = v25
    }
    var v26 int64 = (v4 & (v4 << ((255) & 63)))
    var v27 uint32 = 729
    var v28 uint16 = 861
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := GenS2{F0: 937}
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: (v28 - 413), F1: 18446744073709551615, F2: v1}
    v32 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
