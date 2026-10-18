package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8; F1 int32 }
func genFn0(p0 int32) string {
    var v6 uint16 = 42
    v7 := GenS1{F0: (1.5 <= 0.0)}
    var v8 uint64 = 18446744073709551615
    v8 -= (v8 << ((42) & 63))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 string, p1 uint16) uint64 {
    switch p0 {
    case "hello":
        switch p0 {
        case "hello":
            v9 := GenS2{F0: 4294967295, F1: false, F2: 100}
            _ = v9
        case "":
            var v10 uint = 255
            _ = v10
        case "a":
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
    case "\n\t":
        for fi12 := 0; fi12 < 20; fi12++ {
            p1 = 255
            _ = fi12
        }
    }
    var v13 int = 568
    _ = p0
    _ = p1
    _ = v13
    return 334
}

func main() {
    var v1 uint16 = 32767
    var v2 bool = false
    var v3 uint8 = 42
    var v4 int32 = -915
    v5 := make(map[int]bool)
    v2 = (!true)
    for fi14 := 0; fi14 < 4; fi14++ {
        v1 += (v1 >> ((2) & 63))
        v15 := GenS3{F0: v2, F1: fi14}
        fmt.Println(v3)
        _ = fi14
        _ = v15
    }
    if ((false && v2) || v2) {
        if (!((1 != 42) && v2)) {
            fmt.Println(v2, v2)
            var v16 string = ("hello" + ("foo bar" + ("foo bar" + ("hello" + ((("a" + "") + "foo bar") + ((((("" + "hello") + "\n\t") + "foo bar") + "") + ("foo bar" + "")))))))
            var v17 int = -990
            _ = v16
            _ = v17
        }
        var v18 bool = (v2 && v2)
        var v19 uint64 = 4294967295
        _ = v18
        _ = v19
    }
    var v20 int = 42
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := []int{fi21, 42}
        v5[v20] = (!v2)
        _ = fi21
        _ = v22
    }
    v20 += v20
    fmt.Println(v20, v3)
    for fi23 := 0; fi23 < 7; fi23++ {
        v1 += (^(v1))
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := GenS2{F0: 285, F1: v2, F2: -128}
            if true {
                delete(v5, (fi23 + fi24))
                switch v20 {
                case 9223372036854775807:
                    var v26 float64 = (-2.25 + float64(3.14159))
                    _ = v26
                case 127:
                    v20 += (v20 & (fi24 | (fi24 ^ fi23)))
                default:
                    switch ("a" + "foo bar") {
                    case "\n\t":
                        delete(v5, (v20 >> ((255) & 63)))
                    }
                }
            }
            _ = fi24
            _ = v25
        }
        switch (v20 & 32767) {
        case 9223372036854775807:
            switch fi23 {
            case -1:
                if (0 == 0) {
                    v27 := []int{fi23, fi23, v20, v20}
                    var v28 uint64 = 0
                    v2 = (v3 < v3)
                    _ = v27
                    _ = v28
                } else {
                    switch "a" {
                    case "hello":
                        v29 := GenS4{F0: -577, F1: (1 == 1), F2: 65535}
                        _ = v29
                    case "":
                        if true {
                            fmt.Println(v1)
                            v2 = true
                        } else {
                            v30 := GenS3{F0: (3.14159 == 0.1), F1: v20}
                            v2 = v2
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 8; fi31++ {
                            v4 -= v4
                            switch "a" {
                            case "":
                                if (42 != -128) {
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = v3
                                        if (1 <= v4) {
                                            v32 := GenS5{F0: v1}
                                            _ = v32
                                        }
                                        var v33 uint16 = 2
                                        _ = v33
                                    case "hello":
                                        v3 -= v3
                                    default:
                                        v34 := GenS3{F0: v2, F1: fi31}
                                        _ = v34
                                    }
                                    v20 += fi23
                                    v35 := []int{v20, (^(v20)), -141, (fi23 | -859)}
                                    _ = v35
                                }
                            case "\n\t":
                                switch ("hello" + "") {
                                case "\n\t":
                                    if (!false) {
                                        fmt.Println(v20, v1, v3)
                                        switch (("" + "") + "hello") {
                                        case "":
                                            var v36 int = fi23
                                            _ = v36
                                        case "\n\t":
                                            v1 = (v1 - (v1 - (v1 | (v1 + v1))))
                                        case "a":
                                            v37 := v5[-1]
                                            _ = v37
                                        }
                                    }
                                case "":
                                    var v38 int8 = 1
                                    _ = v38
                                default:
                                    fmt.Println(v20)
                                }
                            default:
                                if (v4 > (v4 * (v4 - v4))) {
                                    var v39 int64 = 858
                                    _ = v39
                                }
                            }
                            v40 := []int{}
                            delete(v5, fi23)
                            var v41 uint = 60
                            switch v20 {
                            case 127:
                                var v42 uint8 = (v3 & 128)
                                _ = v42
                            case 32768:
                                v43 := v5[(-(v20))]
                                _ = v43
                            case -2147483648:
                                if (0.1 < float32(1.5)) {
                                    v41 = (v41 / ((uint(v1)) | 1))
                                }
                            }
                            var v44 uint64 = 0
                            v45 := GenS6{F0: int(v44)}
                            var v46 uint32 = 65535
                            _ = fi31
                            _ = v40
                            _ = v41
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                    }
                    v3 = v3
                }
            case 9223372036854775807:
                var v47 uint8 = (^(v3))
                _ = v47
            case 2147483647:
                v48 := GenS7{F0: -1, F1: 42}
                _ = v48
            default:
                v2 = v2
            }
        case 128:
            if ((!v2) && (1 >= 4294967295)) {
                var v49 uint16 = v1
                fmt.Println(v3, fi23, v4)
                _ = v49
            }
        case 42:
            v3 += v3
        default:
            v2 = (false || v2)
        }
        _ = fi23
    }
    v50 := GenS2{F0: 295, F1: (-2.25 > (0.1 / 1.5)), F2: 0}
    if (935 != 1) {
        v1 -= (v1 | 42)
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        switch "\n\t" {
        case "\n\t":
            switch ((("hello" + "") + "\n\t") + "foo bar") {
            case "\n\t":
                switch (-(v20)) {
                case -2147483648:
                    var v52 int = -805
                    _ = v52
                default:
                    if (!v2) {
                        v20 += fi51
                        fmt.Println(v1, v4)
                        if (("\n\t" >= "a") || v2) {
                            if (((("hello" + (("\n\t" + "foo bar") + "")) + ("" + "hello")) + "") < "\n\t") {
                                if (714 != genFn1("hello", (^(v1)))) {
                                    var v53 float32 = 0.1
                                    fmt.Println(v53, fi51)
                                    _ = v53
                                } else {
                                    if (!v2) {
                                        var v54 uint32 = 403
                                        for fi55 := 0; fi55 < 1; fi55++ {
                                            switch "foo bar" {
                                            case "foo bar":
                                                v2 = v2
                                            case "a":
                                                var v56 uint8 = v3
                                                _ = v56
                                            default:
                                                v57 := GenS8{F0: (v3 & 1)}
                                                _ = v57
                                            }
                                            _ = fi55
                                        }
                                        _ = v54
                                    } else {
                                        v58 := []int{107, 42}
                                        for fi59 := 0; fi59 < 18; fi59++ {
                                            fmt.Println(v2, fi59)
                                            _ = fi59
                                        }
                                        switch v20 {
                                        case 32767:
                                            var v60 int64 = 0
                                            _ = v60
                                        case 4294967295:
                                            var v61 string = ""
                                            _ = v61
                                        case 9223372036854775807:
                                            fmt.Println(fi51, v2)
                                        }
                                        _ = v58
                                    }
                                    if (156 >= 1) {
                                        v20 = (fi51 & fi51)
                                        fmt.Println(v3, fi51)
                                        v5[-2147483648] = v2
                                    }
                                    var v62 uint8 = v3
                                    _ = v62
                                }
                                v63 := GenS9{F0: 1, F1: (v4 + (v4 & v4))}
                                var v64 int8 = 127
                                _ = v63
                                _ = v64
                            }
                            var v65 uint64 = 4294967295
                            _ = v65
                        } else {
                            switch "foo bar" {
                            case "":
                                fmt.Println(fi51, v4, v4)
                            case "hello":
                                v66 := []int{(^(fi51)), fi51, (v20 % ((int(v3)) | 1))}
                                _ = v66
                            case "foo bar":
                                if true {
                                    v67 := GenS9{F0: -1, F1: (v4 * v4)}
                                    var v68 int = 1
                                    var v69 int8 = 1
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                } else {
                                    v20 = 0
                                    for fi70 := 0; fi70 < 8; fi70++ {
                                        v3 += (v3 >> ((65535) & 63))
                                        var v71 int8 = -128
                                        v20 += fi51
                                        v72 := GenS6{F0: fi51}
                                        _ = fi70
                                        _ = v71
                                        _ = v72
                                    }
                                    for fi73 := 0; fi73 < 20; fi73++ {
                                        var v74 int32 = v4
                                        v75 := GenS6{F0: v20}
                                        switch fi73 {
                                        case 1:
                                            var v76 int64 = -1
                                            _ = v76
                                        }
                                        _ = fi73
                                        _ = v74
                                        _ = v75
                                    }
                                }
                            default:
                                if v2 {
                                    fmt.Println(v2, v4)
                                    var v77 uint = 1
                                    v3 -= (^(v3))
                                    _ = v77
                                } else {
                                    v78 := []int{(v20 + v20), fi51, v20, v20}
                                    v79 := v5[fi51]
                                    for fi80 := 0; fi80 < 17; fi80++ {
                                        if v2 {
                                            if false {
                                                var v81 uint64 = 18446744073709551615
                                                _ = v81
                                            }
                                            var v82 uint32 = 2
                                            _ = v82
                                        }
                                        fmt.Println(v20)
                                        v4 = -2147483648
                                        _ = fi80
                                    }
                                    _ = v78
                                    _ = v79
                                }
                            }
                            v83 := []int{(v20 / (((v20 ^ fi51)) | 1))}
                            v4 += (v4 | v4)
                            _ = v83
                        }
                    }
                }
            default:
                v84 := v5[(fi51 - (fi51 % (((fi51 << ((uint(v1)) & 63))) | 1)))]
                _ = v84
            }
        case "hello":
            v4 -= 1
        case "a":
            v2 = (false && v2)
        }
        for fi85 := 0; fi85 < 18; fi85++ {
            v4 -= (-(v4))
            _ = fi85
        }
        var v86 string = ("\n\t" + (("" + "foo bar") + ("a" + "a")))
        _ = fi51
        _ = v86
    }
    if v2 {
        v1 -= (v1 * 2)
        var v87 uint64 = 1
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
