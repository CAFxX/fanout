package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 int8; F2 int16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 int; F1 int32; F2 int16 }
type GenS7 struct { F0 uint16; F1 uint; F2 uint }
type GenS8 struct { F0 int64; F1 uint; F2 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS11 struct { F0 uint8; F1 bool; F2 uint64 }
func genFn0() uint {
    v7 := GenS1{F0: -533, F1: 814, F2: (!true)}
    var v8 int = 127
    v9 := []int{}
    v10 := GenS1{F0: (v8 % ((v8) | 1)), F1: 572, F2: (false || (1 != 502))}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 869
}
func genFn1() uint32 {
    v11 := []int{-9223372036854775808, -58, 76, 2147483647}
    var v12 int32 = -534
    v13 := GenS1{F0: -2147483648, F1: 1, F2: (false || false)}
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int8 = 116
    var v2 int16 = -32768
    var v3 string = ""
    var v4 int32 = 42
    var v5 int64 = 42
    var v6 int = 204
    v4 = v4
    v5 += -1
    if ((v6 == (v6 << ((898) & 63))) && (!(625 >= v4))) {
        var v14 int8 = (v1 | 1)
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint32 = 65535
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v17 uint64 = 18446744073709551615
    var v18 uint32 = genFn1()
    if false {
        v3 = ("\n\t" + "foo bar")
        var v19 string = (v3 + (v3 + v3))
        for fi20 := 0; fi20 < 7; fi20++ {
            if (65535 == 733) {
                fmt.Println(v17, v18)
            } else {
                if (((v19 + v3) <= "\n\t") && true) {
                    v21 := []int{v6, (v6 + v6), (fi20 + int(v5)), -9223372036854775808}
                    _ = v21
                }
            }
            switch int(v5) {
            case 0:
                for fi22 := 0; fi22 < 20; fi22++ {
                    v23 := GenS2{F0: v19, F1: v1, F2: v2}
                    v2 += 42
                    switch fi20 {
                    case -531:
                        v24 := []int{v6, -271, (fi20 >> ((2) & 63)), (v6 % ((584) | 1))}
                        var v25 uint16 = 255
                        var v26 uint32 = genFn1()
                        fmt.Println(v4)
                        fmt.Println(v4)
                        if (genFn0() <= 4294967295) {
                            fmt.Println(v3, v3)
                            v27 := GenS1{F0: fi22, F1: v25, F2: ((v2 == v2) && false)}
                            _ = v27
                        } else {
                            if false {
                                if (!(v5 != 9223372036854775807)) {
                                    v24 = append(v24, v6)
                                    if false {
                                        v28 := GenS3{F0: (v1 ^ (v1 << ((1) & 63))), F1: v19}
                                        _ = v28
                                    } else {
                                        v29 := GenS4{F0: (v4 & v4)}
                                        v30 := GenS4{F0: v4}
                                        if (-1 > int8(v5)) {
                                            v2 -= v2
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                        } else {
                                            switch (v6 & 442) {
                                            case -458:
                                                switch fi22 {
                                                case 32768:
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                case 2147483647:
                                                    v31 := GenS3{F0: v1, F1: v19}
                                                    _ = v31
                                                default:
                                                    v6 += 1
                                                }
                                            case 127:
                                                var v32 int16 = int16(v5)
                                                _ = v32
                                            }
                                            v1 -= v1
                                            var v33 uint = 52
                                            _ = v33
                                        }
                                        if false {
                                            v24 = append(v24, fi20)
                                            v24 = append(v24, fi22)
                                            v4 = (-(v4))
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                    v34 := GenS5{F0: v5, F1: 18446744073709551615}
                                    _ = v34
                                }
                            }
                        }
                        if false {
                            var v35 int64 = v5
                            v2 -= (v2 - 1)
                            var v36 bool = (fi20 >= v6)
                            _ = v35
                            _ = v36
                        } else {
                            switch int(v25) {
                            case -915:
                                if false {
                                    if true {
                                        if len(v24) > 0 {
                                            v24[0] = (^(v6))
                                        }
                                    } else {
                                        if (v5 <= (v5 << ((771) & 63))) {
                                            fmt.Println(fi22)
                                            var v37 uint = 884
                                            _ = v37
                                        } else {
                                            v4 += (v4 / (((v4 * 410)) | 1))
                                        }
                                        if (!true) {
                                            v18 -= (v18 % (((v26 >> ((942) & 63))) | 1))
                                            var v38 int32 = -1
                                            var v39 int64 = v5
                                            _ = v38
                                            _ = v39
                                        } else {
                                            if (v26 == (^(v18))) {
                                                var v40 int32 = 1
                                                v1 = (v1 | v1)
                                                _ = len(v24)
                                                _ = v40
                                            } else {
                                                switch (v19 + "foo bar") {
                                                case "hello":
                                                    var v41 int32 = (v4 % ((-1) | 1))
                                                    _ = v41
                                                case "a":
                                                    v42 := GenS6{F0: int(v2), F1: (v4 >> ((9) & 63)), F2: v2}
                                                    _ = v42
                                                }
                                                if false {
                                                    v2 -= 32767
                                                    v2 += 32767
                                                } else {
                                                    if ((v25 >> ((genFn0()) & 63)) <= (v25 * (v25 ^ (^(v25))))) {
                                                        v2 += (v2 >> ((genFn0()) & 63))
                                                        _ = len(v24)
                                                        if true {
                                                            v6 = v6
                                                        }
                                                        var v43 float64 = ((0.1 * float64(1.5)) * 1e10)
                                                        v5 = v5
                                                        v17 += (v17 % ((v17) | 1))
                                                        _ = v43
                                                    }
                                                }
                                                var v44 int32 = (v4 ^ (v4 + 2147483647))
                                                _ = v44
                                            }
                                            fmt.Println(fi20, v25)
                                        }
                                        var v45 uint = 255
                                        _ = v45
                                    }
                                    if true {
                                        if false {
                                            if false {
                                                if false {
                                                    var v46 int32 = (^(v4))
                                                    if (false && (((true || ((((v25 | (^(v25))) >= 32767) && (int16(v1) == -32768)) || true)) && true) || true)) {
                                                        fmt.Println(v5)
                                                        fmt.Println(v18)
                                                    }
                                                    if ((3.14159 + 1.5) > float64(1e10)) {
                                                        if ((1.5 > 0.0) || (0 <= 0)) {
                                                            v1 = v1
                                                            if len(v24) >= 2 {
                                                                v24 = v24[1:len(v24)]
                                                            }
                                                            v24 = append(v24, v6)
                                                        }
                                                    }
                                                    _ = v46
                                                }
                                                _ = len(v24)
                                                v47 := GenS4{F0: 2147483647}
                                                fmt.Println(v6, v4)
                                                _ = v47
                                            }
                                            fmt.Println(v2, fi20)
                                            var v48 int = fi22
                                            _ = v48
                                        } else {
                                            switch (v6 ^ fi20) {
                                            case 1:
                                                var v49 uint16 = 2
                                                _ = v49
                                            case -1:
                                                var v50 uint64 = 2
                                                _ = v50
                                            default:
                                                switch v3 {
                                                case "a":
                                                    v51 := GenS7{F0: v25, F1: genFn0(), F2: uint(v25)}
                                                    _ = v51
                                                case "foo bar":
                                                    var v52 float32 = -2.25
                                                    _ = v52
                                                case "\n\t":
                                                    v53 := GenS8{F0: (v5 + 51), F1: 778, F2: fi20}
                                                    _ = v53
                                                }
                                            }
                                            if len(v24) > 0 {
                                                v24[0] = (-(v6))
                                            }
                                            v26 += 1
                                        }
                                        var v54 uint64 = v17
                                        v24 = append(v24, (fi22 & fi22))
                                        _ = v54
                                    }
                                } else {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    var v55 uint16 = (^(v25))
                                    _ = len(v24)
                                    if true {
                                        if ((v5 | -360) == int64(v2)) {
                                            if (v17 < v17) {
                                                v24 = append(v24, (fi20 % ((-491) | 1)))
                                                if len(v24) > 0 {
                                                    v24[0] = 128
                                                }
                                            } else {
                                                var v56 uint64 = 42
                                                if (581 >= genFn0()) {
                                                    var v57 uint = 381
                                                    var v58 float32 = 1.5
                                                    var v59 int64 = v5
                                                    v4 = (-(v4))
                                                    switch (fi22 << ((genFn0()) & 63)) {
                                                    case 270:
                                                        v5 += v59
                                                    case 9223372036854775807:
                                                        if (0.1 != float64(v58)) {
                                                            fmt.Println(fi22)
                                                            fmt.Println(v3, v57, fi20)
                                                            fmt.Println(v4)
                                                        } else {
                                                            v25 = v25
                                                            switch (fi20 % (((v6 & fi22)) | 1)) {
                                                            case 128:
                                                                v60 := GenS4{F0: (v4 + v4)}
                                                                _ = v60
                                                            case -935:
                                                                v18 -= v18
                                                            }
                                                            v24 = append(v24, (v6 * (v6 & -973)))
                                                        }
                                                    }
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                }
                                                _ = v56
                                            }
                                            v18 = 560
                                        }
                                        v19 = (("a" + "foo bar") + "\n\t")
                                        fmt.Println(v6, v19)
                                        v25 -= v25
                                        if false {
                                            fmt.Println(v17, v19, v5)
                                            _ = len(v24)
                                            var v61 uint64 = (v17 | 2)
                                            _ = v61
                                        } else {
                                            v62 := GenS9{F0: (v17 & v17)}
                                            var v63 uint16 = (v25 >> ((2) & 63))
                                            v64 := GenS10{F0: v2, F1: genFn1(), F2: 146}
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                    } else {
                                        v55 += (v55 & v25)
                                        var v65 int = 657
                                        switch v6 {
                                        case 32767:
                                            var v66 float64 = (1.5 * (0.1 * 1e10))
                                            _ = v66
                                        case 42:
                                            v55 -= (v55 ^ v55)
                                        default:
                                            if len(v24) > 0 {
                                                v24[0] = v65
                                            }
                                        }
                                        _ = v65
                                    }
                                    fmt.Println(v19)
                                    _ = v55
                                }
                            case -2147483648:
                                switch fi22 {
                                case 32767:
                                    var v67 int8 = v1
                                    _ = v67
                                default:
                                    var v68 int32 = (^(v4))
                                    _ = v68
                                }
                            default:
                                v69 := GenS11{F0: 251, F1: false, F2: v17}
                                _ = v69
                            }
                            _ = len(v24)
                            _ = len(v24)
                            v2 -= -1
                            var v70 float64 = 3.14159
                            _ = v70
                        }
                        if (!false) {
                            if (((v18 + v18) < v18) && false) {
                                var v71 string = v3
                                _ = v71
                            }
                            if len(v24) > 0 {
                                v24[0] = (^(fi20))
                            }
                            fmt.Println(v26, fi22)
                            v19 = "a"
                        } else {
                            var v72 string = "\n\t"
                            _ = v72
                        }
                        fmt.Println(v1)
                        v18 += (v18 * genFn1())
                        _ = v24
                        _ = v25
                        _ = v26
                    case -881:
                        v18 -= (v18 | v18)
                    }
                    _ = fi22
                    _ = v23
                }
            case -76:
                switch fi20 {
                case 241:
                    var v73 bool = true
                    _ = v73
                default:
                    for fi74 := 0; fi74 < 6; fi74++ {
                        v1 += 26
                        var v75 bool = false
                        _ = fi74
                        _ = v75
                    }
                }
            default:
                if ((v1 << ((18446744073709551615) & 63)) == 0) {
                    switch "a" {
                    case "a":
                        if false {
                            v1 = (v1 + v1)
                            v76 := []int{(^(v6)), -9223372036854775808, int(v17), (fi20 / ((-2147483648) | 1))}
                            _ = v76
                        }
                    default:
                        v77 := []int{(fi20 % ((fi20) | 1)), int(v2), v6, 42}
                        _ = v77
                    }
                }
            }
            fmt.Println(v19)
            _ = fi20
        }
        _ = v19
    } else {
        var v78 float64 = (((float64(0.1) * (1e10 + ((-2.25 + 0.0) - (-2.25 / 1e10)))) / (((float64((1.5 + 0.0)) - 3.14159) * (float64(1e10) + float64(((3.14159 - (1.5 * 1.5)) * 0.0)))) - 1.5)) * (0.1 - 1e10))
        var v79 uint32 = v18
        _ = v78
        _ = v79
    }
    v6 -= -356
    fmt.Println(v17)
    var v80 string = v3
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
