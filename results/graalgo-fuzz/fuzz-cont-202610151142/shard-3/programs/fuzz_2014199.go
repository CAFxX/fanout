package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 string; F1 uint32; F2 uint64 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 string; F2 int64 }
type GenS10 struct { F0 uint32; F1 bool; F2 int32 }
type GenS11 struct { F0 string; F1 int32; F2 uint }
func genFn0(p0 uint, p1 uint16, p2 int32) int32 {
    v4 := GenS1{F0: "\n\t", F1: (-(p2))}
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 int32 = 694
    if false {
        var v5 int = 42
        _ = v5
    }
    if v2 {
        v2 = (!true)
        v6 := []int{813, 9223372036854775807, 1}
        fmt.Println(v3, v3)
        _ = v6
    }
    if false {
        if v2 {
            var v7 uint64 = 1
            _ = v7
        } else {
            v8 := GenS1{F0: (("a" + "hello") + "foo bar"), F1: (v3 / (((v3 ^ v3)) | 1))}
            v3 = 42
            v9 := []int{0, 127, 32768}
            _ = v8
            _ = v9
        }
    } else {
        var v10 int32 = v3
        _ = v10
    }
    v2 = (2 <= 255)
    var v11 bool = v2
    if ((42 >= 485) && true) {
        v12 := GenS1{F0: ("a" + "a"), F1: v3}
        v13 := GenS2{F0: 255, F1: -128}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v3 = 45
        for fi15 := 0; fi15 < 2; fi15++ {
            var v16 float32 = 0.0
            v17 := []int{(fi15 | (fi15 ^ fi15)), fi14, (fi14 & fi15)}
            v3 += 0
            v2 = (true || v11)
            var v18 int64 = 1
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        switch "a" {
        case "foo bar":
            for fi19 := 0; fi19 < 8; fi19++ {
                if false {
                    var v20 string = ""
                    if v2 {
                        v3 += 0
                        var v21 uint8 = 42
                        var v22 int = (fi14 >> ((413) & 63))
                        v23 := []int{v22, v22, 197}
                        if (v11 && v2) {
                            v23 = append(v23, fi19)
                            v24 := GenS2{F0: uint16(v21), F1: -79}
                            v21 -= (v21 | (v21 - 68))
                            _ = v24
                        } else {
                            if (!(v2 && v11)) {
                                if (1 <= 2147483647) {
                                    var v25 uint = 255
                                    _ = v25
                                } else {
                                    v26 := GenS3{F0: (fi19 ^ v22), F1: v3}
                                    _ = v26
                                }
                                fmt.Println(v2, v21, v22)
                            }
                        }
                        if ((2 <= 790) || v2) {
                            var v27 uint8 = 127
                            _ = v27
                        } else {
                            v28 := GenS3{F0: 537, F1: 0}
                            if (!v2) {
                                switch v22 {
                                case 151:
                                    v22 -= 32767
                                }
                                var v29 string = v20
                                v30 := GenS4{F0: 0, F1: (47 >= -1)}
                                _ = v29
                                _ = v30
                            } else {
                                var v31 float64 = 3.14159
                                _ = v31
                            }
                            _ = v28
                        }
                        fmt.Println(v2, fi19)
                        if (9223372036854775807 <= 0) {
                            switch 680 {
                            case -870:
                                fmt.Println(v11, v22, v3)
                            case 32767:
                                if (false || v2) {
                                    switch (fi19 / ((1) | 1)) {
                                    case -1:
                                        var v32 bool = v2
                                        _ = v32
                                    case 32767:
                                        if (42 >= 42) {
                                            switch (fi14 * (fi14 << ((65535) & 63))) {
                                            case 2147483647:
                                                var v33 bool = (false || true)
                                                var v34 uint16 = 450
                                                v35 := GenS5{F0: "a", F1: 300, F2: 0}
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            default:
                                                if ((1.5 - 1.5) > 3.14159) {
                                                    v3 += genFn0(4294967295, 2, v3)
                                                }
                                            }
                                            v3 = (^(v3))
                                            v20 = "foo bar"
                                            v20 = v20
                                            switch v20 {
                                            case "hello":
                                                if true {
                                                    var v36 uint32 = 1
                                                    v23 = append(v23, 32767)
                                                    switch v22 {
                                                    case 127:
                                                        if v2 {
                                                            v2 = v11
                                                            v20 = v20
                                                        }
                                                    }
                                                    _ = v36
                                                } else {
                                                    v1[32768] = fi14
                                                    switch v20 {
                                                    case "\n\t":
                                                        if v11 {
                                                            var v37 string = v20
                                                            fmt.Println(v2, v21, v37)
                                                            _ = v37
                                                        } else {
                                                            if (v11 || v2) {
                                                                switch (v20 + v20) {
                                                                case "\n\t":
                                                                    fmt.Println(v11, v11)
                                                                case "":
                                                                    if (-32768 <= -195) {
                                                                        var v38 uint8 = v21
                                                                        switch v20 {
                                                                        case "foo bar":
                                                                            fmt.Println(v38, fi14)
                                                                            v21 = v21
                                                                        case "a":
                                                                            v1[(v22 - -242)] = (fi19 | (v22 - 1))
                                                                            if v2 {
                                                                                if true {
                                                                                    var v39 bool = v11
                                                                                    _ = v39
                                                                                }
                                                                            }
                                                                            v38 = v21
                                                                            var v40 int = -1
                                                                            var v41 int16 = 1000
                                                                            _ = v40
                                                                            _ = v41
                                                                        case "":
                                                                            var v42 uint = 1
                                                                            _ = v42
                                                                        }
                                                                        _ = v38
                                                                    } else {
                                                                        v3 = -175
                                                                        v20 = (v20 + v20)
                                                                        v2 = true
                                                                    }
                                                                case "a":
                                                                    if len(v23) >= 2 {
                                                                        v23 = v23[1:len(v23)]
                                                                    }
                                                                default:
                                                                    switch v20 {
                                                                    case "foo bar":
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if len(v23) > 0 {
                                                                                v23[0] = fi14
                                                                            }
                                                                        default:
                                                                            if ((^(v21)) >= 42) {
                                                                                _ = len(v23)
                                                                                v11 = v11
                                                                            }
                                                                            switch v20 {
                                                                            case "a":
                                                                                fmt.Println(fi19, v2, v21)
                                                                            case "foo bar":
                                                                                var v43 float64 = ((float64((3.14159 - 0.0)) / ((3.14159 / -2.25) + 0.1)) * (3.14159 / ((3.14159 / (float64(1.5) / 0.1)) - 0.0)))
                                                                                _ = v43
                                                                            }
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v22)
                                                                        var v44 uint64 = 0
                                                                        _ = v44
                                                                    case "a":
                                                                        v45 := v1[(fi14 >> ((255) & 63))]
                                                                        _ = v45
                                                                    }
                                                                }
                                                            } else {
                                                                var v46 int32 = v3
                                                                v47 := GenS2{F0: 780, F1: 110}
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                            var v48 bool = false
                                                            _ = v48
                                                        }
                                                    case "a":
                                                        v20 = v20
                                                    case "foo bar":
                                                        v22 = 875
                                                    }
                                                    switch v22 {
                                                    case 9223372036854775807:
                                                        v21 -= v21
                                                    case -9223372036854775808:
                                                        v2 = v2
                                                    case 32768:
                                                        _ = len(v23)
                                                    }
                                                    v2 = ((v22 * 73) != (fi19 % (((fi19 * (v22 | 42))) | 1)))
                                                    if v11 {
                                                        if false {
                                                            if (v11 || (4294967295 < 2147483647)) {
                                                                v2 = v2
                                                            } else {
                                                                fmt.Println(v3, v20, v21)
                                                            }
                                                            var v49 int64 = -1
                                                            v1[fi14] = 127
                                                            v50 := GenS6{F0: v11, F1: 1, F2: 79}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v51 := v1[fi14]
                                                            fmt.Println(v3)
                                                            var v52 uint = 0
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                var v53 uint16 = 688
                                                _ = v53
                                            }
                                        } else {
                                            v54 := v1[2147483647]
                                            _ = v54
                                        }
                                    case 128:
                                        var v55 float64 = float64((1.5 * float32(1.5)))
                                        _ = v55
                                    default:
                                        var v56 string = "a"
                                        _ = v56
                                    }
                                }
                            case 0:
                                switch v22 {
                                case -2147483648:
                                    if v2 {
                                        if (v2 && (0 == -860)) {
                                            var v57 string = (v20 + (((v20 + ("" + "foo bar")) + v20) + (v20 + (v20 + v20))))
                                            v11 = (-2.25 < float32((3.14159 - 1.5)))
                                            fmt.Println(v2, v22, v3)
                                            v58 := GenS2{F0: 152, F1: 0}
                                            _ = v57
                                            _ = v58
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch 2147483647 {
                                            case 188:
                                                fmt.Println(v20, fi14, v11)
                                            default:
                                                var v59 int16 = 42
                                                _ = v59
                                            }
                                        case "\n\t":
                                            v21 -= (^(v21))
                                            var v60 string = v20
                                            _ = v60
                                        }
                                        _ = len(v23)
                                    }
                                case 128:
                                    v11 = (int8(v21) != -1)
                                case -128:
                                    _ = len(v23)
                                    v61 := GenS7{F0: 42}
                                    _ = v61
                                default:
                                    _ = len(v23)
                                }
                            default:
                                v62 := GenS3{F0: fi14, F1: (v3 ^ genFn0(1, 809, (v3 | (v3 + 0))))}
                                _ = v62
                            }
                            if true {
                                if v2 {
                                    v63 := GenS7{F0: 0}
                                    _ = v63
                                }
                            }
                            var v64 int64 = 9223372036854775807
                            _ = v64
                        } else {
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            switch (v22 ^ v22) {
                            case -1:
                                if ((1.5 - 1e10) >= 0.0) {
                                    var v65 int = fi19
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    var v66 int8 = 0
                                    _ = v65
                                    _ = v66
                                } else {
                                    switch (v20 + v20) {
                                    case "a":
                                        v20 = "a"
                                    case "foo bar":
                                        delete(v1, -68)
                                        if v11 {
                                            if ((v3 != v3) && false) {
                                                v67 := v1[fi19]
                                                _ = v67
                                            }
                                        } else {
                                            v3 = v3
                                        }
                                        if v2 {
                                            var v68 uint16 = 21
                                            v21 = (v21 & v21)
                                            _ = v68
                                        } else {
                                            var v69 int64 = -9223372036854775808
                                            _ = v69
                                        }
                                    case "hello":
                                        var v70 int32 = int32(v21)
                                        _ = v70
                                    }
                                }
                            case 32768:
                                v3 -= v3
                            case 734:
                                v20 = ""
                            }
                        }
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        switch fi19 {
                        case 128:
                            var v71 int = fi19
                            _ = v71
                        case -1:
                            if (v2 && false) {
                                fmt.Println(v2, v3, v20)
                                if (v2 && false) {
                                    var v72 int16 = 1000
                                    v73 := GenS8{F0: v3}
                                    _ = v72
                                    _ = v73
                                }
                            }
                        default:
                            v11 = ("\n\t" <= v20)
                        }
                    }
                    var v74 bool = v11
                    _ = v20
                    _ = v74
                }
                switch fi14 {
                case 32768:
                    v75 := []int{}
                    if v2 {
                        v76 := GenS4{F0: 520, F1: (v2 || (v11 || false))}
                        _ = v76
                    }
                    _ = v75
                case 128:
                    switch "a" {
                    case "a":
                        if (32767 < 256) {
                            v2 = v2
                            v11 = v2
                            fmt.Println(v11, v3, v2)
                        }
                    default:
                        v77 := GenS9{F0: ("a" + "foo bar"), F1: ((("\n\t" + (("\n\t" + "\n\t") + (("" + "a") + ("" + ("foo bar" + (((("\n\t" + ("a" + ("a" + "hello"))) + ("" + "a")) + "hello") + "\n\t")))))) + "") + "foo bar"), F2: 0}
                        _ = v77
                    }
                }
                v78 := GenS7{F0: 4294967295}
                _ = fi19
                _ = v78
            }
        case "\n\t":
            v79 := GenS3{F0: (fi14 / ((492) | 1)), F1: 42}
            _ = v79
        default:
            v1[(^(fi14))] = 128
        }
        _ = fi14
    }
    v11 = true
    switch "a" {
    case "\n\t":
        var v80 bool = (0.0 <= 1.5)
        _ = v80
    case "hello":
        switch "\n\t" {
        case "a":
            v81 := []int{2147483647, -9223372036854775808, 32767, -2147483648}
            _ = v81
        case "hello":
            v82 := GenS10{F0: 42, F1: true, F2: (v3 & (v3 << ((65535) & 63)))}
            _ = v82
        default:
            if (-530 == 42) {
                v83 := GenS3{F0: -310, F1: 1}
                _ = v83
            }
        }
    }
    var v84 bool = v11
    var v85 float64 = (float64(float32(float64(((3.14159 * 1.5) - 0.0)))) - 3.14159)
    v86 := GenS8{F0: -827}
    fmt.Println(v3, v2)
    if (v2 || (2 > 0)) {
        var v87 uint32 = 2147483648
        if v84 {
            if (v2 || true) {
                fmt.Println(v2, v85)
                var v88 int = 128
                v89 := GenS11{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: (v3 % ((genFn0(0, 65535, v3)) | 1)), F2: 2}
                _ = v88
                _ = v89
            } else {
                switch 2147483647 {
                case 244:
                    v87 += (^(v87))
                case 32767:
                    fmt.Println(v87, v85)
                case 32768:
                    switch "\n\t" {
                    case "a":
                        v90 := GenS3{F0: -9223372036854775808, F1: genFn0(255, 258, v3)}
                        _ = v90
                    default:
                        for fi91 := 0; fi91 < 16; fi91++ {
                            var v92 int16 = 32767
                            for fi93 := 0; fi93 < 5; fi93++ {
                                fmt.Println(v2, fi93, v87)
                                _ = fi93
                            }
                            _ = fi91
                            _ = v92
                        }
                    }
                }
            }
            var v94 float32 = float32(0.0)
            _ = v94
        } else {
            var v95 int16 = 1
            _ = v95
        }
        _ = v87
    } else {
        _ = len(v1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v84
    _ = v85
    _ = v86
}
