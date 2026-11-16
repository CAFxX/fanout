package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    var v3 int = 32767
    var v4 uint = 0
    delete(v1, "\n\t")
    v3 -= v3
    fmt.Println(v3)
    for fi5 := 0; fi5 < 12; fi5++ {
        if false {
            v6 := GenS1{F0: 255, F1: (("foo bar" + "\n\t") + ""), F2: 32767}
            if (true || (false || (873 > 2))) {
                var v7 uint32 = 377
                if false {
                    for fi8 := 0; fi8 < 20; fi8++ {
                        var v9 int64 = 311
                        switch "foo bar" {
                        case "":
                            var v10 string = "hello"
                            _ = v10
                        default:
                            var v11 int64 = v9
                            _ = v11
                        }
                        _ = fi8
                        _ = v9
                    }
                }
                if (true && false) {
                    v4 += (v4 & v4)
                    switch fi5 {
                    case 961:
                        v4 = v4
                    }
                    var v12 float32 = -2.25
                    _ = v12
                }
                _ = v7
            }
            for fi13 := 0; fi13 < 18; fi13++ {
                if true {
                    if (32767 >= -401) {
                        if (false && (true || false)) {
                            _ = len(v2)
                            v14 := GenS1{F0: 255, F1: ("\n\t" + "\n\t"), F2: 2}
                            fmt.Println(v4, v3)
                            if (343 >= 42) {
                                fmt.Println(v4, fi13, v3)
                                var v15 uint8 = 255
                                _ = v15
                            } else {
                                var v16 uint32 = 42
                                var v17 int16 = 967
                                _ = v16
                                _ = v17
                            }
                            _ = v14
                        }
                        var v18 float64 = (((3.14159 * -2.25) - (((-2.25 / 1.5) * (0.1 * (float64(0.0) - -2.25))) / (((1e10 * 1e10) - 1e10) / ((3.14159 / float64(-2.25)) - (0.0 + 0.1))))) * (((0.1 * (3.14159 - -2.25)) * 0.0) + -2.25))
                        fmt.Println(fi5, v18)
                        _ = v18
                    } else {
                        v4 = v4
                        v19 := GenS1{F0: 1, F1: "foo bar", F2: 867}
                        switch v3 {
                        case 32768:
                            v4 = v4
                        case 9223372036854775807:
                            _ = len(v2)
                        case 128:
                            fmt.Println(fi5, fi13, fi5)
                        }
                        delete(v1, ("foo bar" + "hello"))
                        v20 := GenS1{F0: 187, F1: "", F2: 0}
                        _ = v19
                        _ = v20
                    }
                    v21 := GenS2{F0: ""}
                    v1[("foo bar" + ("foo bar" + ""))] = (v3 != fi5)
                    _ = v21
                } else {
                    _ = len(v2)
                }
                if (-895 >= -486) {
                    v4 += 955
                    var v22 bool = true
                    _ = v22
                } else {
                    v23 := GenS1{F0: 42, F1: ("\n\t" + "hello"), F2: 32767}
                    switch 909 {
                    case -2147483648:
                        var v24 int64 = -9223372036854775808
                        _ = v24
                    case 9223372036854775807:
                        var v25 int = v3
                        _ = v25
                    case 32767:
                        var v26 uint64 = 4294967295
                        _ = v26
                    default:
                        v27 := v1["\n\t"]
                        _ = v27
                    }
                    _ = v23
                }
                v2 = append(v2, fi13)
                if false {
                    if (0.1 <= (((float32(float64((1e10 / -2.25))) - (3.14159 - 1.5)) * float32(3.14159)) * 1.5)) {
                        fmt.Println(v4, v3, fi5)
                        switch (fi13 >> ((v4) & 63)) {
                        case 128:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case -1:
                            v3 = v3
                        default:
                            _ = len(v2)
                        }
                        fmt.Println(fi5, fi5)
                    }
                    v4 -= v4
                }
                _ = fi13
            }
            _ = v6
        }
        switch ("hello" + "") {
        case "a":
            fmt.Println(v3, v4, v4)
        default:
            switch (^(v3)) {
            case 9223372036854775807:
                for fi28 := 0; fi28 < 17; fi28++ {
                    if (0 >= -5) {
                        v4 += 303
                        fmt.Println(v3)
                    }
                    if false {
                        switch "hello" {
                        case "a":
                            fmt.Println(fi28)
                        default:
                            v3 = v3
                        }
                    }
                    v4 = 18446744073709551615
                    var v29 int32 = 829
                    _ = fi28
                    _ = v29
                }
            }
        }
        _ = fi5
    }
    if len(v2) > 0 {
        v2[0] = v3
    }
    fmt.Println(v4)
    v4 = (v4 ^ v4)
    fmt.Println(v4)
    switch "a" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = (v3 % (((v3 | v3)) | 1))
        }
    case "a":
        fmt.Println(v3, v3)
    default:
        v30 := GenS1{F0: 0, F1: (("a" + "a") + "foo bar"), F2: 687}
        _ = v30
    }
    var v31 uint = 18446744073709551615
    if true {
        if true {
            v32 := v1["foo bar"]
            if (-2147483648 >= 836) {
                for fi33 := 0; fi33 < 9; fi33++ {
                    _ = len(v2)
                    for fi34 := 0; fi34 < 13; fi34++ {
                        switch (("\n\t" + "") + "\n\t") {
                        case "\n\t":
                            var v35 uint64 = 85
                            v36 := GenS3{F0: 610, F1: 32767}
                            switch (fi33 << ((v31) & 63)) {
                            case -9223372036854775808:
                                fmt.Println(v32, v3, v35)
                            case 4294967295:
                                if ((((!(!((v32 || (v32 && (v31 < (v4 - (v4 ^ v31))))) || false))) || false) && false) && ((true && (!v32)) || ((fi34 * 32767) == (fi33 + (v3 | (fi33 * 0)))))) {
                                    if true {
                                        fmt.Println(v32, fi33)
                                        v37 := GenS1{F0: 1, F1: "", F2: 354}
                                        _ = v37
                                    }
                                    v3 -= (^(v3))
                                    if true {
                                        v4 -= v4
                                        v3 += (fi34 + (v3 - fi33))
                                    } else {
                                        if (v32 && v32) {
                                            v31 = v4
                                            v3 -= (fi34 | -9223372036854775808)
                                            delete(v1, "")
                                        } else {
                                            var v38 uint32 = 991
                                            if v32 {
                                                _ = len(v2)
                                            } else {
                                                v31 = (^(v4))
                                                var v39 uint64 = v35
                                                v40 := GenS3{F0: 768, F1: uint16(v4)}
                                                _ = v39
                                                _ = v40
                                            }
                                            fmt.Println(v31, v38, v32)
                                            _ = v38
                                        }
                                        v31 = v31
                                        v32 = false
                                    }
                                }
                            }
                            _ = v35
                            _ = v36
                        case "hello":
                            fmt.Println(v4, v32)
                        }
                        v4 -= v31
                        _ = fi34
                    }
                    switch ("" + ("\n\t" + (((("" + "foo bar") + (("foo bar" + "\n\t") + "\n\t")) + "foo bar") + ""))) {
                    case "foo bar":
                        var v41 uint32 = 2
                        _ = v41
                    case "hello":
                        v42 := GenS4{F0: 42, F1: 758, F2: 2}
                        _ = v42
                    }
                    _ = fi33
                }
            } else {
                var v43 uint16 = 486
                _ = v43
            }
            _ = v32
        } else {
            v4 += v4
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch "hello" {
        case "\n\t":
            var v44 uint64 = 2
            _ = v44
        case "a":
            v45 := GenS4{F0: 42, F1: 42, F2: 826}
            _ = v45
        case "hello":
            var v46 uint = v31
            _ = v46
        }
    }
    var v47 int64 = int64(v3)
    v31 += 0
    v48 := GenS5{F0: (v31 / (((^(v4))) | 1))}
    v3 -= (v3 & (-(v3)))
    if ((42 <= 18446744073709551615) || true) {
        var v49 int64 = (v47 + v47)
        v31 = (^(v4))
        v2 = append(v2, v3)
        _ = v49
    } else {
        v47 = -1
        v50 := GenS6{F0: 553}
        if false {
            v47 = v47
            if false {
                if false {
                    var v51 uint8 = 128
                    if ("" <= "foo bar") {
                        v51 += 2
                    }
                    _ = v51
                } else {
                    var v52 uint8 = 127
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v52
                }
                v2 = append(v2, v3)
            } else {
                v3 -= v3
                if (false || true) {
                    if true {
                        v3 += v3
                    }
                    v2 = append(v2, v3)
                    var v53 float64 = 1.5
                    _ = v53
                } else {
                    if true {
                        var v54 uint32 = 2
                        var v55 uint8 = 51
                        var v56 uint64 = 819
                        _ = v54
                        _ = v55
                        _ = v56
                    } else {
                        var v57 int = 119
                        _ = v57
                    }
                    v58 := GenS7{F0: 4294967295, F1: 32767}
                    _ = v58
                }
            }
        }
        _ = v50
    }
    v4 = 42
    fmt.Println(v4, v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v47
    _ = v48
}
