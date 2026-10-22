package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 uint32; F1 int }
type GenS10 struct { F0 uint; F1 int }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int64; F1 uint16 }
type GenS16 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int64 }
type GenS19 struct { F0 uint64; F1 bool; F2 uint }
func genFn0(p0 uint32, p1 uint32) int64 {
    if false {
        switch 42 {
        case 4294967295:
            v4 := GenS1{F0: 185, F1: 584, F2: true}
            _ = v4
        case 0:
            var v5 float32 = 0.0
            _ = v5
        case 128:
            var v6 float32 = 3.14159
            _ = v6
        }
        var v7 float32 = 1.5
        _ = v7
    } else {
        var v8 int64 = 42
        fmt.Println(p0, v8)
        fmt.Println(p1, p0)
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if ((!false) || false) {
            v10 := GenS2{F0: 42}
            v11 := GenS3{F0: 104, F1: p0}
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        p1 += (p1 * p1)
        v13 := []int{}
        fmt.Println(p0, p0, p1)
        _ = fi12
        _ = v13
    }
    if true {
        if (true || true) {
            v14 := []int{1}
            v15 := GenS1{F0: 4294967295, F1: 0, F2: false}
            p1 += p0
            _ = v14
            _ = v15
        }
        p0 -= (p0 & (p0 >> ((4294967295) & 63)))
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int = 1
            _ = fi16
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        _ = v18
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int = 32768
    v2 := make(map[string]uint32)
    var v3 int = (-(v1))
    v19 := GenS4{F0: -388, F1: genFn0(1, 0)}
    var v20 int32 = 42
    if (false || true) {
        if (!true) {
            var v21 int8 = 1
            fmt.Println(v3, v1, v20)
            var v22 float32 = (0.0 + (0.0 * 0.0))
            _ = v21
            _ = v22
        }
        v23 := []int{42, 127, (v3 >> ((4294967295) & 63)), (^(v3))}
        _ = v23
    } else {
        v24 := []int{(v3 / ((v3) | 1)), v1, (v1 ^ 848), (v1 + v1)}
        v24 = append(v24, (-(v1)))
        _ = v24
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if true {
            if false {
                v2[(("foo bar" + ("" + ("foo bar" + "hello"))) + ((("" + "hello") + "\n\t") + (("hello" + "\n\t") + ("\n\t" + ("foo bar" + "foo bar")))))] = 2147483648
            } else {
                if ((128 >= 62) || false) {
                    switch (fi25 | (v1 << ((1) & 63))) {
                    case -2147483648:
                        v26 := GenS3{F0: 1, F1: 523}
                        _ = v26
                    default:
                        if true {
                            var v27 uint8 = 42
                            _ = v27
                        }
                    }
                    v28 := []int{fi25, (v1 / ((v3) | 1)), 42}
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        switch fi25 {
                        case -331:
                            var v30 float64 = 1e10
                            v30 += (v30 * 3.14159)
                            if true {
                                if false {
                                    v1 += fi25
                                    switch "a" {
                                    case "foo bar":
                                        v31 := []int{fi29}
                                        fmt.Println(fi29)
                                        if (!true) {
                                            var v32 uint64 = 405
                                            _ = len(v31)
                                            _ = v32
                                        }
                                        _ = v31
                                    case "a":
                                        v20 -= v20
                                    case "\n\t":
                                        v20 += v20
                                    default:
                                        var v33 bool = true
                                        _ = v33
                                    }
                                } else {
                                    v34 := GenS5{F0: genFn0(2147483647, 65535), F1: "foo bar"}
                                    var v35 int64 = -1
                                    if true {
                                        switch int(v35) {
                                        case 1:
                                            switch v1 {
                                            case 0:
                                                fmt.Println(v35, v20)
                                            }
                                        case -373:
                                            v3 += -2147483648
                                        case -512:
                                            if ((float64((0.0 + -2.25)) + float64((0.0 / float32(v30)))) > 0.1) {
                                                fmt.Println(fi25)
                                                var v36 uint8 = 128
                                                _ = v36
                                            }
                                        default:
                                            if false {
                                                var v37 uint32 = 786
                                                v38 := GenS4{F0: (-(fi25)), F1: (v35 + -9223372036854775808)}
                                                _ = v37
                                                _ = v38
                                            } else {
                                                v39 := []int{}
                                                var v40 uint16 = 40
                                                var v41 uint32 = 4294967295
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            }
                                        }
                                        v30 = v30
                                    } else {
                                        if (false || false) {
                                            var v42 int = fi29
                                            if ((!true) && false) {
                                                delete(v2, "foo bar")
                                                v43 := GenS3{F0: 100, F1: 4294967295}
                                                fmt.Println(v1, v3)
                                                _ = v43
                                            } else {
                                                v44 := GenS6{F0: v20}
                                                fmt.Println(v1, v3)
                                                v30 -= (v30 * (((3.14159 - float64(float32((1e10 - (v30 - v30))))) / 1e10) - (v30 * v30)))
                                                _ = v44
                                            }
                                            _ = v42
                                        } else {
                                            if false {
                                                v45 := GenS7{F0: v35}
                                                v30 -= (v30 + (1.5 - 1.5))
                                                v46 := GenS1{F0: 1, F1: 626, F2: false}
                                                delete(v2, "")
                                                _ = v45
                                                _ = v46
                                            }
                                            var v47 int = (v1 % (((v3 | 762)) | 1))
                                            v20 += (^(v20))
                                            switch "hello" {
                                            case "":
                                                v48 := GenS8{F0: 255, F1: (v47 << ((42) & 63))}
                                                _ = v48
                                            case "\n\t":
                                                v49 := GenS5{F0: (v35 - 42), F1: "a"}
                                                _ = v49
                                            }
                                            v50 := GenS9{F0: 728, F1: 42}
                                            v3 = fi29
                                            v47 = v1
                                            v2[""] = 42
                                            switch fi29 {
                                            case -1:
                                                if (((1 <= 42) && (false && true)) && false) {
                                                    var v51 int64 = genFn0(42, 2)
                                                    if true {
                                                        fmt.Println(v1, fi29)
                                                        v1 -= (fi25 ^ fi25)
                                                        var v52 string = "hello"
                                                        _ = v52
                                                    }
                                                    _ = v51
                                                }
                                            case 2147483647:
                                                var v53 uint8 = 255
                                                _ = v53
                                            case 985:
                                                v54 := GenS10{F0: 0, F1: 711}
                                                _ = v54
                                            }
                                            fmt.Println(v35, v1, v1)
                                            fmt.Println(v20)
                                            _ = v47
                                            _ = v50
                                        }
                                        switch -259 {
                                        case -2147483648:
                                            if false {
                                                _ = len(v2)
                                                v55 := GenS7{F0: v35}
                                                var v56 uint = 860
                                                _ = v55
                                                _ = v56
                                            } else {
                                                var v57 int64 = v35
                                                _ = v57
                                            }
                                        default:
                                            if true {
                                                _ = len(v2)
                                                if false {
                                                    var v58 int32 = 42
                                                    v59 := GenS8{F0: 42, F1: (v1 ^ (v3 / (((v1 << ((651) & 63))) | 1)))}
                                                    fmt.Println(v58)
                                                    v60 := []int{(fi29 & v3)}
                                                    var v61 int64 = v35
                                                    v60 = append(v60, (^(v3)))
                                                    var v62 uint32 = 690
                                                    v30 = v30
                                                    _ = v58
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    v63 := v2[("a" + "")]
                                                    if (294 <= 4294967295) {
                                                        v64 := []int{v3}
                                                        if true {
                                                            fmt.Println(v63, v35, v35)
                                                            if false {
                                                                v3 = (fi29 | v3)
                                                                var v65 uint32 = 770
                                                                delete(v2, ("foo bar" + ""))
                                                                var v66 bool = (!((false && false) || false))
                                                                _ = v65
                                                                _ = v66
                                                            } else {
                                                                var v67 uint16 = 165
                                                                fmt.Println(v20, v67)
                                                                if (true || false) {
                                                                    switch ("hello" + "") {
                                                                    case "hello":
                                                                        fmt.Println(v3)
                                                                    default:
                                                                        var v68 int32 = (-(v20))
                                                                        _ = v68
                                                                    }
                                                                    var v69 uint8 = 42
                                                                    v35 = (v35 / (((v35 * (v35 % ((0) | 1)))) | 1))
                                                                    _ = v69
                                                                }
                                                                _ = v67
                                                            }
                                                        } else {
                                                            fmt.Println(v3, v35, fi25)
                                                            v70 := GenS4{F0: 32768, F1: v35}
                                                            if (!(-17 <= -1)) {
                                                                switch 0 {
                                                                case 9223372036854775807:
                                                                    if (412 >= -32768) {
                                                                        var v71 string = ("\n\t" + "hello")
                                                                        _ = v71
                                                                    }
                                                                case -1:
                                                                    fmt.Println(v1, fi29, v20)
                                                                }
                                                                v72 := GenS11{F0: 65535}
                                                                v73 := GenS11{F0: 235}
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            _ = v70
                                                        }
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(v63, v35)
                                                        case "a":
                                                            switch -1 {
                                                            case 9223372036854775807:
                                                                _ = len(v2)
                                                            case -903:
                                                                v35 -= (v35 * (v35 ^ (v35 / (((v35 >> ((2) & 63))) | 1))))
                                                            case -311:
                                                                if len(v64) > 0 {
                                                                    v64[0] = fi25
                                                                }
                                                            }
                                                            var v74 bool = true
                                                            _ = v74
                                                        default:
                                                            var v75 int8 = 1
                                                            v35 = -9223372036854775808
                                                            _ = v75
                                                        }
                                                        if (false || (-43 == -1)) {
                                                            v3 = v3
                                                            switch (fi25 - 32768) {
                                                            case -1:
                                                                if (!true) {
                                                                    v76 := GenS12{F0: -117}
                                                                    if false {
                                                                        switch "" {
                                                                        case "hello":
                                                                            fmt.Println(v35, v35, v3)
                                                                        case "a":
                                                                            v77 := GenS13{F0: (-(fi25))}
                                                                            _ = v77
                                                                        case "foo bar":
                                                                            var v78 int32 = (^(v20))
                                                                            _ = v78
                                                                        default:
                                                                            v20 = (v20 & (v20 + v20))
                                                                        }
                                                                        delete(v2, ((("" + "a") + "hello") + "hello"))
                                                                        if true {
                                                                            switch "" {
                                                                            case "a":
                                                                                if len(v64) > 0 {
                                                                                    v64[0] = (-(fi29))
                                                                                }
                                                                            }
                                                                        }
                                                                        v79 := GenS11{F0: 65535}
                                                                        v30 = v30
                                                                        _ = v79
                                                                    }
                                                                    _ = v76
                                                                } else {
                                                                    v80 := GenS14{F0: -32}
                                                                    v63 = 34
                                                                    v81 := GenS15{F0: -991, F1: 42}
                                                                    _ = v80
                                                                    _ = v81
                                                                }
                                                            case 32767:
                                                                v30 += (float64(1e10) / (v30 + 0.1))
                                                            case 1:
                                                                if false {
                                                                    fmt.Println(v3, v63, fi25)
                                                                }
                                                            default:
                                                                fmt.Println(v35)
                                                            }
                                                            if (true || false) {
                                                                if true {
                                                                    var v82 int = 0
                                                                    _ = v82
                                                                }
                                                                if (v20 > v20) {
                                                                    v2["\n\t"] = (v63 ^ 50)
                                                                }
                                                                fmt.Println(v30, v30, v1)
                                                            } else {
                                                                v63 += v63
                                                            }
                                                            var v83 int32 = v20
                                                            switch (v1 / ((fi29) | 1)) {
                                                            case -47:
                                                                v64 = append(v64, (-(fi25)))
                                                                if ((-1 <= 42) || (-1 < 1)) {
                                                                    if false {
                                                                        v3 -= v3
                                                                        var v84 int16 = 0
                                                                        _ = v84
                                                                    }
                                                                }
                                                            case -519:
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    if (687 != 42) {
                                                                        v85 := GenS8{F0: 2, F1: -1}
                                                                        v2["hello"] = (v63 * v63)
                                                                        _ = v85
                                                                    } else {
                                                                        var v86 float64 = v30
                                                                        _ = v86
                                                                    }
                                                                }
                                                            }
                                                            _ = v83
                                                        }
                                                        _ = v64
                                                    } else {
                                                        v20 += 116
                                                        var v87 uint16 = 1
                                                        _ = v87
                                                    }
                                                    switch fi29 {
                                                    case -321:
                                                        if false {
                                                            v88 := GenS1{F0: 2, F1: 42, F2: (696 != 461)}
                                                            _ = v88
                                                        }
                                                    default:
                                                        fmt.Println(v30)
                                                    }
                                                    v63 += v63
                                                    _ = v63
                                                }
                                                if true {
                                                    v20 = v20
                                                    var v89 int16 = 42
                                                    v90 := GenS3{F0: -1, F1: 42}
                                                    v20 = (v20 >> ((1) & 63))
                                                    _ = v89
                                                    _ = v90
                                                } else {
                                                    switch -1 {
                                                    case 4294967295:
                                                        switch (("a" + ("hello" + ((("" + "foo bar") + ("hello" + "\n\t")) + ((("" + "\n\t") + "") + "foo bar")))) + "") {
                                                        case "foo bar":
                                                            v35 = (v35 - int64(v20))
                                                            v30 += ((v30 * (float64(-2.25) + (1.5 + v30))) + v30)
                                                        case "a":
                                                            v35 = (v35 - genFn0(482, 2147483647))
                                                        case "":
                                                            fmt.Println(v3, v3)
                                                        }
                                                    case 812:
                                                        var v91 bool = false
                                                        _ = v91
                                                    case 127:
                                                        v92 := GenS2{F0: -32768}
                                                        _ = v92
                                                    default:
                                                        if false {
                                                            v93 := GenS16{F0: 0, F1: v35, F2: 1}
                                                            _ = v93
                                                        }
                                                    }
                                                    v30 = ((v30 - float64((3.14159 + (3.14159 + 1e10)))) - (v30 / 1.5))
                                                }
                                            }
                                        }
                                        v2["foo bar"] = 65535
                                        var v94 int64 = (v35 - genFn0(323, 1))
                                        if (!false) {
                                            switch "\n\t" {
                                            case "hello":
                                                v95 := []int{fi25, v3}
                                                _ = v95
                                            }
                                            if (27 != 2) {
                                                fmt.Println(fi25, v35)
                                                v94 -= v35
                                            } else {
                                                v2["a"] = 87
                                                switch -9223372036854775808 {
                                                case 109:
                                                    switch "" {
                                                    case "\n\t":
                                                        v30 = (v30 / v30)
                                                    case "hello":
                                                        var v96 int64 = -1
                                                        _ = v96
                                                    }
                                                default:
                                                    var v97 string = "hello"
                                                    _ = v97
                                                }
                                            }
                                        }
                                        _ = v94
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                _ = len(v2)
                                v98 := GenS17{F0: "\n\t"}
                                _ = v98
                            }
                            _ = v30
                        }
                        _ = fi29
                    }
                    fmt.Println(v1, v20)
                }
                v1 = (v3 & (fi25 << ((255) & 63)))
                fmt.Println(v1, v1)
            }
            fmt.Println(v3, fi25, v20)
            v20 = v20
        } else {
            if false {
                var v99 uint16 = 65535
                v20 = v20
                var v100 uint64 = 4294967295
                _ = v100
                _ = v99
            } else {
                v101 := []int{}
                v102 := GenS18{F0: -9223372036854775808}
                _ = v101
                _ = v102
            }
            v1 -= (v1 - v3)
            v3 += fi25
        }
        if (!(float64(0.0) < (((float64(1.5) - 0.0) - 0.0) / -2.25))) {
            var v103 uint8 = 128
            var v104 int16 = 1000
            _ = v103
            _ = v104
        }
        delete(v2, ("\n\t" + "a"))
        _ = fi25
    }
    var v105 bool = ("\n\t" < "foo bar")
    v20 = int32(v3)
    var v106 uint8 = 42
    var v107 uint16 = 0
    if v105 {
        for fi108 := 0; fi108 < 14; fi108++ {
            for fi109 := 0; fi109 < 18; fi109++ {
                switch (fi109 * fi109) {
                case 0:
                    var v110 int16 = 1000
                    _ = v110
                }
                var v111 string = "foo bar"
                v112 := GenS19{F0: 4294967295, F1: false, F2: 65535}
                v113 := v2[v111]
                _ = fi109
                _ = v111
                _ = v112
                _ = v113
            }
            fmt.Println(v3)
            _ = fi108
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v105
    _ = v106
    _ = v107
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
