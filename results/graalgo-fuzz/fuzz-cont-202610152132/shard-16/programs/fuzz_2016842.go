package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 int; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 bool; F2 uint16 }
func genFn0(p0 bool, p1 int) uint32 {
    fmt.Println(p1, p1)
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{-2147483648, 32767, (^(fi7))}
            _ = fi8
            _ = v9
        }
        if ((-101 < 1) || p0) {
            p0 = p0
            for fi10 := 0; fi10 < 2; fi10++ {
                v11 := GenS1{F0: 723, F1: -2147483648, F2: 42}
                switch p1 {
                case 1:
                    if p0 {
                        if true {
                            if p0 {
                                v12 := []int{(p1 ^ p1), (fi7 << ((65535) & 63))}
                                _ = v12
                            } else {
                                if false {
                                    if p0 {
                                        p1 = 2147483647
                                    } else {
                                        var v13 float32 = ((1.5 + 0.1) / (1.5 / ((1.5 - (3.14159 - float32(float64(3.14159)))) - float32(1e10))))
                                        _ = v13
                                    }
                                }
                                var v14 uint32 = 1
                                _ = v14
                            }
                            fmt.Println(fi10, p1, p0)
                        }
                        p1 = fi10
                        var v15 int = (fi7 << ((18446744073709551615) & 63))
                        v16 := []int{}
                        _ = v15
                        _ = v16
                    } else {
                        fmt.Println(fi7, fi7, fi10)
                    }
                }
                _ = fi10
                _ = v11
            }
            switch "a" {
            case "hello":
                p0 = true
            case "":
                var v17 uint = 2
                _ = v17
            case "\n\t":
                p1 -= 32768
            default:
                for fi18 := 0; fi18 < 15; fi18++ {
                    p0 = true
                    v19 := []int{}
                    var v20 uint8 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
        _ = fi7
    }
    p0 = false
    _ = p0
    _ = p1
    return 65535
}
func genFn1() int {
    var v21 string = "foo bar"
    var v22 uint = 2
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -170
    v3 := []string{}
    var v4 int32 = -2147483648
    var v5 uint64 = 42
    v6 := make(map[string]uint32)
    var v23 int64 = (v2 ^ (v2 - 42))
    var v24 int8 = 100
    if true {
        if true {
            v2 = (v2 << ((18446744073709551615) & 63))
            if true {
                v4 -= 2147483647
                v24 -= v24
            }
            switch "foo bar" {
            case "hello":
                var v25 uint8 = 36
                _ = v25
            }
        } else {
            v2 = v2
            if (true && ((3.14159 != -2.25) || ((65535 != 205) && (false || (926 >= 0))))) {
                if false {
                    var v26 uint32 = 976
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 5; fi27++ {
                        if ((v24 + v24) < (v24 | v24)) {
                            switch "a" {
                            case "\n\t":
                                if (1000 == -1) {
                                    if true {
                                        fmt.Println(v5, v2)
                                        v28 := GenS2{F0: "a", F1: fi27, F2: 4294967295}
                                        _ = v28
                                    } else {
                                        var v29 int16 = 865
                                        var v30 int16 = v29
                                        _ = v29
                                        _ = v30
                                    }
                                    for fi31 := 0; fi31 < 14; fi31++ {
                                        if ((((("hello" + "a") + "\n\t") >= "a") && (true && (641 > 32767))) && true) {
                                            v24 -= (v24 >> ((4294967295) & 63))
                                            if false {
                                                fmt.Println(fi27, fi31)
                                                var v32 float32 = 0.1
                                                v33 := GenS1{F0: (v4 | v4), F1: (v4 >> ((42) & 63)), F2: 0}
                                                v34 := GenS3{F0: (v4 | (v4 * 0))}
                                                v5 = v5
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                        } else {
                                            var v35 int16 = 1000
                                            fmt.Println(v4, v24)
                                            _ = v35
                                        }
                                        var v36 int8 = (v24 + (-(v24)))
                                        var v37 uint = 65535
                                        v4 = v4
                                        if (!false) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if true {
                                                    fmt.Println(v37)
                                                    v38 := GenS3{F0: v4}
                                                    v23 -= 1
                                                    _ = v38
                                                } else {
                                                    v39 := GenS1{F0: 108, F1: (v4 & (v4 % ((-1) | 1))), F2: 32767}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    v40 := GenS4{F0: 42, F1: (v5 << ((v37) & 63)), F2: v36}
                                                    var v41 int = genFn1()
                                                    _ = v39
                                                    _ = v40
                                                    _ = v41
                                                }
                                            case "a":
                                                var v42 float32 = 0.0
                                                v24 += -71
                                                _ = v42
                                            case "":
                                                v43 := GenS5{F0: 111, F1: 816}
                                                _ = v43
                                            default:
                                                if false {
                                                    var v44 uint64 = 404
                                                    var v45 float32 = (((-2.25 / (1e10 / 1.5)) / ((1.5 - (1e10 + 0.0)) - 1e10)) + ((1e10 - (0.1 * 1.5)) - (3.14159 * 0.0)))
                                                    _ = v44
                                                    _ = v45
                                                }
                                                v36 = v36
                                            }
                                            var v46 int8 = v24
                                            _ = v46
                                        } else {
                                            switch ("hello" + "foo bar") {
                                            case "a":
                                                var v47 uint = (^(v37))
                                                _ = v47
                                            case "":
                                                if true {
                                                    switch ("\n\t" + (("foo bar" + "a") + "a")) {
                                                    case "a":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "":
                                                        var v48 uint8 = 255
                                                        _ = v48
                                                    }
                                                    v49 := GenS6{F0: false, F1: 1, F2: 127}
                                                    _ = v49
                                                } else {
                                                    v50 := GenS2{F0: "foo bar", F1: (^(fi31)), F2: v5}
                                                    _ = v50
                                                }
                                            default:
                                                v5 = v5
                                            }
                                            var v51 int = fi31
                                            v52 := GenS1{F0: -1, F1: v4, F2: 32767}
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = fi31
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            case "a":
                                v24 -= (v24 & v24)
                            }
                            fmt.Println(v24, v23, fi27)
                            if (false || ((((float32(0.1) / 0.1) / (1e10 * 3.14159)) + float32(3.14159)) == float32((1.5 + 1e10)))) {
                                if (42 >= 255) {
                                    for fi53 := 0; fi53 < 7; fi53++ {
                                        fmt.Println(v5, v2, fi53)
                                        _ = fi53
                                    }
                                    var v54 uint16 = 322
                                    var v55 int8 = (v24 / ((-128) | 1))
                                    _ = v54
                                    _ = v55
                                }
                            }
                        } else {
                            fmt.Println(v23)
                            switch -247 {
                            case 0:
                                fmt.Println(v24)
                            case -1:
                                for fi56 := 0; fi56 < 9; fi56++ {
                                    v24 -= v24
                                    v23 = -329
                                    _ = len(v3)
                                    _ = fi56
                                }
                            case 9223372036854775807:
                                v5 = v5
                            }
                        }
                        fmt.Println(v2, fi27, v24)
                        _ = fi27
                    }
                    v3 = append(v3, "hello")
                }
                v4 += -445
            }
        }
    }
    v1[127] = (18446744073709551615 <= 42)
    var v57 uint16 = 94
    var v58 uint32 = 65535
    var v59 int = 1
    v2 -= v2
    v3 = append(v3, ("hello" + ((("" + (("\n\t" + "foo bar") + ("\n\t" + ("hello" + ("" + (("a" + "a") + "a")))))) + ("" + "hello")) + "hello")))
    v60 := v6[""]
    if (1 != 616) {
        switch "hello" {
        case "hello":
            v6["hello"] = (^(v58))
        default:
            v4 = -967
        }
    } else {
        switch "foo bar" {
        case "hello":
            if (!(-1 != v4)) {
                if (v4 >= v4) {
                    v2 = (v2 * v23)
                } else {
                    var v61 uint32 = (v60 * v60)
                    for fi62 := 0; fi62 < 10; fi62++ {
                        var v63 bool = true
                        v64 := GenS5{F0: (v5 - (v5 >> ((255) & 63))), F1: int16(v24)}
                        for fi65 := 0; fi65 < 1; fi65++ {
                            v1[fi62] = v63
                            _ = fi65
                        }
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                    _ = v61
                }
                var v66 bool = true
                _ = v66
            }
        }
        switch (v59 * v59) {
        case -360:
            switch (v59 | v59) {
            case 735:
                switch (v59 + v59) {
                case 127:
                    if (42 < 42) {
                        for fi67 := 0; fi67 < 6; fi67++ {
                            v60 += (v58 + v60)
                            v68 := GenS2{F0: ("hello" + "hello"), F1: 872, F2: (^(v5))}
                            switch "" {
                            case "a":
                                fmt.Println(v2, fi67, v57)
                            default:
                                v69 := GenS5{F0: 0, F1: -1}
                                _ = v69
                            }
                            _ = fi67
                            _ = v68
                        }
                        if (v57 < (v57 >> ((1) & 63))) {
                            switch (^(v59)) {
                            case 4294967295:
                                v58 = uint32(v57)
                            case 0:
                                fmt.Println(v23, v24)
                            }
                            switch 669 {
                            case -1:
                                v70 := GenS7{F0: (v59 % ((v59) | 1))}
                                _ = v70
                            }
                            _ = len(v3)
                        }
                    }
                case 2147483647:
                    v2 += -408
                }
            case -665:
                fmt.Println(v24)
            default:
                switch (-(v59)) {
                case -60:
                    for fi71 := 0; fi71 < 3; fi71++ {
                        for fi72 := 0; fi72 < 10; fi72++ {
                            var v73 float64 = 1e10
                            _ = fi72
                            _ = v73
                        }
                        var v74 float64 = (float64((1.5 * 3.14159)) * (0.1 + float64((float32(float64(0.0)) * (1e10 / -2.25)))))
                        _ = fi71
                        _ = v74
                    }
                case -2147483648:
                    v3 = append(v3, "foo bar")
                case 128:
                    _ = len(v1)
                }
            }
        case 0:
            switch ("hello" + (("\n\t" + "foo bar") + ("" + ("foo bar" + "")))) {
            case "":
                fmt.Println(v60, v5, v57)
            case "foo bar":
                v24 += v24
            case "a":
                v24 += v24
            default:
                var v75 int16 = -32768
                _ = v75
            }
        case 2147483647:
            var v76 float32 = (((3.14159 / 0.1) + (0.1 - float32(1e10))) + float32(-2.25))
            _ = v76
        default:
            v60 += 2
        }
    }
    var v77 float32 = 1e10
    var v78 int16 = 42
    fmt.Println(v5, v78)
    v3 = append(v3, "hello")
    if false {
        v79 := GenS3{F0: v4}
        if (((v59 << ((0) & 63)) < int(v57)) && (false || ((v77 * (0.0 * v77)) <= (v77 - (3.14159 - v77))))) {
            switch (v59 | (v59 | genFn1())) {
            case 0:
                v3 = append(v3, ("\n\t" + "hello"))
            case -1:
                v3 = append(v3, "foo bar")
            case 32768:
                v3 = append(v3, "a")
            }
            v60 = genFn0((0 > 1), v59)
            if (127 < (v24 & v24)) {
                var v80 int64 = -1
                _ = len(v1)
                _ = v80
            } else {
                fmt.Println(v58)
            }
        }
        if ((v24 ^ (v24 + v24)) < (^(v24))) {
            _ = len(v1)
            var v81 uint16 = v57
            v24 -= 42
            _ = v81
        } else {
            for fi82 := 0; fi82 < 3; fi82++ {
                if (float32(0.0) == (3.14159 - v77)) {
                    switch "a" {
                    case "\n\t":
                        if false {
                            for fi83 := 0; fi83 < 17; fi83++ {
                                fmt.Println(v2, v23)
                                fmt.Println(v78, v57, v60)
                                var v84 int = genFn1()
                                _ = fi83
                                _ = v84
                            }
                        }
                    case "hello":
                        for fi85 := 0; fi85 < 12; fi85++ {
                            var v86 bool = ((v78 | (v78 - int16(v4))) == (v78 | v78))
                            v87 := GenS6{F0: (v86 && ((-(fi82)) >= fi82)), F1: 32767, F2: 1}
                            var v88 float64 = (1e10 * 0.0)
                            var v89 bool = v86
                            _ = fi85
                            _ = v86
                            _ = v87
                            _ = v88
                            _ = v89
                        }
                    case "":
                        v60 -= genFn0(false, genFn1())
                    default:
                        switch (-(v59)) {
                        case -806:
                            for fi90 := 0; fi90 < 7; fi90++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v60 += genFn0((!false), fi82)
                                var v91 uint16 = (v57 & (v57 / ((32767) | 1)))
                                _ = fi90
                                _ = v91
                            }
                        case -905:
                            v92 := GenS3{F0: int32(v58)}
                            _ = v92
                        case -1:
                            v57 = v57
                        }
                    }
                } else {
                    v93 := GenS8{F0: 2, F1: (v23 << ((4294967295) & 63))}
                    for fi94 := 0; fi94 < 6; fi94++ {
                        _ = len(v1)
                        _ = fi94
                    }
                    for fi95 := 0; fi95 < 3; fi95++ {
                        var v96 float32 = -2.25
                        var v97 int = fi95
                        _ = len(v3)
                        var v98 float64 = float64(float32(float64(float32(3.14159))))
                        if true {
                            if (1 != 1) {
                                if (3.14159 <= (float32(3.14159) + (v77 + v77))) {
                                    var v99 int64 = v2
                                    _ = v99
                                }
                                if len(v3) > 0 {
                                    v3[0] = "foo bar"
                                }
                            } else {
                                switch ("" + "hello") {
                                case "hello":
                                    v77 = 3.14159
                                default:
                                    v100 := GenS6{F0: (!false), F1: (^(v57)), F2: fi82}
                                    _ = v100
                                }
                                v101 := GenS9{F0: "foo bar"}
                                _ = v101
                            }
                            var v102 uint64 = (v5 * 42)
                            _ = v102
                        } else {
                            v77 = 3.14159
                        }
                        _ = fi95
                        _ = v96
                        _ = v97
                        _ = v98
                    }
                    _ = v93
                }
                if true {
                    fmt.Println(v24, v77)
                    v103 := GenS4{F0: (v5 ^ (v5 * 0)), F1: v5, F2: (v24 % ((v24) | 1))}
                    v104 := GenS10{F0: (v4 * (v4 ^ v4)), F1: (0 < 238), F2: 32767}
                    _ = v103
                    _ = v104
                }
                v1[fi82] = true
                _ = fi82
            }
        }
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v77
    _ = v78
}
