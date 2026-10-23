package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 int32; F1 int32; F2 int16 }
type GenS4 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 uint8; F1 uint }
type GenS13 struct { F0 int8; F1 int }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int16 }
func genFn0() uint64 {
    v6 := []int{-9223372036854775808, -9223372036854775808}
    var v7 int32 = -890
    fmt.Println(v7)
    if false {
        fmt.Println(v7)
        v7 -= 0
        v7 = (v7 - -1)
    } else {
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                v7 = (-(v7))
            } else {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 string = ""
                    var v11 int = (fi9 & 1)
                    _ = len(v6)
                    var v12 int = (fi8 << ((2) & 63))
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi8
        }
        v7 = -741
        if (true || true) {
            if true {
                if true {
                    if (42 == 0) {
                        v7 = (v7 & v7)
                        var v13 string = "a"
                        _ = v13
                    } else {
                        v7 = v7
                    }
                    var v14 uint32 = 1
                    v15 := GenS1{F0: 253, F1: v7}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int32 = v7
                    for fi17 := 0; fi17 < 16; fi17++ {
                        v16 += (v16 << ((1) & 63))
                        var v18 int16 = 42
                        v18 = 1
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                var v19 int16 = -32768
                var v20 uint = 593
                _ = v19
                _ = v20
            }
        }
    }
    _ = v6
    _ = v7
    return 796
}

func main() {
    var v1 int64 = 71
    v2 := []int{4294967295}
    var v3 uint8 = 93
    var v4 int = -827
    v5 := []string{("" + (("foo bar" + ((("foo bar" + "") + "") + "hello")) + "hello")), ""}
    v3 -= 255
    switch ("foo bar" + "") {
    case "hello":
        var v21 bool = true
        _ = v21
    case "a":
        v1 -= v1
    default:
        v22 := GenS2{F0: genFn0(), F1: 9223372036854775807, F2: true}
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        for fi24 := 0; fi24 < 10; fi24++ {
            switch (v4 | (fi24 << ((255) & 63))) {
            case -948:
                v25 := GenS1{F0: 200, F1: 42}
                _ = v25
            case 128:
                v26 := GenS1{F0: v3, F1: 42}
                v27 := GenS3{F0: 1, F1: 42, F2: 32767}
                if false {
                    v28 := GenS2{F0: 0, F1: v4, F2: false}
                    _ = v28
                } else {
                    v3 -= v3
                    var v29 int64 = (v1 / ((v1) | 1))
                    if (!true) {
                        v3 -= (^(v3))
                        switch "a" {
                        case "hello":
                            fmt.Println(fi23)
                        case "":
                            var v30 string = "hello"
                            _ = v30
                        case "\n\t":
                            fmt.Println(v3, v1)
                        default:
                            v5 = append(v5, "\n\t")
                        }
                        _ = len(v5)
                    }
                    switch ("\n\t" + "hello") {
                    case "a":
                        var v31 float64 = ((float64(-2.25) - (3.14159 - float64(float32((-2.25 / 1.5))))) / -2.25)
                        _ = v31
                    case "\n\t":
                        v32 := GenS3{F0: 339, F1: 42, F2: -1}
                        _ = v32
                    case "":
                        var v33 int64 = v1
                        _ = v33
                    }
                    fmt.Println(v4, v3)
                    _ = v29
                }
                var v34 uint8 = v3
                _ = v26
                _ = v27
                _ = v34
            default:
                if false {
                    if (((((0.1 / -2.25) - -2.25) * ((1.5 + -2.25) + 0.1)) * (1.5 + (-2.25 - 0.0))) <= float64((1e10 * 0.1))) {
                        var v35 string = ""
                        v36 := GenS1{F0: 127, F1: 1}
                        v37 := GenS4{F0: 0, F1: 1000, F2: 524}
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    switch fi23 {
                    case 0:
                        var v38 bool = false
                        _ = v38
                    case -140:
                        v39 := GenS5{F0: -221, F1: (true || false)}
                        _ = v39
                    case -2147483648:
                        fmt.Println(v1, fi24, v1)
                    default:
                        if false {
                            var v40 bool = false
                            if len(v2) > 0 {
                                v2[0] = fi24
                            }
                            v41 := GenS6{F0: v1}
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                            _ = v40
                            _ = v41
                        }
                    }
                    if (1000 < 0) {
                        v3 -= (^(v3))
                        v4 += fi24
                    }
                }
                if false {
                    fmt.Println(v1)
                    switch (v4 >> ((4294967295) & 63)) {
                    case 127:
                        var v42 uint16 = 2
                        _ = v42
                    case 2147483647:
                        var v43 uint16 = 32767
                        _ = v43
                    default:
                        v4 = 128
                    }
                    var v44 float64 = -2.25
                    _ = v44
                } else {
                    var v45 int64 = (v1 ^ (v1 * v1))
                    v1 = v45
                    if true {
                        fmt.Println(v4)
                    } else {
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        if true {
                            var v46 float64 = ((3.14159 + 0.1) + 1.5)
                            _ = v46
                        } else {
                            var v47 int = (fi23 / ((v4) | 1))
                            _ = v47
                        }
                        if (0 < 0) {
                            _ = len(v5)
                            v3 -= v3
                        } else {
                            if false {
                                v48 := GenS7{F0: (fi23 + fi23)}
                                _ = v48
                            } else {
                                if false {
                                    fmt.Println(v45)
                                }
                                v4 -= (fi24 & v4)
                            }
                            var v49 int32 = -401
                            if (!false) {
                                switch fi23 {
                                case 32767:
                                    if true {
                                        if true {
                                            fmt.Println(v3)
                                            var v50 float64 = 0.1
                                            v45 = v1
                                            _ = v50
                                        }
                                        if true {
                                            fmt.Println(v1, v45)
                                        }
                                        switch (v4 - v4) {
                                        case 32768:
                                            var v51 uint = 255
                                            _ = v51
                                        }
                                    } else {
                                        var v52 bool = false
                                        v52 = (255 >= uint(v3))
                                        _ = v52
                                    }
                                }
                                v2 = append(v2, (v4 & 9223372036854775807))
                                v1 = 263
                            }
                            v4 = 32768
                            var v53 float32 = 1e10
                            _ = v49
                            _ = v53
                        }
                    }
                    if (-2.25 < 1e10) {
                        switch "a" {
                        case "foo bar":
                            switch "\n\t" {
                            case "":
                                v54 := GenS8{F0: 1, F1: 18446744073709551615}
                                _ = v54
                            default:
                                _ = len(v5)
                            }
                        case "a":
                            v3 = v3
                        case "":
                            v2 = append(v2, -1)
                        default:
                            fmt.Println(v45, v1)
                        }
                        fmt.Println(v45, fi24, v45)
                        v1 -= (v1 << ((394) & 63))
                    } else {
                        v45 = (v1 << ((0) & 63))
                        if true {
                            switch (^(fi23)) {
                            case 127:
                                v55 := GenS9{F0: 0}
                                _ = v55
                            default:
                                var v56 float64 = 0.1
                                _ = v56
                            }
                        } else {
                            if false {
                                v57 := GenS10{F0: 93}
                                v3 -= v3
                                v5 = append(v5, "")
                                fmt.Println(fi23, v45, fi24)
                                _ = v57
                            } else {
                                var v58 float64 = (((0.1 - ((-2.25 / (3.14159 + 3.14159)) + float64(((((-2.25 - -2.25) * 3.14159) - 0.0) * 0.1)))) / 1.5) * float64(3.14159))
                                _ = v58
                            }
                            switch -2147483648 {
                            case 32767:
                                v59 := GenS2{F0: 1, F1: 9223372036854775807, F2: false}
                                _ = v59
                            default:
                                if true {
                                    if false {
                                        v45 -= 9223372036854775807
                                        v60 := GenS7{F0: fi24}
                                        _ = len(v5)
                                        _ = v60
                                    } else {
                                        fmt.Println(v45, fi23, fi23)
                                        var v61 float32 = 1e10
                                        var v62 int32 = -858
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    v1 = (v45 | -56)
                                }
                            }
                        }
                    }
                    fmt.Println(fi23, fi24, fi24)
                    _ = v45
                }
            }
            switch fi23 {
            case 313:
                fmt.Println(v4)
            case 430:
                fmt.Println(v4, v4)
            case 128:
                var v63 uint8 = (v3 ^ 2)
                var v64 float64 = 0.1
                var v65 uint64 = 0
                _ = v63
                _ = v64
                _ = v65
            }
            var v66 int16 = 0
            if true {
                fmt.Println(v3, v1, v66)
            }
            _ = fi24
            _ = v66
        }
        v1 = 1
        _ = fi23
    }
    if ((((true && true) || true) && false) || ("hello" == "foo bar")) {
        if true {
            var v67 int32 = 1
            if true {
                var v68 uint = 42
                fmt.Println(v4, v1)
                if (!(v67 == v67)) {
                    for fi69 := 0; fi69 < 19; fi69++ {
                        fmt.Println(v3)
                        _ = fi69
                    }
                }
                _ = v68
            }
            _ = v67
        } else {
            if ((!(false || (!true))) || false) {
                if (true || ((2 < 32767) && false)) {
                    var v70 string = ("\n\t" + "a")
                    _ = v70
                } else {
                    var v71 int8 = 0
                    var v72 uint64 = 104
                    _ = v71
                    _ = v72
                }
            }
        }
        if true {
            for fi73 := 0; fi73 < 4; fi73++ {
                switch (v4 % (((^(fi73))) | 1)) {
                case 0:
                    var v74 int32 = 42
                    _ = v74
                default:
                    if false {
                        for fi75 := 0; fi75 < 16; fi75++ {
                            var v76 uint32 = 512
                            var v77 uint64 = 0
                            var v78 float32 = (-2.25 / float32(1.5))
                            _ = fi75
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                v79 := GenS11{F0: 255, F1: 42}
                _ = fi73
                _ = v79
            }
            var v80 uint8 = v3
            _ = v80
        } else {
            var v81 uint64 = uint64(v3)
            if true {
                v82 := GenS1{F0: v3, F1: 0}
                var v83 bool = (true && (42 < 4294967295))
                _ = v82
                _ = v83
            } else {
                switch -1 {
                case -2147483648:
                    if (v3 != v3) {
                        for fi84 := 0; fi84 < 9; fi84++ {
                            var v85 int8 = 105
                            _ = fi84
                            _ = v85
                        }
                        v81 = (^(v81))
                        switch "a" {
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = v4
                            }
                        case "\n\t":
                            v4 += v4
                        }
                    }
                case -78:
                    if true {
                        fmt.Println(v3, v3)
                        v3 -= (v3 & v3)
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                    } else {
                        v86 := GenS4{F0: 65535, F1: -32768, F2: 2}
                        _ = v86
                    }
                }
                v2 = append(v2, (v4 << ((994) & 63)))
            }
            _ = v81
        }
    }
    var v87 int8 = -60
    switch "a" {
    case "foo bar":
        v3 -= v3
    }
    fmt.Println(v3)
    for fi88 := 0; fi88 < 4; fi88++ {
        var v89 uint16 = 32767
        fmt.Println(v3)
        fmt.Println(fi88, v1, v3)
        _ = fi88
        _ = v89
    }
    var v90 uint = 255
    if false {
        v4 -= (^(v4))
    } else {
        switch (("" + "") + ("foo bar" + "")) {
        case "\n\t":
            switch int(v1) {
            case 840:
                if (994 > 606) {
                    if (v1 == 42) {
                        if len(v2) > 0 {
                            v2[0] = v4
                        }
                    } else {
                        if (((533 > 365) || ((-1 == -58) || (32767 > 0))) || (0.0 == (1.5 / 3.14159))) {
                            v91 := GenS11{F0: v3, F1: (v4 % ((v4) | 1))}
                            var v92 int32 = -2147483648
                            _ = v91
                            _ = v92
                        } else {
                            switch (v4 >> ((v90) & 63)) {
                            case -2147483648:
                                if true {
                                    var v93 uint8 = (v3 + (v3 * (v3 % ((v3) | 1))))
                                    for fi94 := 0; fi94 < 19; fi94++ {
                                        v95 := GenS2{F0: 42, F1: (fi94 - 0), F2: (-2147483648 <= -373)}
                                        var v96 float32 = (((float32(((-2.25 + (1e10 / 1.0)) * 0.0)) - ((-2.25 + -2.25) / 1.5)) - (((0.1 / 1.0) * 0.1) + 0.1)) + 3.14159)
                                        _ = fi94
                                        _ = v95
                                        _ = v96
                                    }
                                    var v97 int16 = 1
                                    _ = v93
                                    _ = v97
                                } else {
                                    v1 = v1
                                    var v98 bool = false
                                    fmt.Println(v4, v1)
                                    _ = v98
                                }
                            case -9223372036854775808:
                                v4 = v4
                            case 415:
                                var v99 float32 = -2.25
                                _ = v99
                            default:
                                var v100 uint8 = v3
                                _ = v100
                            }
                            var v101 int64 = (v1 >> (((v90 | v90)) & 63))
                            _ = v101
                        }
                    }
                }
            case 849:
                var v102 bool = (false && false)
                _ = v102
            default:
                switch v4 {
                case 2147483647:
                    var v103 string = ""
                    _ = v103
                case -360:
                    v1 -= (v1 + v1)
                }
            }
        case "a":
            if (false || false) {
                v104 := GenS12{F0: 1, F1: v90}
                var v105 string = "a"
                v106 := GenS13{F0: (-(v87)), F1: 0}
                _ = v104
                _ = v105
                _ = v106
            }
        case "":
            if true {
                var v107 int32 = -2147483648
                v90 = v90
                _ = v107
            } else {
                var v108 float64 = (-2.25 + 3.14159)
                _ = v108
            }
        default:
            v109 := GenS14{F0: 1}
            _ = v109
        }
    }
    for fi110 := 0; fi110 < 8; fi110++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi110
    }
    v111 := GenS15{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v87:", v87)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
    _ = v90
}
