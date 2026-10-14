package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint; F1 int; F2 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 0.1
        _ = fi7
        _ = v8
    }
    var v9 int16 = 1000
    if (!(!(true && false))) {
        v9 += 477
        var v10 string = "\n\t"
        v11 := GenS1{F0: 192, F1: 255}
        _ = v10
        _ = v11
    }
    _ = v9
    return 255
}
func genFn1(p0 int16, p1 uint64) int32 {
    switch 4294967295 {
    case -2147483648:
        if (!(false || (!(4294967295 > p1)))) {
            p1 = (p1 >> ((4294967295) & 63))
            var v12 float32 = 1.5
            if ("\n\t" < ("" + "")) {
                var v13 uint16 = 42
                var v14 int16 = (p0 / ((p0) | 1))
                switch ("a" + ("\n\t" + "")) {
                case "foo bar":
                    var v15 float64 = (((0.0 - 0.0) - (0.1 / 1e10)) - 3.14159)
                    _ = v15
                case "a":
                    var v16 int16 = (p0 / ((32767) | 1))
                    _ = v16
                case "":
                    v13 = 2
                }
                _ = v13
                _ = v14
            }
            _ = v12
        } else {
            v17 := GenS2{F0: "a", F1: 2147483647}
            _ = v17
        }
    case 1:
        var v18 uint = 2
        _ = v18
    default:
        v19 := []int{32767}
        _ = v19
    }
    v20 := []int{707, -2147483648, 127}
    for fi21 := 0; fi21 < 13; fi21++ {
        if (2147483648 != 2) {
            for fi22 := 0; fi22 < 10; fi22++ {
                p0 = p0
                _ = fi22
            }
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 uint16 = 429
                var v25 uint64 = (p1 << ((0) & 63))
                if (!((!false) || false)) {
                    var v26 uint64 = 42
                    var v27 float64 = 1.5
                    switch (fi23 % (((fi21 | int(p0))) | 1)) {
                    case -1:
                        if (int8(v24) < 79) {
                            if ((!(true && (1 <= 0))) && true) {
                                switch "a" {
                                case "a":
                                    var v28 string = ("foo bar" + ("foo bar" + ("foo bar" + "")))
                                    _ = v28
                                case "":
                                    if len(v20) > 0 {
                                        v20[0] = fi21
                                    }
                                case "hello":
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                    v29 := GenS3{F0: p0, F1: 65535, F2: 43}
                                    _ = len(v20)
                                    _ = v29
                                }
                                fmt.Println(fi21)
                                p0 += 574
                                p0 = p0
                            }
                            v25 = v26
                        } else {
                            fmt.Println(v25)
                            v24 = (v24 / ((331) | 1))
                        }
                    default:
                        var v30 int = -447
                        _ = v30
                    }
                    _ = v26
                    _ = v27
                }
                fmt.Println(v25, fi21, p1)
                v31 := GenS4{F0: 1, F1: 127, F2: -936}
                var v32 int8 = -1
                _ = fi23
                _ = v24
                _ = v25
                _ = v31
                _ = v32
            }
            switch "a" {
            case "hello":
                if true {
                    v33 := GenS5{F0: 4294967295}
                    v20 = append(v20, fi21)
                    var v34 int32 = 2147483647
                    _ = v33
                    _ = v34
                }
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    var v35 float64 = 0.1
                    _ = v35
                case "hello":
                    v36 := GenS5{F0: fi21}
                    _ = v36
                case "\n\t":
                    p1 += (p1 - (p1 & p1))
                }
            default:
                if ((((42 != 9223372036854775807) || (p1 == (p1 % (((p1 / (((p1 << ((42) & 63))) | 1))) | 1)))) && false) || (false && false)) {
                    p0 += p0
                    for fi37 := 0; fi37 < 17; fi37++ {
                        var v38 bool = (1 != 42)
                        var v39 uint8 = 0
                        v20 = append(v20, (^(fi37)))
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
            }
        }
        var v40 string = "a"
        _ = fi21
        _ = v40
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = v20
    return -1
}

func main() {
    v1 := []string{}
    var v2 uint8 = 42
    v3 := []string{"a", "hello", ("a" + "foo bar")}
    v4 := []uint32{2147483648, 2147483647, 109}
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    if false {
        v2 = (v2 % ((v2) | 1))
        if false {
            var v41 int16 = 78
            if false {
                _ = len(v6)
            }
            fmt.Println(v41, v2)
            _ = v41
        }
    }
    v2 -= 27
    fmt.Println(v2)
    v2 = (v2 >> ((18446744073709551615) & 63))
    var v42 int = 32767
    var v43 string = ("\n\t" + ("\n\t" + (("foo bar" + ("a" + ("foo bar" + ("a" + "hello")))) + ("\n\t" + "a"))))
    switch (("a" + ((("\n\t" + (v43 + "foo bar")) + "\n\t") + ("foo bar" + "hello"))) + ("" + (("a" + ((v43 + v43) + "a")) + (v43 + ((v43 + "\n\t") + v43))))) {
    case "\n\t":
        if false {
            v44 := GenS3{F0: 0, F1: 65535, F2: (v2 - (v2 ^ (^(v2))))}
            _ = v44
        }
    case "":
        var v45 uint32 = 2
        _ = v45
    }
    var v46 bool = true
    var v47 string = v43
    if len(v1) > 0 {
        v1[0] = (v47 + "hello")
    }
    v2 = v2
    for fi48 := 0; fi48 < 10; fi48++ {
        if (!(fi48 <= fi48)) {
            if (!false) {
                fmt.Println(v43, v47)
                if len(v1) > 0 {
                    v1[0] = v47
                }
                var v49 float64 = 1.5
                _ = v49
            } else {
                for fi50 := 0; fi50 < 3; fi50++ {
                    var v51 uint = 239
                    switch ((v43 + v43) + v47) {
                    case "a":
                        fmt.Println(v42, v46)
                    case "\n\t":
                        v47 = ("foo bar" + "\n\t")
                    }
                    var v52 uint32 = 2147483648
                    var v53 uint = v51
                    _ = fi50
                    _ = v51
                    _ = v52
                    _ = v53
                }
                switch v43 {
                case "hello":
                    v54 := GenS2{F0: "\n\t", F1: v42}
                    _ = v54
                case "a":
                    for fi55 := 0; fi55 < 6; fi55++ {
                        var v56 uint64 = 0
                        _ = fi55
                        _ = v56
                    }
                case "foo bar":
                    if v46 {
                        var v57 uint64 = 1
                        _ = v57
                    } else {
                        var v58 uint8 = v2
                        switch v47 {
                        case "":
                            fmt.Println(v46, v58, v2)
                        case "a":
                            var v59 uint32 = 2
                            _ = v59
                        case "\n\t":
                            var v60 float32 = float32(float64(0.1))
                            _ = v60
                        default:
                            v46 = (v46 && v46)
                        }
                        for fi61 := 0; fi61 < 8; fi61++ {
                            v6[v43] = 20
                            _ = fi61
                        }
                        _ = v58
                    }
                }
            }
        } else {
            if (0.0 < (0.1 - 1e10)) {
                switch v43 {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    var v62 uint32 = 491
                    _ = v62
                default:
                    v63 := v6[v47]
                    _ = v63
                }
                _ = len(v3)
                switch ("" + "foo bar") {
                case "":
                    switch v42 {
                    case 729:
                        v42 += (v42 * fi48)
                    case 9223372036854775807:
                        fmt.Println(fi48, v2, v47)
                    }
                case "foo bar":
                    var v64 uint = 65535
                    _ = v64
                case "hello":
                    v65 := GenS2{F0: v43, F1: (fi48 << ((0) & 63))}
                    _ = v65
                }
            } else {
                v1 = append(v1, v43)
                switch "\n\t" {
                case "hello":
                    var v66 string = "hello"
                    _ = v66
                }
                for fi67 := 0; fi67 < 8; fi67++ {
                    if (!((v46 && v46) && v46)) {
                        v68 := GenS1{F0: (v2 & v2), F1: v2}
                        _ = v68
                    }
                    _ = fi67
                }
            }
            if true {
                v42 += fi48
                if ((^(v2)) >= (v2 ^ 42)) {
                    v2 += (v2 % ((v2) | 1))
                    if (0.0 != (1e10 * 1.5)) {
                        v69 := GenS6{F0: 629}
                        var v70 int16 = 1
                        for fi71 := 0; fi71 < 15; fi71++ {
                            fmt.Println(v2)
                            var v72 int64 = 9223372036854775807
                            if (!false) {
                                var v73 uint64 = 978
                                _ = v73
                            } else {
                                v46 = (!(!(float64(1e10) == 3.14159)))
                            }
                            _ = fi71
                            _ = v72
                        }
                        _ = v69
                        _ = v70
                    } else {
                        v43 = "foo bar"
                        var v74 uint8 = 249
                        _ = v74
                    }
                    fmt.Println(v43, v47, v42)
                }
            }
        }
        _ = fi48
    }
    v75 := v6["hello"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v75
}
