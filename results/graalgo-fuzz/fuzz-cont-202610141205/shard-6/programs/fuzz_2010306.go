package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 int32; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint32; F2 bool }
type GenS5 struct { F0 bool; F1 int64; F2 uint8 }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 int16; F1 string }
func genFn0(p0 int16, p1 int16) int8 {
    fmt.Println(p1, p1)
    var v4 uint8 = 127
    var v5 uint32 = 2147483647
    var v6 uint64 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int) uint32 {
    v7 := []int{2147483647, p0, (^(p0)), 542}
    p0 = 570
    p0 -= 4294967295
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = false
    var v3 uint32 = 547
    v2 = ((-116 > -128) || (true && true))
    fmt.Println(v3)
    var v8 uint16 = 260
    if (v2 && ((128 > 0) || v2)) {
        if (-1 >= -1) {
            var v9 float32 = 1.5
            _ = v9
        }
    }
    if true {
        v10 := GenS1{F0: -196}
        switch ("\n\t" + "foo bar") {
        case "a":
            switch 4294967295 {
            case 32767:
                for fi11 := 0; fi11 < 17; fi11++ {
                    switch "\n\t" {
                    case "":
                        fmt.Println(fi11)
                    case "a":
                        switch "\n\t" {
                        case "\n\t":
                            v1[""] = (v3 << ((490) & 63))
                        case "foo bar":
                            switch -945 {
                            case -2147483648:
                                if (38 != 127) {
                                    var v12 string = "\n\t"
                                    _ = v12
                                } else {
                                    v8 -= v8
                                    v2 = (int32(v3) <= -524)
                                    var v13 bool = (3.14159 < 0.1)
                                    _ = v13
                                }
                            case 32767:
                                if false {
                                    v8 = uint16(v3)
                                    var v14 uint8 = 0
                                    switch "hello" {
                                    case "hello":
                                        for fi15 := 0; fi15 < 9; fi15++ {
                                            v16 := []int{fi15, (fi11 >> ((18446744073709551615) & 63)), (fi11 ^ (-(fi11)))}
                                            switch (fi11 | fi11) {
                                            case 0:
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                            case 2147483647:
                                                v17 := GenS1{F0: -1}
                                                _ = v17
                                            }
                                            _ = fi15
                                            _ = v16
                                        }
                                    case "\n\t":
                                        var v18 int64 = -713
                                        _ = v18
                                    case "foo bar":
                                        fmt.Println(fi11, v3)
                                    }
                                    _ = v14
                                } else {
                                    if (v2 || (!true)) {
                                        var v19 bool = (0.0 < 1e10)
                                        _ = v19
                                    } else {
                                        for fi20 := 0; fi20 < 13; fi20++ {
                                            v21 := GenS1{F0: -1}
                                            if v2 {
                                                var v22 float64 = (0.0 / 1.0)
                                                _ = v22
                                            }
                                            delete(v1, ("\n\t" + ("hello" + "a")))
                                            _ = fi20
                                            _ = v21
                                        }
                                        for fi23 := 0; fi23 < 1; fi23++ {
                                            if true {
                                                if ((v8 << ((1) & 63)) < (v8 % ((0) | 1))) {
                                                    v24 := GenS2{F0: 2, F1: (uint64(v3) == 18446744073709551615)}
                                                    v25 := GenS3{F0: -2147483648, F1: 2, F2: -128}
                                                    switch (fi11 / ((0) | 1)) {
                                                    case 9223372036854775807:
                                                        if (!true) {
                                                            v3 -= (v3 % ((v3) | 1))
                                                            var v26 uint8 = 203
                                                            _ = v26
                                                        } else {
                                                            var v27 int16 = -32768
                                                            v3 = v3
                                                            var v28 float64 = 1.5
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    case 127:
                                                        switch "hello" {
                                                        case "a":
                                                            switch (^(fi23)) {
                                                            case 127:
                                                                if v2 {
                                                                    if (1e10 <= 1e10) {
                                                                        v2 = v2
                                                                        v29 := []int{fi23, (fi23 | (-(fi23))), fi23}
                                                                        switch (fi23 / ((42) | 1)) {
                                                                        case 42:
                                                                            var v30 uint8 = 42
                                                                            _ = v30
                                                                        case -533:
                                                                            var v31 uint8 = 42
                                                                            _ = v31
                                                                        default:
                                                                            _ = len(v1)
                                                                        }
                                                                        _ = v29
                                                                    }
                                                                    if (((^(v3)) <= 974) && (v2 || true)) {
                                                                        if (1 <= 42) {
                                                                            v8 = (v8 % ((v8) | 1))
                                                                            if v2 {
                                                                                v3 += v3
                                                                            }
                                                                            fmt.Println(fi11, fi23, v3)
                                                                        } else {
                                                                            fmt.Println(fi11, v8, v3)
                                                                            v8 = 105
                                                                        }
                                                                    }
                                                                    switch ("\n\t" + ("" + "")) {
                                                                    case "foo bar":
                                                                        v32 := []int{42}
                                                                        fmt.Println(fi23)
                                                                        _ = v32
                                                                    case "\n\t":
                                                                        fmt.Println(v3)
                                                                    default:
                                                                        var v33 int8 = -21
                                                                        _ = v33
                                                                    }
                                                                    v34 := GenS4{F0: (v3 | v3), F1: v3, F2: v2}
                                                                    v35 := []int{(fi23 ^ fi11), (-(fi11)), (fi23 << ((791) & 63)), fi11}
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            case 1:
                                                                if true {
                                                                    v36 := GenS1{F0: 32767}
                                                                    var v37 uint16 = v8
                                                                    switch "foo bar" {
                                                                    case "\n\t":
                                                                        switch 4294967295 {
                                                                        case -108:
                                                                            var v38 uint8 = 147
                                                                            _ = v38
                                                                        case 128:
                                                                            v3 = (^(v3))
                                                                        default:
                                                                            v8 += v8
                                                                        }
                                                                    default:
                                                                        v37 = (v8 / ((v8) | 1))
                                                                    }
                                                                    _ = v36
                                                                    _ = v37
                                                                } else {
                                                                    var v39 float32 = -2.25
                                                                    fmt.Println(v3, fi11)
                                                                    _ = v39
                                                                }
                                                            default:
                                                                v40 := []int{(fi23 + 816)}
                                                                _ = v40
                                                            }
                                                        case "":
                                                            if v2 {
                                                                v8 -= v8
                                                            } else {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    fmt.Println(fi23, v3, fi23)
                                                                case "\n\t":
                                                                    var v41 uint16 = v8
                                                                    _ = v41
                                                                default:
                                                                    v42 := GenS5{F0: false, F1: -1, F2: 42}
                                                                    _ = v42
                                                                }
                                                                v43 := GenS2{F0: 2, F1: v2}
                                                                _ = v43
                                                            }
                                                        default:
                                                            v3 = 93
                                                        }
                                                    case -9223372036854775808:
                                                        var v44 uint16 = (v8 + v8)
                                                        _ = v44
                                                    default:
                                                        var v45 uint8 = 155
                                                        _ = v45
                                                    }
                                                    v8 += (v8 - (v8 ^ v8))
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    var v46 uint16 = (v8 >> ((1) & 63))
                                                    fmt.Println(v8, v3, v2)
                                                    _ = v46
                                                }
                                            } else {
                                                v47 := []int{fi23, fi11, fi23}
                                                switch ("foo bar" + "") {
                                                case "a":
                                                    v48 := GenS4{F0: v3, F1: v3, F2: (((3.14159 - (0.1 - -2.25)) - float64(float32(0.1))) <= 1.5)}
                                                    if true {
                                                        v3 -= genFn1((^(fi11)))
                                                    }
                                                    _ = v48
                                                }
                                                v3 -= (v3 >> ((583) & 63))
                                                _ = v47
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                if v2 {
                                                    switch -2147483648 {
                                                    case 757:
                                                        v49 := GenS1{F0: -32768}
                                                        _ = v49
                                                    case -2147483648:
                                                        switch "foo bar" {
                                                        case "a":
                                                            v50 := GenS6{F0: (v8 ^ v8), F1: 0}
                                                            _ = v50
                                                        case "foo bar":
                                                            fmt.Println(v2, v3, v2)
                                                        default:
                                                            fmt.Println(fi11, fi11)
                                                        }
                                                    case 32768:
                                                        v2 = true
                                                    }
                                                } else {
                                                    var v51 string = (("hello" + (("hello" + "") + "hello")) + "\n\t")
                                                    v52 := GenS7{F0: (v8 * v8)}
                                                    v3 += (v3 * 2147483648)
                                                    _ = v51
                                                    _ = v52
                                                }
                                            }
                                            v53 := GenS8{F0: 1, F1: (^(v3))}
                                            _ = fi23
                                            _ = v53
                                        }
                                    }
                                    fmt.Println(v3, fi11)
                                    v54 := GenS8{F0: -700, F1: genFn1((fi11 << ((42) & 63)))}
                                    _ = v54
                                }
                            default:
                                v55 := []int{(^(fi11)), fi11, fi11}
                                _ = v55
                            }
                        case "hello":
                            fmt.Println(v2, v3, v8)
                        }
                    }
                    _ = fi11
                }
            case 120:
                v3 = (^(v3))
            case 1:
                var v56 float32 = 3.14159
                _ = v56
            }
        case "":
            if (false && (true || v2)) {
                var v57 int = 9223372036854775807
                if ((!false) && (3.14159 < 1.5)) {
                    v58 := GenS5{F0: (!true), F1: -9223372036854775808, F2: 2}
                    _ = v58
                } else {
                    v3 -= (^(v3))
                }
                v8 -= 618
                _ = v57
            }
        default:
            v1[""] = v3
        }
        v59 := GenS9{F0: (v3 | 42), F1: (((("" + "") + (("\n\t" + "foo bar") + "")) + "\n\t") + "")}
        _ = v10
        _ = v59
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        var v61 int32 = 232
        v61 = v61
        var v62 float32 = 0.1
        _ = fi60
        _ = v61
        _ = v62
    }
    v3 = v3
    v2 = true
    fmt.Println(v3, v2, v8)
    var v63 int64 = 9223372036854775807
    var v64 uint16 = v8
    fmt.Println(v64, v64, v64)
    var v65 uint16 = v8
    for fi66 := 0; fi66 < 17; fi66++ {
        var v67 int16 = 42
        _ = fi66
        _ = v67
    }
    switch "a" {
    case "hello":
        fmt.Println(v65, v8, v2)
    default:
        var v68 int = -29
        _ = v68
    }
    var v69 bool = (712 < 4294967295)
    v64 = v64
    if (v2 && (("" + ("" + "hello")) > "hello")) {
        var v70 float64 = 0.0
        switch (("" + "hello") + "\n\t") {
        case "\n\t":
            if (v2 && v69) {
                var v71 bool = false
                v65 -= v65
                for fi72 := 0; fi72 < 19; fi72++ {
                    v70 = v70
                    switch fi72 {
                    case 2147483647:
                        for fi73 := 0; fi73 < 11; fi73++ {
                            delete(v1, "a")
                            fmt.Println(v63, v71)
                            switch "" {
                            case "\n\t":
                                var v74 uint16 = (v8 << ((293) & 63))
                                _ = v74
                            }
                            _ = fi73
                        }
                    case 4294967295:
                        switch ("hello" + "\n\t") {
                        case "hello":
                            fmt.Println(v63, v70, v8)
                        case "\n\t":
                            if v71 {
                                v75 := GenS1{F0: 42}
                                v2 = v69
                                _ = v75
                            }
                        case "foo bar":
                            v70 = (v70 - float64(float32((v70 - (-2.25 / v70)))))
                        }
                    default:
                        fmt.Println(v69, v65)
                    }
                    var v76 uint8 = 127
                    _ = fi72
                    _ = v76
                }
                _ = v71
            } else {
                for fi77 := 0; fi77 < 18; fi77++ {
                    var v78 int32 = 1
                    v70 -= (1.5 / float64((float32(1e10) / (1e10 + 0.1))))
                    _ = fi77
                    _ = v78
                }
                if v69 {
                    for fi79 := 0; fi79 < 14; fi79++ {
                        v80 := GenS4{F0: (v3 & (v3 / ((genFn1(fi79)) | 1))), F1: (v3 >> ((65535) & 63)), F2: (fi79 >= (fi79 & 4294967295))}
                        _ = fi79
                        _ = v80
                    }
                    if (42 == -2147483648) {
                        var v81 uint32 = (v3 << ((4294967295) & 63))
                        var v82 string = ""
                        _ = v81
                        _ = v82
                    } else {
                        for fi83 := 0; fi83 < 3; fi83++ {
                            v8 -= v8
                            var v84 float32 = (3.14159 * (0.1 + (0.0 / -2.25)))
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(fi83, v3)
                            default:
                                var v85 uint8 = 48
                                _ = v85
                            }
                            _ = fi83
                            _ = v84
                        }
                        v8 = 255
                        for fi86 := 0; fi86 < 1; fi86++ {
                            if v2 {
                                if (1000 >= -32768) {
                                    var v87 uint32 = genFn1((fi86 ^ fi86))
                                    var v88 float64 = (v70 / (v70 * v70))
                                    v63 = (v63 - 9223372036854775807)
                                    _ = v87
                                    _ = v88
                                }
                                v89 := GenS4{F0: v3, F1: 1, F2: v69}
                                var v90 uint8 = 255
                                _ = v89
                                _ = v90
                            } else {
                                v91 := GenS8{F0: (v63 - 455), F1: v3}
                                switch ("foo bar" + "a") {
                                case "":
                                    v92 := GenS8{F0: (v63 & 904), F1: (v3 >> ((899) & 63))}
                                    _ = v92
                                case "foo bar":
                                    if ((((v70 - v70) / (v70 - v70)) - 1.5) < v70) {
                                        v3 = (v3 + (^(v3)))
                                        if (!true) {
                                            v2 = ("hello" == "foo bar")
                                        }
                                        v63 = v63
                                    } else {
                                        switch -2147483648 {
                                        case 4294967295:
                                            v70 -= (v70 + ((v70 * ((1.5 - ((float64(1e10) / (v70 / v70)) - ((0.0 - (3.14159 / (0.1 - v70))) * v70))) / ((v70 + v70) * v70))) - ((v70 * ((1e10 + (v70 + 1e10)) / v70)) * float64(((3.14159 - 0.0) * (0.0 * 1.5))))))
                                        default:
                                            switch (fi86 - fi86) {
                                            case 9223372036854775807:
                                                fmt.Println(v63, v70)
                                            case 2147483647:
                                                v8 = 0
                                            case 127:
                                                v93 := GenS1{F0: -1}
                                                _ = v93
                                            default:
                                                switch 128 {
                                                case 127:
                                                    v94 := GenS3{F0: -1, F1: uint8(v65), F2: 0}
                                                    _ = v94
                                                default:
                                                    v95 := v1[(("\n\t" + "a") + "foo bar")]
                                                    _ = v95
                                                }
                                            }
                                        }
                                        v69 = (32767 == 370)
                                    }
                                default:
                                    v65 += 65535
                                }
                                _ = v91
                            }
                            v64 = (v65 & (^(v8)))
                            var v96 uint = 65535
                            _ = fi86
                            _ = v96
                        }
                    }
                    var v97 uint64 = 42
                    _ = v97
                }
            }
        default:
            v63 = (v63 % ((-1) | 1))
        }
        v98 := GenS10{F0: -1, F1: ("" + "foo bar")}
        _ = v70
        _ = v98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v69:", v69)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v69
    _ = v8
}
