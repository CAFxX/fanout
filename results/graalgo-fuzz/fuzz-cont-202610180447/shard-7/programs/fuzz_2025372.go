package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint }
type GenS3 struct { F0 int8; F1 int32; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 string; F1 int16; F2 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int; F1 uint; F2 string }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 int64; F1 uint64 }
type GenS20 struct { F0 uint; F1 string }
type GenS21 struct { F0 int16; F1 int; F2 uint8 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 uint8; F1 int64 }
type GenS24 struct { F0 int32; F1 int16; F2 int64 }
type GenS25 struct { F0 uint64; F1 int8 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 int8; F1 int16; F2 int16 }
type GenS28 struct { F0 uint16; F1 int16 }
type GenS29 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS30 struct { F0 uint32; F1 uint }
type GenS31 struct { F0 uint32; F1 int }
type GenS32 struct { F0 int16 }
type GenS33 struct { F0 uint64 }
type GenS34 struct { F0 int; F1 int64; F2 int16 }
type GenS35 struct { F0 uint8; F1 bool }
type GenS36 struct { F0 uint16; F1 bool; F2 int }
type GenS37 struct { F0 int16; F1 uint }
type GenS38 struct { F0 int }
type GenS39 struct { F0 string; F1 string; F2 uint16 }
type GenS40 struct { F0 uint; F1 int32; F2 int16 }
type GenS41 struct { F0 int16 }
func genFn0(p0 int64) int8 {
    p0 = p0
    p0 += (p0 & (p0 | (p0 ^ (-(p0)))))
    _ = p0
    return 42
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{"a"}
    var v3 int32 = 1
    v4 := make(map[string]uint32)
    v5 := []uint32{0, 0}
    var v6 bool = ((true && false) && false)
    fmt.Println(v3, v3, v3)
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            fmt.Println(v1, fi8)
            if (("\n\t" + (("a" + "") + "")) > "hello") {
                v1 += ((float32(((0.0 * 0.1) + 0.1)) / float32((1e10 / 1.5))) * v1)
            }
            fmt.Println(v6)
            _ = fi8
        }
        v9 := GenS1{F0: 42}
        var v10 int8 = 1
        _ = fi7
        _ = v10
        _ = v9
    }
    if (0 >= 18446744073709551615) {
        switch 4294967295 {
        case -9223372036854775808:
            v11 := GenS2{F0: 100, F1: 394, F2: 4294967295}
            _ = v11
        default:
            var v12 int16 = 1000
            _ = v12
        }
        v6 = v6
        v6 = v6
    }
    v3 += 0
    switch 42 {
    case 127:
        if v6 {
            v13 := GenS3{F0: 13, F1: (^(v3)), F2: (4294967295 > 2)}
            var v14 int32 = 42
            _ = v13
            _ = v14
        }
    }
    v3 += (v3 >> ((65535) & 63))
    var v15 uint = 42
    v16 := GenS2{F0: genFn0(795), F1: 4294967295, F2: v15}
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 int = (fi17 >> ((v15) & 63))
        v4[("hello" + (("hello" + (("a" + ("\n\t" + "hello")) + "")) + ""))] = 2
        if v6 {
            v1 -= (v1 + (0.0 - (v1 + v1)))
            fmt.Println(v1, v6, v1)
            var v19 bool = v6
            _ = v19
        } else {
            switch "a" {
            case "hello":
                var v20 uint32 = 2147483648
                _ = v20
            case "":
                switch "hello" {
                case "foo bar":
                    v1 = v1
                default:
                    for fi21 := 0; fi21 < 6; fi21++ {
                        if true {
                            var v22 int16 = 488
                            _ = v22
                        } else {
                            if v6 {
                                fmt.Println(v15, fi17, v15)
                            } else {
                                v23 := GenS3{F0: -37, F1: -1, F2: (0 == 2)}
                                _ = v23
                            }
                            if v6 {
                                switch int(v15) {
                                case 127:
                                    if (v3 != (v3 & v3)) {
                                        switch fi21 {
                                        case 0:
                                            fmt.Println(v3)
                                        }
                                        v24 := GenS2{F0: 15, F1: 60, F2: v15}
                                        v6 = (!(3.14159 == ((v1 * ((v1 * 0.1) * v1)) * 1e10)))
                                        _ = v24
                                    } else {
                                        v25 := GenS1{F0: (v3 | (v3 + v3))}
                                        var v26 bool = v6
                                        v18 = (v18 / ((-2147483648) | 1))
                                        _ = v25
                                        _ = v26
                                    }
                                    var v27 int = (fi21 + (^(v18)))
                                    if false {
                                        if v6 {
                                            var v28 int16 = 1000
                                            _ = v28
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v29 := v4["foo bar"]
                                            switch (fi21 >> (((v15 << (((^(v15))) & 63))) & 63)) {
                                            case 701:
                                                v30 := GenS2{F0: -1, F1: 2147483647, F2: v15}
                                                _ = v30
                                            case -1:
                                                if (((1 == uint16(v15)) || v6) || (-32768 > 1)) {
                                                    if v6 {
                                                        v31 := GenS3{F0: 0, F1: int32(v15), F2: v6}
                                                        fmt.Println(fi17, v6, v1)
                                                        var v32 int32 = (v3 * (v3 / ((327) | 1)))
                                                        _ = v31
                                                        _ = v32
                                                    } else {
                                                        if (!(!(9223372036854775807 <= 0))) {
                                                            v18 += (fi21 - (-(v18)))
                                                        }
                                                        if true {
                                                            v33 := GenS4{F0: 32767}
                                                            var v34 int = fi21
                                                            v18 = fi17
                                                            _ = v33
                                                            _ = v34
                                                        } else {
                                                            v2 = append(v2, "")
                                                            v35 := GenS3{F0: 1, F1: (v3 / ((1) | 1)), F2: ((v3 > v3) && (!(v6 || v6)))}
                                                            if (!(!(true && ((^(v15)) == v15)))) {
                                                                var v36 uint8 = 42
                                                                if v6 {
                                                                    switch -9223372036854775808 {
                                                                    case 9223372036854775807:
                                                                        v37 := GenS1{F0: v3}
                                                                        _ = v37
                                                                    }
                                                                    var v38 int8 = -70
                                                                    var v39 uint32 = (v29 & v29)
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            var v40 uint8 = (v36 * (v36 | v36))
                                                                            if (-838 <= 42) {
                                                                                delete(v4, "\n\t")
                                                                                var v41 int8 = 1
                                                                                _ = v41
                                                                            }
                                                                            v3 = 2147483647
                                                                            _ = v40
                                                                        }
                                                                        v42 := GenS1{F0: 1}
                                                                        if len(v5) > 0 {
                                                                            v5[0] = (v29 | 2147483647)
                                                                        }
                                                                        v43 := GenS5{F0: (v36 >> ((48) & 63))}
                                                                        var v44 string = "hello"
                                                                        switch ((v44 + v44) + ((v44 + "\n\t") + v44)) {
                                                                        case "hello":
                                                                            _ = len(v4)
                                                                        case "a":
                                                                            v6 = v6
                                                                        default:
                                                                            v1 -= ((float32(-2.25) / v1) - (v1 + v1))
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v29
                                                                            }
                                                                            switch (("a" + (v44 + "foo bar")) + ("" + v44)) {
                                                                            case "hello":
                                                                                var v45 bool = v6
                                                                                _ = v45
                                                                            case "foo bar":
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "hello"
                                                                                }
                                                                            case "\n\t":
                                                                                var v46 uint64 = 2
                                                                                _ = v46
                                                                            }
                                                                        }
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        var v47 int32 = int32(v36)
                                                                        if false {
                                                                            fmt.Println(v3, v6)
                                                                        } else {
                                                                            var v48 uint16 = 42
                                                                            fmt.Println(v18)
                                                                            v27 = (fi21 / ((-9223372036854775808) | 1))
                                                                            _ = v48
                                                                        }
                                                                        if v6 {
                                                                            var v49 int8 = 0
                                                                            _ = v49
                                                                        }
                                                                        if ((!(!(!true))) && (0.0 == float64((((3.14159 * (v1 / float32((0.0 - 0.1)))) / (v1 + v1)) / v1)))) {
                                                                            if v6 {
                                                                                _ = len(v2)
                                                                                v5 = append(v5, (v29 & 246))
                                                                            }
                                                                            v47 = v47
                                                                            fmt.Println(v15, v3, v6)
                                                                            fmt.Println(fi17, v3)
                                                                        } else {
                                                                            if ((0.0 / (-2.25 / (float64(v1) - ((float64((v1 - 0.0)) / 0.1) + 1e10)))) != 1e10) {
                                                                                var v50 int32 = (v3 + v47)
                                                                                switch (^(fi21)) {
                                                                                case -396:
                                                                                    v27 = (fi17 & fi17)
                                                                                case 42:
                                                                                    v3 -= v50
                                                                                }
                                                                                v50 -= (-(v47))
                                                                                var v51 uint16 = 1
                                                                                _ = v50
                                                                                _ = v51
                                                                            } else {
                                                                                var v52 uint = 65535
                                                                                fmt.Println(v1, v29)
                                                                                v53 := GenS2{F0: 1, F1: v29, F2: v15}
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                        }
                                                                        if v6 {
                                                                            var v54 int32 = v3
                                                                            fmt.Println(v27, v3)
                                                                            var v55 uint64 = 18446744073709551615
                                                                            _ = v54
                                                                            _ = v55
                                                                        } else {
                                                                            var v56 uint32 = (v29 & (v29 / ((v29) | 1)))
                                                                            var v57 bool = v6
                                                                            if true {
                                                                                if v6 {
                                                                                    if (v6 || (-9223372036854775808 >= 0)) {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        if (32767 > -1) {
                                                                                            _ = len(v4)
                                                                                            v58 := GenS6{F0: 44, F1: 2}
                                                                                            switch (v27 << ((v15) & 63)) {
                                                                                            case -2147483648:
                                                                                                v59 := GenS7{F0: 255, F1: true}
                                                                                                v4[(("hello" + (("\n\t" + "foo bar") + "foo bar")) + ("" + "hello"))] = (v29 ^ v29)
                                                                                                _ = v59
                                                                                            case 848:
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    var v60 uint16 = 1
                                                                                                    v61 := GenS1{F0: -2147483648}
                                                                                                    v57 = v6
                                                                                                    _ = v60
                                                                                                    _ = v61
                                                                                                case "":
                                                                                                    v62 := GenS2{F0: 55, F1: v56, F2: (v15 << (((v15 / (((^(v15))) | 1))) & 63))}
                                                                                                    v1 -= float32(float64((float32(1.5) / v1)))
                                                                                                    var v63 uint8 = 2
                                                                                                    switch "" {
                                                                                                    case "":
                                                                                                        v64 := GenS8{F0: v63}
                                                                                                        _ = v64
                                                                                                    default:
                                                                                                        switch ("hello" + "hello") {
                                                                                                        case "":
                                                                                                            var v65 float32 = (1.5 * 0.0)
                                                                                                            _ = v65
                                                                                                        case "foo bar":
                                                                                                            if (!(false || v57)) {
                                                                                                                v66 := v4[(("foo bar" + "\n\t") + "a")]
                                                                                                                _ = v66
                                                                                                            } else {
                                                                                                                v67 := GenS9{F0: 42, F1: v47}
                                                                                                                if (32767 < -32768) {
                                                                                                                    var v68 string = ""
                                                                                                                    fmt.Println(v15)
                                                                                                                    if (!true) {
                                                                                                                        var v69 int = (fi17 % ((int(v29)) | 1))
                                                                                                                        switch fi21 {
                                                                                                                        case -1:
                                                                                                                            var v70 uint8 = 0
                                                                                                                            _ = v70
                                                                                                                        }
                                                                                                                        var v71 bool = v6
                                                                                                                        _ = v69
                                                                                                                        _ = v71
                                                                                                                    }
                                                                                                                    _ = v68
                                                                                                                } else {
                                                                                                                    var v72 int64 = 42
                                                                                                                    v15 = (v15 >> ((4294967295) & 63))
                                                                                                                    v18 = (v18 % ((4294967295) | 1))
                                                                                                                    _ = v72
                                                                                                                }
                                                                                                                v36 = (v63 | (v63 ^ (^(v63))))
                                                                                                                _ = v67
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v73 string = "\n\t"
                                                                                                            _ = v73
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                case "\n\t":
                                                                                                    var v74 int32 = 0
                                                                                                    _ = v74
                                                                                                default:
                                                                                                    v6 = true
                                                                                                }
                                                                                            default:
                                                                                                if (v18 > (v27 & fi17)) {
                                                                                                    v75 := v4[("foo bar" + "\n\t")]
                                                                                                    if v57 {
                                                                                                        if true {
                                                                                                            if v6 {
                                                                                                                v76 := GenS10{F0: ("foo bar" + "hello"), F1: 32767, F2: 100}
                                                                                                                if len(v5) > 0 {
                                                                                                                    v5[0] = v29
                                                                                                                }
                                                                                                                _ = v76
                                                                                                            }
                                                                                                            v36 = (v36 >> ((4294967295) & 63))
                                                                                                        } else {
                                                                                                            v77 := v4["a"]
                                                                                                            switch fi17 {
                                                                                                            case 127:
                                                                                                                v78 := v4["\n\t"]
                                                                                                                _ = v78
                                                                                                            case -441:
                                                                                                                v79 := GenS3{F0: 1, F1: -2147483648, F2: v57}
                                                                                                                _ = v79
                                                                                                            }
                                                                                                            switch "a" {
                                                                                                            case "foo bar":
                                                                                                                _ = len(v4)
                                                                                                            case "hello":
                                                                                                                if v57 {
                                                                                                                    v4["hello"] = 0
                                                                                                                    var v80 uint8 = 127
                                                                                                                    var v81 uint = (v15 + (v15 % (((v15 % (((v15 * (v15 * v15))) | 1))) | 1)))
                                                                                                                    _ = v80
                                                                                                                    _ = v81
                                                                                                                } else {
                                                                                                                    v82 := GenS4{F0: 863}
                                                                                                                    _ = v82
                                                                                                                }
                                                                                                                var v83 float32 = v1
                                                                                                                _ = v83
                                                                                                            }
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        var v84 int16 = -312
                                                                                                        _ = v84
                                                                                                    } else {
                                                                                                        v85 := GenS11{F0: 12}
                                                                                                        var v86 int8 = 14
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    }
                                                                                                    fmt.Println(v15, v29, v15)
                                                                                                    _ = v75
                                                                                                }
                                                                                            }
                                                                                            _ = v58
                                                                                        } else {
                                                                                            delete(v4, "foo bar")
                                                                                            var v87 uint = (v15 + (v15 * v15))
                                                                                            fmt.Println(v1, v3)
                                                                                            v6 = ((v56 - (v29 << ((432) & 63))) > (v56 | v56))
                                                                                            _ = v87
                                                                                        }
                                                                                    } else {
                                                                                        v36 += v36
                                                                                    }
                                                                                    if false {
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            v88 := GenS12{F0: 18446744073709551615}
                                                                                            _ = v88
                                                                                        default:
                                                                                            v27 += (fi17 / (((v27 / ((v18) | 1))) | 1))
                                                                                        }
                                                                                        v47 = v47
                                                                                        if (false || (v57 || ("foo bar" == ("a" + "\n\t")))) {
                                                                                            if (!(v6 || (true && (v57 && v6)))) {
                                                                                                var v89 uint16 = 42
                                                                                                _ = v89
                                                                                            } else {
                                                                                                var v90 uint64 = 952
                                                                                                _ = v90
                                                                                            }
                                                                                            if (v6 || v57) {
                                                                                                if (v57 || false) {
                                                                                                    var v91 int16 = 844
                                                                                                    fmt.Println(v1)
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    if (false && (!v6)) {
                                                                                                        if v57 {
                                                                                                            fmt.Println(v3)
                                                                                                            v92 := GenS13{F0: 326}
                                                                                                            var v93 int32 = v47
                                                                                                            var v94 float32 = 1.5
                                                                                                            _ = v92
                                                                                                            _ = v93
                                                                                                            _ = v94
                                                                                                        } else {
                                                                                                            fmt.Println(fi21, v3)
                                                                                                        }
                                                                                                        if true {
                                                                                                            switch fi17 {
                                                                                                            case 1:
                                                                                                                v4["\n\t"] = (^(v29))
                                                                                                            case 127:
                                                                                                                if (2 <= 909) {
                                                                                                                    v56 += uint32(v36)
                                                                                                                    delete(v4, "a")
                                                                                                                    switch "a" {
                                                                                                                    case "\n\t":
                                                                                                                        v18 = fi17
                                                                                                                    case "":
                                                                                                                        switch "a" {
                                                                                                                        case "hello":
                                                                                                                            var v95 int32 = (^(v47))
                                                                                                                            _ = v95
                                                                                                                        default:
                                                                                                                            v96 := GenS9{F0: 100, F1: v3}
                                                                                                                            _ = v96
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        v97 := GenS8{F0: uint8(v15)}
                                                                                                                        v36 += v36
                                                                                                                        _ = v97
                                                                                                                    default:
                                                                                                                        switch "hello" {
                                                                                                                        case "\n\t":
                                                                                                                            if v6 {
                                                                                                                                if (((v15 & (^(v15))) <= 4294967295) && v57) {
                                                                                                                                    v98 := GenS8{F0: v36}
                                                                                                                                    v1 = float32(3.14159)
                                                                                                                                    _ = v98
                                                                                                                                } else {
                                                                                                                                    var v99 int = (v18 >> (((v15 ^ 1)) & 63))
                                                                                                                                    v15 = 535
                                                                                                                                    v100 := GenS8{F0: 66}
                                                                                                                                    _ = v100
                                                                                                                                    _ = v99
                                                                                                                                }
                                                                                                                                var v101 uint16 = 0
                                                                                                                                v101 -= 575
                                                                                                                                _ = v101
                                                                                                                            } else {
                                                                                                                                v56 += 2147483648
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            v15 -= v15
                                                                                                                        case "foo bar":
                                                                                                                            var v102 int32 = (v47 + (^(v47)))
                                                                                                                            _ = v102
                                                                                                                        default:
                                                                                                                            var v103 string = ((("" + ("a" + (((("foo bar" + ("" + "a")) + "\n\t") + (("\n\t" + (("\n\t" + "") + ("hello" + "a"))) + ("foo bar" + "\n\t"))) + "\n\t"))) + "hello") + "\n\t")
                                                                                                                            _ = v103
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v104 := GenS14{F0: (v18 % ((-9223372036854775808) | 1)), F1: (v15 * (^(v15))), F2: "hello"}
                                                                                                                    if ("" != "") {
                                                                                                                        switch (fi17 | (fi17 & 32767)) {
                                                                                                                        case 1:
                                                                                                                            var v105 uint32 = v29
                                                                                                                            _ = v105
                                                                                                                        case 32768:
                                                                                                                            if (32767 >= 681) {
                                                                                                                                var v106 string = "a"
                                                                                                                                if ((!true) || (!((^(fi21)) > (v18 ^ v27)))) {
                                                                                                                                    v107 := GenS10{F0: v106, F1: 664, F2: 42}
                                                                                                                                    if v57 {
                                                                                                                                        v108 := v4["foo bar"]
                                                                                                                                        var v109 uint16 = 570
                                                                                                                                        _ = v108
                                                                                                                                        _ = v109
                                                                                                                                    } else {
                                                                                                                                        v110 := GenS15{F0: (true && true)}
                                                                                                                                        v15 = (v15 >> (((v15 - v15)) & 63))
                                                                                                                                        _ = v110
                                                                                                                                    }
                                                                                                                                    v111 := GenS10{F0: "foo bar", F1: 217, F2: 0}
                                                                                                                                    _ = v107
                                                                                                                                    _ = v111
                                                                                                                                }
                                                                                                                                var v112 int16 = -32768
                                                                                                                                _ = v106
                                                                                                                                _ = v112
                                                                                                                            } else {
                                                                                                                                v57 = ("a" >= "foo bar")
                                                                                                                            }
                                                                                                                        case -850:
                                                                                                                            if ("\n\t" < "hello") {
                                                                                                                                v113 := GenS16{F0: 516}
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        if (v18 == 2147483647) {
                                                                                                                                            delete(v4, "foo bar")
                                                                                                                                        }
                                                                                                                                        if true {
                                                                                                                                            v15 += (v15 << (((^(v15))) & 63))
                                                                                                                                            v15 += v15
                                                                                                                                        } else {
                                                                                                                                            v18 += fi21
                                                                                                                                            v27 = (^(v27))
                                                                                                                                            fmt.Println(v47, v18, v57)
                                                                                                                                        }
                                                                                                                                        switch (v18 * (fi17 + 425)) {
                                                                                                                                        case 0:
                                                                                                                                            var v114 uint32 = v56
                                                                                                                                            _ = v114
                                                                                                                                        default:
                                                                                                                                            switch v27 {
                                                                                                                                            case 0:
                                                                                                                                                v36 -= (v36 + v36)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v115 := GenS6{F0: 42, F1: v36}
                                                                                                                                        if v57 {
                                                                                                                                            v27 = (fi21 + 128)
                                                                                                                                            v2 = append(v2, "\n\t")
                                                                                                                                            var v116 uint8 = 255
                                                                                                                                            _ = v116
                                                                                                                                        }
                                                                                                                                        _ = v115
                                                                                                                                    }
                                                                                                                                    if len(v5) > 0 {
                                                                                                                                        v5[0] = v29
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    v117 := GenS17{F0: 1}
                                                                                                                                    var v118 bool = true
                                                                                                                                    var v119 bool = (!v118)
                                                                                                                                    v4["a"] = v56
                                                                                                                                    if ((v36 + v36) > v36) {
                                                                                                                                        var v120 float32 = v1
                                                                                                                                        var v121 bool = (!false)
                                                                                                                                        var v122 int8 = genFn0(9223372036854775807)
                                                                                                                                        _ = v120
                                                                                                                                        _ = v121
                                                                                                                                        _ = v122
                                                                                                                                    }
                                                                                                                                    var v123 int32 = 0
                                                                                                                                    if v118 {
                                                                                                                                        if v119 {
                                                                                                                                            v56 = v29
                                                                                                                                            if false {
                                                                                                                                                v124 := GenS17{F0: 1}
                                                                                                                                                var v125 uint = v15
                                                                                                                                                _ = v124
                                                                                                                                                _ = v125
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            var v126 int32 = (v123 * v123)
                                                                                                                                            var v127 uint = 255
                                                                                                                                            switch (("\n\t" + "hello") + "") {
                                                                                                                                            case "":
                                                                                                                                                var v128 uint = (v15 + (^(v127)))
                                                                                                                                                _ = v128
                                                                                                                                            case "hello":
                                                                                                                                                fmt.Println(v18)
                                                                                                                                            case "a":
                                                                                                                                                switch "foo bar" {
                                                                                                                                                case "foo bar":
                                                                                                                                                    v129 := GenS18{F0: 42}
                                                                                                                                                    _ = v129
                                                                                                                                                case "":
                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                    }
                                                                                                                                                case "hello":
                                                                                                                                                    var v130 string = "foo bar"
                                                                                                                                                    var v131 int32 = 829
                                                                                                                                                    _ = len(v4)
                                                                                                                                                    switch v130 {
                                                                                                                                                    case "":
                                                                                                                                                        if false {
                                                                                                                                                            fmt.Println(v29)
                                                                                                                                                            if true {
                                                                                                                                                                v132 := GenS3{F0: genFn0(-9223372036854775808), F1: v126, F2: v118}
                                                                                                                                                                if (v56 != (v56 + (v56 % ((442) | 1)))) {
                                                                                                                                                                    fmt.Println(fi21, v27, v127)
                                                                                                                                                                } else {
                                                                                                                                                                    fmt.Println(v36)
                                                                                                                                                                    v27 -= (fi21 | (fi21 * (v27 / ((v27) | 1))))
                                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                v18 -= (-(v18))
                                                                                                                                                                _ = v132
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            var v133 int8 = 42
                                                                                                                                                            _ = v133
                                                                                                                                                        }
                                                                                                                                                        v134 := v4[v130]
                                                                                                                                                        v119 = (v36 >= v36)
                                                                                                                                                        v29 = v134
                                                                                                                                                        v36 = (^(v36))
                                                                                                                                                        var v135 int64 = 42
                                                                                                                                                        _ = v134
                                                                                                                                                        _ = v135
                                                                                                                                                    case "a":
                                                                                                                                                        v1 -= 3.14159
                                                                                                                                                        switch v130 {
                                                                                                                                                        case "foo bar":
                                                                                                                                                            if true {
                                                                                                                                                                var v136 float32 = 3.14159
                                                                                                                                                                _ = v136
                                                                                                                                                            } else {
                                                                                                                                                                v36 -= 238
                                                                                                                                                            }
                                                                                                                                                        case "":
                                                                                                                                                            v119 = (212 <= (^(v36)))
                                                                                                                                                        default:
                                                                                                                                                            v4["\n\t"] = (v56 << ((v15) & 63))
                                                                                                                                                        }
                                                                                                                                                        v137 := GenS2{F0: 96, F1: v29, F2: uint(v56)}
                                                                                                                                                        _ = v137
                                                                                                                                                    default:
                                                                                                                                                        switch fi21 {
                                                                                                                                                        case 4294967295:
                                                                                                                                                            _ = len(v4)
                                                                                                                                                        case 484:
                                                                                                                                                            if v118 {
                                                                                                                                                                if true {
                                                                                                                                                                    if v118 {
                                                                                                                                                                        switch v130 {
                                                                                                                                                                        case "":
                                                                                                                                                                            var v138 float32 = 0.1
                                                                                                                                                                            _ = v138
                                                                                                                                                                        default:
                                                                                                                                                                            var v139 string = "foo bar"
                                                                                                                                                                            v140 := v4[v130]
                                                                                                                                                                            _ = v139
                                                                                                                                                                            _ = v140
                                                                                                                                                                        }
                                                                                                                                                                        v2 = append(v2, "")
                                                                                                                                                                    } else {
                                                                                                                                                                        var v141 uint16 = 255
                                                                                                                                                                        _ = v141
                                                                                                                                                                    }
                                                                                                                                                                    switch fi21 {
                                                                                                                                                                    case 32767:
                                                                                                                                                                        var v142 float32 = (((((v1 + 1.5) * v1) / (0.0 - v1)) - (v1 * (0.0 + 1.5))) - (float32((((((((0.1 + 1.5) * 0.1) * 0.0) + 0.0) - -2.25) + 0.1) + (1e10 + 1.5))) * 0.0))
                                                                                                                                                                        _ = v142
                                                                                                                                                                    case 9223372036854775807:
                                                                                                                                                                        var v143 uint16 = 266
                                                                                                                                                                        _ = v143
                                                                                                                                                                    case 127:
                                                                                                                                                                        if (v56 <= 4294967295) {
                                                                                                                                                                            fmt.Println(v131, v126, v47)
                                                                                                                                                                            if v118 {
                                                                                                                                                                                fmt.Println(v29)
                                                                                                                                                                            } else {
                                                                                                                                                                                fmt.Println(v56)
                                                                                                                                                                                v144 := GenS11{F0: 127}
                                                                                                                                                                                _ = v144
                                                                                                                                                                            }
                                                                                                                                                                            var v145 int = 32768
                                                                                                                                                                            v146 := GenS19{F0: 97, F1: 1}
                                                                                                                                                                            _ = v145
                                                                                                                                                                            _ = v146
                                                                                                                                                                        } else {
                                                                                                                                                                            var v147 int = v18
                                                                                                                                                                            v127 -= 596
                                                                                                                                                                            if len(v2) >= 2 {
                                                                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                                                                            }
                                                                                                                                                                            _ = v147
                                                                                                                                                                        }
                                                                                                                                                                        if (583 != -32768) {
                                                                                                                                                                            if len(v2) > 0 {
                                                                                                                                                                                v2[0] = v130
                                                                                                                                                                            }
                                                                                                                                                                            v148 := GenS9{F0: 100, F1: v47}
                                                                                                                                                                            v118 = (1 >= 509)
                                                                                                                                                                            _ = v148
                                                                                                                                                                        } else {
                                                                                                                                                                            v15 = 18446744073709551615
                                                                                                                                                                            var v149 string = ""
                                                                                                                                                                            _ = v149
                                                                                                                                                                        }
                                                                                                                                                                        var v150 int8 = int8(v56)
                                                                                                                                                                        _ = v150
                                                                                                                                                                    default:
                                                                                                                                                                        var v151 uint16 = 935
                                                                                                                                                                        _ = v151
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                if v6 {
                                                                                                                                                                    fmt.Println(v1)
                                                                                                                                                                } else {
                                                                                                                                                                    v6 = (298 <= 42)
                                                                                                                                                                }
                                                                                                                                                                if (v6 || true) {
                                                                                                                                                                    if v6 {
                                                                                                                                                                        var v152 string = v130
                                                                                                                                                                        var v153 uint8 = v36
                                                                                                                                                                        delete(v4, v152)
                                                                                                                                                                        _ = v152
                                                                                                                                                                        _ = v153
                                                                                                                                                                    } else {
                                                                                                                                                                        v4["a"] = (v29 / ((4294967295) | 1))
                                                                                                                                                                        switch "hello" {
                                                                                                                                                                        case "hello":
                                                                                                                                                                            var v154 int16 = 384
                                                                                                                                                                            var v155 uint8 = (v36 % (((v36 * 128)) | 1))
                                                                                                                                                                            var v156 uint8 = (v155 + v36)
                                                                                                                                                                            _ = v154
                                                                                                                                                                            _ = v155
                                                                                                                                                                            _ = v156
                                                                                                                                                                        default:
                                                                                                                                                                            if ((v36 << ((18446744073709551615) & 63)) < v36) {
                                                                                                                                                                                v2 = append(v2, ("foo bar" + v130))
                                                                                                                                                                            } else {
                                                                                                                                                                                _ = len(v4)
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    delete(v4, v130)
                                                                                                                                                                    switch (fi21 * fi21) {
                                                                                                                                                                    case -647:
                                                                                                                                                                        v18 += (^(v18))
                                                                                                                                                                    case 32768:
                                                                                                                                                                        v15 = v127
                                                                                                                                                                    case 2147483647:
                                                                                                                                                                        var v157 uint8 = v36
                                                                                                                                                                        _ = v157
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    v158 := GenS20{F0: v127, F1: ("\n\t" + v130)}
                                                                                                                                                                    _ = v158
                                                                                                                                                                }
                                                                                                                                                                if v6 {
                                                                                                                                                                    v159 := GenS21{F0: 32767, F1: (^(v18)), F2: v36}
                                                                                                                                                                    _ = v159
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v130
                                                                                                                                                    _ = v131
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                if (v47 > 1) {
                                                                                                                                                    if true {
                                                                                                                                                        var v160 int = (fi17 / (((v27 >> ((v127) & 63))) | 1))
                                                                                                                                                        fmt.Println(fi21)
                                                                                                                                                        if v119 {
                                                                                                                                                            fmt.Println(v3, v118, fi17)
                                                                                                                                                            switch fi21 {
                                                                                                                                                            case 4294967295:
                                                                                                                                                                v123 = (v126 & v126)
                                                                                                                                                            case 0:
                                                                                                                                                                v4["a"] = (^(v29))
                                                                                                                                                            default:
                                                                                                                                                                var v161 int32 = -288
                                                                                                                                                                _ = v161
                                                                                                                                                            }
                                                                                                                                                            switch -592 {
                                                                                                                                                            case -830:
                                                                                                                                                                v162 := GenS22{F0: genFn0(0)}
                                                                                                                                                                _ = v162
                                                                                                                                                            default:
                                                                                                                                                                fmt.Println(v3, v160)
                                                                                                                                                            }
                                                                                                                                                            v163 := GenS1{F0: (v126 | -2147483648)}
                                                                                                                                                            if false {
                                                                                                                                                                if (1 == 410) {
                                                                                                                                                                    v36 += v36
                                                                                                                                                                    var v164 string = "\n\t"
                                                                                                                                                                    _ = v164
                                                                                                                                                                } else {
                                                                                                                                                                    switch "a" {
                                                                                                                                                                    case "":
                                                                                                                                                                        if (!v119) {
                                                                                                                                                                            v165 := v4["a"]
                                                                                                                                                                            v126 += (v123 & 902)
                                                                                                                                                                            var v166 uint16 = 65535
                                                                                                                                                                            _ = v165
                                                                                                                                                                            _ = v166
                                                                                                                                                                        }
                                                                                                                                                                    case "hello":
                                                                                                                                                                        if v6 {
                                                                                                                                                                            v1 = (v1 - (float32(float64((v1 * 0.0))) - (3.14159 + v1)))
                                                                                                                                                                        } else {
                                                                                                                                                                            v27 = fi17
                                                                                                                                                                            var v167 float32 = 1e10
                                                                                                                                                                            v3 = v3
                                                                                                                                                                            var v168 float64 = ((-2.25 + (3.14159 + 0.0)) - (3.14159 * (3.14159 + (-2.25 + (float64(v167) + float64((v1 - v1)))))))
                                                                                                                                                                            _ = v167
                                                                                                                                                                            _ = v168
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    var v169 float64 = -2.25
                                                                                                                                                                    var v170 int64 = int64(v56)
                                                                                                                                                                    if ((^(v126)) != 2147483647) {
                                                                                                                                                                        fmt.Println(v27)
                                                                                                                                                                    } else {
                                                                                                                                                                        if false {
                                                                                                                                                                            var v171 float32 = (v1 + v1)
                                                                                                                                                                            _ = v171
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    _ = len(v2)
                                                                                                                                                                    _ = v169
                                                                                                                                                                    _ = v170
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                if v6 {
                                                                                                                                                                    var v172 uint16 = 0
                                                                                                                                                                    _ = len(v4)
                                                                                                                                                                    _ = v172
                                                                                                                                                                }
                                                                                                                                                                v36 = (v36 & (v36 & 128))
                                                                                                                                                                if false {
                                                                                                                                                                    var v173 uint64 = 2
                                                                                                                                                                    _ = v173
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            v174 := GenS14{F0: (-(v27)), F1: v15, F2: ("a" + ((("hello" + (("\n\t" + "\n\t") + ("foo bar" + ("hello" + "")))) + "") + "foo bar"))}
                                                                                                                                                            v47 = v3
                                                                                                                                                            v175 := v4[""]
                                                                                                                                                            _ = v163
                                                                                                                                                            _ = v174
                                                                                                                                                            _ = v175
                                                                                                                                                        }
                                                                                                                                                        _ = v160
                                                                                                                                                    }
                                                                                                                                                    if (v36 == v36) {
                                                                                                                                                        var v176 bool = ((v126 << ((v127) & 63)) > (v3 + -1))
                                                                                                                                                        var v177 int64 = 0
                                                                                                                                                        v178 := GenS23{F0: (v36 - uint8(v15)), F1: (-(v177))}
                                                                                                                                                        switch v27 {
                                                                                                                                                        case 0:
                                                                                                                                                            v179 := GenS24{F0: v47, F1: 32767, F2: (v177 >> ((uint(v36)) & 63))}
                                                                                                                                                            _ = v179
                                                                                                                                                        case 1:
                                                                                                                                                            v29 += (v56 / ((745) | 1))
                                                                                                                                                        }
                                                                                                                                                        v123 += 0
                                                                                                                                                        if true {
                                                                                                                                                            switch -9223372036854775808 {
                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                fmt.Println(v18, v126, v118)
                                                                                                                                                            case -85:
                                                                                                                                                                if true {
                                                                                                                                                                    if ("\n\t" >= "\n\t") {
                                                                                                                                                                        if true {
                                                                                                                                                                            var v180 int64 = int64(v127)
                                                                                                                                                                            switch ("foo bar" + "") {
                                                                                                                                                                            case "a":
                                                                                                                                                                                var v181 int64 = (v180 % ((v177) | 1))
                                                                                                                                                                                _ = v181
                                                                                                                                                                            case "foo bar":
                                                                                                                                                                                var v182 int32 = v3
                                                                                                                                                                                _ = v182
                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                switch ("\n\t" + (("\n\t" + "\n\t") + "hello")) {
                                                                                                                                                                                case "\n\t":
                                                                                                                                                                                    v183 := GenS25{F0: 1, F1: 100}
                                                                                                                                                                                    _ = v183
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            var v184 float64 = float64(v1)
                                                                                                                                                                            _ = v180
                                                                                                                                                                            _ = v184
                                                                                                                                                                        }
                                                                                                                                                                        switch (("\n\t" + "foo bar") + "\n\t") {
                                                                                                                                                                        case "a":
                                                                                                                                                                            var v185 float32 = -2.25
                                                                                                                                                                            _ = v185
                                                                                                                                                                        }
                                                                                                                                                                        v1 = v1
                                                                                                                                                                        v186 := GenS26{F0: false}
                                                                                                                                                                        _ = v186
                                                                                                                                                                    } else {
                                                                                                                                                                        switch "" {
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            var v187 int = 42
                                                                                                                                                                            v188 := GenS4{F0: 463}
                                                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                                                v5[0] = (v29 + (^(v56)))
                                                                                                                                                                            }
                                                                                                                                                                            _ = v187
                                                                                                                                                                            _ = v188
                                                                                                                                                                        case "a":
                                                                                                                                                                            var v189 uint = (v15 % ((4294967295) | 1))
                                                                                                                                                                            _ = v189
                                                                                                                                                                        default:
                                                                                                                                                                            fmt.Println(v6, v36, v29)
                                                                                                                                                                            delete(v4, "hello")
                                                                                                                                                                        }
                                                                                                                                                                        var v190 int16 = 321
                                                                                                                                                                        _ = v190
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            case 32768:
                                                                                                                                                                v47 += v3
                                                                                                                                                            default:
                                                                                                                                                                v15 = v15
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v176
                                                                                                                                                        _ = v177
                                                                                                                                                        _ = v178
                                                                                                                                                    } else {
                                                                                                                                                        if (((v127 ^ v127) != (v127 * 4294967295)) && (!((v15 & v127) > v15))) {
                                                                                                                                                            v2 = append(v2, "foo bar")
                                                                                                                                                            v123 += v47
                                                                                                                                                            fmt.Println(v47, v47)
                                                                                                                                                        }
                                                                                                                                                        var v191 int64 = 877
                                                                                                                                                        v192 := GenS15{F0: ((v1 * v1) <= 0.0)}
                                                                                                                                                        _ = v191
                                                                                                                                                        _ = v192
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    if len(v2) > 0 {
                                                                                                                                                        v2[0] = ("a" + ("" + ("hello" + "foo bar")))
                                                                                                                                                    }
                                                                                                                                                    v36 -= v36
                                                                                                                                                }
                                                                                                                                                v4["foo bar"] = (^(v56))
                                                                                                                                                if (int16(v127) < 32767) {
                                                                                                                                                    var v193 int = (fi21 / (((fi17 + 1)) | 1))
                                                                                                                                                    v126 = 42
                                                                                                                                                    v194 := GenS18{F0: 1}
                                                                                                                                                    _ = v193
                                                                                                                                                    _ = v194
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v195 float32 = ((v1 / v1) / v1)
                                                                                                                                            fmt.Println(v47, v118)
                                                                                                                                            _ = v126
                                                                                                                                            _ = v127
                                                                                                                                            _ = v195
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v196 := GenS27{F0: genFn0(9223372036854775807), F1: 1000, F2: 1000}
                                                                                                                                    var v197 bool = (float64(v1) < -2.25)
                                                                                                                                    _ = v117
                                                                                                                                    _ = v118
                                                                                                                                    _ = v119
                                                                                                                                    _ = v123
                                                                                                                                    _ = v196
                                                                                                                                    _ = v197
                                                                                                                                }
                                                                                                                                _ = v113
                                                                                                                            } else {
                                                                                                                                v57 = true
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            var v198 int = fi17
                                                                                                                            _ = len(v4)
                                                                                                                            v29 -= v56
                                                                                                                            _ = v198
                                                                                                                        }
                                                                                                                        v29 -= v29
                                                                                                                        v2 = append(v2, "foo bar")
                                                                                                                        switch "" {
                                                                                                                        case "foo bar":
                                                                                                                            if (("" + "foo bar") != "foo bar") {
                                                                                                                                v56 += (v56 << ((v15) & 63))
                                                                                                                                var v199 uint32 = 373
                                                                                                                                v200 := GenS12{F0: 2}
                                                                                                                                _ = v199
                                                                                                                                _ = v200
                                                                                                                            } else {
                                                                                                                                var v201 bool = ((!v6) && (-1 >= 1000))
                                                                                                                                v202 := GenS13{F0: 65535}
                                                                                                                                _ = v201
                                                                                                                                _ = v202
                                                                                                                            }
                                                                                                                        case "a":
                                                                                                                            if len(v5) > 0 {
                                                                                                                                v5[0] = v29
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            fmt.Println(v47, v27)
                                                                                                                        }
                                                                                                                        switch "a" {
                                                                                                                        case "\n\t":
                                                                                                                            var v203 int16 = 42
                                                                                                                            _ = v203
                                                                                                                        case "hello":
                                                                                                                            var v204 uint8 = (v36 >> ((v15) & 63))
                                                                                                                            _ = v204
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v205 int8 = -1
                                                                                                                        _ = v205
                                                                                                                    }
                                                                                                                    v1 += v1
                                                                                                                    _ = v104
                                                                                                                }
                                                                                                            case 2147483647:
                                                                                                                if true {
                                                                                                                    v15 -= 330
                                                                                                                    var v206 int = 42
                                                                                                                    switch 2147483647 {
                                                                                                                    case 9223372036854775807:
                                                                                                                        v57 = v6
                                                                                                                    case 32768:
                                                                                                                        v207 := GenS28{F0: 65535, F1: -1}
                                                                                                                        _ = v207
                                                                                                                    }
                                                                                                                    _ = v206
                                                                                                                }
                                                                                                            default:
                                                                                                                v27 += 1
                                                                                                            }
                                                                                                            var v208 int64 = -9223372036854775808
                                                                                                            var v209 string = ("foo bar" + "\n\t")
                                                                                                            _ = v208
                                                                                                            _ = v209
                                                                                                        }
                                                                                                        if false {
                                                                                                            v57 = v57
                                                                                                        }
                                                                                                        v15 = v15
                                                                                                        v210 := v4["foo bar"]
                                                                                                        _ = v210
                                                                                                    }
                                                                                                    var v211 bool = (-756 != int32(v15))
                                                                                                    var v212 uint32 = 42
                                                                                                    _ = v211
                                                                                                    _ = v212
                                                                                                }
                                                                                                switch v18 {
                                                                                                case 1:
                                                                                                    v47 -= v3
                                                                                                case 9223372036854775807:
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = v29
                                                                                                    }
                                                                                                default:
                                                                                                    delete(v4, "\n\t")
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "":
                                                                                                    v213 := GenS27{F0: genFn0(42), F1: -140, F2: -686}
                                                                                                    _ = v213
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v56 = 435
                                                                                    }
                                                                                }
                                                                                v214 := GenS27{F0: -27, F1: -32768, F2: 42}
                                                                                v56 -= v29
                                                                                _ = v214
                                                                            } else {
                                                                                v57 = (912 != (^(v18)))
                                                                                var v215 int8 = int8(v47)
                                                                                v27 = -833
                                                                                _ = v215
                                                                            }
                                                                            var v216 float32 = v1
                                                                            v216 = float32(0.1)
                                                                            _ = v216
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        var v217 bool = v6
                                                                        v218 := v4[(("" + "foo bar") + "hello")]
                                                                        var v219 uint32 = 2147483647
                                                                        _ = v217
                                                                        _ = v218
                                                                        _ = v219
                                                                        _ = v47
                                                                    }
                                                                    var v220 int8 = genFn0(0)
                                                                    _ = v220
                                                                }
                                                                _ = v36
                                                            }
                                                            _ = v35
                                                        }
                                                        delete(v4, ("" + "foo bar"))
                                                    }
                                                    v3 -= -2147483648
                                                } else {
                                                    v29 -= 596
                                                    var v221 uint = (v15 % ((v15) | 1))
                                                    if (0 >= 332) {
                                                        v221 = (v15 % ((v221) | 1))
                                                        v222 := GenS26{F0: v6}
                                                        _ = v222
                                                    } else {
                                                        var v223 float64 = (1e10 * 3.14159)
                                                        _ = v223
                                                    }
                                                    _ = v221
                                                }
                                            }
                                            _ = v29
                                        }
                                        if v6 {
                                            fmt.Println(fi17, v27)
                                            v1 += (v1 - v1)
                                            v27 = v18
                                        }
                                    } else {
                                        v15 = v15
                                    }
                                    var v224 int16 = -433
                                    _ = v224
                                    _ = v27
                                case 366:
                                    if true {
                                        v225 := GenS29{F0: 42, F1: v6, F2: 65535}
                                        switch "hello" {
                                        case "foo bar":
                                            if v6 {
                                                fmt.Println(v18)
                                                v226 := v4["\n\t"]
                                                v227 := GenS29{F0: 255, F1: v6, F2: 0}
                                                v228 := GenS30{F0: v226, F1: (v15 & (v15 * v15))}
                                                if (-43 < 103) {
                                                    if v6 {
                                                        v3 += (v3 * v3)
                                                        fmt.Println(v3, v18)
                                                        v226 -= 0
                                                    }
                                                    var v229 int64 = 42
                                                    v230 := GenS31{F0: 1, F1: int(v15)}
                                                    _ = v229
                                                    _ = v230
                                                }
                                                v4[""] = (v226 & v226)
                                                if len(v2) > 0 {
                                                    v2[0] = "hello"
                                                }
                                                _ = v226
                                                _ = v227
                                                _ = v228
                                            } else {
                                                if (32767 >= 1) {
                                                    v15 -= v15
                                                    switch "foo bar" {
                                                    case "":
                                                        v231 := GenS2{F0: 1, F1: 889, F2: (v15 % ((v15) | 1))}
                                                        _ = v231
                                                    default:
                                                        var v232 int16 = -1
                                                        _ = v232
                                                    }
                                                } else {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v233 int8 = genFn0(-9223372036854775808)
                                                    _ = v233
                                                }
                                            }
                                        default:
                                            fmt.Println(v18)
                                        }
                                        _ = v225
                                    } else {
                                        v2 = append(v2, ("\n\t" + ("a" + (("foo bar" + "") + ("" + "a")))))
                                        switch "a" {
                                        case "hello":
                                            v234 := v4["foo bar"]
                                            _ = v234
                                        case "a":
                                            if v6 {
                                                v1 += ((v1 + v1) * (((1e10 - (float32(1.5) / v1)) * v1) * 0.1))
                                                v235 := GenS32{F0: 1}
                                                var v236 int32 = v3
                                                _ = v235
                                                _ = v236
                                            }
                                        default:
                                            switch (fi17 >> ((v15) & 63)) {
                                            case 1:
                                                var v237 uint = v15
                                                _ = v237
                                            }
                                        }
                                        v4["a"] = 4294967295
                                    }
                                default:
                                    if (1e10 <= (1e10 + ((0.0 - 1.5) / 1e10))) {
                                        v3 = (v3 / ((0) | 1))
                                    }
                                }
                                if (((!(!false)) || v6) || (-1 == v3)) {
                                    fmt.Println(v6, fi17)
                                    if (1 < -9223372036854775808) {
                                        _ = len(v5)
                                        if ("hello" < (("foo bar" + "foo bar") + "\n\t")) {
                                            v238 := GenS29{F0: 32767, F1: (320 == v15), F2: 582}
                                            if true {
                                                if len(v5) > 0 {
                                                    v5[0] = 42
                                                }
                                                v3 = v3
                                                if false {
                                                    var v239 int = 127
                                                    v240 := v4[("foo bar" + ((("a" + "foo bar") + "hello") + "a"))]
                                                    fmt.Println(fi21, fi17, fi21)
                                                    _ = v239
                                                    _ = v240
                                                } else {
                                                    v1 = float32(3.14159)
                                                    v6 = v6
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = ("hello" + "foo bar")
                                                }
                                            } else {
                                                switch (fi17 * 629) {
                                                case -452:
                                                    if v6 {
                                                        if v6 {
                                                            switch (fi21 & (^(fi17))) {
                                                            case 32768:
                                                                v241 := GenS16{F0: (fi21 % ((fi21) | 1))}
                                                                _ = v241
                                                            case -378:
                                                                v15 += v15
                                                            case 9223372036854775807:
                                                                switch v18 {
                                                                case -9223372036854775808:
                                                                    v242 := GenS33{F0: 1}
                                                                    _ = v242
                                                                case 128:
                                                                    v243 := GenS34{F0: (v18 << (((v15 % ((4294967295) | 1))) & 63)), F1: 0, F2: 5}
                                                                    _ = v243
                                                                case 440:
                                                                    if v6 {
                                                                        v244 := GenS35{F0: 0, F1: false}
                                                                        _ = v244
                                                                    }
                                                                    if (120 >= 0) {
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 42
                                                                        }
                                                                    }
                                                                    var v245 float64 = (1e10 / 1e10)
                                                                    _ = v245
                                                                }
                                                            }
                                                            fmt.Println(v1, fi17, fi21)
                                                        }
                                                    }
                                                default:
                                                    v6 = v6
                                                }
                                                if (0.1 == 1e10) {
                                                    if (!false) {
                                                        switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                        case "a":
                                                            switch v18 {
                                                            case 2147483647:
                                                                delete(v4, "hello")
                                                            case -715:
                                                                var v246 uint16 = 255
                                                                _ = v246
                                                            }
                                                            switch fi21 {
                                                            case -2147483648:
                                                                var v247 float32 = v1
                                                                _ = v247
                                                            default:
                                                                delete(v4, "hello")
                                                            }
                                                            v248 := GenS36{F0: 82, F1: v6, F2: -948}
                                                            _ = v248
                                                        case "hello":
                                                            var v249 float32 = -2.25
                                                            _ = v249
                                                        }
                                                        v250 := GenS15{F0: v6}
                                                        var v251 uint16 = 0
                                                        _ = v250
                                                        _ = v251
                                                    } else {
                                                        v3 = (-(v3))
                                                    }
                                                    var v252 bool = (1000 == -85)
                                                    fmt.Println(fi21)
                                                    _ = v252
                                                } else {
                                                    switch fi17 {
                                                    case 32767:
                                                        v253 := v4[(((("hello" + "\n\t") + "hello") + (("foo bar" + "foo bar") + "a")) + (("hello" + "foo bar") + "foo bar"))]
                                                        _ = v253
                                                    case 127:
                                                        var v254 int = (v18 & 1)
                                                        _ = v254
                                                    }
                                                    v6 = true
                                                }
                                            }
                                            _ = v238
                                        } else {
                                            switch "" {
                                            case "hello":
                                                if (v6 || true) {
                                                    if (1.5 <= 1.5) {
                                                        var v255 uint32 = 545
                                                        _ = v255
                                                    }
                                                    var v256 uint8 = 58
                                                    _ = v256
                                                }
                                            case "\n\t":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                var v257 float32 = ((v1 + v1) / float32((1e10 + 3.14159)))
                                                _ = v257
                                            }
                                            if v6 {
                                                var v258 int64 = 997
                                                _ = v258
                                            }
                                            switch int(v3) {
                                            case 42:
                                                v15 += (v15 % ((0) | 1))
                                            }
                                        }
                                        v4[("" + "a")] = 4294967295
                                    } else {
                                        var v259 uint16 = 1
                                        switch -581 {
                                        case -139:
                                            switch (fi17 + (^(v18))) {
                                            case 955:
                                                var v260 uint64 = 125
                                                v3 += v3
                                                _ = v260
                                            }
                                        case 42:
                                            if true {
                                                if true {
                                                    var v261 int8 = genFn0(-810)
                                                    _ = v261
                                                }
                                            }
                                        }
                                        _ = v259
                                    }
                                    if v6 {
                                        v15 += (^(v15))
                                        fmt.Println(fi21, v3, v3)
                                        v4["foo bar"] = 65535
                                    } else {
                                        if (-1 >= 9223372036854775807) {
                                            var v262 string = "\n\t"
                                            if len(v5) > 0 {
                                                v5[0] = 4294967295
                                            }
                                            _ = v262
                                        } else {
                                            if v6 {
                                                fmt.Println(v1, v6, fi21)
                                                v263 := GenS37{F0: 32767, F1: (v15 / ((482) | 1))}
                                                v264 := GenS38{F0: -86}
                                                _ = v263
                                                _ = v264
                                            } else {
                                                v1 = 0.0
                                                fmt.Println(v18)
                                                var v265 int = fi17
                                                _ = v265
                                            }
                                            if (v6 || v6) {
                                                if v6 {
                                                    if ((650 < 1) || (v6 && true)) {
                                                        if len(v2) > 0 {
                                                            v2[0] = "foo bar"
                                                        }
                                                    }
                                                }
                                                var v266 uint32 = 375
                                                v1 = float32(float64(1e10))
                                                _ = v266
                                            } else {
                                                fmt.Println(v3, fi21, v15)
                                                delete(v4, ("" + "hello"))
                                            }
                                            switch fi17 {
                                            case 681:
                                                if ((!true) && ((2 > 942) && v6)) {
                                                    v3 = (v3 ^ v3)
                                                    fmt.Println(v6)
                                                    if (127 != 128) {
                                                        var v267 int = -9223372036854775808
                                                        v267 += (v18 - 32768)
                                                        _ = v267
                                                    }
                                                } else {
                                                    v3 = v3
                                                }
                                            case -859:
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                            default:
                                                v4[("hello" + "")] = 0
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v4["foo bar"] = 592
                                    }
                                }
                            }
                            v1 -= (float32(0.1) + 1e10)
                        }
                        var v268 int16 = -32768
                        var v269 bool = (937 >= 18446744073709551615)
                        _ = fi21
                        _ = v268
                        _ = v269
                    }
                }
            }
            for fi270 := 0; fi270 < 2; fi270++ {
                if (!(!false)) {
                    v6 = v6
                    v1 = v1
                    var v271 uint64 = 748
                    var v272 string = "foo bar"
                    v1 -= v1
                    v273 := GenS39{F0: (((((v272 + (v272 + v272)) + v272) + (v272 + (v272 + v272))) + v272) + ("hello" + v272)), F1: "hello", F2: 32767}
                    _ = v271
                    _ = v272
                    _ = v273
                }
                v1 += 3.14159
                switch (^(v18)) {
                case 127:
                    if (!(v6 && (v15 != v15))) {
                        var v274 uint16 = 2
                        _ = v274
                    } else {
                        v275 := GenS34{F0: fi17, F1: -1, F2: 42}
                        var v276 uint = (v15 * v15)
                        _ = v275
                        _ = v276
                    }
                case 32767:
                    switch (("\n\t" + "\n\t") + "\n\t") {
                    case "hello":
                        if v6 {
                            v277 := GenS9{F0: 0, F1: -2147483648}
                            var v278 uint32 = 1
                            var v279 uint16 = 65535
                            _ = v277
                            _ = v278
                            _ = v279
                        } else {
                            var v280 uint32 = 42
                            _ = v280
                        }
                    case "":
                        v1 -= v1
                    }
                case 9223372036854775807:
                    v281 := GenS40{F0: 42, F1: v3, F2: 901}
                    _ = v281
                }
                switch v18 {
                case 9223372036854775807:
                    var v282 int64 = -9223372036854775808
                    _ = v282
                case 0:
                    v3 += v3
                    v15 -= v15
                default:
                    v15 += v15
                    _ = len(v4)
                }
                switch -1000 {
                case -9223372036854775808:
                    v283 := GenS16{F0: fi270}
                    _ = v283
                case -526:
                    v284 := GenS10{F0: (((("a" + "a") + "a") + "\n\t") + ((("hello" + "\n\t") + "hello") + "a")), F1: -280, F2: genFn0(-93)}
                    _ = v284
                case -631:
                    v3 = v3
                default:
                    v15 += (^(v15))
                }
                _ = fi270
            }
        }
        _ = fi17
        _ = v18
    }
    var v285 uint16 = 32767
    for fi286 := 0; fi286 < 13; fi286++ {
        var v287 float64 = 0.1
        fmt.Println(v287, v3)
        v287 = -2.25
        _ = fi286
        _ = v287
    }
    switch -9223372036854775808 {
    case 9223372036854775807:
        v288 := GenS29{F0: v285, F1: true, F2: 2}
        _ = v288
    case -31:
        v285 = 65535
    default:
        v15 = v15
    }
    if v6 {
        if ((0.1 - -2.25) >= (float32(1e10) - (v1 / ((float32(3.14159) + (1e10 + v1)) * v1)))) {
            switch 835 {
            case 0:
                for fi289 := 0; fi289 < 9; fi289++ {
                    if v6 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        if v6 {
                            if ((((3.14159 - float64(v1)) / 1.0) + 1e10) == 1e10) {
                                var v290 int32 = (v3 % ((-1) | 1))
                                if v6 {
                                    fmt.Println(v285)
                                }
                                v291 := v4[(("" + ("hello" + ((("foo bar" + (("foo bar" + "a") + "a")) + "") + ("\n\t" + (("a" + "hello") + ("a" + ("hello" + (("hello" + "hello") + "")))))))) + ("hello" + ""))]
                                _ = v290
                                _ = v291
                            } else {
                                fmt.Println(fi289)
                                v3 = v3
                                v15 += 1
                            }
                            switch fi289 {
                            case 42:
                                var v292 string = "foo bar"
                                _ = v292
                            }
                            switch ("a" + (("hello" + ("hello" + "\n\t")) + ("" + (("" + "foo bar") + "a")))) {
                            case "":
                                var v293 int64 = 0
                                _ = v293
                            default:
                                _ = len(v4)
                            }
                        }
                        switch (fi289 | 214) {
                        case 9223372036854775807:
                            var v294 float64 = -2.25
                            _ = v294
                        case 33:
                            if (!(true && (v15 <= uint(v285)))) {
                                if true {
                                    v285 = v285
                                } else {
                                    fmt.Println(v6, v6, v285)
                                }
                            } else {
                                var v295 uint32 = 400
                                _ = v295
                            }
                        default:
                            if v6 {
                                var v296 int = (fi289 * -2147483648)
                                _ = v296
                            }
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi289
                }
            default:
                var v297 float64 = 0.0
                _ = v297
            }
            var v298 int = 127
            v285 -= v285
            _ = v298
        }
    } else {
        v299 := GenS41{F0: 0}
        for fi300 := 0; fi300 < 9; fi300++ {
            var v301 uint32 = 2147483647
            _ = fi300
            _ = v301
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v299
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v285:", v285)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v285
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
