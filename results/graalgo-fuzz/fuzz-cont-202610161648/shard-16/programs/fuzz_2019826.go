package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 int64 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 uint32; F1 int8; F2 string }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 bool; F1 int64; F2 string }
type GenS8 struct { F0 uint32; F1 bool }
type GenS9 struct { F0 int; F1 int32; F2 string }
type GenS10 struct { F0 uint; F1 int16; F2 uint32 }
type GenS11 struct { F0 int64; F1 int8; F2 uint }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 uint32; F1 bool }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint; F1 string }
type GenS16 struct { F0 int64; F1 uint8 }
type GenS17 struct { F0 int64 }

func main() {
    var v1 int32 = 42
    var v2 uint16 = 562
    v3 := make(map[uint32]string)
    fmt.Println(v1, v2)
    v1 -= 767
    v2 += 65535
    v4 := GenS1{F0: -9223372036854775808}
    var v5 uint = 1
    v6 := []int{}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = (v2 >> (((v5 | v5)) & 63))
        _ = fi7
        _ = v8
    }
    if true {
        v5 = (^(v5))
    } else {
        v9 := GenS1{F0: 9223372036854775807}
        var v10 uint8 = 186
        v6 = append(v6, 32767)
        _ = v10
        _ = v9
    }
    if (false && (1 <= 42)) {
        if false {
            switch -9223372036854775808 {
            case 128:
                if true {
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v1 += (v1 / (((v1 - v1)) | 1))
                        for fi12 := 0; fi12 < 11; fi12++ {
                            var v13 uint8 = 128
                            var v14 uint64 = 0
                            fmt.Println(v1)
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        _ = fi11
                    }
                }
            case 127:
                v5 = 1
            case 1:
                v5 += v5
            default:
                v5 = (v5 >> ((v5) & 63))
            }
        } else {
            for fi15 := 0; fi15 < 2; fi15++ {
                v2 = (v2 | 353)
                v5 += (v5 ^ (v5 * (v5 | (v5 << ((v5) & 63)))))
                _ = fi15
            }
            v16 := GenS1{F0: -732}
            v5 = (v5 % (((v5 ^ 1)) | 1))
            _ = v16
        }
        v17 := GenS2{F0: (v2 % ((v2) | 1))}
        _ = v17
    } else {
        v1 += (v1 / ((0) | 1))
        if (true && ((((!((true || true) && ((false || false) || true))) && true) && ((445 != 61) && true)) && (float32(((-2.25 + (3.14159 - (((-2.25 / 1e10) / (0.0 - (1e10 / (3.14159 + 0.0)))) + 3.14159))) * 1.5)) != (((0.0 - (3.14159 * -2.25)) + 0.0) + 1e10)))) {
            var v18 int16 = 1000
            _ = v18
        } else {
            if len(v6) > 0 {
                v6[0] = 128
            }
        }
    }
    v5 += (v5 << ((v5) & 63))
    v19 := GenS3{F0: v1, F1: 255, F2: 9223372036854775807}
    var v20 bool = false
    var v21 int16 = -568
    if (!(1 >= 0)) {
        fmt.Println(v1, v1, v21)
    }
    switch 4294967295 {
    case -9223372036854775808:
        v22 := GenS3{F0: 42, F1: 42, F2: 0}
        _ = v22
    case 128:
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 float64 = (3.14159 / 1e10)
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 bool = true
                if (v21 > (v21 % ((v21) | 1))) {
                    if v20 {
                        switch 2147483647 {
                        case 32768:
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v5)
                            case "\n\t":
                                var v27 uint16 = v2
                                _ = v27
                            case "":
                                if (((!(v5 == v5)) && v20) || v20) {
                                    var v28 int16 = v21
                                    switch int(v2) {
                                    case 605:
                                        switch fi25 {
                                        case 42:
                                            v5 += (^(v5))
                                        case -498:
                                            if (v26 && ("a" > ("hello" + "\n\t"))) {
                                                v29 := GenS4{F0: -125, F1: 42}
                                                _ = v29
                                            }
                                        case 745:
                                            var v30 int = fi25
                                            _ = v30
                                        }
                                    default:
                                        var v31 uint8 = 255
                                        _ = v31
                                    }
                                    _ = v28
                                }
                            default:
                                var v32 bool = v26
                                _ = v32
                            }
                        case 840:
                            var v33 bool = v20
                            _ = v33
                        default:
                            v20 = v20
                        }
                        if false {
                            v34 := GenS2{F0: (v2 ^ v2)}
                            v21 -= (v21 + (v21 * (^(v21))))
                            _ = v34
                        } else {
                            var v35 uint16 = v2
                            if ((-(fi23)) == fi23) {
                                if v26 {
                                    if v20 {
                                        if (-1 >= (v1 << (((v5 | (v5 & (v5 ^ (^(v5)))))) & 63))) {
                                            if v26 {
                                                fmt.Println(v20, fi25)
                                                var v36 int32 = (^(v1))
                                                _ = v36
                                            } else {
                                                v21 += -32768
                                            }
                                            if false {
                                                var v37 bool = true
                                                _ = v37
                                            } else {
                                                if (float32((v24 * 1.5)) > (0.1 / 1.5)) {
                                                    switch "" {
                                                    case "":
                                                        var v38 float64 = -2.25
                                                        _ = v38
                                                    default:
                                                        v39 := GenS3{F0: int32(v21), F1: 8, F2: -9223372036854775808}
                                                        fmt.Println(v5, v5, v1)
                                                        _ = v39
                                                    }
                                                } else {
                                                    fmt.Println(fi25, v5)
                                                    fmt.Println(v21)
                                                }
                                                var v40 int32 = v1
                                                if (!v20) {
                                                    _ = len(v6)
                                                    v41 := GenS5{F0: 371, F1: 1, F2: ""}
                                                    _ = v41
                                                }
                                                _ = v40
                                            }
                                        }
                                        v42 := GenS4{F0: 127, F1: 42}
                                        v3[2147483648] = ""
                                        _ = len(v6)
                                        _ = v42
                                    }
                                    fmt.Println(v5)
                                    v43 := GenS3{F0: v1, F1: uint8(v2), F2: -1}
                                    _ = v43
                                }
                            } else {
                                v2 = uint16(v5)
                            }
                            _ = v35
                        }
                        v5 = (^(v5))
                    } else {
                        if (1.5 < 1.5) {
                            var v44 int8 = 127
                            v45 := GenS5{F0: 42, F1: int8(v5), F2: "hello"}
                            _ = v44
                            _ = v45
                        }
                        v24 += ((v24 * (1e10 - v24)) - float64(0.0))
                        if (v20 || false) {
                            v46 := GenS6{F0: v26, F1: "a"}
                            var v47 uint16 = 191
                            _ = v46
                            _ = v47
                        } else {
                            if (!false) {
                                var v48 float32 = (float32(v24) + float32((0.0 * 3.14159)))
                                var v49 uint = 255
                                _ = len(v6)
                                _ = v48
                                _ = v49
                            } else {
                                v50 := GenS4{F0: -103, F1: 9223372036854775807}
                                v51 := GenS4{F0: -102, F1: -1}
                                _ = v50
                                _ = v51
                            }
                            v3[2147483648] = (("\n\t" + "") + "\n\t")
                        }
                    }
                    switch (fi25 * fi23) {
                    case 128:
                        v21 -= v21
                    case 32767:
                        v2 += 255
                    case 1:
                        if v26 {
                            var v52 uint = v5
                            _ = v52
                        } else {
                            v21 = v21
                            _ = len(v3)
                            v53 := v3[2147483647]
                            v53 = ((("foo bar" + "hello") + v53) + "a")
                            if (false || (!v26)) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            _ = v53
                        }
                    default:
                        var v54 uint = v5
                        _ = v54
                    }
                    v55 := GenS7{F0: (116 != -104), F1: 176, F2: "\n\t"}
                    fmt.Println(v21, v21)
                    v21 = -720
                    _ = v55
                } else {
                    if v20 {
                        switch (-(fi23)) {
                        case 32768:
                            v20 = (-128 >= 110)
                        default:
                            fmt.Println(v1, v20)
                            switch (^(fi25)) {
                            case -2147483648:
                                v56 := GenS7{F0: (0 > 2), F1: -268, F2: "a"}
                                _ = v56
                            case 128:
                                switch ("" + "") {
                                case "hello":
                                    _ = len(v6)
                                    var v57 int64 = int64(v21)
                                    fmt.Println(v1, v26, v5)
                                    if true {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if (42 <= 0) {
                                            var v58 int32 = int32(v5)
                                            v2 = (v2 - 612)
                                            _ = v58
                                        } else {
                                            var v59 uint16 = v2
                                            switch "a" {
                                            case "a":
                                                v6 = append(v6, fi25)
                                            default:
                                                v6 = append(v6, fi25)
                                            }
                                            var v60 float32 = 0.1
                                            _ = v59
                                            _ = v60
                                        }
                                        v6 = append(v6, fi23)
                                        if v26 {
                                            switch ("a" + "hello") {
                                            case "foo bar":
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            case "\n\t":
                                                fmt.Println(v2, v5)
                                            case "":
                                                fmt.Println(fi23)
                                            default:
                                                var v61 int32 = (v1 >> ((v5) & 63))
                                                _ = v61
                                            }
                                            fmt.Println(v26, fi25, v2)
                                        } else {
                                            var v62 int16 = v21
                                            _ = v62
                                        }
                                    } else {
                                        v63 := GenS4{F0: -112, F1: v57}
                                        switch (fi23 / ((fi23) | 1)) {
                                        case 1:
                                            var v64 string = ("a" + "")
                                            _ = v64
                                        default:
                                            var v65 float64 = v24
                                            _ = v65
                                        }
                                        _ = v63
                                    }
                                    v21 += (v21 << ((v5) & 63))
                                    v66 := GenS8{F0: 1, F1: v26}
                                    v21 -= 1
                                    v67 := GenS9{F0: -195, F1: 2147483647, F2: ""}
                                    _ = v57
                                    _ = v66
                                    _ = v67
                                case "\n\t":
                                    v20 = v26
                                case "foo bar":
                                    switch "\n\t" {
                                    case "hello":
                                        v68 := GenS3{F0: (v1 << ((uint(v2)) & 63)), F1: 2, F2: 9223372036854775807}
                                        _ = v68
                                    case "a":
                                        if v20 {
                                            var v69 int32 = int32(fi23)
                                            fmt.Println(v69)
                                            var v70 uint32 = 0
                                            _ = v69
                                            _ = v70
                                        } else {
                                            if v20 {
                                                switch (fi23 * fi25) {
                                                case 0:
                                                    var v71 int8 = 64
                                                    _ = v71
                                                default:
                                                    v72 := GenS1{F0: -9223372036854775808}
                                                    _ = v72
                                                }
                                                var v73 uint16 = v2
                                                _ = v73
                                            }
                                            v74 := GenS10{F0: (v5 >> (((v5 / (((^(v5))) | 1))) & 63)), F1: (v21 * (v21 & (-(v21)))), F2: 300}
                                            _ = v74
                                        }
                                    case "":
                                        v24 = v24
                                    }
                                }
                            case 713:
                                fmt.Println(v5, v1, v2)
                            }
                            v21 = (v21 * v21)
                        }
                    } else {
                        v2 = 583
                        var v75 float32 = 0.1
                        if v20 {
                            v76 := GenS1{F0: -9223372036854775808}
                            switch int(v5) {
                            case 4294967295:
                                if false {
                                    fmt.Println(v21)
                                } else {
                                    v26 = v20
                                    v26 = v20
                                    v26 = v26
                                }
                            case 1:
                                v3[0] = ""
                            case 32768:
                                v77 := GenS11{F0: int64(v1), F1: 100, F2: (v5 * 141)}
                                _ = v77
                            default:
                                v5 += v5
                            }
                            if true {
                                v20 = (v2 == v2)
                                v2 += v2
                                var v78 int = 128
                                _ = v78
                            } else {
                                v79 := GenS12{F0: -103, F1: 1}
                                v2 += (v2 - v2)
                                _ = v79
                            }
                            _ = v76
                        }
                        _ = v75
                    }
                    fmt.Println(v1, v26, v2)
                }
                v80 := GenS13{F0: 0, F1: v26}
                _ = fi25
                _ = v26
                _ = v80
            }
            _ = fi23
            _ = v24
        }
    }
    for fi81 := 0; fi81 < 14; fi81++ {
        switch "foo bar" {
        case "hello":
            if v20 {
                switch "a" {
                case "foo bar":
                    if v20 {
                        if (v1 > 42) {
                            var v82 string = ""
                            v83 := GenS14{F0: v20}
                            v21 += (v21 % (((v21 >> ((v5) & 63))) | 1))
                            _ = v82
                            _ = v83
                        } else {
                            for fi84 := 0; fi84 < 4; fi84++ {
                                var v85 uint64 = 18446744073709551615
                                var v86 int32 = v1
                                var v87 uint = 964
                                _ = fi84
                                _ = v85
                                _ = v86
                                _ = v87
                            }
                        }
                        v88 := GenS9{F0: fi81, F1: v1, F2: "foo bar"}
                        for fi89 := 0; fi89 < 7; fi89++ {
                            v6 = append(v6, fi89)
                            _ = fi89
                        }
                        _ = v88
                    } else {
                        v5 = uint(v2)
                        switch 9223372036854775807 {
                        case 42:
                            v2 = (v2 ^ v2)
                        case 728:
                            v20 = (0.0 != 0.1)
                        default:
                            switch (fi81 & -1) {
                            case 577:
                                if false {
                                    fmt.Println(v2)
                                } else {
                                    if false {
                                        _ = len(v6)
                                        v90 := GenS12{F0: -17, F1: 252}
                                        v3[218] = "\n\t"
                                        _ = v90
                                    } else {
                                        var v91 uint8 = 117
                                        v1 = v1
                                        _ = v91
                                    }
                                    fmt.Println(v5, v5, v5)
                                    v92 := GenS15{F0: 18446744073709551615, F1: "foo bar"}
                                    _ = v92
                                }
                            case 1:
                                if (("hello" + "a") >= (("\n\t" + "hello") + ("foo bar" + "\n\t"))) {
                                    var v93 int16 = v21
                                    var v94 int8 = 42
                                    _ = v93
                                    _ = v94
                                } else {
                                    for fi95 := 0; fi95 < 5; fi95++ {
                                        switch "\n\t" {
                                        case "hello":
                                            v1 += 42
                                        case "a":
                                            var v96 string = "foo bar"
                                            _ = v96
                                        default:
                                            var v97 int32 = 215
                                            _ = v97
                                        }
                                        if (!(!false)) {
                                            v98 := GenS9{F0: fi81, F1: -762, F2: "foo bar"}
                                            _ = v98
                                        } else {
                                            v99 := GenS16{F0: 42, F1: 65}
                                            switch fi81 {
                                            case -415:
                                                v2 += 2
                                            default:
                                                v100 := GenS16{F0: -9223372036854775808, F1: 128}
                                                _ = v100
                                            }
                                            _ = v99
                                        }
                                        var v101 float32 = 0.0
                                        _ = fi95
                                        _ = v101
                                    }
                                    if (!((42 == 0) || v20)) {
                                        if v20 {
                                            v5 -= (v5 % ((2) | 1))
                                            for fi102 := 0; fi102 < 8; fi102++ {
                                                v20 = (float32(((-2.25 * 1.5) - 1e10)) > 0.1)
                                                _ = fi102
                                            }
                                            fmt.Println(v21)
                                        }
                                        fmt.Println(v2, v2, v21)
                                    }
                                }
                            default:
                                v21 -= v21
                            }
                        }
                    }
                default:
                    v2 = (^(v2))
                }
                v20 = v20
                v103 := GenS17{F0: 42}
                _ = v103
            }
        case "a":
            var v104 uint8 = 162
            _ = v104
        default:
            v105 := GenS2{F0: (v2 >> ((v5) & 63))}
            _ = v105
        }
        _ = fi81
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
