package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 int }
type GenS6 struct { F0 int32; F1 uint32; F2 int }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int64; F1 uint; F2 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 bool; F1 int; F2 int }
func genFn0(p0 int64) int {
    v6 := GenS1{F0: (p0 & p0)}
    fmt.Println(p0)
    var v7 uint16 = 2
    if true {
        switch "hello" {
        case "":
            v7 = (^(v7))
        case "\n\t":
            fmt.Println(v7, v7)
        case "a":
            v7 -= v7
        default:
            if (!(("foo bar" + ("\n\t" + "a")) <= "")) {
                fmt.Println(v7)
                fmt.Println(p0, v7)
                if (-2.25 <= 1e10) {
                    if (!false) {
                        var v8 float64 = 3.14159
                        v9 := GenS1{F0: 0}
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v7 = 255
                            var v11 uint = 2
                            _ = fi10
                            _ = v11
                        }
                        _ = v8
                        _ = v9
                    } else {
                        var v12 uint8 = 95
                        v12 = (v12 >> ((121) & 63))
                        _ = v12
                    }
                }
            } else {
                v13 := GenS1{F0: (p0 ^ (-(p0)))}
                p0 += (p0 - (p0 - p0))
                switch "a" {
                case "foo bar":
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v15 := []int{1, -728}
                        for fi16 := 0; fi16 < 3; fi16++ {
                            var v17 uint16 = v7
                            var v18 uint16 = v7
                            var v19 uint = 190
                            _ = fi16
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        if false {
                            var v20 uint = 40
                            _ = v20
                        } else {
                            v7 += 136
                        }
                        _ = fi14
                        _ = v15
                    }
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        v21 := []int{}
                        _ = v21
                    case "\n\t":
                        if ((-2.25 / float64(1.5)) >= 0.0) {
                            fmt.Println(v7)
                            fmt.Println(p0)
                        }
                    case "a":
                        var v22 int64 = p0
                        _ = v22
                    default:
                        v23 := GenS1{F0: p0}
                        _ = v23
                    }
                case "":
                    switch -1 {
                    case -2147483648:
                        p0 = 0
                    case -839:
                        var v24 uint32 = 2147483647
                        _ = v24
                    }
                }
                _ = v13
            }
        }
        var v25 uint8 = 255
        _ = v25
    }
    v26 := GenS2{F0: 42, F1: 4294967295}
    _ = p0
    _ = v26
    _ = v6
    _ = v7
    return 967
}
func genFn1(p0 int8) bool {
    v27 := GenS1{F0: 9223372036854775807}
    if false {
        p0 -= p0
        v28 := GenS1{F0: -9223372036854775808}
        _ = v28
    }
    p0 += (p0 >> ((320) & 63))
    v29 := []int{127}
    v30 := GenS3{F0: 697, F1: 529}
    _ = p0
    _ = v27
    _ = v29
    _ = v30
    return ((541 > 462) && false)
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]uint32)
    var v3 int16 = 885
    v4 := make(map[string]bool)
    var v5 uint16 = 65535
    v5 += v5
    v31 := []int{}
    switch genFn0(0) {
    case 0:
        switch genFn0(0) {
        case 42:
            switch 128 {
            case 32768:
                var v32 string = ("a" + ((("" + "") + ("foo bar" + "\n\t")) + ""))
                _ = v32
            case 9223372036854775807:
                if false {
                    fmt.Println(v5)
                    for fi33 := 0; fi33 < 12; fi33++ {
                        switch "foo bar" {
                        case "\n\t":
                            v34 := GenS4{F0: (v1 & (v1 << ((v1) & 63))), F1: 0}
                            _ = v34
                        default:
                            if (1 > uint32(v5)) {
                                var v35 float64 = 0.0
                                var v36 uint8 = 223
                                v3 += (v3 % ((1000) | 1))
                                _ = v35
                                _ = v36
                            }
                        }
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        _ = fi33
                    }
                    _ = len(v31)
                } else {
                    var v37 float64 = 0.0
                    v38 := GenS5{F0: -1, F1: "foo bar", F2: 1}
                    v37 += (v37 / 1.0)
                    _ = v37
                    _ = v38
                }
            default:
                if (51 == 127) {
                    switch 32767 {
                    case -2147483648:
                        var v39 string = "hello"
                        _ = v39
                    case 32767:
                        var v40 uint8 = 1
                        _ = v40
                    default:
                        for fi41 := 0; fi41 < 20; fi41++ {
                            v42 := GenS1{F0: 1}
                            if true {
                                v5 = v5
                                if len(v31) > 0 {
                                    v31[0] = genFn0(761)
                                }
                            } else {
                                fmt.Println(v1, v3, fi41)
                            }
                            _ = fi41
                            _ = v42
                        }
                    }
                }
            }
        }
    case -9223372036854775808:
        fmt.Println(v5, v5)
    default:
        var v43 int16 = (v3 & 1)
        _ = v43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        for fi45 := 0; fi45 < 12; fi45++ {
            fmt.Println(fi44)
            v46 := GenS6{F0: 1, F1: 1, F2: 127}
            _ = len(v2)
            _ = fi45
            _ = v46
        }
        if len(v31) > 0 {
            v31[0] = (fi44 % ((32767) | 1))
        }
        var v47 bool = true
        _ = fi44
        _ = v47
    }
    switch "a" {
    case "":
        switch 436 {
        case 2147483647:
            var v48 float64 = 1e10
            _ = v48
        }
    }
    v1 = (v1 * (v1 >> ((v1) & 63)))
    v5 = v5
    if (42 < 4294967295) {
        if (!true) {
            v49 := GenS5{F0: 127, F1: ("a" + ("\n\t" + ((("hello" + "\n\t") + "hello") + "a"))), F2: 42}
            var v50 uint64 = 575
            v51 := GenS5{F0: 0, F1: ("a" + "foo bar"), F2: -9223372036854775808}
            _ = v49
            _ = v50
            _ = v51
        }
        var v52 string = "hello"
        _ = v52
    }
    v1 = v1
    if true {
        v53 := GenS7{F0: 200, F1: -2147483648}
        var v54 uint16 = uint16(v1)
        for fi55 := 0; fi55 < 10; fi55++ {
            if ((v3 * (v3 >> ((338) & 63))) > -1) {
                var v56 uint16 = (^(v5))
                v57 := v4[""]
                _ = v56
                _ = v57
            } else {
                for fi58 := 0; fi58 < 3; fi58++ {
                    if (!false) {
                        v59 := GenS7{F0: 255, F1: -1}
                        var v60 int8 = 0
                        switch fi55 {
                        case 32767:
                            v3 -= v3
                        case 79:
                            v61 := GenS8{F0: 42}
                            _ = v61
                        }
                        v60 = (^(v60))
                        _ = v59
                        _ = v60
                    }
                    var v62 bool = (false && ((v54 ^ (v5 >> ((v1) & 63))) < v54))
                    _ = fi58
                    _ = v62
                }
                switch fi55 {
                case 0:
                    v3 -= 528
                default:
                    v63 := GenS1{F0: 9223372036854775807}
                    _ = v63
                }
                v1 = v1
            }
            _ = fi55
        }
        _ = v53
        _ = v54
    }
    fmt.Println(v5, v1)
    for fi64 := 0; fi64 < 7; fi64++ {
        v2["a"] = 665
        v65 := GenS9{F0: -1, F1: 502, F2: -9223372036854775808}
        switch "" {
        case "":
            var v66 bool = ((0 > 573) || false)
            _ = v66
        case "foo bar":
            v67 := GenS10{F0: uint16(v1)}
            _ = v67
        }
        _ = fi64
        _ = v65
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        fmt.Println(fi68)
        v5 = (^(v5))
        switch 32768 {
        case 127:
            v1 = (v1 | v1)
        case 0:
            for fi69 := 0; fi69 < 14; fi69++ {
                var v70 int32 = 2147483647
                switch "" {
                case "foo bar":
                    v5 -= v5
                default:
                    var v71 int64 = -530
                    _ = v71
                }
                _ = len(v31)
                _ = fi69
                _ = v70
            }
        case 2147483647:
            v5 -= (^(v5))
        }
        _ = fi68
    }
    var v72 string = (("hello" + "\n\t") + ("foo bar" + ""))
    var v73 uint = v1
    v74 := GenS11{F0: true, F1: genFn0(-9223372036854775808), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
}
