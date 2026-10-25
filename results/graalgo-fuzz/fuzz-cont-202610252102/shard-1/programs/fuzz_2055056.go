package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 int8; F1 uint8; F2 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint64; F2 int }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 uint64; F1 uint; F2 int16 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 int64; F1 int16 }
func genFn0(p0 int8) uint {
    var v5 uint16 = 0
    if false {
        var v6 uint16 = v5
        v7 := []int{}
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        if (!true) {
            if false {
                v9 := GenS1{F0: (-(p0)), F1: 817}
                _ = v9
            }
            v10 := GenS1{F0: (p0 | (p0 << ((213) & 63))), F1: 773}
            _ = v10
        }
        v11 := []int{2147483647, 565}
        _ = fi8
        _ = v11
    }
    var v12 uint16 = v5
    v13 := []int{-2147483648}
    _ = p0
    _ = v12
    _ = v13
    _ = v5
    return 959
}

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 0
    v3 := []string{("" + "a")}
    var v4 float32 = -2.25
    var v14 uint16 = 255
    var v15 string = ""
    var v16 int64 = -1
    if false {
        switch 9223372036854775807 {
        case -1:
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := v1[v15]
                if (2147483647 > 892) {
                    switch v15 {
                    case "hello":
                        for fi19 := 0; fi19 < 16; fi19++ {
                            v2 += v2
                            v2 = -1
                            _ = fi19
                        }
                    case "a":
                        var v20 uint = 65535
                        _ = v20
                    }
                    v21 := v1["hello"]
                    _ = v21
                } else {
                    fmt.Println(fi17)
                    v4 += ((float32(-2.25) * 1.5) + (v4 * float32(0.0)))
                }
                _ = fi17
                _ = v18
            }
        case 1:
            var v22 string = "\n\t"
            _ = v22
        case -755:
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        default:
            v15 = "\n\t"
        }
        switch -140 {
        case 127:
            switch "\n\t" {
            case "foo bar":
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = -168
                }
            case "a":
                if ((0.0 - ((0.0 - ((float64(v4) * (((0.1 - 0.1) - 0.0) * 0.1)) + (float64(v4) * 0.1))) * 0.1)) == -2.25) {
                    switch -9223372036854775808 {
                    case -1:
                        v23 := v1[(("foo bar" + (v15 + v15)) + ("\n\t" + "\n\t"))]
                        _ = v23
                    case 1:
                        var v24 uint32 = 0
                        _ = v24
                    default:
                        _ = len(v3)
                    }
                } else {
                    v15 = v15
                }
            default:
                var v25 int8 = 65
                _ = v25
            }
        }
        fmt.Println(v14, v2)
    } else {
        v2 = (v2 >> ((uint(v14)) & 63))
        fmt.Println(v4, v15)
    }
    v4 = v4
    if (!((v2 % ((42) | 1)) != -721)) {
        if (false && false) {
            var v26 uint32 = 835
            _ = v26
        }
        var v27 string = (("foo bar" + (v15 + (v15 + (("hello" + (v15 + v15)) + (("hello" + v15) + "hello"))))) + v15)
        _ = v27
    } else {
        v28 := GenS1{F0: 42, F1: uint32(v14)}
        _ = v28
    }
    switch (v15 + "hello") {
    case "a":
        v29 := GenS1{F0: 127, F1: 370}
        _ = v29
    }
    v4 -= v4
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 float64 = (-2.25 - 3.14159)
        switch (-(fi30)) {
        case 127:
            if false {
                var v32 uint = genFn0(127)
                for fi33 := 0; fi33 < 9; fi33++ {
                    if false {
                        v14 = v14
                        var v34 int16 = 0
                        v35 := GenS2{F0: 127, F1: (^(v14)), F2: v14}
                        _ = v34
                        _ = v35
                    }
                    _ = fi33
                }
                switch fi30 {
                case 727:
                    v36 := GenS1{F0: -105, F1: 24}
                    _ = v36
                case 4294967295:
                    fmt.Println(v4)
                }
                _ = v32
            } else {
                var v37 uint8 = 46
                if true {
                    var v38 uint = genFn0(0)
                    var v39 int8 = -79
                    _ = v38
                    _ = v39
                }
                v40 := GenS3{F0: true, F1: 1000}
                _ = v37
                _ = v40
            }
        case -2147483648:
            fmt.Println(v15, fi30)
        case 32768:
            var v41 int8 = 127
            _ = v41
        default:
            fmt.Println(v15, v15, fi30)
        }
        _ = fi30
        _ = v31
    }
    v42 := GenS4{F0: (v14 + (v14 * (v14 & v14))), F1: 9223372036854775807}
    v43 := GenS1{F0: 0, F1: 1}
    var v44 bool = ((167 <= 128) && false)
    v45 := GenS2{F0: 0, F1: 65535, F2: v14}
    if false {
        if (-217 == 1) {
            v46 := GenS5{F0: -704, F1: 42}
            var v47 int8 = 1
            if (-984 < v2) {
                v16 -= v16
                if v44 {
                    var v48 uint32 = uint32(v14)
                    var v49 int64 = (v16 * 9223372036854775807)
                    switch (v15 + (v15 + "hello")) {
                    case "a":
                        v44 = v44
                    case "foo bar":
                        for fi50 := 0; fi50 < 7; fi50++ {
                            switch (v15 + ("a" + (v15 + v15))) {
                            case "\n\t":
                                _ = len(v1)
                            case "a":
                                switch v15 {
                                case "":
                                    _ = len(v1)
                                }
                            case "foo bar":
                                v16 = -9223372036854775808
                            default:
                                for fi51 := 0; fi51 < 13; fi51++ {
                                    fmt.Println(v44, v14)
                                    if (!v44) {
                                        v47 = v47
                                        if v44 {
                                            v16 -= -851
                                        } else {
                                            if v44 {
                                                switch ("\n\t" + "foo bar") {
                                                case "foo bar":
                                                    if (1 < 4294967295) {
                                                        v44 = (((v14 | uint16(v48)) != v14) || v44)
                                                        v52 := GenS5{F0: fi51, F1: 0}
                                                        _ = v52
                                                    }
                                                    v48 = (v48 & v48)
                                                case "a":
                                                    v15 = v15
                                                case "hello":
                                                    if v44 {
                                                        v2 = (v2 | (^(v2)))
                                                        switch int(v14) {
                                                        case 721:
                                                            var v53 string = v15
                                                            fmt.Println(v53, v49)
                                                            if v44 {
                                                                if (v44 || true) {
                                                                    if v44 {
                                                                        var v54 float64 = float64(v4)
                                                                        v15 = v53
                                                                        _ = v54
                                                                    } else {
                                                                        fmt.Println(v49)
                                                                        var v55 int16 = -32768
                                                                        v56 := GenS3{F0: (42 == 18446744073709551615), F1: (v55 ^ 32767)}
                                                                        var v57 float64 = 3.14159
                                                                        _ = v55
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                } else {
                                                                    v44 = ((21 < 122) || v44)
                                                                }
                                                            }
                                                            v2 -= v2
                                                            var v58 uint8 = 0
                                                            switch fi51 {
                                                            case 1:
                                                                fmt.Println(v58, v14)
                                                            default:
                                                                var v59 int = fi51
                                                                _ = v59
                                                            }
                                                            _ = v53
                                                            _ = v58
                                                        default:
                                                            v60 := GenS6{F0: (v47 / (((^(v47))) | 1)), F1: 1, F2: (fi51 << ((4294967295) & 63))}
                                                            _ = v60
                                                        }
                                                        v4 += v4
                                                        if v44 {
                                                            if len(v3) > 0 {
                                                                v3[0] = ((((("a" + "foo bar") + (v15 + ("foo bar" + ((((((v15 + ((v15 + "hello") + "hello")) + "foo bar") + v15) + v15) + "") + "\n\t")))) + (("foo bar" + "") + (v15 + ("foo bar" + "a")))) + ((v15 + "") + v15)) + ((v15 + v15) + (v15 + v15)))
                                                            }
                                                            var v61 uint64 = 826
                                                            _ = v61
                                                        } else {
                                                            v62 := GenS7{F0: ""}
                                                            _ = v62
                                                        }
                                                        var v63 uint8 = 94
                                                        _ = v63
                                                    } else {
                                                        v16 -= v49
                                                    }
                                                default:
                                                    v64 := GenS8{F0: (v2 & v2), F1: 18446744073709551615, F2: fi50}
                                                    _ = v64
                                                }
                                                if ((fi50 << ((uint(v14)) & 63)) == (fi50 ^ fi50)) {
                                                    var v65 uint64 = 42
                                                    if v44 {
                                                        switch fi51 {
                                                        case -303:
                                                            var v66 float32 = v4
                                                            _ = v66
                                                        }
                                                        var v67 int = fi50
                                                        var v68 uint32 = v48
                                                        _ = v67
                                                        _ = v68
                                                    } else {
                                                        var v69 uint8 = 2
                                                        var v70 int32 = v2
                                                        switch fi50 {
                                                        case 27:
                                                            if (v4 == (v4 - (1e10 - (1.5 - 1e10)))) {
                                                                var v71 float32 = 0.0
                                                                _ = v71
                                                            } else {
                                                                v70 = 654
                                                            }
                                                        default:
                                                            v3 = append(v3, v15)
                                                        }
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                    v49 -= (v49 * 42)
                                                    _ = v65
                                                }
                                            }
                                        }
                                    } else {
                                        if false {
                                            if ((42 >= 679) || (-32768 != 584)) {
                                                v2 += (^(v2))
                                                v72 := GenS9{F0: (v15 + v15)}
                                                _ = v72
                                            } else {
                                                if (0 <= 463) {
                                                    var v73 int = (fi50 << ((2) & 63))
                                                    _ = v73
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if (-929 > v49) {
                                                    if (16 >= 0) {
                                                        if true {
                                                            if (v44 || true) {
                                                                v44 = false
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        var v74 bool = (-783 > fi50)
                                                        if (32767 <= fi51) {
                                                            var v75 int8 = v47
                                                            var v76 int32 = (v2 >> ((776) & 63))
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            if v44 {
                                                                v47 = (-(v47))
                                                                v4 = ((((float32(((((1e10 - (1e10 / 0.1)) / 1e10) / float64(v4)) + ((((1e10 / (0.0 * (float64(v4) / (1.5 + 0.0)))) / 1.5) / -2.25) + (1e10 - 1e10)))) / (float32(-2.25) / (float32(-2.25) / (((3.14159 - 3.14159) - ((3.14159 * (v4 * (v4 + (v4 * (0.0 + v4))))) + ((0.0 + (v4 + float32(0.1))) - 1.5))) * v4)))) * (v4 * (v4 / v4))) * ((0.0 / v4) - v4)) * v4)
                                                                if (v2 != v2) {
                                                                    if (v44 && true) {
                                                                        if (0 <= 42) {
                                                                            if (!((v16 * v49) != v16)) {
                                                                                switch (fi50 << ((413) & 63)) {
                                                                                case 128:
                                                                                    _ = len(v1)
                                                                                    v48 = v48
                                                                                    fmt.Println(v47, v15)
                                                                                default:
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        _ = len(v3)
                                                                                    case "":
                                                                                        v44 = v44
                                                                                        v15 = v15
                                                                                    case "foo bar":
                                                                                        fmt.Println(v48)
                                                                                    default:
                                                                                        v4 -= 0.1
                                                                                    }
                                                                                    var v77 int = (fi51 - (fi51 << ((genFn0(v47)) & 63)))
                                                                                    if (!v44) {
                                                                                        v49 += 635
                                                                                        switch 32768 {
                                                                                        case 9223372036854775807:
                                                                                            if v44 {
                                                                                                var v78 int64 = v49
                                                                                                v79 := GenS10{F0: 721}
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                        case 32767:
                                                                                            delete(v1, (v15 + "\n\t"))
                                                                                        case 128:
                                                                                            v77 = fi51
                                                                                        default:
                                                                                            v49 = (v49 & (v49 % ((v16) | 1)))
                                                                                        }
                                                                                    }
                                                                                    _ = v77
                                                                                }
                                                                            }
                                                                            var v80 uint16 = 80
                                                                            _ = v80
                                                                        } else {
                                                                            var v81 float32 = ((v4 * v4) * (v4 * (v4 - -2.25)))
                                                                            v2 = 42
                                                                            var v82 uint64 = 18446744073709551615
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        delete(v1, v15)
                                                                        v49 = (v49 * (^(v16)))
                                                                        switch fi51 {
                                                                        case 146:
                                                                            fmt.Println(v2)
                                                                        default:
                                                                            v83 := GenS11{F0: (v16 + v49)}
                                                                            _ = v83
                                                                        }
                                                                    }
                                                                } else {
                                                                    _ = len(v3)
                                                                    var v84 uint64 = 0
                                                                    fmt.Println(v49)
                                                                    _ = v84
                                                                }
                                                            }
                                                            if len(v3) > 0 {
                                                                v3[0] = ("hello" + v15)
                                                            }
                                                        }
                                                        _ = v74
                                                    }
                                                }
                                            }
                                        } else {
                                            switch ((v15 + v15) + "\n\t") {
                                            case "":
                                                fmt.Println(v16)
                                            case "a":
                                                if (!(!true)) {
                                                    if false {
                                                        var v85 bool = true
                                                        v2 = (^(v2))
                                                        var v86 int8 = v47
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                } else {
                                                    v87 := GenS9{F0: "\n\t"}
                                                    var v88 uint8 = uint8(v48)
                                                    _ = v87
                                                    _ = v88
                                                }
                                                if ((297 >= 42) && (537 > 708)) {
                                                    delete(v1, "foo bar")
                                                    fmt.Println(v15, v4, v44)
                                                    if true {
                                                        v89 := GenS11{F0: (v49 >> ((genFn0((v47 ^ (v47 % ((-128) | 1))))) & 63))}
                                                        fmt.Println(v14, fi50)
                                                        _ = v89
                                                    }
                                                } else {
                                                    var v90 uint16 = 489
                                                    if (462 != 2) {
                                                        if (v44 || false) {
                                                            v47 = v47
                                                            var v91 int8 = 112
                                                            v92 := GenS12{F0: 0, F1: 127}
                                                            _ = v91
                                                            _ = v92
                                                        } else {
                                                            switch (fi51 % ((fi51) | 1)) {
                                                            case 4294967295:
                                                                if ((!((v14 - v90) <= v14)) || (!true)) {
                                                                    if ((^(v90)) < v14) {
                                                                        v4 += v4
                                                                        switch (v15 + v15) {
                                                                        case "\n\t":
                                                                            fmt.Println(v48)
                                                                        }
                                                                        v3 = append(v3, ("hello" + ("" + v15)))
                                                                    } else {
                                                                        _ = len(v3)
                                                                    }
                                                                    var v93 uint = 255
                                                                    v14 = (v14 % (((v90 / ((v14) | 1))) | 1))
                                                                    _ = v93
                                                                } else {
                                                                    if (3.14159 < (0.1 - (1.5 / 3.14159))) {
                                                                        var v94 int = int(v16)
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v95 := GenS5{F0: (fi50 | -717), F1: -315}
                                                                            _ = v95
                                                                        case "\n\t":
                                                                            v15 = (v15 + "hello")
                                                                            switch v15 {
                                                                            case "a":
                                                                                v4 = v4
                                                                            case "\n\t":
                                                                                v47 = (v47 & (-(v47)))
                                                                            case "hello":
                                                                                var v96 bool = v44
                                                                                _ = v96
                                                                            }
                                                                        default:
                                                                            fmt.Println(v48, v90)
                                                                        }
                                                                        _ = v94
                                                                    } else {
                                                                        var v97 int8 = -80
                                                                        v98 := GenS13{F0: 4294967295, F1: genFn0((^(v47))), F2: 1}
                                                                        _ = v97
                                                                        _ = v98
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v15 + (v15 + (((((v15 + v15) + v15) + "foo bar") + "a") + v15)))
                                                                    }
                                                                }
                                                            case 863:
                                                                fmt.Println(v47)
                                                            case 754:
                                                                var v99 uint8 = 0
                                                                _ = v99
                                                            default:
                                                                v14 = 32767
                                                            }
                                                        }
                                                        var v100 uint64 = 0
                                                        _ = v100
                                                    }
                                                    fmt.Println(v49, v48)
                                                    _ = v90
                                                }
                                            }
                                            var v101 int8 = (v47 >> ((0) & 63))
                                            _ = v101
                                        }
                                        switch v15 {
                                        case "\n\t":
                                            fmt.Println(v4, v49, v2)
                                        case "":
                                            switch fi51 {
                                            case 128:
                                                var v102 int32 = int32(v47)
                                                _ = v102
                                            case -50:
                                                _ = len(v3)
                                                v47 += v47
                                            case -1:
                                                fmt.Println(fi50)
                                            default:
                                                if false {
                                                    var v103 int32 = 2147483647
                                                    v104 := GenS6{F0: v47, F1: 42, F2: 127}
                                                    _ = v103
                                                    _ = v104
                                                }
                                            }
                                        default:
                                            if (42 == 80) {
                                                v47 = v47
                                                fmt.Println(v16, v4, v4)
                                            }
                                            v105 := GenS13{F0: 1, F1: 294, F2: int16(fi51)}
                                            switch fi50 {
                                            case -2147483648:
                                                fmt.Println(v14)
                                            case -1:
                                                switch (v15 + v15) {
                                                case "a":
                                                    fmt.Println(v4, v44)
                                                case "\n\t":
                                                    v2 += (v2 & -2147483648)
                                                case "":
                                                    if v44 {
                                                        var v106 int64 = (v49 * v16)
                                                        v107 := GenS7{F0: v15}
                                                        fmt.Println(fi50, v47)
                                                        var v108 uint = 18446744073709551615
                                                        _ = v106
                                                        _ = v107
                                                        _ = v108
                                                    } else {
                                                        if true {
                                                            v3 = append(v3, "")
                                                        }
                                                        _ = len(v1)
                                                        var v109 uint16 = v14
                                                        _ = v109
                                                    }
                                                }
                                            }
                                            _ = v105
                                        }
                                    }
                                    _ = fi51
                                }
                            }
                            switch v15 {
                            case "a":
                                fmt.Println(v16)
                            default:
                                delete(v1, ("\n\t" + ((((v15 + "\n\t") + v15) + v15) + "hello")))
                            }
                            for fi110 := 0; fi110 < 17; fi110++ {
                                fmt.Println(v4, v14)
                                _ = fi110
                            }
                            _ = fi50
                        }
                    }
                    _ = v48
                    _ = v49
                } else {
                    var v111 float64 = (((1e10 * 1e10) - 0.1) * (3.14159 * 1.5))
                    switch 16 {
                    case 2147483647:
                        fmt.Println(v15)
                    case 834:
                        v3 = append(v3, v15)
                    default:
                        v112 := v1[v15]
                        _ = v112
                    }
                    fmt.Println(v16, v47, v111)
                    _ = v111
                }
                v113 := GenS13{F0: 0, F1: 4294967295, F2: 1}
                _ = v113
            } else {
                var v114 uint64 = 42
                v114 += (^(v114))
                v115 := GenS10{F0: 1}
                _ = v114
                _ = v115
            }
            _ = v46
            _ = v47
        }
        v116 := GenS9{F0: v15}
        _ = v116
    } else {
        for fi117 := 0; fi117 < 5; fi117++ {
            v44 = false
            switch int(v2) {
            case 4294967295:
                if v44 {
                    var v118 bool = (!v44)
                    if v44 {
                        var v119 uint64 = uint64(v14)
                        for fi120 := 0; fi120 < 20; fi120++ {
                            var v121 float32 = v4
                            var v122 uint8 = 2
                            var v123 uint16 = v14
                            v118 = (-128 > 0)
                            fmt.Println(fi117, v121, v121)
                            _ = fi120
                            _ = v121
                            _ = v122
                            _ = v123
                        }
                        _ = v119
                    }
                    _ = v118
                } else {
                    v15 = ((v15 + v15) + v15)
                }
            case 32767:
                for fi124 := 0; fi124 < 6; fi124++ {
                    switch (-(fi117)) {
                    case -2147483648:
                        var v125 bool = v44
                        _ = v125
                    case 32767:
                        switch (fi124 * -1) {
                        case 127:
                            var v126 int64 = v16
                            _ = v126
                        case -902:
                            v127 := GenS10{F0: 288}
                            _ = v127
                        case 9223372036854775807:
                            v128 := GenS14{F0: v14, F1: 0}
                            _ = v128
                        }
                    case 42:
                        v129 := GenS15{F0: (v16 ^ (v16 >> ((42) & 63))), F1: 1}
                        _ = v129
                    }
                    var v130 uint32 = 1
                    fmt.Println(v130, v44)
                    _ = fi124
                    _ = v130
                }
            case 9223372036854775807:
                v44 = (v44 || v44)
            default:
                var v131 int8 = 100
                _ = v131
            }
            v44 = (!v44)
            _ = fi117
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
