package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }
type GenS3 struct { F0 uint32; F1 int16; F2 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 int64 }
type GenS11 struct { F0 int8; F1 int8; F2 int }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int8; F1 uint }
type GenS14 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int32; F1 int32 }
func genFn0(p0 int16, p1 int8, p2 string) bool {
    var v7 uint = 517
    var v8 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (false && (!(v7 < 706)))
}
func genFn1(p0 uint32, p1 int64, p2 string) uint8 {
    switch 518 {
    case 825:
        p1 -= (-(p1))
    }
    if (!((p0 * p0) >= 191)) {
        p1 -= 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = p2
    return 180
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = 3.14159
    var v3 uint64 = 184
    var v4 uint32 = 65535
    var v5 int = 32767
    var v6 float64 = ((-2.25 - (v2 - -2.25)) - v2)
    switch (v5 - (v5 + 384)) {
    case -549:
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{2147483647, (^(fi9)), 1, v5}
            _ = fi9
            _ = v10
        }
    case 42:
        var v11 uint64 = 18446744073709551615
        _ = v11
    case 127:
        var v12 int64 = -1
        _ = v12
    }
    v3 -= 1
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint = 364
        if (12 != 120) {
            v15 := GenS1{F0: 1}
            if (false && (false || true)) {
                for fi16 := 0; fi16 < 20; fi16++ {
                    v17 := []int{fi16}
                    fmt.Println(v1, v5)
                    switch "a" {
                    case "foo bar":
                        var v18 string = ""
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
                var v19 uint8 = uint8(v1)
                if (42 >= 0) {
                    var v20 int64 = 1
                    v1 += v3
                    _ = v20
                } else {
                    if true {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            if false {
                                if true {
                                    v5 -= (fi13 / ((fi21) | 1))
                                }
                            }
                            if ("a" != ("\n\t" + "\n\t")) {
                                switch "hello" {
                                case "a":
                                    v6 -= (3.14159 * v6)
                                case "hello":
                                    v4 += (v4 | (^(v4)))
                                case "":
                                    v14 = v14
                                }
                                v22 := []int{0, fi13, (fi13 % ((42) | 1)), 2147483647}
                                if ((v19 | genFn1(42, -1, "\n\t")) > v19) {
                                    if len(v22) > 0 {
                                        v22[0] = fi13
                                    }
                                    fmt.Println(v6, v14, v5)
                                    var v23 uint16 = 1
                                    fmt.Println(v19)
                                    v5 = 9223372036854775807
                                    _ = v23
                                } else {
                                    v14 -= (^(v14))
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v1, v2)
                                    default:
                                        var v24 uint8 = v19
                                        var v25 uint64 = v3
                                        v14 -= (v14 - (v14 * (v14 / ((v14) | 1))))
                                        fmt.Println(v6, v3)
                                        _ = v24
                                        _ = v25
                                    }
                                }
                                if false {
                                    v26 := GenS1{F0: 42}
                                    _ = v26
                                } else {
                                    v3 = 2
                                    v4 += 2
                                    var v27 int16 = -32768
                                    _ = v27
                                }
                                v2 = 0.1
                                v4 += (^(v4))
                                fmt.Println(v4)
                                var v28 int8 = 42
                                switch (fi13 | (v5 - (v5 >> ((v14) & 63)))) {
                                case 614:
                                    fmt.Println(v6)
                                case -1:
                                    v19 += (v19 & 31)
                                    switch fi13 {
                                    case -1:
                                        switch "foo bar" {
                                        case "foo bar":
                                            v19 = genFn1(v4, 745, "a")
                                        case "a":
                                            v29 := GenS2{F0: (v14 * (v14 & (v14 | (v14 << ((0) & 63))))), F1: v3, F2: true}
                                            _ = v29
                                        case "\n\t":
                                            v30 := GenS3{F0: v4, F1: 0, F2: (v28 % ((127) | 1))}
                                            _ = v30
                                        default:
                                            var v31 int32 = -1
                                            _ = v31
                                        }
                                        switch v5 {
                                        case -9223372036854775808:
                                            _ = len(v22)
                                            v14 += (^(v14))
                                            fmt.Println(v4, v28)
                                            var v32 float32 = 0.0
                                            if len(v22) > 0 {
                                                v22[0] = int(v4)
                                            }
                                            if (true && false) {
                                                v5 -= (-(v5))
                                                var v33 uint64 = v1
                                                if (true && true) {
                                                    fmt.Println(v28, v4)
                                                }
                                                _ = v33
                                            } else {
                                                if len(v22) >= 2 {
                                                    v22 = v22[1:len(v22)]
                                                }
                                            }
                                            _ = v32
                                        case 32767:
                                            var v34 uint8 = (^(v19))
                                            _ = v34
                                        }
                                        v35 := GenS2{F0: (v14 ^ v14), F1: 4294967295, F2: ("foo bar" != "hello")}
                                        _ = v35
                                    case 9223372036854775807:
                                        var v36 uint32 = v4
                                        _ = v36
                                    default:
                                        v28 += v28
                                    }
                                    if (!(!true)) {
                                        switch (-(v5)) {
                                        case 1:
                                            _ = len(v22)
                                        }
                                    } else {
                                        v37 := GenS1{F0: 816}
                                        v38 := GenS2{F0: v14, F1: (v3 | (v1 / ((v3) | 1))), F2: true}
                                        if true {
                                            if (!false) {
                                                var v39 int16 = -698
                                                v2 = v2
                                                v19 -= 136
                                                _ = v39
                                            } else {
                                                fmt.Println(v2)
                                                v40 := GenS1{F0: 819}
                                                _ = v40
                                            }
                                            fmt.Println(v5, v14)
                                            v22 = append(v22, fi21)
                                            if len(v22) >= 2 {
                                                v22 = v22[1:len(v22)]
                                            }
                                            v5 = fi13
                                            var v41 float64 = v2
                                            _ = v41
                                        } else {
                                            v22 = append(v22, (fi13 ^ (fi13 % ((fi13) | 1))))
                                            v5 += (v5 << ((v14) & 63))
                                        }
                                        if len(v22) > 0 {
                                            v22[0] = 1
                                        }
                                        _ = v37
                                        _ = v38
                                    }
                                }
                                _ = v22
                                _ = v28
                            } else {
                                var v42 uint32 = v4
                                switch -9223372036854775808 {
                                case 536:
                                    v43 := GenS4{F0: "\n\t", F1: 1}
                                    _ = v43
                                case 2147483647:
                                    var v44 int16 = 1000
                                    _ = v44
                                case 128:
                                    v45 := []int{(fi21 | 32768), v5}
                                    _ = v45
                                default:
                                    v46 := GenS4{F0: ("a" + (("a" + "a") + "a")), F1: 42}
                                    _ = v46
                                }
                                if true {
                                    if false {
                                        switch "a" {
                                        case "":
                                            v47 := []int{(v5 * v5), fi21, -560}
                                            _ = v47
                                        default:
                                            var v48 float32 = ((1e10 - 1.5) * float32((1.5 / v2)))
                                            _ = v48
                                        }
                                        if false {
                                            var v49 int64 = -9223372036854775808
                                            v4 = (v4 / ((1) | 1))
                                            v6 += v2
                                            _ = v49
                                        } else {
                                            switch ("foo bar" + ("a" + "")) {
                                            case "a":
                                                v50 := []int{}
                                                switch (fi21 - int(v14)) {
                                                case 128:
                                                    if (((3.14159 + 0.1) - float32(((v2 * v2) * 0.1))) < (((1.5 + 0.0) - 1.5) / 1.5)) {
                                                        var v51 uint64 = uint64(v4)
                                                        v3 = (v51 | (v1 | v51))
                                                        _ = v51
                                                    } else {
                                                        if len(v50) > 0 {
                                                            v50[0] = fi13
                                                        }
                                                    }
                                                default:
                                                    v52 := GenS5{F0: (v5 & (v5 & fi13))}
                                                    v53 := GenS6{F0: (^(v14)), F1: 127}
                                                    _ = v52
                                                    _ = v53
                                                }
                                                _ = v50
                                            case "\n\t":
                                                v54 := GenS7{F0: (v19 + (^(v19)))}
                                                _ = v54
                                            case "":
                                                v3 -= (v1 & v1)
                                            }
                                        }
                                    }
                                    switch ("a" + ("foo bar" + "")) {
                                    case "a":
                                        var v55 uint8 = (v19 - (^(v19)))
                                        _ = v55
                                    case "":
                                        v56 := []int{v5, -401}
                                        _ = v56
                                    default:
                                        v19 = v19
                                    }
                                    v57 := GenS7{F0: (^(v19))}
                                    _ = v57
                                } else {
                                    v4 -= v4
                                    var v58 float32 = ((0.0 - 1e10) + ((float32(float64(3.14159)) - -2.25) - 1.5))
                                    var v59 int = fi21
                                    _ = v58
                                    _ = v59
                                }
                                _ = v42
                            }
                            v60 := []int{}
                            switch fi21 {
                            case -954:
                                fmt.Println(v14, v5, fi13)
                            case -9223372036854775808:
                                fmt.Println(v1, v5)
                            default:
                                v61 := GenS8{F0: "\n\t"}
                                _ = v61
                            }
                            v62 := GenS9{F0: (v5 << ((v14) & 63))}
                            _ = fi21
                            _ = v60
                            _ = v62
                        }
                        for fi63 := 0; fi63 < 5; fi63++ {
                            fmt.Println(v6, v5, v14)
                            v64 := GenS10{F0: v1, F1: -9223372036854775808}
                            v2 = v6
                            _ = fi63
                            _ = v64
                        }
                        var v65 int32 = -2147483648
                        _ = v65
                    }
                }
                _ = v19
            }
            if true {
                v4 += 2147483647
            } else {
                v1 = (v1 + v1)
                if ((true || (("hello" + "") >= "hello")) || (42 != 42)) {
                    v3 = 42
                    if false {
                        if (-1 > 435) {
                            fmt.Println(v3)
                            for fi66 := 0; fi66 < 18; fi66++ {
                                v14 += (v14 - (v14 + v14))
                                fmt.Println(v14)
                                if (-118 == -1) {
                                    v67 := GenS9{F0: -513}
                                    v6 = v6
                                    v68 := []int{fi13}
                                    _ = v67
                                    _ = v68
                                }
                                _ = fi66
                            }
                            v69 := []int{(-(v5)), fi13}
                            _ = v69
                        }
                    } else {
                        v70 := GenS11{F0: 68, F1: -47, F2: -587}
                        v71 := GenS12{F0: (v3 != v1)}
                        _ = v70
                        _ = v71
                    }
                }
            }
            _ = v15
        }
        v72 := []int{362, (v5 * v5)}
        _ = fi13
        _ = v14
        _ = v72
    }
    v73 := []int{(v5 * v5), (^(v5))}
    v74 := GenS13{F0: -1, F1: 2}
    var v75 bool = (true || (977 <= 42))
    v5 -= (-(v5))
    v76 := GenS13{F0: 127, F1: 0}
    if true {
        if false {
            var v77 uint = 65535
            var v78 int16 = 0
            _ = v77
            _ = v78
        }
    }
    v4 = v4
    if true {
        for fi79 := 0; fi79 < 11; fi79++ {
            var v80 bool = ((^(v4)) != v4)
            _ = fi79
            _ = v80
        }
        fmt.Println(v3, v3, v3)
    } else {
        v3 -= (v3 / ((1) | 1))
    }
    var v81 bool = v75
    v82 := GenS14{F0: (!true), F1: (v4 % ((v4) | 1)), F2: v4}
    v73 = append(v73, (v5 & v5))
    v3 += 42
    v83 := GenS15{F0: 0}
    v84 := GenS16{F0: -1, F1: -2147483648}
    v75 = (2 > 2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v75:", v75)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
