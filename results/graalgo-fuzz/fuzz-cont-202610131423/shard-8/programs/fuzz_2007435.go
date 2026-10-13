package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint; F1 bool; F2 uint16 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 uint; F1 uint8; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool; F1 int8; F2 int }
type GenS10 struct { F0 string }
func genFn0(p0 uint32) int8 {
    p0 -= (p0 & (p0 * 4294967295))
    fmt.Println(p0)
    v7 := GenS1{F0: 164, F1: 65535}
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(p0, p0)
        _ = fi8
    }
    _ = p0
    _ = v7
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]int)
    var v3 int8 = 24
    var v4 string = "a"
    var v5 int16 = 1
    v6 := make(map[uint32]string)
    switch 0 {
    case 42:
        delete(v2, 901)
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(fi10, v5, v1)
            _ = fi10
        }
        var v11 uint64 = 42
        v12 := []int{(-(fi9))}
        _ = fi9
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 32767}
    var v14 uint32 = 887
    for fi15 := 0; fi15 < 16; fi15++ {
        switch 812 {
        case -383:
            if true {
                fmt.Println(v5, fi15)
                for fi16 := 0; fi16 < 13; fi16++ {
                    fmt.Println(v4, v3, v5)
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        switch v4 {
                        case "a":
                            if (("foo bar" + (("\n\t" + v4) + v4)) == v4) {
                                v1 -= v1
                                var v17 int16 = (v5 ^ v5)
                                _ = v17
                            } else {
                                v18 := GenS3{F0: v3, F1: 18446744073709551615}
                                switch "hello" {
                                case "\n\t":
                                    var v19 int16 = (v5 ^ 42)
                                    _ = v19
                                default:
                                    var v20 uint64 = 18446744073709551615
                                    _ = v20
                                }
                                _ = v18
                            }
                        case "\n\t":
                            var v21 string = v4
                            _ = v21
                        }
                    case "":
                        v4 = (v4 + ("\n\t" + (((v4 + "\n\t") + "hello") + v4)))
                    default:
                        var v22 string = ""
                        _ = v22
                    }
                    _ = fi16
                }
            } else {
                if ((3.14159 / 1.0) > float32(1.5)) {
                    if false {
                        v23 := v6[v14]
                        for fi24 := 0; fi24 < 4; fi24++ {
                            switch "a" {
                            case "foo bar":
                                var v25 int32 = 1
                                _ = v25
                            }
                            if true {
                                if true {
                                    v26 := GenS2{F0: 282}
                                    v27 := v2[(v14 / (((^(v14))) | 1))]
                                    _ = v26
                                    _ = v27
                                } else {
                                    if false {
                                        var v28 int8 = genFn0((v14 % (((v14 / ((v14) | 1))) | 1)))
                                        _ = v28
                                    } else {
                                        switch v23 {
                                        case "":
                                            var v29 float64 = float64(1.5)
                                            _ = v29
                                        case "\n\t":
                                            var v30 float32 = float32(1.5)
                                            _ = v30
                                        case "hello":
                                            var v31 int32 = 944
                                            _ = v31
                                        }
                                        v32 := []int{(-(fi24))}
                                        if false {
                                            v4 = "a"
                                        }
                                        _ = v32
                                    }
                                    var v33 int16 = 1
                                    _ = v33
                                }
                                var v34 float32 = (float32(0.0) - -2.25)
                                v35 := GenS4{F0: 0, F1: v23}
                                _ = v34
                                _ = v35
                            }
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        fmt.Println(v4, fi15, v1)
                        switch fi15 {
                        case 767:
                            for fi36 := 0; fi36 < 14; fi36++ {
                                v37 := GenS2{F0: 28}
                                v38 := GenS4{F0: 0, F1: ("\n\t" + v4)}
                                if false {
                                    fmt.Println(v5, fi36, v1)
                                    v1 += (v1 + v1)
                                } else {
                                    switch v4 {
                                    case "\n\t":
                                        var v39 int8 = (v3 % (((v3 * v3)) | 1))
                                        _ = v39
                                    }
                                    switch 128 {
                                    case 1:
                                        v1 -= -67
                                    case 2147483647:
                                        v3 += v3
                                    default:
                                        if (true || (!(v4 >= v4))) {
                                            _ = len(v6)
                                        }
                                    }
                                    var v40 uint32 = (v14 >> ((493) & 63))
                                    var v41 uint64 = 1
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi36
                                _ = v37
                                _ = v38
                            }
                        case 127:
                            fmt.Println(v14, fi15)
                        default:
                            for fi42 := 0; fi42 < 5; fi42++ {
                                v43 := []int{}
                                _ = fi42
                                _ = v43
                            }
                        }
                        var v44 int32 = 42
                        _ = v44
                    }
                } else {
                    var v45 int32 = -1
                    switch (v4 + "\n\t") {
                    case "hello":
                        for fi46 := 0; fi46 < 13; fi46++ {
                            if false {
                                switch (^(fi15)) {
                                case -602:
                                    if (!(-880 == v1)) {
                                        v47 := v2[432]
                                        v4 = "a"
                                        var v48 float64 = 3.14159
                                        _ = v47
                                        _ = v48
                                    } else {
                                        var v49 uint32 = v14
                                        _ = v49
                                    }
                                }
                                fmt.Println(v4)
                                v50 := []int{(fi46 * 32768)}
                                v51 := GenS5{F0: 749, F1: false, F2: 2}
                                var v52 float64 = 1e10
                                _ = v50
                                _ = v51
                                _ = v52
                            } else {
                                fmt.Println(fi15, v1)
                                switch "a" {
                                case "":
                                    v1 -= v1
                                }
                                var v53 float32 = ((-2.25 / ((1.5 - 0.1) * 1.5)) * ((((-2.25 + 1.5) * (((0.0 / float32((3.14159 / 1.0))) / 0.1) - 1.5)) + 3.14159) / 1e10))
                                _ = v53
                            }
                            _ = fi46
                        }
                    case "a":
                        v54 := GenS1{F0: int(v1), F1: 0}
                        _ = v54
                    case "foo bar":
                        v1 += 192
                    default:
                        var v55 int16 = (-(v5))
                        _ = v55
                    }
                    v3 = (^(v3))
                    _ = v45
                }
            }
        case -9223372036854775808:
            v56 := []int{}
            _ = v56
        case 2147483647:
            v3 = v3
        default:
            v57 := []int{(-(fi15)), (^(fi15))}
            _ = v57
        }
        v5 += v5
        _ = fi15
    }
    if false {
        if (3.14159 == 1e10) {
            if (v14 != 138) {
                switch v4 {
                case "":
                    v5 -= v5
                case "a":
                    v58 := GenS3{F0: genFn0(65535), F1: 65535}
                    _ = v58
                }
                v59 := []int{-9223372036854775808, 2147483647, 4294967295}
                v60 := GenS5{F0: 965, F1: true, F2: uint16(v14)}
                _ = v59
                _ = v60
            }
        } else {
            v61 := GenS5{F0: 255, F1: false, F2: 0}
            var v62 float64 = 1e10
            _ = v61
            _ = v62
        }
        v63 := GenS6{F0: 9223372036854775807, F1: v5}
        v64 := GenS1{F0: -545, F1: 549}
        _ = v63
        _ = v64
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        v66 := GenS5{F0: 18446744073709551615, F1: false, F2: 32767}
        fmt.Println(v5, fi65, v4)
        v67 := GenS7{F0: 2, F1: 0, F2: (fi65 - (fi65 | fi65))}
        _ = fi65
        _ = v66
        _ = v67
    }
    v68 := GenS1{F0: 72, F1: 610}
    fmt.Println(v5, v3, v3)
    if (!(false && false)) {
        for fi69 := 0; fi69 < 18; fi69++ {
            for fi70 := 0; fi70 < 20; fi70++ {
                if false {
                    var v71 uint64 = 42
                    var v72 float32 = (-2.25 * 0.0)
                    _ = v71
                    _ = v72
                } else {
                    switch ((v4 + v4) + (v4 + v4)) {
                    case "\n\t":
                        fmt.Println(fi69, fi69)
                    case "hello":
                        v5 += (v5 ^ -855)
                    default:
                        delete(v2, 42)
                    }
                    _ = len(v6)
                    if ((-2.25 <= ((1.5 / 0.1) - 1.5)) && (true || true)) {
                        v1 = -1
                        var v73 float32 = 3.14159
                        var v74 uint64 = 0
                        var v75 uint32 = (v14 + v14)
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                }
                switch v4 {
                case "\n\t":
                    fmt.Println(fi69)
                }
                _ = fi70
            }
            var v76 int32 = 1
            if true {
                v77 := GenS4{F0: 1, F1: ("foo bar" + v4)}
                for fi78 := 0; fi78 < 9; fi78++ {
                    var v79 uint32 = 2147483648
                    if false {
                        if true {
                            _ = len(v6)
                            if (!true) {
                                v80 := GenS8{F0: (2 < 18446744073709551615)}
                                fmt.Println(v76)
                                var v81 string = (v4 + ("\n\t" + ""))
                                _ = v80
                                _ = v81
                            }
                        }
                        switch (fi78 | fi69) {
                        case 32768:
                            var v82 uint64 = 2
                            _ = v82
                        case -522:
                            v76 += (v76 << ((704) & 63))
                        default:
                            v5 = (v5 | (v5 + v5))
                        }
                    } else {
                        v14 = v79
                        v14 = v79
                    }
                    _ = fi78
                    _ = v79
                }
                _ = v77
            }
            _ = fi69
            _ = v76
        }
    }
    v83 := GenS9{F0: true, F1: v3, F2: 32768}
    if true {
        if (!(4294967295 != 18446744073709551615)) {
            v4 = ""
            var v84 float32 = (-2.25 / 1.5)
            v85 := GenS2{F0: 29}
            _ = v84
            _ = v85
        }
        switch ("\n\t" + ("\n\t" + "")) {
        case "a":
            var v86 int32 = -1
            _ = v86
        case "":
            if false {
                var v87 uint32 = v14
                fmt.Println(v87, v5, v5)
                var v88 uint8 = 128
                _ = v87
                _ = v88
            }
        case "foo bar":
            v3 += 100
        }
    }
    for fi89 := 0; fi89 < 7; fi89++ {
        if false {
            if (!true) {
                var v90 int64 = (v1 * (v1 * (v1 & v1)))
                for fi91 := 0; fi91 < 7; fi91++ {
                    v92 := GenS3{F0: (v3 - (v3 - (-(v3)))), F1: 18446744073709551615}
                    _ = fi91
                    _ = v92
                }
                var v93 bool = (180 > 42)
                _ = v90
                _ = v93
            } else {
                fmt.Println(fi89, v1, v5)
                fmt.Println(fi89, v3)
            }
            switch "hello" {
            case "\n\t":
                var v94 uint64 = 1
                _ = v94
            default:
                var v95 uint32 = v14
                _ = v95
            }
            var v96 float32 = 1e10
            _ = v96
        }
        switch v4 {
        case "\n\t":
            fmt.Println(v14, v5)
        case "hello":
            v97 := []int{fi89, (fi89 ^ fi89), fi89}
            _ = v97
        default:
            var v98 float32 = 3.14159
            _ = v98
        }
        v99 := []int{int(v3)}
        _ = fi89
        _ = v99
    }
    v100 := GenS10{F0: v4}
    v101 := []int{113, -1, 128, 32767}
    var v102 int32 = 2147483647
    if (v4 != v4) {
        for fi103 := 0; fi103 < 10; fi103++ {
            v101 = append(v101, fi103)
            fmt.Println(v3, v4, v4)
            v5 -= 138
            _ = fi103
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v102:", v102)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v83
}
