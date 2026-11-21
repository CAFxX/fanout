package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 uint16; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 uint }
func genFn0(p0 bool, p1 int64, p2 uint32) int {
    for fi4 := 0; fi4 < 19; fi4++ {
        v5 := []int{(^(fi4)), 447}
        _ = fi4
        _ = v5
    }
    p1 = (^(p1))
    p2 = (^(p2))
    if (38 >= 947) {
        var v6 int8 = -128
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int64 = p1
            fmt.Println(p0, p2)
            _ = fi7
            _ = v8
        }
        var v9 int32 = 1
        _ = v6
        _ = v9
    }
    var v10 int = int(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (v10 | v10)
}

func main() {
    var v1 float32 = (0.0 - 1e10)
    var v2 uint64 = 2
    var v3 int16 = -1
    switch "" {
    case "foo bar":
        var v11 uint = 707
        _ = v11
    default:
        if true {
            var v12 int64 = 42
            _ = v12
        } else {
            v13 := []int{462, 2147483647}
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint64 = v2
                v2 += 42
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    }
    v16 := GenS1{F0: v2, F1: "foo bar", F2: 9223372036854775807}
    for fi17 := 0; fi17 < 14; fi17++ {
        if true {
            v1 += v1
            var v18 uint = 4294967295
            _ = v18
        } else {
            v19 := GenS2{F0: "a", F1: (-(v3)), F2: 0}
            for fi20 := 0; fi20 < 7; fi20++ {
                switch "a" {
                case "hello":
                    var v21 int = (fi20 + (fi20 | fi17))
                    fmt.Println(v21, fi20, v1)
                    switch (("\n\t" + "a") + "hello") {
                    case "hello":
                        if true {
                            v2 = (v2 ^ 601)
                        }
                    case "\n\t":
                        if false {
                            var v22 bool = false
                            var v23 int16 = (v3 * v3)
                            var v24 int = 32767
                            v23 = -996
                            v25 := GenS3{F0: 0, F1: (("foo bar" + "hello") + (("" + ("a" + "hello")) + ("hello" + ("a" + ""))))}
                            _ = v22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    default:
                        v26 := GenS4{F0: fi20}
                        _ = v26
                    }
                    fmt.Println(v21, fi17)
                    _ = v21
                case "\n\t":
                    v27 := []int{}
                    _ = v27
                case "a":
                    v1 = (v1 / v1)
                }
                var v28 uint64 = v2
                _ = fi20
                _ = v28
            }
            _ = v19
        }
        _ = fi17
    }
    switch 127 {
    case -2147483648:
        for fi29 := 0; fi29 < 2; fi29++ {
            for fi30 := 0; fi30 < 9; fi30++ {
                switch fi29 {
                case -842:
                    var v31 int = genFn0((!true), -148, 1)
                    _ = v31
                }
                _ = fi30
            }
            v2 -= v2
            _ = fi29
        }
    }
    var v32 uint16 = 65535
    for fi33 := 0; fi33 < 1; fi33++ {
        switch ("foo bar" + (("foo bar" + "") + "foo bar")) {
        case "a":
            if true {
                switch "hello" {
                case "\n\t":
                    switch "hello" {
                    case "\n\t":
                        v34 := GenS2{F0: ("a" + (("a" + ("\n\t" + ("" + "\n\t"))) + "hello")), F1: 1, F2: -1}
                        _ = v34
                    }
                default:
                    var v35 int64 = 42
                    _ = v35
                }
                v36 := GenS3{F0: 963, F1: "\n\t"}
                v37 := GenS5{F0: ("hello" + "\n\t"), F1: (^(v32)), F2: (fi33 - (fi33 / ((fi33) | 1)))}
                _ = v36
                _ = v37
            }
        case "":
            var v38 float32 = -2.25
            _ = v38
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            v40 := GenS6{F0: 65535}
            switch "\n\t" {
            case "\n\t":
                v41 := GenS5{F0: ("a" + "hello"), F1: (v32 | 65535), F2: (^(fi33))}
                _ = v41
            default:
                if (-128 > -1) {
                    v42 := []int{fi33}
                    _ = v42
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi33
    }
    if ((v32 << ((255) & 63)) == v32) {
        var v43 float32 = float32(0.0)
        if ((v32 ^ v32) != v32) {
            if ((2 >= 204) || true) {
                for fi44 := 0; fi44 < 15; fi44++ {
                    if (false && false) {
                        v45 := GenS4{F0: genFn0(true, 1, 4294967295)}
                        var v46 int64 = 1
                        for fi47 := 0; fi47 < 2; fi47++ {
                            switch "foo bar" {
                            case "\n\t":
                                var v48 float64 = 0.1
                                v49 := []int{}
                                _ = v48
                                _ = v49
                            case "":
                                var v50 uint64 = 584
                                _ = v50
                            case "hello":
                                v1 = v1
                            }
                            if true {
                                v3 -= (-(v3))
                                fmt.Println(v43, fi47)
                                v51 := GenS3{F0: 2147483648, F1: (("\n\t" + "") + "")}
                                _ = v51
                            }
                            _ = fi47
                        }
                        _ = v45
                        _ = v46
                    }
                    var v52 int8 = -35
                    _ = fi44
                    _ = v52
                }
                fmt.Println(v1, v2, v1)
                fmt.Println(v2, v1)
            } else {
                v53 := GenS7{F0: (false && (true && false)), F1: -128}
                _ = v53
            }
        }
        _ = v43
    } else {
        fmt.Println(v1, v3)
        var v54 int64 = -627
        switch "hello" {
        case "a":
            v32 = v32
        }
        _ = v54
    }
    v55 := GenS5{F0: "", F1: 700, F2: 536}
    var v56 float32 = (3.14159 + v1)
    v57 := []int{128, -236, genFn0((false || false), -1, 2), 127}
    if true {
        v32 = v32
    }
    for fi58 := 0; fi58 < 1; fi58++ {
        var v59 float32 = (-2.25 / (1.5 / ((float32(-2.25) / (((1e10 / 1.0) + (v56 - v1)) - (float32(((0.0 * float64(v56)) * (1e10 / (1.5 / 1e10)))) + ((v1 + float32((0.0 - 0.0))) + (v1 * float32(-2.25)))))) / v56)))
        var v60 float64 = 3.14159
        _ = fi58
        _ = v59
        _ = v60
    }
    var v61 int16 = (v3 - (v3 * v3))
    v62 := GenS8{F0: 2147483648, F1: 0}
    switch "\n\t" {
    case "\n\t":
        if (9223372036854775807 >= 409) {
            if true {
                var v63 float32 = v1
                switch 0 {
                case 42:
                    for fi64 := 0; fi64 < 4; fi64++ {
                        v61 = v61
                        var v65 int64 = -9223372036854775808
                        _ = fi64
                        _ = v65
                    }
                case -1:
                    var v66 int64 = 0
                    _ = v66
                case 192:
                    for fi67 := 0; fi67 < 9; fi67++ {
                        v63 = (float32((1e10 * 0.1)) / ((0.0 * v56) * (3.14159 + 0.0)))
                        v1 -= (v56 * v1)
                        v61 = (v3 | v3)
                        _ = fi67
                    }
                }
                _ = v63
            } else {
                if (255 <= 0) {
                    var v68 int8 = -52
                    _ = v68
                }
            }
            if len(v57) > 0 {
                v57[0] = 0
            }
        } else {
            v32 = (v32 ^ v32)
            if (!((2 > 2147483647) && false)) {
                if (!(1.5 >= ((1.5 + 1e10) / 1.0))) {
                    var v69 int32 = -1
                    _ = v69
                }
            } else {
                v70 := GenS9{F0: 18446744073709551615}
                _ = v70
            }
            switch 128 {
            case 742:
                if false {
                    switch 128 {
                    case -152:
                        var v71 uint8 = 0
                        _ = v71
                    }
                } else {
                    if ((("\n\t" + ("a" + "hello")) >= "") && true) {
                        if ((!true) || true) {
                            fmt.Println(v61)
                            if (!((3.14159 + (-2.25 + 0.1)) != ((((float64(v56) * 1e10) * float64(float32(-2.25))) / 1e10) + 0.1))) {
                                fmt.Println(v32, v61)
                                var v72 int64 = 42
                                _ = len(v57)
                                _ = v72
                            }
                        }
                        for fi73 := 0; fi73 < 16; fi73++ {
                            v61 += (^(v61))
                            _ = fi73
                        }
                    } else {
                        var v74 uint8 = 42
                        v56 = (1e10 - v1)
                        if false {
                            var v75 bool = (127 < -167)
                            v2 -= v2
                            v76 := GenS6{F0: 2147483648}
                            _ = v75
                            _ = v76
                        }
                        _ = v74
                    }
                    fmt.Println(v61, v2, v32)
                }
            case 0:
                v3 -= 1
            case -1:
                _ = len(v57)
            default:
                v2 = (^(v2))
            }
        }
    case "a":
        fmt.Println(v2)
    case "foo bar":
        for fi77 := 0; fi77 < 14; fi77++ {
            if false {
                for fi78 := 0; fi78 < 8; fi78++ {
                    v79 := GenS4{F0: fi78}
                    v1 += ((v1 + (((v1 / (v56 / 0.1)) / 3.14159) - 0.0)) / (3.14159 / v1))
                    if true {
                        v3 = (v61 >> ((42) & 63))
                        switch (fi78 * fi77) {
                        case -2147483648:
                            fmt.Println(fi77, v56, v32)
                        }
                    }
                    _ = fi78
                    _ = v79
                }
                fmt.Println(v61, v61)
            }
            _ = fi77
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v55
    _ = v56
    _ = v57
    _ = v61
    _ = v62
}
