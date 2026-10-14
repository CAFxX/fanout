package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 uint32 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int; F1 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 uint; F2 int8 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 string; F1 int16; F2 int32 }
type GenS11 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint8, p1 uint64, p2 uint16) uint {
    v7 := GenS1{F0: "\n\t", F1: 1}
    v8 := GenS1{F0: "", F1: 1}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS2{F0: 0}
        _ = fi9
        _ = v10
    }
    switch 0 {
    case 42:
        p1 += (p1 - p1)
    case -922:
        p1 = 1
    default:
        p2 = p2
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 int8 = 1
        v13 := GenS3{F0: 0, F1: (-312 == 585), F2: 4294967295}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 1
    var v3 float64 = 1.5
    v4 := make(map[int]uint32)
    var v5 int = 2147483647
    var v6 uint8 = 255
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            fmt.Println(v5)
            if (false || false) {
                var v16 int32 = 1
                var v17 int64 = 1
                if (!false) {
                    var v18 uint32 = 978
                    switch "\n\t" {
                    case "a":
                        switch (fi14 | 791) {
                        case -9223372036854775808:
                            if (189 < (v6 * 1)) {
                                var v19 string = ("hello" + "foo bar")
                                var v20 int = (fi15 % ((-2147483648) | 1))
                                _ = v19
                                _ = v20
                            } else {
                                v4[(fi15 >> ((1) & 63))] = v18
                                var v21 uint32 = (^(v18))
                                _ = v21
                            }
                        }
                    case "":
                        v17 += (v17 >> ((438) & 63))
                    default:
                        v16 = (v16 >> ((42) & 63))
                    }
                    _ = v18
                }
                var v22 bool = false
                v22 = (v5 != (fi15 | fi15))
                var v23 float64 = 3.14159
                fmt.Println(v2, v22)
                v23 += 1.5
                _ = v16
                _ = v17
                _ = v22
                _ = v23
            } else {
                v5 = (-(fi14))
                v24 := GenS4{F0: "hello", F1: false}
                var v25 uint32 = 320
                _ = v24
                _ = v25
            }
            v26 := []int{(^(fi14)), (^(fi14))}
            _ = fi15
            _ = v26
        }
        _ = fi14
    }
    var v27 uint = 782
    v28 := GenS5{F0: -624, F1: v5}
    var v29 int16 = 1000
    v3 = v3
    if (v6 != 2) {
        delete(v1, 0)
        for fi30 := 0; fi30 < 2; fi30++ {
            var v31 int32 = 0
            for fi32 := 0; fi32 < 14; fi32++ {
                v27 = uint(v6)
                var v33 uint = (v27 - v27)
                switch "hello" {
                case "":
                    v34 := v4[(fi30 << ((v27) & 63))]
                    _ = v34
                case "\n\t":
                    var v35 bool = false
                    _ = v35
                case "hello":
                    v6 -= (v6 ^ (v6 | uint8(v27)))
                    switch ("a" + (("" + ("a" + "")) + "foo bar")) {
                    case "\n\t":
                        fmt.Println(v31)
                    case "a":
                        switch "\n\t" {
                        case "":
                            var v36 uint64 = v2
                            _ = v36
                        case "\n\t":
                            var v37 string = "hello"
                            _ = v37
                        case "foo bar":
                            if (false && (0.0 <= 0.1)) {
                                if (!true) {
                                    v5 -= (fi30 * (fi32 >> ((v27) & 63)))
                                    delete(v1, 964)
                                    if ((v3 != 0.1) || (false || true)) {
                                        v38 := GenS6{F0: (fi32 + fi30), F1: v31}
                                        switch "\n\t" {
                                        case "foo bar":
                                            v2 -= (v2 ^ (^(v2)))
                                        case "a":
                                            v39 := []int{fi30, fi30, (-(v5))}
                                            _ = v39
                                        case "\n\t":
                                            if (("\n\t" + ("a" + "hello")) <= "hello") {
                                                if (!true) {
                                                    if false {
                                                        delete(v1, 1)
                                                    } else {
                                                        var v40 uint = (^(v33))
                                                        v41 := GenS7{F0: 65535}
                                                        _ = v40
                                                        _ = v41
                                                    }
                                                }
                                            }
                                        }
                                        _ = v38
                                    }
                                }
                            }
                        default:
                            v2 += 2
                        }
                    case "":
                        fmt.Println(v27)
                    }
                default:
                    v42 := GenS8{F0: 2, F1: 42, F2: -125}
                    _ = v42
                }
                _ = fi32
                _ = v33
            }
            if (false || (true || false)) {
                v3 -= v3
                var v43 float64 = (v3 * (1.5 - (1e10 * (float64(3.14159) - (v3 - (0.0 - ((v3 + 1.5) / ((v3 + (v3 + 0.0)) * ((((((float64(1e10) + v3) * (v3 + v3)) + v3) / v3) + float64(-2.25)) * (((0.0 / (v3 - v3)) * (v3 + v3)) * v3))))))))))
                var v44 uint = (v27 - (v27 * v27))
                _ = v43
                _ = v44
            }
            _ = fi30
            _ = v31
        }
        fmt.Println(v2, v27)
    }
    delete(v4, v5)
    var v45 int = v5
    v29 = (v29 * (v29 / ((1) | 1)))
    v3 = v3
    for fi46 := 0; fi46 < 10; fi46++ {
        if false {
            for fi47 := 0; fi47 < 14; fi47++ {
                if true {
                    switch "foo bar" {
                    case "":
                        var v48 uint8 = v6
                        _ = v48
                    case "foo bar":
                        var v49 float64 = v3
                        _ = v49
                    }
                    var v50 uint64 = (v2 >> (((v27 - v27)) & 63))
                    v51 := []int{(fi47 ^ fi47), v45}
                    _ = v50
                    _ = v51
                } else {
                    if true {
                        switch "" {
                        case "a":
                            v52 := GenS5{F0: 0, F1: fi47}
                            _ = v52
                        case "\n\t":
                            v2 = v2
                        }
                        fmt.Println(v6, fi47, v27)
                        switch 127 {
                        case -436:
                            v53 := GenS7{F0: 1}
                            _ = v53
                        case -309:
                            var v54 uint = 65535
                            _ = v54
                        case 128:
                            var v55 int = v5
                            _ = v55
                        default:
                            v56 := GenS9{F0: (v29 | (v29 >> ((genFn0(0, v2, 552)) & 63))), F1: v2}
                            if (18446744073709551615 != v2) {
                                if (0 != -1) {
                                    if false {
                                        var v57 int8 = 127
                                        _ = v57
                                    } else {
                                        var v58 uint64 = (v2 - (^(v2)))
                                        fmt.Println(v2, v58)
                                        v59 := []int{}
                                        _ = v58
                                        _ = v59
                                    }
                                    if true {
                                        v45 -= v5
                                        v45 = v45
                                        v2 -= v2
                                        var v60 int32 = 1
                                        _ = v60
                                    }
                                    v5 += 128
                                } else {
                                    var v61 bool = false
                                    if true {
                                        v62 := GenS8{F0: 65535, F1: v27, F2: 1}
                                        _ = v62
                                    } else {
                                        var v63 int8 = 123
                                        v45 = (^(fi46))
                                        v4[-425] = 2147483647
                                        _ = v63
                                    }
                                    _ = v61
                                }
                                fmt.Println(v3, v5, v27)
                            }
                            v5 = fi47
                            v64 := GenS10{F0: "a", F1: (v29 / (((-(v29))) | 1)), F2: -381}
                            var v65 int64 = -9223372036854775808
                            _ = v56
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v45)
                    }
                    switch "a" {
                    case "":
                        fmt.Println(v2, v3, fi46)
                    }
                    fmt.Println(v29, v5, v3)
                }
                _ = fi47
            }
            v66 := []int{}
            _ = v66
        }
        v29 -= (v29 ^ (v29 >> ((v27) & 63)))
        var v67 uint8 = v6
        _ = fi46
        _ = v67
    }
    v3 += float64(0.0)
    for fi68 := 0; fi68 < 11; fi68++ {
        v69 := []int{v45, -734}
        switch (v5 & (v5 / ((v5) | 1))) {
        case 2147483647:
            if false {
                for fi70 := 0; fi70 < 5; fi70++ {
                    v71 := v4[int(v6)]
                    v6 -= v6
                    v72 := GenS11{F0: 1, F1: 560}
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                v3 += ((v3 * v3) * (float64(-2.25) / ((v3 + v3) / 1e10)))
                delete(v4, (v5 % ((v5) | 1)))
            }
        }
        v2 -= v2
        _ = fi68
        _ = v69
    }
    v29 += 260
    for fi73 := 0; fi73 < 15; fi73++ {
        v74 := []int{}
        v3 += ((-2.25 - v3) - (0.0 - -2.25))
        _ = fi73
        _ = v74
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
}
