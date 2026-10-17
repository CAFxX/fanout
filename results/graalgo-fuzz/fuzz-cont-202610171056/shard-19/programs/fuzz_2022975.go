package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int; F1 int16; F2 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 string; F1 uint32 }
type GenS12 struct { F0 int8; F1 uint32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS15 struct { F0 uint8; F1 int32 }
type GenS16 struct { F0 uint8; F1 uint }
type GenS17 struct { F0 uint64; F1 int8; F2 int8 }
type GenS18 struct { F0 int8; F1 int32 }
type GenS19 struct { F0 int8; F1 int32; F2 string }
type GenS20 struct { F0 int16; F1 uint; F2 uint64 }
func genFn0(p0 uint16) uint16 {
    v6 := GenS1{F0: 127, F1: (!(2147483647 < 65535))}
    v7 := []int{518, 128}
    if (9223372036854775807 < 1) {
        fmt.Println(p0)
        switch 2147483647 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 17; fi8++ {
                var v9 int32 = 545
                v10 := GenS2{F0: (fi8 + 4294967295)}
                switch "\n\t" {
                case "hello":
                    fmt.Println(v9)
                case "a":
                    var v11 int32 = v9
                    _ = v11
                case "foo bar":
                    if (927 < 1) {
                        v12 := GenS2{F0: fi8}
                        if false {
                            fmt.Println(fi8, fi8)
                            fmt.Println(v9)
                            p0 -= (^(p0))
                        }
                        _ = v12
                    } else {
                        v13 := GenS3{F0: 502}
                        _ = v13
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            for fi14 := 0; fi14 < 11; fi14++ {
                p0 -= 65535
                _ = fi14
            }
        }
    } else {
        v15 := GenS1{F0: 100, F1: false}
        if ((^(p0)) >= (^(p0))) {
            if len(v7) > 0 {
                v7[0] = -1
            }
        } else {
            if ((-1 < 1000) || (!(!(127 < 42)))) {
                v16 := GenS2{F0: 0}
                fmt.Println(p0)
                _ = v16
            }
        }
        if (!(false && false)) {
            switch 127 {
            case 0:
                p0 -= (p0 / ((p0) | 1))
            case 32768:
                if (986 < 912) {
                    _ = len(v7)
                    for fi17 := 0; fi17 < 16; fi17++ {
                        for fi18 := 0; fi18 < 20; fi18++ {
                            p0 -= 2
                            var v19 uint64 = 18446744073709551615
                            var v20 uint = 42
                            fmt.Println(fi17, fi18)
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        if (1 == 4294967295) {
                            if (float64(-2.25) <= (1e10 / 1e10)) {
                                v21 := GenS4{F0: 1, F1: 255}
                                p0 += p0
                                p0 += p0
                                _ = v21
                            }
                            var v22 string = (((("hello" + "\n\t") + "foo bar") + "foo bar") + "hello")
                            v22 = "hello"
                            _ = v22
                        } else {
                            v23 := GenS3{F0: 42}
                            p0 = (p0 | (p0 ^ 42))
                            _ = v23
                        }
                        _ = fi17
                    }
                    if (p0 == (p0 % ((p0) | 1))) {
                        v24 := GenS2{F0: int(p0)}
                        _ = v24
                    }
                }
            case -9223372036854775808:
                var v25 bool = true
                _ = v25
            }
        } else {
            var v26 float32 = (1.5 + 1.5)
            v26 += v26
            _ = v26
        }
        _ = v15
    }
    _ = p0
    _ = v6
    _ = v7
    return (p0 * (p0 % ((255) | 1)))
}
func genFn1(p0 string) bool {
    v27 := []int{2147483647, 9223372036854775807, -1}
    p0 = (p0 + p0)
    _ = p0
    _ = v27
    return (false || true)
}

func main() {
    var v1 uint32 = 42
    var v2 int32 = 42
    v3 := []uint32{}
    var v4 uint16 = 32767
    var v5 bool = ((true && false) && false)
    v28 := GenS2{F0: -776}
    v1 += v1
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            switch int(v1) {
            case 127:
                fmt.Println(v5, v2)
            case 4294967295:
                if v5 {
                    if false {
                        var v31 uint8 = 167
                        var v32 uint64 = 18446744073709551615
                        v33 := GenS5{F0: (v1 - (v1 << ((65535) & 63))), F1: 1}
                        if v5 {
                            if (genFn0(genFn0(genFn0((v4 / ((v4) | 1))))) <= v4) {
                                v34 := GenS6{F0: -128, F1: (v31 >= 1)}
                                _ = len(v3)
                                _ = v34
                            } else {
                                var v35 float64 = (float64(((0.1 + (0.1 + 3.14159)) - (0.1 * 1.5))) + (((3.14159 + 0.0) - float64((0.1 * -2.25))) / 3.14159))
                                v2 += 848
                                _ = v35
                            }
                            var v36 int8 = 6
                            _ = v36
                        } else {
                            v37 := GenS4{F0: -980, F1: (v31 % (((v31 % (((v31 ^ v31)) | 1))) | 1))}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v2)
                            _ = v37
                        }
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v38 := GenS7{F0: 690}
                    _ = v38
                }
            case -2147483648:
                if v5 {
                    v39 := GenS8{F0: fi30, F1: 0, F2: 496}
                    if v5 {
                        if ("\n\t" < "a") {
                            if ((v2 | v2) > (v2 - -1)) {
                                v4 -= (v4 - 255)
                            }
                            v40 := GenS5{F0: v1, F1: -1}
                            _ = v40
                        } else {
                            var v41 uint16 = 968
                            var v42 uint16 = (v41 | v41)
                            v43 := GenS9{F0: 0}
                            v44 := GenS10{F0: 232, F1: genFn0(v41)}
                            _ = len(v3)
                            _ = v41
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        switch (^(fi29)) {
                        case -9223372036854775808:
                            var v45 uint32 = (^(v1))
                            _ = v45
                        case -21:
                            v4 += v4
                        default:
                            switch 9223372036854775807 {
                            case 2147483647:
                                if (1 > 2) {
                                    switch 4294967295 {
                                    case 827:
                                        var v46 string = ("foo bar" + (("\n\t" + "hello") + (("foo bar" + ((("hello" + "a") + "") + "")) + "hello")))
                                        _ = v46
                                    case 1:
                                        if (true && ((!(!v5)) && (65535 < 65535))) {
                                            v4 += 32767
                                        }
                                    default:
                                        v4 = 313
                                        var v47 float32 = (3.14159 * 1e10)
                                        _ = v47
                                    }
                                } else {
                                    switch ("\n\t" + "") {
                                    case "a":
                                        var v48 int8 = 42
                                        _ = v48
                                    case "foo bar":
                                        var v49 int64 = -9223372036854775808
                                        _ = v49
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = (v1 ^ (v1 + (v1 << ((4294967295) & 63))))
                                        }
                                    default:
                                        switch (fi29 / ((fi30) | 1)) {
                                        case 1:
                                            if false {
                                                if ((v1 % ((v1) | 1)) > v1) {
                                                    fmt.Println(fi30)
                                                    if false {
                                                        fmt.Println(v5, v2, v1)
                                                        var v50 int16 = 32767
                                                        switch "a" {
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        }
                                                        _ = v50
                                                    } else {
                                                        v5 = ("hello" == ("\n\t" + "foo bar"))
                                                        v51 := GenS2{F0: (fi30 >> ((4294967295) & 63))}
                                                        v4 -= (v4 * (v4 / (((v4 & uint16(v1))) | 1)))
                                                        _ = v51
                                                    }
                                                } else {
                                                    var v52 uint32 = (v1 % (((v1 | (v1 * (v1 + v1)))) | 1))
                                                    if (false || v5) {
                                                        fmt.Println(v2, fi30)
                                                        if (-421 >= -695) {
                                                            v2 = 42
                                                            switch -9223372036854775808 {
                                                            case 9223372036854775807:
                                                                var v53 uint32 = (v52 << ((255) & 63))
                                                                _ = v53
                                                            case 32767:
                                                                v1 -= (v1 >> ((271) & 63))
                                                            case 4294967295:
                                                                switch fi30 {
                                                                case 1:
                                                                    if (339 == (^(v4))) {
                                                                        v1 += 2
                                                                    } else {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v54 := GenS11{F0: "", F1: (v52 - (v52 & v52))}
                                                                            _ = v54
                                                                        case "hello":
                                                                            if v5 {
                                                                                switch ("hello" + "\n\t") {
                                                                                case "\n\t":
                                                                                    var v55 float64 = ((3.14159 - 0.0) * (((1.5 * 3.14159) / 1e10) / (float64((float32(1.5) * -2.25)) * 1.5)))
                                                                                    _ = v55
                                                                                }
                                                                                fmt.Println(v52, v4)
                                                                            }
                                                                        case "":
                                                                            v56 := GenS12{F0: -1, F1: (v52 << ((255) & 63))}
                                                                            _ = v56
                                                                        }
                                                                        _ = len(v3)
                                                                        if len(v3) > 0 {
                                                                            v3[0] = 1
                                                                        }
                                                                    }
                                                                case 828:
                                                                    v57 := GenS13{F0: ((-(v2)) > -2147483648)}
                                                                    _ = v57
                                                                default:
                                                                    var v58 int16 = 1
                                                                    _ = v58
                                                                }
                                                            }
                                                            if v5 {
                                                                v5 = (255 < 127)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                switch (fi30 ^ fi29) {
                                                                case -9223372036854775808:
                                                                    var v59 int32 = 228
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v60 int64 = 1
                                                                        _ = v60
                                                                    case "a":
                                                                        if false {
                                                                            switch fi30 {
                                                                            case -258:
                                                                                if true {
                                                                                    v61 := GenS13{F0: (-457 >= fi30)}
                                                                                    _ = v61
                                                                                } else {
                                                                                    switch (^(fi29)) {
                                                                                    case 9223372036854775807:
                                                                                        fmt.Println(v52, v59, v2)
                                                                                    case 602:
                                                                                        if (!(!v5)) {
                                                                                            var v62 int16 = 390
                                                                                            var v63 int16 = v62
                                                                                            if v5 {
                                                                                                v1 = v52
                                                                                                var v64 uint = 255
                                                                                                v65 := GenS9{F0: 2}
                                                                                                _ = v64
                                                                                                _ = v65
                                                                                            }
                                                                                            _ = v62
                                                                                            _ = v63
                                                                                        } else {
                                                                                            var v66 string = "foo bar"
                                                                                            _ = v66
                                                                                        }
                                                                                    case 2147483647:
                                                                                        v67 := GenS6{F0: int8(v4), F1: (644 < v52)}
                                                                                        _ = v67
                                                                                    default:
                                                                                        if v5 {
                                                                                            var v68 int16 = 0
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (v52 + (v1 ^ (v1 & v52)))
                                                                                            }
                                                                                            _ = v68
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case 71:
                                                                                fmt.Println(fi30, v52, v2)
                                                                            }
                                                                            v52 = v52
                                                                            switch ("hello" + "foo bar") {
                                                                            case "":
                                                                                _ = len(v3)
                                                                            case "a":
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (v1 | v52)
                                                                                }
                                                                            case "\n\t":
                                                                                var v69 int32 = (-(v2))
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v70 := GenS4{F0: -9223372036854775808, F1: 65}
                                                                        switch (fi30 | int(v4)) {
                                                                        case -188:
                                                                            var v71 int16 = 1000
                                                                            v5 = (32767 > (v71 - (v71 % (((v71 ^ (v71 ^ v71))) | 1))))
                                                                            _ = v71
                                                                        case -165:
                                                                            v2 = int32(v52)
                                                                        }
                                                                        v59 -= v59
                                                                        _ = v70
                                                                    }
                                                                    _ = v59
                                                                }
                                                            } else {
                                                                v72 := GenS6{F0: 127, F1: (fi30 != -38)}
                                                                _ = v72
                                                            }
                                                            var v73 uint = 0
                                                            _ = v73
                                                        } else {
                                                            if true {
                                                                var v74 uint64 = 1
                                                                var v75 int32 = 107
                                                                _ = v74
                                                                _ = v75
                                                            } else {
                                                                var v76 int8 = 1
                                                                _ = v76
                                                            }
                                                        }
                                                        v5 = (-487 != fi29)
                                                    }
                                                    _ = v52
                                                }
                                                v77 := GenS14{F0: 42, F1: 4294967295, F2: 100}
                                                v4 = 2
                                                _ = v77
                                            } else {
                                                if v5 {
                                                    v1 -= v1
                                                    v3 = append(v3, v1)
                                                }
                                            }
                                        }
                                    }
                                    v2 += 2147483647
                                }
                            case 1:
                                fmt.Println(v1, fi29, v2)
                            default:
                                v4 += (v4 ^ v4)
                            }
                        }
                    } else {
                        v2 += (-(v2))
                        v78 := GenS15{F0: 14, F1: v2}
                        _ = v78
                    }
                    if len(v3) > 0 {
                        v3[0] = v1
                    }
                    fmt.Println(fi29, fi30)
                    _ = v39
                } else {
                    switch 127 {
                    case 353:
                        v79 := GenS16{F0: 1, F1: 35}
                        _ = v79
                    case 2147483647:
                        _ = len(v3)
                    }
                    _ = len(v3)
                }
            }
            v3 = append(v3, (v1 | v1))
            _ = fi30
        }
        v80 := GenS2{F0: fi29}
        v4 -= 177
        _ = fi29
        _ = v80
    }
    var v81 bool = (!((1.5 / 1.0) > (float32(0.0) * 1.5)))
    v3 = append(v3, (v1 | 412))
    fmt.Println(v4, v81, v81)
    var v82 uint32 = v1
    if (v81 && v81) {
        for fi83 := 0; fi83 < 9; fi83++ {
            v84 := GenS17{F0: 0, F1: 1, F2: -75}
            var v85 uint8 = uint8(v4)
            _ = fi83
            _ = v84
            _ = v85
        }
        for fi86 := 0; fi86 < 4; fi86++ {
            if v5 {
                for fi87 := 0; fi87 < 13; fi87++ {
                    fmt.Println(v1, v1, v4)
                    var v88 bool = v5
                    _ = fi87
                    _ = v88
                }
                var v89 uint8 = 127
                _ = v89
            } else {
                var v90 uint32 = v82
                _ = v90
            }
            var v91 uint8 = 42
            _ = fi86
            _ = v91
        }
    } else {
        v2 += v2
        if (float64(((0.1 + 3.14159) * 1e10)) != 3.14159) {
            v92 := GenS18{F0: -128, F1: (v2 >> ((0) & 63))}
            _ = v92
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v81, v82)
    for fi93 := 0; fi93 < 5; fi93++ {
        if v5 {
            switch -1 {
            case 9223372036854775807:
                v3 = append(v3, 0)
            case 140:
                v1 = (v82 & 1)
            case 127:
                v94 := GenS16{F0: 247, F1: uint(v4)}
                _ = v94
            default:
                v2 = v2
            }
            v3 = append(v3, 2)
        } else {
            if v5 {
                var v95 float64 = (0.1 - 0.1)
                _ = v95
            } else {
                v96 := GenS17{F0: 100, F1: 0, F2: 42}
                _ = v96
            }
        }
        var v97 int8 = -128
        switch fi93 {
        case 2147483647:
            switch (fi93 * 4294967295) {
            case -1:
                fmt.Println(v1, v4)
            case -9223372036854775808:
                fmt.Println(v82, v4, v81)
            case 586:
                if len(v3) > 0 {
                    v3[0] = (v82 % ((2147483648) | 1))
                }
            }
        case 9223372036854775807:
            var v98 uint8 = 42
            _ = v98
        }
        _ = fi93
        _ = v97
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v99 := GenS6{F0: 0, F1: true}
    for fi100 := 0; fi100 < 18; fi100++ {
        if ((((!((v4 - genFn0(343)) < 32767)) || v81) || false) || false) {
            for fi101 := 0; fi101 < 6; fi101++ {
                v2 -= int32(fi101)
                if (fi100 > (fi100 | (fi100 >> ((66) & 63)))) {
                    v102 := GenS15{F0: 208, F1: -2147483648}
                    _ = len(v3)
                    var v103 uint16 = (v4 << ((18446744073709551615) & 63))
                    _ = v102
                    _ = v103
                } else {
                    switch 582 {
                    case 2147483647:
                        v2 = v2
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (128 != uint8(v82)) {
                            v104 := GenS19{F0: -29, F1: (v2 % (((v2 | (v2 - v2))) | 1)), F2: "hello"}
                            v105 := GenS3{F0: 0}
                            v106 := GenS20{F0: 1000, F1: 18446744073709551615, F2: 42}
                            _ = v104
                            _ = v105
                            _ = v106
                        }
                        if ((^(v4)) >= (v4 << ((4294967295) & 63))) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch (fi100 | -1) {
                            case -962:
                                var v107 uint8 = 2
                                _ = v107
                            case 4294967295:
                                var v108 uint8 = 168
                                var v109 float64 = (0.0 + (-2.25 + 3.14159))
                                var v110 int = 9223372036854775807
                                _ = v108
                                _ = v109
                                _ = v110
                            }
                            v81 = v5
                        }
                    case 128:
                        _ = len(v3)
                    case 1:
                        var v111 int16 = 1000
                        _ = v111
                    default:
                        switch (("hello" + "foo bar") + "hello") {
                        case "\n\t":
                            v1 = v1
                        case "hello":
                            var v112 int64 = 459
                            var v113 int32 = (v2 >> ((65535) & 63))
                            _ = v112
                            _ = v113
                        default:
                            var v114 float64 = (1.5 + 0.1)
                            _ = v114
                        }
                        var v115 int32 = v2
                        _ = v115
                    }
                }
                v4 += 1
                _ = fi101
            }
            var v116 int8 = 42
            _ = v116
        } else {
            if len(v3) > 0 {
                v3[0] = v82
            }
        }
        if len(v3) > 0 {
            v3[0] = (^(v82))
        }
        v1 += (^(v82))
        _ = fi100
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v81
    _ = v82
    _ = v99
}
