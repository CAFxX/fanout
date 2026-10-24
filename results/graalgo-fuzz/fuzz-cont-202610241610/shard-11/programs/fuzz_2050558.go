package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 bool; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int32 }
func genFn0() uint8 {
    var v6 uint8 = 42
    if ((19 <= 108) || false) {
        v6 -= 255
        var v7 uint8 = v6
        var v8 uint16 = 32767
        _ = v7
        _ = v8
    } else {
        var v9 uint = 2
        _ = v9
    }
    v10 := GenS1{F0: 42}
    if false {
        var v11 uint64 = 549
        v11 = (v11 % ((v11) | 1))
        _ = v11
    }
    v6 = (v6 - (v6 * v6))
    _ = v10
    _ = v6
    return 1
}
func genFn1(p0 int64, p1 uint8) int8 {
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(fi12, p0, p1)
        if false {
            var v13 uint = uint(p1)
            var v14 float64 = (3.14159 - 0.0)
            var v15 uint64 = 42
            _ = v13
            _ = v14
            _ = v15
        } else {
            p1 = p1
            fmt.Println(fi12, fi12)
            if false {
                if false {
                    if false {
                        if true {
                            for fi16 := 0; fi16 < 17; fi16++ {
                                if (float64((0.1 - -2.25)) != 1e10) {
                                    if (65535 > 178) {
                                        p1 += 1
                                        v17 := []int{}
                                        _ = v17
                                    } else {
                                        v18 := []int{(-(fi12)), fi16, (fi12 << ((255) & 63))}
                                        p1 = genFn0()
                                        _ = v18
                                    }
                                    var v19 int64 = p0
                                    _ = v19
                                }
                                v20 := []int{}
                                _ = fi16
                                _ = v20
                            }
                        } else {
                            var v21 int8 = -128
                            var v22 int = 9223372036854775807
                            _ = v21
                            _ = v22
                        }
                        p0 = (^(p0))
                    }
                    for fi23 := 0; fi23 < 7; fi23++ {
                        var v24 int = (fi12 / (((fi12 >> ((18446744073709551615) & 63))) | 1))
                        if (true && false) {
                            var v25 uint16 = 869
                            p1 = genFn0()
                            v25 = v25
                            p0 = (p0 | p0)
                            _ = v25
                        }
                        switch (-(fi23)) {
                        case -9223372036854775808:
                            p1 += p1
                        case 9223372036854775807:
                            v26 := GenS2{F0: (v24 >> ((42) & 63)), F1: 920, F2: p1}
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                p0 += (p0 - p0)
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    var v28 uint32 = 4294967295
                    p0 = p0
                    var v29 uint16 = 65535
                    v30 := []int{}
                    if len(v30) > 0 {
                        v30[0] = (fi12 * 42)
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
        }
        var v31 int64 = p0
        _ = fi12
        _ = v31
    }
    p0 += (^(p0))
    var v32 uint8 = p1
    p0 = p0
    _ = p0
    _ = p1
    _ = v32
    return -1
}

func main() {
    var v1 uint8 = 2
    v2 := []int{1}
    v3 := make(map[string]int)
    var v4 uint64 = 2
    var v5 uint64 = (v4 % (((v4 / ((v4) | 1))) | 1))
    v2 = append(v2, 2147483647)
    for fi33 := 0; fi33 < 1; fi33++ {
        v34 := v3["hello"]
        _ = fi33
        _ = v34
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v1 -= v1
    if ("foo bar" != ("foo bar" + (("a" + ("\n\t" + ("a" + "hello"))) + ((("a" + ("foo bar" + ("hello" + "hello"))) + ("foo bar" + "a")) + ("" + (("\n\t" + ("\n\t" + (("\n\t" + "foo bar") + "hello"))) + "a")))))) {
        if false {
            v1 = v1
        } else {
            v2 = append(v2, 128)
            var v35 uint8 = (v1 << ((2) & 63))
            v5 = v4
            _ = v35
        }
    }
    v36 := GenS2{F0: 32768, F1: 0, F2: v1}
    v37 := v3[("" + ((("" + "hello") + "\n\t") + "foo bar"))]
    v3[(("" + "a") + "")] = v37
    v5 -= 210
    for fi38 := 0; fi38 < 13; fi38++ {
        var v39 uint = 4294967295
        for fi40 := 0; fi40 < 1; fi40++ {
            switch v37 {
            case 0:
                _ = len(v2)
            }
            if false {
                v1 = 153
                var v41 uint = (^(v39))
                v42 := GenS1{F0: uint32(v1)}
                _ = v41
                _ = v42
            } else {
                var v43 int32 = 0
                _ = v43
            }
            v44 := GenS2{F0: fi40, F1: 42, F2: 42}
            _ = fi40
            _ = v44
        }
        var v45 uint64 = 18446744073709551615
        _ = fi38
        _ = v39
        _ = v45
    }
    if false {
        var v46 int32 = -1
        var v47 uint16 = 2
        _ = v46
        _ = v47
    } else {
        if true {
            if len(v2) > 0 {
                v2[0] = (v37 >> ((439) & 63))
            }
            delete(v3, "")
            v48 := GenS1{F0: 4294967295}
            _ = v48
        } else {
            if false {
                if true {
                    var v49 int64 = -197
                    if false {
                        v49 += v49
                        v50 := v3["hello"]
                        _ = v50
                    } else {
                        if (false || (!true)) {
                            if (!(!(true && ("a" < ("" + "a"))))) {
                                for fi51 := 0; fi51 < 2; fi51++ {
                                    v4 -= (v5 & 116)
                                    _ = fi51
                                }
                                for fi52 := 0; fi52 < 11; fi52++ {
                                    var v53 uint32 = 2
                                    _ = fi52
                                    _ = v53
                                }
                            }
                            v54 := v3[""]
                            _ = v54
                        } else {
                            var v55 uint = 255
                            if (!(int8(v1) != 0)) {
                                delete(v3, "hello")
                                v4 = (v5 | v4)
                                fmt.Println(v37, v5)
                            } else {
                                var v56 int16 = 0
                                v1 += (^(v1))
                                v57 := GenS1{F0: 42}
                                _ = v56
                                _ = v57
                            }
                            var v58 int16 = 1000
                            _ = v55
                            _ = v58
                        }
                        var v59 bool = (-414 != v49)
                        _ = v59
                    }
                    v5 += (v4 ^ (v5 << ((703) & 63)))
                    _ = v49
                }
                var v60 uint = uint(v1)
                _ = v60
            }
        }
    }
    v4 += (v4 << ((784) & 63))
    switch v37 {
    case 32768:
        switch (("\n\t" + "\n\t") + "foo bar") {
        case "a":
            switch "foo bar" {
            case "foo bar":
                var v61 int = v37
                _ = v61
            default:
                _ = len(v3)
            }
        case "hello":
            if (!false) {
                if false {
                    if (2 == 2) {
                        if (!false) {
                            switch (v37 >> ((uint(v1)) & 63)) {
                            case 32767:
                                for fi62 := 0; fi62 < 19; fi62++ {
                                    if (100 == 127) {
                                        if true {
                                            var v63 int64 = -393
                                            _ = v63
                                        }
                                        var v64 uint8 = v1
                                        v1 -= v1
                                        _ = v64
                                    } else {
                                        fmt.Println(v1, v4, fi62)
                                    }
                                    for fi65 := 0; fi65 < 6; fi65++ {
                                        switch "a" {
                                        case "foo bar":
                                            if (0 >= 1) {
                                                var v66 uint16 = 32767
                                                var v67 int = (fi65 % (((v37 | 4294967295)) | 1))
                                                v68 := GenS2{F0: 32767, F1: 866, F2: (v1 % ((v1) | 1))}
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                if true {
                                                    if false {
                                                        switch "" {
                                                        case "\n\t":
                                                            fmt.Println(v5, v4)
                                                        case "hello":
                                                            v5 += (v5 | 18446744073709551615)
                                                        case "foo bar":
                                                            v69 := GenS2{F0: (-(fi65)), F1: 65535, F2: (v1 % ((genFn0()) | 1))}
                                                            _ = v69
                                                        }
                                                        if true {
                                                            var v70 int32 = -580
                                                            var v71 uint8 = 42
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            fmt.Println(v37, v5)
                                                            v72 := GenS1{F0: 4294967295}
                                                            v5 = (v4 / (((v4 * (v5 >> ((4294967295) & 63)))) | 1))
                                                            _ = v72
                                                        }
                                                        v4 -= 4294967295
                                                    }
                                                } else {
                                                    var v73 uint16 = 0
                                                    fmt.Println(v73, v4, v5)
                                                    switch (fi62 % ((fi65) | 1)) {
                                                    case -389:
                                                        fmt.Println(v37, fi65)
                                                    case 90:
                                                        v73 = (v73 - v73)
                                                    case 2147483647:
                                                        v1 -= v1
                                                    default:
                                                        var v74 bool = true
                                                        _ = v74
                                                    }
                                                    _ = v73
                                                }
                                            }
                                        case "\n\t":
                                            if true {
                                                v75 := GenS1{F0: 2}
                                                var v76 uint8 = genFn0()
                                                _ = v75
                                                _ = v76
                                            }
                                        default:
                                            var v77 bool = false
                                            _ = v77
                                        }
                                        _ = fi65
                                    }
                                    _ = fi62
                                }
                            case -382:
                                var v78 uint = 2
                                _ = v78
                            case 127:
                                for fi79 := 0; fi79 < 3; fi79++ {
                                    v37 -= (v37 & 260)
                                    switch (v37 >> ((42) & 63)) {
                                    case -972:
                                        switch int(v4) {
                                        case -586:
                                            if true {
                                                if (((1.5 * 1e10) - 0.1) > 0.1) {
                                                    v80 := GenS3{F0: 0, F1: true, F2: fi79}
                                                    _ = v80
                                                }
                                                for fi81 := 0; fi81 < 4; fi81++ {
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(v1, fi79, fi81)
                                                    case "foo bar":
                                                        if false {
                                                            var v82 int8 = 100
                                                            switch fi81 {
                                                            case 9223372036854775807:
                                                                var v83 int8 = (v82 + (v82 | -77))
                                                                _ = v83
                                                            case 1:
                                                                var v84 uint32 = 0
                                                                _ = v84
                                                            default:
                                                                if (!(!(539 < 0))) {
                                                                    v85 := GenS3{F0: 4294967295, F1: (("a" + "a") == "foo bar"), F2: -428}
                                                                    _ = v85
                                                                } else {
                                                                    switch "a" {
                                                                    case "":
                                                                        switch (("" + "\n\t") + ("\n\t" + "\n\t")) {
                                                                        case "foo bar":
                                                                            var v86 uint = 604
                                                                            _ = v86
                                                                        default:
                                                                            v4 -= 77
                                                                        }
                                                                    default:
                                                                        v37 = (fi79 * v37)
                                                                    }
                                                                    fmt.Println(v1)
                                                                }
                                                            }
                                                            _ = v82
                                                        }
                                                    }
                                                    v37 = 598
                                                    v5 = (v4 | v4)
                                                    _ = fi81
                                                }
                                                if false {
                                                    v4 = (v4 / ((496) | 1))
                                                    var v87 uint16 = 42
                                                    _ = v87
                                                }
                                            }
                                        case 32768:
                                            var v88 int16 = -1
                                            _ = v88
                                        default:
                                            var v89 int16 = -1
                                            _ = v89
                                        }
                                    case 127:
                                        v90 := GenS4{F0: (fi79 % ((9223372036854775807) | 1))}
                                        _ = v90
                                    }
                                    _ = fi79
                                }
                            }
                            _ = len(v3)
                        }
                        _ = len(v2)
                        var v91 uint = 572
                        _ = v91
                    } else {
                        v92 := GenS5{F0: uint(v1), F1: -254}
                        if ((v1 & v1) < v1) {
                            switch (v37 << ((18446744073709551615) & 63)) {
                            case 4294967295:
                                switch ((("a" + ("a" + "a")) + ("" + ("foo bar" + "\n\t"))) + ("a" + ("foo bar" + ("hello" + "a")))) {
                                case "\n\t":
                                    v93 := v3["\n\t"]
                                    _ = v93
                                default:
                                    var v94 int64 = -1
                                    _ = v94
                                }
                            case 960:
                                var v95 uint16 = 2
                                _ = v95
                            case 42:
                                var v96 uint64 = (^(v4))
                                _ = v96
                            default:
                                fmt.Println(v4, v1)
                            }
                            switch "a" {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = 127
                                }
                            case "foo bar":
                                if ((true && (!(v37 > v37))) && (!false)) {
                                    delete(v3, ("foo bar" + ""))
                                    fmt.Println(v37, v5, v37)
                                    var v97 uint64 = v5
                                    _ = v97
                                }
                            case "hello":
                                v5 = 0
                            }
                        }
                        var v98 uint64 = (v4 + v4)
                        _ = v92
                        _ = v98
                    }
                } else {
                    if true {
                        var v99 float32 = 1e10
                        for fi100 := 0; fi100 < 6; fi100++ {
                            v99 = 3.14159
                            v4 -= v4
                            _ = fi100
                        }
                        _ = v99
                    }
                }
            }
        case "\n\t":
            for fi101 := 0; fi101 < 14; fi101++ {
                _ = len(v3)
                for fi102 := 0; fi102 < 5; fi102++ {
                    v4 = v5
                    v103 := GenS2{F0: (fi101 & (fi101 % ((2147483647) | 1))), F1: 186, F2: v1}
                    _ = fi102
                    _ = v103
                }
                var v104 uint64 = v4
                _ = fi101
                _ = v104
            }
        default:
            if false {
                for fi105 := 0; fi105 < 20; fi105++ {
                    v106 := GenS4{F0: 32767}
                    for fi107 := 0; fi107 < 11; fi107++ {
                        var v108 int = fi105
                        fmt.Println(fi105)
                        _ = fi107
                        _ = v108
                    }
                    _ = fi105
                    _ = v106
                }
                for fi109 := 0; fi109 < 14; fi109++ {
                    fmt.Println(v5)
                    _ = fi109
                }
            }
        }
    }
    fmt.Println(v1, v37)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
