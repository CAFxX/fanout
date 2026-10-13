package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int) bool {
    v5 := []int{(p0 / ((p0) | 1)), (p0 >> ((18446744073709551615) & 63)), p0, p0}
    if len(v5) > 0 {
        v5[0] = (-(p0))
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        p0 = (fi6 - p0)
        v5 = append(v5, 470)
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v5) > 0 {
            v5[0] = p0
        }
        fmt.Println(fi7, fi7)
        v5 = append(v5, (fi7 % ((42) | 1)))
        _ = fi7
    }
    p0 = (p0 * p0)
    _ = p0
    _ = v5
    return (false && true)
}

func main() {
    var v1 bool = false
    var v2 int32 = 42
    var v3 bool = (561 > 4294967295)
    var v4 int16 = -1
    var v8 int16 = v4
    v9 := GenS1{F0: 32768, F1: (v8 | (v8 >> ((759) & 63))), F2: (v2 % ((v2) | 1))}
    var v10 string = (("" + "foo bar") + (("\n\t" + "") + "\n\t"))
    var v11 int64 = 42
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float32 = 1e10
        for fi14 := 0; fi14 < 5; fi14++ {
            switch -621 {
            case 128:
                var v15 uint = 926
                _ = v15
            case -370:
                v13 -= (v13 / float32((-2.25 * (0.1 / 3.14159))))
            default:
                v3 = v3
                if v3 {
                    v1 = (18446744073709551615 > 42)
                    if (!false) {
                        v10 = v10
                        v16 := []int{(-(fi14)), (fi12 & (fi12 / (((-(fi12))) | 1))), 127, (fi12 / (((fi12 & fi12)) | 1))}
                        if (!false) {
                            v17 := GenS2{F0: (v4 % (((v8 | v4)) | 1)), F1: 117}
                            _ = v17
                        }
                        _ = v16
                    }
                    fmt.Println(v13)
                    var v18 uint8 = 127
                    _ = v18
                } else {
                    if (!(v4 > v4)) {
                        v19 := []int{fi12}
                        if v3 {
                            if (!(1 > 32767)) {
                                if len(v19) > 0 {
                                    v19[0] = (fi12 | fi12)
                                }
                                if len(v19) > 0 {
                                    v19[0] = fi14
                                }
                            } else {
                                v4 = v8
                                if (v10 < v10) {
                                    v20 := GenS1{F0: fi14, F1: (^(v4)), F2: (v2 & (v2 * v2))}
                                    v1 = ((545 < 255) || (!("hello" == ("" + "\n\t"))))
                                    _ = v20
                                } else {
                                    v4 = v4
                                    v10 = "foo bar"
                                    v10 = ("foo bar" + "\n\t")
                                }
                                var v21 string = (v10 + v10)
                                _ = v21
                            }
                            var v22 int32 = v2
                            var v23 float32 = ((1.5 * ((v13 / v13) * 1e10)) / v13)
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS2{F0: v8, F1: 255}
                        fmt.Println(v4, v10, v1)
                        _ = v19
                        _ = v24
                    } else {
                        if v3 {
                            v25 := []int{42}
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                            _ = v25
                        }
                        var v26 uint = 42
                        _ = v26
                    }
                    v27 := []int{}
                    _ = v27
                }
                if v1 {
                    var v28 int16 = (v4 & 586)
                    v28 = (v28 * v28)
                    if false {
                        var v29 int = fi12
                        v10 = "a"
                        _ = v29
                    }
                    switch (fi14 | (fi12 + (fi12 + 42))) {
                    case 0:
                        v1 = (!(2147483648 <= 0))
                    default:
                        var v30 uint64 = 1
                        var v31 bool = ((-2.25 - -2.25) <= ((0.0 + (float64((v13 + (v13 - 0.0))) / ((3.14159 / 1e10) / ((-2.25 / -2.25) * (1e10 * 1e10))))) - (0.0 / 3.14159)))
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS3{F0: (^(v11)), F1: (^(fi12))}
                    _ = v28
                    _ = v32
                }
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    switch 280 {
    case -2147483648:
        switch (v10 + v10) {
        case "foo bar":
            for fi33 := 0; fi33 < 1; fi33++ {
                var v34 float64 = -2.25
                switch 1 {
                case 1:
                    v35 := []int{fi33}
                    _ = v35
                case 890:
                    for fi36 := 0; fi36 < 9; fi36++ {
                        var v37 float64 = (v34 / v34)
                        v2 = (v2 + (v2 + 1))
                        if (v8 != 0) {
                            if (982 == 42) {
                                switch "a" {
                                case "hello":
                                    var v38 int = (fi33 % ((1) | 1))
                                    if (!v3) {
                                        fmt.Println(fi36)
                                    }
                                    var v39 bool = v1
                                    _ = v38
                                    _ = v39
                                case "\n\t":
                                    if v3 {
                                        if (v34 == ((v34 + ((v37 * v34) * v37)) / v34)) {
                                            var v40 int32 = v2
                                            var v41 string = v10
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if v3 {
                                                v42 := GenS3{F0: (v11 & -1), F1: -2147483648}
                                                v43 := GenS2{F0: v4, F1: 766}
                                                _ = v42
                                                _ = v43
                                            }
                                            if (0 == 807) {
                                                v44 := GenS4{F0: (v1 || v3)}
                                                switch (fi33 & (fi36 << ((0) & 63))) {
                                                case 0:
                                                    v37 -= (1e10 / v34)
                                                case 42:
                                                    v8 = -32768
                                                case -9223372036854775808:
                                                    if (v3 || (!false)) {
                                                        if v3 {
                                                            switch ("\n\t" + ((v10 + (v10 + "foo bar")) + "\n\t")) {
                                                            case "a":
                                                                switch (v10 + v10) {
                                                                case "foo bar":
                                                                    fmt.Println(v2)
                                                                case "\n\t":
                                                                    var v45 int32 = v2
                                                                    _ = v45
                                                                case "hello":
                                                                    if (0.1 != 1e10) {
                                                                        fmt.Println(fi33)
                                                                        v11 = 159
                                                                    }
                                                                default:
                                                                    v46 := []int{fi36, fi36}
                                                                    _ = v46
                                                                }
                                                            default:
                                                                switch 1 {
                                                                case 9223372036854775807:
                                                                    v47 := []int{9223372036854775807}
                                                                    _ = v47
                                                                case 568:
                                                                    v3 = v1
                                                                case -2147483648:
                                                                    switch "" {
                                                                    case "hello":
                                                                        v48 := []int{9223372036854775807, fi33, (fi36 & (fi36 >> ((863) & 63))), (fi33 * fi33)}
                                                                        _ = v48
                                                                    case "foo bar":
                                                                        v49 := []int{}
                                                                        _ = v49
                                                                    default:
                                                                        v4 -= (v4 / (((v8 | 163)) | 1))
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            var v50 int16 = 0
                                                            v4 -= 1
                                                            _ = v50
                                                        }
                                                        v51 := GenS5{F0: (fi36 << ((42) & 63)), F1: v1}
                                                        v52 := GenS3{F0: v11, F1: fi33}
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                }
                                                v53 := []int{fi36, fi33, int(v2)}
                                                _ = v44
                                                _ = v53
                                            } else {
                                                if v1 {
                                                    v3 = (float32(v37) >= 3.14159)
                                                    v54 := GenS6{F0: 407, F1: v4}
                                                    if (!v3) {
                                                        var v55 int = fi36
                                                        _ = v55
                                                    } else {
                                                        if (!v3) {
                                                            v56 := GenS7{F0: 18446744073709551615, F1: (0.1 < -2.25)}
                                                            v34 += v37
                                                            _ = v56
                                                        }
                                                        var v57 uint8 = 55
                                                        _ = v57
                                                    }
                                                    _ = v54
                                                }
                                                var v58 uint32 = 921
                                                _ = v58
                                            }
                                            if (v11 <= v11) {
                                                v59 := GenS8{F0: 127}
                                                if (268 <= 356) {
                                                    if v3 {
                                                        var v60 int64 = v11
                                                        var v61 uint16 = 1
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                } else {
                                                    if (18446744073709551615 >= 255) {
                                                        if v3 {
                                                            v62 := []int{fi36, fi36, -183, (fi36 & 4294967295)}
                                                            _ = v62
                                                        } else {
                                                            v8 = (-(v8))
                                                            fmt.Println(v1)
                                                        }
                                                        switch (fi36 * 32768) {
                                                        case -1:
                                                            var v63 uint = 2
                                                            _ = v63
                                                        case -435:
                                                            v64 := []int{(fi33 | (fi36 ^ 250)), fi33, (fi33 ^ fi36)}
                                                            _ = v64
                                                        }
                                                        v65 := GenS7{F0: 0, F1: (!(255 != 255))}
                                                        _ = v65
                                                    }
                                                    switch fi33 {
                                                    case 32768:
                                                        if (v1 && (false || (391 < (-(v2))))) {
                                                            var v66 uint = 619
                                                            _ = v66
                                                        } else {
                                                            fmt.Println(v34, v1)
                                                            fmt.Println(v11)
                                                        }
                                                    case 1:
                                                        v8 = 32767
                                                    }
                                                    var v67 int32 = (v2 | (v2 << ((2) & 63)))
                                                    _ = v67
                                                }
                                                _ = v59
                                            } else {
                                                fmt.Println(v10, v11)
                                            }
                                            v68 := GenS1{F0: (fi33 / ((fi33) | 1)), F1: (v4 % (((v8 + v4)) | 1)), F2: v2}
                                            _ = v68
                                        }
                                        v69 := []int{(fi33 / (((fi36 << ((280) & 63))) | 1)), fi36, 4294967295}
                                        _ = v69
                                    }
                                }
                                if (473 >= 56) {
                                    if ((v2 - -2147483648) != (^(v2))) {
                                        if v1 {
                                            fmt.Println(v2, v2, v3)
                                        } else {
                                            v70 := GenS2{F0: v8, F1: 481}
                                            v37 += ((((v34 + float64(float32(((v37 - v37) + v34)))) + v37) + 3.14159) + (float64((0.1 + 1e10)) * float64(0.0)))
                                            v4 += 32767
                                            _ = v70
                                        }
                                    } else {
                                        if v1 {
                                            var v71 string = (v10 + "a")
                                            v72 := GenS1{F0: 1, F1: (v8 * v8), F2: v2}
                                            if (0 >= 127) {
                                                var v73 int16 = (v8 << ((1) & 63))
                                                _ = v73
                                            } else {
                                                v1 = true
                                                var v74 int32 = -2147483648
                                                if true {
                                                    var v75 float64 = ((1e10 / -2.25) * 1.5)
                                                    _ = v75
                                                }
                                                _ = v74
                                            }
                                            var v76 uint32 = 2
                                            _ = v71
                                            _ = v72
                                            _ = v76
                                        } else {
                                            var v77 float32 = 0.1
                                            v78 := []int{}
                                            var v79 uint8 = 117
                                            _ = v77
                                            _ = v78
                                            _ = v79
                                        }
                                        if (v2 == v2) {
                                            var v80 uint32 = 2147483647
                                            fmt.Println(fi36, v2, v37)
                                            _ = v80
                                        }
                                        v81 := []int{fi36, 9223372036854775807, (fi36 | fi36)}
                                        _ = v81
                                    }
                                    var v82 int8 = -57
                                    _ = v82
                                } else {
                                    if ((!(!false)) || (367 > 255)) {
                                        switch "hello" {
                                        case "a":
                                            var v83 bool = (18446744073709551615 != 250)
                                            _ = v83
                                        }
                                        var v84 int16 = (v4 % (((v4 * (v4 / (((-(v4))) | 1)))) | 1))
                                        var v85 float64 = v37
                                        _ = v84
                                        _ = v85
                                    }
                                    var v86 uint8 = 42
                                    v1 = v3
                                    _ = v86
                                }
                            }
                        } else {
                            var v87 uint32 = 4294967295
                            _ = v87
                        }
                        _ = fi36
                        _ = v37
                    }
                case 2147483647:
                    v11 = v11
                }
                fmt.Println(v1)
                _ = fi33
                _ = v34
            }
        }
    case 2147483647:
        var v88 uint64 = 42
        _ = v88
    case -889:
        var v89 int16 = (v8 ^ 809)
        _ = v89
    default:
        v90 := GenS8{F0: 247}
        _ = v90
    }
    v91 := GenS8{F0: 60}
    switch (("" + (v10 + v10)) + (v10 + v10)) {
    case "foo bar":
        fmt.Println(v8)
    case "hello":
        v4 = v8
    case "a":
        for fi92 := 0; fi92 < 9; fi92++ {
            v3 = (1 >= 255)
            _ = fi92
        }
    default:
        for fi93 := 0; fi93 < 16; fi93++ {
            if true {
                if true {
                    var v94 int8 = 127
                    _ = v94
                }
                if v3 {
                    v95 := GenS9{F0: 707}
                    v96 := GenS3{F0: v11, F1: 128}
                    _ = v95
                    _ = v96
                }
                var v97 uint = 266
                _ = v97
            }
            v8 += v8
            _ = fi93
        }
    }
    if v1 {
        switch 773 {
        case 9223372036854775807:
            var v98 int = 2147483647
            _ = v98
        default:
            if ((v11 + v11) == 42) {
                fmt.Println(v11, v10)
            } else {
                v99 := GenS5{F0: 127, F1: (!((812 <= 499) && v1))}
                var v100 int16 = 866
                _ = v100
                _ = v99
            }
        }
        fmt.Println(v4, v8)
    }
    v101 := []int{29, 9223372036854775807}
    if (v3 && v1) {
        if len(v101) > 0 {
            v101[0] = 128
        }
    } else {
        var v102 string = v10
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v101:", len(v101), v101)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v101
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
    _ = v91
}
