package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int64; F2 bool }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint64; F2 int16 }
func genFn0(p0 int64, p1 string) int {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{-2147483648}
        var v8 int32 = 2147483647
        switch (p1 + ("" + p1)) {
        case "":
            fmt.Println(p0, p0, p1)
        case "foo bar":
            for fi9 := 0; fi9 < 18; fi9++ {
                v7 = append(v7, fi9)
                fmt.Println(p1, v8)
                _ = fi9
            }
        case "a":
            _ = len(v7)
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v10 := []int{128, 891, 0, 32768}
    var v11 uint32 = 65535
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 915
}
func genFn1(p0 int, p1 int, p2 int64) int {
    if (-128 > -1) {
        if false {
            p2 += (p2 >> ((2) & 63))
        } else {
            switch "\n\t" {
            case "\n\t":
                p1 -= (^(p0))
            case "":
                fmt.Println(p0)
            case "foo bar":
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: (p0 & fi12), F1: 255}
                    _ = fi12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                switch "\n\t" {
                case "hello":
                    p1 = p0
                case "foo bar":
                    var v15 int32 = 1
                    _ = v15
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint64 = 42
                    fmt.Println(fi14, fi16, v17)
                    _ = fi16
                    _ = v17
                }
                _ = fi14
            }
            p2 += p2
        }
        var v18 int8 = 120
        fmt.Println(p2, p1, p0)
        _ = v18
    } else {
        var v19 uint8 = 127
        _ = v19
    }
    p2 -= 42
    for fi20 := 0; fi20 < 19; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 float64 = 1.5
            var v23 uint = 18446744073709551615
            var v24 uint64 = 2
            v25 := GenS2{F0: 4294967295}
            if true {
                var v26 int64 = 1
                v22 = (3.14159 + float64(-2.25))
                v24 = v24
                _ = v26
            }
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
    }
    p1 = (p0 / ((8) | 1))
    v27 := []int{p1, p0, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return p1
}

