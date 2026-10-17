package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int32; F1 bool; F2 int64 }
type GenS6 struct { F0 uint8; F1 string; F2 int64 }
type GenS7 struct { F0 int16; F1 int64; F2 int }
type GenS8 struct { F0 int64; F1 bool; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS12 struct { F0 int32; F1 int8; F2 uint }

func main() {
    var v1 bool = true
    var v2 int = 9223372036854775807
    var v3 float32 = 1.5
    var v4 int32 = 1
    var v5 float32 = (float32((3.14159 / 3.14159)) / (v3 * (3.14159 * (float32(3.14159) * v3))))
    v6 := make(map[string]string)
    var v7 int8 = 127
    v1 = v1
    v8 := []int{v2}
    if (42 == 2) {
        switch 42 {
        case 9223372036854775807:
            fmt.Println(v3)
        }
    }
    if v1 {
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 11; fi10++ {
                v2 = v2
                _ = len(v8)
                _ = fi10
            }
            if ((-2147483648 < 42) || true) {
                if (v4 >= 2147483647) {
                    fmt.Println(v5)
                    for fi11 := 0; fi11 < 4; fi11++ {
                        if (42 >= 1) {
                            var v12 float64 = float64(0.0)
                            _ = v12
                        } else {
                            v13 := GenS1{F0: (v4 & v4), F1: 65535}
                            _ = v13
                        }
                        var v14 int32 = (v4 - v4)
                        if (65535 <= 32767) {
                            fmt.Println(fi11)
                            v4 -= v14
                            var v15 uint16 = 2
                            _ = v15
                        } else {
                            _ = len(v6)
                            var v16 int16 = 32767
                            _ = v16
                        }
                        _ = fi11
                        _ = v14
                    }
                    v6["foo bar"] = "a"
                } else {
                    v4 -= 1
                    fmt.Println(v1, v1, v4)
                    if v1 {
                        var v17 int64 = 42
                        _ = v17
                    }
                }
                fmt.Println(v2, v5)
            } else {
                if v1 {
                    v18 := v6["\n\t"]
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 int = (-(fi9))
                        v21 := GenS2{F0: (v4 / (((v4 & v4)) | 1)), F1: -960, F2: (v4 | v4)}
                        var v22 uint8 = 127
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    if (0 == 595) {
                        fmt.Println(v7, fi9, v4)
                        var v23 string = v18
                        if (0 == 2) {
                            if v1 {
                                var v24 int8 = v7
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    v3 = 1e10
                                    fmt.Println(v4, v4, fi9)
                                    v8 = append(v8, (fi25 << ((255) & 63)))
                                    _ = fi25
                                }
                                _ = v24
                            }
                            v2 += fi9
                        }
                        _ = v23
                    } else {
                        var v26 int = v2
                        v27 := GenS2{F0: v4, F1: 0, F2: (v4 | v4)}
                        _ = v26
                        _ = v27
                    }
                    _ = v18
                } else {
                    v7 -= v7
                    var v28 int16 = 32767
                    _ = v28
                }
            }
            _ = fi9
        }
    } else {
        for fi29 := 0; fi29 < 9; fi29++ {
            if ((0.0 <= float64((-2.25 * (((v3 / v5) + (1e10 / (v3 / 0.1))) / (v3 / -2.25))))) || (1 >= 539)) {
                var v30 uint64 = 151
                _ = v30
            }
            fmt.Println(v7)
            _ = fi29
        }
    }
    var v31 int16 = 688
    if false {
        for fi32 := 0; fi32 < 10; fi32++ {
            fmt.Println(v1, v5)
            fmt.Println(fi32)
            _ = fi32
        }
        var v33 float32 = 3.14159
        var v34 float32 = v33
        _ = v33
        _ = v34
    } else {
        fmt.Println(v31, v31)
        v5 = ((1e10 * float32((0.1 / 1.5))) * (v5 / (v3 - -2.25)))
    }
    v35 := GenS3{F0: 2147483647, F1: (v4 | (v4 / ((v4) | 1))), F2: 162}
    fmt.Println(v2)
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint8 = 41
        v38 := GenS1{F0: (^(v4)), F1: 386}
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := GenS4{F0: 1, F1: 0}
    switch v2 {
    case 278:
        if false {
            v31 -= 42
        }
    case 0:
        v7 -= (v7 << ((18446744073709551615) & 63))
    case 128:
        for fi40 := 0; fi40 < 15; fi40++ {
            for fi41 := 0; fi41 < 10; fi41++ {
                var v42 string = "foo bar"
                if (3.14159 >= (0.0 - 1.5)) {
                    var v43 uint8 = 1
                    delete(v6, v42)
                    switch v2 {
                    case 449:
                        v44 := GenS4{F0: 42, F1: 0}
                        _ = v44
                    case 128:
                        v42 = ("" + ((v42 + v42) + v42))
                    case 2147483647:
                        v45 := GenS4{F0: 0, F1: 0}
                        fmt.Println(v4, fi40, v43)
                        if (!(v1 || (v4 < v4))) {
                            v46 := GenS5{F0: 92, F1: (!(!v1)), F2: 627}
                            _ = v46
                        } else {
                            v42 = ""
                        }
                        switch (fi40 % ((514) | 1)) {
                        case 1:
                            switch "" {
                            case "a":
                                var v47 int16 = v31
                                _ = v47
                            default:
                                v42 = "a"
                                if v1 {
                                    switch fi41 {
                                    case 1:
                                        if (v43 <= (v43 << ((4294967295) & 63))) {
                                            switch v42 {
                                            case "hello":
                                                delete(v6, ("\n\t" + "foo bar"))
                                            case "foo bar":
                                                v3 = v5
                                            default:
                                                v4 += v4
                                            }
                                            v42 = "foo bar"
                                        }
                                    default:
                                        v7 += -128
                                    }
                                    if false {
                                        var v48 int = (v2 & (fi40 & (^(v2))))
                                        if (42 != 255) {
                                            var v49 string = ((("a" + v42) + v42) + (v42 + v42))
                                            if ((true || v1) || v1) {
                                                var v50 float32 = v3
                                                _ = len(v8)
                                                var v51 uint64 = 0
                                                _ = v50
                                                _ = v51
                                            }
                                            if true {
                                                v5 -= float32(0.0)
                                            }
                                            v43 += 2
                                            var v52 bool = true
                                            _ = v49
                                            _ = v52
                                        }
                                        _ = v48
                                    } else {
                                        if (v1 && false) {
                                            var v53 uint16 = 65535
                                            _ = v53
                                        }
                                        fmt.Println(fi41, fi40, fi40)
                                    }
                                    if false {
                                        switch (v42 + "\n\t") {
                                        case "":
                                            delete(v6, (v42 + v42))
                                        case "a":
                                            if v1 {
                                                if (v3 != -2.25) {
                                                    v54 := GenS6{F0: (v43 * v43), F1: "", F2: 9223372036854775807}
                                                    v55 := GenS5{F0: 0, F1: (!((v43 | (v43 + (v43 / ((v43) | 1)))) == v43)), F2: 355}
                                                    switch v42 {
                                                    case "\n\t":
                                                        var v56 int16 = (v31 | (v31 ^ (v31 & (v31 | -625))))
                                                        _ = v56
                                                    }
                                                    _ = v54
                                                    _ = v55
                                                }
                                            } else {
                                                var v57 uint = 0
                                                _ = v57
                                            }
                                        }
                                        _ = len(v8)
                                    }
                                }
                            }
                        case 32767:
                            var v58 bool = v1
                            _ = v58
                        case 127:
                            if v1 {
                                var v59 float64 = 1e10
                                fmt.Println(v7)
                                _ = v59
                            } else {
                                if v1 {
                                    var v60 int16 = (^(v31))
                                    _ = v60
                                } else {
                                    v43 = (v43 - 127)
                                    var v61 uint8 = (^(v43))
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    switch (fi41 - fi40) {
                                    case 145:
                                        v6[v42] = v42
                                    case 1:
                                        v1 = v1
                                    case 32767:
                                        var v62 float64 = (-2.25 - (-2.25 - 1.5))
                                        _ = v62
                                    default:
                                        fmt.Println(v3)
                                    }
                                    if v1 {
                                        v63 := GenS7{F0: (^(v31)), F1: 42, F2: 128}
                                        v2 += (fi40 * fi40)
                                        _ = v63
                                    } else {
                                        v7 += 40
                                        switch (fi40 / ((int(v4)) | 1)) {
                                        case 313:
                                            if len(v8) > 0 {
                                                v8[0] = (v2 ^ (^(v2)))
                                            }
                                        case 128:
                                            switch ("foo bar" + (v42 + v42)) {
                                            case "\n\t":
                                                v43 = 64
                                            case "a":
                                                v6[("\n\t" + v42)] = v42
                                            case "foo bar":
                                                if ((!false) && false) {
                                                    v2 = (fi41 | 9223372036854775807)
                                                    if ((!((-(v7)) > 42)) || ((0 >= v43) && v1)) {
                                                        if true {
                                                            var v64 uint16 = 255
                                                            fmt.Println(v4, v64, v43)
                                                            v65 := GenS7{F0: 949, F1: -1, F2: -482}
                                                            var v66 uint32 = 458
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                        } else {
                                                            switch "" {
                                                            case "foo bar":
                                                                v31 -= (v31 * 1)
                                                            default:
                                                                var v67 uint16 = 275
                                                                _ = v67
                                                            }
                                                            v61 -= 2
                                                            if v1 {
                                                                v8 = append(v8, (fi40 ^ 32767))
                                                                v7 -= (^(v7))
                                                            } else {
                                                                var v68 int8 = v7
                                                                if v1 {
                                                                    switch (("" + v42) + v42) {
                                                                    case "hello":
                                                                        v4 = -265
                                                                    case "a":
                                                                        fmt.Println(v68, v7, v43)
                                                                        var v69 uint8 = v43
                                                                        var v70 bool = (!true)
                                                                        v71 := GenS8{F0: 896, F1: v1, F2: -9223372036854775808}
                                                                        _ = v69
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                }
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                                _ = v68
                                                            }
                                                        }
                                                        var v72 uint = uint(v61)
                                                        switch (^(v2)) {
                                                        case 128:
                                                            switch fi40 {
                                                            case 128:
                                                                _ = len(v6)
                                                                v73 := GenS9{F0: v42}
                                                                _ = v73
                                                            case -780:
                                                                if (42 < v4) {
                                                                    v2 = fi41
                                                                    var v74 int8 = v7
                                                                    v31 += -591
                                                                    _ = v74
                                                                } else {
                                                                    v2 -= v2
                                                                    var v75 uint32 = 2
                                                                    switch fi41 {
                                                                    case -364:
                                                                        switch v2 {
                                                                        case 42:
                                                                            if ((^(v75)) >= v75) {
                                                                                fmt.Println(v3, fi40)
                                                                            } else {
                                                                                v72 = (v72 & 42)
                                                                                v76 := GenS6{F0: (v61 + 107), F1: (v42 + "\n\t"), F2: -9223372036854775808}
                                                                                var v77 bool = ((float32(1e10) * v3) < 3.14159)
                                                                                v78 := v6[(v42 + v42)]
                                                                                var v79 int64 = -825
                                                                                _ = v76
                                                                                _ = v77
                                                                                _ = v78
                                                                                _ = v79
                                                                            }
                                                                        case 0:
                                                                            if true {
                                                                                v80 := GenS10{F0: -888, F1: 2147483647}
                                                                                v72 -= 42
                                                                                switch fi40 {
                                                                                case 32768:
                                                                                    var v81 uint32 = v75
                                                                                    _ = v81
                                                                                case -9223372036854775808:
                                                                                    fmt.Println(v2, v75, v72)
                                                                                }
                                                                                _ = v80
                                                                            } else {
                                                                                var v82 bool = (false && (v1 || v1))
                                                                                v83 := GenS5{F0: (v4 | 1), F1: v82, F2: -1}
                                                                                _ = v82
                                                                                _ = v83
                                                                            }
                                                                        case -1:
                                                                            switch 717 {
                                                                            case 127:
                                                                                var v84 int64 = -9223372036854775808
                                                                                _ = v84
                                                                            case 32767:
                                                                                var v85 int = fi41
                                                                                _ = v85
                                                                            case -9223372036854775808:
                                                                                v86 := GenS11{F0: 4294967295, F1: 65535, F2: (-(v4))}
                                                                                _ = v86
                                                                            }
                                                                        }
                                                                        var v87 int8 = v7
                                                                        if ((!v1) || ((false && true) || (v42 >= (v42 + "a")))) {
                                                                            v61 -= (v61 << ((v72) & 63))
                                                                            v8 = append(v8, v2)
                                                                            _ = len(v6)
                                                                        } else {
                                                                            fmt.Println(v75, v87)
                                                                            var v88 string = ("a" + "a")
                                                                            v8 = append(v8, (v2 >> ((v72) & 63)))
                                                                            _ = v88
                                                                        }
                                                                        _ = v87
                                                                    case 9223372036854775807:
                                                                        var v89 uint8 = (v43 << ((v72) & 63))
                                                                        _ = v89
                                                                    default:
                                                                        var v90 uint8 = (v43 & (^(v43)))
                                                                        _ = v90
                                                                    }
                                                                    _ = v75
                                                                }
                                                            default:
                                                                if v1 {
                                                                    fmt.Println(v31)
                                                                    v7 = v7
                                                                }
                                                            }
                                                        case 127:
                                                            var v91 uint = (v72 | (v72 / (((^(v72))) | 1)))
                                                            _ = v91
                                                        case -203:
                                                            v31 += (v31 * (v31 >> (((v72 | v72)) & 63)))
                                                        }
                                                        _ = v72
                                                    } else {
                                                        v2 -= v2
                                                        var v92 uint64 = 18446744073709551615
                                                        v93 := GenS4{F0: (v92 | 18446744073709551615), F1: 784}
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    switch v42 {
                                                    case "\n\t":
                                                        var v94 uint8 = v61
                                                        _ = v94
                                                    case "foo bar":
                                                        v5 = (v5 - (float32(-2.25) - ((0.0 / ((float32(3.14159) / v5) / v3)) + v3)))
                                                        switch 9223372036854775807 {
                                                        case 127:
                                                            v8 = append(v8, fi41)
                                                        }
                                                        v95 := GenS10{F0: (v4 << ((546) & 63)), F1: (v2 - (-(fi41)))}
                                                        fmt.Println(v43)
                                                        _ = v95
                                                    case "a":
                                                        v42 = "hello"
                                                    default:
                                                        v1 = (!false)
                                                        fmt.Println(v2, v42)
                                                        v96 := GenS12{F0: (v4 & (v4 >> ((18446744073709551615) & 63))), F1: 0, F2: 1}
                                                        _ = v96
                                                    }
                                                    _ = len(v6)
                                                    _ = len(v6)
                                                }
                                            default:
                                                v6[v42] = "foo bar"
                                            }
                                        case 127:
                                            var v97 float64 = (0.0 * 0.1)
                                            _ = v97
                                        }
                                    }
                                    _ = v61
                                }
                                v98 := GenS10{F0: (v4 - v4), F1: fi41}
                                v1 = v1
                                delete(v6, v42)
                                _ = v98
                            }
                        }
                        var v99 bool = (255 < 0)
                        _ = v45
                        _ = v99
                    }
                    _ = v43
                }
                _ = fi41
                _ = v42
            }
            v100 := GenS8{F0: 9223372036854775807, F1: v1, F2: 0}
            _ = fi40
            _ = v100
        }
    default:
        var v101 int32 = v4
        _ = v101
    }
    var v102 float32 = (v3 - (float32((0.0 * 1e10)) + ((3.14159 * (v5 + -2.25)) - 1.5)))
    if true {
        fmt.Println(v7, v7, v2)
        var v103 uint16 = 854
        if v1 {
            fmt.Println(v4)
        }
        _ = v103
    } else {
        v102 -= v5
        v104 := GenS10{F0: (v4 >> ((65535) & 63)), F1: v2}
        _ = v104
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
