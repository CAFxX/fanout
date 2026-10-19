package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int64; F2 int32 }
type GenS7 struct { F0 string; F1 uint32; F2 uint32 }
type GenS8 struct { F0 bool; F1 int8; F2 uint64 }
type GenS9 struct { F0 bool; F1 int8 }
type GenS10 struct { F0 int; F1 uint8; F2 int }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int8; F1 int16 }
func genFn0(p0 uint16, p1 uint64) uint64 {
    v4 := GenS1{F0: 2, F1: 4294967295, F2: "\n\t"}
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := GenS2{F0: 639}
        var v7 int = (fi5 & (fi5 % (((fi5 + (fi5 & fi5))) | 1)))
        switch "" {
        case "a":
            v8 := []int{1, 4294967295, (fi5 << ((65535) & 63))}
            _ = v8
        case "":
            for fi9 := 0; fi9 < 14; fi9++ {
                if (-128 < 127) {
                    v7 -= fi9
                }
                p0 = p0
                p1 = 4294967295
                if true {
                    if (-111 == 100) {
                        p0 = (p0 - 65535)
                        var v10 float32 = 0.1
                        if false {
                            v11 := GenS3{F0: ("\n\t" + "a"), F1: true, F2: -229}
                            var v12 int16 = -32768
                            var v13 int64 = -1
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        _ = v10
                    } else {
                        switch "hello" {
                        case "hello":
                            p1 += p1
                        case "foo bar":
                            v14 := []int{32767}
                            _ = v14
                        case "":
                            if false {
                                var v15 float64 = 0.1
                                p0 += 994
                                var v16 uint = 543
                                _ = v15
                                _ = v16
                            }
                        }
                        var v17 bool = (((3.14159 / 1e10) == 0.0) || false)
                        _ = v17
                    }
                    switch fi5 {
                    case -2147483648:
                        v18 := GenS4{F0: 118, F1: 4294967295}
                        _ = v18
                    case 89:
                        switch "hello" {
                        case "foo bar":
                            p0 -= 0
                        case "":
                            fmt.Println(fi5, fi9)
                        }
                    default:
                        v19 := GenS2{F0: (fi9 / ((int(p0)) | 1))}
                        _ = v19
                    }
                    var v20 uint64 = (p1 | p1)
                    fmt.Println(v7, v7)
                    fmt.Println(fi5, p1, p1)
                    _ = v20
                } else {
                    if true {
                        v21 := GenS4{F0: 2147483647, F1: 1}
                        p0 = p0
                        p1 = (p1 + p1)
                        _ = v21
                    }
                    fmt.Println(p1, fi9, fi9)
                    if false {
                        p1 -= p1
                        var v22 int64 = int64(p1)
                        p1 += p1
                        _ = v22
                    } else {
                        fmt.Println(p0, fi5, p1)
                        p1 += (p1 / ((p1) | 1))
                    }
                }
                _ = fi9
            }
        default:
            var v23 bool = ((0.0 / float32(3.14159)) > (1e10 + float32(3.14159)))
            _ = v23
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 % (((p1 | p1)) | 1))
}
func genFn1() int {
    var v24 int = 4294967295
    var v25 string = ""
    var v26 float64 = 0.1
    _ = v24
    _ = v25
    _ = v26
    return 0
}

func main() {
    var v1 string = "a"
    var v2 uint = 719
    v3 := []string{(((v1 + v1) + v1) + v1), v1, v1, "\n\t"}
    var v27 string = v1
    var v28 uint64 = 42
    fmt.Println(v2, v1)
    if false {
        for fi29 := 0; fi29 < 9; fi29++ {
            if true {
                for fi30 := 0; fi30 < 14; fi30++ {
                    var v31 string = v1
                    _ = fi30
                    _ = v31
                }
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                v2 += (v2 ^ v2)
            } else {
                var v32 uint32 = 2147483647
                var v33 bool = (-1 <= 9223372036854775807)
                if (v28 > (^(v28))) {
                    var v34 int32 = -986
                    var v35 uint = (v2 ^ v2)
                    v27 = v27
                    _ = v34
                    _ = v35
                }
                _ = v32
                _ = v33
            }
            if true {
                v36 := GenS2{F0: genFn1()}
                if true {
                    for fi37 := 0; fi37 < 7; fi37++ {
                        v1 = "foo bar"
                        if (2147483647 >= 1) {
                            if false {
                                var v38 uint16 = 690
                                var v39 int = 32768
                                var v40 int16 = -1
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            v28 -= v28
                            v41 := GenS1{F0: v2, F1: (fi37 << ((v2) & 63)), F2: ""}
                            _ = v41
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                if (127 == 100) {
                                    v42 := GenS4{F0: 1, F1: 21}
                                    _ = v42
                                } else {
                                    v43 := GenS5{F0: 32767}
                                    v44 := GenS6{F0: 4294967295, F1: -1, F2: -1}
                                    _ = v43
                                    _ = v44
                                }
                            } else {
                                v28 = v28
                                if false {
                                    fmt.Println(fi29)
                                } else {
                                    v28 = 42
                                    v27 = v1
                                }
                                v28 += genFn0(336, (v28 << (((v2 | (v2 - 4294967295))) & 63)))
                            }
                            if false {
                                v2 = (^(v2))
                                var v45 int16 = 345
                                _ = v45
                            } else {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v46 uint8 = 255
                                    if true {
                                        var v47 bool = (!((-2.25 - 3.14159) == (3.14159 - 0.0)))
                                        v48 := GenS7{F0: v1, F1: 529, F2: 4294967295}
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if false {
                                            var v49 uint32 = 96
                                            _ = v49
                                        } else {
                                            v28 = 339
                                            var v50 int16 = -471
                                            switch (fi37 | fi37) {
                                            case -481:
                                                v46 = uint8(v2)
                                            case 127:
                                                v27 = v27
                                                if (558 != 9223372036854775807) {
                                                    var v51 int32 = 2147483647
                                                    _ = v51
                                                }
                                            default:
                                                v28 -= v28
                                            }
                                            _ = v50
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = v27
                                        }
                                    }
                                    _ = v46
                                } else {
                                    if false {
                                        switch fi37 {
                                        case 728:
                                            v52 := GenS8{F0: (false || (0 < 666)), F1: 127, F2: v28}
                                            _ = v52
                                        default:
                                            switch (fi37 >> ((v2) & 63)) {
                                            case 4294967295:
                                                var v53 int64 = 9223372036854775807
                                                _ = v53
                                            default:
                                                v1 = "foo bar"
                                            }
                                            var v54 float32 = 0.0
                                            v27 = ""
                                            _ = v54
                                        }
                                        if (-2.25 == 3.14159) {
                                            switch "a" {
                                            case "":
                                                switch fi37 {
                                                case -684:
                                                    var v55 int32 = 2147483647
                                                    _ = v55
                                                case -1:
                                                    var v56 int32 = -974
                                                    _ = v56
                                                }
                                            }
                                        }
                                        if true {
                                            v28 -= (v28 & (v28 - v28))
                                        }
                                        var v57 uint64 = v28
                                        _ = v57
                                    } else {
                                        v2 -= 2
                                        v2 = v2
                                        if ((true || true) && true) {
                                            v58 := GenS4{F0: 820, F1: 4294967295}
                                            var v59 int64 = -851
                                            _ = v58
                                            _ = v59
                                        } else {
                                            if true {
                                                if (v2 <= (^(v2))) {
                                                    v3 = append(v3, (v27 + (v27 + "hello")))
                                                    v27 = "hello"
                                                } else {
                                                    var v60 int32 = 0
                                                    if (uint8(v2) == 132) {
                                                        v61 := GenS3{F0: v27, F1: (!(!((true || false) || (-81 > 0)))), F2: 557}
                                                        var v62 string = "foo bar"
                                                        var v63 uint8 = 128
                                                        _ = v61
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        var v64 int = (fi37 % (((fi29 ^ (-(fi29)))) | 1))
                                                        var v65 float32 = (1.5 - 1e10)
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    var v66 uint = v2
                                                    var v67 int8 = -128
                                                    v2 += (v66 + v2)
                                                    _ = v60
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v68 := GenS9{F0: true, F1: 52}
                                                v69 := GenS1{F0: v2, F1: 128, F2: ""}
                                                _ = v68
                                                _ = v69
                                            }
                                        }
                                    }
                                    v3 = append(v3, v1)
                                }
                                v28 = 4294967295
                                var v70 int32 = -104
                                _ = v70
                            }
                        }
                        _ = fi37
                    }
                    v71 := GenS10{F0: (fi29 * fi29), F1: 2, F2: fi29}
                    _ = v71
                } else {
                    fmt.Println(v28, fi29, v27)
                }
                v27 = (v27 + ("hello" + v1))
                _ = v36
            }
            v1 = v27
            _ = fi29
        }
        var v72 int8 = 22
        if false {
            v3 = append(v3, v27)
        }
        _ = v72
    } else {
        var v73 int = genFn1()
        var v74 int = (v73 % ((2147483647) | 1))
        for fi75 := 0; fi75 < 4; fi75++ {
            if (255 < 42) {
                var v76 int8 = 0
                for fi77 := 0; fi77 < 19; fi77++ {
                    fmt.Println(v28, v73)
                    var v78 int16 = 1
                    _ = fi77
                    _ = v78
                }
                for fi79 := 0; fi79 < 10; fi79++ {
                    switch v74 {
                    case 32767:
                        fmt.Println(fi79, v28)
                    case 127:
                        v74 -= (fi75 % ((genFn1()) | 1))
                    case -1:
                        var v80 string = (v27 + v1)
                        _ = v80
                    default:
                        switch "foo bar" {
                        case "a":
                            if true {
                                v81 := GenS5{F0: uint16(v2)}
                                if false {
                                    v2 += 52
                                    if true {
                                        v74 += 0
                                    }
                                }
                                fmt.Println(v27, v2)
                                _ = v81
                            } else {
                                if ("a" == "") {
                                    var v82 uint = v2
                                    _ = v82
                                } else {
                                    v2 += v2
                                }
                                var v83 int8 = v76
                                _ = v83
                            }
                        case "\n\t":
                            if true {
                                v73 -= 642
                                var v84 int16 = 1000
                                fmt.Println(v73, fi75, v73)
                                _ = v84
                            } else {
                                v85 := GenS11{F0: 42}
                                _ = v85
                            }
                        case "hello":
                            switch (v74 << (((v2 >> ((v2) & 63))) & 63)) {
                            case 9223372036854775807:
                                fmt.Println(v28, v76, v74)
                            default:
                                if true {
                                    v86 := GenS1{F0: (v2 - (v2 | 0)), F1: (fi79 + (^(v73))), F2: v27}
                                    _ = v86
                                }
                            }
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (((169 > 128) || (true || (223 >= v28))) || false) {
                        v76 = (^(v76))
                        v3 = append(v3, v1)
                    }
                    _ = fi79
                }
                _ = v76
            }
            _ = fi75
        }
        _ = v73
        _ = v74
    }
    fmt.Println(v28, v1, v27)
    if true {
        if (127 != 2) {
            var v87 uint16 = 65535
            v88 := GenS6{F0: 2147483648, F1: -288, F2: 42}
            _ = v87
            _ = v88
        } else {
            var v89 uint64 = (^(v28))
            if len(v3) > 0 {
                v3[0] = "a"
            }
            _ = v89
        }
    }
    v1 = (v27 + (v1 + "a"))
    var v90 int8 = 10
    if ((0 != (v90 * (v90 - v90))) || (-2147483648 < 301)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi91 := 0; fi91 < 16; fi91++ {
            var v92 bool = (true && (!(1 < 1)))
            var v93 string = v1
            for fi94 := 0; fi94 < 14; fi94++ {
                if (v1 == "hello") {
                    if len(v3) > 0 {
                        v3[0] = (v27 + v1)
                    }
                    fmt.Println(v92, v27)
                } else {
                    switch (-(fi91)) {
                    case 0:
                        if (v92 || true) {
                            var v95 int = (fi94 * genFn1())
                            v90 += v90
                            _ = v95
                        } else {
                            v90 += (v90 >> ((v2) & 63))
                            var v96 float32 = ((float32(float64(((3.14159 / 3.14159) + 3.14159))) / 1e10) / -2.25)
                            v93 = (("foo bar" + (v1 + "foo bar")) + (("hello" + "foo bar") + (((v27 + (((v1 + (v1 + v27)) + v93) + ("a" + (((v93 + v27) + "a") + ("a" + ((v93 + v93) + "foo bar")))))) + (v1 + v1)) + ((v1 + "") + ("" + "")))))
                            _ = v96
                        }
                    case -2147483648:
                        if (1000 < 416) {
                            v27 = "a"
                        }
                    default:
                        v92 = ((("\n\t" + "foo bar") + (v27 + v93)) == v93)
                    }
                    v90 = v90
                    if ((^(fi94)) != int(v2)) {
                        v97 := GenS12{F0: (v90 + v90), F1: 27}
                        _ = v97
                    }
                }
                var v98 int = fi91
                v90 = (v90 - (v90 + v90))
                _ = fi94
                _ = v98
            }
            _ = fi91
            _ = v92
            _ = v93
        }
    }
    switch -1 {
    case -2147483648:
        var v99 int64 = 0
        _ = v99
    case 4294967295:
        v2 -= 104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v90
}
