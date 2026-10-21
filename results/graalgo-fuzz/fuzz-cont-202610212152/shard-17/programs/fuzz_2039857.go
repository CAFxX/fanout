package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 string; F1 int64; F2 int64 }
type GenS7 struct { F0 int32; F1 bool }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 uint16; F1 int64; F2 uint64 }

func main() {
    var v1 uint64 = 660
    v2 := make(map[string]int)
    var v3 int16 = 703
    v4 := make(map[string]bool)
    var v5 float64 = (0.1 + 0.1)
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{}
            var v10 uint8 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
        if false {
            v11 := GenS1{F0: 1, F1: 187}
            if true {
                if true {
                    if false {
                        fmt.Println(v5, v1)
                        if false {
                            for fi12 := 0; fi12 < 10; fi12++ {
                                v13 := []int{(fi7 / ((fi12) | 1)), (fi7 % (((fi12 - fi12)) | 1)), fi12, fi12}
                                var v14 float64 = v5
                                if (((((((0.1 + -2.25) - (v5 / v14)) / v5) * v5) + v14) - (0.0 * (((1.5 + v14) - v14) - -2.25))) < ((float64((1e10 * (1.5 + 0.1))) * v5) / (v14 + 0.1))) {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v1 = v1
                                    v15 := GenS1{F0: 2147483647, F1: 185}
                                    _ = v15
                                } else {
                                    var v16 uint = 633
                                    var v17 string = ("" + (("foo bar" + "hello") + ("\n\t" + (("hello" + "") + ""))))
                                    if false {
                                        switch (fi7 / ((fi12) | 1)) {
                                        case 128:
                                            v1 += v1
                                        case -40:
                                            v18 := v2[v17]
                                            _ = v18
                                        case 1:
                                            if true {
                                                var v19 string = "\n\t"
                                                v20 := GenS1{F0: 263, F1: 255}
                                                _ = v19
                                                _ = v20
                                            }
                                        }
                                    }
                                    _ = v16
                                    _ = v17
                                }
                                if (127 != (fi7 % (((-(fi12))) | 1))) {
                                    if false {
                                        switch (("" + "") + "a") {
                                        case "foo bar":
                                            v5 -= 1.5
                                        case "a":
                                            if true {
                                                fmt.Println(v5, v1)
                                                var v21 int32 = 764
                                                _ = v21
                                            } else {
                                                v3 = v3
                                                v2[("" + (("" + ("a" + "")) + "foo bar"))] = int(v3)
                                            }
                                        case "\n\t":
                                            if (!false) {
                                                var v22 float32 = float32(v14)
                                                var v23 int32 = 42
                                                _ = len(v6)
                                                _ = v22
                                                _ = v23
                                            } else {
                                                var v24 int32 = 1
                                                v25 := GenS1{F0: v24, F1: 705}
                                                var v26 float32 = (0.1 - float32(-2.25))
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                            }
                                        default:
                                            v1 += 426
                                        }
                                        v27 := GenS1{F0: 42, F1: 1}
                                        _ = v27
                                    } else {
                                        switch ("\n\t" + "\n\t") {
                                        case "foo bar":
                                            switch "a" {
                                            case "":
                                                v1 -= v1
                                            case "a":
                                                if true {
                                                    var v28 int16 = v3
                                                    v29 := GenS2{F0: 1, F1: 284, F2: 654}
                                                    fmt.Println(v5, v3)
                                                    switch "" {
                                                    case "\n\t":
                                                        v30 := GenS3{F0: 42, F1: 42, F2: 30}
                                                        _ = v30
                                                    case "foo bar":
                                                        if true {
                                                            v28 = v28
                                                        }
                                                    case "hello":
                                                        var v31 bool = true
                                                        _ = v31
                                                    }
                                                    _ = v28
                                                    _ = v29
                                                }
                                            case "hello":
                                                v32 := GenS3{F0: -1, F1: 0, F2: 222}
                                                _ = v32
                                            }
                                        case "a":
                                            v1 -= 0
                                        }
                                    }
                                    v1 = (v1 + (v1 << ((1) & 63)))
                                } else {
                                    var v33 uint8 = 179
                                    var v34 int32 = 937
                                    _ = v33
                                    _ = v34
                                }
                                var v35 uint8 = 128
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    v35 -= v35
                                }
                                _ = fi12
                                _ = v13
                                _ = v14
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 17; fi36++ {
                                switch (("\n\t" + "\n\t") + ("hello" + "foo bar")) {
                                case "foo bar":
                                    v37 := GenS1{F0: 2147483647, F1: 65535}
                                    _ = v37
                                case "hello":
                                    v38 := []int{}
                                    _ = v38
                                case "":
                                    v39 := v2["a"]
                                    _ = v39
                                default:
                                    v5 -= ((v5 + 0.1) * 1.5)
                                }
                                var v40 int8 = 1
                                var v41 int64 = -265
                                _ = fi36
                                _ = v40
                                _ = v41
                            }
                        } else {
                            if true {
                                for fi42 := 0; fi42 < 16; fi42++ {
                                    v1 += (^(v1))
                                    _ = fi42
                                }
                            } else {
                                var v43 string = "\n\t"
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    if true {
                                        v3 -= v3
                                        var v45 float64 = ((float64((-2.25 + 0.1)) + v5) + (v5 * ((v5 * v5) + v5)))
                                        if true {
                                            v5 = (v5 - float64(-2.25))
                                            if true {
                                                if true {
                                                    v46 := GenS4{F0: 18446744073709551615, F1: 128, F2: 255}
                                                    if false {
                                                        switch -2147483648 {
                                                        case 912:
                                                            v3 = 676
                                                            v47 := []int{(fi7 | (fi44 - (^(fi7)))), (fi44 / ((-49) | 1)), fi7}
                                                            _ = v47
                                                        case 128:
                                                            if false {
                                                                v48 := []int{}
                                                                var v49 int16 = 32767
                                                                v50 := GenS5{F0: (v1 >> ((65535) & 63)), F1: (v49 ^ (v49 / ((v3) | 1)))}
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                v51 := GenS5{F0: (^(v1)), F1: (^(v3))}
                                                                var v52 uint64 = 4294967295
                                                                if true {
                                                                    var v53 int64 = -1
                                                                    delete(v2, v43)
                                                                    switch fi44 {
                                                                    case 0:
                                                                        v1 = 420
                                                                    }
                                                                    _ = v53
                                                                } else {
                                                                    v54 := GenS4{F0: uint(v1), F1: 208, F2: 0}
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        if (2 != 65535) {
                                                                            v55 := GenS1{F0: 2147483647, F1: 0}
                                                                            var v56 uint16 = 32767
                                                                            _ = v55
                                                                            _ = v56
                                                                        } else {
                                                                            var v57 int64 = 9223372036854775807
                                                                            if true {
                                                                                v58 := GenS5{F0: v1, F1: -1}
                                                                                _ = v58
                                                                            } else {
                                                                                switch (fi44 + fi44) {
                                                                                case 32767:
                                                                                    var v59 string = "\n\t"
                                                                                    fmt.Println(v3, v1)
                                                                                    v60 := []int{}
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            var v61 int32 = 962
                                                                            _ = v57
                                                                            _ = v61
                                                                        }
                                                                    default:
                                                                        v1 = v52
                                                                    }
                                                                    if false {
                                                                        if (!(127 == 0)) {
                                                                            v62 := GenS1{F0: 1, F1: 2}
                                                                            v45 -= -2.25
                                                                            fmt.Println(v52)
                                                                            var v63 uint = uint(v52)
                                                                            _ = v62
                                                                            _ = v63
                                                                        } else {
                                                                            var v64 uint = 42
                                                                            v1 = (v1 & (v52 / (((^(v1))) | 1)))
                                                                            v65 := GenS6{F0: (v43 + ""), F1: -9223372036854775808, F2: 9223372036854775807}
                                                                            var v66 int8 = 100
                                                                            fmt.Println(v5, v64, v45)
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    _ = v54
                                                                }
                                                                if (42 > -31) {
                                                                    var v67 float32 = 1e10
                                                                    _ = v67
                                                                }
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                        v68 := GenS3{F0: 42, F1: 693, F2: 253}
                                                        _ = v68
                                                    }
                                                    _ = v46
                                                }
                                                v5 -= v5
                                                v43 = v43
                                                if false {
                                                    if (!false) {
                                                        switch v43 {
                                                        case "hello":
                                                            var v69 float32 = 0.1
                                                            _ = v69
                                                        case "":
                                                            if (!true) {
                                                                fmt.Println(fi44, v3)
                                                            } else {
                                                                switch "hello" {
                                                                case "":
                                                                    v3 -= (v3 + v3)
                                                                default:
                                                                    var v70 uint8 = 1
                                                                    _ = v70
                                                                }
                                                                v3 -= v3
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = v45
                                    } else {
                                        if false {
                                            v1 = 93
                                            var v71 uint32 = 2
                                            var v72 string = "a"
                                            _ = v71
                                            _ = v72
                                        }
                                        var v73 int8 = -79
                                        _ = v73
                                    }
                                    fmt.Println(v43)
                                    _ = fi44
                                }
                                _ = v43
                            }
                        }
                        v3 = 720
                    }
                    var v74 int32 = 2147483647
                    if (!(((^(v3)) != (v3 << ((255) & 63))) || false)) {
                        fmt.Println(v74)
                        if true {
                            for fi75 := 0; fi75 < 1; fi75++ {
                                var v76 int32 = v74
                                if true {
                                    if true {
                                        fmt.Println(fi75, v76)
                                    } else {
                                        fmt.Println(v76)
                                        v77 := GenS5{F0: 4294967295, F1: (v3 + 0)}
                                        _ = v77
                                    }
                                    var v78 uint16 = 42
                                    _ = v78
                                }
                                _ = len(v6)
                                _ = fi75
                                _ = v76
                            }
                        } else {
                            if (!true) {
                                if false {
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v5)
                                    case "\n\t":
                                        var v79 uint = 255
                                        _ = v79
                                    case "a":
                                        fmt.Println(v3)
                                    default:
                                        v80 := GenS3{F0: 42, F1: 2147483648, F2: 128}
                                        _ = v80
                                    }
                                    for fi81 := 0; fi81 < 10; fi81++ {
                                        v82 := GenS7{F0: v74, F1: true}
                                        v83 := v2[(("foo bar" + ("\n\t" + (("a" + "a") + ("\n\t" + "")))) + "hello")]
                                        v84 := GenS4{F0: 39, F1: 127, F2: 255}
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                } else {
                                    for fi85 := 0; fi85 < 15; fi85++ {
                                        var v86 float32 = (1.5 / 1.0)
                                        _ = fi85
                                        _ = v86
                                    }
                                }
                                var v87 uint64 = v1
                                var v88 int32 = v74
                                _ = v87
                                _ = v88
                            }
                            var v89 int32 = 42
                            switch ("hello" + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                for fi90 := 0; fi90 < 17; fi90++ {
                                    var v91 uint8 = 198
                                    v5 += ((-2.25 - float64((float32((((0.0 - ((float64((0.1 * 0.0)) + -2.25) + 0.0)) * float64(0.0)) * v5)) * 1.5))) / v5)
                                    _ = fi90
                                    _ = v91
                                }
                            case "a":
                                switch (fi7 >> ((894) & 63)) {
                                case 127:
                                    v92 := GenS8{F0: 4294967295, F1: 2}
                                    _ = v92
                                default:
                                    var v93 int32 = -530
                                    _ = v93
                                }
                            }
                            _ = v89
                        }
                        var v94 uint32 = 2147483647
                        _ = v94
                    }
                    _ = v74
                }
            }
            _ = v11
        } else {
            v3 = v3
        }
        _ = fi7
    }
    v95 := GenS1{F0: -2147483648, F1: 226}
    v96 := []int{42}
    fmt.Println(v1, v3)
    var v97 float64 = v5
    v98 := GenS1{F0: 1, F1: 2}
    var v99 bool = ((int32(v1) != 828) && false)
    v4[""] = true
    if len(v96) > 0 {
        v96[0] = -627
    }
    var v100 int = -9223372036854775808
    var v101 int64 = -1
    if v99 {
        var v102 uint16 = 42
        delete(v4, "foo bar")
        _ = v102
    } else {
        var v103 bool = v99
        var v104 uint16 = 255
        _ = v103
        _ = v104
    }
    switch (v100 - 42) {
    case 2147483647:
        var v105 uint8 = 127
        _ = v105
    case 136:
        for fi106 := 0; fi106 < 9; fi106++ {
            for fi107 := 0; fi107 < 10; fi107++ {
                v108 := GenS9{F0: 2, F1: v101, F2: 4294967295}
                v99 = true
                _ = fi107
                _ = v108
            }
            v109 := GenS1{F0: 2147483647, F1: 32767}
            _ = fi106
            _ = v109
        }
    }
    switch (v100 << ((2) & 63)) {
    case 9223372036854775807:
        v96 = append(v96, v100)
    case 42:
        for fi110 := 0; fi110 < 3; fi110++ {
            v101 += (^(v101))
            _ = fi110
        }
    default:
        for fi111 := 0; fi111 < 16; fi111++ {
            v1 += v1
            var v112 int32 = 2147483647
            _ = fi111
            _ = v112
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v96:", len(v96), v96)
    fmt.Println("v97:", v97)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
