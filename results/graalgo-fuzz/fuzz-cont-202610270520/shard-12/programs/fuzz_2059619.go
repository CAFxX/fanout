package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64; F1 int16; F2 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 uint32; F1 uint16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint; F1 int8 }
func genFn0(p0 uint32, p1 uint8) bool {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 int16 = 32767
        var v6 int8 = 67
        p1 = (p1 & p1)
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{-1, 660, 314}
    v8 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 string) int32 {
    switch ("" + p0) {
    case "foo bar":
        fmt.Println(p0)
    case "\n\t":
        switch "a" {
        case "\n\t":
            for fi9 := 0; fi9 < 15; fi9++ {
                p0 = "a"
                var v10 int32 = -632
                fmt.Println(p0, fi9)
                _ = fi9
                _ = v10
            }
        case "foo bar":
            p0 = p0
        case "hello":
            switch ((p0 + ("foo bar" + "hello")) + ((p0 + ((p0 + ((("a" + "a") + ((p0 + p0) + "foo bar")) + "foo bar")) + (((p0 + ("\n\t" + "hello")) + p0) + "a"))) + (("a" + ("hello" + ("hello" + "a"))) + ""))) {
            case "":
                fmt.Println(p0)
            case "a":
                p0 = p0
            default:
                p0 = (((p0 + p0) + "hello") + ("hello" + (p0 + p0)))
            }
        default:
            v11 := []int{406, 4294967295, 127}
            _ = v11
        }
    case "hello":
        switch 466 {
        case -1:
            p0 = p0
        case 247:
            v12 := []int{}
            _ = v12
        }
    }
    fmt.Println(p0)
    var v13 uint = 65535
    _ = p0
    _ = v13
    return 2147483647
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]int)
    var v3 float32 = 1.5
    delete(v2, "")
    v14 := GenS2{F0: 182, F1: -9223372036854775808, F2: 128}
    v15 := GenS3{F0: 255}
    switch ("hello" + "\n\t") {
    case "foo bar":
        v3 += 0.0
    default:
        v1 += ((v1 - 1.5) * v1)
    }
    v3 = v3
    v16 := []int{128, -320}
    var v17 float32 = v3
    if false {
        if false {
            switch 32767 {
            case 120:
                _ = len(v16)
            default:
                var v18 int8 = 0
                _ = v18
            }
            var v19 int8 = 0
            _ = v19
        }
        if true {
            _ = len(v2)
            switch "" {
            case "\n\t":
                v20 := GenS2{F0: 2, F1: 9223372036854775807, F2: -921}
                _ = v20
            case "":
                switch -816 {
                case 223:
                    var v21 uint64 = 325
                    _ = v21
                default:
                    v3 -= ((0.1 + 0.1) - (v17 + (((float32(v1) * float32(v1)) + v17) - v17)))
                }
            case "foo bar":
                var v22 uint32 = 61
                _ = v22
            default:
                v1 -= v1
            }
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            v3 += ((v3 - (3.14159 + v17)) / float32(-2.25))
            _ = fi23
        }
    }
    if true {
        var v24 float32 = float32((v1 - v1))
        v3 += v24
        if (!true) {
            for fi25 := 0; fi25 < 3; fi25++ {
                v17 -= (v3 - v3)
                fmt.Println(v24)
                v26 := v2["hello"]
                _ = fi25
                _ = v26
            }
            v24 = v24
            v24 -= v3
        } else {
            var v27 int = 32768
            _ = v27
        }
        _ = v24
    } else {
        for fi28 := 0; fi28 < 20; fi28++ {
            for fi29 := 0; fi29 < 12; fi29++ {
                switch fi29 {
                case 2147483647:
                    var v30 float64 = v1
                    _ = v30
                case 4294967295:
                    v1 = v1
                    switch fi29 {
                    case 1:
                        fmt.Println(v17, fi29)
                    case 42:
                        v31 := GenS1{F0: 1}
                        _ = v31
                    case 9223372036854775807:
                        var v32 int64 = 1
                        _ = v32
                    default:
                        var v33 uint16 = 65535
                        _ = v33
                    }
                default:
                    _ = len(v16)
                }
                fmt.Println(fi29, v3, v1)
                switch ("a" + "") {
                case "":
                    v3 -= (1e10 / 1.5)
                case "foo bar":
                    var v34 float32 = (v17 - ((((3.14159 / (v17 - 0.1)) + 3.14159) + 1e10) / float32(1e10)))
                    _ = v34
                default:
                    if false {
                        var v35 float64 = (float64(((v3 - -2.25) / (v3 / 1.0))) / (v1 - (1e10 / v1)))
                        v1 = v1
                        var v36 bool = (32767 > 946)
                        _ = v35
                        _ = v36
                    }
                }
                v1 = 0.1
                if true {
                    v3 -= (v3 + v3)
                    _ = len(v16)
                    v1 = -2.25
                    var v37 string = "foo bar"
                    if false {
                        if (true || false) {
                            if true {
                                v1 = (v1 + (v1 - v1))
                                switch fi28 {
                                case 42:
                                    if true {
                                        v17 -= (v3 / (v3 + ((((float32(float64(v3)) / v3) + -2.25) / v3) - (v3 - 0.1))))
                                        if (332 != 1) {
                                            v1 = v1
                                            var v38 uint = 2
                                            _ = v38
                                        } else {
                                            var v39 uint8 = 127
                                            _ = v39
                                        }
                                        switch fi29 {
                                        case 127:
                                            v37 = "foo bar"
                                        case 119:
                                            if true {
                                                var v40 int32 = 0
                                                v1 = v1
                                                switch 9223372036854775807 {
                                                case -600:
                                                    v1 -= 0.1
                                                }
                                                _ = v40
                                            } else {
                                                var v41 float64 = v1
                                                if len(v16) > 0 {
                                                    v16[0] = fi29
                                                }
                                                var v42 int32 = -2147483648
                                                _ = v41
                                                _ = v42
                                            }
                                        }
                                        v1 = 1e10
                                        var v43 uint8 = 197
                                        _ = v43
                                    }
                                case -193:
                                    if (true && true) {
                                        var v44 int16 = 0
                                        var v45 int32 = 222
                                        switch 9223372036854775807 {
                                        case 42:
                                            v44 -= v44
                                        case -368:
                                            if (fi28 > 42) {
                                                fmt.Println(v3)
                                                v46 := v2[v37]
                                                if false {
                                                    var v47 string = v37
                                                    switch (v46 << ((18446744073709551615) & 63)) {
                                                    case 2147483647:
                                                        fmt.Println(fi29, fi28, v45)
                                                    case 571:
                                                        if true {
                                                            var v48 float32 = float32(1.5)
                                                            if (("foo bar" + v37) >= ("\n\t" + ("foo bar" + "a"))) {
                                                                v44 -= 0
                                                                var v49 bool = false
                                                                v1 -= (0.1 * v1)
                                                                _ = v49
                                                            }
                                                            if true {
                                                                v50 := GenS1{F0: (v45 | genFn1("a"))}
                                                                switch -510 {
                                                                case 838:
                                                                    fmt.Println(v3, fi29, fi28)
                                                                case 1:
                                                                    v51 := GenS4{F0: 127, F1: 1, F2: (v45 | (v45 % (((v45 % (((v45 * v45)) | 1))) | 1)))}
                                                                    _ = v51
                                                                default:
                                                                    var v52 int8 = 42
                                                                    _ = v52
                                                                }
                                                                _ = v50
                                                            } else {
                                                                v3 -= v3
                                                            }
                                                            _ = v48
                                                        }
                                                        var v53 uint8 = 4
                                                        if ((v44 < (v44 | v44)) && ((v45 + v45) < (v45 % ((-587) | 1)))) {
                                                            v54 := GenS5{F0: v45}
                                                            _ = len(v16)
                                                            var v55 int32 = v45
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        _ = v53
                                                    case -1:
                                                        v56 := GenS1{F0: genFn1("foo bar")}
                                                        _ = v56
                                                    default:
                                                        v57 := GenS5{F0: (v45 * v45)}
                                                        _ = v57
                                                    }
                                                    v17 = 3.14159
                                                    if false {
                                                        if false {
                                                            switch -119 {
                                                            case 127:
                                                                if (-92 <= 0) {
                                                                    v45 -= (v45 << ((18446744073709551615) & 63))
                                                                    v58 := GenS3{F0: 42}
                                                                    if (635 > 772) {
                                                                        v59 := GenS6{F0: 614, F1: v44, F2: 617}
                                                                        _ = v59
                                                                    } else {
                                                                        var v60 uint16 = 0
                                                                        v3 -= 1e10
                                                                        _ = v60
                                                                    }
                                                                    _ = v58
                                                                } else {
                                                                    fmt.Println(fi29, v45)
                                                                    if false {
                                                                        if false {
                                                                            fmt.Println(v45)
                                                                            fmt.Println(fi29, v47, v1)
                                                                        }
                                                                        v61 := GenS4{F0: -1, F1: 255, F2: (-(v45))}
                                                                        if (127 == -1) {
                                                                            var v62 int8 = 0
                                                                            _ = v62
                                                                        } else {
                                                                            switch (v47 + ("\n\t" + v47)) {
                                                                            case "":
                                                                                v63 := GenS7{F0: -1}
                                                                                _ = v63
                                                                            case "a":
                                                                                switch (fi28 & fi29) {
                                                                                case 42:
                                                                                    switch v37 {
                                                                                    case "":
                                                                                        var v64 uint = 4294967295
                                                                                        _ = v64
                                                                                    case "a":
                                                                                        v44 += v44
                                                                                    default:
                                                                                        var v65 int = -909
                                                                                        _ = v65
                                                                                    }
                                                                                default:
                                                                                    var v66 uint32 = 65535
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                        }
                                                                        switch ("" + v47) {
                                                                        case "hello":
                                                                            if (v45 < 1) {
                                                                                fmt.Println(v45, v46, fi29)
                                                                                v67 := GenS8{F0: v46}
                                                                                var v68 uint32 = 991
                                                                                _ = v67
                                                                                _ = v68
                                                                            } else {
                                                                                var v69 int64 = -9223372036854775808
                                                                                _ = v69
                                                                            }
                                                                        case "":
                                                                            v3 = v3
                                                                        }
                                                                        _ = v61
                                                                    }
                                                                    v70 := GenS9{F0: 0}
                                                                    _ = v70
                                                                }
                                                            case -1:
                                                                if false {
                                                                    var v71 int32 = -607
                                                                    _ = v71
                                                                } else {
                                                                    v17 -= 0.0
                                                                    v3 -= (1.5 / 1.5)
                                                                    v47 = (("\n\t" + v47) + v47)
                                                                }
                                                            default:
                                                                v72 := GenS10{F0: (^(v45)), F1: 2147483647}
                                                                _ = v72
                                                            }
                                                            v1 = (1e10 - 0.0)
                                                            _ = len(v2)
                                                            v73 := GenS10{F0: (v45 / (((-(v45))) | 1)), F1: 2147483648}
                                                            _ = v73
                                                        }
                                                    }
                                                    v74 := GenS1{F0: -265}
                                                    _ = v47
                                                    _ = v74
                                                } else {
                                                    var v75 uint64 = 76
                                                    _ = v75
                                                }
                                                _ = len(v2)
                                                _ = v46
                                            } else {
                                                v76 := GenS11{F0: -128, F1: 42}
                                                v45 = v45
                                                _ = v76
                                            }
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                case 128:
                                    if len(v16) > 0 {
                                        v16[0] = (^(fi29))
                                    }
                                default:
                                    v77 := GenS1{F0: 42}
                                    _ = v77
                                }
                                v17 += float32(v1)
                            }
                            if (!true) {
                                var v78 int32 = 42
                                v79 := GenS12{F0: 2, F1: 255}
                                switch ("a" + ("" + (v37 + v37))) {
                                case "\n\t":
                                    var v80 int16 = 0
                                    _ = v80
                                }
                                _ = v78
                                _ = v79
                            }
                            v17 = ((float32((3.14159 * 0.1)) - v17) + v3)
                            var v81 float32 = v17
                            _ = v81
                        } else {
                            var v82 int16 = 123
                            _ = v82
                        }
                    } else {
                        var v83 uint16 = 1
                        var v84 int8 = 0
                        v85 := GenS12{F0: 2147483648, F1: (^(v83))}
                        v1 = 3.14159
                        v86 := v2["a"]
                        _ = v83
                        _ = v84
                        _ = v85
                        _ = v86
                    }
                    var v87 float32 = v17
                    _ = v37
                    _ = v87
                } else {
                    switch "a" {
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            var v88 float32 = v3
                            _ = v88
                        default:
                            var v89 int32 = -1
                            _ = v89
                        }
                    default:
                        delete(v2, "")
                        delete(v2, "\n\t")
                    }
                }
                v1 -= v1
                _ = fi29
            }
            _ = fi28
        }
        for fi90 := 0; fi90 < 6; fi90++ {
            v91 := GenS13{F0: -1}
            _ = fi90
            _ = v91
        }
        switch 1 {
        case 127:
            if (1 <= 1) {
                var v92 uint32 = 780
                v93 := GenS14{F0: 0, F1: 42}
                var v94 float64 = v1
                _ = v92
                _ = v93
                _ = v94
            }
        }
    }
    for fi95 := 0; fi95 < 19; fi95++ {
        v2[(("\n\t" + "a") + "")] = (^(fi95))
        v96 := v2["\n\t"]
        _ = fi95
        _ = v96
    }
    var v97 int64 = -94
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v97
}