func main() {
    v1 := []int{}
    var v2 int8 = -108
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    if len(v1) > 0 {
        v1[0] = 735
    }
    v2 -= (v2 / (((v2 | v2)) | 1))
    if true {
        if false {
            if (true || ((0.1 * float64(1.5)) >= 3.14159)) {
                var v28 uint64 = 0
                v2 = (^(v2))
                _ = v28
            } else {
                v29 := GenS3{F0: 127}
                fmt.Println(v3, v2)
                switch "foo bar" {
                case "hello":
                    var v30 uint8 = 0
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 17; fi31++ {
                        switch "" {
                        case "":
                            if false {
                                fmt.Println(fi31, v2, v2)
                                v3 = v3
                            }
                        case "a":
                            for fi32 := 0; fi32 < 2; fi32++ {
                                if (v2 >= (^(v2))) {
                                    fmt.Println(fi31)
                                    switch (("foo bar" + ("a" + "\n\t")) + "\n\t") {
                                    case "\n\t":
                                        switch (fi31 + 32767) {
                                        case -1:
                                            var v33 int16 = 42
                                            _ = v33
                                        case 128:
                                            var v34 float32 = 0.0
                                            _ = v34
                                        case 32768:
                                            delete(v5, (("hello" + "") + ("" + ("" + ""))))
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                fmt.Println(fi32)
                                            case "\n\t":
                                                var v35 float64 = 3.14159
                                                _ = v35
                                            case "":
                                                fmt.Println(fi31, fi32)
                                            }
                                        }
                                    case "hello":
                                        v2 = (v2 | 127)
                                    case "":
                                        fmt.Println(fi31, v2, v3)
                                    }
                                } else {
                                    var v36 uint8 = 2
                                    if (0 <= 255) {
                                        switch fi32 {
                                        case 666:
                                            switch "\n\t" {
                                            case "a":
                                                v3 = v3
                                            }
                                        case 4294967295:
                                            if true {
                                                v3 = uint32(v36)
                                                if (!(0 == -754)) {
                                                    var v37 bool = (1 <= 0)
                                                    var v38 uint16 = 32767
                                                    _ = v37
                                                    _ = v38
                                                }
                                            } else {
                                                v39 := v5["a"]
                                                _ = v39
                                            }
                                        default:
                                            v2 = -89
                                        }
                                        if (false && true) {
                                            v40 := GenS3{F0: (v36 | v36)}
                                            v3 -= (v3 ^ v3)
                                            _ = v40
                                        } else {
                                            if ((!((18446744073709551615 == 2) && ((!true) && true))) && false) {
                                                fmt.Println(fi31, fi31, fi32)
                                                _ = len(v5)
                                            } else {
                                                v2 -= v2
                                                delete(v4, v3)
                                                if (false && ((true || false) && ((!(485 >= 65535)) || true))) {
                                                    var v41 uint8 = uint8(v3)
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi31 ^ 1)
                                                    }
                                                    _ = v41
                                                }
                                            }
                                        }
                                    }
                                    _ = v36
                                }
                                fmt.Println(fi31, v2, v2)
                                fmt.Println(v2, v3, v2)
                                _ = fi32
                            }
                        case "\n\t":
                            v1 = append(v1, genFn0(9223372036854775807, (("" + "") + "foo bar")))
                        default:
                            var v42 int64 = 0
                            _ = v42
                        }
                        fmt.Println(v2)
                        v3 += 42
                        _ = fi31
                    }
                }
                _ = v29
            }
            v2 -= v2
            v43 := GenS2{F0: 2}
            _ = v43
        }
        v44 := GenS4{F0: -627, F1: 42, F2: true}
        _ = v44
    } else {
        v2 -= v2
        v2 = v2
        fmt.Println(v2)
    }
    var v45 int16 = 347
    v2 = v2
    var v46 int = genFn1(42, -2147483648, 9223372036854775807)
    var v47 uint16 = 87
    for fi48 := 0; fi48 < 5; fi48++ {
        var v49 int16 = v45
        if false {
            switch "" {
            case "a":
                _ = len(v5)
            case "hello":
                switch v46 {
                case 128:
                    v49 -= v49
                }
            }
            if (!false) {
                v50 := GenS5{F0: 2, F1: (v47 & (v47 + 218)), F2: v47}
                if len(v1) > 0 {
                    v1[0] = (fi48 * genFn1(fi48, (v46 + fi48), 9223372036854775807))
                }
                v46 -= -1
                _ = v50
            }
        } else {
            fmt.Println(v47)
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(fi48)
        default:
            if false {
                v49 -= (v49 % (((v49 % ((int16(v2)) | 1))) | 1))
            } else {
                if (float32(3.14159) > 1e10) {
                    for fi51 := 0; fi51 < 1; fi51++ {
                        v52 := v4[v3]
                        switch v46 {
                        case 234:
                            var v53 uint32 = v3
                            _ = v53
                        }
                        if v52 {
                            if true {
                                v52 = (!(v52 || ("hello" <= "\n\t")))
                            } else {
                                _ = len(v5)
                                v46 = (fi48 - fi48)
                                delete(v4, v3)
                            }
                            v1 = append(v1, 4294967295)
                            v46 += v46
                        } else {
                            switch "foo bar" {
                            case "":
                                v1 = append(v1, genFn1(32767, (^(v46)), -412))
                            }
                        }
                        _ = fi51
                        _ = v52
                    }
                    var v54 uint32 = 1
                    _ = v54
                } else {
                    v55 := GenS4{F0: 0, F1: 9223372036854775807, F2: false}
                    v2 -= (-(v2))
                    _ = v55
                }
                v56 := GenS4{F0: genFn1((fi48 & 0), genFn1(v46, genFn0(1, "hello"), int64(v47)), 76), F1: 0, F2: (46 == 2)}
                _ = v56
            }
        }
        _ = fi48
        _ = v49
    }
    if true {
        v46 = (^(v46))
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        v3 = (v3 / ((4294967295) | 1))
    case "\n\t":
        var v57 bool = true
        _ = v57
    }
    var v58 uint8 = 42
    v58 -= v58
    v1 = append(v1, genFn1((v46 << ((723) & 63)), v46, -1))
    for fi59 := 0; fi59 < 10; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            var v61 uint32 = v3
            if (v46 > -1) {
                fmt.Println(fi59, v58, fi60)
            }
            switch "hello" {
            case "a":
                var v62 uint64 = 1
                _ = v62
            case "\n\t":
                v1 = append(v1, 0)
            default:
                var v63 int8 = 42
                _ = v63
            }
            _ = fi60
            _ = v61
        }
        v58 += (v58 + uint8(v47))
        v64 := GenS2{F0: 193}
        _ = fi59
        _ = v64
    }
    if false {
        v2 = (v2 & v2)
        var v65 int64 = int64(v46)
        _ = v65
    } else {
        v66 := GenS2{F0: 42}
        v67 := GenS6{F0: (v58 + (^(v58)))}
        _ = v66
        _ = v67
    }
    if false {
        if false {
            for fi68 := 0; fi68 < 7; fi68++ {
                v69 := GenS5{F0: 598, F1: 1, F2: 2}
                var v70 float32 = 1.5
                v71 := GenS7{F0: fi68, F1: 260, F2: -933}
                _ = fi68
                _ = v69
                _ = v70
                _ = v71
            }
        }
        v72 := GenS2{F0: 1}
        _ = v72
    }
    v73 := GenS3{F0: v58}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v58
    _ = v73
}
