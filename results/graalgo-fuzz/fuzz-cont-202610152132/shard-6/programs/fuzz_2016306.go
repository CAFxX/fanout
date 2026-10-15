package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 int32; F2 string }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 int, p1 int) int16 {
    fmt.Println(p1, p1)
    v5 := []int{}
    v6 := GenS1{F0: -651, F1: 861, F2: p1}
    if true {
        if true {
            p1 = (p1 - (^(p1)))
        } else {
            var v7 int32 = -1
            _ = v7
        }
    } else {
        fmt.Println(p0, p0)
        var v8 string = ""
        for fi9 := 0; fi9 < 2; fi9++ {
            v5 = append(v5, p1)
            v10 := GenS2{F0: 794}
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int16) uint {
    p0 = p0
    v11 := []int{}
    var v12 uint32 = 266
    _ = p0
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 uint8 = 1
    var v2 uint8 = v1
    v3 := make(map[int]uint32)
    var v4 float32 = 1.5
    v13 := []int{}
    if (!true) {
        var v14 uint16 = 43
        if (!(-2147483648 != -764)) {
            switch "hello" {
            case "hello":
                for fi15 := 0; fi15 < 6; fi15++ {
                    var v16 uint = 42
                    switch "\n\t" {
                    case "hello":
                        switch 2147483647 {
                        case 32768:
                            v4 -= v4
                        case 234:
                            var v17 uint32 = 914
                            _ = v17
                        }
                    case "":
                        fmt.Println(v1, v4, v2)
                    case "foo bar":
                        v16 = 4294967295
                    }
                    switch "a" {
                    case "":
                        v2 = (^(v2))
                    }
                    _ = fi15
                    _ = v16
                }
            case "":
                fmt.Println(v1, v1, v14)
            case "a":
                for fi18 := 0; fi18 < 13; fi18++ {
                    if false {
                        v19 := GenS1{F0: -945, F1: (-(fi18)), F2: (fi18 % ((fi18) | 1))}
                        var v20 int64 = 42
                        _ = v19
                        _ = v20
                    } else {
                        v4 = v4
                        v14 += 0
                        switch (fi18 / (((fi18 / ((42) | 1))) | 1)) {
                        case -969:
                            v21 := GenS2{F0: 2}
                            _ = v21
                        case 42:
                            var v22 int = (fi18 & -96)
                            _ = v22
                        case 4294967295:
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v4 -= 1.5
                                _ = fi23
                            }
                        default:
                            var v24 float32 = v4
                            _ = v24
                        }
                    }
                    var v25 int32 = -2147483648
                    v2 = 42
                    _ = fi18
                    _ = v25
                }
            }
            var v26 int32 = -428
            _ = v26
        }
        _ = v14
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 uint = 65535
        switch "\n\t" {
        case "\n\t":
            for fi29 := 0; fi29 < 19; fi29++ {
                if true {
                    delete(v3, fi29)
                    switch (^(fi27)) {
                    case 32768:
                        var v30 float32 = 0.1
                        v28 = v28
                        _ = v30
                    case 0:
                        var v31 int64 = 946
                        _ = v31
                    }
                    switch "foo bar" {
                    case "hello":
                        switch ("" + ("" + "foo bar")) {
                        case "":
                            delete(v3, (fi27 | (fi27 * fi29)))
                        case "\n\t":
                            switch fi29 {
                            case -584:
                                var v32 int8 = 82
                                _ = v32
                            case 32768:
                                v33 := GenS1{F0: -289, F1: fi27, F2: (fi27 * 4294967295)}
                                _ = v33
                            case 0:
                                var v34 int16 = -25
                                _ = v34
                            default:
                                _ = len(v13)
                            }
                        case "a":
                            var v35 uint32 = 596
                            _ = v35
                        default:
                            v28 += genFn1(-632)
                        }
                    }
                    var v36 float32 = ((float32((((1e10 - 3.14159) - float64(v4)) / ((0.0 + 1.5) - 1e10))) + -2.25) - -2.25)
                    _ = v36
                } else {
                    if len(v13) > 0 {
                        v13[0] = fi29
                    }
                }
                switch "hello" {
                case "a":
                    var v37 int8 = -20
                    switch (fi29 & fi27) {
                    case 42:
                        v38 := v3[42]
                        _ = v38
                    default:
                        fmt.Println(v1)
                    }
                    _ = v37
                default:
                    if ((!true) && true) {
                        var v39 bool = false
                        switch "hello" {
                        case "foo bar":
                            v3[(fi29 - (fi29 >> (((v28 / ((genFn1(0)) | 1))) & 63)))] = 1
                        case "\n\t":
                            if v39 {
                                v1 -= 82
                            } else {
                                v39 = false
                                var v40 int16 = genFn0(fi27, int(v2))
                                if len(v13) > 0 {
                                    v13[0] = fi29
                                }
                                _ = v40
                            }
                        case "":
                            v41 := GenS1{F0: -2147483648, F1: (fi29 % (((fi27 << ((uint(v2)) & 63))) | 1)), F2: fi27}
                            _ = v41
                        default:
                            v13 = append(v13, (-(fi27)))
                            switch fi29 {
                            case -2147483648:
                                switch fi27 {
                                case -672:
                                    var v42 bool = (!((v39 && v39) && (!(v28 >= v28))))
                                    _ = v42
                                default:
                                    delete(v3, (fi27 * 32768))
                                }
                            }
                        }
                        _ = v39
                    }
                }
                v2 += v2
                var v43 string = ("foo bar" + ((("\n\t" + ("a" + ("foo bar" + "a"))) + "hello") + "hello"))
                v43 = (v43 + v43)
                v44 := GenS3{F0: -1}
                v1 = 127
                if false {
                    _ = len(v13)
                }
                v1 += (v1 * (v2 - (v1 << (((^(v28))) & 63))))
                var v45 int16 = -376
                var v46 uint64 = 2
                _ = fi29
                _ = v43
                _ = v44
                _ = v45
                _ = v46
            }
        case "a":
            v47 := GenS4{F0: "\n\t", F1: 1, F2: "a"}
            _ = v47
        case "hello":
            var v48 int32 = 0
            _ = v48
        }
        _ = fi27
        _ = v28
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        v50 := GenS3{F0: -30}
        _ = fi49
        _ = v50
    }
    switch "foo bar" {
    case "":
        fmt.Println(v2)
    case "foo bar":
        var v51 uint16 = 2
        _ = v51
    case "hello":
        var v52 int8 = -1
        _ = v52
    }
    v53 := GenS2{F0: uint32(v1)}
    v2 = (v2 % (((v2 ^ (v1 << ((239) & 63)))) | 1))
    if (32767 == 65535) {
        v54 := v3[-1]
        _ = v54
    } else {
        for fi55 := 0; fi55 < 14; fi55++ {
            v1 = 255
            v2 = 255
            _ = fi55
        }
        v3[2147483647] = uint32(v2)
        for fi56 := 0; fi56 < 16; fi56++ {
            switch "a" {
            case "a":
                var v57 int8 = 127
                _ = v57
            case "":
                var v58 float32 = float32((3.14159 / 3.14159))
                _ = v58
            case "hello":
                v59 := GenS2{F0: 2147483648}
                _ = v59
            }
            var v60 uint8 = 127
            _ = fi56
            _ = v60
        }
    }
    _ = len(v3)
    var v61 int64 = 0
    var v62 int64 = (v61 | v61)
    var v63 uint32 = 1
    fmt.Println(v1, v62)
    _ = len(v13)
    for fi64 := 0; fi64 < 4; fi64++ {
        fmt.Println(v61, v1)
        if true {
            if false {
                for fi65 := 0; fi65 < 15; fi65++ {
                    v4 = (3.14159 * float32(1.5))
                    var v66 int = (fi64 << ((42) & 63))
                    if true {
                        switch fi64 {
                        case 128:
                            if ((v61 / ((-617) | 1)) == (-(v62))) {
                                var v67 int16 = 89
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v1, v67, v61)
                                case "":
                                    v68 := GenS1{F0: fi65, F1: (fi64 + -1), F2: fi64}
                                    _ = v68
                                case "foo bar":
                                    if false {
                                        fmt.Println(fi64, v66)
                                    } else {
                                        switch v66 {
                                        case 0:
                                            switch fi65 {
                                            case 2147483647:
                                                fmt.Println(v61)
                                            case 128:
                                                fmt.Println(v61)
                                            case -9223372036854775808:
                                                var v69 uint16 = 0
                                                _ = v69
                                            default:
                                                v67 = (-(v67))
                                            }
                                        default:
                                            var v70 string = "foo bar"
                                            _ = v70
                                        }
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                v71 := GenS2{F0: 955}
                                                var v72 int8 = -42
                                                var v73 int8 = -128
                                                _ = v71
                                                _ = v72
                                                _ = v73
                                            } else {
                                                v4 -= (v4 + (0.1 * v4))
                                                var v74 string = "hello"
                                                var v75 uint = 1
                                                _ = v74
                                                _ = v75
                                            }
                                        case "a":
                                            fmt.Println(v67, v63, v1)
                                        case "foo bar":
                                            switch (fi64 >> ((328) & 63)) {
                                            case 1:
                                                _ = len(v3)
                                            default:
                                                if ((true || true) && false) {
                                                    fmt.Println(v2, fi65, v1)
                                                } else {
                                                    fmt.Println(v67, fi65, v61)
                                                    var v76 int8 = 102
                                                    v77 := GenS3{F0: (v61 + (v62 ^ v62))}
                                                    fmt.Println(v1, v63)
                                                    switch v66 {
                                                    case 286:
                                                        var v78 int32 = -1
                                                        _ = v78
                                                    case 0:
                                                        fmt.Println(v61, v2)
                                                        if ((1e10 * (((-2.25 - float64(v4)) - 1.5) - 3.14159)) <= (1.5 * (0.1 - ((((0.0 / 0.1) - float64((0.1 - v4))) * -2.25) + 3.14159)))) {
                                                            var v79 uint8 = v2
                                                            fmt.Println(v63, v66, v66)
                                                            var v80 uint64 = 42
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        if (false && true) {
                                                            var v81 uint8 = (v1 & (v1 | (v2 << ((4294967295) & 63))))
                                                            var v82 string = "a"
                                                            var v83 string = v82
                                                            v84 := GenS3{F0: (-(v61))}
                                                            _ = v81
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        v85 := GenS2{F0: v63}
                                                        var v86 int = v66
                                                        var v87 int16 = 600
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                    _ = v76
                                                    _ = v77
                                                }
                                            }
                                        }
                                        if false {
                                            v3[(fi65 | (fi64 * (^(v66))))] = 409
                                        } else {
                                            var v88 float64 = (0.0 + 0.0)
                                            v89 := GenS3{F0: 306}
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                }
                                _ = v67
                            }
                        }
                        v1 = v2
                        var v90 bool = true
                        v2 += (^(v2))
                        switch (v66 << ((89) & 63)) {
                        case 582:
                            var v91 uint = 915
                            _ = v91
                        case -2147483648:
                            v92 := GenS1{F0: 32767, F1: (v66 - (v66 * -2147483648)), F2: fi64}
                            _ = v92
                        case 201:
                            v93 := GenS5{F0: 0, F1: "\n\t"}
                            _ = v93
                        default:
                            v94 := GenS3{F0: (v61 ^ (v62 % ((v61) | 1)))}
                            _ = v94
                        }
                        if v90 {
                            var v95 int = (fi65 ^ (fi65 / ((9223372036854775807) | 1)))
                            var v96 uint64 = 4294967295
                            _ = v95
                            _ = v96
                        } else {
                            v61 += (v61 | v62)
                            v97 := GenS1{F0: fi65, F1: -9223372036854775808, F2: fi64}
                            _ = v97
                        }
                        v2 += v2
                        _ = v90
                    } else {
                        if (0 <= 762) {
                            v1 = v2
                            _ = len(v3)
                            var v98 string = (("a" + "hello") + "\n\t")
                            switch 4294967295 {
                            case 128:
                                v98 = (v98 + (("\n\t" + "hello") + ""))
                            case -672:
                                var v99 uint64 = 42
                                _ = v99
                            case 1:
                                v63 -= 1
                            default:
                                fmt.Println(fi65, v61, v63)
                            }
                            var v100 int16 = 32767
                            switch v66 {
                            case 32768:
                                fmt.Println(v100, fi65, fi65)
                            case -718:
                                if (-2.25 >= 1e10) {
                                    var v101 uint32 = (^(v63))
                                    if true {
                                        _ = len(v3)
                                        if (v98 < v98) {
                                            var v102 bool = false
                                            _ = v102
                                        } else {
                                            if true {
                                                _ = len(v13)
                                                var v103 int16 = (v100 * 135)
                                                var v104 int16 = (v100 | (v103 * v100))
                                                _ = v103
                                                _ = v104
                                            }
                                        }
                                    } else {
                                        fmt.Println(v98, v4)
                                        var v105 uint = 223
                                        _ = v105
                                    }
                                    var v106 uint = 0
                                    v4 = (v4 + ((v4 * float32((((float64(v4) / 1.0) - -2.25) - 1.5))) + v4))
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v62 -= v62
                                    if false {
                                        fmt.Println(v2, fi65, v1)
                                    } else {
                                        var v107 uint64 = 42
                                        var v108 float32 = v4
                                        var v109 uint64 = v107
                                        v101 -= (v63 + (v101 & 2147483647))
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    }
                                    if (!true) {
                                        v110 := GenS4{F0: "foo bar", F1: int32(v61), F2: v98}
                                        v66 += (fi64 >> ((v106) & 63))
                                        if true {
                                            delete(v3, fi65)
                                            fmt.Println(v66)
                                            fmt.Println(v98, v63, v106)
                                            v61 = v61
                                        } else {
                                            var v111 int16 = (v100 - (v100 + (v100 | 42)))
                                            var v112 bool = false
                                            _ = v111
                                            _ = v112
                                        }
                                        _ = v110
                                    }
                                    v113 := GenS1{F0: (fi65 - (v66 % ((fi65) | 1))), F1: 42, F2: fi65}
                                    v101 = (v101 & uint32(v1))
                                    _ = v101
                                    _ = v106
                                    _ = v113
                                }
                            }
                            _ = v100
                            _ = v98
                        }
                        fmt.Println(v66)
                    }
                    _ = fi65
                    _ = v66
                }
                if len(v13) > 0 {
                    v13[0] = fi64
                }
                v114 := v3[32767]
                _ = v114
            }
        } else {
            fmt.Println(v1)
        }
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v61
    _ = v62
    _ = v63
}
