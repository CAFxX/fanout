package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint32 }
func genFn0(p0 int16, p1 int64) int16 {
    p0 = 1000
    v7 := GenS1{F0: -302, F1: 610, F2: 2147483647}
    _ = p0
    _ = p1
    _ = v7
    return p0
}

func main() {
    var v1 uint8 = 1
    var v2 int32 = 42
    var v3 uint8 = 255
    var v4 uint16 = 65535
    var v5 uint16 = v4
    var v6 uint16 = 226
    for fi8 := 0; fi8 < 20; fi8++ {
        v4 += v6
        _ = fi8
    }
    var v9 float32 = 1e10
    if (v9 <= v9) {
        var v10 uint8 = 170
        _ = v10
    }
    if (!(!(v2 >= (v2 - (v2 + (-(v2))))))) {
        v1 -= 128
    }
    v11 := GenS1{F0: 0, F1: 18446744073709551615, F2: 0}
    fmt.Println(v5, v6)
    var v12 int32 = (v2 % ((v2) | 1))
    if (0.1 < -2.25) {
        switch -9223372036854775808 {
        case 4294967295:
            if true {
                switch 128 {
                case 2147483647:
                    var v13 uint8 = (v1 / ((v3) | 1))
                    _ = v13
                case 42:
                    var v14 uint32 = 2
                    _ = v14
                }
            } else {
                v2 = (-(v12))
                switch ("" + ("foo bar" + ("foo bar" + "foo bar"))) {
                case "hello":
                    v15 := []int{9223372036854775807, -349, 141, 42}
                    _ = v15
                }
                for fi16 := 0; fi16 < 10; fi16++ {
                    v3 = 127
                    _ = fi16
                }
            }
        case 9223372036854775807:
            if false {
                switch 32767 {
                case 9223372036854775807:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if (1000 <= -360) {
                            v6 -= 2
                            v3 += uint8(v6)
                            switch -2147483648 {
                            case 42:
                                var v18 uint32 = 2
                                _ = v18
                            case -9223372036854775808:
                                if false {
                                    v19 := GenS1{F0: 0, F1: 71, F2: 382}
                                    for fi20 := 0; fi20 < 8; fi20++ {
                                        v5 += 2
                                        switch "" {
                                        case "\n\t":
                                            var v21 int16 = genFn0(910, 0)
                                            _ = v21
                                        case "a":
                                            switch -9223372036854775808 {
                                            case 850:
                                                v22 := GenS1{F0: -948, F1: 18446744073709551615, F2: 2147483648}
                                                var v23 uint = 65535
                                                v4 -= (v4 * v6)
                                                _ = v22
                                                _ = v23
                                            case -502:
                                                v5 = (v4 & v6)
                                            default:
                                                v2 -= 514
                                                v24 := []int{}
                                                _ = v24
                                            }
                                        case "foo bar":
                                            v3 -= v3
                                        }
                                        _ = fi20
                                    }
                                    v25 := GenS1{F0: 1, F1: 4294967295, F2: 4294967295}
                                    _ = v19
                                    _ = v25
                                } else {
                                    fmt.Println(v4, v1, v12)
                                    var v26 int64 = 1
                                    v1 = v1
                                    _ = v26
                                }
                            case 32767:
                                if (-1 > -67) {
                                    var v27 uint32 = 458
                                    v6 += 1
                                    _ = v27
                                }
                            default:
                                v6 += v4
                            }
                        }
                        _ = fi17
                    }
                }
            }
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            if true {
                fmt.Println(v9, v2)
                v9 = float32(-2.25)
                switch ("\n\t" + "") {
                case "a":
                    var v29 float64 = (((0.1 / 1.0) * 1e10) * float64((3.14159 - (3.14159 - float32(0.0)))))
                    _ = v29
                case "":
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v1 = (v3 >> ((18446744073709551615) & 63))
                        _ = fi30
                    }
                default:
                    var v31 float32 = ((1.5 / (v9 - v9)) - ((v9 * (v9 - 1e10)) / 1e10))
                    _ = v31
                }
            }
            _ = fi28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
