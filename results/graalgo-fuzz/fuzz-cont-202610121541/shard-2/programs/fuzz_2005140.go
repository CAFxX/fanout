package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 string; F2 int16 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 int16; F1 int; F2 bool }
type GenS10 struct { F0 uint16; F1 int16; F2 string }
type GenS11 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint; F1 uint32; F2 int8 }
type GenS16 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 uint) uint32 {
    var v5 uint8 = 128
    p0 += (p0 % ((p0) | 1))
    p0 = p0
    p0 = p0
    var v6 uint64 = 42
    _ = p0
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{42, 65535}
    var v3 int16 = 32767
    v4 := make(map[uint32]uint32)
    var v7 uint64 = 2
    switch ("\n\t" + ("hello" + "hello")) {
    case "a":
        var v8 uint32 = uint32(v7)
        _ = v8
    case "":
        for fi9 := 0; fi9 < 15; fi9++ {
            v1["\n\t"] = true
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 255, F1: 255}
                if true {
                    v7 -= v7
                } else {
                    if ((-1 >= -2147483648) && true) {
                        fmt.Println(v3, v7)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                    } else {
                        if ((v3 - 32767) != -32768) {
                            v3 -= v3
                        } else {
                            var v12 int8 = 19
                            _ = v12
                        }
                        var v13 int32 = -2147483648
                        _ = v13
                    }
                    if true {
                        if false {
                            v7 = (v7 / ((v7) | 1))
                            fmt.Println(v7, v3, v7)
                            var v14 uint8 = 1
                            _ = v14
                        }
                    }
                    if (127 <= -83) {
                        var v15 float32 = float32(3.14159)
                        v16 := GenS1{F0: 42, F1: 255}
                        _ = v15
                        _ = v16
                    }
                    switch "" {
                    case "hello":
                        if len(v2) > 0 {
                            v2[0] = genFn0(255)
                        }
                    }
                    v3 -= v3
                    var v17 int64 = 9223372036854775807
                    _ = v17
                }
                if (4294967295 != 42) {
                    v7 -= v7
                    v18 := GenS1{F0: 42, F1: 42}
                    fmt.Println(fi10)
                    var v19 int8 = -128
                    _ = v18
                    _ = v19
                } else {
                    switch 128 {
                    case -1:
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            _ = len(v1)
                        case "":
                            if true {
                                _ = len(v1)
                                var v20 int32 = 531
                                _ = v20
                            } else {
                                _ = len(v2)
                                if (0 == 0) {
                                    var v21 bool = (118 <= 0)
                                    fmt.Println(v7, v21, v21)
                                    v2 = append(v2, 65535)
                                    _ = v21
                                }
                            }
                            if (!((-(v3)) <= v3)) {
                                v7 = (v7 >> ((255) & 63))
                                var v22 int16 = 0
                                switch (fi10 << ((1) & 63)) {
                                case 32768:
                                    v23 := v1["foo bar"]
                                    _ = v23
                                }
                                _ = v22
                            } else {
                                delete(v1, ("\n\t" + "hello"))
                                if true {
                                    if (118 >= 127) {
                                        switch (fi10 * fi10) {
                                        case -2147483648:
                                            v24 := GenS1{F0: 17, F1: 2}
                                            _ = v24
                                        case -9223372036854775808:
                                            var v25 uint16 = 42
                                            _ = v25
                                        case -65:
                                            if true {
                                                v7 -= v7
                                                var v26 uint = 805
                                                var v27 int = (fi10 << (((v26 | v26)) & 63))
                                                _ = v26
                                                _ = v27
                                            }
                                            if true {
                                                switch "a" {
                                                case "\n\t":
                                                    v7 -= 1
                                                case "foo bar":
                                                    v3 += (-(v3))
                                                case "":
                                                    fmt.Println(v7)
                                                default:
                                                    if ("foo bar" < ("a" + "a")) {
                                                        v3 = (v3 / (((v3 | v3)) | 1))
                                                        var v28 uint16 = 255
                                                        fmt.Println(fi9, v7)
                                                        v28 += (v28 - (^(v28)))
                                                        if false {
                                                            fmt.Println(fi9)
                                                            v28 = (v28 ^ 42)
                                                            v4[uint32(v28)] = genFn0(1)
                                                        } else {
                                                            var v29 float64 = 0.1
                                                            if (true && (825 < 2147483647)) {
                                                                var v30 int64 = -932
                                                                _ = v30
                                                            } else {
                                                                v28 = (v28 - (v28 >> ((255) & 63)))
                                                                var v31 uint64 = (v7 >> ((2) & 63))
                                                                _ = v31
                                                            }
                                                            _ = v29
                                                        }
                                                        fmt.Println(v3, v3, fi10)
                                                        _ = v28
                                                    } else {
                                                        switch fi10 {
                                                        case -859:
                                                            v3 -= -32768
                                                        case 32767:
                                                            _ = len(v2)
                                                        case 4294967295:
                                                            _ = len(v4)
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                    }
                                                }
                                                _ = len(v4)
                                                switch 0 {
                                                case -2147483648:
                                                    if true {
                                                        v3 += (v3 * v3)
                                                        v32 := GenS1{F0: 91, F1: 42}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if true {
                                                            if (true && false) {
                                                                fmt.Println(fi9, v7)
                                                                v7 -= (v7 / ((96) | 1))
                                                                v2 = append(v2, 942)
                                                                if (false || false) {
                                                                    fmt.Println(fi10)
                                                                } else {
                                                                    if true {
                                                                        if (!false) {
                                                                            v33 := v4[genFn0(1)]
                                                                            _ = v33
                                                                        }
                                                                        fmt.Println(v3)
                                                                    } else {
                                                                        fmt.Println(fi10, fi9)
                                                                        var v34 string = "hello"
                                                                        var v35 uint32 = 0
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                    if true {
                                                                        _ = len(v4)
                                                                    } else {
                                                                        if ("foo bar" > "a") {
                                                                            v3 -= 802
                                                                        } else {
                                                                            fmt.Println(v7, v3, v7)
                                                                            fmt.Println(fi9, fi9, fi10)
                                                                            if false {
                                                                                fmt.Println(fi9)
                                                                            } else {
                                                                                switch fi9 {
                                                                                case -575:
                                                                                    var v36 uint = 118
                                                                                    _ = v36
                                                                                case 1:
                                                                                    v3 = (v3 % ((v3) | 1))
                                                                                case 2147483647:
                                                                                    switch fi9 {
                                                                                    case -875:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    case 1:
                                                                                        switch fi9 {
                                                                                        case 109:
                                                                                            v37 := GenS1{F0: 1, F1: 160}
                                                                                            _ = v37
                                                                                        case 2147483647:
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                v3 -= 1
                                                                                            case "":
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = 65535
                                                                                                }
                                                                                            case "foo bar":
                                                                                                if (v3 == (v3 >> ((523) & 63))) {
                                                                                                    fmt.Println(v7, v3)
                                                                                                    var v38 uint = 255
                                                                                                    var v39 int16 = 32767
                                                                                                    _ = v38
                                                                                                    _ = v39
                                                                                                } else {
                                                                                                    v3 = (^(v3))
                                                                                                    v40 := GenS1{F0: 128, F1: 0}
                                                                                                    fmt.Println(v3, fi10, v7)
                                                                                                    _ = v40
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    v7 = 42
                                                                                                case "hello":
                                                                                                    fmt.Println(v3)
                                                                                                case "a":
                                                                                                    v41 := GenS2{F0: true, F1: -9223372036854775808}
                                                                                                    _ = v41
                                                                                                }
                                                                                            default:
                                                                                                if true {
                                                                                                    switch fi10 {
                                                                                                    case 32768:
                                                                                                        var v42 uint8 = 107
                                                                                                        v3 = v3
                                                                                                        fmt.Println(v7)
                                                                                                        _ = v42
                                                                                                    }
                                                                                                } else {
                                                                                                    var v43 int64 = 42
                                                                                                    var v44 int32 = -220
                                                                                                    _ = v43
                                                                                                    _ = v44
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v45 int8 = 90
                                                                                    _ = v45
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v46 uint16 = 2
                                                                    _ = v46
                                                                }
                                                                if false {
                                                                    switch -403 {
                                                                    case 1:
                                                                        _ = len(v4)
                                                                    case 42:
                                                                        v7 = v7
                                                                    default:
                                                                        if false {
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                var v47 uint16 = 32767
                                                                                _ = v47
                                                                            case "foo bar":
                                                                                var v48 int = (^(fi9))
                                                                                v48 = fi9
                                                                                _ = v48
                                                                            }
                                                                        }
                                                                    }
                                                                    v3 -= v3
                                                                    v49 := GenS1{F0: 2, F1: 156}
                                                                    _ = v49
                                                                }
                                                            } else {
                                                                switch (("\n\t" + "a") + "hello") {
                                                                case "\n\t":
                                                                    v7 -= v7
                                                                default:
                                                                    v50 := GenS2{F0: false, F1: 0}
                                                                    _ = v50
                                                                }
                                                                if (255 != 937) {
                                                                    var v51 uint64 = (v7 & v7)
                                                                    var v52 uint32 = genFn0(65535)
                                                                    if ("" > "") {
                                                                        if true {
                                                                            fmt.Println(fi9, v52, v51)
                                                                            var v53 uint32 = (^(v52))
                                                                            v54 := GenS2{F0: (true || true), F1: 441}
                                                                            _ = v53
                                                                            _ = v54
                                                                        } else {
                                                                            v2 = append(v2, v52)
                                                                            var v55 int8 = 127
                                                                            v7 = v51
                                                                            v56 := GenS2{F0: false, F1: 42}
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                        v57 := GenS1{F0: 0, F1: 255}
                                                                        _ = v57
                                                                    } else {
                                                                        if true {
                                                                            if false {
                                                                                if (!false) {
                                                                                    fmt.Println(v52, v51)
                                                                                }
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (^(v52))
                                                                            }
                                                                            v3 -= v3
                                                                            var v58 uint = 0
                                                                            v3 = v3
                                                                            _ = v58
                                                                        }
                                                                        var v59 uint = 2
                                                                        var v60 int32 = 487
                                                                        if len(v2) > 0 {
                                                                            v2[0] = genFn0(uint(v51))
                                                                        }
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                    if (806 >= 9223372036854775807) {
                                                                        v7 = 2
                                                                        var v61 float64 = ((1.5 - ((float64(0.0) * 1e10) / 1.0)) * (0.1 - (float64(1.5) / float64((1.5 + 1.5)))))
                                                                        v52 += (v52 << ((42) & 63))
                                                                        if true {
                                                                            switch 32768 {
                                                                            case -365:
                                                                                var v62 string = ("hello" + "a")
                                                                                _ = v62
                                                                            default:
                                                                                switch ("hello" + ("foo bar" + ("a" + "a"))) {
                                                                                case "foo bar":
                                                                                    v52 -= genFn0(0)
                                                                                case "\n\t":
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        _ = len(v2)
                                                                                    case "":
                                                                                        v63 := GenS1{F0: uint8(v7), F1: 128}
                                                                                        _ = v63
                                                                                    }
                                                                                case "hello":
                                                                                    v61 -= 1.5
                                                                                default:
                                                                                    _ = len(v2)
                                                                                }
                                                                            }
                                                                            switch ((("foo bar" + "hello") + "") + (("a" + "hello") + "hello")) {
                                                                            case "\n\t":
                                                                                var v64 uint64 = 18446744073709551615
                                                                                _ = v64
                                                                            case "":
                                                                                v65 := GenS1{F0: 127, F1: 128}
                                                                                _ = v65
                                                                            case "a":
                                                                                delete(v1, "")
                                                                            default:
                                                                                v2 = append(v2, 2147483648)
                                                                            }
                                                                        }
                                                                        var v66 string = "a"
                                                                        _ = v61
                                                                        _ = v66
                                                                    } else {
                                                                        _ = len(v2)
                                                                        v51 = v51
                                                                        v67 := GenS3{F0: "\n\t"}
                                                                        _ = v67
                                                                    }
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                            if false {
                                                                var v68 bool = false
                                                                v68 = v68
                                                                _ = v68
                                                            }
                                                            if false {
                                                                var v69 uint = 901
                                                                _ = v69
                                                            } else {
                                                                fmt.Println(fi10)
                                                                var v70 int64 = 407
                                                                v70 = v70
                                                                v70 += (v70 + v70)
                                                                _ = v70
                                                            }
                                                        } else {
                                                            switch "\n\t" {
                                                            case "":
                                                                if false {
                                                                    v3 = 0
                                                                } else {
                                                                    v3 += (v3 - -67)
                                                                    if (!(int16(v7) > v3)) {
                                                                        var v71 uint16 = 2
                                                                        if false {
                                                                            if (v3 < v3) {
                                                                                var v72 int64 = -1
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    var v73 uint8 = 2
                                                                                    _ = v73
                                                                                case "foo bar":
                                                                                    var v74 uint16 = 1
                                                                                    _ = v74
                                                                                case "a":
                                                                                    if false {
                                                                                        v75 := GenS4{F0: 0, F1: 923, F2: 255}
                                                                                        var v76 int = -823
                                                                                        var v77 uint64 = 18446744073709551615
                                                                                        _ = v75
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    }
                                                                                default:
                                                                                    switch (fi10 & (fi9 & fi10)) {
                                                                                    case 670:
                                                                                        v78 := GenS5{F0: (v72 + (v72 >> ((18446744073709551615) & 63)))}
                                                                                        _ = v78
                                                                                    case 32768:
                                                                                        v79 := GenS6{F0: 0, F1: ("a" + (("foo bar" + "a") + "a")), F2: v3}
                                                                                        _ = v79
                                                                                    case 46:
                                                                                        v71 += (v71 - v71)
                                                                                    default:
                                                                                        var v80 bool = (!(84 <= 1))
                                                                                        _ = v80
                                                                                    }
                                                                                }
                                                                                _ = v72
                                                                            } else {
                                                                                var v81 uint64 = v7
                                                                                _ = v81
                                                                            }
                                                                            var v82 int8 = -119
                                                                            _ = len(v2)
                                                                            _ = v82
                                                                        }
                                                                        _ = v71
                                                                    }
                                                                    v83 := GenS7{F0: 42, F1: v7}
                                                                    if (100 > int8(v3)) {
                                                                        v84 := GenS2{F0: false, F1: -284}
                                                                        _ = v84
                                                                    } else {
                                                                        v3 = (v3 | v3)
                                                                    }
                                                                    _ = v83
                                                                }
                                                            case "foo bar":
                                                                switch -9223372036854775808 {
                                                                case 128:
                                                                    var v85 bool = (true || false)
                                                                    _ = v85
                                                                default:
                                                                    fmt.Println(v7)
                                                                }
                                                                switch fi10 {
                                                                case 2147483647:
                                                                    if true {
                                                                        v86 := GenS8{F0: 882, F1: 4294967295}
                                                                        var v87 uint16 = 2
                                                                        _ = v86
                                                                        _ = v87
                                                                    } else {
                                                                        v88 := v1["a"]
                                                                        if true {
                                                                            v3 -= v3
                                                                            var v89 uint32 = 42
                                                                            _ = v89
                                                                        }
                                                                        v90 := GenS9{F0: (v3 % ((v3) | 1)), F1: (fi10 % ((127) | 1)), F2: true}
                                                                        v91 := GenS1{F0: 255, F1: 127}
                                                                        v92 := GenS6{F0: 0, F1: "foo bar", F2: int16(fi10)}
                                                                        fmt.Println(v7, v3, v88)
                                                                        v3 = 14
                                                                        if v88 {
                                                                            var v93 int8 = -128
                                                                            v3 = (v3 / (((v3 + v3)) | 1))
                                                                            v94 := GenS10{F0: 2, F1: (v3 + v3), F2: "foo bar"}
                                                                            _ = v93
                                                                            _ = v94
                                                                        }
                                                                        var v95 float32 = (1.5 + 3.14159)
                                                                        _ = v88
                                                                        _ = v90
                                                                        _ = v91
                                                                        _ = v92
                                                                        _ = v95
                                                                    }
                                                                case 32767:
                                                                    v7 -= 1
                                                                }
                                                                v96 := GenS2{F0: false, F1: 992}
                                                                _ = v96
                                                            default:
                                                                v97 := GenS10{F0: 42, F1: (v3 - 1000), F2: ("a" + "\n\t")}
                                                                _ = v97
                                                            }
                                                        }
                                                        v7 = 4294967295
                                                        switch ("hello" + "foo bar") {
                                                        case "foo bar":
                                                            if false {
                                                                if (!false) {
                                                                    v3 = v3
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                } else {
                                                                    switch fi10 {
                                                                    case -806:
                                                                        _ = len(v1)
                                                                    case 375:
                                                                        var v98 int64 = -9223372036854775808
                                                                        _ = v98
                                                                    case 4294967295:
                                                                        if false {
                                                                            var v99 float64 = ((1.5 + (0.1 + 1.5)) / 1.5)
                                                                            _ = v99
                                                                        }
                                                                    default:
                                                                        v100 := GenS5{F0: 42}
                                                                        _ = v100
                                                                    }
                                                                    var v101 int64 = -9223372036854775808
                                                                    _ = v101
                                                                }
                                                                v102 := GenS8{F0: 0, F1: 4294967295}
                                                                _ = v102
                                                            } else {
                                                                if true {
                                                                    v103 := GenS11{F0: ((!true) || true), F1: 2, F2: 510}
                                                                    var v104 bool = ((3.14159 > 0.1) || true)
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 842
                                                                    }
                                                                    _ = v103
                                                                    _ = v104
                                                                } else {
                                                                    v7 = (v7 - (v7 << ((0) & 63)))
                                                                    var v105 float32 = float32((3.14159 + (1.5 / (3.14159 + -2.25))))
                                                                    _ = v105
                                                                }
                                                                v7 = v7
                                                                var v106 float32 = -2.25
                                                                _ = v106
                                                            }
                                                        }
                                                        if true {
                                                            _ = len(v2)
                                                            v107 := GenS12{F0: 112}
                                                            _ = v107
                                                        } else {
                                                            if false {
                                                                var v108 int32 = -370
                                                                _ = v108
                                                            } else {
                                                                v109 := GenS1{F0: 243, F1: 0}
                                                                if len(v2) > 0 {
                                                                    v2[0] = 42
                                                                }
                                                                _ = v109
                                                            }
                                                            v110 := GenS13{F0: -273}
                                                            v3 += v3
                                                            _ = v110
                                                        }
                                                        _ = v32
                                                    }
                                                case 1:
                                                    v3 += 1000
                                                }
                                                if false {
                                                    v111 := GenS14{F0: v3}
                                                    v112 := GenS10{F0: 2, F1: v3, F2: ""}
                                                    v7 = v7
                                                    _ = v111
                                                    _ = v112
                                                } else {
                                                    switch ("a" + (("hello" + "hello") + "foo bar")) {
                                                    case "hello":
                                                        v7 += v7
                                                    default:
                                                        v113 := GenS10{F0: 65535, F1: v3, F2: ""}
                                                        _ = v113
                                                    }
                                                }
                                                if (!true) {
                                                    v114 := GenS15{F0: 28, F1: genFn0(476), F2: 1}
                                                    _ = v114
                                                } else {
                                                    switch (fi10 ^ (fi10 * fi9)) {
                                                    case -1:
                                                        var v115 string = ("foo bar" + "hello")
                                                        _ = v115
                                                    case -377:
                                                        var v116 uint64 = 680
                                                        _ = v116
                                                    case -167:
                                                        if false {
                                                            v3 -= 42
                                                            if (4294967295 == 577) {
                                                                v2 = append(v2, genFn0(867))
                                                                v3 += 0
                                                            } else {
                                                                fmt.Println(fi9)
                                                                var v117 uint16 = 866
                                                                delete(v4, 4294967295)
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v117
                                                            }
                                                        }
                                                    default:
                                                        var v118 float64 = float64(((float32(-2.25) * 0.0) - (0.1 * (1e10 * 3.14159))))
                                                        _ = v118
                                                    }
                                                    v119 := GenS12{F0: 0}
                                                    _ = v119
                                                }
                                                fmt.Println(v7, fi9, fi9)
                                            }
                                            v120 := GenS14{F0: -32768}
                                            _ = v120
                                        default:
                                            _ = len(v4)
                                        }
                                        var v121 float32 = 1.5
                                        v3 = 32767
                                        _ = v121
                                    } else {
                                        var v122 int8 = -128
                                        _ = len(v2)
                                        _ = v122
                                    }
                                } else {
                                    _ = len(v1)
                                    fmt.Println(v7)
                                }
                                v123 := GenS11{F0: false, F1: 189, F2: genFn0(634)}
                                _ = v123
                            }
                            v7 = 18446744073709551615
                        case "foo bar":
                            v7 = 42
                        }
                    case -9223372036854775808:
                        v7 += 4294967295
                    }
                }
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case "hello":
        if false {
            v4[65535] = 4294967295
            switch 4294967295 {
            case 127:
                v7 = v7
            case 2147483647:
                if (2 < 18446744073709551615) {
                    var v124 uint = 793
                    var v125 uint64 = uint64(v124)
                    _ = v124
                    _ = v125
                } else {
                    for fi126 := 0; fi126 < 14; fi126++ {
                        for fi127 := 0; fi127 < 14; fi127++ {
                            v7 -= v7
                            var v128 int8 = 42
                            v7 = v7
                            _ = fi127
                            _ = v128
                        }
                        fmt.Println(v7, v3, v7)
                        _ = fi126
                    }
                }
            case 351:
                _ = len(v2)
            default:
                for fi129 := 0; fi129 < 8; fi129++ {
                    if (42 < 128) {
                        fmt.Println(fi129, v3)
                        for fi130 := 0; fi130 < 8; fi130++ {
                            v7 = (v7 / ((v7) | 1))
                            var v131 string = "\n\t"
                            v3 = v3
                            _ = fi130
                            _ = v131
                        }
                        v2 = append(v2, 2)
                    }
                    var v132 string = ("foo bar" + "foo bar")
                    _ = fi129
                    _ = v132
                }
            }
            var v133 uint32 = 2
            _ = v133
        }
    default:
        switch "" {
        case "a":
            if (65535 != 668) {
                switch -450 {
                case -1:
                    v134 := GenS13{F0: 680}
                    _ = v134
                default:
                    if false {
                        for fi135 := 0; fi135 < 16; fi135++ {
                            for fi136 := 0; fi136 < 15; fi136++ {
                                v7 += (v7 ^ v7)
                                _ = fi136
                            }
                            v2 = append(v2, 65535)
                            _ = fi135
                        }
                        fmt.Println(v7, v7)
                    } else {
                        var v137 int8 = -69
                        _ = v137
                    }
                }
                v1["a"] = true
            }
        case "foo bar":
            v7 = (^(v7))
        }
    }
    var v138 bool = false
    fmt.Println(v3, v7)
    v139 := GenS16{F0: 137, F1: 226, F2: 1}
    v7 += (^(v7))
    var v140 int8 = -128
    var v141 int8 = -52
    var v142 uint16 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v138:", v138)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v142:", v142)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v138
    _ = v139
    _ = v140
    _ = v141
    _ = v142
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
