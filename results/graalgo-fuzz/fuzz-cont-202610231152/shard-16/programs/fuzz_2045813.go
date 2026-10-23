package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 bool }
type GenS4 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 int32; F1 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64; F1 uint32; F2 bool }
type GenS13 struct { F0 uint8; F1 string }
type GenS14 struct { F0 int; F1 int32 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int16; F1 int64; F2 int }
func genFn0(p0 int64, p1 int32, p2 uint) int64 {
    var v4 uint = p2
    var v5 float64 = 0.0
    v6 := []int{-548, 2147483647}
    var v7 int8 = 127
    var v8 uint8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int16, p1 int64) int32 {
    var v9 float32 = 0.0
    v9 -= ((float32((3.14159 / -2.25)) * -2.25) / (0.1 - v9))
    p0 = p0
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("" + (("foo bar" + ("hello" + "\n\t")) + "hello")), ("\n\t" + "hello"), "hello"}
    v3 := make(map[string]bool)
    delete(v1, 1)
    if false {
        switch -123 {
        case 1:
            v10 := GenS1{F0: false, F1: -842}
            _ = v10
        case 32767:
            var v11 int32 = 42
            _ = v11
        default:
            var v12 int32 = -2147483648
            _ = v12
        }
    } else {
        v13 := GenS1{F0: (!(1 <= 42)), F1: 42}
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        var v14 int32 = 0
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        switch fi15 {
        case -2147483648:
            _ = len(v3)
        case 411:
            var v16 int32 = 270
            _ = v16
        }
        v17 := GenS2{F0: 65535, F1: -1}
        _ = fi15
        _ = v17
    }
    var v18 uint8 = 255
    if false {
        var v19 bool = (!(-553 >= 9223372036854775807))
        v20 := GenS2{F0: 602, F1: -2147483648}
        switch "\n\t" {
        case "hello":
            delete(v1, 989)
        case "":
            v21 := GenS3{F0: -60, F1: -504, F2: (0 > -1)}
            _ = v21
        default:
            v18 = v18
        }
        _ = v19
        _ = v20
    } else {
        if (true || true) {
            v18 += (^(v18))
            var v22 uint16 = 255
            _ = v22
        } else {
            v23 := GenS2{F0: 32767, F1: 315}
            var v24 uint16 = 0
            switch 127 {
            case 42:
                v25 := GenS1{F0: true, F1: int64(v18)}
                _ = v25
            case 589:
                if (v18 != (^(v18))) {
                    v18 -= 87
                    v24 = v24
                } else {
                    var v26 float32 = -2.25
                    _ = len(v3)
                    _ = v26
                }
            case -814:
                for fi27 := 0; fi27 < 6; fi27++ {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        v18 += v18
                        var v29 uint8 = v18
                        fmt.Println(fi27, fi27)
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                var v30 string = "a"
                                _ = v30
                            case "":
                                if (fi28 < (fi27 + fi27)) {
                                    var v31 uint = 42
                                    if (fi27 < (fi27 << (((v31 * 1)) & 63))) {
                                        v24 += (v24 & (v24 & v24))
                                        if (44 < -128) {
                                            var v32 uint32 = 0
                                            switch -2147483648 {
                                            case -197:
                                                v29 = v18
                                            default:
                                                v32 = v32
                                            }
                                            v18 -= 138
                                            _ = len(v2)
                                            _ = v32
                                        } else {
                                            switch -968 {
                                            case 0:
                                                v33 := GenS4{F0: (65535 <= 905), F1: 65535, F2: 919}
                                                v31 = (v31 % (((v31 ^ 42)) | 1))
                                                _ = v33
                                            case 128:
                                                var v34 bool = (true && false)
                                                _ = v34
                                            case 845:
                                                v35 := GenS1{F0: ((float32(0.1) + -2.25) != (-2.25 * (-2.25 - (float32(0.1) - 0.0)))), F1: 9223372036854775807}
                                                _ = v35
                                            }
                                        }
                                        var v36 bool = true
                                        var v37 uint32 = 580
                                        switch "hello" {
                                        case "hello":
                                            if (fi28 > fi27) {
                                                v38 := GenS5{F0: (v24 - 1), F1: 634}
                                                _ = v38
                                            } else {
                                                switch (fi27 % ((127) | 1)) {
                                                case 1:
                                                    v39 := GenS5{F0: v24, F1: fi27}
                                                    _ = v39
                                                case -2147483648:
                                                    _ = len(v3)
                                                }
                                                v36 = v36
                                            }
                                        case "\n\t":
                                            if v36 {
                                                v40 := GenS5{F0: (v24 & (v24 >> ((2) & 63))), F1: fi28}
                                                if v36 {
                                                    v37 -= (v37 & (v37 + v37))
                                                    if false {
                                                        v41 := v1[(-(fi28))]
                                                        v37 -= v37
                                                        _ = v41
                                                    }
                                                    var v42 bool = true
                                                    _ = v42
                                                } else {
                                                    v29 -= (^(v18))
                                                }
                                                _ = v40
                                            }
                                        case "foo bar":
                                            if v36 {
                                                v43 := GenS1{F0: (1 >= -2147483648), F1: genFn0(1, 0, v31)}
                                                fmt.Println(v29, fi28)
                                                _ = v43
                                            } else {
                                                fmt.Println(fi28, v36)
                                                var v44 bool = true
                                                _ = v44
                                            }
                                        default:
                                            var v45 bool = v36
                                            _ = v45
                                        }
                                        _ = v36
                                        _ = v37
                                    } else {
                                        v46 := GenS6{F0: "hello", F1: (v24 >> ((v31) & 63))}
                                        if true {
                                            var v47 int32 = 2147483647
                                            var v48 uint8 = v29
                                            var v49 uint16 = 0
                                            v50 := GenS7{F0: (290 != 42)}
                                            v51 := GenS8{F0: -199}
                                            if (1 == -32768) {
                                                v52 := GenS5{F0: (v24 % (((^(v49))) | 1)), F1: (fi28 - (fi27 ^ (fi28 * (fi28 - fi27))))}
                                                _ = v52
                                            } else {
                                                if true {
                                                    fmt.Println(fi28, v18)
                                                } else {
                                                    var v53 bool = ((!((!(false && false)) || (v24 <= (v49 - v49)))) || (1 <= 0))
                                                    v24 += (v49 & (v24 / ((v24) | 1)))
                                                    if len(v2) > 0 {
                                                        v2[0] = ""
                                                    }
                                                    var v54 uint16 = (v24 / (((^(v49))) | 1))
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v55 uint32 = 984
                                                _ = v55
                                            }
                                            v48 -= (^(v29))
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                            _ = v51
                                        } else {
                                            var v56 int64 = 9223372036854775807
                                            _ = v56
                                        }
                                        _ = v46
                                    }
                                    _ = v31
                                } else {
                                    v29 = v29
                                    if true {
                                        var v57 uint = 65535
                                        delete(v3, "foo bar")
                                        if true {
                                            v18 = (v18 / ((v29) | 1))
                                            v57 = v57
                                            v18 -= 0
                                        } else {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v58 uint = (v57 >> (((^(v57))) & 63))
                                            v58 = (^(v57))
                                            _ = v58
                                        }
                                        _ = v57
                                    }
                                    fmt.Println(v29)
                                }
                            case "foo bar":
                                switch ("\n\t" + ("\n\t" + "foo bar")) {
                                case "foo bar":
                                    var v59 int = (fi27 % ((fi28) | 1))
                                    _ = v59
                                case "hello":
                                    var v60 int = fi28
                                    _ = v60
                                default:
                                    v2 = append(v2, "hello")
                                }
                            default:
                                if (-1 < 0) {
                                    fmt.Println(fi27, fi28)
                                }
                            }
                        } else {
                            switch fi27 {
                            case 32767:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case 4294967295:
                                var v61 string = ("\n\t" + "hello")
                                _ = v61
                            case -699:
                                switch "hello" {
                                case "a":
                                    v62 := v1[32768]
                                    _ = v62
                                case "foo bar":
                                    if false {
                                        v63 := GenS9{F0: v24, F1: v24}
                                        var v64 bool = (v24 == (^(v24)))
                                        _ = v63
                                        _ = v64
                                    }
                                case "\n\t":
                                    var v65 bool = (((!true) && false) && (1000 >= 0))
                                    _ = v65
                                default:
                                    var v66 int = (fi27 & fi27)
                                    _ = v66
                                }
                                var v67 uint32 = 751
                                switch "" {
                                case "\n\t":
                                    if true {
                                        v24 -= v24
                                        if ((!(1 > -128)) || ((1.5 * 3.14159) == ((-2.25 * ((((1.5 - 3.14159) - (1e10 / ((-2.25 / -2.25) / 3.14159))) / 3.14159) - 0.1)) + float64((1.5 * float32(0.1)))))) {
                                            var v68 int8 = 0
                                            v24 -= v24
                                            if true {
                                                var v69 int32 = 568
                                                v70 := GenS10{F0: -809, F1: (v18 >> ((4294967295) & 63))}
                                                var v71 int8 = (v68 * (v68 * v68))
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                            _ = v68
                                        } else {
                                            _ = len(v2)
                                            switch 1 {
                                            case 498:
                                                if true {
                                                    var v72 float32 = -2.25
                                                    switch 128 {
                                                    case 32768:
                                                        if ((((("\n\t" + "") + "foo bar") + "a") + "\n\t") < (("" + ((("foo bar" + ("\n\t" + "a")) + ("a" + "a")) + "a")) + ("foo bar" + (("foo bar" + ("" + "a")) + ("" + (("a" + "\n\t") + "foo bar")))))) {
                                                            var v73 bool = true
                                                            _ = v73
                                                        }
                                                    default:
                                                        v74 := GenS2{F0: v24, F1: -796}
                                                        _ = v74
                                                    }
                                                    _ = v72
                                                }
                                            case 127:
                                                v75 := GenS11{F0: v29}
                                                _ = v75
                                            case -2147483648:
                                                var v76 bool = true
                                                _ = v76
                                            }
                                        }
                                        var v77 uint8 = (v29 >> ((42) & 63))
                                        _ = v77
                                    }
                                case "a":
                                    var v78 int64 = 9223372036854775807
                                    _ = v78
                                case "foo bar":
                                    switch (fi27 / ((fi27) | 1)) {
                                    case 864:
                                        if false {
                                            var v79 int = (fi27 * (fi28 << ((469) & 63)))
                                            fmt.Println(v29)
                                            _ = v79
                                        }
                                    case -9223372036854775808:
                                        var v80 uint16 = 255
                                        _ = v80
                                    case -2147483648:
                                        if false {
                                            v24 = (v24 ^ 65535)
                                            fmt.Println(v18)
                                        }
                                    default:
                                        var v81 uint16 = (v24 * v24)
                                        v18 = 130
                                        fmt.Println(v24, fi27, v67)
                                        _ = v81
                                    }
                                }
                                _ = v67
                            default:
                                if true {
                                    v24 += v24
                                    var v82 uint8 = v18
                                    _ = v82
                                } else {
                                    if false {
                                        v83 := GenS12{F0: 255, F1: 282, F2: (3.14159 == 0.0)}
                                        _ = v83
                                    }
                                    fmt.Println(fi28, fi27)
                                }
                            }
                            v29 -= 90
                            switch 0 {
                            case 32767:
                                switch "hello" {
                                case "\n\t":
                                    var v84 bool = (!(2147483648 == 2147483648))
                                    _ = v84
                                }
                            case 425:
                                fmt.Println(fi28, fi28, fi28)
                            }
                            switch ("" + "hello") {
                            case "foo bar":
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            case "":
                                v29 = (v18 - 127)
                            case "\n\t":
                                v24 = (^(v24))
                                switch fi28 {
                                case 624:
                                    switch (fi27 - fi28) {
                                    case 291:
                                        var v85 int8 = -128
                                        _ = v85
                                    case -586:
                                        v29 = v18
                                    default:
                                        v29 -= v18
                                    }
                                case 490:
                                    delete(v3, ("a" + "foo bar"))
                                }
                                v86 := v1[fi27]
                                _ = v86
                            default:
                                if (113 > 830) {
                                    v24 -= (v24 * v24)
                                } else {
                                    v24 -= v24
                                    v87 := GenS1{F0: false, F1: -9223372036854775808}
                                    _ = v87
                                }
                            }
                        }
                        _ = fi28
                        _ = v29
                    }
                    if (0 > -1) {
                        if false {
                            switch (fi27 % ((-9223372036854775808) | 1)) {
                            case -259:
                                v24 = (v24 * v24)
                            case 32767:
                                for fi88 := 0; fi88 < 9; fi88++ {
                                    v18 -= v18
                                    v18 -= v18
                                    var v89 int = (fi27 / ((4294967295) | 1))
                                    _ = fi88
                                    _ = v89
                                }
                            default:
                                switch fi27 {
                                case 32768:
                                    for fi90 := 0; fi90 < 11; fi90++ {
                                        var v91 uint64 = 42
                                        _ = fi90
                                        _ = v91
                                    }
                                case 0:
                                    if false {
                                        v92 := GenS13{F0: v18, F1: ""}
                                        _ = len(v2)
                                        v24 = v24
                                        _ = v92
                                    }
                                case 517:
                                    var v93 int16 = 39
                                    _ = v93
                                default:
                                    fmt.Println(v18, v18)
                                }
                            }
                            for fi94 := 0; fi94 < 20; fi94++ {
                                v95 := GenS14{F0: (fi27 - int(v24)), F1: 1}
                                _ = fi94
                                _ = v95
                            }
                        }
                        for fi96 := 0; fi96 < 2; fi96++ {
                            v97 := v3[("a" + "foo bar")]
                            _ = fi96
                            _ = v97
                        }
                    }
                    _ = fi27
                }
            default:
                if len(v2) > 0 {
                    v2[0] = (("foo bar" + "") + (("hello" + ((("\n\t" + "\n\t") + "a") + ("foo bar" + "foo bar"))) + "\n\t"))
                }
            }
            _ = v23
            _ = v24
        }
        var v98 uint16 = 255
        _ = v98
    }
    switch (("\n\t" + (("a" + "foo bar") + ("" + "foo bar"))) + "a") {
    case "hello":
        for fi99 := 0; fi99 < 14; fi99++ {
            if true {
                for fi100 := 0; fi100 < 14; fi100++ {
                    var v101 int = (fi99 + (fi99 ^ fi100))
                    _ = fi100
                    _ = v101
                }
                fmt.Println(v18)
                v18 -= (^(v18))
            } else {
                for fi102 := 0; fi102 < 3; fi102++ {
                    var v103 string = ""
                    v103 = v103
                    v18 -= (v18 * 143)
                    switch fi102 {
                    case 128:
                        var v104 uint32 = 653
                        _ = v104
                    case 32767:
                        v105 := GenS15{F0: 32767}
                        _ = v105
                    case 42:
                        switch ("foo bar" + "foo bar") {
                        case "\n\t":
                            v103 = "\n\t"
                        case "":
                            v106 := GenS4{F0: (32767 == 661), F1: 884, F2: 0}
                            _ = v106
                        }
                    }
                    switch ("" + (("a" + ("foo bar" + ("a" + v103))) + ("a" + v103))) {
                    case "\n\t":
                        if (2147483647 <= -792) {
                            var v107 uint8 = (v18 | (v18 + (v18 % ((255) | 1))))
                            _ = v107
                        }
                    case "foo bar":
                        var v108 uint64 = 2
                        _ = v108
                    case "hello":
                        var v109 int = (-(fi99))
                        _ = v109
                    }
                    _ = fi102
                    _ = v103
                }
                v18 = v18
            }
            _ = fi99
        }
    case "foo bar":
        v110 := v1[-543]
        _ = v110
    }
    v18 = 2
    var v111 int32 = 2147483647
    var v112 float64 = 1e10
    if (32767 <= 32767) {
        delete(v3, ("a" + "hello"))
    } else {
        if (genFn0(genFn0(-520, (-(v111)), 255), v111, 255) <= 0) {
            v18 += (^(v18))
            v18 = v18
            var v113 uint = 679
            _ = v113
        }
        v18 -= (v18 << ((1) & 63))
        v111 -= (-(v111))
    }
    switch 807 {
    case 9223372036854775807:
        v18 += v18
    case 1:
        var v114 bool = (1 != 4294967295)
        _ = v114
    case 0:
        v1[2147483647] = (1 < -1)
    default:
        var v115 string = "hello"
        _ = v115
    }
    for fi116 := 0; fi116 < 13; fi116++ {
        v111 += (v111 >> ((4294967295) & 63))
        fmt.Println(v111)
        _ = fi116
    }
    v117 := GenS16{F0: 1, F1: 42, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v111
    _ = v112
    _ = v117
    _ = v18
    _ = v2
    _ = v3
}
