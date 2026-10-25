package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 int }
func genFn0(p0 string, p1 int64) uint64 {
    p1 -= (p1 + p1)
    v4 := []int{42, -9223372036854775808, 4294967295}
    switch "" {
    case "hello":
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 uint = 42
            var v7 uint64 = 0
            var v8 bool = true
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
    case "foo bar":
        v9 := GenS1{F0: 2147483648, F1: 2}
        _ = v9
    case "\n\t":
        p0 = p0
    default:
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 string = ("foo bar" + p0)
            switch "hello" {
            case "foo bar":
                if true {
                    v12 := GenS2{F0: 180, F1: 128}
                    var v13 uint32 = 1
                    _ = v12
                    _ = v13
                }
            }
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 bool = ((false || (((true && (true || (!(((true && ((((3.14159 - 1.5) - ((1.5 / 3.14159) + 3.14159)) - (0.0 - ((3.14159 / 1.0) - 1e10))) != (0.0 / ((1e10 + -2.25) + 0.0)))) && true) || ((-9223372036854775808 < 0) && true))))) && (652 > 42)) || true)) && (-128 >= 0))
    var v2 uint8 = 128
    var v3 int64 = 42
    var v14 int16 = 898
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS2{F0: 2, F1: 42}
        v17 := GenS2{F0: 757, F1: v2}
        _ = fi15
        _ = v16
        _ = v17
    }
    switch ("" + "\n\t") {
    case "\n\t":
        var v18 uint = 42
        _ = v18
    case "a":
        var v19 int16 = (-(v14))
        _ = v19
    case "hello":
        var v20 int = 1
        _ = v20
    default:
        var v21 uint32 = 42
        _ = v21
    }
    if (-413 < 1) {
        var v22 int64 = (v3 * (v3 ^ 9223372036854775807))
        if v1 {
            var v23 int16 = 32767
            _ = v23
        } else {
            var v24 int8 = -128
            v25 := GenS3{F0: (-(v24)), F1: 1}
            var v26 uint = 0
            _ = v24
            _ = v25
            _ = v26
        }
        if v1 {
            v27 := []int{1}
            for fi28 := 0; fi28 < 17; fi28++ {
                switch ("foo bar" + "") {
                case "hello":
                    var v29 string = "foo bar"
                    _ = v29
                case "foo bar":
                    v14 += v14
                default:
                    var v30 float64 = (1e10 / ((0.1 / -2.25) / 1.0))
                    _ = v30
                }
                _ = fi28
            }
            _ = v27
        }
        _ = v22
    }
    if v1 {
        if (700 <= 265) {
            switch 2147483647 {
            case 128:
                v3 = v3
            default:
                for fi31 := 0; fi31 < 17; fi31++ {
                    v14 = v14
                    v32 := GenS4{F0: fi31, F1: 1}
                    _ = fi31
                    _ = v32
                }
            }
            v33 := GenS2{F0: 353, F1: 198}
            for fi34 := 0; fi34 < 12; fi34++ {
                v2 = (v2 - (^(v2)))
                _ = fi34
            }
            _ = v33
        }
        v35 := []int{}
        _ = v35
    } else {
        var v36 int8 = -128
        for fi37 := 0; fi37 < 14; fi37++ {
            v38 := []int{}
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    var v39 float32 = (((3.14159 / 1.0) / float32(0.1)) - (1.5 - 0.1))
    if v1 {
        var v40 int = -111
        if v1 {
            v40 += 9223372036854775807
            if (1 == 18446744073709551615) {
                for fi41 := 0; fi41 < 11; fi41++ {
                    v1 = true
                    switch fi41 {
                    case -9223372036854775808:
                        var v42 uint64 = 18446744073709551615
                        _ = v42
                    default:
                        if (float32(float64(3.14159)) < 0.0) {
                            v43 := []int{(fi41 & v40), v40, fi41}
                            v43 = append(v43, 42)
                            _ = v43
                        } else {
                            if ("\n\t" >= "") {
                                v39 += 3.14159
                                for fi44 := 0; fi44 < 16; fi44++ {
                                    v45 := []int{(v40 + v40), v40, (v40 | -795), v40}
                                    var v46 uint = 65535
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            for fi47 := 0; fi47 < 6; fi47++ {
                                v48 := []int{(^(v40)), fi47}
                                var v49 int = (v40 >> ((42) & 63))
                                if (813 == 403) {
                                    fmt.Println(fi47, v40)
                                }
                                _ = fi47
                                _ = v48
                                _ = v49
                            }
                            if false {
                                var v50 bool = (!(65535 > 1))
                                v2 = (v2 << ((255) & 63))
                                _ = v50
                            }
                        }
                    }
                    for fi51 := 0; fi51 < 12; fi51++ {
                        v52 := []int{(v40 % ((212) | 1))}
                        _ = fi51
                        _ = v52
                    }
                    _ = fi41
                }
            }
        } else {
            if (!(!true)) {
                v3 += (^(v3))
                var v53 uint64 = 4294967295
                _ = v53
            } else {
                var v54 float64 = 0.1
                var v55 int64 = v3
                v55 = -1
                _ = v54
                _ = v55
            }
            if (18446744073709551615 != 42) {
                v14 = v14
                v39 += v39
                if (v1 && true) {
                    if ((v40 >> ((1) & 63)) > 32768) {
                        v56 := GenS5{F0: 42, F1: "hello"}
                        v57 := []int{v40, v40, v40, -368}
                        v58 := GenS5{F0: (-(v3)), F1: "hello"}
                        _ = v56
                        _ = v57
                        _ = v58
                    } else {
                        if v1 {
                            v1 = true
                        } else {
                            if true {
                                v59 := GenS6{F0: "hello"}
                                var v60 float64 = ((-2.25 * (1e10 - float64((0.0 + ((((float32((1.5 + 1.5)) - (((3.14159 + 1e10) / v39) / 1.0)) * float32((float64((float32(1.5) - (v39 - v39))) / 1.0))) - 3.14159) * 1e10))))) - 0.0)
                                fmt.Println(v39)
                                _ = v59
                                _ = v60
                            }
                        }
                        if (v1 || (1e10 >= 0.0)) {
                            v39 = 3.14159
                            v61 := []int{-2147483648}
                            _ = v61
                        } else {
                            for fi62 := 0; fi62 < 1; fi62++ {
                                for fi63 := 0; fi63 < 13; fi63++ {
                                    var v64 string = (("\n\t" + ("a" + "\n\t")) + "")
                                    switch 128 {
                                    case -302:
                                        var v65 float32 = v39
                                        _ = v65
                                    case 283:
                                        v66 := GenS1{F0: 65535, F1: 172}
                                        _ = v66
                                    }
                                    v39 = float32(1.5)
                                    _ = fi63
                                    _ = v64
                                }
                                if (v1 && (v1 || (true || (!(!(0 < v3)))))) {
                                    var v67 uint32 = 65535
                                    _ = v67
                                } else {
                                    var v68 uint = 1
                                    _ = v68
                                }
                                _ = fi62
                            }
                        }
                    }
                    v40 -= (v40 >> ((65535) & 63))
                } else {
                    v40 -= v40
                }
            }
            fmt.Println(v2, v3, v40)
        }
        fmt.Println(v2, v14, v14)
        _ = v40
    } else {
        var v69 uint = 713
        fmt.Println(v69, v2, v1)
        var v70 uint8 = v2
        _ = v69
        _ = v70
    }
    fmt.Println(v1, v1)
    var v71 float32 = 3.14159
    switch 1 {
    case -9223372036854775808:
        v72 := []int{-2147483648}
        _ = v72
    default:
        v39 -= (v39 + -2.25)
    }
    var v73 uint64 = 2
    if (!false) {
        var v74 uint32 = 2
        v75 := []int{}
        _ = v74
        _ = v75
    } else {
        if false {
            var v76 float64 = (1e10 + -2.25)
            v77 := GenS5{F0: (v3 / ((490) | 1)), F1: (("a" + "") + "foo bar")}
            _ = v76
            _ = v77
        }
        v78 := GenS7{F0: 4294967295, F1: 32768}
        v79 := []int{32767, 32767, 1, 127}
        _ = v78
        _ = v79
    }
    v1 = true
    switch "a" {
    case "":
        if v1 {
            var v80 uint64 = (v73 & v73)
            var v81 uint = 255
            switch ("hello" + "foo bar") {
            case "hello":
                switch ("foo bar" + "") {
                case "":
                    v82 := GenS8{F0: 0}
                    _ = v82
                case "\n\t":
                    v83 := []int{}
                    _ = v83
                case "hello":
                    v1 = v1
                default:
                    for fi84 := 0; fi84 < 5; fi84++ {
                        v85 := GenS8{F0: -840}
                        var v86 float64 = 0.1
                        _ = fi84
                        _ = v85
                        _ = v86
                    }
                }
            }
            _ = v80
            _ = v81
        }
    case "\n\t":
        switch (("foo bar" + "hello") + "\n\t") {
        case "hello":
            v71 = -2.25
        case "\n\t":
            v71 -= v71
        }
    case "a":
        for fi87 := 0; fi87 < 14; fi87++ {
            if v1 {
                if ((float32((0.1 / (((0.0 + ((((((1.5 / 1.5) + 1.5) * (3.14159 / ((float64(0.1) + 1e10) - float64(v39)))) / 3.14159) * 1e10) * 0.0)) * 0.1) - 0.1))) * (v71 - ((v39 * float32((0.1 * (1e10 / (1.5 - (((float64((float32((float64(((float32((1.5 - 1.5)) + v71) * (0.0 + v71))) - (((-2.25 / ((-2.25 * 1e10) - ((((1e10 / 1.5) + 1.5) / (3.14159 - float64((((v71 + -2.25) + v71) / 3.14159)))) / 1e10))) * (((float64(v39) * 3.14159) * 1.5) / float64(v71))) / 0.1))) / v39)) * -2.25) / (0.1 / 1.5)) - float64(float32(float64(3.14159))))))))) * v39))) > v71) {
                    fmt.Println(v3, v2, v2)
                } else {
                    for fi88 := 0; fi88 < 18; fi88++ {
                        var v89 uint16 = uint16(v73)
                        var v90 uint32 = 165
                        switch fi88 {
                        case -457:
                            v3 = -220
                        }
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                    if (!(!v1)) {
                        fmt.Println(v73)
                        v91 := []int{}
                        _ = v91
                    } else {
                        v92 := GenS9{F0: fi87, F1: 32767}
                        v71 = v39
                        v14 += v14
                        _ = v92
                    }
                }
                v1 = (v1 && v1)
                v14 -= (v14 | (v14 * v14))
            }
            var v93 int32 = 42
            v14 = (v14 * v14)
            _ = fi87
            _ = v93
        }
    default:
        fmt.Println(v73, v39)
    }
    if v1 {
        for fi94 := 0; fi94 < 13; fi94++ {
            v3 = 9223372036854775807
            fmt.Println(fi94, v2, v14)
            if (!v1) {
                v73 = v73
                var v95 int = (fi94 % (((fi94 % ((fi94) | 1))) | 1))
                fmt.Println(v39)
                _ = v95
            }
            _ = fi94
        }
        v96 := []int{32768}
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v71
    _ = v73
}
