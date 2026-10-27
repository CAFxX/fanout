package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint16; F2 string }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int8; F1 int32; F2 int8 }
type GenS8 struct { F0 uint16; F1 string; F2 int8 }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint32; F1 int16 }
type GenS12 struct { F0 string; F1 bool; F2 int32 }
type GenS13 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -82
    var v3 int16 = -1
    switch int(v2) {
    case -807:
        for fi4 := 0; fi4 < 16; fi4++ {
            v3 = -152
            if true {
                v5 := GenS1{F0: (1 > 2), F1: 32767, F2: 4294967295}
                fmt.Println(fi4)
                _ = v5
            }
            v6 := []int{}
            _ = fi4
            _ = v6
        }
    case 32767:
        v2 -= v2
    case -2147483648:
        for fi7 := 0; fi7 < 1; fi7++ {
            if (42 < v2) {
                v8 := GenS1{F0: true, F1: 32767, F2: 709}
                v9 := GenS2{F0: v3}
                _ = v8
                _ = v9
            } else {
                v2 += v2
                for fi10 := 0; fi10 < 12; fi10++ {
                    v11 := []int{(fi7 | (^(fi10)))}
                    if false {
                        v12 := v1[-256]
                        var v13 int32 = 376
                        _ = v12
                        _ = v13
                    } else {
                        fmt.Println(v2, v2)
                        if false {
                            var v14 int64 = 9223372036854775807
                            var v15 float64 = (1.5 - 3.14159)
                            _ = v14
                            _ = v15
                        } else {
                            if (((float64(float32(float64((1.5 / ((0.1 / (((1e10 + (float32(1e10) - 0.1)) / 1.0) - (0.0 - 3.14159))) * 3.14159))))) / 1e10) != 3.14159) || (!false)) {
                                v16 := GenS3{F0: 767, F1: 127, F2: 18446744073709551615}
                                var v17 uint64 = 4294967295
                                v18 := GenS1{F0: false, F1: 0, F2: 2}
                                var v19 float32 = 1.5
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        }
                    }
                    var v20 int = (fi10 / ((127) | 1))
                    _ = fi10
                    _ = v11
                    _ = v20
                }
            }
            var v21 uint8 = 255
            _ = fi7
            _ = v21
        }
    }
    v22 := GenS1{F0: (!false), F1: 124, F2: 408}
    v2 += (v2 * (v2 | v2))
    if false {
        for fi23 := 0; fi23 < 7; fi23++ {
            switch fi23 {
            case 4294967295:
                var v24 int16 = -449
                _ = v24
            case 2147483647:
                var v25 int32 = 42
                _ = v25
            case 42:
                v26 := []int{}
                _ = v26
            }
            v27 := []int{fi23, fi23, fi23, fi23}
            _ = fi23
            _ = v27
        }
        fmt.Println(v2)
    } else {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := v1[fi28]
            var v30 uint8 = 42
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := GenS4{F0: -1, F1: 1, F2: "a"}
        for fi32 := 0; fi32 < 3; fi32++ {
            var v33 uint64 = 42
            fmt.Println(fi32, v3)
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v2 = v2
    v34 := []int{794, 691, 838, -7}
    v2 = (v2 >> ((443) & 63))
    v35 := v1[128]
    switch "hello" {
    case "foo bar":
        switch 128 {
        case 4294967295:
            _ = len(v1)
        case 0:
            v3 = v3
        case 32767:
            v35 = (v35 % ((42) | 1))
        }
    case "":
        var v36 int8 = 1
        _ = v36
    case "\n\t":
        v37 := GenS5{F0: 572, F1: -9223372036854775808}
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        v35 = (v35 ^ 673)
        _ = fi38
    }
    switch -156 {
    case 32768:
        switch -1 {
        case -295:
            switch 42 {
            case 32768:
                for fi39 := 0; fi39 < 14; fi39++ {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        var v41 int64 = 0
                        v1[fi40] = 1
                        var v42 int16 = (v3 + int16(v35))
                        _ = len(v34)
                        switch fi39 {
                        case 32767:
                            var v43 int32 = -398
                            _ = v43
                        case -315:
                            if (209 >= 224) {
                                var v44 uint = 192
                                _ = v44
                            } else {
                                v35 = 901
                                _ = len(v34)
                                v41 -= (v41 << ((42) & 63))
                            }
                            v41 = (v41 & v41)
                            if false {
                                fmt.Println(v2, fi40, v41)
                                v45 := GenS6{F0: fi40}
                                if (0.0 > float64((0.1 + 3.14159))) {
                                    delete(v1, (fi40 | (fi39 >> ((131) & 63))))
                                    if (true || (true || (v3 > (v3 + -22)))) {
                                        if true {
                                            var v46 bool = true
                                            _ = v46
                                        } else {
                                            v42 = (v3 * int16(v41))
                                            if len(v34) > 0 {
                                                v34[0] = (^(fi39))
                                            }
                                            v41 -= v41
                                            v35 = 4294967295
                                        }
                                    } else {
                                        if (354 <= 2) {
                                            v3 += 694
                                        } else {
                                            if true {
                                                if len(v34) > 0 {
                                                    v34[0] = fi39
                                                }
                                                v47 := GenS7{F0: 100, F1: 2147483647, F2: 112}
                                                _ = v47
                                            } else {
                                                v48 := GenS8{F0: 65535, F1: "foo bar", F2: (v2 * int8(v41))}
                                                _ = v48
                                            }
                                            var v49 int8 = v2
                                            var v50 string = "\n\t"
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                }
                                _ = v45
                            } else {
                                v2 = int8(v41)
                                if true {
                                    _ = len(v34)
                                    switch int(v41) {
                                    case 2147483647:
                                        v42 += (v42 / (((v42 >> ((2) & 63))) | 1))
                                    case 32768:
                                        v42 = 420
                                    case -9223372036854775808:
                                        v34 = append(v34, fi39)
                                    }
                                    v2 = 91
                                } else {
                                    switch ((("hello" + "") + "a") + "\n\t") {
                                    case "hello":
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v51 uint64 = 0
                                            var v52 uint16 = 32767
                                            switch 128 {
                                            case 42:
                                                switch ((("\n\t" + "hello") + "a") + "foo bar") {
                                                case "\n\t":
                                                    if len(v34) > 0 {
                                                        v34[0] = fi39
                                                    }
                                                default:
                                                    v3 = v3
                                                }
                                            case -992:
                                                if (0 != 0) {
                                                    v53 := GenS7{F0: v2, F1: 0, F2: v2}
                                                    v51 += (v51 & v51)
                                                    var v54 int64 = (v41 | (-(v41)))
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    var v55 uint8 = 255
                                                    switch fi39 {
                                                    case 4294967295:
                                                        fmt.Println(v51)
                                                    }
                                                    _ = v55
                                                }
                                            default:
                                                if false {
                                                    switch int(v35) {
                                                    case 42:
                                                        switch (("\n\t" + (("a" + "") + "a")) + "a") {
                                                        case "a":
                                                            if false {
                                                                var v56 int64 = (-(v41))
                                                                fmt.Println(v52, v51, v52)
                                                                _ = v56
                                                            }
                                                        case "foo bar":
                                                            var v57 string = ("foo bar" + "")
                                                            _ = v57
                                                        case "":
                                                            v58 := GenS7{F0: 77, F1: -595, F2: (v2 - (v2 | (^(v2))))}
                                                            _ = v58
                                                        default:
                                                            var v59 uint8 = 2
                                                            _ = v59
                                                        }
                                                    case 32767:
                                                        var v60 int16 = (v3 * (v3 - (v42 >> ((0) & 63))))
                                                        _ = v60
                                                    case 333:
                                                        fmt.Println(v2, v41)
                                                    default:
                                                        v42 = (v3 | (v3 >> ((0) & 63)))
                                                    }
                                                }
                                            }
                                            _ = v51
                                            _ = v52
                                        case "\n\t":
                                            if true {
                                                var v61 int = (fi39 >> ((2) & 63))
                                                if (true || true) {
                                                    fmt.Println(v2)
                                                } else {
                                                    if len(v34) >= 2 {
                                                        v34 = v34[1:len(v34)]
                                                    }
                                                }
                                                _ = v61
                                            } else {
                                                var v62 int64 = (v41 % ((9223372036854775807) | 1))
                                                _ = v62
                                            }
                                        case "hello":
                                            v63 := GenS9{F0: 317, F1: false}
                                            _ = v63
                                        default:
                                            if ((v41 % (((v41 | 334)) | 1)) >= (v41 >> ((18446744073709551615) & 63))) {
                                                v2 = int8(fi39)
                                                var v64 uint64 = 4294967295
                                                _ = v64
                                            }
                                            var v65 int = (fi39 << ((950) & 63))
                                            _ = v65
                                        }
                                    case "foo bar":
                                        var v66 string = "a"
                                        _ = v66
                                    case "\n\t":
                                        switch -817 {
                                        case 0:
                                            var v67 int64 = (v41 + v41)
                                            _ = v67
                                        case 2147483647:
                                            switch 4294967295 {
                                            case 9223372036854775807:
                                                if ((-(v42)) >= (v42 | (-(v42)))) {
                                                    v42 -= (-(v42))
                                                }
                                            case 53:
                                                var v68 uint64 = 0
                                                switch (fi39 % ((fi40) | 1)) {
                                                case 32768:
                                                    if (((841 != 32767) && false) && false) {
                                                        v69 := GenS8{F0: 145, F1: ((("foo bar" + "") + (((("foo bar" + ("a" + ("" + "\n\t"))) + "hello") + "") + "\n\t")) + ("hello" + ("foo bar" + "\n\t"))), F2: 100}
                                                        v70 := v1[(fi39 - (fi39 * (fi40 * (fi40 | 32768))))]
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                case 32767:
                                                    v2 -= v2
                                                default:
                                                    v71 := v1[-445]
                                                    _ = v71
                                                }
                                                var v72 int8 = (v2 >> ((153) & 63))
                                                _ = v68
                                                _ = v72
                                            default:
                                                v41 = 0
                                            }
                                        case 945:
                                            v42 = 42
                                        default:
                                            v73 := GenS10{F0: 42}
                                            _ = v73
                                        }
                                    }
                                    if (v3 >= (v42 | int16(v2))) {
                                        fmt.Println(v41)
                                        v74 := GenS11{F0: (v35 - v35), F1: v42}
                                        if (1 >= 65535) {
                                            if true {
                                                v41 += 0
                                            }
                                        } else {
                                            var v75 int = (fi39 & 9223372036854775807)
                                            _ = v75
                                        }
                                        _ = v74
                                    }
                                    switch (-(fi40)) {
                                    case -9223372036854775808:
                                        if true {
                                            var v76 int32 = -2147483648
                                            _ = v76
                                        } else {
                                            var v77 int = int(v35)
                                            _ = v77
                                        }
                                    case -1:
                                        var v78 int8 = v2
                                        _ = v78
                                    case -156:
                                        var v79 int32 = -872
                                        _ = v79
                                    }
                                    v35 += 1
                                    var v80 float64 = ((float64(3.14159) / ((1.5 / (1.5 / -2.25)) - 3.14159)) / 0.1)
                                    v1[fi39] = v35
                                    v81 := GenS12{F0: "foo bar", F1: true, F2: int32(v35)}
                                    _ = v80
                                    _ = v81
                                }
                                var v82 uint = 806
                                v1[(fi39 << (((v82 | v82)) & 63))] = uint32(v82)
                                _ = v82
                            }
                        default:
                            fmt.Println(v41, v41, fi40)
                        }
                        fmt.Println(v41, v42, v41)
                        v3 = v42
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v83 string = "\n\t"
                    var v84 int64 = -172
                    _ = fi39
                    _ = v83
                    _ = v84
                }
            case 98:
                v85 := GenS13{F0: -695, F1: -2147483648}
                _ = v85
            case 9223372036854775807:
                var v86 int16 = 716
                _ = v86
            }
        case 1:
            var v87 int8 = v2
            _ = v87
        case 32768:
            for fi88 := 0; fi88 < 18; fi88++ {
                var v89 uint = 0
                var v90 uint = v89
                _ = fi88
                _ = v89
                _ = v90
            }
        }
    default:
        v91 := GenS13{F0: -1, F1: 1}
        _ = v91
    }
    switch ("a" + (("\n\t" + "") + ("foo bar" + "foo bar"))) {
    case "":
        switch 348 {
        case 42:
            v35 -= (v35 << ((65535) & 63))
        case -188:
            _ = len(v34)
        case -1:
            var v92 bool = false
            _ = v92
        default:
            switch "" {
            case "foo bar":
                switch int(v2) {
                case 318:
                    for fi93 := 0; fi93 < 6; fi93++ {
                        var v94 int64 = -9223372036854775808
                        _ = fi93
                        _ = v94
                    }
                case 32767:
                    switch ("hello" + ("" + "a")) {
                    case "\n\t":
                        if len(v34) > 0 {
                            v34[0] = 723
                        }
                    default:
                        var v95 uint64 = 18446744073709551615
                        _ = v95
                    }
                }
            case "":
                fmt.Println(v3, v2, v35)
            case "\n\t":
                var v96 float32 = (-2.25 * 0.0)
                _ = v96
            }
        }
    case "a":
        v97 := GenS12{F0: (("a" + "") + (((("hello" + "hello") + "a") + "a") + "")), F1: (-2147483648 > 1), F2: 42}
        _ = v97
    case "hello":
        v2 += (v2 - (v2 / (((v2 | (v2 * v2))) | 1)))
    default:
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
}
