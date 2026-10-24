package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint8; F2 int }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint; F2 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 bool; F1 string; F2 string }
type GenS10 struct { F0 int; F1 string; F2 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 int32 }
func genFn0() int16 {
    var v4 int = -2147483648
    var v5 uint = 65535
    var v6 float32 = (-2.25 * (0.1 + 0.1))
    v4 = -9223372036854775808
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string) int8 {
    p0 = ((p0 + ("hello" + p0)) + p0)
    var v7 int = 2147483647
    var v8 uint8 = 1
    if (!true) {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int = (^(v7))
            v11 := []int{4294967295, fi9, (v10 * v10), (v10 | fi9)}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := []int{fi12, v7}
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v7
    _ = v8
    return -128
}

func main() {
    var v1 uint64 = 340
    v2 := []int{}
    var v3 uint = 65535
    if false {
        v3 = 2
        var v14 uint64 = 2
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                if ("" < "foo bar") {
                    if true {
                        v16 := GenS1{F0: fi15}
                        v14 = v1
                        v1 = 2
                        _ = v16
                    } else {
                        v1 += 0
                    }
                }
                v17 := GenS2{F0: 65535, F1: 128, F2: (fi15 * int(v1))}
                _ = v17
            } else {
                var v18 uint64 = 0
                var v19 uint64 = v1
                v20 := GenS3{F0: 1, F1: true}
                _ = v18
                _ = v19
                _ = v20
            }
            if true {
                if ("foo bar" < "a") {
                    fmt.Println(v14, fi15, v14)
                } else {
                    if true {
                        var v21 int32 = 2147483647
                        switch (fi15 + (fi15 >> (((v3 << (((v3 * (v3 ^ (v3 % ((18446744073709551615) | 1))))) & 63))) & 63))) {
                        case 2147483647:
                            switch (fi15 + -586) {
                            case 128:
                                var v22 uint = 255
                                _ = v22
                            case 2147483647:
                                switch ("a" + (("foo bar" + (("" + "foo bar") + "foo bar")) + "foo bar")) {
                                case "hello":
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        if false {
                                            if false {
                                                switch "" {
                                                case "hello":
                                                    switch fi23 {
                                                    case 128:
                                                        v3 -= v3
                                                    case -231:
                                                        var v24 float64 = (3.14159 - (-2.25 - (3.14159 + -2.25)))
                                                        _ = v24
                                                    case 440:
                                                        fmt.Println(v1, v21)
                                                    default:
                                                        switch (fi23 / ((fi15) | 1)) {
                                                        case 9223372036854775807:
                                                            if (("hello" != "") && (4294967295 == 312)) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            } else {
                                                                v21 -= -1
                                                                switch "a" {
                                                                case "foo bar":
                                                                    if true {
                                                                        var v25 int64 = -1
                                                                        _ = v25
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(fi15)
                                                                        }
                                                                        var v26 int = fi15
                                                                        _ = v26
                                                                    }
                                                                default:
                                                                    var v27 int32 = -2147483648
                                                                    _ = v27
                                                                }
                                                            }
                                                        case 4294967295:
                                                            var v28 uint8 = 127
                                                            _ = v28
                                                        case -2147483648:
                                                            var v29 string = (("foo bar" + ("hello" + "")) + "foo bar")
                                                            _ = v29
                                                        }
                                                    }
                                                case "":
                                                    if (genFn1("") < -128) {
                                                        if (true && (("hello" + "") != "foo bar")) {
                                                            var v30 int8 = 13
                                                            v31 := GenS4{F0: (v14 ^ v1)}
                                                            v32 := GenS4{F0: 4294967295}
                                                            if false {
                                                                var v33 string = "foo bar"
                                                                var v34 bool = false
                                                                v35 := GenS4{F0: v14}
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            var v36 float64 = 0.0
                                                            _ = v30
                                                            _ = v31
                                                            _ = v32
                                                            _ = v36
                                                        }
                                                        var v37 uint16 = 42
                                                        _ = v37
                                                    } else {
                                                        v21 += (v21 + v21)
                                                        if (4294967295 >= (v1 / (((^(v14))) | 1))) {
                                                            v1 = 18446744073709551615
                                                            if (!(-862 >= 9223372036854775807)) {
                                                                fmt.Println(v21)
                                                            } else {
                                                                if ((v1 + 4294967295) > (v1 * (^(v14)))) {
                                                                    var v38 bool = true
                                                                    v3 += 18446744073709551615
                                                                    _ = v38
                                                                }
                                                                if false {
                                                                    var v39 int8 = -128
                                                                    fmt.Println(v21, v21, v3)
                                                                    v40 := GenS3{F0: 42, F1: false}
                                                                    _ = v39
                                                                    _ = v40
                                                                }
                                                                var v41 int16 = 1000
                                                                _ = v41
                                                            }
                                                            if false {
                                                                v21 -= (v21 + v21)
                                                            }
                                                            var v42 uint8 = 2
                                                            _ = v42
                                                        }
                                                    }
                                                default:
                                                    switch ("foo bar" + "foo bar") {
                                                    case "foo bar":
                                                        v21 = (v21 | (v21 << ((152) & 63)))
                                                    default:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                fmt.Println(fi15)
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v3)
                                                case "foo bar":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                default:
                                                    switch (fi15 + 128) {
                                                    case -2147483648:
                                                        v43 := GenS5{F0: fi15}
                                                        _ = v43
                                                    case 127:
                                                        if len(v2) > 0 {
                                                            v2[0] = fi23
                                                        }
                                                    case 1:
                                                        var v44 float32 = -2.25
                                                        v3 = v3
                                                        if ((!(fi15 != fi15)) || true) {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v44 += (v44 / 3.14159)
                                                            case "":
                                                                v1 -= v14
                                                            }
                                                            var v45 bool = (!false)
                                                            _ = v45
                                                        } else {
                                                            v44 -= 1.5
                                                            if false {
                                                                fmt.Println(v1)
                                                            }
                                                            v1 -= v1
                                                            var v46 uint32 = 2147483647
                                                            var v47 uint16 = 0
                                                            var v48 int64 = 0
                                                            var v49 uint16 = 42
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                        _ = v44
                                                    default:
                                                        if false {
                                                            v3 = v3
                                                            var v50 float32 = 3.14159
                                                            _ = v50
                                                        } else {
                                                            v21 = (v21 ^ -594)
                                                            fmt.Println(v1, v1, fi23)
                                                        }
                                                    }
                                                }
                                                var v51 int8 = 97
                                                switch "a" {
                                                case "foo bar":
                                                    v52 := GenS2{F0: uint16(v14), F1: 0, F2: (fi23 % ((int(v14)) | 1))}
                                                    _ = v52
                                                case "a":
                                                    v53 := GenS6{F0: "\n\t"}
                                                    _ = v53
                                                case "hello":
                                                    if (!true) {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        var v54 int = fi23
                                                        var v55 int32 = (v21 * (v21 + (v21 << ((v3) & 63))))
                                                        v1 += 277
                                                        _ = v54
                                                        _ = v55
                                                    }
                                                default:
                                                    v56 := GenS5{F0: 127}
                                                    _ = v56
                                                }
                                                _ = v51
                                            }
                                            v57 := GenS3{F0: 630, F1: true}
                                            var v58 uint64 = v14
                                            _ = v57
                                            _ = v58
                                        }
                                        v1 += 777
                                        v21 = int32(v14)
                                        _ = fi23
                                    }
                                case "foo bar":
                                    v21 += 42
                                case "":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v59 float64 = float64(float32((((-2.25 * 3.14159) / 1.0) - 1e10)))
                                    _ = v59
                                }
                            case 76:
                                if true {
                                    var v60 float64 = -2.25
                                    var v61 string = "a"
                                    _ = v60
                                    _ = v61
                                } else {
                                    fmt.Println(v14)
                                    v62 := GenS7{F0: 9223372036854775807, F1: v3, F2: 74}
                                    _ = v62
                                }
                            }
                        case 32768:
                            _ = len(v2)
                        case 32767:
                            fmt.Println(v1, v14, fi15)
                        }
                        _ = len(v2)
                        _ = v21
                    } else {
                        if (!(((-2.25 + (1.5 - 1e10)) / 0.1) >= 0.0)) {
                            v14 += (v1 + v14)
                            var v63 int64 = -391
                            _ = v63
                        } else {
                            fmt.Println(v14, v3)
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                            if false {
                                v64 := GenS4{F0: 2}
                                if (!(32767 <= genFn0())) {
                                    switch "a" {
                                    case "foo bar":
                                        v2 = append(v2, (fi15 - (fi15 << (((v3 & v3)) & 63))))
                                    }
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = v64
                            }
                        }
                    }
                    var v65 int8 = -128
                    fmt.Println(v14, v14)
                    _ = v65
                }
            }
            if (1.5 <= (float32((0.0 + 1.5)) / float32(-2.25))) {
                for fi66 := 0; fi66 < 5; fi66++ {
                    v67 := GenS4{F0: 2}
                    var v68 uint64 = 42
                    if false {
                        switch "a" {
                        case "":
                            var v69 uint16 = 623
                            v2 = append(v2, (fi15 / ((fi66) | 1)))
                            if (!(!(0 == 233))) {
                                var v70 string = "foo bar"
                                v71 := GenS8{F0: uint8(v1)}
                                v3 -= 42
                                _ = v70
                                _ = v71
                            } else {
                                v69 = v69
                            }
                            _ = v69
                        case "\n\t":
                            v14 += 2
                        }
                    }
                    var v72 uint = 2
                    v68 += v14
                    _ = fi66
                    _ = v67
                    _ = v68
                    _ = v72
                }
            } else {
                _ = len(v2)
                var v73 int16 = 32767
                for fi74 := 0; fi74 < 20; fi74++ {
                    if true {
                        v3 += (v3 - (v3 % ((v3) | 1)))
                    }
                    _ = fi74
                }
                _ = v73
            }
            _ = fi15
        }
        _ = v14
    } else {
        for fi75 := 0; fi75 < 6; fi75++ {
            for fi76 := 0; fi76 < 17; fi76++ {
                if false {
                    _ = len(v2)
                } else {
                    var v77 uint = 255
                    _ = v77
                }
                v3 -= 18446744073709551615
                if len(v2) > 0 {
                    v2[0] = fi76
                }
                _ = fi76
            }
            fmt.Println(fi75, v1)
            _ = fi75
        }
        var v78 float32 = 1.5
        var v79 string = "a"
        _ = v78
        _ = v79
    }
    if false {
        v80 := GenS9{F0: false, F1: "\n\t", F2: ("foo bar" + ("" + "foo bar"))}
        _ = len(v2)
        for fi81 := 0; fi81 < 16; fi81++ {
            for fi82 := 0; fi82 < 20; fi82++ {
                fmt.Println(fi82, v1, v3)
                _ = fi82
            }
            v83 := GenS10{F0: -1, F1: "foo bar", F2: -1}
            var v84 int8 = 0
            _ = fi81
            _ = v83
            _ = v84
        }
        _ = v80
    }
    for fi85 := 0; fi85 < 19; fi85++ {
        v2 = append(v2, fi85)
        fmt.Println(fi85, v3, v1)
        fmt.Println(v3, fi85, v3)
        _ = fi85
    }
    if (1.5 > 3.14159) {
        v86 := GenS10{F0: 408, F1: "a", F2: 9223372036854775807}
        _ = v86
    }
    switch "a" {
    case "hello":
        v87 := GenS11{F0: (false || (9223372036854775807 >= -716))}
        _ = v87
    case "a":
        v1 -= 4294967295
    default:
        switch 2147483647 {
        case 127:
            switch ((("a" + "") + (("hello" + "a") + (("a" + "hello") + "a"))) + "a") {
            case "hello":
                fmt.Println(v3, v1)
            case "foo bar":
                switch -9223372036854775808 {
                case -465:
                    v1 += 2
                }
            }
        case 562:
            fmt.Println(v1, v1)
        case 32767:
            v3 = 876
        default:
            if false {
                var v88 float32 = -2.25
                fmt.Println(v88, v88, v3)
                _ = v88
            } else {
                for fi89 := 0; fi89 < 14; fi89++ {
                    var v90 float32 = float32(((1e10 * 1e10) * 3.14159))
                    v1 += 865
                    if (fi89 <= (fi89 << ((v3) & 63))) {
                        v90 -= 1.5
                    } else {
                        v3 -= (v3 | v3)
                    }
                    _ = fi89
                    _ = v90
                }
                for fi91 := 0; fi91 < 10; fi91++ {
                    v3 -= v3
                    v1 += 0
                    for fi92 := 0; fi92 < 9; fi92++ {
                        fmt.Println(v3, fi92, fi92)
                        if ((-32768 <= 42) || false) {
                            var v93 uint64 = v1
                            _ = v93
                        } else {
                            var v94 int64 = -9223372036854775808
                            _ = v94
                        }
                        _ = fi92
                    }
                    _ = fi91
                }
            }
        }
    }
    switch ("" + "hello") {
    case "hello":
        for fi95 := 0; fi95 < 16; fi95++ {
            var v96 float64 = -2.25
            v97 := GenS1{F0: fi95}
            var v98 uint8 = 2
            _ = fi95
            _ = v96
            _ = v97
            _ = v98
        }
    case "foo bar":
        for fi99 := 0; fi99 < 3; fi99++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v100 := GenS3{F0: 341, F1: (!(v3 > v3))}
            _ = fi99
            _ = v100
        }
    case "\n\t":
        v101 := GenS12{F0: 0, F1: -2147483648}
        _ = v101
    }
    var v102 int = 531
    v102 = v102
    v1 += (v1 >> ((4294967295) & 63))
    if len(v2) > 0 {
        v2[0] = v102
    }
    v103 := GenS13{F0: 0}
    var v104 uint = v3
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v104:", v104)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v102
    _ = v103
    _ = v104
    _ = v2
    _ = v3
}
