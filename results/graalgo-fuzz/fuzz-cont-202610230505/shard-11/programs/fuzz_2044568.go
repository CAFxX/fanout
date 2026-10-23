package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 int64; F1 uint; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 int32 }
func genFn0(p0 string, p1 uint32) uint {
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            switch (fi7 / ((-372) | 1)) {
            case 32768:
                v8 := []int{(^(fi6))}
                _ = v8
            case 183:
                if false {
                    p0 = p0
                } else {
                    v9 := GenS1{F0: 291}
                    if true {
                        var v10 int = -9223372036854775808
                        v11 := []int{-687, (fi6 & fi7), v10, v10}
                        _ = len(v11)
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(p1, p1, p0)
                        var v12 string = "hello"
                        p0 = ((v12 + (("hello" + (v12 + p0)) + "a")) + "a")
                        _ = v12
                    }
                    switch "a" {
                    case "a":
                        var v13 uint64 = 1
                        _ = v13
                    case "\n\t":
                        switch fi6 {
                        case -495:
                            fmt.Println(fi6, p1, p0)
                        case 4294967295:
                            v14 := GenS1{F0: 0}
                            _ = v14
                        }
                    case "":
                        switch "a" {
                        case "hello":
                            p1 -= (p1 | (p1 >> ((42) & 63)))
                        case "":
                            switch fi7 {
                            case 1:
                                v15 := GenS2{F0: p0}
                                _ = v15
                            case -563:
                                p1 += (p1 >> ((135) & 63))
                            }
                        case "foo bar":
                            var v16 int32 = 1
                            _ = v16
                        }
                    default:
                        var v17 string = p0
                        _ = v17
                    }
                    _ = v9
                }
            case -580:
                p1 += 2147483647
            default:
                if false {
                    p0 = (p0 + p0)
                    if ((fi7 - (fi6 * (^(fi7)))) < (fi7 % ((fi7) | 1))) {
                        fmt.Println(fi7, fi7)
                        switch 922 {
                        case 32767:
                            var v18 int8 = 0
                            _ = v18
                        case 304:
                            fmt.Println(fi6, p1)
                        case -1:
                            var v19 string = p0
                            if true {
                                v20 := []int{9223372036854775807, -9223372036854775808, (-(fi6))}
                                fmt.Println(fi6, fi7)
                                p1 = 615
                                _ = v20
                            } else {
                                var v21 float32 = 1.5
                                _ = v21
                            }
                            _ = v19
                        }
                        v22 := GenS3{F0: -37, F1: 839}
                        p0 = ((p0 + p0) + "hello")
                        _ = v22
                    }
                } else {
                    if (!true) {
                        p0 = p0
                    } else {
                        v23 := []int{}
                        p0 = p0
                        var v24 uint = uint(p1)
                        _ = v23
                        _ = v24
                    }
                    var v25 int = (-(fi6))
                    _ = v25
                }
            }
            _ = fi7
        }
        switch "\n\t" {
        case "a":
            for fi26 := 0; fi26 < 1; fi26++ {
                p0 = ((p0 + ("a" + (p0 + ("foo bar" + (p0 + "hello"))))) + "foo bar")
                v27 := []int{-257}
                var v28 string = (p0 + p0)
                _ = fi26
                _ = v27
                _ = v28
            }
        case "foo bar":
            var v29 int = 32768
            _ = v29
        }
        var v30 int32 = 1
        _ = fi6
        _ = v30
    }
    var v31 uint = uint(p1)
    var v32 int64 = -9223372036854775808
    var v33 int8 = 42
    if (1.5 <= (0.0 / 1.5)) {
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 uint8 = 2
            if (v31 > v31) {
                if false {
                    if (false || ((p1 <= (^(p1))) && false)) {
                        v31 = (v31 | (v31 | (v31 | (v31 << (((v31 << (((^(v31))) & 63))) & 63)))))
                        p0 = "hello"
                    }
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v35 += v35
                        v37 := GenS4{F0: (^(p1)), F1: 42}
                        var v38 uint64 = 233
                        v39 := []int{-403}
                        var v40 int8 = 0
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    var v41 int = (fi34 >> ((v31) & 63))
                    _ = v41
                }
                switch p0 {
                case "foo bar":
                    v42 := GenS4{F0: (^(p1)), F1: 652}
                    _ = v42
                case "a":
                    p1 = p1
                case "\n\t":
                    fmt.Println(fi34)
                }
            }
            v43 := GenS5{F0: 1, F1: v31, F2: true}
            _ = fi34
            _ = v35
            _ = v43
        }
    }
    _ = p0
    _ = p1
    _ = v31
    _ = v32
    _ = v33
    return (v31 % (((v31 << ((v31) & 63))) | 1))
}
func genFn1(p0 uint8, p1 uint8, p2 uint64) int {
    p0 += p1
    var v44 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v44
    return 32767
}

