package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 uint64; F1 uint32; F2 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 int; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64; F1 int16 }
type GenS11 struct { F0 int32 }
func genFn0(p0 uint32) int64 {
    var v6 int8 = 100
    switch 9223372036854775807 {
    case 9223372036854775807:
        fmt.Println(v6)
    case -765:
        v7 := GenS1{F0: -2147483648, F1: 32767, F2: ("" + "a")}
        _ = v7
    }
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return -783
}
func genFn1(p0 int16, p1 uint32) int16 {
    var v8 uint8 = 0
    var v9 int64 = genFn0(65535)
    v10 := GenS2{F0: "", F1: 0}
    var v11 int16 = p0
    if (((!(false || false)) || false) && true) {
        var v12 uint16 = 50
        p0 += -1
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := []int{fi13, fi13}
            if len(v14) > 0 {
                v14[0] = fi13
            }
            var v15 float32 = 1.5
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        v11 -= v11
        var v16 uint32 = 2147483648
        v17 := GenS1{F0: -1, F1: 0, F2: "foo bar"}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (^(v11))
}

func main() {
    v1 := []string{"\n\t", "foo bar", (("a" + ("" + "\n\t")) + (("foo bar" + "foo bar") + "a"))}
    var v2 int32 = 2147483647
    v3 := make(map[int]int)
    var v4 int16 = 42
    v5 := make(map[int]int)
    v18 := GenS3{F0: 0, F1: 194}
    switch ("" + "") {
    case "foo bar":
        fmt.Println(v4, v4)
    case "a":
        v2 = (v2 % ((v2) | 1))
    default:
        v19 := GenS2{F0: "a", F1: 42}
        _ = v19
    }
    if true {
        v4 += (v4 >> ((255) & 63))
    }
    switch 2147483647 {
    case 32767:
        for fi20 := 0; fi20 < 16; fi20++ {
            v4 += (v4 - 32767)
            _ = fi20
        }
    case 4294967295:
        var v21 int = -969
        _ = v21
    case 127:
        delete(v3, 32768)
    default:
        fmt.Println(v4, v4)
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint = 129
            _ = fi23
            _ = v24
        }
        fmt.Println(fi22, fi22, v4)
        switch (-(fi22)) {
        case 128:
            var v25 int = fi22
            _ = v25
        default:
            for fi26 := 0; fi26 < 10; fi26++ {
                if (!(!false)) {
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            var v27 int32 = v2
                            _ = v27
                        } else {
                            if (!(1 <= (v2 + (v2 >> ((42) & 63))))) {
                                var v28 int32 = (-(v2))
                                v4 = genFn1((v4 * -1), 2147483647)
                                var v29 float32 = 0.1
                                _ = v28
                                _ = v29
                            }
                        }
                    case "foo bar":
                        fmt.Println(fi26, fi26)
                    case "a":
                        fmt.Println(v2)
                    }
                    var v30 uint = 215
                    _ = v30
                } else {
                    switch fi22 {
                    case -347:
                        if true {
                            v31 := GenS3{F0: 205, F1: 216}
                            _ = v31
                        } else {
                            fmt.Println(fi26, fi26, v4)
                            v4 -= genFn1((v4 ^ v4), 1)
                            var v32 bool = false
                            _ = v32
                        }
                    case 32767:
                        var v33 uint64 = 1
                        _ = v33
                    case 42:
                        var v34 int16 = (v4 / ((1) | 1))
                        _ = v34
                    default:
                        fmt.Println(fi22)
                    }
                }
                v35 := GenS1{F0: (fi22 << ((255) & 63)), F1: fi26, F2: ("hello" + "")}
                _ = len(v1)
                switch ("hello" + "foo bar") {
                case "a":
                    switch ("hello" + ((("\n\t" + "hello") + ("" + "foo bar")) + (("a" + (((("hello" + "a") + "a") + "") + ((("hello" + (("\n\t" + "\n\t") + ("foo bar" + "foo bar"))) + "a") + ("\n\t" + "\n\t")))) + ((("a" + "a") + "hello") + (("" + "") + ""))))) {
                    case "hello":
                        var v36 int64 = -9223372036854775808
                        switch 1 {
                        case 881:
                            if ((0.1 - 1e10) != (0.0 / (((3.14159 / 1.5) / 1.0) + 0.0))) {
                                if true {
                                    if (!false) {
                                        if false {
                                            var v37 int32 = -209
                                            fmt.Println(v2, v37)
                                            _ = v37
                                        }
                                    } else {
                                        switch fi22 {
                                        case 127:
                                            _ = len(v1)
                                        }
                                        var v38 int16 = v4
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                        _ = v38
                                    }
                                } else {
                                    _ = len(v1)
                                }
                            } else {
                                fmt.Println(v2)
                            }
                        case -700:
                            v1 = append(v1, "\n\t")
                        default:
                            if false {
                                fmt.Println(v36)
                            }
                        }
                        switch (-(fi26)) {
                        case 42:
                            var v39 uint32 = 136
                            _ = v39
                        default:
                            v40 := v3[fi26]
                            _ = v40
                        }
                        _ = v36
                    }
                case "foo bar":
                    if false {
                        v41 := v3[314]
                        switch "\n\t" {
                        case "hello":
                            v4 = (-(v4))
                        case "":
                            delete(v3, (v41 >> ((571) & 63)))
                        case "foo bar":
                            var v42 int32 = (v2 - (v2 ^ v2))
                            _ = v42
                        default:
                            if ((!false) && false) {
                                if ("foo bar" != "a") {
                                    var v43 bool = (!false)
                                    var v44 bool = (2147483647 == fi26)
                                    _ = len(v3)
                                    _ = v43
                                    _ = v44
                                }
                                v2 = 42
                            } else {
                                var v45 uint8 = 0
                                switch "" {
                                case "\n\t":
                                    var v46 float32 = (1.5 + -2.25)
                                    _ = v46
                                case "hello":
                                    v1 = append(v1, "a")
                                }
                                _ = v45
                            }
                        }
                        _ = v41
                    } else {
                        var v47 uint32 = 338
                        if false {
                            if false {
                                v4 -= (v4 % (((v4 - genFn1(v4, (v47 | 988)))) | 1))
                                v47 = (v47 << ((609) & 63))
                            }
                            v48 := GenS1{F0: (fi26 >> ((uint(v47)) & 63)), F1: 719, F2: ""}
                            _ = len(v3)
                            _ = v48
                        } else {
                            var v49 bool = true
                            var v50 int64 = 9223372036854775807
                            v5[-1] = int(v4)
                            _ = v49
                            _ = v50
                        }
                        var v51 uint16 = 0
                        v52 := GenS1{F0: fi26, F1: fi26, F2: "a"}
                        if false {
                            fmt.Println(fi22, v4, v47)
                            switch 163 {
                            case 586:
                                v2 = v2
                            default:
                                if len(v1) > 0 {
                                    v1[0] = ("a" + "\n\t")
                                }
                            }
                            v53 := GenS3{F0: (^(v47)), F1: 255}
                            _ = v53
                        } else {
                            var v54 uint64 = 18446744073709551615
                            if (3.14159 < 1e10) {
                                switch "\n\t" {
                                case "\n\t":
                                    v2 = v2
                                }
                                if false {
                                    if (false || true) {
                                        var v55 int8 = 1
                                        v54 += v54
                                        fmt.Println(fi26, v55, v54)
                                        _ = v55
                                    }
                                    var v56 uint = 4294967295
                                    _ = v56
                                } else {
                                    if (true && ((v51 >> ((986) & 63)) >= (v51 % (((^(v51))) | 1)))) {
                                        if ((1e10 / 3.14159) != (0.1 + 0.0)) {
                                            var v57 string = "foo bar"
                                            if ((fi22 >> ((65535) & 63)) == 32768) {
                                                if ((v4 & genFn1((v4 ^ -1), (v47 << ((65535) & 63)))) >= v4) {
                                                    if true {
                                                        v58 := GenS4{F0: 0}
                                                        v57 = "\n\t"
                                                        var v59 int64 = -343
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        fmt.Println(v54)
                                                        var v60 uint64 = v54
                                                        v61 := GenS5{F0: (genFn0(v47) == 9223372036854775807), F1: (!(!(((v51 ^ 32767) > (v51 + (^(v51)))) || true)))}
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                    fmt.Println(v54, fi22, v54)
                                                } else {
                                                    var v62 int64 = 0
                                                    if false {
                                                        switch (v57 + v57) {
                                                        case "":
                                                            var v63 int8 = 42
                                                            switch (fi22 - (^(fi22))) {
                                                            case 32768:
                                                                switch "" {
                                                                case "foo bar":
                                                                    switch v57 {
                                                                    case "":
                                                                        v51 += v51
                                                                    case "foo bar":
                                                                        v47 -= (v47 - 134)
                                                                    }
                                                                default:
                                                                    v4 += genFn1(int16(v47), (v47 / ((v47) | 1)))
                                                                }
                                                            default:
                                                                var v64 bool = (v63 != (-(v63)))
                                                                _ = v64
                                                            }
                                                            _ = v63
                                                        case "a":
                                                            if (((float32(1.5) * 1.5) / 1.5) <= 1.5) {
                                                                v65 := GenS3{F0: (v47 / (((^(v47))) | 1)), F1: uint8(v51)}
                                                                var v66 uint = 255
                                                                delete(v3, 128)
                                                                v67 := GenS6{F0: 541, F1: (v47 << ((1) & 63)), F2: (fi26 - fi22)}
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v57 + v57)
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                var v68 string = ("foo bar" + ("" + v57))
                                                                if (false && ("a" == v68)) {
                                                                    if false {
                                                                        v4 = 0
                                                                    }
                                                                    var v69 uint32 = 0
                                                                    fmt.Println(v54, v68, v51)
                                                                    _ = v69
                                                                }
                                                                v57 = (v68 + "hello")
                                                                _ = v68
                                                            }
                                                        case "\n\t":
                                                            var v70 int32 = v2
                                                            var v71 uint16 = v51
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    }
                                                    var v72 int = (fi22 ^ fi26)
                                                    switch "foo bar" {
                                                    case "a":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v57 = v57
                                                    default:
                                                        v73 := GenS3{F0: v47, F1: 0}
                                                        if true {
                                                            var v74 int = 32767
                                                            _ = v74
                                                        } else {
                                                            v54 = v54
                                                            v75 := GenS4{F0: 883}
                                                            if (((1e10 + (0.1 / 1.0)) * 0.1) != (1.5 / 1.5)) {
                                                                v4 -= -129
                                                                v72 = fi22
                                                                if (true || (true || false)) {
                                                                    v57 = "\n\t"
                                                                }
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            _ = v75
                                                        }
                                                        var v76 int8 = 92
                                                        fmt.Println(v72, v4, v4)
                                                        var v77 uint = 1
                                                        _ = v73
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    v4 = (v4 + v4)
                                                    _ = v62
                                                    _ = v72
                                                }
                                            }
                                            fmt.Println(v47, v54)
                                            _ = v57
                                        }
                                        switch (fi26 >> ((601) & 63)) {
                                        case 9223372036854775807:
                                            v78 := GenS7{F0: 992}
                                            delete(v5, -9223372036854775808)
                                            if false {
                                                if false {
                                                    v79 := GenS5{F0: (100 >= 0), F1: false}
                                                    _ = v79
                                                }
                                            }
                                            _ = v78
                                        case -1:
                                            var v80 uint8 = 128
                                            v81 := GenS5{F0: (-9223372036854775808 <= -1), F1: (uint8(v51) <= (^(v80)))}
                                            _ = v80
                                            _ = v81
                                        case -16:
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                switch "" {
                                                case "\n\t":
                                                    v82 := GenS8{F0: "\n\t", F1: fi22, F2: v54}
                                                    _ = v82
                                                }
                                            case "a":
                                                fmt.Println(fi26, v51)
                                            case "":
                                                v83 := GenS7{F0: genFn0(v47)}
                                                _ = v83
                                            default:
                                                var v84 uint16 = v51
                                                _ = v84
                                            }
                                        }
                                        v85 := GenS9{F0: "hello"}
                                        _ = v85
                                    } else {
                                        v2 = v2
                                        var v86 int8 = 1
                                        _ = v86
                                    }
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            var v87 int32 = (v2 >> ((2) & 63))
                                            _ = v87
                                        case "":
                                            switch (fi22 / ((int(v47)) | 1)) {
                                            case 42:
                                                v47 = uint32(v51)
                                            }
                                        case "foo bar":
                                            v88 := GenS10{F0: int64(fi26), F1: 1}
                                            _ = v88
                                        default:
                                            v1 = append(v1, "\n\t")
                                        }
                                    }
                                    _ = len(v1)
                                }
                            }
                            _ = v54
                        }
                        _ = v47
                        _ = v51
                        _ = v52
                    }
                case "\n\t":
                    var v89 int16 = v4
                    _ = v89
                }
                var v90 uint8 = 37
                var v91 int32 = v2
                _ = fi26
                _ = v35
                _ = v90
                _ = v91
            }
        }
        _ = fi22
    }
    switch 0 {
    case 32767:
        v2 += (v2 / ((v2) | 1))
    case 127:
        var v92 uint32 = 42
        _ = v92
    case 474:
        v2 = (v2 / ((v2) | 1))
    default:
        v93 := GenS11{F0: -2147483648}
        _ = v93
    }
    fmt.Println(v2)
    var v94 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v94:", v94)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v94
}
