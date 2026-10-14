package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint; F1 bool; F2 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int; F2 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32 }
func genFn0(p0 bool) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        if p0 {
            v7 := []int{127, fi6, -1, (fi6 << ((175) & 63))}
            for fi8 := 0; fi8 < 13; fi8++ {
                p0 = (0 > -1)
                _ = fi8
            }
            _ = v7
        } else {
            v9 := GenS1{F0: -9223372036854775808, F1: 127}
            fmt.Println(fi6)
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int8 = 127
            v12 := GenS2{F0: "hello", F1: 1}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    if ((-2147483648 > -1) && p0) {
        if (!((1 != 127) && (p0 && p0))) {
            switch 301 {
            case -154:
                for fi13 := 0; fi13 < 16; fi13++ {
                    v14 := []int{(fi13 | fi13), fi13, 0}
                    fmt.Println(p0)
                    switch "a" {
                    case "\n\t":
                        var v15 bool = (p0 || p0)
                        _ = v15
                    case "":
                        switch (fi13 & 862) {
                        case 0:
                            v16 := GenS1{F0: (fi13 ^ (fi13 >> ((65535) & 63))), F1: 203}
                            _ = v16
                        case 4294967295:
                            p0 = (!((42 > 0) && p0))
                        default:
                            if p0 {
                                var v17 float32 = float32(0.0)
                                _ = v17
                            } else {
                                if (!(2147483647 == 65535)) {
                                    v18 := GenS3{F0: 112, F1: false, F2: int16(fi13)}
                                    _ = v18
                                }
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            default:
                var v19 int = 32768
                _ = v19
            }
        }
        p0 = (p0 || false)
        p0 = p0
    }
    v20 := []int{}
    if true {
        var v21 int = -51
        if len(v20) > 0 {
            v20[0] = v21
        }
        var v22 int16 = 1
        _ = v21
        _ = v22
    } else {
        if (-613 > 42) {
            v20 = append(v20, 300)
            p0 = p0
        }
    }
    _ = p0
    _ = v20
    return 214
}

func main() {
    v1 := []uint32{998, 287, 1, 2147483647}
    var v2 uint8 = 127
    v3 := []uint32{2147483647, 42, 78}
    v4 := make(map[int]bool)
    var v5 uint = 42
    fmt.Println(v5)
    v5 = (v5 + (v5 % (((v5 / (((v5 / ((704) | 1))) | 1))) | 1)))
    if (3.14159 == -2.25) {
        v4[4294967295] = false
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        if true {
            v5 += (v5 * v5)
            if (!false) {
                v5 -= (v5 % ((v5) | 1))
            }
            v5 = (v5 >> ((2) & 63))
        } else {
            fmt.Println(v2, v5, v2)
            switch "hello" {
            case "\n\t":
                v24 := GenS3{F0: v5, F1: true, F2: 32767}
                _ = v24
            case "":
                var v25 uint32 = 2147483648
                _ = v25
            }
            if ((1e10 * 0.0) <= 1e10) {
                var v26 int32 = 0
                _ = v26
            } else {
                v2 -= v2
                if (42 != 850) {
                    v5 -= (v5 % (((v5 & 42)) | 1))
                    v27 := GenS1{F0: (fi23 % ((fi23) | 1)), F1: v2}
                    for fi28 := 0; fi28 < 9; fi28++ {
                        var v29 uint64 = 18446744073709551615
                        v30 := GenS1{F0: (fi28 >> ((v5) & 63)), F1: 16}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = v27
                }
            }
        }
        var v31 float64 = 1e10
        _ = fi23
        _ = v31
    }
    fmt.Println(v5, v2)
    var v32 uint64 = 5
    if (!true) {
        for fi33 := 0; fi33 < 12; fi33++ {
            v34 := GenS4{F0: v32, F1: 32768}
            if (3.14159 < (float32(0.0) * 0.0)) {
                var v35 int64 = 153
                for fi36 := 0; fi36 < 8; fi36++ {
                    if len(v1) > 0 {
                        v1[0] = 58
                    }
                    switch "foo bar" {
                    case "":
                        switch ((("\n\t" + "foo bar") + ("foo bar" + "\n\t")) + "hello") {
                        case "a":
                            fmt.Println(v5, v5)
                        case "\n\t":
                            fmt.Println(v32, v32)
                        case "":
                            v5 -= v5
                        default:
                            var v37 uint16 = genFn0(((v2 * v2) >= 0))
                            _ = v37
                        }
                    default:
                        switch fi36 {
                        case -9223372036854775808:
                            v5 = v5
                        case -1:
                            var v38 uint = (^(v5))
                            _ = v38
                        default:
                            v35 -= (-(v35))
                        }
                    }
                    _ = fi36
                }
                _ = v35
            }
            _ = fi33
            _ = v34
        }
    } else {
        v39 := GenS2{F0: "\n\t", F1: 42}
        switch "\n\t" {
        case "\n\t":
            v40 := GenS1{F0: 4294967295, F1: (v2 % (((v2 + (v2 >> ((255) & 63)))) | 1))}
            _ = v40
        case "":
            for fi41 := 0; fi41 < 18; fi41++ {
                var v42 int = fi41
                _ = fi41
                _ = v42
            }
        case "foo bar":
            if true {
                v5 = 42
            } else {
                v3 = append(v3, 257)
                var v43 bool = false
                _ = v43
            }
        }
        fmt.Println(v5, v2)
        _ = v39
    }
    if (true || ((!false) && (172 < (v2 >> (((v5 ^ v5)) & 63))))) {
        v44 := GenS3{F0: v5, F1: false, F2: 1}
        var v45 int8 = 127
        _ = v44
        _ = v45
    } else {
        fmt.Println(v32, v5)
        if (42 < int32(v5)) {
            v46 := GenS5{F0: "hello"}
            v47 := GenS6{F0: 552}
            _ = v46
            _ = v47
        }
        v5 += (v5 % (((v5 * (v5 % (((v5 << ((v5) & 63))) | 1)))) | 1))
    }
    v48 := v4[-1]
    fmt.Println(v32)
    for fi49 := 0; fi49 < 8; fi49++ {
        if (255 >= genFn0(false)) {
            for fi50 := 0; fi50 < 12; fi50++ {
                v4[(fi49 & (fi50 << ((v5) & 63)))] = false
                if false {
                    _ = len(v4)
                } else {
                    var v51 int32 = -1
                    _ = v51
                }
                _ = fi50
            }
            var v52 float32 = (3.14159 / float32(1e10))
            _ = v52
        }
        var v53 bool = (v48 && (false && (((42 > 32767) && v48) && false)))
        if v53 {
            switch fi49 {
            case -719:
                v54 := GenS6{F0: 1}
                _ = v54
            case -28:
                v55 := GenS7{F0: 2147483647, F1: (fi49 % ((fi49) | 1)), F2: v2}
                _ = v55
            case 0:
                v48 = (!false)
            default:
                var v56 string = "a"
                _ = v56
            }
            for fi57 := 0; fi57 < 16; fi57++ {
                if false {
                    if v48 {
                        v3 = append(v3, 6)
                    } else {
                        switch fi57 {
                        case 32768:
                            fmt.Println(v32)
                        case 4294967295:
                            switch "hello" {
                            case "":
                                switch "hello" {
                                case "hello":
                                    v58 := GenS8{F0: v32}
                                    _ = v58
                                }
                            case "\n\t":
                                delete(v4, -189)
                            }
                        case -478:
                            var v59 int64 = 404
                            _ = v59
                        default:
                            v32 = (v32 % (((v32 >> (((v5 * v5)) & 63))) | 1))
                            v32 = 42
                        }
                    }
                }
                if false {
                    v2 = (^(v2))
                    v60 := v4[fi57]
                    v60 = (32767 == 65535)
                    _ = v60
                } else {
                    v48 = v48
                    var v61 string = "\n\t"
                    _ = v61
                }
                switch (fi49 * (fi49 % ((fi49) | 1))) {
                case 1:
                    v62 := GenS7{F0: -2147483648, F1: (fi57 & 32767), F2: v2}
                    if (-2.25 != 0.1) {
                        fmt.Println(fi49)
                        var v63 uint64 = (v32 | (v32 % ((uint64(v5)) | 1)))
                        switch ("" + "foo bar") {
                        case "hello":
                            if (!false) {
                                var v64 uint64 = v63
                                _ = v64
                            } else {
                                var v65 float32 = 1e10
                                var v66 string = "foo bar"
                                var v67 float32 = v65
                                _ = v65
                                _ = v66
                                _ = v67
                            }
                        case "foo bar":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                        }
                        _ = v63
                    } else {
                        var v68 int16 = 1000
                        _ = v68
                    }
                    if true {
                        fmt.Println(v53)
                    } else {
                        fmt.Println(v5)
                        if (v48 && v48) {
                            var v69 string = ("a" + "a")
                            switch v69 {
                            case "hello":
                                var v70 int = fi49
                                _ = v70
                            default:
                                var v71 uint32 = 2
                                _ = v71
                            }
                            fmt.Println(v48, v48, v53)
                            _ = v69
                        } else {
                            if (!(74 < 42)) {
                                v72 := GenS9{F0: 42, F1: 1, F2: 32767}
                                delete(v4, fi49)
                                _ = v72
                            }
                        }
                        var v73 uint32 = 0
                        _ = v73
                    }
                    _ = v62
                case 9223372036854775807:
                    fmt.Println(v32, v5)
                case -604:
                    v53 = (!v53)
                }
                _ = fi57
            }
        } else {
            for fi74 := 0; fi74 < 1; fi74++ {
                fmt.Println(v48)
                if v48 {
                    switch "hello" {
                    case "\n\t":
                        v75 := GenS10{F0: v53}
                        _ = v75
                    case "a":
                        var v76 uint = 0
                        _ = v76
                    default:
                        if (!v53) {
                            switch ("" + "hello") {
                            case "a":
                                switch (fi74 * fi74) {
                                case 4294967295:
                                    v5 = (v5 / (((v5 | (v5 * v5))) | 1))
                                }
                            case "hello":
                                v53 = false
                            case "foo bar":
                                v32 += (v32 | (^(v32)))
                            }
                        }
                    }
                    fmt.Println(v5, fi49, v32)
                }
                _ = fi74
            }
            v77 := v4[fi49]
            v78 := GenS11{F0: 564}
            _ = v77
            _ = v78
        }
        _ = fi49
        _ = v53
    }
    v79 := GenS9{F0: uint16(v2), F1: (v32 / (((^(v32))) | 1)), F2: -32768}
    v2 -= v2
    if (uint32(v5) > 125) {
        v48 = (110 <= 0)
    } else {
        var v80 uint64 = (v32 * 751)
        delete(v4, 110)
        v81 := GenS5{F0: "a"}
        _ = v80
        _ = v81
    }
    if (42 >= 0) {
        var v82 float64 = 1.5
        if true {
            fmt.Println(v82)
            v48 = (1e10 != 0.0)
            v83 := GenS1{F0: -874, F1: uint8(v5)}
            _ = v83
        } else {
            var v84 float64 = ((v82 * v82) * v82)
            fmt.Println(v32, v48)
            _ = v84
        }
        if ((42 <= -9223372036854775808) || ((v48 || v48) || true)) {
            switch ("a" + "") {
            case "\n\t":
                v32 += (v32 * (v32 + v32))
            case "":
                switch 127 {
                case 9223372036854775807:
                    var v85 int32 = -1
                    _ = v85
                case 404:
                    var v86 uint = (v5 ^ (v5 - (v5 * (v5 % ((v5) | 1)))))
                    _ = v86
                default:
                    var v87 int64 = int64(v32)
                    _ = v87
                }
            default:
                var v88 int8 = 1
                _ = v88
            }
            v89 := GenS10{F0: true}
            fmt.Println(v82, v32, v5)
            _ = v89
        }
        _ = v82
    } else {
        var v90 int16 = -1
        _ = v90
    }
    v91 := GenS10{F0: ((-548 > -1) && (v48 && (v5 < 65535)))}
    v92 := GenS4{F0: 42, F1: 128}
    for fi93 := 0; fi93 < 2; fi93++ {
        var v94 int16 = 1000
        _ = fi93
        _ = v94
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v48
    _ = v5
    _ = v79
    _ = v91
    _ = v92
}
