package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint8; F1 int32; F2 uint }
func genFn0(p0 int32, p1 int32) int8 {
    for fi6 := 0; fi6 < 14; fi6++ {
        if ((0 <= fi6) || (!(!(2147483647 < 42)))) {
            fmt.Println(p1, p1, p1)
            if true {
                v7 := GenS1{F0: 42, F1: 2}
                _ = v7
            } else {
                var v8 uint32 = 850
                _ = v8
            }
            var v9 uint16 = 66
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            switch "a" {
            case "foo bar":
                v11 := GenS1{F0: 153, F1: 486}
                _ = v11
            case "hello":
                var v12 uint16 = 300
                _ = v12
            case "a":
                if false {
                    v13 := []int{}
                    _ = v13
                }
            }
            p0 += (p0 >> ((65535) & 63))
            if (42 >= 1) {
                p1 += p0
                p0 = -1
                var v14 int16 = -1
                _ = v14
            } else {
                fmt.Println(fi6)
            }
            _ = fi10
        }
        _ = fi6
    }
    var v15 int8 = -35
    _ = p0
    _ = p1
    _ = v15
    return (v15 >> ((18446744073709551615) & 63))
}
func genFn1(p0 bool, p1 string, p2 int64) int64 {
    v16 := []int{32768}
    v17 := GenS2{F0: 42, F1: 9, F2: 255}
    var v18 int64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return (p2 * (p2 >> ((42) & 63)))
}

