package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int8; F1 int16 }
type GenS10 struct { F0 uint; F1 int32; F2 uint }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 uint; F1 int8; F2 uint32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 uint64; F2 uint }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool; F1 int }
type GenS18 struct { F0 int8; F1 uint }
type GenS19 struct { F0 int; F1 int; F2 uint64 }
type GenS20 struct { F0 int16; F1 string }
type GenS21 struct { F0 string }
type GenS22 struct { F0 int16; F1 bool; F2 int }
type GenS23 struct { F0 uint; F1 bool }
func genFn0(p0 int) uint16 {
    fmt.Println(p0)
    switch p0 {
    case 4:
        p0 += p0
    case 32768:
        p0 += -1
    default:
        var v6 int16 = -1
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 278}
        var v9 int16 = 1000
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += p0
    _ = p0
    return 42
}
func genFn1(p0 int8) int8 {
    if (!false) {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS2{F0: "", F1: true, F2: -98}
                v13 := []int{(fi10 >> ((99) & 63)), fi10}
                fmt.Println(p0, fi10, fi10)
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        fmt.Println(p0)
        var v14 uint16 = 42
        _ = v14
    }
    var v15 int8 = 104
    _ = p0
    _ = v15
    return -23
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    v3 := []uint32{}
    v4 := make(map[int]bool)
    var v5 string = ""
    if false {
        v1[74] = (false && false)
    } else {
        delete(v1, 2147483648)
        var v16 int8 = -128
        _ = v16
    }
    v5 = ""
    fmt.Println(v2, v2)
    v2 += v2
    var v17 uint32 = 1
    switch (v5 + "hello") {
    case "hello":
        v18 := GenS2{F0: v5, F1: false, F2: 127}
        _ = v18
    case "a":
        var v19 string = v5
        _ = v19
    }
    fmt.Println(v17, v17, v2)
    v2 = -2.25
    v5 = ("a" + v5)
    if false {
        fmt.Println(v5, v17)
        v20 := GenS2{F0: "", F1: (127 > genFn1(1)), F2: 121}
        switch v5 {
        case "a":
            for fi21 := 0; fi21 < 4; fi21++ {
                switch v5 {
                case "\n\t":
                    fmt.Println(v2, v5)
                case "hello":
                    switch v5 {
                    case "":
                        fmt.Println(v2, v5, v17)
                    case "hello":
                        var v22 uint8 = 127
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v1[v17] = (v17 > (^(v17)))
                            var v24 uint16 = 42
                            if false {
                                fmt.Println(fi23)
                            } else {
                                v25 := GenS3{F0: 65535}
                                _ = v25
                            }
                            _ = fi23
                            _ = v24
                        }
                    }
                case "a":
                    v26 := GenS4{F0: 2147483647}
                    _ = v26
                }
                for fi27 := 0; fi27 < 7; fi27++ {
                    v5 = ("foo bar" + (v5 + v5))
                    switch v5 {
                    case "":
                        if false {
                            switch ((v5 + "hello") + "") {
                            case "a":
                                if (772 < 1) {
                                    var v28 float32 = (1.5 - v2)
                                    delete(v1, (v17 + (v17 * v17)))
                                    if (true || true) {
                                        _ = len(v3)
                                        v29 := GenS2{F0: v5, F1: (true && false), F2: 0}
                                        var v30 int8 = 42
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v31 := GenS3{F0: 376}
                                        if ("hello" < v5) {
                                            v28 += (v28 / v28)
                                            if (65535 == 2) {
                                                if false {
                                                    if false {
                                                        var v32 uint32 = (v17 >> ((255) & 63))
                                                        _ = v32
                                                    } else {
                                                        _ = len(v4)
                                                        var v33 string = ("a" + (((v5 + "foo bar") + "foo bar") + "hello"))
                                                        _ = v33
                                                    }
                                                }
                                                var v34 uint32 = 250
                                                _ = v34
                                            } else {
                                                var v35 int8 = 42
                                                v36 := GenS2{F0: (v5 + ((v5 + v5) + v5)), F1: true, F2: v35}
                                                if true {
                                                    delete(v4, -413)
                                                    var v37 uint8 = 42
                                                    _ = v37
                                                } else {
                                                    var v38 uint64 = 889
                                                    _ = v38
                                                }
                                                _ = v35
                                                _ = v36
                                            }
                                        } else {
                                            var v39 bool = (false && (v5 >= ("hello" + v5)))
                                            v39 = false
                                            _ = v39
                                        }
                                        _ = v31
                                    }
                                    _ = v28
                                } else {
                                    v40 := GenS3{F0: 416}
                                    _ = v40
                                }
                            case "foo bar":
                                v41 := v4[(^(fi21))]
                                _ = v41
                            default:
                                switch (("foo bar" + v5) + "hello") {
                                case "":
                                    v42 := GenS3{F0: 2}
                                    _ = v42
                                default:
                                    v2 = ((v2 - ((v2 / (float32(1e10) + v2)) * v2)) - (v2 + v2))
                                    v3 = append(v3, (v17 * v17))
                                    if (fi27 != (fi27 * (fi21 >> ((444) & 63)))) {
                                        fmt.Println(v17, fi27)
                                        if ((^(v17)) < v17) {
                                            if false {
                                                fmt.Println(v5, v17)
                                                switch (fi27 * fi21) {
                                                case -1:
                                                    if (false && true) {
                                                        var v43 uint8 = 66
                                                        _ = v43
                                                    } else {
                                                        v5 = "a"
                                                        v44 := GenS5{F0: 738}
                                                        var v45 int8 = 1
                                                        var v46 float64 = 1.5
                                                        if true {
                                                            switch (-(fi27)) {
                                                            case 32767:
                                                                if false {
                                                                    fmt.Println(v2, v17, v5)
                                                                }
                                                            case 32768:
                                                                v2 += v2
                                                            case -546:
                                                                var v47 int8 = (^(v45))
                                                                _ = v47
                                                            default:
                                                                if false {
                                                                    v48 := v4[(fi27 | 32768)]
                                                                    v45 -= (v45 - v45)
                                                                    v49 := GenS4{F0: -323}
                                                                    var v50 bool = false
                                                                    var v51 uint16 = 42
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                    _ = v51
                                                                }
                                                                if false {
                                                                    v52 := GenS5{F0: 10}
                                                                    v53 := GenS5{F0: 1}
                                                                    var v54 string = v5
                                                                    var v55 uint64 = 4294967295
                                                                    var v56 string = v54
                                                                    v46 -= (v46 - v46)
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                    _ = v55
                                                                    _ = v56
                                                                } else {
                                                                    fmt.Println(v5, fi27)
                                                                    var v57 uint8 = 84
                                                                    if false {
                                                                        var v58 int64 = 1
                                                                        if false {
                                                                            switch fi21 {
                                                                            case 32767:
                                                                                delete(v1, v17)
                                                                            case 917:
                                                                                v45 += genFn1(127)
                                                                            case 1:
                                                                                v59 := GenS6{F0: v17, F1: (v58 / ((v58) | 1))}
                                                                                _ = v59
                                                                            }
                                                                            if (!(false || (((v58 ^ 42) > (^(v58))) && (252 <= 2)))) {
                                                                                var v60 bool = ((v45 < v45) && false)
                                                                                if (!v60) {
                                                                                    fmt.Println(v45, fi21, v17)
                                                                                    v57 = v57
                                                                                }
                                                                                _ = v60
                                                                            } else {
                                                                                if false {
                                                                                    v45 += v45
                                                                                    fmt.Println(v58)
                                                                                    switch (fi27 * fi21) {
                                                                                    case -316:
                                                                                        v58 = v58
                                                                                    case -280:
                                                                                        fmt.Println(fi21)
                                                                                    case 0:
                                                                                        var v61 uint = 752
                                                                                        _ = v61
                                                                                    default:
                                                                                        if false {
                                                                                            var v62 uint8 = uint8(v17)
                                                                                            _ = v62
                                                                                        }
                                                                                        var v63 int8 = (v45 | genFn1(genFn1(100)))
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = (^(v17))
                                                                                        }
                                                                                        _ = v63
                                                                                    }
                                                                                    if true {
                                                                                        v2 = 3.14159
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi21, v5, fi27)
                                                                                }
                                                                            }
                                                                            delete(v1, v17)
                                                                            if false {
                                                                                switch fi27 {
                                                                                case 127:
                                                                                    v45 += (v45 ^ (v45 / ((genFn1((v45 << ((65535) & 63)))) | 1)))
                                                                                case 595:
                                                                                    _ = len(v3)
                                                                                    v64 := GenS7{F0: (fi21 * fi21)}
                                                                                    if (!false) {
                                                                                        switch -172 {
                                                                                        case 887:
                                                                                            v58 -= v58
                                                                                        case 4294967295:
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                        switch (fi27 & fi21) {
                                                                                        case 564:
                                                                                            v65 := GenS8{F0: v57, F1: 2}
                                                                                            _ = v65
                                                                                        case 2147483647:
                                                                                            var v66 float32 = v2
                                                                                            _ = v66
                                                                                        case -827:
                                                                                            var v67 int = 32768
                                                                                            _ = v67
                                                                                        default:
                                                                                            switch ((v5 + v5) + "\n\t") {
                                                                                            case "hello":
                                                                                                if (!true) {
                                                                                                    _ = len(v3)
                                                                                                    v46 = (3.14159 / ((float64(v2) * (3.14159 - (v46 * 3.14159))) - v46))
                                                                                                } else {
                                                                                                    switch int(v17) {
                                                                                                    case 128:
                                                                                                        var v68 int16 = 738
                                                                                                        _ = v68
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        v57 += 127
                                                                                    } else {
                                                                                        if (true && false) {
                                                                                            if (!(v2 >= float32(-2.25))) {
                                                                                                v69 := GenS9{F0: 1, F1: 1}
                                                                                                var v70 uint32 = v17
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            }
                                                                                        } else {
                                                                                            v58 += 42
                                                                                            if true {
                                                                                                var v71 float64 = float64(((3.14159 + v2) * v2))
                                                                                                if (v57 != (v57 >> ((42) & 63))) {
                                                                                                    var v72 int32 = 0
                                                                                                    v57 += (v57 | v57)
                                                                                                    _ = v72
                                                                                                }
                                                                                                switch 32768 {
                                                                                                case -787:
                                                                                                    v73 := GenS10{F0: 459, F1: 0, F2: 2}
                                                                                                    _ = v73
                                                                                                case -1:
                                                                                                    if len(v3) > 0 {
                                                                                                        v3[0] = 42
                                                                                                    }
                                                                                                    if false {
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                        fmt.Println(fi21, v17)
                                                                                                        v2 -= v2
                                                                                                    } else {
                                                                                                        var v74 bool = (18446744073709551615 > 18446744073709551615)
                                                                                                        v75 := GenS1{F0: (v17 | (v17 % ((1) | 1)))}
                                                                                                        _ = v74
                                                                                                        _ = v75
                                                                                                    }
                                                                                                default:
                                                                                                    v58 = (v58 % (((-(v58))) | 1))
                                                                                                }
                                                                                                _ = v71
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi21, fi21, fi21)
                                                                                    _ = v64
                                                                                case 32767:
                                                                                    var v76 float32 = -2.25
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            var v77 uint64 = 928
                                                                            _ = v77
                                                                        }
                                                                        _ = v58
                                                                    } else {
                                                                        var v78 string = "a"
                                                                        v57 -= v57
                                                                        v3 = append(v3, (v17 & 841))
                                                                        _ = v78
                                                                    }
                                                                    _ = v57
                                                                }
                                                                var v79 bool = false
                                                                _ = v79
                                                            }
                                                            var v80 int64 = 125
                                                            v46 -= v46
                                                            _ = v80
                                                        } else {
                                                            v46 += float64(1.5)
                                                            switch v5 {
                                                            case "\n\t":
                                                                var v81 int64 = -1
                                                                _ = v81
                                                            case "a":
                                                                var v82 string = (v5 + ("a" + ((((v5 + "a") + "a") + v5) + (((v5 + (v5 + v5)) + v5) + v5))))
                                                                _ = v82
                                                            default:
                                                                if ((true || (0 <= 42)) || (true && (3.14159 <= (0.1 * ((v46 + 1.5) * float64(v2)))))) {
                                                                    fmt.Println(v5, v17)
                                                                    if ((32767 != 32767) || false) {
                                                                        v83 := v1[v17]
                                                                        var v84 uint64 = 42
                                                                        v84 = 18446744073709551615
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                    if (-170 != 42) {
                                                                        _ = len(v1)
                                                                    } else {
                                                                        if true {
                                                                            switch fi21 {
                                                                            case -2147483648:
                                                                                var v85 float64 = v46
                                                                                _ = v85
                                                                            case 42:
                                                                                _ = len(v1)
                                                                                if (!true) {
                                                                                    switch "\n\t" {
                                                                                    case "a":
                                                                                        v86 := GenS11{F0: (v17 | v17), F1: 254}
                                                                                        _ = v86
                                                                                    case "\n\t":
                                                                                        fmt.Println(v46, v45, v2)
                                                                                        _ = len(v4)
                                                                                    }
                                                                                    var v87 float32 = v2
                                                                                    fmt.Println(v46, v17, v2)
                                                                                    _ = v87
                                                                                }
                                                                                v88 := GenS12{F0: 784, F1: (v45 + (v45 | v45)), F2: (v17 % (((v17 << ((615) & 63))) | 1))}
                                                                                v1[(v17 ^ v17)] = (636 < -562)
                                                                                if true {
                                                                                    v89 := GenS4{F0: -258}
                                                                                    _ = v89
                                                                                }
                                                                                if false {
                                                                                    if (1 >= 1) {
                                                                                        v2 = v2
                                                                                        if true {
                                                                                            if false {
                                                                                                fmt.Println(v46, v45)
                                                                                            }
                                                                                            v17 -= (^(v17))
                                                                                            v46 = 1e10
                                                                                            switch 2147483647 {
                                                                                            case 1:
                                                                                                if false {
                                                                                                    v2 += (v2 - ((v2 + v2) / v2))
                                                                                                    v2 = v2
                                                                                                } else {
                                                                                                    v90 := GenS1{F0: v17}
                                                                                                    v17 = 65535
                                                                                                    _ = v90
                                                                                                }
                                                                                            case 4294967295:
                                                                                                v4[(-(fi27))] = (((1 <= 255) || true) || false)
                                                                                            case -73:
                                                                                                var v91 uint = 4294967295
                                                                                                _ = v91
                                                                                            default:
                                                                                                v17 += (v17 * v17)
                                                                                            }
                                                                                        }
                                                                                        var v92 uint16 = 32767
                                                                                        var v93 uint64 = 2
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 116
                                                                                        }
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                    var v94 float64 = (v46 - 0.1)
                                                                                    if ((true || true) || false) {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            var v95 int = fi27
                                                                                            _ = v95
                                                                                        case "foo bar":
                                                                                            if ((!true) && true) {
                                                                                                if (false || false) {
                                                                                                    fmt.Println(v45, fi27, fi21)
                                                                                                    if true {
                                                                                                        var v96 float32 = v2
                                                                                                        var v97 float64 = float64((-2.25 + (v2 * (v2 - v2))))
                                                                                                        var v98 int32 = int32(v45)
                                                                                                        v4[(fi27 + (fi27 & fi21))] = ((-(v98)) >= v98)
                                                                                                        _ = v96
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    v99 := GenS4{F0: 42}
                                                                                                    _ = v99
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        switch ((v5 + "a") + (v5 + v5)) {
                                                                                        case "hello":
                                                                                            if (v5 == v5) {
                                                                                                var v100 uint32 = 2
                                                                                                switch (^(fi21)) {
                                                                                                case -167:
                                                                                                    var v101 int64 = 189
                                                                                                    _ = v101
                                                                                                case 1:
                                                                                                    var v102 bool = false
                                                                                                    _ = v102
                                                                                                case -327:
                                                                                                    var v103 float32 = v2
                                                                                                    _ = v103
                                                                                                }
                                                                                                _ = v100
                                                                                            } else {
                                                                                                var v104 uint = 1
                                                                                                _ = v104
                                                                                            }
                                                                                        case "":
                                                                                            _ = len(v4)
                                                                                        case "foo bar":
                                                                                            switch (v5 + ("\n\t" + v5)) {
                                                                                            case "\n\t":
                                                                                                v45 = -1
                                                                                            case "":
                                                                                                var v105 int16 = 700
                                                                                                _ = v105
                                                                                            default:
                                                                                                switch fi27 {
                                                                                                case 642:
                                                                                                    if (v2 > 1.5) {
                                                                                                        v5 = (v5 + "\n\t")
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v2 = float32(-2.25)
                                                                                                        if false {
                                                                                                            switch fi27 {
                                                                                                            case 128:
                                                                                                                if true {
                                                                                                                    v106 := GenS13{F0: genFn0((fi21 | 0))}
                                                                                                                    v107 := GenS14{F0: (255 >= 42), F1: 685, F2: 4294967295}
                                                                                                                    _ = v106
                                                                                                                    _ = v107
                                                                                                                } else {
                                                                                                                    v108 := GenS4{F0: 1}
                                                                                                                    fmt.Println(fi27, fi21)
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                            case 0:
                                                                                                                switch (fi21 ^ (fi21 * fi21)) {
                                                                                                                case 127:
                                                                                                                    v2 -= v2
                                                                                                                    v109 := v4[(^(fi27))]
                                                                                                                    var v110 int16 = 0
                                                                                                                    _ = v109
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                            case 32768:
                                                                                                                var v111 float64 = ((v46 * v46) * v94)
                                                                                                                _ = v111
                                                                                                            }
                                                                                                        }
                                                                                                        if true {
                                                                                                            v94 = ((v46 - 1.5) - v94)
                                                                                                        } else {
                                                                                                            if (!false) {
                                                                                                                switch v5 {
                                                                                                                case "\n\t":
                                                                                                                    if (131 <= 127) {
                                                                                                                        v112 := GenS7{F0: (-(fi27))}
                                                                                                                        v94 = (v94 - v94)
                                                                                                                        v46 = v94
                                                                                                                        _ = v112
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v46 += v46
                                                                                                                }
                                                                                                                v45 += -112
                                                                                                                fmt.Println(v5)
                                                                                                            }
                                                                                                        }
                                                                                                        switch 1 {
                                                                                                        case 806:
                                                                                                            fmt.Println(fi27)
                                                                                                        default:
                                                                                                            switch (fi21 - fi21) {
                                                                                                            case 4294967295:
                                                                                                                if true {
                                                                                                                    switch fi27 {
                                                                                                                    case -892:
                                                                                                                        v113 := GenS5{F0: 190}
                                                                                                                        _ = v113
                                                                                                                    case 2147483647:
                                                                                                                        var v114 uint = 42
                                                                                                                        _ = v114
                                                                                                                    case -382:
                                                                                                                        _ = len(v3)
                                                                                                                        var v115 int16 = 1000
                                                                                                                        _ = v115
                                                                                                                    default:
                                                                                                                        var v116 uint32 = (v17 >> ((468) & 63))
                                                                                                                        _ = v116
                                                                                                                    }
                                                                                                                }
                                                                                                            case 0:
                                                                                                                v1[(v17 + (^(v17)))] = false
                                                                                                            default:
                                                                                                                v117 := GenS10{F0: 1, F1: 0, F2: 889}
                                                                                                                _ = v117
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            v94 = v46
                                                                                            v118 := GenS5{F0: 82}
                                                                                            v5 = "foo bar"
                                                                                            if len(v3) > 0 {
                                                                                                v3[0] = (v17 >> ((0) & 63))
                                                                                            }
                                                                                            var v119 uint16 = 990
                                                                                            _ = v118
                                                                                            _ = v119
                                                                                        } else {
                                                                                            if true {
                                                                                                var v120 string = v5
                                                                                                _ = v120
                                                                                            } else {
                                                                                                var v121 string = v5
                                                                                                if (-1 != 0) {
                                                                                                    v121 = "a"
                                                                                                } else {
                                                                                                    v17 = v17
                                                                                                    var v122 int32 = 502
                                                                                                    v123 := GenS8{F0: 1, F1: 450}
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                }
                                                                                                v121 = v121
                                                                                                _ = v121
                                                                                            }
                                                                                            var v124 int16 = -467
                                                                                            var v125 int = (fi21 << ((255) & 63))
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                        }
                                                                                    }
                                                                                    v46 = (-2.25 - 1.5)
                                                                                    var v126 int = (fi27 >> ((uint(v17)) & 63))
                                                                                    _ = v126
                                                                                    _ = v94
                                                                                } else {
                                                                                    v5 = v5
                                                                                }
                                                                                v5 = v5
                                                                                _ = v88
                                                                            case -9223372036854775808:
                                                                                fmt.Println(v2)
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(v45)
                                                                                if true {
                                                                                    if false {
                                                                                        fmt.Println(fi21, v2)
                                                                                        if false {
                                                                                            v127 := GenS15{F0: (^(fi27))}
                                                                                            v17 = v17
                                                                                            fmt.Println(v17, v46, v5)
                                                                                            _ = v127
                                                                                        } else {
                                                                                            var v128 int32 = 949
                                                                                            var v129 int = fi27
                                                                                            _ = v128
                                                                                            _ = v129
                                                                                        }
                                                                                        var v130 string = (v5 + "")
                                                                                        _ = v130
                                                                                    } else {
                                                                                        var v131 uint8 = 160
                                                                                        v132 := GenS2{F0: v5, F1: (0 != -1), F2: (v45 / ((int8(v17)) | 1))}
                                                                                        v1[(v17 | (v17 << ((18446744073709551615) & 63)))] = (((0.1 / 0.1) <= float64((float32(v46) / (-2.25 - float32((float64(v2) / ((((0.1 - 0.0) / 1e10) * v46) - -2.25))))))) || (v2 > 0.0))
                                                                                        _ = v131
                                                                                        _ = v132
                                                                                    }
                                                                                    var v133 uint16 = 255
                                                                                    var v134 uint32 = (v17 - v17)
                                                                                    if (!(!(true && (!(!true))))) {
                                                                                        v45 = (^(v45))
                                                                                        fmt.Println(fi27)
                                                                                    }
                                                                                    if false {
                                                                                        if true {
                                                                                            v134 = 42
                                                                                            v46 -= float64(((0.0 - v2) + (v2 - v2)))
                                                                                            _ = len(v3)
                                                                                        } else {
                                                                                            v4[(-(fi21))] = (v17 != (^(v17)))
                                                                                            var v135 int32 = -434
                                                                                            v136 := GenS16{F0: (v45 - (v45 << ((42) & 63)))}
                                                                                            _ = v135
                                                                                            _ = v136
                                                                                        }
                                                                                        v5 = v5
                                                                                    } else {
                                                                                        if (-1 <= 1000) {
                                                                                            var v137 string = "a"
                                                                                            var v138 uint8 = 128
                                                                                            var v139 uint32 = 499
                                                                                            if true {
                                                                                                v5 = v5
                                                                                            }
                                                                                            var v140 int32 = 2147483647
                                                                                            _ = v137
                                                                                            _ = v138
                                                                                            _ = v139
                                                                                            _ = v140
                                                                                        }
                                                                                    }
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                }
                                                                                if (255 > 65535) {
                                                                                    fmt.Println(v17, fi21)
                                                                                }
                                                                            }
                                                                            v141 := GenS17{F0: (2 != 0), F1: fi21}
                                                                            _ = v141
                                                                        }
                                                                        var v142 int8 = 1
                                                                        var v143 int16 = -32768
                                                                        _ = v142
                                                                        _ = v143
                                                                    }
                                                                    v2 += (v2 / v2)
                                                                }
                                                            }
                                                            switch (fi21 / (((-(fi21))) | 1)) {
                                                            case 0:
                                                                var v144 uint = 573
                                                                if (166 < 127) {
                                                                    fmt.Println(v46, v17, fi27)
                                                                    var v145 uint16 = genFn0(9223372036854775807)
                                                                    _ = v145
                                                                }
                                                                v146 := GenS10{F0: (^(v144)), F1: 1, F2: v144}
                                                                var v147 uint16 = 1
                                                                _ = v144
                                                                _ = v146
                                                                _ = v147
                                                            case 2147483647:
                                                                switch ("\n\t" + v5) {
                                                                case "\n\t":
                                                                    v45 -= v45
                                                                case "hello":
                                                                    v2 += 1e10
                                                                case "":
                                                                    if (0 == 1) {
                                                                        v5 = v5
                                                                    }
                                                                    var v148 int64 = 9223372036854775807
                                                                    _ = v148
                                                                }
                                                            default:
                                                                if false {
                                                                    v149 := GenS6{F0: (^(v17)), F1: int64(v17)}
                                                                    if false {
                                                                        if true {
                                                                            if false {
                                                                                if ((964 >= 42) && (!true)) {
                                                                                    v3 = append(v3, (v17 | (v17 | v17)))
                                                                                }
                                                                                v150 := GenS16{F0: v45}
                                                                                var v151 int16 = 42
                                                                                _ = v150
                                                                                _ = v151
                                                                            }
                                                                        } else {
                                                                            var v152 int8 = 57
                                                                            switch (fi27 % ((fi27) | 1)) {
                                                                            case -2147483648:
                                                                                fmt.Println(v45, v152, v2)
                                                                            case 462:
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                delete(v4, fi21)
                                                                            }
                                                                            v46 += v46
                                                                            _ = v152
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi27, v2, v17)
                                                                        fmt.Println(v5, v45)
                                                                        if (genFn0((fi27 & -9223372036854775808)) < 410) {
                                                                            switch v5 {
                                                                            case "hello":
                                                                                if (!false) {
                                                                                    var v153 float64 = (((v46 - 1.5) + v46) / float64(v2))
                                                                                    v3 = append(v3, v17)
                                                                                    _ = v153
                                                                                }
                                                                                v17 = v17
                                                                                v4[32768] = (!(genFn0(int(v45)) > 0))
                                                                                if (-32768 > 0) {
                                                                                    switch v5 {
                                                                                    case "foo bar":
                                                                                        v5 = "foo bar"
                                                                                    }
                                                                                } else {
                                                                                    v5 = (v5 + v5)
                                                                                }
                                                                                if (v45 != v45) {
                                                                                    if (true && false) {
                                                                                        v5 = ("\n\t" + v5)
                                                                                        var v154 uint16 = 533
                                                                                        if true {
                                                                                            var v155 bool = (v45 < (v45 - (^(v45))))
                                                                                            v156 := GenS18{F0: 100, F1: 525}
                                                                                            v45 -= v45
                                                                                            _ = v155
                                                                                            _ = v156
                                                                                        } else {
                                                                                            v157 := GenS6{F0: 65535, F1: 9223372036854775807}
                                                                                            v158 := GenS5{F0: 0}
                                                                                            _ = v157
                                                                                            _ = v158
                                                                                        }
                                                                                        _ = v154
                                                                                    } else {
                                                                                        v159 := GenS3{F0: 940}
                                                                                        _ = v159
                                                                                    }
                                                                                }
                                                                                fmt.Println(v5, v46, v5)
                                                                            case "a":
                                                                                v45 = (v45 & 42)
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v149
                                                                }
                                                            }
                                                        }
                                                        v2 -= float32((v46 * (0.0 / (v46 / float64(float32(v46))))))
                                                        if (v17 != (v17 << ((1) & 63))) {
                                                            v3 = append(v3, v17)
                                                            switch "\n\t" {
                                                            case "a":
                                                                v17 = 2
                                                            case "\n\t":
                                                                v160 := GenS19{F0: (^(fi27)), F1: (fi27 * (fi27 << ((312) & 63))), F2: 0}
                                                                _ = v160
                                                            default:
                                                                v2 = 1.5
                                                            }
                                                            var v161 int = fi27
                                                            _ = v161
                                                        } else {
                                                            var v162 float32 = float32(((v46 + ((float64(float32(1.5)) + (v46 + (1.5 + v46))) * ((v46 / 0.1) + (3.14159 - v46)))) - (v46 - (v46 / v46))))
                                                            v163 := GenS3{F0: 139}
                                                            _ = v162
                                                            _ = v163
                                                        }
                                                        _ = v44
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case 0:
                                                    if (false && false) {
                                                        var v164 uint16 = 119
                                                        var v165 bool = true
                                                        switch 0 {
                                                        case 2147483647:
                                                            fmt.Println(v164, fi27, v5)
                                                        case -2147483648:
                                                            var v166 int = 0
                                                            var v167 int8 = 127
                                                            _ = v166
                                                            _ = v167
                                                        case -607:
                                                            if (-9223372036854775808 >= 0) {
                                                                var v168 string = (("hello" + "a") + "foo bar")
                                                                var v169 uint = 1
                                                                _ = v168
                                                                _ = v169
                                                            }
                                                        default:
                                                            var v170 int32 = -2147483648
                                                            _ = v170
                                                        }
                                                        _ = v164
                                                        _ = v165
                                                    } else {
                                                        delete(v1, v17)
                                                        fmt.Println(v2, fi21, fi27)
                                                    }
                                                case -2147483648:
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(v17))
                                                    }
                                                default:
                                                    if ((fi27 / ((486) | 1)) != fi27) {
                                                        v171 := GenS20{F0: 984, F1: (v5 + ("a" + "\n\t"))}
                                                        fmt.Println(v5)
                                                        v172 := GenS21{F0: ("foo bar" + v5)}
                                                        _ = v171
                                                        _ = v172
                                                    } else {
                                                        var v173 int8 = -128
                                                        v174 := GenS7{F0: -1}
                                                        _ = v173
                                                        _ = v174
                                                    }
                                                }
                                            }
                                            v175 := GenS22{F0: 0, F1: false, F2: fi21}
                                            switch (-(fi27)) {
                                            case 412:
                                                var v176 int16 = -171
                                                _ = v176
                                            case -1:
                                                var v177 uint16 = 162
                                                _ = v177
                                            case -36:
                                                v178 := GenS10{F0: 65535, F1: 42, F2: 1}
                                                _ = v178
                                            }
                                            _ = v175
                                        }
                                        v179 := GenS9{F0: genFn1(32), F1: -32768}
                                        var v180 int16 = -1
                                        v181 := GenS23{F0: 18446744073709551615, F1: true}
                                        v180 -= v180
                                        _ = v179
                                        _ = v180
                                        _ = v181
                                    } else {
                                        fmt.Println(fi27, fi27, fi27)
                                        var v182 uint64 = 1
                                        _ = v182
                                    }
                                }
                            }
                            if (!(!true)) {
                                var v183 uint8 = 255
                                _ = v183
                            }
                            switch 804 {
                            case 16:
                                v4[(fi21 >> ((65535) & 63))] = (-1 < 2147483647)
                            case -589:
                                _ = len(v1)
                            default:
                                v2 = (v2 * v2)
                            }
                        } else {
                            fmt.Println(v17, fi21)
                            switch v5 {
                            case "hello":
                                switch fi27 {
                                case 565:
                                    v184 := GenS9{F0: 0, F1: -1}
                                    _ = v184
                                }
                            case "\n\t":
                                if (((((v5 + "\n\t") + "hello") + (v5 + v5)) + "") < (("a" + "a") + "foo bar")) {
                                    var v185 int32 = -332
                                    _ = v185
                                }
                            }
                        }
                    case "a":
                        var v186 string = (((((v5 + ("a" + ((("foo bar" + v5) + "") + v5))) + "a") + v5) + v5) + ("foo bar" + (v5 + ("" + (((((("\n\t" + "foo bar") + v5) + v5) + "hello") + ("foo bar" + v5)) + "")))))
                        _ = v186
                    default:
                        v5 = ("\n\t" + ((v5 + (((v5 + v5) + v5) + (v5 + "\n\t"))) + v5))
                    }
                    var v187 uint32 = v17
                    _ = fi27
                    _ = v187
                }
                v17 = (v17 >> ((uint(v17)) & 63))
                _ = fi21
            }
        case "":
            v188 := GenS14{F0: (true || (uint16(v17) <= genFn0(32768))), F1: 42, F2: 903}
            _ = v188
        case "hello":
            for fi189 := 0; fi189 < 14; fi189++ {
                switch 42 {
                case 9223372036854775807:
                    for fi190 := 0; fi190 < 16; fi190++ {
                        var v191 float32 = (v2 + v2)
                        var v192 uint16 = 32767
                        var v193 int64 = 0
                        _ = fi190
                        _ = v191
                        _ = v192
                        _ = v193
                    }
                case 32768:
                    fmt.Println(v5, v2, v2)
                default:
                    v2 += 0.1
                }
                _ = fi189
            }
        }
        _ = v20
    }
    v2 -= (v2 + v2)
    v194 := v1[(v17 >> ((2) & 63))]
    for fi195 := 0; fi195 < 10; fi195++ {
        if true {
            v5 = v5
        }
        for fi196 := 0; fi196 < 7; fi196++ {
            var v197 uint64 = 857
            _ = fi196
            _ = v197
        }
        _ = fi195
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v194:", v194)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v194
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
