package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 int64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint32 }
func genFn0(p0 bool, p1 string) uint32 {
    if true {
        if (!p0) {
            fmt.Println(p0, p0)
            v7 := GenS1{F0: 483, F1: 100}
            if (!(true && (1.5 > (3.14159 * ((0.1 * 0.1) / 1.0))))) {
                if true {
                    var v8 uint8 = 127
                    var v9 int16 = -1
                    p0 = (!("\n\t" >= ("\n\t" + (p1 + p1))))
                    _ = v8
                    _ = v9
                } else {
                    var v10 uint32 = 42
                    v11 := []int{-1}
                    _ = v10
                    _ = v11
                }
                v12 := GenS2{F0: -1, F1: -1, F2: 26}
                _ = v12
            }
            _ = v7
        } else {
            var v13 uint8 = 249
            _ = v13
        }
        if (2 == 2) {
            var v14 uint64 = 0
            _ = v14
        }
        p1 = "hello"
    } else {
        if ("" < ("foo bar" + "hello")) {
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(fi15, p0, fi15)
                _ = fi15
            }
            p1 = p1
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                if p0 {
                    p1 = p1
                    if true {
                        v18 := GenS2{F0: 42, F1: -219, F2: 191}
                        if p0 {
                            p1 = p1
                        }
                        var v19 float32 = 0.1
                        v20 := GenS2{F0: 1, F1: 9223372036854775807, F2: -9223372036854775808}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                }
                v21 := []int{(fi17 & fi17), (fi16 - 32767)}
                p1 = "a"
                _ = fi17
                _ = v21
            }
            p0 = p0
            v22 := GenS1{F0: 32767, F1: -97}
            _ = fi16
            _ = v22
        }
    }
    var v23 string = p1
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 int8 = 1
            v27 := GenS1{F0: 32767, F1: 42}
            var v28 uint = 2
            var v29 float32 = (float32((3.14159 + 1.5)) - (3.14159 - (1e10 / 1e10)))
            _ = fi25
            _ = v26
            _ = v27
            _ = v28
            _ = v29
        }
        if p0 {
            switch (v23 + (("" + ((v23 + p1) + ("foo bar" + ((v23 + "a") + (p1 + p1))))) + ("hello" + v23))) {
            case "hello":
                switch (-(fi24)) {
                case -9223372036854775808:
                    var v30 bool = false
                    _ = v30
                case 1:
                    fmt.Println(p0)
                default:
                    for fi31 := 0; fi31 < 3; fi31++ {
                        switch fi31 {
                        case -1:
                            if p0 {
                                var v32 float64 = 0.0
                                p0 = p0
                                _ = v32
                            } else {
                                var v33 float32 = (3.14159 * (0.1 - -2.25))
                                _ = v33
                            }
                        }
                        _ = fi31
                    }
                }
            case "\n\t":
                var v34 uint16 = 569
                _ = v34
            default:
                var v35 int16 = -1
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 19; fi36++ {
            fmt.Println(p1, fi36, fi24)
            v23 = p1
            p1 = (("a" + v23) + p1)
            _ = fi36
        }
        _ = fi24
    }
    v37 := GenS1{F0: 757, F1: -128}
    for fi38 := 0; fi38 < 11; fi38++ {
        for fi39 := 0; fi39 < 11; fi39++ {
            fmt.Println(p1)
            fmt.Println(p0, fi39, p0)
            var v40 uint8 = 2
            if (127 == 42) {
                v40 -= v40
                v41 := []int{(fi39 % ((42) | 1)), fi38, (^(fi38))}
                var v42 float64 = (0.1 * ((1e10 * 0.1) - float64(1e10)))
                p0 = (p0 && false)
                _ = v41
                _ = v42
            } else {
                v43 := GenS3{F0: 787, F1: 956}
                _ = v43
            }
            if p0 {
                p1 = v23
                v44 := []int{fi38, (^(fi38)), (^(fi38)), -179}
                _ = v44
            }
            _ = fi39
            _ = v40
        }
        for fi45 := 0; fi45 < 7; fi45++ {
            v46 := GenS3{F0: 42, F1: 653}
            switch (fi45 ^ fi38) {
            case 9223372036854775807:
                if p0 {
                    var v47 float64 = 1.5
                    _ = v47
                }
            }
            _ = fi45
            _ = v46
        }
        v48 := []int{4294967295, 127, (fi38 >> ((65535) & 63)), 35}
        _ = fi38
        _ = v48
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v37
    return 2147483648
}
func genFn1(p0 uint8) uint16 {
    v49 := GenS1{F0: 65535, F1: -128}
    p0 += (p0 + p0)
    if (42 > 0) {
        v50 := []int{}
        v51 := GenS2{F0: 1, F1: -334, F2: -437}
        _ = v50
        _ = v51
    }
    var v52 int8 = -1
    if true {
        var v53 string = "\n\t"
        switch "\n\t" {
        case "\n\t":
            for fi54 := 0; fi54 < 11; fi54++ {
                v55 := GenS4{F0: int32(fi54), F1: 32767, F2: 580}
                if (-342 <= 1) {
                    var v56 int64 = 1
                    _ = v56
                }
                _ = fi54
                _ = v55
            }
        case "foo bar":
            v57 := GenS5{F0: 42, F1: 0}
            _ = v57
        case "":
            fmt.Println(v52)
        }
        if (-2.25 > -2.25) {
            var v58 uint8 = (p0 >> ((621) & 63))
            v59 := []int{128}
            _ = v58
            _ = v59
        }
        _ = v53
    } else {
        for fi60 := 0; fi60 < 4; fi60++ {
            if ((!true) && true) {
                p0 = p0
                switch 42 {
                case -1:
                    for fi61 := 0; fi61 < 9; fi61++ {
                        if ("" >= ("a" + "a")) {
                            p0 = (p0 * (p0 ^ (^(p0))))
                        } else {
                            switch -642 {
                            case 48:
                                v52 -= v52
                            }
                            v52 += (^(v52))
                            v52 = -1
                        }
                        fmt.Println(fi61, v52, fi61)
                        _ = fi61
                    }
                default:
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v52 = (v52 * (v52 << ((42) & 63)))
                        if true {
                            if true {
                                if ("a" != "hello") {
                                    v63 := GenS3{F0: 42, F1: 1}
                                    _ = v63
                                }
                                var v64 int64 = 1
                                _ = v64
                            }
                        } else {
                            v65 := GenS1{F0: 32767, F1: v52}
                            v52 = v52
                            _ = v65
                        }
                        v52 += v52
                        _ = fi62
                    }
                }
            } else {
                if (-2147483648 < -653) {
                    fmt.Println(fi60)
                } else {
                    fmt.Println(p0)
                }
            }
            fmt.Println(v52, fi60, v52)
            for fi66 := 0; fi66 < 4; fi66++ {
                switch "foo bar" {
                case "\n\t":
                    v67 := GenS4{F0: 522, F1: 42, F2: 9223372036854775807}
                    _ = v67
                case "hello":
                    v52 -= v52
                }
                switch fi60 {
                case 42:
                    if (!true) {
                        v68 := GenS6{F0: 902, F1: 340}
                        var v69 uint = 18446744073709551615
                        v70 := []int{}
                        _ = v68
                        _ = v69
                        _ = v70
                    } else {
                        p0 = 127
                        var v71 uint8 = (p0 % (((p0 ^ (p0 - p0))) | 1))
                        v52 -= 127
                        _ = v71
                    }
                }
                var v72 string = "hello"
                _ = fi66
                _ = v72
            }
            _ = fi60
        }
        v52 -= (^(v52))
        v73 := GenS3{F0: 18446744073709551615, F1: 4294967295}
        _ = v73
    }
    _ = p0
    _ = v49
    _ = v52
    return 65535
}

