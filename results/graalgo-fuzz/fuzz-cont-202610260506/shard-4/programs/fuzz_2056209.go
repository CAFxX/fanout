package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int32; F1 uint; F2 string }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int64) int32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 0.1
        var v9 string = ("foo bar" + "foo bar")
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{1}
    _ = p0
    _ = v10
    return -383
}
func genFn1(p0 uint8, p1 string) uint8 {
    v11 := []int{}
    var v12 string = ("\n\t" + p1)
    v11 = append(v11, 2147483647)
    var v13 float64 = 0.0
    v14 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return (^(p0))
}

func main() {
    var v1 string = ("\n\t" + "")
    var v2 float32 = -2.25
    var v3 int = -2147483648
    var v4 uint = 18446744073709551615
    var v5 uint32 = uint32(v4)
    v6 := []uint32{}
    _ = len(v6)
    if true {
        v2 = v2
    }
    var v15 int8 = -97
    v1 = v1
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            if true {
                var v18 int = v3
                _ = v18
            } else {
                v4 = 556
                switch 1 {
                case 32767:
                    var v19 int16 = 0
                    _ = v19
                case 1:
                    if len(v6) > 0 {
                        v6[0] = 1
                    }
                default:
                    fmt.Println(v1, fi16)
                }
                if true {
                    v20 := GenS2{F0: v4, F1: (!(false || false))}
                    if (true || (!false)) {
                        if (!(!(1e10 >= (1.5 - (1.5 * 0.1))))) {
                            var v21 uint16 = 0
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch v1 {
                            case "hello":
                                var v22 string = v1
                                _ = v22
                            case "":
                                v21 = uint16(v4)
                            default:
                                v21 = v21
                            }
                            var v23 bool = ((!false) || true)
                            var v24 int32 = -155
                            v24 += (v24 - genFn0(-9223372036854775808))
                            v25 := GenS2{F0: (v4 + v4), F1: true}
                            _ = v21
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v5 -= (v5 + 0)
                            if ((!(false || true)) || (v2 != float32(0.1))) {
                                v26 := GenS2{F0: v4, F1: false}
                                _ = v26
                            }
                            var v27 uint = (v4 ^ v4)
                            var v28 int32 = 0
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS2{F0: 466, F1: false}
                        _ = v29
                    }
                    _ = v20
                }
            }
            v4 += v4
            v2 += v2
            _ = fi17
        }
        fmt.Println(v2)
        _ = fi16
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        if false {
            v3 -= v3
            if false {
                v15 = (-(v15))
                for fi31 := 0; fi31 < 3; fi31++ {
                    switch v1 {
                    case "":
                        v32 := GenS2{F0: v4, F1: (-9223372036854775808 <= 9223372036854775807)}
                        var v33 float32 = ((v2 - v2) + (v2 - 0.0))
                        _ = v32
                        _ = v33
                    case "foo bar":
                        v34 := GenS3{F0: -990, F1: 0}
                        _ = v34
                    case "\n\t":
                        if len(v6) > 0 {
                            v6[0] = 0
                        }
                    default:
                        v6 = append(v6, v5)
                    }
                    _ = fi31
                }
            }
        } else {
            v3 -= fi30
            if (356 == -785) {
                v4 += v4
                switch ("a" + (v1 + (v1 + ("a" + "hello")))) {
                case "":
                    if (!(!true)) {
                        if true {
                            v35 := GenS4{F0: 1}
                            _ = v35
                        }
                        v36 := GenS5{F0: 1, F1: true}
                        fmt.Println(v15)
                        _ = v36
                    }
                case "hello":
                    switch (v1 + v1) {
                    case "":
                        if true {
                            if (((0.1 + 0.0) - 3.14159) < -2.25) {
                                if true {
                                    if true {
                                        var v37 bool = (127 > 127)
                                        v15 = 127
                                        var v38 float32 = v2
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if (4294967295 >= (v5 + v5)) {
                                            if (255 < 1) {
                                                for fi39 := 0; fi39 < 14; fi39++ {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v40 := GenS3{F0: -1, F1: 1}
                                                        _ = v40
                                                    case "a":
                                                        if true {
                                                            var v41 int64 = 0
                                                            v6 = append(v6, (v5 ^ v5))
                                                            v41 = 1
                                                            _ = v41
                                                        } else {
                                                            fmt.Println(v1, v5)
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v15, v3)
                                                        v42 := GenS3{F0: fi30, F1: v4}
                                                        _ = v42
                                                    default:
                                                        v4 = v4
                                                    }
                                                    var v43 uint = (v4 * (v4 / ((4294967295) | 1)))
                                                    var v44 bool = true
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    _ = fi39
                                                    _ = v43
                                                    _ = v44
                                                }
                                            }
                                            switch v1 {
                                            case "a":
                                                var v45 int8 = -67
                                                _ = v45
                                            case "\n\t":
                                                v46 := GenS2{F0: 389, F1: (-2.25 <= (((1.5 / 1.0) * 1.5) + 1e10))}
                                                _ = v46
                                            case "":
                                                for fi47 := 0; fi47 < 17; fi47++ {
                                                    v48 := GenS4{F0: 1}
                                                    if len(v6) > 0 {
                                                        v6[0] = v5
                                                    }
                                                    v3 -= fi30
                                                    fmt.Println(v1, v2)
                                                    if false {
                                                        var v49 int32 = -1
                                                        _ = v49
                                                    } else {
                                                        if len(v6) >= 2 {
                                                            v6 = v6[1:len(v6)]
                                                        }
                                                    }
                                                    var v50 int16 = 870
                                                    _ = fi47
                                                    _ = v48
                                                    _ = v50
                                                }
                                            }
                                        }
                                        for fi51 := 0; fi51 < 15; fi51++ {
                                            v6 = append(v6, 42)
                                            if (false && ((v4 + v4) < 42)) {
                                                var v52 float64 = 0.0
                                                v2 = (float32(float64((v2 * ((v2 + 1e10) / v2)))) + 1e10)
                                                v53 := GenS2{F0: v4, F1: false}
                                                fmt.Println(v3)
                                                v4 = (v4 % ((v4) | 1))
                                                _ = v52
                                                _ = v53
                                            }
                                            _ = fi51
                                        }
                                        var v54 int16 = 42
                                        _ = v54
                                    }
                                    v6 = append(v6, v5)
                                    if (("" != (v1 + v1)) && true) {
                                        var v55 uint16 = 255
                                        v4 += (v4 % (((v4 + 1)) | 1))
                                        _ = v55
                                    } else {
                                        for fi56 := 0; fi56 < 18; fi56++ {
                                            v57 := GenS6{F0: -1, F1: v4, F2: "hello"}
                                            fmt.Println(v4, fi30)
                                            fmt.Println(fi56, v1)
                                            var v58 bool = ((v2 - v2) == (((v2 - v2) / float32(-2.25)) * (v2 + v2)))
                                            if len(v6) > 0 {
                                                v6[0] = v5
                                            }
                                            _ = fi56
                                            _ = v57
                                            _ = v58
                                        }
                                        for fi59 := 0; fi59 < 15; fi59++ {
                                            _ = len(v6)
                                            v60 := GenS4{F0: 880}
                                            fmt.Println(fi59, v2, v1)
                                            _ = len(v6)
                                            _ = fi59
                                            _ = v60
                                        }
                                    }
                                }
                            }
                        } else {
                            fmt.Println(v15, v4)
                        }
                    case "foo bar":
                        if (!(3.14159 > 1.5)) {
                            if true {
                                if false {
                                    v6 = append(v6, 746)
                                } else {
                                    v4 -= v4
                                }
                                if false {
                                    var v61 uint8 = 1
                                    _ = v61
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = (v5 >> ((v4) & 63))
                                    }
                                    v62 := GenS4{F0: genFn0(562)}
                                    _ = v62
                                }
                                v4 -= (^(v4))
                            }
                            fmt.Println(fi30)
                            switch (fi30 >> ((v4) & 63)) {
                            case -1:
                                _ = len(v6)
                            default:
                                v63 := GenS4{F0: 584}
                                _ = v63
                            }
                        }
                    default:
                        v1 = v1
                    }
                }
            } else {
                switch (fi30 / ((int(v4)) | 1)) {
                case 0:
                    v5 = (v5 & (v5 ^ 42))
                case 32767:
                    v64 := GenS7{F0: ((!false) && (!(1 >= -1))), F1: (^(v5))}
                    _ = v64
                case -2147483648:
                    var v65 int8 = -49
                    _ = v65
                }
            }
            if false {
                switch (-(v3)) {
                case 970:
                    _ = len(v6)
                case -454:
                    var v66 bool = (!true)
                    _ = v66
                case -314:
                    v67 := GenS8{F0: genFn0(-571), F1: 32767}
                    _ = v67
                }
                if len(v6) > 0 {
                    v6[0] = (v5 << ((v4) & 63))
                }
                switch "" {
                case "":
                    v4 = (v4 << ((v4) & 63))
                }
            } else {
                for fi68 := 0; fi68 < 20; fi68++ {
                    var v69 float32 = (0.1 / (v2 * float32(1.5)))
                    _ = fi68
                    _ = v69
                }
            }
        }
        _ = fi30
    }
    if (true || (1 == 42)) {
        for fi70 := 0; fi70 < 11; fi70++ {
            for fi71 := 0; fi71 < 20; fi71++ {
                if true {
                    if true {
                        _ = len(v6)
                        _ = len(v6)
                        if (!true) {
                            _ = len(v6)
                            if ((float64(v2) + 0.0) >= float64(v2)) {
                                if len(v6) > 0 {
                                    v6[0] = 499
                                }
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = v5
                                    }
                                    v6 = append(v6, v5)
                                } else {
                                    fmt.Println(v4, v15, v15)
                                    switch ("\n\t" + (v1 + v1)) {
                                    case "\n\t":
                                        v5 = (^(v5))
                                    case "":
                                        v72 := GenS8{F0: 0, F1: v3}
                                        _ = v72
                                    case "hello":
                                        var v73 int64 = 42
                                        _ = v73
                                    }
                                    var v74 string = v1
                                    _ = v74
                                }
                            }
                        }
                    }
                    switch (fi71 | 926) {
                    case 40:
                        var v75 uint16 = uint16(v5)
                        _ = v75
                    default:
                        if (-79 < v15) {
                            switch v1 {
                            case "a":
                                fmt.Println(fi71, fi70, v15)
                            case "hello":
                                if (("" + (v1 + (v1 + (v1 + (v1 + v1))))) > (v1 + (("a" + ((v1 + v1) + v1)) + v1))) {
                                    var v76 uint64 = 42
                                    if true {
                                        if (-2.25 > ((0.1 - 0.1) - (1e10 * 0.0))) {
                                            v77 := GenS9{F0: 2147483647, F1: 930}
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v78 float64 = 0.0
                                            fmt.Println(v78)
                                            if len(v6) > 0 {
                                                v6[0] = v5
                                            }
                                            _ = v77
                                            _ = v78
                                        }
                                        v2 = v2
                                        var v79 int8 = (v15 >> (((^(v4))) & 63))
                                        _ = v79
                                    } else {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        v4 += v4
                                    }
                                    if false {
                                        var v80 int8 = v15
                                        var v81 uint16 = 2
                                        _ = v80
                                        _ = v81
                                    }
                                    _ = v76
                                }
                            case "\n\t":
                                var v82 uint64 = 2
                                _ = v82
                            default:
                                v1 = ("" + v1)
                            }
                            if (299 < 589) {
                                var v83 int8 = (-(v15))
                                switch (fi71 % (((fi70 << (((v4 / (((^(v4))) | 1))) & 63))) | 1)) {
                                case 42:
                                    var v84 string = v1
                                    _ = v84
                                case -252:
                                    v85 := GenS10{F0: 4294967295}
                                    _ = v85
                                case -585:
                                    var v86 int16 = -32768
                                    _ = v86
                                default:
                                    var v87 uint16 = 2
                                    if (!((-2147483648 == 395) && (true && true))) {
                                        v88 := GenS1{F0: (v5 & v5)}
                                        _ = v88
                                    }
                                    _ = v87
                                }
                                _ = v83
                            } else {
                                var v89 int64 = 42
                                _ = v89
                            }
                            v15 = (v15 | 42)
                        } else {
                            _ = len(v6)
                            var v90 uint64 = 1
                            v1 = v1
                            _ = v90
                        }
                    }
                } else {
                    var v91 uint64 = 1
                    var v92 uint32 = v5
                    _ = v91
                    _ = v92
                }
                v93 := GenS8{F0: 618, F1: 127}
                _ = fi71
                _ = v93
            }
            for fi94 := 0; fi94 < 11; fi94++ {
                v95 := GenS4{F0: int32(fi94)}
                _ = fi94
                _ = v95
            }
            var v96 float64 = 1.5
            _ = fi70
            _ = v96
        }
    } else {
        fmt.Println(v5, v5)
    }
    var v97 uint8 = 128
    if true {
        switch ((v1 + ((v1 + (((v1 + (("foo bar" + v1) + v1)) + v1) + v1)) + v1)) + v1) {
        case "":
            var v98 float32 = 1e10
            _ = v98
        default:
            v97 = (^(v97))
        }
        var v99 float64 = 0.0
        _ = v99
    }
    fmt.Println(v97)
    for fi100 := 0; fi100 < 15; fi100++ {
        var v101 uint8 = 255
        _ = fi100
        _ = v101
    }
    v6 = append(v6, (v5 % (((v5 / ((2147483648) | 1))) | 1)))
    if false {
        v2 += ((v2 / float32((0.1 * 0.1))) - -2.25)
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v97
}
