package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8; F1 int16; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 int32; F2 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{1, 636, 2147483647}
    var v3 float32 = (((3.14159 / 1.0) + (float32((float64(1e10) + ((((1.5 - (0.0 + 1.5)) / 0.1) * ((3.14159 * (float64((((((-2.25 - 0.0) * ((((-2.25 + 1.5) * 3.14159) * (1e10 * -2.25)) * ((-2.25 + ((0.0 / 3.14159) - ((0.0 * 0.0) - ((1e10 - 0.1) / 1.0)))) * (0.0 * 0.0)))) * float32(((3.14159 * float64(1.5)) / -2.25))) * ((-2.25 / (1.5 - (0.1 / 1.5))) + (-2.25 - (3.14159 + 0.0)))) + 0.1)) - 1.5)) * 3.14159)) / (0.1 / -2.25)))) * (3.14159 * (3.14159 * 1e10)))) / (((0.0 - (1.5 * (0.0 - (3.14159 + -2.25)))) + 3.14159) - 0.0))
    var v4 int64 = -549
    v1 = (v1 % ((1) | 1))
    v5 := GenS1{F0: 127, F1: 740}
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint = 255
        fmt.Println(v3, v3)
        fmt.Println(v4, v4, v1)
        _ = fi6
        _ = v7
    }
    switch 9223372036854775807 {
    case -117:
        v3 -= v3
    case -315:
        v8 := GenS2{F0: 1, F1: 1, F2: ("" + "foo bar")}
        _ = v8
    case 1:
        var v9 uint = 1
        _ = v9
    default:
        if true {
            if false {
                var v10 int64 = -700
                _ = v10
            } else {
                v3 = (((v3 * 3.14159) * v3) - (float32(1e10) * (v3 + ((v3 * v3) / ((0.0 - float32(float64(float32(0.1)))) / v3)))))
                if false {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        _ = len(v2)
                        var v12 uint16 = 255
                        _ = fi11
                        _ = v12
                    }
                    var v13 float64 = 3.14159
                    _ = v13
                }
                v14 := GenS2{F0: 113, F1: 0, F2: ""}
                _ = v14
            }
            var v15 uint8 = 0
            if ((630 == 65535) || ((-9223372036854775808 >= 840) && true)) {
                fmt.Println(v1, v4)
                v16 := GenS1{F0: 1, F1: 555}
                _ = v16
            } else {
                if true {
                    for fi17 := 0; fi17 < 9; fi17++ {
                        v3 += 0.1
                        v1 = v1
                        _ = fi17
                    }
                    v18 := GenS3{F0: 32767}
                    _ = v18
                }
                var v19 uint64 = (^(v1))
                _ = v19
            }
            _ = v15
        } else {
            if false {
                if ((v4 == (v4 ^ (-(v4)))) && true) {
                    switch "" {
                    case "a":
                        for fi20 := 0; fi20 < 3; fi20++ {
                            _ = len(v2)
                            var v21 uint8 = 0
                            _ = fi20
                            _ = v21
                        }
                    case "\n\t":
                        var v22 uint8 = 0
                        _ = v22
                    }
                } else {
                    var v23 int = 0
                    _ = v23
                }
                for fi24 := 0; fi24 < 5; fi24++ {
                    fmt.Println(v1)
                    var v25 float32 = float32(3.14159)
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    _ = fi24
                    _ = v25
                }
                switch 32768 {
                case 1:
                    switch -492 {
                    case 127:
                        switch ("a" + "") {
                        case "a":
                            for fi26 := 0; fi26 < 7; fi26++ {
                                var v27 uint64 = (v1 / ((910) | 1))
                                _ = fi26
                                _ = v27
                            }
                        default:
                            v28 := GenS3{F0: -478}
                            _ = v28
                        }
                    case 497:
                        if ((0 < 0) || false) {
                            v29 := GenS4{F0: (v4 & 878)}
                            switch 4294967295 {
                            case -231:
                                switch 127 {
                                case 93:
                                    var v30 uint = 190
                                    _ = v30
                                case 911:
                                    v31 := GenS5{F0: 0, F1: 42, F2: "\n\t"}
                                    _ = v31
                                default:
                                    if true {
                                        fmt.Println(v3, v3, v1)
                                        switch "\n\t" {
                                        case "foo bar":
                                            v4 -= (v4 / ((v4) | 1))
                                        case "a":
                                            v3 += 0.1
                                        }
                                        if (!(0.1 != float64((0.1 / v3)))) {
                                            var v32 int32 = -2147483648
                                            v33 := GenS1{F0: 42, F1: 0}
                                            _ = v32
                                            _ = v33
                                        }
                                    } else {
                                        v34 := GenS5{F0: -2147483648, F1: 0, F2: ""}
                                        var v35 int32 = 42
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v34
                                        _ = v35
                                    }
                                }
                            default:
                                if (!true) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v4, v1, v3)
                                } else {
                                    for fi36 := 0; fi36 < 17; fi36++ {
                                        v3 = 0.1
                                        for fi37 := 0; fi37 < 15; fi37++ {
                                            _ = len(v2)
                                            _ = fi37
                                        }
                                        var v38 uint = 393
                                        _ = fi36
                                        _ = v38
                                    }
                                }
                            }
                            v1 = v1
                            _ = v29
                        } else {
                            var v39 float32 = ((-2.25 / -2.25) / (v3 * -2.25))
                            _ = v39
                        }
                    }
                case -475:
                    v3 = v3
                case 4294967295:
                    v40 := GenS3{F0: -1}
                    _ = v40
                }
            }
        }
    }
    if ((!(v3 >= 3.14159)) && false) {
        v3 -= v3
        var v41 int64 = (v4 / ((v4) | 1))
        _ = v41
    } else {
        v42 := GenS3{F0: 0}
        _ = v42
    }
    if true {
        v4 = v4
        fmt.Println(v1)
        var v43 uint64 = v1
        _ = v43
    }
    _ = len(v2)
    if (-1 > 0) {
        fmt.Println(v3, v3, v4)
        var v44 int16 = 486
        _ = v44
    }
    _ = len(v2)
    if (false || (-686 == -1)) {
        v45 := GenS4{F0: (v4 >> ((754) & 63))}
        _ = v45
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v46 uint64 = v1
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
