package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    v1 := []int{128, 9223372036854775807, 128, -9223372036854775808}
    var v2 int32 = -208
    v3 := []int{32767, -982}
    switch "\n\t" {
    case "foo bar":
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v2 += (^(v2))
            if len(v3) > 0 {
                v3[0] = 370
            }
            var v4 uint32 = 191
            _ = v4
        }
    }
    if ((!(true && true)) || ((0.0 / -2.25) <= 0.1)) {
        v5 := GenS1{F0: 32767, F1: 308}
        _ = v5
    } else {
        _ = len(v3)
    }
    v2 += v2
    var v6 bool = true
    var v7 uint16 = 336
    var v8 int32 = (v2 + -2147483648)
    var v9 float64 = 0.0
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint8 = 128
        var v12 float32 = float32((-2.25 + (1.5 - -2.25)))
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = 153
    switch ("hello" + "foo bar") {
    case "foo bar":
        v8 = 0
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        switch "\n\t" {
        case "foo bar":
            var v15 uint16 = v7
            _ = v15
        }
        _ = fi14
    }
    switch "hello" {
    case "":
        v3 = append(v3, v13)
    case "hello":
        v13 += 32768
    case "\n\t":
        if (0 > uint8(v7)) {
            _ = len(v3)
            if false {
                fmt.Println(v2, v8)
                if v6 {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        var v17 int8 = -18
                        if false {
                            if (uint64(v7) < 1) {
                                v2 += v2
                            } else {
                                v8 += (v2 - v2)
                            }
                            v18 := GenS1{F0: int16(v8), F1: 18446744073709551615}
                            _ = v18
                        } else {
                            if v6 {
                                v1 = append(v1, (fi16 * v13))
                            } else {
                                v13 -= 32768
                                v6 = (v13 < fi16)
                                v8 -= v2
                            }
                        }
                        if ((fi16 | 443) == (fi16 - v13)) {
                            if (!(v6 && false)) {
                                for fi19 := 0; fi19 < 8; fi19++ {
                                    switch v13 {
                                    case 32767:
                                        _ = len(v1)
                                    }
                                    v8 = (^(v8))
                                    fmt.Println(v8, v8)
                                    _ = fi19
                                }
                            } else {
                                if ((65535 >= 751) || v6) {
                                    var v20 int = v13
                                    _ = v20
                                }
                                v21 := GenS1{F0: -924, F1: 255}
                                _ = v21
                            }
                            var v22 float64 = (v9 / -2.25)
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 7; fi23++ {
                                fmt.Println(v6, v6)
                                var v24 int32 = 1
                                if (!v6) {
                                    if len(v3) > 0 {
                                        v3[0] = (v13 + (-(fi16)))
                                    }
                                } else {
                                    v25 := GenS1{F0: -549, F1: 18446744073709551615}
                                    v26 := GenS1{F0: 115, F1: 500}
                                    var v27 bool = true
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                                _ = fi23
                                _ = v24
                            }
                        }
                        _ = fi16
                        _ = v17
                    }
                    v13 = (-(v13))
                    fmt.Println(v9, v7)
                }
                v28 := GenS1{F0: -32768, F1: 1}
                _ = v28
            } else {
                if ("foo bar" < ("hello" + ("\n\t" + "a"))) {
                    fmt.Println(v13)
                } else {
                    var v29 uint64 = 4294967295
                    _ = v29
                }
            }
        } else {
            for fi30 := 0; fi30 < 2; fi30++ {
                v3 = append(v3, 221)
                for fi31 := 0; fi31 < 20; fi31++ {
                    if ((!false) || v6) {
                        v9 = float64(0.1)
                        var v32 int64 = -1
                        _ = v32
                    }
                    v33 := GenS1{F0: 498, F1: 18446744073709551615}
                    _ = fi31
                    _ = v33
                }
                var v34 float64 = -2.25
                _ = fi30
                _ = v34
            }
        }
    default:
        fmt.Println(v13)
    }
    v6 = (-2.25 != 0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
