package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 string; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 bool; F2 int16 }
type GenS11 struct { F0 int; F1 string }
type GenS12 struct { F0 uint8; F1 uint8 }
type GenS13 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS14 struct { F0 int32; F1 uint16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int32; F1 string }
func genFn0(p0 bool, p1 uint, p2 string) int64 {
    var v7 int = -9223372036854775808
    for fi8 := 0; fi8 < 11; fi8++ {
        v7 = 42
        if p0 {
            for fi9 := 0; fi9 < 15; fi9++ {
                p1 -= (p1 | 908)
                p2 = "foo bar"
                _ = fi9
            }
        }
        _ = fi8
    }
    var v10 uint16 = 637
    v7 -= (v7 % ((v7) | 1))
    v7 = (v7 + v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 int16 = 42
    var v2 uint16 = 749
    v3 := []int{808}
    var v4 float64 = 3.14159
    var v5 uint32 = 42
    v6 := make(map[string]string)
    var v11 int8 = 113
    var v12 float32 = 0.0
    switch "foo bar" {
    case "a":
        if (v2 == (v2 / ((1) | 1))) {
            var v13 uint32 = (v5 | 0)
            for fi14 := 0; fi14 < 20; fi14++ {
                for fi15 := 0; fi15 < 4; fi15++ {
                    v5 -= (v13 * v5)
                    _ = fi15
                }
                _ = fi14
            }
            var v16 int32 = 42
            _ = v13
            _ = v16
        } else {
            v1 -= 1000
        }
    default:
        switch 32767 {
        case -21:
            switch 601 {
            case 0:
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 int = (fi17 & (-(fi17)))
                    v5 = 4294967295
                    _ = fi17
                    _ = v18
                }
            case 2147483647:
                switch 436 {
                case 9223372036854775807:
                    var v19 int = 2147483647
                    _ = v19
                case 128:
                    if false {
                        v2 += (^(v2))
                        var v20 string = "hello"
                        var v21 int64 = 395
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v4 = v4
                            var v23 int32 = 0
                            _ = fi22
                            _ = v23
                        }
                        fmt.Println(v4, v4, v11)
                    }
                }
            default:
                switch 127 {
                case -195:
                    fmt.Println(v11, v1)
                case 32767:
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v12 = v12
                        if ((v12 == (v12 + v12)) || true) {
                            if false {
                                if true {
                                    fmt.Println(v12, v12, v2)
                                    v6[""] = ""
                                    fmt.Println(v5, v5)
                                }
                                if true {
                                    switch -1 {
                                    case 32767:
                                        switch 9223372036854775807 {
                                        case -85:
                                            _ = len(v6)
                                        case 32767:
                                            var v24 uint64 = 30
                                            _ = v24
                                        case 771:
                                            v25 := GenS1{F0: v11, F1: (^(v2)), F2: int(v11)}
                                            _ = v25
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v3 = append(v3, -345)
                                }
                                for fi26 := 0; fi26 < 19; fi26++ {
                                    if ((-(fi26)) <= -419) {
                                        if false {
                                            for fi27 := 0; fi27 < 1; fi27++ {
                                                v12 += ((v12 - v12) - 0.1)
                                                var v28 float64 = 0.1
                                                _ = fi27
                                                _ = v28
                                            }
                                            for fi29 := 0; fi29 < 15; fi29++ {
                                                var v30 float32 = v12
                                                v31 := GenS2{F0: "foo bar", F1: (v2 - v2)}
                                                if (!true) {
                                                    v32 := GenS2{F0: "\n\t", F1: (v2 ^ (v2 << ((65535) & 63)))}
                                                    v33 := GenS1{F0: v11, F1: v2, F2: -196}
                                                    var v34 int8 = (v11 % ((v11) | 1))
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                }
                                                _ = fi29
                                                _ = v30
                                                _ = v31
                                            }
                                            if true {
                                                v4 = ((-2.25 - v4) * v4)
                                                var v35 int16 = v1
                                                switch (^(fi26)) {
                                                case 9223372036854775807:
                                                    for fi36 := 0; fi36 < 7; fi36++ {
                                                        if (!(121 != -778)) {
                                                            v5 -= (v5 % (((v5 / ((v5) | 1))) | 1))
                                                        } else {
                                                            v35 = -840
                                                        }
                                                        v35 += v35
                                                        switch (fi36 + fi26) {
                                                        case 9223372036854775807:
                                                            v12 = (v12 + v12)
                                                        }
                                                        _ = fi36
                                                    }
                                                case -329:
                                                    if (!false) {
                                                        v2 += (v2 % ((42) | 1))
                                                        v6["a"] = "foo bar"
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = fi26
                                                        }
                                                        var v37 int32 = -93
                                                        _ = v37
                                                    }
                                                default:
                                                    switch ((("foo bar" + "\n\t") + ("foo bar" + "")) + "a") {
                                                    case "hello":
                                                        v38 := GenS3{F0: v35, F1: uint64(v2), F2: 109}
                                                        _ = v38
                                                    }
                                                }
                                                _ = v35
                                            }
                                        }
                                        if ((v2 * 32767) != v2) {
                                            for fi39 := 0; fi39 < 20; fi39++ {
                                                v40 := GenS3{F0: v1, F1: 42, F2: 0}
                                                if false {
                                                    delete(v6, "hello")
                                                    v41 := GenS4{F0: 1}
                                                    _ = v41
                                                }
                                                var v42 bool = ((((v1 ^ v1) >= (^(v1))) || ((!false) && false)) || ((((v2 + (v2 | 65535)) < v2) && true) || ((v2 ^ v2) <= v2)))
                                                _ = fi39
                                                _ = v40
                                                _ = v42
                                            }
                                            switch "a" {
                                            case "hello":
                                                switch -611 {
                                                case 18:
                                                    if (!(!false)) {
                                                        v11 = (v11 / (((v11 / (((v11 % ((v11) | 1))) | 1))) | 1))
                                                        switch 128 {
                                                        case -9223372036854775808:
                                                            var v43 uint32 = (v5 | (v5 ^ v5))
                                                            _ = v43
                                                        default:
                                                            var v44 int32 = 2147483647
                                                            _ = v44
                                                        }
                                                    }
                                                }
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    var v45 int64 = 1
                                                    _ = v45
                                                }
                                            }
                                        } else {
                                            switch fi26 {
                                            case 54:
                                                var v46 int8 = (v11 - -1)
                                                _ = v46
                                            default:
                                                for fi47 := 0; fi47 < 2; fi47++ {
                                                    v48 := GenS5{F0: (v2 / (((v2 % ((v2) | 1))) | 1))}
                                                    if true {
                                                        var v49 int64 = -132
                                                        v4 += v4
                                                        _ = v49
                                                    }
                                                    v50 := GenS6{F0: (v1 + (^(v1))), F1: "", F2: (-(v11))}
                                                    _ = fi47
                                                    _ = v48
                                                    _ = v50
                                                }
                                            }
                                            for fi51 := 0; fi51 < 3; fi51++ {
                                                v52 := v6["\n\t"]
                                                _ = fi51
                                                _ = v52
                                            }
                                        }
                                        var v53 float64 = (v4 * (float64(((v12 * (v12 - float32(v4))) + v12)) - (1.5 * v4)))
                                        _ = v53
                                    } else {
                                        fmt.Println(v12, v11)
                                    }
                                    _ = len(v3)
                                    _ = fi26
                                }
                            }
                            v54 := GenS7{F0: v1}
                            if (!(!false)) {
                                v55 := GenS4{F0: 127}
                                _ = v55
                            } else {
                                v12 += float32(0.0)
                            }
                            _ = v54
                        }
                    }
                }
            }
        default:
            v5 = (v5 >> ((42) & 63))
        }
    }
    switch 9223372036854775807 {
    case -647:
        var v56 int = 127
        _ = v56
    case -707:
        if len(v3) > 0 {
            v3[0] = 485
        }
    case -251:
        var v57 uint64 = 4294967295
        _ = v57
    default:
        for fi58 := 0; fi58 < 13; fi58++ {
            switch ("foo bar" + "hello") {
            case "\n\t":
                fmt.Println(fi58, v5, v2)
            default:
                if false {
                    var v59 int64 = -1
                    switch fi58 {
                    case 127:
                        fmt.Println(v59, v2)
                    case 327:
                        v60 := GenS8{F0: 38, F1: 42, F2: -128}
                        _ = v60
                    case 2147483647:
                        v61 := GenS9{F0: ((fi58 >= -967) && (v59 <= -419))}
                        _ = v61
                    default:
                        switch fi58 {
                        case 32767:
                            v62 := GenS10{F0: -46, F1: ((v5 ^ v5) < (v5 | v5)), F2: 967}
                            _ = v62
                        case 42:
                            for fi63 := 0; fi63 < 4; fi63++ {
                                var v64 uint64 = 4294967295
                                if false {
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        var v65 float32 = v12
                                        if (v1 >= 1000) {
                                            var v66 int = fi63
                                            _ = v66
                                        } else {
                                            if (true || true) {
                                                v4 = (v4 - v4)
                                            } else {
                                                var v67 int32 = 42
                                                v68 := v6[""]
                                                fmt.Println(fi58, v65, v5)
                                                _ = v67
                                                _ = v68
                                            }
                                            if true {
                                                var v69 float64 = 3.14159
                                                fmt.Println(v2, v12, v59)
                                                if ((v69 < 1e10) && (true || ("foo bar" < "\n\t"))) {
                                                    var v70 uint8 = 14
                                                    v1 -= 42
                                                    _ = v70
                                                }
                                                _ = v69
                                            }
                                        }
                                        if (1 < -2147483648) {
                                            v11 -= 127
                                        }
                                        fmt.Println(v65, v65)
                                        if (739 <= 496) {
                                            var v71 uint8 = 253
                                            _ = v71
                                        } else {
                                            v2 += (v2 - uint16(v5))
                                            var v72 string = ("" + ("a" + "hello"))
                                            var v73 float32 = 1e10
                                            _ = v72
                                            _ = v73
                                        }
                                        _ = v65
                                    case "a":
                                        v74 := GenS9{F0: false}
                                        _ = v74
                                    default:
                                        fmt.Println(v11)
                                    }
                                    v75 := GenS11{F0: fi58, F1: "\n\t"}
                                    v2 -= v2
                                    _ = v75
                                }
                                var v76 int16 = 1000
                                if (v4 != -2.25) {
                                    _ = len(v3)
                                }
                                if false {
                                    if true {
                                        fmt.Println(v4)
                                    }
                                    var v77 int8 = (v11 % ((v11) | 1))
                                    if (v77 < (v11 >> ((42) & 63))) {
                                        if true {
                                            v78 := GenS3{F0: 42, F1: (v64 - (v64 >> ((1) & 63))), F2: 42}
                                            var v79 string = "\n\t"
                                            _ = len(v3)
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        _ = len(v6)
                                    case "foo bar":
                                        var v80 int8 = (^(v77))
                                        _ = v80
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = fi58
                                        }
                                    }
                                    _ = len(v3)
                                    if (((false && (!((v5 >> ((151) & 63)) < 0))) && true) && ((((("foo bar" + "\n\t") + ("\n\t" + ("" + ("a" + "a")))) + "a") + "") > ("a" + ""))) {
                                        if true {
                                            fmt.Println(v5, v5)
                                        } else {
                                            v4 = v4
                                            _ = len(v6)
                                            if true {
                                                var v81 bool = false
                                                v4 = v4
                                                var v82 float32 = v12
                                                _ = v81
                                                _ = v82
                                            } else {
                                                if (1 == 18446744073709551615) {
                                                    v83 := GenS11{F0: fi58, F1: "foo bar"}
                                                    v64 = (v64 | (v64 >> ((0) & 63)))
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v4 = (1.5 + (((1.5 * ((v4 / v4) * v4)) + ((v4 - v4) / v4)) * ((float64((float32(1e10) * (v12 - (1e10 * (0.1 / 1.5))))) + v4) / (3.14159 - (v4 / 1.5)))))
                                                    v77 = (v77 - int8(v2))
                                                    _ = v83
                                                }
                                                if false {
                                                    v84 := v6["hello"]
                                                    v85 := GenS12{F0: 42, F1: 1}
                                                    _ = v84
                                                    _ = v85
                                                }
                                                switch fi63 {
                                                case -9223372036854775808:
                                                    v12 = v12
                                                case 2147483647:
                                                    v86 := GenS13{F0: 615, F1: v2, F2: (^(v2))}
                                                    _ = v86
                                                case 1:
                                                    v3 = append(v3, (fi58 | (fi63 * int(v59))))
                                                }
                                                if (v5 <= v5) {
                                                    fmt.Println(v5)
                                                    fmt.Println(v4)
                                                }
                                                v64 = v64
                                            }
                                        }
                                    }
                                    if false {
                                        if true {
                                            v1 = 1000
                                            fmt.Println(fi63)
                                        }
                                        v1 -= v76
                                    }
                                    _ = v77
                                }
                                v1 += 1000
                                _ = fi63
                                _ = v64
                                _ = v76
                            }
                        case 127:
                            fmt.Println(v5, v59, v12)
                        }
                    }
                    _ = v59
                } else {
                    for fi87 := 0; fi87 < 14; fi87++ {
                        switch fi87 {
                        case -603:
                            var v88 int16 = -32768
                            _ = v88
                        default:
                            v89 := GenS14{F0: 42, F1: (v2 + v2)}
                            _ = v89
                        }
                        _ = fi87
                    }
                }
            }
            if false {
                v90 := GenS15{F0: 2}
                fmt.Println(v11)
                _ = v90
            } else {
                v2 -= 32767
            }
            var v91 int32 = 794
            _ = fi58
            _ = v91
        }
    }
    for fi92 := 0; fi92 < 4; fi92++ {
        v93 := GenS16{F0: -2147483648, F1: "\n\t"}
        fmt.Println(v2, v4, fi92)
        var v94 uint16 = v2
        _ = fi92
        _ = v93
        _ = v94
    }
    if false {
        var v95 uint32 = (^(v5))
        v95 = (^(v5))
        _ = v95
    }
    v4 = v4
    var v96 uint16 = (v2 % ((v2) | 1))
    if true {
        fmt.Println(v5, v12, v12)
    } else {
        switch 0 {
        case 42:
            var v97 int16 = (v1 * 42)
            _ = v97
        default:
            v12 = (v12 / 0.1)
        }
        v96 -= (^(v96))
    }
    v4 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v96:", v96)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v96
}
