package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 int64; F1 int64 }
type GenS9 struct { F0 bool; F1 int32; F2 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int64; F1 int8; F2 int32 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 uint; F1 string }
type GenS17 struct { F0 uint8; F1 int8 }
type GenS18 struct { F0 int; F1 uint16; F2 int8 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 uint8; F1 int64 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 int16; F1 int }
type GenS25 struct { F0 bool; F1 uint16 }
type GenS26 struct { F0 string; F1 int64 }
type GenS27 struct { F0 int; F1 int; F2 bool }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 int8 }
type GenS30 struct { F0 uint32; F1 string }
type GenS31 struct { F0 int64; F1 int32 }
type GenS32 struct { F0 int8; F1 int64; F2 uint64 }
type GenS33 struct { F0 uint }
func genFn0(p0 bool, p1 uint8, p2 bool) uint {
    switch "a" {
    case "\n\t":
        var v7 int8 = 1
        _ = v7
    case "a":
        p0 = (32767 <= 1)
    }
    var v8 int32 = 1
    v9 := GenS1{F0: v8, F1: -639, F2: "foo bar"}
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v8, p1)
        var v11 uint = 42
        p1 += (p1 / ((p1) | 1))
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1() string {
    var v12 int16 = 1
    v13 := GenS1{F0: int32(v12), F1: (^(v12)), F2: "hello"}
    if false {
        v12 = v12
        v12 = (v12 / (((v12 * 1000)) | 1))
    }
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 int64 = 0
    var v2 uint8 = 0
    v3 := make(map[int]bool)
    var v4 uint64 = 42
    var v5 float64 = 3.14159
    var v6 uint = 683
    var v14 uint8 = (^(v2))
    v15 := GenS2{F0: true, F1: 2147483647}
    switch "\n\t" {
    case "":
        fmt.Println(v1)
    case "hello":
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            if true {
                v14 += v14
                if (!((v4 % (((v4 % (((v4 % (((v4 - 221)) | 1))) | 1))) | 1)) >= v4)) {
                    v16 := GenS1{F0: 1, F1: int16(v2), F2: ""}
                    v2 -= (^(v14))
                    _ = v16
                }
                if false {
                    delete(v3, 128)
                    var v17 bool = ((v5 / (v5 * ((float64(3.14159) + v5) - float64(-2.25)))) >= (v5 / 0.1))
                    if (true || (v17 || v17)) {
                        if (!(-2.25 >= v5)) {
                            var v18 int8 = -128
                            _ = v18
                        } else {
                            fmt.Println(v17)
                            var v19 uint16 = 65535
                            _ = v19
                        }
                        v20 := GenS2{F0: (!v17), F1: 219}
                        for fi21 := 0; fi21 < 17; fi21++ {
                            for fi22 := 0; fi22 < 8; fi22++ {
                                fmt.Println(v4, fi21)
                                _ = fi22
                            }
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v24 := GenS3{F0: 2147483648, F1: v14}
                            var v25 uint32 = 675
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 int16 = 42
                        _ = v26
                    }
                    _ = v17
                } else {
                    v27 := GenS2{F0: (42 <= 2), F1: 42}
                    _ = v27
                }
            }
        case "\n\t":
            v4 = v4
        case "a":
            v2 = (^(v2))
        }
    }
    var v28 uint32 = 0
    for fi29 := 0; fi29 < 16; fi29++ {
        if false {
            for fi30 := 0; fi30 < 10; fi30++ {
                v6 += (v6 / (((^(v6))) | 1))
                v31 := GenS4{F0: int64(v2), F1: (v28 << ((v6) & 63))}
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                v33 := GenS1{F0: -843, F1: 42, F2: "foo bar"}
                _ = fi32
                _ = v33
            }
            if (-199 > int32(v2)) {
                switch "" {
                case "foo bar":
                    v14 += v2
                case "hello":
                    for fi34 := 0; fi34 < 5; fi34++ {
                        switch ("\n\t" + "foo bar") {
                        case "":
                            if false {
                                var v35 bool = true
                                _ = v35
                            }
                        case "a":
                            v1 += v1
                        }
                        switch fi34 {
                        case -9223372036854775808:
                            switch (fi29 % ((fi34) | 1)) {
                            case 0:
                                v36 := GenS5{F0: 2, F1: (^(v28)), F2: -128}
                                _ = v36
                            case 42:
                                if ("foo bar" == ("\n\t" + "foo bar")) {
                                    v5 = float64(-2.25)
                                    var v37 uint32 = v28
                                    v5 = -2.25
                                    if (22 <= 42) {
                                        if false {
                                            switch (fi29 | (fi34 + (fi34 % ((-803) | 1)))) {
                                            case 782:
                                                v28 += v37
                                            case 0:
                                                if true {
                                                    delete(v3, fi29)
                                                    v38 := GenS2{F0: ((v2 + v14) >= (v2 ^ v2)), F1: fi34}
                                                    v5 = (v5 - v5)
                                                    _ = v38
                                                } else {
                                                    v37 = (v37 | uint32(v4))
                                                }
                                            case -2147483648:
                                                v5 += (3.14159 / ((1.5 + 1.5) - (((1e10 / 1.5) / v5) / float64((1.5 - (((3.14159 / 0.1) * 1.5) - 1e10))))))
                                                fmt.Println(v1, v28, v14)
                                                v1 = int64(v28)
                                                if (!(!true)) {
                                                    var v39 int16 = -32768
                                                    switch "\n\t" {
                                                    case "":
                                                        v40 := []int{fi29}
                                                        _ = v40
                                                    case "foo bar":
                                                        switch 32768 {
                                                        case 42:
                                                            v41 := GenS5{F0: v28, F1: 65535, F2: -1}
                                                            _ = v41
                                                        case -781:
                                                            if true {
                                                                var v42 uint16 = 2
                                                                v43 := []int{fi29, fi29, -621}
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                v44 := GenS5{F0: (^(v37)), F1: 509, F2: -128}
                                                                switch 0 {
                                                                case -594:
                                                                    v14 -= 128
                                                                case 914:
                                                                    var v45 uint32 = (v28 >> (((v6 + v6)) & 63))
                                                                    if (!((0.0 + (float64((0.1 * -2.25)) * v5)) <= (((v5 + (0.0 + 3.14159)) - 1.5) + v5))) {
                                                                        var v46 uint64 = v4
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v47 int16 = v39
                                                                            var v48 float32 = 1.5
                                                                            var v49 float32 = 3.14159
                                                                            v39 += v39
                                                                            _ = v47
                                                                            _ = v48
                                                                            _ = v49
                                                                        }
                                                                        _ = v46
                                                                    } else {
                                                                        delete(v3, fi34)
                                                                        v50 := GenS6{F0: (-111 == int8(v4)), F1: v1}
                                                                        _ = v50
                                                                    }
                                                                    v4 += 2
                                                                    _ = v45
                                                                case -9223372036854775808:
                                                                    if (1 < 936) {
                                                                        v3[fi29] = false
                                                                        if ((1.5 / (1e10 - 1.5)) <= (0.1 / 3.14159)) {
                                                                            var v51 uint = (v6 & v6)
                                                                            var v52 float32 = 0.0
                                                                            var v53 int8 = 42
                                                                            v54 := GenS1{F0: 2147483647, F1: (v39 & (v39 % ((0) | 1))), F2: ""}
                                                                            switch fi34 {
                                                                            case 42:
                                                                                if false {
                                                                                    switch fi29 {
                                                                                    case -1:
                                                                                        switch fi29 {
                                                                                        case -2147483648:
                                                                                            if (!false) {
                                                                                                v55 := GenS1{F0: 1, F1: (v39 << ((42) & 63)), F2: ""}
                                                                                                switch 4294967295 {
                                                                                                case -479:
                                                                                                    if false {
                                                                                                        v5 += 0.0
                                                                                                    }
                                                                                                case 2147483647:
                                                                                                    if ((int(v14) != fi34) && true) {
                                                                                                        switch ("" + "") {
                                                                                                        case "\n\t":
                                                                                                            v28 -= (^(v28))
                                                                                                        default:
                                                                                                            var v56 uint64 = v4
                                                                                                            _ = v56
                                                                                                        }
                                                                                                        fmt.Println(v28, v37)
                                                                                                        if true {
                                                                                                            v5 -= ((v5 + (v5 + v5)) / ((float64(v52) + 3.14159) / v5))
                                                                                                        } else {
                                                                                                            v57 := []int{42, fi34, fi29}
                                                                                                            if ((v37 & (^(v28))) != (^(v28))) {
                                                                                                                var v58 uint32 = 0
                                                                                                                _ = v58
                                                                                                            }
                                                                                                            _ = v57
                                                                                                        }
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                        fmt.Println(fi29, v37)
                                                                                                    }
                                                                                                default:
                                                                                                    if ((((("\n\t" + "\n\t") + ("hello" + "foo bar")) + "a") + "a") != "a") {
                                                                                                        fmt.Println(v53, v1)
                                                                                                        var v59 uint32 = 4294967295
                                                                                                        _ = v59
                                                                                                    }
                                                                                                }
                                                                                                var v60 uint8 = 100
                                                                                                _ = v55
                                                                                                _ = v60
                                                                                            }
                                                                                        default:
                                                                                            var v61 int64 = v1
                                                                                            _ = v61
                                                                                        }
                                                                                    case 213:
                                                                                        var v62 int8 = v53
                                                                                        _ = v62
                                                                                    case 127:
                                                                                        var v63 int8 = 51
                                                                                        _ = v63
                                                                                    }
                                                                                } else {
                                                                                    var v64 int16 = (v39 ^ v39)
                                                                                    if ((v1 * (v1 * int64(v6))) <= v1) {
                                                                                        v65 := []int{int(v1)}
                                                                                        v2 = (^(v14))
                                                                                        _ = len(v3)
                                                                                        _ = v65
                                                                                    }
                                                                                    v51 = (v6 ^ 554)
                                                                                    if (65535 >= (v51 + v6)) {
                                                                                        if false {
                                                                                            v66 := GenS4{F0: v1, F1: (v28 << ((genFn0(false, (^(v2)), ("foo bar" < "foo bar"))) & 63))}
                                                                                            v14 = 1
                                                                                            switch (("hello" + "\n\t") + ("a" + "hello")) {
                                                                                            case "foo bar":
                                                                                                delete(v3, fi34)
                                                                                                v1 -= v1
                                                                                                v28 = (v37 * v28)
                                                                                            case "a":
                                                                                                var v67 bool = (!false)
                                                                                                _ = v67
                                                                                            case "\n\t":
                                                                                                if (1 < v4) {
                                                                                                    if (!true) {
                                                                                                        v68 := []int{-9223372036854775808, fi29, (fi34 / (((fi34 / (((-(fi29))) | 1))) | 1))}
                                                                                                        fmt.Println(v37)
                                                                                                        v69 := GenS7{F0: v53, F1: v28}
                                                                                                        _ = v68
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    v28 = (^(v37))
                                                                                                    if (v1 < (v1 + 791)) {
                                                                                                        var v70 int8 = (v53 ^ v53)
                                                                                                        fmt.Println(v70)
                                                                                                        _ = v70
                                                                                                    }
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if ((!(v39 != v64)) || false) {
                                                                                                            var v71 int32 = 1
                                                                                                            _ = v71
                                                                                                        } else {
                                                                                                            v2 += (v2 * (^(v2)))
                                                                                                            v72 := GenS8{F0: v1, F1: (v1 - (v1 / ((-1) | 1)))}
                                                                                                            _ = v72
                                                                                                        }
                                                                                                        var v73 bool = false
                                                                                                        _ = v73
                                                                                                    } else {
                                                                                                        fmt.Println(v53, v53)
                                                                                                        v74 := []int{(fi29 >> (((v6 << ((uint(v4)) & 63))) & 63)), 4294967295}
                                                                                                        _ = v74
                                                                                                    }
                                                                                                    v64 = (v64 % ((v39) | 1))
                                                                                                    if true {
                                                                                                        fmt.Println(v2, v37)
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                v64 += v39
                                                                                            }
                                                                                            _ = v66
                                                                                        } else {
                                                                                            if true {
                                                                                                var v75 float64 = float64(v52)
                                                                                                if (65 > v53) {
                                                                                                    if ((v52 > v52) && false) {
                                                                                                        switch (fi29 % ((fi29) | 1)) {
                                                                                                        case 731:
                                                                                                            var v76 int = fi29
                                                                                                            _ = v76
                                                                                                        case 687:
                                                                                                            if false {
                                                                                                                switch fi29 {
                                                                                                                case 9223372036854775807:
                                                                                                                    v77 := []int{int(v1)}
                                                                                                                    _ = v77
                                                                                                                }
                                                                                                                delete(v3, fi34)
                                                                                                                if (!false) {
                                                                                                                    v78 := GenS1{F0: -2147483648, F1: v39, F2: "a"}
                                                                                                                    if ((0.0 < (v52 * float32((float64(v52) / v5)))) && false) {
                                                                                                                        var v79 int8 = v53
                                                                                                                        var v80 int32 = 0
                                                                                                                        _ = v79
                                                                                                                        _ = v80
                                                                                                                    } else {
                                                                                                                        if ((true && true) || (v14 == (v2 + uint8(v4)))) {
                                                                                                                            v81 := GenS9{F0: false, F1: int32(v14), F2: (v4 & uint64(v37))}
                                                                                                                            _ = v81
                                                                                                                        }
                                                                                                                        v3[(fi34 + fi29)] = false
                                                                                                                        var v82 float32 = v52
                                                                                                                        _ = v82
                                                                                                                    }
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                            } else {
                                                                                                                fmt.Println(fi29)
                                                                                                                if (true || false) {
                                                                                                                    v83 := GenS4{F0: 1, F1: (^(v37))}
                                                                                                                    _ = v83
                                                                                                                }
                                                                                                                v84 := GenS6{F0: (v4 <= (v4 + (v4 / (((v4 | (^(v4)))) | 1)))), F1: (v1 << ((18446744073709551615) & 63))}
                                                                                                                _ = v84
                                                                                                            }
                                                                                                        }
                                                                                                        v85 := v3[-2147483648]
                                                                                                        if ((v51 % ((v51) | 1)) < v51) {
                                                                                                            var v86 uint8 = (v2 & v2)
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        _ = v85
                                                                                                    }
                                                                                                    var v87 uint8 = 128
                                                                                                    fmt.Println(v51)
                                                                                                    switch "a" {
                                                                                                    case "\n\t":
                                                                                                        v6 = (^(v6))
                                                                                                    }
                                                                                                    _ = v87
                                                                                                }
                                                                                                v5 += 1e10
                                                                                                _ = v75
                                                                                            } else {
                                                                                                if true {
                                                                                                    v39 = -872
                                                                                                } else {
                                                                                                    v88 := GenS10{F0: true}
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            v89 := []int{(-(fi29)), 705}
                                                                                                            _ = v89
                                                                                                        } else {
                                                                                                            v90 := []int{(fi34 * (^(fi34))), (fi29 * fi34), int(v51)}
                                                                                                            _ = v90
                                                                                                        }
                                                                                                        fmt.Println(v64, v51)
                                                                                                        fmt.Println(v2, v28, v51)
                                                                                                    }
                                                                                                    v91 := v3[int(v4)]
                                                                                                    _ = v88
                                                                                                    _ = v91
                                                                                                }
                                                                                                v92 := GenS11{F0: -1, F1: (v53 * 0), F2: -2147483648}
                                                                                                var v93 uint64 = v4
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                            }
                                                                                        }
                                                                                        if true {
                                                                                            if true {
                                                                                                var v94 bool = true
                                                                                                v14 -= (v2 | (v14 / (((^(v14))) | 1)))
                                                                                                _ = v94
                                                                                            } else {
                                                                                                var v95 uint16 = 2
                                                                                                v6 += v51
                                                                                                v5 = ((float64((0.0 / v52)) * 1e10) / v5)
                                                                                                _ = v95
                                                                                            }
                                                                                            if true {
                                                                                                v39 -= v64
                                                                                                if ((v1 <= (v1 - (-(v1)))) && true) {
                                                                                                    if (42 != 2) {
                                                                                                        if (!true) {
                                                                                                            if true {
                                                                                                                var v96 string = ""
                                                                                                                v97 := GenS12{F0: 32767}
                                                                                                                _ = v96
                                                                                                                _ = v97
                                                                                                            } else {
                                                                                                                v4 = uint64(v28)
                                                                                                            }
                                                                                                            if false {
                                                                                                                v51 += 18446744073709551615
                                                                                                            } else {
                                                                                                                fmt.Println(v64, v5)
                                                                                                                v98 := v3[2147483647]
                                                                                                                v99 := GenS13{F0: -421}
                                                                                                                v100 := []int{(fi34 / (((^(fi29))) | 1)), int(v14), (fi29 ^ fi29)}
                                                                                                                v101 := GenS14{F0: (v1 % ((v1) | 1))}
                                                                                                                var v102 uint32 = (v37 & 0)
                                                                                                                _ = v100
                                                                                                                _ = v101
                                                                                                                _ = v102
                                                                                                                _ = v98
                                                                                                                _ = v99
                                                                                                            }
                                                                                                            var v103 int8 = v53
                                                                                                            _ = v103
                                                                                                        } else {
                                                                                                            v51 = (v6 / ((genFn0(true, v2, false)) | 1))
                                                                                                            var v104 uint32 = 705
                                                                                                            v6 += 773
                                                                                                            _ = v104
                                                                                                        }
                                                                                                        v105 := GenS15{F0: (v4 % (((v4 | uint64(v2))) | 1)), F1: (v1 % ((25) | 1))}
                                                                                                        _ = v105
                                                                                                    } else {
                                                                                                        fmt.Println(v51, v39)
                                                                                                        if true {
                                                                                                            var v106 uint64 = 2
                                                                                                            v107 := GenS14{F0: v1}
                                                                                                            _ = v106
                                                                                                            _ = v107
                                                                                                        } else {
                                                                                                            v108 := GenS16{F0: genFn0((((0.1 - v52) - ((v52 + (v52 / v52)) - float32((-2.25 * v5)))) == 3.14159), uint8(v4), true), F1: (("\n\t" + ("" + "")) + "a")}
                                                                                                            _ = v108
                                                                                                        }
                                                                                                        var v109 uint32 = v28
                                                                                                        var v110 uint8 = v14
                                                                                                        v3[fi29] = (false || (((v14 % ((127) | 1)) < v110) && false))
                                                                                                        _ = v109
                                                                                                        _ = v110
                                                                                                    }
                                                                                                    switch (fi34 * (fi34 << ((v51) & 63))) {
                                                                                                    case 2147483647:
                                                                                                        v111 := GenS17{F0: v2, F1: v53}
                                                                                                        _ = v111
                                                                                                    default:
                                                                                                        switch ((("a" + ("\n\t" + "")) + ("a" + "\n\t")) + ("hello" + ("" + ((((("\n\t" + "a") + "hello") + "foo bar") + "hello") + ("a" + "foo bar"))))) {
                                                                                                        case "foo bar":
                                                                                                            v112 := GenS18{F0: int(v53), F1: uint16(v6), F2: 1}
                                                                                                            v113 := []int{(^(fi34)), (fi34 & fi34), (fi34 ^ fi34), fi34}
                                                                                                            _ = v112
                                                                                                            _ = v113
                                                                                                        default:
                                                                                                            v114 := v3[fi29]
                                                                                                            _ = v114
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v2, fi34)
                                                                                                }
                                                                                            }
                                                                                            v51 = (v51 | genFn0((0 <= (v39 - (v64 | (v64 ^ 1)))), uint8(v4), true))
                                                                                        }
                                                                                    } else {
                                                                                        switch ((("foo bar" + "\n\t") + "\n\t") + (("\n\t" + "hello") + "a")) {
                                                                                        case "a":
                                                                                            var v115 int16 = v39
                                                                                            _ = v115
                                                                                        }
                                                                                    }
                                                                                    _ = v64
                                                                                }
                                                                            case -1:
                                                                                v6 -= 691
                                                                            }
                                                                            if true {
                                                                                v2 -= v14
                                                                            }
                                                                            if true {
                                                                                v116 := GenS5{F0: (v37 % ((v37) | 1)), F1: (^(v28)), F2: (^(v53))}
                                                                                _ = v116
                                                                            }
                                                                            v117 := v3[(-(fi29))]
                                                                            _ = v117
                                                                            _ = v51
                                                                            _ = v52
                                                                            _ = v53
                                                                            _ = v54
                                                                        } else {
                                                                            var v118 string = "foo bar"
                                                                            _ = v118
                                                                        }
                                                                        if (!((2 >= 1) && (2 < v2))) {
                                                                            var v119 uint16 = 2
                                                                            var v120 int16 = v39
                                                                            _ = v119
                                                                            _ = v120
                                                                        } else {
                                                                            v121 := GenS9{F0: (((v6 << ((uint(v4)) & 63)) <= (v6 % ((v6) | 1))) && true), F1: 33, F2: v4}
                                                                            v37 = 4294967295
                                                                            v1 += int64(v6)
                                                                            var v122 int16 = int16(v1)
                                                                            _ = v121
                                                                            _ = v122
                                                                        }
                                                                    } else {
                                                                        v1 -= v1
                                                                        v1 = v1
                                                                    }
                                                                    v28 = (v28 & v37)
                                                                    delete(v3, (fi29 / ((-1) | 1)))
                                                                default:
                                                                    if true {
                                                                        fmt.Println(fi34, v28)
                                                                    }
                                                                }
                                                                v123 := GenS18{F0: int(v28), F1: 255, F2: -103}
                                                                _ = v123
                                                                _ = v44
                                                            }
                                                        case 9223372036854775807:
                                                            if (1.5 != ((float32(v5) / 1e10) + (0.0 + 0.1))) {
                                                                var v124 int8 = 0
                                                                v125 := GenS9{F0: true, F1: 0, F2: v4}
                                                                if ((-(v1)) != -1) {
                                                                    switch int(v6) {
                                                                    case -346:
                                                                        v126 := GenS4{F0: (^(v1)), F1: (v37 & v28)}
                                                                        _ = v126
                                                                    case 128:
                                                                        if (false && (-369 >= -1)) {
                                                                            if false {
                                                                                v127 := GenS19{F0: (v1 + int64(fi29))}
                                                                                fmt.Println(v14)
                                                                                v128 := GenS3{F0: (v37 - 0), F1: 58}
                                                                                if false {
                                                                                    var v129 float32 = 0.0
                                                                                    v37 = v28
                                                                                    _ = v129
                                                                                } else {
                                                                                    var v130 uint = (v6 / (((v6 % (((v6 << (((v6 / ((v6) | 1))) & 63))) | 1))) | 1))
                                                                                    _ = v130
                                                                                }
                                                                                var v131 int64 = 373
                                                                                _ = v127
                                                                                _ = v128
                                                                                _ = v131
                                                                            }
                                                                            var v132 float32 = ((1.5 + -2.25) / 1.5)
                                                                            v133 := GenS2{F0: true, F1: -757}
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                if false {
                                                                                    v6 = v6
                                                                                    var v134 int = 32767
                                                                                    v134 = (^(fi34))
                                                                                    _ = v134
                                                                                } else {
                                                                                    v135 := []int{fi34, fi34}
                                                                                    v136 := GenS16{F0: v6, F1: ""}
                                                                                    var v137 float32 = float32(((1e10 / v5) / 1.0))
                                                                                    _ = v135
                                                                                    _ = v136
                                                                                    _ = v137
                                                                                }
                                                                            case "a":
                                                                                v138 := GenS20{F0: (v14 / (((v14 - v14)) | 1)), F1: 1}
                                                                                _ = v138
                                                                            case "\n\t":
                                                                                var v139 uint16 = 0
                                                                                _ = v139
                                                                            }
                                                                            _ = v132
                                                                            _ = v133
                                                                        } else {
                                                                            v39 = (v39 | -1)
                                                                        }
                                                                    case -1:
                                                                        var v140 bool = (v124 < (^(v124)))
                                                                        _ = v140
                                                                    default:
                                                                        if ((v39 << ((65535) & 63)) != (v39 >> ((v6) & 63))) {
                                                                            v124 = (v124 / (((v124 / ((-128) | 1))) | 1))
                                                                            v6 = v6
                                                                        } else {
                                                                            var v141 uint32 = 0
                                                                            _ = v141
                                                                        }
                                                                    }
                                                                    v14 += (v14 + v2)
                                                                    if (462 > 1) {
                                                                        v142 := GenS21{F0: (!true)}
                                                                        if ((v28 ^ 262) >= (v37 + (v37 - (v37 << ((218) & 63))))) {
                                                                            switch (fi29 - (fi34 & (fi29 | (fi29 + (-(fi29)))))) {
                                                                            case 32768:
                                                                                v143 := GenS14{F0: (v1 / (((v1 ^ v1)) | 1))}
                                                                                _ = v143
                                                                            case 32767:
                                                                                var v144 int = 42
                                                                                _ = v144
                                                                            case 9223372036854775807:
                                                                                v2 -= 255
                                                                            }
                                                                            var v145 int8 = int8(fi29)
                                                                            if ((!(!true)) || (0 == 0)) {
                                                                                var v146 uint16 = 0
                                                                                v147 := GenS22{F0: 65535}
                                                                                _ = v146
                                                                                _ = v147
                                                                            } else {
                                                                                fmt.Println(v124, v2)
                                                                                var v148 float32 = (1e10 + -2.25)
                                                                                _ = v148
                                                                            }
                                                                            _ = v145
                                                                        }
                                                                        _ = v142
                                                                    }
                                                                } else {
                                                                    var v149 int32 = -1
                                                                    _ = v149
                                                                }
                                                                _ = v124
                                                                _ = v125
                                                            } else {
                                                                v37 = v28
                                                            }
                                                        }
                                                        var v150 uint = (v6 + v6)
                                                        _ = v150
                                                    case "\n\t":
                                                        v4 -= v4
                                                    }
                                                    _ = v39
                                                } else {
                                                    var v151 string = ""
                                                    if true {
                                                        var v152 int = fi34
                                                        var v153 string = ("" + "hello")
                                                        if true {
                                                            if (!(6 < uint8(v4))) {
                                                                v154 := GenS23{F0: 351}
                                                                _ = v154
                                                            }
                                                            var v155 float32 = ((0.0 - float32((3.14159 / 0.1))) - 0.1)
                                                            v156 := GenS7{F0: 46, F1: v37}
                                                            v157 := []int{-1, v152, (fi34 & fi29)}
                                                            fmt.Println(fi29, v5)
                                                            _ = v155
                                                            _ = v156
                                                            _ = v157
                                                        } else {
                                                            fmt.Println(v1)
                                                        }
                                                        _ = v152
                                                        _ = v153
                                                    } else {
                                                        v4 = 2
                                                    }
                                                    _ = v151
                                                }
                                            }
                                            v158 := []int{-2147483648, fi29, (fi29 & fi29)}
                                            v5 = (0.0 / v5)
                                            switch fi34 {
                                            case -1:
                                                fmt.Println(v1)
                                            case 195:
                                                var v159 float64 = float64((((1.5 / (((0.0 + (3.14159 / (0.1 - (-2.25 * 0.1)))) * (-2.25 / -2.25)) * 1e10)) - 1.5) - -2.25))
                                                _ = v159
                                            case -225:
                                                var v160 float32 = float32(0.0)
                                                _ = v160
                                            default:
                                                if (v6 > 2) {
                                                    v28 = v37
                                                    v14 = 151
                                                    if (true && true) {
                                                        _ = len(v3)
                                                        v161 := GenS22{F0: uint(v14)}
                                                        if (v37 < v37) {
                                                            fmt.Println(v4, fi34)
                                                            v6 += 65535
                                                        } else {
                                                            var v162 float64 = v5
                                                            if (((!((((0.1 + (v162 / ((float64(1e10) * (v162 / v5)) * v162))) + (((v162 + v162) - (float64(1e10) + v5)) + v162)) - (float64((1.5 * 1e10)) - v5)) >= (((-2.25 - (v5 - -2.25)) - 0.1) + 3.14159))) || true) || true) {
                                                                var v163 int32 = -1
                                                                v164 := v3[-95]
                                                                switch int(v163) {
                                                                case -9223372036854775808:
                                                                    var v165 float64 = (v162 / (3.14159 - v5))
                                                                    _ = v165
                                                                case -454:
                                                                    var v166 int16 = -724
                                                                    v167 := GenS24{F0: 1, F1: (fi34 >> ((2) & 63))}
                                                                    fmt.Println(fi29, v5)
                                                                    _ = v166
                                                                    _ = v167
                                                                case -745:
                                                                    var v168 int8 = 48
                                                                    _ = v168
                                                                default:
                                                                    v4 -= v4
                                                                }
                                                                _ = v163
                                                                _ = v164
                                                            } else {
                                                                var v169 int = fi29
                                                                if true {
                                                                    v14 -= v14
                                                                }
                                                                v170 := GenS25{F0: true, F1: 65535}
                                                                _ = v169
                                                                _ = v170
                                                            }
                                                            _ = v162
                                                        }
                                                        v6 += genFn0((1 < uint16(v6)), (^(v2)), (true || true))
                                                        var v171 uint32 = (v28 % ((42) | 1))
                                                        _ = v161
                                                        _ = v171
                                                    } else {
                                                        v172 := GenS26{F0: "", F1: v1}
                                                        _ = v172
                                                    }
                                                    switch (fi34 + (fi34 - fi34)) {
                                                    case 4294967295:
                                                        var v173 int = fi34
                                                        _ = v173
                                                    case -565:
                                                        if true {
                                                            v174 := GenS27{F0: -9223372036854775808, F1: (fi29 ^ fi29), F2: ((-(v1)) == v1)}
                                                            v158 = append(v158, (-(fi34)))
                                                            switch 42 {
                                                            case -821:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v14 += uint8(v37)
                                                                case "":
                                                                    v6 = v6
                                                                }
                                                            case 32768:
                                                                if ((fi34 / (((fi29 - fi29)) | 1)) >= (^(fi34))) {
                                                                    if ((1 == 99) && false) {
                                                                        v4 -= (v4 ^ v4)
                                                                    }
                                                                }
                                                            }
                                                            _ = v174
                                                        }
                                                        switch 999 {
                                                        case -751:
                                                            var v175 int8 = -1
                                                            _ = v175
                                                        }
                                                        fmt.Println(v1, v14, v2)
                                                    case 42:
                                                        switch "foo bar" {
                                                        case "":
                                                            var v176 int16 = 32767
                                                            _ = v176
                                                        default:
                                                            if len(v158) >= 2 {
                                                                v158 = v158[1:len(v158)]
                                                            }
                                                        }
                                                    }
                                                    var v177 int32 = 42
                                                    _ = v177
                                                }
                                            }
                                            _ = v158
                                        }
                                        if true {
                                            if (v4 >= v4) {
                                                var v178 bool = (95 <= 673)
                                                _ = v178
                                            }
                                            var v179 int16 = 42
                                            v14 = (v2 ^ 24)
                                            _ = v179
                                        } else {
                                            var v180 bool = (true && (-1 > 42))
                                            _ = v180
                                        }
                                        var v181 uint64 = uint64(v37)
                                        _ = v181
                                    } else {
                                        if true {
                                            if ((!(!true)) && true) {
                                                fmt.Println(v37, fi29)
                                                if true {
                                                    fmt.Println(fi29, v1, v4)
                                                    if (fi29 < 127) {
                                                        if true {
                                                            switch "" {
                                                            case "hello":
                                                                if true {
                                                                    var v182 bool = true
                                                                    var v183 uint32 = (v37 >> (((v6 << ((v6) & 63))) & 63))
                                                                    v184 := v3[fi34]
                                                                    fmt.Println(v14, v2, v28)
                                                                    _ = v182
                                                                    _ = v183
                                                                    _ = v184
                                                                }
                                                            case "foo bar":
                                                                if false {
                                                                    v185 := GenS10{F0: false}
                                                                    var v186 uint16 = uint16(v6)
                                                                    switch (fi29 << (((v6 / (((v6 >> ((v6) & 63))) | 1))) & 63)) {
                                                                    case 1:
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            delete(v3, (fi34 | (fi34 / ((-1) | 1))))
                                                                        default:
                                                                            v187 := GenS28{F0: (v1 % ((int64(v6)) | 1))}
                                                                            _ = v187
                                                                        }
                                                                    case 128:
                                                                        v188 := v3[(fi29 * (-(fi29)))]
                                                                        _ = v188
                                                                    default:
                                                                        delete(v3, (-(fi34)))
                                                                    }
                                                                    _ = v185
                                                                    _ = v186
                                                                }
                                                            case "":
                                                                fmt.Println(v2, v1)
                                                            }
                                                        }
                                                    } else {
                                                        var v189 bool = (false || (!false))
                                                        fmt.Println(v4, v14, v4)
                                                        _ = v189
                                                    }
                                                    switch (fi34 + -298) {
                                                    case 4294967295:
                                                        var v190 int16 = -32768
                                                        _ = v190
                                                    default:
                                                        if true {
                                                            v5 -= v5
                                                            v191 := GenS28{F0: -9223372036854775808}
                                                            v28 = (^(v28))
                                                            switch "a" {
                                                            case "foo bar":
                                                                v1 = (v1 / ((v1) | 1))
                                                            default:
                                                                var v192 string = (("hello" + "foo bar") + "a")
                                                                _ = v192
                                                            }
                                                            _ = len(v3)
                                                            _ = v191
                                                        }
                                                    }
                                                }
                                                v1 -= v1
                                            }
                                            if (v14 >= (v14 | v2)) {
                                                v193 := GenS29{F0: 1}
                                                switch "hello" {
                                                case "a":
                                                    v194 := v3[fi29]
                                                    _ = v194
                                                case "":
                                                    v1 -= (v1 | v1)
                                                case "hello":
                                                    if true {
                                                        switch "" {
                                                        case "foo bar":
                                                            var v195 float64 = -2.25
                                                            _ = v195
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v14 += v14
                                                            }
                                                        case "hello":
                                                            v196 := v3[(fi34 * fi29)]
                                                            _ = v196
                                                        }
                                                        v5 = float64(3.14159)
                                                        v2 += v14
                                                    }
                                                default:
                                                    v14 -= uint8(v28)
                                                }
                                                _ = v193
                                            } else {
                                                v197 := GenS9{F0: true, F1: 0, F2: v4}
                                                var v198 int32 = 42
                                                _ = v197
                                                _ = v198
                                            }
                                            switch (fi34 ^ fi34) {
                                            case 2147483647:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v4 -= 472
                                                case "hello":
                                                    v1 = 42
                                                default:
                                                    v199 := GenS14{F0: v1}
                                                    _ = v199
                                                }
                                            }
                                        }
                                    }
                                    v1 = 197
                                    v37 = (v37 << ((v6) & 63))
                                    var v200 bool = ((fi34 & fi34) != (fi29 << ((genFn0(((v5 * (float64((float32(1.5) - 3.14159)) - v5)) != 0.0), (^(v14)), false)) & 63)))
                                    _ = v200
                                    _ = v37
                                }
                            case 32768:
                                if true {
                                    var v201 uint = 2
                                    var v202 bool = (!true)
                                    v14 = v2
                                    _ = v201
                                    _ = v202
                                } else {
                                    v2 += (v14 - v2)
                                    v203 := GenS30{F0: (v28 * (v28 & v28)), F1: ("hello" + "\n\t")}
                                    _ = v203
                                }
                            default:
                                delete(v3, fi34)
                            }
                        case 1:
                            fmt.Println(v28)
                        case 69:
                            var v204 string = "\n\t"
                            _ = v204
                        }
                        _ = fi34
                    }
                case "":
                    for fi205 := 0; fi205 < 7; fi205++ {
                        v28 = v28
                        v2 += (v2 ^ 2)
                        _ = fi205
                    }
                }
            } else {
                v206 := GenS2{F0: false, F1: fi29}
                var v207 int32 = 2147483647
                var v208 int = fi29
                _ = v206
                _ = v207
                _ = v208
            }
        } else {
            v209 := GenS31{F0: (v1 + v1), F1: 0}
            _ = v209
        }
        v210 := GenS2{F0: false, F1: fi29}
        fmt.Println(v2)
        _ = fi29
        _ = v210
    }
    if false {
        fmt.Println(v14)
    } else {
        for fi211 := 0; fi211 < 20; fi211++ {
            var v212 int = (fi211 & fi211)
            v213 := GenS32{F0: 5, F1: v1, F2: (v4 - v4)}
            v214 := v3[fi211]
            _ = fi211
            _ = v212
            _ = v213
            _ = v214
        }
        var v215 int64 = 0
        _ = v215
    }
    v216 := GenS8{F0: 0, F1: 9223372036854775807}
    v217 := []int{32767, 127, 325, 42}
    v14 = v2
    for fi218 := 0; fi218 < 1; fi218++ {
        if ((v5 - 0.0) <= v5) {
            if ((3.14159 + (3.14159 + (-2.25 + 1.5))) > 0.0) {
                if len(v217) > 0 {
                    v217[0] = fi218
                }
                switch "hello" {
                case "hello":
                    var v219 float32 = 1e10
                    _ = v219
                case "":
                    var v220 float32 = (1e10 + 1.5)
                    _ = v220
                case "foo bar":
                    v6 = v6
                }
                var v221 uint = genFn0((("hello" >= ((("hello" + ("a" + (("" + ("\n\t" + "\n\t")) + ("\n\t" + ("\n\t" + "foo bar"))))) + "hello") + "\n\t")) || (!(1e10 == 0.1))), (^(v14)), (((((1.5 / float64(3.14159)) + v5) - (v5 * float64((0.0 + 0.1)))) + v5) < v5))
                _ = v221
            }
        }
        _ = fi218
    }
    var v222 int16 = 32767
    if true {
        if false {
            var v223 float64 = ((1.5 * v5) * v5)
            switch int(v1) {
            case 9223372036854775807:
                switch int(v6) {
                case -2147483648:
                    if (!((!false) || (true && false))) {
                        var v224 uint8 = (v14 ^ v14)
                        _ = len(v217)
                        if false {
                            v222 += (v222 >> (((v6 + v6)) & 63))
                            if (v28 <= uint32(v4)) {
                                if (1.5 <= 1e10) {
                                    _ = len(v3)
                                    var v225 uint = v6
                                    fmt.Println(v4, v223, v4)
                                    _ = v225
                                } else {
                                    var v226 int8 = 0
                                    _ = v226
                                }
                                var v227 int64 = (^(v1))
                                _ = v227
                            } else {
                                var v228 uint = 255
                                _ = v228
                            }
                            v6 = v6
                        }
                        _ = v224
                    } else {
                        if false {
                            var v229 float32 = 1e10
                            switch int(v14) {
                            case 499:
                                var v230 uint32 = uint32(v6)
                                _ = v230
                            default:
                                switch -2147483648 {
                                case 127:
                                    if true {
                                        if len(v217) >= 2 {
                                            v217 = v217[1:len(v217)]
                                        }
                                    }
                                default:
                                    v231 := GenS1{F0: -846, F1: 32767, F2: (("a" + "a") + "a")}
                                    _ = v231
                                }
                            }
                            fmt.Println(v1)
                            _ = v229
                        } else {
                            for fi232 := 0; fi232 < 12; fi232++ {
                                if len(v217) > 0 {
                                    v217[0] = 4294967295
                                }
                                _ = fi232
                            }
                            if (((v4 * 660) <= v4) || (!true)) {
                                if (!(v28 <= (v28 << ((v6) & 63)))) {
                                    var v233 int32 = 229
                                    v222 = v222
                                    _ = v233
                                }
                                if len(v217) >= 2 {
                                    v217 = v217[1:len(v217)]
                                }
                            }
                        }
                        var v234 int = -159
                        _ = v234
                    }
                }
            }
            if ((v6 >> ((v6) & 63)) != genFn0((381 == 2), (^(v2)), false)) {
                var v235 float64 = (v223 - 1e10)
                v6 = v6
                if false {
                    v235 = (v223 * v5)
                    fmt.Println(v2, v28, v1)
                    v236 := GenS33{F0: (^(v6))}
                    _ = v236
                } else {
                    if len(v217) >= 2 {
                        v217 = v217[1:len(v217)]
                    }
                    var v237 int8 = -1
                    var v238 float32 = float32(v223)
                    _ = v237
                    _ = v238
                }
                _ = v235
            } else {
                switch "\n\t" {
                case "foo bar":
                    v28 -= (v28 & v28)
                }
                fmt.Println(v5, v14)
            }
            _ = v223
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v217:", len(v217), v217)
    fmt.Println("v222:", v222)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v216
    _ = v217
    _ = v222
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
