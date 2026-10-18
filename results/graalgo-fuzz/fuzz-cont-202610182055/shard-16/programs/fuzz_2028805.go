package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 uint32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 int; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint16; F2 bool }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS11 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 uint8; F1 int64 }
type GenS15 struct { F0 int64; F1 uint32; F2 int16 }
type GenS16 struct { F0 uint64 }
func genFn0(p0 int16) uint {
    v7 := GenS1{F0: 419, F1: 0, F2: 115}
    for fi8 := 0; fi8 < 9; fi8++ {
        switch -596 {
        case -84:
            p0 = p0
        case 128:
            v9 := GenS1{F0: 295, F1: fi8, F2: 6}
            _ = v9
        case 32767:
            if (42 <= 0) {
                v10 := []int{32767, (fi8 % ((fi8) | 1)), 42, (fi8 + (fi8 % ((fi8) | 1)))}
                switch "" {
                case "foo bar":
                    fmt.Println(p0)
                case "a":
                    p0 = p0
                case "hello":
                    var v11 bool = false
                    _ = v11
                default:
                    var v12 int = (fi8 - (fi8 + 4294967295))
                    _ = v12
                }
                v13 := GenS1{F0: 1, F1: (fi8 * (fi8 / (((fi8 << ((18446744073709551615) & 63))) | 1))), F2: 1}
                _ = v10
                _ = v13
            } else {
                if (0 < 87) {
                    var v14 int = 128
                    _ = v14
                }
                switch "" {
                case "":
                    v15 := []int{(fi8 % ((-9223372036854775808) | 1)), (fi8 / ((fi8) | 1)), fi8}
                    _ = v15
                case "a":
                    fmt.Println(p0)
                case "foo bar":
                    if true {
                        var v16 bool = (65535 <= 255)
                        _ = v16
                    } else {
                        fmt.Println(p0)
                    }
                default:
                    v17 := GenS1{F0: 65535, F1: -1, F2: 127}
                    _ = v17
                }
            }
        }
        p0 += 1
        for fi18 := 0; fi18 < 8; fi18++ {
            v19 := GenS2{F0: 360, F1: 140, F2: 2147483647}
            var v20 int64 = 0
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 283
}
func genFn1() uint64 {
    v21 := GenS3{F0: "\n\t", F1: 32767}
    var v22 int32 = -505
    var v23 bool = (v22 >= (v22 * -2147483648))
    var v24 uint8 = 89
    switch "hello" {
    case "\n\t":
        var v25 bool = v23
        _ = v25
    case "foo bar":
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 5; fi27++ {
                if ((42 > 2147483647) || false) {
                    v28 := GenS4{F0: 65535, F1: (((4294967295 == fi27) && true) && true)}
                    var v29 uint8 = (^(v24))
                    _ = v28
                    _ = v29
                } else {
                    if v23 {
                        v24 += (v24 >> ((65535) & 63))
                        var v30 uint32 = 4294967295
                        var v31 uint32 = 700
                        _ = v30
                        _ = v31
                    } else {
                        v22 = v22
                        switch ("hello" + "") {
                        case "\n\t":
                            v22 = (-(v22))
                        case "foo bar":
                            switch fi26 {
                            case 42:
                                var v32 int8 = -1
                                if v23 {
                                    v33 := GenS5{F0: 65535, F1: 1, F2: (^(v22))}
                                    if true {
                                        v22 = (-(v22))
                                        v34 := []int{}
                                        var v35 int16 = 1000
                                        _ = v34
                                        _ = v35
                                    } else {
                                        var v36 float32 = (0.1 * 0.0)
                                        v37 := GenS1{F0: 65535, F1: (^(fi27)), F2: (v32 + (v32 * (v32 * -1)))}
                                        if (!true) {
                                            var v38 bool = v23
                                            var v39 string = "foo bar"
                                            v32 = (v32 * -26)
                                            fmt.Println(v24, fi27)
                                            v40 := []int{(fi27 & (fi27 << ((genFn0(42)) & 63))), fi27, 4294967295, fi27}
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                        v41 := []int{-9223372036854775808, (fi27 | fi26), fi26, int(v24)}
                                        _ = v36
                                        _ = v37
                                        _ = v41
                                    }
                                    v22 += v22
                                    _ = v33
                                } else {
                                    fmt.Println(v24, v23, v32)
                                    fmt.Println(fi27, v22, v22)
                                }
                                _ = v32
                            case -319:
                                v42 := []int{fi27, fi26, 173, (fi26 ^ 183)}
                                _ = v42
                            case -758:
                                var v43 uint = 65535
                                _ = v43
                            }
                        }
                    }
                }
                v22 += (v22 << ((4294967295) & 63))
                switch ("hello" + "a") {
                case "foo bar":
                    v23 = (((3.14159 + 3.14159) / (((float32(-2.25) / 0.1) * 3.14159) + (-2.25 * 1.5))) < (0.0 * float32(-2.25)))
                case "\n\t":
                    if (!v23) {
                        v44 := GenS6{F0: 128, F1: 280, F2: v23}
                        switch (("foo bar" + "hello") + "foo bar") {
                        case "\n\t":
                            var v45 int32 = v22
                            _ = v45
                        }
                        _ = v44
                    } else {
                        v23 = true
                        v46 := GenS7{F0: 1, F1: fi26}
                        if (!v23) {
                            v47 := GenS3{F0: (("hello" + ("foo bar" + "a")) + "hello"), F1: 187}
                            _ = v47
                        } else {
                            v22 += 0
                            v23 = (((true || (42 < int64(fi26))) && (0 == 44)) || v23)
                            if v23 {
                                v23 = v23
                                fmt.Println(v23, fi26)
                            }
                        }
                        _ = v46
                    }
                case "a":
                    v24 = v24
                default:
                    v22 = v22
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    var v2 int64 = 1
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    var v6 int32 = 42
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println(v2, v2)
    if true {
        v6 = (v6 / ((v6) | 1))
    } else {
        if (!false) {
            switch "foo bar" {
            case "\n\t":
                v2 = (v2 * v2)
            default:
                if (false || true) {
                    for fi48 := 0; fi48 < 4; fi48++ {
                        v6 = v6
                        v49 := GenS3{F0: "hello", F1: 1}
                        _ = fi48
                        _ = v49
                    }
                    var v50 uint8 = 181
                    v50 += (v50 * 226)
                    _ = v50
                }
            }
        } else {
            fmt.Println(v2, v2)
            for fi51 := 0; fi51 < 19; fi51++ {
                v2 = (v2 + 42)
                _ = fi51
            }
        }
    }
    v6 -= v6
    var v52 float32 = -2.25
    v53 := GenS8{F0: true, F1: 207}
    v54 := GenS4{F0: 65535, F1: (0 >= 65535)}
    var v55 int = -2147483648
    v56 := GenS8{F0: ((-(v2)) > (v2 << ((4294967295) & 63))), F1: 2}
    var v57 uint = 65535
    switch v55 {
    case 9223372036854775807:
        v2 = (v2 * v2)
    case 32767:
        switch v55 {
        case 32768:
            for fi58 := 0; fi58 < 2; fi58++ {
                var v59 uint8 = 2
                for fi60 := 0; fi60 < 8; fi60++ {
                    var v61 int16 = int16(v57)
                    _ = fi60
                    _ = v61
                }
                switch "a" {
                case "foo bar":
                    var v62 uint16 = 653
                    _ = v62
                case "hello":
                    v59 += (^(v59))
                }
                _ = fi58
                _ = v59
            }
        case 101:
            if ((v6 == (v6 ^ v6)) || false) {
                if ((0.0 + -2.25) >= (1.5 * 1.5)) {
                    v63 := v4["hello"]
                    v64 := GenS9{F0: 418, F1: 42}
                    if (v63 && (53 >= 42)) {
                        var v65 bool = v63
                        _ = v65
                    }
                    _ = v63
                    _ = v64
                }
                switch 494 {
                case 2147483647:
                    if (42 > -128) {
                        fmt.Println(v57, v6)
                        var v66 int64 = v2
                        _ = v66
                    }
                default:
                    v6 = v6
                }
                v67 := GenS10{F0: 2147483648, F1: (v2 & (v2 << ((v57) & 63))), F2: 1}
                _ = v67
            }
        }
    default:
        delete(v4, "a")
    }
    v68 := GenS11{F0: 631, F1: 20, F2: 255}
    if false {
        if (361 == (^(v57))) {
            var v69 int8 = 1
            _ = v69
        } else {
            var v70 bool = ((v2 / ((0) | 1)) != (v2 & 9223372036854775807))
            var v71 uint64 = 809
            v72 := GenS1{F0: 1, F1: v55, F2: 100}
            _ = v70
            _ = v71
            _ = v72
        }
        v73 := GenS8{F0: true, F1: 32767}
        v2 = (v2 - (v2 * v2))
        _ = v73
    }
    var v74 int8 = -111
    fmt.Println(v55, v2, v57)
    delete(v3, "foo bar")
    for fi75 := 0; fi75 < 17; fi75++ {
        if true {
            for fi76 := 0; fi76 < 9; fi76++ {
                switch "a" {
                case "a":
                    v77 := GenS8{F0: false, F1: 844}
                    _ = v77
                case "hello":
                    v57 -= 255
                }
                v78 := GenS12{F0: "\n\t", F1: (^(v74))}
                v79 := GenS13{F0: 1, F1: 0}
                _ = fi76
                _ = v78
                _ = v79
            }
            if false {
                if (float32(((1e10 / 1e10) + (0.0 / 1.0))) < v52) {
                    fmt.Println(v57, v55, v74)
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            if (!(!((v57 | v57) <= 800))) {
                                switch int(v57) {
                                case 32767:
                                    fmt.Println(v2, v74)
                                case 4294967295:
                                    _ = len(v3)
                                case -9223372036854775808:
                                    var v80 uint64 = genFn1()
                                    _ = v80
                                }
                                var v81 int = (v55 + v55)
                                v6 += -276
                                _ = v81
                            }
                        default:
                            switch fi75 {
                            case -932:
                                for fi82 := 0; fi82 < 17; fi82++ {
                                    if false {
                                        var v83 int = (fi75 | fi75)
                                        _ = v83
                                    }
                                    _ = fi82
                                }
                            case 127:
                                var v84 uint16 = 255
                                _ = v84
                            }
                        }
                        if (false && true) {
                            v85 := GenS14{F0: 127, F1: v2}
                            v86 := GenS3{F0: "", F1: 42}
                            var v87 float64 = 0.0
                            _ = v85
                            _ = v86
                            _ = v87
                        } else {
                            delete(v4, "hello")
                            if (-1 < 0) {
                                v88 := GenS3{F0: "hello", F1: 1000}
                                var v89 float64 = ((float64(v52) + -2.25) / 1e10)
                                var v90 int = (-(fi75))
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                        }
                    }
                    if ((v57 + v57) <= v57) {
                        for fi91 := 0; fi91 < 13; fi91++ {
                            switch fi75 {
                            case 9223372036854775807:
                                var v92 int = fi91
                                _ = v92
                            case 32768:
                                if (false && false) {
                                    fmt.Println(v74)
                                    if ((fi75 >> (((^(v57))) & 63)) < 933) {
                                        v6 = 525
                                    } else {
                                        _ = len(v4)
                                    }
                                } else {
                                    if true {
                                        var v93 float32 = (v52 * v52)
                                        v94 := v3["hello"]
                                        v3[""] = (fi75 << (((v57 / ((v57) | 1))) & 63))
                                        _ = v93
                                        _ = v94
                                    }
                                }
                            }
                            fmt.Println(v57, v52, fi75)
                            _ = fi91
                        }
                    } else {
                        _ = len(v4)
                        v52 -= (v52 * (v52 / v52))
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                    }
                } else {
                    switch "hello" {
                    case "foo bar":
                        v74 = (v74 + v74)
                    default:
                        v95 := GenS6{F0: 21, F1: 65535, F2: (2 != 65535)}
                        _ = v95
                    }
                    switch (^(fi75)) {
                    case 2147483647:
                        delete(v5, (^(fi75)))
                    case 9223372036854775807:
                        v57 = (^(v57))
                    }
                }
            } else {
                for fi96 := 0; fi96 < 16; fi96++ {
                    fmt.Println(v55, v74)
                    _ = fi96
                }
                v52 = v52
            }
            var v97 uint32 = 65535
            _ = v97
        } else {
            switch -112 {
            case 215:
                v57 = 1
            case 127:
                var v98 int32 = (v6 ^ (v6 ^ (v6 ^ (v6 >> ((2) & 63)))))
                _ = v98
            }
            v99 := v4[""]
            _ = v99
        }
        var v100 uint = v57
        _ = fi75
        _ = v100
    }
    for fi101 := 0; fi101 < 12; fi101++ {
        switch 42 {
        case 216:
            var v102 int8 = (-(v74))
            _ = v102
        case -9223372036854775808:
            v103 := GenS15{F0: (v2 / ((9223372036854775807) | 1)), F1: 2147483648, F2: -32768}
            _ = v103
        case -790:
            v104 := GenS16{F0: genFn1()}
            _ = v104
        }
        fmt.Println(v52)
        for fi105 := 0; fi105 < 7; fi105++ {
            v57 -= (^(v57))
            v106 := GenS5{F0: 100, F1: 0, F2: (-(v6))}
            var v107 float64 = -2.25
            _ = fi105
            _ = v106
            _ = v107
        }
        _ = fi101
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v6
    _ = v68
    _ = v74
}
