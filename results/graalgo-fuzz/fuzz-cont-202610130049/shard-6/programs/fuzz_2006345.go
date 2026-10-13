package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int64; F2 uint8 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 int64; F1 bool; F2 uint32 }
type GenS9 struct { F0 int8; F1 uint16; F2 uint8 }
func genFn0() uint {
    if true {
        var v4 int16 = -1
        v4 += 42
        _ = v4
    } else {
        var v5 uint16 = 2
        var v6 uint8 = 0
        var v7 uint16 = v5
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 42}
    var v9 string = ("foo bar" + "\n\t")
    var v10 int64 = 42
    var v11 int = 128
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int8, p2 int32) int16 {
    p1 += 65
    switch ((((("" + "") + "foo bar") + "") + ("" + "foo bar")) + ("\n\t" + ("\n\t" + ("a" + "a")))) {
    case "\n\t":
        p0 = 988
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            if ((2147483647 == 42) || false) {
                var v13 int8 = (p1 & 93)
                if (18446744073709551615 < 0) {
                    switch "\n\t" {
                    case "a":
                        var v14 int64 = 0
                        _ = v14
                    case "foo bar":
                        if true {
                            var v15 uint64 = 42
                            p0 = int16(v15)
                            fmt.Println(v15, p0, v13)
                            _ = v15
                        }
                    }
                    var v16 float32 = 3.14159
                    for fi17 := 0; fi17 < 19; fi17++ {
                        v13 += (p1 * p1)
                        _ = fi17
                    }
                    _ = v16
                }
                v18 := GenS2{F0: 1, F1: (^(v13))}
                _ = v13
                _ = v18
            }
            var v19 uint = 327
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = (^(fi12))
                v22 := []int{}
                _ = len(v22)
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi12
            _ = v19
        }
    case "foo bar":
        switch (("hello" + "a") + (("hello" + "a") + ("hello" + ""))) {
        case "hello":
            var v23 uint32 = 1
            _ = v23
        case "a":
            for fi24 := 0; fi24 < 7; fi24++ {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := []int{fi24, fi24, (fi25 | (fi24 % ((fi25) | 1)))}
                    if (!true) {
                        switch (fi25 << ((228) & 63)) {
                        case 2147483647:
                            p1 += (p1 + 1)
                        }
                        p2 = p2
                        p1 -= int8(p2)
                    }
                    var v27 uint8 = 1
                    if len(v26) > 0 {
                        v26[0] = 127
                    }
                    switch fi25 {
                    case 548:
                        v28 := GenS2{F0: -1, F1: 127}
                        _ = v28
                    default:
                        switch "" {
                        case "foo bar":
                            switch (("\n\t" + "a") + "foo bar") {
                            case "\n\t":
                                switch (fi25 + 128) {
                                case 0:
                                    if (1e10 != (0.1 + -2.25)) {
                                        var v29 uint16 = 103
                                        _ = v29
                                    } else {
                                        var v30 uint8 = 1
                                        v27 += (v27 & (v30 ^ v30))
                                        _ = v30
                                    }
                                case 738:
                                    if false {
                                        if true {
                                            if true {
                                                var v31 uint = 18446744073709551615
                                                var v32 int8 = p1
                                                if true {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v33 float32 = 1.5
                                                        _ = v33
                                                    }
                                                    fmt.Println(v27, p0)
                                                    v34 := GenS3{F0: (true || false)}
                                                    _ = v34
                                                } else {
                                                    if (fi25 > (fi24 - fi25)) {
                                                        if (int(v31) <= (-(fi25))) {
                                                            fmt.Println(v31)
                                                            fmt.Println(fi25, fi24)
                                                            v35 := GenS2{F0: 9223372036854775807, F1: (^(p1))}
                                                            fmt.Println(v31, v31, fi25)
                                                            if true {
                                                                p0 += (p0 & p0)
                                                                var v36 float64 = 1.5
                                                                _ = v36
                                                            }
                                                            v37 := GenS4{F0: p1, F1: 9223372036854775807, F2: 2}
                                                            _ = v35
                                                            _ = v37
                                                        } else {
                                                            if true {
                                                                if true {
                                                                    var v38 uint64 = 0
                                                                    fmt.Println(v38, v38, v31)
                                                                    v39 := GenS5{F0: (p2 >> (((v31 | 703)) & 63)), F1: "a"}
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    p1 += int8(fi25)
                                                                }
                                                            } else {
                                                                fmt.Println(p0, v27, v27)
                                                            }
                                                            v40 := GenS3{F0: ((true || false) || (818 <= 1))}
                                                            _ = v40
                                                        }
                                                    } else {
                                                        fmt.Println(v32, v32)
                                                        if (true || false) {
                                                            v32 -= (^(p1))
                                                            switch 4294967295 {
                                                            case 707:
                                                                var v41 float32 = 0.1
                                                                _ = v41
                                                            case -9223372036854775808:
                                                                var v42 float32 = (float32(0.0) - -2.25)
                                                                _ = v42
                                                            }
                                                        } else {
                                                            p1 += v32
                                                            var v43 uint = (v31 ^ v31)
                                                            v27 = (v27 / (((v27 * (^(v27)))) | 1))
                                                            _ = v43
                                                        }
                                                        if false {
                                                            var v44 uint8 = v27
                                                            fmt.Println(v31)
                                                            v44 = 204
                                                            _ = v44
                                                        } else {
                                                            var v45 uint = v31
                                                            _ = v45
                                                        }
                                                    }
                                                    var v46 int16 = p0
                                                    var v47 bool = (!(32767 == 617))
                                                    _ = v46
                                                    _ = v47
                                                }
                                                _ = v31
                                                _ = v32
                                            } else {
                                                switch fi24 {
                                                case 32768:
                                                    var v48 int32 = 885
                                                    if (genFn0() != 369) {
                                                        if true {
                                                            p0 = p0
                                                        } else {
                                                            if ((fi25 << ((18446744073709551615) & 63)) >= fi24) {
                                                                v49 := GenS4{F0: p1, F1: 9223372036854775807, F2: v27}
                                                                v50 := GenS6{F0: (fi24 << ((518) & 63))}
                                                                v51 := GenS1{F0: 42}
                                                                switch fi24 {
                                                                case -63:
                                                                    var v52 uint = 4294967295
                                                                    _ = v52
                                                                case 9223372036854775807:
                                                                    if ((v27 / ((v27) | 1)) <= v27) {
                                                                        switch (fi24 / ((fi25) | 1)) {
                                                                        case 0:
                                                                            p1 = int8(p2)
                                                                        }
                                                                        v26 = append(v26, fi25)
                                                                    } else {
                                                                        v48 -= (p2 - (p2 % ((v48) | 1)))
                                                                    }
                                                                case 227:
                                                                    fmt.Println(v48)
                                                                default:
                                                                    v53 := GenS6{F0: fi25}
                                                                    _ = v53
                                                                }
                                                                var v54 uint64 = 1
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                                _ = v54
                                                            } else {
                                                                _ = len(v26)
                                                            }
                                                        }
                                                    }
                                                    v55 := GenS4{F0: p1, F1: 42, F2: v27}
                                                    v26 = append(v26, 32768)
                                                    var v56 uint = genFn0()
                                                    _ = v48
                                                    _ = v55
                                                    _ = v56
                                                case 0:
                                                    fmt.Println(p2, fi24)
                                                case 42:
                                                    fmt.Println(fi24, p0)
                                                }
                                            }
                                        } else {
                                            _ = len(v26)
                                            var v57 uint32 = 4294967295
                                            _ = v57
                                        }
                                        var v58 uint = uint(v27)
                                        _ = v58
                                    } else {
                                        fmt.Println(fi25, fi24, p2)
                                        v26 = append(v26, 128)
                                    }
                                case -399:
                                    if len(v26) > 0 {
                                        v26[0] = fi25
                                    }
                                default:
                                    v59 := GenS3{F0: (-339 >= p0)}
                                    _ = v59
                                }
                            default:
                                v27 = (v27 << ((1) & 63))
                                if false {
                                    var v60 uint16 = uint16(v27)
                                    var v61 int16 = p0
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v62 bool = true
                                    _ = v62
                                }
                                p1 += 62
                                fmt.Println(p2, fi24, p1)
                                fmt.Println(fi24)
                            }
                        }
                    }
                    v63 := GenS1{F0: genFn0()}
                    v64 := GenS6{F0: (^(fi25))}
                    fmt.Println(p1, p0, fi24)
                    p0 -= 32767
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v63
                    _ = v64
                }
                switch ("a" + "foo bar") {
                case "a":
                    switch "foo bar" {
                    case "a":
                        var v65 bool = false
                        _ = v65
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            switch "hello" {
                            case "a":
                                v66 := GenS2{F0: -9223372036854775808, F1: p1}
                                _ = v66
                            case "foo bar":
                                p1 += p1
                            case "":
                                p2 = (p2 >> ((4294967295) & 63))
                            default:
                                switch -890 {
                                case 9223372036854775807:
                                    if false {
                                        var v67 int32 = 926
                                        v68 := []int{fi24, fi24}
                                        var v69 bool = false
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    }
                                case -547:
                                    p2 = (p2 - 2147483647)
                                case -674:
                                    var v70 int = -984
                                    _ = v70
                                default:
                                    p1 += 0
                                }
                            }
                        case "a":
                            fmt.Println(p1)
                        }
                    }
                case "foo bar":
                    v71 := GenS1{F0: 0}
                    _ = v71
                case "":
                    if (true && (!false)) {
                        p2 = (p2 * p2)
                    } else {
                        v72 := GenS7{F0: -332, F1: 2}
                        v73 := []int{fi24, -55}
                        _ = v72
                        _ = v73
                    }
                }
                for fi74 := 0; fi74 < 16; fi74++ {
                    v75 := GenS2{F0: -445, F1: p1}
                    v76 := GenS6{F0: (fi74 + -916)}
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                _ = fi24
            }
        case "foo bar":
            p1 = (p1 * p1)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 - p0)
}

