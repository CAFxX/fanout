package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int; F2 uint32 }
type GenS7 struct { F0 int; F1 int }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 bool; F1 uint }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 uint32; F1 int16; F2 uint }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint64 }
func genFn0(p0 uint16, p1 string) int64 {
    v4 := GenS1{F0: 42, F1: 42, F2: true}
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{0, 406}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 0, F2: false}
    if (0 < 128) {
        v8 := GenS1{F0: 1, F1: uint32(p0), F2: true}
        for fi9 := 0; fi9 < 7; fi9++ {
            p1 = p1
            v10 := []int{(fi9 - (fi9 << ((255) & 63)))}
            _ = fi9
            _ = v10
        }
        p0 = p0
        _ = v8
    } else {
        p0 += 65535
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint64 = 1
            var v13 uint32 = 65535
            var v14 uint8 = 215
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        if true {
            v15 := GenS1{F0: 2, F1: 2, F2: true}
            var v16 uint8 = 70
            _ = v15
            _ = v16
        } else {
            v17 := []int{-693, 127}
            if (0 == 1) {
                var v18 int32 = 196
                _ = v18
            }
            p1 = p1
            _ = v17
        }
    }
    v19 := GenS2{F0: 4294967295, F1: (p0 * p0), F2: true}
    _ = p0
    _ = p1
    _ = v19
    _ = v4
    _ = v7
    return 0
}

