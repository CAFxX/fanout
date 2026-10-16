package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int64 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int64 = 0
    if (-2.25 != 1.5) {
        var v5 uint = 2
        var v6 bool = true
        _ = v5
        _ = v6
    } else {
        var v7 int64 = v4
        _ = v7
    }
    switch "" {
    case "":
        if true {
            delete(v2, 2147483648)
            var v8 float32 = 1.5
            if ((v4 * (v4 * (v4 >> ((255) & 63)))) <= (v4 - v4)) {
                v9 := []int{0, 128}
                v10 := GenS1{F0: -1, F1: 99}
                _ = v10
                _ = v9
            } else {
                var v11 bool = (!(128 >= 2))
                var v12 uint16 = 42
                for fi13 := 0; fi13 < 1; fi13++ {
                    if false {
                        if v11 {
                            var v14 int8 = -1
                            var v15 float64 = 1e10
                            v16 := GenS1{F0: (v14 & v14), F1: v14}
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v18 := []int{fi13, (^(fi17)), fi13, fi13}
                                v12 -= v12
                                v19 := GenS1{F0: -18, F1: 127}
                                if (!true) {
                                    var v20 int = 1
                                    v21 := GenS2{F0: false, F1: 18446744073709551615}
                                    _ = v20
                                    _ = v21
                                }
                                v11 = ((v11 && ((float64(v8) * ((0.1 / 1e10) / float64(v8))) >= (0.0 + (0.0 * 0.0)))) || (255 > 0))
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            v22 := GenS1{F0: -128, F1: 0}
                            _ = v22
                        }
                    }
                    switch fi13 {
                    case -542:
                        var v23 int8 = -15
                        _ = v23
                    case 32768:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 uint16 = 328
                            var v26 uint32 = 65535
                            v27 := GenS1{F0: -15, F1: 124}
                            _ = fi24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case 128:
                        var v28 int32 = 42
                        _ = v28
                    }
                    switch "hello" {
                    case "hello":
                        var v29 int32 = -228
                        _ = v29
                    case "":
                        switch -382 {
                        case -9223372036854775808:
                            v11 = v11
                        }
                    case "a":
                        v30 := GenS1{F0: 1, F1: -1}
                        _ = v30
                    }
                    _ = fi13
                }
                _ = v11
                _ = v12
            }
            _ = v8
        } else {
            var v31 float64 = 0.0
            fmt.Println(v4, v31)
            _ = v31
        }
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v32 uint16 = 255
            _ = v32
        }
    case "\n\t":
        if ((false || ((18446744073709551615 <= 4294967295) && true)) || (!true)) {
            v33 := GenS2{F0: (!(false || true)), F1: 42}
            _ = v33
        } else {
            var v34 int16 = -864
            _ = v34
        }
    }
    var v35 int8 = -1
    v4 = v4
    var v36 uint16 = 42
    if (true && true) {
        if true {
            for fi37 := 0; fi37 < 2; fi37++ {
                v4 = (v4 & 1)
                var v38 float32 = 1.5
                _ = fi37
                _ = v38
            }
            fmt.Println(v4, v4, v35)
        } else {
            var v39 int64 = v4
            _ = v39
        }
    }
    if (4294967295 < 18446744073709551615) {
        _ = len(v1)
    } else {
        for fi40 := 0; fi40 < 16; fi40++ {
            v36 += v36
            v35 -= v35
            _ = fi40
        }
        v41 := GenS1{F0: v35, F1: v35}
        switch ("foo bar" + "\n\t") {
        case "hello":
            switch "a" {
            case "hello":
                var v42 uint16 = (v36 | 1)
                _ = v42
            }
        case "":
            v43 := GenS2{F0: true, F1: 1}
            _ = v43
        case "\n\t":
            for fi44 := 0; fi44 < 20; fi44++ {
                for fi45 := 0; fi45 < 3; fi45++ {
                    fmt.Println(v36, v36, v4)
                    _ = fi45
                }
                _ = fi44
            }
        }
        _ = v41
    }
    var v46 int8 = v35
    v36 = 2
    var v47 string = ("\n\t" + "")
    if (!(-2147483648 == -990)) {
        v48 := GenS3{F0: 4294967295, F1: 2, F2: -350}
        fmt.Println(v36, v47, v47)
        _ = v48
    }
    v49 := GenS3{F0: 4294967295, F1: 830, F2: v4}
    v47 = "foo bar"
    switch v47 {
    case "hello":
        v50 := GenS1{F0: (v46 << ((4294967295) & 63)), F1: -39}
        _ = v50
    }
    v46 = v46
    var v51 float64 = (1e10 / 1.5)
    switch 889 {
    case 575:
        if (!(0 == -2147483648)) {
            var v52 bool = false
            v53 := v2[2]
            switch 32768 {
            case -959:
                if v53 {
                    fmt.Println(v4, v4, v51)
                    switch ("hello" + "hello") {
                    case "\n\t":
                        var v54 bool = (!(127 == (v46 | (v46 << ((392) & 63)))))
                        _ = v54
                    case "":
                        for fi55 := 0; fi55 < 15; fi55++ {
                            var v56 int = 32768
                            _ = fi55
                            _ = v56
                        }
                    default:
                        if (3 > (^(v36))) {
                            if (v52 || v52) {
                                v57 := []int{9223372036854775807, 2147483647, 804}
                                for fi58 := 0; fi58 < 17; fi58++ {
                                    var v59 int32 = -1
                                    for fi60 := 0; fi60 < 14; fi60++ {
                                        v36 -= v36
                                        switch (fi60 * fi60) {
                                        case -9:
                                            v36 -= 0
                                        case 4294967295:
                                            if v53 {
                                                var v61 uint16 = 0
                                                fmt.Println(v52, fi58)
                                                _ = v61
                                            }
                                        case 2147483647:
                                            v62 := GenS3{F0: 267, F1: 42, F2: (v4 | (v4 + (v4 & 9223372036854775807)))}
                                            _ = v62
                                        }
                                        var v63 float64 = float64(1e10)
                                        _ = fi60
                                        _ = v63
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                                _ = v57
                            }
                            v64 := GenS1{F0: v35, F1: 100}
                            switch v47 {
                            case "":
                                v35 -= (v46 - (v46 >> ((18446744073709551615) & 63)))
                            case "a":
                                v65 := v3[0]
                                _ = v65
                            }
                            _ = v64
                        } else {
                            v66 := GenS2{F0: v52, F1: 255}
                            _ = v66
                        }
                    }
                    for fi67 := 0; fi67 < 19; fi67++ {
                        var v68 int8 = v46
                        v68 -= v68
                        switch v47 {
                        case "a":
                            v4 = (v4 | (v4 & v4))
                        case "":
                            for fi69 := 0; fi69 < 11; fi69++ {
                                var v70 int8 = v68
                                _ = fi69
                                _ = v70
                            }
                        case "foo bar":
                            v51 = v51
                        default:
                            var v71 float64 = v51
                            _ = v71
                        }
                        _ = fi67
                        _ = v68
                    }
                }
            case 9223372036854775807:
                var v72 bool = false
                _ = v72
            }
            _ = v52
            _ = v53
        } else {
            v36 += (v36 % (((^(v36))) | 1))
            v1[(v47 + ("foo bar" + v47))] = 42
            var v73 int8 = -128
            _ = v73
        }
    case 294:
        v47 = v47
    }
    if false {
        if (-636 > 32767) {
            switch v47 {
            case "foo bar":
                v47 = (v47 + ("a" + v47))
            }
            fmt.Println(v35)
            for fi74 := 0; fi74 < 17; fi74++ {
                v75 := GenS2{F0: true, F1: 255}
                _ = fi74
                _ = v75
            }
        } else {
            if (false || true) {
                for fi76 := 0; fi76 < 5; fi76++ {
                    for fi77 := 0; fi77 < 10; fi77++ {
                        v46 += v46
                        var v78 uint = 255
                        switch (-(fi77)) {
                        case -2147483648:
                            if ((fi77 - fi77) == fi77) {
                                v51 += 0.1
                                if (true && (true && ((!false) && (v51 >= (-2.25 + v51))))) {
                                    v79 := []int{(fi77 >> (((v78 & (v78 >> (((v78 ^ 2)) & 63)))) & 63))}
                                    var v80 uint8 = 1
                                    _ = v79
                                    _ = v80
                                }
                            } else {
                                fmt.Println(v51)
                                v81 := []int{fi76, (fi77 << (((v78 - v78)) & 63))}
                                _ = v81
                            }
                        case 9223372036854775807:
                            v82 := GenS1{F0: 100, F1: (v46 & -128)}
                            var v83 float32 = 3.14159
                            v46 = (v46 + (v35 - (^(v46))))
                            _ = v82
                            _ = v83
                        default:
                            switch ((v47 + v47) + (v47 + ("hello" + v47))) {
                            case "":
                                var v84 float32 = float32(float64(3.14159))
                                _ = v84
                            case "a":
                                if (v4 != v4) {
                                    var v85 int16 = 790
                                    _ = v85
                                }
                            default:
                                v86 := []int{(fi76 / ((-468) | 1)), fi76}
                                _ = v86
                            }
                        }
                        _ = fi77
                        _ = v78
                    }
                    _ = fi76
                }
                v87 := []int{-9223372036854775808, -2147483648, -2147483648}
                fmt.Println(v47, v46)
                _ = v87
            } else {
                v2[2147483648] = (18446744073709551615 >= 42)
                var v88 uint = 255
                if false {
                    v4 = (-(v4))
                    var v89 int = 2147483647
                    for fi90 := 0; fi90 < 2; fi90++ {
                        fmt.Println(v88, v88)
                        _ = fi90
                    }
                    _ = v89
                }
                _ = v88
            }
            _ = len(v3)
            for fi91 := 0; fi91 < 3; fi91++ {
                var v92 float32 = 1.5
                var v93 int8 = -84
                _ = fi91
                _ = v92
                _ = v93
            }
        }
    } else {
        if (((((v51 - v51) - (0.1 * 1e10)) * ((v51 - ((1e10 * v51) + -2.25)) + (float64(1e10) / v51))) != (v51 - 3.14159)) || false) {
            v94 := GenS4{F0: 444}
            v95 := []int{-1, 0}
            v96 := GenS2{F0: true, F1: 4294967295}
            _ = v94
            _ = v95
            _ = v96
        }
        v97 := GenS2{F0: (1 < 4294967295), F1: 0}
        fmt.Println(v51)
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v51
}
