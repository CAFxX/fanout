package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 uint; F2 int32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 string; F1 int16; F2 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 bool; F1 uint32; F2 int }
type GenS13 struct { F0 int16; F1 int32; F2 int32 }
type GenS14 struct { F0 uint64; F1 bool }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS17 struct { F0 uint64; F1 uint64; F2 bool }
func genFn0(p0 uint64, p1 uint32, p2 int32) int {
    p1 += p1
    var v5 float64 = -2.25
    if true {
        if (0 > -9223372036854775808) {
            v6 := []int{}
            _ = v6
        }
        v5 = v5
        v5 = 1e10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 9223372036854775807
}
func genFn1() int16 {
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 42
        if (!(2 <= 343)) {
            v10 := GenS1{F0: (false || (195 < 0))}
            fmt.Println(v7)
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 1; fi11++ {
        if (!false) {
            for fi12 := 0; fi12 < 19; fi12++ {
                v7 -= v7
                _ = fi12
            }
            v13 := []int{fi11, (fi11 - int(v7)), 223}
            v7 += v7
            _ = v13
        } else {
            for fi14 := 0; fi14 < 13; fi14++ {
                if ((false && (-2.25 < float32(3.14159))) || true) {
                    fmt.Println(fi14, fi14, fi14)
                } else {
                    var v15 uint = 42
                    switch fi11 {
                    case 9223372036854775807:
                        v16 := GenS1{F0: false}
                        _ = v16
                    default:
                        var v17 bool = false
                        _ = v17
                    }
                    if true {
                        var v18 float64 = (0.1 * 0.1)
                        v19 := []int{fi14}
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS2{F0: (v7 ^ v7)}
                        v21 := GenS3{F0: "hello"}
                        var v22 uint64 = 0
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    switch fi14 {
                    case -9223372036854775808:
                        v23 := []int{32767, genFn0(560, 65535, -2147483648), fi14}
                        _ = v23
                    case -2147483648:
                        if (true || false) {
                            v15 -= v15
                            fmt.Println(v15, v15, v7)
                            var v24 uint16 = v7
                            _ = v24
                        } else {
                            switch (fi11 << ((65535) & 63)) {
                            case 128:
                                v25 := GenS4{F0: 1000}
                                _ = v25
                            }
                        }
                    case -70:
                        if (-9223372036854775808 < -9223372036854775808) {
                            var v26 bool = true
                            v15 += (v15 | 249)
                            v27 := GenS2{F0: v7}
                            _ = v26
                            _ = v27
                        } else {
                            var v28 uint32 = 303
                            _ = v28
                        }
                    default:
                        v7 -= 32767
                    }
                    v15 = (^(v15))
                    _ = v15
                }
                v7 += 32767
                _ = fi14
            }
            v29 := []int{(^(fi11)), genFn0(18446744073709551615, 815, -1), -2147483648}
            _ = v29
        }
        if (42 < -203) {
            var v30 int8 = 127
            if true {
                v31 := []int{}
                _ = v31
            }
            v30 += v30
            _ = v30
        } else {
            var v32 uint8 = 42
            _ = v32
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            var v34 int64 = -1
            var v35 int8 = 42
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi11
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        switch (("a" + "hello") + "hello") {
        case "hello":
            v37 := GenS2{F0: v7}
            _ = v37
        case "a":
            v38 := []int{genFn0(4294967295, 740, 1), int(v7), fi36, genFn0(282, 2, -2147483648)}
            _ = v38
        default:
            for fi39 := 0; fi39 < 19; fi39++ {
                switch genFn0(0, 2, 42) {
                case -1:
                    v40 := []int{(fi36 + (fi39 >> ((2) & 63))), fi39, 42, 4294967295}
                    _ = v40
                default:
                    v7 = 2
                }
                switch "hello" {
                case "":
                    var v41 int64 = 0
                    _ = v41
                default:
                    v42 := GenS5{F0: 127, F1: 425, F2: 788}
                    _ = v42
                }
                _ = fi39
            }
        }
        var v43 float64 = ((1e10 + 1e10) / 1e10)
        _ = fi36
        _ = v43
    }
    _ = v7
    return 1000
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 int32 = 921
    for fi44 := 0; fi44 < 2; fi44++ {
        fmt.Println(v4, v4)
        _ = fi44
    }
    v4 += v4
    var v45 bool = (v2 && (255 >= 26))
    if len(v1) > 0 {
        v1[0] = ("a" + "")
    }
    var v46 int64 = 572
    fmt.Println(v45, v4, v2)
    if ("" <= "a") {
        var v47 float32 = (0.1 + 0.0)
        if (!false) {
            for fi48 := 0; fi48 < 18; fi48++ {
                switch fi48 {
                case 32767:
                    var v49 int8 = 30
                    _ = v49
                case -1:
                    fmt.Println(v2, v2)
                case 81:
                    if (65535 == 65535) {
                        if (v4 > 42) {
                            v3[""] = v45
                            if (true || v45) {
                                if v45 {
                                    var v50 int16 = -460
                                    _ = v50
                                }
                            }
                        }
                        v2 = false
                    } else {
                        v51 := GenS6{F0: 32767}
                        _ = v51
                    }
                default:
                    var v52 int32 = 1
                    _ = v52
                }
                if false {
                    var v53 uint = 42
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    var v54 uint8 = 42
                    _ = v53
                    _ = v54
                }
                var v55 string = ""
                _ = fi48
                _ = v55
            }
            for fi56 := 0; fi56 < 6; fi56++ {
                _ = len(v1)
                if (((fi56 >> ((65535) & 63)) <= genFn0(42, 473, v4)) && v2) {
                    switch (("" + "\n\t") + "hello") {
                    case "hello":
                        v45 = ((0 >= 250) || (553 > -1))
                    default:
                        v47 += 3.14159
                    }
                    _ = len(v3)
                }
                _ = len(v3)
                _ = fi56
            }
            v45 = false
        } else {
            v2 = (65535 != 2)
            var v57 uint = 18446744073709551615
            _ = v57
        }
        for fi58 := 0; fi58 < 4; fi58++ {
            if ((-(v4)) != v4) {
                fmt.Println(v4, v47)
                v3["hello"] = v2
            }
            _ = fi58
        }
        _ = v47
    } else {
        v1 = append(v1, "a")
        var v59 int16 = -32768
        _ = v59
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        if (2 == 2) {
            for fi61 := 0; fi61 < 2; fi61++ {
                var v62 uint32 = 2147483648
                fmt.Println(v46)
                if false {
                    v63 := GenS2{F0: 1}
                    _ = v63
                } else {
                    v4 += -2147483648
                    switch 9223372036854775807 {
                    case 0:
                        v62 += 189
                    case 929:
                        fmt.Println(v45, fi61)
                    case 240:
                        v64 := GenS7{F0: "a", F1: (1e10 == 1e10)}
                        _ = v64
                    }
                }
                _ = fi61
                _ = v62
            }
        }
        var v65 float32 = 1.5
        var v66 string = (("\n\t" + "foo bar") + ("" + (("\n\t" + ("" + "hello")) + ("" + "a"))))
        _ = fi60
        _ = v65
        _ = v66
    }
    v67 := v3["\n\t"]
    v4 += (-(v4))
    _ = len(v1)
    var v68 uint16 = 554
    if (v67 || v67) {
        if v2 {
            v68 += (v68 + (v68 | v68))
        }
        var v69 int8 = -128
        var v70 uint8 = 1
        _ = v69
        _ = v70
    } else {
        var v71 string = "hello"
        v72 := GenS7{F0: (v71 + v71), F1: v45}
        if (255 <= 611) {
            var v73 uint32 = 1
            fmt.Println(v71, v4, v2)
            _ = v73
        } else {
            if v45 {
                v4 += 576
            }
            switch (((("" + "a") + "hello") + v71) + v71) {
            case "foo bar":
                if v45 {
                    v74 := GenS8{F0: uint32(v68)}
                    _ = v74
                } else {
                    var v75 uint16 = v68
                    v68 = (v68 << ((672) & 63))
                    _ = v75
                }
            }
            fmt.Println(v45, v67, v71)
        }
        _ = v71
        _ = v72
    }
    for fi76 := 0; fi76 < 9; fi76++ {
        var v77 uint8 = 2
        fmt.Println(v2, v77, v68)
        if (1 == (v46 ^ v46)) {
            if (1 != 0) {
                v46 = (v46 * (v46 + (v46 - v46)))
                if (!v45) {
                    if v67 {
                        var v78 int8 = 5
                        _ = v78
                    } else {
                        v79 := v3["foo bar"]
                        fmt.Println(v67, v77)
                        _ = v79
                    }
                }
                v46 = (v46 ^ (v46 / (((v46 * v46)) | 1)))
            } else {
                var v80 int64 = v46
                fmt.Println(v77, v80)
                switch (fi76 << ((4294967295) & 63)) {
                case -2147483648:
                    if v45 {
                        for fi81 := 0; fi81 < 4; fi81++ {
                            var v82 float32 = 0.1
                            var v83 bool = (53 != -1)
                            v84 := GenS9{F0: -19, F1: v4}
                            _ = fi81
                            _ = v82
                            _ = v83
                            _ = v84
                        }
                        if false {
                            switch "" {
                            case "a":
                                var v85 int16 = -32768
                                _ = v85
                            case "foo bar":
                                if false {
                                    v86 := GenS10{F0: "\n\t", F1: -157, F2: 285}
                                    v80 += v46
                                    var v87 uint8 = v77
                                    _ = v86
                                    _ = v87
                                } else {
                                    v88 := GenS1{F0: true}
                                    v89 := GenS1{F0: (uint(v77) == 42)}
                                    _ = v88
                                    _ = v89
                                }
                            }
                        } else {
                            if (1 <= 2) {
                                var v90 int32 = v4
                                fmt.Println(v2, v67, v2)
                                var v91 float32 = 1.5
                                _ = v90
                                _ = v91
                            } else {
                                fmt.Println(v45, v45)
                                delete(v3, "hello")
                                v92 := GenS11{F0: 13}
                                _ = v92
                            }
                            var v93 int = genFn0(4294967295, 0, 42)
                            fmt.Println(v46, v93, v77)
                            _ = v93
                        }
                        v94 := GenS12{F0: (0 > 7), F1: uint32(v77), F2: (fi76 << ((18446744073709551615) & 63))}
                        _ = v94
                    }
                case 814:
                    v95 := GenS13{F0: 1, F1: v4, F2: (^(v4))}
                    _ = v95
                case -88:
                    v96 := GenS6{F0: (^(fi76))}
                    _ = v96
                }
                _ = v80
            }
            v97 := GenS14{F0: 97, F1: false}
            if v45 {
                var v98 int64 = v46
                if (v2 && (true && true)) {
                    v2 = v45
                    v99 := GenS11{F0: v4}
                    for fi100 := 0; fi100 < 12; fi100++ {
                        var v101 float64 = float64((0.0 - 0.1))
                        _ = fi100
                        _ = v101
                    }
                    _ = v99
                }
                _ = v98
            } else {
                v2 = (true || v67)
            }
            _ = v97
        } else {
            switch (fi76 ^ genFn0(631, 657, v4)) {
            case -190:
                for fi102 := 0; fi102 < 18; fi102++ {
                    if v45 {
                        v45 = (v77 >= (v77 << ((814) & 63)))
                        v4 = (v4 % (((v4 & (-(v4)))) | 1))
                    } else {
                        v103 := GenS15{F0: 4294967295}
                        _ = len(v1)
                        v104 := GenS16{F0: 255, F1: 42, F2: 100}
                        _ = v103
                        _ = v104
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = ((1e10 > 1e10) || (!v2))
                    v67 = false
                    v105 := v3[""]
                    _ = fi102
                    _ = v105
                }
            default:
                fmt.Println(fi76)
            }
            v106 := GenS17{F0: uint64(v68), F1: 42, F2: (32767 > -141)}
            _ = v106
        }
        _ = fi76
        _ = v77
    }
    v2 = (2147483648 <= 2147483648)
    var v107 string = "hello"
    var v108 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v67
    _ = v68
}
