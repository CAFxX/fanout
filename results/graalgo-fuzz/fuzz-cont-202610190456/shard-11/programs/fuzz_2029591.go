package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16; F1 int; F2 int8 }
type GenS8 struct { F0 uint64; F1 int32; F2 int32 }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint32; F1 int64 }
type GenS15 struct { F0 int8; F1 uint8 }
func genFn0(p0 int8, p1 int32) bool {
    var v4 uint = 42
    if true {
        var v5 uint = (v4 << ((4294967295) & 63))
        _ = v5
    }
    _ = p0
    _ = p1
    _ = v4
    return true
}
func genFn1(p0 uint16, p1 int32, p2 int16) string {
    if false {
        if true {
            if (true || false) {
                var v6 uint8 = 1
                p1 = 2147483647
                v6 += 0
                _ = v6
            }
        } else {
            switch "" {
            case "foo bar":
                v7 := GenS1{F0: p0, F1: 4294967295}
                _ = v7
            case "hello":
                v8 := []int{}
                _ = v8
            case "a":
                if false {
                    p0 -= 42
                    switch int(p1) {
                    case -2147483648:
                        if false {
                            for fi9 := 0; fi9 < 15; fi9++ {
                                v10 := GenS2{F0: 857}
                                p1 += 103
                                _ = fi9
                                _ = v10
                            }
                            fmt.Println(p0, p2)
                            for fi11 := 0; fi11 < 12; fi11++ {
                                p1 -= (p1 + (p1 * 42))
                                v12 := []int{int(p2), -1, (fi11 / (((^(fi11))) | 1))}
                                _ = fi11
                                _ = v12
                            }
                        }
                    case 128:
                        if true {
                            switch "\n\t" {
                            case "\n\t":
                                v13 := []int{}
                                _ = v13
                            default:
                                v14 := GenS3{F0: p2, F1: (("" + "") <= ("\n\t" + ("foo bar" + ((("hello" + "\n\t") + ("hello" + "a")) + ("\n\t" + "\n\t")))))}
                                _ = v14
                            }
                            v15 := GenS2{F0: 65535}
                            _ = v15
                        }
                    case 798:
                        if (!false) {
                            if (p2 >= (p2 / (((p2 / ((-32768) | 1))) | 1))) {
                                var v16 int64 = 1
                                var v17 uint16 = p0
                                v18 := GenS4{F0: (-7 < 0)}
                                _ = v16
                                _ = v17
                                _ = v18
                            }
                            if (false || (1 <= 0)) {
                                v19 := GenS4{F0: (!((^(p0)) == (p0 % ((p0) | 1))))}
                                _ = v19
                            }
                            p1 -= (-(p1))
                        } else {
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 uint8 = 127
                                _ = fi20
                                _ = v21
                            }
                            fmt.Println(p2, p1)
                            v22 := []int{128}
                            _ = v22
                        }
                    }
                    var v23 int32 = p1
                    _ = v23
                } else {
                    v24 := GenS3{F0: (p2 + p2), F1: false}
                    var v25 uint = 2
                    _ = v24
                    _ = v25
                }
            }
            v26 := GenS5{F0: 2, F1: -9223372036854775808}
            var v27 int16 = (p2 * 32767)
            _ = v26
            _ = v27
        }
    } else {
        p1 = p1
        switch (("a" + "") + "a") {
        case "a":
            v28 := []int{2147483647}
            _ = v28
        case "hello":
            v29 := GenS6{F0: 911}
            _ = v29
        case "\n\t":
            if (1 < -221) {
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := []int{}
                    _ = fi30
                    _ = v31
                }
            }
        }
        v32 := []int{1}
        _ = v32
    }
    switch "foo bar" {
    case "\n\t":
        var v33 int64 = -290
        _ = v33
    case "foo bar":
        v34 := []int{4294967295}
        _ = v34
    default:
        for fi35 := 0; fi35 < 7; fi35++ {
            var v36 bool = true
            v37 := GenS7{F0: p2, F1: (fi35 * int(p1)), F2: 42}
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    switch "foo bar" {
    case "":
        for fi38 := 0; fi38 < 3; fi38++ {
            if (fi38 <= (fi38 | -1)) {
                for fi39 := 0; fi39 < 5; fi39++ {
                    if true {
                        fmt.Println(p1)
                    } else {
                        if (false && true) {
                            v40 := GenS6{F0: (p2 * (p2 | (p2 << ((18446744073709551615) & 63))))}
                            switch ("foo bar" + ((("a" + ("hello" + "\n\t")) + (("a" + "a") + "hello")) + "\n\t")) {
                            case "a":
                                var v41 float64 = float64(1e10)
                                _ = v41
                            }
                            var v42 string = ("foo bar" + "foo bar")
                            _ = v40
                            _ = v42
                        } else {
                            var v43 float64 = (((3.14159 / (float64(((float32(0.1) + ((3.14159 * (1e10 - 0.1)) - 0.0)) / ((0.1 - (1e10 / -2.25)) * 1e10))) * ((1e10 - (1.5 + -2.25)) - (-2.25 - 3.14159)))) - 0.1) / ((0.0 / 0.1) - 0.1))
                            if (1e10 == 0.1) {
                                v43 = v43
                                if false {
                                    v44 := GenS8{F0: uint64(p0), F1: (-(p1)), F2: 2147483647}
                                    if (!(true && (false && (!true)))) {
                                        v45 := GenS9{F0: -1, F1: "hello"}
                                        var v46 int = (fi38 >> ((448) & 63))
                                        switch "\n\t" {
                                        case "hello":
                                            if (17 != 1) {
                                                var v47 int8 = 69
                                                v48 := GenS10{F0: 255}
                                                fmt.Println(p1, fi38, p1)
                                                v49 := []int{fi38, fi38, (v46 & (fi39 - v46)), 128}
                                                _ = v47
                                                _ = v48
                                                _ = v49
                                            }
                                        case "\n\t":
                                            var v50 int8 = 100
                                            _ = v50
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    switch 1 {
                                    case 281:
                                        var v51 uint32 = 4294967295
                                        _ = v51
                                    case -899:
                                        fmt.Println(p2, p0)
                                    case 128:
                                        if false {
                                            var v52 int8 = -22
                                            v53 := GenS11{F0: p1}
                                            v54 := []int{(fi38 / (((fi38 % (((fi38 >> ((856) & 63))) | 1))) | 1))}
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        } else {
                                            switch ("\n\t" + ("a" + "a")) {
                                            case "\n\t":
                                                var v55 int8 = 1
                                                _ = v55
                                            case "foo bar":
                                                if (-1 == -54) {
                                                    var v56 uint64 = 320
                                                    if false {
                                                        var v57 float32 = -2.25
                                                        switch (("foo bar" + "") + "foo bar") {
                                                        case "hello":
                                                            var v58 uint32 = 2147483648
                                                            _ = v58
                                                        case "":
                                                            fmt.Println(v56, fi39)
                                                        default:
                                                            switch (fi38 & fi38) {
                                                            case -930:
                                                                v59 := GenS12{F0: 255}
                                                                _ = v59
                                                            default:
                                                                fmt.Println(v57, v56, fi38)
                                                            }
                                                        }
                                                        v43 -= (1e10 / float64(v57))
                                                        _ = v57
                                                    }
                                                    _ = v56
                                                } else {
                                                    var v60 uint = 776
                                                    v61 := GenS8{F0: 852, F1: p1, F2: 42}
                                                    switch ("hello" + ("foo bar" + (("hello" + "a") + "hello"))) {
                                                    case "foo bar":
                                                        v62 := GenS10{F0: 945}
                                                        _ = v62
                                                    case "\n\t":
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v63 := GenS13{F0: p0}
                                                            _ = v63
                                                        case "a":
                                                            v64 := []int{}
                                                            _ = v64
                                                        default:
                                                            v43 = 0.0
                                                        }
                                                    case "a":
                                                        var v65 uint32 = 42
                                                        _ = v65
                                                    default:
                                                        switch "foo bar" {
                                                        case "a":
                                                            v66 := GenS13{F0: p0}
                                                            _ = v66
                                                        case "\n\t":
                                                            var v67 uint = (v60 + 18446744073709551615)
                                                            _ = v67
                                                        case "":
                                                            var v68 bool = (415 > -9223372036854775808)
                                                            _ = v68
                                                        }
                                                    }
                                                    _ = v60
                                                    _ = v61
                                                }
                                            case "hello":
                                                if false {
                                                    v69 := []int{32767}
                                                    v43 -= 1.5
                                                    switch ("foo bar" + "a") {
                                                    case "\n\t":
                                                        var v70 string = ""
                                                        var v71 uint16 = p0
                                                        _ = v70
                                                        _ = v71
                                                    case "a":
                                                        var v72 string = "\n\t"
                                                        _ = v72
                                                    case "hello":
                                                        if len(v69) >= 2 {
                                                            v69 = v69[1:len(v69)]
                                                        }
                                                    default:
                                                        fmt.Println(fi39)
                                                    }
                                                    v73 := GenS11{F0: (p1 - p1)}
                                                    _ = v69
                                                    _ = v73
                                                }
                                            }
                                            if (!false) {
                                                v74 := []int{fi38, fi38}
                                                v43 = float64(0.1)
                                                var v75 int16 = (p2 << ((103) & 63))
                                                _ = v74
                                                _ = v75
                                            } else {
                                                if (1 > uint32(p0)) {
                                                    switch ("\n\t" + "a") {
                                                    case "":
                                                        var v76 int = fi39
                                                        _ = v76
                                                    default:
                                                        var v77 bool = false
                                                        _ = v77
                                                    }
                                                }
                                                var v78 int64 = -743
                                                _ = v78
                                            }
                                            fmt.Println(fi39, p0)
                                        }
                                    default:
                                        v79 := []int{(fi39 * int(p2)), fi38, fi38, fi39}
                                        _ = v79
                                    }
                                    _ = v44
                                } else {
                                    p1 -= (p1 / ((504) | 1))
                                    var v80 int8 = -113
                                    var v81 string = "\n\t"
                                    _ = v80
                                    _ = v81
                                }
                            }
                            var v82 uint8 = 55
                            _ = v43
                            _ = v82
                        }
                        v83 := []int{fi39, -613, fi38, (fi39 - fi38)}
                        p2 = p2
                        _ = v83
                    }
                    var v84 int64 = -299
                    var v85 bool = (!(42 > 0))
                    v86 := []int{fi38, (fi39 / ((fi38) | 1)), fi38, fi39}
                    _ = fi39
                    _ = v84
                    _ = v85
                    _ = v86
                }
            }
            _ = fi38
        }
    case "a":
        v87 := GenS14{F0: 56, F1: 9223372036854775807}
        _ = v87
    case "foo bar":
        p1 = 2147483647
    }
    v88 := []int{128}
    _ = p0
    _ = p1
    _ = p2
    _ = v88
    return "a"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 688
    v3 := []string{"hello", ((("a" + "hello") + "\n\t") + ("" + "\n\t"))}
    for fi89 := 0; fi89 < 16; fi89++ {
        v1["foo bar"] = 42
        for fi90 := 0; fi90 < 7; fi90++ {
            var v91 uint16 = 0
            v2 = (v2 ^ (^(v2)))
            _ = fi90
            _ = v91
        }
        _ = fi89
    }
    v2 = (v2 ^ 42)
    var v92 string = "foo bar"
    fmt.Println(v92)
    switch 9223372036854775807 {
    case -935:
        switch 4294967295 {
        case 1:
            v2 += v2
        case 9223372036854775807:
            if len(v3) > 0 {
                v3[0] = (v92 + "\n\t")
            }
        default:
            if (127 < 49) {
                var v93 uint32 = 540
                v3 = append(v3, "hello")
                if true {
                    v1[""] = v93
                } else {
                    var v94 int8 = 100
                    v92 = v92
                    delete(v1, v92)
                    _ = v94
                }
                _ = v93
            }
        }
    default:
        v92 = v92
    }
    var v95 int64 = 168
    v2 += v2
    var v96 string = v92
    if false {
        v96 = v96
        fmt.Println(v92, v96)
        v95 = 9223372036854775807
    }
    if (false && false) {
        var v97 int = 1
        v98 := GenS15{F0: -128, F1: 42}
        _ = v97
        _ = v98
    }
    var v99 float32 = (0.1 * 3.14159)
    v1[((v92 + v96) + "")] = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v92:", v92)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v92
    _ = v95
    _ = v96
    _ = v99
}
