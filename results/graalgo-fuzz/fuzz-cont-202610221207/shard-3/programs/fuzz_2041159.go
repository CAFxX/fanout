package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS4 struct { F0 int; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int32; F1 uint }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 int8; F1 uint16 }
type GenS14 struct { F0 int16; F1 uint; F2 uint }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int16; F1 int; F2 uint }
func genFn0(p0 uint32, p1 bool, p2 uint64) int64 {
    v4 := GenS1{F0: 208, F1: (p2 + 662), F2: p2}
    var v5 int8 = 15
    p1 = (!p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 string, p1 uint16, p2 int) uint8 {
    var v6 uint = 4294967295
    for fi7 := 0; fi7 < 9; fi7++ {
        if false {
            v8 := []int{-9223372036854775808, int(v6)}
            switch p2 {
            case -266:
                p2 -= -9223372036854775808
            default:
                v6 += 60
            }
            switch fi7 {
            case -9223372036854775808:
                _ = len(v8)
            case -1:
                fmt.Println(p2)
            case 226:
                p1 -= p1
            }
            _ = v8
        }
        _ = fi7
    }
    v6 = v6
    var v9 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 127
}

func main() {
    var v1 uint32 = 966
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]uint32)
    switch 42 {
    case 32768:
        v1 = 1
    }
    v10 := GenS2{F0: 128, F1: -1, F2: 399}
    var v11 float32 = float32((1.5 * (-2.25 / 1.5)))
    var v12 uint16 = 1
    var v13 int64 = genFn0(v1, false, 4294967295)
    v14 := GenS3{F0: v12, F1: 0, F2: 65535}
    if true {
        v15 := GenS3{F0: v12, F1: 1, F2: 298}
        if (0 <= 0) {
            v2 = v13
            for fi16 := 0; fi16 < 12; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := GenS4{F0: 127, F1: 1, F2: (v1 % (((v1 << ((2) & 63))) | 1))}
                    v19 := GenS5{F0: 65535}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v11 = float32((3.14159 / 1.0))
                _ = fi16
            }
        } else {
            var v20 uint = 0
            if true {
                v21 := GenS1{F0: int16(v1), F1: 42, F2: uint64(v20)}
                if (!true) {
                    if (v2 >= (v13 / ((v2) | 1))) {
                        fmt.Println(v1, v2)
                        fmt.Println(v1, v12)
                        fmt.Println(v1, v20, v11)
                    }
                    switch "" {
                    case "\n\t":
                        var v22 uint64 = 4294967295
                        _ = v22
                    case "a":
                        if true {
                            var v23 uint64 = 365
                            v23 = (v23 | v23)
                            for fi24 := 0; fi24 < 9; fi24++ {
                                switch (("" + "a") + "") {
                                case "hello":
                                    for fi25 := 0; fi25 < 14; fi25++ {
                                        v26 := []int{(fi24 << ((0) & 63))}
                                        switch (fi25 ^ 9223372036854775807) {
                                        case 2147483647:
                                            if false {
                                                var v27 int8 = 1
                                                if len(v26) > 0 {
                                                    v26[0] = fi25
                                                }
                                                v20 -= (v20 - (v20 * v20))
                                                _ = v27
                                            } else {
                                                if (false && (1 == v12)) {
                                                    v20 = uint(v23)
                                                    var v28 uint = v20
                                                    _ = v28
                                                } else {
                                                    v29 := GenS6{F0: 1, F1: ("a" + "\n\t")}
                                                    v11 = v11
                                                    _ = v29
                                                }
                                            }
                                        case -2147483648:
                                            v2 = v13
                                        default:
                                            v30 := GenS3{F0: v12, F1: 127, F2: 32767}
                                            _ = v30
                                        }
                                        if false {
                                            v13 = (v2 / (((v13 & v2)) | 1))
                                            var v31 string = "foo bar"
                                            v32 := GenS7{F0: 54, F1: (v23 ^ (v23 * v23)), F2: 718}
                                            _ = v31
                                            _ = v32
                                        }
                                        v33 := GenS2{F0: (fi25 - 128), F1: 127, F2: 2}
                                        v34 := GenS7{F0: 2, F1: uint64(v12), F2: 42}
                                        _ = fi25
                                        _ = v26
                                        _ = v33
                                        _ = v34
                                    }
                                case "a":
                                    var v35 int = (^(fi24))
                                    _ = v35
                                case "\n\t":
                                    fmt.Println(fi24, v1, v23)
                                default:
                                    v12 = (v12 * 0)
                                }
                                for fi36 := 0; fi36 < 20; fi36++ {
                                    var v37 uint32 = 1
                                    var v38 float32 = ((1.5 + (0.1 / v11)) - v11)
                                    var v39 uint64 = (^(v23))
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                _ = fi24
                            }
                            _ = v23
                        }
                    }
                } else {
                    if true {
                        if (!((!(111 != 684)) || (1000 != 1))) {
                            switch 1 {
                            case 669:
                                v1 -= (v1 % (((v1 % ((v1) | 1))) | 1))
                            case -179:
                                var v40 int8 = -1
                                _ = v40
                            }
                            var v41 uint32 = v1
                            for fi42 := 0; fi42 < 16; fi42++ {
                                for fi43 := 0; fi43 < 19; fi43++ {
                                    v44 := GenS2{F0: (^(fi42)), F1: -1, F2: v20}
                                    _ = fi43
                                    _ = v44
                                }
                                switch -311 {
                                case 32768:
                                    delete(v3, "a")
                                }
                                v13 -= (v2 | v13)
                                _ = fi42
                            }
                            _ = v41
                        } else {
                            _ = len(v3)
                        }
                    }
                    var v45 int8 = -26
                    v46 := GenS2{F0: -1, F1: 0, F2: (^(v20))}
                    _ = v45
                    _ = v46
                }
                _ = v21
            }
            fmt.Println(v12, v2)
            _ = v20
        }
        _ = v15
    }
    v47 := GenS8{F0: 61}
    v13 = v2
    var v48 bool = true
    var v49 int = 128
    var v50 uint64 = 140
    switch "foo bar" {
    case "foo bar":
        var v51 int32 = 0
        _ = v51
    case "":
        for fi52 := 0; fi52 < 15; fi52++ {
            var v53 bool = v48
            var v54 uint16 = (v12 / (((v12 * v12)) | 1))
            var v55 int64 = (-(v13))
            _ = fi52
            _ = v53
            _ = v54
            _ = v55
        }
    case "a":
        switch v49 {
        case 32767:
            for fi56 := 0; fi56 < 20; fi56++ {
                v57 := GenS9{F0: v48, F1: 65535}
                delete(v3, "foo bar")
                v58 := GenS7{F0: v50, F1: 42, F2: 42}
                _ = fi56
                _ = v57
                _ = v58
            }
        case -9223372036854775808:
            var v59 uint = 18446744073709551615
            _ = v59
        default:
            switch 1 {
            case 42:
                for fi60 := 0; fi60 < 11; fi60++ {
                    var v61 int16 = -375
                    _ = fi60
                    _ = v61
                }
            case -2147483648:
                if v48 {
                    var v62 uint32 = v1
                    if true {
                        if v48 {
                            fmt.Println(v1, v50)
                            v63 := GenS10{F0: 42, F1: 4294967295}
                            switch "hello" {
                            case "foo bar":
                                v64 := []int{(v49 >> ((382) & 63)), (v49 & (v49 >> ((89) & 63)))}
                                _ = v64
                            case "":
                                v65 := GenS11{F0: 994, F1: -598}
                                _ = v65
                            case "\n\t":
                                v3[("hello" + "foo bar")] = 543
                            }
                            _ = v63
                        } else {
                            v49 = (v49 << ((59) & 63))
                            var v66 uint = uint(v62)
                            if v48 {
                                var v67 bool = v48
                                for fi68 := 0; fi68 < 19; fi68++ {
                                    v69 := []int{-405, v49, (v49 | (v49 * -439)), 2147483647}
                                    if len(v69) >= 2 {
                                        v69 = v69[1:len(v69)]
                                    }
                                    _ = fi68
                                    _ = v69
                                }
                                fmt.Println(v48)
                                _ = v67
                            } else {
                                if (!v48) {
                                    v1 += v62
                                }
                            }
                            _ = v66
                        }
                        switch v49 {
                        case 128:
                            v70 := GenS11{F0: 42, F1: 58}
                            _ = v70
                        case 127:
                            v71 := GenS12{F0: 696, F1: 42}
                            _ = v71
                        default:
                            for fi72 := 0; fi72 < 11; fi72++ {
                                var v73 uint32 = (v62 ^ (^(v62)))
                                _ = fi72
                                _ = v73
                            }
                        }
                    }
                    _ = v62
                } else {
                    for fi74 := 0; fi74 < 10; fi74++ {
                        v75 := GenS13{F0: 98, F1: (v12 << ((367) & 63))}
                        _ = fi74
                        _ = v75
                    }
                    var v76 uint = uint(v50)
                    _ = v76
                }
            case 282:
                if v48 {
                    v49 += int(v13)
                } else {
                    v13 += 1
                    if v48 {
                        fmt.Println(v49, v1)
                        v77 := GenS7{F0: 42, F1: (v50 % (((v50 * uint64(v12))) | 1)), F2: 827}
                        v1 = uint32(v12)
                        _ = v77
                    }
                    var v78 int32 = -805
                    _ = v78
                }
            default:
                switch 32767 {
                case 127:
                    if false {
                        v79 := GenS2{F0: v49, F1: 0, F2: 0}
                        fmt.Println(v50, v13)
                        v80 := []int{v49}
                        _ = v79
                        _ = v80
                    }
                case 4294967295:
                    if v48 {
                        var v81 int8 = 0
                        v82 := GenS5{F0: 701}
                        _ = v81
                        _ = v82
                    } else {
                        if (!((248 >= 2) && (v49 < (v49 | 934)))) {
                            var v83 float32 = (0.1 / 1e10)
                            v1 -= 65535
                            v12 -= (v12 * 42)
                            _ = v83
                        }
                        var v84 int64 = genFn0((v1 >> ((354) & 63)), v48, v50)
                        _ = v84
                    }
                default:
                    switch 128 {
                    case -400:
                        var v85 float64 = 3.14159
                        _ = v85
                    default:
                        switch "" {
                        case "foo bar":
                            for fi86 := 0; fi86 < 12; fi86++ {
                                v13 += 9223372036854775807
                                switch ("\n\t" + "foo bar") {
                                case "\n\t":
                                    v87 := []int{fi86, (fi86 * 42), -889, (v49 * 0)}
                                    _ = v87
                                case "a":
                                    for fi88 := 0; fi88 < 3; fi88++ {
                                        if v48 {
                                            switch (-(fi88)) {
                                            case 42:
                                                _ = len(v3)
                                            case 0:
                                                if true {
                                                    v12 -= (v12 * (v12 - (v12 >> ((0) & 63))))
                                                    v49 = (fi86 / ((v49) | 1))
                                                    if true {
                                                        switch "a" {
                                                        case "foo bar":
                                                            if ((((0.1 / 0.1) == 1e10) && false) || (true || false)) {
                                                                v89 := GenS5{F0: 42}
                                                                var v90 uint = 927
                                                                var v91 int8 = 127
                                                                _ = v89
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                        case "hello":
                                                            var v92 uint32 = uint32(v50)
                                                            _ = v92
                                                        default:
                                                            switch "" {
                                                            case "hello":
                                                                fmt.Println(v12, v1, v12)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v93 float32 = (v11 * (0.0 + v11))
                                                    v49 = v49
                                                    _ = v93
                                                }
                                            case 161:
                                                var v94 uint8 = 2
                                                _ = v94
                                            default:
                                                var v95 int64 = (^(v2))
                                                _ = v95
                                            }
                                        } else {
                                            v96 := GenS1{F0: 0, F1: (v50 << ((1) & 63)), F2: v50}
                                            v11 += 1.5
                                            v12 -= (^(v12))
                                            _ = v96
                                        }
                                        switch fi86 {
                                        case 884:
                                            switch (fi88 - 32768) {
                                            case 184:
                                                v97 := []int{}
                                                _ = v97
                                            case 880:
                                                v98 := GenS14{F0: 1000, F1: 65535, F2: 255}
                                                _ = v98
                                            case 0:
                                                v2 = (v2 << ((65535) & 63))
                                            default:
                                                fmt.Println(v12, v1)
                                            }
                                        }
                                        _ = fi88
                                    }
                                default:
                                    if (0 == 0) {
                                        v48 = ((v48 || false) && v48)
                                        if v48 {
                                            v1 = (^(v1))
                                            if (((fi86 - (v49 % ((128) | 1))) >= fi86) && v48) {
                                                fmt.Println(v1, v12)
                                                v2 -= 42
                                            } else {
                                                if v48 {
                                                    v99 := GenS12{F0: -32768, F1: -128}
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        _ = len(v3)
                                                    default:
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            v100 := []int{fi86, 42, v49}
                                                            _ = v100
                                                        case "a":
                                                            for fi101 := 0; fi101 < 16; fi101++ {
                                                                switch (("a" + "hello") + ("a" + "a")) {
                                                                case "a":
                                                                    v102 := GenS9{F0: ((32767 == v12) && (255 > genFn1("\n\t", 274, -2147483648))), F1: v12}
                                                                    _ = v102
                                                                case "\n\t":
                                                                    v48 = true
                                                                default:
                                                                    fmt.Println(fi101)
                                                                }
                                                                var v103 uint = 65535
                                                                _ = fi101
                                                                _ = v103
                                                            }
                                                        case "hello":
                                                            if v48 {
                                                                var v104 int64 = (v13 * (^(v2)))
                                                                var v105 bool = v48
                                                                v106 := GenS15{F0: 127}
                                                                _ = v104
                                                                _ = v105
                                                                _ = v106
                                                            }
                                                        default:
                                                            v48 = v48
                                                        }
                                                    }
                                                    _ = v99
                                                } else {
                                                    var v107 uint8 = 2
                                                    _ = v107
                                                }
                                                fmt.Println(v11, v12)
                                            }
                                            var v108 bool = (!(0 <= 1000))
                                            _ = v108
                                        }
                                        if (1e10 >= 1e10) {
                                            delete(v3, ("" + "foo bar"))
                                        }
                                    } else {
                                        for fi109 := 0; fi109 < 19; fi109++ {
                                            v13 = -92
                                            _ = fi109
                                        }
                                        v110 := GenS16{F0: 912, F1: (v49 / ((int(v50)) | 1)), F2: 519}
                                        var v111 int32 = 42
                                        _ = v110
                                        _ = v111
                                    }
                                }
                                _ = fi86
                            }
                        }
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
