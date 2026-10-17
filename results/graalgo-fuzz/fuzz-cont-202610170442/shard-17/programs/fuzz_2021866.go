package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 string; F2 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 string; F2 bool }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 bool; F1 bool; F2 uint8 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint8; F1 uint16; F2 bool }
type GenS11 struct { F0 uint32; F1 int16 }
func genFn0() bool {
    var v7 uint8 = 85
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 477}
        _ = fi8
        _ = v9
    }
    var v10 int8 = 95
    _ = v10
    _ = v7
    return (0 < 1)
}
func genFn1(p0 int16, p1 uint8) uint {
    v11 := []int{}
    if true {
        var v12 uint32 = 150
        _ = v12
    } else {
        var v13 uint64 = 0
        fmt.Println(p0, p1)
        _ = v13
    }
    switch ("foo bar" + "a") {
    case "":
        for fi14 := 0; fi14 < 9; fi14++ {
            if (1 != 0) {
                for fi15 := 0; fi15 < 6; fi15++ {
                    _ = len(v11)
                    v16 := GenS2{F0: ("a" + "hello"), F1: (("\n\t" + "a") + "\n\t"), F2: 1}
                    p1 = (p1 - 77)
                    _ = fi15
                    _ = v16
                }
                var v17 float32 = 1e10
                v18 := GenS2{F0: "", F1: ("hello" + ""), F2: 0}
                _ = v17
                _ = v18
            } else {
                var v19 int32 = -2147483648
                var v20 int8 = 100
                _ = v19
                _ = v20
            }
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 602
}

func main() {
    var v1 int64 = -958
    v2 := make(map[int]uint32)
    var v3 uint64 = 42
    v4 := []string{}
    v5 := make(map[int]string)
    var v6 string = ""
    var v21 bool = false
    if (0 < 1) {
        fmt.Println(v6, v21)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v22 int64 = v1
    if (-1 >= 2147483647) {
        if (!(!v21)) {
            switch 9223372036854775807 {
            case 4294967295:
                var v23 uint64 = v3
                _ = v23
            case 1:
                v4 = append(v4, (("" + v6) + v6))
            default:
                v24 := GenS3{F0: v21}
                _ = v24
            }
        } else {
            fmt.Println(v1)
            v25 := GenS3{F0: (!v21)}
            var v26 uint16 = 1
            _ = v25
            _ = v26
        }
        v27 := GenS1{F0: 69}
        _ = v27
    } else {
        v28 := GenS4{F0: true, F1: v6, F2: ((v3 & v3) < v3)}
        v29 := GenS5{F0: 42, F1: 1}
        _ = v28
        _ = v29
    }
    switch 42 {
    case 128:
        v30 := GenS6{F0: (!v21), F1: v21, F2: 255}
        _ = v30
    case -959:
        var v31 uint16 = 1
        _ = v31
    default:
        switch -9223372036854775808 {
        case 42:
            switch "a" {
            case "a":
                var v32 bool = ((116 >= 2) || (!(!v21)))
                _ = v32
            case "foo bar":
                var v33 bool = v21
                _ = v33
            default:
                if v21 {
                    var v34 bool = v21
                    _ = v34
                }
            }
        case -179:
            switch 873 {
            case 128:
                v3 -= v3
            }
        case 9223372036854775807:
            if false {
                for fi35 := 0; fi35 < 9; fi35++ {
                    for fi36 := 0; fi36 < 1; fi36++ {
                        fmt.Println(v21, fi35)
                        _ = fi36
                    }
                    for fi37 := 0; fi37 < 10; fi37++ {
                        if true {
                            var v38 uint8 = 42
                            v3 = (v3 + 2)
                            _ = v38
                        }
                        if true {
                            var v39 uint8 = 255
                            _ = v39
                        }
                        v40 := GenS3{F0: true}
                        _ = fi37
                        _ = v40
                    }
                    v1 = (v22 - v1)
                    _ = fi35
                }
                v1 = (v22 % (((-(v22))) | 1))
                for fi41 := 0; fi41 < 18; fi41++ {
                    var v42 int16 = 1
                    v43 := GenS4{F0: (v3 != v3), F1: "\n\t", F2: (v6 <= (v6 + v6))}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 19; fi44++ {
                    var v45 uint32 = 614
                    v21 = (((3.14159 + 0.0) >= ((float64(3.14159) * -2.25) * -2.25)) && v21)
                    _ = fi44
                    _ = v45
                }
            }
        }
    }
    var v46 uint16 = 42
    switch 4294967295 {
    case 4294967295:
        var v47 float64 = ((3.14159 / 1.0) - -2.25)
        _ = v47
    case 1:
        for fi48 := 0; fi48 < 13; fi48++ {
            v3 += v3
            if (-115 <= (v1 << ((genFn1(683, 1)) & 63))) {
                var v49 int8 = 127
                var v50 uint16 = 937
                switch (v6 + ("hello" + "\n\t")) {
                case "\n\t":
                    for fi51 := 0; fi51 < 3; fi51++ {
                        if len(v4) > 0 {
                            v4[0] = ("a" + (v6 + v6))
                        }
                        if (!(v3 <= v3)) {
                            switch (-(fi48)) {
                            case -268:
                                var v52 int16 = 1000
                                _ = v52
                            case 128:
                                var v53 float32 = 1.5
                                _ = v53
                            }
                        } else {
                            var v54 uint32 = uint32(v3)
                            _ = v54
                        }
                        if (((fi48 >> ((42) & 63)) <= -120) && (0 > 42)) {
                            var v55 int64 = 1
                            _ = v55
                        } else {
                            if true {
                                if (-368 != -176) {
                                    v56 := GenS5{F0: 739, F1: 132}
                                    switch (("\n\t" + v6) + v6) {
                                    case "a":
                                        if v21 {
                                            if len(v4) > 0 {
                                                v4[0] = ("a" + (v6 + v6))
                                            }
                                            switch (v6 + v6) {
                                            case "hello":
                                                v57 := GenS7{F0: "a", F1: v3}
                                                _ = v57
                                            case "":
                                                var v58 float64 = ((float64(0.0) + (1e10 + 0.1)) * 1.5)
                                                _ = v58
                                            case "\n\t":
                                                v22 = v22
                                            default:
                                                v59 := GenS5{F0: 255, F1: 2147483647}
                                                _ = v59
                                            }
                                            _ = len(v2)
                                            if (-1 != 32767) {
                                                if (969 != 42) {
                                                    v22 = -9223372036854775808
                                                    fmt.Println(v3, v49, v1)
                                                    v60 := GenS4{F0: true, F1: "hello", F2: true}
                                                    var v61 float32 = 3.14159
                                                    v62 := GenS1{F0: 42}
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    switch v6 {
                                                    case "a":
                                                        switch 1 {
                                                        case 128:
                                                            v63 := GenS8{F0: (v22 / (((-(v1))) | 1)), F1: ((v50 - 65535) <= (^(v50)))}
                                                            _ = v63
                                                        case 42:
                                                            if v21 {
                                                                v6 = ""
                                                                if v21 {
                                                                    fmt.Println(v21)
                                                                    var v64 int16 = -1
                                                                    _ = v64
                                                                }
                                                            }
                                                        default:
                                                            v1 = v1
                                                        }
                                                    case "\n\t":
                                                        v5[(fi48 * fi51)] = "a"
                                                        v1 += v22
                                                    }
                                                    v65 := GenS9{F0: 126, F1: 0}
                                                    _ = v65
                                                }
                                                var v66 int16 = 343
                                                switch (fi51 ^ (fi51 % (((-(fi48))) | 1))) {
                                                case -516:
                                                    v67 := v2[(fi48 / ((9223372036854775807) | 1))]
                                                    fmt.Println(v50)
                                                    if ((-2.25 / -2.25) > float32((0.1 - (0.0 + 1e10)))) {
                                                        switch fi51 {
                                                        case 32767:
                                                            var v68 uint32 = 2147483648
                                                            v68 = (v68 | v67)
                                                            fmt.Println(fi51, v49)
                                                            _ = v68
                                                        case -9223372036854775808:
                                                            switch fi48 {
                                                            case 9223372036854775807:
                                                                v6 = ((v6 + v6) + v6)
                                                            case 32768:
                                                                var v69 uint = 146
                                                                _ = v69
                                                            }
                                                        case 694:
                                                            var v70 int = 4294967295
                                                            _ = v70
                                                        }
                                                        v6 = "\n\t"
                                                        delete(v2, -955)
                                                        var v71 string = v6
                                                        _ = v71
                                                    } else {
                                                        fmt.Println(v21)
                                                    }
                                                    _ = len(v2)
                                                    _ = v67
                                                case 9223372036854775807:
                                                    fmt.Println(v1)
                                                case 4294967295:
                                                    v6 = v6
                                                }
                                                v6 = "\n\t"
                                                var v72 float64 = 3.14159
                                                if (v21 && v21) {
                                                    v4 = append(v4, (("foo bar" + v6) + "\n\t"))
                                                }
                                                _ = v66
                                                _ = v72
                                            } else {
                                                var v73 int16 = -6
                                                var v74 float64 = (-2.25 / ((3.14159 * -2.25) - 1e10))
                                                v3 = v3
                                                fmt.Println(v73, v50, v74)
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                    case "foo bar":
                                        var v75 int = fi48
                                        _ = v75
                                    case "hello":
                                        var v76 int16 = 1
                                        _ = v76
                                    default:
                                        v46 -= v50
                                    }
                                    v1 += (-(v22))
                                    _ = v56
                                }
                            } else {
                                var v77 string = (((v6 + "a") + "a") + v6)
                                _ = v77
                            }
                            var v78 float32 = (((0.0 + 1.5) / 0.1) + -2.25)
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            _ = v78
                        }
                        _ = fi51
                    }
                case "foo bar":
                    var v79 int16 = int16(v50)
                    _ = v79
                case "":
                    var v80 uint = 255
                    _ = v80
                }
                _ = v49
                _ = v50
            }
            fmt.Println(fi48, v3, v22)
            _ = fi48
        }
    }
    if ((2147483647 != -396) && (-125 >= -1)) {
        switch (((v6 + v6) + ("foo bar" + "foo bar")) + (v6 + "foo bar")) {
        case "\n\t":
            var v81 float32 = float32(1.5)
            _ = v81
        case "a":
            v22 += (v22 << ((genFn1(-382, 128)) & 63))
        case "foo bar":
            if (true && (1 != 1000)) {
                v3 -= v3
                var v82 uint8 = uint8(v3)
                _ = v82
            }
        default:
            if (88 > 887) {
                fmt.Println(v6, v21, v46)
            }
        }
    }
    if v21 {
        for fi83 := 0; fi83 < 8; fi83++ {
            for fi84 := 0; fi84 < 18; fi84++ {
                fmt.Println(fi83, v1)
                if (0 <= 811) {
                    _ = len(v5)
                } else {
                    v85 := GenS7{F0: ("" + v6), F1: (^(v3))}
                    v86 := GenS10{F0: 2, F1: v46, F2: v21}
                    var v87 int = (fi83 + 128)
                    _ = v85
                    _ = v86
                    _ = v87
                }
                _ = fi84
            }
            _ = len(v4)
            _ = fi83
        }
    }
    v88 := GenS11{F0: 399, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v88
}
