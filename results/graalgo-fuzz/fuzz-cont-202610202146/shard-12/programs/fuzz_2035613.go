package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int16; F1 int64; F2 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int) int8 {
    fmt.Println(p0)
    p0 = (p0 ^ p0)
    var v4 int16 = 446
    fmt.Println(v4)
    v5 := []int{p0, p0, p0, -963}
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int16) string {
    if false {
        p0 -= (p0 ^ 1000)
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int32 = 220
            _ = fi6
            _ = v7
        }
    } else {
        var v8 bool = (!(-108 > 100))
        _ = v8
    }
    v9 := []int{32767, 521}
    var v10 uint16 = 0
    switch "" {
    case "hello":
        var v11 uint8 = 127
        _ = v11
    case "a":
        v12 := GenS1{F0: 42, F1: 0, F2: (("foo bar" + "\n\t") + ("foo bar" + "\n\t"))}
        _ = v12
    case "":
        fmt.Println(p0, v10)
    default:
        p0 = p0
    }
    v10 = 865
    _ = p0
    _ = v10
    _ = v9
    return ("hello" + "a")
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]bool)
    var v3 uint32 = 761
    if false {
        v13 := []int{}
        v14 := GenS2{F0: (v1 % ((255) | 1)), F1: false, F2: int64(v1)}
        _ = v13
        _ = v14
    }
    v1 = 154
    var v15 string = ("a" + "\n\t")
    var v16 uint16 = 0
    v17 := GenS1{F0: 0, F1: -97, F2: v15}
    fmt.Println(v3)
    for fi18 := 0; fi18 < 9; fi18++ {
        switch fi18 {
        case 4294967295:
            if ((false || true) && false) {
                if false {
                    fmt.Println(v15, v16, v16)
                    v3 -= v3
                    for fi19 := 0; fi19 < 14; fi19++ {
                        var v20 uint64 = 18446744073709551615
                        _ = fi19
                        _ = v20
                    }
                } else {
                    fmt.Println(fi18, v3, v15)
                    v1 -= (v1 & 128)
                    v21 := []int{(fi18 + (fi18 ^ fi18)), -9223372036854775808}
                    _ = v21
                }
                if false {
                    for fi22 := 0; fi22 < 7; fi22++ {
                        v23 := GenS3{F0: v15}
                        v3 = v3
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS4{F0: -1}
                    _ = v24
                } else {
                    var v25 float64 = -2.25
                    switch "foo bar" {
                    case "":
                        if (42 <= 0) {
                            var v26 int16 = 41
                            var v27 uint32 = v3
                            _ = v26
                            _ = v27
                        }
                    case "foo bar":
                        var v28 int32 = -623
                        _ = v28
                    case "hello":
                        v29 := []int{fi18, (^(fi18)), (fi18 & fi18), fi18}
                        _ = v29
                    }
                    if (((v16 / ((v16) | 1)) == v16) || (v25 <= ((1.5 - (3.14159 + float64(3.14159))) - v25))) {
                        var v30 int = -345
                        _ = v30
                    }
                    _ = v25
                }
                if ((!true) || (true || (float64(1.5) == 0.0))) {
                    var v31 string = v15
                    _ = v31
                } else {
                    v1 += (^(v1))
                    if ((-(fi18)) <= (fi18 >> ((uint(v16)) & 63))) {
                        v15 = v15
                        v16 += (^(v16))
                    } else {
                        v1 -= (v1 | 53)
                    }
                    switch "\n\t" {
                    case "":
                        for fi32 := 0; fi32 < 8; fi32++ {
                            if true {
                                fmt.Println(v3)
                                switch "" {
                                case "":
                                    switch "" {
                                    case "\n\t":
                                        var v33 uint64 = 245
                                        _ = v33
                                    default:
                                        v34 := []int{fi32, (fi18 ^ int(v16)), 137}
                                        _ = v34
                                    }
                                default:
                                    if (-128 < -128) {
                                        var v35 uint64 = 326
                                        _ = v35
                                    } else {
                                        v3 -= 4294967295
                                        if (!false) {
                                            delete(v2, 736)
                                            if false {
                                                switch (fi32 ^ (fi32 - (fi32 | (fi32 ^ (-(fi32)))))) {
                                                case -473:
                                                    fmt.Println(v1)
                                                default:
                                                    var v36 uint = 1
                                                    _ = v36
                                                }
                                                var v37 uint = 2
                                                fmt.Println(v15, v15, fi32)
                                                _ = v37
                                            }
                                            if false {
                                                v1 += (v1 ^ v1)
                                            }
                                        } else {
                                            var v38 uint8 = 255
                                            _ = v38
                                        }
                                    }
                                }
                                var v39 int = fi32
                                _ = v39
                            } else {
                                if true {
                                    v15 = ((v15 + "foo bar") + ("" + "a"))
                                }
                            }
                            var v40 int32 = 0
                            v41 := v2[int(v40)]
                            _ = fi32
                            _ = v40
                            _ = v41
                        }
                    }
                }
            } else {
                var v42 int8 = 25
                var v43 int8 = genFn0((fi18 >> ((766) & 63)))
                _ = v42
                _ = v43
            }
        case 1:
            var v44 uint16 = (v16 / (((v16 + uint16(v3))) | 1))
            _ = v44
        case 298:
            if (true || true) {
                if true {
                    switch 284 {
                    case 9223372036854775807:
                        if true {
                            switch "\n\t" {
                            case "a":
                                v45 := GenS4{F0: -2}
                                _ = v45
                            case "":
                                v46 := []int{fi18, 2147483647, (^(fi18)), 127}
                                _ = v46
                            case "\n\t":
                                switch "" {
                                case "foo bar":
                                    var v47 int16 = 1000
                                    _ = v47
                                default:
                                    v48 := GenS5{F0: 42, F1: v3, F2: v3}
                                    _ = v48
                                }
                            default:
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    fmt.Println(fi49, v1, v16)
                                    if false {
                                        fmt.Println(fi18)
                                        v3 += v3
                                    } else {
                                        if false {
                                            switch fi49 {
                                            case 0:
                                                v3 += (v3 % ((v3) | 1))
                                            case 9223372036854775807:
                                                v50 := GenS3{F0: v15}
                                                _ = v50
                                            default:
                                                fmt.Println(fi49, v3)
                                            }
                                        } else {
                                            var v51 uint64 = 0
                                            _ = v51
                                        }
                                    }
                                    if false {
                                        v1 = v1
                                        v52 := GenS2{F0: (v1 >> ((0) & 63)), F1: false, F2: 0}
                                        fmt.Println(v1, fi49)
                                        _ = v52
                                    }
                                    v15 = (("hello" + (v15 + v15)) + (("a" + ("" + v15)) + (v15 + (("\n\t" + ("a" + ("foo bar" + v15))) + "hello"))))
                                    _ = fi49
                                }
                            }
                            if (true || false) {
                                if (0 != 1) {
                                    switch (fi18 << ((255) & 63)) {
                                    case -1:
                                        for fi53 := 0; fi53 < 5; fi53++ {
                                            if false {
                                                var v54 int16 = 1
                                                fmt.Println(v54, v16, v15)
                                                switch "foo bar" {
                                                case "a":
                                                    switch v15 {
                                                    case "":
                                                        switch 32768 {
                                                        case 0:
                                                            var v55 int8 = 86
                                                            _ = v55
                                                        default:
                                                            fmt.Println(v3, v3, v15)
                                                        }
                                                    case "a":
                                                        v56 := []int{(fi53 & fi18), fi18, 1, 128}
                                                        _ = v56
                                                    case "foo bar":
                                                        switch v15 {
                                                        case "foo bar":
                                                            var v57 uint16 = v16
                                                            _ = v57
                                                        default:
                                                            var v58 string = ""
                                                            _ = v58
                                                        }
                                                    default:
                                                        switch (fi53 + fi53) {
                                                        case 42:
                                                            var v59 int = 0
                                                            _ = v59
                                                        }
                                                    }
                                                case "\n\t":
                                                    if (!(!true)) {
                                                        v2[(fi18 >> ((2) & 63))] = true
                                                        v60 := []int{fi53, 128}
                                                        _ = v60
                                                    }
                                                    switch (fi18 + fi53) {
                                                    case -617:
                                                        v61 := GenS4{F0: 1}
                                                        _ = v61
                                                    case -629:
                                                        v62 := []int{-1, (-(fi18)), (fi53 ^ (-(fi18)))}
                                                        _ = v62
                                                    case 2147483647:
                                                        v15 = ("foo bar" + (v15 + v15))
                                                        switch v15 {
                                                        case "":
                                                            v63 := GenS6{F0: 42, F1: 0, F2: (^(fi18))}
                                                            _ = v63
                                                        case "hello":
                                                            var v64 float64 = (0.0 * 0.0)
                                                            _ = v64
                                                        }
                                                        var v65 uint32 = v3
                                                        _ = v65
                                                    default:
                                                        v66 := GenS3{F0: ""}
                                                        _ = v66
                                                    }
                                                    v67 := []int{fi18, (fi18 >> ((390) & 63)), fi18, fi18}
                                                    v68 := GenS7{F0: v3}
                                                    _ = v67
                                                    _ = v68
                                                }
                                                _ = v54
                                            }
                                            v69 := []int{}
                                            v16 -= v16
                                            _ = fi53
                                            _ = v69
                                        }
                                    case 42:
                                        var v70 float32 = 1.5
                                        _ = v70
                                    }
                                } else {
                                    for fi71 := 0; fi71 < 5; fi71++ {
                                        v72 := []int{822, fi18, (fi18 | (fi18 % ((fi18) | 1))), fi71}
                                        v1 = v1
                                        v15 = v15
                                        var v73 float32 = -2.25
                                        _ = fi71
                                        _ = v72
                                        _ = v73
                                    }
                                }
                                for fi74 := 0; fi74 < 16; fi74++ {
                                    v2[(fi74 - fi18)] = false
                                    _ = fi74
                                }
                                v1 = 1
                            }
                        }
                    case 32768:
                        var v75 string = ("hello" + (("\n\t" + (("foo bar" + (v15 + ("" + "foo bar"))) + ("" + "hello"))) + v15))
                        _ = v75
                    default:
                        v16 -= (v16 >> ((uint(v3)) & 63))
                    }
                    var v76 string = v15
                    v77 := []int{int(v16), fi18, fi18, (fi18 & -9223372036854775808)}
                    _ = v76
                    _ = v77
                } else {
                    v78 := []int{128}
                    var v79 int = (fi18 & fi18)
                    _ = v78
                    _ = v79
                }
                if false {
                    v15 = v15
                    v80 := []int{}
                    _ = v80
                }
            }
        }
        v15 = v15
        v1 = (v1 >> ((605) & 63))
        _ = fi18
    }
    switch 127 {
    case 127:
        v16 += (^(v16))
    case -721:
        fmt.Println(v15)
    case 42:
        v81 := GenS8{F0: 38, F1: v1}
        _ = v81
    default:
        var v82 bool = true
        _ = v82
    }
    var v83 int16 = 42
    var v84 int8 = 0
    v85 := []int{-671, -16}
    for fi86 := 0; fi86 < 3; fi86++ {
        if ((!false) && true) {
            var v87 bool = ((v1 + v1) >= (^(v1)))
            if len(v85) >= 2 {
                v85 = v85[1:len(v85)]
            }
            _ = v87
        }
        _ = fi86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("slicelen v85:", len(v85), v85)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v83
    _ = v84
    _ = v85
}
