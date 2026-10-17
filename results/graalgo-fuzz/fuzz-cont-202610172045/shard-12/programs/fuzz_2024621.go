package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint16; F2 int16 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int; F1 string; F2 string }
type GenS10 struct { F0 string; F1 uint32 }
type GenS11 struct { F0 int16; F1 int }
func genFn0(p0 uint64, p1 int32, p2 string) string {
    fmt.Println(p1)
    v6 := GenS1{F0: -2147483648, F1: -743}
    fmt.Println(p0, p1, p0)
    v7 := GenS2{F0: -948, F1: 410}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = (v1 - v1)
    var v3 int16 = -369
    var v4 int8 = 100
    var v5 string = "\n\t"
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v2, v4, v2)
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int = 862
            _ = fi9
            _ = v10
        }
        v3 += (v3 | v3)
        _ = fi8
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v1 = (v2 << ((0) & 63))
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int16 = v3
            var v14 uint16 = 57
            var v15 uint32 = 65535
            var v16 uint = 255
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v16
        }
        if true {
            var v17 float32 = ((0.0 - ((0.1 + 0.1) / float32(((-2.25 + 3.14159) - -2.25)))) / 0.1)
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS3{F0: 255, F1: 1, F2: 652}
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = fi11
    }
    v20 := []int{0, -979, -1}
    var v21 int16 = v3
    var v22 uint8 = v2
    fmt.Println(v5)
    switch (v5 + ((v5 + v5) + v5)) {
    case "\n\t":
        v23 := GenS4{F0: "\n\t"}
        _ = v23
    case "a":
        v24 := GenS4{F0: v5}
        _ = v24
    case "":
        switch -258 {
        case 1:
            var v25 uint8 = 2
            _ = v25
        case 42:
            _ = len(v20)
        case 32767:
            if ((((v5 + "hello") + v5) + "\n\t") <= v5) {
                var v26 uint32 = 42
                if false {
                    var v27 uint8 = (v2 | (v1 >> ((uint(v2)) & 63)))
                    _ = len(v20)
                    v26 = (^(v26))
                    _ = v27
                }
                _ = v26
            } else {
                var v28 uint = 0
                _ = v28
            }
        }
    default:
        if true {
            if len(v20) > 0 {
                v20[0] = 396
            }
            var v29 int32 = 1
            _ = v29
        } else {
            v30 := GenS5{F0: v5, F1: uint16(v22), F2: v21}
            var v31 int64 = 1
            var v32 uint64 = 4294967295
            _ = v30
            _ = v31
            _ = v32
        }
    }
    if (false || false) {
        var v33 float64 = 0.0
        _ = v33
    }
    var v34 float32 = (1.5 - 0.1)
    for fi35 := 0; fi35 < 6; fi35++ {
        if false {
            for fi36 := 0; fi36 < 20; fi36++ {
                if ((v34 >= (float32(float64(v34)) * ((-2.25 * (1.5 / v34)) - (float32(0.0) / float32((3.14159 / -2.25)))))) || false) {
                    if false {
                        var v37 int8 = (v4 + 0)
                        v20 = append(v20, 1)
                        var v38 uint8 = (v22 << ((18446744073709551615) & 63))
                        _ = v37
                        _ = v38
                    }
                    switch -2147483648 {
                    case 791:
                        v34 += 3.14159
                    default:
                        var v39 uint16 = 2
                        _ = v39
                    }
                    v2 -= 127
                } else {
                    switch (fi36 / ((fi35) | 1)) {
                    case -897:
                        switch (fi35 % (((fi36 | 0)) | 1)) {
                        case -2147483648:
                            fmt.Println(v34, v21, v21)
                        default:
                            switch (fi36 | (fi35 ^ (fi36 & fi35))) {
                            case 0:
                                v3 += v21
                            default:
                                if (68 < 42) {
                                    switch fi35 {
                                    case -2147483648:
                                        var v40 uint32 = 255
                                        _ = v40
                                    case -325:
                                        fmt.Println(v22, v3, v34)
                                    case 808:
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                    default:
                                        switch (fi35 & fi36) {
                                        case -9223372036854775808:
                                            fmt.Println(v1)
                                        case 4294967295:
                                            var v41 int64 = 9223372036854775807
                                            _ = v41
                                        }
                                    }
                                    v42 := GenS2{F0: -333, F1: 971}
                                    _ = v42
                                } else {
                                    v2 -= (^(v22))
                                }
                            }
                        }
                    case 42:
                        if ((!false) || true) {
                            var v43 int8 = v4
                            _ = v43
                        }
                    }
                }
                v44 := GenS6{F0: (v3 * v3), F1: -410}
                _ = fi36
                _ = v44
            }
            switch int(v1) {
            case -2147483648:
                var v45 int32 = -2147483648
                _ = v45
            case 128:
                switch (-(fi35)) {
                case 1:
                    var v46 float32 = ((0.0 * v34) + 1.5)
                    _ = v46
                case -2147483648:
                    switch ("foo bar" + v5) {
                    case "a":
                        var v47 float32 = float32((0.0 + -2.25))
                        _ = v47
                    case "hello":
                        v48 := GenS7{F0: 42, F1: 1}
                        _ = v48
                    case "":
                        var v49 bool = false
                        _ = v49
                    default:
                        v3 = v21
                    }
                case -1:
                    v2 = v1
                }
            default:
                switch "\n\t" {
                case "":
                    _ = len(v20)
                case "a":
                    if true {
                        var v50 bool = (!(((true || false) && true) && false))
                        _ = v50
                    } else {
                        if false {
                            v5 = "\n\t"
                            v3 = int16(fi35)
                        }
                        var v51 int8 = v4
                        v1 -= v2
                        _ = v51
                    }
                }
            }
            v21 = 32767
        } else {
            if (!true) {
                var v52 float64 = 1e10
                switch fi35 {
                case 9223372036854775807:
                    v5 = "a"
                case 32768:
                    for fi53 := 0; fi53 < 7; fi53++ {
                        var v54 int64 = -9223372036854775808
                        fmt.Println(fi53, v34, v1)
                        v1 -= v1
                        _ = fi53
                        _ = v54
                    }
                case 1:
                    if true {
                        v2 -= 42
                    } else {
                        if (v22 <= 192) {
                            fmt.Println(v21, v4)
                            var v55 string = "hello"
                            _ = v55
                        } else {
                            for fi56 := 0; fi56 < 17; fi56++ {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                v34 += v34
                                var v57 float64 = (((v52 + float64((v34 / float32(v52)))) + ((v52 * v52) + v52)) / v52)
                                _ = fi56
                                _ = v57
                            }
                            v58 := GenS8{F0: v4, F1: 887}
                            _ = v58
                        }
                        var v59 int8 = -1
                        for fi60 := 0; fi60 < 13; fi60++ {
                            v61 := GenS6{F0: 250, F1: 0}
                            if len(v20) > 0 {
                                v20[0] = 32768
                            }
                            _ = fi60
                            _ = v61
                        }
                        _ = v59
                    }
                }
                _ = v52
            }
            fmt.Println(v21)
        }
        v21 += v3
        _ = fi35
    }
    for fi62 := 0; fi62 < 5; fi62++ {
        v34 += v34
        for fi63 := 0; fi63 < 1; fi63++ {
            v64 := GenS8{F0: 47, F1: 2147483647}
            v65 := GenS9{F0: fi62, F1: v5, F2: (v5 + (v5 + v5))}
            _ = len(v20)
            _ = fi63
            _ = v64
            _ = v65
        }
        _ = fi62
    }
    var v66 uint16 = 32767
    for fi67 := 0; fi67 < 7; fi67++ {
        var v68 int32 = 42
        var v69 int8 = 100
        _ = fi67
        _ = v68
        _ = v69
    }
    switch ((("foo bar" + v5) + v5) + "\n\t") {
    case "a":
        v70 := GenS3{F0: v22, F1: 205, F2: (^(v66))}
        _ = v70
    case "foo bar":
        var v71 uint32 = 65535
        _ = v71
    }
    if ((42 > 1) && true) {
        if false {
            for fi72 := 0; fi72 < 3; fi72++ {
                v20 = append(v20, 0)
                var v73 int8 = 72
                _ = fi72
                _ = v73
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        } else {
            var v74 uint = 0
            if true {
                v22 -= 1
                var v75 int16 = (v21 - v3)
                _ = v75
            } else {
                if true {
                    switch v5 {
                    case "\n\t":
                        v76 := GenS10{F0: ((v5 + ((v5 + v5) + v5)) + ("a" + v5)), F1: 612}
                        _ = v76
                    case "foo bar":
                        var v77 uint = (v74 & 65535)
                        _ = v77
                    default:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                }
                if false {
                    switch "" {
                    case "\n\t":
                        for fi78 := 0; fi78 < 9; fi78++ {
                            fmt.Println(v74)
                            var v79 uint32 = 2
                            _ = fi78
                            _ = v79
                        }
                    case "foo bar":
                        v4 -= v4
                    default:
                        v5 = (v5 + ("" + "\n\t"))
                    }
                }
            }
            var v80 float32 = 1e10
            _ = v74
            _ = v80
        }
        switch 558 {
        case 32767:
            if false {
                var v81 uint = 255
                var v82 int32 = -2147483648
                v4 = (v4 & -44)
                _ = v81
                _ = v82
            }
        default:
            var v83 uint = 0
            _ = v83
        }
    }
    v84 := GenS11{F0: v21, F1: 32768}
    var v85 int8 = (v4 >> ((0) & 63))
    v86 := GenS9{F0: -2147483648, F1: v5, F2: ("a" + ("a" + v5))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v66
    _ = v84
    _ = v85
    _ = v86
}
