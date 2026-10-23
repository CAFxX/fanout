package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int16 }
type GenS8 struct { F0 int; F1 string }
func genFn0(p0 int64, p1 int64, p2 int16) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 0
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 float64 = -2.25
        v10 := []int{-9223372036854775808, 32767, fi8, 4294967295}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (!(!false)) {
        var v11 int16 = -32768
        p1 = 0
        v12 := GenS1{F0: 9223372036854775807, F1: 495, F2: 579}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 string = (("a" + ("" + (("hello" + "") + ("" + ("" + "\n\t"))))) + "foo bar")
    var v2 int16 = 36
    var v3 int8 = -128
    var v4 string = v1
    var v5 uint16 = 139
    var v13 uint8 = 42
    if false {
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: fi14, F1: fi14, F2: (v5 % ((1) | 1))}
            if false {
                v16 := []int{fi14, (fi14 / (((^(fi14))) | 1))}
                if true {
                    var v17 uint32 = 42
                    var v18 bool = false
                    var v19 bool = (74 >= v13)
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v3 -= v3
                    fmt.Println(v5)
                }
                _ = v16
            }
            v20 := []int{1}
            _ = fi14
            _ = v15
            _ = v20
        }
        v21 := []int{}
        _ = v21
    } else {
        v13 -= (v13 << ((406) & 63))
        var v22 uint64 = 2
        for fi23 := 0; fi23 < 3; fi23++ {
            switch v1 {
            case "hello":
                if (v3 < (v3 * -1)) {
                    v24 := GenS2{F0: (^(v3))}
                    _ = v24
                }
            case "\n\t":
                var v25 uint = 653
                _ = v25
            case "a":
                v26 := GenS3{F0: (^(v22))}
                _ = v26
            default:
                var v27 uint32 = 520
                _ = v27
            }
            v28 := []int{(^(fi23)), fi23, 127, 32767}
            _ = fi23
            _ = v28
        }
        _ = v22
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 8; fi30++ {
            if (0 <= -935) {
                switch v1 {
                case "":
                    v31 := []int{fi29, (-(fi30)), fi30, fi30}
                    _ = v31
                case "foo bar":
                    v3 = v3
                }
                v32 := []int{}
                _ = v32
            }
            v33 := []int{}
            _ = fi30
            _ = v33
        }
        _ = fi29
    }
    v5 += (^(v5))
    v34 := GenS4{F0: 865, F1: 253}
    switch int(v3) {
    case 128:
        for fi35 := 0; fi35 < 8; fi35++ {
            var v36 bool = true
            if ((v2 >> ((0) & 63)) > v2) {
                if false {
                    var v37 int = -756
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v39 := GenS5{F0: v2, F1: 910}
                        var v40 uint16 = (v5 - v5)
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = v37
                } else {
                    v2 -= v2
                }
            }
            _ = fi35
            _ = v36
        }
    case -52:
        v2 = (^(v2))
    default:
        v41 := []int{}
        _ = v41
    }
    v42 := GenS4{F0: (v5 - (v5 & v5)), F1: 2147483647}
    if false {
        switch (v4 + v1) {
        case "foo bar":
            v43 := []int{4294967295, -9223372036854775808, -820}
            _ = v43
        case "a":
            v44 := []int{3}
            _ = v44
        case "\n\t":
            if false {
                var v45 uint8 = 255
                v46 := []int{2147483647, 0, -638, 32767}
                _ = v45
                _ = v46
            }
        default:
            v2 = int16(v5)
        }
        var v47 uint = 1
        if true {
            switch v4 {
            case "hello":
                v2 = (v2 >> (((v47 * (v47 % (((v47 | 565)) | 1)))) & 63))
            case "a":
                var v48 int64 = 1
                _ = v48
            default:
                var v49 uint16 = (v5 % ((32767) | 1))
                _ = v49
            }
            fmt.Println(v3, v1)
        }
        _ = v47
    }
    for fi50 := 0; fi50 < 15; fi50++ {
        switch "foo bar" {
        case "hello":
            if true {
                v2 = 1
                for fi51 := 0; fi51 < 12; fi51++ {
                    var v52 int16 = 175
                    var v53 bool = (!true)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                for fi54 := 0; fi54 < 14; fi54++ {
                    v5 = 255
                    v13 = (v13 & (v13 | v13))
                    _ = fi54
                }
            } else {
                if (0.1 != 0.0) {
                    var v55 float64 = 1e10
                    v56 := []int{}
                    _ = v55
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 9; fi57++ {
                        var v58 uint32 = 2147483648
                        v59 := []int{fi50, fi50, (^(fi50))}
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                    v60 := GenS6{F0: (v3 + v3)}
                    fmt.Println(v4, v13)
                    _ = v60
                }
            }
        }
        for fi61 := 0; fi61 < 3; fi61++ {
            if false {
                if (v5 < v5) {
                    if true {
                        v62 := GenS7{F0: false, F1: (v2 * 1)}
                        v1 = v1
                        _ = v62
                    } else {
                        if false {
                            var v63 uint = 2
                            if (fi61 > (fi61 >> (((v63 % ((v63) | 1))) & 63))) {
                                fmt.Println(v13)
                            }
                            _ = v63
                        }
                        v64 := []int{fi50, (fi61 & (-(fi50))), -9223372036854775808}
                        var v65 float32 = 1e10
                        _ = v64
                        _ = v65
                    }
                } else {
                    fmt.Println(fi61, fi61)
                    v66 := GenS5{F0: v2, F1: 4294967295}
                    _ = v66
                }
                if (false || true) {
                    v67 := []int{}
                    _ = v67
                } else {
                    switch "" {
                    case "":
                        var v68 uint8 = 128
                        _ = v68
                    case "hello":
                        var v69 float64 = (1e10 - ((0.1 / 1.0) / (-2.25 - (1e10 + (1.5 / 0.1)))))
                        _ = v69
                    case "a":
                        v5 += (v5 | 2)
                    default:
                        v70 := GenS8{F0: (fi61 << ((65535) & 63)), F1: v4}
                        _ = v70
                    }
                }
            } else {
                v71 := []int{}
                _ = v71
            }
            v13 += (v13 | v13)
            var v72 float64 = 0.1
            _ = fi61
            _ = v72
        }
        fmt.Println(v3, fi50)
        _ = fi50
    }
    v5 -= (v5 | (v5 - (v5 % ((0) | 1))))
    v73 := []int{425}
    v74 := GenS5{F0: -1, F1: 18446744073709551615}
    if false {
        v2 -= 42
    } else {
        for fi75 := 0; fi75 < 15; fi75++ {
            if ((v13 + (v13 ^ 0)) <= v13) {
                fmt.Println(fi75)
                fmt.Println(v4, v4, v4)
            }
            var v76 int64 = 9223372036854775807
            _ = fi75
            _ = v76
        }
        if len(v73) >= 2 {
            v73 = v73[1:len(v73)]
        }
    }
    switch ("foo bar" + ("foo bar" + ("\n\t" + ""))) {
    case "":
        v5 = (v5 >> ((18446744073709551615) & 63))
    case "a":
        if (!true) {
            switch ("foo bar" + v4) {
            case "hello":
                switch 2147483647 {
                case 32768:
                    v13 = 128
                case 9223372036854775807:
                    v1 = "\n\t"
                }
            }
        } else {
            v77 := GenS8{F0: int(v2), F1: v4}
            for fi78 := 0; fi78 < 6; fi78++ {
                v79 := GenS8{F0: (fi78 + (fi78 / ((fi78) | 1))), F1: v1}
                v1 = ((v1 + v4) + (v1 + ("foo bar" + v4)))
                for fi80 := 0; fi80 < 13; fi80++ {
                    if len(v73) >= 2 {
                        v73 = v73[1:len(v73)]
                    }
                    var v81 int32 = -1
                    _ = fi80
                    _ = v81
                }
                _ = fi78
                _ = v79
            }
            _ = v77
        }
    }
    var v82 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v73:", len(v73), v73)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v73
    _ = v74
    _ = v82
}
