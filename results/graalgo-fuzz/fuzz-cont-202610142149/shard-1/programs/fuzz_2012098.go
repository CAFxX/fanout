package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 int16; F1 uint32 }
type GenS8 struct { F0 uint; F1 int32 }

func main() {
    var v1 int16 = 1
    var v2 float32 = 1e10
    var v3 int64 = 700
    var v4 uint8 = 0
    v5 := make(map[uint32]bool)
    var v6 int8 = int8(v4)
    v7 := GenS1{F0: 412, F1: ((true || true) || true), F2: v4}
    for fi8 := 0; fi8 < 2; fi8++ {
        if false {
            if false {
                var v9 string = "hello"
                v10 := v5[133]
                switch v9 {
                case "hello":
                    v11 := GenS2{F0: v3, F1: (fi8 ^ fi8), F2: v1}
                    _ = v11
                }
                _ = v10
                _ = v9
            } else {
                v12 := []int{fi8, 1, -153, 32768}
                if (!false) {
                    fmt.Println(v3, v3, v2)
                }
                _ = v12
            }
            if ((v2 - v2) < v2) {
                var v13 int8 = -72
                var v14 int8 = 0
                _ = v13
                _ = v14
            }
        }
        _ = fi8
    }
    if false {
        v3 = (v3 * -1)
    } else {
        var v15 int64 = (v3 - v3)
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        switch fi16 {
        case 0:
            v4 = (v4 / (((v4 >> ((uint(v4)) & 63))) | 1))
        case -2147483648:
            var v17 float64 = 0.1
            _ = v17
        case -436:
            v18 := GenS1{F0: 0, F1: false, F2: (^(v4))}
            _ = v18
        }
        _ = fi16
    }
    var v19 int64 = v3
    v4 -= v4
    if (-1 == 789) {
        v20 := GenS2{F0: (v19 % ((91) | 1)), F1: 1, F2: v1}
        _ = v20
    }
    v2 += ((1.5 / v2) - v2)
    if true {
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int32 = -15
            if false {
                if (false || false) {
                    v4 -= 127
                    for fi23 := 0; fi23 < 11; fi23++ {
                        v24 := GenS3{F0: 42}
                        v2 = 0.0
                        if ((1e10 == (float64((3.14159 / (v2 - (v2 / v2)))) / (0.1 / 1.0))) || (482 != 18446744073709551615)) {
                            delete(v5, 1)
                            v25 := GenS4{F0: (v3 / (((-(v19))) | 1))}
                            v26 := GenS4{F0: (v19 * (v3 << ((984) & 63)))}
                            _ = v25
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                var v27 int32 = (v22 / (((v22 ^ int32(v1))) | 1))
                _ = v27
            } else {
                fmt.Println(v4, v19, v19)
            }
            _ = fi21
            _ = v22
        }
    } else {
        if true {
            fmt.Println(v1, v3)
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 int = fi28
            var v30 float64 = (0.1 + float64(v2))
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if false {
        v31 := []int{}
        v6 += 42
        _ = v31
    } else {
        fmt.Println(v6)
        v6 += (v6 - v6)
        if false {
            switch 2147483647 {
            case 127:
                for fi32 := 0; fi32 < 14; fi32++ {
                    var v33 int16 = (v1 / ((int16(v19)) | 1))
                    v34 := []int{433}
                    v19 += (v19 * int64(v33))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            case 128:
                for fi35 := 0; fi35 < 8; fi35++ {
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v37 := GenS4{F0: v19}
                        v38 := GenS4{F0: (v19 - 0)}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                    v19 -= v3
                    for fi39 := 0; fi39 < 7; fi39++ {
                        var v40 string = ("" + ("a" + "hello"))
                        v41 := GenS3{F0: 913}
                        v42 := v5[42]
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    _ = fi35
                }
            case 750:
                if (!(false || (int32(v1) == 2147483647))) {
                    if (((2147483647 > -649) || (!(v4 <= (v4 & 2)))) && true) {
                        fmt.Println(v3, v4, v4)
                        for fi43 := 0; fi43 < 8; fi43++ {
                            v44 := GenS5{F0: 42, F1: (v1 << ((1) & 63))}
                            v45 := []int{}
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                        v46 := GenS1{F0: 4294967295, F1: true, F2: v4}
                        _ = v46
                    }
                    if false {
                        delete(v5, 65535)
                    } else {
                        var v47 int = 9223372036854775807
                        if (("hello" <= "\n\t") && false) {
                            var v48 int8 = (v6 ^ 91)
                            _ = v48
                        } else {
                            var v49 int8 = 0
                            _ = v49
                        }
                        if ((0 != 18446744073709551615) || false) {
                            var v50 bool = (2 <= 2)
                            var v51 float32 = v2
                            _ = v50
                            _ = v51
                        } else {
                            if (100 > (v6 % (((v6 | v6)) | 1))) {
                                switch -9223372036854775808 {
                                case 550:
                                    var v52 bool = false
                                    _ = v52
                                case -543:
                                    switch ("hello" + "a") {
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v3, v19)
                                            v1 = v1
                                            switch "\n\t" {
                                            case "":
                                                v4 = (^(v4))
                                            default:
                                                if true {
                                                    v5[0] = true
                                                } else {
                                                    v19 += 9223372036854775807
                                                }
                                            }
                                        }
                                    case "":
                                        fmt.Println(v1)
                                    }
                                }
                            } else {
                                var v53 int8 = 0
                                _ = v53
                            }
                            switch v47 {
                            case 574:
                                var v54 uint = 42
                                _ = v54
                            case 2147483647:
                                v55 := GenS2{F0: -9223372036854775808, F1: v47, F2: (v1 | 1)}
                                _ = v55
                            case -915:
                                v47 = v47
                            default:
                                if (0 == 2) {
                                    switch ("foo bar" + "") {
                                    case "":
                                        if true {
                                            v56 := GenS6{F0: (-(v19)), F1: 438}
                                            if (false && (!true)) {
                                                v4 = (^(v4))
                                                if (true && (1.5 >= v2)) {
                                                    var v57 int32 = 1
                                                    v58 := GenS5{F0: uint32(v4), F1: (^(v1))}
                                                    _ = v57
                                                    _ = v58
                                                }
                                            }
                                            _ = v56
                                        } else {
                                            var v59 int32 = 268
                                            if false {
                                                if (((v6 / (((-(v6))) | 1)) > 0) && false) {
                                                    v2 -= (float32(0.1) / float32((((0.0 * 0.0) * (1.5 - (0.0 - -2.25))) + 0.1)))
                                                    for fi60 := 0; fi60 < 4; fi60++ {
                                                        var v61 string = "a"
                                                        _ = fi60
                                                        _ = v61
                                                    }
                                                    var v62 uint = 254
                                                    _ = v62
                                                } else {
                                                    v59 = v59
                                                    v63 := GenS7{F0: (v1 + v1), F1: 1}
                                                    v3 -= v19
                                                    _ = v63
                                                }
                                                if true {
                                                    var v64 int = -9223372036854775808
                                                    v65 := GenS3{F0: 42}
                                                    _ = v64
                                                    _ = v65
                                                }
                                                for fi66 := 0; fi66 < 9; fi66++ {
                                                    var v67 bool = false
                                                    var v68 uint64 = 4294967295
                                                    _ = fi66
                                                    _ = v67
                                                    _ = v68
                                                }
                                            }
                                            v5[2] = (547 < 0)
                                            _ = v59
                                        }
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v47, v19)
                                            fmt.Println(v3, v2, v1)
                                            if (v1 > v1) {
                                                v69 := GenS4{F0: v19}
                                                _ = v69
                                            } else {
                                                var v70 float32 = float32(1e10)
                                                if true {
                                                    v1 = 42
                                                    var v71 int64 = 1
                                                    _ = v71
                                                }
                                                v72 := GenS7{F0: (^(v1)), F1: 2147483647}
                                                _ = v70
                                                _ = v72
                                            }
                                        }
                                    case "foo bar":
                                        switch v47 {
                                        case -2147483648:
                                            if true {
                                                v73 := []int{(v47 >> ((42) & 63))}
                                                v3 += (v19 & 42)
                                                _ = v73
                                            }
                                        case 2147483647:
                                            v47 = (v47 >> ((470) & 63))
                                        }
                                    }
                                }
                            }
                        }
                        _ = v47
                    }
                    var v74 float32 = v2
                    _ = v74
                } else {
                    var v75 uint32 = 675
                    var v76 uint16 = 2
                    _ = v75
                    _ = v76
                }
            }
            for fi77 := 0; fi77 < 1; fi77++ {
                var v78 float32 = v2
                if (!true) {
                    if (1e10 == ((0.0 * 1.5) - 1e10)) {
                        var v79 bool = true
                        _ = v79
                    }
                    v80 := GenS1{F0: 1, F1: (false || false), F2: (v4 & (^(v4)))}
                    _ = v80
                } else {
                    if ((!false) || ((float64(0.0) / 1.0) > 0.1)) {
                        delete(v5, 65535)
                    }
                    v81 := GenS4{F0: -848}
                    _ = v81
                }
                v6 = -58
                _ = fi77
                _ = v78
            }
        }
    }
    v82 := []int{1, 127}
    if (v4 > v4) {
        switch "" {
        case "hello":
            v4 = (v4 - 181)
        default:
            v83 := GenS8{F0: 0, F1: 312}
            _ = v83
        }
        fmt.Println(v1, v4, v6)
        var v84 uint8 = 128
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v82
}
