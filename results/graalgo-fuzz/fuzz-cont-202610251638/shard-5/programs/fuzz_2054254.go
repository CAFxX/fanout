package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint; F1 bool }
type GenS15 struct { F0 uint32; F1 int }
type GenS16 struct { F0 int; F1 uint }
func genFn0(p0 uint) int16 {
    var v6 bool = (!true)
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            if (20 > 108) {
                var v8 uint = 2
                for fi9 := 0; fi9 < 10; fi9++ {
                    if ((42 <= 42) || v6) {
                        fmt.Println(v6, fi9)
                        v10 := []int{}
                        var v11 uint = (p0 << ((65535) & 63))
                        _ = v10
                        _ = v11
                    }
                    var v12 uint8 = 1
                    v6 = false
                    v12 += (^(v12))
                    _ = fi9
                    _ = v12
                }
                _ = v8
            }
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 bool = ((v6 && v6) || true)
                v15 := []int{}
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 int32 = 42
            _ = v16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            if (2 <= 65535) {
                v18 := []int{fi17}
                fmt.Println(fi17, v6, p0)
                _ = v18
            } else {
                var v19 uint16 = 255
                _ = v19
            }
            v20 := GenS1{F0: 65535}
            v21 := []int{(fi17 >> (((p0 & (p0 | 4294967295))) & 63)), (fi7 * (fi7 | -1)), (fi17 | fi7), fi17}
            v22 := GenS1{F0: 1}
            if false {
                switch "hello" {
                case "a":
                    v23 := GenS2{F0: 0, F1: 42, F2: 1}
                    _ = v23
                default:
                    v6 = true
                }
            } else {
                if (0 != -9223372036854775808) {
                    if true {
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                    } else {
                        v24 := GenS2{F0: 42, F1: 228, F2: 2}
                        _ = v24
                    }
                    v25 := GenS2{F0: 127, F1: 2147483647, F2: 136}
                    _ = v25
                }
                p0 -= p0
                switch "\n\t" {
                case "\n\t":
                    v6 = v6
                case "a":
                    if (fi7 <= int(p0)) {
                        var v26 int16 = 0
                        switch ("\n\t" + "a") {
                        case "hello":
                            if len(v21) > 0 {
                                v21[0] = (fi17 * (^(fi17)))
                            }
                        case "foo bar":
                            fmt.Println(v6, v26, fi17)
                        }
                        v27 := GenS1{F0: 2}
                        _ = v26
                        _ = v27
                    }
                }
            }
            _ = fi17
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi7
    }
    v28 := []int{404, -9223372036854775808}
    v29 := GenS1{F0: 32767}
    for fi30 := 0; fi30 < 17; fi30++ {
        if (65535 < 543) {
            v31 := GenS1{F0: 388}
            _ = v31
        } else {
            for fi32 := 0; fi32 < 20; fi32++ {
                if (42 != 20) {
                    v33 := GenS3{F0: "a"}
                    _ = v33
                }
                if v6 {
                    if true {
                        var v34 uint64 = 0
                        fmt.Println(fi30, fi30, fi30)
                        _ = v34
                    } else {
                        v28 = append(v28, fi30)
                    }
                    v28 = append(v28, (fi32 / (((fi32 / ((-1) | 1))) | 1)))
                    var v35 float64 = 1e10
                    if v6 {
                        if (-32768 != 401) {
                            var v36 uint64 = 600
                            if (1 <= 111) {
                                if v6 {
                                    v37 := GenS1{F0: 42}
                                    var v38 int16 = 42
                                    _ = v37
                                    _ = v38
                                } else {
                                    var v39 int = (fi32 / (((fi32 / (((fi30 ^ (fi32 - fi32))) | 1))) | 1))
                                    if (36 == 127) {
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(p0, fi32)
                                        case "foo bar":
                                            if len(v28) > 0 {
                                                v28[0] = (fi32 * 0)
                                            }
                                            v40 := GenS4{F0: 1, F1: 960}
                                            _ = v40
                                        case "hello":
                                            if v6 {
                                                v41 := GenS3{F0: ("foo bar" + "a")}
                                                v36 += (v36 - 57)
                                                if (v6 && (v6 || v6)) {
                                                    var v42 uint16 = 1
                                                    if len(v28) > 0 {
                                                        v28[0] = (v39 ^ fi32)
                                                    }
                                                    _ = v42
                                                }
                                                _ = v41
                                            }
                                            v43 := GenS1{F0: 255}
                                            var v44 int32 = -504
                                            if v6 {
                                                v36 = 0
                                            } else {
                                                p0 += p0
                                                var v45 string = ""
                                                _ = v45
                                            }
                                            v46 := GenS5{F0: 1}
                                            _ = v43
                                            _ = v44
                                            _ = v46
                                        }
                                        fmt.Println(v6, fi30)
                                    } else {
                                        fmt.Println(v39, fi30)
                                        v35 -= 0.0
                                        if (!(v6 && v6)) {
                                            v6 = ((0.0 / (1e10 * (1.5 + -2.25))) < ((1e10 - 1e10) - 0.0))
                                            switch "" {
                                            case "foo bar":
                                                if len(v28) >= 2 {
                                                    v28 = v28[1:len(v28)]
                                                }
                                            }
                                        } else {
                                            var v47 int16 = 0
                                            if len(v28) >= 2 {
                                                v28 = v28[1:len(v28)]
                                            }
                                            _ = v47
                                        }
                                        if v6 {
                                            var v48 int32 = -1
                                            _ = v48
                                        }
                                    }
                                    var v49 uint16 = 2
                                    _ = v39
                                    _ = v49
                                }
                            }
                            if (!((-128 != int8(fi30)) || (v6 || v6))) {
                                var v50 uint8 = 42
                                _ = v50
                            } else {
                                p0 = (p0 % (((p0 & 525)) | 1))
                                if (v6 && ((1e10 <= float64((1.5 * (0.1 * -2.25)))) || false)) {
                                    v35 = 0.1
                                    if (float32(float64(((3.14159 * 1.5) + 1e10))) == 1e10) {
                                        var v51 int32 = -1
                                        switch "a" {
                                        case "a":
                                            v51 = v51
                                        case "":
                                            v6 = v6
                                        case "foo bar":
                                            v6 = (962 == 353)
                                        }
                                        _ = v51
                                    }
                                    _ = len(v28)
                                }
                            }
                            _ = v36
                        }
                    }
                    var v52 uint64 = 18446744073709551615
                    v6 = (v6 || true)
                    _ = v35
                    _ = v52
                } else {
                    switch 2147483647 {
                    case -768:
                        v53 := GenS6{F0: (p0 % ((p0) | 1))}
                        _ = v53
                    default:
                        if (("hello" + "\n\t") > ("\n\t" + ("\n\t" + ("hello" + ("\n\t" + "foo bar"))))) {
                            v54 := GenS2{F0: 1, F1: 2, F2: 42}
                            _ = v54
                        }
                    }
                    switch "foo bar" {
                    case "":
                        fmt.Println(p0)
                        fmt.Println(fi32)
                        v6 = (0 >= 0)
                    case "\n\t":
                        if v6 {
                            switch "a" {
                            case "":
                                if len(v28) > 0 {
                                    v28[0] = (-(fi32))
                                }
                            case "\n\t":
                                v55 := GenS1{F0: 625}
                                _ = v55
                            case "a":
                                v56 := GenS7{F0: 65535, F1: (fi32 ^ 4294967295)}
                                _ = v56
                            default:
                                switch "a" {
                                case "foo bar":
                                    if v6 {
                                        v57 := GenS8{F0: (p0 << ((952) & 63)), F1: 0}
                                        _ = v57
                                    } else {
                                        switch (^(fi32)) {
                                        case 9223372036854775807:
                                            if false {
                                                fmt.Println(fi32, fi30)
                                                var v58 int64 = 0
                                                v59 := GenS9{F0: 1, F1: -631, F2: 4294967295}
                                                _ = v58
                                                _ = v59
                                            }
                                        }
                                        v60 := GenS9{F0: 255, F1: -32768, F2: 1}
                                        if (1.5 >= ((0.1 - 1.5) + (1.5 * -2.25))) {
                                            p0 += p0
                                            switch -167 {
                                            case 32767:
                                                v61 := GenS4{F0: -793, F1: (p0 | p0)}
                                                _ = v61
                                            }
                                        } else {
                                            if ("foo bar" == "") {
                                                var v62 int = 32768
                                                _ = v62
                                            }
                                        }
                                        _ = v60
                                    }
                                    if v6 {
                                        fmt.Println(fi30)
                                        v6 = ("a" > "hello")
                                        v63 := GenS1{F0: 425}
                                        _ = v63
                                    }
                                    if len(v28) > 0 {
                                        v28[0] = 817
                                    }
                                default:
                                    v64 := GenS10{F0: (p0 * 18446744073709551615)}
                                    _ = v64
                                }
                            }
                            if ((-996 != -451) && v6) {
                                var v65 uint8 = 200
                                _ = v65
                            }
                            if (((v6 && v6) && v6) || (42 > 9223372036854775807)) {
                                var v66 int64 = 1
                                _ = v66
                            }
                        }
                    default:
                        switch fi30 {
                        case 32768:
                            if (((((1.5 / (1e10 * ((0.0 + -2.25) * 0.1))) - 0.1) + 0.1) * ((0.0 - -2.25) + 0.0)) <= 0.0) {
                                v6 = ("" != ("foo bar" + "hello"))
                                v67 := GenS11{F0: 2147483647}
                                _ = v67
                            }
                        case 1:
                            v6 = (!v6)
                        }
                    }
                }
                if v6 {
                    if (v6 || v6) {
                        p0 += 65535
                    }
                    v6 = v6
                } else {
                    p0 = (p0 ^ p0)
                    if (9223372036854775807 > -594) {
                        v6 = (284 == 4294967295)
                        switch (fi30 >> (((p0 % ((18446744073709551615) | 1))) & 63)) {
                        case 128:
                            v68 := GenS12{F0: (("hello" + "foo bar") + "foo bar"), F1: ""}
                            _ = v68
                        case 4294967295:
                            v28 = append(v28, (fi30 & (^(fi30))))
                        default:
                            v69 := GenS2{F0: 0, F1: 4294967295, F2: 65535}
                            _ = v69
                        }
                        _ = len(v28)
                    }
                }
                _ = fi32
            }
            v6 = false
            p0 = p0
        }
        _ = fi30
    }
    _ = p0
    _ = v28
    _ = v29
    _ = v6
    return 0
}

