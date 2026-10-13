package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 bool; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32; F1 bool; F2 int64 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 uint) uint16 {
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 18; fi7++ {
        v6 = 255
        p0 = p0
        v8 := GenS1{F0: false}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: (!false)}
        switch 751 {
        case 2147483647:
            v11 := []int{32767, -35, 73}
            _ = v11
        case -9223372036854775808:
            var v12 uint8 = v6
            _ = v12
        case 695:
            switch "\n\t" {
            case "\n\t":
                var v13 int8 = 81
                _ = v13
            case "hello":
                for fi14 := 0; fi14 < 12; fi14++ {
                    var v15 float64 = (0.0 * 0.1)
                    switch fi9 {
                    case -974:
                        switch "foo bar" {
                        case "a":
                            v16 := []int{2147483647, fi14, -37, fi9}
                            _ = v16
                        default:
                            v17 := []int{(fi14 | fi14), 0}
                            _ = v17
                        }
                    case 0:
                        switch 0 {
                        case 32767:
                            fmt.Println(fi14, v15)
                        }
                    default:
                        v18 := []int{fi14, fi14}
                        _ = v18
                    }
                    v19 := GenS1{F0: false}
                    _ = fi14
                    _ = v15
                    _ = v19
                }
            default:
                var v20 uint32 = 568
                _ = v20
            }
        default:
            var v21 int = fi9
            _ = v21
        }
        _ = fi9
        _ = v10
    }
    p0 = 255
    _ = p0
    _ = v6
    return 343
}
func genFn1(p0 uint16, p1 bool, p2 bool) uint32 {
    var v22 int32 = -1
    if p1 {
        v23 := GenS2{F0: 0, F1: 699, F2: 80}
        v24 := GenS1{F0: p1}
        fmt.Println(p0)
        _ = v23
        _ = v24
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 0
}

func main() {
    var v1 string = ("hello" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t"))))
    var v2 int64 = -1
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    var v5 uint = 255
    fmt.Println(v2, v1)
    for fi25 := 0; fi25 < 17; fi25++ {
        v26 := GenS2{F0: 0, F1: uint64(v5), F2: 729}
        var v27 uint64 = 0
        for fi28 := 0; fi28 < 2; fi28++ {
            if ((^(v5)) > v5) {
                switch (v1 + "hello") {
                case "\n\t":
                    switch v1 {
                    case "\n\t":
                        if true {
                            switch (v1 + v1) {
                            case "a":
                                var v29 int8 = 0
                                _ = v29
                            }
                            v30 := GenS2{F0: -1, F1: (v27 | 1), F2: 42}
                            _ = v30
                        }
                    case "hello":
                        if (65535 > 65535) {
                            v31 := GenS3{F0: -128, F1: (v27 - 18446744073709551615), F2: (v27 * v27)}
                            v32 := GenS2{F0: 32767, F1: v27, F2: 19}
                            v33 := []int{fi28}
                            v27 = (^(v27))
                            var v34 int16 = 1000
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS4{F0: -106}
                            var v36 float32 = float32((3.14159 - -2.25))
                            var v37 float64 = 1e10
                            switch (fi25 * 128) {
                            case 2147483647:
                                if false {
                                    v3["foo bar"] = 42
                                } else {
                                    if (255 == 169) {
                                        fmt.Println(fi28, fi28, fi25)
                                        switch ((v1 + "hello") + v1) {
                                        case "a":
                                            var v38 uint = v5
                                            _ = v38
                                        default:
                                            fmt.Println(v27)
                                        }
                                        var v39 int8 = int8(v27)
                                        _ = v39
                                    }
                                    if false {
                                        v2 += 9223372036854775807
                                    }
                                }
                            case 1:
                                var v40 string = v1
                                _ = v40
                            case 32767:
                                v41 := []int{fi28, fi28}
                                _ = v41
                            }
                            var v42 int16 = 154
                            var v43 int64 = (v2 % ((v2) | 1))
                            _ = v35
                            _ = v36
                            _ = v37
                            _ = v42
                            _ = v43
                        }
                    case "foo bar":
                        v4[((v1 + v1) + "a")] = fi25
                    default:
                        if true {
                            var v44 int32 = 42
                            var v45 uint8 = 25
                            switch fi28 {
                            case -744:
                                _ = len(v3)
                            case 2147483647:
                                switch 2147483647 {
                                case -853:
                                    v5 += v5
                                case -9223372036854775808:
                                    fmt.Println(v5, v27, v1)
                                default:
                                    _ = len(v4)
                                }
                            case 4294967295:
                                if false {
                                    v46 := GenS5{F0: (1 > 127), F1: (1e10 < float32(0.0)), F2: false}
                                    v47 := GenS1{F0: ((0 == 0) || (false && (v27 != (^(v27)))))}
                                    _ = v46
                                    _ = v47
                                }
                                fmt.Println(fi25, v44, fi25)
                                v44 = v44
                            }
                            if true {
                                v1 = v1
                                var v48 bool = (((1e10 / 1e10) / (3.14159 + 0.1)) > -2.25)
                                if v48 {
                                    if (((1.5 * 0.1) / 1e10) == ((1e10 / -2.25) * 1.5)) {
                                        v49 := []int{fi25}
                                        fmt.Println(v45, v27, v48)
                                        _ = v49
                                    } else {
                                        if ((v48 || v48) || v48) {
                                            v50 := GenS3{F0: -44, F1: 0, F2: 1}
                                            if (((fi25 << (((^(v5))) & 63)) > fi25) && false) {
                                                switch (fi25 % ((fi28) | 1)) {
                                                case 32768:
                                                    v1 = "foo bar"
                                                case -1:
                                                    _ = len(v3)
                                                    fmt.Println(v2)
                                                case 4294967295:
                                                    switch ((((v1 + v1) + v1) + "foo bar") + v1) {
                                                    case "hello":
                                                        v1 = "a"
                                                    case "\n\t":
                                                        v51 := GenS3{F0: 1, F1: v27, F2: 42}
                                                        _ = v51
                                                    case "foo bar":
                                                        var v52 uint8 = 255
                                                        if v48 {
                                                            v45 = v52
                                                        } else {
                                                            var v53 uint32 = 2147483648
                                                            _ = v53
                                                        }
                                                        fmt.Println(v2)
                                                        _ = v52
                                                    default:
                                                        if v48 {
                                                            v45 = 226
                                                            _ = len(v3)
                                                            v54 := GenS2{F0: 0, F1: (v27 << ((v5) & 63)), F2: 65535}
                                                            _ = v54
                                                        } else {
                                                            if false {
                                                                switch (v1 + v1) {
                                                                case "":
                                                                    v55 := GenS2{F0: 0, F1: v27, F2: 2}
                                                                    _ = v55
                                                                }
                                                                v56 := GenS5{F0: true, F1: (127 > -1), F2: true}
                                                                if v48 {
                                                                    var v57 int64 = v2
                                                                    _ = v57
                                                                }
                                                                _ = v56
                                                            } else {
                                                                var v58 uint8 = 42
                                                                var v59 int32 = (v44 >> ((v5) & 63))
                                                                fmt.Println(fi28, fi28)
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            v5 += 0
                                                        }
                                                    }
                                                default:
                                                    v2 -= (v2 + v2)
                                                }
                                                switch (v1 + (((v1 + v1) + v1) + "a")) {
                                                case "\n\t":
                                                    v60 := GenS6{F0: v5}
                                                    _ = v60
                                                case "a":
                                                    var v61 float32 = 1.5
                                                    _ = v61
                                                }
                                                if v48 {
                                                    fmt.Println(v1, v44, v5)
                                                    if v48 {
                                                        switch fi28 {
                                                        case 32767:
                                                            v62 := GenS3{F0: -128, F1: v27, F2: (v27 ^ 223)}
                                                            _ = v62
                                                        }
                                                        if false {
                                                            v48 = (v48 && (fi28 <= (fi28 << ((v5) & 63))))
                                                            v48 = true
                                                            delete(v4, v1)
                                                        } else {
                                                            fmt.Println(v1)
                                                            var v63 string = (((v1 + v1) + v1) + v1)
                                                            if (fi28 != fi28) {
                                                                v64 := []int{(fi25 & (fi28 >> ((326) & 63))), (fi25 % ((fi25) | 1)), (fi28 % ((-1) | 1))}
                                                                v65 := GenS4{F0: 100}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            _ = v63
                                                        }
                                                    } else {
                                                        var v66 bool = true
                                                        var v67 string = "\n\t"
                                                        v68 := []int{fi28, (fi25 << ((v5) & 63)), 32767}
                                                        _ = v66
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                    var v69 uint64 = v27
                                                    _ = v69
                                                } else {
                                                    v70 := GenS2{F0: -32768, F1: v27, F2: 0}
                                                    _ = v70
                                                }
                                            }
                                            _ = v50
                                        }
                                    }
                                    fmt.Println(v5, v5, v45)
                                    fmt.Println(v45)
                                    var v71 int8 = 42
                                    var v72 int64 = 9223372036854775807
                                    switch "a" {
                                    case "foo bar":
                                        if ((fi25 | (fi25 | fi28)) > fi25) {
                                            var v73 int16 = -32768
                                            v27 -= (v27 << ((166) & 63))
                                            v27 = v27
                                            _ = v73
                                        }
                                    }
                                    var v74 int = fi28
                                    _ = v71
                                    _ = v72
                                    _ = v74
                                }
                                _ = v48
                            }
                            _ = v44
                            _ = v45
                        }
                    }
                default:
                    if (4294967295 < 0) {
                        v75 := []int{(fi28 << (((v5 - (^(v5)))) & 63)), 0}
                        v76 := GenS2{F0: -32768, F1: v27, F2: 1}
                        _ = v75
                        _ = v76
                    }
                }
            }
            _ = fi28
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    v77 := v4[v1]
    v5 += v5
    switch v77 {
    case 763:
        v5 += v5
    case -1:
        if (42 <= 222) {
            if false {
                var v78 float32 = float32(3.14159)
                if (255 == 42) {
                    var v79 string = (("foo bar" + (v1 + (v1 + v1))) + (v1 + v1))
                    switch ((v79 + "a") + "hello") {
                    case "\n\t":
                        if false {
                            var v80 bool = true
                            _ = len(v4)
                            _ = v80
                        }
                    case "":
                        for fi81 := 0; fi81 < 8; fi81++ {
                            for fi82 := 0; fi82 < 17; fi82++ {
                                var v83 uint = v5
                                var v84 uint8 = 82
                                v85 := []int{v77, (v77 & (v77 & (fi81 | fi82))), (^(fi81))}
                                _ = fi82
                                _ = v83
                                _ = v84
                                _ = v85
                            }
                            v86 := GenS4{F0: 70}
                            v78 = (1.5 - (float32(0.1) + v78))
                            _ = fi81
                            _ = v86
                        }
                    case "foo bar":
                        _ = len(v3)
                    default:
                        v87 := GenS7{F0: 1}
                        _ = v87
                    }
                    if false {
                        if true {
                            if false {
                                v2 += v2
                                v88 := []int{42}
                                _ = v88
                            } else {
                                var v89 uint16 = 65535
                                var v90 uint8 = 255
                                _ = v89
                                _ = v90
                            }
                            var v91 uint16 = 0
                            _ = v91
                        }
                    }
                    _ = v79
                }
                _ = v78
            } else {
                if ((genFn0((v5 & 255)) <= 32767) || (false && false)) {
                    _ = len(v4)
                }
                var v92 uint16 = 0
                _ = v92
            }
        }
    }
    v2 += 1
    if true {
        var v93 int16 = -32768
        if (false && false) {
            v94 := GenS8{F0: 0, F1: false, F2: int64(v5)}
            var v95 float64 = 0.1
            v96 := []int{v77, v77, v77}
            _ = v94
            _ = v95
            _ = v96
        }
        _ = v93
    }
    for fi97 := 0; fi97 < 15; fi97++ {
        for fi98 := 0; fi98 < 8; fi98++ {
            v5 += 1
            var v99 float32 = 0.1
            _ = fi98
            _ = v99
        }
        _ = fi97
    }
    v2 += v2
    v100 := GenS5{F0: true, F1: false, F2: (v1 > (v1 + (v1 + "foo bar")))}
    v101 := GenS9{F0: -538, F1: 0}
    switch v77 {
    case -1:
        if false {
            switch v1 {
            case "\n\t":
                v2 -= 1
            default:
                for fi102 := 0; fi102 < 8; fi102++ {
                    v103 := []int{(v77 % ((v77) | 1)), (^(v77))}
                    v104 := v4[((v1 + v1) + v1)]
                    fmt.Println(v1)
                    _ = fi102
                    _ = v103
                    _ = v104
                }
            }
            v1 = "foo bar"
        }
    case 0:
        v105 := GenS10{F0: v1, F1: false, F2: genFn0(104)}
        _ = v105
    }
    switch v77 {
    case 1:
        v4[v1] = 26
    }
    var v106 uint16 = genFn0(v5)
    var v107 uint16 = v106
    switch 32768 {
    case 1:
        var v108 uint64 = 140
        _ = v108
    case 2147483647:
        delete(v3, v1)
    case 4294967295:
        fmt.Println(v77, v1, v2)
    }
    v77 = (v77 | v77)
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v100
    _ = v101
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
}
