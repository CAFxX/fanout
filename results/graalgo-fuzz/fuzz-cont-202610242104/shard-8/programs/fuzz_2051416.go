package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 bool; F2 string }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 32767
    var v3 string = "\n\t"
    var v4 uint64 = 42
    v5 := make(map[string]int)
    var v6 int32 = -2147483648
    var v7 uint = 1
    v6 -= -1
    v8 := GenS1{F0: (v4 | (^(v4))), F1: (v4 + 986)}
    v7 += v7
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            if true {
                v4 = 4294967295
            }
            v2 -= v2
            v11 := GenS1{F0: v4, F1: v4}
            if (!((true && true) || true)) {
                _ = len(v5)
                if true {
                    v12 := GenS1{F0: v4, F1: 2}
                    if false {
                        v4 = v4
                        var v13 float32 = 1e10
                        v3 = "foo bar"
                        _ = v13
                    }
                    switch v3 {
                    case "hello":
                        v14 := []int{}
                        _ = v14
                    case "foo bar":
                        switch 32768 {
                        case 9223372036854775807:
                            var v15 float64 = (-2.25 * ((0.1 - 3.14159) * 3.14159))
                            _ = v15
                        case 32768:
                            if false {
                                switch (-(fi9)) {
                                case -958:
                                    if false {
                                        v3 = "\n\t"
                                    } else {
                                        if (2147483647 < 371) {
                                            if true {
                                                if false {
                                                    if ((v7 * (^(v7))) <= v7) {
                                                        var v16 float32 = 1e10
                                                        var v17 int8 = 17
                                                        if (((v6 * v6) <= v6) || true) {
                                                            v18 := []int{42}
                                                            _ = v18
                                                        } else {
                                                            switch "hello" {
                                                            case "":
                                                                var v19 uint32 = 118
                                                                switch 32767 {
                                                                case -1:
                                                                    var v20 uint = (v7 ^ v7)
                                                                    _ = v20
                                                                case 127:
                                                                    var v21 uint8 = 255
                                                                    _ = v21
                                                                case 1:
                                                                    v16 = v16
                                                                }
                                                                _ = v19
                                                            case "foo bar":
                                                                fmt.Println(v16)
                                                            case "a":
                                                                v22 := []int{(fi9 >> ((65535) & 63)), -2147483648, 4294967295}
                                                                _ = v22
                                                            }
                                                            v23 := []int{1, fi9, fi10}
                                                            var v24 uint16 = 403
                                                            _ = v23
                                                            _ = v24
                                                        }
                                                        _ = v16
                                                        _ = v17
                                                    }
                                                    if (v7 == 255) {
                                                        v7 += 1
                                                    } else {
                                                        var v25 int64 = 791
                                                        _ = v25
                                                    }
                                                    v26 := GenS1{F0: v4, F1: (v4 ^ v4)}
                                                    v7 -= 255
                                                    delete(v5, v3)
                                                    _ = v26
                                                }
                                                v6 = v6
                                                v27 := []int{}
                                                v28 := GenS1{F0: (v4 + 302), F1: 635}
                                                var v29 int16 = 1
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            } else {
                                                switch "" {
                                                case "a":
                                                    v30 := GenS1{F0: v4, F1: (^(v4))}
                                                    _ = v30
                                                case "":
                                                    switch -2147483648 {
                                                    case 0:
                                                        v6 = (v6 << ((v7) & 63))
                                                    case -562:
                                                        switch fi10 {
                                                        case 32768:
                                                            v2 -= (v2 << ((uint(v4)) & 63))
                                                            var v31 int64 = -827
                                                            _ = v31
                                                        case 128:
                                                            var v32 uint64 = v4
                                                            if (42 <= 255) {
                                                                fmt.Println(fi10, fi10, fi10)
                                                                v4 -= v4
                                                            } else {
                                                                v33 := []int{(fi10 % (((fi9 / (((fi10 << ((v7) & 63))) | 1))) | 1))}
                                                                var v34 float32 = (3.14159 + 1e10)
                                                                var v35 float64 = 0.1
                                                                var v36 uint64 = (v32 | (v32 - 158))
                                                                _ = v33
                                                                _ = v34
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                            if false {
                                                                _ = len(v5)
                                                                var v37 bool = (v3 > v3)
                                                                if v37 {
                                                                    v1[""] = (((v2 * (^(v2))) > (v2 ^ 42)) || ((!v37) || (v37 && v37)))
                                                                    v38 := GenS1{F0: (v4 - 4294967295), F1: 42}
                                                                    var v39 int16 = v2
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    var v40 uint64 = 441
                                                                    _ = v40
                                                                }
                                                                _ = v37
                                                            } else {
                                                                delete(v1, "hello")
                                                                var v41 uint32 = 0
                                                                _ = v41
                                                            }
                                                            _ = v32
                                                        default:
                                                            v6 = int32(v2)
                                                        }
                                                    case -267:
                                                        v6 += -1
                                                        v3 = v3
                                                        var v42 int64 = 100
                                                        _ = v42
                                                    }
                                                }
                                                switch fi9 {
                                                case -255:
                                                    if (-2.25 < 0.1) {
                                                        if (!(!(2147483648 > 65535))) {
                                                            var v43 bool = false
                                                            _ = v43
                                                        }
                                                        fmt.Println(v7)
                                                    }
                                                default:
                                                    fmt.Println(v6)
                                                }
                                                v6 -= 87
                                            }
                                            v2 = (v2 / (((v2 + (v2 & (^(v2))))) | 1))
                                        }
                                    }
                                default:
                                    if false {
                                        v44 := GenS1{F0: 18446744073709551615, F1: 0}
                                        v45 := GenS1{F0: (v4 * (v4 ^ v4)), F1: (v4 << ((18446744073709551615) & 63))}
                                        _ = v44
                                        _ = v45
                                    } else {
                                        var v46 string = v3
                                        if true {
                                            var v47 float64 = 0.1
                                            if (v6 != (v6 - 0)) {
                                                _ = len(v5)
                                                switch "" {
                                                case "hello":
                                                    var v48 uint8 = 255
                                                    _ = v48
                                                default:
                                                    v49 := GenS2{F0: 128, F1: true, F2: ("a" + v3)}
                                                    _ = v49
                                                }
                                                var v50 int16 = v2
                                                _ = v50
                                            } else {
                                                v4 += (^(v4))
                                                if true {
                                                    v6 = 42
                                                    var v51 float64 = ((v47 / 1.0) * 1e10)
                                                    v52 := GenS2{F0: 32, F1: false, F2: ("foo bar" + ("a" + v3))}
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        delete(v5, (("a" + "a") + v46))
                                                    case "":
                                                        v7 = (v7 >> (((v7 * 18446744073709551615)) & 63))
                                                    case "a":
                                                        var v53 uint32 = 2147483647
                                                        fmt.Println(v7, v53, v46)
                                                        v54 := GenS2{F0: 128, F1: true, F2: ((("\n\t" + v46) + v46) + v46)}
                                                        if (!true) {
                                                            fmt.Println(v4, v46)
                                                            v51 += v51
                                                            v2 = (v2 | (v2 << (((v7 + 0)) & 63)))
                                                            switch -953 {
                                                            case 531:
                                                                var v55 uint16 = 32767
                                                                fmt.Println(v51, fi10)
                                                                var v56 uint64 = (^(v4))
                                                                _ = v55
                                                                _ = v56
                                                            }
                                                        } else {
                                                            v4 = uint64(v53)
                                                            v47 -= v51
                                                        }
                                                        var v57 uint32 = v53
                                                        _ = v53
                                                        _ = v54
                                                        _ = v57
                                                    }
                                                    _ = v51
                                                    _ = v52
                                                } else {
                                                    if (v7 < (v7 % (((^(v7))) | 1))) {
                                                        if (fi9 >= -1) {
                                                            var v58 int32 = -1
                                                            _ = v58
                                                        } else {
                                                            v59 := []int{0, (fi9 - fi10)}
                                                            var v60 int8 = 26
                                                            v46 = v3
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                        if (219 < 128) {
                                                            if true {
                                                                var v61 bool = (125 > 2)
                                                                v62 := GenS2{F0: 42, F1: (false || true), F2: v3}
                                                                v46 = (v3 + v46)
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                var v63 float64 = (-2.25 + float64(3.14159))
                                                                fmt.Println(v46)
                                                                v64 := GenS3{F0: (^(v4)), F1: -128, F2: 32}
                                                                var v65 bool = false
                                                                _ = v63
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            if (!false) {
                                                                var v66 uint32 = 42
                                                                _ = v66
                                                            }
                                                        }
                                                        fmt.Println(v3)
                                                        v67 := GenS3{F0: v4, F1: 1, F2: 32}
                                                        var v68 int = (fi9 * fi10)
                                                        v6 -= (v6 * v6)
                                                        v4 = 707
                                                        v2 += (v2 * (v2 & v2))
                                                        var v69 uint16 = 1
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                }
                                            }
                                            _ = v47
                                        }
                                        v2 += v2
                                        v4 = 171
                                        v70 := GenS4{F0: 391, F1: 9223372036854775807, F2: 42}
                                        _ = v46
                                        _ = v70
                                    }
                                }
                                v71 := []int{fi10, (-(fi10)), 128}
                                _ = v71
                            } else {
                                v72 := GenS2{F0: 134, F1: (!false), F2: (("hello" + v3) + v3)}
                                _ = v72
                            }
                        default:
                            if false {
                                v73 := GenS5{F0: fi10, F1: v4, F2: (v6 % (((v6 + v6)) | 1))}
                                _ = v73
                            }
                        }
                    case "\n\t":
                        switch 9223372036854775807 {
                        case 4294967295:
                            fmt.Println(fi10, fi9)
                        default:
                            v4 += v4
                        }
                    default:
                        v74 := []int{(fi10 | 42), fi10, (fi10 / (((fi9 * 2147483647)) | 1))}
                        _ = v74
                    }
                    _ = v12
                }
                fmt.Println(v7, v4)
            }
            v75 := []int{fi9, fi10}
            _ = fi10
            _ = v11
            _ = v75
        }
        _ = fi9
    }
    v76 := GenS6{F0: 1}
    var v77 int32 = (v6 % (((v6 << (((^(v7))) & 63))) | 1))
    var v78 uint16 = 42
    v79 := GenS7{F0: int8(v7), F1: (^(v2))}
    for fi80 := 0; fi80 < 19; fi80++ {
        v78 = v78
        v81 := GenS4{F0: v4, F1: 9223372036854775807, F2: v78}
        v77 += (v6 + 0)
        _ = fi80
        _ = v81
    }
    v3 = v3
    for fi82 := 0; fi82 < 18; fi82++ {
        if (false && (!false)) {
            v83 := GenS8{F0: true}
            _ = v83
        }
        var v84 uint32 = 355
        var v85 uint8 = 66
        _ = fi82
        _ = v84
        _ = v85
    }
    switch 32768 {
    case 32767:
        fmt.Println(v7, v78, v77)
    case 28:
        v2 = (v2 + v2)
    case -860:
        switch "foo bar" {
        case "hello":
            fmt.Println(v4)
        }
    default:
        v86 := GenS9{F0: v3}
        _ = v86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v8
}
