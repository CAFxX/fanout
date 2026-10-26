package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS7 struct { F0 uint; F1 int8; F2 int64 }
type GenS8 struct { F0 uint16; F1 int64; F2 uint }

func main() {
    var v1 int = 128
    var v2 uint = 255
    var v3 string = ("hello" + "foo bar")
    if true {
        switch ("" + v3) {
        case "hello":
            v4 := GenS1{F0: v2, F1: uint16(v2)}
            _ = v4
        }
    } else {
        v3 = v3
    }
    v2 -= (^(v2))
    var v5 bool = false
    if (!(3.14159 <= float64(1e10))) {
        if v5 {
            v3 = v3
            var v6 uint64 = 2
            switch (v1 << (((v2 & 2)) & 63)) {
            case 2147483647:
                v5 = (v5 && v5)
            }
            _ = v6
        } else {
            v7 := GenS2{F0: -1, F1: 32768}
            v8 := GenS2{F0: -2147483648, F1: (v1 - (v1 | v1))}
            _ = v7
            _ = v8
        }
        var v9 uint64 = 669
        v1 -= -355
        _ = v9
    } else {
        var v10 uint16 = 255
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                switch v3 {
                case "foo bar":
                    switch fi11 {
                    case 2147483647:
                        v13 := []int{}
                        _ = v13
                    case -151:
                        v14 := GenS1{F0: v2, F1: 0}
                        _ = v14
                    case -130:
                        v10 += v10
                        v15 := []int{(fi12 | fi11)}
                        var v16 string = v3
                        v2 -= 65535
                        _ = v15
                        _ = v16
                    default:
                        v17 := []int{}
                        _ = v17
                    }
                default:
                    v18 := GenS2{F0: 329, F1: (fi12 ^ fi11)}
                    _ = v18
                }
                if (v2 < v2) {
                    var v19 int64 = -1
                    switch 4294967295 {
                    case 1:
                        v20 := []int{}
                        _ = v20
                    default:
                        v5 = v5
                    }
                    v2 += 484
                    _ = v19
                } else {
                    v21 := GenS1{F0: v2, F1: (v10 >> ((v2) & 63))}
                    _ = v21
                }
                v22 := []int{fi11}
                var v23 float64 = (1.5 * ((3.14159 + 0.0) * 1e10))
                fmt.Println(v3)
                _ = fi12
                _ = v22
                _ = v23
            }
            var v24 uint64 = 42
            _ = fi11
            _ = v24
        }
        switch (v3 + v3) {
        case "\n\t":
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint16 = v10
                if v5 {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v28 := []int{(fi27 / (((v1 & fi27)) | 1)), (v1 ^ (fi27 - v1))}
                        _ = fi27
                        _ = v28
                    }
                    v29 := []int{(v1 >> ((v2) & 63)), int(v10), fi25, v1}
                    _ = v29
                } else {
                    var v30 bool = ((!(2147483647 == 2147483647)) || v5)
                    var v31 uint16 = (v10 | v26)
                    _ = v30
                    _ = v31
                }
                switch (fi25 ^ -388) {
                case 128:
                    v32 := GenS3{F0: 0}
                    _ = v32
                case 127:
                    var v33 int8 = 86
                    _ = v33
                case 0:
                    v26 -= v26
                }
                _ = fi25
                _ = v26
            }
        case "":
            if v5 {
                v34 := GenS1{F0: v2, F1: 2}
                if (-1 == 90) {
                    v10 -= 2
                }
                if (0 != 100) {
                    if (((v3 + v3) <= "hello") || v5) {
                        switch ("foo bar" + ((v3 + ((v3 + (("foo bar" + "hello") + ("" + v3))) + ("" + (("hello" + "hello") + "a")))) + v3)) {
                        case "a":
                            if (v5 || (!v5)) {
                                v35 := []int{}
                                v2 = (v2 | (v2 | (v2 + (v2 >> (((^(v2))) & 63)))))
                                if (!v5) {
                                    v35 = append(v35, v1)
                                    fmt.Println(v3, v2, v5)
                                    switch v3 {
                                    case "foo bar":
                                        var v36 float32 = (3.14159 - 0.0)
                                        _ = v36
                                    case "":
                                        v3 = ("hello" + "hello")
                                    }
                                } else {
                                    var v37 bool = (v5 || v5)
                                    v38 := GenS1{F0: (v2 % ((443) | 1)), F1: (v10 * 2)}
                                    v39 := GenS4{F0: 42, F1: 65535, F2: 42}
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                _ = v35
                            }
                        default:
                            v1 -= v1
                        }
                        v1 -= (^(v1))
                    } else {
                        v40 := []int{(-(v1)), v1}
                        _ = v40
                    }
                    v10 += v10
                    v41 := []int{}
                    _ = v41
                }
                _ = v34
            } else {
                switch -217 {
                case -1:
                    fmt.Println(v3)
                case -421:
                    v3 = v3
                }
                var v42 uint32 = 2147483647
                _ = v42
            }
        case "foo bar":
            v43 := []int{(v1 & (v1 ^ v1)), v1, v1}
            _ = v43
        default:
            fmt.Println(v3)
        }
        _ = v10
    }
    switch v1 {
    case -2147483648:
        v44 := GenS5{F0: -885}
        _ = v44
    case 4294967295:
        var v45 int64 = 42
        _ = v45
    }
    v2 = (v2 % ((v2) | 1))
    var v46 uint64 = 886
    var v47 uint = 511
    if (42 > 100) {
        fmt.Println(v5)
        for fi48 := 0; fi48 < 6; fi48++ {
            var v49 int8 = -122
            v2 = (v47 ^ v47)
            v50 := []int{int(v49), v1, 32767}
            _ = fi48
            _ = v49
            _ = v50
        }
        var v51 uint32 = 42
        _ = v51
    }
    var v52 uint16 = 255
    v3 = v3
    if (v1 <= (^(v1))) {
        v53 := []int{(^(v1))}
        _ = v53
    } else {
        v54 := []int{}
        switch (v1 << ((65535) & 63)) {
        case 617:
            var v55 uint16 = v52
            _ = v55
        case 606:
            v56 := GenS1{F0: v2, F1: 255}
            _ = v56
        case 535:
            if len(v54) >= 2 {
                v54 = v54[1:len(v54)]
            }
        default:
            switch v3 {
            case "foo bar":
                fmt.Println(v3, v3)
            case "hello":
                _ = len(v54)
            }
        }
        if len(v54) >= 2 {
            v54 = v54[1:len(v54)]
        }
        _ = v54
    }
    var v57 int16 = -1
    var v58 int64 = -9223372036854775808
    v59 := []int{(v1 % ((v1) | 1))}
    if (v5 && (!v5)) {
        for fi60 := 0; fi60 < 11; fi60++ {
            for fi61 := 0; fi61 < 15; fi61++ {
                var v62 int8 = -1
                switch v3 {
                case "":
                    var v63 uint = 307
                    _ = v63
                default:
                    var v64 float64 = 0.1
                    _ = v64
                }
                _ = fi61
                _ = v62
            }
            _ = fi60
        }
        switch (v3 + "") {
        case "foo bar":
            if (!((v58 + 9223372036854775807) == 9223372036854775807)) {
                v52 -= (v52 - v52)
            }
        }
    } else {
        if v5 {
            if len(v59) > 0 {
                v59[0] = (v1 & (v1 << ((v2) & 63)))
            }
            for fi65 := 0; fi65 < 10; fi65++ {
                if (v5 && (!((float64(float32((float64((0.1 + 0.1)) + 1e10))) != 3.14159) || (42 < -2147483648)))) {
                    for fi66 := 0; fi66 < 10; fi66++ {
                        if (false || (!(v5 || v5))) {
                            _ = len(v59)
                            var v67 uint = v47
                            v5 = true
                            if v5 {
                                if len(v59) >= 2 {
                                    v59 = v59[1:len(v59)]
                                }
                                if v5 {
                                    if (!v5) {
                                        v3 = v3
                                        v67 += v67
                                        switch ((v3 + "foo bar") + v3) {
                                        case "":
                                            _ = len(v59)
                                        case "\n\t":
                                            switch ("a" + "\n\t") {
                                            case "foo bar":
                                                v68 := GenS3{F0: -186}
                                                _ = v68
                                            case "hello":
                                                var v69 int = v1
                                                _ = v69
                                            case "a":
                                                v3 = (v3 + (("foo bar" + v3) + "foo bar"))
                                            }
                                        default:
                                            switch 1 {
                                            case 2147483647:
                                                if (42 > -1) {
                                                    var v70 int64 = -1
                                                    v58 += (-(v58))
                                                    _ = v70
                                                }
                                            default:
                                                fmt.Println(fi65, v3)
                                            }
                                        }
                                    }
                                    var v71 uint16 = v52
                                    _ = v71
                                }
                            } else {
                                if ((0.0 - -2.25) < -2.25) {
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                } else {
                                    if v5 {
                                        v1 += 36
                                    }
                                    v3 = v3
                                }
                                var v72 string = v3
                                _ = v72
                            }
                            _ = v67
                        } else {
                            switch (fi65 * 0) {
                            case -951:
                                if ((!v5) && (!(v5 && false))) {
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v47, fi65)
                                    }
                                }
                            case -149:
                                fmt.Println(v58)
                            default:
                                v5 = v5
                            }
                            if true {
                                v57 -= (v57 >> ((255) & 63))
                                if v5 {
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                    var v73 uint16 = 425
                                    v47 = v2
                                    var v74 int8 = -26
                                    v75 := GenS6{F0: v58, F1: v46, F2: 37}
                                    v74 -= (v74 & v74)
                                    _ = v73
                                    _ = v74
                                    _ = v75
                                }
                            } else {
                                if true {
                                    _ = len(v59)
                                    var v76 string = v3
                                    switch v76 {
                                    case "hello":
                                        if len(v59) >= 2 {
                                            v59 = v59[1:len(v59)]
                                        }
                                    case "":
                                        var v77 bool = v5
                                        _ = v77
                                    }
                                    var v78 int = fi65
                                    var v79 int32 = 77
                                    if len(v59) >= 2 {
                                        v59 = v59[1:len(v59)]
                                    }
                                    if (v52 <= v52) {
                                        v3 = "\n\t"
                                        v80 := GenS4{F0: 42, F1: 550, F2: int32(v2)}
                                        _ = len(v59)
                                        _ = v80
                                    } else {
                                        switch v1 {
                                        case 128:
                                            _ = len(v59)
                                        case 647:
                                            var v81 float32 = 3.14159
                                            _ = v81
                                        case 4294967295:
                                            if len(v59) >= 2 {
                                                v59 = v59[1:len(v59)]
                                            }
                                        }
                                        switch 9223372036854775807 {
                                        case 0:
                                            if (v78 != v1) {
                                                v78 = v78
                                                v59 = append(v59, fi65)
                                            } else {
                                                if v5 {
                                                    if (false && ((float32(1e10) + (1.5 + 3.14159)) != 3.14159)) {
                                                        v46 -= v46
                                                        v46 = 42
                                                    }
                                                    v82 := GenS6{F0: v58, F1: (v46 | 0), F2: 191}
                                                    v46 = (v46 * 798)
                                                    var v83 bool = (!v5)
                                                    v2 += v2
                                                    var v84 bool = ((((v76 + v76) <= "\n\t") || (v46 != 423)) || (v5 || (v46 < 4294967295)))
                                                    if v5 {
                                                        fmt.Println(v3)
                                                    } else {
                                                        fmt.Println(v1)
                                                    }
                                                    if len(v59) >= 2 {
                                                        v59 = v59[1:len(v59)]
                                                    }
                                                    fmt.Println(fi65, v84, v47)
                                                    var v85 string = v3
                                                    var v86 float32 = 0.1
                                                    if false {
                                                        v3 = "hello"
                                                        var v87 uint32 = 1
                                                        _ = v87
                                                    }
                                                    v84 = ((v46 % (((v46 / ((4294967295) | 1))) | 1)) > (v46 ^ v46))
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    v76 = v76
                                                    var v88 int16 = (^(v57))
                                                    v89 := GenS7{F0: 1, F1: 100, F2: v58}
                                                    _ = v88
                                                    _ = v89
                                                }
                                                switch v3 {
                                                case "foo bar":
                                                    var v90 int32 = v79
                                                    _ = v90
                                                case "hello":
                                                    var v91 int16 = (v57 & v57)
                                                    _ = v91
                                                }
                                                var v92 int8 = 0
                                                _ = v92
                                            }
                                        case -725:
                                            var v93 int = int(v47)
                                            _ = v93
                                        }
                                        v94 := GenS2{F0: (^(v79)), F1: v1}
                                        _ = v94
                                    }
                                    _ = v76
                                    _ = v78
                                    _ = v79
                                }
                            }
                            if false {
                                var v95 bool = false
                                var v96 float32 = 1e10
                                if true {
                                    fmt.Println(v57)
                                    v96 -= (float32(1.5) / v96)
                                    v97 := GenS8{F0: v52, F1: (v58 | 718), F2: 395}
                                    _ = v97
                                } else {
                                    var v98 uint64 = (v46 - uint64(v2))
                                    v99 := GenS6{F0: v58, F1: 4294967295, F2: 195}
                                    _ = v98
                                    _ = v99
                                }
                                _ = v95
                                _ = v96
                            }
                        }
                        switch fi66 {
                        case 2147483647:
                            v57 = (v57 % ((int16(fi66)) | 1))
                        }
                        var v100 uint32 = 13
                        _ = fi66
                        _ = v100
                    }
                } else {
                    v59 = append(v59, -9223372036854775808)
                    v101 := GenS2{F0: -2147483648, F1: fi65}
                    _ = v101
                }
                _ = fi65
            }
        } else {
            for fi102 := 0; fi102 < 15; fi102++ {
                for fi103 := 0; fi103 < 6; fi103++ {
                    v2 = 42
                    v46 = v46
                    _ = fi103
                }
                _ = len(v59)
                fmt.Println(fi102, fi102, v1)
                _ = fi102
            }
            for fi104 := 0; fi104 < 8; fi104++ {
                fmt.Println(v1, v58)
                _ = fi104
            }
        }
        v1 = (v1 + (v1 - v1))
    }
    v105 := GenS7{F0: v47, F1: 114, F2: v58}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v57
    _ = v58
    _ = v59
}
