package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint16; F1 int64 }

func main() {
    var v1 float32 = float32(1.5)
    var v2 float64 = 1.5
    var v3 float64 = float64(v1)
    v4 := []int{-183, -1, -471, -2147483648}
    var v5 int32 = 42
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            if (493 > 18446744073709551615) {
                var v7 int32 = (v5 + -2147483648)
                _ = v7
            }
            if (false || false) {
                v4 = append(v4, fi6)
                v2 = 1e10
            } else {
                switch (fi6 / ((-2147483648) | 1)) {
                case -159:
                    v4 = append(v4, fi6)
                case 2147483647:
                    v1 += v1
                default:
                    v8 := GenS1{F0: 2147483647, F1: 742}
                    _ = v8
                }
                v2 += (((float64((((((v1 / v1) + 1e10) + (v1 - (v1 * ((v1 * (((v1 - ((v1 - (((0.0 + (v1 * (v1 * (v1 / v1)))) / float32(v2)) * v1)) / 1e10)) * (v1 * v1)) - 1.5)) + v1)))) * ((-2.25 - float32((-2.25 - float64(0.0)))) * 1e10)) / float32(v2))) - v3) - v2) + (((float64(((float32(1.5) / (((v1 * 0.1) / (v1 + 1.5)) - v1)) * 1e10)) * 1e10) + v3) + v3))
            }
        }
        v9 := GenS2{F0: 51}
        _ = fi6
        _ = v9
    }
    if (!false) {
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                v3 += v2
                _ = fi10
            }
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(v2, v3, v2)
                switch "" {
                case "\n\t":
                    v12 := GenS3{F0: 225, F1: 128}
                    _ = v12
                case "a":
                    v13 := GenS1{F0: 0, F1: 1000}
                    _ = v13
                case "foo bar":
                    v3 = (0.1 * v2)
                default:
                    if false {
                        for fi14 := 0; fi14 < 11; fi14++ {
                            v2 += -2.25
                            _ = fi14
                        }
                    } else {
                        switch 0 {
                        case 32768:
                            v2 = ((((v2 - v3) - -2.25) + v2) - v3)
                        case 128:
                            v2 = 1e10
                        default:
                            var v15 uint8 = 175
                            _ = v15
                        }
                        switch ("" + ("" + "\n\t")) {
                        case "foo bar":
                            switch fi11 {
                            case 0:
                                fmt.Println(v1)
                            case 363:
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    if false {
                                        if false {
                                            var v17 float64 = 0.1
                                            v18 := GenS4{F0: 255}
                                            v5 = int32(fi16)
                                            _ = v17
                                            _ = v18
                                        }
                                        v3 = v2
                                        var v19 int64 = 165
                                        _ = v19
                                    }
                                    _ = len(v4)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v1 = (0.1 / ((float32(v3) * 3.14159) + 0.1))
                                    _ = fi16
                                }
                            }
                        case "a":
                            var v20 uint32 = 65535
                            _ = v20
                        case "":
                            switch (("a" + "a") + "foo bar") {
                            case "foo bar":
                                var v21 string = "\n\t"
                                _ = v21
                            default:
                                v2 = (1e10 * 1.5)
                            }
                        default:
                            var v22 int32 = 111
                            _ = v22
                        }
                    }
                }
                _ = fi11
            }
            v4 = append(v4, -396)
            fmt.Println(v5, v2)
        }
        v1 = v1
        switch 32768 {
        case 913:
            _ = len(v4)
        }
    } else {
        v1 += ((0.0 - v1) - (v1 / v1))
        if (93 >= 23) {
            if true {
                v23 := GenS5{F0: 127, F1: 32768}
                var v24 int64 = 0
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 2, F1: -32768}
            _ = v25
        }
    }
    v26 := GenS3{F0: 127, F1: -50}
    var v27 bool = (50 > 65535)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v28 := GenS6{F0: 880, F1: -9223372036854775808}
    var v29 int64 = -9223372036854775808
    v30 := GenS3{F0: 0, F1: 128}
    fmt.Println(v1, v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
