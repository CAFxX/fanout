package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int8; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int, p1 int8, p2 uint) uint16 {
    v6 := GenS1{F0: 999, F1: (p2 % ((p2) | 1))}
    var v7 uint16 = 32767
    if false {
        p2 = p2
    }
    var v8 int64 = -414
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 770
    var v2 int32 = 1
    var v3 int64 = 1
    var v4 int16 = 0
    var v5 float32 = 0.0
    if (v5 <= ((-2.25 / ((v5 + (v5 / v5)) * v5)) + ((float32(1.5) + v5) + v5))) {
        var v9 uint16 = 65535
        _ = v9
    } else {
        v5 += 1.5
        var v10 uint16 = genFn0(int(v4), -128, 820)
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 int16 = 32767
            v13 := []int{(fi11 & 1)}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    if true {
        v14 := GenS1{F0: 584, F1: 2}
        v4 -= int16(v3)
        _ = v14
    } else {
        var v15 int = -2147483648
        fmt.Println(v2, v4)
        v2 += 0
        _ = v15
    }
    var v16 float32 = v5
    v4 = v4
    v17 := []int{}
    if (v1 > (v1 + v1)) {
        v16 = (v5 - 3.14159)
        if true {
            v16 -= v5
            if true {
                v16 -= v16
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS2{F0: "foo bar", F1: 634, F2: 2147483648}
                    if (("a" + ("a" + "")) <= "foo bar") {
                        _ = len(v17)
                        v20 := GenS3{F0: "hello", F1: (^(fi18))}
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
            }
            if true {
                v21 := GenS2{F0: "", F1: 495, F2: 0}
                var v22 int32 = (v2 * v2)
                var v23 uint32 = 65535
                _ = v21
                _ = v22
                _ = v23
            } else {
                var v24 int = -2147483648
                _ = v24
            }
        }
        var v25 int16 = 1000
        _ = v25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS4{F0: fi26, F1: 65535}
        switch fi26 {
        case 4294967295:
            _ = len(v17)
        case -971:
            if (((3.14159 - 0.1) + (3.14159 - 0.1)) == 3.14159) {
                switch 42 {
                case 127:
                    var v28 float32 = v16
                    _ = v28
                case 996:
                    v29 := GenS2{F0: "", F1: (^(v1)), F2: 42}
                    _ = v29
                case -9223372036854775808:
                    for fi30 := 0; fi30 < 8; fi30++ {
                        switch ("foo bar" + "") {
                        case "hello":
                            if true {
                                var v31 string = "\n\t"
                                _ = v31
                            }
                            var v32 uint = 255
                            var v33 int64 = 539
                            _ = v32
                            _ = v33
                        default:
                            if true {
                                switch 1 {
                                case -931:
                                    var v34 uint = 255
                                    _ = v34
                                case 32767:
                                    var v35 uint8 = 2
                                    _ = v35
                                case 32768:
                                    v36 := GenS5{F0: (v1 % ((0) | 1)), F1: -56, F2: 9223372036854775807}
                                    _ = v36
                                default:
                                    v5 += (1e10 / v5)
                                }
                                if (!(true && (42 >= 42))) {
                                    v37 := GenS6{F0: 201, F1: 65535}
                                    var v38 int64 = 9223372036854775807
                                    var v39 uint16 = 42
                                    switch (fi26 - (fi26 | fi30)) {
                                    case -2147483648:
                                        var v40 uint16 = v39
                                        _ = v40
                                    }
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                                var v41 string = (("foo bar" + "\n\t") + ((("foo bar" + (("" + "\n\t") + "foo bar")) + "hello") + ""))
                                _ = v41
                            }
                        }
                        v42 := GenS2{F0: "", F1: 42, F2: 65535}
                        _ = fi30
                        _ = v42
                    }
                }
                v16 = v5
            }
        case 1:
            for fi43 := 0; fi43 < 14; fi43++ {
                var v44 uint8 = 0
                var v45 bool = true
                v5 -= v16
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        v1 -= (v1 << ((1) & 63))
        _ = fi26
        _ = v27
    }
    switch "foo bar" {
    case "a":
        var v46 uint8 = 175
        _ = v46
    case "foo bar":
        v5 += v5
    case "\n\t":
        for fi47 := 0; fi47 < 14; fi47++ {
            v4 -= v4
            if (4294967295 != 18446744073709551615) {
                for fi48 := 0; fi48 < 7; fi48++ {
                    v49 := GenS5{F0: 358, F1: 25, F2: v3}
                    v2 -= v2
                    v2 -= v2
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi48
                    _ = v49
                }
            } else {
                v2 = v2
                for fi50 := 0; fi50 < 10; fi50++ {
                    switch "hello" {
                    case "foo bar":
                        var v51 int64 = (^(v3))
                        _ = v51
                    default:
                        v4 = (v4 * (v4 / (((v4 * (v4 >> ((388) & 63)))) | 1)))
                    }
                    var v52 int8 = -128
                    _ = fi50
                    _ = v52
                }
            }
            if false {
                var v53 uint16 = 0
                switch "hello" {
                case "a":
                    v54 := GenS5{F0: 42, F1: 100, F2: int64(v4)}
                    _ = v54
                case "foo bar":
                    v55 := GenS3{F0: ("" + "foo bar"), F1: fi47}
                    _ = v55
                default:
                    if len(v17) > 0 {
                        v17[0] = 0
                    }
                }
                for fi56 := 0; fi56 < 11; fi56++ {
                    if true {
                        if (127 < 255) {
                            v57 := GenS4{F0: fi56, F1: 42}
                            var v58 string = ("hello" + "foo bar")
                            fmt.Println(v2, v58)
                            if (!false) {
                                if (2 != v53) {
                                    var v59 string = ""
                                    v53 += v53
                                    v1 -= 787
                                    var v60 int32 = v2
                                    _ = v59
                                    _ = v60
                                } else {
                                    v17 = append(v17, (^(fi47)))
                                    var v61 int64 = v3
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    _ = len(v17)
                                    _ = v61
                                }
                                v5 -= 0.0
                                var v62 uint = 65535
                                _ = v62
                            }
                            _ = v57
                            _ = v58
                        } else {
                            var v63 int = (^(fi56))
                            v53 = genFn0((fi56 & (v63 % ((fi56) | 1))), -128, 282)
                            _ = v63
                        }
                        var v64 uint64 = 979
                        v65 := GenS2{F0: "foo bar", F1: (^(v64)), F2: 213}
                        _ = v64
                        _ = v65
                    } else {
                        switch 0 {
                        case 128:
                            v53 += (v53 - v53)
                        case 4294967295:
                            v66 := GenS3{F0: "a", F1: 32767}
                            _ = v66
                        case 127:
                            var v67 uint = 2
                            _ = v67
                        default:
                            switch ("\n\t" + "foo bar") {
                            case "hello":
                                if true {
                                    if (!false) {
                                        v5 += (0.0 - 3.14159)
                                        var v68 string = ""
                                        _ = v68
                                    }
                                    switch -2147483648 {
                                    case 32768:
                                        switch "a" {
                                        case "a":
                                            switch "hello" {
                                            case "hello":
                                                switch (fi56 * -631) {
                                                case 4294967295:
                                                    if len(v17) > 0 {
                                                        v17[0] = fi56
                                                    }
                                                }
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    var v69 int64 = 654
                                                    _ = v69
                                                case "\n\t":
                                                    switch ("" + "hello") {
                                                    case "":
                                                        var v70 int8 = int8(v1)
                                                        _ = v70
                                                    case "hello":
                                                        v2 += v2
                                                    case "foo bar":
                                                        v53 += 0
                                                    }
                                                case "foo bar":
                                                    var v71 uint8 = 127
                                                    _ = v71
                                                default:
                                                    var v72 uint32 = 79
                                                    _ = v72
                                                }
                                            }
                                        case "foo bar":
                                            fmt.Println(v1)
                                        case "":
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                        default:
                                            var v73 uint16 = 65535
                                            _ = v73
                                        }
                                    }
                                    if ((!false) && (v4 > v4)) {
                                        var v74 float32 = v5
                                        _ = v74
                                    }
                                    if len(v17) > 0 {
                                        v17[0] = 32768
                                    }
                                    if ((v2 & -2147483648) > v2) {
                                        v75 := GenS7{F0: v1}
                                        _ = v75
                                    } else {
                                        if false {
                                            var v76 string = ((("a" + "\n\t") + "hello") + "")
                                            _ = v76
                                        } else {
                                            v1 = v1
                                            var v77 float32 = -2.25
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            _ = v77
                                        }
                                        switch (fi47 ^ (fi56 + (fi47 % ((501) | 1)))) {
                                        case 32768:
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                        case -430:
                                            if true {
                                                if len(v17) > 0 {
                                                    v17[0] = fi47
                                                }
                                            }
                                        default:
                                            var v78 float64 = (-2.25 - -2.25)
                                            _ = v78
                                        }
                                        var v79 int16 = v4
                                        var v80 uint32 = 42
                                        var v81 uint = 842
                                        var v82 int64 = (v3 - (^(v3)))
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                }
                            case "":
                                if (!false) {
                                    if false {
                                        var v83 int64 = v3
                                        var v84 int8 = 42
                                        v85 := GenS1{F0: 65535, F1: 18446744073709551615}
                                        if false {
                                            v4 -= (v4 >> ((61) & 63))
                                            var v86 float32 = v5
                                            var v87 string = "hello"
                                            _ = v86
                                            _ = v87
                                        } else {
                                            switch 128 {
                                            case 9:
                                                if true {
                                                    var v88 float64 = (1.5 * (1e10 + 0.1))
                                                    var v89 uint32 = 443
                                                    v89 = v89
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    v53 = 255
                                                    var v90 float32 = (float32(3.14159) - (v5 - v5))
                                                    _ = v90
                                                }
                                            default:
                                                var v91 uint16 = (v53 >> ((0) & 63))
                                                _ = v91
                                            }
                                        }
                                        _ = len(v17)
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    v92 := GenS7{F0: (v1 ^ (v1 >> ((18446744073709551615) & 63)))}
                                    _ = v92
                                }
                            case "a":
                                if false {
                                    v93 := GenS3{F0: "hello", F1: (fi47 / (((fi47 & fi47)) | 1))}
                                    var v94 int32 = 42
                                    var v95 uint32 = 4294967295
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                } else {
                                    if false {
                                        var v96 uint64 = v1
                                        _ = v96
                                    } else {
                                        v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
                                        v4 = (-(v4))
                                        v53 = v53
                                    }
                                    fmt.Println(v5, v16, v2)
                                    v1 += 1
                                }
                            }
                        }
                        if (((v3 | v3) == 687) && (1.5 > 3.14159)) {
                            _ = len(v17)
                        }
                    }
                    _ = fi56
                }
                _ = v53
            }
            _ = fi47
        }
    }
    fmt.Println(v3, v16, v1)
    _ = len(v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