func main() {
    v1 := []int{}
    var v2 string = ""
    v3 := make(map[int]bool)
    var v77 float32 = ((1.5 / 3.14159) / 1.0)
    for fi78 := 0; fi78 < 6; fi78++ {
        v77 = (3.14159 * (v77 / (v77 / float32(float64((float32(1.5) * v77))))))
        _ = fi78
    }
    if (1000 <= 393) {
        for fi79 := 0; fi79 < 18; fi79++ {
            if true {
                fmt.Println(v77, v77, fi79)
                v80 := GenS8{F0: -9223372036854775808, F1: (fi79 < (fi79 << ((2) & 63))), F2: 123}
                for fi81 := 0; fi81 < 13; fi81++ {
                    if (false && false) {
                        if (genFn1(-97, -128, 0) != genFn1(-1, 28, -1)) {
                            switch (("foo bar" + v2) + (v2 + v2)) {
                            case "\n\t":
                                if (1e10 < 0.1) {
                                    v2 = ""
                                }
                            case "":
                                var v82 uint = 90
                                _ = v82
                            case "a":
                                var v83 uint8 = 153
                                _ = v83
                            default:
                                var v84 bool = true
                                _ = v84
                            }
                            switch v2 {
                            case "hello":
                                v2 = "\n\t"
                            case "a":
                                v2 = ""
                                v85 := GenS9{F0: 42, F1: 1, F2: 103}
                                var v86 uint = genFn0()
                                _ = v85
                                _ = v86
                            case "":
                                if len(v1) > 0 {
                                    v1[0] = fi81
                                }
                            default:
                                v2 = v2
                            }
                        }
                    }
                    v77 = (float32(3.14159) / v77)
                    var v87 float64 = ((1.5 + float64((v77 + v77))) - float64(float32((-2.25 + 0.1))))
                    _ = fi81
                    _ = v87
                }
                _ = v80
            }
            _ = fi79
        }
        _ = len(v3)
        var v88 int16 = 42
        _ = v88
    } else {
        if (1 != 2) {
            var v89 int = 42
            v89 = v89
            _ = v89
        }
        for fi90 := 0; fi90 < 11; fi90++ {
            var v91 uint64 = 0
            _ = fi90
            _ = v91
        }
    }
    v92 := v3[-9223372036854775808]
    for fi93 := 0; fi93 < 13; fi93++ {
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
        delete(v3, fi93)
        var v94 uint8 = 255
        _ = fi93
        _ = v94
    }
    _ = len(v1)
    if false {
        switch 127 {
        case 128:
            switch 42 {
            case -2147483648:
                v77 += v77
            default:
                for fi95 := 0; fi95 < 18; fi95++ {
                    var v96 uint8 = 255
                    _ = len(v1)
                    var v97 uint64 = 4294967295
                    _ = fi95
                    _ = v96
                    _ = v97
                }
            }
        case 127:
            for fi98 := 0; fi98 < 20; fi98++ {
                v3[fi98] = (!true)
                _ = len(v3)
                _ = fi98
            }
        }
        switch v2 {
        case "\n\t":
            v99 := v3[0]
            _ = v99
        case "a":
            switch 127 {
            case -1:
                var v100 int16 = 0
                _ = v100
            }
        case "hello":
            for fi101 := 0; fi101 < 3; fi101++ {
                v1 = append(v1, 128)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v102 uint8 = 0
                _ = fi101
                _ = v102
            }
        }
    } else {
        var v103 uint16 = 2
        v104 := GenS6{F0: 452}
        switch v2 {
        case "foo bar":
            var v105 int32 = 0
            _ = v105
        case "\n\t":
            for fi106 := 0; fi106 < 6; fi106++ {
                v77 = (((v77 + v77) / 1e10) + ((v77 / 0.1) * float32(0.1)))
                v107 := GenS1{F0: 385}
                fmt.Println(v92, v77)
                _ = fi106
                _ = v107
            }
        case "hello":
            for fi108 := 0; fi108 < 11; fi108++ {
                v109 := GenS5{F0: 1, F1: ("" + "")}
                _ = fi108
                _ = v109
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = v103
        _ = v104
    }
    var v110 uint8 = 170
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v77:", v77)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v110
    _ = v2
    _ = v3
    _ = v77
    _ = v92
}
