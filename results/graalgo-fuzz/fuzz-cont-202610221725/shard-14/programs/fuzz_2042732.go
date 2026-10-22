package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint16; F1 int; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 bool; F2 int64 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32 }
func genFn0(p0 int, p1 string, p2 int64) int {
    v6 := GenS1{F0: (p0 | (p0 & p0))}
    switch ("foo bar" + "foo bar") {
    case "":
        for fi7 := 0; fi7 < 17; fi7++ {
            p2 = (p2 ^ -1)
            _ = fi7
        }
    }
    v8 := []int{266, (^(p0))}
    if (!true) {
        p0 -= (p0 % ((-9223372036854775808) | 1))
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return p0
}
func genFn1(p0 uint16, p1 int, p2 int16) int8 {
    var v9 uint64 = 0
    var v10 uint32 = 65535
    var v11 uint64 = (v9 & (v9 - (v9 << ((65535) & 63))))
    var v12 uint16 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -108
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []string{("hello" + "a"), "", "hello", "hello"}
    v3 := []string{"", (("a" + "hello") + "foo bar")}
    var v4 float32 = ((0.1 * ((0.0 / 1.0) + float32(-2.25))) / ((1e10 + 0.1) + ((3.14159 - -2.25) / -2.25)))
    v5 := make(map[uint32]bool)
    for fi13 := 0; fi13 < 19; fi13++ {
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = fi13
    }
    var v14 string = (("\n\t" + "") + "")
    v14 = "foo bar"
    var v15 int64 = 9223372036854775807
    var v16 bool = (false || (814 != 65535))
    if (genFn1(773, genFn0(genFn0(9223372036854775807, v14, v15), v14, v15), 42) < 1) {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int8 = 107
            var v19 string = ((v14 + (v14 + v14)) + "hello")
            switch v14 {
            case "foo bar":
                v20 := GenS1{F0: fi17}
                _ = v20
            case "a":
                var v21 uint16 = 2
                _ = v21
            case "\n\t":
                if v16 {
                    fmt.Println(v15)
                    switch ("a" + v14) {
                    case "a":
                        v14 = v19
                    case "":
                        fmt.Println(v18, v16, v19)
                    default:
                        var v22 int16 = -32768
                        _ = v22
                    }
                }
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS1{F0: 569}
                v26 := GenS2{F0: fi23, F1: 1}
                var v27 uint8 = 152
                v28 := GenS3{F0: v16, F1: (fi24 % (((^(fi24))) | 1))}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                if v16 {
                    v16 = v16
                    v30 := v5[65535]
                    fmt.Println(fi29, v30, v30)
                    _ = v30
                } else {
                    v31 := GenS1{F0: 32768}
                    var v32 int = fi29
                    _ = v31
                    _ = v32
                }
                v4 = (((v4 + v4) + float32((float64((v4 * ((3.14159 + 3.14159) + v4))) - 0.1))) * (0.0 - ((v4 - v4) + ((0.1 - v4) / float32(-2.25)))))
                var v33 bool = ((fi29 % (((fi23 << ((1) & 63))) | 1)) <= fi23)
                v34 := GenS4{F0: 255, F1: genFn0(genFn0(fi23, v14, v15), "hello", int64(fi23)), F2: -128}
                _ = fi29
                _ = v33
                _ = v34
            }
            _ = fi23
        }
        if (982 != 2147483647) {
            var v35 int8 = -67
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v35
        }
    } else {
        _ = len(v5)
        if (float64(float32(((1e10 + (float64(0.0) / 1.0)) / ((((1e10 - ((((3.14159 - 3.14159) + float64(v4)) / float64(((-2.25 - (v4 / -2.25)) - 0.1))) * 0.0)) + 0.1) - 0.0) - 0.1)))) < (-2.25 * (-2.25 - 1.5))) {
            fmt.Println(v15, v15, v16)
            if true {
                switch 4294967295 {
                case -584:
                    switch v14 {
                    case "a":
                        v36 := GenS5{F0: 381}
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 string = (v14 + v14)
                            _ = fi37
                            _ = v38
                        }
                    }
                case 2147483647:
                    v3 = append(v3, ("hello" + ""))
                case -337:
                    delete(v5, 171)
                }
                switch ("" + "") {
                case "a":
                    var v39 uint = 65535
                    _ = v39
                default:
                    v40 := GenS6{F0: -32768, F1: v16, F2: v15}
                    _ = v40
                }
                delete(v5, 334)
            }
        }
        _ = len(v5)
    }
    if v16 {
        v41 := GenS1{F0: 791}
        v42 := GenS4{F0: 1, F1: 0, F2: 0}
        _ = v41
        _ = v42
    } else {
        v16 = v16
        if (false && (!(472 == 2))) {
            for fi43 := 0; fi43 < 1; fi43++ {
                var v44 int = 32768
                _ = fi43
                _ = v44
            }
            if (0.0 < ((0.1 / -2.25) / ((3.14159 + float64(0.0)) * -2.25))) {
                for fi45 := 0; fi45 < 6; fi45++ {
                    v16 = v16
                    if v16 {
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        for fi46 := 0; fi46 < 5; fi46++ {
                            if v16 {
                                v16 = (!(v16 && (!v16)))
                                if false {
                                    var v47 float64 = 1e10
                                    _ = v47
                                } else {
                                    if v16 {
                                        if ((-2.25 > v4) || (398 == 1)) {
                                            if false {
                                                v48 := GenS3{F0: v16, F1: fi46}
                                                switch v14 {
                                                case "a":
                                                    switch ((v14 + v14) + "foo bar") {
                                                    case "":
                                                        v4 += v4
                                                    case "\n\t":
                                                        v49 := v5[65535]
                                                        fmt.Println(v16)
                                                        _ = v49
                                                    case "foo bar":
                                                        fmt.Println(fi45, v16)
                                                    }
                                                case "\n\t":
                                                    v15 = 9223372036854775807
                                                default:
                                                    fmt.Println(v4, v15, v4)
                                                }
                                                _ = v48
                                            } else {
                                                if v16 {
                                                    v14 = ((v14 + v14) + "foo bar")
                                                    if true {
                                                        v50 := GenS1{F0: 9223372036854775807}
                                                        _ = v50
                                                    } else {
                                                        var v51 int8 = -48
                                                        var v52 int64 = (^(v15))
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if true {
                                                                switch (-(fi46)) {
                                                                case 1:
                                                                    var v53 int8 = v51
                                                                    v54 := GenS7{F0: 0, F1: (v14 + "foo bar")}
                                                                    _ = v53
                                                                    _ = v54
                                                                case -9223372036854775808:
                                                                    v55 := v5[0]
                                                                    _ = v55
                                                                case 9223372036854775807:
                                                                    v56 := GenS7{F0: 2, F1: "foo bar"}
                                                                    _ = v56
                                                                }
                                                                v57 := GenS5{F0: 1}
                                                                v5[4294967295] = (!(true || ((((v4 / (v4 / v4)) / float32((((float64(v4) / 1.5) + -2.25) + float64(float32((3.14159 * float64(v4))))))) - 1e10) > -2.25)))
                                                                v16 = (false || (59 <= 255))
                                                                _ = v57
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(v51, fi45, v16)
                                                        case "hello":
                                                            fmt.Println(v51, v51, v14)
                                                        default:
                                                            v58 := v5[8]
                                                            _ = v58
                                                        }
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    delete(v5, 262)
                                                    if v16 {
                                                        var v59 int64 = v15
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v60 uint64 = 548
                                                        v4 -= v4
                                                        v5[2] = ((!(!v16)) && (48 >= -128))
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    v4 -= float32(float64(0.1))
                                                    fmt.Println(v14)
                                                    var v61 int = fi45
                                                    _ = v61
                                                } else {
                                                    switch fi46 {
                                                    case -2147483648:
                                                        if (!(!v16)) {
                                                            var v62 int = fi45
                                                            v16 = v16
                                                            _ = v62
                                                        } else {
                                                            if (127 == 237) {
                                                                switch v14 {
                                                                case "a":
                                                                    switch fi46 {
                                                                    case 155:
                                                                        if false {
                                                                            v63 := GenS8{F0: v14}
                                                                            _ = v63
                                                                        }
                                                                    case -1:
                                                                        v16 = ((!(42 < 127)) || ((^(v15)) >= v15))
                                                                    case 2147483647:
                                                                        v14 = v14
                                                                    default:
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case 32767:
                                                        if ((!v16) && true) {
                                                            var v64 int16 = -32768
                                                            _ = v64
                                                        } else {
                                                            var v65 int8 = 1
                                                            var v66 uint64 = 2
                                                            if len(v2) > 0 {
                                                                v2[0] = ((v14 + "foo bar") + v14)
                                                            }
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    }
                                                    var v67 string = v14
                                                    _ = v67
                                                }
                                                v1 = append(v1, 2)
                                            }
                                        } else {
                                            if (299 > 65535) {
                                                fmt.Println(v4, v14)
                                                fmt.Println(v15)
                                                var v68 uint = 4294967295
                                                fmt.Println(v14, v15)
                                                if len(v1) > 0 {
                                                    v1[0] = 271
                                                }
                                                fmt.Println(fi46)
                                                _ = v68
                                            } else {
                                                var v69 int8 = 0
                                                v14 = ""
                                                _ = v69
                                            }
                                            if v16 {
                                                v14 = v14
                                                fmt.Println(v4, v14)
                                                v16 = (1 >= 42)
                                            }
                                            if (!(v16 || v16)) {
                                                var v70 uint8 = 255
                                                v14 = "\n\t"
                                                var v71 uint32 = 28
                                                _ = v70
                                                _ = v71
                                            } else {
                                                if ((v15 / (((-(v15))) | 1)) >= (v15 - (v15 & (v15 / ((v15) | 1))))) {
                                                    if v16 {
                                                        if v16 {
                                                            if (42 > 174) {
                                                                v14 = v14
                                                                var v72 int32 = -1
                                                                var v73 bool = (fi45 < genFn0(fi46, "hello", (v15 ^ v15)))
                                                                _ = v72
                                                                _ = v73
                                                            } else {
                                                                _ = len(v5)
                                                            }
                                                            var v74 float32 = (0.1 + v4)
                                                            v16 = v16
                                                            _ = v74
                                                        } else {
                                                            var v75 bool = v16
                                                            v14 = ("foo bar" + v14)
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if false {
                                                                var v76 int = 0
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v76
                                                            } else {
                                                                if v16 {
                                                                    switch (("" + v14) + "a") {
                                                                    case "a":
                                                                        v16 = v16
                                                                    case "":
                                                                        switch fi46 {
                                                                        case 32767:
                                                                            v15 -= v15
                                                                            switch 963 {
                                                                            case -779:
                                                                                fmt.Println(fi45, v14)
                                                                            case -570:
                                                                                var v77 int64 = (v15 << ((42) & 63))
                                                                                _ = v77
                                                                            default:
                                                                                var v78 uint = 1
                                                                                _ = v78
                                                                            }
                                                                            if ((float64((v4 - 3.14159)) * 0.1) > ((((((0.1 / -2.25) * 3.14159) + 1.5) / 1.0) + -2.25) + 0.0)) {
                                                                                v79 := GenS9{F0: 255}
                                                                                _ = v79
                                                                            }
                                                                        default:
                                                                            v16 = true
                                                                        }
                                                                    case "\n\t":
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            if false {
                                                                                var v80 int8 = 127
                                                                                var v81 string = ("a" + "hello")
                                                                                var v82 uint8 = 66
                                                                                _ = v80
                                                                                _ = v81
                                                                                _ = v82
                                                                            } else {
                                                                                var v83 int8 = -128
                                                                                _ = v83
                                                                            }
                                                                        default:
                                                                            var v84 float32 = v4
                                                                            _ = v84
                                                                        }
                                                                    default:
                                                                        v15 -= 1
                                                                    }
                                                                    v85 := GenS4{F0: 65535, F1: genFn0(4294967295, ((v14 + "") + ("a" + (v14 + "hello"))), (v15 + -535)), F2: 100}
                                                                    _ = v85
                                                                }
                                                                if (85 <= 122) {
                                                                    if (574 > 140) {
                                                                        var v86 float64 = 0.0
                                                                        _ = v86
                                                                    } else {
                                                                        v14 = (v14 + ("a" + v14))
                                                                    }
                                                                    fmt.Println(v15, v15)
                                                                    v87 := GenS10{F0: 1}
                                                                    _ = v87
                                                                } else {
                                                                    if (0.1 != (-2.25 / 3.14159)) {
                                                                        fmt.Println(fi46, v75, v75)
                                                                    } else {
                                                                        v15 -= (v15 % (((v15 + v15)) | 1))
                                                                        switch (v14 + v14) {
                                                                        case "":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = v14
                                                                            }
                                                                        case "foo bar":
                                                                            v4 -= 3.14159
                                                                        case "a":
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                        default:
                                                                            var v88 float32 = v4
                                                                            if v75 {
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    v89 := GenS10{F0: 2}
                                                                                    _ = v89
                                                                                case "a":
                                                                                    switch 654 {
                                                                                    case -773:
                                                                                        var v90 uint64 = 2
                                                                                        _ = v90
                                                                                    case 9223372036854775807:
                                                                                        if v75 {
                                                                                            if (!v16) {
                                                                                                v4 = (float32(0.1) - (v88 + (v4 + v4)))
                                                                                            } else {
                                                                                                v16 = false
                                                                                                if false {
                                                                                                    v91 := GenS3{F0: (-1 == genFn1(42, fi46, -1)), F1: genFn0(9223372036854775807, v14, (v15 & (^(v15))))}
                                                                                                    v92 := v5[65535]
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    _ = v91
                                                                                                    _ = v92
                                                                                                }
                                                                                            }
                                                                                            v93 := GenS11{F0: 42}
                                                                                            v94 := GenS11{F0: -128}
                                                                                            v4 += float32(0.1)
                                                                                            var v95 bool = (!(!(2 == 0)))
                                                                                            _ = v93
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        } else {
                                                                                            var v96 bool = v75
                                                                                            v96 = v16
                                                                                            _ = v96
                                                                                        }
                                                                                    case -9223372036854775808:
                                                                                        _ = len(v5)
                                                                                    }
                                                                                case "hello":
                                                                                    v97 := GenS11{F0: 100}
                                                                                    _ = v97
                                                                                }
                                                                            }
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                    var v98 string = ("foo bar" + ((v14 + v14) + (("hello" + "\n\t") + (("\n\t" + v14) + v14))))
                                                                    if (18446744073709551615 == 4294967295) {
                                                                        if v16 {
                                                                            fmt.Println(v14, v16)
                                                                        } else {
                                                                            var v99 int16 = -1
                                                                            var v100 int = fi45
                                                                            fmt.Println(v15, v4)
                                                                            _ = v100
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                    _ = v98
                                                                }
                                                                _ = len(v1)
                                                                v75 = (0 >= 18446744073709551615)
                                                            }
                                                            switch fi46 {
                                                            case 2147483647:
                                                                var v101 int16 = -32768
                                                                _ = v101
                                                            case 128:
                                                                var v102 int32 = -2147483648
                                                                _ = v102
                                                            }
                                                            _ = v75
                                                        }
                                                        var v103 int = (fi46 ^ (fi45 | genFn0((fi46 + 9223372036854775807), v14, -1)))
                                                        switch v14 {
                                                        case "\n\t":
                                                            _ = len(v5)
                                                        }
                                                        v14 = "a"
                                                        _ = v103
                                                    }
                                                    v104 := v5[256]
                                                    fmt.Println(fi46, fi45, fi45)
                                                    _ = v104
                                                } else {
                                                    var v105 bool = v16
                                                    _ = v105
                                                }
                                                switch (fi45 + fi45) {
                                                case 127:
                                                    if (0 != 1) {
                                                        v16 = v16
                                                        v106 := GenS1{F0: (fi46 / ((genFn0((fi45 << ((65535) & 63)), (v14 + v14), (v15 * -1))) | 1))}
                                                        _ = v106
                                                    }
                                                case 4294967295:
                                                    switch v14 {
                                                    case "\n\t":
                                                        v107 := GenS10{F0: 0}
                                                        v5[2] = (-2.25 <= 0.1)
                                                        _ = v107
                                                    case "":
                                                        switch 127 {
                                                        case -9223372036854775808:
                                                            if ((42 < 2147483647) && (!(54 > 1))) {
                                                                fmt.Println(v14)
                                                                _ = len(v5)
                                                            }
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                    }
                                                }
                                                v108 := GenS11{F0: 127}
                                                _ = v108
                                            }
                                            var v109 uint64 = 2
                                            _ = v109
                                        }
                                    } else {
                                        switch v14 {
                                        case "\n\t":
                                            var v110 uint16 = 541
                                            _ = v110
                                        case "a":
                                            v14 = (("" + "a") + "")
                                        }
                                    }
                                    v111 := GenS12{F0: 1}
                                    var v112 bool = (v16 && (2 <= 0))
                                    _ = v111
                                    _ = v112
                                }
                                switch fi45 {
                                case 1:
                                    if (0 >= -1) {
                                        if (1 == 1) {
                                            v15 -= -955
                                            v14 = ("\n\t" + v14)
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v113 uint64 = 0
                                                _ = v113
                                            case "a":
                                                if v16 {
                                                    v3 = append(v3, ("foo bar" + ((v14 + (((v14 + v14) + ("\n\t" + (v14 + ("foo bar" + ("" + v14))))) + (v14 + ((v14 + v14) + v14)))) + v14)))
                                                    var v114 int64 = v15
                                                    if (true && v16) {
                                                        switch (fi45 / ((genFn0(fi46, ("" + ""), v114)) | 1)) {
                                                        case -316:
                                                            delete(v5, 0)
                                                        }
                                                    }
                                                    _ = v114
                                                } else {
                                                    v4 = v4
                                                }
                                            case "":
                                                var v115 float64 = 1.5
                                                _ = v115
                                            }
                                        } else {
                                            fmt.Println(v14)
                                            v14 = ("" + ((v14 + v14) + v14))
                                        }
                                    }
                                case 9223372036854775807:
                                    if ((255 < 192) && false) {
                                        _ = len(v2)
                                        fmt.Println(fi46, v14)
                                    } else {
                                        v16 = (-23 > 42)
                                        if (v14 > ((v14 + v14) + v14)) {
                                            v16 = true
                                            var v116 int32 = -961
                                            if (18446744073709551615 >= 873) {
                                                v116 -= (-(v116))
                                            }
                                            _ = v116
                                        }
                                        v4 += 0.1
                                        var v117 int16 = 42
                                        _ = v117
                                    }
                                case -829:
                                    _ = len(v1)
                                }
                            } else {
                                var v118 int64 = v15
                                switch v14 {
                                case "hello":
                                    delete(v5, 65535)
                                case "foo bar":
                                    fmt.Println(v4)
                                }
                                fmt.Println(v15)
                                _ = v118
                            }
                            fmt.Println(v4, v16)
                            v4 += (1.5 + float32(-2.25))
                            _ = fi46
                        }
                        v16 = (v15 > v15)
                    }
                    _ = fi45
                }
                v3 = append(v3, (v14 + "hello"))
                fmt.Println(v4, v14)
            }
            v16 = v16
        }
    }
    for fi119 := 0; fi119 < 12; fi119++ {
        v14 = v14
        _ = fi119
    }
    var v120 int64 = v15
    v120 -= v15
    v5[4294967295] = (!true)
    v16 = (!true)
    fmt.Println(v14)
    v15 = v120
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v120:", v120)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v120
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
