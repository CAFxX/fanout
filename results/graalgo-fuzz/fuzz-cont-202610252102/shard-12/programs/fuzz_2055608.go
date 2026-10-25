package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 uint32; F2 bool }
type GenS9 struct { F0 int32; F1 int16; F2 uint }
type GenS10 struct { F0 int8; F1 bool; F2 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16; F1 int16 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 int16; F1 uint }
func genFn0(p0 int, p1 uint32) uint {
    if true {
        v6 := []int{(p0 / ((p0) | 1)), (p0 << ((255) & 63)), p0, p0}
        var v7 int = p0
        _ = v6
        _ = v7
    } else {
        p1 += p1
        p1 = (^(p1))
    }
    v8 := GenS1{F0: 1}
    p1 = (p1 << ((1) & 63))
    fmt.Println(p1, p0)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := []int{p0, 75, fi9}
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 2}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = -371
    v3 := []string{"", ((("foo bar" + "foo bar") + "\n\t") + "foo bar")}
    v4 := make(map[string]uint32)
    var v5 int = 913
    var v13 uint32 = 65535
    var v14 uint = genFn0(v5, (v13 | (v13 % ((2) | 1))))
    v15 := GenS2{F0: true, F1: 310}
    if (!false) {
        var v16 uint32 = (^(v13))
        _ = v16
    } else {
        var v17 float64 = float64(((3.14159 * -2.25) - (-2.25 / -2.25)))
        var v18 float64 = v17
        v19 := v1[(-(v5))]
        _ = v17
        _ = v18
        _ = v19
    }
    if false {
        fmt.Println(v13)
        for fi20 := 0; fi20 < 8; fi20++ {
            switch (v5 << ((65535) & 63)) {
            case -615:
                switch v5 {
                case 32768:
                    v21 := v4[(("hello" + "\n\t") + "")]
                    _ = v21
                default:
                    delete(v1, -552)
                }
            }
            v5 += (v5 ^ (v5 & fi20))
            _ = fi20
        }
        if ((v2 | v2) <= v2) {
            for fi22 := 0; fi22 < 5; fi22++ {
                v5 += fi22
                _ = fi22
            }
            fmt.Println(v2, v5)
        } else {
            for fi23 := 0; fi23 < 9; fi23++ {
                for fi24 := 0; fi24 < 6; fi24++ {
                    var v25 int16 = 1
                    _ = fi24
                    _ = v25
                }
                v26 := GenS3{F0: 65535}
                _ = fi23
                _ = v26
            }
            switch "" {
            case "a":
                for fi27 := 0; fi27 < 17; fi27++ {
                    if (1.5 != 3.14159) {
                        var v28 uint64 = 2
                        _ = v28
                    }
                    v4[((("\n\t" + "a") + "a") + "\n\t")] = 526
                    _ = fi27
                }
            case "foo bar":
                v29 := GenS4{F0: 32767}
                _ = v29
            case "hello":
                v30 := GenS5{F0: 0, F1: uint32(v14)}
                _ = v30
            default:
                delete(v1, 32768)
            }
        }
    }
    switch (v5 % ((32767) | 1)) {
    case -9223372036854775808:
        if (42 > 0) {
            var v31 uint64 = 326
            _ = v31
        } else {
            if ((v5 + v5) < -231) {
                fmt.Println(v5)
                var v32 bool = true
                _ = v32
            }
            var v33 uint64 = 18446744073709551615
            _ = v33
        }
    default:
        _ = len(v1)
    }
    v34 := GenS3{F0: 2}
    v35 := GenS6{F0: false, F1: 0}
    if true {
        if false {
            if (true && ((^(v5)) > (v5 % (((v5 / (((v5 >> ((v14) & 63))) | 1))) | 1)))) {
                var v36 int = -406
                if false {
                    if false {
                        if ((v2 >> ((genFn0(v5, v13)) & 63)) <= v2) {
                            switch ("a" + ("foo bar" + "hello")) {
                            case "a":
                                v37 := GenS7{F0: 42}
                                _ = v37
                            case "foo bar":
                                v14 += v14
                            default:
                                if (607 < 833) {
                                    v5 = (^(v5))
                                }
                            }
                        }
                        fmt.Println(v36, v5, v5)
                        switch v5 {
                        case 2147483647:
                            _ = len(v4)
                        case -2147483648:
                            v13 = (v13 >> (((^(v14))) & 63))
                        case 522:
                            for fi38 := 0; fi38 < 6; fi38++ {
                                _ = len(v1)
                                _ = fi38
                            }
                        default:
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                    }
                    for fi39 := 0; fi39 < 15; fi39++ {
                        if (!(!true)) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            if (-32768 != -530) {
                                if true {
                                    v2 -= -1
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v5, fi39, v5)
                                    case "\n\t":
                                        if (32767 == 111) {
                                            for fi40 := 0; fi40 < 6; fi40++ {
                                                v41 := GenS3{F0: 255}
                                                if (42 < -32768) {
                                                    v42 := GenS1{F0: 1}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v42
                                                } else {
                                                    v43 := GenS2{F0: (42 <= 1), F1: 0}
                                                    switch "hello" {
                                                    case "\n\t":
                                                        switch (fi39 % (((^(fi39))) | 1)) {
                                                        case 127:
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                        case -1:
                                                            if (1 != 474) {
                                                                v4["foo bar"] = (v13 & 630)
                                                                if true {
                                                                    var v44 float64 = -2.25
                                                                    switch "hello" {
                                                                    case "":
                                                                        v13 -= 2147483648
                                                                    default:
                                                                        var v45 float64 = (v44 * v44)
                                                                        _ = v45
                                                                    }
                                                                    fmt.Println(v13, fi40, v5)
                                                                    _ = v44
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v46 string = "hello"
                                                                        _ = v46
                                                                    case "":
                                                                        var v47 uint32 = 65535
                                                                        _ = v47
                                                                    default:
                                                                        v48 := GenS8{F0: -23, F1: 2, F2: (!(!true))}
                                                                        _ = v48
                                                                    }
                                                                    v49 := GenS5{F0: 0, F1: (v13 % ((v13) | 1))}
                                                                    v1[(v5 | v36)] = false
                                                                    _ = v49
                                                                }
                                                                var v50 uint = 18446744073709551615
                                                                v5 = v5
                                                                v13 = (v13 << ((genFn0(int(v2), 65535)) & 63))
                                                                _ = v50
                                                            } else {
                                                                fmt.Println(v14, v36)
                                                                _ = len(v4)
                                                            }
                                                        case 328:
                                                            var v51 int32 = 2147483647
                                                            _ = v51
                                                        default:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v52 bool = (!(false && (true && (42 < 1))))
                                                                _ = v52
                                                            default:
                                                                v14 += (v14 >> ((v14) & 63))
                                                            }
                                                        }
                                                    default:
                                                        v53 := GenS4{F0: 450}
                                                        _ = v53
                                                    }
                                                    _ = v43
                                                }
                                                if true {
                                                    var v54 int = (v5 | (-(fi40)))
                                                    _ = v54
                                                } else {
                                                    var v55 float64 = 0.0
                                                    _ = v55
                                                }
                                                _ = fi40
                                                _ = v41
                                            }
                                        } else {
                                            v13 = v13
                                            v56 := GenS9{F0: -1, F1: -245, F2: 1}
                                            _ = v56
                                        }
                                    default:
                                        v57 := GenS8{F0: 100, F1: (v13 + uint32(v14)), F2: (1 != 42)}
                                        _ = v57
                                    }
                                    v36 = 233
                                } else {
                                    v13 += v13
                                    if (42 != 1) {
                                        v4[("\n\t" + "hello")] = (v13 ^ (v13 << ((v14) & 63)))
                                        for fi58 := 0; fi58 < 1; fi58++ {
                                            v59 := GenS10{F0: 0, F1: true, F2: 133}
                                            _ = fi58
                                            _ = v59
                                        }
                                    }
                                    v13 = v13
                                }
                            } else {
                                fmt.Println(v36, v5, v2)
                                v60 := GenS7{F0: -2147483648}
                                v36 -= -9223372036854775808
                                _ = v60
                            }
                            if (uint8(v13) <= 1) {
                                delete(v1, 336)
                                v2 = v2
                                fmt.Println(v5, v2)
                            } else {
                                v1[(v5 / ((v36) | 1))] = false
                                var v61 uint16 = 0
                                for fi62 := 0; fi62 < 17; fi62++ {
                                    v63 := GenS10{F0: -1, F1: (-1 == 1), F2: 170}
                                    _ = fi62
                                    _ = v63
                                }
                                _ = v61
                            }
                            fmt.Println(v5, v13)
                        }
                        _ = fi39
                    }
                } else {
                    v36 = (-(v36))
                    for fi64 := 0; fi64 < 11; fi64++ {
                        _ = len(v1)
                        v5 += (fi64 ^ (v36 / (((fi64 ^ v36)) | 1)))
                        _ = fi64
                    }
                }
                v14 = (^(v14))
                _ = v36
            } else {
                v65 := GenS11{F0: (!true)}
                v66 := GenS12{F0: 32767, F1: 32767}
                _ = len(v4)
                _ = v65
                _ = v66
            }
            v3 = append(v3, "")
        } else {
            for fi67 := 0; fi67 < 3; fi67++ {
                for fi68 := 0; fi68 < 8; fi68++ {
                    v5 += (fi68 | (v5 ^ v5))
                    var v69 uint = uint(v13)
                    _ = fi68
                    _ = v69
                }
                var v70 uint16 = 32767
                fmt.Println(v14, fi67)
                _ = fi67
                _ = v70
            }
            v5 -= -9223372036854775808
            v71 := GenS13{F0: "\n\t", F1: 42}
            _ = v71
        }
        if false {
            fmt.Println(v2, v14)
            v13 = 0
            v3 = append(v3, "a")
        } else {
            var v72 int32 = -1
            var v73 string = (("foo bar" + "foo bar") + "hello")
            v5 += 1
            _ = v72
            _ = v73
        }
    } else {
        if false {
            fmt.Println(v13, v2, v5)
        } else {
            for fi74 := 0; fi74 < 13; fi74++ {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                fmt.Println(v5)
                v1[2147483647] = false
                _ = fi74
            }
            delete(v1, v5)
            if (-72 <= 1) {
                v3 = append(v3, "\n\t")
                for fi75 := 0; fi75 < 1; fi75++ {
                    v5 += fi75
                    _ = fi75
                }
            }
        }
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v13)
    case "hello":
        switch "\n\t" {
        case "":
            var v76 uint = (v14 / ((18446744073709551615) | 1))
            _ = v76
        default:
            var v77 uint = 0
            _ = v77
        }
    case "":
        var v78 int64 = -360
        _ = v78
    }
    v79 := GenS3{F0: 42}
    var v80 int32 = 42
    v81 := GenS7{F0: v80}
    var v82 uint64 = 2
    for fi83 := 0; fi83 < 11; fi83++ {
        v84 := GenS14{F0: 42, F1: (^(v14))}
        _ = fi83
        _ = v84
    }
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v80:", v80)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
