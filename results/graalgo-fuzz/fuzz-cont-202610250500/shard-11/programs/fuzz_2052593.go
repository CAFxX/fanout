package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int8; F1 int32; F2 uint32 }
type GenS8 struct { F0 uint16; F1 uint }
func genFn0(p0 uint16) uint64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        if true {
            var v8 string = "hello"
            for fi9 := 0; fi9 < 14; fi9++ {
                fmt.Println(fi9, fi9)
                v8 = v8
                var v10 int16 = 42
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint64 = 2
                if (1 <= 127) {
                    var v13 uint32 = 42
                    var v14 int = (-(fi11))
                    _ = v13
                    _ = v14
                }
                v15 := GenS1{F0: (fi7 ^ int(p0)), F1: 0, F2: v8}
                _ = fi11
                _ = v12
                _ = v15
            }
            _ = v8
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: fi7, F1: -2147483648, F2: ("hello" + "")}
                var v18 float32 = 0.0
                if (true && ((!(127 >= -1)) || (9223372036854775807 >= 929))) {
                    if (32767 <= 1) {
                        if true {
                            var v19 int8 = 42
                            var v20 uint64 = 18446744073709551615
                            v18 = v18
                            v19 += -1
                            _ = v19
                            _ = v20
                        } else {
                            v18 -= v18
                            var v21 float64 = 3.14159
                            if false {
                                v22 := GenS2{F0: "foo bar"}
                                _ = v22
                            } else {
                                var v23 int8 = -128
                                v23 = v23
                                _ = v23
                            }
                            _ = v21
                        }
                        var v24 float32 = float32((-2.25 / -2.25))
                        v25 := []int{fi16, (^(fi16)), -2147483648, (fi16 - fi7)}
                        p0 -= p0
                        var v26 bool = ("hello" <= "\n\t")
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        var v27 float32 = v18
                        v18 = 0.0
                        _ = v27
                    }
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            v28 := []int{128}
            p0 = 251
            _ = v28
        }
        _ = fi7
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 uint64 = 18446744073709551615
        for fi31 := 0; fi31 < 15; fi31++ {
            v30 = 2
            var v32 uint = 480
            switch "a" {
            case "hello":
                switch ("hello" + "a") {
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v32 -= (^(v32))
                    case "a":
                        if false {
                            if false {
                                p0 += (p0 ^ p0)
                            }
                        } else {
                            v33 := GenS2{F0: ("" + "")}
                            var v34 uint32 = 1
                            _ = v33
                            _ = v34
                        }
                    case "hello":
                        var v35 uint8 = 0
                        switch (fi29 << ((v32) & 63)) {
                        case 210:
                            p0 = (p0 - (p0 << (((v32 << ((uint(v30)) & 63))) & 63)))
                        }
                        v36 := []int{-1, 32767, fi29}
                        _ = v35
                        _ = v36
                    }
                case "hello":
                    v30 -= v30
                case "":
                    v37 := []int{}
                    _ = v37
                }
                var v38 uint64 = 1
                _ = v38
            case "foo bar":
                v39 := []int{}
                _ = v39
            }
            _ = fi31
            _ = v32
        }
        if false {
            for fi40 := 0; fi40 < 15; fi40++ {
                var v41 int64 = 0
                if false {
                    v42 := GenS3{F0: 2147483647}
                    _ = v42
                } else {
                    if (!false) {
                        if false {
                            var v43 uint64 = uint64(p0)
                            v44 := GenS4{F0: -915}
                            _ = v43
                            _ = v44
                        } else {
                            if (!((false || false) && (!(3.14159 < (0.1 / 1.0))))) {
                                v45 := GenS5{F0: -49}
                                if false {
                                    v41 = 599
                                    var v46 int32 = -782
                                    _ = v46
                                } else {
                                    switch (fi29 >> ((255) & 63)) {
                                    case 4294967295:
                                        var v47 bool = (true || true)
                                        _ = v47
                                    default:
                                        var v48 uint8 = uint8(p0)
                                        _ = v48
                                    }
                                    p0 = 255
                                }
                                var v49 uint16 = (p0 | p0)
                                _ = v45
                                _ = v49
                            } else {
                                switch "a" {
                                case "":
                                    v50 := []int{-9223372036854775808, 9223372036854775807, fi29}
                                    _ = v50
                                case "hello":
                                    fmt.Println(fi29, v41, v30)
                                case "a":
                                    switch fi40 {
                                    case 9223372036854775807:
                                        v41 = (v41 | int64(p0))
                                    case 42:
                                        v51 := GenS4{F0: v41}
                                        _ = v51
                                    case -316:
                                        v52 := []int{2147483647, (fi29 + fi29)}
                                        _ = v52
                                    }
                                }
                            }
                            fmt.Println(fi29)
                            var v53 bool = true
                            _ = v53
                        }
                        var v54 int64 = (v41 | v41)
                        if false {
                            v55 := GenS6{F0: 97, F1: -1}
                            _ = v55
                        } else {
                            v54 += (v54 >> ((2) & 63))
                            if false {
                                v56 := []int{(fi40 & fi40)}
                                var v57 uint64 = (^(v30))
                                _ = v56
                                _ = v57
                            }
                            p0 = p0
                        }
                        _ = v54
                    }
                }
                _ = fi40
                _ = v41
            }
            v58 := GenS5{F0: -100}
            _ = v58
        }
        _ = fi29
        _ = v30
    }
    var v59 int16 = 577
    if false {
        v59 = (v59 ^ -319)
        switch 32767 {
        case -2147483648:
            if false {
                var v60 int8 = 127
                v60 = v60
                _ = v60
            } else {
                v61 := []int{128, 2147483647}
                switch "" {
                case "foo bar":
                    v59 -= -1
                case "a":
                    v62 := GenS7{F0: -1, F1: 169, F2: 2}
                    _ = v62
                case "hello":
                    v61 = append(v61, 591)
                default:
                    var v63 uint16 = p0
                    _ = v63
                }
                _ = v61
            }
        case 9223372036854775807:
            if true {
                v64 := []int{-2147483648, 9223372036854775807}
                _ = v64
            } else {
                p0 = (p0 / (((p0 / ((p0) | 1))) | 1))
                var v65 uint64 = 4294967295
                _ = v65
            }
        case 32768:
            fmt.Println(p0, p0)
        }
        p0 -= p0
    } else {
        var v66 float32 = (1e10 - 3.14159)
        if false {
            v67 := []int{9223372036854775807, -151, -1, 354}
            var v68 int32 = -2147483648
            v67 = append(v67, 228)
            _ = v67
            _ = v68
        } else {
            v69 := GenS7{F0: 127, F1: 2147483647, F2: 2147483648}
            _ = v69
        }
        _ = v66
    }
    _ = p0
    _ = v59
    return 0
}