func main() {
    var v1 uint16 = 302
    var v2 float64 = 0.0
    var v3 uint64 = 4294967295
    fmt.Println(v1, v1)
    v20 := GenS2{F0: 791, F1: v1, F2: (!false)}
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := []int{(fi21 % (((-(fi21))) | 1)), fi21, (fi21 << ((uint(v3)) & 63))}
        v1 = (v1 ^ (v1 - 835))
        _ = fi21
        _ = v22
    }
    v23 := []int{}
    if (!true) {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    }
    if (!false) {
        if (128 >= 42) {
            if false {
                if (1 >= 443) {
                    v2 += 1e10
                    for fi24 := 0; fi24 < 20; fi24++ {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            fmt.Println(v3, v2, v2)
                            _ = fi25
                        }
                        v23 = append(v23, fi24)
                        switch fi24 {
                        case 9223372036854775807:
                            if true {
                                v26 := GenS1{F0: 42, F1: 42, F2: (true && true)}
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 16; fi27++ {
                                    v2 += v2
                                    _ = len(v23)
                                    _ = fi27
                                }
                                var v28 bool = false
                                _ = v28
                            }
                        }
                        _ = fi24
                    }
                } else {
                    if false {
                        var v29 int = 128
                        v30 := GenS2{F0: 4294967295, F1: (v1 / ((v1) | 1)), F2: (!(916 < 9223372036854775807))}
                        v3 += (v3 >> ((652) & 63))
                        _ = v29
                        _ = v30
                    } else {
                        v31 := GenS2{F0: 2, F1: (v1 - 2), F2: true}
                        v32 := GenS3{F0: 65535}
                        _ = v31
                        _ = v32
                    }
                    if ((v3 + v3) != 18446744073709551615) {
                        _ = len(v23)
                        var v33 uint16 = 32767
                        v34 := GenS3{F0: 65535}
                        _ = v33
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 3; fi35++ {
                            if (!(v2 < (((0.0 * v2) + -2.25) - 3.14159))) {
                                var v36 float32 = ((1.5 * 1e10) * (3.14159 * -2.25))
                                v37 := GenS4{F0: 910, F1: 42, F2: 65535}
                                _ = v36
                                _ = v37
                            } else {
                                v1 -= (v1 << ((4294967295) & 63))
                                var v38 int = -9223372036854775808
                                _ = v38
                            }
                            v23 = append(v23, 127)
                            fmt.Println(v2)
                            _ = fi35
                        }
                        v39 := GenS3{F0: 711}
                        v23 = append(v23, 109)
                        _ = v39
                    }
                    v23 = append(v23, 0)
                }
                var v40 bool = false
                if len(v23) >= 2 {
                    v23 = v23[1:len(v23)]
                }
                _ = v40
            } else {
                v41 := GenS5{F0: -2147483648}
                var v42 uint32 = 2147483647
                v2 = (3.14159 * v2)
                _ = v41
                _ = v42
            }
        } else {
            var v43 uint16 = 32767
            _ = v43
        }
        if true {
            if true {
                _ = len(v23)
                if (false && ((-2147483648 <= 9223372036854775807) || false)) {
                    for fi44 := 0; fi44 < 2; fi44++ {
                        if (false || (-2147483648 >= 0)) {
                            if false {
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    if len(v23) > 0 {
                                        v23[0] = 42
                                    }
                                default:
                                    if (87 == 61) {
                                        for fi45 := 0; fi45 < 7; fi45++ {
                                            v46 := GenS5{F0: 684}
                                            switch (("hello" + "a") + "a") {
                                            case "a":
                                                var v47 int64 = 473
                                                _ = v47
                                            }
                                            if ((v3 > (v3 & v3)) || false) {
                                                v48 := GenS4{F0: 67, F1: 42, F2: 828}
                                                _ = v48
                                            }
                                            _ = fi45
                                            _ = v46
                                        }
                                    }
                                }
                                if false {
                                    var v49 string = ("\n\t" + (("hello" + "foo bar") + "\n\t"))
                                    _ = v49
                                }
                            }
                        } else {
                            var v50 int8 = 1
                            if ((-792 <= -147) && false) {
                                if true {
                                    v2 -= v2
                                    v51 := GenS5{F0: -360}
                                    for fi52 := 0; fi52 < 11; fi52++ {
                                        var v53 float64 = v2
                                        v54 := GenS6{F0: 4294967295, F1: (fi44 | 513), F2: 2}
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                    _ = v51
                                }
                                v55 := GenS7{F0: (fi44 >> ((627) & 63)), F1: 1}
                                _ = v55
                            }
                            if ((0 == 42) || (v2 <= v2)) {
                                v50 -= (v50 + v50)
                                fmt.Println(fi44, v50)
                            }
                            _ = v50
                        }
                        var v56 string = ("a" + (("a" + "hello") + ("\n\t" + "a")))
                        _ = fi44
                        _ = v56
                    }
                }
            }
        }
        if ((!(792 <= 1)) || (!true)) {
            var v57 int32 = -338
            v1 = (^(v1))
            _ = v57
        }
    } else {
        if len(v23) > 0 {
            v23[0] = 42
        }
        if (false || true) {
            fmt.Println(v3)
        }
    }
    var v58 uint32 = 4294967295
    if false {
        v1 = (v1 & v1)
        for fi59 := 0; fi59 < 9; fi59++ {
            var v60 bool = false
            var v61 uint64 = 4294967295
            if (-981 > (^(fi59))) {
                if true {
                    var v62 int = fi59
                    fmt.Println(v2, v58)
                    if len(v23) > 0 {
                        v23[0] = (-(v62))
                    }
                    _ = v62
                }
                switch "\n\t" {
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        if false {
                            v63 := GenS8{F0: false, F1: uint(v1)}
                            _ = v63
                        }
                    case "hello":
                        if v60 {
                            for fi64 := 0; fi64 < 15; fi64++ {
                                v60 = ((!v60) || v60)
                                _ = len(v23)
                                _ = fi64
                            }
                            switch fi59 {
                            case -343:
                                v60 = (!v60)
                            case 1:
                                switch (("\n\t" + "foo bar") + "foo bar") {
                                case "":
                                    if (!(27 != 1)) {
                                        v65 := GenS9{F0: (!v60), F1: 18446744073709551615}
                                        switch fi59 {
                                        case 32767:
                                            for fi66 := 0; fi66 < 10; fi66++ {
                                                if (!(!(!(!v60)))) {
                                                    v61 -= (^(v61))
                                                    v67 := GenS10{F0: -2147483648, F1: -1}
                                                    v2 -= float64(float32((v2 * 0.0)))
                                                    _ = v67
                                                }
                                                if len(v23) > 0 {
                                                    v23[0] = fi66
                                                }
                                                _ = fi66
                                            }
                                        case -9223372036854775808:
                                            v3 = 0
                                        case 2147483647:
                                            _ = len(v23)
                                        }
                                        v3 += (v61 | uint64(v58))
                                        _ = v65
                                    } else {
                                        fmt.Println(v2, v2, fi59)
                                        var v68 uint = 255
                                        _ = v68
                                    }
                                case "foo bar":
                                    switch -9223372036854775808 {
                                    case 0:
                                        v58 -= v58
                                    case -657:
                                        var v69 int = fi59
                                        _ = v69
                                    case 9223372036854775807:
                                        var v70 bool = v60
                                        _ = v70
                                    }
                                }
                            default:
                                v60 = ("\n\t" >= ("a" + "foo bar"))
                            }
                            var v71 int16 = -355
                            _ = v71
                        }
                    case "\n\t":
                        _ = len(v23)
                    default:
                        for fi72 := 0; fi72 < 8; fi72++ {
                            switch (("\n\t" + "a") + (("foo bar" + ("\n\t" + ("a" + "hello"))) + "a")) {
                            case "foo bar":
                                v23 = append(v23, (fi72 | 32767))
                            case "\n\t":
                                var v73 float64 = v2
                                _ = v73
                            case "":
                                v1 = v1
                            }
                            _ = fi72
                        }
                    }
                case "a":
                    v61 -= 18446744073709551615
                }
                v74 := GenS11{F0: "", F1: "a"}
                _ = v74
            } else {
                v75 := GenS12{F0: v58, F1: 1, F2: 2}
                for fi76 := 0; fi76 < 2; fi76++ {
                    var v77 int16 = -1
                    if false {
                        if v60 {
                            v78 := GenS13{F0: 0}
                            if (1 != 9223372036854775807) {
                                if (-2.25 != 1.5) {
                                    fmt.Println(v3)
                                    if (fi76 > 32767) {
                                        if v60 {
                                            var v79 int64 = -9223372036854775808
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            v61 = v3
                                            _ = v79
                                        } else {
                                            if (v60 && v60) {
                                                _ = len(v23)
                                            } else {
                                                v80 := GenS14{F0: v3}
                                                fmt.Println(fi59)
                                                _ = v80
                                            }
                                            v3 = (v3 - 2)
                                        }
                                        if v60 {
                                            v61 -= 415
                                        }
                                    }
                                }
                                switch (^(fi76)) {
                                case -1:
                                    if v60 {
                                        v61 -= 2
                                        v23 = append(v23, (fi76 * fi59))
                                        if (-190 == fi59) {
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                        }
                                    }
                                default:
                                    if (549 != v77) {
                                        var v81 string = "hello"
                                        var v82 uint32 = 608
                                        _ = v81
                                        _ = v82
                                    }
                                }
                            }
                            var v83 bool = (!false)
                            _ = v78
                            _ = v83
                        }
                        var v84 int = fi76
                        var v85 uint16 = 0
                        _ = v84
                        _ = v85
                    }
                    _ = fi76
                    _ = v77
                }
                _ = v75
            }
            _ = fi59
            _ = v60
            _ = v61
        }
    } else {
        if false {
            v3 = v3
        }
        v2 -= (v2 * v2)
        v2 = (float64(3.14159) / float64((3.14159 + float32(1e10))))
    }
    if len(v23) > 0 {
        v23[0] = 529
    }
    v2 -= v2
    var v86 uint16 = v1
    for fi87 := 0; fi87 < 6; fi87++ {
        if len(v23) > 0 {
            v23[0] = fi87
        }
        _ = fi87
    }
    if (128 < 150) {
        v58 = (v58 >> ((360) & 63))
        var v88 int32 = 1
        var v89 int8 = 42
        _ = v88
        _ = v89
    }
    fmt.Println(v2)
    for fi90 := 0; fi90 < 11; fi90++ {
        var v91 float64 = 1.5
        if (true || (true || false)) {
            var v92 uint16 = (v86 - 32767)
            v93 := GenS1{F0: (v3 / ((uint64(v1)) | 1)), F1: (v58 ^ (v58 / ((42) | 1))), F2: (!(!true))}
            _ = v92
            _ = v93
        } else {
            v58 = 2
        }
        _ = fi90
        _ = v91
    }
    var v94 uint32 = v58
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v58:", v58)
    fmt.Println("v86:", v86)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v58
    _ = v86
    _ = v94
}
