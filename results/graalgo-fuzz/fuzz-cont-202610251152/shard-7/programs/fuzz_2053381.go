package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint8; F2 int }
type GenS11 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int8; F1 bool; F2 uint8 }
type GenS14 struct { F0 int; F1 uint64; F2 string }
func genFn0(p0 uint64) uint {
    var v6 float32 = (0.0 / (0.0 + 1.5))
    v7 := []int{127, 32767, 1}
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 int64, p2 string) uint16 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = 254
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    fmt.Println(p0, p0)
    var v10 string = "foo bar"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 42
    var v3 int64 = (^(v2))
    var v4 uint16 = 42
    v5 := make(map[string]bool)
    if false {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 string = ("foo bar" + ("" + "foo bar"))
            _ = fi11
            _ = v12
        }
        var v13 uint = 255
        _ = v13
    }
    v14 := []int{-816, 32767, 0, 32767}
    var v15 bool = true
    fmt.Println(v3)
    for fi16 := 0; fi16 < 6; fi16++ {
        delete(v5, "")
        _ = fi16
    }
    if false {
        v4 -= (v4 + (v4 | (^(v4))))
    }
    v2 += (v2 & (v2 ^ (^(v2))))
    switch ("a" + "hello") {
    case "\n\t":
        if v15 {
            v17 := GenS1{F0: 0, F1: ((-(v2)) <= 42)}
            v2 = v3
            v18 := GenS1{F0: 42, F1: (4294967295 >= 152)}
            _ = v17
            _ = v18
        } else {
            switch "a" {
            case "a":
                if v15 {
                    _ = len(v5)
                    fmt.Println(v15, v2)
                }
            case "foo bar":
                v19 := GenS1{F0: 42, F1: (42 > 32)}
                _ = v19
            case "":
                for fi20 := 0; fi20 < 15; fi20++ {
                    var v21 uint32 = 0
                    _ = fi20
                    _ = v21
                }
            }
            v4 = v4
            for fi22 := 0; fi22 < 7; fi22++ {
                fmt.Println(v3, v2, v2)
                _ = fi22
            }
        }
    case "hello":
        switch "" {
        case "":
            if v15 {
                if true {
                    var v23 float32 = 0.1
                    _ = v23
                }
                fmt.Println(v3)
                v15 = (1e10 != 3.14159)
            }
        case "hello":
            if true {
                v24 := GenS1{F0: 42, F1: v15}
                v3 += 0
                var v25 string = ("hello" + "foo bar")
                _ = v24
                _ = v25
            } else {
                var v26 float32 = float32((float64(-2.25) / 1.0))
                var v27 int32 = 1
                v15 = false
                _ = v26
                _ = v27
            }
        }
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 int = 606
        _ = fi28
        _ = v29
    }
    var v30 uint16 = (v4 << ((genFn0(417)) & 63))
    fmt.Println(v15, v3, v4)
    v15 = v15
    for fi31 := 0; fi31 < 3; fi31++ {
        switch (("\n\t" + ("" + "")) + "a") {
        case "":
            var v32 uint = 541
            _ = v32
        case "foo bar":
            v2 += 9223372036854775807
        case "\n\t":
            if v15 {
                if ((!false) && (((v30 % ((v4) | 1)) != (v30 << ((genFn0(2)) & 63))) || v15)) {
                    if v15 {
                        for fi33 := 0; fi33 < 4; fi33++ {
                            if v15 {
                                fmt.Println(fi33, v15, v30)
                            }
                            switch "foo bar" {
                            case "foo bar":
                                switch fi31 {
                                case 42:
                                    v15 = false
                                case 9223372036854775807:
                                    var v34 uint16 = (v4 + (v4 & 1))
                                    _ = v34
                                case -143:
                                    var v35 bool = v15
                                    var v36 int32 = int32(v3)
                                    v4 = genFn1(false, v3, ("foo bar" + "foo bar"))
                                    _ = v35
                                    _ = v36
                                }
                            }
                            v37 := GenS2{F0: v3}
                            _ = fi33
                            _ = v37
                        }
                        var v38 int32 = 1
                        v1[fi31] = (286 >= 2147483647)
                        _ = v38
                    } else {
                        var v39 uint32 = 65535
                        if (-609 <= 0) {
                            var v40 uint16 = v4
                            _ = v40
                        }
                        _ = v39
                    }
                    if v15 {
                        if v15 {
                            v4 -= 1
                        } else {
                            if len(v14) > 0 {
                                v14[0] = 1
                            }
                            _ = len(v14)
                        }
                    } else {
                        v4 -= v30
                        v30 -= (v30 % (((^(v30))) | 1))
                        _ = len(v14)
                    }
                } else {
                    _ = len(v1)
                    for fi41 := 0; fi41 < 20; fi41++ {
                        switch "foo bar" {
                        case "foo bar":
                            if true {
                                v42 := GenS3{F0: (v3 >> ((18446744073709551615) & 63)), F1: v15}
                                v43 := GenS1{F0: 100, F1: v15}
                                v30 -= v30
                                _ = v42
                                _ = v43
                            } else {
                                if false {
                                    switch "" {
                                    case "a":
                                        v44 := GenS4{F0: 255, F1: (^(v4)), F2: 91}
                                        _ = v44
                                    case "\n\t":
                                        switch (fi31 % ((fi41) | 1)) {
                                        case -761:
                                            switch ("\n\t" + "") {
                                            case "hello":
                                                v45 := GenS2{F0: (v3 | (v2 % (((v2 - -52)) | 1)))}
                                                _ = v45
                                            }
                                        case -9223372036854775808:
                                            switch ("" + "") {
                                            case "\n\t":
                                                _ = len(v14)
                                            default:
                                                if (((-2.25 - 0.1) - 0.1) > ((0.1 - (1.5 - 1.5)) + (float64(0.0) / (3.14159 + 0.0)))) {
                                                    var v46 float32 = (1.5 * -2.25)
                                                    v3 -= (v3 >> ((484) & 63))
                                                    v5["a"] = (1 > 742)
                                                    _ = v46
                                                }
                                            }
                                        case 817:
                                            if v15 {
                                                if false {
                                                    var v47 int16 = 1000
                                                    var v48 int8 = 100
                                                    switch ("hello" + ("hello" + "foo bar")) {
                                                    case "a":
                                                        if (v15 && (0 == (v30 ^ v30))) {
                                                            v49 := GenS5{F0: 0, F1: 18446744073709551615, F2: 42}
                                                            v50 := GenS2{F0: (-(v2))}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v47 = v47
                                                        }
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                }
                                                var v51 uint16 = (v4 + (v30 * (^(v4))))
                                                _ = v51
                                            }
                                        }
                                    case "foo bar":
                                        var v52 float64 = 0.0
                                        _ = v52
                                    }
                                } else {
                                    v53 := GenS3{F0: v2, F1: ((^(fi31)) != fi31)}
                                    v15 = v15
                                    v30 += v30
                                    _ = v53
                                }
                            }
                        case "a":
                            fmt.Println(v15, v4, v4)
                        case "":
                            if (v3 <= 483) {
                                var v54 uint16 = genFn1(false, 46, (("\n\t" + ("a" + "\n\t")) + ""))
                                switch "a" {
                                case "hello":
                                    var v55 string = "a"
                                    _ = v55
                                case "foo bar":
                                    v30 = (v54 % (((v54 << ((255) & 63))) | 1))
                                case "\n\t":
                                    if false {
                                        v4 = (v4 << ((18446744073709551615) & 63))
                                        v56 := GenS4{F0: 42, F1: (v4 + v54), F2: -128}
                                        switch int(v3) {
                                        case 42:
                                            var v57 bool = (0.1 != -2.25)
                                            _ = v57
                                        case 32768:
                                            switch (fi41 + (fi31 | 42)) {
                                            case 0:
                                                if (((!((!v15) && true)) && (-2147483648 == 42)) || v15) {
                                                    if ((128 < 42) || (!(v15 && (!v15)))) {
                                                        if v15 {
                                                            v58 := GenS2{F0: v2}
                                                            v59 := GenS3{F0: 0, F1: (!true)}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        v60 := GenS6{F0: (v54 - genFn1((v15 || false), -1, "foo bar"))}
                                                        _ = v60
                                                    } else {
                                                        v2 = -362
                                                        var v61 uint64 = 4294967295
                                                        _ = v61
                                                    }
                                                    _ = len(v14)
                                                } else {
                                                    v62 := GenS2{F0: (v3 ^ 9223372036854775807)}
                                                    if (true && false) {
                                                        v63 := v1[630]
                                                        v4 -= 190
                                                        if false {
                                                            v3 = (v2 >> ((4294967295) & 63))
                                                        } else {
                                                            v15 = v63
                                                            var v64 int64 = 1
                                                            _ = v64
                                                        }
                                                        _ = v63
                                                    } else {
                                                        v65 := GenS1{F0: -78, F1: v15}
                                                        if true {
                                                            v66 := GenS4{F0: 128, F1: (v4 >> ((135) & 63)), F2: 71}
                                                            _ = v66
                                                        }
                                                        if v15 {
                                                            var v67 float32 = ((((1.5 - ((3.14159 * 1.5) * 3.14159)) / 3.14159) / -2.25) / float32(1e10))
                                                            var v68 uint16 = (^(v54))
                                                            _ = v67
                                                            _ = v68
                                                        } else {
                                                            if len(v14) > 0 {
                                                                v14[0] = (fi41 * fi31)
                                                            }
                                                            if (!v15) {
                                                                if (!false) {
                                                                    var v69 uint16 = genFn1((!(183 >= 64)), v2, ("\n\t" + (("\n\t" + "a") + "hello")))
                                                                    if len(v14) > 0 {
                                                                        v14[0] = (fi41 ^ (fi31 - fi41))
                                                                    }
                                                                    v70 := GenS1{F0: 1, F1: v15}
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    fmt.Println(v15, v30)
                                                                    var v71 float32 = 3.14159
                                                                    v71 += 0.1
                                                                    v4 -= (v4 | 42)
                                                                    if len(v14) >= 2 {
                                                                        v14 = v14[1:len(v14)]
                                                                    }
                                                                    _ = v71
                                                                }
                                                                switch ((("hello" + "a") + ("" + "a")) + "") {
                                                                case "":
                                                                    var v72 bool = v15
                                                                    _ = v72
                                                                default:
                                                                    fmt.Println(v3, v3, v2)
                                                                }
                                                                v73 := GenS6{F0: (v4 * 80)}
                                                                _ = len(v14)
                                                                var v74 float32 = -2.25
                                                                v54 -= v54
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                fmt.Println(fi41, fi31, v30)
                                                                switch (fi41 << ((4294967295) & 63)) {
                                                                case 32767:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v15 = (!(!(!(v15 || (v30 <= genFn1(true, -76, "a"))))))
                                                                    }
                                                                case 707:
                                                                    fmt.Println(v15, v54)
                                                                }
                                                                v2 = v2
                                                            }
                                                            if (v15 || (18446744073709551615 != 2)) {
                                                                var v75 uint16 = genFn1(v15, (v2 << ((2) & 63)), "")
                                                                var v76 uint32 = 2147483647
                                                                v75 = v75
                                                                _ = v75
                                                                _ = v76
                                                            } else {
                                                                if true {
                                                                    v3 -= v3
                                                                } else {
                                                                    v77 := GenS7{F0: uint64(v30), F1: "a"}
                                                                    v78 := GenS3{F0: 42, F1: v15}
                                                                    v79 := v1[fi41]
                                                                    var v80 uint8 = 255
                                                                    _ = v77
                                                                    _ = v78
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                v81 := GenS8{F0: 0, F1: 719}
                                                                if false {
                                                                    v82 := GenS5{F0: 968, F1: 2, F2: 2}
                                                                    _ = len(v14)
                                                                    _ = v82
                                                                } else {
                                                                    if true {
                                                                        v54 = v30
                                                                        v83 := GenS9{F0: int32(fi31), F1: (v4 & genFn1(false, v3, ""))}
                                                                        _ = v83
                                                                    }
                                                                }
                                                                _ = v81
                                                            }
                                                        }
                                                        var v84 int64 = v2
                                                        switch (("\n\t" + "foo bar") + "\n\t") {
                                                        case "hello":
                                                            var v85 bool = true
                                                            _ = v85
                                                        default:
                                                            v5["hello"] = true
                                                        }
                                                        var v86 int16 = -1
                                                        v84 = v84
                                                        _ = v65
                                                        _ = v84
                                                        _ = v86
                                                    }
                                                    _ = v62
                                                }
                                            }
                                        case 127:
                                            switch 643 {
                                            case 32767:
                                                var v87 string = ("a" + ("foo bar" + "hello"))
                                                _ = v87
                                            case 696:
                                                v2 = (-(v3))
                                            default:
                                                v88 := GenS10{F0: v2, F1: 1, F2: fi31}
                                                _ = v88
                                            }
                                        }
                                        _ = v56
                                    }
                                default:
                                    v15 = v15
                                }
                                v89 := GenS11{F0: uint8(v30), F1: (v30 - (v4 << ((942) & 63))), F2: genFn1(v15, 42, "foo bar")}
                                _ = v54
                                _ = v89
                            } else {
                                if v15 {
                                    v90 := GenS12{F0: -253}
                                    _ = v90
                                }
                                _ = len(v5)
                                v30 = (v30 * (v4 / ((v30) | 1)))
                            }
                        default:
                            if v15 {
                                v15 = v15
                                v2 = (v3 << ((4294967295) & 63))
                                v91 := v5["a"]
                                _ = v91
                            } else {
                                v92 := GenS9{F0: -805, F1: genFn1(v15, v3, "\n\t")}
                                _ = v92
                            }
                            var v93 uint16 = genFn1((!v15), v2, "")
                            switch "" {
                            case "a":
                                var v94 string = "a"
                                _ = v94
                            case "":
                                v95 := GenS10{F0: (^(v3)), F1: 128, F2: fi41}
                                _ = v95
                            case "hello":
                                v30 = (^(v93))
                            }
                            _ = v93
                        }
                        var v96 uint = 42
                        _ = fi41
                        _ = v96
                    }
                    _ = len(v14)
                }
            }
        }
        if v15 {
            v97 := GenS13{F0: 0, F1: (!((("a" + "\n\t") + (((("foo bar" + "") + "foo bar") + (((("" + "foo bar") + "") + "\n\t") + "")) + ("\n\t" + "\n\t"))) != "\n\t")), F2: 228}
            _ = v97
        } else {
            if v15 {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            for fi98 := 0; fi98 < 13; fi98++ {
                switch "hello" {
                case "\n\t":
                    v14 = append(v14, 0)
                case "foo bar":
                    v99 := GenS8{F0: 2147483647, F1: 42}
                    _ = v99
                }
                v4 = 32767
                _ = fi98
            }
        }
        _ = fi31
    }
    v15 = v15
    if v15 {
        v100 := GenS14{F0: 536, F1: 2, F2: ""}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v101 int32 = 0
        _ = v100
        _ = v101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
