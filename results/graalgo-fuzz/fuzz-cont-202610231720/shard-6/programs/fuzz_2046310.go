package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint64; F2 int8 }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS11 struct { F0 uint }
func genFn0(p0 int, p1 bool) uint32 {
    v7 := GenS1{F0: 36}
    p1 = (!(-105 >= 127))
    v8 := GenS2{F0: -1, F1: ("hello" + "\n\t"), F2: 0}
    if (!(127 == 1)) {
        switch (p0 | p0) {
        case 2147483647:
            var v9 bool = p1
            _ = v9
        case 42:
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 uint16 = 255
                v12 := GenS2{F0: -914, F1: ("" + "foo bar"), F2: -1}
                _ = fi10
                _ = v11
                _ = v12
            }
        case -9223372036854775808:
            v13 := []int{}
            _ = v13
        }
    } else {
        var v14 float32 = 3.14159
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int, p1 int32) uint64 {
    p0 = p0
    v15 := []int{}
    p0 = (p0 << ((1) & 63))
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = v15
    return 2
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]string)
    v3 := []string{("hello" + "a"), ((("foo bar" + "") + "a") + "foo bar"), (("" + "\n\t") + "a"), "hello"}
    var v4 int32 = 707
    v5 := make(map[uint32]uint32)
    var v6 string = "\n\t"
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS2{F0: 42, F1: v6, F2: 0}
        switch fi16 {
        case 0:
            v1 = 1e10
        }
        _ = fi16
        _ = v17
    }
    switch (v6 + "") {
    case "foo bar":
        for fi18 := 0; fi18 < 19; fi18++ {
            if false {
                v4 += v4
                if false {
                    if (!(1 <= 1)) {
                        v4 -= v4
                    } else {
                        if (genFn1(9223372036854775807, (v4 * v4)) >= 2) {
                            v19 := GenS3{F0: 2147483647}
                            switch (fi18 >> ((937) & 63)) {
                            case -601:
                                fmt.Println(fi18)
                            case 616:
                                v20 := GenS3{F0: genFn0(fi18, false)}
                                _ = v20
                            case -1:
                                for fi21 := 0; fi21 < 10; fi21++ {
                                    v3 = append(v3, v6)
                                    if false {
                                        fmt.Println(v4, fi21)
                                        v22 := v2[(fi18 << ((396) & 63))]
                                        if (1 <= 4294967295) {
                                            fmt.Println(fi21, fi21, v4)
                                            switch v22 {
                                            case "":
                                                v5[2] = 0
                                            case "foo bar":
                                                fmt.Println(fi21)
                                            case "hello":
                                                v1 += v1
                                            }
                                            if ((-(v4)) == (v4 & v4)) {
                                                var v23 int8 = 42
                                                if (false && (v23 <= (v23 * (v23 >> ((371) & 63))))) {
                                                    v5[genFn0(32767, true)] = genFn0(fi21, (false && (1e10 != (v1 + float32((float64(-2.25) - 0.1))))))
                                                    v24 := GenS2{F0: 0, F1: "\n\t", F2: 833}
                                                    v25 := GenS3{F0: 873}
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    var v26 int16 = 1
                                                    _ = v26
                                                }
                                                v22 = ""
                                                _ = v23
                                            } else {
                                                if true {
                                                    v22 = v22
                                                    if ((144 != 2) || true) {
                                                        switch -2147483648 {
                                                        case -1:
                                                            if (!(!true)) {
                                                                var v27 uint64 = genFn1(-972, (v4 + -2147483648))
                                                                var v28 int8 = 20
                                                                var v29 uint8 = 2
                                                                _ = v27
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        case -9223372036854775808:
                                                            var v30 bool = false
                                                            _ = v30
                                                        default:
                                                            if (true || false) {
                                                                if true {
                                                                    delete(v2, (fi18 & 42))
                                                                } else {
                                                                    fmt.Println(v4)
                                                                    var v31 float32 = v1
                                                                    if (4294967295 == 633) {
                                                                        var v32 int64 = -413
                                                                        fmt.Println(fi21, v4)
                                                                        fmt.Println(v31)
                                                                        _ = v32
                                                                    } else {
                                                                        v6 = v6
                                                                    }
                                                                    _ = v31
                                                                }
                                                                if (9223372036854775807 <= 0) {
                                                                    v33 := GenS3{F0: 210}
                                                                    v34 := GenS2{F0: 0, F1: v22, F2: -1}
                                                                    _ = v33
                                                                    _ = v34
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (v6 + (((("a" + "a") + v6) + (v22 + v6)) + ((v6 + v6) + "\n\t")))
                                                                    }
                                                                    v35 := GenS1{F0: 89}
                                                                    _ = len(v5)
                                                                    _ = v35
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v36 := GenS2{F0: 32767, F1: (v6 + v6), F2: 0}
                                                    _ = v36
                                                }
                                                _ = len(v5)
                                                v37 := GenS4{F0: (^(v4))}
                                                var v38 int8 = 0
                                                _ = v37
                                                _ = v38
                                            }
                                        }
                                        _ = v22
                                    }
                                    _ = len(v5)
                                    _ = fi21
                                }
                            default:
                                if true {
                                    v39 := GenS5{F0: fi18, F1: 460}
                                    _ = v39
                                }
                            }
                            v6 = v6
                            _ = v19
                        }
                    }
                    var v40 float32 = 3.14159
                    switch v6 {
                    case "":
                        v41 := GenS1{F0: 1}
                        _ = v41
                    case "\n\t":
                        for fi42 := 0; fi42 < 16; fi42++ {
                            var v43 uint8 = 1
                            v44 := GenS3{F0: 696}
                            v43 = v43
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                    case "a":
                        fmt.Println(v40)
                    default:
                        switch (fi18 >> ((0) & 63)) {
                        case 128:
                            fmt.Println(v1, fi18)
                        default:
                            v1 -= float32(1e10)
                        }
                    }
                    _ = v40
                }
                v45 := GenS6{F0: 0, F1: "a"}
                _ = v45
            } else {
                switch (v6 + v6) {
                case "hello":
                    var v46 float64 = 3.14159
                    _ = v46
                case "foo bar":
                    var v47 uint = 18446744073709551615
                    _ = v47
                case "a":
                    var v48 float64 = (((3.14159 + float64(v1)) - 1e10) * 1.5)
                    _ = v48
                }
                var v49 int64 = 42
                _ = v49
            }
            var v50 int16 = -745
            _ = fi18
            _ = v50
        }
    case "a":
        v1 -= ((((v1 / (v1 * v1)) - v1) / -2.25) + ((3.14159 * ((v1 - (v1 - (float32(0.1) / float32(1.5)))) + float32(1.5))) + v1))
    case "":
        if (v4 != 1) {
            fmt.Println(v1, v6, v4)
            switch -661 {
            case -746:
                v6 = v6
            default:
                var v51 uint32 = 2147483647
                _ = v51
            }
            var v52 string = "\n\t"
            _ = v52
        }
    }
    fmt.Println(v6, v6)
    if (!true) {
        if ((1e10 * 0.0) != -2.25) {
            var v53 string = "a"
            v54 := GenS7{F0: ((false && (v1 == ((v1 / ((1.5 - v1) / v1)) - -2.25))) || (127 <= 100))}
            _ = v53
            _ = v54
        }
    } else {
        var v55 uint8 = 88
        fmt.Println(v4, v6)
        v6 = v6
        _ = v55
    }
    if false {
        var v56 string = "\n\t"
        _ = v56
    } else {
        if false {
            switch -2147483648 {
            case 9223372036854775807:
                v6 = (v6 + (((v6 + "hello") + (v6 + ("\n\t" + ((v6 + v6) + "hello")))) + v6))
            case 127:
                var v57 uint64 = genFn1(-170, (v4 % (((v4 | (v4 >> ((2) & 63)))) | 1)))
                _ = v57
            }
            var v58 uint32 = 2
            v59 := v5[(v58 >> ((0) & 63))]
            _ = v58
            _ = v59
        } else {
            if (("a" == "\n\t") || true) {
                var v60 int8 = 127
                fmt.Println(v6, v1, v60)
                var v61 int16 = 0
                _ = v60
                _ = v61
            }
            if (42 < -128) {
                v5[65535] = 2147483647
                v62 := GenS8{F0: -95, F1: 1, F2: 100}
                if true {
                    v4 += (-(v4))
                    if ((true && (!(9223372036854775807 <= -33))) || false) {
                        fmt.Println(v6)
                        v1 = v1
                        switch 530 {
                        case 127:
                            for fi63 := 0; fi63 < 1; fi63++ {
                                var v64 float32 = 0.1
                                if false {
                                    switch fi63 {
                                    case 0:
                                        var v65 float32 = v1
                                        _ = v65
                                    case 4294967295:
                                        switch 4294967295 {
                                        case 32768:
                                            if len(v3) > 0 {
                                                v3[0] = v6
                                            }
                                        default:
                                            v66 := GenS7{F0: (0 >= 127)}
                                            _ = v66
                                        }
                                    case 127:
                                        var v67 uint64 = 2
                                        _ = v67
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = (((((v6 + (((v6 + v6) + v6) + ("\n\t" + "a"))) + v6) + "a") + (v6 + v6)) + "hello")
                                        }
                                    }
                                    var v68 int8 = 1
                                    _ = v68
                                }
                                var v69 string = "foo bar"
                                _ = fi63
                                _ = v64
                                _ = v69
                            }
                        case 128:
                            v70 := GenS7{F0: true}
                            _ = v70
                        case -9223372036854775808:
                            v71 := GenS1{F0: 128}
                            _ = v71
                        default:
                            if true {
                                _ = len(v3)
                                var v72 uint32 = 284
                                if false {
                                    for fi73 := 0; fi73 < 12; fi73++ {
                                        switch 397 {
                                        case -429:
                                            var v74 float64 = (((1.5 * (1.5 * float64((((0.0 + ((((0.1 * float32((0.0 + 1e10))) + (0.1 - 1e10)) * 0.0) - 0.1)) * (1e10 / float32(float64(-2.25)))) / float32((((((float64(((float32(0.0) / (v1 - float32(((1.5 + 3.14159) / -2.25)))) + v1)) - 0.0) + (3.14159 - (-2.25 - 0.1))) - 1.5) * 3.14159) - 1.5)))))) * 3.14159) + 3.14159)
                                            _ = v74
                                        }
                                        fmt.Println(v1, v6)
                                        _ = fi73
                                    }
                                }
                                _ = v72
                            } else {
                                var v75 uint16 = 2
                                _ = len(v5)
                                v76 := GenS9{F0: 1, F1: 2}
                                _ = v75
                                _ = v76
                            }
                        }
                    } else {
                        var v77 int = 489
                        var v78 uint16 = 42
                        switch (v77 * v77) {
                        case -1:
                            if false {
                                var v79 float64 = 1.5
                                fmt.Println(v6, v77, v6)
                                _ = v79
                            }
                        case 850:
                            v77 = 9223372036854775807
                        case 32768:
                            var v80 int8 = 100
                            _ = v80
                        }
                        _ = v77
                        _ = v78
                    }
                    if (121 != 816) {
                        v6 = v6
                    } else {
                        if len(v3) > 0 {
                            v3[0] = ((v6 + "a") + v6)
                        }
                        switch "\n\t" {
                        case "\n\t":
                            var v81 int16 = 32767
                            _ = v81
                        case "hello":
                            v1 -= 1.5
                        }
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                }
                _ = v62
            } else {
                var v82 uint = 2
                v83 := GenS10{F0: 0, F1: 100, F2: 1}
                _ = v82
                _ = v83
            }
        }
        v1 = (v1 * (v1 + v1))
        v84 := GenS5{F0: 32768, F1: 42}
        _ = v84
    }
    v2[127] = v6
    delete(v2, 4294967295)
    for fi85 := 0; fi85 < 20; fi85++ {
        v86 := GenS11{F0: 4294967295}
        v3 = append(v3, (v6 + (("hello" + v6) + v6)))
        v2[fi85] = "\n\t"
        _ = fi85
        _ = v86
    }
    v3 = append(v3, v6)
    if true {
        v4 = (v4 / ((v4) | 1))
    }
    fmt.Println(v6, v4)
    switch 92 {
    case -9223372036854775808:
        _ = len(v3)
    default:
        v87 := GenS3{F0: 2147483648}
        _ = v87
    }
    v1 -= (-2.25 + v1)
    var v88 float32 = (v1 * (((float32(0.1) / 3.14159) * v1) + (0.0 / ((1.5 / ((v1 - 3.14159) / v1)) - v1))))
    if (!false) {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
}
