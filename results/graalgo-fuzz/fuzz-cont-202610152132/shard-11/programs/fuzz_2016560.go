package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 bool }

func main() {
    v1 := []uint32{2147483648}
    v2 := []int{1, 9223372036854775807, -9223372036854775808}
    var v3 uint64 = 1
    if (!true) {
        var v4 int32 = 42
        switch (("" + "a") + ("hello" + "\n\t")) {
        case "a":
            var v5 int8 = 37
            _ = v5
        default:
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS1{F0: (^(v4)), F1: 1}
                v1 = append(v1, 598)
                _ = fi6
                _ = v7
            }
        }
        _ = v4
    }
    v8 := GenS1{F0: 2147483647, F1: 42}
    if true {
        var v9 uint8 = 42
        fmt.Println(v3, v9)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v9
    }
    var v10 int = -2147483648
    _ = len(v1)
    switch (-(v10)) {
    case -9223372036854775808:
        v10 += (v10 / ((v10) | 1))
    default:
        if ((2 > 0) || true) {
            if false {
                v3 = 781
                if (true && (13 == 127)) {
                    v10 -= (v10 | v10)
                }
                fmt.Println(v10, v3)
            }
            fmt.Println(v3, v10)
        }
    }
    if true {
        if false {
            var v11 uint8 = 0
            for fi12 := 0; fi12 < 14; fi12++ {
                _ = len(v1)
                v13 := GenS2{F0: fi12, F1: -128, F2: (!(561 < 966))}
                v14 := GenS2{F0: -9223372036854775808, F1: -128, F2: (768 <= 948)}
                _ = fi12
                _ = v13
                _ = v14
            }
            if (((-2.25 - 3.14159) / 0.1) <= (1e10 - 1e10)) {
                fmt.Println(v10, v11, v10)
                v3 += (v3 + v3)
                if (false && ((65535 < 2) && (((1e10 * 3.14159) == 3.14159) || (!false)))) {
                    switch "\n\t" {
                    case "\n\t":
                        if (9223372036854775807 >= v10) {
                            fmt.Println(v10)
                            var v15 bool = (((((1e10 - -2.25) * 3.14159) + (0.1 * -2.25)) >= (float64(0.0) + 0.0)) || (v10 == (^(v10))))
                            if (1 > -1) {
                                switch -424 {
                                case -2147483648:
                                    for fi16 := 0; fi16 < 1; fi16++ {
                                        if (v3 == 42) {
                                            fmt.Println(v3, v10)
                                            fmt.Println(fi16, v10, v10)
                                        }
                                        _ = len(v2)
                                        _ = fi16
                                    }
                                }
                                var v17 bool = (2147483648 == 4294967295)
                                fmt.Println(v11, v10, v17)
                                _ = v17
                            } else {
                                v11 -= (v11 / ((v11) | 1))
                                for fi18 := 0; fi18 < 3; fi18++ {
                                    v15 = v15
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    switch int(v3) {
                                    case 9223372036854775807:
                                        _ = len(v1)
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = (v10 - fi18)
                                        }
                                    }
                                    _ = fi18
                                }
                                v3 = (v3 - (v3 % ((v3) | 1)))
                            }
                            _ = v15
                        } else {
                            v19 := GenS1{F0: 0, F1: 42}
                            var v20 float32 = (1.5 + 1e10)
                            _ = v19
                            _ = v20
                        }
                    case "":
                        v11 = v11
                    case "hello":
                        var v21 int32 = -1
                        _ = v21
                    }
                    var v22 float64 = float64(3.14159)
                    _ = v22
                }
            }
            _ = v11
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            v3 -= 1
            v10 = v10
            _ = fi23
        }
        if (3.14159 > (-2.25 * (((1e10 - (1.5 * 1.5)) * 0.0) / (1e10 - 0.0)))) {
            var v24 int16 = 0
            _ = v24
        }
    } else {
        var v25 int16 = -32768
        _ = v25
    }
    v1 = append(v1, 553)
    var v26 uint8 = 189
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v8
}