func main() {
    var v1 float32 = (-2.25 - 1e10)
    var v2 string = "hello"
    var v3 float64 = ((float64(v1) - ((0.1 * 1e10) - ((1.5 / (float64(float32(((0.0 / 1.5) + 3.14159))) * float64(v1))) + (0.1 - 1.5)))) * ((1e10 * (1.5 * 0.0)) + ((3.14159 / (0.1 - 1.5)) - float64(v1))))
    var v4 float64 = 1.5
    var v5 int16 = 42
    var v45 int32 = int32(v5)
    if (-9223372036854775808 == -9223372036854775808) {
        var v46 string = "a"
        v47 := GenS4{F0: 4294967295, F1: v45}
        _ = v46
        _ = v47
    } else {
        v48 := GenS6{F0: 0}
        if true {
            v49 := []int{4294967295, genFn1(252, 255, 2), -1}
            v50 := GenS7{F0: -32768}
            _ = v49
            _ = v50
        }
        v2 = (v2 + "hello")
        _ = v48
    }
    v51 := []int{}
    if (float64(1.5) >= v4) {
        for fi52 := 0; fi52 < 11; fi52++ {
            _ = len(v51)
            if ((v45 / ((v45) | 1)) < (v45 + (v45 ^ 1))) {
                for fi53 := 0; fi53 < 12; fi53++ {
                    var v54 uint8 = 255
                    switch (fi53 & 9223372036854775807) {
                    case -1:
                        switch (fi52 / ((fi53) | 1)) {
                        case 409:
                            v55 := GenS8{F0: 9223372036854775807, F1: (v5 + 1000)}
                            _ = v55
                        default:
                            _ = len(v51)
                        }
                    case -427:
                        if len(v51) >= 2 {
                            v51 = v51[1:len(v51)]
                        }
                    default:
                        fmt.Println(v4, v1)
                    }
                    if true {
                        if false {
                            fmt.Println(v3, fi53, v4)
                        }
                        var v56 int16 = v5
                        v51 = append(v51, 458)
                        _ = v56
                    } else {
                        switch v2 {
                        case "foo bar":
                            v2 = ("foo bar" + "hello")
                        case "\n\t":
                            v5 = -32768
                        default:
                            v57 := GenS8{F0: -695, F1: v5}
                            _ = v57
                        }
                        var v58 int16 = v5
                        _ = v58
                    }
                    _ = fi53
                    _ = v54
                }
                if false {
                    fmt.Println(v3, v1)
                } else {
                    v3 = (v3 - ((float64(0.1) - (v4 * v4)) - (v3 * float64((-2.25 / 1.0)))))
                    if true {
                        v51 = append(v51, fi52)
                        for fi59 := 0; fi59 < 12; fi59++ {
                            var v60 uint64 = 2
                            var v61 string = (v2 + "a")
                            var v62 float32 = (1e10 * float32(v4))
                            switch v61 {
                            case "\n\t":
                                var v63 int = (fi59 & -544)
                                _ = v63
                            case "":
                                var v64 float32 = v62
                                _ = v64
                            case "a":
                                if (32767 == 644) {
                                    switch 851 {
                                    case -206:
                                        v5 += (v5 + (v5 << ((666) & 63)))
                                    case 2147483647:
                                        if true {
                                            v62 = 1.5
                                            if (!true) {
                                                v45 -= v45
                                                v2 = ("" + v2)
                                                v61 = (v2 + v2)
                                            }
                                            if len(v51) > 0 {
                                                v51[0] = fi59
                                            }
                                            v1 -= (1.5 - ((v1 / 1.0) / (v62 / (v1 / v1))))
                                        }
                                    case 1:
                                        v51 = append(v51, fi52)
                                    }
                                    v1 = (3.14159 / v1)
                                    if (4294967295 >= 65535) {
                                        fmt.Println(v61, v1, v61)
                                        v65 := GenS4{F0: 65535, F1: v45}
                                        if true {
                                            if (165 > 255) {
                                                v45 = v45
                                                if false {
                                                    v3 = -2.25
                                                    v2 = v61
                                                } else {
                                                    switch (fi59 >> ((444) & 63)) {
                                                    case 4294967295:
                                                        var v66 uint32 = 869
                                                        _ = v66
                                                    case 32768:
                                                        switch 4294967295 {
                                                        case -3:
                                                            v1 = 0.0
                                                        case -727:
                                                            var v67 int32 = (v45 & (v45 * -1))
                                                            _ = v67
                                                        }
                                                    case -117:
                                                        fmt.Println(v4, fi59, v60)
                                                    }
                                                    var v68 uint16 = 42
                                                    _ = v68
                                                }
                                            }
                                            switch (("" + (v2 + v2)) + v61) {
                                            case "hello":
                                                switch 897 {
                                                case 32767:
                                                    if false {
                                                        if true {
                                                            v69 := GenS2{F0: "a"}
                                                            if true {
                                                                var v70 int64 = 42
                                                                if len(v51) >= 2 {
                                                                    v51 = v51[1:len(v51)]
                                                                }
                                                                v3 = v3
                                                                var v71 uint64 = v60
                                                                v5 -= 1000
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            var v72 int8 = 1
                                                            _ = v69
                                                            _ = v72
                                                        } else {
                                                            var v73 uint = 499
                                                            var v74 float32 = v1
                                                            v51 = append(v51, (fi52 + fi59))
                                                            _ = v73
                                                            _ = v74
                                                        }
                                                    } else {
                                                        v5 = v5
                                                    }
                                                case -914:
                                                    v1 = (float32((v3 + (((v4 + 1.5) + 0.1) + (0.1 - (float64(v62) * (v3 + v4)))))) - v62)
                                                case -1:
                                                    v1 = (v62 - v1)
                                                default:
                                                    if true {
                                                        var v75 int8 = 100
                                                        v61 = v2
                                                        var v76 string = v61
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                }
                                            case "":
                                                if len(v51) >= 2 {
                                                    v51 = v51[1:len(v51)]
                                                }
                                            case "a":
                                                switch (fi59 * fi52) {
                                                case 0:
                                                    var v77 uint16 = 32767
                                                    _ = v77
                                                default:
                                                    var v78 int64 = 1
                                                    _ = v78
                                                }
                                            }
                                            var v79 string = v61
                                            fmt.Println(v60, v62)
                                            _ = v79
                                        } else {
                                            var v80 bool = (82 >= v60)
                                            _ = v80
                                        }
                                        fmt.Println(v45, v45)
                                        var v81 int = (fi52 - (fi52 % (((fi59 - fi59)) | 1)))
                                        _ = len(v51)
                                        var v82 string = ""
                                        var v83 bool = (false || (!(v45 >= v45)))
                                        v84 := GenS9{F0: 0}
                                        _ = v65
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    } else {
                                        if false {
                                            var v85 float32 = v62
                                            _ = v85
                                        }
                                        v51 = append(v51, fi59)
                                    }
                                    v86 := GenS3{F0: 0, F1: 418}
                                    v45 = (v45 << ((25) & 63))
                                    _ = v86
                                } else {
                                    v3 = -2.25
                                }
                            default:
                                var v87 int64 = -1
                                _ = v87
                            }
                            var v88 uint32 = 2147483647
                            v51 = append(v51, (fi59 & int(v45)))
                            v2 = v2
                            _ = fi59
                            _ = v60
                            _ = v61
                            _ = v62
                            _ = v88
                        }
                        var v89 float64 = v3
                        _ = v89
                    }
                    var v90 uint8 = 127
                    _ = v90
                }
            }
            _ = fi52
        }
    }
    var v91 float64 = v3
    if false {
        for fi92 := 0; fi92 < 12; fi92++ {
            fmt.Println(v3, v1, v1)
            _ = fi92
        }
    }
    v2 = (("foo bar" + (v2 + "")) + "\n\t")
    if len(v51) > 0 {
        v51[0] = genFn1(238, 254, 2)
    }
    if (0 > 274) {
        _ = len(v51)
    } else {
        fmt.Println(v5)
        fmt.Println(v3, v5)
        var v93 int8 = -1
        _ = v93
    }
    switch 42 {
    case 1:
        fmt.Println(v1, v5, v1)
    case 4294967295:
        v51 = append(v51, -942)
    }
    v3 += 0.1
    _ = len(v51)
    for fi94 := 0; fi94 < 13; fi94++ {
        var v95 uint = 1
        _ = fi94
        _ = v95
    }
    v96 := GenS6{F0: 18446744073709551615}
    fmt.Println(v5, v4)
    var v97 int8 = -25
    if false {
        var v98 int16 = (v5 % ((v5) | 1))
        if (409 <= 65535) {
            var v99 float64 = ((0.0 * v4) + 3.14159)
            _ = len(v51)
            _ = v99
        } else {
            var v100 bool = true
            _ = v100
        }
        v45 += 42
        _ = v98
    } else {
        var v101 bool = (true && (127 != 127))
        var v102 string = (v2 + "hello")
        _ = v101
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v91:", v91)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v51
    _ = v91
    _ = v96
    _ = v97
}
