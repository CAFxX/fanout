package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint16; F2 int64 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
func genFn0(p0 uint64) int32 {
    if false {
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 uint = 1
            _ = fi4
            _ = v5
        }
        switch "hello" {
        case "hello":
            v6 := []int{-861, 126, 1, -942}
            _ = v6
        default:
            switch 32767 {
            case -47:
                var v7 int16 = 1000
                _ = v7
            }
        }
        v8 := GenS1{F0: 127}
        _ = v8
    } else {
        p0 += (p0 & 1)
    }
    p0 -= p0
    _ = p0
    return 50
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 1
    var v3 float32 = (((0.1 * 1.5) * (0.1 - 0.0)) / -2.25)
    for fi9 := 0; fi9 < 18; fi9++ {
        switch "hello" {
        case "foo bar":
            if false {
                var v10 uint = 2
                if true {
                    if ((false || (544 > 987)) && (36 >= 255)) {
                        if false {
                            var v11 int32 = 1
                            v12 := GenS2{F0: 42, F1: v11}
                            fmt.Println(v10, v10)
                            _ = v11
                            _ = v12
                        }
                        v13 := GenS3{F0: v10}
                        if (4294967295 <= 2) {
                            v14 := GenS2{F0: 25, F1: 370}
                            _ = v14
                        } else {
                            if (!(625 >= 42)) {
                                var v15 uint16 = 32767
                                if false {
                                    if (1 <= -9223372036854775808) {
                                        v16 := []int{}
                                        _ = v16
                                    }
                                    switch (("\n\t" + "foo bar") + ("hello" + ("a" + ""))) {
                                    case "hello":
                                        v17 := []int{fi9, (fi9 + fi9)}
                                        _ = v17
                                    case "foo bar":
                                        for fi18 := 0; fi18 < 10; fi18++ {
                                            if (v2 <= v2) {
                                                delete(v1, "a")
                                            } else {
                                                switch fi9 {
                                                case 218:
                                                    v10 += 2
                                                case 32768:
                                                    fmt.Println(fi18)
                                                case 128:
                                                    v19 := GenS3{F0: v10}
                                                    _ = v19
                                                }
                                                var v20 int16 = -32768
                                                _ = v20
                                            }
                                            var v21 float32 = ((v3 / -2.25) * -2.25)
                                            _ = fi18
                                            _ = v21
                                        }
                                    }
                                }
                                fmt.Println(v15, v3, v15)
                                _ = v15
                            }
                            var v22 int16 = -1
                            _ = v22
                        }
                        _ = v13
                    } else {
                        if (1e10 >= 0.1) {
                            if true {
                                v3 += (0.0 * (-2.25 * float32(((0.1 + float64(0.1)) / 0.1))))
                                v23 := GenS4{F0: fi9, F1: 42, F2: -1}
                                _ = v23
                            } else {
                                delete(v1, "a")
                                var v24 uint16 = 813
                                _ = v24
                            }
                            v3 -= v3
                        } else {
                            fmt.Println(v3, fi9)
                            v10 += v10
                        }
                    }
                } else {
                    for fi25 := 0; fi25 < 2; fi25++ {
                        v26 := []int{v2}
                        var v27 bool = (967 == 1)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                v3 += 1e10
                _ = v10
            } else {
                v28 := GenS3{F0: 0}
                v29 := []int{}
                _ = v28
                _ = v29
            }
        case "hello":
            var v30 int16 = 1000
            _ = v30
        default:
            var v31 string = "foo bar"
            _ = v31
        }
        var v32 float32 = (1e10 * ((v3 + float32(3.14159)) - 1e10))
        _ = fi9
        _ = v32
    }
    v3 = (((1.5 / (-2.25 + v3)) * v3) * ((0.0 - (((((3.14159 / 0.1) * 3.14159) / 1.5) / (float32(((0.1 * (-2.25 * (float64(3.14159) / ((float64(v3) * 0.1) * -2.25)))) + (0.0 / 0.1))) + v3)) / v3)) / 3.14159))
    var v33 int32 = 2147483647
    var v34 bool = true
    if false {
        v35 := GenS5{F0: v2, F1: ""}
        v34 = true
        _ = v35
    } else {
        var v36 int16 = 1000
        if (v34 || ((float64((0.0 * (3.14159 - (((1.5 - v3) + v3) / v3)))) * 1.5) < (0.1 + (0.1 * (((0.0 - (float64(((0.1 + (v3 * (1.5 - v3))) + (v3 + float32(0.0)))) * 1.5)) / float64((float32((1e10 * 3.14159)) + (v3 + (float32(float64(v3)) + v3))))) - (((0.1 * (((-2.25 + (0.0 / (float64(v3) + -2.25))) - -2.25) - 0.0)) + (float64(0.0) * 1e10)) / 0.1)))))) {
            for fi37 := 0; fi37 < 12; fi37++ {
                v38 := GenS6{F0: "a", F1: -128}
                for fi39 := 0; fi39 < 3; fi39++ {
                    if false {
                        v34 = v34
                        if true {
                            v40 := GenS4{F0: fi37, F1: 32767, F2: 1}
                            _ = v40
                        } else {
                            fmt.Println(v36, fi37, fi39)
                        }
                    } else {
                        v34 = ((v36 | (^(v36))) == v36)
                    }
                    if (!(2 >= 65535)) {
                        v41 := v1[("" + "\n\t")]
                        v42 := GenS6{F0: "foo bar", F1: -55}
                        var v43 float64 = 0.1
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        v3 = (v3 / -2.25)
                        fmt.Println(v2, fi37, v2)
                    }
                    _ = fi39
                }
                _ = fi37
                _ = v38
            }
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 int = (v2 >> ((255) & 63))
                _ = fi44
                _ = v45
            }
        } else {
            switch v2 {
            case -894:
                var v46 int64 = 9223372036854775807
                _ = v46
            case 1:
                v47 := []int{(v2 & v2)}
                _ = v47
            case 4294967295:
                v48 := []int{v2, int(v33), 9223372036854775807}
                _ = v48
            default:
                fmt.Println(v34, v2)
            }
            if true {
                v36 -= v36
                v3 = v3
            } else {
                v36 = int16(v33)
            }
            switch "" {
            case "":
                fmt.Println(v36, v33, v34)
            case "a":
                switch 4294967295 {
                case 127:
                    v49 := GenS7{F0: v2}
                    _ = v49
                default:
                    if false {
                        if true {
                            v2 -= -1
                            v50 := []int{}
                            _ = v50
                        }
                        v51 := v1["foo bar"]
                        _ = v51
                    }
                }
            case "foo bar":
                v52 := []int{}
                _ = v52
            default:
                if (!true) {
                    var v53 int32 = (v33 * -1)
                    _ = v53
                } else {
                    v54 := []int{v2, 0, (^(v2)), v2}
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                    _ = v54
                }
            }
        }
        _ = v36
    }
    var v55 int8 = 127
    switch ("" + "a") {
    case "a":
        var v56 uint8 = 0
        _ = v56
    case "\n\t":
        if true {
            v33 -= (v33 << ((25) & 63))
            if true {
                for fi57 := 0; fi57 < 5; fi57++ {
                    if v34 {
                        v55 += v55
                    } else {
                        v3 += 1e10
                    }
                    v55 = -128
                    switch "a" {
                    case "":
                        if (65535 > 172) {
                            v33 = -256
                            v58 := GenS7{F0: (v2 - 1)}
                            _ = v58
                        } else {
                            var v59 uint64 = 1
                            if ((1.5 - 1e10) != float64(((v3 + v3) - (v3 + ((float32(float64(v3)) / (float32(float64(1e10)) / v3)) + 0.1))))) {
                                var v60 string = "a"
                                fmt.Println(v59, v59)
                                _ = v60
                            } else {
                                var v61 int8 = (v55 / ((v55) | 1))
                                var v62 uint8 = 42
                                _ = v61
                                _ = v62
                            }
                            v55 -= v55
                            _ = v59
                        }
                    case "foo bar":
                        var v63 string = "a"
                        _ = v63
                    }
                    _ = fi57
                }
                var v64 int64 = 42
                _ = v64
            } else {
                v65 := []int{v2, (v2 ^ -9223372036854775808)}
                v34 = (!true)
                v55 += (-(v55))
                _ = v65
            }
            fmt.Println(v33, v33)
        }
    case "":
        fmt.Println(v3)
    default:
        var v66 int32 = (-(v33))
        _ = v66
    }
    switch v2 {
    case 128:
        var v67 bool = (287 > 840)
        _ = v67
    case -2147483648:
        var v68 uint32 = 1
        _ = v68
    case 4294967295:
        fmt.Println(v33, v34, v55)
    default:
        for fi69 := 0; fi69 < 7; fi69++ {
            fmt.Println(v3)
            v70 := []int{-957, fi69}
            _ = fi69
            _ = v70
        }
    }
    fmt.Println(v34, v55, v34)
    for fi71 := 0; fi71 < 8; fi71++ {
        if (!true) {
            fmt.Println(v33, v2)
            v72 := []int{(^(v2)), 128, 32767, v2}
            if false {
                for fi73 := 0; fi73 < 6; fi73++ {
                    switch fi71 {
                    case -65:
                        _ = len(v1)
                    case 388:
                        fmt.Println(v3)
                    case 42:
                        var v74 string = "foo bar"
                        _ = v74
                    }
                    if (v3 >= v3) {
                        var v75 float32 = v3
                        v76 := v1["a"]
                        switch (fi73 | (fi71 / ((9223372036854775807) | 1))) {
                        case 2147483647:
                            var v77 int8 = v55
                            v78 := v1[(((("hello" + (("a" + ((((("\n\t" + "hello") + (("foo bar" + (("" + "\n\t") + "")) + ("hello" + ("hello" + "a")))) + "\n\t") + "\n\t") + ("foo bar" + ("foo bar" + "")))) + ("foo bar" + ""))) + "\n\t") + "") + "foo bar")]
                            _ = v77
                            _ = v78
                        default:
                            _ = len(v1)
                            fmt.Println(v75, v34)
                            if ((!(2 > 680)) || v34) {
                                v75 += float32(1e10)
                                var v79 uint16 = 484
                                if (647 < 0) {
                                    fmt.Println(v79, fi73, v76)
                                    v80 := GenS8{F0: "foo bar"}
                                    v55 = v55
                                    _ = v80
                                }
                                _ = v79
                            } else {
                                switch v76 {
                                case -1:
                                    v81 := GenS1{F0: -2147483648}
                                    _ = v81
                                case 32767:
                                    fmt.Println(fi71, fi73, v33)
                                default:
                                    var v82 uint32 = 65535
                                    _ = v82
                                }
                                var v83 int32 = v33
                                _ = v83
                            }
                        }
                        _ = v75
                        _ = v76
                    }
                    switch 653 {
                    case 9223372036854775807:
                        var v84 string = ("hello" + "foo bar")
                        _ = v84
                    }
                    fmt.Println(v33, v34)
                    _ = fi73
                }
                var v85 int8 = 0
                _ = v85
            }
            _ = v72
        } else {
            switch (-(fi71)) {
            case 2147483647:
                var v86 uint64 = 163
                _ = v86
            case -712:
                fmt.Println(v55, v55, fi71)
            case 32767:
                v87 := []int{v2, (-(v2))}
                _ = v87
            }
            v88 := v1[("a" + "hello")]
            if v34 {
                v33 += genFn0(650)
            } else {
                if (!v34) {
                    var v89 int8 = 100
                    _ = v89
                }
                v90 := []int{int(v55), (v88 + (v88 / ((v88) | 1)))}
                _ = v90
            }
            _ = v88
        }
        v91 := []int{(v2 << ((156) & 63)), (fi71 * (v2 - (v2 % ((-550) | 1))))}
        fmt.Println(v34, v55, v2)
        _ = fi71
        _ = v91
    }
    for fi92 := 0; fi92 < 10; fi92++ {
        var v93 uint32 = 739
        _ = fi92
        _ = v93
    }
    var v94 int64 = 9223372036854775807
    v94 = v94
    for fi95 := 0; fi95 < 11; fi95++ {
        var v96 uint8 = 128
        for fi97 := 0; fi97 < 13; fi97++ {
            v2 -= fi97
            _ = len(v1)
            _ = fi97
        }
        _ = fi95
        _ = v96
    }
    v98 := GenS9{F0: (v55 >> ((519) & 63))}
    var v99 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v55:", v55)
    fmt.Println("v94:", v94)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v55
    _ = v94
    _ = v98
    _ = v99
}
