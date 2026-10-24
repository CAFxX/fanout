package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 uint; F2 int64 }
type GenS6 struct { F0 uint; F1 int8 }
func genFn0(p0 int, p1 uint64, p2 uint16) int8 {
    var v5 uint64 = (p1 ^ (p1 % ((p1) | 1)))
    v6 := []int{(p0 * p0), (p0 ^ -341)}
    if false {
        if true {
            p2 = p2
        }
        v7 := GenS1{F0: (!((1 < 4294967295) || false)), F1: 0}
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p1 -= 475
        switch (p0 << ((1) & 63)) {
        case -2147483648:
            if len(v6) > 0 {
                v6[0] = p0
            }
        case -9223372036854775808:
            for fi9 := 0; fi9 < 2; fi9++ {
                p1 += (p1 & (p1 + p1))
                var v10 uint16 = (p2 * (p2 ^ uint16(v5)))
                switch fi9 {
                case 42:
                    p2 -= (p2 * v10)
                case 128:
                    if false {
                        var v11 int16 = 1
                        _ = v11
                    } else {
                        var v12 bool = true
                        _ = v12
                    }
                case 127:
                    var v13 uint64 = (p1 / ((42) | 1))
                    _ = v13
                default:
                    var v14 bool = (uint64(v10) == 0)
                    _ = v14
                }
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v15 := GenS2{F0: 0}
                _ = fi9
                _ = v10
                _ = v15
            }
        case 2147483647:
            fmt.Println(p1, p2)
        default:
            var v16 int8 = 100
            _ = v16
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 84
}
func genFn1(p0 int32) int {
    v17 := GenS2{F0: -498}
    v18 := []int{2147483647, -1, 2147483647}
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 int = fi19
        _ = fi19
        _ = v20
    }
    p0 = 42
    _ = p0
    _ = v17
    _ = v18
    return 42
}

func main() {
    var v1 string = ""
    var v2 uint8 = 2
    var v3 uint8 = 126
    var v4 uint8 = v2
    v4 = (v3 >> ((4294967295) & 63))
    var v21 uint = 65535
    switch "a" {
    case "\n\t":
        for fi22 := 0; fi22 < 8; fi22++ {
            if true {
                if (-1 == 9223372036854775807) {
                    v2 = (v4 % ((78) | 1))
                    fmt.Println(v1)
                    var v23 uint16 = 328
                    _ = v23
                }
                v3 += (v3 << (((v21 + 2)) & 63))
            } else {
                var v24 uint16 = 966
                _ = v24
            }
            _ = fi22
        }
    case "a":
        fmt.Println(v2, v2)
    case "":
        v25 := []int{783, genFn1(1)}
        _ = v25
    }
    if false {
        v26 := []int{-771, -194, -2147483648}
        if len(v26) > 0 {
            v26[0] = 630
        }
        _ = v26
    } else {
        v27 := []int{127, int(v4)}
        if true {
            for fi28 := 0; fi28 < 17; fi28++ {
                fmt.Println(v4, v3, v21)
                _ = fi28
            }
            v27 = append(v27, 9223372036854775807)
        } else {
            if true {
                if len(v27) > 0 {
                    v27[0] = genFn1(-2147483648)
                }
            } else {
                for fi29 := 0; fi29 < 19; fi29++ {
                    v30 := GenS2{F0: 1}
                    var v31 int16 = 1
                    switch v1 {
                    case "a":
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v4 += (v3 & v3)
                            switch fi32 {
                            case 709:
                                var v33 float64 = (1e10 + 0.1)
                                _ = v33
                            default:
                                v34 := GenS3{F0: 2}
                                _ = v34
                            }
                            v35 := GenS2{F0: 42}
                            _ = fi32
                            _ = v35
                        }
                    default:
                        _ = len(v27)
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                if true {
                    for fi36 := 0; fi36 < 13; fi36++ {
                        for fi37 := 0; fi37 < 9; fi37++ {
                            var v38 bool = true
                            if (!true) {
                                v39 := GenS3{F0: 4294967295}
                                var v40 int = genFn1(750)
                                var v41 uint64 = 770
                                _ = v39
                                _ = v40
                                _ = v41
                            } else {
                                var v42 float32 = float32(3.14159)
                                var v43 bool = (0 != 0)
                                switch ((v1 + ("a" + v1)) + v1) {
                                case "":
                                    v38 = (fi36 != fi36)
                                }
                                _ = v42
                                _ = v43
                            }
                            _ = fi37
                            _ = v38
                        }
                        _ = fi36
                    }
                    var v44 uint32 = 1
                    var v45 uint16 = 255
                    _ = v44
                    _ = v45
                }
            }
            fmt.Println(v2, v3, v1)
        }
        _ = v27
    }
    if false {
        v46 := []int{-753, 1}
        _ = v46
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        fmt.Println(v21, v1)
        var v48 int64 = -825
        _ = fi47
        _ = v48
    }
    var v49 uint32 = 975
    var v50 int32 = 235
    switch int(v21) {
    case -405:
        for fi51 := 0; fi51 < 12; fi51++ {
            var v52 string = v1
            v53 := GenS4{F0: (v2 & 2)}
            v49 += (v49 / ((911) | 1))
            _ = fi51
            _ = v52
            _ = v53
        }
    default:
        fmt.Println(v1, v49)
    }
    if false {
        v54 := GenS1{F0: false, F1: int16(v49)}
        v55 := []int{-953, -1, 4294967295}
        _ = v54
        _ = v55
    }
    v56 := []int{genFn1((-(v50))), -1, -350}
    v49 += (v49 % (((v49 << ((v21) & 63))) | 1))
    v2 = 255
    v57 := GenS1{F0: (((false && true) && false) || (42 == v21)), F1: 1}
    v58 := GenS5{F0: 127, F1: (^(v21)), F2: 985}
    if true {
        for fi59 := 0; fi59 < 14; fi59++ {
            var v60 int32 = v50
            switch "foo bar" {
            case "\n\t":
                _ = len(v56)
            case "":
                for fi61 := 0; fi61 < 13; fi61++ {
                    v60 = (v60 >> (((v21 ^ 0)) & 63))
                    if true {
                        var v62 uint = v21
                        v63 := GenS6{F0: v21, F1: 119}
                        v64 := GenS2{F0: 0}
                        v1 = "hello"
                        _ = len(v56)
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    var v65 uint = v21
                    switch (fi59 ^ genFn1(int32(v3))) {
                    case 429:
                        var v66 string = v1
                        v3 -= v2
                        _ = v66
                    }
                    if len(v56) >= 2 {
                        v56 = v56[1:len(v56)]
                    }
                    _ = fi61
                    _ = v65
                }
            default:
                fmt.Println(fi59)
            }
            _ = fi59
            _ = v60
        }
    }
    v67 := GenS3{F0: (v49 << ((v21) & 63))}
    for fi68 := 0; fi68 < 9; fi68++ {
        switch genFn1((v50 + v50)) {
        case -597:
            if false {
                var v69 int64 = 1
                _ = v69
            }
        case 42:
            switch ("a" + "a") {
            case "":
                v1 = ""
            case "\n\t":
                fmt.Println(v3)
            case "foo bar":
                fmt.Println(v3)
            default:
                var v70 uint8 = v2
                _ = v70
            }
        default:
            for fi71 := 0; fi71 < 20; fi71++ {
                if false {
                    if true {
                        var v72 uint64 = 287
                        if len(v56) > 0 {
                            v56[0] = fi68
                        }
                        var v73 uint16 = 255
                        _ = v72
                        _ = v73
                    }
                    var v74 uint64 = 0
                    var v75 int64 = 0
                    if (!true) {
                        v2 = (v3 | (v4 | 127))
                    }
                    var v76 uint = (v21 % ((v21) | 1))
                    _ = v74
                    _ = v75
                    _ = v76
                }
                _ = fi71
            }
        }
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v56
    _ = v57
    _ = v58
    _ = v67
}
