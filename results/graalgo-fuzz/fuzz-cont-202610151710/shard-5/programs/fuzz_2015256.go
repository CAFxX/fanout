package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int8; F1 uint64; F2 int32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 int; F2 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS12 struct { F0 uint64 }

func main() {
    var v1 uint = 255
    var v2 uint16 = 117
    v3 := make(map[int]string)
    var v4 uint32 = 1
    v5 := make(map[string]uint32)
    var v6 int64 = 28
    fmt.Println(v2)
    var v7 uint32 = (v4 >> ((v1) & 63))
    fmt.Println(v1, v1, v1)
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 11; fi9++ {
        v4 += v4
        for fi10 := 0; fi10 < 17; fi10++ {
            v6 = v6
            _ = fi10
        }
        _ = fi9
    }
    v11 := v3[32767]
    switch "hello" {
    case "hello":
        v12 := GenS1{F0: 32767, F1: 42, F2: 1}
        _ = v12
    case "foo bar":
        v13 := []int{-9223372036854775808, 128}
        _ = v13
    case "\n\t":
        v14 := GenS2{F0: v8, F1: (v8 & (v8 ^ v8)), F2: (("foo bar" + "") + ("a" + "a"))}
        _ = v14
    default:
        v4 -= 65535
    }
    if ((!((!false) && true)) || false) {
        v2 -= 2
        fmt.Println(v7, v4, v1)
    } else {
        if false {
            var v15 int = 32767
            if ((9 > 42) && (!false)) {
                var v16 int64 = 9223372036854775807
                _ = v16
            }
            if (255 != v1) {
                if true {
                    var v17 int16 = 1
                    _ = v17
                } else {
                    v18 := []int{(v15 >> ((v1) & 63)), v15}
                    if (uint8(v2) > 127) {
                        if (995 != (v1 / ((42) | 1))) {
                            v19 := GenS1{F0: 220, F1: 0, F2: 0}
                            _ = v19
                        }
                        var v20 uint32 = 1
                        _ = v20
                    } else {
                        v21 := GenS1{F0: -365, F1: 1000, F2: 18446744073709551615}
                        if (0 > (v8 / ((1) | 1))) {
                            v22 := GenS3{F0: (v4 | (^(v4))), F1: true, F2: 127}
                            v23 := GenS4{F0: (v4 << ((12) & 63))}
                            var v24 bool = true
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            var v25 uint64 = 0
                            if len(v18) > 0 {
                                v18[0] = -2147483648
                            }
                            _ = v25
                        }
                        v26 := GenS3{F0: 4294967295, F1: false, F2: 0}
                        _ = v21
                        _ = v26
                    }
                    _ = v18
                }
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 float64 = float64(3.14159)
                    _ = fi27
                    _ = v28
                }
                if (!true) {
                    var v29 float32 = (-2.25 / ((float32(0.0) - 0.1) + -2.25))
                    switch (v15 ^ v15) {
                    case 32767:
                        v30 := []int{v15}
                        _ = v30
                    case 32768:
                        switch v11 {
                        case "\n\t":
                            fmt.Println(v2, v2, v2)
                        case "hello":
                            switch (^(v15)) {
                            case 2147483647:
                                v4 += 42
                            case 4294967295:
                                if (!false) {
                                    switch v11 {
                                    case "foo bar":
                                        if true {
                                            switch (v15 ^ 2147483647) {
                                            case 42:
                                                v31 := GenS4{F0: v4}
                                                _ = v31
                                            case -1:
                                                v32 := GenS2{F0: (v8 >> ((0) & 63)), F1: (v8 | v8), F2: v11}
                                                _ = v32
                                            }
                                        } else {
                                            var v33 float32 = -2.25
                                            var v34 uint16 = (^(v2))
                                            _ = v33
                                            _ = v34
                                        }
                                    case "hello":
                                        v35 := []int{v15, v15, (v15 / ((811) | 1)), -9223372036854775808}
                                        _ = v35
                                    case "a":
                                        for fi36 := 0; fi36 < 18; fi36++ {
                                            switch v15 {
                                            case 0:
                                                for fi37 := 0; fi37 < 9; fi37++ {
                                                    fmt.Println(v6, fi37, v11)
                                                    v2 = (v2 * (v2 | uint16(v1)))
                                                    v38 := GenS1{F0: int16(v8), F1: 42, F2: 1}
                                                    _ = fi37
                                                    _ = v38
                                                }
                                            case 542:
                                                v39 := []int{(^(v15)), fi36, fi36}
                                                _ = v39
                                            case 128:
                                                fmt.Println(v4, v8, v8)
                                            }
                                            _ = fi36
                                        }
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        v29 += float32(((0.1 - 0.0) + 0.0))
                                    case "hello":
                                        v40 := GenS5{F0: v15, F1: -128}
                                        _ = v40
                                    }
                                }
                            case -1:
                                v41 := []int{}
                                _ = v41
                            default:
                                for fi42 := 0; fi42 < 2; fi42++ {
                                    switch ((("hello" + v11) + v11) + v11) {
                                    case "foo bar":
                                        switch v11 {
                                        case "a":
                                            v11 = v11
                                        case "":
                                            v43 := GenS6{F0: 622, F1: -1}
                                            _ = v43
                                        default:
                                            v2 -= (^(v2))
                                        }
                                    case "":
                                        v44 := []int{2147483647, fi42, (fi42 * v15), 4294967295}
                                        _ = v44
                                    default:
                                        if false {
                                            switch "foo bar" {
                                            case "\n\t":
                                                v6 -= 0
                                            case "":
                                                v45 := GenS4{F0: 65535}
                                                _ = v45
                                            default:
                                                var v46 uint64 = 18446744073709551615
                                                _ = v46
                                            }
                                            v47 := []int{fi42}
                                            for fi48 := 0; fi48 < 10; fi48++ {
                                                var v49 bool = false
                                                v7 = 2
                                                var v50 uint32 = 4294967295
                                                _ = fi48
                                                _ = v49
                                                _ = v50
                                            }
                                            _ = v47
                                        }
                                    }
                                    _ = fi42
                                }
                            }
                        case "a":
                            switch "" {
                            case "a":
                                var v51 uint8 = 128
                                _ = v51
                            default:
                                var v52 bool = (0 <= 2)
                                _ = v52
                            }
                        }
                    case 127:
                        var v53 int64 = -96
                        _ = v53
                    }
                    v8 = v8
                    _ = v29
                } else {
                    v54 := GenS7{F0: 42, F1: 1, F2: v8}
                    v1 += v1
                    _ = v54
                }
            }
            _ = v15
        } else {
            if (42 < 100) {
                var v55 uint32 = 4294967295
                if false {
                    v56 := GenS7{F0: 42, F1: 0, F2: 1}
                    _ = v56
                }
                _ = v55
            }
        }
        var v57 uint8 = 128
        if (1 <= -1) {
            v58 := GenS8{F0: 18446744073709551615}
            if (true && (42 > (v6 | (v6 + -597)))) {
                for fi59 := 0; fi59 < 3; fi59++ {
                    var v60 int8 = -128
                    _ = fi59
                    _ = v60
                }
                fmt.Println(v6, v11)
            } else {
                v61 := GenS9{F0: 0, F1: -9223372036854775808, F2: v57}
                var v62 uint32 = v4
                _ = v61
                _ = v62
            }
            v7 = 658
            _ = v58
        }
        _ = v57
    }
    for fi63 := 0; fi63 < 13; fi63++ {
        v5["\n\t"] = (v4 + (v7 & v7))
        var v64 float64 = -2.25
        _ = fi63
        _ = v64
    }
    fmt.Println(v2, v4)
    var v65 uint64 = 181
    var v66 float64 = (1e10 / -2.25)
    if (0.1 <= (1.5 / 3.14159)) {
        v67 := []int{4294967295, 78}
        _ = v67
    } else {
        v68 := GenS10{F0: 4294967295}
        v2 = uint16(v1)
        var v69 int = 861
        _ = v68
        _ = v69
    }
    _ = len(v3)
    var v70 float64 = v66
    if true {
        delete(v5, "foo bar")
        var v71 int = 0
        if false {
            if ((v71 - v71) != 729) {
                for fi72 := 0; fi72 < 13; fi72++ {
                    for fi73 := 0; fi73 < 4; fi73++ {
                        var v74 uint8 = 0
                        _ = fi73
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 4; fi75++ {
                        v7 -= (v7 / (((^(v4))) | 1))
                        v76 := []int{0, 526, (fi72 * (v71 ^ 32768))}
                        _ = fi75
                        _ = v76
                    }
                    for fi77 := 0; fi77 < 4; fi77++ {
                        v78 := GenS11{F0: 525, F1: v4, F2: v65}
                        v11 = (v11 + v11)
                        _ = fi77
                        _ = v78
                    }
                    _ = fi72
                }
            } else {
                for fi79 := 0; fi79 < 3; fi79++ {
                    var v80 uint8 = 255
                    _ = fi79
                    _ = v80
                }
                if true {
                    fmt.Println(v71, v2)
                } else {
                    v2 = 65535
                    if (!false) {
                        for fi81 := 0; fi81 < 11; fi81++ {
                            fmt.Println(v71, v65, v1)
                            v71 = (fi81 << (((v1 & v1)) & 63))
                            var v82 int32 = (v8 / ((-139) | 1))
                            _ = fi81
                            _ = v82
                        }
                        var v83 uint8 = 79
                        _ = v83
                    }
                    v1 -= (^(v1))
                }
            }
        }
        _ = v71
    } else {
        for fi84 := 0; fi84 < 7; fi84++ {
            switch (v11 + (v11 + "")) {
            case "":
                v65 += (v65 % ((42) | 1))
            case "\n\t":
                switch (fi84 << ((1) & 63)) {
                case 4294967295:
                    var v85 float64 = float64(1.5)
                    _ = v85
                case 974:
                    var v86 uint32 = 4294967295
                    _ = v86
                default:
                    v87 := v3[fi84]
                    _ = v87
                }
            case "foo bar":
                v4 = (v7 - v4)
            default:
                delete(v5, v11)
            }
            _ = fi84
        }
        for fi88 := 0; fi88 < 11; fi88++ {
            v89 := GenS12{F0: (v65 ^ 1)}
            v90 := []int{fi88, (fi88 - (fi88 << ((42) & 63))), 42, fi88}
            _ = fi88
            _ = v89
            _ = v90
        }
    }
    v91 := GenS9{F0: 1, F1: 1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v7
    _ = v70
    _ = v8
    _ = v91
}
