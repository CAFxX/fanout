package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 string }
type GenS2 struct { F0 int16; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8; F1 uint; F2 uint64 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 int16; F1 int64 }
func genFn0(p0 uint8) uint64 {
    switch ("hello" + "") {
    case "foo bar":
        switch ("" + "a") {
        case "\n\t":
            fmt.Println(p0)
        case "hello":
            p0 += p0
        }
    case "hello":
        v6 := GenS1{F0: 115, F1: 0, F2: ("hello" + ("" + ""))}
        _ = v6
    case "\n\t":
        fmt.Println(p0)
    }
    if true {
        if ((!(2147483648 == 88)) && ((!false) && false)) {
            var v7 int16 = 32767
            p0 += p0
            _ = v7
        } else {
            var v8 int64 = 0
            if false {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v8 -= (^(v8))
                    var v10 int16 = -1
                    var v11 bool = false
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 8; fi12++ {
                switch "" {
                case "":
                    fmt.Println(fi12)
                }
                _ = fi12
            }
            _ = v8
        }
        fmt.Println(p0)
        v13 := []int{-36}
        _ = v13
    }
    var v14 uint16 = 255
    v15 := GenS1{F0: 9223372036854775807, F1: -128, F2: (("foo bar" + "a") + "a")}
    _ = p0
    _ = v14
    _ = v15
    return 42
}
func genFn1(p0 uint8, p1 int) int32 {
    v16 := []int{p1}
    v16 = append(v16, p1)
    for fi17 := 0; fi17 < 15; fi17++ {
        p1 = p1
        v18 := GenS1{F0: 0, F1: 0, F2: ((("\n\t" + ((("\n\t" + ("hello" + ((("a" + "hello") + "hello") + ("\n\t" + "a")))) + "foo bar") + "hello")) + ("\n\t" + "\n\t")) + ("\n\t" + "foo bar"))}
        v19 := GenS1{F0: 42, F1: 0, F2: "foo bar"}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v16
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 352
    var v3 float32 = -2.25
    v4 := make(map[int]uint32)
    var v5 uint = 1
    v2 = v2
    if (-1 <= 0) {
        if ((v3 + 1e10) != (v3 / ((v3 / v3) * (v3 + 3.14159)))) {
            if (0 <= 32767) {
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch 4294967295 {
                    case 690:
                        v21 := []int{1, fi20, 4294967295}
                        _ = v21
                    default:
                        v22 := []int{fi20, (fi20 | -473), (fi20 ^ (-(fi20)))}
                        _ = v22
                    }
                    var v23 int16 = -32768
                    var v24 float64 = -2.25
                    _ = fi20
                    _ = v23
                    _ = v24
                }
                switch "\n\t" {
                case "a":
                    v25 := []int{4294967295}
                    _ = v25
                default:
                    v3 = v3
                }
            }
        } else {
            if (true || true) {
                fmt.Println(v5, v3, v3)
                if ((!false) && true) {
                    if (!((("\n\t" + v1) <= v1) || (-2147483648 <= genFn1(128, 2147483647)))) {
                        if true {
                            var v26 int32 = 2147483647
                            var v27 float32 = (0.1 / 3.14159)
                            _ = v26
                            _ = v27
                        } else {
                            v28 := GenS2{F0: -749, F1: 216, F2: 42}
                            switch v1 {
                            case "a":
                                switch v1 {
                                case "foo bar":
                                    v29 := GenS3{F0: v5, F1: (^(v5)), F2: -9223372036854775808}
                                    _ = v29
                                case "hello":
                                    for fi30 := 0; fi30 < 20; fi30++ {
                                        for fi31 := 0; fi31 < 9; fi31++ {
                                            v32 := []int{(fi31 & fi30)}
                                            v32 = append(v32, (fi31 | (fi30 * fi31)))
                                            _ = fi31
                                            _ = v32
                                        }
                                        _ = fi30
                                    }
                                }
                            case "hello":
                                if (-623 != -1) {
                                    v2 -= v2
                                } else {
                                    v5 = v5
                                }
                            case "foo bar":
                                v33 := []int{-676, 906, 42, 32768}
                                _ = v33
                            }
                            _ = v28
                        }
                    }
                    v3 -= v3
                    v34 := GenS2{F0: 1, F1: 158, F2: 2147483647}
                    _ = v34
                } else {
                    v4[127] = v2
                    v2 = (v2 - (v2 / ((v2) | 1)))
                }
            } else {
                switch -833 {
                case 413:
                    var v35 int32 = -603
                    _ = v35
                }
                switch (v1 + v1) {
                case "a":
                    v36 := GenS1{F0: int64(v2), F1: 116, F2: "a"}
                    _ = v36
                default:
                    v37 := GenS1{F0: 268, F1: -128, F2: ("" + v1)}
                    _ = v37
                }
            }
        }
    } else {
        for fi38 := 0; fi38 < 9; fi38++ {
            v39 := GenS4{F0: ("" + v1)}
            var v40 float64 = (1.5 - 0.0)
            _ = fi38
            _ = v39
            _ = v40
        }
        v5 = 4294967295
        if ((((1 > 128) || (!(((v5 | v5) <= 18446744073709551615) || (0 <= 0)))) && (!false)) || (106 == 2)) {
            for fi41 := 0; fi41 < 5; fi41++ {
                switch (fi41 - -2147483648) {
                case 32768:
                    var v42 uint16 = 1
                    _ = v42
                case -2147483648:
                    v2 = v2
                }
                var v43 int8 = 0
                v44 := GenS5{F0: 65535, F1: v2}
                _ = fi41
                _ = v43
                _ = v44
            }
        }
    }
    switch 0 {
    case -263:
        if (174 < -350) {
            var v45 uint16 = 32767
            for fi46 := 0; fi46 < 7; fi46++ {
                switch v1 {
                case "hello":
                    var v47 float64 = (1.5 + (1.5 / (-2.25 / 1.0)))
                    _ = v47
                }
                v48 := GenS6{F0: (fi46 * 32767)}
                v49 := GenS3{F0: v5, F1: 1, F2: -1}
                _ = fi46
                _ = v48
                _ = v49
            }
            v2 += v2
            _ = v45
        }
    }
    fmt.Println(v5, v2, v1)
    switch v1 {
    case "a":
        if true {
            fmt.Println(v5, v1)
        }
    }
    v5 = (v5 << ((670) & 63))
    switch v1 {
    case "":
        v5 -= v5
    case "hello":
        var v50 string = (v1 + "\n\t")
        _ = v50
    case "foo bar":
        v2 -= 224
    default:
        for fi51 := 0; fi51 < 8; fi51++ {
            var v52 int = (fi51 >> ((v5) & 63))
            v53 := GenS7{F0: (^(v2))}
            if true {
                var v54 uint = v5
                switch v1 {
                case "":
                    fmt.Println(v3, v3, v5)
                case "a":
                    var v55 float32 = 1.5
                    _ = v55
                case "hello":
                    if ((v1 + v1) < "a") {
                        var v56 uint16 = 1
                        v4[(-(v52))] = uint32(v54)
                        if true {
                            if true {
                                v57 := []int{fi51, (v52 % ((fi51) | 1))}
                                if false {
                                    var v58 float64 = float64(0.1)
                                    var v59 string = (v1 + "hello")
                                    _ = v58
                                    _ = v59
                                }
                                _ = v57
                            } else {
                                for fi60 := 0; fi60 < 15; fi60++ {
                                    fmt.Println(v54)
                                    v56 = v56
                                    if (((-2.25 + (0.1 / 1e10)) >= 3.14159) || true) {
                                        v61 := v4[(fi60 % ((fi51) | 1))]
                                        fmt.Println(v5, v61)
                                        _ = v61
                                    }
                                    _ = fi60
                                }
                                fmt.Println(v52, fi51, fi51)
                                v62 := GenS8{F0: 42, F1: 18446744073709551615, F2: 18446744073709551615}
                                _ = v62
                            }
                            if (!false) {
                                if (true && true) {
                                    delete(v4, (v52 - 32767))
                                    switch -1 {
                                    case 32767:
                                        v3 = 1e10
                                    case 128:
                                        v63 := []int{695}
                                        _ = v63
                                    default:
                                        for fi64 := 0; fi64 < 17; fi64++ {
                                            v3 -= v3
                                            v65 := GenS9{F0: 0, F1: (!(396 > int64(v52)))}
                                            switch 4294967295 {
                                            case 212:
                                                v52 = fi51
                                            case -2147483648:
                                                v66 := GenS6{F0: fi51}
                                                _ = v66
                                            }
                                            _ = fi64
                                            _ = v65
                                        }
                                    }
                                } else {
                                    switch v1 {
                                    case "":
                                        var v67 int8 = 42
                                        _ = v67
                                    case "hello":
                                        v2 = (v2 * v2)
                                    case "a":
                                        fmt.Println(v5, v3)
                                    }
                                }
                                var v68 int64 = -9223372036854775808
                                if false {
                                    var v69 string = (v1 + (v1 + v1))
                                    for fi70 := 0; fi70 < 1; fi70++ {
                                        switch v69 {
                                        case "\n\t":
                                            switch (fi70 << ((v5) & 63)) {
                                            case 842:
                                                var v71 int16 = 719
                                                _ = v71
                                            case -677:
                                                v56 = (^(v56))
                                            case 42:
                                                if true {
                                                    v72 := []int{fi70, (fi70 ^ (v52 | (-(fi70)))), fi70}
                                                    if (3.14159 == ((3.14159 + 1e10) * (0.0 + 1.5))) {
                                                        var v73 float32 = (v3 * (0.1 * v3))
                                                        switch fi51 {
                                                        case 597:
                                                            if len(v72) >= 2 {
                                                                v72 = v72[1:len(v72)]
                                                            }
                                                        case -1:
                                                            var v74 int = fi70
                                                            _ = v74
                                                        }
                                                        _ = v73
                                                    }
                                                    _ = v72
                                                } else {
                                                    var v75 int8 = 1
                                                    v2 = v2
                                                    _ = v75
                                                }
                                            default:
                                                switch fi51 {
                                                case -1:
                                                    v76 := GenS8{F0: 1, F1: v5, F2: 18446744073709551615}
                                                    _ = v76
                                                }
                                            }
                                        case "hello":
                                            if (677 < genFn0(249)) {
                                                var v77 int8 = 127
                                                v78 := GenS2{F0: -32768, F1: 42, F2: 2147483647}
                                                if false {
                                                    v68 += -1
                                                    v2 -= v2
                                                } else {
                                                    if (!true) {
                                                        v3 += 0.0
                                                        v52 -= 42
                                                    }
                                                    var v79 uint = (^(v5))
                                                    _ = v79
                                                }
                                                _ = v77
                                                _ = v78
                                            } else {
                                                switch v1 {
                                                case "\n\t":
                                                    switch (^(fi70)) {
                                                    case 2147483647:
                                                        var v80 uint16 = (v56 ^ (v56 - uint16(v2)))
                                                        _ = v80
                                                    }
                                                case "foo bar":
                                                    v81 := GenS9{F0: uint64(v54), F1: true}
                                                    _ = v81
                                                case "a":
                                                    if (46 != 1) {
                                                        var v82 uint32 = v2
                                                        v52 = (fi70 << ((v54) & 63))
                                                        var v83 string = (("foo bar" + "foo bar") + v1)
                                                        var v84 int = -192
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                default:
                                                    switch fi51 {
                                                    case 754:
                                                        switch v52 {
                                                        case -994:
                                                            v3 = v3
                                                        case 131:
                                                            var v85 uint16 = v56
                                                            _ = v85
                                                        }
                                                    case 127:
                                                        v2 -= v2
                                                    case 472:
                                                        var v86 int64 = v68
                                                        _ = v86
                                                    default:
                                                        if (false && false) {
                                                            var v87 uint8 = 255
                                                            fmt.Println(v2)
                                                            v54 = (v54 + (^(v5)))
                                                            _ = v87
                                                        } else {
                                                            if (128 <= 42) {
                                                                fmt.Println(fi51, v68, v2)
                                                                if true {
                                                                    var v88 int64 = (v68 + (v68 >> ((1) & 63)))
                                                                    _ = v88
                                                                } else {
                                                                    var v89 int64 = v68
                                                                    _ = v89
                                                                }
                                                                if (((v3 / 1.5) > (((v3 + v3) + float32(float64(float32((((1e10 / 1e10) * 1e10) * 0.0))))) / 1e10)) || false) {
                                                                    fmt.Println(fi70, fi70, v5)
                                                                } else {
                                                                    v69 = ((("\n\t" + v1) + (v69 + v1)) + "a")
                                                                    var v90 bool = ((v3 * ((v3 + v3) * (v3 / v3))) != v3)
                                                                    fmt.Println(v1, v3, v1)
                                                                    _ = v90
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        case "":
                                            v4[(fi70 ^ fi51)] = 65535
                                        default:
                                            switch v69 {
                                            case "a":
                                                v69 = "hello"
                                            }
                                        }
                                        fmt.Println(v3, v56)
                                        var v91 float32 = 1.5
                                        _ = fi70
                                        _ = v91
                                    }
                                    _ = v69
                                } else {
                                    v92 := GenS10{F0: 1, F1: v68}
                                    _ = v92
                                }
                                _ = v68
                            }
                        }
                        _ = v56
                    }
                default:
                    switch (v52 * (fi51 ^ (v52 << ((42) & 63)))) {
                    case -9223372036854775808:
                        v93 := []int{127, fi51, 487}
                        _ = v93
                    default:
                        v94 := GenS2{F0: -402, F1: 47, F2: 380}
                        _ = v94
                    }
                }
                fmt.Println(v5)
                _ = v54
            }
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    v1 = v1
    v5 -= v5
    var v95 uint64 = uint64(v2)
    for fi96 := 0; fi96 < 14; fi96++ {
        v3 = (-2.25 / v3)
        v2 += v2
        var v97 uint32 = 42
        _ = fi96
        _ = v97
    }
    var v98 uint32 = v2
    delete(v4, 128)
    v2 += (v98 >> ((42) & 63))
    v2 -= (v98 / ((v98) | 1))
    var v99 int = -514
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
    _ = v98
    _ = v99
}
