package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 uint }
type GenS3 struct { F0 int64; F1 uint; F2 int }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 uint8; F2 bool }
type GenS7 struct { F0 uint; F1 int32; F2 uint8 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0() int8 {
    var v4 uint64 = 730
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    v4 -= v4
    v4 -= (v4 ^ (^(v4)))
    v4 = 2
    _ = v4
    return 1
}
func genFn1(p0 uint, p1 int16) int16 {
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            if (127 >= genFn0()) {
                p1 = -1
            } else {
                v8 := []int{-108, (fi6 | fi6), (fi7 * fi6), fi7}
                _ = len(v8)
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        if false {
            v10 := []int{fi9, int(p1), -9223372036854775808, fi9}
            _ = v10
        } else {
            fmt.Println(p1, p1)
            if false {
                if ((false || true) || (((false || false) || ((((0.0 - ((0.0 * 0.1) + 3.14159)) * (((1.5 / 0.1) / (3.14159 / 1.5)) + 1e10)) + (float64(float32(1e10)) * float64(0.1))) < 1e10)) && ((p0 << ((2) & 63)) != p0))) {
                    p1 = (p1 + (^(p1)))
                    for fi11 := 0; fi11 < 10; fi11++ {
                        var v12 bool = true
                        _ = fi11
                        _ = v12
                    }
                }
            }
        }
        _ = fi9
    }
    if (!false) {
        v13 := GenS1{F0: false, F1: 2147483647, F2: uint16(p0)}
        if true {
            v14 := GenS1{F0: false, F1: 9223372036854775807, F2: 32767}
            if true {
                var v15 uint8 = 2
                if (42 < 42) {
                    var v16 uint16 = 42
                    v17 := GenS2{F0: true, F1: -1, F2: p0}
                    _ = v16
                    _ = v17
                } else {
                    var v18 int8 = -128
                    _ = v18
                }
                _ = v15
            } else {
                var v19 int8 = 0
                var v20 int = 128
                _ = v19
                _ = v20
            }
            _ = v14
        }
        _ = v13
    } else {
        var v21 uint8 = 127
        if true {
            for fi22 := 0; fi22 < 19; fi22++ {
                switch ("" + (("hello" + ("hello" + "")) + ("hello" + ("foo bar" + "\n\t")))) {
                case "hello":
                    if (!false) {
                        if (!(42 >= genFn0())) {
                            var v23 uint = p0
                            v24 := GenS1{F0: (float64((((0.0 * 1e10) + 1e10) / 1.0)) <= 0.1), F1: (fi22 + (^(fi22))), F2: 0}
                            _ = v23
                            _ = v24
                        } else {
                            var v25 uint16 = 48
                            v26 := GenS3{F0: int64(v21), F1: p0, F2: fi22}
                            _ = v25
                            _ = v26
                        }
                    } else {
                        for fi27 := 0; fi27 < 16; fi27++ {
                            p1 = (p1 + -32768)
                            v28 := []int{-73, 128, fi22}
                            if (3.14159 >= 0.1) {
                                var v29 float32 = (1.5 + 1e10)
                                fmt.Println(v29, fi22)
                                var v30 uint16 = 42
                                _ = v29
                                _ = v30
                            } else {
                                switch (fi22 & -2147483648) {
                                case -2147483648:
                                    var v31 int8 = int8(fi27)
                                    _ = v31
                                default:
                                    var v32 float64 = 3.14159
                                    _ = v32
                                }
                                if (!(!(v21 > (v21 + 0)))) {
                                    switch fi22 {
                                    case 795:
                                        v33 := GenS4{F0: "\n\t", F1: 128}
                                        v21 = v21
                                        v34 := GenS4{F0: "", F1: 32768}
                                        _ = v33
                                        _ = v34
                                    case 1:
                                        v28 = append(v28, (fi27 % ((fi22) | 1)))
                                        switch -385 {
                                        case 42:
                                            if true {
                                                var v35 bool = false
                                                v36 := GenS2{F0: v35, F1: (p1 >> (((p0 ^ p0)) & 63)), F2: 1}
                                                _ = v35
                                                _ = v36
                                            } else {
                                                _ = len(v28)
                                            }
                                        case 1:
                                            if (true && (true || false)) {
                                                fmt.Println(p1)
                                                v37 := GenS3{F0: 0, F1: 255, F2: 9223372036854775807}
                                                var v38 uint64 = 583
                                                _ = v37
                                                _ = v38
                                            } else {
                                                switch (^(fi22)) {
                                                case 1:
                                                    v21 -= 231
                                                case 128:
                                                    fmt.Println(p0, fi27, fi27)
                                                default:
                                                    _ = len(v28)
                                                }
                                                fmt.Println(p1, v21, p0)
                                                var v39 uint8 = v21
                                                var v40 float64 = -2.25
                                                _ = v39
                                                _ = v40
                                            }
                                        default:
                                            if true {
                                                p0 = p0
                                                v41 := GenS5{F0: 2}
                                                if (1 > -2147483648) {
                                                    var v42 string = "hello"
                                                    if true {
                                                        switch ((v42 + "") + v42) {
                                                        case "foo bar":
                                                            p1 = p1
                                                        default:
                                                            fmt.Println(p0)
                                                            var v43 int16 = (p1 % ((p1) | 1))
                                                            _ = v43
                                                        }
                                                        fmt.Println(v21, fi22, v42)
                                                    } else {
                                                        var v44 float32 = -2.25
                                                        if len(v28) >= 2 {
                                                            v28 = v28[1:len(v28)]
                                                        }
                                                        p0 += p0
                                                        _ = v44
                                                    }
                                                    _ = v42
                                                }
                                                if (!true) {
                                                    v45 := GenS3{F0: 42, F1: p0, F2: 614}
                                                    _ = v45
                                                }
                                                fmt.Println(p1)
                                                _ = v41
                                            } else {
                                                if (!((fi27 | int(p0)) <= fi22)) {
                                                    p1 = -992
                                                } else {
                                                    v21 -= 42
                                                }
                                                if false {
                                                    var v46 uint8 = (v21 + 255)
                                                    var v47 int = int(p1)
                                                    v48 := GenS6{F0: 2147483647, F1: v21, F2: (!false)}
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    var v49 uint64 = 18446744073709551615
                                                    v50 := GenS1{F0: (1.5 > 0.1), F1: fi22, F2: 0}
                                                    v49 = 4294967295
                                                    _ = v49
                                                    _ = v50
                                                }
                                                var v51 uint8 = (v21 & v21)
                                                _ = v51
                                            }
                                        }
                                    default:
                                        _ = len(v28)
                                    }
                                    if len(v28) > 0 {
                                        v28[0] = 82
                                    }
                                } else {
                                    if len(v28) > 0 {
                                        v28[0] = 297
                                    }
                                }
                                _ = len(v28)
                            }
                            if (!false) {
                                v28 = append(v28, (fi22 ^ (fi27 - fi22)))
                                v21 = (v21 - (v21 & (v21 % ((127) | 1))))
                                if true {
                                    p0 = 902
                                    var v52 uint64 = 42
                                    _ = v52
                                }
                                p1 = p1
                                v53 := GenS3{F0: -1, F1: (p0 + 639), F2: (fi27 ^ (^(fi22)))}
                                _ = v53
                            }
                            fmt.Println(fi27, p1)
                            if true {
                                fmt.Println(fi27, p1)
                            } else {
                                v21 = v21
                                v54 := GenS7{F0: 65535, F1: -2147483648, F2: v21}
                                _ = v54
                            }
                            var v55 uint16 = 1
                            _ = fi27
                            _ = v28
                            _ = v55
                        }
                    }
                case "foo bar":
                    var v56 uint64 = 445
                    _ = v56
                default:
                    var v57 float32 = 1.5
                    _ = v57
                }
                _ = fi22
            }
        }
        fmt.Println(p1, v21, p0)
        _ = v21
    }
    v58 := []int{128, int(p1), 0}
    _ = p0
    _ = p1
    _ = v58
    return (p1 >> (((p0 - p0)) & 63))
}

func main() {
    var v1 uint8 = 0
    var v2 bool = true
    var v3 int16 = 32767
    v59 := []int{42, 2147483647}
    for fi60 := 0; fi60 < 20; fi60++ {
        var v61 int32 = 1
        v1 += 128
        _ = fi60
        _ = v61
    }
    switch 0 {
    case 821:
        v2 = (298 <= 638)
    case 1:
        switch "foo bar" {
        case "foo bar":
            v2 = true
        case "hello":
            fmt.Println(v2)
        case "\n\t":
            switch 4294967295 {
            case -118:
                v62 := GenS8{F0: 248, F1: -2147483648, F2: 1}
                _ = v62
            }
        default:
            for fi63 := 0; fi63 < 10; fi63++ {
                var v64 int64 = 487
                switch -700 {
                case -2147483648:
                    _ = len(v59)
                case 32767:
                    v59 = append(v59, -9223372036854775808)
                default:
                    _ = len(v59)
                }
                if (42 >= 2147483648) {
                    if v2 {
                        var v65 float64 = 0.0
                        v65 = v65
                        var v66 int32 = -2147483648
                        _ = v65
                        _ = v66
                    }
                    v67 := GenS9{F0: v64}
                    _ = v67
                }
                _ = fi63
                _ = v64
            }
        }
    default:
        v68 := GenS8{F0: v1, F1: 1, F2: 42}
        _ = v68
    }
    var v69 uint = 0
    switch "hello" {
    case "":
        for fi70 := 0; fi70 < 13; fi70++ {
            v1 += v1
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
            var v71 int16 = 32767
            _ = fi70
            _ = v71
        }
    case "a":
        v72 := GenS10{F0: ("foo bar" + ("" + "foo bar")), F1: 4294967295, F2: 1}
        _ = v72
    case "hello":
        switch "hello" {
        case "a":
            v59 = append(v59, 4294967295)
        case "hello":
            switch ("foo bar" + "a") {
            case "":
                if v2 {
                    v2 = (v2 && (!v2))
                    var v73 uint64 = 2
                    _ = v73
                }
            case "foo bar":
                v1 = v1
            default:
                var v74 uint8 = (v1 - (v1 ^ (v1 - 209)))
                _ = v74
            }
        }
    default:
        fmt.Println(v2, v2)
    }
    var v75 uint16 = 0
    v76 := GenS4{F0: "\n\t", F1: 2147483647}
    var v77 float64 = 0.0
    if false {
        fmt.Println(v1, v69)
        for fi78 := 0; fi78 < 8; fi78++ {
            v3 -= v3
            v3 += v3
            _ = fi78
        }
        _ = len(v59)
    }
    if ((1 < -1) && v2) {
        if ((v3 - 32767) >= (v3 & 1)) {
            var v79 uint = v69
            for fi80 := 0; fi80 < 17; fi80++ {
                v81 := GenS11{F0: v79, F1: 2147483648, F2: (v1 / (((v1 * (^(v1)))) | 1))}
                var v82 int = 2147483647
                if (((!(42 <= 65535)) || (-1 < 1)) || v2) {
                    var v83 float32 = 3.14159
                    var v84 int32 = 42
                    var v85 string = "a"
                    _ = v83
                    _ = v84
                    _ = v85
                }
                _ = fi80
                _ = v81
                _ = v82
            }
            _ = v79
        } else {
            fmt.Println(v75, v69)
            var v86 uint16 = (v75 % ((42) | 1))
            _ = v86
        }
    }
    for fi87 := 0; fi87 < 17; fi87++ {
        var v88 uint32 = 42
        _ = fi87
        _ = v88
    }
    for fi89 := 0; fi89 < 5; fi89++ {
        v90 := GenS5{F0: (v1 / (((^(v1))) | 1))}
        _ = fi89
        _ = v90
    }
    v2 = (!v2)
    var v91 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v69:", v69)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v69
    _ = v75
    _ = v76
    _ = v77
    _ = v91
}
