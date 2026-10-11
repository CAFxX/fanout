package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint; F1 int8; F2 bool }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int16; F1 int32 }
func genFn0() uint32 {
    var v7 float64 = (float64((float32(0.0) * 0.0)) + 0.0)
    var v8 int16 = 789
    v9 := GenS1{F0: 2}
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{(fi10 & (fi10 << ((791) & 63))), (fi10 * (fi10 ^ (fi10 / ((fi10) | 1))))}
        for fi12 := 0; fi12 < 11; fi12++ {
            v8 = (v8 / ((v8) | 1))
            v11 = append(v11, fi10)
            var v13 string = "hello"
            fmt.Println(fi12, v7)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 439
}
func genFn1(p0 bool, p1 int16, p2 uint) int {
    if (p0 || p0) {
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    if p0 {
        v15 := []int{32768, 32768}
        for fi16 := 0; fi16 < 3; fi16++ {
            if (321 < 2147483647) {
                var v17 uint64 = 18446744073709551615
                if p0 {
                    v17 = (v17 * 2)
                    p0 = p0
                    var v18 uint = 343
                    _ = v18
                } else {
                    if p0 {
                        _ = len(v15)
                        p2 = 179
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                    }
                }
                var v19 int8 = -21
                _ = v17
                _ = v19
            } else {
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 int16 = p1
                    _ = fi20
                    _ = v21
                }
                if p0 {
                    fmt.Println(p0, fi16, p2)
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if (p0 && ((true || p0) || p0)) {
                            p0 = p0
                        }
                        _ = fi22
                    }
                } else {
                    fmt.Println(p0, fi16, p1)
                }
                var v23 uint32 = 878
                _ = v23
            }
            _ = fi16
        }
        var v24 string = "foo bar"
        _ = v15
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 int = -2147483648
    v2 := make(map[uint32]string)
    var v3 int16 = 0
    var v4 float64 = 1e10
    var v5 int32 = 429
    v6 := make(map[uint32]int)
    if true {
        var v25 int8 = 0
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 uint16 = 2
            _ = fi26
            _ = v27
        }
        if false {
            v5 += 1
            fmt.Println(v25)
        } else {
            v1 = (v1 | (-(v1)))
        }
        _ = v25
    } else {
        var v28 int32 = (v5 % ((-2147483648) | 1))
        _ = v28
    }
    delete(v6, 449)
    var v29 uint32 = genFn0()
    switch "a" {
    case "":
        var v30 bool = (!(127 != -128))
        _ = v30
    case "foo bar":
        v5 = v5
    default:
        var v31 float32 = 1.5
        _ = v31
    }
    if false {
        var v32 int16 = 42
        v33 := v2[v29]
        _ = v32
        _ = v33
    } else {
        var v34 uint = 932
        if (!true) {
            var v35 int64 = 0
            for fi36 := 0; fi36 < 19; fi36++ {
                v37 := GenS1{F0: 1}
                _ = fi36
                _ = v37
            }
            v38 := v6[v29]
            _ = v35
            _ = v38
        }
        var v39 uint64 = 1
        _ = v34
        _ = v39
    }
    v40 := GenS2{F0: 42}
    delete(v2, genFn0())
    switch (v1 ^ v1) {
    case 287:
        var v41 uint32 = 325
        _ = v41
    case 516:
        v42 := v2[(v29 >> ((608) & 63))]
        _ = v42
    case 32768:
        v43 := v6[(v29 & 2147483647)]
        _ = v43
    default:
        var v44 uint = 18446744073709551615
        _ = v44
    }
    if (true && true) {
        var v45 int = genFn1(true, (v3 * (v3 + v3)), 255)
        var v46 int32 = (v5 | -459)
        _ = v45
        _ = v46
    } else {
        if true {
            switch genFn1(false, (v3 - (v3 - (v3 ^ (v3 | v3)))), 303) {
            case -633:
                if false {
                    if true {
                        v47 := v2[genFn0()]
                        fmt.Println(v29, v1)
                        _ = v47
                    } else {
                        switch (^(v1)) {
                        case -1:
                            fmt.Println(v5, v4, v1)
                        }
                        v3 -= (v3 - (v3 - v3))
                        var v48 uint32 = (v29 << ((2) & 63))
                        _ = v48
                    }
                } else {
                    v4 -= ((0.0 * (0.1 + ((v4 + v4) - -2.25))) + 3.14159)
                    v5 = (v5 | (v5 / (((-(v5))) | 1)))
                    if ((-(v5)) != (v5 >> ((65535) & 63))) {
                        v3 = (v3 % ((v3) | 1))
                        switch (v1 & -183) {
                        case 42:
                            v29 = v29
                        case -9223372036854775808:
                            v1 -= v1
                        }
                    }
                }
            }
        } else {
            if true {
                var v49 uint64 = 214
                var v50 uint16 = 65535
                _ = v49
                _ = v50
            } else {
                fmt.Println(v1, v3, v3)
                var v51 uint = 859
                _ = v51
            }
            if ((3.14159 * (-2.25 - (float32(v4) + (3.14159 * float32((v4 + v4)))))) != 3.14159) {
                switch "hello" {
                case "foo bar":
                    v52 := []int{v1, v1}
                    _ = v52
                }
                switch (v1 >> ((18446744073709551615) & 63)) {
                case 1:
                    fmt.Println(v3, v1)
                }
                v53 := GenS3{F0: 883, F1: 0, F2: 9223372036854775807}
                _ = v53
            } else {
                v1 = (^(v1))
                if (128 > 240) {
                    for fi54 := 0; fi54 < 11; fi54++ {
                        for fi55 := 0; fi55 < 13; fi55++ {
                            delete(v2, v29)
                            v56 := GenS3{F0: (v5 - (v5 << ((42) & 63))), F1: 42, F2: 9223372036854775807}
                            _ = fi55
                            _ = v56
                        }
                        _ = fi54
                    }
                    for fi57 := 0; fi57 < 1; fi57++ {
                        switch "\n\t" {
                        case "":
                            var v58 int = v1
                            _ = v58
                        case "hello":
                            var v59 uint32 = (^(v29))
                            _ = v59
                        case "a":
                            switch v1 {
                            case 4294967295:
                                if true {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        var v60 int32 = (v5 / ((337) | 1))
                                        _ = v60
                                    case "foo bar":
                                        switch -1 {
                                        case -533:
                                            if (1 > 0) {
                                                v61 := GenS2{F0: 1}
                                                for fi62 := 0; fi62 < 14; fi62++ {
                                                    v3 += v3
                                                    v3 = v3
                                                    if false {
                                                        if (1 == 42) {
                                                            if false {
                                                                v63 := []int{(v1 & (v1 & genFn1(false, v3, 607)))}
                                                                _ = v63
                                                            } else {
                                                                if (255 > 2) {
                                                                    _ = len(v6)
                                                                    switch "a" {
                                                                    case "":
                                                                        if true {
                                                                            var v64 uint32 = (^(v29))
                                                                            _ = v64
                                                                        } else {
                                                                            v65 := GenS2{F0: 0}
                                                                            _ = v65
                                                                        }
                                                                        v3 += v3
                                                                        v29 += (v29 / ((4294967295) | 1))
                                                                    case "hello":
                                                                        switch (fi62 & (v1 % ((-953) | 1))) {
                                                                        case 212:
                                                                            var v66 int32 = 42
                                                                            _ = v66
                                                                        case 0:
                                                                            if false {
                                                                                v67 := GenS4{F0: 0, F1: v3, F2: 18446744073709551615}
                                                                                _ = len(v2)
                                                                                if (fi62 != (fi57 - -869)) {
                                                                                    if false {
                                                                                        v4 -= 0.1
                                                                                        v1 -= v1
                                                                                        var v68 uint64 = 0
                                                                                        _ = v68
                                                                                    }
                                                                                } else {
                                                                                    var v69 uint8 = 0
                                                                                    if true {
                                                                                        v69 -= uint8(v29)
                                                                                    }
                                                                                    _ = v69
                                                                                }
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        v70 := []int{}
                                                                        _ = v70
                                                                    }
                                                                    v5 = -827
                                                                    v71 := GenS4{F0: 0, F1: (v3 ^ (v3 * v3)), F2: 1}
                                                                    _ = v71
                                                                }
                                                                var v72 int = 32767
                                                                switch ("\n\t" + "") {
                                                                case "\n\t":
                                                                    v29 -= v29
                                                                }
                                                                _ = v72
                                                            }
                                                        }
                                                        v1 -= (v1 / ((v1) | 1))
                                                        v5 = v5
                                                    }
                                                    _ = fi62
                                                }
                                                _ = v61
                                            }
                                        }
                                    default:
                                        for fi73 := 0; fi73 < 16; fi73++ {
                                            var v74 int16 = v3
                                            var v75 string = "a"
                                            _ = fi73
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    switch (fi57 << ((617) & 63)) {
                                    case 128:
                                        v4 -= ((float64(0.0) / 1.5) * float64(1e10))
                                    default:
                                        var v76 int32 = (v5 | -235)
                                        _ = v76
                                    }
                                    v77 := GenS5{F0: 255, F1: -128, F2: (((-2.25 * -2.25) / 3.14159) <= float32((v4 * ((-2.25 / v4) * v4))))}
                                    _ = v77
                                }
                            default:
                                v3 = v3
                            }
                        }
                        v5 -= (^(v5))
                        v78 := []int{genFn1((float64(float32((v4 / v4))) <= 3.14159), (v3 + (v3 / ((v3) | 1))), 567), genFn1((255 < 1), v3, 4294967295)}
                        _ = fi57
                        _ = v78
                    }
                    if true {
                        v29 -= 785
                        switch (v1 / (((v1 | v1)) | 1)) {
                        case -479:
                            for fi79 := 0; fi79 < 16; fi79++ {
                                var v80 float32 = 1e10
                                for fi81 := 0; fi81 < 5; fi81++ {
                                    var v82 int = fi81
                                    fmt.Println(v80)
                                    v83 := []int{(-(fi79)), fi81, v82}
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                }
                                _ = fi79
                                _ = v80
                            }
                        case 42:
                            var v84 float32 = 1.5
                            _ = v84
                        case -1:
                            var v85 uint32 = v29
                            _ = v85
                        default:
                            v4 = 0.0
                        }
                    }
                } else {
                    fmt.Println(v1, v5)
                }
                v3 = -282
            }
            if true {
                if true {
                    v4 = (v4 + (float64(3.14159) + ((v4 / v4) + (((((1.5 + v4) * 3.14159) - -2.25) + 0.0) / ((v4 + (0.0 / (((v4 / v4) - (float64(3.14159) - 0.0)) * (float64((0.0 - (3.14159 - (1e10 * 0.1)))) * (v4 - v4))))) * (((float64(0.1) - (v4 - 0.1)) / v4) * 1.5))))))
                    var v86 uint64 = 171
                    var v87 uint = 2
                    _ = v86
                    _ = v87
                }
                for fi88 := 0; fi88 < 8; fi88++ {
                    v89 := GenS6{F0: 32767, F1: 18446744073709551615, F2: 115}
                    if (18446744073709551615 == 1) {
                        v90 := []int{v1, fi88}
                        _ = v90
                    } else {
                        fmt.Println(fi88, v5)
                    }
                    _ = fi88
                    _ = v89
                }
                if (!(true && false)) {
                    for fi91 := 0; fi91 < 20; fi91++ {
                        if (2 == uint64(v29)) {
                            var v92 uint64 = uint64(v29)
                            v93 := []int{1, 4294967295, v1}
                            _ = v92
                            _ = v93
                        } else {
                            var v94 uint64 = 1
                            v29 = genFn0()
                            if true {
                                v94 = v94
                                v95 := GenS6{F0: 2, F1: (^(v94)), F2: 29}
                                _ = v95
                            }
                            _ = v94
                        }
                        for fi96 := 0; fi96 < 11; fi96++ {
                            if (!(-2.25 != 3.14159)) {
                                v29 += v29
                            }
                            var v97 float64 = -2.25
                            v98 := GenS7{F0: 42, F1: 954}
                            _ = fi96
                            _ = v97
                            _ = v98
                        }
                        v1 += genFn1((!(false || false)), (v3 << ((2) & 63)), 18446744073709551615)
                        _ = fi91
                    }
                } else {
                    var v99 int64 = 1
                    v100 := GenS2{F0: 127}
                    if false {
                        var v101 float32 = 1e10
                        var v102 int16 = v3
                        _ = v101
                        _ = v102
                    } else {
                        v103 := GenS6{F0: 42, F1: 2, F2: 128}
                        v104 := GenS8{F0: (v3 >> ((1) & 63)), F1: (^(v5))}
                        if false {
                            var v105 uint = 2
                            _ = v105
                        }
                        _ = v103
                        _ = v104
                    }
                    _ = v100
                    _ = v99
                }
            }
        }
    }
    v4 = (v4 * v4)
    v106 := GenS8{F0: (v3 << ((528) & 63)), F1: (v5 & (v5 & (v5 ^ (-(v5)))))}
    v4 += (((v4 + 0.0) * float64(0.0)) * 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v106
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
