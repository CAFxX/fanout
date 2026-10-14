package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 int) uint {
    v7 := GenS1{F0: "a"}
    v8 := []int{-9223372036854775808, (-(p0)), p0}
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 uint8, p1 int8, p2 uint16) int16 {
    if true {
        switch "foo bar" {
        case "":
            v9 := []int{-9223372036854775808, 32768, -9223372036854775808}
            _ = v9
        case "foo bar":
            p1 = (-(p1))
        }
    } else {
        var v10 uint16 = (p2 << ((genFn0(2147483647)) & 63))
        for fi11 := 0; fi11 < 4; fi11++ {
            if (p2 < 42) {
                p1 = 100
            } else {
                for fi12 := 0; fi12 < 10; fi12++ {
                    var v13 uint8 = 1
                    _ = fi12
                    _ = v13
                }
                fmt.Println(p1, p0, fi11)
                if (-440 >= 0) {
                    var v14 bool = (!false)
                    switch (fi11 >> ((831) & 63)) {
                    case 32768:
                        var v15 uint32 = 4294967295
                        _ = v15
                    case 32767:
                        switch fi11 {
                        case -831:
                            p1 = p1
                        }
                    }
                    p1 += (p1 & (p1 >> ((2) & 63)))
                    _ = v14
                }
            }
            _ = fi11
        }
        _ = v10
    }
    p1 += -1
    _ = p0
    _ = p1
    _ = p2
    return -880
}