func main() {
    var v1 float32 = (3.14159 / 1.0)
    var v2 int16 = 0
    var v3 uint16 = 2
    v4 := make(map[string]int)
    var v5 uint8 = 42
    v3 = 2
    v19 := GenS3{F0: 335, F1: 0}
    v2 += 1000
    _ = len(v4)
    var v20 uint64 = 569
    v3 += 42
    var v21 string = "foo bar"
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 uint8 = (v5 - v5)
            v25 := GenS4{F0: 4294967295}
            var v26 float32 = -2.25
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        var v27 float32 = v1
        switch (^(fi22)) {
        case 0:
            switch v21 {
            case "\n\t":
                var v28 int16 = (v2 + v2)
                _ = v28
            case "":
                if (false || true) {
                    if (42 == 2147483647) {
                        v27 -= (v1 - float32(-2.25))
                        switch (fi22 / ((32768) | 1)) {
                        case 32768:
                            if (0 != 127) {
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    v30 := GenS5{F0: true, F1: (^(fi22)), F2: (fi22 << ((65535) & 63))}
                                    fmt.Println(v3, v20, v27)
                                    _ = fi29
                                    _ = v30
                                }
                                var v31 int64 = -9223372036854775808
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 4; fi32++ {
                                    switch 32768 {
                                    case -78:
                                        if false {
                                            var v33 string = (v21 + v21)
                                            var v34 uint8 = (v5 | v5)
                                            fmt.Println(fi22)
                                            _ = v33
                                            _ = v34
                                        }
                                    case -573:
                                        v3 = v3
                                    case 4294967295:
                                        v35 := []int{42, fi22, fi32}
                                        _ = v35
                                    }
                                    v27 += -2.25
                                    _ = fi32
                                }
                            }
                        }
                    }
                    if (32767 <= (v2 * (^(v2)))) {
                        for fi36 := 0; fi36 < 17; fi36++ {
                            _ = len(v4)
                            var v37 float32 = -2.25
                            if true {
                                var v38 bool = true
                                var v39 bool = (false && (v20 != v20))
                                _ = v38
                                _ = v39
                            }
                            if true {
                                if false {
                                    v40 := v4[v21]
                                    v41 := GenS6{F0: 32767}
                                    fmt.Println(v21, v27)
                                    _ = v40
                                    _ = v41
                                } else {
                                    v42 := GenS7{F0: 37, F1: (v5 ^ v5), F2: 99}
                                    v43 := GenS8{F0: (v5 >> ((18446744073709551615) & 63))}
                                    _ = v42
                                    _ = v43
                                }
                            }
                            delete(v4, ("\n\t" + v21))
                            _ = fi36
                            _ = v37
                        }
                        var v44 int16 = 32767
                        _ = v44
                    } else {
                        switch fi22 {
                        case 9223372036854775807:
                            if (3.14159 != -2.25) {
                                switch ("a" + v21) {
                                case "":
                                    v45 := []int{}
                                    _ = v45
                                case "a":
                                    v20 = (^(v20))
                                default:
                                    v20 += (v20 % (((^(v20))) | 1))
                                }
                                switch (fi22 & (fi22 / ((fi22) | 1))) {
                                case 32768:
                                    for fi46 := 0; fi46 < 3; fi46++ {
                                        if (!(255 <= 255)) {
                                            fmt.Println(v20)
                                        } else {
                                            if false {
                                                fmt.Println(fi22, v20)
                                                switch (fi46 * (-(fi46))) {
                                                case 741:
                                                    var v47 int64 = 1
                                                    _ = v47
                                                case -9223372036854775808:
                                                    fmt.Println(v21, fi22, fi22)
                                                default:
                                                    if (true && false) {
                                                        v48 := GenS5{F0: true, F1: (fi22 & (fi46 * 645)), F2: fi46}
                                                        var v49 int64 = -9223372036854775808
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        v20 = (v20 + v20)
                                                        var v50 int16 = (v2 ^ v2)
                                                        _ = v50
                                                    }
                                                }
                                                switch v21 {
                                                case "foo bar":
                                                    v5 = (v5 & 44)
                                                case "hello":
                                                    v5 -= (v5 | v5)
                                                    var v51 int64 = -1
                                                    if true {
                                                        v21 = "a"
                                                        switch v21 {
                                                        case "a":
                                                            var v52 uint64 = 42
                                                            _ = v52
                                                        default:
                                                            switch ("" + v21) {
                                                            case "\n\t":
                                                                if ((0.1 <= v27) && false) {
                                                                    var v53 int16 = (v2 * v2)
                                                                    fmt.Println(v51)
                                                                    v21 = ""
                                                                    var v54 int16 = (v53 / (((^(v53))) | 1))
                                                                    v55 := v4[((v21 + (v21 + (v21 + ""))) + (("hello" + (v21 + v21)) + v21))]
                                                                    var v56 uint32 = 42
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    if true {
                                                                        var v57 int16 = 1000
                                                                        _ = v57
                                                                    } else {
                                                                        v51 = genFn1((-46 > -128), (v21 + "foo bar"), v51)
                                                                    }
                                                                    _ = len(v4)
                                                                    v58 := GenS8{F0: v5}
                                                                    v3 += 42
                                                                    v59 := []int{9223372036854775807, fi22, int(v3)}
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            case "foo bar":
                                                                var v60 bool = (v3 <= v3)
                                                                _ = v60
                                                            case "hello":
                                                                v1 += (float32(3.14159) - v1)
                                                            }
                                                            v20 = (v20 * (^(v20)))
                                                            v21 = "a"
                                                            switch "\n\t" {
                                                            case "a":
                                                                switch fi46 {
                                                                case -398:
                                                                    var v61 uint16 = v3
                                                                    _ = v61
                                                                default:
                                                                    fmt.Println(v5, v3)
                                                                }
                                                            }
                                                            if (true || (1 != 255)) {
                                                                v27 += v27
                                                                v1 += ((1e10 + 1e10) - v1)
                                                                v62 := GenS5{F0: ((v21 + v21) >= "a"), F1: fi22, F2: fi46}
                                                                _ = v62
                                                            } else {
                                                                if (!true) {
                                                                    v3 -= (v3 % ((1) | 1))
                                                                    v4[(("" + ("a" + "\n\t")) + (("hello" + ("hello" + v21)) + v21))] = (-(fi46))
                                                                    fmt.Println(v1, v3, v51)
                                                                    v2 -= -140
                                                                }
                                                            }
                                                            v5 += v5
                                                            switch (-(fi22)) {
                                                            case 32768:
                                                                var v63 string = v21
                                                                _ = v63
                                                            case 1:
                                                                v1 = 1.5
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v51, v27, v21)
                                                        v64 := GenS8{F0: 128}
                                                        _ = v64
                                                    }
                                                    _ = v51
                                                default:
                                                    if (!(128 >= fi22)) {
                                                        v27 -= (v27 - v1)
                                                        if (int(v20) > fi22) {
                                                            var v65 uint = 18446744073709551615
                                                            var v66 uint16 = v3
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    } else {
                                                        v67 := GenS3{F0: 0, F1: v20}
                                                        v68 := GenS3{F0: 0, F1: (v20 * (v20 << ((2) & 63)))}
                                                        switch ((((((v21 + "hello") + ("" + "foo bar")) + "hello") + v21) + (v21 + "foo bar")) + "foo bar") {
                                                        case "":
                                                            v1 -= float32(-2.25)
                                                        case "foo bar":
                                                            switch (-(fi22)) {
                                                            case 9223372036854775807:
                                                                v27 = v27
                                                            case -9223372036854775808:
                                                                _ = len(v4)
                                                                v4[v21] = -9223372036854775808
                                                                var v69 int32 = 1
                                                                _ = v69
                                                            default:
                                                                v1 = ((v1 - v1) + v27)
                                                            }
                                                        case "hello":
                                                            fmt.Println(v2, v2, fi22)
                                                            if false {
                                                                if (v5 < 128) {
                                                                    if (-1 < (v2 + 1)) {
                                                                        fmt.Println(v1)
                                                                        fmt.Println(v5, v3, fi22)
                                                                        if false {
                                                                            var v70 int16 = (v2 << ((4294967295) & 63))
                                                                            v70 = v2
                                                                            v71 := []int{fi46, fi22, -2147483648, 2147483647}
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                    } else {
                                                                        if (453 != 65535) {
                                                                            v1 -= 3.14159
                                                                        } else {
                                                                            switch (fi22 % (((fi22 ^ (-(fi46)))) | 1)) {
                                                                            case -9223372036854775808:
                                                                                var v72 uint32 = 645
                                                                                _ = v72
                                                                            }
                                                                        }
                                                                        if (((float32((0.0 + -2.25)) / (v1 * (v1 * 0.1))) > (((3.14159 * 0.1) - v1) / v27)) && false) {
                                                                            v21 = v21
                                                                            v73 := GenS4{F0: 1}
                                                                            var v74 int64 = 406
                                                                            v21 = (v21 + ((("" + (((("a" + "foo bar") + v21) + "hello") + v21)) + v21) + (v21 + v21)))
                                                                            v75 := []int{fi22}
                                                                            _ = v73
                                                                            _ = v74
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 string = ((v21 + v21) + v21)
                                                                            v77 := GenS1{F0: v2, F1: 460}
                                                                            var v78 int32 = -2147483648
                                                                            v79 := GenS8{F0: 42}
                                                                            var v80 int = fi22
                                                                            _ = v76
                                                                            _ = v77
                                                                            _ = v78
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                }
                                                                v81 := GenS9{F0: true, F1: 42, F2: 32767}
                                                                var v82 float32 = v1
                                                                v83 := []int{fi46, (fi22 ^ (^(fi22))), -2147483648, fi46}
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            } else {
                                                                v2 += v2
                                                            }
                                                            fmt.Println(v21, v20, v21)
                                                        }
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            }
                                            if false {
                                                switch fi22 {
                                                case -844:
                                                    v84 := v4[((("\n\t" + ("a" + (v21 + v21))) + (v21 + v21)) + v21)]
                                                    _ = v84
                                                case 128:
                                                    var v85 uint16 = 42
                                                    _ = v85
                                                case -686:
                                                    var v86 int16 = -665
                                                    _ = v86
                                                }
                                            }
                                        }
                                        _ = fi46
                                    }
                                case -2147483648:
                                    v87 := GenS10{F0: (^(v5)), F1: 0, F2: 2}
                                    _ = v87
                                default:
                                    v1 = (v1 / v27)
                                }
                            }
                        case 445:
                            for fi88 := 0; fi88 < 12; fi88++ {
                                v1 -= ((v27 - (v1 - 0.1)) / -2.25)
                                var v89 int32 = 42
                                _ = fi88
                                _ = v89
                            }
                        }
                        var v90 uint64 = 360
                        _ = v90
                    }
                } else {
                    for fi91 := 0; fi91 < 1; fi91++ {
                        var v92 uint = 18446744073709551615
                        _ = fi91
                        _ = v92
                    }
                    v93 := GenS8{F0: v5}
                    _ = v93
                }
            }
        case 32767:
            var v94 uint64 = 981
            _ = v94
        default:
            v95 := []int{1, fi22}
            _ = v95
        }
        _ = fi22
        _ = v27
    }
    var v96 uint64 = (v20 | (v20 * uint64(v3)))
    fmt.Println(v3, v2)
    v97 := []int{-1, -21, -880}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    fmt.Println("slicelen v97:", len(v97), v97)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v96
    _ = v97
}
