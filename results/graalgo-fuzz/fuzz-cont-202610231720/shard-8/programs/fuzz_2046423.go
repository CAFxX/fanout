package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint64; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 int32; F2 uint16 }
func genFn0(p0 uint8, p1 bool, p2 uint) uint {
    if (0 > -9223372036854775808) {
        var v7 float64 = float64((-2.25 / 1.0))
        p2 += 65535
        for fi8 := 0; fi8 < 2; fi8++ {
            if (-1 != 9) {
                if ((1 < 42) || false) {
                    var v9 uint32 = 2
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v7 = v7
                        if (p0 > (p0 * p0)) {
                            v11 := []int{fi8, fi10, fi10, fi8}
                            p1 = true
                            v11 = append(v11, -2147483648)
                            _ = v11
                        } else {
                            v12 := GenS1{F0: 247, F1: 2147483647}
                            if (100 < -92) {
                                p2 -= p2
                                v13 := []int{-2147483648, fi8, fi10}
                                fmt.Println(p2, p2)
                                _ = v13
                            }
                            fmt.Println(p0, v7, p1)
                            var v14 float64 = (float64((0.1 * float32(v7))) - v7)
                            if true {
                                var v15 uint16 = 32767
                                v16 := GenS1{F0: -650, F1: 2147483648}
                                var v17 int = (fi8 + 42)
                                _ = v15
                                _ = v16
                                _ = v17
                            }
                            _ = v12
                            _ = v14
                        }
                        _ = fi10
                    }
                    p2 += 42
                }
                var v18 uint64 = 1
                _ = v18
            } else {
                p2 = (p2 & (^(p2)))
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                switch "\n\t" {
                case "a":
                    var v20 uint8 = (p0 << (((p2 << (((p2 >> (((p2 * 2)) & 63))) & 63))) & 63))
                    _ = v20
                default:
                    var v21 uint64 = 2
                    _ = v21
                }
                fmt.Println(v7)
                var v22 int = (fi8 >> ((p2) & 63))
                v23 := []int{}
                _ = fi19
                _ = v22
                _ = v23
            }
            _ = fi8
        }
        _ = v7
    } else {
        switch "foo bar" {
        case "foo bar":
            if p1 {
                v24 := GenS1{F0: -1, F1: 2147483647}
                switch "hello" {
                case "":
                    p0 -= 2
                default:
                    p2 = (^(p2))
                }
                _ = v24
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    var v26 uint8 = 255
                    _ = fi25
                    _ = v26
                }
            }
        case "":
            p1 = p1
        case "\n\t":
            p0 += p0
        }
        if p1 {
            var v27 int32 = -1
            if ((("" + "hello") + ("\n\t" + "foo bar")) >= ("" + "hello")) {
                for fi28 := 0; fi28 < 14; fi28++ {
                    v29 := GenS2{F0: p2, F1: p2, F2: 836}
                    _ = fi28
                    _ = v29
                }
                fmt.Println(p0)
                for fi30 := 0; fi30 < 17; fi30++ {
                    v31 := GenS2{F0: p2, F1: p2, F2: 219}
                    switch (fi30 + (fi30 << (((p2 ^ (p2 % ((p2) | 1)))) & 63))) {
                    case 128:
                        var v32 uint32 = 2
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 float32 = 0.0
                            fmt.Println(fi30, v27)
                            _ = fi33
                            _ = v34
                        }
                    }
                    for fi35 := 0; fi35 < 9; fi35++ {
                        var v36 int8 = -1
                        fmt.Println(v36, p1)
                        var v37 bool = (!true)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                for fi38 := 0; fi38 < 2; fi38++ {
                    p2 += p2
                    _ = fi38
                }
                var v39 int32 = 1
                _ = v39
            }
            switch "\n\t" {
            case "\n\t":
                switch "" {
                case "\n\t":
                    for fi40 := 0; fi40 < 19; fi40++ {
                        v41 := []int{}
                        p1 = true
                        for fi42 := 0; fi42 < 18; fi42++ {
                            v43 := GenS3{F0: 81}
                            var v44 uint32 = 636
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        _ = fi40
                        _ = v41
                    }
                default:
                    for fi45 := 0; fi45 < 14; fi45++ {
                        v46 := []int{(fi45 - -1), fi45, -1, (fi45 & fi45)}
                        if p1 {
                            var v47 uint32 = 773
                            _ = v47
                        }
                        _ = fi45
                        _ = v46
                    }
                }
            }
            _ = v27
        }
        v48 := GenS4{F0: -32768}
        _ = v48
    }
    var v49 float32 = (float32(1.5) - (float32(0.0) / 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v49
    return (p2 % ((p2) | 1))
}
func genFn1(p0 int8, p1 int16) string {
    for fi50 := 0; fi50 < 15; fi50++ {
        var v51 uint16 = 2
        p1 -= (p1 >> ((518) & 63))
        if ((2 != 32767) && (false || false)) {
            v52 := []int{}
            v51 = 664
            var v53 int8 = -1
            _ = v52
            _ = v53
        }
        _ = fi50
        _ = v51
    }
    var v54 uint = 1
    v54 += genFn0(124, true, 284)
    _ = p0
    _ = p1
    _ = v54
    return ("foo bar" + ("" + (("a" + "a") + ("foo bar" + "foo bar"))))
}

func main() {
    v1 := []uint32{883}
    var v2 uint8 = 255
    v3 := []string{"\n\t", "", (("a" + "foo bar") + "")}
    v4 := []uint32{2147483647, 2147483647, 4294967295, 42}
    v5 := make(map[string]bool)
    var v6 int = 32767
    for fi55 := 0; fi55 < 17; fi55++ {
        if (("a" > ("a" + "\n\t")) || (("foo bar" + "a") > "\n\t")) {
            if (true && ((("" + "a") + "foo bar") >= "")) {
                v56 := GenS5{F0: 127, F1: 566}
                _ = v56
            } else {
                v57 := GenS4{F0: 831}
                var v58 int64 = 514
                _ = v57
                _ = v58
            }
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi55
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v3)
    case "hello":
        for fi59 := 0; fi59 < 20; fi59++ {
            var v60 int16 = 838
            fmt.Println(v60, v60, v2)
            v61 := GenS6{F0: v60, F1: 18446744073709551615, F2: 42}
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    v62 := GenS6{F0: -1, F1: 705, F2: 42}
    fmt.Println(v6)
    var v63 uint64 = 18446744073709551615
    for fi64 := 0; fi64 < 19; fi64++ {
        fmt.Println(v2, fi64)
        var v65 uint = 18446744073709551615
        v63 -= (^(v63))
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 12; fi66++ {
        fmt.Println(v6, v6)
        if len(v3) > 0 {
            v3[0] = ("hello" + "foo bar")
        }
        for fi67 := 0; fi67 < 1; fi67++ {
            fmt.Println(fi66, fi66, fi66)
            var v68 string = "a"
            _ = fi67
            _ = v68
        }
        _ = fi66
    }
    if false {
        delete(v5, (((("a" + "\n\t") + "hello") + "a") + "a"))
    }
    v63 += (v63 / (((v63 ^ v63)) | 1))
    switch "" {
    case "":
        v2 = 255
    }
    for fi69 := 0; fi69 < 7; fi69++ {
        v70 := v5["\n\t"]
        for fi71 := 0; fi71 < 6; fi71++ {
            _ = len(v5)
            _ = fi71
        }
        switch "hello" {
        case "a":
            if v70 {
                var v72 float32 = float32(((float64(0.0) + (1.5 + -2.25)) * 0.1))
                var v73 uint64 = v63
                _ = v72
                _ = v73
            }
        default:
            if len(v4) > 0 {
                v4[0] = 0
            }
        }
        _ = fi69
        _ = v70
    }
    switch "foo bar" {
    case "a":
        switch (v6 % (((^(v6))) | 1)) {
        case 32767:
            fmt.Println(v63, v6)
        }
    case "hello":
        switch (("hello" + ("\n\t" + "")) + ("\n\t" + ("" + ("a" + "")))) {
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "foo bar":
        v2 = 255
    }
    var v74 int8 = 5
    if false {
        _ = len(v5)
    }
    var v75 float64 = 3.14159
    if false {
        v76 := GenS3{F0: 2}
        if (true || ((false && (!(v63 <= (v63 >> ((255) & 63))))) || true)) {
            for fi77 := 0; fi77 < 18; fi77++ {
                v63 = 2
                _ = fi77
            }
            v2 = (v2 ^ v2)
            var v78 uint = 42
            _ = v78
        } else {
            v75 += v75
            v6 = (v6 & -1)
        }
        _ = v76
    } else {
        var v79 int32 = int32(v6)
        _ = v79
    }
    if true {
        if (int8(v6) != v74) {
            fmt.Println(v63, v2, v75)
        } else {
            var v80 uint16 = 65535
            if false {
                if false {
                    fmt.Println(v6, v6)
                }
                delete(v5, "a")
                if true {
                    var v81 int = int(v80)
                    _ = v81
                } else {
                    fmt.Println(v6)
                }
            } else {
                switch ("foo bar" + "a") {
                case "\n\t":
                    switch 480 {
                    case 2147483647:
                        var v82 uint64 = (v63 % ((v63) | 1))
                        _ = v82
                    case 32768:
                        for fi83 := 0; fi83 < 19; fi83++ {
                            v75 -= v75
                            v6 = (v6 + fi83)
                            _ = fi83
                        }
                    }
                case "foo bar":
                    _ = len(v3)
                case "":
                    v84 := GenS3{F0: (v80 & v80)}
                    _ = v84
                default:
                    var v85 uint = genFn0(v2, (("foo bar" == "a") && (v63 > uint64(v80))), 719)
                    _ = v85
                }
            }
            _ = v80
        }
        if len(v4) > 0 {
            v4[0] = 269
        }
        var v86 int = (-(v6))
        _ = v86
    } else {
        if true {
            v87 := GenS1{F0: 42, F1: 4294967295}
            v88 := GenS6{F0: -1, F1: v63, F2: 516}
            fmt.Println(v6, v6, v2)
            _ = v87
            _ = v88
        } else {
            v4 = append(v4, 4294967295)
            switch ("hello" + "foo bar") {
            case "a":
                v89 := GenS7{F0: "hello"}
                _ = v89
            default:
                if (!(636 > 511)) {
                    switch (v6 ^ v6) {
                    case -2147483648:
                        for fi90 := 0; fi90 < 20; fi90++ {
                            var v91 float64 = v75
                            v5[("hello" + "\n\t")] = false
                            _ = fi90
                            _ = v91
                        }
                    case 32768:
                        if (true || (!false)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = len(v5)
                            v75 = 0.0
                        } else {
                            fmt.Println(v75, v63)
                        }
                    }
                    for fi92 := 0; fi92 < 5; fi92++ {
                        var v93 string = ("\n\t" + "\n\t")
                        v94 := GenS5{F0: v74, F1: 2147483647}
                        delete(v5, "")
                        _ = fi92
                        _ = v93
                        _ = v94
                    }
                    v95 := GenS8{F0: true}
                    _ = v95
                } else {
                    if ((!(v6 != (v6 << ((42) & 63)))) || (((0 != v63) && false) && false)) {
                        v96 := GenS8{F0: (!(652 == 42))}
                        _ = v96
                    } else {
                        var v97 int = int(v63)
                        v98 := GenS6{F0: 118, F1: v63, F2: -1}
                        _ = v97
                        _ = v98
                    }
                    switch 225 {
                    case 9223372036854775807:
                        v6 -= 2147483647
                    case 17:
                        var v99 uint16 = 42
                        _ = v99
                    case 128:
                        v75 -= v75
                    default:
                        v100 := GenS9{F0: 65535, F1: -980, F2: 417}
                        _ = v100
                    }
                }
            }
            v101 := GenS4{F0: 32767}
            _ = v101
        }
        if false {
            var v102 int64 = -1
            _ = v102
        } else {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
        v6 += v6
    }
    v6 = (v6 | 4294967295)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v74
    _ = v75
}