func main() {
    var v1 bool = false
    var v2 int32 = -1
    var v3 uint8 = 186
    v4 := make(map[int]int)
    var v5 uint8 = v3
    var v6 int64 = 1
    for fi16 := 0; fi16 < 12; fi16++ {
        switch ("" + "hello") {
        case "\n\t":
            var v17 uint32 = 4294967295
            _ = v17
        case "foo bar":
            v18 := []int{(-(fi16)), fi16}
            _ = v18
        case "":
            for fi19 := 0; fi19 < 9; fi19++ {
                switch "foo bar" {
                case "":
                    switch "\n\t" {
                    case "\n\t":
                        var v20 int32 = 600
                        _ = v20
                    case "":
                        v1 = true
                    case "foo bar":
                        fmt.Println(v3, fi16)
                    }
                case "hello":
                    v21 := []int{fi19, int(v2), fi19, fi16}
                    _ = v21
                case "a":
                    v22 := GenS1{F0: ("a" + "hello")}
                    _ = v22
                }
                v23 := GenS1{F0: ""}
                _ = fi19
                _ = v23
            }
        default:
            v24 := GenS2{F0: int(v2), F1: (^(v3))}
            _ = v24
        }
        v25 := GenS1{F0: ""}
        _ = fi16
        _ = v25
    }
    v26 := GenS1{F0: "\n\t"}
    if false {
        fmt.Println(v5)
        if (962 <= 0) {
            if v1 {
                fmt.Println(v5)
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := v4[fi27]
                    _ = fi27
                    _ = v28
                }
            } else {
                switch "a" {
                case "a":
                    if (!(743 < 255)) {
                        var v29 uint64 = 4294967295
                        fmt.Println(v3, v29)
                        if v1 {
                            v5 = (v5 % (((v3 + uint8(v29))) | 1))
                            var v30 uint16 = 2
                            _ = v30
                        }
                        _ = v29
                    }
                case "foo bar":
                    var v31 int8 = 100
                    _ = v31
                case "hello":
                    v1 = (-128 > 100)
                default:
                    v6 -= v6
                }
                if ((1e10 != -2.25) || false) {
                    v32 := []int{-483, -277, int(v6), 4294967295}
                    fmt.Println(v1, v5, v3)
                    _ = v32
                } else {
                    v33 := GenS2{F0: 127, F1: v5}
                    _ = v33
                }
            }
            switch "foo bar" {
            case "foo bar":
                v5 = (^(v3))
            case "":
                if v1 {
                    v1 = (!v1)
                    for fi34 := 0; fi34 < 1; fi34++ {
                        v35 := GenS3{F0: (v6 | 9223372036854775807)}
                        _ = fi34
                        _ = v35
                    }
                    if (false && (127 < -1)) {
                        switch (("" + "foo bar") + "a") {
                        case "":
                            fmt.Println(v5)
                        default:
                            v6 += v6
                        }
                        switch -102 {
                        case 2147483647:
                            var v36 int64 = 1
                            _ = v36
                        }
                        var v37 float64 = ((1.5 / (0.1 * float64(((-2.25 + float32((3.14159 - -2.25))) / 1.0)))) * (0.0 - 1e10))
                        _ = v37
                    } else {
                        fmt.Println(v2, v6, v6)
                        for fi38 := 0; fi38 < 2; fi38++ {
                            v2 = v2
                            if false {
                                switch (fi38 - (fi38 | (fi38 * -1))) {
                                case 32767:
                                    var v39 uint = 32
                                    _ = v39
                                case 16:
                                    fmt.Println(v3)
                                default:
                                    if true {
                                        for fi40 := 0; fi40 < 12; fi40++ {
                                            switch ("foo bar" + "hello") {
                                            case "":
                                                v41 := []int{fi40, (fi38 & fi38)}
                                                _ = v41
                                            case "\n\t":
                                                fmt.Println(v2)
                                            default:
                                                v42 := v4[0]
                                                _ = v42
                                            }
                                            v43 := GenS2{F0: fi40, F1: v3}
                                            var v44 float32 = (1e10 * 0.1)
                                            switch fi38 {
                                            case -2147483648:
                                                v6 += (v6 ^ (^(v6)))
                                            }
                                            switch fi38 {
                                            case -894:
                                                fmt.Println(v5, v3)
                                            }
                                            v5 -= v3
                                            fmt.Println(v5)
                                            if v1 {
                                                _ = len(v4)
                                                var v45 uint = 929
                                                _ = v45
                                            } else {
                                                v46 := []int{(fi38 % ((32767) | 1)), fi38, 1, (fi38 + 4294967295)}
                                                fmt.Println(v3, fi40, v6)
                                                _ = v46
                                            }
                                            _ = fi40
                                            _ = v43
                                            _ = v44
                                        }
                                        for fi47 := 0; fi47 < 20; fi47++ {
                                            switch fi38 {
                                            case -2147483648:
                                                var v48 uint8 = (v3 ^ 100)
                                                _ = v48
                                            case -533:
                                                fmt.Println(v2, fi47, v3)
                                            case 9223372036854775807:
                                                v5 += (^(v3))
                                            default:
                                                var v49 uint16 = 1
                                                _ = v49
                                            }
                                            _ = fi47
                                        }
                                    } else {
                                        var v50 string = "hello"
                                        _ = v50
                                    }
                                }
                            }
                            for fi51 := 0; fi51 < 18; fi51++ {
                                var v52 uint64 = 0
                                v53 := []int{fi38, (fi51 & (fi51 + fi51))}
                                v1 = ((v3 < (^(v3))) || false)
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            _ = fi38
                        }
                    }
                }
            default:
                if v1 {
                    if (3.14159 != 1.5) {
                        v1 = false
                        _ = len(v4)
                        v54 := GenS4{F0: 1, F1: 42, F2: v2}
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 7; fi55++ {
                            v56 := GenS2{F0: (fi55 | 540), F1: v3}
                            for fi57 := 0; fi57 < 14; fi57++ {
                                if false {
                                    v58 := GenS3{F0: 9223372036854775807}
                                    _ = v58
                                }
                                v6 = (^(v6))
                                _ = fi57
                            }
                            _ = fi55
                            _ = v56
                        }
                        v59 := GenS3{F0: v6}
                        switch -557 {
                        case 4294967295:
                            if v1 {
                                var v60 float32 = 0.0
                                _ = v60
                            } else {
                                for fi61 := 0; fi61 < 3; fi61++ {
                                    var v62 uint8 = (^(v3))
                                    _ = fi61
                                    _ = v62
                                }
                                v3 = (v3 ^ 255)
                                var v63 int = -575
                                _ = v63
                            }
                        default:
                            var v64 bool = v1
                            _ = v64
                        }
                        _ = v59
                    }
                }
            }
            var v65 uint32 = 2147483648
            _ = v65
        } else {
            for fi66 := 0; fi66 < 12; fi66++ {
                var v67 float64 = -2.25
                v68 := GenS3{F0: 0}
                _ = fi66
                _ = v67
                _ = v68
            }
            if true {
                v69 := GenS4{F0: 26, F1: -908, F2: v2}
                v70 := GenS2{F0: -9223372036854775808, F1: v3}
                var v71 float64 = (3.14159 - -2.25)
                _ = v69
                _ = v70
                _ = v71
            }
        }
    }
    v72 := GenS2{F0: 767, F1: v3}
    v5 += v5
    v1 = (genFn1(v3, 1, 32767) <= genFn1(v5, 0, 920))
    switch (("hello" + "") + "\n\t") {
    case "foo bar":
        v3 = (^(v3))
    case "a":
        v73 := []int{}
        _ = v73
    case "\n\t":
        if (v1 || ((187 >= 721) || ((int16(v2) == -32768) && (float32(1e10) > 0.0)))) {
            for fi74 := 0; fi74 < 18; fi74++ {
                fmt.Println(v2)
                if (0 == 753) {
                    var v75 uint16 = 0
                    var v76 bool = (v75 == (v75 >> ((18446744073709551615) & 63)))
                    for fi77 := 0; fi77 < 8; fi77++ {
                        var v78 int64 = -1
                        var v79 uint8 = (^(v5))
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                    _ = v75
                    _ = v76
                } else {
                    var v80 uint8 = (v3 >> ((755) & 63))
                    switch fi74 {
                    case -9223372036854775808:
                        v6 -= v6
                    }
                    v1 = true
                    _ = v80
                }
                for fi81 := 0; fi81 < 16; fi81++ {
                    v82 := []int{173, (fi74 & (fi74 % ((fi74) | 1))), fi81, -478}
                    v3 += (v3 * (v3 ^ v5))
                    if false {
                        v82 = append(v82, (fi74 ^ (fi81 & (fi74 * 32767))))
                    } else {
                        switch ("\n\t" + "hello") {
                        case "hello":
                            v2 -= (v2 / ((v2) | 1))
                        }
                        if ((v3 << ((18446744073709551615) & 63)) == (v3 << ((804) & 63))) {
                            if v1 {
                                v82 = append(v82, fi81)
                                switch (fi81 + (-(fi81))) {
                                case 419:
                                    if v1 {
                                        var v83 uint = 42
                                        _ = v83
                                    } else {
                                        v2 -= v2
                                        v84 := GenS4{F0: 4294967295, F1: genFn1((v5 << ((genFn0(fi74)) & 63)), 85, 202), F2: v2}
                                        v1 = (-2.25 != 0.1)
                                        _ = v84
                                    }
                                default:
                                    v85 := v4[(fi81 % ((fi74) | 1))]
                                    _ = v85
                                }
                            }
                            v1 = (!(v1 || (124 < 2147483648)))
                            fmt.Println(fi81, fi74)
                        }
                    }
                    _ = fi81
                    _ = v82
                }
                _ = fi74
            }
            var v86 int = 2147483647
            _ = v86
        } else {
            v3 = (v3 + 1)
            switch -2147483648 {
            case 128:
                v6 -= (v6 / (((v6 | v6)) | 1))
            case -9223372036854775808:
                if false {
                    if false {
                        if (32767 > -32768) {
                            if (!false) {
                                if (!(v1 || v1)) {
                                    if v1 {
                                        v2 = 1
                                        v6 -= (^(v6))
                                        for fi87 := 0; fi87 < 3; fi87++ {
                                            v88 := GenS2{F0: (fi87 & fi87), F1: v3}
                                            _ = fi87
                                            _ = v88
                                        }
                                    } else {
                                        v1 = (("" + "\n\t") > ("" + "hello"))
                                        var v89 int32 = (-(v2))
                                        v1 = (42 != 32767)
                                        _ = v89
                                    }
                                    v3 = v3
                                }
                                v6 = v6
                            }
                            v90 := []int{-421, 127, 2147483647, -2147483648}
                            v4[-1] = 833
                            _ = v90
                        }
                    } else {
                        for fi91 := 0; fi91 < 12; fi91++ {
                            v2 = v2
                            fmt.Println(v5, v2, v2)
                            _ = fi91
                        }
                    }
                } else {
                    var v92 uint8 = 128
                    _ = v92
                }
            default:
                v1 = true
            }
        }
    }
    var v93 uint = genFn0(-1)
    var v94 uint = v93
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
    _ = v93
    _ = v94
}
