package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int64; F1 int64; F2 uint16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 uint16; F1 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint16; F1 int64 }
func genFn0(p0 bool, p1 uint64) int64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{int(p1)}
        fmt.Println(p0, p1)
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 32
    var v10 int32 = 42
    var v11 int64 = -847
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return v11
}
func genFn1(p0 uint16, p1 int, p2 int8) int {
    for fi12 := 0; fi12 < 13; fi12++ {
        p0 = p0
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 uint16 = (^(p0))
            p2 += (p2 * p2)
            _ = fi13
            _ = v14
        }
        fmt.Println(p0, p2)
        _ = fi12
    }
    if (0 <= 2) {
        p0 += (p0 & (^(p0)))
        p0 = (p0 ^ 0)
        switch "" {
        case "\n\t":
            switch p1 {
            case 9223372036854775807:
                for fi15 := 0; fi15 < 9; fi15++ {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v17 := GenS1{F0: (p2 * p2)}
                        var v18 int64 = genFn0((0 >= 42), 42)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    switch ("\n\t" + (("a" + "foo bar") + "hello")) {
                    case "foo bar":
                        p2 = (p2 + (p2 / (((p2 * p2)) | 1)))
                    }
                    _ = fi15
                }
            }
        case "foo bar":
            v19 := GenS1{F0: 65}
            _ = v19
        case "":
            if false {
                v20 := GenS1{F0: p2}
                _ = v20
            } else {
                var v21 uint16 = (p0 * p0)
                if (true || true) {
                    var v22 uint16 = v21
                    p2 = p2
                    _ = v22
                }
                if true {
                    p2 -= (p2 ^ p2)
                    p1 -= (p1 & (p1 + -9223372036854775808))
                }
                _ = v21
            }
        default:
            switch ("foo bar" + (("foo bar" + "foo bar") + "hello")) {
            case "a":
                fmt.Println(p1, p2)
            case "\n\t":
                p0 = 0
            default:
                v23 := GenS2{F0: 2}
                _ = v23
            }
        }
    } else {
        if (true && true) {
            if (!false) {
                for fi24 := 0; fi24 < 17; fi24++ {
                    if ((!(699 == 2147483648)) && (((128 != 0) && (true || (1000 == -636))) && true)) {
                        fmt.Println(p2, p1, p2)
                        var v25 bool = (772 <= 907)
                        _ = v25
                    } else {
                        var v26 uint16 = 2
                        v26 = p0
                        v27 := []int{p1}
                        _ = v26
                        _ = v27
                    }
                    switch (("foo bar" + ("foo bar" + ("hello" + "hello"))) + "") {
                    case "foo bar":
                        fmt.Println(p2, p1)
                    case "":
                        v28 := GenS3{F0: -604, F1: 65535, F2: ("hello" + "a")}
                        _ = v28
                    default:
                        if ((p2 & 42) == p2) {
                            var v29 string = ("" + ("hello" + ""))
                            switch p1 {
                            case -1:
                                var v30 int = (-(p1))
                                _ = v30
                            default:
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    v32 := GenS4{F0: 29}
                                    var v33 uint64 = 18446744073709551615
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                            }
                            switch fi24 {
                            case 42:
                                v29 = ((("foo bar" + v29) + (("\n\t" + v29) + ((v29 + "") + "a"))) + "hello")
                            }
                            _ = v29
                        } else {
                            var v34 string = "hello"
                            _ = v34
                        }
                    }
                    _ = fi24
                }
                if true {
                    p0 = p0
                    p0 += (p0 ^ p0)
                }
                var v35 uint16 = 353
                _ = v35
            }
            p2 = p2
        }
        p0 = 0
    }
    if (false && true) {
        p2 = (p2 ^ p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 ^ p1)
}

