package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 string; F2 bool }
type GenS5 struct { F0 uint; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint64; F1 int8; F2 int16 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 5; fi5++ {
            var v6 int64 = 9223372036854775807
            var v7 uint32 = 913
            _ = fi5
            _ = v6
            _ = v7
        }
        _ = fi4
    }
    var v8 int8 = 127
    _ = v8
    return 2147483648
}
func genFn1(p0 uint32) uint64 {
    if false {
        v9 := GenS1{F0: 18446744073709551615, F1: 1}
        v10 := []int{}
        _ = v10
        _ = v9
    } else {
        p0 = genFn0()
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(p0, p0)
                if (!(65535 >= 18446744073709551615)) {
                    v12 := []int{246, 32767, 684, fi11}
                    p0 = p0
                    _ = v12
                } else {
                    p0 = 2147483648
                    var v13 uint8 = 0
                    if (!false) {
                        v14 := []int{int(p0), 4294967295, (fi11 * 42)}
                        _ = v14
                    } else {
                        v15 := GenS1{F0: 1, F1: -2147483648}
                        if (1 == -704) {
                            fmt.Println(fi11)
                            p0 += p0
                            p0 -= p0
                        }
                        var v16 int = (fi11 << ((487) & 63))
                        _ = v15
                        _ = v16
                    }
                    _ = v13
                }
                if true {
                    if (2147483647 != -1) {
                        if false {
                            v17 := []int{(fi11 | fi11), fi11}
                            _ = v17
                        } else {
                            v18 := GenS1{F0: 351, F1: 1}
                            if (true && false) {
                                v19 := []int{fi11, fi11}
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 14; fi20++ {
                                v21 := GenS1{F0: 0, F1: -834}
                                _ = fi20
                                _ = v21
                            }
                            _ = v18
                        }
                        p0 = (p0 + genFn0())
                    } else {
                        if (true && (-508 > -2147483648)) {
                            v22 := GenS2{F0: fi11, F1: 57}
                            _ = v22
                        }
                        p0 += (p0 >> ((65535) & 63))
                    }
                    p0 -= 173
                    var v23 uint16 = 65535
                    _ = v23
                } else {
                    var v24 bool = (p0 != (p0 << ((255) & 63)))
                    _ = v24
                }
                _ = fi11
            }
        case "\n\t":
            v25 := GenS1{F0: 42, F1: -894}
            _ = v25
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(p0)
            _ = fi26
        }
    }
    p0 = 2147483648
    _ = p0
    return uint64(p0)
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[uint32]string)
    var v3 int16 = 0
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 10; fi28++ {
            var v29 int8 = -86
            _ = fi28
            _ = v29
        }
        var v30 int64 = 9223372036854775807
        _ = fi27
        _ = v30
    }
    fmt.Println(v1, v3)
    var v31 int32 = 931
    v2[857] = ("foo bar" + "\n\t")
    for fi32 := 0; fi32 < 11; fi32++ {
        if (-2.25 >= 1.5) {
            if (!true) {
                for fi33 := 0; fi33 < 11; fi33++ {
                    var v34 uint64 = 18446744073709551615
                    v31 += 2147483647
                    v34 += 876
                    var v35 int8 = 100
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 float64 = 0.1
                    v3 = v3
                    _ = fi36
                    _ = v37
                }
                for fi38 := 0; fi38 < 15; fi38++ {
                    var v39 float64 = 3.14159
                    v40 := GenS2{F0: fi38, F1: -128}
                    var v41 int64 = 947
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
            }
            v31 = v31
        } else {
            var v42 float32 = 0.1
            for fi43 := 0; fi43 < 13; fi43++ {
                switch "a" {
                case "hello":
                    var v44 int8 = 127
                    _ = v44
                default:
                    var v45 float64 = 0.0
                    _ = v45
                }
                v46 := GenS3{F0: fi43}
                var v47 float64 = ((3.14159 - float64((((float32(3.14159) / v42) * (v42 - v42)) - (1e10 / 1.0)))) - (3.14159 + -2.25))
                _ = fi43
                _ = v46
                _ = v47
            }
            v42 = (0.0 + 3.14159)
            _ = v42
        }
        var v48 int16 = 0
        for fi49 := 0; fi49 < 6; fi49++ {
            var v50 int64 = -1
            v3 -= -1
            v31 = v31
            _ = fi49
            _ = v50
        }
        _ = fi32
        _ = v48
    }
    for fi51 := 0; fi51 < 5; fi51++ {
        v31 = v31
        var v52 int32 = (v31 | (v31 - (v31 & v31)))
        v3 = (v3 * 0)
        _ = fi51
        _ = v52
    }
    v53 := GenS1{F0: 4294967295, F1: 0}
    var v54 uint16 = (v1 / ((v1) | 1))
    _ = len(v2)
    fmt.Println(v3, v54, v31)
    if false {
        switch "\n\t" {
        case "a":
            if false {
                v54 = v54
            } else {
                switch 0 {
                case 32767:
                    if true {
                        v55 := GenS4{F0: -9223372036854775808, F1: "a", F2: true}
                        for fi56 := 0; fi56 < 9; fi56++ {
                            v3 += (v3 + v3)
                            fmt.Println(v31)
                            _ = fi56
                        }
                        if false {
                            v3 = (v3 ^ (^(v3)))
                            var v57 int8 = 100
                            _ = v57
                        } else {
                            for fi58 := 0; fi58 < 17; fi58++ {
                                var v59 int32 = (v31 >> ((0) & 63))
                                _ = fi58
                                _ = v59
                            }
                            var v60 int = 4294967295
                            if (true && false) {
                                var v61 uint = 2
                                for fi62 := 0; fi62 < 3; fi62++ {
                                    for fi63 := 0; fi63 < 3; fi63++ {
                                        var v64 float32 = (((1e10 / 1.5) * 0.1) + 0.0)
                                        v65 := []int{fi63, int(v31)}
                                        var v66 int = -267
                                        _ = fi63
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    _ = fi62
                                }
                                _ = v61
                            } else {
                                if ((v31 - int32(v60)) != v31) {
                                    if true {
                                        v67 := GenS5{F0: 255, F1: 42, F2: -949}
                                        if true {
                                            for fi68 := 0; fi68 < 20; fi68++ {
                                                v69 := []int{(-(fi68)), -2147483648, (v60 & fi68)}
                                                fmt.Println(v1)
                                                _ = fi68
                                                _ = v69
                                            }
                                            var v70 uint32 = genFn0()
                                            fmt.Println(v60, v3, v31)
                                            _ = v70
                                        }
                                        _ = v67
                                    }
                                    v71 := GenS1{F0: 0, F1: 2147483647}
                                    for fi72 := 0; fi72 < 20; fi72++ {
                                        v60 = 2147483647
                                        _ = fi72
                                    }
                                    _ = v71
                                } else {
                                    fmt.Println(v60, v60, v31)
                                }
                            }
                            _ = v60
                        }
                        _ = v55
                    }
                default:
                    v1 += (v1 / ((v54) | 1))
                }
                var v73 uint32 = 2147483648
                switch ("foo bar" + ("hello" + ("a" + ""))) {
                case "\n\t":
                    var v74 float64 = (3.14159 + (0.0 + 1e10))
                    _ = v74
                default:
                    v1 = 255
                }
                _ = v73
            }
        }
        var v75 int64 = -9223372036854775808
        v76 := GenS4{F0: -456, F1: (("a" + ("" + ("hello" + ("a" + "\n\t")))) + "hello"), F2: false}
        _ = v75
        _ = v76
    } else {
        var v77 int = 9223372036854775807
        _ = v77
    }
    if (4294967295 >= 1) {
        fmt.Println(v3)
        switch "\n\t" {
        case "foo bar":
            v3 = (v3 - (v3 * (-(v3))))
        }
        var v78 float64 = 3.14159
        _ = v78
    } else {
        for fi79 := 0; fi79 < 20; fi79++ {
            switch fi79 {
            case -28:
                var v80 uint16 = (v1 * 376)
                _ = v80
            case 2147483647:
                switch fi79 {
                case -500:
                    if ((1 != int8(fi79)) && false) {
                        var v81 uint64 = genFn1(336)
                        var v82 int32 = (^(v31))
                        _ = v81
                        _ = v82
                    }
                case 127:
                    fmt.Println(v54, fi79)
                default:
                    v83 := GenS6{F0: 1, F1: 100, F2: (^(v3))}
                    _ = v83
                }
            default:
                fmt.Println(v1, v54, v1)
            }
            _ = fi79
        }
        var v84 int64 = 1
        v1 = (v54 % (((v1 | (^(v1)))) | 1))
        _ = v84
    }
    for fi85 := 0; fi85 < 3; fi85++ {
        for fi86 := 0; fi86 < 4; fi86++ {
            var v87 uint64 = 4294967295
            v3 -= (v3 ^ 986)
            _ = fi86
            _ = v87
        }
        v1 = (v1 >> ((0) & 63))
        var v88 uint8 = 209
        _ = fi85
        _ = v88
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v53
    _ = v54
}
