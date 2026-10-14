package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint8, p1 uint) bool {
    var v7 uint = 18446744073709551615
    var v8 uint32 = 2147483647
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 int64 = 1
        v11 := []int{}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: 32767}
        var v14 int = fi12
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int64) int {
    var v15 uint32 = 28
    if false {
        v15 = 2147483648
        var v16 int8 = 108
        v17 := GenS1{F0: 448}
        _ = v16
        _ = v17
    }
    p0 = -1
    var v18 bool = true
    if false {
        fmt.Println(p0, v15, v18)
        if true {
            v19 := GenS2{F0: -338, F1: 371, F2: -137}
            v15 = v15
            _ = v19
        } else {
            v20 := []int{639, 128, 9223372036854775807}
            _ = v20
        }
        var v21 float32 = 3.14159
        _ = v21
    }
    _ = p0
    _ = v15
    _ = v18
    return 2147483647
}

func main() {
    var v1 string = "hello"
    var v2 int = -506
    var v3 float64 = 0.1
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    var v6 int64 = 9223372036854775807
    switch v1 {
    case "hello":
        switch (((v1 + v1) + "a") + "") {
        case "\n\t":
            switch (v2 & (v2 - v2)) {
            case 856:
                v2 = v2
            case 564:
                switch ((v1 + v1) + (v1 + ((v1 + v1) + v1))) {
                case "":
                    v6 += v6
                case "foo bar":
                    fmt.Println(v6, v2, v6)
                }
            }
        }
    }
    switch v1 {
    case "hello":
        fmt.Println(v3)
    }
    switch v1 {
    case "a":
        v2 += v2
    default:
        v6 -= v6
    }
    switch v1 {
    case "\n\t":
        var v22 int8 = 1
        _ = v22
    case "":
        v3 = (0.0 + ((v3 * float64(0.1)) - v3))
    default:
        if true {
            if (1000 >= 1000) {
                fmt.Println(v1, v6)
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 uint8 = 127
                    _ = fi23
                    _ = v24
                }
                fmt.Println(v3, v6)
            } else {
                v25 := v4[1]
                switch ("\n\t" + v25) {
                case "hello":
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v2 -= (v2 + 2147483647)
                        _ = len(v4)
                        _ = fi26
                    }
                case "a":
                    if ((1 != 0) || (true || true)) {
                        var v27 uint16 = 768
                        _ = v27
                    } else {
                        var v28 uint8 = 127
                        var v29 uint16 = 42
                        if (110 > -871) {
                            v6 = v6
                            v30 := GenS3{F0: (-1 >= 0), F1: 32, F2: 41}
                            switch v1 {
                            case "\n\t":
                                v31 := GenS3{F0: ((32768 == 127) || false), F1: 762, F2: 43}
                                _ = v31
                            default:
                                if ((!false) && true) {
                                    v2 -= genFn1((v6 << ((2) & 63)))
                                } else {
                                    if false {
                                        var v32 int = v2
                                        v33 := GenS4{F0: 2147483648}
                                        _ = v32
                                        _ = v33
                                    }
                                }
                            }
                            _ = v30
                        } else {
                            var v34 uint = 42
                            v28 = (^(v28))
                            v35 := GenS3{F0: ("hello" <= v1), F1: 4294967295, F2: 100}
                            _ = v34
                            _ = v35
                        }
                        _ = v28
                        _ = v29
                    }
                }
                if (!true) {
                    var v36 uint = 18446744073709551615
                    _ = v36
                } else {
                    var v37 uint32 = 65535
                    var v38 uint = 255
                    if ((960 <= 0) && false) {
                        fmt.Println(v38)
                    } else {
                        v39 := []int{(v2 >> (((^(v38))) & 63)), genFn1((v6 | (v6 & -1)))}
                        _ = v39
                    }
                    _ = v37
                    _ = v38
                }
                _ = v25
            }
            v1 = "foo bar"
            v40 := GenS1{F0: 65535}
            _ = v40
        } else {
            v2 += v2
            var v41 uint16 = 32767
            _ = v41
        }
    }
    v1 = (((v1 + v1) + (v1 + v1)) + v1)
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 uint16 = 32767
        _ = fi42
        _ = v43
    }
    if (42 != -64) {
        v1 = (v1 + (((v1 + "foo bar") + ("a" + v1)) + v1))
    } else {
        v3 = 0.0
        if (0 != 42) {
            var v44 uint64 = 2
            v45 := []int{(v2 % ((v2) | 1)), (v2 >> ((4294967295) & 63)), 1}
            for fi46 := 0; fi46 < 5; fi46++ {
                var v47 uint8 = 0
                _ = fi46
                _ = v47
            }
            _ = v44
            _ = v45
        }
    }
    switch "a" {
    case "foo bar":
        for fi48 := 0; fi48 < 5; fi48++ {
            if (1 != 911) {
                for fi49 := 0; fi49 < 7; fi49++ {
                    v50 := GenS4{F0: 2}
                    _ = fi49
                    _ = v50
                }
                if (!((25 >= 2) && (false && true))) {
                    v51 := []int{}
                    v2 -= genFn1((^(v6)))
                    _ = v51
                }
            }
            _ = fi48
        }
    case "hello":
        fmt.Println(v2)
    }
    var v52 int64 = v6
    switch "\n\t" {
    case "":
        v53 := []int{v2}
        _ = v53
    case "a":
        var v54 uint32 = 65535
        _ = v54
    }
    switch v2 {
    case -2147483648:
        var v55 int64 = -434
        _ = v55
    default:
        if true {
            if false {
                if (131 == 106) {
                    v6 = v52
                    for fi56 := 0; fi56 < 1; fi56++ {
                        var v57 int8 = -69
                        v58 := []int{(fi56 * (fi56 * genFn1(v6))), -978, fi56, (fi56 / ((fi56) | 1))}
                        v1 = "\n\t"
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                } else {
                    v2 -= (v2 % ((127) | 1))
                    v6 -= (v52 / (((v6 & 0)) | 1))
                }
                for fi59 := 0; fi59 < 3; fi59++ {
                    v60 := GenS2{F0: 142, F1: 485, F2: 2147483647}
                    if (42 > 1) {
                        if true {
                            for fi61 := 0; fi61 < 2; fi61++ {
                                fmt.Println(fi59, fi61, v1)
                                var v62 int64 = v6
                                delete(v5, ((v1 + v1) + v1))
                                _ = fi61
                                _ = v62
                            }
                            var v63 int8 = int8(v2)
                            var v64 bool = ((0.1 - (0.1 + -2.25)) != 3.14159)
                            _ = v63
                            _ = v64
                        }
                    } else {
                        for fi65 := 0; fi65 < 6; fi65++ {
                            v6 += (-(v6))
                            if (0 > 614) {
                                switch v2 {
                                case 4294967295:
                                    if (!(128 == 26)) {
                                        v66 := GenS5{F0: 32767}
                                        var v67 uint32 = 1
                                        v67 = (v67 & v67)
                                        _ = v66
                                        _ = v67
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                v68 := GenS3{F0: true, F1: 4294967295, F2: -1}
                                                _ = v68
                                            } else {
                                                if (0.0 <= 0.0) {
                                                    v69 := GenS4{F0: 2147483648}
                                                    _ = v69
                                                }
                                                v2 = 0
                                            }
                                        default:
                                            if false {
                                                if false {
                                                    var v70 uint8 = 138
                                                    v71 := GenS6{F0: 42}
                                                    _ = v70
                                                    _ = v71
                                                }
                                            }
                                        }
                                        if true {
                                            if (128 < 22) {
                                                fmt.Println(v6)
                                            }
                                            if false {
                                                v3 = v3
                                            }
                                            v2 += genFn1((v6 + -1))
                                            if false {
                                                v72 := []int{(fi59 * (fi65 | -2147483648)), (v2 % ((genFn1(-9223372036854775808)) | 1)), (fi59 | (v2 - (fi59 + genFn1((-(v6)))))), fi65}
                                                v6 = (v6 - v52)
                                                _ = v72
                                            }
                                            if false {
                                                var v73 string = v1
                                                var v74 string = "\n\t"
                                                _ = v73
                                                _ = v74
                                            }
                                        } else {
                                            v1 = ("a" + "\n\t")
                                            v4[476] = (v1 + v1)
                                        }
                                        delete(v5, v1)
                                        if false {
                                            switch (fi59 - (-(v2))) {
                                            case 193:
                                                v75 := GenS5{F0: 32767}
                                                _ = v75
                                            case 128:
                                                var v76 string = ""
                                                _ = v76
                                            case 645:
                                                var v77 bool = (42 >= 42)
                                                _ = v77
                                            default:
                                                var v78 uint32 = 531
                                                _ = v78
                                            }
                                            fmt.Println(v1, v2, v52)
                                            if (127 == 2) {
                                                var v79 bool = (true || (v52 <= v52))
                                                var v80 int64 = v52
                                                var v81 int32 = 566
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            } else {
                                                var v82 int32 = -2147483648
                                                v52 = v52
                                                fmt.Println(v6, fi65, v6)
                                                _ = v82
                                            }
                                        }
                                    }
                                case 283:
                                    if true {
                                        v3 += v3
                                        fmt.Println(v52)
                                    } else {
                                        v1 = v1
                                        v83 := GenS7{F0: "foo bar", F1: 191}
                                        _ = v83
                                    }
                                default:
                                    var v84 int = fi65
                                    _ = v84
                                }
                                var v85 uint = 0
                                _ = v85
                            }
                            var v86 int32 = -245
                            _ = fi65
                            _ = v86
                        }
                    }
                    _ = fi59
                    _ = v60
                }
                v87 := GenS5{F0: 930}
                _ = v87
            } else {
                fmt.Println(v52, v2, v52)
                switch (^(v2)) {
                case 212:
                    switch ((v1 + v1) + "hello") {
                    case "a":
                        v2 += (^(v2))
                    case "hello":
                        v2 -= (v2 - (v2 * genFn1((v52 & v52))))
                    default:
                        var v88 uint16 = 42
                        _ = v88
                    }
                case 235:
                    var v89 uint16 = 0
                    _ = v89
                case -305:
                    if true {
                        v52 += (v6 & (v52 / (((v52 - v52)) | 1)))
                        v90 := GenS7{F0: (v1 + v1), F1: 4294967295}
                        _ = v90
                    }
                default:
                    fmt.Println(v6, v52)
                }
                for fi91 := 0; fi91 < 15; fi91++ {
                    var v92 int64 = 9223372036854775807
                    _ = fi91
                    _ = v92
                }
            }
        } else {
            v93 := GenS8{F0: 40, F1: 1}
            v94 := GenS7{F0: "hello", F1: 2}
            _ = v93
            _ = v94
        }
    }
    if (true && (!(42 != 99))) {
        var v95 bool = true
        var v96 uint32 = 15
        if ((v96 << ((0) & 63)) != (v96 | 2147483647)) {
            fmt.Println(v1)
            if false {
                for fi97 := 0; fi97 < 5; fi97++ {
                    _ = len(v5)
                    _ = fi97
                }
            }
            for fi98 := 0; fi98 < 17; fi98++ {
                var v99 uint64 = 2
                _ = fi98
                _ = v99
            }
        }
        _ = v95
        _ = v96
    }
    if (0 == 1) {
        _ = len(v4)
    }
    var v100 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
