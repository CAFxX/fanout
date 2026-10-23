package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int; F2 uint8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS11 struct { F0 int64; F1 string }

func main() {
    v1 := []string{"a", ("foo bar" + "foo bar"), "hello", "foo bar"}
    v2 := make(map[int]string)
    var v3 float32 = (float32(-2.25) - 0.0)
    var v4 int64 = 0
    v5 := make(map[string]uint32)
    var v6 uint64 = 0
    if true {
        v6 = 4294967295
        v4 = v4
        if true {
            switch 1 {
            case -9223372036854775808:
                v6 -= (v6 - (^(v6)))
            case 32768:
                v3 -= (v3 - ((v3 * v3) - ((1.5 + 0.0) * 3.14159)))
            default:
                _ = len(v2)
            }
        } else {
            for fi7 := 0; fi7 < 13; fi7++ {
                v8 := GenS1{F0: "hello", F1: 0, F2: 18446744073709551615}
                for fi9 := 0; fi9 < 11; fi9++ {
                    if ((false && (false && (-128 > 1))) || (32767 <= 972)) {
                        var v10 uint32 = 1
                        _ = v10
                    } else {
                        v11 := GenS2{F0: -2147483648, F1: (^(v6))}
                        _ = v11
                    }
                    _ = len(v5)
                    v12 := GenS3{F0: v4, F1: fi9}
                    _ = fi9
                    _ = v12
                }
                if true {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 uint16 = 1
                        if false {
                            v4 += (-(v4))
                            var v15 float64 = 1.5
                            _ = v15
                        } else {
                            v3 += 1e10
                            if true {
                                var v16 int64 = v4
                                switch "\n\t" {
                                case "a":
                                    v3 += -2.25
                                }
                                _ = v16
                            }
                            if false {
                                switch "a" {
                                case "foo bar":
                                    v17 := GenS4{F0: -9223372036854775808}
                                    _ = v17
                                default:
                                    v3 = v3
                                }
                                if (-599 < -1) {
                                    var v18 int8 = 1
                                    var v19 int8 = v18
                                    _ = v18
                                    _ = v19
                                } else {
                                    fmt.Println(v6, v14, v3)
                                    v20 := GenS2{F0: 42, F1: 0}
                                    var v21 string = "foo bar"
                                    _ = v20
                                    _ = v21
                                }
                                fmt.Println(v4, fi7)
                            }
                        }
                        v4 = int64(fi13)
                        _ = fi13
                        _ = v14
                    }
                    var v22 int64 = v4
                    _ = v22
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v23 := GenS3{F0: v4, F1: (fi7 << ((1) & 63))}
                        _ = v23
                    default:
                        fmt.Println(v4)
                    }
                    var v24 uint16 = 42
                    fmt.Println(fi7, v24)
                    _ = v24
                }
                _ = fi7
                _ = v8
            }
        }
    } else {
        v6 += (v6 / ((871) | 1))
        var v25 uint = 2
        _ = v25
    }
    var v26 uint = 0
    v27 := GenS5{F0: true}
    var v28 string = "\n\t"
    if false {
        switch -93 {
        case -1:
            if (true && true) {
                v3 = v3
                v28 = ("hello" + (v28 + v28))
            } else {
                v29 := GenS4{F0: (v4 % ((42) | 1))}
                if true {
                    for fi30 := 0; fi30 < 13; fi30++ {
                        for fi31 := 0; fi31 < 2; fi31++ {
                            var v32 float32 = ((float32(0.0) + v3) * v3)
                            v33 := GenS1{F0: v28, F1: (v6 | uint64(v26)), F2: v6}
                            _ = len(v1)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        v3 = v3
                        v26 = v26
                        _ = fi30
                    }
                    v34 := GenS5{F0: false}
                    _ = v34
                }
                _ = v29
            }
        default:
            v3 += (v3 + v3)
        }
    } else {
        switch 32767 {
        case -546:
            var v35 float32 = 1e10
            _ = v35
        case 142:
            v3 -= float32(0.1)
        case 9223372036854775807:
            for fi36 := 0; fi36 < 6; fi36++ {
                var v37 int = fi36
                switch ((v28 + v28) + "") {
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                default:
                    var v38 int64 = 9223372036854775807
                    _ = v38
                }
                if false {
                    delete(v2, 9223372036854775807)
                    v6 += uint64(v26)
                }
                _ = fi36
                _ = v37
            }
        default:
            var v39 float64 = 0.0
            _ = v39
        }
        var v40 int64 = v4
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            fmt.Println(v40, v3, v26)
            v41 := GenS3{F0: v4, F1: 127}
            _ = v41
        }
        _ = v40
    }
    v5["a"] = 2147483648
    var v42 uint = (v26 * (v26 >> ((v26) & 63)))
    fmt.Println(v26, v42, v3)
    var v43 uint8 = 127
    for fi44 := 0; fi44 < 1; fi44++ {
        if true {
            if true {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                var v45 uint64 = v6
                fmt.Println(v28, v45)
                _ = v45
            } else {
                fmt.Println(v6, v42)
                v4 -= 0
            }
            fmt.Println(v43, v28)
            if (false && true) {
                fmt.Println(v4, v6)
                v28 = v28
            }
        } else {
            var v46 uint16 = 32767
            _ = v46
        }
        if false {
            if (128 == (v43 ^ (v43 ^ (^(v43))))) {
                v3 = 0.0
                v4 += (v4 * (v4 * -9223372036854775808))
                var v47 uint16 = 65535
                _ = v47
            } else {
                v26 = 18446744073709551615
                if false {
                    switch "a" {
                    case "foo bar":
                        for fi48 := 0; fi48 < 7; fi48++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v49 := GenS4{F0: (v4 / (((v4 << ((v26) & 63))) | 1))}
                            if (((v3 * ((((float32(-2.25) / (v3 / (v3 - v3))) + 3.14159) / -2.25) - v3)) / (v3 * 0.0)) >= ((v3 * v3) * ((v3 + v3) + 1.5))) {
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    var v50 int8 = -37
                                    _ = v50
                                case "":
                                    v51 := GenS2{F0: -150, F1: uint64(v42)}
                                    _ = v51
                                case "hello":
                                    var v52 float32 = -2.25
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v42)
                                v53 := GenS3{F0: -1, F1: (fi48 & -573)}
                                var v54 uint64 = 502
                                _ = v53
                                _ = v54
                            }
                            var v55 int = (fi44 >> ((v42) & 63))
                            _ = fi48
                            _ = v49
                            _ = v55
                        }
                    case "hello":
                        v3 -= v3
                    default:
                        v56 := GenS6{F0: 32767, F1: fi44, F2: 255}
                        _ = v56
                    }
                    var v57 uint = (v42 | (^(v26)))
                    _ = v57
                }
            }
        } else {
            v58 := GenS7{F0: (v28 + "hello"), F1: v28}
            _ = v58
        }
        _ = fi44
    }
    if (!(v3 > (3.14159 * float32(0.1)))) {
        fmt.Println(v4)
        var v59 int = -136
        if true {
            fmt.Println(v4, v3, v6)
        } else {
            if (-497 <= -878) {
                v60 := GenS8{F0: (v59 & (v59 >> ((v26) & 63)))}
                for fi61 := 0; fi61 < 11; fi61++ {
                    var v62 uint16 = 42
                    fmt.Println(v62, v59, v28)
                    _ = fi61
                    _ = v62
                }
                var v63 float64 = 1e10
                _ = v60
                _ = v63
            } else {
                fmt.Println(v4, v28, v43)
                var v64 string = v28
                _ = len(v1)
                _ = v64
            }
            fmt.Println(v42, v4)
        }
        _ = v59
    }
    switch ("" + "a") {
    case "a":
        delete(v2, 90)
    case "\n\t":
        if ((v4 == v4) && (v4 >= -748)) {
            if ((true || (v42 == v26)) && (!((3.14159 <= 0.1) && (42 != uint16(v43))))) {
                fmt.Println(v3)
                if false {
                    for fi65 := 0; fi65 < 2; fi65++ {
                        v66 := GenS9{F0: v43, F1: 127}
                        v67 := GenS7{F0: ("a" + "a"), F1: v28}
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                } else {
                    v3 -= v3
                }
                switch v28 {
                case "\n\t":
                    for fi68 := 0; fi68 < 3; fi68++ {
                        v26 += (v26 | v42)
                        for fi69 := 0; fi69 < 17; fi69++ {
                            v26 += v42
                            v70 := v5[v28]
                            _ = fi69
                            _ = v70
                        }
                        var v71 float32 = 1e10
                        _ = fi68
                        _ = v71
                    }
                default:
                    switch "a" {
                    case "a":
                        var v72 float32 = (v3 - (v3 - v3))
                        _ = v72
                    }
                }
            }
        }
    case "":
        v42 = (v26 / ((v26) | 1))
    }
    for fi73 := 0; fi73 < 16; fi73++ {
        fmt.Println(fi73, v43)
        _ = fi73
    }
    v74 := GenS3{F0: v4, F1: 32768}
    for fi75 := 0; fi75 < 5; fi75++ {
        v76 := GenS4{F0: v4}
        _ = fi75
        _ = v76
    }
    if false {
        v1 = append(v1, v28)
        if false {
            v28 = v28
        }
        v6 = (v6 | (^(v6)))
    }
    switch v28 {
    case "\n\t":
        switch 32768 {
        case 2147483647:
            for fi77 := 0; fi77 < 4; fi77++ {
                var v78 int64 = (v4 & v4)
                v78 += (v4 % (((v4 & 702)) | 1))
                if len(v1) > 0 {
                    v1[0] = (((v28 + v28) + v28) + (("foo bar" + v28) + "a"))
                }
                _ = fi77
                _ = v78
            }
        case 127:
            v79 := GenS5{F0: true}
            _ = v79
        case 80:
            switch 4294967295 {
            case 128:
                v6 = 42
            case 32767:
                delete(v2, 312)
            case 32768:
                var v80 uint = (v26 + 750)
                _ = v80
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case "":
        if false {
            for fi81 := 0; fi81 < 15; fi81++ {
                var v82 int16 = 32767
                var v83 int32 = 952
                for fi84 := 0; fi84 < 12; fi84++ {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        case "hello":
                            v83 = 0
                        default:
                            v83 = v83
                        }
                        v26 = (^(v42))
                    }
                    _ = fi84
                }
                _ = fi81
                _ = v82
                _ = v83
            }
            if true {
                var v85 uint8 = 220
                _ = v85
            }
            v2[424] = (v28 + v28)
        } else {
            fmt.Println(v42, v43)
            v86 := GenS10{F0: 2, F1: -128, F2: uint8(v6)}
            if false {
                _ = len(v1)
            }
            _ = v86
        }
    case "a":
        v87 := GenS11{F0: (-(v4)), F1: ("foo bar" + (v28 + v28))}
        _ = v87
    }
    var v88 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v74
    _ = v88
}
