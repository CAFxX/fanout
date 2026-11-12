package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 int; F1 int32; F2 uint16 }
type GenS9 struct { F0 int }
func genFn0(p0 uint64, p1 uint) uint32 {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 bool = (2147483648 <= 4294967295)
        v6 := GenS1{F0: (0 != 212)}
        v7 := []int{int(p1), (fi4 ^ 9223372036854775807), fi4}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        p0 = p0
        var v9 float64 = (0.1 + 1.5)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{fi10, (fi10 ^ 9223372036854775807), fi10, fi10}
        switch "a" {
        case "foo bar":
            p1 += (p1 ^ (p1 & p1))
        case "\n\t":
            var v12 uint16 = 255
            _ = v12
        case "":
            p0 += 1
        default:
            var v13 uint = 18446744073709551615
            _ = v13
        }
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS2{F0: "a", F1: 1}
            var v16 uint8 = 0
            switch (fi14 ^ (fi14 * (fi10 | -228))) {
            case 4294967295:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            default:
                var v17 uint64 = 1
                _ = v17
            }
            p0 -= (p0 & 42)
            v18 := GenS2{F0: "foo bar", F1: 2147483647}
            _ = fi14
            _ = v15
            _ = v16
            _ = v18
        }
        _ = fi10
        _ = v11
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        p0 -= p0
        var v20 int32 = -903
        v21 := GenS3{F0: 0}
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 255}
    _ = p0
    _ = p1
    _ = v22
    return 42
}
func genFn1(p0 string, p1 int64, p2 uint32) string {
    switch "hello" {
    case "\n\t":
        var v23 float64 = ((1.5 * 3.14159) * (1.5 * (1e10 - (-2.25 + (((1e10 / 1e10) * ((-2.25 / 1.5) / 1.0)) - (1.5 / (0.1 + 3.14159)))))))
        _ = v23
    case "":
        v24 := GenS1{F0: true}
        _ = v24
    }
    p1 = 9223372036854775807
    switch p0 {
    case "a":
        switch p0 {
        case "hello":
            p1 += -1
        case "\n\t":
            if (false || (!true)) {
                v25 := []int{}
                for fi26 := 0; fi26 < 2; fi26++ {
                    p0 = (p0 + p0)
                    v27 := GenS4{F0: 1, F1: 0}
                    for fi28 := 0; fi28 < 12; fi28++ {
                        p2 += genFn0(1, 4294967295)
                        var v29 float64 = -2.25
                        var v30 float32 = 1e10
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
                _ = v25
            }
        default:
            var v31 int16 = -1
            _ = v31
        }
    case "hello":
        v32 := GenS5{F0: ("hello" + p0)}
        _ = v32
    default:
        v33 := []int{32768, 32767}
        _ = v33
    }
    v34 := GenS2{F0: "", F1: -2147483648}
    var v35 uint32 = (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v35
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 819
    var v3 int16 = 32767
    _ = len(v1)
    v3 -= (v3 - 1)
    switch 909 {
    case -1:
        switch 127 {
        case -1:
            v36 := GenS4{F0: 675, F1: 0}
            _ = v36
        }
    default:
        var v37 int16 = 32767
        _ = v37
    }
    for fi38 := 0; fi38 < 10; fi38++ {
        switch "a" {
        case "\n\t":
            if (255 == 228) {
                v2 += v2
            }
        case "foo bar":
            v3 = (v3 & v3)
        default:
            v2 = v2
        }
        for fi39 := 0; fi39 < 19; fi39++ {
            var v40 uint16 = 381
            if true {
                var v41 uint64 = 857
                v42 := []int{}
                _ = v41
                _ = v42
            } else {
                var v43 float32 = 3.14159
                _ = v43
            }
            switch "\n\t" {
            case "":
                v2 = v2
            default:
                v44 := GenS6{F0: -1}
                _ = v44
            }
            _ = fi39
            _ = v40
        }
        var v45 int8 = -1
        _ = fi38
        _ = v45
    }
    _ = len(v1)
    _ = len(v1)
    fmt.Println(v2, v3)
    var v46 uint64 = 1
    if false {
        var v47 uint8 = 42
        _ = v47
    } else {
        switch -14 {
        case -2147483648:
            for fi48 := 0; fi48 < 10; fi48++ {
                v49 := []int{1, fi48, (fi48 >> ((18446744073709551615) & 63))}
                var v50 int64 = -269
                if (116 > 65535) {
                    fmt.Println(v50, v50, v50)
                    if (!(!(v46 == v46))) {
                        v46 += v46
                    }
                    var v51 string = ((((("hello" + ("\n\t" + ((("hello" + "") + "") + "hello"))) + "hello") + "hello") + "\n\t") + "")
                    _ = v51
                } else {
                    if true {
                        switch (fi48 | fi48) {
                        case 0:
                            v1[95] = 2147483648
                        case -1:
                            switch fi48 {
                            case 32767:
                                switch 0 {
                                case 42:
                                    for fi52 := 0; fi52 < 3; fi52++ {
                                        var v53 int = (fi52 % ((32767) | 1))
                                        fmt.Println(v3)
                                        v54 := GenS7{F0: 65535, F1: ("foo bar" + "")}
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                    }
                                case 654:
                                    v46 += v46
                                default:
                                    if len(v49) > 0 {
                                        v49[0] = fi48
                                    }
                                }
                            }
                        default:
                            if (1 == 488) {
                                switch (("hello" + ("\n\t" + "")) + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v50)
                                }
                            }
                        }
                    }
                    switch (fi48 | 297) {
                    case 271:
                        v50 = -9223372036854775808
                    case 739:
                        v3 = 1
                    default:
                        v3 = (v3 - v3)
                    }
                }
                _ = fi48
                _ = v49
                _ = v50
            }
        case 0:
            if true {
                for fi55 := 0; fi55 < 16; fi55++ {
                    v56 := []int{fi55}
                    _ = fi55
                    _ = v56
                }
                var v57 uint8 = 128
                _ = v57
            } else {
                v58 := []int{127, -1}
                _ = v58
            }
        }
        v46 += 42
    }
    var v59 float32 = 1e10
    if ((v59 - v59) > 1e10) {
        var v60 int = 9223372036854775807
        switch (^(v60)) {
        case 4294967295:
            v3 = 32767
        default:
            for fi61 := 0; fi61 < 13; fi61++ {
                v46 += (^(v46))
                if ((v59 / 1e10) != 0.0) {
                    for fi62 := 0; fi62 < 2; fi62++ {
                        if (-6 != 127) {
                            if true {
                                if false {
                                    var v63 int64 = 9223372036854775807
                                    v64 := v1[761]
                                    if (v64 >= (v64 & 42)) {
                                        v65 := GenS6{F0: v3}
                                        _ = v65
                                    }
                                    switch fi62 {
                                    case -754:
                                        v66 := []int{127, fi62}
                                        _ = v66
                                    }
                                    if true {
                                        v63 -= (v63 & 0)
                                        var v67 bool = ((-1 == v63) || true)
                                        _ = v67
                                    }
                                    _ = v63
                                    _ = v64
                                } else {
                                    v68 := GenS2{F0: ("" + "\n\t"), F1: v2}
                                    var v69 uint = 255
                                    _ = v68
                                    _ = v69
                                }
                                var v70 float64 = 3.14159
                                v59 = ((1.5 / v59) + v59)
                                if false {
                                    var v71 int = (-(fi61))
                                    _ = v71
                                } else {
                                    var v72 int32 = int32(v46)
                                    var v73 uint16 = 65535
                                    fmt.Println(v46)
                                    _ = v72
                                    _ = v73
                                }
                                v2 -= 0
                                _ = v70
                            }
                        } else {
                            v74 := []int{int(v2), -1, fi61, fi61}
                            delete(v1, 42)
                            var v75 int16 = (v3 ^ (v3 * v3))
                            switch v60 {
                            case -9223372036854775808:
                                if false {
                                    v75 -= (v3 - -1)
                                    fmt.Println(v60, v46, v46)
                                } else {
                                    var v76 float32 = 1.5
                                    _ = v76
                                }
                                var v77 uint = 42
                                v78 := GenS8{F0: (fi61 + 338), F1: v2, F2: 238}
                                _ = v77
                                _ = v78
                            case 42:
                                var v79 uint64 = v46
                                _ = v79
                            case 4294967295:
                                if len(v74) >= 2 {
                                    v74 = v74[1:len(v74)]
                                }
                            default:
                                if ((("" + "") >= "a") || false) {
                                    if (831 == 147) {
                                        var v80 int8 = -59
                                        if (((0.1 - 0.0) != (-2.25 / 1.0)) && false) {
                                            if len(v74) >= 2 {
                                                v74 = v74[1:len(v74)]
                                            }
                                            fmt.Println(v46, v80)
                                        } else {
                                            var v81 uint64 = 0
                                            var v82 int32 = v2
                                            _ = v81
                                            _ = v82
                                        }
                                        v83 := GenS5{F0: "\n\t"}
                                        v84 := GenS9{F0: fi61}
                                        v46 -= (v46 / ((v46) | 1))
                                        _ = v80
                                        _ = v83
                                        _ = v84
                                    }
                                } else {
                                    var v85 int64 = 9223372036854775807
                                    v1[42] = genFn0((v46 * v46), 1)
                                    _ = v85
                                }
                                v86 := GenS5{F0: "\n\t"}
                                v87 := GenS7{F0: uint16(v46), F1: "\n\t"}
                                _ = v86
                                _ = v87
                            }
                            _ = v74
                            _ = v75
                        }
                        v88 := GenS3{F0: 2}
                        _ = fi62
                        _ = v88
                    }
                } else {
                    switch ("\n\t" + (("hello" + "") + "\n\t")) {
                    case "\n\t":
                        var v89 uint32 = 742
                        _ = v89
                    }
                    var v90 uint16 = 643
                    _ = v90
                }
                _ = fi61
            }
        }
        _ = v60
    } else {
        if (true && (!(42 != 42))) {
            v91 := []int{}
            fmt.Println(v2, v59, v59)
            for fi92 := 0; fi92 < 4; fi92++ {
                v93 := GenS3{F0: 246}
                _ = fi92
                _ = v93
            }
            _ = v91
        } else {
            v59 += (1e10 - -2.25)
        }
    }
    if ((((740 < int64(v46)) || (true || ((-9223372036854775808 > 9223372036854775807) || (!true)))) && true) && (58 > 100)) {
        v1[2147483648] = genFn0(v46, 137)
        for fi94 := 0; fi94 < 15; fi94++ {
            var v95 int16 = (-(v3))
            _ = fi94
            _ = v95
        }
        var v96 int32 = (v2 ^ 1)
        _ = v96
    }
    v2 = (v2 & int32(v3))
    var v97 int8 = -128
    var v98 string = "foo bar"
    fmt.Println(v46, v46, v59)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v59:", v59)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v59
    _ = v97
    _ = v98
}
