package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint64; F2 string }
type GenS6 struct { F0 int32; F1 int32; F2 int16 }
type GenS7 struct { F0 uint; F1 uint64 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    var v3 bool = (1 <= 100)
    v4 := make(map[string]int)
    var v5 int16 = 42
    if false {
        v5 = (v5 >> ((42) & 63))
        switch -899 {
        case 42:
            v6 := GenS1{F0: -1, F1: 1, F2: 127}
            _ = v6
        default:
            var v7 uint = 65535
            _ = v7
        }
    } else {
        fmt.Println(v5)
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        switch 4294967295 {
        case 265:
            v3 = true
        }
        _ = fi8
    }
    switch ("a" + "") {
    case "":
        if ("\n\t" != "foo bar") {
            fmt.Println(v3, v3)
            v9 := GenS2{F0: 1, F1: 233, F2: (("foo bar" + (("a" + ("" + ("\n\t" + "foo bar"))) + ("a" + "hello"))) + "\n\t")}
            v3 = true
            _ = v9
        } else {
            v10 := []int{}
            v5 = v5
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := v2["a"]
                var v13 int32 = -79
                v10 = append(v10, (fi11 << ((749) & 63)))
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
    case "hello":
        delete(v2, ("hello" + ("a" + "foo bar")))
    case "a":
        if v3 {
            var v14 uint16 = 255
            v15 := GenS1{F0: -807, F1: -551, F2: 4294967295}
            if v3 {
                v16 := GenS3{F0: -9223372036854775808}
                _ = v16
            }
            _ = v14
            _ = v15
        }
    }
    switch 32768 {
    case -784:
        var v17 bool = true
        _ = v17
    case 9223372036854775807:
        switch 1 {
        case 1:
            v5 = v5
        case -2147483648:
            v1 += (0.0 / (float64((3.14159 + 1e10)) / v1))
        }
    default:
        var v18 uint64 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := []int{(fi19 / (((fi19 + 1)) | 1))}
        var v21 string = ""
        switch 2147483647 {
        case 753:
            v1 += (v1 / (v1 + (v1 + v1)))
        default:
            v22 := GenS1{F0: -401, F1: -9223372036854775808, F2: (fi19 % ((fi19) | 1))}
            _ = v22
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        v1 = 1e10
        _ = fi23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 14; fi25++ {
            if v3 {
                var v26 bool = (127 >= 126)
                _ = len(v4)
                v5 -= 42
                _ = v26
            }
            var v27 float32 = 3.14159
            _ = fi25
            _ = v27
        }
        if true {
            v5 = 42
            var v28 uint64 = 18446744073709551615
            _ = v28
        } else {
            var v29 uint16 = 65535
            v29 -= v29
            var v30 string = ""
            _ = v29
            _ = v30
        }
        var v31 uint64 = 0
        _ = fi24
        _ = v31
    }
    fmt.Println(v1, v3)
    v32 := GenS2{F0: 923, F1: 1, F2: "foo bar"}
    for fi33 := 0; fi33 < 16; fi33++ {
        v3 = (v3 || (2 != 4294967295))
        fmt.Println(v3)
        _ = fi33
    }
    switch 127 {
    case 1:
        var v34 int = -449
        _ = v34
    case 42:
        for fi35 := 0; fi35 < 13; fi35++ {
            if (2147483647 != 2147483647) {
                v36 := GenS3{F0: 588}
                if (128 > 241) {
                    switch "hello" {
                    case "":
                        v5 = 0
                    case "a":
                        switch "foo bar" {
                        case "\n\t":
                            v37 := []int{2147483647}
                            _ = v37
                        }
                    case "foo bar":
                        for fi38 := 0; fi38 < 13; fi38++ {
                            if (v3 || v3) {
                                if ((899 == 0) && (608 != 2147483648)) {
                                    if (!v3) {
                                        if (v3 || (2147483648 > 42)) {
                                            switch "\n\t" {
                                            case "a":
                                                if (42 >= 405) {
                                                    var v39 int32 = -2147483648
                                                    v40 := []int{7, (fi38 % ((fi38) | 1)), 2147483647, (fi35 ^ int(v5))}
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    v5 = (v5 ^ v5)
                                                }
                                            }
                                            v1 = (v1 * (float64((0.1 - -2.25)) - v1))
                                        }
                                        if v3 {
                                            fmt.Println(fi38, v3, v3)
                                            v1 = v1
                                        }
                                    } else {
                                        v3 = true
                                        _ = len(v4)
                                    }
                                } else {
                                    v3 = false
                                    var v41 float64 = (v1 / float64(((0.1 - float32(((0.1 - (v1 - (1.5 + float64(1e10)))) / (v1 / v1)))) * 3.14159)))
                                    var v42 bool = ((fi38 >> ((2) & 63)) == 714)
                                    _ = v41
                                    _ = v42
                                }
                                if v3 {
                                    var v43 uint64 = 2
                                    _ = v43
                                }
                                var v44 int64 = -1
                                _ = v44
                            } else {
                                v45 := GenS1{F0: -442, F1: 9223372036854775807, F2: fi38}
                                _ = v45
                            }
                            _ = fi38
                        }
                    default:
                        if true {
                            if v3 {
                                var v46 string = "a"
                                fmt.Println(v3, v1)
                                _ = v46
                            } else {
                                fmt.Println(v1, fi35, v1)
                            }
                            v5 = v5
                            for fi47 := 0; fi47 < 10; fi47++ {
                                fmt.Println(fi47)
                                _ = fi47
                            }
                        } else {
                            v3 = (0 == 2147483648)
                        }
                    }
                } else {
                    if (v3 && (!(255 <= 255))) {
                        for fi48 := 0; fi48 < 20; fi48++ {
                            _ = len(v4)
                            fmt.Println(fi48, fi48, fi35)
                            v3 = v3
                            _ = fi48
                        }
                        v1 -= 3.14159
                    } else {
                        switch "hello" {
                        case "":
                            for fi49 := 0; fi49 < 9; fi49++ {
                                if (65 < 405) {
                                    var v50 uint64 = 0
                                    v51 := GenS4{F0: (false && v3)}
                                    var v52 float32 = -2.25
                                    _ = len(v2)
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                }
                                v3 = true
                                if true {
                                    v5 = 561
                                }
                                v1 = ((((float64((0.1 + ((-2.25 + -2.25) * (1.5 - 0.0)))) + (v1 + v1)) * v1) / (v1 * 0.1)) - float64(((float32(float64(1e10)) * 1.5) / 0.1)))
                                _ = fi49
                            }
                        case "a":
                            v5 = -264
                        case "hello":
                            if (!v3) {
                                v53 := GenS3{F0: 0}
                                _ = v53
                            }
                        }
                    }
                    if false {
                        var v54 int64 = -611
                        if (float32((v1 / 1.0)) < ((0.0 + 1.5) / 1.5)) {
                            var v55 uint = 18446744073709551615
                            _ = v55
                        }
                        var v56 bool = (true && ((-2.25 == v1) && v3))
                        _ = v54
                        _ = v56
                    } else {
                        v4[""] = fi35
                    }
                }
                _ = v36
            }
            _ = fi35
        }
    default:
        v57 := []int{-976, 710, 32768}
        _ = v57
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        v59 := GenS3{F0: -9223372036854775808}
        _ = fi58
        _ = v59
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        if (0 == 0) {
            for fi60 := 0; fi60 < 19; fi60++ {
                if v3 {
                    fmt.Println(v5, v1, fi60)
                    var v61 uint32 = 4294967295
                    _ = v61
                }
                if true {
                    if ((("foo bar" + "a") + "a") != "foo bar") {
                        v62 := GenS5{F0: "\n\t", F1: 1, F2: "a"}
                        _ = v62
                    } else {
                        fmt.Println(fi60, v5)
                        v63 := GenS1{F0: 0, F1: 9223372036854775807, F2: (fi60 & fi60)}
                        _ = v63
                    }
                    v64 := GenS6{F0: 0, F1: 42, F2: 228}
                    _ = v64
                } else {
                    fmt.Println(v5, fi60, v1)
                    v65 := []int{(fi60 * 4294967295)}
                    _ = v65
                }
                _ = fi60
            }
            var v66 uint64 = 2
            if v3 {
                if true {
                    v67 := GenS7{F0: 629, F1: (^(v66))}
                    _ = v67
                }
                v68 := GenS6{F0: -115, F1: -1, F2: (v5 - -32768)}
                _ = v68
            }
            _ = v66
        }
    case "foo bar":
        v69 := GenS5{F0: "a", F1: 910, F2: ""}
        _ = v69
    case "":
        if v3 {
            fmt.Println(v3, v5, v1)
            v70 := []int{-1, 4294967295}
            for fi71 := 0; fi71 < 1; fi71++ {
                if v3 {
                    if (-1 == 1) {
                        v3 = (3.14159 == (v1 - ((1e10 + (v1 * v1)) + 1e10)))
                        var v72 int = (-(fi71))
                        if (int8(v5) >= 42) {
                            v72 -= v72
                            for fi73 := 0; fi73 < 8; fi73++ {
                                v74 := GenS7{F0: 922, F1: 35}
                                var v75 uint32 = 478
                                fmt.Println(fi71)
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                            v3 = (!(119 > 828))
                        }
                        _ = v72
                    } else {
                        if len(v70) > 0 {
                            v70[0] = 127
                        }
                    }
                } else {
                    fmt.Println(v1, v5)
                    var v76 int64 = -9223372036854775808
                    v3 = (v1 == ((v1 - 1.5) - 0.0))
                    _ = v76
                }
                for fi77 := 0; fi77 < 6; fi77++ {
                    switch (fi77 | (fi77 << ((42) & 63))) {
                    case -9223372036854775808:
                        if true {
                            v78 := GenS8{F0: 65535}
                            var v79 float64 = 1e10
                            v5 = (v5 / (((v5 >> ((42) & 63))) | 1))
                            v79 = (((1.5 - 0.1) / 1e10) / 3.14159)
                            v80 := v2["foo bar"]
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            v5 = (v5 * (v5 - v5))
                        }
                    default:
                        if (42 > 42) {
                            v3 = v3
                            if (0 < 1) {
                                v4[("hello" + "foo bar")] = fi71
                                var v81 int16 = (v5 | v5)
                                v1 = 3.14159
                                _ = v81
                            } else {
                                switch (fi71 + fi71) {
                                case -889:
                                    switch "" {
                                    case "foo bar":
                                        var v82 int = fi77
                                        _ = v82
                                    }
                                case 32768:
                                    v5 = 1000
                                }
                            }
                            if v3 {
                                v1 -= 0.0
                            } else {
                                var v83 uint = 2
                                fmt.Println(fi71, v83, fi77)
                                switch 100 {
                                case -119:
                                    if (42 > 2147483647) {
                                        v2[("\n\t" + "")] = "hello"
                                        v3 = v3
                                    }
                                }
                                v84 := GenS7{F0: v83, F1: 4294967295}
                                _ = v83
                                _ = v84
                            }
                            v85 := GenS8{F0: 675}
                            _ = v85
                        } else {
                            v1 += (0.1 + v1)
                            if len(v70) > 0 {
                                v70[0] = (fi71 - 4294967295)
                            }
                            fmt.Println(v5, v5, fi77)
                        }
                    }
                    var v86 float64 = (1e10 - (float64(((((((1.5 / (1e10 / 0.1)) * 0.1) - -2.25) / 1e10) - (0.1 / (((1.5 / 3.14159) - ((-2.25 + (3.14159 - -2.25)) / 1.0)) + 0.1))) * (((1e10 + (0.1 + 0.1)) / (1.5 / 1e10)) / ((1e10 - 1e10) - 1e10)))) * (v1 * ((3.14159 - 3.14159) * (v1 + -2.25)))))
                    _ = fi77
                    _ = v86
                }
                var v87 int = (fi71 ^ -2147483648)
                _ = fi71
                _ = v87
            }
            _ = v70
        } else {
            v88 := GenS4{F0: (!(!(!v3)))}
            var v89 uint32 = 42
            _ = v88
            _ = v89
        }
    default:
        v90 := GenS8{F0: 4294967295}
        _ = v90
    }
    v91 := []int{-1}
    for fi92 := 0; fi92 < 4; fi92++ {
        v4[("hello" + "\n\t")] = fi92
        switch (^(fi92)) {
        case -2147483648:
            for fi93 := 0; fi93 < 16; fi93++ {
                var v94 uint = 1
                _ = fi93
                _ = v94
            }
        default:
            fmt.Println(v3)
        }
        _ = fi92
    }
    var v95 float32 = (0.0 + 0.1)
    var v96 uint8 = 4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v91
    _ = v95
    _ = v96
}
