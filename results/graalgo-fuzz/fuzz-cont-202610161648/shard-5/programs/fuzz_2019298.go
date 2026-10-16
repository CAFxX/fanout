package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int8; F1 int64; F2 uint64 }
type GenS13 struct { F0 uint16; F1 int }
type GenS14 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS15 struct { F0 int16; F1 int32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint32; F1 bool }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int32; F1 int32; F2 string }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int8; F1 int8; F2 int16 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 uint16; F1 string }
type GenS25 struct { F0 string; F1 uint16 }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 int16; F1 bool }
type GenS29 struct { F0 uint8; F1 string; F2 uint16 }
type GenS30 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS31 struct { F0 bool; F1 uint8; F2 int8 }
type GenS32 struct { F0 int16; F1 uint64 }
type GenS33 struct { F0 uint64 }
type GenS34 struct { F0 uint64 }
type GenS35 struct { F0 uint; F1 uint64 }
type GenS36 struct { F0 int64; F1 int16 }
type GenS37 struct { F0 int16; F1 bool }
type GenS38 struct { F0 uint16; F1 uint8 }
type GenS39 struct { F0 uint16 }
type GenS40 struct { F0 uint16; F1 uint8; F2 uint8 }
func genFn0(p0 uint) uint32 {
    switch "" {
    case "foo bar":
        v7 := []int{4294967295, -964, -311, -563}
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return 2147483648
}
func genFn1(p0 int16, p1 int16, p2 uint16) uint64 {
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{}
        p2 = 1
        _ = fi8
        _ = v9
    }
    p0 = p0
    for fi10 := 0; fi10 < 18; fi10++ {
        if (0.1 < (0.0 / 0.1)) {
            if (18446744073709551615 < uint64(p2)) {
                p0 -= 625
            } else {
                p1 = p1
            }
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                fmt.Println(p0, fi11, p0)
                p0 += p0
                _ = fi11
            }
        }
        _ = fi10
    }
    var v12 bool = false
    v13 := []int{-1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 string = (("a" + "") + "foo bar")
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    var v4 uint8 = 255
    var v5 int = 0
    var v6 int8 = 100
    var v14 float32 = (((0.1 / ((-2.25 - (0.1 - 1.5)) + 0.0)) * 1e10) + 1e10)
    delete(v2, genFn0(65535))
    fmt.Println(v5, v4, v4)
    v15 := GenS1{F0: 0, F1: 67, F2: int16(v6)}
    v6 = v6
    var v16 string = v1
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            if ((739 <= 945) || true) {
                v19 := GenS1{F0: 42, F1: 1, F2: -32768}
                if false {
                    var v20 float64 = 3.14159
                    _ = v20
                }
                _ = v19
            }
            switch (v5 / (((fi18 | fi17)) | 1)) {
            case 32768:
                if (!true) {
                    var v21 uint32 = 1
                    if false {
                        v22 := []int{fi17, (v5 & fi18), v5, fi17}
                        var v23 int = (v5 | v5)
                        if true {
                            v24 := GenS2{F0: (^(v4)), F1: 1, F2: 42}
                            var v25 int8 = int8(v4)
                            switch v16 {
                            case "\n\t":
                                if true {
                                    var v26 bool = (-3 > fi18)
                                    _ = v26
                                } else {
                                    if false {
                                        if false {
                                            v16 = "\n\t"
                                        } else {
                                            v6 = v6
                                            v27 := GenS1{F0: 0, F1: v4, F2: 0}
                                            v6 = v6
                                            _ = v27
                                        }
                                        if false {
                                            switch -876 {
                                            case 32768:
                                                var v28 float32 = v14
                                                _ = v28
                                            case 710:
                                                var v29 int16 = 1000
                                                _ = v29
                                            case 4294967295:
                                                switch v16 {
                                                case "hello":
                                                    if (((v16 + v1) + (v16 + v1)) != "\n\t") {
                                                        v30 := GenS3{F0: true, F1: 376}
                                                        v4 = v4
                                                        fmt.Println(v5)
                                                        _ = v30
                                                    }
                                                default:
                                                    if true {
                                                        delete(v2, (v21 - 65535))
                                                        var v31 int32 = -1
                                                        _ = v31
                                                    }
                                                }
                                            default:
                                                v32 := GenS2{F0: v4, F1: 255, F2: -1}
                                                _ = v32
                                            }
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                        }
                                    } else {
                                        var v33 int16 = 1
                                        v34 := GenS1{F0: 695, F1: (v4 % ((v4) | 1)), F2: (v33 >> ((2) & 63))}
                                        switch (v16 + v1) {
                                        case "":
                                            v22 = append(v22, fi18)
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := GenS2{F0: v4, F1: 65535, F2: 367}
                                    var v36 float32 = float32(((0.1 - 1.5) * 1.5))
                                    v37 := GenS4{F0: -300, F1: -1, F2: ""}
                                    v38 := GenS5{F0: genFn0(179)}
                                    fmt.Println(v25, v6)
                                    if ((42 >= v4) || ((genFn1(1000, -415, 255) >= 0) || (int8(v21) != v6))) {
                                        v6 = (-(v6))
                                        v5 -= (fi18 & (v23 ^ (v23 & fi17)))
                                        v5 -= v23
                                    }
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                            case "":
                                fmt.Println(v21)
                            case "foo bar":
                                var v39 string = (v1 + (("\n\t" + "\n\t") + v16))
                                _ = v39
                            default:
                                if (0 <= 2) {
                                    v40 := GenS2{F0: 42, F1: 940, F2: 42}
                                    _ = v40
                                }
                            }
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v23)
                            if true {
                                v23 -= (fi18 << ((48) & 63))
                                if (!true) {
                                    v41 := v2[42]
                                    v21 = (v21 + (v21 ^ 4294967295))
                                    v6 = v6
                                    _ = v41
                                } else {
                                    fmt.Println(fi17, v4)
                                    var v42 uint16 = 255
                                    _ = v42
                                }
                                var v43 int8 = (v6 | v6)
                                _ = v43
                            } else {
                                if (((^(v21)) != (v21 | (v21 ^ v21))) && (2 < 700)) {
                                    v44 := GenS6{F0: (v6 % ((v6) | 1))}
                                    var v45 uint16 = 32767
                                    v16 = v16
                                    v46 := GenS7{F0: 42}
                                    fmt.Println(v16, v16, v1)
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v14 = -2.25
                                    var v47 uint16 = 87
                                    _ = v47
                                }
                                if true {
                                    v16 = ((v16 + "hello") + v16)
                                    switch fi18 {
                                    case 968:
                                        v48 := GenS7{F0: -978}
                                        _ = v48
                                    case -9223372036854775808:
                                        var v49 uint16 = 0
                                        _ = v49
                                    case -258:
                                        fmt.Println(fi18, v14)
                                    default:
                                        v4 -= 128
                                    }
                                    var v50 uint64 = uint64(v4)
                                    _ = v50
                                }
                                fmt.Println(v14)
                            }
                            if true {
                                v16 = v16
                                v22 = append(v22, (^(fi17)))
                            }
                        }
                        _ = len(v2)
                        _ = v22
                        _ = v23
                    } else {
                        var v51 uint = 65535
                        _ = v51
                    }
                    _ = v21
                }
            case -1:
                fmt.Println(v1, v5, v5)
            case -9223372036854775808:
                if true {
                    if ((!((!(!true)) && false)) && true) {
                        if (1.5 >= (float64(-2.25) + 0.1)) {
                            v5 = fi18
                        }
                        if false {
                            v52 := []int{fi18}
                            v4 = v4
                            var v53 bool = true
                            if v53 {
                                var v54 uint = 2
                                if (!(v53 && v53)) {
                                    if v53 {
                                        _ = len(v2)
                                        v55 := v2[0]
                                        switch -9223372036854775808 {
                                        case 4294967295:
                                            var v56 uint32 = genFn0(1)
                                            fmt.Println(v4, v6, v55)
                                            if (687 != 320) {
                                                v57 := GenS8{F0: 0, F1: 287}
                                                _ = v57
                                            } else {
                                                v16 = v16
                                                fmt.Println(v5, v4, v5)
                                            }
                                            _ = v56
                                        default:
                                            if (42 < 9223372036854775807) {
                                                v1 = "foo bar"
                                                v14 = v14
                                                v52 = append(v52, int(v6))
                                            } else {
                                                var v58 uint64 = 803
                                                _ = v58
                                            }
                                        }
                                        v5 = fi18
                                        _ = v55
                                    } else {
                                        if false {
                                            if (979 != 881) {
                                                var v59 int = fi18
                                                v60 := GenS9{F0: v59}
                                                _ = v59
                                                _ = v60
                                            }
                                            v6 = (v6 ^ v6)
                                            switch (v1 + "") {
                                            case "\n\t":
                                                v6 -= v6
                                            default:
                                                var v61 int32 = -1
                                                _ = v61
                                            }
                                        } else {
                                            v6 = v6
                                            v1 = v1
                                            switch "a" {
                                            case "":
                                                v62 := GenS4{F0: 42, F1: -664, F2: "foo bar"}
                                                _ = v62
                                            case "hello":
                                                v63 := GenS2{F0: (v4 >> ((v54) & 63)), F1: 0, F2: -845}
                                                if len(v52) > 0 {
                                                    v52[0] = fi17
                                                }
                                                _ = v63
                                            case "foo bar":
                                                var v64 string = ((((v16 + ((v1 + v16) + "a")) + v16) + (v1 + "\n\t")) + "a")
                                                _ = v64
                                            default:
                                                var v65 string = v16
                                                _ = v65
                                            }
                                            if (!((true && false) && ((v14 - (1.5 + v14)) == (v14 + v14)))) {
                                                var v66 float32 = 1.5
                                                v67 := GenS10{F0: uint16(v54)}
                                                var v68 int8 = 42
                                                switch (fi17 - (fi18 & (v5 >> ((uint(v4)) & 63)))) {
                                                case -2147483648:
                                                    v4 += 0
                                                case 0:
                                                    _ = len(v2)
                                                default:
                                                    switch v16 {
                                                    case "a":
                                                        if v53 {
                                                            v69 := GenS7{F0: 42}
                                                            v16 = v1
                                                            var v70 int64 = 835
                                                            _ = v69
                                                            _ = v70
                                                        } else {
                                                            switch -9223372036854775808 {
                                                            case 127:
                                                                v4 += (v4 & v4)
                                                            default:
                                                                switch "" {
                                                                case "\n\t":
                                                                    v66 += ((v66 * ((v14 * v66) * (3.14159 + -2.25))) * v66)
                                                                case "hello":
                                                                    v71 := GenS11{F0: genFn1(0, 1, 0)}
                                                                    _ = v71
                                                                case "foo bar":
                                                                    if ((-1 == 0) && (genFn0(v54) == 2147483648)) {
                                                                        v72 := GenS12{F0: v68, F1: -296, F2: 42}
                                                                        v73 := GenS13{F0: 449, F1: (v5 + fi18)}
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        v52 = append(v52, (fi17 % ((fi18) | 1)))
                                                                    }
                                                                default:
                                                                    if v53 {
                                                                        v6 = (v68 << (((^(v54))) & 63))
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        v68 = v68
                                                    case "\n\t":
                                                        if (genFn0(743) >= 65535) {
                                                            fmt.Println(fi17)
                                                            fmt.Println(v68, v1)
                                                        }
                                                        if (v54 <= v54) {
                                                            v54 = 18446744073709551615
                                                            v74 := v2[0]
                                                            var v75 uint64 = 18446744073709551615
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                                switch fi18 {
                                                case 0:
                                                    v76 := GenS2{F0: v4, F1: 42, F2: 42}
                                                    _ = v76
                                                case -612:
                                                    v77 := GenS14{F0: 4294967295, F1: v6, F2: 1}
                                                    _ = v77
                                                }
                                                v78 := GenS13{F0: 2, F1: (fi18 * v5)}
                                                v4 += 255
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                                _ = v78
                                            }
                                            v6 -= v6
                                        }
                                        if v53 {
                                            if false {
                                                v54 += 621
                                                if false {
                                                    if (4294967295 == (v54 >> ((v54) & 63))) {
                                                        v79 := GenS15{F0: -791, F1: -2147483648}
                                                        v4 = v4
                                                        var v80 uint16 = 2
                                                        if v53 {
                                                            switch v5 {
                                                            case 202:
                                                                var v81 uint64 = 991
                                                                _ = v81
                                                            default:
                                                                v1 = "\n\t"
                                                            }
                                                            fmt.Println(fi17, v16)
                                                        }
                                                        if v53 {
                                                            v4 = 0
                                                            var v82 float64 = float64(v14)
                                                            var v83 int16 = -32768
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        var v84 uint32 = 0
                                                        var v85 int64 = -9223372036854775808
                                                        if v53 {
                                                            v86 := GenS16{F0: v6}
                                                            if (false || v53) {
                                                                v14 = v14
                                                                v53 = false
                                                                if v53 {
                                                                    if v53 {
                                                                        v1 = v16
                                                                        if ((!(v53 || v53)) && (((false && v53) || v53) || (true || (false && v53)))) {
                                                                            if (0 != v85) {
                                                                                var v87 string = ("a" + (v1 + ((v16 + v16) + "hello")))
                                                                                if (!false) {
                                                                                    v88 := GenS17{F0: genFn0(v54), F1: (!(((3.14159 + ((0.0 * ((-2.25 / 1e10) * 0.1)) / float64((0.0 / ((float32((float64(((1.5 * v14) * float32(((((1.5 - 1.5) + 1.5) / -2.25) * float64(v14))))) - 0.0)) / 1.5) / ((((1.5 - ((float32((((float64(-2.25) + (float64((3.14159 - 0.0)) * ((float64(1e10) + -2.25) / 1.5))) + 0.1) * 3.14159)) / (v14 + (3.14159 - v14))) * -2.25)) * float32((3.14159 * float64(v14)))) / (1e10 - (v14 / (v14 * v14)))) * v14)))))) + 0.0) <= -2.25))}
                                                                                    var v89 int16 = -427
                                                                                    if (!v53) {
                                                                                        var v90 float32 = 1.5
                                                                                        if (true && (v89 < v89)) {
                                                                                            v14 += v90
                                                                                        }
                                                                                        _ = v90
                                                                                    }
                                                                                    var v91 uint8 = 42
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                    _ = v91
                                                                                } else {
                                                                                    var v92 bool = v53
                                                                                    _ = v92
                                                                                }
                                                                                var v93 uint64 = genFn1(-32768, 32767, 255)
                                                                                _ = v87
                                                                                _ = v93
                                                                            }
                                                                            var v94 int32 = 569
                                                                            v52 = append(v52, (fi18 % ((fi17) | 1)))
                                                                            _ = v94
                                                                        } else {
                                                                            if (true && v53) {
                                                                                if v53 {
                                                                                    if len(v52) >= 2 {
                                                                                        v52 = v52[1:len(v52)]
                                                                                    }
                                                                                    v52 = append(v52, (fi18 << ((v54) & 63)))
                                                                                    v84 += v84
                                                                                    if v53 {
                                                                                        var v95 string = "hello"
                                                                                        _ = len(v2)
                                                                                        v95 = v16
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                            }
                                                                            switch v16 {
                                                                            case "\n\t":
                                                                                fmt.Println(v84, fi17, v54)
                                                                            case "hello":
                                                                                if (!((v53 && false) || v53)) {
                                                                                    v14 += 0.0
                                                                                } else {
                                                                                    fmt.Println(v14, v54)
                                                                                    v52 = append(v52, (fi18 ^ (fi17 + (-(fi18)))))
                                                                                }
                                                                            default:
                                                                                if (v85 >= v85) {
                                                                                    fmt.Println(v5)
                                                                                    switch (v16 + "\n\t") {
                                                                                    case "hello":
                                                                                        v54 = (v54 / ((42) | 1))
                                                                                    case "\n\t":
                                                                                        v96 := GenS8{F0: (v85 ^ (v85 - v85)), F1: 1}
                                                                                        _ = v96
                                                                                    case "a":
                                                                                        v6 += -1
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        if (!(false || ((fi17 - v5) < 1))) {
                                                                            v97 := GenS18{F0: v53}
                                                                            fmt.Println(v53, v85)
                                                                            var v98 float64 = 1.5
                                                                            _ = v97
                                                                            _ = v98
                                                                        }
                                                                        var v99 bool = (-740 != (fi17 + fi17))
                                                                        if v99 {
                                                                            var v100 int32 = -396
                                                                            fmt.Println(v16, v6, v5)
                                                                            v4 = v4
                                                                            var v101 int32 = 0
                                                                            if len(v52) >= 2 {
                                                                                v52 = v52[1:len(v52)]
                                                                            }
                                                                            _ = v100
                                                                            _ = v101
                                                                        }
                                                                        _ = v99
                                                                    } else {
                                                                        if (2147483647 >= fi17) {
                                                                            var v102 bool = true
                                                                            var v103 int8 = 0
                                                                            _ = v102
                                                                            _ = v103
                                                                        } else {
                                                                            _ = len(v3)
                                                                        }
                                                                        v16 = ("foo bar" + v16)
                                                                    }
                                                                    if (!(!false)) {
                                                                        v4 = (v4 | 78)
                                                                        if (!v53) {
                                                                            switch v5 {
                                                                            case -146:
                                                                                switch fi18 {
                                                                                case 229:
                                                                                    v14 = 1.5
                                                                                case -9223372036854775808:
                                                                                    v104 := GenS17{F0: v84, F1: true}
                                                                                    var v105 uint16 = uint16(v84)
                                                                                    _ = v104
                                                                                    _ = v105
                                                                                case -1:
                                                                                    var v106 uint64 = 2
                                                                                    _ = v106
                                                                                default:
                                                                                    var v107 uint16 = 255
                                                                                    _ = v107
                                                                                }
                                                                            case 860:
                                                                                v1 = v16
                                                                            case -9223372036854775808:
                                                                                v5 -= (-(v5))
                                                                            }
                                                                            v108 := GenS19{F0: -2147483648, F1: 336, F2: (((("a" + v16) + v16) + v1) + v1)}
                                                                            if true {
                                                                                var v109 uint64 = 42
                                                                                _ = v109
                                                                            }
                                                                            _ = v108
                                                                        } else {
                                                                            if v53 {
                                                                                var v110 uint32 = v84
                                                                                _ = v110
                                                                            }
                                                                            if len(v52) > 0 {
                                                                                v52[0] = (v5 * fi17)
                                                                            }
                                                                        }
                                                                    }
                                                                    switch v16 {
                                                                    case "hello":
                                                                        v111 := GenS18{F0: false}
                                                                        _ = v111
                                                                    case "":
                                                                        var v112 uint32 = (v84 >> ((v54) & 63))
                                                                        _ = v112
                                                                    case "a":
                                                                        v85 = v85
                                                                    default:
                                                                        v1 = ("foo bar" + "\n\t")
                                                                    }
                                                                    v16 = (("\n\t" + "") + "\n\t")
                                                                    var v113 bool = false
                                                                    switch (v16 + "hello") {
                                                                    case "foo bar":
                                                                        v54 += 213
                                                                    default:
                                                                        v114 := GenS5{F0: (^(v84))}
                                                                        _ = v114
                                                                    }
                                                                    _ = v113
                                                                } else {
                                                                    if ((3.14159 >= v14) && v53) {
                                                                        v4 += (v4 >> (((v54 * (v54 % ((530) | 1)))) & 63))
                                                                        var v115 string = v16
                                                                        _ = v115
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                            } else {
                                                                var v116 int64 = 42
                                                                fmt.Println(v14)
                                                                switch -9223372036854775808 {
                                                                case -2147483648:
                                                                    v117 := GenS14{F0: 512, F1: v6, F2: 1}
                                                                    _ = v117
                                                                case 260:
                                                                    v118 := GenS7{F0: -1}
                                                                    _ = v118
                                                                case 4294967295:
                                                                    delete(v2, 4294967295)
                                                                default:
                                                                    if v53 {
                                                                        switch (-(fi17)) {
                                                                        case 212:
                                                                            if true {
                                                                                v14 = 1e10
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v119 := GenS20{F0: 729}
                                                                                    _ = v119
                                                                                case "\n\t":
                                                                                    switch fi18 {
                                                                                    case 4294967295:
                                                                                        v120 := GenS21{F0: (v6 - v6), F1: int8(v54), F2: 42}
                                                                                        _ = v120
                                                                                    default:
                                                                                        var v121 uint16 = 65535
                                                                                        _ = v121
                                                                                    }
                                                                                case "a":
                                                                                    v84 += v84
                                                                                    if len(v52) >= 2 {
                                                                                        v52 = v52[1:len(v52)]
                                                                                    }
                                                                                    switch v5 {
                                                                                    case -2147483648:
                                                                                        fmt.Println(v54, v16, v1)
                                                                                    case 4294967295:
                                                                                        var v122 uint64 = genFn1(0, 32767, 32767)
                                                                                        _ = v122
                                                                                    case 0:
                                                                                        var v123 string = v1
                                                                                        _ = v123
                                                                                    }
                                                                                default:
                                                                                    v16 = "\n\t"
                                                                                }
                                                                                fmt.Println(v16, v6, v54)
                                                                            }
                                                                        default:
                                                                            if (v4 >= (v4 << ((v54) & 63))) {
                                                                                v124 := GenS12{F0: -16, F1: v85, F2: 36}
                                                                                _ = v124
                                                                            } else {
                                                                                if v53 {
                                                                                    v85 += -9223372036854775808
                                                                                } else {
                                                                                    fmt.Println(v5)
                                                                                }
                                                                                var v125 float64 = (-2.25 - 1e10)
                                                                                _ = v125
                                                                            }
                                                                        }
                                                                        _ = len(v52)
                                                                        if false {
                                                                            fmt.Println(v1, v14)
                                                                            v85 = (v85 | 661)
                                                                        }
                                                                    }
                                                                }
                                                                _ = v116
                                                            }
                                                            v1 = ""
                                                            _ = v86
                                                        }
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                    if (2147483647 < -261) {
                                                        v126 := GenS22{F0: 770}
                                                        _ = v126
                                                    }
                                                    if v53 {
                                                        switch (((v16 + (("" + v1) + "")) + "\n\t") + "a") {
                                                        case "\n\t":
                                                            v4 = v4
                                                        case "hello":
                                                            fmt.Println(v4, v53, fi18)
                                                        case "foo bar":
                                                            if v53 {
                                                                if (3.14159 <= ((((1e10 * 3.14159) * (-2.25 * 0.0)) / ((((-2.25 / 1.0) * ((-2.25 - float64(v14)) / 3.14159)) - 1e10) - float64((v14 * v14)))) - 3.14159)) {
                                                                    v53 = (-2.25 != ((((1e10 / 1.0) + -2.25) / (3.14159 * 1.5)) / (3.14159 - (0.1 / (1.5 / 1e10)))))
                                                                    if (1 < 0) {
                                                                        if (uint32(v54) <= 4294967295) {
                                                                            var v127 int8 = v6
                                                                            var v128 string = v16
                                                                            _ = v127
                                                                            _ = v128
                                                                        }
                                                                        v129 := GenS23{F0: 1}
                                                                        v130 := GenS3{F0: (835 != 65535), F1: 837}
                                                                        v6 -= 102
                                                                        _ = v129
                                                                        _ = v130
                                                                    } else {
                                                                        var v131 int32 = -295
                                                                        v4 -= 109
                                                                        v132 := GenS9{F0: -1}
                                                                        _ = v131
                                                                        _ = v132
                                                                    }
                                                                    _ = len(v3)
                                                                }
                                                                var v133 int = fi17
                                                                if false {
                                                                    v14 = ((0.0 / -2.25) / (1.5 * ((1e10 * v14) + ((v14 * v14) * (float32((3.14159 - (-2.25 / float64(v14)))) + v14)))))
                                                                    var v134 int64 = 42
                                                                    if (true || (v54 <= v54)) {
                                                                        v135 := GenS19{F0: int32(v54), F1: 758, F2: ""}
                                                                        _ = v135
                                                                    } else {
                                                                        delete(v3, v133)
                                                                        var v136 uint8 = (v4 % (((v4 | (v4 << ((0) & 63)))) | 1))
                                                                        delete(v3, (v5 >> ((1) & 63)))
                                                                        _ = v136
                                                                    }
                                                                    _ = v134
                                                                } else {
                                                                    fmt.Println(fi17)
                                                                }
                                                                _ = v133
                                                            } else {
                                                                var v137 uint = (^(v54))
                                                                var v138 string = v1
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                        }
                                                        v139 := v2[genFn0(v54)]
                                                        _ = len(v2)
                                                        _ = v139
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if len(v52) >= 2 {
                                                                v52 = v52[1:len(v52)]
                                                            }
                                                        default:
                                                            v140 := GenS24{F0: 26, F1: "\n\t"}
                                                            _ = v140
                                                        }
                                                        if true {
                                                            fmt.Println(v1, v4)
                                                        } else {
                                                            switch v1 {
                                                            case "a":
                                                                if ((^(fi17)) >= fi17) {
                                                                    v141 := v2[65535]
                                                                    _ = v141
                                                                } else {
                                                                    if false {
                                                                        if v53 {
                                                                            if ((1.5 * ((1e10 * 1.5) + -2.25)) >= 1.5) {
                                                                                if true {
                                                                                    var v142 float32 = 1.5
                                                                                    if (1 == int32(v4)) {
                                                                                        var v143 int32 = -2147483648
                                                                                        _ = len(v52)
                                                                                        var v144 bool = v53
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    }
                                                                                    _ = v142
                                                                                } else {
                                                                                    v3[(v5 % (((fi18 | fi17)) | 1))] = (((!(int64(v4) > 0)) || v53) && true)
                                                                                    v5 -= fi18
                                                                                    var v145 int16 = 1
                                                                                    _ = v145
                                                                                }
                                                                                v54 = v54
                                                                                if (v53 || v53) {
                                                                                    var v146 int = (v5 - 42)
                                                                                    _ = len(v52)
                                                                                    var v147 float64 = 3.14159
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                } else {
                                                                                    v6 += (v6 / ((int8(v4)) | 1))
                                                                                    if (523 < 4294967295) {
                                                                                        if v53 {
                                                                                            v4 += 127
                                                                                            if (float32(0.0) > -2.25) {
                                                                                                v148 := GenS25{F0: v1, F1: 65535}
                                                                                                switch (fi18 & v5) {
                                                                                                case -955:
                                                                                                    v1 = ""
                                                                                                }
                                                                                                _ = v148
                                                                                            } else {
                                                                                                v3[v5] = ((!(3.14159 >= 0.1)) || (int(v54) <= (^(fi18))))
                                                                                            }
                                                                                            if (!(32767 == 1)) {
                                                                                                v1 = v1
                                                                                                v5 += (v5 << (((v54 / (((^(v54))) | 1))) & 63))
                                                                                                switch (v1 + v1) {
                                                                                                case "":
                                                                                                    v3[fi17] = (!v53)
                                                                                                case "\n\t":
                                                                                                    if (v53 || ((v54 | 4294967295) == (v54 * v54))) {
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            v54 -= (v54 & (v54 | uint(v4)))
                                                                                                        case "a":
                                                                                                            var v149 string = (v1 + (v16 + v1))
                                                                                                            _ = v149
                                                                                                        default:
                                                                                                            var v150 uint32 = 187
                                                                                                            _ = v150
                                                                                                        }
                                                                                                        var v151 uint16 = 2
                                                                                                        _ = v151
                                                                                                    }
                                                                                                default:
                                                                                                    if v53 {
                                                                                                        if (2 > (v4 % ((v4) | 1))) {
                                                                                                            switch 1 {
                                                                                                            case 4294967295:
                                                                                                                v152 := GenS26{F0: 42}
                                                                                                                _ = v152
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v52 = append(v52, (v5 % (((v5 / ((1) | 1))) | 1)))
                                                                                                        switch "\n\t" {
                                                                                                        case "foo bar":
                                                                                                            fmt.Println(v14, v5)
                                                                                                        case "hello":
                                                                                                            var v153 uint64 = 551
                                                                                                            _ = v153
                                                                                                        case "":
                                                                                                            v154 := GenS2{F0: v4, F1: 2, F2: -212}
                                                                                                            _ = v154
                                                                                                        default:
                                                                                                            if (((0.0 * 0.1) != (((3.14159 * float64(0.1)) - (float64((0.1 / 3.14159)) - 0.1)) / (-2.25 * (1e10 * 0.1)))) || v53) {
                                                                                                                if v53 {
                                                                                                                    var v155 uint = v54
                                                                                                                    _ = len(v52)
                                                                                                                    var v156 uint32 = 1
                                                                                                                    _ = v155
                                                                                                                    _ = v156
                                                                                                                }
                                                                                                            } else {
                                                                                                                v1 = "\n\t"
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(fi18, v1)
                                                                                                }
                                                                                                var v157 uint64 = 4294967295
                                                                                                _ = v157
                                                                                            } else {
                                                                                                fmt.Println(v1, fi18)
                                                                                                v158 := v2[42]
                                                                                                _ = v158
                                                                                            }
                                                                                            if (!(((!(!(v4 == v4))) || v53) || v53)) {
                                                                                                var v159 int32 = 0
                                                                                                v1 = (v1 + v16)
                                                                                                v52 = append(v52, (fi17 * (-(v5))))
                                                                                                _ = v159
                                                                                            }
                                                                                            v160 := GenS12{F0: -124, F1: 496, F2: 42}
                                                                                            var v161 uint64 = 4294967295
                                                                                            _ = v160
                                                                                            _ = v161
                                                                                        } else {
                                                                                            v6 -= v6
                                                                                            fmt.Println(v5, v4)
                                                                                            fmt.Println(v14, v5, v6)
                                                                                        }
                                                                                        if true {
                                                                                            var v162 uint8 = (^(v4))
                                                                                            v14 = float32(((((1e10 / float64(v14)) / -2.25) / 3.14159) * float64(1.5)))
                                                                                            switch "" {
                                                                                            case "":
                                                                                                v163 := GenS1{F0: 32767, F1: v4, F2: 5}
                                                                                                _ = v163
                                                                                            case "hello":
                                                                                                var v164 uint8 = (v162 ^ (v4 + uint8(v54)))
                                                                                                _ = v164
                                                                                            default:
                                                                                                v165 := GenS22{F0: 610}
                                                                                                _ = v165
                                                                                            }
                                                                                            _ = v162
                                                                                        } else {
                                                                                            if v53 {
                                                                                                var v166 string = v16
                                                                                                switch (fi17 >> (((v54 - (v54 * v54))) & 63)) {
                                                                                                case 127:
                                                                                                    v167 := GenS18{F0: v53}
                                                                                                    switch 0 {
                                                                                                    case 127:
                                                                                                        var v168 uint = v54
                                                                                                        _ = v168
                                                                                                    default:
                                                                                                        v52 = append(v52, (fi18 * (v5 - fi18)))
                                                                                                    }
                                                                                                    var v169 uint32 = 65535
                                                                                                    var v170 string = (v166 + "a")
                                                                                                    v5 = fi18
                                                                                                    _ = v167
                                                                                                    _ = v169
                                                                                                    _ = v170
                                                                                                case 981:
                                                                                                    if v53 {
                                                                                                        fmt.Println(fi18)
                                                                                                        v6 = (^(v6))
                                                                                                    } else {
                                                                                                        fmt.Println(v53, v14, v14)
                                                                                                        v171 := GenS27{F0: 1}
                                                                                                        fmt.Println(v6, v5, fi17)
                                                                                                        _ = v171
                                                                                                    }
                                                                                                default:
                                                                                                    if len(v52) >= 2 {
                                                                                                        v52 = v52[1:len(v52)]
                                                                                                    }
                                                                                                }
                                                                                                _ = len(v52)
                                                                                                _ = v166
                                                                                            } else {
                                                                                                switch 32768 {
                                                                                                case 4294967295:
                                                                                                    switch v1 {
                                                                                                    case "hello":
                                                                                                        switch v5 {
                                                                                                        case -360:
                                                                                                            if len(v52) > 0 {
                                                                                                                v52[0] = 9223372036854775807
                                                                                                            }
                                                                                                        case -1:
                                                                                                            v16 = v1
                                                                                                        default:
                                                                                                            v54 -= (v54 ^ 264)
                                                                                                        }
                                                                                                    case "":
                                                                                                        v172 := v3[int(v6)]
                                                                                                        _ = v172
                                                                                                    case "a":
                                                                                                        fmt.Println(v16, fi17, fi17)
                                                                                                    }
                                                                                                case 127:
                                                                                                    fmt.Println(v6, v54, fi18)
                                                                                                default:
                                                                                                    var v173 bool = (1 < 32767)
                                                                                                    _ = v173
                                                                                                }
                                                                                                v174 := v3[(v5 + (^(fi18)))]
                                                                                                if (0 == 41) {
                                                                                                    var v175 uint8 = (v4 + 2)
                                                                                                    if (483 == 708) {
                                                                                                        var v176 int8 = -1
                                                                                                        var v177 uint64 = 1
                                                                                                        var v178 uint8 = 1
                                                                                                        _ = v176
                                                                                                        _ = v177
                                                                                                        _ = v178
                                                                                                    } else {
                                                                                                        v16 = v16
                                                                                                        var v179 uint16 = 255
                                                                                                        _ = v179
                                                                                                    }
                                                                                                    _ = v175
                                                                                                } else {
                                                                                                    var v180 string = v1
                                                                                                    v180 = v180
                                                                                                    _ = v180
                                                                                                }
                                                                                                _ = v174
                                                                                            }
                                                                                            switch (fi18 & fi18) {
                                                                                            case 1:
                                                                                                if len(v52) >= 2 {
                                                                                                    v52 = v52[1:len(v52)]
                                                                                                }
                                                                                            case 362:
                                                                                                if (-32768 != -502) {
                                                                                                    var v181 uint32 = 42
                                                                                                    v14 += ((float32((1e10 / 3.14159)) + float32(float64(1.5))) - (float32((1.5 / 1.0)) / v14))
                                                                                                    v4 -= 2
                                                                                                    switch ("foo bar" + "a") {
                                                                                                    case "a":
                                                                                                        fmt.Println(v14)
                                                                                                        if (42 == -617) {
                                                                                                            var v182 uint16 = 0
                                                                                                            var v183 int64 = -356
                                                                                                            _ = v182
                                                                                                            _ = v183
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        if (!((!v53) && (uint16(v54) != 255))) {
                                                                                                            var v184 uint16 = 1
                                                                                                            _ = v184
                                                                                                        } else {
                                                                                                            if ((3.14159 - -2.25) <= -2.25) {
                                                                                                                switch 1 {
                                                                                                                case -2147483648:
                                                                                                                    fmt.Println(v6, v1, v6)
                                                                                                                case -142:
                                                                                                                    v185 := GenS28{F0: 502, F1: false}
                                                                                                                    _ = v185
                                                                                                                }
                                                                                                                v6 += (v6 >> ((v54) & 63))
                                                                                                                v186 := GenS2{F0: v4, F1: 255, F2: 131}
                                                                                                                _ = v186
                                                                                                            }
                                                                                                            var v187 uint = (v54 ^ (^(v54)))
                                                                                                            _ = v187
                                                                                                        }
                                                                                                    default:
                                                                                                        v188 := GenS25{F0: "a", F1: 42}
                                                                                                        _ = v188
                                                                                                    }
                                                                                                    v14 = 3.14159
                                                                                                    _ = v181
                                                                                                } else {
                                                                                                    fmt.Println(v14)
                                                                                                    var v189 int16 = -813
                                                                                                    v53 = ((^(v6)) >= v6)
                                                                                                    v190 := GenS29{F0: v4, F1: ("\n\t" + v1), F2: 42}
                                                                                                    _ = v189
                                                                                                    _ = v190
                                                                                                }
                                                                                            default:
                                                                                                switch (fi17 >> (((v54 / ((v54) | 1))) & 63)) {
                                                                                                case -1:
                                                                                                    var v191 int32 = 1
                                                                                                    _ = v191
                                                                                                case -176:
                                                                                                    var v192 int8 = 100
                                                                                                    _ = v192
                                                                                                case 1:
                                                                                                    if v53 {
                                                                                                        v193 := GenS22{F0: 18446744073709551615}
                                                                                                        if v53 {
                                                                                                            v194 := GenS15{F0: 1000, F1: 2147483647}
                                                                                                            if (float64((v14 / 1.5)) == -2.25) {
                                                                                                                var v195 uint8 = 255
                                                                                                                fmt.Println(fi17, fi17, v4)
                                                                                                                _ = v195
                                                                                                            }
                                                                                                            _ = v194
                                                                                                        } else {
                                                                                                            v14 += v14
                                                                                                            v16 = (v1 + "hello")
                                                                                                        }
                                                                                                        _ = v193
                                                                                                    } else {
                                                                                                        switch (v1 + v16) {
                                                                                                        case "":
                                                                                                            var v196 bool = (!(!(!true)))
                                                                                                            _ = v196
                                                                                                        case "\n\t":
                                                                                                            v197 := GenS2{F0: v4, F1: 42, F2: -2147483648}
                                                                                                            _ = v197
                                                                                                        case "a":
                                                                                                            v198 := GenS4{F0: -9223372036854775808, F1: 1, F2: (v16 + (v1 + v1))}
                                                                                                            _ = v198
                                                                                                        default:
                                                                                                            if true {
                                                                                                                v5 += -1
                                                                                                                v54 = (^(v54))
                                                                                                                v4 = v4
                                                                                                            } else {
                                                                                                                if (-729 < 67) {
                                                                                                                    switch -1 {
                                                                                                                    case 42:
                                                                                                                        fmt.Println(v16, v1)
                                                                                                                    case -564:
                                                                                                                        fmt.Println(v1)
                                                                                                                    default:
                                                                                                                        var v199 uint16 = 620
                                                                                                                        _ = v199
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v200 := GenS30{F0: -1, F1: 18446744073709551615, F2: 177}
                                                                                                                    var v201 int8 = -128
                                                                                                                    v14 = 3.14159
                                                                                                                    _ = v200
                                                                                                                    _ = v201
                                                                                                                }
                                                                                                                if false {
                                                                                                                    if v53 {
                                                                                                                        var v202 float64 = 0.0
                                                                                                                        _ = v202
                                                                                                                    }
                                                                                                                    switch (^(v5)) {
                                                                                                                    case 9223372036854775807:
                                                                                                                        v3[fi17] = (uint8(v54) == v4)
                                                                                                                    case 4294967295:
                                                                                                                        v203 := GenS30{F0: 1000, F1: genFn1(-992, -32768, 32767), F2: 32767}
                                                                                                                        _ = v203
                                                                                                                    default:
                                                                                                                        _ = len(v2)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v204 int32 = 0
                                                                                                                    switch ((v16 + (((v1 + v16) + ("hello" + ((v16 + v16) + "a"))) + v16)) + "\n\t") {
                                                                                                                    case "":
                                                                                                                        v5 += (fi18 >> ((v54) & 63))
                                                                                                                    case "foo bar":
                                                                                                                        if ((-(v6)) < v6) {
                                                                                                                            switch (v5 * -1) {
                                                                                                                            case -2147483648:
                                                                                                                                if v53 {
                                                                                                                                    if v53 {
                                                                                                                                        var v205 string = v16
                                                                                                                                        _ = v205
                                                                                                                                    }
                                                                                                                                    var v206 uint64 = genFn1(269, 1000, 1)
                                                                                                                                    _ = v206
                                                                                                                                }
                                                                                                                            case -9223372036854775808:
                                                                                                                                switch (v1 + (v1 + v1)) {
                                                                                                                                case "foo bar":
                                                                                                                                    if (129 > (v4 >> ((v54) & 63))) {
                                                                                                                                        var v207 string = (v1 + "\n\t")
                                                                                                                                        switch v1 {
                                                                                                                                        case "\n\t":
                                                                                                                                            if (("foo bar" + v1) != v1) {
                                                                                                                                                v4 += (v4 << (((v54 + v54)) & 63))
                                                                                                                                                fmt.Println(fi18)
                                                                                                                                            } else {
                                                                                                                                                v53 = ((v204 <= v204) || (!(65535 <= 254)))
                                                                                                                                                var v208 uint8 = (v4 % ((0) | 1))
                                                                                                                                                var v209 string = v1
                                                                                                                                                _ = v208
                                                                                                                                                _ = v209
                                                                                                                                            }
                                                                                                                                            switch 2147483647 {
                                                                                                                                            case 32768:
                                                                                                                                                var v210 int64 = -116
                                                                                                                                                _ = v210
                                                                                                                                            default:
                                                                                                                                                var v211 string = v1
                                                                                                                                                _ = v211
                                                                                                                                            }
                                                                                                                                            v14 = (v14 * (0.1 * v14))
                                                                                                                                        case "a":
                                                                                                                                            var v212 int = fi17
                                                                                                                                            _ = v212
                                                                                                                                        case "foo bar":
                                                                                                                                            var v213 string = ""
                                                                                                                                            _ = v213
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v16, v4, v16)
                                                                                                                                        }
                                                                                                                                        _ = v207
                                                                                                                                    }
                                                                                                                                case "\n\t":
                                                                                                                                    v53 = ((v204 + v204) < v204)
                                                                                                                                case "hello":
                                                                                                                                    v214 := GenS31{F0: (1.5 == -2.25), F1: (^(v4)), F2: -128}
                                                                                                                                    _ = v214
                                                                                                                                }
                                                                                                                            case 190:
                                                                                                                                var v215 int16 = -156
                                                                                                                                _ = v215
                                                                                                                            }
                                                                                                                            _ = len(v3)
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v4)
                                                                                                                    default:
                                                                                                                        var v216 bool = ((v204 / ((int32(v54)) | 1)) != v204)
                                                                                                                        _ = v216
                                                                                                                    }
                                                                                                                    _ = v204
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        v217 := v2[42]
                                                                                                        _ = v217
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v218 uint16 = 869
                                                                                            _ = v218
                                                                                        }
                                                                                        var v219 int8 = (^(v6))
                                                                                        _ = v219
                                                                                    }
                                                                                }
                                                                                v220 := GenS32{F0: -817, F1: 372}
                                                                                delete(v3, (fi18 << (((v54 | (v54 << ((4294967295) & 63)))) & 63)))
                                                                                _ = v220
                                                                            }
                                                                            _ = len(v52)
                                                                        }
                                                                        v1 = v16
                                                                        v54 = 0
                                                                        _ = len(v3)
                                                                        fmt.Println(v14, v4, fi17)
                                                                        var v221 int = (v5 % ((int(v54)) | 1))
                                                                        _ = v221
                                                                    } else {
                                                                        v222 := GenS32{F0: 0, F1: 2}
                                                                        v223 := GenS7{F0: 700}
                                                                        _ = v222
                                                                        _ = v223
                                                                    }
                                                                    fmt.Println(fi17, fi18, fi18)
                                                                    fmt.Println(v4, fi17, v14)
                                                                }
                                                            case "hello":
                                                                switch v1 {
                                                                case "a":
                                                                    switch -1 {
                                                                    case -1:
                                                                        if v53 {
                                                                            var v224 uint64 = 574
                                                                            var v225 int64 = 681
                                                                            if ((v16 + ((("a" + v1) + "hello") + v1)) <= "a") {
                                                                                switch v16 {
                                                                                case "":
                                                                                    if (v1 != v1) {
                                                                                        var v226 uint64 = (v224 / ((v224) | 1))
                                                                                        var v227 int = (fi17 % ((fi18) | 1))
                                                                                        v228 := v2[788]
                                                                                        _ = v226
                                                                                        _ = v227
                                                                                        _ = v228
                                                                                    }
                                                                                case "a":
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        if v53 {
                                                                                            v229 := GenS33{F0: genFn1(-483, -32768, 47)}
                                                                                            v54 += (v54 >> (((v54 + v54)) & 63))
                                                                                            _ = len(v52)
                                                                                            _ = v229
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                            var v230 int = (-(fi17))
                                                                                            fmt.Println(v225)
                                                                                            _ = v230
                                                                                        }
                                                                                    default:
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            switch (-(v5)) {
                                                                                            case -9223372036854775808:
                                                                                                var v231 bool = true
                                                                                                _ = v231
                                                                                            }
                                                                                        case "hello":
                                                                                            var v232 string = (v16 + "hello")
                                                                                            _ = v232
                                                                                        case "foo bar":
                                                                                            v233 := v2[65535]
                                                                                            _ = v233
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    v16 = v1
                                                                                default:
                                                                                    var v234 int32 = 42
                                                                                    _ = v234
                                                                                }
                                                                                fmt.Println(v5, v4, v16)
                                                                            } else {
                                                                                var v235 int = fi17
                                                                                v236 := GenS13{F0: 42, F1: fi17}
                                                                                v224 += (v224 << ((588) & 63))
                                                                                v1 = v1
                                                                                if ((!(v4 >= 129)) || v53) {
                                                                                    v237 := GenS4{F0: int64(v6), F1: -379, F2: ""}
                                                                                    fmt.Println(v5, v53, v53)
                                                                                    if (!(-2.25 < 0.1)) {
                                                                                        fmt.Println(v235, v4)
                                                                                    } else {
                                                                                        var v238 uint16 = 145
                                                                                        var v239 uint64 = 82
                                                                                        switch v16 {
                                                                                        case "":
                                                                                            fmt.Println(fi17, v53)
                                                                                        case "hello":
                                                                                            if ((^(v238)) != v238) {
                                                                                                delete(v3, 128)
                                                                                            }
                                                                                        }
                                                                                        _ = v238
                                                                                        _ = v239
                                                                                    }
                                                                                    v235 += 32767
                                                                                    switch (-(v5)) {
                                                                                    case 2147483647:
                                                                                        var v240 int64 = (v225 / ((v225) | 1))
                                                                                        _ = v240
                                                                                    case 0:
                                                                                        var v241 bool = false
                                                                                        _ = v241
                                                                                    case 557:
                                                                                        v2[genFn0(v54)] = (!((42 >= -831) || (v225 < 0)))
                                                                                    default:
                                                                                        var v242 float64 = 1e10
                                                                                        _ = v242
                                                                                    }
                                                                                    v1 = v16
                                                                                    var v243 int16 = -829
                                                                                    _ = v237
                                                                                    _ = v243
                                                                                } else {
                                                                                    var v244 int8 = v6
                                                                                    _ = v244
                                                                                }
                                                                                if (!(v53 || v53)) {
                                                                                    var v245 int = (v235 ^ 4294967295)
                                                                                    if (v4 >= v4) {
                                                                                        fmt.Println(v224, v6)
                                                                                        fmt.Println(v5, v16, v1)
                                                                                        v235 = (fi18 >> (((v54 - v54)) & 63))
                                                                                    }
                                                                                    _ = v245
                                                                                } else {
                                                                                    v1 = ("a" + ((v16 + (v1 + ((v16 + ("hello" + v1)) + "a"))) + v1))
                                                                                    v246 := GenS34{F0: v224}
                                                                                    v6 = 100
                                                                                    _ = v246
                                                                                }
                                                                                if false {
                                                                                    switch (v235 << ((v54) & 63)) {
                                                                                    case 1:
                                                                                        var v247 uint8 = v4
                                                                                        _ = v247
                                                                                    case 768:
                                                                                        v224 -= (^(v224))
                                                                                        var v248 string = (("foo bar" + "a") + "\n\t")
                                                                                        var v249 int = -9223372036854775808
                                                                                        _ = v248
                                                                                        _ = v249
                                                                                    default:
                                                                                        var v250 float64 = (0.0 + (1e10 + 0.1))
                                                                                        _ = v250
                                                                                    }
                                                                                    v251 := GenS17{F0: 940, F1: (!("hello" != ""))}
                                                                                    v53 = ((3.14159 - 1.5) < 3.14159)
                                                                                    _ = v251
                                                                                } else {
                                                                                    v252 := GenS35{F0: (v54 / ((v54) | 1)), F1: (v224 * genFn1(-1, 15, 65535))}
                                                                                    v253 := GenS36{F0: (v225 % ((v225) | 1)), F1: 717}
                                                                                    var v254 int16 = 1
                                                                                    _ = v252
                                                                                    _ = v253
                                                                                    _ = v254
                                                                                }
                                                                                _ = v235
                                                                                _ = v236
                                                                            }
                                                                            _ = v224
                                                                            _ = v225
                                                                        }
                                                                    default:
                                                                        v255 := GenS37{F0: 1000, F1: v53}
                                                                        _ = v255
                                                                    }
                                                                case "\n\t":
                                                                    var v256 uint = v54
                                                                    _ = v256
                                                                case "hello":
                                                                    switch -345 {
                                                                    case 440:
                                                                        v16 = "a"
                                                                    case -9223372036854775808:
                                                                        v54 -= (v54 / ((931) | 1))
                                                                    case 9223372036854775807:
                                                                        switch v5 {
                                                                        case -586:
                                                                            if v53 {
                                                                                fmt.Println(v4)
                                                                            }
                                                                        }
                                                                    default:
                                                                        fmt.Println(fi17, v14, fi18)
                                                                    }
                                                                }
                                                            default:
                                                                _ = len(v52)
                                                            }
                                                        }
                                                        v54 = (v54 & (^(v54)))
                                                    }
                                                }
                                                if (!(!v53)) {
                                                    v257 := GenS11{F0: 2}
                                                    fmt.Println(fi17, v16)
                                                    switch v5 {
                                                    case 0:
                                                        v14 += (0.0 * v14)
                                                    }
                                                    _ = v257
                                                }
                                            } else {
                                                var v258 bool = (!v53)
                                                if ((fi18 + fi17) < fi17) {
                                                    fmt.Println(fi18)
                                                    fmt.Println(v54)
                                                    var v259 float64 = ((-2.25 - 0.0) - 1.5)
                                                    _ = v259
                                                } else {
                                                    v6 += v6
                                                }
                                                delete(v2, 0)
                                                v260 := GenS6{F0: v6}
                                                var v261 float32 = v14
                                                _ = v258
                                                _ = v260
                                                _ = v261
                                            }
                                        }
                                        _ = len(v52)
                                    }
                                    if v53 {
                                        var v262 uint = v54
                                        _ = v262
                                    }
                                }
                                if (!true) {
                                    switch (v16 + "hello") {
                                    case "\n\t":
                                        v263 := GenS38{F0: 1, F1: v4}
                                        _ = v263
                                    case "a":
                                        v4 = 154
                                    }
                                    var v264 uint = uint(v4)
                                    _ = v264
                                }
                                _ = v54
                            }
                            v14 -= (v14 * v14)
                            _ = v52
                            _ = v53
                        }
                    } else {
                        v1 = ""
                        var v265 int32 = 654
                        v266 := []int{}
                        _ = v265
                        _ = v266
                    }
                } else {
                    var v267 int = (fi17 + fi18)
                    v268 := GenS39{F0: 65535}
                    v267 = v5
                    _ = v267
                    _ = v268
                }
            default:
                v4 = v4
            }
            var v269 int32 = 2147483647
            _ = fi18
            _ = v269
        }
        v5 -= fi17
        _ = fi17
    }
    v5 -= (v5 / (((v5 - v5)) | 1))
    v5 += (v5 | (^(v5)))
    var v270 int64 = 0
    for fi271 := 0; fi271 < 6; fi271++ {
        var v272 float64 = -2.25
        v273 := GenS40{F0: 42, F1: 132, F2: (v4 - (v4 * 2))}
        _ = fi271
        _ = v272
        _ = v273
    }
    var v274 float32 = v14
    var v275 int32 = 929
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v270:", v270)
    fmt.Println("v274:", v274)
    fmt.Println("v275:", v275)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v270
    _ = v274
    _ = v275
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
