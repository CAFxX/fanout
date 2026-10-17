package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 uint; F1 int64; F2 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint64; F2 int32 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint; F1 uint }
type GenS9 struct { F0 uint; F1 int32 }
type GenS10 struct { F0 int; F1 uint }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string }
type GenS16 struct { F0 bool; F1 int; F2 int8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 bool; F1 uint16 }
type GenS19 struct { F0 uint32; F1 int16 }
type GenS20 struct { F0 uint32; F1 int32; F2 int16 }
type GenS21 struct { F0 int64; F1 int64; F2 string }
type GenS22 struct { F0 int32; F1 int64 }
type GenS23 struct { F0 uint32; F1 uint; F2 int64 }
type GenS24 struct { F0 int; F1 int64; F2 int64 }
type GenS25 struct { F0 int16; F1 int32 }
type GenS26 struct { F0 int16 }
type GenS27 struct { F0 string; F1 uint16; F2 string }
type GenS28 struct { F0 bool; F1 uint16 }
type GenS29 struct { F0 int16 }
type GenS30 struct { F0 int }
type GenS31 struct { F0 uint16 }
type GenS32 struct { F0 string; F1 uint }
type GenS33 struct { F0 uint }
type GenS34 struct { F0 int32; F1 string }
type GenS35 struct { F0 int8; F1 uint32; F2 int64 }
func genFn0() int16 {
    v7 := GenS1{F0: "\n\t", F1: -2147483648, F2: -1}
    var v8 int8 = 127
    v8 += (v8 - (v8 >> ((386) & 63)))
    if (true && true) {
        fmt.Println(v8)
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint8 = 0
    var v3 uint32 = 4294967295
    v4 := make(map[string]uint32)
    var v5 uint = 1
    var v6 int16 = -573
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint64 = 42
        switch 4294967295 {
        case 566:
            v2 -= 0
        default:
            switch "" {
            case "a":
                var v11 string = "foo bar"
                _ = v11
            default:
                v6 = (v6 - (v6 - int16(v5)))
            }
        }
        v5 += (^(v5))
        _ = fi9
        _ = v10
    }
    var v12 uint64 = 940
    fmt.Println(v6, v3, v2)
    if false {
        fmt.Println(v5, v12, v5)
    } else {
        var v13 int = -477
        var v14 float32 = (0.1 / -2.25)
        _ = v13
        _ = v14
    }
    if false {
        if (!false) {
            switch ("" + "hello") {
            case "a":
                var v15 int64 = 9223372036854775807
                _ = v15
            }
            var v16 bool = true
            if v16 {
                if v16 {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v17 := GenS2{F0: -1, F1: v12, F2: "a"}
                        _ = v17
                    }
                    v18 := v4[""]
                    delete(v4, "")
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v20 := GenS3{F0: 65535, F1: -1, F2: 0}
                        switch fi19 {
                        case -625:
                            if (!v16) {
                                v21 := GenS3{F0: (v5 & uint(v3)), F1: 1, F2: 100}
                                fmt.Println(v2)
                                switch 2147483647 {
                                case 127:
                                    for fi22 := 0; fi22 < 16; fi22++ {
                                        var v23 int16 = v6
                                        var v24 uint8 = (v2 + (v2 - 42))
                                        switch fi22 {
                                        case 222:
                                            if v16 {
                                                v25 := GenS4{F0: -9223372036854775808, F1: v3}
                                                v2 = 2
                                                _ = v25
                                            }
                                        default:
                                            if false {
                                                v23 = (v6 | genFn0())
                                            }
                                        }
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                }
                                _ = v21
                            }
                        case -2147483648:
                            v26 := GenS4{F0: fi19, F1: (v3 % (((v3 << (((v5 << (((v5 / ((v5) | 1))) & 63))) & 63))) | 1))}
                            _ = v26
                        case -636:
                            switch "hello" {
                            case "foo bar":
                                delete(v4, "a")
                            default:
                                for fi27 := 0; fi27 < 1; fi27++ {
                                    fmt.Println(v5, v3, v12)
                                    v5 = (^(v5))
                                    _ = fi27
                                }
                            }
                        default:
                            v28 := v4[("hello" + "a")]
                            _ = v28
                        }
                        v29 := GenS2{F0: -4, F1: (v12 & 1), F2: (("" + "\n\t") + ("a" + "\n\t"))}
                        _ = fi19
                        _ = v20
                        _ = v29
                    }
                    var v30 float32 = (0.1 - float32((0.0 * -2.25)))
                    for fi31 := 0; fi31 < 7; fi31++ {
                        fmt.Println(v5, v6)
                        var v32 string = ("" + "a")
                        for fi33 := 0; fi33 < 19; fi33++ {
                            if v16 {
                                if v16 {
                                    if v16 {
                                        v4[v32] = v3
                                        if (1 > (fi31 ^ (fi31 + int(v3)))) {
                                            var v34 uint16 = 42
                                            delete(v4, ((v32 + (v32 + ("" + (("foo bar" + v32) + ((v32 + v32) + "foo bar"))))) + v32))
                                            _ = v34
                                        }
                                        var v35 float64 = 0.0
                                        fmt.Println(v3, v35, v32)
                                        v16 = v16
                                        if v16 {
                                            v5 -= v5
                                            if (v6 != (-(v6))) {
                                                if v16 {
                                                    if v16 {
                                                        if v16 {
                                                            var v36 uint32 = 2147483648
                                                            _ = v36
                                                        }
                                                        _ = len(v1)
                                                        if v16 {
                                                            _ = len(v1)
                                                            var v37 int16 = (v6 << (((v5 + v5)) & 63))
                                                            fmt.Println(v16)
                                                            _ = v37
                                                        } else {
                                                            v38 := GenS3{F0: (v5 - v5), F1: -1, F2: int8(v5)}
                                                            if (3.14159 >= 0.0) {
                                                                v35 += float64((float32(0.1) - (v30 * 3.14159)))
                                                                if false {
                                                                    if (!true) {
                                                                        if v16 {
                                                                            var v39 int16 = genFn0()
                                                                            _ = v39
                                                                        }
                                                                        if ((fi33 + (fi31 | (fi31 * (fi31 - (fi33 - fi31))))) <= 32768) {
                                                                            v40 := GenS5{F0: 1000, F1: (v12 << (((^(v5))) & 63)), F2: 42}
                                                                            var v41 int16 = v6
                                                                            _ = v40
                                                                            _ = v41
                                                                        } else {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v42 int16 = (^(v6))
                                                                            v12 = v12
                                                                            _ = v42
                                                                        }
                                                                        fmt.Println(v12)
                                                                    } else {
                                                                        if true {
                                                                            if (!(100 <= -82)) {
                                                                                v43 := GenS4{F0: 512, F1: (v3 % (((^(v3))) | 1))}
                                                                                if (1e10 >= v30) {
                                                                                    if ((0.1 - (v35 * -2.25)) == (v35 * float64(v30))) {
                                                                                        var v44 float32 = (v30 * float32((v35 / (((v35 - -2.25) * v35) * 0.1))))
                                                                                        var v45 int32 = -2147483648
                                                                                        var v46 uint8 = v2
                                                                                        if v16 {
                                                                                            v47 := GenS6{F0: v6, F1: 688}
                                                                                            _ = v47
                                                                                        }
                                                                                        _ = v44
                                                                                        _ = v45
                                                                                        _ = v46
                                                                                    } else {
                                                                                        var v48 int16 = (v6 << ((v5) & 63))
                                                                                        _ = v48
                                                                                    }
                                                                                    _ = len(v1)
                                                                                }
                                                                                var v49 bool = v16
                                                                                _ = v43
                                                                                _ = v49
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v32 + ((v32 + "hello") + "hello"))
                                                                            }
                                                                            var v50 int8 = -54
                                                                            _ = v50
                                                                        }
                                                                        if (!v16) {
                                                                            v51 := GenS1{F0: ("\n\t" + v32), F1: fi33, F2: (-(v6))}
                                                                            _ = v51
                                                                        } else {
                                                                            v52 := GenS7{F0: (((((v32 + (v32 + ((((v32 + (("a" + "a") + "foo bar")) + v32) + (v32 + v32)) + v32))) + v32) + v32) + v32) + "hello")}
                                                                            v53 := GenS8{F0: v5, F1: 65535}
                                                                            var v54 uint16 = 255
                                                                            v55 := v4[v32]
                                                                            if v16 {
                                                                                v2 = 181
                                                                                fmt.Println(fi33, v16, fi33)
                                                                                var v56 float32 = v30
                                                                                _ = v56
                                                                            } else {
                                                                                v3 = v3
                                                                                v4[""] = v3
                                                                                var v57 int32 = 1
                                                                                _ = v57
                                                                            }
                                                                            var v58 bool = (0 == 182)
                                                                            var v59 uint = 65535
                                                                            v60 := GenS9{F0: (v59 & 0), F1: int32(v54)}
                                                                            v59 = 568
                                                                            if false {
                                                                                switch fi31 {
                                                                                case 0:
                                                                                    v35 -= -2.25
                                                                                default:
                                                                                    v61 := v4[("foo bar" + ((("a" + "foo bar") + "\n\t") + "a"))]
                                                                                    _ = v61
                                                                                }
                                                                                v12 = (v12 << ((v5) & 63))
                                                                                var v62 float64 = (0.1 * 0.1)
                                                                                fmt.Println(fi33)
                                                                                _ = v62
                                                                            }
                                                                            switch int(v54) {
                                                                            case 42:
                                                                                var v63 int16 = v6
                                                                                _ = v63
                                                                            case 2147483647:
                                                                                var v64 uint8 = uint8(v12)
                                                                                _ = v64
                                                                            case 32767:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            default:
                                                                                var v65 int16 = genFn0()
                                                                                _ = v65
                                                                            }
                                                                            _ = v52
                                                                            _ = v53
                                                                            _ = v54
                                                                            _ = v55
                                                                            _ = v58
                                                                            _ = v59
                                                                            _ = v60
                                                                        }
                                                                        v66 := GenS10{F0: (^(fi31)), F1: 211}
                                                                        _ = v66
                                                                    }
                                                                    var v67 uint32 = (v3 / (((v3 / (((v3 - 0)) | 1))) | 1))
                                                                    _ = v67
                                                                } else {
                                                                    var v68 int32 = 1
                                                                    v12 += 851
                                                                    if len(v1) > 0 {
                                                                        v1[0] = v32
                                                                    }
                                                                    _ = v68
                                                                }
                                                                if (!(!(3.14159 <= 0.0))) {
                                                                    var v69 uint64 = v12
                                                                    _ = v69
                                                                } else {
                                                                    switch 9223372036854775807 {
                                                                    case 32768:
                                                                        v32 = v32
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                }
                                                            } else {
                                                                if v16 {
                                                                    _ = len(v1)
                                                                    var v70 int16 = -552
                                                                    _ = v70
                                                                }
                                                                v71 := GenS10{F0: fi31, F1: 377}
                                                                _ = v71
                                                            }
                                                            v72 := GenS4{F0: fi33, F1: v3}
                                                            _ = v38
                                                            _ = v72
                                                        }
                                                    } else {
                                                        v4[v32] = 0
                                                        v73 := GenS9{F0: (v5 << ((v5) & 63)), F1: 771}
                                                        switch v32 {
                                                        case "hello":
                                                            var v74 int64 = 811
                                                            _ = v74
                                                        case "":
                                                            if (v16 || v16) {
                                                                fmt.Println(v16)
                                                                fmt.Println(v5)
                                                                var v75 int32 = -810
                                                                _ = v75
                                                            }
                                                        default:
                                                            v4[v32] = v3
                                                        }
                                                        _ = v73
                                                    }
                                                    fmt.Println(v5)
                                                    var v76 int64 = -419
                                                    _ = v76
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v4)
                                                    v30 -= (v30 / v30)
                                                    var v77 uint64 = v12
                                                    var v78 int = (fi33 * (fi33 | 740))
                                                    v79 := GenS8{F0: (v5 | uint(v3)), F1: v5}
                                                    var v80 uint16 = 1
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                }
                                                v2 -= v2
                                            }
                                            switch fi33 {
                                            case 127:
                                                fmt.Println(v6)
                                            case 128:
                                                v81 := v4[v32]
                                                _ = v81
                                            }
                                        }
                                        _ = v35
                                    } else {
                                        if false {
                                            if ((2 != 100) && false) {
                                                switch "foo bar" {
                                                case "":
                                                    switch v32 {
                                                    case "a":
                                                        v30 -= (v30 / (v30 + v30))
                                                    case "":
                                                        switch (^(fi31)) {
                                                        case 9223372036854775807:
                                                            v82 := GenS11{F0: (fi33 + -2147483648)}
                                                            _ = v82
                                                        case 1:
                                                            switch "" {
                                                            case "":
                                                                var v83 uint64 = v12
                                                                _ = v83
                                                            case "\n\t":
                                                                var v84 int16 = v6
                                                                _ = v84
                                                            case "foo bar":
                                                                var v85 int64 = 42
                                                                _ = v85
                                                            }
                                                        case -879:
                                                            fmt.Println(fi31, fi31)
                                                        }
                                                    }
                                                case "hello":
                                                    if v16 {
                                                        var v86 bool = (genFn0() <= genFn0())
                                                        v16 = v86
                                                        _ = v86
                                                    }
                                                case "foo bar":
                                                    v6 -= (-(v6))
                                                default:
                                                    fmt.Println(v3, v16)
                                                }
                                                switch (fi31 & (fi33 >> (((^(v5))) & 63))) {
                                                case 9223372036854775807:
                                                    var v87 uint8 = (v2 - v2)
                                                    _ = v87
                                                case 1:
                                                    v88 := GenS12{F0: fi31, F1: (v12 ^ v12)}
                                                    _ = v88
                                                case 429:
                                                    var v89 int64 = -1
                                                    if v16 {
                                                        _ = len(v1)
                                                        fmt.Println(v30)
                                                    }
                                                    _ = v89
                                                default:
                                                    if true {
                                                        switch ("hello" + (((v32 + v32) + v32) + ("" + ("a" + v32)))) {
                                                        case "hello":
                                                            v2 -= v2
                                                        case "a":
                                                            var v90 uint64 = (^(v12))
                                                            _ = v90
                                                        default:
                                                            v6 += genFn0()
                                                        }
                                                        var v91 int8 = 42
                                                        var v92 int16 = -1
                                                        _ = v91
                                                        _ = v92
                                                    }
                                                }
                                                switch "a" {
                                                case "":
                                                    delete(v4, v32)
                                                default:
                                                    if v16 {
                                                        var v93 int64 = 9223372036854775807
                                                        var v94 float32 = v30
                                                        if (!true) {
                                                            if false {
                                                                var v95 float32 = v30
                                                                v94 += ((v95 + v95) / ((v30 / v30) * ((1e10 + 1e10) - v30)))
                                                                _ = v95
                                                            } else {
                                                                if false {
                                                                    if v16 {
                                                                        v96 := GenS13{F0: v3, F1: "\n\t"}
                                                                        switch -1 {
                                                                        case 967:
                                                                            _ = len(v4)
                                                                        case 0:
                                                                            fmt.Println(v2, v12, v30)
                                                                        case -2147483648:
                                                                            v97 := GenS14{F0: 128}
                                                                            _ = v97
                                                                        default:
                                                                            if (!(555 > 42)) {
                                                                                if false {
                                                                                    if true {
                                                                                        v3 = (v3 | v3)
                                                                                    } else {
                                                                                        var v98 bool = v16
                                                                                        _ = v98
                                                                                    }
                                                                                    var v99 uint32 = 2
                                                                                    _ = v99
                                                                                }
                                                                                delete(v4, v32)
                                                                            } else {
                                                                                if false {
                                                                                    var v100 int8 = -97
                                                                                    if (v94 != -2.25) {
                                                                                        v101 := GenS8{F0: 1, F1: v5}
                                                                                        switch fi33 {
                                                                                        case 598:
                                                                                            v12 += v12
                                                                                        }
                                                                                        if false {
                                                                                            if (!(!v16)) {
                                                                                                switch v32 {
                                                                                                case "hello":
                                                                                                    var v102 int8 = (v100 & (v100 + (v100 ^ 42)))
                                                                                                    _ = v102
                                                                                                case "\n\t":
                                                                                                    var v103 uint = v5
                                                                                                    _ = v103
                                                                                                case "":
                                                                                                    v104 := GenS6{F0: (^(v6)), F1: -787}
                                                                                                    _ = v104
                                                                                                default:
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v12, v12, fi33)
                                                                                                    case "":
                                                                                                        v105 := GenS15{F0: ""}
                                                                                                        v16 = v16
                                                                                                        v32 = v32
                                                                                                        v1 = append(v1, (((v32 + v32) + "\n\t") + (((v32 + v32) + v32) + "hello")))
                                                                                                        _ = v105
                                                                                                    case "a":
                                                                                                        if true {
                                                                                                            v106 := GenS11{F0: (fi33 ^ 32768)}
                                                                                                            _ = v106
                                                                                                        }
                                                                                                    default:
                                                                                                        if v16 {
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                            switch (^(fi31)) {
                                                                                                            case 4294967295:
                                                                                                                if v16 {
                                                                                                                    var v107 int = (^(fi31))
                                                                                                                    v30 = float32(1e10)
                                                                                                                    _ = v107
                                                                                                                } else {
                                                                                                                    if v16 {
                                                                                                                        _ = len(v4)
                                                                                                                        if (v32 < "\n\t") {
                                                                                                                            v108 := GenS12{F0: (fi33 % (((fi31 + fi31)) | 1)), F1: v12}
                                                                                                                            _ = v108
                                                                                                                        } else {
                                                                                                                            fmt.Println(v32, v12)
                                                                                                                            var v109 uint = v5
                                                                                                                            v32 = "\n\t"
                                                                                                                            var v110 uint8 = v2
                                                                                                                            if v16 {
                                                                                                                                _ = len(v4)
                                                                                                                                switch fi33 {
                                                                                                                                case -731:
                                                                                                                                    v111 := GenS3{F0: v5, F1: (v93 / ((v93) | 1)), F2: (v100 << ((2) & 63))}
                                                                                                                                    v2 = v2
                                                                                                                                    switch (fi33 / ((1) | 1)) {
                                                                                                                                    case 32768:
                                                                                                                                        switch (fi33 * int(v3)) {
                                                                                                                                        case 32767:
                                                                                                                                            v94 = (1e10 * 0.1)
                                                                                                                                        default:
                                                                                                                                            v100 -= (v100 * (^(v100)))
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v112 float32 = v94
                                                                                                                                        _ = v112
                                                                                                                                    }
                                                                                                                                    if (true || (0 != 678)) {
                                                                                                                                        if (v94 > v94) {
                                                                                                                                            var v113 int32 = int32(v12)
                                                                                                                                            var v114 int32 = -971
                                                                                                                                            _ = v113
                                                                                                                                            _ = v114
                                                                                                                                        } else {
                                                                                                                                            var v115 uint = v109
                                                                                                                                            var v116 int = (fi31 - fi31)
                                                                                                                                            v115 -= uint(v12)
                                                                                                                                            _ = v115
                                                                                                                                            _ = v116
                                                                                                                                        }
                                                                                                                                        v117 := GenS3{F0: 42, F1: 42, F2: (v100 >> (((v5 | v109)) & 63))}
                                                                                                                                        switch 1 {
                                                                                                                                        case 1:
                                                                                                                                            if len(v1) >= 2 {
                                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                                            }
                                                                                                                                        case 508:
                                                                                                                                            v118 := GenS16{F0: v16, F1: fi31, F2: (v100 / ((v100) | 1))}
                                                                                                                                            _ = v118
                                                                                                                                        default:
                                                                                                                                            var v119 int64 = -412
                                                                                                                                            _ = v119
                                                                                                                                        }
                                                                                                                                        if v16 {
                                                                                                                                            if (v110 == v110) {
                                                                                                                                                var v120 string = v32
                                                                                                                                                if (((fi31 + fi31) < (fi33 % ((fi33) | 1))) || v16) {
                                                                                                                                                    v121 := GenS17{F0: (v110 & (v110 + v2))}
                                                                                                                                                    _ = v121
                                                                                                                                                }
                                                                                                                                                _ = v120
                                                                                                                                            }
                                                                                                                                            if (!((-(v100)) >= v100)) {
                                                                                                                                                switch 9223372036854775807 {
                                                                                                                                                case -2147483648:
                                                                                                                                                    switch (fi31 << (((^(v109))) & 63)) {
                                                                                                                                                    case 953:
                                                                                                                                                        v12 -= (v12 >> (((v109 ^ v5)) & 63))
                                                                                                                                                    case 2147483647:
                                                                                                                                                        switch v32 {
                                                                                                                                                        case "hello":
                                                                                                                                                            if false {
                                                                                                                                                                switch (fi33 | fi33) {
                                                                                                                                                                case -2147483648:
                                                                                                                                                                    fmt.Println(fi33, v94)
                                                                                                                                                                case 32767:
                                                                                                                                                                    var v122 string = (v32 + "a")
                                                                                                                                                                    _ = v122
                                                                                                                                                                }
                                                                                                                                                                if (genFn0() <= int16(v3)) {
                                                                                                                                                                    var v123 uint16 = 25
                                                                                                                                                                    fmt.Println(v109, v93, v93)
                                                                                                                                                                    v5 = 4294967295
                                                                                                                                                                    _ = v123
                                                                                                                                                                } else {
                                                                                                                                                                    v32 = (("foo bar" + "foo bar") + (v32 + ((("hello" + v32) + v32) + ("hello" + v32))))
                                                                                                                                                                    v32 = ("a" + v32)
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                if (true && ((-(v6)) > v6)) {
                                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                                    }
                                                                                                                                                                    v124 := GenS18{F0: (((v110 / (((v2 ^ uint8(v3))) | 1)) < (v110 & v2)) || v16), F1: 0}
                                                                                                                                                                    _ = v124
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        default:
                                                                                                                                                            fmt.Println(v93, v100, v6)
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            var v125 int64 = v93
                                                                                                                                            switch (fi33 / ((128) | 1)) {
                                                                                                                                            case 768:
                                                                                                                                                if true {
                                                                                                                                                    var v126 string = ("hello" + (v32 + ((((v32 + "\n\t") + (v32 + (((v32 + v32) + "") + ""))) + ("hello" + ((v32 + ((("" + v32) + (("\n\t" + (v32 + v32)) + "\n\t")) + "")) + (((("hello" + (("a" + "") + v32)) + "") + "a") + "a")))) + ("foo bar" + (v32 + (v32 + ("" + v32)))))))
                                                                                                                                                    _ = v126
                                                                                                                                                }
                                                                                                                                            case -2147483648:
                                                                                                                                                v125 += -1
                                                                                                                                            case 157:
                                                                                                                                                var v127 int = (^(fi31))
                                                                                                                                                _ = v127
                                                                                                                                            default:
                                                                                                                                                if true {
                                                                                                                                                    if (v5 >= v109) {
                                                                                                                                                        v128 := GenS19{F0: v3, F1: (v6 % ((1) | 1))}
                                                                                                                                                        if v16 {
                                                                                                                                                            switch fi31 {
                                                                                                                                                            case 42:
                                                                                                                                                                v129 := GenS14{F0: (fi33 << (((v109 >> ((756) & 63))) & 63))}
                                                                                                                                                                _ = v129
                                                                                                                                                            case -93:
                                                                                                                                                                v30 = v30
                                                                                                                                                            default:
                                                                                                                                                                var v130 int64 = v93
                                                                                                                                                                _ = v130
                                                                                                                                                            }
                                                                                                                                                            v100 = v100
                                                                                                                                                            if ("" <= ((v32 + v32) + v32)) {
                                                                                                                                                                v12 -= (v12 - 2)
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v128
                                                                                                                                                    } else {
                                                                                                                                                        fmt.Println(v125, v6)
                                                                                                                                                        fmt.Println(v100, v32, v30)
                                                                                                                                                        v110 -= 169
                                                                                                                                                    }
                                                                                                                                                    _ = len(v1)
                                                                                                                                                } else {
                                                                                                                                                    fmt.Println(v16, v12)
                                                                                                                                                    var v131 float64 = 0.1
                                                                                                                                                    _ = v131
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            if v16 {
                                                                                                                                                v12 = (^(v12))
                                                                                                                                                v32 = ("a" + ("hello" + ((v32 + (("foo bar" + v32) + (("\n\t" + v32) + "\n\t"))) + "a")))
                                                                                                                                                var v132 int32 = 42
                                                                                                                                                _ = v132
                                                                                                                                            } else {
                                                                                                                                                var v133 uint = (v109 * v109)
                                                                                                                                                _ = v133
                                                                                                                                            }
                                                                                                                                            _ = v125
                                                                                                                                        }
                                                                                                                                        if (47 > 42) {
                                                                                                                                            switch (fi31 ^ 127) {
                                                                                                                                            case 191:
                                                                                                                                                if (!v16) {
                                                                                                                                                    fmt.Println(v12)
                                                                                                                                                } else {
                                                                                                                                                    v100 = (v100 % ((32) | 1))
                                                                                                                                                }
                                                                                                                                            case 128:
                                                                                                                                                if len(v1) >= 2 {
                                                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                v30 += 0.1
                                                                                                                                            }
                                                                                                                                            v6 = -32768
                                                                                                                                        }
                                                                                                                                        _ = v117
                                                                                                                                    }
                                                                                                                                    _ = v111
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                switch (("foo bar" + "\n\t") + v32) {
                                                                                                                                case "a":
                                                                                                                                    if true {
                                                                                                                                        v134 := GenS20{F0: (^(v3)), F1: 526, F2: (^(v6))}
                                                                                                                                        _ = v134
                                                                                                                                    } else {
                                                                                                                                        var v135 int64 = (v93 + v93)
                                                                                                                                        fmt.Println(v12, v100)
                                                                                                                                        var v136 int = fi31
                                                                                                                                        _ = v135
                                                                                                                                        _ = v136
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v137 uint64 = (v12 ^ (v12 + v12))
                                                                                                                                v138 := GenS2{F0: -386, F1: v12, F2: "\n\t"}
                                                                                                                                _ = v137
                                                                                                                                _ = v138
                                                                                                                            }
                                                                                                                            _ = v109
                                                                                                                            _ = v110
                                                                                                                        }
                                                                                                                        v16 = v16
                                                                                                                        var v139 bool = v16
                                                                                                                        _ = v139
                                                                                                                    } else {
                                                                                                                        if ((1.5 - 0.1) != (0.0 + 0.1)) {
                                                                                                                            var v140 float64 = 1.5
                                                                                                                            var v141 string = v32
                                                                                                                            v94 = v94
                                                                                                                            _ = v140
                                                                                                                            _ = v141
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            var v142 int8 = v100
                                                                                                                            v100 = -71
                                                                                                                            if (v93 >= v93) {
                                                                                                                                if (-511 <= -2147483648) {
                                                                                                                                    v143 := GenS9{F0: (v5 + v5), F1: 42}
                                                                                                                                    var v144 uint64 = v12
                                                                                                                                    _ = len(v4)
                                                                                                                                    _ = v143
                                                                                                                                    _ = v144
                                                                                                                                } else {
                                                                                                                                    if v16 {
                                                                                                                                        v145 := GenS21{F0: v93, F1: v93, F2: "hello"}
                                                                                                                                        _ = v145
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v6 -= int16(v93)
                                                                                                                            }
                                                                                                                            switch ("" + v32) {
                                                                                                                            case "":
                                                                                                                                var v146 float64 = ((float64(v30) - ((float64(((((((v30 + v94) + 1.5) / v30) / v30) - 1e10) * float32(0.1))) / (1e10 - (1e10 * (0.1 / -2.25)))) - float64(1.5))) + (-2.25 * ((1e10 * -2.25) * (-2.25 * 0.1))))
                                                                                                                                _ = v146
                                                                                                                            }
                                                                                                                            _ = v142
                                                                                                                        } else {
                                                                                                                            var v147 uint64 = 18446744073709551615
                                                                                                                            var v148 uint = uint(v2)
                                                                                                                            if (!(1.5 == ((v30 - 1.5) * -2.25))) {
                                                                                                                                if v16 {
                                                                                                                                    v149 := GenS22{F0: -579, F1: int64(v6)}
                                                                                                                                    _ = v149
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    switch fi33 {
                                                                                                                                    case 751:
                                                                                                                                        if (!v16) {
                                                                                                                                            fmt.Println(v94, v3, v147)
                                                                                                                                            v150 := GenS13{F0: (v3 * (v3 & 42)), F1: (((v32 + "hello") + v32) + ((v32 + "") + v32))}
                                                                                                                                            _ = v150
                                                                                                                                        }
                                                                                                                                    case 42:
                                                                                                                                        var v151 uint64 = v12
                                                                                                                                        _ = v151
                                                                                                                                    case 152:
                                                                                                                                        v3 += 899
                                                                                                                                    default:
                                                                                                                                        var v152 int16 = -707
                                                                                                                                        _ = v152
                                                                                                                                    }
                                                                                                                                    var v153 float32 = v94
                                                                                                                                    v16 = ((2 > (v5 / ((v5) | 1))) || v16)
                                                                                                                                    if (32768 == (^(fi31))) {
                                                                                                                                        fmt.Println(v147, v3, v94)
                                                                                                                                        if false {
                                                                                                                                            v154 := GenS9{F0: 1, F1: -622}
                                                                                                                                            var v155 int8 = v100
                                                                                                                                            _ = v154
                                                                                                                                            _ = v155
                                                                                                                                        }
                                                                                                                                        var v156 int = 128
                                                                                                                                        _ = v156
                                                                                                                                    }
                                                                                                                                    if v16 {
                                                                                                                                        v157 := GenS5{F0: (v6 | (-(v6))), F1: (v12 - 2), F2: 1}
                                                                                                                                        if (v93 <= v93) {
                                                                                                                                            v6 += -1
                                                                                                                                            v6 -= genFn0()
                                                                                                                                            v5 = 42
                                                                                                                                        } else {
                                                                                                                                            v2 -= (v2 / ((uint8(v147)) | 1))
                                                                                                                                            v158 := GenS20{F0: (v3 - (v3 << (((v148 % ((v5) | 1))) & 63))), F1: -762, F2: v6}
                                                                                                                                            _ = v158
                                                                                                                                        }
                                                                                                                                        v93 += (^(v93))
                                                                                                                                        v153 = ((3.14159 + (v153 / 0.1)) + (v94 * 3.14159))
                                                                                                                                        switch v32 {
                                                                                                                                        case "\n\t":
                                                                                                                                            fmt.Println(v12)
                                                                                                                                        case "a":
                                                                                                                                            if v16 {
                                                                                                                                                v159 := GenS23{F0: v3, F1: (^(v5)), F2: 42}
                                                                                                                                                _ = v159
                                                                                                                                            }
                                                                                                                                        case "":
                                                                                                                                            switch v32 {
                                                                                                                                            case "hello":
                                                                                                                                                v6 = (v6 % ((int16(v147)) | 1))
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        _ = v157
                                                                                                                                    } else {
                                                                                                                                        if v16 {
                                                                                                                                            var v160 uint16 = 255
                                                                                                                                            if (v100 <= (v100 - (v100 / ((-58) | 1)))) {
                                                                                                                                                var v161 uint16 = (v160 % (((v160 + 783)) | 1))
                                                                                                                                                _ = v161
                                                                                                                                            }
                                                                                                                                            v162 := GenS24{F0: int(v93), F1: (v93 * (v93 & (v93 | v93))), F2: 9223372036854775807}
                                                                                                                                            _ = v160
                                                                                                                                            _ = v162
                                                                                                                                        }
                                                                                                                                        v4[v32] = v3
                                                                                                                                        if (!(v100 <= (-(v100)))) {
                                                                                                                                            if true {
                                                                                                                                                if (((float32(((1e10 * (3.14159 * -2.25)) - 0.0)) - (v94 + float32((1.5 / (1e10 / 0.1))))) - 0.0) >= v153) {
                                                                                                                                                    if (!(0.0 > (0.0 * float64(0.1)))) {
                                                                                                                                                        fmt.Println(v100, v93, v6)
                                                                                                                                                        fmt.Println(v148, v94, v148)
                                                                                                                                                        if (42 <= 2147483647) {
                                                                                                                                                            if true {
                                                                                                                                                                if (v2 >= (v2 + (v2 | (v2 + v2)))) {
                                                                                                                                                                    v163 := GenS4{F0: (fi33 * fi31), F1: v3}
                                                                                                                                                                    var v164 uint8 = (v2 * (v2 / ((42) | 1)))
                                                                                                                                                                    var v165 int64 = v93
                                                                                                                                                                    _ = v163
                                                                                                                                                                    _ = v164
                                                                                                                                                                    _ = v165
                                                                                                                                                                } else {
                                                                                                                                                                    var v166 uint64 = (v12 / (((^(v147))) | 1))
                                                                                                                                                                    v167 := GenS25{F0: v6, F1: 2147483647}
                                                                                                                                                                    if false {
                                                                                                                                                                        var v168 int64 = (v93 % (((v93 >> ((255) & 63))) | 1))
                                                                                                                                                                        _ = v168
                                                                                                                                                                    } else {
                                                                                                                                                                        switch fi31 {
                                                                                                                                                                        case 32767:
                                                                                                                                                                            if (32768 >= 128) {
                                                                                                                                                                                delete(v4, "hello")
                                                                                                                                                                                var v169 int8 = (v100 % (((-(v100))) | 1))
                                                                                                                                                                                _ = v169
                                                                                                                                                                            }
                                                                                                                                                                        case -979:
                                                                                                                                                                            v94 = v94
                                                                                                                                                                        case 9223372036854775807:
                                                                                                                                                                            var v170 uint8 = uint8(v3)
                                                                                                                                                                            _ = v170
                                                                                                                                                                        default:
                                                                                                                                                                            var v171 uint = 18446744073709551615
                                                                                                                                                                            _ = v171
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    v172 := GenS16{F0: v16, F1: fi31, F2: int8(v6)}
                                                                                                                                                                    _ = v166
                                                                                                                                                                    _ = v167
                                                                                                                                                                    _ = v172
                                                                                                                                                                }
                                                                                                                                                                v173 := GenS26{F0: v6}
                                                                                                                                                                _ = len(v1)
                                                                                                                                                                v153 -= 1e10
                                                                                                                                                                _ = v173
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    if v16 {
                                                                                                                                                        var v174 int32 = 2147483647
                                                                                                                                                        if len(v1) >= 2 {
                                                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                                                        }
                                                                                                                                                        var v175 int8 = (v100 & (v100 >> (((^(v148))) & 63)))
                                                                                                                                                        _ = v174
                                                                                                                                                        _ = v175
                                                                                                                                                    } else {
                                                                                                                                                        var v176 float32 = (v94 - v94)
                                                                                                                                                        v177 := GenS27{F0: v32, F1: 109, F2: v32}
                                                                                                                                                        if v16 {
                                                                                                                                                            switch 128 {
                                                                                                                                                            case 0:
                                                                                                                                                                var v178 uint = v5
                                                                                                                                                                _ = v178
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v176
                                                                                                                                                        _ = v177
                                                                                                                                                    }
                                                                                                                                                    switch fi31 {
                                                                                                                                                    case -9223372036854775808:
                                                                                                                                                        v153 += 1.5
                                                                                                                                                    default:
                                                                                                                                                        var v179 int = fi31
                                                                                                                                                        _ = v179
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                var v180 float32 = (0.1 + 0.0)
                                                                                                                                                switch 42 {
                                                                                                                                                case 42:
                                                                                                                                                    switch (fi31 >> (((v5 % ((981) | 1))) & 63)) {
                                                                                                                                                    case -772:
                                                                                                                                                        var v181 uint16 = 606
                                                                                                                                                        _ = v181
                                                                                                                                                    default:
                                                                                                                                                        var v182 float32 = ((((((0.0 * (v94 / ((3.14159 + v180) / (3.14159 + v94)))) / v30) + v94) * float32(-2.25)) + ((v30 - float32(float64(v180))) - 1e10)) / -2.25)
                                                                                                                                                        _ = v182
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    var v183 uint32 = (v3 | (v3 / ((2147483648) | 1)))
                                                                                                                                                    _ = v183
                                                                                                                                                }
                                                                                                                                                if (v3 < (v3 << ((v148) & 63))) {
                                                                                                                                                    if (uint16(v12) <= 661) {
                                                                                                                                                        v184 := GenS28{F0: (v2 == v2), F1: 2}
                                                                                                                                                        _ = len(v4)
                                                                                                                                                        _ = v184
                                                                                                                                                    }
                                                                                                                                                    v100 = -128
                                                                                                                                                    var v185 int8 = v100
                                                                                                                                                    var v186 uint32 = (v3 ^ 553)
                                                                                                                                                    _ = v185
                                                                                                                                                    _ = v186
                                                                                                                                                } else {
                                                                                                                                                    var v187 int64 = int64(v100)
                                                                                                                                                    v94 += v180
                                                                                                                                                    switch ((v32 + v32) + v32) {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        if ((((v148 ^ v148) < v148) && (v16 && (0 < 0))) && (!(true || true))) {
                                                                                                                                                            fmt.Println(v30, v12, v2)
                                                                                                                                                            v188 := GenS20{F0: v3, F1: 0, F2: (v6 * 1000)}
                                                                                                                                                            fmt.Println(v3, v100)
                                                                                                                                                            v189 := GenS1{F0: (("\n\t" + "hello") + v32), F1: 32767, F2: v6}
                                                                                                                                                            _ = v188
                                                                                                                                                            _ = v189
                                                                                                                                                        }
                                                                                                                                                    case "\n\t":
                                                                                                                                                        var v190 uint = (v148 + (v5 / (((v5 % ((0) | 1))) | 1)))
                                                                                                                                                        _ = v190
                                                                                                                                                    case "a":
                                                                                                                                                        var v191 int64 = v187
                                                                                                                                                        _ = v191
                                                                                                                                                    default:
                                                                                                                                                        fmt.Println(v6, v32, v6)
                                                                                                                                                    }
                                                                                                                                                    _ = v187
                                                                                                                                                }
                                                                                                                                                _ = v180
                                                                                                                                            }
                                                                                                                                            if len(v1) >= 2 {
                                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        fmt.Println(v100, v153, v3)
                                                                                                                                        _ = len(v1)
                                                                                                                                        var v192 uint16 = 65535
                                                                                                                                        if len(v1) > 0 {
                                                                                                                                            v1[0] = v32
                                                                                                                                        }
                                                                                                                                        _ = v192
                                                                                                                                    }
                                                                                                                                    switch 32768 {
                                                                                                                                    case 9223372036854775807:
                                                                                                                                        var v193 uint8 = v2
                                                                                                                                        _ = v193
                                                                                                                                    case 128:
                                                                                                                                        if (true || v16) {
                                                                                                                                            delete(v4, "\n\t")
                                                                                                                                            v153 = -2.25
                                                                                                                                            switch (^(fi33)) {
                                                                                                                                            case 42:
                                                                                                                                                var v194 uint8 = (v2 * 42)
                                                                                                                                                _ = v194
                                                                                                                                            case -604:
                                                                                                                                                switch v32 {
                                                                                                                                                case "a":
                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                    }
                                                                                                                                                case "foo bar":
                                                                                                                                                    v32 = v32
                                                                                                                                                }
                                                                                                                                            case 443:
                                                                                                                                                var v195 float64 = 0.1
                                                                                                                                                _ = v195
                                                                                                                                            default:
                                                                                                                                                if false {
                                                                                                                                                    var v196 string = "\n\t"
                                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                                    }
                                                                                                                                                    v93 = 871
                                                                                                                                                    _ = v196
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            v1 = append(v1, v32)
                                                                                                                                            fmt.Println(fi31, v16, v3)
                                                                                                                                            var v197 float32 = v94
                                                                                                                                            _ = v197
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        if (197 < 255) {
                                                                                                                                            var v198 uint64 = v12
                                                                                                                                            _ = v198
                                                                                                                                        } else {
                                                                                                                                            var v199 bool = false
                                                                                                                                            _ = v199
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    _ = v153
                                                                                                                                } else {
                                                                                                                                    if (v93 != int64(fi31)) {
                                                                                                                                        if (18446744073709551615 >= 1) {
                                                                                                                                            var v200 string = "\n\t"
                                                                                                                                            if (1 > (v148 + v5)) {
                                                                                                                                                v30 += float32(1e10)
                                                                                                                                            }
                                                                                                                                            _ = v200
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if len(v1) >= 2 {
                                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                                    }
                                                                                                                                    if ((!v16) && ((false || true) || v16)) {
                                                                                                                                        v2 += 155
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                fmt.Println(v32, v12)
                                                                                                                            }
                                                                                                                            v201 := GenS12{F0: 4294967295, F1: uint64(v3)}
                                                                                                                            _ = v147
                                                                                                                            _ = v148
                                                                                                                            _ = v201
                                                                                                                        }
                                                                                                                        if ((v93 * (v93 * (v93 & v93))) <= (^(v93))) {
                                                                                                                            var v202 uint16 = 255
                                                                                                                            _ = v202
                                                                                                                        }
                                                                                                                    }
                                                                                                                    if ((v32 + "a") <= (v32 + v32)) {
                                                                                                                        var v203 uint32 = (^(v3))
                                                                                                                        _ = v203
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                var v204 int16 = genFn0()
                                                                                                                _ = v204
                                                                                                            }
                                                                                                            v5 -= 433
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v205 int64 = v93
                                                                                                if v16 {
                                                                                                    v4["hello"] = v3
                                                                                                } else {
                                                                                                    var v206 uint8 = v2
                                                                                                    var v207 uint64 = (^(v12))
                                                                                                    _ = v206
                                                                                                    _ = v207
                                                                                                }
                                                                                                _ = v205
                                                                                            }
                                                                                        }
                                                                                        _ = v101
                                                                                    } else {
                                                                                        var v208 uint8 = 1
                                                                                        _ = v208
                                                                                    }
                                                                                    _ = v100
                                                                                } else {
                                                                                    var v209 uint = v5
                                                                                    _ = v209
                                                                                }
                                                                                var v210 float32 = float32(1.5)
                                                                                _ = v210
                                                                            }
                                                                        }
                                                                        _ = v96
                                                                    } else {
                                                                        var v211 float64 = (3.14159 * 0.0)
                                                                        v16 = ((^(v6)) > int16(v5))
                                                                        v2 += 128
                                                                        _ = v211
                                                                    }
                                                                    switch (fi33 + fi33) {
                                                                    case -2147483648:
                                                                        v212 := GenS28{F0: (v12 > v12), F1: 717}
                                                                        _ = v212
                                                                    default:
                                                                        fmt.Println(v94, v2)
                                                                    }
                                                                }
                                                                v213 := GenS29{F0: (v6 ^ v6)}
                                                                v5 = 455
                                                                _ = v213
                                                            }
                                                            if false {
                                                                var v214 int32 = 42
                                                                _ = v214
                                                            }
                                                        } else {
                                                            var v215 int8 = -9
                                                            switch fi33 {
                                                            case 42:
                                                                var v216 int32 = -2147483648
                                                                _ = v216
                                                            case -2147483648:
                                                                v217 := GenS30{F0: 127}
                                                                _ = v217
                                                            case -9223372036854775808:
                                                                if (v215 == v215) {
                                                                    var v218 uint8 = v2
                                                                    if false {
                                                                        switch (^(fi33)) {
                                                                        case -2147483648:
                                                                            v2 += (v2 + 205)
                                                                        case 4294967295:
                                                                            var v219 uint32 = 0
                                                                            _ = v219
                                                                        case 42:
                                                                            fmt.Println(fi31, v12, v16)
                                                                        }
                                                                    } else {
                                                                        var v220 uint16 = 42
                                                                        var v221 uint64 = (v12 + (v12 | (v12 & (^(v12)))))
                                                                        v94 = (3.14159 / (3.14159 * 1e10))
                                                                        _ = v220
                                                                        _ = v221
                                                                    }
                                                                    var v222 int32 = -457
                                                                    _ = v218
                                                                    _ = v222
                                                                }
                                                            }
                                                            _ = v215
                                                        }
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v4[v32] = (v3 >> ((v5) & 63))
                                                        v223 := v4[""]
                                                        _ = v223
                                                    }
                                                }
                                            }
                                            v224 := GenS16{F0: false, F1: (fi31 * (fi33 + 631)), F2: -128}
                                            switch fi31 {
                                            case -1:
                                                var v225 uint32 = 2147483648
                                                _ = v225
                                            case 281:
                                                var v226 uint = 255
                                                if true {
                                                    var v227 uint = 346
                                                    v2 += (^(v2))
                                                    if false {
                                                        switch (fi31 + (fi31 & fi33)) {
                                                        case 127:
                                                            v3 += (^(v3))
                                                        }
                                                        v228 := GenS8{F0: v227, F1: v226}
                                                        _ = v228
                                                    }
                                                    _ = v227
                                                } else {
                                                    if false {
                                                        if false {
                                                            v229 := GenS26{F0: genFn0()}
                                                            _ = v229
                                                        }
                                                        v230 := GenS31{F0: 32767}
                                                        _ = v230
                                                    } else {
                                                        if v16 {
                                                            v12 += (v12 | v12)
                                                            if (829 <= uint16(v226)) {
                                                                v231 := GenS32{F0: ((v32 + v32) + (((v32 + (((v32 + v32) + "hello") + (v32 + v32))) + v32) + v32)), F1: (v5 ^ v5)}
                                                                var v232 uint64 = (^(v12))
                                                                var v233 string = ("a" + "hello")
                                                                _ = v231
                                                                _ = v232
                                                                _ = v233
                                                            }
                                                            if (841 < int32(v3)) {
                                                                v234 := GenS30{F0: fi33}
                                                                _ = v234
                                                            }
                                                        } else {
                                                            v30 = float32(-2.25)
                                                            v16 = ("" < "a")
                                                        }
                                                        if len(v1) > 0 {
                                                            v1[0] = v32
                                                        }
                                                    }
                                                    v1 = append(v1, ("hello" + v32))
                                                    v226 -= 2
                                                }
                                                var v235 int64 = 1
                                                _ = v226
                                                _ = v235
                                            case 755:
                                                var v236 bool = (1 != -772)
                                                _ = v236
                                            }
                                            _ = v224
                                        }
                                        v1 = append(v1, "foo bar")
                                        v12 -= (^(v12))
                                    }
                                    var v237 float64 = ((float64(((v30 + v30) - 3.14159)) + 0.0) + (0.0 - float64(v30)))
                                    switch v32 {
                                    case "":
                                        var v238 float64 = v237
                                        _ = v238
                                    case "hello":
                                        v5 = (v5 + uint(v3))
                                    }
                                    _ = v237
                                }
                                var v239 int = (fi33 + fi31)
                                _ = v239
                            }
                            _ = fi33
                        }
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
                var v240 int = -227
                _ = v240
            } else {
                for fi241 := 0; fi241 < 16; fi241++ {
                    fmt.Println(v3)
                    _ = fi241
                }
                var v242 int32 = 664
                v1 = append(v1, "foo bar")
                _ = v242
            }
            _ = v16
        } else {
            for fi243 := 0; fi243 < 9; fi243++ {
                v6 = 32767
                if false {
                    var v244 string = (("hello" + "hello") + "foo bar")
                    if (false && ((-1 != -1) && ((fi243 / ((32768) | 1)) <= (-(fi243))))) {
                        v2 = v2
                        v3 = 0
                        for fi245 := 0; fi245 < 17; fi245++ {
                            if false {
                                var v246 float64 = 1.5
                                _ = v246
                            }
                            v247 := GenS33{F0: (^(v5))}
                            var v248 int = fi243
                            _ = fi245
                            _ = v247
                            _ = v248
                        }
                    }
                    _ = v244
                }
                v6 += 42
                _ = fi243
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    v249 := GenS31{F0: 42}
    v12 = 667
    v2 -= v2
    v250 := GenS34{F0: -2147483648, F1: "hello"}
    fmt.Println(v5, v3)
    for fi251 := 0; fi251 < 12; fi251++ {
        v3 = 2147483648
        _ = fi251
    }
    if false {
        fmt.Println(v3)
        if true {
            v12 += (v12 % ((1) | 1))
            _ = len(v1)
            var v252 string = ("a" + "a")
            _ = v252
        }
        v12 = 82
    }
    v4["foo bar"] = (v3 >> ((v5) & 63))
    v253 := GenS35{F0: 1, F1: 65535, F2: 1}
    v254 := GenS1{F0: ("hello" + ""), F1: 32767, F2: v6}
    v12 -= v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v249
    _ = v250
    _ = v253
    _ = v254
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
