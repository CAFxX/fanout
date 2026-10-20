package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int32 = -1
        var v6 float32 = -2.25
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 1
    v8 := GenS1{F0: 9223372036854775807}
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := []int{-9223372036854775808, 1, 4294967295, 4294967295}
    var v3 int8 = 86
    fmt.Println(v1)
    v9 := GenS1{F0: -369}
    if true {
        if len(v2) > 0 {
            v2[0] = 42
        }
        var v10 float64 = 3.14159
        v11 := GenS1{F0: 719}
        _ = v10
        _ = v11
    }
    if (false && true) {
        v1 = (v1 + v1)
    }
    v12 := GenS2{F0: 4294967295, F1: 772}
    v3 += v3
    var v13 uint16 = 0
    _ = len(v2)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint8 = 42
        if (true && (!true)) {
            if (285 < -2147483648) {
                switch fi14 {
                case 238:
                    var v16 uint = 482
                    _ = v16
                case 2147483647:
                    if (788 < -32768) {
                        v1 = v1
                    } else {
                        var v17 float32 = 3.14159
                        var v18 uint8 = (v15 % ((v15) | 1))
                        _ = v17
                        _ = v18
                    }
                case -1:
                    switch fi14 {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        switch ((v1 + "") + v1) {
                        case "foo bar":
                            v3 += v3
                        default:
                            var v19 int32 = int32(fi14)
                            _ = v19
                        }
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v20 float32 = float32(1.5)
                if false {
                    var v21 uint8 = 250
                    if (!true) {
                        v3 = v3
                    } else {
                        if (!true) {
                            var v22 uint16 = v13
                            var v23 int32 = 0
                            _ = v22
                            _ = v23
                        } else {
                            if (false || ((1 < 42) && true)) {
                                v15 -= v21
                                var v24 uint32 = uint32(v21)
                                _ = v24
                            } else {
                                _ = len(v2)
                            }
                        }
                    }
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 int = 9223372036854775807
                        _ = fi25
                        _ = v26
                    }
                    _ = v21
                } else {
                    var v27 uint64 = 42
                    _ = v27
                }
                _ = v20
            } else {
                v13 = (v13 | (v13 | (v13 ^ 0)))
                var v28 int8 = v3
                _ = v28
            }
        }
        _ = fi14
        _ = v15
    }
    v13 += 65535
    if false {
        var v29 int16 = 374
        _ = v29
    } else {
        _ = len(v2)
        var v30 float32 = (0.0 * 1.5)
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        fmt.Println(fi31)
        for fi32 := 0; fi32 < 5; fi32++ {
            v1 = v1
            v33 := GenS1{F0: -8}
            v34 := GenS2{F0: 18446744073709551615, F1: v13}
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
