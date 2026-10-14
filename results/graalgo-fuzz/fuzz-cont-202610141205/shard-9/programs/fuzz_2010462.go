package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int32; F1 uint8; F2 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 bool; F1 int }
func genFn0() uint32 {
    var v7 float64 = -2.25
    var v8 string = ""
    if true {
        if (!true) {
            for fi9 := 0; fi9 < 5; fi9++ {
                switch v8 {
                case "a":
                    var v10 uint16 = 32767
                    _ = v10
                case "foo bar":
                    var v11 uint64 = 893
                    _ = v11
                case "hello":
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v13 := []int{fi12, (fi12 << ((0) & 63)), 127}
                        if true {
                            var v14 int64 = 1
                            v15 := GenS1{F0: (v14 * v14), F1: 31, F2: -2147483648}
                            _ = v14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                default:
                    if (0 > 29) {
                        v16 := GenS2{F0: 1000, F1: 65535, F2: int64(fi9)}
                        v17 := []int{(fi9 ^ (fi9 >> ((255) & 63))), fi9, fi9, fi9}
                        switch ((((v8 + "") + v8) + "hello") + (((v8 + "\n\t") + "a") + (v8 + "hello"))) {
                        case "":
                            var v18 bool = ((true || true) || false)
                            _ = v18
                        }
                        _ = v16
                        _ = v17
                    }
                }
                for fi19 := 0; fi19 < 18; fi19++ {
                    switch fi19 {
                    case 127:
                        if false {
                            v7 += 0.0
                        } else {
                            if (42 != 222) {
                                v20 := GenS1{F0: -1, F1: 42, F2: -2147483648}
                                switch -9223372036854775808 {
                                case 42:
                                    v8 = (v8 + "\n\t")
                                case 2147483647:
                                    v21 := GenS3{F0: 2, F1: 0, F2: 42}
                                    _ = v21
                                }
                                _ = v20
                            } else {
                                var v22 float64 = float64((0.0 / 1.5))
                                if (true || true) {
                                    v23 := []int{fi19, 32767}
                                    _ = v23
                                }
                                _ = v22
                            }
                        }
                    case 9223372036854775807:
                        v24 := GenS1{F0: 0, F1: 127, F2: 0}
                        fmt.Println(fi19)
                        v7 += ((v7 * (v7 - float64((float32(3.14159) + 3.14159)))) - float64(1.5))
                        var v25 uint8 = 20
                        fmt.Println(fi9)
                        _ = v24
                        _ = v25
                    default:
                        if false {
                            switch v8 {
                            case "a":
                                v7 += v7
                            case "":
                                var v26 bool = (127 > 255)
                                _ = v26
                            default:
                                if true {
                                    v8 = v8
                                    switch fi9 {
                                    case 2147483647:
                                        v27 := GenS1{F0: 42, F1: 128, F2: -2147483648}
                                        _ = v27
                                    }
                                } else {
                                    v28 := GenS4{F0: 4294967295}
                                    v7 -= v7
                                    v29 := GenS1{F0: -263, F1: 53, F2: 1}
                                    _ = v28
                                    _ = v29
                                }
                                v30 := GenS1{F0: -317, F1: 127, F2: 2147483647}
                                _ = v30
                            }
                            if true {
                                var v31 uint64 = 2
                                var v32 int64 = -934
                                _ = v31
                                _ = v32
                            } else {
                                fmt.Println(fi19, fi19, v8)
                                var v33 uint64 = 2
                                v33 += (v33 / (((v33 << ((473) & 63))) | 1))
                                _ = v33
                            }
                            if true {
                                v34 := GenS5{F0: 32767, F1: 41}
                                if true {
                                    if true {
                                        if true {
                                            var v35 uint16 = 228
                                            v36 := []int{(fi19 % ((fi9) | 1))}
                                            fmt.Println(fi19)
                                            _ = v35
                                            _ = v36
                                        } else {
                                            v37 := GenS4{F0: 727}
                                            _ = v37
                                        }
                                        v8 = ""
                                    }
                                    v38 := GenS2{F0: 32767, F1: 42, F2: -896}
                                    _ = v38
                                } else {
                                    v39 := []int{(fi19 ^ fi19), 32768, fi19}
                                    fmt.Println(fi9, v7)
                                    fmt.Println(v8, fi19)
                                    _ = v39
                                }
                                _ = v34
                            } else {
                                v8 = (v8 + "foo bar")
                            }
                            v7 += (v7 / v7)
                        } else {
                            v40 := GenS6{F0: -834, F1: 143, F2: 9223372036854775807}
                            var v41 int64 = 0
                            switch ((("hello" + "") + (v8 + (v8 + v8))) + v8) {
                            case "a":
                                v42 := GenS7{F0: -117}
                                _ = v42
                            case "":
                                v43 := GenS8{F0: 127, F1: 65535}
                                _ = v43
                            case "hello":
                                v7 = -2.25
                            default:
                                v8 = "\n\t"
                            }
                            _ = v40
                            _ = v41
                        }
                    }
                    _ = fi19
                }
                _ = fi9
            }
            v7 += v7
            if (!true) {
                v44 := []int{127, -2147483648, 9223372036854775807, 4294967295}
                if (!(!false)) {
                    if len(v44) > 0 {
                        v44[0] = 588
                    }
                    v8 = v8
                } else {
                    v8 = "\n\t"
                    for fi45 := 0; fi45 < 1; fi45++ {
                        var v46 float64 = v7
                        if (!(!true)) {
                            fmt.Println(fi45)
                        } else {
                            fmt.Println(v46)
                            var v47 int32 = -1
                            _ = v47
                        }
                        var v48 int64 = -1
                        _ = fi45
                        _ = v46
                        _ = v48
                    }
                }
                v49 := GenS2{F0: 83, F1: 65535, F2: -9223372036854775808}
                _ = v44
                _ = v49
            } else {
                for fi50 := 0; fi50 < 8; fi50++ {
                    v51 := []int{fi50, fi50}
                    if len(v51) > 0 {
                        v51[0] = fi50
                    }
                    if len(v51) >= 2 {
                        v51 = v51[1:len(v51)]
                    }
                    _ = fi50
                    _ = v51
                }
            }
        }
        v52 := GenS1{F0: 9223372036854775807, F1: 2, F2: 0}
        for fi53 := 0; fi53 < 2; fi53++ {
            var v54 uint32 = 176
            switch v8 {
            case "hello":
                for fi55 := 0; fi55 < 15; fi55++ {
                    v7 += v7
                    if (509 > 255) {
                        var v56 float64 = v7
                        switch v8 {
                        case "":
                            v57 := []int{-2147483648, (fi53 << ((41) & 63))}
                            var v58 uint16 = 0
                            var v59 uint64 = 1
                            var v60 int32 = -81
                            _ = v57
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        v61 := GenS8{F0: 255, F1: v54}
                        _ = v56
                        _ = v61
                    }
                    switch fi53 {
                    case -613:
                        v62 := []int{fi53, 956, fi53, (fi53 >> ((783) & 63))}
                        _ = v62
                    case 32768:
                        if true {
                            switch v8 {
                            case "foo bar":
                                v63 := GenS6{F0: 608, F1: 0, F2: 882}
                                _ = v63
                            case "hello":
                                v54 -= 65535
                                switch (v8 + v8) {
                                case "a":
                                    switch v8 {
                                    case "a":
                                        v64 := GenS9{F0: -1}
                                        _ = v64
                                    }
                                case "hello":
                                    if true {
                                        v65 := GenS10{F0: 4294967295, F1: 255}
                                        _ = v65
                                    }
                                default:
                                    var v66 bool = false
                                    _ = v66
                                }
                            case "a":
                                if (0 <= 18446744073709551615) {
                                    var v67 uint = 0
                                    v68 := []int{fi53, fi53}
                                    v69 := GenS9{F0: 42}
                                    if (true && ((v67 | 163) < (v67 + 2))) {
                                        v70 := GenS11{F0: false}
                                        _ = v70
                                    }
                                    var v71 uint = v67
                                    _ = v67
                                    _ = v68
                                    _ = v69
                                    _ = v71
                                } else {
                                    fmt.Println(fi55)
                                }
                            }
                        } else {
                            v54 = v54
                            var v72 int64 = 0
                            _ = v72
                        }
                    default:
                        var v73 float32 = 3.14159
                        _ = v73
                    }
                    _ = fi55
                }
            case "foo bar":
                v74 := []int{fi53}
                _ = v74
            default:
                var v75 int = (fi53 | fi53)
                _ = v75
            }
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    if (965 == 0) {
        if false {
            v8 = (v8 + ("\n\t" + v8))
            v7 += 3.14159
            switch 32768 {
            case 61:
                var v76 int = 106
                _ = v76
            }
        } else {
            if false {
                for fi77 := 0; fi77 < 6; fi77++ {
                    switch fi77 {
                    case 419:
                        var v78 uint16 = 255
                        _ = v78
                    }
                    for fi79 := 0; fi79 < 3; fi79++ {
                        v8 = v8
                        v7 -= ((v7 - float64(float32(1e10))) - float64(-2.25))
                        v7 -= v7
                        _ = fi79
                    }
                    _ = fi77
                }
                v80 := []int{-9223372036854775808, 2147483647}
                _ = v80
            }
        }
        if (true || true) {
            v81 := []int{127, -2147483648, 128}
            _ = v81
        } else {
            v7 -= v7
            switch 9223372036854775807 {
            case -2147483648:
                v82 := GenS12{F0: v8, F1: 830}
                _ = v82
            case 32768:
                v8 = "a"
            case -491:
                v7 -= -2.25
            default:
                var v83 int32 = 0
                _ = v83
            }
            v84 := []int{}
            _ = v84
        }
        var v85 string = (("a" + "\n\t") + v8)
        _ = v85
    }
    v7 = (1e10 / float64((((((0.0 + 1.5) / (3.14159 / 1.0)) - float32(float64(0.0))) * 0.1) / 0.1)))
    _ = v7
    _ = v8
    return 780
}
func genFn1(p0 int) uint32 {
    p0 += 9223372036854775807
    v86 := []int{(-(p0)), p0, p0, (p0 << ((520) & 63))}
    if (871 >= 9223372036854775807) {
        p0 += 2147483647
    } else {
        p0 = (p0 % (((p0 & (p0 / ((p0) | 1)))) | 1))
        var v87 int = p0
        _ = v87
    }
    fmt.Println(p0)
    _ = p0
    _ = v86
    return 4294967295
}

func main() {
    var v1 uint8 = 42
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    var v4 string = "hello"
    var v5 uint32 = 1
    v6 := make(map[uint32]string)
    _ = len(v3)
    var v88 int64 = -9223372036854775808
    switch v4 {
    case "":
        v88 -= (v88 / ((v88) | 1))
    case "foo bar":
        v1 = 2
    case "\n\t":
        v5 = (v5 ^ v5)
    }
    v89 := []int{}
    v1 = v1
    if true {
        for fi90 := 0; fi90 < 9; fi90++ {
            v2 = 1
            var v91 int = fi90
            _ = fi90
            _ = v91
        }
    }
    v92 := v6[(^(v5))]
    v89 = append(v89, -1)
    _ = len(v89)
    var v93 uint16 = 1
    if ((!(1 < -1)) && false) {
        var v94 string = v4
        _ = v94
    } else {
        v5 = v5
        v95 := GenS2{F0: -275, F1: (v93 | v93), F2: -9223372036854775808}
        v2 += (v2 % ((-640) | 1))
        _ = v95
    }
    v96 := GenS13{F0: true, F1: 216}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v88:", v88)
    fmt.Println("slicelen v89:", len(v89), v89)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v89
    _ = v92
    _ = v93
    _ = v96
}
