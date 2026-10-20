package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint16; F1 int64; F2 int16 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int8; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 bool }
type GenS12 struct { F0 uint64; F1 int32; F2 int }
type GenS13 struct { F0 int16; F1 int16; F2 uint16 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int64; F1 int }
func genFn0(p0 string, p1 uint) int64 {
    if (!("foo bar" != ("\n\t" + ("" + "foo bar")))) {
        v6 := GenS1{F0: true}
        if false {
            var v7 int32 = 0
            _ = v7
        }
        var v8 int16 = 1000
        _ = v6
        _ = v8
    }
    v9 := GenS1{F0: true}
    p1 = p1
    _ = p0
    _ = p1
    _ = v9
    return -1
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    var v3 int = 4294967295
    v4 := []int{v3, 1, v3}
    v5 := make(map[string]int)
    var v10 bool = ((!false) && (!(0 < 1000)))
    v10 = (1 <= 1)
    v11 := GenS2{F0: 1, F1: true}
    fmt.Println(v1)
    v3 += (v3 / ((v3) | 1))
    switch (("a" + "\n\t") + "") {
    case "":
        v4 = append(v4, v3)
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v1, v1)
    default:
        v5[(("" + "hello") + "")] = v3
    }
    v12 := GenS3{F0: 947, F1: -282, F2: v1}
    var v13 uint16 = 42
    if v10 {
        var v14 uint64 = 772
        v3 = v3
        var v15 uint16 = v13
        _ = v14
        _ = v15
    } else {
        switch 32768 {
        case -713:
            var v16 string = ("\n\t" + ("hello" + "foo bar"))
            _ = v16
        default:
            v10 = ("" > "\n\t")
        }
    }
    if true {
        var v17 int64 = 124
        _ = v17
    } else {
        v4 = append(v4, (-(v3)))
        v18 := GenS4{F0: v1, F1: 32767}
        _ = v18
    }
    v19 := GenS5{F0: 1, F1: 106}
    var v20 int8 = -128
    v20 += v20
    v21 := GenS2{F0: (^(v13)), F1: (!("hello" > "\n\t"))}
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS6{F0: v13, F1: 0}
        for fi24 := 0; fi24 < 13; fi24++ {
            v1 = (v1 * (v1 << ((743) & 63)))
            var v25 int8 = v20
            _ = fi24
            _ = v25
        }
        _ = fi22
        _ = v23
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 string = ""
        if (570 <= 1) {
            fmt.Println(v20, v13, v3)
        } else {
            v28 := GenS1{F0: (true || v10)}
            var v29 bool = true
            switch (fi26 | (-(fi26))) {
            case -1:
                v4 = append(v4, -1)
            case 0:
                for fi30 := 0; fi30 < 7; fi30++ {
                    var v31 float64 = float64(((float32(((float64(3.14159) * 1.5) + 1.5)) + (-2.25 / (1e10 - ((0.0 - ((1e10 - 0.0) * 1e10)) * 0.0)))) / 1e10))
                    var v32 bool = (!true)
                    var v33 float64 = v31
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
            default:
                v34 := GenS6{F0: (v13 | v13), F1: 255}
                _ = v34
            }
            _ = v28
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    if (!true) {
        var v35 int = 32768
        v20 += (-(v20))
        if (32768 > -687) {
            switch int(v13) {
            case 9223372036854775807:
                if true {
                    var v36 float64 = -2.25
                    if (!(0 != 42)) {
                        fmt.Println(v13, v36, v10)
                        fmt.Println(v13, v3, v20)
                        if v10 {
                            v10 = ((!(v3 > -968)) && false)
                        }
                    } else {
                        v35 -= (v3 << ((uint(v13)) & 63))
                    }
                    var v37 float32 = 1.5
                    _ = v36
                    _ = v37
                }
            case -309:
                if false {
                    v38 := GenS5{F0: 42, F1: 1}
                    for fi39 := 0; fi39 < 13; fi39++ {
                        _ = len(v2)
                        for fi40 := 0; fi40 < 6; fi40++ {
                            var v41 int16 = (v1 & (v1 ^ v1))
                            var v42 int8 = v20
                            v10 = (v10 || ("\n\t" != "a"))
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                        switch (v3 - (v35 | (-(fi39)))) {
                        case 32768:
                            switch "foo bar" {
                            case "hello":
                                v1 = (v1 << ((1) & 63))
                            case "foo bar":
                                var v43 int8 = -128
                                _ = v43
                            case "a":
                                var v44 uint8 = 139
                                _ = v44
                            default:
                                switch "hello" {
                                case "hello":
                                    fmt.Println(v3, v3)
                                case "\n\t":
                                    switch int(v13) {
                                    case 128:
                                        if true {
                                            if false {
                                                v1 += v1
                                                if (-1 > (v1 / ((v1) | 1))) {
                                                    if len(v4) > 0 {
                                                        v4[0] = v3
                                                    }
                                                    if v10 {
                                                        v45 := GenS7{F0: v20, F1: 2147483647}
                                                        v46 := GenS3{F0: (v13 % ((0) | 1)), F1: 1, F2: 32767}
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        if (!false) {
                                                            v13 = v13
                                                        } else {
                                                            v47 := GenS6{F0: 65535, F1: 1}
                                                            _ = v47
                                                        }
                                                        v48 := GenS5{F0: 42, F1: 237}
                                                        _ = v48
                                                    }
                                                } else {
                                                    var v49 uint = 255
                                                    _ = v49
                                                }
                                            } else {
                                                _ = len(v4)
                                                switch v3 {
                                                case 9223372036854775807:
                                                    switch "a" {
                                                    case "foo bar":
                                                        for fi50 := 0; fi50 < 11; fi50++ {
                                                            v13 = (v13 / (((v13 + 1)) | 1))
                                                            v20 += (v20 % (((v20 ^ v20)) | 1))
                                                            _ = fi50
                                                        }
                                                    case "a":
                                                        v3 = (v35 / ((v35) | 1))
                                                    case "\n\t":
                                                        if v10 {
                                                            for fi51 := 0; fi51 < 8; fi51++ {
                                                                fmt.Println(v13, v35, v10)
                                                                v20 = 127
                                                                if true {
                                                                    switch (^(fi51)) {
                                                                    case 109:
                                                                        v52 := GenS3{F0: v13, F1: int64(fi39), F2: (v1 & v1)}
                                                                        _ = v52
                                                                    case -1:
                                                                        if v10 {
                                                                            v13 = v13
                                                                        } else {
                                                                            v1 += (v1 >> ((2) & 63))
                                                                            v53 := GenS5{F0: 42, F1: 66}
                                                                            var v54 float64 = -2.25
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                    case 32767:
                                                                        _ = len(v4)
                                                                        if (!(!false)) {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if (!v10) {
                                                                                    fmt.Println(fi39, v3, fi39)
                                                                                } else {
                                                                                    if v10 {
                                                                                        v10 = (false || (412 < 251))
                                                                                        if v10 {
                                                                                            var v55 float32 = 1e10
                                                                                            v35 = 128
                                                                                            _ = v55
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v56 uint16 = 42
                                                                        v20 = v20
                                                                        v13 -= v56
                                                                        _ = v56
                                                                    }
                                                                    v20 += (v20 / ((42) | 1))
                                                                } else {
                                                                    v57 := GenS8{F0: 511, F1: 0, F2: 231}
                                                                    v58 := v2[2]
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                v4 = append(v4, fi39)
                                                                var v59 uint8 = 127
                                                                _ = fi51
                                                                _ = v59
                                                            }
                                                            fmt.Println(v35)
                                                        } else {
                                                            for fi60 := 0; fi60 < 5; fi60++ {
                                                                var v61 float32 = 0.1
                                                                if true {
                                                                    if false {
                                                                        v3 = 4294967295
                                                                    } else {
                                                                        var v62 uint8 = 255
                                                                        fmt.Println(fi39, v3, fi60)
                                                                        _ = v62
                                                                    }
                                                                }
                                                                if len(v4) > 0 {
                                                                    v4[0] = v3
                                                                }
                                                                _ = fi60
                                                                _ = v61
                                                            }
                                                        }
                                                    default:
                                                        v63 := GenS9{F0: 128, F1: 2147483648}
                                                        _ = v63
                                                    }
                                                default:
                                                    v64 := GenS4{F0: (v1 % (((^(v1))) | 1)), F1: v1}
                                                    _ = v64
                                                }
                                            }
                                            v65 := GenS3{F0: v13, F1: 670, F2: (v1 << ((18446744073709551615) & 63))}
                                            v66 := GenS9{F0: 2, F1: 65535}
                                            _ = v65
                                            _ = v66
                                        } else {
                                            var v67 int8 = 1
                                            _ = v67
                                        }
                                    case -2147483648:
                                        if (v10 || false) {
                                            delete(v2, 399)
                                        }
                                    case -1:
                                        v68 := GenS9{F0: 187, F1: 2147483647}
                                        _ = v68
                                    default:
                                        if v10 {
                                            if ((^(v20)) != (v20 - 127)) {
                                                v69 := GenS4{F0: 0, F1: v1}
                                                var v70 bool = (65535 >= 663)
                                                v71 := GenS10{F0: 42}
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            } else {
                                                v1 += v1
                                            }
                                            var v72 float64 = ((0.1 * float64(3.14159)) / 1.0)
                                            var v73 int8 = (v20 ^ 0)
                                            _ = v72
                                            _ = v73
                                        }
                                    }
                                case "foo bar":
                                    v1 = (v1 | (^(v1)))
                                }
                            }
                        case -76:
                            for fi74 := 0; fi74 < 11; fi74++ {
                                var v75 string = "\n\t"
                                _ = fi74
                                _ = v75
                            }
                        case 0:
                            for fi76 := 0; fi76 < 9; fi76++ {
                                if true {
                                    var v77 uint8 = 1
                                    var v78 int = 2147483647
                                    fmt.Println(v77)
                                    _ = v77
                                    _ = v78
                                } else {
                                    switch "a" {
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            v1 += v1
                                        case "\n\t":
                                            if ((-39 <= 1) || v10) {
                                                switch ((((("" + "\n\t") + (("a" + "hello") + "hello")) + "a") + "") + "\n\t") {
                                                case "":
                                                    var v79 int8 = (v20 / (((v20 << ((1) & 63))) | 1))
                                                    _ = v79
                                                default:
                                                    if (!v10) {
                                                        v80 := GenS11{F0: 42, F1: ((v10 || v10) && true)}
                                                        var v81 uint8 = 1
                                                        _ = len(v4)
                                                        var v82 uint8 = (v81 | v81)
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                    } else {
                                                        var v83 int8 = v20
                                                        var v84 int32 = 0
                                                        v35 = 32768
                                                        v85 := GenS12{F0: 2, F1: 2147483647, F2: 32767}
                                                        v5["hello"] = -1
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                }
                                                var v86 int16 = (v1 << ((831) & 63))
                                                v86 += 626
                                                _ = v86
                                            }
                                        }
                                    case "a":
                                        v87 := GenS13{F0: v1, F1: v1, F2: v13}
                                        _ = v87
                                    case "\n\t":
                                        v88 := GenS2{F0: v13, F1: v10}
                                        var v89 uint16 = v13
                                        if true {
                                            var v90 uint16 = v13
                                            var v91 uint32 = 65535
                                            var v92 uint = 1
                                            fmt.Println(fi39)
                                            var v93 int16 = (-(v1))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v35 = v35
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                            _ = v93
                                        } else {
                                            v94 := GenS14{F0: v89}
                                            fmt.Println(v89)
                                            _ = v94
                                        }
                                        _ = v88
                                        _ = v89
                                    default:
                                        v35 -= (fi39 << ((18446744073709551615) & 63))
                                    }
                                    var v95 string = "\n\t"
                                    v96 := GenS2{F0: 447, F1: (int8(v1) == v20)}
                                    var v97 uint8 = 1
                                    var v98 uint64 = 42
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                                _ = fi76
                            }
                        }
                        _ = fi39
                    }
                    _ = v38
                } else {
                    v1 += -891
                    v99 := GenS3{F0: (v13 & v13), F1: -529, F2: -545}
                    v100 := GenS15{F0: -459, F1: v35}
                    _ = v100
                    _ = v99
                }
            default:
                v4 = append(v4, v35)
            }
            var v101 int = v3
            delete(v5, "")
            _ = v101
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
