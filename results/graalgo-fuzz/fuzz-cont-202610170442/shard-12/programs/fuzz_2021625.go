package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 int; F2 int64 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS13 struct { F0 int; F1 int }
type GenS14 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    v4 := []uint32{926, 2147483647, 500}
    v5 := []uint32{42}
    v6 := GenS1{F0: 18446744073709551615, F1: -1, F2: (1.5 <= -2.25)}
    if (2 != 65535) {
        delete(v1, (("foo bar" + ("foo bar" + v2)) + ("foo bar" + v2)))
        if true {
            if true {
                var v7 uint16 = 804
                if false {
                    v8 := GenS2{F0: 218}
                    var v9 int64 = 0
                    if true {
                        v10 := GenS3{F0: 1, F1: ("hello" + ((v2 + "foo bar") + v2)), F2: 42}
                        _ = v10
                    }
                    _ = v8
                    _ = v9
                }
                _ = v7
            } else {
                v11 := GenS4{F0: "\n\t"}
                v12 := GenS5{F0: -2147483648}
                _ = len(v1)
                _ = v11
                _ = v12
            }
        }
        v13 := GenS5{F0: 2147483647}
        _ = v13
    } else {
        delete(v3, -915)
        if true {
            v14 := GenS5{F0: -1}
            switch ("" + v2) {
            case "foo bar":
                var v15 int8 = 42
                _ = v15
            case "":
                fmt.Println(v2)
            default:
                v2 = (("a" + ("\n\t" + "")) + v2)
            }
            var v16 string = v2
            _ = v14
            _ = v16
        }
    }
    v2 = "foo bar"
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                var v19 uint8 = 42
                v2 = (v2 + "foo bar")
                _ = v19
            } else {
                if ((!true) || true) {
                    fmt.Println(fi18, fi17, fi18)
                    delete(v1, v2)
                }
                if true {
                    _ = len(v5)
                    v20 := GenS6{F0: 419}
                    _ = v20
                }
                var v21 int = fi17
                _ = v21
            }
            v2 = v2
            if (false && true) {
                var v22 int16 = 42
                _ = v22
            }
            _ = fi18
        }
        _ = fi17
    }
    switch 127 {
    case -426:
        v2 = v2
    case -425:
        v3[127] = 2
    default:
        if true {
            v4 = append(v4, 0)
            fmt.Println(v2)
        }
    }
    v23 := GenS7{F0: 443, F1: 42, F2: 0}
    var v24 uint = 255
    var v25 string = ((v2 + ("a" + "")) + "foo bar")
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v25, fi26)
        var v27 uint64 = 4294967295
        _ = fi26
        _ = v27
    }
    v25 = v25
    var v28 float64 = 0.0
    if (!false) {
        for fi29 := 0; fi29 < 6; fi29++ {
            switch (v2 + "foo bar") {
            case "foo bar":
                v5 = append(v5, 688)
            default:
                v28 = 1e10
            }
            if ((v24 >> (((v24 << (((v24 * v24)) & 63))) & 63)) >= 18446744073709551615) {
                if (!true) {
                    for fi30 := 0; fi30 < 1; fi30++ {
                        v28 += (3.14159 * 1e10)
                        var v31 uint8 = 128
                        if (584 <= 326) {
                            if (!(!(9223372036854775807 >= 9223372036854775807))) {
                                var v32 uint32 = 65535
                                v32 -= v32
                                var v33 int32 = 0
                                _ = v32
                                _ = v33
                            } else {
                                var v34 string = v2
                                fmt.Println(v28, v2, v25)
                                _ = v34
                            }
                            switch v25 {
                            case "hello":
                                v2 = (v25 + "foo bar")
                            case "":
                                var v35 float64 = v28
                                _ = v35
                            case "\n\t":
                                v36 := v1[""]
                                _ = v36
                            default:
                                if (((0.1 / 1e10) <= 1e10) && false) {
                                    if true {
                                        fmt.Println(v25, v31)
                                        v24 = v24
                                    }
                                    delete(v3, (fi29 * (fi29 + fi29)))
                                    var v37 int8 = 85
                                    _ = v37
                                }
                            }
                        } else {
                            switch v25 {
                            case "\n\t":
                                v28 -= v28
                            case "foo bar":
                                if (418 < 42) {
                                    var v38 uint32 = 1
                                    v31 += (v31 & 248)
                                    _ = v38
                                } else {
                                    v28 = (3.14159 - v28)
                                    var v39 float64 = ((v28 * v28) + v28)
                                    _ = v39
                                }
                            default:
                                var v40 float64 = float64(0.1)
                                _ = v40
                            }
                            if (!false) {
                                if len(v4) > 0 {
                                    v4[0] = 409
                                }
                                delete(v3, fi30)
                                if ((true && true) || true) {
                                    if (false && ((v31 + (v31 ^ (^(v31)))) <= (^(v31)))) {
                                        if (false || true) {
                                            v24 += 143
                                        }
                                    }
                                } else {
                                    fmt.Println(v31, v2, v2)
                                    fmt.Println(v31)
                                }
                            } else {
                                v28 += 0.0
                                if (0 > 626) {
                                    if (!false) {
                                        var v41 int16 = -278
                                        if (true && false) {
                                            var v42 uint8 = (^(v31))
                                            v42 = v42
                                            _ = v42
                                        } else {
                                            var v43 int64 = -1
                                            _ = v43
                                        }
                                        _ = v41
                                    } else {
                                        if ((1e10 * (1.5 * 3.14159)) <= 1e10) {
                                            v24 += 875
                                            var v44 int = fi30
                                            if false {
                                                if ((42 == 1) && false) {
                                                    v2 = v25
                                                    var v45 int64 = -1
                                                    v2 = ""
                                                    _ = v45
                                                } else {
                                                    if (((v24 >> ((42) & 63)) > v24) && (!true)) {
                                                        if false {
                                                            v46 := GenS1{F0: uint64(v31), F1: 531, F2: (v2 <= (v2 + ("hello" + v25)))}
                                                            v25 = v2
                                                            _ = v46
                                                        } else {
                                                            var v47 uint = (v24 ^ 18446744073709551615)
                                                            var v48 float64 = ((((((0.1 * v28) + -2.25) * 3.14159) / 1.0) + (v28 + v28)) - ((((v28 + 1e10) - v28) + (-2.25 / ((v28 + (1.5 + float64(((1e10 - -2.25) * (0.1 - 3.14159))))) - v28))) * 0.1))
                                                            if true {
                                                                var v49 int8 = 105
                                                                var v50 float64 = (1.5 + v48)
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            fmt.Println(fi30)
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        v51 := GenS8{F0: v2, F1: v2}
                                                        v24 = v24
                                                        var v52 uint16 = 2
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    if true {
                                                        var v53 bool = ((((!(!true)) || ((fi29 / ((fi29) | 1)) == fi30)) || false) && (-1 != 32767))
                                                        fmt.Println(v28, fi30, v28)
                                                        _ = v53
                                                    } else {
                                                        v54 := GenS7{F0: 0, F1: v44, F2: -978}
                                                        v55 := GenS9{F0: -2147483648}
                                                        var v56 float64 = float64(0.1)
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    v1[v2] = 0
                                                }
                                                switch 4294967295 {
                                                case -2147483648:
                                                    v57 := GenS10{F0: 15}
                                                    _ = v57
                                                case 648:
                                                    v2 = ""
                                                case 0:
                                                    v44 = (v44 | 286)
                                                default:
                                                    var v58 int8 = 0
                                                    _ = v58
                                                }
                                            }
                                            _ = v44
                                        }
                                        if false {
                                            var v59 int64 = 0
                                            _ = v59
                                        }
                                    }
                                    var v60 int16 = -1
                                    var v61 uint32 = 1
                                    _ = v60
                                    _ = v61
                                } else {
                                    if true {
                                        v2 = v2
                                        v62 := GenS11{F0: 42}
                                        fmt.Println(v28)
                                        if (-2.25 > 3.14159) {
                                            var v63 bool = (65535 == (^(v24)))
                                            var v64 int = (fi30 >> ((v24) & 63))
                                            v64 += fi29
                                            _ = v63
                                            _ = v64
                                        } else {
                                            var v65 int = (fi30 % (((-(fi29))) | 1))
                                            v28 += v28
                                            v65 = fi30
                                            _ = v65
                                        }
                                        var v66 bool = true
                                        _ = v62
                                        _ = v66
                                    } else {
                                        fmt.Println(v2)
                                    }
                                    delete(v1, v25)
                                }
                            }
                        }
                        v24 = 255
                        var v67 int = fi29
                        v68 := GenS1{F0: 235, F1: 1000, F2: (65535 >= uint16(v31))}
                        fmt.Println(v31, v31)
                        _ = fi30
                        _ = v31
                        _ = v67
                        _ = v68
                    }
                    v2 = ((v2 + "hello") + v25)
                    fmt.Println(v25, v2, v2)
                }
                v28 = 1e10
            }
            _ = fi29
        }
        if false {
            fmt.Println(v2)
            if (true && false) {
                v24 = v24
                for fi69 := 0; fi69 < 16; fi69++ {
                    v24 -= v24
                    fmt.Println(v25, v25, v28)
                    _ = fi69
                }
                var v70 int = 925
                _ = v70
            }
            if ((v24 <= v24) && true) {
                switch v25 {
                case "hello":
                    for fi71 := 0; fi71 < 1; fi71++ {
                        _ = len(v3)
                        var v72 float64 = 1.5
                        var v73 string = ("hello" + v25)
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                case "\n\t":
                    switch -9223372036854775808 {
                    case 127:
                        var v74 string = v25
                        _ = v74
                    case 128:
                        v75 := v3[1]
                        _ = v75
                    case 185:
                        var v76 int32 = 1
                        _ = v76
                    default:
                        v24 -= v24
                    }
                case "foo bar":
                    if ((((218 <= 0) || false) || false) || false) {
                        fmt.Println(v28, v28, v25)
                        v77 := GenS12{F0: 2147483647, F1: 4294967295, F2: 2}
                        _ = v77
                    } else {
                        _ = len(v5)
                        switch "" {
                        case "\n\t":
                            fmt.Println(v28)
                        default:
                            var v78 uint = v24
                            _ = v78
                        }
                    }
                }
                for fi79 := 0; fi79 < 18; fi79++ {
                    v80 := GenS13{F0: (-(fi79)), F1: (fi79 / ((int(v24)) | 1))}
                    var v81 string = v2
                    var v82 int8 = 96
                    _ = fi79
                    _ = v80
                    _ = v81
                    _ = v82
                }
            }
        } else {
            v2 = (("foo bar" + ((v2 + v25) + "a")) + (v2 + "hello"))
        }
    } else {
        v24 -= 42
        for fi83 := 0; fi83 < 8; fi83++ {
            v84 := GenS10{F0: 999}
            if (false && true) {
                v4 = append(v4, 446)
            } else {
                v24 += (v24 ^ v24)
                var v85 uint8 = 18
                _ = v85
            }
            v86 := GenS14{F0: -343, F1: 255}
            _ = fi83
            _ = v84
            _ = v86
        }
        var v87 uint32 = 4294967295
        _ = v87
    }
    _ = len(v5)
    if true {
        if false {
            switch 32768 {
            case -1:
                v88 := GenS1{F0: 42, F1: 32767, F2: (!false)}
                _ = v88
            case 645:
                if true {
                    v24 -= v24
                } else {
                    for fi89 := 0; fi89 < 17; fi89++ {
                        if (4294967295 >= 522) {
                            var v90 bool = ((v24 << ((v24) & 63)) == 42)
                            _ = v90
                        }
                        _ = fi89
                    }
                    v24 += v24
                    for fi91 := 0; fi91 < 13; fi91++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi91
                    }
                }
            }
        }
    }
    _ = len(v1)
    switch v25 {
    case "hello":
        var v92 float32 = -2.25
        _ = v92
    case "a":
        var v93 uint16 = 975
        _ = v93
    default:
        var v94 bool = (-128 == -55)
        _ = v94
    }
    if false {
        v25 = v25
    }
    switch (v2 + v25) {
    case "\n\t":
        if (0 < -9223372036854775808) {
            fmt.Println(v28, v25)
            fmt.Println(v25)
        }
    case "a":
        var v95 bool = (!((true && (!true)) && (v28 != ((((3.14159 * v28) - v28) / 1e10) / v28))))
        _ = v95
    case "foo bar":
        _ = len(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
