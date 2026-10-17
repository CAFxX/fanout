package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint8; F2 int16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS8 struct { F0 bool; F1 int32; F2 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 int8 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 bool, p1 int64, p2 int32) uint {
    for fi6 := 0; fi6 < 12; fi6++ {
        if (842 > 2) {
            for fi7 := 0; fi7 < 18; fi7++ {
                p2 = 42
                _ = fi7
            }
            p0 = p0
            if true {
                switch (("a" + (("a" + "") + "")) + (("" + "a") + ("a" + "a"))) {
                case "hello":
                    var v8 int = 391
                    _ = v8
                case "\n\t":
                    fmt.Println(p1)
                }
            } else {
                v9 := GenS1{F0: (fi6 * (fi6 | fi6)), F1: 127}
                v10 := []int{32767}
                _ = v10
                _ = v9
            }
        } else {
            var v11 uint8 = 127
            if (!true) {
                v12 := GenS1{F0: (fi6 | -164), F1: v11}
                var v13 bool = true
                v14 := []int{(-(fi6)), 32767}
                _ = v12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        _ = fi6
    }
    var v15 string = "hello"
    var v16 uint64 = 865
    switch "a" {
    case "a":
        var v17 int64 = -1
        _ = v17
    case "hello":
        if p0 {
            v18 := GenS2{F0: 68}
            var v19 float64 = (0.1 + 0.0)
            _ = v18
            _ = v19
        }
    case "\n\t":
        p2 -= -1
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        switch fi20 {
        case -171:
            v21 := []int{fi20, -2147483648}
            _ = v21
        case -774:
            if (uint(v16) >= 1) {
                switch int(p1) {
                case -129:
                    switch (fi20 % ((fi20) | 1)) {
                    case -815:
                        switch v15 {
                        case "a":
                            v22 := GenS1{F0: (^(fi20)), F1: 255}
                            _ = v22
                        case "hello":
                            v23 := []int{fi20, fi20}
                            _ = v23
                        default:
                            var v24 uint64 = (v16 << ((4294967295) & 63))
                            _ = v24
                        }
                    }
                }
                v25 := GenS1{F0: fi20, F1: 13}
                _ = v25
            }
        case 4294967295:
            p2 += (p2 / (((p2 * -603)) | 1))
        default:
            v26 := GenS3{F0: 413}
            _ = v26
        }
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 665
}
func genFn1(p0 int8, p1 int) uint32 {
    if (-9223372036854775808 > 9223372036854775807) {
        switch "hello" {
        case "foo bar":
            var v27 uint = genFn0(false, 42, 572)
            _ = v27
        case "":
            if (false && (true || false)) {
                v28 := GenS4{F0: p1, F1: 238, F2: 32767}
                var v29 float64 = 0.1
                v29 = v29
                _ = v28
                _ = v29
            } else {
                p0 = (^(p0))
                v30 := []int{}
                if ((-556 > 0) || (1 > -315)) {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        for fi32 := 0; fi32 < 5; fi32++ {
                            switch fi32 {
                            case 1:
                                v33 := GenS5{F0: 1, F1: -1}
                                _ = v33
                            case 2147483647:
                                fmt.Println(p0)
                            case -2147483648:
                                switch fi31 {
                                case -2147483648:
                                    var v34 string = (("hello" + ("\n\t" + "a")) + "\n\t")
                                    p0 = 100
                                    v35 := GenS6{F0: ("a" + ((v34 + ("\n\t" + "hello")) + v34)), F1: 42, F2: 179}
                                    _ = v34
                                    _ = v35
                                default:
                                    var v36 int32 = -750
                                    _ = v36
                                }
                            default:
                                v37 := GenS6{F0: "a", F1: -700, F2: 487}
                                _ = v37
                            }
                            var v38 int32 = 42
                            var v39 float64 = -2.25
                            p0 = 73
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                            if len(v30) > 0 {
                                v30[0] = int(v38)
                            }
                            var v40 string = "foo bar"
                            var v41 uint16 = 255
                            _ = fi32
                            _ = v38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        var v42 float32 = 0.0
                        _ = fi31
                        _ = v42
                    }
                    var v43 float32 = 3.14159
                    _ = v43
                } else {
                    p1 -= (^(p1))
                }
                _ = v30
            }
        case "\n\t":
            switch "a" {
            case "\n\t":
                var v44 uint32 = 177
                _ = v44
            case "a":
                v45 := []int{p1, p1}
                _ = v45
            case "foo bar":
                if false {
                    p1 += p1
                } else {
                    p0 = (p0 & (p0 / (((p0 * (p0 | (p0 >> ((0) & 63))))) | 1)))
                    v46 := []int{}
                    _ = v46
                }
            default:
                fmt.Println(p1)
            }
        default:
            v47 := []int{p1, (p1 << ((4294967295) & 63)), (-(p1)), (p1 & p1)}
            _ = v47
        }
        if false {
            switch "a" {
            case "":
                fmt.Println(p0, p1)
            case "\n\t":
                if false {
                    p1 = p1
                }
            case "foo bar":
                var v48 int64 = 9223372036854775807
                _ = v48
            default:
                p0 = (p0 | p0)
            }
            for fi49 := 0; fi49 < 8; fi49++ {
                p0 -= (p0 * (p0 | 1))
                if true {
                    p0 += 127
                } else {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        var v51 uint16 = 0
                        if (42 > 2) {
                            switch fi50 {
                            case -444:
                                var v52 int32 = 2147483647
                                _ = v52
                            case -9223372036854775808:
                                v53 := []int{(p1 ^ 164)}
                                _ = v53
                            case 4294967295:
                                switch fi50 {
                                case 9223372036854775807:
                                    p0 = 1
                                case 369:
                                    var v54 uint64 = 920
                                    _ = v54
                                case 273:
                                    v51 = (v51 << ((377) & 63))
                                }
                                v55 := []int{fi49, fi50, (p1 >> ((761) & 63))}
                                v51 = v51
                                _ = v55
                            default:
                                switch "\n\t" {
                                case "hello":
                                    p1 += fi49
                                case "\n\t":
                                    v56 := []int{}
                                    _ = v56
                                }
                            }
                            v57 := []int{fi50, (fi50 / (((fi49 / ((4294967295) | 1))) | 1)), fi50, (fi49 >> ((491) & 63))}
                            if false {
                                switch "hello" {
                                case "":
                                    var v58 string = "hello"
                                    _ = v58
                                case "foo bar":
                                    p0 += (p0 + p0)
                                default:
                                    p0 = -1
                                }
                                if true {
                                    if (-32768 >= 32767) {
                                        p0 += (p0 << ((genFn0(((float32((float64((0.0 + float32((0.1 - 3.14159)))) + (0.1 - 0.0))) == (1.5 - 3.14159)) && false), 9223372036854775807, 259)) & 63))
                                    }
                                    p1 = (p1 % ((fi49) | 1))
                                } else {
                                    if true {
                                        var v59 uint = 2
                                        var v60 int32 = -1
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 bool = true
                                        _ = v61
                                    }
                                    var v62 int16 = -32768
                                    _ = v62
                                }
                                if false {
                                    if false {
                                        fmt.Println(p1)
                                        v57 = append(v57, 2147483647)
                                    }
                                    if ((42 < 42) && (0 < 440)) {
                                        var v63 uint = 65535
                                        _ = v63
                                    } else {
                                        if ("hello" <= ("hello" + ("foo bar" + ((("foo bar" + (("hello" + ((("\n\t" + "") + "hello") + "foo bar")) + "hello")) + "") + ("\n\t" + (("hello" + ("" + ("\n\t" + "\n\t"))) + "foo bar")))))) {
                                            if len(v57) > 0 {
                                                v57[0] = -588
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                var v64 int64 = 0
                                                _ = v64
                                            case "foo bar":
                                                v65 := GenS7{F0: 2, F1: 4294967295, F2: v51}
                                                _ = v65
                                            }
                                            fmt.Println(fi50)
                                        } else {
                                            fmt.Println(p0, p0, fi49)
                                        }
                                    }
                                    if ("\n\t" >= "foo bar") {
                                        var v66 int64 = 42
                                        _ = v66
                                    } else {
                                        var v67 int = 218
                                        switch 42 {
                                        case 583:
                                            p0 -= p0
                                        case -1:
                                            v51 = (v51 % (((v51 << ((830) & 63))) | 1))
                                        case 165:
                                            v67 += p1
                                        }
                                        _ = v67
                                    }
                                }
                                if (!(9 <= (p0 ^ p0))) {
                                    p1 += (fi50 % (((p1 | fi49)) | 1))
                                    v68 := GenS8{F0: (false && (p0 >= (p0 | 124))), F1: 554, F2: 868}
                                    _ = v68
                                }
                                if (4294967295 < 4294967295) {
                                    if len(v57) >= 2 {
                                        v57 = v57[1:len(v57)]
                                    }
                                    var v69 uint = 42
                                    _ = v69
                                }
                                if false {
                                    var v70 int64 = -1
                                    _ = v70
                                }
                            } else {
                                v71 := GenS9{F0: (p0 << ((905) & 63))}
                                _ = v71
                            }
                            _ = v57
                        } else {
                            v51 += (v51 & (v51 + v51))
                            v72 := []int{fi50}
                            p1 = 1
                            _ = v72
                        }
                        _ = fi50
                        _ = v51
                    }
                }
                var v73 int64 = 42
                _ = fi49
                _ = v73
            }
        } else {
            for fi74 := 0; fi74 < 12; fi74++ {
                var v75 int8 = 59
                var v76 int = fi74
                if true {
                    var v77 string = ("\n\t" + "")
                    var v78 uint = 4294967295
                    var v79 int16 = 32767
                    _ = v77
                    _ = v78
                    _ = v79
                } else {
                    v80 := []int{}
                    var v81 int16 = 1000
                    _ = v80
                    _ = v81
                }
                _ = fi74
                _ = v75
                _ = v76
            }
            var v82 int64 = 9223372036854775807
            switch "" {
            case "\n\t":
                var v83 uint16 = 15
                _ = v83
            case "a":
                v84 := GenS10{F0: 42, F1: p0}
                _ = v84
            }
            _ = v82
        }
    }
    var v85 uint32 = 42
    _ = p0
    _ = p1
    _ = v85
    return v85
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = (1e10 - 3.14159)
    var v3 uint32 = 2
    var v4 float64 = v2
    var v5 int = 0
    var v86 float64 = 0.1
    v3 = v3
    v87 := GenS7{F0: uint8(v3), F1: genFn0(true, -1, 2147483647), F2: 2}
    var v88 int8 = 127
    v2 = float64((1e10 - -2.25))
    switch -319 {
    case -716:
        switch ("\n\t" + "") {
        case "":
            var v89 int16 = 923
            _ = v89
        case "hello":
            var v90 uint16 = 0
            _ = v90
        case "\n\t":
            for fi91 := 0; fi91 < 19; fi91++ {
                for fi92 := 0; fi92 < 10; fi92++ {
                    var v93 int = fi91
                    _ = fi92
                    _ = v93
                }
                _ = fi91
            }
        }
    case -2147483648:
        v94 := GenS11{F0: 67}
        _ = v94
    }
    if true {
        var v95 int16 = 42
        v4 -= v86
        var v96 bool = true
        _ = v95
        _ = v96
    } else {
        fmt.Println(v2)
    }
    v3 -= (v3 + v3)
    var v97 uint = uint(v3)
    v1[4294967295] = (v5 << ((v97) & 63))
    v97 = (v97 + v97)
    v86 = v86
    if (false || ("foo bar" <= "a")) {
        var v98 uint = v97
        fmt.Println(v4)
        _ = v98
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    fmt.Println("v88:", v88)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
    _ = v88
    _ = v97
}
