package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 bool) bool {
    var v4 uint = 766
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 = v4
        _ = fi5
    }
    var v6 float64 = -2.25
    v4 -= (v4 << ((65535) & 63))
    p0 = (1e10 < (float64(float32(v6)) - v6))
    _ = p0
    _ = v4
    _ = v6
    return (((v6 - float64(((1e10 - 3.14159) / 1e10))) * v6) >= float64((float32(3.14159) / 3.14159)))
}

func main() {
    var v1 int64 = 0
    var v2 uint32 = 2147483647
    v3 := []uint32{}
    var v7 int32 = -365
    fmt.Println(v7, v7, v7)
    if false {
        for fi8 := 0; fi8 < 12; fi8++ {
            switch int(v2) {
            case -485:
                v9 := GenS1{F0: "", F1: 255}
                _ = v9
            case 220:
                _ = len(v3)
            case 42:
                v10 := GenS2{F0: 1000, F1: 2}
                _ = v10
            }
            if true {
                fmt.Println(v7, v2)
                v3 = append(v3, (v2 >> ((18446744073709551615) & 63)))
            }
            _ = fi8
        }
    } else {
        var v11 bool = false
        switch "" {
        case "":
            switch "\n\t" {
            case "a":
                var v12 int32 = (-(v7))
                _ = v12
            case "":
                var v13 uint8 = 42
                _ = v13
            default:
                if ((v1 / (((v1 >> ((18446744073709551615) & 63))) | 1)) == (v1 << ((0) & 63))) {
                    if (v11 || (!(v1 > 9223372036854775807))) {
                        v14 := GenS1{F0: "\n\t", F1: uint16(v2)}
                        for fi15 := 0; fi15 < 1; fi15++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v16 uint64 = 715
                            v16 -= (^(v16))
                            _ = fi15
                            _ = v16
                        }
                        switch -46 {
                        case 32767:
                            for fi17 := 0; fi17 < 6; fi17++ {
                                if len(v3) > 0 {
                                    v3[0] = (v2 >> ((uint(v2)) & 63))
                                }
                                _ = fi17
                            }
                        case 4294967295:
                            var v18 string = ("" + (("\n\t" + "foo bar") + "hello"))
                            _ = v18
                        default:
                            var v19 uint8 = 11
                            _ = v19
                        }
                        _ = v14
                    }
                    if ((!(-216 == -2147483648)) && true) {
                        for fi20 := 0; fi20 < 6; fi20++ {
                            v21 := GenS1{F0: ("hello" + ""), F1: 42}
                            _ = fi20
                            _ = v21
                        }
                        if true {
                            var v22 uint16 = 1
                            _ = v22
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        for fi23 := 0; fi23 < 10; fi23++ {
                            _ = len(v3)
                            _ = fi23
                        }
                    } else {
                        var v24 int64 = v1
                        _ = v24
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                }
            }
        case "\n\t":
            var v25 float32 = 3.14159
            _ = v25
        }
        if true {
            if (0.0 > 1.5) {
                v26 := GenS3{F0: -1}
                v7 += (v7 >> ((1) & 63))
                _ = v26
            }
        } else {
            v1 -= v1
        }
        _ = v11
    }
    var v27 uint64 = 0
    _ = len(v3)
    fmt.Println(v2, v2, v7)
    v1 = (-(v1))
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint = 65535
        _ = fi28
        _ = v29
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v7
}
