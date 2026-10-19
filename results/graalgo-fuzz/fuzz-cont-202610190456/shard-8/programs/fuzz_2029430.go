package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint64; F2 int }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint16; F1 uint64; F2 bool }
type GenS8 struct { F0 string; F1 int32; F2 bool }
type GenS9 struct { F0 int64; F1 int32; F2 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint64 = 542
    var v2 bool = (1.5 >= 1.5)
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    if ((1.5 / 1e10) != (1e10 + 1e10)) {
        var v5 int = -732
        v6 := []int{-1, 42, (v5 & v5)}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 367, F1: 1, F2: ""}
    v1 = 2
    v1 += (v1 >> ((823) & 63))
    var v8 int8 = 0
    v9 := []int{127, 128}
    for fi10 := 0; fi10 < 9; fi10++ {
        _ = len(v9)
        for fi11 := 0; fi11 < 8; fi11++ {
            v1 -= v1
            switch fi11 {
            case 32767:
                var v12 uint16 = 2
                if (-32768 > -343) {
                    var v13 uint8 = 155
                    if true {
                        var v14 float64 = 1.5
                        if v2 {
                            v15 := GenS2{F0: false}
                            if v2 {
                                switch fi11 {
                                case 414:
                                    fmt.Println(v1)
                                default:
                                    var v16 int32 = 0
                                    _ = v16
                                }
                                v9 = append(v9, (fi10 % ((1) | 1)))
                                v17 := GenS2{F0: (("\n\t" + ("a" + ("a" + (("foo bar" + "\n\t") + ("hello" + ("foo bar" + ("hello" + ""))))))) <= (("foo bar" + ("" + "foo bar")) + "hello"))}
                                _ = v17
                            } else {
                                if (42 == 4294967295) {
                                    var v18 int64 = 1
                                    _ = len(v4)
                                    var v19 uint8 = (v13 % ((137) | 1))
                                    switch fi11 {
                                    case 129:
                                        switch "hello" {
                                        case "a":
                                            _ = len(v3)
                                        case "":
                                            var v20 uint64 = v1
                                            _ = v20
                                        default:
                                            v1 = 2
                                        }
                                        var v21 uint = 929
                                        _ = v21
                                    case -9223372036854775808:
                                        if v2 {
                                            if (1.5 < (3.14159 + 3.14159)) {
                                                if (!v2) {
                                                    var v22 int8 = (v8 * (v8 % ((-17) | 1)))
                                                    switch "hello" {
                                                    case "":
                                                        v23 := GenS3{F0: ((true || v2) && false)}
                                                        _ = v23
                                                    case "foo bar":
                                                        v19 -= v19
                                                        var v24 uint8 = (v13 << ((2) & 63))
                                                        var v25 float64 = v14
                                                        if v2 {
                                                            if len(v9) > 0 {
                                                                v9[0] = (fi11 % ((9223372036854775807) | 1))
                                                            }
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                            var v26 uint8 = 128
                                                            if len(v9) > 0 {
                                                                v9[0] = int(v19)
                                                            }
                                                            _ = v26
                                                        }
                                                        _ = v24
                                                        _ = v25
                                                    case "a":
                                                        v27 := GenS3{F0: false}
                                                        _ = v27
                                                    }
                                                    var v28 uint32 = 4294967295
                                                    _ = v22
                                                    _ = v28
                                                } else {
                                                    var v29 string = "hello"
                                                    v1 -= v1
                                                    _ = v29
                                                }
                                                v8 += v8
                                            }
                                            var v30 int8 = (v8 / ((v8) | 1))
                                            var v31 uint = 1
                                            _ = v30
                                            _ = v31
                                        }
                                        switch (fi11 / (((fi10 & int(v18))) | 1)) {
                                        case 83:
                                            fmt.Println(v2, fi11, v14)
                                        default:
                                            if v2 {
                                                if (!((!(!(-9223372036854775808 <= v18))) || true)) {
                                                    v18 = (v18 + int64(v12))
                                                    switch fi11 {
                                                    case 32768:
                                                        switch "hello" {
                                                        case "hello":
                                                            switch fi11 {
                                                            case 128:
                                                                v32 := GenS4{F0: (v13 & v19), F1: (v1 << ((4294967295) & 63)), F2: (fi10 << ((2) & 63))}
                                                                _ = v32
                                                            case -1:
                                                                if (((!v2) && ((v14 + (-2.25 + v14)) <= ((v14 + v14) - v14))) || ((float64(float32(v14)) / (v14 * 1e10)) >= v14)) {
                                                                    if (true && (!true)) {
                                                                        if (v18 < (v18 % ((v18) | 1))) {
                                                                            v19 -= (v19 ^ v19)
                                                                            var v33 int32 = int32(v13)
                                                                            _ = v33
                                                                        } else {
                                                                            v34 := v4[(^(fi10))]
                                                                            _ = v34
                                                                        }
                                                                    } else {
                                                                        var v35 uint32 = 1
                                                                        var v36 float32 = (float32((v14 / (1e10 + v14))) / (-2.25 + 1.5))
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                    if len(v9) > 0 {
                                                                        v9[0] = fi11
                                                                    }
                                                                    v37 := GenS3{F0: (true || (9223372036854775807 < fi11))}
                                                                    if v2 {
                                                                        v38 := GenS5{F0: 2147483648, F1: v8}
                                                                        v2 = (3.14159 >= 3.14159)
                                                                        v39 := GenS6{F0: (^(v18)), F1: 1000}
                                                                        switch (fi10 / (((-(fi10))) | 1)) {
                                                                        case 2147483647:
                                                                            if v2 {
                                                                                switch (fi10 - (-(fi10))) {
                                                                                case 776:
                                                                                    var v40 uint8 = v19
                                                                                    _ = v40
                                                                                case -9223372036854775808:
                                                                                    v1 = (v1 & (v1 & v1))
                                                                                    v1 = v1
                                                                                    var v41 bool = (!v2)
                                                                                    v13 += (v19 & v13)
                                                                                    v4[fi10] = (("" + (("\n\t" + ("foo bar" + "\n\t")) + "a")) + "")
                                                                                    var v42 int8 = v8
                                                                                    fmt.Println(v12, v13)
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                case 128:
                                                                                    var v43 int16 = -32768
                                                                                    _ = v43
                                                                                default:
                                                                                    v14 = (v14 / (v14 * v14))
                                                                                }
                                                                                _ = len(v9)
                                                                            }
                                                                        case 42:
                                                                            var v44 string = "a"
                                                                            _ = v44
                                                                        case 32768:
                                                                            if (47 > 150) {
                                                                                delete(v3, fi10)
                                                                                v8 -= -21
                                                                            }
                                                                        default:
                                                                            v18 += v18
                                                                        }
                                                                        fmt.Println(v8)
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                    if ((2147483648 >= 562) && v2) {
                                                                        fmt.Println(fi11)
                                                                    } else {
                                                                        switch 9223372036854775807 {
                                                                        case -923:
                                                                            v1 = 2
                                                                        }
                                                                        v45 := GenS7{F0: uint16(v19), F1: (v1 % ((v1) | 1)), F2: true}
                                                                        v18 = 1
                                                                        _ = v45
                                                                    }
                                                                    _ = len(v9)
                                                                    _ = v37
                                                                } else {
                                                                    fmt.Println(v13)
                                                                    var v46 int64 = -9223372036854775808
                                                                    var v47 uint16 = v12
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            case 1:
                                                                fmt.Println(v8, v13, fi10)
                                                            default:
                                                                v48 := GenS8{F0: "a", F1: -2147483648, F2: (((522 == 65535) || ((0.1 < float64((1.5 / 1.0))) || (v2 || ((("a" + "hello") < "\n\t") && v2)))) && v2)}
                                                                _ = v48
                                                            }
                                                        case "":
                                                            fmt.Println(v13, v12, v19)
                                                        case "a":
                                                            v19 = (v13 * 128)
                                                        }
                                                    case -9223372036854775808:
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                    case -1:
                                                        var v49 uint = 176
                                                        _ = v49
                                                    default:
                                                        v50 := GenS5{F0: 340, F1: v8}
                                                        _ = v50
                                                    }
                                                    v12 = v12
                                                } else {
                                                    var v51 int64 = (v18 + (v18 / (((v18 | v18)) | 1)))
                                                    if (("foo bar" < "foo bar") && (4294967295 != 42)) {
                                                        v1 = v1
                                                        var v52 bool = v2
                                                        _ = v52
                                                    } else {
                                                        var v53 int32 = -969
                                                        fmt.Println(fi10, v53)
                                                        _ = v53
                                                    }
                                                    fmt.Println(v51, v18)
                                                    _ = v51
                                                }
                                            }
                                        }
                                    case -47:
                                        v9 = append(v9, 42)
                                    default:
                                        var v54 uint64 = (v1 << ((65535) & 63))
                                        _ = v54
                                    }
                                    _ = v18
                                    _ = v19
                                } else {
                                    var v55 float64 = (v14 / v14)
                                    _ = v55
                                }
                            }
                            _ = v15
                        } else {
                            var v56 string = "foo bar"
                            var v57 uint = uint(v13)
                            v9 = append(v9, (fi11 / (((^(fi11))) | 1)))
                            _ = v56
                            _ = v57
                        }
                        _ = v14
                    }
                    _ = v13
                } else {
                    switch (fi10 + (fi11 ^ (fi10 % ((fi11) | 1)))) {
                    case -715:
                        var v58 int16 = 1
                        _ = v58
                    case -2147483648:
                        v59 := GenS8{F0: "\n\t", F1: -767, F2: ((("\n\t" + "\n\t") + "hello") > "")}
                        if v2 {
                            fmt.Println(fi10)
                        } else {
                            v9 = append(v9, fi11)
                            var v60 int32 = -2147483648
                            _ = v60
                        }
                        if v2 {
                            v61 := GenS2{F0: (32767 != -32768)}
                            if (-2147483648 > 1) {
                                switch (fi11 & (fi10 & (^(fi11)))) {
                                case 32767:
                                    v12 = v12
                                case 1:
                                    var v62 int32 = 0
                                    _ = v62
                                }
                                v8 = (v8 ^ (v8 * (v8 | v8)))
                                if (-1 == -2147483648) {
                                    switch (("" + "") + "hello") {
                                    case "hello":
                                        v12 -= (v12 / ((v12) | 1))
                                    case "\n\t":
                                        v63 := GenS4{F0: 127, F1: v1, F2: fi10}
                                        _ = v63
                                    }
                                    v1 = 42
                                    var v64 int8 = v8
                                    _ = v64
                                }
                            } else {
                                v9 = append(v9, (fi10 + fi11))
                                v12 = (^(v12))
                                v1 += (v1 / (((v1 ^ 2)) | 1))
                            }
                            _ = v61
                        } else {
                            if (float64((float32((3.14159 / 1.0)) + 0.1)) >= 1e10) {
                                if (fi10 > 960) {
                                    v1 += v1
                                    var v65 int = -967
                                    v66 := GenS6{F0: 337, F1: 0}
                                    _ = v65
                                    _ = v66
                                }
                                fmt.Println(v8, fi10, v8)
                                var v67 bool = (42 <= 727)
                                _ = v67
                            }
                            var v68 uint64 = v1
                            if (!(v2 && ((v68 - (v68 & v1)) <= (^(v68))))) {
                                switch fi10 {
                                case 1:
                                    var v69 uint64 = v1
                                    _ = v69
                                }
                            } else {
                                var v70 int64 = -9223372036854775808
                                if len(v9) > 0 {
                                    v9[0] = -574
                                }
                                switch "\n\t" {
                                case "":
                                    v68 += (v68 - (v68 & (v68 + v1)))
                                default:
                                    var v71 int32 = -618
                                    _ = v71
                                }
                                v72 := GenS8{F0: (("" + (("\n\t" + (((("hello" + "a") + ("foo bar" + "hello")) + ((((("\n\t" + "hello") + ("\n\t" + "\n\t")) + "") + "a") + ((("" + "a") + "foo bar") + "foo bar"))) + "\n\t")) + "a")) + "\n\t"), F1: int32(v12), F2: ("\n\t" <= ("" + "hello"))}
                                fmt.Println(v8, v2)
                                var v73 float32 = -2.25
                                _ = v70
                                _ = v72
                                _ = v73
                            }
                            _ = v68
                        }
                        if (!((416 > 671) && (0.0 > (0.1 * 1.5)))) {
                            v74 := GenS2{F0: v2}
                            if true {
                                v12 = (v12 * v12)
                                if false {
                                    if ((42 < -32768) && true) {
                                        if (((^(v1)) != (v1 / ((1) | 1))) || false) {
                                            var v75 bool = v2
                                            _ = v75
                                        } else {
                                            var v76 int32 = 1
                                            fmt.Println(v1, fi10, v1)
                                            _ = v76
                                        }
                                        var v77 uint16 = 1
                                        var v78 int64 = 9223372036854775807
                                        _ = v77
                                        _ = v78
                                    } else {
                                        v79 := GenS9{F0: -1, F1: -213, F2: v8}
                                        if false {
                                            var v80 uint8 = 166
                                            v4[-521] = "\n\t"
                                            _ = v80
                                        }
                                        switch fi11 {
                                        case -2147483648:
                                            fmt.Println(fi11)
                                        }
                                        v1 += v1
                                        _ = v79
                                    }
                                } else {
                                    if (224 <= 127) {
                                        v9 = append(v9, fi11)
                                        var v81 uint32 = 2
                                        _ = v81
                                    }
                                    v1 = (v1 / (((v1 & (^(v1)))) | 1))
                                    var v82 float64 = 1.5
                                    _ = v82
                                }
                            }
                            v83 := GenS6{F0: 42, F1: -1}
                            _ = v74
                            _ = v83
                        } else {
                            _ = len(v9)
                            fmt.Println(v12)
                        }
                        var v84 uint64 = 872
                        _ = v59
                        _ = v84
                    }
                }
                v2 = false
                _ = v12
            case -412:
                switch (^(fi10)) {
                case -9223372036854775808:
                    v8 = (^(v8))
                case -2147483648:
                    v2 = v2
                    fmt.Println(v8)
                }
            }
            _ = len(v9)
            _ = fi11
        }
        v85 := GenS6{F0: 9223372036854775807, F1: 32767}
        _ = fi10
        _ = v85
    }
    v2 = v2
    for fi86 := 0; fi86 < 5; fi86++ {
        v87 := GenS9{F0: -563, F1: 2147483647, F2: v8}
        switch fi86 {
        case 344:
            _ = len(v9)
        case -419:
            fmt.Println(fi86)
        default:
            v8 = v8
        }
        switch (fi86 & (^(fi86))) {
        case 2147483647:
            v8 = (v8 / ((v8) | 1))
        }
        _ = fi86
        _ = v87
    }
    v1 -= (v1 ^ 977)
    delete(v3, -2147483648)
    v88 := GenS10{F0: -72}
    var v89 int = 42
    switch v89 {
    case -630:
        var v90 bool = (509 < 300)
        _ = v90
    default:
        switch "" {
        case "foo bar":
            v8 -= v8
        case "a":
            switch -1 {
            case -9223372036854775808:
                for fi91 := 0; fi91 < 2; fi91++ {
                    v92 := GenS11{F0: -278}
                    var v93 int16 = 1
                    _ = fi91
                    _ = v92
                    _ = v93
                }
            case -1:
                fmt.Println(v2)
            case 42:
                switch "hello" {
                case "\n\t":
                    var v94 uint16 = 2
                    _ = v94
                }
            }
        case "":
            switch 4294967295 {
            case 128:
                switch 0 {
                case 0:
                    var v95 float32 = 0.0
                    _ = v95
                case 4294967295:
                    v89 = v89
                case 42:
                    v96 := GenS11{F0: -1}
                    _ = v96
                }
            case 754:
                v97 := GenS12{F0: 127, F1: 42}
                _ = v97
            case 42:
                v98 := GenS13{F0: 1, F1: 0}
                _ = v98
            }
        default:
            v2 = (-1 < 0)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v89:", v89)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v88
    _ = v89
    _ = v9
}
