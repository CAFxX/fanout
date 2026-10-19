package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint16 }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS16 struct { F0 int32; F1 int32 }
type GenS17 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS18 struct { F0 int; F1 uint8 }
func genFn0(p0 int64, p1 uint16) int16 {
    fmt.Println(p1, p1)
    switch "a" {
    case "":
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 610}
            p1 = p1
            _ = fi7
            _ = v8
        }
    case "a":
        v9 := GenS1{F0: 42}
        _ = v9
    default:
        p1 = p1
    }
    var v10 uint16 = (p1 * (p1 << ((2) & 63)))
    _ = p0
    _ = p1
    _ = v10
    return 492
}
func genFn1(p0 uint) int32 {
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 1
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(p0, fi13)
        v14 := GenS2{F0: (!true), F1: (fi13 << ((255) & 63)), F2: 0}
        _ = fi13
        _ = v14
    }
    if true {
        var v15 int64 = 970
        v15 += v15
        v16 := GenS3{F0: 0}
        _ = v15
        _ = v16
    } else {
        v17 := []int{42, 9223372036854775807, -1}
        p0 -= 2
        _ = v17
    }
    _ = p0
    return -2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 127
    v3 := make(map[string]string)
    var v4 int16 = 1000
    var v5 bool = false
    v6 := []string{"\n\t", "\n\t", (("hello" + "foo bar") + ("foo bar" + ("\n\t" + ("\n\t" + ""))))}
    switch ("\n\t" + "") {
    case "\n\t":
        var v18 string = "a"
        _ = v18
    case "":
        v5 = false
    default:
        v4 = (^(v4))
    }
    delete(v1, 4294967295)
    if v5 {
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 float32 = 1.5
            v5 = false
            _ = fi19
            _ = v20
        }
    }
    var v21 uint16 = 732
    fmt.Println(v2)
    if v5 {
        for fi22 := 0; fi22 < 4; fi22++ {
            fmt.Println(v4, fi22, v21)
            var v23 bool = (v5 || ((("foo bar" + "\n\t") + "hello") <= "foo bar"))
            v24 := GenS4{F0: false, F1: v2}
            _ = fi22
            _ = v23
            _ = v24
        }
        switch "a" {
        case "":
            fmt.Println(v4)
        }
        _ = len(v6)
    }
    if (-1 == 9223372036854775807) {
        var v25 int32 = 1
        var v26 uint = 0
        v2 = v2
        _ = v25
        _ = v26
    } else {
        v5 = false
        switch "hello" {
        case "a":
            var v27 int32 = -115
            _ = v27
        case "hello":
            if v5 {
                v5 = (128 >= 128)
                var v28 uint32 = 2147483647
                _ = v28
            }
        }
        v4 = (v4 / ((-32768) | 1))
    }
    v5 = (2147483647 != 42)
    if ((-1 < -9) && (false || (-2147483648 >= 404))) {
        fmt.Println(v4)
        v29 := GenS5{F0: 2147483647, F1: "foo bar"}
        switch (v2 - (v2 * (v2 ^ (v2 & v2)))) {
        case 32767:
            if len(v6) > 0 {
                v6[0] = "a"
            }
        case -390:
            for fi30 := 0; fi30 < 3; fi30++ {
                switch "\n\t" {
                case "hello":
                    v31 := GenS5{F0: 1, F1: (("a" + "hello") + "\n\t")}
                    _ = v31
                case "foo bar":
                    if (127 != uint8(v21)) {
                        var v32 int = fi30
                        switch (v2 % (((v32 << ((255) & 63))) | 1)) {
                        case 42:
                            if false {
                                delete(v1, 2147483647)
                                fmt.Println(v32)
                            }
                        }
                        if (59 == 18) {
                            v33 := GenS6{F0: 2}
                            var v34 uint32 = 422
                            var v35 string = "a"
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            v4 = (v4 / (((v4 % ((-32768) | 1))) | 1))
                        }
                        _ = v32
                    } else {
                        if v5 {
                            for fi36 := 0; fi36 < 17; fi36++ {
                                if false {
                                    var v37 uint = 255
                                    _ = v37
                                } else {
                                    v38 := GenS7{F0: 1, F1: v21}
                                    _ = v38
                                }
                                _ = fi36
                            }
                        } else {
                            v39 := GenS1{F0: 691}
                            _ = v39
                        }
                    }
                default:
                    for fi40 := 0; fi40 < 2; fi40++ {
                        if (!(false || v5)) {
                            _ = len(v1)
                            v41 := GenS2{F0: ((4294967295 >= 42) || (!(128 == 128))), F1: v2, F2: 127}
                            _ = v41
                        } else {
                            var v42 int = fi40
                            var v43 int32 = -2147483648
                            var v44 int8 = -12
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        v21 += (v21 % (((v21 + v21)) | 1))
                        v2 = v2
                        _ = fi40
                    }
                }
                _ = fi30
            }
        case -2147483648:
            v45 := GenS8{F0: ((1e10 / 1e10) >= 0.0), F1: (v4 & int16(v2))}
            _ = v45
        default:
            v46 := GenS9{F0: (-9223372036854775808 <= -9223372036854775808)}
            _ = v46
        }
        _ = v29
    } else {
        if false {
            v47 := GenS10{F0: false}
            if v5 {
                if v5 {
                    v5 = (65535 > 42)
                } else {
                    v4 = v4
                    var v48 bool = false
                    _ = v48
                }
                for fi49 := 0; fi49 < 8; fi49++ {
                    var v50 uint32 = 0
                    _ = fi49
                    _ = v50
                }
                for fi51 := 0; fi51 < 3; fi51++ {
                    var v52 int32 = 42
                    for fi53 := 0; fi53 < 5; fi53++ {
                        if (!(((3.14159 - ((0.0 / float64(1.5)) - (-2.25 - 1e10))) + 1.5) < 1e10)) {
                            if v5 {
                                if true {
                                    fmt.Println(v52, v5)
                                    var v54 string = ""
                                    _ = v54
                                }
                                var v55 uint = 4294967295
                                _ = v55
                            } else {
                                var v56 int16 = (v4 / ((int16(v21)) | 1))
                                var v57 float64 = ((0.0 / 1.0) / float64(float32(float64(-2.25))))
                                _ = v56
                                _ = v57
                            }
                            switch "" {
                            case "":
                                var v58 uint = 645
                                _ = v58
                            case "\n\t":
                                var v59 int = (-(v2))
                                _ = v59
                            case "hello":
                                var v60 uint = 0
                                _ = v60
                            default:
                                if false {
                                    switch (fi51 * (fi51 % (((^(fi51))) | 1))) {
                                    case 128:
                                        var v61 int = (v2 | fi53)
                                        _ = v61
                                    case 4294967295:
                                        var v62 string = ("\n\t" + "foo bar")
                                        _ = v62
                                    default:
                                        var v63 int16 = v4
                                        _ = v63
                                    }
                                    _ = len(v1)
                                    delete(v3, "hello")
                                } else {
                                    v5 = false
                                }
                                var v64 int16 = v4
                                _ = v64
                            }
                        } else {
                            var v65 uint8 = 0
                            _ = v65
                        }
                        v66 := GenS11{F0: 2, F1: (v4 & (v4 | 32767))}
                        v21 += (v21 << ((218) & 63))
                        _ = fi53
                        _ = v66
                    }
                    if (((2147483648 != 65535) && (true && (2 <= 413))) && (!v5)) {
                        v5 = false
                        if (65535 != 42) {
                            var v67 float64 = float64(0.1)
                            switch "" {
                            case "hello":
                                var v68 bool = ((v5 || v5) && (v5 && v5))
                                _ = v68
                            case "a":
                                v69 := GenS10{F0: (!(float64(float32(v67)) == v67))}
                                _ = v69
                            case "foo bar":
                                v70 := GenS12{F0: (^(v2))}
                                _ = v70
                            default:
                                var v71 uint = 2
                                _ = v71
                            }
                            _ = v67
                        } else {
                            v72 := GenS13{F0: 16}
                            for fi73 := 0; fi73 < 10; fi73++ {
                                if false {
                                    v52 = v52
                                    var v74 int = 2147483647
                                    v5 = v5
                                    if true {
                                        switch "a" {
                                        case "a":
                                            var v75 int = (v2 & (fi73 | (v74 - fi73)))
                                            _ = v75
                                        case "hello":
                                            fmt.Println(v2, v2)
                                        default:
                                            var v76 int16 = -707
                                            _ = v76
                                        }
                                    } else {
                                        fmt.Println(v4)
                                        if len(v6) > 0 {
                                            v6[0] = "hello"
                                        }
                                        var v77 int8 = 100
                                        _ = v77
                                    }
                                    var v78 int = fi73
                                    _ = v74
                                    _ = v78
                                }
                                _ = fi73
                            }
                            _ = len(v6)
                            _ = v72
                        }
                        if true {
                            v2 = v2
                        } else {
                            switch (v2 << ((255) & 63)) {
                            case 128:
                                var v79 uint8 = 127
                                _ = v79
                            case -2147483648:
                                v5 = true
                            default:
                                v21 -= v21
                            }
                            v6 = append(v6, ("\n\t" + "hello"))
                            if (836 != 411) {
                                if (-1 < -934) {
                                    var v80 uint8 = 127
                                    fmt.Println(v2)
                                    _ = v80
                                }
                                if len(v6) > 0 {
                                    v6[0] = "a"
                                }
                            } else {
                                if true {
                                    switch fi51 {
                                    case -2147483648:
                                        for fi81 := 0; fi81 < 11; fi81++ {
                                            v52 = (-(v52))
                                            _ = fi81
                                        }
                                    }
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if false {
                                    var v82 uint = 0
                                    for fi83 := 0; fi83 < 17; fi83++ {
                                        if v5 {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = len(v1)
                                            v52 += (v52 - (v52 ^ (v52 + v52)))
                                        }
                                        _ = fi83
                                    }
                                    _ = v82
                                }
                            }
                        }
                    } else {
                        for fi84 := 0; fi84 < 16; fi84++ {
                            switch (v2 | -1) {
                            case 2147483647:
                                switch fi84 {
                                case 32767:
                                    if v5 {
                                        v4 -= (v4 & v4)
                                        fmt.Println(v5)
                                    }
                                default:
                                    if (false || v5) {
                                        var v85 int64 = -1
                                        var v86 int32 = (v52 ^ (v52 - (^(v52))))
                                        _ = v85
                                        _ = v86
                                    } else {
                                        fmt.Println(v21, v52, fi84)
                                        v87 := GenS11{F0: 664, F1: v4}
                                        _ = v87
                                    }
                                }
                                var v88 uint32 = 1
                                if (!v5) {
                                    if false {
                                        var v89 uint64 = 187
                                        v90 := GenS9{F0: (false && ((4294967295 > 42) || (100 == 0)))}
                                        _ = v89
                                        _ = v90
                                    }
                                    var v91 string = ""
                                    v2 -= 128
                                    _ = v91
                                }
                                _ = v88
                            case -1:
                                delete(v3, "")
                            case 32767:
                                if v5 {
                                    v92 := GenS14{F0: ((("hello" + ("hello" + "hello")) + "") + "foo bar"), F1: 42}
                                    if len(v6) > 0 {
                                        v6[0] = "a"
                                    }
                                    fmt.Println(v5)
                                    _ = v92
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    v93 := GenS15{F0: 2147483648, F1: 394, F2: 42}
                                    switch "" {
                                    case "foo bar":
                                        v94 := GenS13{F0: 127}
                                        _ = v94
                                    default:
                                        v21 = v21
                                    }
                                    switch fi51 {
                                    case 9223372036854775807:
                                        var v95 int64 = -9223372036854775808
                                        _ = v95
                                    default:
                                        switch (v2 / ((fi51) | 1)) {
                                        case -2147483648:
                                            _ = len(v6)
                                        case -407:
                                            var v96 string = "a"
                                            _ = v96
                                        default:
                                            fmt.Println(v21)
                                        }
                                    }
                                    delete(v3, "")
                                    _ = v93
                                }
                            default:
                                v97 := GenS16{F0: (v52 % (((^(v52))) | 1)), F1: genFn1(255)}
                                _ = v97
                            }
                            if (199 != 127) {
                                var v98 int32 = (v52 >> ((42) & 63))
                                switch ("hello" + ("a" + "hello")) {
                                case "\n\t":
                                    v6 = append(v6, ("" + "hello"))
                                case "hello":
                                    var v99 int16 = (v4 + -75)
                                    _ = v99
                                }
                                _ = v98
                            }
                            _ = fi84
                        }
                    }
                    _ = fi51
                    _ = v52
                }
            } else {
                for fi100 := 0; fi100 < 5; fi100++ {
                    if false {
                        for fi101 := 0; fi101 < 12; fi101++ {
                            var v102 int32 = 42
                            fmt.Println(v4, v5, v4)
                            if ("\n\t" != ("" + "hello")) {
                                var v103 uint64 = 18446744073709551615
                                switch ("" + (((("" + "foo bar") + "hello") + "\n\t") + ("" + "hello"))) {
                                case "foo bar":
                                    if ((!(float64(3.14159) < (1.5 / 0.1))) && (!((v102 % (((v102 - genFn1(255))) | 1)) <= int32(v21)))) {
                                        fmt.Println(v102)
                                        v4 = (-(v4))
                                    }
                                }
                                _ = v103
                            }
                            _ = fi101
                            _ = v102
                        }
                        v104 := GenS17{F0: 1, F1: 0, F2: 382}
                        v5 = v5
                        _ = v104
                    }
                    _ = fi100
                }
            }
            _ = v47
        }
        if v5 {
            switch (("a" + "a") + "a") {
            case "":
                if v5 {
                    v21 -= 102
                    _ = len(v1)
                    var v105 int8 = 78
                    _ = v105
                }
            case "foo bar":
                v5 = (!(v4 != (v4 >> ((42) & 63))))
            default:
                fmt.Println(v21)
            }
            fmt.Println(v4)
        }
    }
    if v5 {
        v4 -= v4
        switch ("" + ("\n\t" + "")) {
        case "\n\t":
            v5 = v5
        case "a":
            var v106 uint32 = 2147483647
            _ = v106
        case "":
            if len(v6) > 0 {
                v6[0] = "hello"
            }
        }
        fmt.Println(v5, v5)
    } else {
        if (301 == 42) {
            var v107 float32 = (0.1 + -2.25)
            _ = v107
        }
        v108 := GenS18{F0: v2, F1: 2}
        _ = v108
    }
    if (v5 && false) {
        _ = len(v3)
    }
    v2 = (^(v2))
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