func main() {
    v1 := []int{0}
    var v2 int8 = -96
    v3 := make(map[int]string)
    var v4 int32 = 42
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]uint32)
    var v74 int32 = v4
    for fi75 := 0; fi75 < 20; fi75++ {
        for fi76 := 0; fi76 < 18; fi76++ {
            var v77 uint = 0
            var v78 float32 = 0.0
            _ = fi76
            _ = v77
            _ = v78
        }
        v2 = v2
        for fi79 := 0; fi79 < 2; fi79++ {
            var v80 float64 = 3.14159
            _ = fi79
            _ = v80
        }
        _ = fi75
    }
    v2 = (v2 >> ((255) & 63))
    if true {
        v81 := GenS7{F0: 65535, F1: -128}
        fmt.Println(v4, v74)
        _ = v81
    } else {
        v82 := GenS4{F0: 42, F1: 456, F2: -9223372036854775808}
        v2 = v2
        _ = v82
    }
    var v83 int64 = 1
    if ((v74 >> ((65535) & 63)) < (v4 ^ v74)) {
        v74 = (^(v74))
    } else {
        v84 := GenS1{F0: 1, F1: v2}
        v4 -= v4
        _ = v84
    }
    if (true || (false || true)) {
        var v85 float32 = 3.14159
        v86 := GenS8{F0: genFn0(true, (("\n\t" + ("hello" + "hello")) + ("\n\t" + ("foo bar" + ""))))}
        var v87 float32 = v85
        _ = v85
        _ = v86
        _ = v87
    } else {
        var v88 int64 = (v83 * v83)
        _ = v88
    }
    v2 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v83
}