func main() {
    var v1 int16 = 1
    var v2 uint8 = 127
    v3 := make(map[int]int)
    var v4 int = 128
    var v5 int16 = (v1 + (v1 | (v1 << ((uint(v2)) & 63))))
    var v6 int = v4
    v2 = v2
    var v70 bool = (!(true || false))
    var v71 bool = v70
    v2 = (v2 << ((529) & 63))
    v72 := []int{v6, (v4 & v4), -18, v6}
    var v73 float32 = (1.5 * 3.14159)
    var v74 bool = ((v4 - (v6 % ((v6) | 1))) <= v4)
    for fi75 := 0; fi75 < 7; fi75++ {
        for fi76 := 0; fi76 < 18; fi76++ {
            var v77 int64 = 9223372036854775807
            fmt.Println(v6, v77)
            _ = fi76
            _ = v77
        }
        switch ("" + "hello") {
        case "a":
            if v74 {
                switch "\n\t" {
                case "\n\t":
                    v3[(-(v4))] = (v6 * (v4 - v4))
                case "":
                    var v78 uint = 0
                    _ = v78
                case "foo bar":
                    v72 = append(v72, int(v1))
                default:
                    switch "foo bar" {
                    case "a":
                        delete(v3, (v6 - v4))
                    case "":
                        v79 := GenS1{F0: (v6 >> ((18446744073709551615) & 63)), F1: 0, F2: "a"}
                        _ = v79
                    default:
                        delete(v3, (-(v6)))
                    }
                }
                for fi80 := 0; fi80 < 5; fi80++ {
                    var v81 string = "foo bar"
                    switch ("a" + ("hello" + v81)) {
                    case "":
                        _ = len(v72)
                    }
                    _ = fi80
                    _ = v81
                }
            } else {
                for fi82 := 0; fi82 < 6; fi82++ {
                    var v83 uint32 = 2147483648
                    _ = fi82
                    _ = v83
                }
                if (0 != 656) {
                    v3[v4] = v4
                }
            }
        case "hello":
            v84 := GenS2{F0: ("a" + ((("" + ("" + "a")) + "\n\t") + "a"))}
            _ = v84
        case "":
            var v85 uint64 = 370
            _ = v85
        default:
            v86 := v3[(fi75 + v4)]
            _ = v86
        }
        v87 := GenS1{F0: (fi75 * fi75), F1: -189, F2: "hello"}
        _ = fi75
        _ = v87
    }
    v88 := GenS8{F0: uint16(v2), F1: 42}
    var v89 int = int(v2)
    for fi90 := 0; fi90 < 12; fi90++ {
        var v91 uint32 = 65535
        var v92 float64 = (3.14159 + 1e10)
        _ = fi90
        _ = v91
        _ = v92
    }
    var v93 int32 = 201
    var v94 float32 = 1e10
    var v95 uint = 65535
    if len(v72) > 0 {
        v72[0] = (v6 & (v6 / ((v6) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    fmt.Println("v89:", v89)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v88
    _ = v89
    _ = v93
    _ = v94
    _ = v95
}
