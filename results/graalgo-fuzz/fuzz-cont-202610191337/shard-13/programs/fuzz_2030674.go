package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint64; F1 string; F2 bool }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 int16; F2 uint16 }

func main() {
    var v1 float32 = (3.14159 * 0.0)
    var v2 bool = false
    var v3 int8 = -103
    v4 := []int{}
    if ((1e10 - 0.1) >= (1.5 + (1e10 / (0.1 - -2.25)))) {
        var v5 uint8 = 237
        var v6 int32 = 42
        _ = v5
        _ = v6
    }
    if (0 < 0) {
        v3 = v3
        if false {
            if (225 != 1) {
                v3 = (v3 / ((v3) | 1))
                fmt.Println(v2, v1, v3)
            } else {
                var v7 uint8 = 1
                _ = v7
            }
            v8 := GenS1{F0: 0, F1: -9223372036854775808}
            _ = v8
        } else {
            v1 = v1
            for fi9 := 0; fi9 < 16; fi9++ {
                for fi10 := 0; fi10 < 15; fi10++ {
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        v11 := GenS1{F0: 1, F1: 127}
                        _ = v11
                    default:
                        var v12 int32 = 2147483647
                        _ = v12
                    }
                    v1 = ((v1 * float32(-2.25)) + (1.5 + v1))
                    if v2 {
                        var v13 uint16 = 288
                        _ = v13
                    }
                    _ = fi10
                }
                var v14 int32 = 2147483647
                _ = fi9
                _ = v14
            }
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 = false
        v2 = v2
        _ = fi15
    }
    fmt.Println(v1, v2)
    fmt.Println(v1)
    if (!(!v2)) {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (("a" + "a") <= ("hello" + ("hello" + "hello"))) {
                if (!true) {
                    if v2 {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            v18 := GenS1{F0: 65535, F1: (fi17 - -79)}
                            fmt.Println(v2, v1)
                            _ = fi17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 12; fi19++ {
                            var v20 int16 = -354
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        var v21 uint16 = 65535
                        for fi22 := 0; fi22 < 17; fi22++ {
                            v23 := GenS2{F0: 451}
                            v2 = v2
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                    if ((true && false) || (4294967295 > 394)) {
                        if v2 {
                            v24 := GenS3{F0: -9223372036854775808, F1: (fi16 * 42)}
                            v2 = v2
                            _ = len(v4)
                            _ = v24
                        } else {
                            v3 -= 1
                            for fi25 := 0; fi25 < 9; fi25++ {
                                v26 := GenS2{F0: 1}
                                var v27 int64 = -9223372036854775808
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                        if false {
                            v28 := GenS4{F0: (((((0.1 - (0.0 * v1)) / (-2.25 / v1)) / v1) * ((0.0 - (v1 - -2.25)) - ((1e10 - (v1 + -2.25)) * v1))) < v1), F1: -2147483648}
                            _ = v28
                        } else {
                            for fi29 := 0; fi29 < 15; fi29++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                if len(v4) > 0 {
                                    v4[0] = fi16
                                }
                                _ = fi29
                            }
                            v1 = float32(0.1)
                            v1 += 0.0
                        }
                        for fi30 := 0; fi30 < 6; fi30++ {
                            if v2 {
                                switch "a" {
                                case "":
                                    v31 := GenS3{F0: -9223372036854775808, F1: fi16}
                                    _ = v31
                                case "hello":
                                    if false {
                                        v32 := GenS1{F0: 42, F1: (fi30 >> ((2) & 63))}
                                        v4 = append(v4, 32767)
                                        var v33 int = (fi16 | fi30)
                                        _ = v32
                                        _ = v33
                                    }
                                case "foo bar":
                                    _ = len(v4)
                                }
                                v34 := GenS2{F0: -262}
                                _ = v34
                            }
                            _ = fi30
                        }
                    } else {
                        switch (("\n\t" + "foo bar") + "") {
                        case "foo bar":
                            v2 = (-313 == 0)
                        case "hello":
                            switch (fi16 * (fi16 + (fi16 | fi16))) {
                            case -1:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 1:
                                for fi35 := 0; fi35 < 14; fi35++ {
                                    if (v1 > (v1 - 0.0)) {
                                        var v36 int8 = -1
                                        _ = v36
                                    }
                                    v2 = v2
                                    v37 := GenS4{F0: false, F1: 416}
                                    v2 = v2
                                    _ = fi35
                                    _ = v37
                                }
                            case 128:
                                v2 = (9223372036854775807 < -459)
                            default:
                                var v38 bool = ((!v2) && v2)
                                _ = v38
                            }
                        }
                        v39 := GenS5{F0: 691, F1: "a", F2: v2}
                        _ = v39
                    }
                }
                var v40 uint = 4294967295
                _ = v40
            } else {
                switch (fi16 | fi16) {
                case 1:
                    v2 = false
                case -539:
                    var v41 int8 = v3
                    _ = v41
                default:
                    v42 := GenS6{F0: -9223372036854775808, F1: 42}
                    _ = v42
                }
                for fi43 := 0; fi43 < 7; fi43++ {
                    switch "a" {
                    case "hello":
                        fmt.Println(fi16, v2, v3)
                    default:
                        v1 += v1
                    }
                    if true {
                        switch (fi16 - fi43) {
                        case 128:
                            v2 = (1 == 0)
                        default:
                            _ = len(v4)
                        }
                        if ((v3 + (v3 % (((-(v3))) | 1))) >= 124) {
                            fmt.Println(v3, fi43, v1)
                        } else {
                            var v44 uint64 = 0
                            _ = v44
                        }
                        _ = len(v4)
                    } else {
                        v2 = (18446744073709551615 > 0)
                        if ("hello" < "\n\t") {
                            v2 = (!true)
                            v1 = v1
                        } else {
                            v2 = true
                            var v45 int16 = -195
                            switch fi16 {
                            case -7:
                                v1 = v1
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v46 uint32 = 2147483648
                                _ = v46
                            case 32768:
                                v47 := GenS7{F0: 1, F1: (v3 - 100), F2: 42}
                                _ = v47
                            case 9223372036854775807:
                                switch "foo bar" {
                                case "":
                                    var v48 uint32 = 0
                                    _ = v48
                                case "foo bar":
                                    if len(v4) > 0 {
                                        v4[0] = (fi16 ^ (fi43 / ((fi16) | 1)))
                                    }
                                case "hello":
                                    fmt.Println(v45, v3)
                                default:
                                    v2 = (269 < 1)
                                }
                            default:
                                v49 := GenS2{F0: -1}
                                _ = v49
                            }
                            _ = v45
                        }
                        fmt.Println(v3, v2, v2)
                    }
                    var v50 int32 = 2147483647
                    _ = fi43
                    _ = v50
                }
            }
            _ = fi16
        }
    }
    switch ("\n\t" + "") {
    case "\n\t":
        for fi51 := 0; fi51 < 20; fi51++ {
            for fi52 := 0; fi52 < 5; fi52++ {
                var v53 uint8 = 255
                v2 = v2
                if true {
                    var v54 string = "hello"
                    fmt.Println(v53, v53, v54)
                    var v55 uint = 0
                    fmt.Println(v2, v53)
                    if v2 {
                        fmt.Println(v2, fi52)
                        switch fi52 {
                        case 541:
                            v55 -= v55
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi51 | 2147483647)
                        }
                    } else {
                        switch 834 {
                        case 1:
                            fmt.Println(v53)
                            var v56 int32 = 1
                            _ = v56
                        }
                        var v57 uint = v55
                        var v58 int = fi51
                        v59 := GenS8{F0: v2}
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    _ = v54
                    _ = v55
                } else {
                    _ = len(v4)
                }
                _ = fi52
                _ = v53
            }
            v3 = v3
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi51
        }
    case "":
        if (293 != -1) {
            v60 := GenS9{F0: -106, F1: -815, F2: 679}
            fmt.Println(v3, v3)
            _ = v60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