func main() {
    var v1 uint = 895
    var v2 bool = false
    var v3 uint = (^(v1))
    v4 := []string{"a"}
    v5 := []string{}
    v1 -= v3
    if (0 == 574) {
        for fi70 := 0; fi70 < 19; fi70++ {
            if (((18446744073709551615 == 1) && (391 > 65535)) && false) {
                v1 = (v3 >> ((v3) & 63))
                v1 -= v1
                v3 -= v1
            }
            _ = fi70
        }
        for fi71 := 0; fi71 < 3; fi71++ {
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            if (v1 != v3) {
                if v2 {
                    var v72 int8 = 0
                    _ = v72
                }
                v73 := GenS4{F0: 0, F1: (^(v3))}
                _ = v73
            } else {
                var v74 uint16 = 1
                for fi75 := 0; fi75 < 9; fi75++ {
                    var v76 int16 = -675
                    switch (fi75 >> (((v3 + v1)) & 63)) {
                    case 128:
                        var v77 uint = (^(v3))
                        v78 := GenS13{F0: -200}
                        v5 = append(v5, "\n\t")
                        _ = v77
                        _ = v78
                    case 127:
                        var v79 float64 = 1e10
                        _ = v79
                    default:
                        v80 := GenS14{F0: (v3 & (v3 * v3)), F1: ("a" >= "foo bar")}
                        _ = v80
                    }
                    _ = fi75
                    _ = v76
                }
                _ = v74
            }
            switch (fi71 & fi71) {
            case 32768:
                v3 += (^(v3))
            default:
                var v81 uint = (v3 ^ 246)
                _ = v81
            }
            _ = fi71
        }
        v82 := GenS15{F0: 2147483648, F1: 128}
        _ = v82
    }
    _ = len(v4)
    if v2 {
        fmt.Println(v2)
        fmt.Println(v2, v2, v3)
    } else {
        var v83 string = "\n\t"
        for fi84 := 0; fi84 < 6; fi84++ {
            fmt.Println(v3, v83, v1)
            _ = fi84
        }
        _ = v83
    }
    switch 1 {
    case -2147483648:
        for fi85 := 0; fi85 < 19; fi85++ {
            v1 = (v1 % ((4294967295) | 1))
            _ = fi85
        }
    }
    v86 := GenS14{F0: 1, F1: (v2 && false)}
    v2 = ("" >= "a")
    var v87 bool = false
    v2 = true
    if false {
        v3 -= (v3 / (((v3 % ((949) | 1))) | 1))
        fmt.Println(v3)
        for fi88 := 0; fi88 < 14; fi88++ {
            var v89 int64 = -934
            var v90 uint16 = 967
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = fi88
            _ = v89
            _ = v90
        }
    }
    if len(v5) > 0 {
        v5[0] = "hello"
    }
    switch ("a" + "a") {
    case "foo bar":
        var v91 uint16 = 0
        _ = v91
    case "":
        if false {
            var v92 float32 = -2.25
            switch "foo bar" {
            case "":
                v4 = append(v4, ("foo bar" + (("hello" + ("\n\t" + "hello")) + ("" + "foo bar"))))
            case "hello":
                var v93 string = "\n\t"
                _ = v93
            case "a":
                v1 -= v1
            default:
                if false {
                    v5 = append(v5, ("" + ("hello" + "a")))
                }
            }
            _ = v92
        }
    }
    fmt.Println(v2)
    if v2 {
        if (127 < 32767) {
            switch "" {
            case "a":
                v3 = v3
            case "\n\t":
                var v94 uint = (v1 - (v3 | (^(v3))))
                _ = v94
            case "hello":
                switch 303 {
                case 42:
                    v1 = v3
                }
            }
            if (!(!false)) {
                var v95 uint = 956
                v87 = false
                if (!(42 < 542)) {
                    v96 := GenS15{F0: 2147483647, F1: 4294967295}
                    _ = v96
                } else {
                    v3 = v95
                    v3 -= (v3 | v95)
                }
                _ = v95
            } else {
                var v97 int = -2147483648
                v98 := GenS7{F0: 0, F1: (v97 << ((v3) & 63))}
                _ = v97
                _ = v98
            }
            if (1 != -32768) {
                if ((v2 || (1e10 >= 3.14159)) || (5 == 57)) {
                    v2 = (255 > v3)
                }
                fmt.Println(v87)
            }
        } else {
            fmt.Println(v3, v1, v87)
            fmt.Println(v2, v3)
            v99 := GenS5{F0: 0}
            _ = v99
        }
        switch -2147483648 {
        case -1:
            v3 += (v3 / ((4294967295) | 1))
        case 9223372036854775807:
            v100 := GenS4{F0: -224, F1: 42}
            _ = v100
        case -467:
            switch "a" {
            case "foo bar":
                var v101 bool = v2
                _ = v101
            case "\n\t":
                v102 := GenS14{F0: v3, F1: (42 == 591)}
                _ = v102
            case "":
                var v103 int8 = 45
                _ = v103
            default:
                v2 = false
            }
        default:
            var v104 int = 127
            _ = v104
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if ((!(((252 != 95) || ((!(0 != 0)) && v2)) || v2)) || (float32((0.1 / 3.14159)) < 3.14159)) {
        v105 := GenS16{F0: 42, F1: (v3 + v1)}
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
}
