package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 uint64 }
func genFn0(p0 int16, p1 int8, p2 uint8) int32 {
    p0 = -184
    var v5 uint = 714
    var v6 int = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -509
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 int16 = -32768
    var v4 uint16 = 42
    v3 = (v3 ^ (v3 | (v3 >> ((255) & 63))))
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 4; fi7++ {
        if true {
            v8 := []int{(fi7 * 128), (fi7 / ((0) | 1)), (fi7 ^ int(v3))}
            fmt.Println(fi7, v3, v4)
            var v9 uint32 = 2147483648
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 2}
        _ = fi7
        _ = v10
    }
    switch 4294967295 {
    case -9223372036854775808:
        if false {
            var v11 uint32 = 42
            v12 := []int{42, 32767, -1, -1}
            var v13 uint16 = v4
            _ = v11
            _ = v12
            _ = v13
        } else {
            v14 := GenS1{F0: (v4 ^ v4)}
            var v15 uint32 = 824
            if false {
                if false {
                    v16 := []int{0, 42, 128}
                    var v17 float32 = 1.5
                    v3 = 1000
                    _ = v16
                    _ = v17
                } else {
                    var v18 int16 = (-(v3))
                    var v19 int64 = 9223372036854775807
                    v20 := GenS1{F0: 32767}
                    _ = v18
                    _ = v19
                    _ = v20
                }
            }
            _ = v14
            _ = v15
        }
    case 128:
        if (341 <= 1) {
            v21 := GenS2{F0: 18446744073709551615, F1: -676}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: (v4 << ((1) & 63))}
                var v24 uint64 = 18446744073709551615
                var v25 uint16 = 1
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
        }
    }
    v4 = v4
    v26 := []int{128, 32767, 4294967295}
    v27 := GenS1{F0: (v4 & v4)}
    if (!(-9223372036854775808 != 137)) {
        var v28 int16 = -391
        v29 := GenS1{F0: v4}
        v30 := GenS3{F0: "foo bar"}
        _ = v28
        _ = v29
        _ = v30
    }
    delete(v1, "foo bar")
    if true {
        switch "" {
        case "":
            if (false || false) {
                for fi31 := 0; fi31 < 8; fi31++ {
                    fmt.Println(v3, v4, v3)
                    _ = fi31
                }
                fmt.Println(v3, v4)
                v3 -= v3
            } else {
                fmt.Println(v4)
                v3 += (v3 & v3)
                v3 += v3
            }
        case "a":
            if true {
                fmt.Println(v3, v3)
                for fi32 := 0; fi32 < 12; fi32++ {
                    v2[(fi32 / ((-43) | 1))] = true
                    _ = fi32
                }
                switch "foo bar" {
                case "\n\t":
                    v33 := GenS1{F0: 65535}
                    _ = v33
                case "hello":
                    v4 = (v4 >> ((255) & 63))
                default:
                    v2[int(v3)] = true
                }
            } else {
                v4 = (^(v4))
            }
        default:
            if true {
                for fi34 := 0; fi34 < 12; fi34++ {
                    var v35 uint64 = 18446744073709551615
                    v36 := GenS4{F0: 0}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
                if (986 <= 1) {
                    v37 := GenS3{F0: ("a" + (("a" + ("\n\t" + ("" + ("hello" + "foo bar")))) + "hello"))}
                    v2[482] = (201 < -709)
                    for fi38 := 0; fi38 < 18; fi38++ {
                        v39 := v1["\n\t"]
                        if false {
                            for fi40 := 0; fi40 < 10; fi40++ {
                                switch "foo bar" {
                                case "":
                                    var v41 int8 = -103
                                    var v42 int16 = -1
                                    _ = v41
                                    _ = v42
                                case "a":
                                    v3 = v3
                                    var v43 float32 = (-2.25 + -2.25)
                                    _ = v43
                                case "\n\t":
                                    var v44 uint32 = 883
                                    v45 := GenS1{F0: v4}
                                    v46 := v2[(^(fi40))]
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v47 uint16 = 0
                                        fmt.Println(fi38, v39, fi40)
                                        fmt.Println(v4, fi40)
                                        _ = v47
                                    case "a":
                                        v48 := GenS4{F0: 2147483648}
                                        _ = v48
                                    default:
                                        if v39 {
                                            var v49 float64 = 1.5
                                            v3 += v3
                                            _ = v49
                                        }
                                    }
                                    if (634 > 65535) {
                                        v50 := GenS2{F0: 65535, F1: -291}
                                        _ = v50
                                    }
                                    v51 := GenS4{F0: 2147483647}
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(fi38, fi40, fi40)
                                    case "a":
                                        v3 = 1
                                    }
                                    _ = v51
                                }
                                _ = fi40
                            }
                            for fi52 := 0; fi52 < 9; fi52++ {
                                v26 = append(v26, fi52)
                                if true {
                                    if (v39 || v39) {
                                        var v53 float32 = (1e10 - 1e10)
                                        v54 := GenS4{F0: 2}
                                        _ = v53
                                        _ = v54
                                    } else {
                                        if (!false) {
                                            v3 += 32767
                                            v55 := GenS5{F0: 232, F1: 808}
                                            switch fi52 {
                                            case 1:
                                                _ = len(v26)
                                            }
                                            _ = v55
                                        } else {
                                            var v56 uint32 = 4294967295
                                            var v57 uint16 = (v4 + uint16(v56))
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = len(v26)
                                    }
                                    switch "" {
                                    case "a":
                                        switch "" {
                                        case "hello":
                                            var v58 uint16 = 42
                                            _ = v58
                                        default:
                                            v59 := GenS6{F0: 42, F1: 4294967295}
                                            _ = v59
                                        }
                                    default:
                                        switch (fi52 / ((fi38) | 1)) {
                                        case -9223372036854775808:
                                            fmt.Println(v3)
                                        default:
                                            if ("" >= ("foo bar" + "")) {
                                                if true {
                                                    fmt.Println(fi52)
                                                    v60 := GenS1{F0: (v4 + (v4 * (v4 - v4)))}
                                                    _ = v60
                                                }
                                            } else {
                                                if (-447 <= 1) {
                                                    v39 = (2 == 703)
                                                } else {
                                                    v61 := GenS7{F0: 1}
                                                    _ = v61
                                                }
                                            }
                                        }
                                    }
                                    var v62 uint8 = 1
                                    _ = v62
                                } else {
                                    var v63 int32 = int32(v4)
                                    fmt.Println(v63, v39)
                                    v3 = (v3 * v3)
                                    _ = v63
                                }
                                _ = fi52
                            }
                            for fi64 := 0; fi64 < 7; fi64++ {
                                switch fi64 {
                                case -106:
                                    var v65 int8 = 0
                                    _ = v65
                                }
                                v3 -= (v3 << ((140) & 63))
                                v3 -= (v3 + (v3 / ((-32768) | 1)))
                                _ = fi64
                            }
                        }
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                } else {
                    var v66 int = 0
                    _ = v66
                }
            } else {
                switch "\n\t" {
                case "a":
                    if false {
                        var v67 string = "hello"
                        v4 -= (v4 % ((42) | 1))
                        v68 := GenS6{F0: -863, F1: 18446744073709551615}
                        _ = v67
                        _ = v68
                    }
                case "":
                    for fi69 := 0; fi69 < 7; fi69++ {
                        var v70 uint16 = v4
                        _ = fi69
                        _ = v70
                    }
                default:
                    var v71 int8 = -118
                    _ = v71
                }
                v4 += v4
            }
        }
    } else {
        for fi72 := 0; fi72 < 16; fi72++ {
            if ((-(fi72)) != 128) {
                v4 = v4
            }
            if len(v26) > 0 {
                v26[0] = fi72
            }
            _ = fi72
        }
        v73 := GenS4{F0: uint32(v4)}
        var v74 string = ("" + "foo bar")
        _ = v73
        _ = v74
    }
    var v75 uint8 = 0
    v3 = v3
    var v76 bool = false
    _ = len(v26)
    switch 2147483647 {
    case 0:
        var v77 string = ("\n\t" + "")
        _ = v77
    case 32768:
        switch 32767 {
        case -2147483648:
            switch "foo bar" {
            case "":
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
            }
        case 9223372036854775807:
            v2[374] = (v76 || ((true || (719 != -1)) && (!v76)))
        default:
            if (180 != 65535) {
                v78 := GenS6{F0: 0, F1: 356}
                switch "hello" {
                case "a":
                    for fi79 := 0; fi79 < 6; fi79++ {
                        _ = len(v1)
                        v75 = (v75 | (v75 + v75))
                        for fi80 := 0; fi80 < 11; fi80++ {
                            v81 := GenS8{F0: 483, F1: 0}
                            v82 := GenS7{F0: 2}
                            if true {
                                var v83 string = "hello"
                                _ = v83
                            } else {
                                switch (fi80 + (^(fi79))) {
                                case 42:
                                    v75 = v75
                                }
                                var v84 int = (fi80 | (fi80 & fi79))
                                var v85 float32 = 1e10
                                _ = v84
                                _ = v85
                            }
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                        _ = fi79
                    }
                case "foo bar":
                    v76 = v76
                default:
                    v4 = (v4 ^ v4)
                }
                if false {
                    fmt.Println(v4)
                    var v86 int64 = 9223372036854775807
                    v3 = 32767
                    _ = v86
                }
                _ = v78
            }
        }
    case 42:
        v4 = v4
    }
    switch "hello" {
    case "foo bar":
        for fi87 := 0; fi87 < 17; fi87++ {
            var v88 int16 = (v3 + (v3 - v3))
            _ = fi87
            _ = v88
        }
    case "\n\t":
        if len(v26) >= 2 {
            v26 = v26[1:len(v26)]
        }
    }
    fmt.Println(v3, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v75
    _ = v76
}