func main() {
    var v1 uint16 = 91
    v2 := []uint32{4294967295, 4294967295, 2147483648, 65535}
    v3 := make(map[int]bool)
    var v4 uint64 = 26
    var v5 string = "foo bar"
    v6 := make(map[uint32]string)
    var v36 int = -515
    switch (v36 ^ genFn1(v1, v36, 29)) {
    case 1:
        v37 := GenS1{F0: 0}
        _ = v37
    case 273:
        fmt.Println(v5)
    default:
        if ((!false) && false) {
            v38 := GenS5{F0: 0}
            _ = v38
        } else {
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 int64 = -9223372036854775808
                for fi41 := 0; fi41 < 14; fi41++ {
                    var v42 int16 = int16(v40)
                    _ = len(v2)
                    if true {
                        if true {
                            v40 += genFn0((!false), v4)
                            v36 = int(v4)
                        }
                        v43 := GenS5{F0: -128}
                        v1 = v1
                        v44 := GenS6{F0: 65535, F1: ((-(v42)) == v42)}
                        switch "a" {
                        case "hello":
                            _ = len(v2)
                        case "foo bar":
                            v1 += (v1 >> ((58) & 63))
                        case "":
                            var v45 string = ((v5 + ((v5 + v5) + (("hello" + (v5 + "")) + v5))) + v5)
                            fmt.Println(v1, fi41)
                            v40 += 1
                            if (v1 > 2) {
                                var v46 float64 = 0.0
                                v1 += 596
                                fmt.Println(v5, v36)
                                _ = v46
                            } else {
                                v1 += v1
                                if (42 >= 0) {
                                    v42 -= 1000
                                    v2 = append(v2, 2)
                                    v47 := GenS4{F0: v4}
                                    var v48 uint32 = 42
                                    _ = v47
                                    _ = v48
                                } else {
                                    switch ("" + (((((v5 + (v5 + v5)) + (("hello" + (v45 + v5)) + (v5 + v45))) + (v45 + (v45 + v45))) + v5) + "a")) {
                                    case "foo bar":
                                        v49 := GenS5{F0: -48}
                                        _ = v49
                                    case "":
                                        var v50 float32 = (float32(1.5) + (1.5 * 0.1))
                                        _ = v50
                                    default:
                                        if (128 > 255) {
                                            var v51 int64 = v40
                                            if (820 == 929) {
                                                var v52 int16 = 32767
                                                _ = v52
                                            }
                                            _ = v51
                                        } else {
                                            v36 = fi39
                                            v53 := GenS7{F0: v40, F1: v40, F2: 255}
                                            v1 -= v1
                                            if false {
                                                v54 := GenS7{F0: v40, F1: (v40 / (((^(v40))) | 1)), F2: 42}
                                                _ = v54
                                            } else {
                                                var v55 int = -9223372036854775808
                                                v3[127] = (2 >= uint32(v1))
                                                switch -9223372036854775808 {
                                                case 894:
                                                    var v56 float32 = 0.1
                                                    _ = v56
                                                }
                                                _ = v55
                                            }
                                            v57 := GenS6{F0: 2, F1: ((v42 / (((v42 ^ v42)) | 1)) > -248)}
                                            _ = v53
                                            _ = v57
                                        }
                                    }
                                }
                                v42 = (v42 << ((42) & 63))
                                var v58 uint8 = 0
                                _ = v58
                            }
                            var v59 uint16 = v1
                            v4 -= v4
                            _ = v45
                            _ = v59
                        default:
                            fmt.Println(fi41, v1, v40)
                        }
                        _ = v43
                        _ = v44
                    }
                    _ = fi41
                    _ = v42
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    if (-32768 > -1) {
                        if (127 >= -1) {
                            var v61 int16 = 1
                            _ = v61
                        } else {
                            if false {
                                fmt.Println(v4, fi39, v40)
                            } else {
                                v62 := GenS8{F0: genFn1((v1 * (v1 ^ v1)), fi39, -77)}
                                v1 += v1
                                _ = v62
                            }
                            v4 -= 0
                        }
                        var v63 uint8 = 128
                        v5 = v5
                        _ = v63
                    } else {
                        if (!false) {
                            v64 := GenS8{F0: fi60}
                            _ = v64
                        }
                        v5 = v5
                    }
                    _ = fi60
                }
                _ = fi39
                _ = v40
            }
        }
    }
    v5 = ""
    v2 = append(v2, 14)
    var v65 bool = (true && true)
    v5 = "a"
    switch v5 {
    case "":
        v5 = ("\n\t" + v5)
    case "foo bar":
        var v66 uint8 = 2
        _ = v66
    }
    for fi67 := 0; fi67 < 14; fi67++ {
        switch (v5 + v5) {
        case "":
            for fi68 := 0; fi68 < 10; fi68++ {
                switch (v36 % ((fi68) | 1)) {
                case -1:
                    if (1 < 0) {
                        v36 = (fi68 * (-(fi67)))
                    }
                case -565:
                    fmt.Println(fi68)
                case 0:
                    v69 := GenS7{F0: genFn0(true, 441), F1: int64(v4), F2: v1}
                    _ = v69
                }
                v1 -= v1
                _ = fi68
            }
        case "foo bar":
            var v70 int8 = 66
            _ = v70
        }
        v1 = (v1 ^ v1)
        _ = fi67
    }
    if (2 < v4) {
        var v71 string = "hello"
        v72 := GenS9{F0: 65535, F1: 42}
        var v73 uint = 938
        _ = v71
        _ = v72
        _ = v73
    }
    var v74 uint = 4294967295
    for fi75 := 0; fi75 < 15; fi75++ {
        var v76 int64 = -1
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 16; fi77++ {
        v78 := GenS10{F0: (v1 ^ (v1 / (((^(v1))) | 1))), F1: (v74 ^ (v74 + 4294967295))}
        v36 = 643
        switch (v5 + ((("hello" + v5) + ("hello" + v5)) + (v5 + v5))) {
        case "foo bar":
            switch 2147483647 {
            case 32767:
                switch v5 {
                case "":
                    _ = len(v2)
                case "a":
                    switch (v36 % (((fi77 / (((v36 % (((-(v36))) | 1))) | 1))) | 1)) {
                    case -9223372036854775808:
                        if v65 {
                            var v79 uint8 = 128
                            _ = v79
                        } else {
                            for fi80 := 0; fi80 < 7; fi80++ {
                                switch v5 {
                                case "hello":
                                    switch fi77 {
                                    case 2147483647:
                                        v74 = uint(v1)
                                    case 42:
                                        var v81 uint32 = 2147483647
                                        _ = v81
                                    default:
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                    }
                                    fmt.Println(v4)
                                    fmt.Println(v74, v36)
                                case "":
                                    v82 := GenS11{F0: v65}
                                    var v83 uint8 = 0
                                    _ = v82
                                    _ = v83
                                case "a":
                                    var v84 int = (fi77 + (v36 >> ((v74) & 63)))
                                    var v85 int64 = -1
                                    switch (v5 + "") {
                                    case "\n\t":
                                        if v65 {
                                            v6[384] = v5
                                            if true {
                                                v1 = (v1 - (^(v1)))
                                                _ = len(v2)
                                                var v86 bool = v65
                                                v5 = v5
                                                _ = v86
                                            } else {
                                                var v87 uint32 = 2147483647
                                                var v88 uint8 = 127
                                                _ = v87
                                                _ = v88
                                            }
                                            switch "" {
                                            case "foo bar":
                                                v89 := GenS9{F0: (v74 / ((v74) | 1)), F1: -128}
                                                _ = v89
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        }
                                    }
                                    _ = v84
                                    _ = v85
                                default:
                                    switch "foo bar" {
                                    case "":
                                        if v65 {
                                            fmt.Println(v5, v5)
                                            var v90 float32 = float32((0.1 - 0.0))
                                            var v91 uint32 = 958
                                            v91 -= v91
                                            v90 += v90
                                            _ = v90
                                            _ = v91
                                        }
                                    }
                                }
                                var v92 int = 2147483647
                                _ = fi80
                                _ = v92
                            }
                            switch v5 {
                            case "a":
                                var v93 uint8 = 0
                                _ = v93
                            }
                            v94 := GenS12{F0: 1}
                            _ = v94
                        }
                    case 891:
                        fmt.Println(v1, v65, v4)
                    case -79:
                        var v95 uint8 = uint8(v4)
                        _ = v95
                    }
                case "hello":
                    v96 := GenS9{F0: v74, F1: -128}
                    _ = v96
                }
            default:
                fmt.Println(v36)
            }
        case "\n\t":
            var v97 int64 = -851
            _ = v97
        case "":
            var v98 bool = ((-10 <= 127) && v65)
            _ = v98
        }
        _ = fi77
        _ = v78
    }
    fmt.Println(v65)
    var v99 float64 = 3.14159
    v100 := GenS13{F0: v1, F1: 1}
    v2 = append(v2, 639)
    for fi101 := 0; fi101 < 8; fi101++ {
        for fi102 := 0; fi102 < 7; fi102++ {
            fmt.Println(v65)
            if v65 {
                var v103 uint8 = 226
                _ = v103
            } else {
                if v65 {
                    v99 = float64(float32(v99))
                    v1 += (v1 % (((v1 / (((^(v1))) | 1))) | 1))
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                }
                switch int(v4) {
                case 32767:
                    _ = len(v6)
                }
                var v104 float32 = 3.14159
                _ = v104
            }
            var v105 uint8 = 1
            _ = fi102
            _ = v105
        }
        _ = fi101
    }
    v106 := GenS12{F0: (v4 | 862)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v65:", v65)
    fmt.Println("v74:", v74)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v106
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v74
    _ = v99
}
