package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint8 = 237
    var v2 uint8 = 14
    v3 := []int{-1, 501, 951}
    var v4 bool = (32767 > 32767)
    v3 = append(v3, -9223372036854775808)
    _ = len(v3)
    var v5 string = ("" + "\n\t")
    if (v4 && true) {
        v1 -= 127
        switch -605 {
        case 4294967295:
            v4 = v4
        }
    } else {
        switch 51 {
        case -398:
            for fi6 := 0; fi6 < 6; fi6++ {
                if (!(((255 >= 1) || (1 > 1000)) || (((float64(1e10) / -2.25) + (((0.0 * float64(float32((((1.5 / 1e10) - 0.1) + (1.5 - 3.14159))))) * -2.25) - 1.5)) <= float64((3.14159 / ((1.5 - (float32(0.0) - (float32(float64(0.0)) * (1e10 + float32((-2.25 * (0.0 - 0.1))))))) / 0.1)))))) {
                    var v7 uint16 = 948
                    fmt.Println(fi6, v7, v4)
                    _ = v7
                }
                _ = fi6
            }
        default:
            v2 -= (v1 / (((v2 ^ v1)) | 1))
        }
    }
    if v4 {
        for fi8 := 0; fi8 < 1; fi8++ {
            switch fi8 {
            case -671:
                v5 = ("" + v5)
            }
            _ = len(v3)
            v9 := GenS1{F0: ((("hello" + v5) + (v5 + v5)) + "a"), F1: 1}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: 0, F1: -9223372036854775808}
    if v4 {
        if true {
            switch -576 {
            case -261:
                if (((((v5 + v5) + "foo bar") + v5) + "foo bar") >= v5) {
                    if false {
                        if false {
                            switch "a" {
                            case "a":
                                v11 := GenS1{F0: v5, F1: 0}
                                _ = v11
                            case "hello":
                                fmt.Println(v2)
                            case "foo bar":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            switch "\n\t" {
                            case "a":
                                var v12 uint = 157
                                _ = v12
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                switch -9223372036854775808 {
                                case 2147483647:
                                    v2 += (v2 - (v1 * (v2 & 2)))
                                }
                            }
                            var v13 int8 = -1
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 2; fi14++ {
                            v15 := GenS3{F0: int64(fi14)}
                            fmt.Println(v4)
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 12; fi16++ {
                            v17 := GenS2{F0: -968, F1: 32768}
                            v5 = v5
                            var v18 int8 = 1
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        fmt.Println(v2, fi19, v1)
                        _ = fi19
                    }
                }
            default:
                _ = len(v3)
            }
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 bool = (1 > 9223372036854775807)
                if len(v3) > 0 {
                    v3[0] = (fi20 | (fi20 % ((fi20) | 1)))
                }
                _ = fi20
                _ = v21
            }
        }
        v1 = v2
        if len(v3) > 0 {
            v3[0] = 32768
        }
    } else {
        if len(v3) > 0 {
            v3[0] = -937
        }
    }
    v22 := GenS2{F0: 1, F1: 4294967295}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 = (v5 + v5)
    v4 = ((^(v1)) == (^(v1)))
    var v23 uint = 65535
    v4 = (-1 > 22)
    _ = len(v3)
    for fi24 := 0; fi24 < 18; fi24++ {
        v4 = (-325 != -764)
        var v25 int8 = -128
        v2 = v2
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v1 -= 42
        switch "" {
        case "hello":
            v5 = "\n\t"
        default:
            var v27 uint16 = 945
            _ = v27
        }
        _ = fi26
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 uint = (v23 >> (((v23 ^ 255)) & 63))
        var v30 bool = (!v4)
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS1{F0: (v5 + (v5 + "foo bar")), F1: (v23 | 65535)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
