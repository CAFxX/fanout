package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int64; F1 int8; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint32 }
type GenS9 struct { F0 string; F1 uint64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 bool; F2 uint }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 uint64 }

func main() {
    var v1 float32 = (float32((0.0 - (0.1 + ((1e10 * 1e10) * 0.0)))) * (1.5 - (0.0 / float32(float64(1.5)))))
    var v2 bool = true
    v3 := []int{194, 127, -1, 32767}
    var v4 float64 = float64(3.14159)
    v5 := []int{32767, 346}
    v6 := []int{0, 255, 128}
    var v7 int8 = -1
    v8 := GenS1{F0: 2147483648, F1: -32768}
    if (v2 || v2) {
        if (127 == v7) {
            v7 = (v7 ^ (v7 % ((v7) | 1)))
        } else {
            v9 := GenS1{F0: 0, F1: -1}
            if (!true) {
                v10 := GenS1{F0: 0, F1: 625}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if v2 {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v1 = 0.0
                    v2 = ((!v2) && ((!(((v1 - 0.0) >= v1) || (!((v2 && (((2147483647 > -9223372036854775808) || ((((v1 / 1.5) * float32((0.0 - v4))) + 3.14159) != (v1 + (v1 * 1.5)))) && false)) || (!("\n\t" == "\n\t")))))) && (!false)))
                }
                _ = v10
            }
            _ = v9
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v1 += v1
            v1 = (v1 / float32(((-2.25 + v4) - 3.14159)))
            _ = fi11
        }
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            if v2 {
                switch ("a" + ("foo bar" + "")) {
                case "a":
                    switch (-(fi12)) {
                    case -2147483648:
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        } else {
                            v1 += float32((1e10 / ((((-2.25 * float64(0.1)) / float64(v1)) / v4) / v4)))
                            var v13 uint16 = 879
                            _ = v13
                        }
                    case 733:
                        var v14 float64 = (float64(float32(-2.25)) + 3.14159)
                        _ = v14
                    case 1:
                        var v15 int32 = -220
                        _ = v15
                    }
                case "":
                    v16 := GenS1{F0: 2147483648, F1: -32768}
                    _ = v16
                case "hello":
                    v17 := GenS2{F0: 255, F1: -1}
                    _ = v17
                default:
                    switch 42 {
                    case 2147483647:
                        _ = len(v3)
                    case -1:
                        _ = len(v6)
                    default:
                        var v18 uint32 = 4294967295
                        _ = v18
                    }
                }
                fmt.Println(v2, v7, v2)
                for fi19 := 0; fi19 < 6; fi19++ {
                    if (2 != 65535) {
                        if (1 <= 255) {
                            v4 = (v4 / float64(((float32(float64(v1)) + (float32((((v4 + 1.5) / (v4 - float64(float32(0.0)))) * -2.25)) / 1.0)) + v1)))
                        }
                        var v20 bool = v2
                        _ = len(v5)
                        switch "foo bar" {
                        case "":
                            v21 := GenS3{F0: -601, F1: 114}
                            _ = v21
                        }
                        v6 = append(v6, -985)
                        if v20 {
                            switch (fi19 % (((fi12 << ((1) & 63))) | 1)) {
                            case 376:
                                switch (fi12 ^ 128) {
                                case 775:
                                    switch fi19 {
                                    case -913:
                                        switch 128 {
                                        case 4294967295:
                                            switch "" {
                                            case "a":
                                                v20 = v2
                                            case "\n\t":
                                                if v20 {
                                                    v20 = (v4 >= (1.5 / ((v4 * v4) * v4)))
                                                } else {
                                                    if v2 {
                                                        v22 := GenS4{F0: 1, F1: -89, F2: 42}
                                                        var v23 int = fi12
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v22
                                                        _ = v23
                                                    } else {
                                                        fmt.Println(v4, fi12, v20)
                                                        fmt.Println(v20, v4)
                                                        v7 += (v7 << ((18446744073709551615) & 63))
                                                    }
                                                    v24 := GenS5{F0: true}
                                                    _ = v24
                                                }
                                                if v20 {
                                                    var v25 float64 = ((v4 / 1.0) - v4)
                                                    var v26 uint8 = 2
                                                    _ = v25
                                                    _ = v26
                                                }
                                            }
                                        case 9223372036854775807:
                                            var v27 int = fi19
                                            _ = v27
                                        case 32768:
                                            var v28 float32 = (float32(v4) * (v1 * 1e10))
                                            _ = v28
                                        default:
                                            v7 -= (^(v7))
                                        }
                                    default:
                                        v29 := GenS6{F0: 255, F1: 32767}
                                        var v30 uint64 = 4294967295
                                        fmt.Println(v7, v4, fi19)
                                        _ = v29
                                        _ = v30
                                    }
                                case 34:
                                    v1 += v1
                                case 128:
                                    v2 = (v7 > v7)
                                }
                            case 0:
                                switch (fi12 - fi19) {
                                case 0:
                                    v31 := GenS5{F0: (29 >= 968)}
                                    _ = v31
                                case 32768:
                                    var v32 int64 = -9223372036854775808
                                    _ = v32
                                case 32767:
                                    switch 4294967295 {
                                    case 2147483647:
                                        v20 = true
                                    case -2147483648:
                                        _ = len(v6)
                                    case 652:
                                        v2 = (42 <= 399)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    default:
                                        var v33 uint8 = 157
                                        _ = v33
                                    }
                                }
                            }
                            v6 = append(v6, 9223372036854775807)
                            var v34 int16 = 55
                            _ = v34
                        }
                        _ = v20
                    } else {
                        if ("hello" <= ("\n\t" + (("hello" + "a") + ("a" + "")))) {
                            v35 := GenS7{F0: v7}
                            _ = v35
                        } else {
                            var v36 uint16 = 65535
                            if true {
                                v4 -= (1e10 + (v4 * v4))
                            } else {
                                var v37 uint16 = v36
                                if (true || (!((v7 & (v7 ^ (v7 >> ((2) & 63)))) < (v7 * (^(v7)))))) {
                                    if v2 {
                                        _ = len(v5)
                                        var v38 uint64 = 4294967295
                                        _ = v38
                                    } else {
                                        if (false && (555 != -326)) {
                                            if len(v3) > 0 {
                                                v3[0] = (fi12 & -2147483648)
                                            }
                                            if false {
                                                var v39 float64 = v4
                                                var v40 bool = (!v2)
                                                _ = v39
                                                _ = v40
                                            }
                                            fmt.Println(v37)
                                        }
                                        if false {
                                            var v41 uint16 = (^(v36))
                                            switch (fi19 << ((65535) & 63)) {
                                            case -1:
                                                if (!true) {
                                                    if ((v2 && (!(!v2))) && (!true)) {
                                                        fmt.Println(v1, fi19)
                                                        if (((42 == 255) && (42 > 42)) || v2) {
                                                            if ("hello" <= ("hello" + "")) {
                                                                v3 = append(v3, fi19)
                                                            } else {
                                                                fmt.Println(fi12)
                                                                v42 := GenS7{F0: 2}
                                                                _ = v42
                                                            }
                                                            v5 = append(v5, fi12)
                                                        }
                                                        if v2 {
                                                            v36 = v36
                                                            if false {
                                                                switch ("foo bar" + (("" + "\n\t") + "")) {
                                                                case "hello":
                                                                    switch ("\n\t" + "foo bar") {
                                                                    case "\n\t":
                                                                        v7 -= v7
                                                                    case "a":
                                                                        v43 := GenS4{F0: -9223372036854775808, F1: 1, F2: -1}
                                                                        fmt.Println(v41)
                                                                        _ = v43
                                                                    }
                                                                }
                                                                var v44 int16 = 42
                                                                switch (fi12 / ((fi12) | 1)) {
                                                                case 0:
                                                                    switch ((("hello" + "") + "") + "a") {
                                                                    case "foo bar":
                                                                        v44 = (v44 + v44)
                                                                    default:
                                                                        v2 = v2
                                                                    }
                                                                default:
                                                                    fmt.Println(v44, v2, v37)
                                                                }
                                                                v45 := GenS5{F0: true}
                                                                _ = len(v3)
                                                                v6 = append(v6, (fi12 & fi19))
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                _ = len(v5)
                                                            }
                                                            switch "a" {
                                                            case "":
                                                                if ((1.5 + 1.5) != v4) {
                                                                    v46 := GenS8{F0: v7, F1: 251}
                                                                    _ = v46
                                                                } else {
                                                                    v36 = v36
                                                                    var v47 uint32 = 123
                                                                    fmt.Println(v37)
                                                                    _ = v47
                                                                }
                                                            case "\n\t":
                                                                v7 += (v7 % ((v7) | 1))
                                                            default:
                                                                if (v2 && ((v41 ^ v36) >= v41)) {
                                                                    var v48 float64 = 0.1
                                                                    var v49 bool = ((^(v36)) != 145)
                                                                    v50 := GenS6{F0: 0, F1: 1000}
                                                                    if (-9223372036854775808 != 9223372036854775807) {
                                                                        v51 := GenS4{F0: 347, F1: -1, F2: -1}
                                                                        v37 -= v37
                                                                        v52 := GenS4{F0: 0, F1: (v7 - 100), F2: -1}
                                                                        _ = v51
                                                                        _ = v52
                                                                    } else {
                                                                        switch 2147483647 {
                                                                        case -9223372036854775808:
                                                                            var v53 int8 = (v7 ^ -40)
                                                                            _ = v53
                                                                        case 42:
                                                                            fmt.Println(fi12)
                                                                        case -1:
                                                                            var v54 int64 = 1
                                                                            _ = v54
                                                                        }
                                                                        if v2 {
                                                                            v7 += (v7 * v7)
                                                                            v55 := GenS3{F0: fi12, F1: (v7 ^ (-(v7)))}
                                                                            if (v4 > 1.5) {
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = -1
                                                                                }
                                                                                v4 = (v4 * v4)
                                                                            }
                                                                            v49 = (!(v4 < (v4 / 3.14159)))
                                                                            if len(v6) >= 2 {
                                                                                v6 = v6[1:len(v6)]
                                                                            }
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            if (v49 || true) {
                                                                                if ((-(fi12)) >= fi12) {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    var v56 float64 = -2.25
                                                                                    if v49 {
                                                                                        switch ("" + "\n\t") {
                                                                                        case "foo bar":
                                                                                            v57 := GenS7{F0: (v7 * 100)}
                                                                                            _ = v57
                                                                                        case "":
                                                                                            v1 = v1
                                                                                        case "\n\t":
                                                                                            v48 -= v48
                                                                                        }
                                                                                    }
                                                                                    _ = v56
                                                                                } else {
                                                                                    v36 = 1
                                                                                    v58 := GenS3{F0: fi12, F1: (v7 | (^(v7)))}
                                                                                    fmt.Println(v49)
                                                                                    _ = v58
                                                                                }
                                                                                v7 += (v7 / (((v7 | 1)) | 1))
                                                                                if ((((fi19 / ((int(v37)) | 1)) <= (-(fi12))) && false) && (false || (4294967295 >= 42))) {
                                                                                    var v59 uint32 = 2
                                                                                    v60 := GenS2{F0: 18446744073709551615, F1: 1}
                                                                                    _ = v59
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            v36 += 0
                                                                            _ = v55
                                                                        } else {
                                                                            v61 := GenS9{F0: "\n\t", F1: 142}
                                                                            fmt.Println(v4)
                                                                            v62 := GenS1{F0: 798, F1: -1}
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v63 uint32 = 0
                                                        v1 += v1
                                                        _ = v63
                                                    }
                                                    if (v41 <= v36) {
                                                        v37 = (v41 + (v36 ^ v41))
                                                        switch fi12 {
                                                        case 752:
                                                            v3 = append(v3, fi12)
                                                        }
                                                    } else {
                                                        v64 := GenS10{F0: ""}
                                                        v36 = 2
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        _ = v64
                                                    }
                                                } else {
                                                    v65 := GenS11{F0: v7, F1: (!(!(v2 || v2))), F2: uint(v36)}
                                                    var v66 uint16 = v36
                                                    v67 := GenS11{F0: (-(v7)), F1: v2, F2: 226}
                                                    fmt.Println(v7)
                                                    if (v2 && false) {
                                                        fmt.Println(v1, fi12)
                                                        switch "hello" {
                                                        case "foo bar":
                                                            fmt.Println(v2)
                                                        default:
                                                            v36 -= 2
                                                        }
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            _ = v41
                                        }
                                        v4 = v4
                                        v68 := GenS12{F0: 18446744073709551615, F1: 0}
                                        _ = v68
                                    }
                                }
                                _ = v37
                            }
                            _ = v36
                        }
                        v2 = (!v2)
                        v69 := GenS13{F0: 65535, F1: 42}
                        _ = v69
                    }
                    v2 = v2
                    _ = fi19
                }
            }
            v70 := GenS7{F0: (v7 | (^(v7)))}
            _ = fi12
            _ = v70
        }
        if v2 {
            var v71 float32 = 0.1
            _ = v71
        } else {
            fmt.Println(v7, v4, v1)
            var v72 uint16 = 14
            for fi73 := 0; fi73 < 10; fi73++ {
                if v2 {
                    var v74 uint = 203
                    if v2 {
                        var v75 int8 = -110
                        var v76 uint16 = v72
                        _ = v75
                        _ = v76
                    } else {
                        var v77 uint32 = 42
                        if len(v5) > 0 {
                            v5[0] = -213
                        }
                        _ = v77
                    }
                    _ = v74
                } else {
                    if len(v6) > 0 {
                        v6[0] = (fi73 | (fi73 - (^(fi73))))
                    }
                }
                _ = fi73
            }
            _ = v72
        }
    }
    var v78 uint = 0
    for fi79 := 0; fi79 < 18; fi79++ {
        v80 := GenS14{F0: v78}
        var v81 uint32 = 2
        _ = fi79
        _ = v80
        _ = v81
    }
    v4 = v4
    fmt.Println(v4, v4, v4)
    v82 := GenS12{F0: (v78 % ((v78) | 1)), F1: 2147483648}
    fmt.Println(v1, v7, v7)
    v83 := GenS15{F0: v78, F1: -2147483648}
    if ((!v2) && v2) {
        if (((632 <= 18446744073709551615) && false) || (v1 == 1.5)) {
            v84 := GenS16{F0: 296}
            v85 := GenS16{F0: 899}
            _ = v84
            _ = v85
        }
    }
    var v86 uint16 = 0
    var v87 string = "a"
    for fi88 := 0; fi88 < 16; fi88++ {
        v4 += v4
        v7 = v7
        if false {
            var v89 int = (-(fi88))
            var v90 string = v87
            _ = v89
            _ = v90
        }
        _ = fi88
    }
    var v91 uint32 = 2
    v92 := GenS14{F0: v78}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v78:", v78)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v78
    _ = v8
    _ = v82
    _ = v83
    _ = v86
    _ = v87
    _ = v91
    _ = v92
}
