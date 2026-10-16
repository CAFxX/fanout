package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int; F2 uint16 }
type GenS7 struct { F0 uint64; F1 int64; F2 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int; F1 int64; F2 uint8 }
type GenS11 struct { F0 int16; F1 int64; F2 string }
func genFn0() uint32 {
    if (65535 > 221) {
        var v4 uint32 = 42
        v4 = (v4 >> ((0) & 63))
        _ = v4
    }
    switch "\n\t" {
    case "a":
        v5 := GenS1{F0: 9223372036854775807, F1: 80, F2: 100}
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 7; fi6++ {
            v7 := []int{fi6}
            fmt.Println(fi6)
            for fi8 := 0; fi8 < 1; fi8++ {
                if true {
                    v9 := GenS1{F0: (fi6 - (-(fi8))), F1: 255, F2: 100}
                    switch ("foo bar" + "a") {
                    case "a":
                        var v10 uint = 324
                        _ = v10
                    case "hello":
                        var v11 string = "foo bar"
                        _ = v11
                    default:
                        switch fi8 {
                        case 0:
                            var v12 int16 = 42
                            _ = v12
                        case 128:
                            fmt.Println(fi8, fi6)
                        case 32768:
                            v13 := GenS2{F0: -417, F1: 18446744073709551615, F2: 54}
                            _ = v13
                        }
                    }
                    var v14 uint = 4294967295
                    _ = v14
                    _ = v9
                } else {
                    v15 := GenS1{F0: (^(fi8)), F1: 2, F2: -128}
                    var v16 int16 = 1000
                    _ = v15
                    _ = v16
                }
                var v17 float64 = (((1.5 + (1e10 / 0.1)) / float64((float32(float64(((3.14159 - 0.0) + 1e10))) - 1e10))) * 3.14159)
                _ = fi8
                _ = v17
            }
            _ = fi6
            _ = v7
        }
    case "\n\t":
        var v18 int16 = 32767
        _ = v18
    default:
        var v19 int64 = -1
        _ = v19
    }
    return 1
}
func genFn1() int32 {
    v20 := GenS3{F0: 2147483648, F1: 0}
    v21 := []int{32768, 42, 127, 32767}
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 string = ("hello" + "")
    var v2 uint = 18446744073709551615
    var v3 float64 = (1.5 - ((float64((1.5 * -2.25)) / (float64(((0.1 + 1e10) / ((((0.0 * 1e10) / 1.0) + 0.1) / float32(-2.25)))) + (0.1 - (1.5 + 1e10)))) * -2.25))
    if (!true) {
        if false {
            for fi22 := 0; fi22 < 9; fi22++ {
                for fi23 := 0; fi23 < 7; fi23++ {
                    if (true || false) {
                        v1 = (v1 + (v1 + "a"))
                        v24 := GenS2{F0: -968, F1: 1, F2: 42}
                        v25 := GenS3{F0: 0, F1: 18446744073709551615}
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
                var v26 float32 = ((0.1 + -2.25) / 0.1)
                fmt.Println(v1, fi22, v1)
                _ = fi22
                _ = v26
            }
        }
    } else {
        switch (v1 + (v1 + "\n\t")) {
        case "a":
            v3 += 1.5
        case "foo bar":
            if (630 >= 4294967295) {
                v27 := []int{1, -9223372036854775808}
                if len(v27) > 0 {
                    v27[0] = 9223372036854775807
                }
                _ = v27
            }
        default:
            var v28 int8 = -1
            _ = v28
        }
    }
    if false {
        if true {
            v29 := GenS1{F0: 32767, F1: (v2 ^ v2), F2: 127}
            switch v1 {
            case "":
                if (910 > -1) {
                    switch 127 {
                    case 486:
                        v30 := []int{-363, 4294967295}
                        _ = v30
                    case 1:
                        v2 = v2
                    case 2147483647:
                        switch (v1 + v1) {
                        case "foo bar":
                            switch v1 {
                            case "a":
                                v2 = 0
                            case "\n\t":
                                var v31 uint64 = 149
                                _ = v31
                            case "foo bar":
                                v32 := GenS4{F0: (v1 + (v1 + v1))}
                                _ = v32
                            default:
                                if (42 <= 506) {
                                    v1 = v1
                                    if (v3 <= v3) {
                                        for fi33 := 0; fi33 < 10; fi33++ {
                                            fmt.Println(v3, v2, fi33)
                                            switch v1 {
                                            case "foo bar":
                                                var v34 string = v1
                                                _ = v34
                                            case "":
                                                for fi35 := 0; fi35 < 11; fi35++ {
                                                    v36 := GenS3{F0: 2147483648, F1: 701}
                                                    _ = fi35
                                                    _ = v36
                                                }
                                            case "a":
                                                v37 := []int{(fi33 & (fi33 % ((fi33) | 1))), -1}
                                                _ = v37
                                            default:
                                                if false {
                                                    for fi38 := 0; fi38 < 2; fi38++ {
                                                        v2 = (v2 % (((v2 ^ 65535)) | 1))
                                                        _ = fi38
                                                    }
                                                    v2 = (v2 + v2)
                                                }
                                            }
                                            switch "a" {
                                            case "hello":
                                                var v39 int16 = 698
                                                _ = v39
                                            default:
                                                var v40 float64 = 1e10
                                                _ = v40
                                            }
                                            _ = fi33
                                        }
                                        for fi41 := 0; fi41 < 3; fi41++ {
                                            if (!false) {
                                                for fi42 := 0; fi42 < 12; fi42++ {
                                                    if (18446744073709551615 == 274) {
                                                        var v43 int8 = 1
                                                        _ = v43
                                                    }
                                                    v44 := []int{fi41, fi42, fi41}
                                                    v1 = v1
                                                    _ = fi42
                                                    _ = v44
                                                }
                                                v45 := []int{(fi41 - (fi41 << ((v2) & 63))), fi41, 1}
                                                _ = v45
                                            } else {
                                                var v46 int64 = 9223372036854775807
                                                var v47 uint16 = 32767
                                                _ = v46
                                                _ = v47
                                            }
                                            v1 = "a"
                                            if false {
                                                fmt.Println(v3, v3)
                                                if true {
                                                    v48 := []int{128, (fi41 % ((fi41) | 1)), (-(fi41)), (fi41 / ((fi41) | 1))}
                                                    _ = v48
                                                } else {
                                                    for fi49 := 0; fi49 < 11; fi49++ {
                                                        var v50 int16 = -1
                                                        var v51 float32 = (3.14159 + (1.5 * ((1e10 - float32(v3)) * 0.0)))
                                                        _ = fi49
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                }
                                            } else {
                                                v2 -= v2
                                                v3 = float64(-2.25)
                                            }
                                            _ = fi41
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(v1, v3, v3)
                                        }
                                        v52 := GenS5{F0: 774}
                                        _ = v52
                                    }
                                }
                            }
                        default:
                            var v53 int32 = 1
                            _ = v53
                        }
                    default:
                        switch ((v1 + ("foo bar" + "hello")) + "a") {
                        case "\n\t":
                            var v54 uint32 = 4294967295
                            _ = v54
                        case "hello":
                            fmt.Println(v1, v2)
                        default:
                            v3 += 3.14159
                        }
                    }
                    var v55 bool = ((v3 + 1e10) == v3)
                    if true {
                        var v56 bool = v55
                        _ = v56
                    } else {
                        v57 := GenS1{F0: 1, F1: v2, F2: -1}
                        _ = v57
                    }
                    _ = v55
                }
            case "foo bar":
                v2 = (v2 % ((v2) | 1))
            case "\n\t":
                v58 := GenS5{F0: 777}
                _ = v58
            default:
                v1 = (v1 + v1)
            }
            _ = v29
        }
        switch v1 {
        case "foo bar":
            if true {
                v59 := []int{42, 4294967295}
                if true {
                    for fi60 := 0; fi60 < 3; fi60++ {
                        var v61 uint = (v2 & v2)
                        _ = fi60
                        _ = v61
                    }
                } else {
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v59 = append(v59, (fi62 ^ (fi62 & fi62)))
                        _ = fi62
                    }
                }
                v1 = "hello"
                _ = v59
            }
        case "hello":
            if false {
                v63 := []int{127, 32767}
                var v64 bool = false
                _ = len(v63)
                _ = v63
                _ = v64
            }
        default:
            fmt.Println(v3)
        }
        v65 := []int{955}
        _ = v65
    } else {
        if true {
            if true {
                var v66 int8 = 100
                for fi67 := 0; fi67 < 15; fi67++ {
                    var v68 int64 = 42
                    _ = fi67
                    _ = v68
                }
                var v69 int64 = 446
                _ = v66
                _ = v69
            }
            for fi70 := 0; fi70 < 6; fi70++ {
                fmt.Println(v2, v3)
                var v71 int8 = 8
                _ = fi70
                _ = v71
            }
        }
        v72 := GenS4{F0: (v1 + v1)}
        v2 = (v2 >> ((v2) & 63))
        _ = v72
    }
    v2 += v2
    fmt.Println(v2, v2)
    if (!true) {
        var v73 uint32 = 42
        fmt.Println(v3, v2)
        switch v1 {
        case "foo bar":
            v74 := []int{826}
            _ = v74
        case "hello":
            v75 := []int{-36}
            _ = v75
        }
        _ = v73
    }
    if false {
        var v76 int8 = -1
        switch -650 {
        case -522:
            v77 := GenS6{F0: 0, F1: 2147483647, F2: 1}
            _ = v77
        case 0:
            v78 := GenS7{F0: 877, F1: 1, F2: (v2 / (((v2 * 18446744073709551615)) | 1))}
            _ = v78
        default:
            var v79 uint32 = 42
            _ = v79
        }
        switch v1 {
        case "foo bar":
            fmt.Println(v76)
        case "hello":
            v80 := []int{}
            _ = v80
        case "\n\t":
            var v81 int16 = -687
            _ = v81
        }
        _ = v76
    }
    v3 = float64(0.0)
    v82 := []int{}
    if len(v82) >= 2 {
        v82 = v82[1:len(v82)]
    }
    if true {
        if (false || true) {
            for fi83 := 0; fi83 < 13; fi83++ {
                if false {
                    for fi84 := 0; fi84 < 14; fi84++ {
                        if (!(1 < 342)) {
                            v85 := GenS8{F0: 255}
                            _ = v85
                        }
                        v86 := GenS9{F0: 42}
                        _ = fi84
                        _ = v86
                    }
                    var v87 float64 = ((v3 * float64(float32((v3 - 0.0)))) * (0.1 * v3))
                    v3 = (v3 + 0.1)
                    _ = v87
                } else {
                    if len(v82) >= 2 {
                        v82 = v82[1:len(v82)]
                    }
                    for fi88 := 0; fi88 < 2; fi88++ {
                        var v89 uint32 = 1
                        if (false && false) {
                            v2 = (v2 | v2)
                        } else {
                            var v90 string = (("hello" + ("a" + v1)) + "")
                            fmt.Println(v3)
                            v89 = (v89 + 2)
                            fmt.Println(v90, fi88, v2)
                            var v91 string = "\n\t"
                            v92 := GenS10{F0: (fi88 ^ 127), F1: 348, F2: 196}
                            _ = v90
                            _ = v91
                            _ = v92
                        }
                        v93 := GenS11{F0: -551, F1: -1, F2: v1}
                        _ = fi88
                        _ = v89
                        _ = v93
                    }
                }
                v94 := GenS11{F0: 42, F1: 1, F2: v1}
                if true {
                    v1 = v1
                    var v95 bool = (!false)
                    _ = v95
                }
                _ = fi83
                _ = v94
            }
            switch v1 {
            case "\n\t":
                var v96 float32 = (-2.25 - 1e10)
                _ = v96
            case "a":
                var v97 int32 = -235
                _ = v97
            }
        }
    } else {
        fmt.Println(v3)
        var v98 uint32 = 2
        if true {
            v99 := GenS6{F0: 128, F1: 9223372036854775807, F2: 1}
            _ = v99
        }
        _ = v98
    }
    v100 := GenS6{F0: 1, F1: 468, F2: 712}
    v1 = v1
    for fi101 := 0; fi101 < 8; fi101++ {
        fmt.Println(fi101, v2, fi101)
        v1 = "\n\t"
        _ = fi101
    }
    v1 = ((v1 + "hello") + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v82:", len(v82), v82)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v82
}
