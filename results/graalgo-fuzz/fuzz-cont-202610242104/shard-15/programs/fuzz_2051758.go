package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 int32; F2 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8; F1 string; F2 uint64 }
type GenS11 struct { F0 bool; F1 bool }
func genFn0(p0 int32, p1 string) int32 {
    var v4 int8 = 127
    p1 = "hello"
    if (!(1 == 127)) {
        var v5 uint = 4294967295
        _ = v5
    } else {
        var v6 int = 1
        v4 += (v4 - (v4 + 127))
        var v7 int = (v6 % ((v6) | 1))
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return (p0 / ((p0) | 1))
}
func genFn1(p0 int16, p1 uint32) uint {
    p0 += p0
    if (true || false) {
        p0 += p0
        fmt.Println(p0, p0)
    } else {
        p1 += (^(p1))
        v8 := []int{-771, 148}
        fmt.Println(p1)
        _ = v8
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{2}
    v2 := []string{"\n\t", "hello", "a", (("hello" + "") + ((((("a" + "\n\t") + "a") + ("" + ((("a" + "foo bar") + "\n\t") + "hello"))) + "hello") + "hello"))}
    var v3 uint16 = 32767
    if false {
        var v9 int32 = -1
        _ = v9
    } else {
        var v10 int64 = 161
        v11 := GenS1{F0: 996, F1: 18446744073709551615}
        switch ("foo bar" + "") {
        case "":
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(fi12, fi12, fi12)
                _ = fi12
            }
        case "hello":
            fmt.Println(v3, v3)
        case "\n\t":
            if true {
                var v13 float32 = float32(((0.0 - 1.5) - 0.1))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v14 float64 = (0.1 - float64(v13))
                _ = v13
                _ = v14
            }
        default:
            var v15 int32 = 42
            _ = v15
        }
        _ = v10
        _ = v11
    }
    switch 39 {
    case -9223372036854775808:
        v3 += 255
    case -2147483648:
        v3 = 1
    }
    v3 += v3
    for fi16 := 0; fi16 < 16; fi16++ {
        v1 = append(v1, 65535)
        if true {
            v3 += (v3 >> ((42) & 63))
        } else {
            fmt.Println(fi16)
            var v17 uint8 = 1
            switch (^(fi16)) {
            case 32768:
                switch "" {
                case "foo bar":
                    if true {
                        v18 := GenS2{F0: (("\n\t" + "hello") + ""), F1: int64(v3), F2: 65535}
                        for fi19 := 0; fi19 < 6; fi19++ {
                            v3 += v3
                            var v20 int8 = 60
                            v21 := GenS3{F0: v17, F1: ("" >= "foo bar")}
                            var v22 int32 = genFn0(int32(v3), "\n\t")
                            var v23 int32 = int32(v20)
                            switch "foo bar" {
                            case "hello":
                                v23 += (v23 + -2147483648)
                            case "\n\t":
                                if (42 != 42) {
                                    if (0 < 4294967295) {
                                        v24 := GenS4{F0: 0, F1: 879, F2: v20}
                                        var v25 int = fi19
                                        v23 = int32(v17)
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 float32 = ((3.14159 + 1.5) * (float32(1.5) + (0.0 * (0.0 + 1.5))))
                                        v22 -= (v22 * (v22 % (((v22 >> ((228) & 63))) | 1)))
                                        _ = v26
                                    }
                                    switch (-(fi16)) {
                                    case -1:
                                        var v27 uint32 = 2147483647
                                        _ = v27
                                    case 128:
                                        switch (fi16 * 32767) {
                                        case -366:
                                            if (true && true) {
                                                v1 = append(v1, 0)
                                                _ = len(v1)
                                                if ((v3 >> ((genFn1(0, 42)) & 63)) <= (^(v3))) {
                                                    v23 = v22
                                                    if ((int(v17) <= fi16) && (v3 == (v3 ^ (v3 | v3)))) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v28 int16 = -975
                                                        v17 = 128
                                                        _ = v28
                                                    } else {
                                                        var v29 bool = true
                                                        _ = v29
                                                    }
                                                    v22 -= (-(v23))
                                                    v2 = append(v2, "foo bar")
                                                    if true {
                                                        v22 = v23
                                                    } else {
                                                        v30 := GenS2{F0: "a", F1: -1, F2: 889}
                                                        v3 -= v3
                                                        _ = len(v1)
                                                        _ = v30
                                                    }
                                                    v23 -= 42
                                                    v23 = (v23 & (^(v23)))
                                                }
                                                var v31 string = "foo bar"
                                                _ = v31
                                            }
                                            v32 := GenS5{F0: v23, F1: -1}
                                            _ = v32
                                        case -9223372036854775808:
                                            if false {
                                                if false {
                                                    v1 = append(v1, 958)
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (32767 != 1) {
                                                            fmt.Println(fi16, v23, v3)
                                                            fmt.Println(v17, v3, fi19)
                                                        }
                                                    case "hello":
                                                        var v33 string = "a"
                                                        _ = v33
                                                    }
                                                }
                                                fmt.Println(v3, v23, v17)
                                                if true {
                                                    var v34 int64 = -207
                                                    _ = v34
                                                } else {
                                                    var v35 bool = (!(!false))
                                                    if true {
                                                        fmt.Println(v35, v22)
                                                        if (!false) {
                                                            v36 := GenS6{F0: 890}
                                                            var v37 int16 = 32767
                                                            if v35 {
                                                                v38 := GenS5{F0: (-(v23)), F1: (v37 | (v37 ^ 1000))}
                                                                v39 := GenS6{F0: 4294967295}
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    } else {
                                                        if false {
                                                            fmt.Println(v35, v35, v20)
                                                        }
                                                        v20 -= (^(v20))
                                                        v3 = v3
                                                    }
                                                    _ = v35
                                                }
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                    default:
                                        switch (("a" + "hello") + ("foo bar" + "foo bar")) {
                                        case "a":
                                            v40 := GenS5{F0: genFn0((v22 * (v22 - (v22 % (((v23 ^ genFn0(v22, "\n\t"))) | 1)))), "foo bar"), F1: 1}
                                            _ = v40
                                        case "\n\t":
                                            v41 := GenS7{F0: fi16}
                                            _ = v41
                                        default:
                                            if true {
                                                v17 += (v17 * (^(v17)))
                                                var v42 int32 = (^(v22))
                                                _ = v42
                                            }
                                        }
                                    }
                                    if true {
                                        var v43 float64 = (0.1 + 1e10)
                                        _ = v43
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            var v44 int8 = (v20 - v20)
                                            _ = v44
                                        case "hello":
                                            var v45 int32 = genFn0(874, "a")
                                            _ = v45
                                        default:
                                            var v46 int8 = (^(v20))
                                            _ = v46
                                        }
                                        v20 += v20
                                        var v47 int16 = 1000
                                        _ = v47
                                    }
                                }
                            default:
                                var v48 uint64 = 2
                                _ = v48
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v18
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                fmt.Println(fi16, fi16, v3)
            }
            _ = v17
        }
        _ = fi16
    }
    switch 127 {
    case 128:
        if ((116 >= 42) && true) {
            var v49 uint64 = 18446744073709551615
            _ = v49
        } else {
            fmt.Println(v3)
            if (-736 != -2147483648) {
                for fi50 := 0; fi50 < 3; fi50++ {
                    switch "hello" {
                    case "hello":
                        switch "foo bar" {
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            if false {
                                var v51 string = ("\n\t" + "")
                                _ = v51
                            } else {
                                var v52 string = ("" + "hello")
                                v52 = ""
                                _ = v52
                            }
                        }
                    case "":
                        var v53 float64 = 0.1
                        _ = v53
                    case "a":
                        if false {
                            v3 += (v3 * 14)
                            fmt.Println(fi50, fi50)
                            v3 -= (v3 & 363)
                        } else {
                            _ = len(v1)
                        }
                    }
                    if (!(-237 == 2147483647)) {
                        for fi54 := 0; fi54 < 12; fi54++ {
                            v3 -= (v3 / ((v3) | 1))
                            v3 = (v3 << ((1) & 63))
                            if true {
                                fmt.Println(fi54, v3)
                            } else {
                                _ = len(v1)
                            }
                            _ = fi54
                        }
                    }
                    v3 = (v3 / (((v3 + (v3 << ((18446744073709551615) & 63)))) | 1))
                    _ = fi50
                }
                for fi55 := 0; fi55 < 11; fi55++ {
                    switch "\n\t" {
                    case "hello":
                        var v56 uint8 = 255
                        _ = v56
                    case "a":
                        v57 := GenS3{F0: 2, F1: false}
                        _ = v57
                    case "":
                        v3 = (v3 + (v3 % ((42) | 1)))
                    default:
                        var v58 string = "a"
                        _ = v58
                    }
                    _ = fi55
                }
                var v59 int32 = 2147483647
                _ = v59
            } else {
                var v60 uint32 = 344
                var v61 uint64 = 662
                for fi62 := 0; fi62 < 11; fi62++ {
                    var v63 string = ""
                    v60 = (v60 - (v60 ^ 36))
                    if (!false) {
                        if false {
                            fmt.Println(v63)
                            v64 := GenS1{F0: (^(v61)), F1: v61}
                            switch fi62 {
                            case -2147483648:
                                if false {
                                    switch fi62 {
                                    case 0:
                                        v65 := GenS5{F0: 0, F1: 998}
                                        _ = v65
                                    case 127:
                                        fmt.Println(fi62, v60, v60)
                                    }
                                    if false {
                                        v60 = v60
                                        if len(v2) > 0 {
                                            v2[0] = v63
                                        }
                                        var v66 uint8 = 250
                                        _ = v66
                                    } else {
                                        v60 = (v60 - uint32(v61))
                                        v67 := GenS8{F0: false, F1: genFn0(genFn0(genFn0(-594, v63), ("hello" + v63)), ("\n\t" + v63)), F2: 32767}
                                        for fi68 := 0; fi68 < 19; fi68++ {
                                            v60 -= v60
                                            _ = fi68
                                        }
                                        _ = v67
                                    }
                                } else {
                                    v69 := GenS1{F0: v61, F1: 1}
                                    var v70 string = "a"
                                    var v71 uint16 = v3
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                            default:
                                if (((!(-9223372036854775808 <= 1)) || true) && (128 >= 73)) {
                                    fmt.Println(v60)
                                    _ = len(v1)
                                } else {
                                    v3 += (^(v3))
                                    for fi72 := 0; fi72 < 8; fi72++ {
                                        v73 := GenS1{F0: (v61 << ((genFn1(int16(v3), (v60 & v60))) & 63)), F1: v61}
                                        _ = fi72
                                        _ = v73
                                    }
                                    var v74 bool = true
                                    _ = v74
                                }
                            }
                            _ = v64
                        } else {
                            for fi75 := 0; fi75 < 9; fi75++ {
                                fmt.Println(fi75, fi62, fi75)
                                if true {
                                    v63 = v63
                                } else {
                                    v63 = ("" + v63)
                                }
                                _ = fi75
                            }
                            var v76 int32 = -1
                            if (1e10 >= 0.1) {
                                v77 := GenS1{F0: 18446744073709551615, F1: (v61 - 1)}
                                v78 := GenS8{F0: (!(!(546 == int64(v76)))), F1: -2147483648, F2: v3}
                                if false {
                                    if false {
                                        v79 := GenS9{F0: (v3 * v3)}
                                        var v80 uint8 = 1
                                        _ = v79
                                        _ = v80
                                    }
                                } else {
                                    for fi81 := 0; fi81 < 20; fi81++ {
                                        v76 = int32(v60)
                                        if (-2.25 > 1.5) {
                                            _ = len(v1)
                                            switch (fi62 & fi81) {
                                            case 32768:
                                                if (false && true) {
                                                    if len(v2) > 0 {
                                                        v2[0] = v63
                                                    }
                                                    var v82 int32 = v76
                                                    v61 -= (^(v61))
                                                    _ = v82
                                                } else {
                                                    fmt.Println(v3, fi81)
                                                    var v83 string = v63
                                                    _ = v83
                                                }
                                            }
                                            v84 := GenS10{F0: 136, F1: v63, F2: (v61 + v61)}
                                            _ = v84
                                        } else {
                                            v61 -= (^(v61))
                                            if len(v2) > 0 {
                                                v2[0] = ("" + v63)
                                            }
                                            switch v63 {
                                            case "\n\t":
                                                if len(v2) > 0 {
                                                    v2[0] = v63
                                                }
                                            case "a":
                                                var v85 uint = 42
                                                _ = v85
                                            }
                                        }
                                        v63 = ("a" + v63)
                                        _ = fi81
                                    }
                                    _ = len(v2)
                                }
                                _ = v77
                                _ = v78
                            } else {
                                v76 += v76
                            }
                            _ = v76
                        }
                    } else {
                        var v86 int16 = -1
                        v87 := GenS3{F0: 28, F1: (float64(float32(float64(0.1))) < 1e10)}
                        switch fi62 {
                        case 9223372036854775807:
                            v63 = "a"
                        default:
                            v88 := GenS1{F0: (v61 ^ v61), F1: v61}
                            _ = v88
                        }
                        _ = v86
                        _ = v87
                    }
                    _ = fi62
                    _ = v63
                }
                _ = v60
                _ = v61
            }
        }
    case -2147483648:
        v3 += (^(v3))
    case 127:
        v3 = v3
    default:
        switch ("foo bar" + ("a" + "\n\t")) {
        case "a":
            var v89 bool = false
            _ = v89
        case "\n\t":
            v90 := GenS8{F0: (100 != 0), F1: 1, F2: 83}
            _ = v90
        default:
            v91 := GenS11{F0: true, F1: (309 == 2)}
            _ = v91
        }
    }
    if ("a" > ("hello" + "\n\t")) {
        var v92 bool = false
        _ = v92
    }
    var v93 uint8 = 0
    var v94 int32 = 979
    _ = len(v2)
    v3 += 32767
    v95 := GenS2{F0: "foo bar", F1: 84, F2: 2147483647}
    for fi96 := 0; fi96 < 12; fi96++ {
        var v97 uint8 = (v93 + v93)
        _ = fi96
        _ = v97
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v93
    _ = v94
    _ = v95
}
