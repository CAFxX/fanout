package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int16; F1 uint; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 int }
func genFn0(p0 int64, p1 uint8) uint {
    if (!false) {
        var v6 uint32 = 2147483648
        _ = v6
    } else {
        v7 := []int{-990, 1}
        _ = v7
    }
    v8 := []int{128}
    var v9 uint8 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 255
}
func genFn1() uint16 {
    v10 := []int{-1}
    for fi11 := 0; fi11 < 19; fi11++ {
        v10 = append(v10, 4294967295)
        _ = fi11
    }
    var v12 uint = genFn0(1, 255)
    _ = v10
    _ = v12
    return 32767
}

func main() {
    var v1 int8 = 1
    var v2 float64 = 1e10
    v3 := []string{"hello", "", "hello"}
    var v4 bool = false
    var v5 uint16 = 65535
    for fi13 := 0; fi13 < 12; fi13++ {
        switch "" {
        case "foo bar":
            var v14 bool = (2 >= 0)
            _ = v14
        case "":
            v15 := GenS1{F0: 1, F1: 42, F2: true}
            _ = v15
        default:
            v16 := GenS2{F0: 0, F1: 18446744073709551615, F2: 42}
            _ = v16
        }
        if (-1 != int16(fi13)) {
            var v17 int64 = -49
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint32 = 0
                var v20 float32 = 3.14159
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
        fmt.Println(fi13)
        _ = fi13
    }
    switch 128 {
    case 182:
        if v4 {
            var v21 int64 = 607
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 12; fi23++ {
                    if (((v1 ^ v1) > (v1 % (((v1 | v1)) | 1))) && v4) {
                        v24 := GenS3{F0: -599}
                        var v25 int = (fi23 & -1)
                        var v26 int64 = -889
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi23
                }
                _ = fi22
            }
            _ = v21
        }
    case 2147483647:
        if (v1 < v1) {
            _ = len(v3)
            for fi27 := 0; fi27 < 17; fi27++ {
                v2 -= v2
                switch "" {
                case "a":
                    v4 = false
                case "hello":
                    var v28 int64 = -1
                    _ = v28
                }
                _ = fi27
            }
            v5 = v5
        }
    case -17:
        var v29 bool = v4
        _ = v29
    default:
        v3 = append(v3, ("" + ("hello" + "")))
    }
    fmt.Println(v2, v4)
    v30 := GenS2{F0: 451, F1: 581, F2: 84}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2, v4)
    v4 = false
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 float64 = v2
            fmt.Println(v4)
            v1 -= (v1 % (((v1 ^ v1)) | 1))
            switch "a" {
            case "foo bar":
                if false {
                    fmt.Println(v33)
                }
            case "a":
                v3 = append(v3, "")
            }
            fmt.Println(v1)
            fmt.Println(fi31, v4, v1)
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    fmt.Println(v2, v2, v1)
    var v34 int32 = 805
    if (!(!((-2.25 + ((float32((((v2 + (((float64(1.5) / v2) + 0.1) + (v2 / v2))) - v2) - 0.1)) * 0.1) + 1.5)) > 0.0))) {
        var v35 int8 = 1
        if (-32768 > -32768) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                for fi36 := 0; fi36 < 6; fi36++ {
                    var v37 uint64 = 1
                    for fi38 := 0; fi38 < 16; fi38++ {
                        fmt.Println(fi38)
                        v5 = v5
                        _ = fi38
                    }
                    _ = fi36
                    _ = v37
                }
                if true {
                    if (float64(1e10) != v2) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v39 := GenS1{F0: 2, F1: 793, F2: (649 == 905)}
                        _ = v39
                    } else {
                        for fi40 := 0; fi40 < 17; fi40++ {
                            v41 := GenS3{F0: -32768}
                            _ = fi40
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 2; fi42++ {
                            v43 := GenS2{F0: -1, F1: 255, F2: 1}
                            _ = fi42
                            _ = v43
                        }
                        fmt.Println(v4, v2)
                    }
                    for fi44 := 0; fi44 < 14; fi44++ {
                        v5 = v5
                        v2 = ((0.1 / (v2 + (v2 * (-2.25 - v2)))) - (v2 * (v2 - ((1e10 - float64(-2.25)) - ((0.0 - (-2.25 + (v2 + 3.14159))) * v2)))))
                        _ = fi44
                    }
                    if v4 {
                        var v45 int = 4294967295
                        for fi46 := 0; fi46 < 5; fi46++ {
                            v2 += (1e10 - (v2 / v2))
                            _ = fi46
                        }
                        _ = v45
                    }
                }
                var v47 float64 = v2
                _ = v47
            }
        }
        switch "foo bar" {
        case "":
            switch -2147483648 {
            case 128:
                if (255 != uint8(v5)) {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    for fi48 := 0; fi48 < 6; fi48++ {
                        switch "" {
                        case "a":
                            fmt.Println(v1)
                        default:
                            v35 = 105
                        }
                        var v49 uint = 0
                        var v50 float32 = 0.1
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    var v51 float32 = 3.14159
                    _ = v51
                }
            case 1:
                for fi52 := 0; fi52 < 17; fi52++ {
                    var v53 int16 = 726
                    switch fi52 {
                    case 981:
                        var v54 string = ("hello" + "hello")
                        _ = v54
                    case -9223372036854775808:
                        var v55 int16 = v53
                        _ = v55
                    case 32767:
                        v34 += (v34 - (v34 * v34))
                    default:
                        if false {
                            if (v4 && (-647 < (fi52 & (fi52 + (fi52 | fi52))))) {
                                switch "hello" {
                                case "foo bar":
                                    var v56 bool = (!v4)
                                    _ = v56
                                default:
                                    v1 = (v35 ^ v35)
                                }
                                var v57 int8 = (v1 << ((1) & 63))
                                var v58 uint32 = 4294967295
                                _ = v57
                                _ = v58
                            }
                            var v59 string = "foo bar"
                            var v60 int64 = -9223372036854775808
                            _ = v59
                            _ = v60
                        } else {
                            var v61 string = "hello"
                            v62 := GenS4{F0: -332, F1: 1, F2: -140}
                            var v63 int32 = (v34 * (v34 + (v34 / ((42) | 1))))
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                    }
                    if v4 {
                        v34 = 1
                        if (fi52 >= fi52) {
                            v2 += v2
                            var v64 float32 = -2.25
                            v1 = 127
                            _ = v64
                        } else {
                            if (80 < v35) {
                                v65 := GenS5{F0: 2, F1: 0, F2: 239}
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v4)
                                case "hello":
                                    for fi66 := 0; fi66 < 16; fi66++ {
                                        if v4 {
                                            switch ("a" + "") {
                                            case "foo bar":
                                                fmt.Println(fi52, fi66)
                                            case "a":
                                                var v67 uint16 = (^(v5))
                                                _ = v67
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                            v2 -= v2
                                            v68 := GenS6{F0: (v34 * (v34 + (v34 - (v34 & v34))))}
                                            _ = v68
                                        }
                                        _ = fi66
                                    }
                                default:
                                    var v69 float32 = -2.25
                                    _ = v69
                                }
                                _ = v65
                            } else {
                                if false {
                                    var v70 int = 938
                                    _ = v70
                                }
                                v71 := GenS7{F0: 18446744073709551615, F1: int(v35)}
                                _ = v71
                            }
                        }
                    } else {
                        fmt.Println(fi52, v53, v1)
                        var v72 float32 = (1e10 + (0.0 - 1e10))
                        fmt.Println(v72, v2)
                        _ = v72
                    }
                    _ = fi52
                    _ = v53
                }
            }
        default:
            for fi73 := 0; fi73 < 8; fi73++ {
                var v74 bool = v4
                v74 = (0.0 == (v2 * 0.0))
                var v75 int = 700
                _ = fi73
                _ = v74
                _ = v75
            }
        }
        _ = v35
    } else {
        fmt.Println(v5, v4)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch -1 {
    case -212:
        var v76 bool = ((v2 * v2) <= 0.1)
        _ = v76
    case 911:
        switch 1 {
        case 2147483647:
            if ((0 >= 9) || (!(118 < (^(v34))))) {
                var v77 uint8 = 255
                v77 = v77
                _ = v77
            }
        case -619:
            var v78 int = 42
            _ = v78
        case 1:
            for fi79 := 0; fi79 < 2; fi79++ {
                var v80 uint64 = 2
                _ = fi79
                _ = v80
            }
        }
    case 127:
        if len(v3) > 0 {
            v3[0] = ("a" + "foo bar")
        }
    }
    v4 = ((((3.14159 - float32(3.14159)) / -2.25) / (float32(1.5) + 1e10)) <= (((1e10 + (1e10 / ((-2.25 * ((-2.25 / 1e10) + float32(v2))) - 0.0))) + (((1.5 / 1.0) - float32((v2 / v2))) - 0.0)) / 1e10))
    v81 := GenS4{F0: -9223372036854775808, F1: 0, F2: 259}
    v82 := GenS2{F0: 827, F1: uint(v5), F2: -629}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v81
    _ = v82
}
