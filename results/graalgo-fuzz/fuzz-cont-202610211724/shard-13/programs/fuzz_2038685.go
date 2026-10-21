package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int32; F1 int; F2 uint }
type GenS6 struct { F0 uint; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32; F1 uint8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 bool }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                v9 := []int{}
                _ = v9
            } else {
                var v10 float32 = 0.1
                if ((18446744073709551615 < 1) || (true || true)) {
                    v11 := []int{-9223372036854775808, fi7}
                    _ = v11
                }
                _ = v10
            }
            var v12 uint16 = 2
            _ = fi8
            _ = v12
        }
        var v13 uint8 = 1
        if false {
            var v14 uint8 = v13
            _ = v14
        } else {
            if false {
                if true {
                    for fi15 := 0; fi15 < 3; fi15++ {
                        v13 -= 0
                        switch 0 {
                        case -63:
                            if true {
                                if (("" == "hello") && (!(125 >= 18446744073709551615))) {
                                    v13 = (v13 - (v13 + 51))
                                } else {
                                    var v16 uint = 170
                                    v13 += (v13 % (((v13 / ((2) | 1))) | 1))
                                    v13 += (v13 >> ((481) & 63))
                                    _ = v16
                                }
                            }
                        }
                        v17 := []int{(fi15 & fi15), 1, fi7}
                        _ = fi15
                        _ = v17
                    }
                }
            } else {
                if false {
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        var v18 float32 = -2.25
                        _ = v18
                    }
                }
            }
        }
        _ = fi7
        _ = v13
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint8 = 45
        _ = fi19
        _ = v20
    }
    if true {
        v21 := []int{-1, -1, 128}
        v22 := GenS1{F0: 2, F1: "", F2: 70}
        switch "" {
        case "\n\t":
            if false {
                for fi23 := 0; fi23 < 17; fi23++ {
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                    _ = fi23
                }
                v21 = append(v21, 32767)
            }
        case "foo bar":
            v24 := GenS1{F0: 4294967295, F1: "", F2: 255}
            _ = v24
        default:
            if len(v21) > 0 {
                v21[0] = -223
            }
        }
        _ = v21
        _ = v22
    }
    return 42
}
func genFn1(p0 int8, p1 int) uint16 {
    v25 := GenS1{F0: 1, F1: "hello", F2: 1}
    for fi26 := 0; fi26 < 10; fi26++ {
        p1 = (-(fi26))
        if false {
            p1 = fi26
            switch (fi26 | -500) {
            case -2147483648:
                v27 := []int{}
                _ = v27
            }
        }
        if (fi26 < p1) {
            if true {
                for fi28 := 0; fi28 < 8; fi28++ {
                    if (("hello" + ("hello" + ("hello" + "\n\t"))) != "\n\t") {
                        v29 := GenS2{F0: 26, F1: 258, F2: 0}
                        fmt.Println(p0, p0)
                        if false {
                            p0 = p0
                        }
                        _ = v29
                    } else {
                        var v30 float64 = 1.5
                        if true {
                            fmt.Println(fi26, p0, fi26)
                            fmt.Println(p0, fi26, v30)
                        }
                        v31 := GenS2{F0: 2, F1: -668, F2: 913}
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS1{F0: 1, F1: "", F2: 40}
                    fmt.Println(p0)
                    _ = fi28
                    _ = v32
                }
                v33 := GenS3{F0: 208}
                _ = v33
            }
        } else {
            p0 -= 45
            if (false || false) {
                v34 := GenS4{F0: "a", F1: -2147483648}
                var v35 uint = 18446744073709551615
                var v36 uint = v35
                _ = v34
                _ = v35
                _ = v36
            } else {
                if false {
                    var v37 int8 = (p0 - (p0 << ((4294967295) & 63)))
                    _ = v37
                }
            }
        }
        _ = fi26
    }
    v38 := GenS1{F0: 667, F1: ("\n\t" + "a"), F2: 128}
    _ = p0
    _ = p1
    _ = v25
    _ = v38
    return 255
}

func main() {
    var v1 int32 = 1
    var v2 float32 = ((1e10 / 0.1) + ((float32(float64(-2.25)) / -2.25) + ((-2.25 * 1e10) + float32(3.14159))))
    var v3 int16 = -1
    var v4 bool = false
    var v5 float64 = (-2.25 - (-2.25 + 3.14159))
    var v6 string = "hello"
    if (v4 && v4) {
        var v39 uint64 = 172
        _ = v39
    }
    v3 += -1
    if v4 {
        var v40 int16 = (v3 - v3)
        _ = v40
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := GenS4{F0: (((("foo bar" + "foo bar") + ((("a" + v6) + v6) + v6)) + ((v6 + "") + "\n\t")) + ("\n\t" + "a")), F1: 0}
        for fi43 := 0; fi43 < 10; fi43++ {
            var v44 uint8 = 0
            v45 := GenS4{F0: v6, F1: (v1 / (((v1 | v1)) | 1))}
            _ = fi43
            _ = v44
            _ = v45
        }
        fmt.Println(v4)
        _ = fi41
        _ = v42
    }
    for fi46 := 0; fi46 < 6; fi46++ {
        var v47 float64 = (((1.5 * float64(v2)) / (float64(v2) + (1e10 / v5))) - (((((v5 + v5) + v5) / -2.25) - v5) + float64((v2 - v2))))
        var v48 uint64 = 18446744073709551615
        _ = fi46
        _ = v47
        _ = v48
    }
    fmt.Println(v4, v6, v6)
    if v4 {
        v1 += (v1 + 0)
    } else {
        v49 := []int{0, 42}
        v2 -= v2
        if len(v49) > 0 {
            v49[0] = -9223372036854775808
        }
        _ = v49
    }
    var v50 uint = 65535
    var v51 uint = (v50 << ((2) & 63))
    var v52 int16 = (v3 % ((1) | 1))
    for fi53 := 0; fi53 < 1; fi53++ {
        if (1 == 1) {
            for fi54 := 0; fi54 < 14; fi54++ {
                if false {
                    if ((true || (-100 == -103)) || v4) {
                        v1 = (v1 + -125)
                    } else {
                        v55 := []int{fi54}
                        switch v6 {
                        case "hello":
                            switch fi54 {
                            case 42:
                                var v56 uint64 = 2
                                _ = v56
                            case -2147483648:
                                v57 := GenS5{F0: (^(v1)), F1: -388, F2: 18446744073709551615}
                                _ = v57
                            default:
                                if v4 {
                                    if (genFn0() < 1) {
                                        var v58 uint16 = 2
                                        v59 := GenS5{F0: v1, F1: fi54, F2: v50}
                                        var v60 string = v6
                                        if (145 < 255) {
                                            var v61 uint16 = v58
                                            fmt.Println(v58, v4)
                                            v51 += (v50 * v51)
                                            switch 42 {
                                            case 32767:
                                                var v62 int64 = 9223372036854775807
                                                _ = v62
                                            case 56:
                                                if v4 {
                                                    v60 = (v6 + v60)
                                                    v63 := GenS6{F0: 65535, F1: uint64(v61), F2: 0}
                                                    _ = v63
                                                } else {
                                                    var v64 uint64 = 18446744073709551615
                                                    v2 = (0.0 + 1.5)
                                                    if len(v55) >= 2 {
                                                        v55 = v55[1:len(v55)]
                                                    }
                                                    _ = v64
                                                }
                                            case -729:
                                                v55 = append(v55, (-(fi53)))
                                            }
                                            v1 += -1
                                            _ = v61
                                        } else {
                                            if v4 {
                                                switch v60 {
                                                case "hello":
                                                    var v65 uint = v50
                                                    _ = v65
                                                default:
                                                    v60 = "a"
                                                }
                                                switch (("" + (v6 + (((("hello" + v6) + v6) + ("\n\t" + (v60 + v60))) + v6))) + v60) {
                                                case "hello":
                                                    v52 = v3
                                                case "":
                                                    var v66 bool = false
                                                    _ = v66
                                                case "\n\t":
                                                    fmt.Println(fi53, v51, fi53)
                                                }
                                                fmt.Println(v5)
                                            } else {
                                                var v67 float64 = float64(float32(v5))
                                                v55 = append(v55, 127)
                                                _ = v67
                                            }
                                            switch v6 {
                                            case "hello":
                                                v68 := GenS3{F0: 1}
                                                _ = v68
                                            }
                                            var v69 int32 = v1
                                            _ = v69
                                        }
                                        if false {
                                            var v70 uint = 546
                                            _ = v70
                                        }
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v71 int64 = 7
                                        switch (("\n\t" + ("hello" + ("" + v6))) + v6) {
                                        case "hello":
                                            v4 = false
                                        case "a":
                                            v71 -= 507
                                        case "\n\t":
                                            if (int8(v52) < -128) {
                                                v5 -= ((1e10 / v5) / (-2.25 / 0.1))
                                            } else {
                                                var v72 int8 = 1
                                                var v73 uint32 = 42
                                                v3 += (v52 & 42)
                                                _ = v72
                                                _ = v73
                                            }
                                        default:
                                            fmt.Println(v51)
                                        }
                                        _ = v71
                                    }
                                    v6 = v6
                                    _ = len(v55)
                                }
                            }
                        default:
                            v1 = -433
                        }
                        v74 := GenS7{F0: 42}
                        _ = v55
                        _ = v74
                    }
                    var v75 int8 = 109
                    v76 := GenS2{F0: 187, F1: v1, F2: (v50 >> ((v50) & 63))}
                    _ = v75
                    _ = v76
                }
                _ = fi54
            }
            var v77 int16 = (v3 / ((1) | 1))
            for fi78 := 0; fi78 < 6; fi78++ {
                switch (-(fi78)) {
                case 32768:
                    if false {
                        v51 += (v51 / ((v51) | 1))
                    } else {
                        v79 := GenS5{F0: v1, F1: fi53, F2: v51}
                        fmt.Println(v2, v2)
                        v50 = v51
                        _ = v79
                    }
                }
                if v4 {
                    v80 := []int{fi78}
                    fmt.Println(v50)
                    switch (v6 + "\n\t") {
                    case "\n\t":
                        v6 = v6
                        v81 := GenS8{F0: (v1 + v1), F1: 255}
                        var v82 uint = 255
                        _ = v81
                        _ = v82
                    case "hello":
                        if len(v80) > 0 {
                            v80[0] = (^(fi78))
                        }
                    default:
                        if len(v80) >= 2 {
                            v80 = v80[1:len(v80)]
                        }
                    }
                    var v83 uint32 = 2
                    _ = len(v80)
                    var v84 int = fi78
                    var v85 string = ("a" + "\n\t")
                    _ = v80
                    _ = v83
                    _ = v84
                    _ = v85
                }
                _ = fi78
            }
            _ = v77
        }
        _ = fi53
    }
    fmt.Println(v4, v3, v2)
    if (v4 && ((v5 > float64(v2)) && (((!v4) && ((v50 - (v51 - 255)) < v51)) && true))) {
        var v86 float64 = 1.5
        var v87 uint8 = 1
        _ = v86
        _ = v87
    }
    var v88 int = -1
    for fi89 := 0; fi89 < 10; fi89++ {
        if false {
            v88 = v88
        }
        if (42 >= 1) {
            var v90 string = (("foo bar" + v6) + "foo bar")
            _ = v90
        } else {
            var v91 uint8 = 89
            v92 := GenS9{F0: v6}
            if (0 < v91) {
                if v4 {
                    fmt.Println(v1, v3, v88)
                } else {
                    for fi93 := 0; fi93 < 5; fi93++ {
                        if (32768 > (fi89 | -9223372036854775808)) {
                            fmt.Println(v5, fi89)
                            v94 := GenS10{F0: genFn1(-100, fi89), F1: (!(!v4))}
                            var v95 int32 = (v1 & 1)
                            _ = v94
                            _ = v95
                        } else {
                            var v96 int8 = -31
                            v52 -= v52
                            v97 := []int{fi93}
                            _ = v96
                            _ = v97
                        }
                        v98 := GenS3{F0: 1}
                        _ = fi93
                        _ = v98
                    }
                    var v99 int64 = 9223372036854775807
                    _ = v99
                }
            }
            _ = v91
            _ = v92
        }
        _ = fi89
    }
    v4 = (v6 < v6)
    v100 := []int{(v88 % (((v88 - v88)) | 1)), (^(v88)), v88}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v100:", len(v100), v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v88
}
