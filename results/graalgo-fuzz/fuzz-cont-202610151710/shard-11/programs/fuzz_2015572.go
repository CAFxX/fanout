package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint16, p1 string) uint64 {
    var v5 uint16 = (p0 | 720)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = -497
        _ = fi6
        _ = v7
    }
    v8 := []int{-420, 338, 127, 4294967295}
    p0 = 42
    fmt.Println(p1, p1, v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    return 851
}
func genFn1(p0 uint64) bool {
    switch "\n\t" {
    case "hello":
        v9 := []int{}
        _ = v9
    }
    if true {
        if true {
            var v10 int = -1
            _ = v10
        } else {
            v11 := GenS1{F0: 127, F1: 27}
            fmt.Println(p0)
            if true {
                var v12 float32 = (0.1 + 3.14159)
                for fi13 := 0; fi13 < 12; fi13++ {
                    if true {
                        v12 += (v12 - v12)
                        v14 := GenS1{F0: 1, F1: 843}
                        _ = v14
                    } else {
                        fmt.Println(fi13, fi13, p0)
                        v15 := GenS2{F0: 2, F1: 38, F2: 9223372036854775807}
                        _ = v15
                    }
                    p0 += 2
                    _ = fi13
                }
                _ = v12
            }
            _ = v11
        }
        if (!true) {
            var v16 uint8 = 239
            fmt.Println(v16)
            _ = v16
        }
        p0 -= 18446744073709551615
    }
    fmt.Println(p0)
    if ((((^(p0)) > (p0 | 99)) && false) && (132 <= -2147483648)) {
        if true {
            if (1 >= genFn0(0, "\n\t")) {
                var v17 string = "foo bar"
                fmt.Println(p0, v17)
                p0 += (p0 % ((p0) | 1))
                _ = v17
            }
            p0 += (p0 - 2)
        } else {
            var v18 uint32 = 2147483648
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            fmt.Println(fi19, p0)
            var v20 uint8 = 255
            _ = fi19
            _ = v20
        }
    }
    _ = p0
    return (0.1 >= 0.1)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    var v3 uint32 = 2147483647
    var v4 int16 = 42
    for fi21 := 0; fi21 < 20; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 uint = 831
            _ = fi22
            _ = v23
        }
        v3 = (v3 | v3)
        var v24 uint16 = 1
        _ = fi21
        _ = v24
    }
    var v25 uint = 65535
    switch 127 {
    case 9223372036854775807:
        var v26 uint16 = 2
        _ = v26
    case 247:
        var v27 int32 = -2147483648
        _ = v27
    case 2147483647:
        if true {
            switch "foo bar" {
            case "\n\t":
                var v28 int = 290
                _ = v28
            case "":
                var v29 uint32 = v3
                _ = v29
            case "hello":
                v25 = 441
            default:
                _ = len(v1)
            }
            switch "" {
            case "hello":
                v2 = append(v2, (("" + "\n\t") + "hello"))
            case "\n\t":
                for fi30 := 0; fi30 < 14; fi30++ {
                    fmt.Println(v4, v3, v4)
                    if (!false) {
                        v31 := GenS1{F0: 1, F1: 2}
                        _ = v31
                    } else {
                        if true {
                            var v32 int16 = (v4 ^ 0)
                            var v33 int = fi30
                            _ = v32
                            _ = v33
                        } else {
                            v25 -= v25
                            if (v25 <= (v25 << (((^(v25))) & 63))) {
                                if true {
                                    v25 += (v25 ^ (v25 % ((v25) | 1)))
                                } else {
                                    var v34 uint64 = 2
                                    _ = v34
                                }
                                var v35 int64 = 9223372036854775807
                                _ = v35
                            } else {
                                v4 += (v4 >> ((20) & 63))
                                switch ("hello" + (("a" + ("" + "")) + "\n\t")) {
                                case "":
                                    switch ((("a" + "a") + "\n\t") + "\n\t") {
                                    case "a":
                                        v3 = (v3 % (((v3 ^ (v3 ^ v3))) | 1))
                                    default:
                                        if ((!((v4 & (v4 | v4)) >= v4)) && (((0.1 / (((1.5 + ((3.14159 - (float64(float32(0.0)) * -2.25)) + (0.1 - (1.5 * 0.1)))) * 0.1) / ((1.5 / 3.14159) * (1.5 + ((-2.25 - (((1.5 - 0.0) * 1e10) / -2.25)) * (0.0 - float64(1e10))))))) / 1e10) > 0.1)) {
                                            v4 -= v4
                                            fmt.Println(v3)
                                            fmt.Println(v4, v25, fi30)
                                        }
                                    }
                                default:
                                    v36 := GenS1{F0: 1, F1: 1}
                                    _ = v36
                                }
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            }
                        }
                    }
                    _ = fi30
                }
            }
            v37 := GenS3{F0: 1}
            _ = v37
        }
    default:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    if (93 < 42) {
        v3 += 42
    } else {
        if true {
            var v38 float32 = 1.5
            _ = v38
        } else {
            v1[(^(v3))] = false
            _ = len(v1)
            for fi39 := 0; fi39 < 14; fi39++ {
                v40 := GenS4{F0: genFn0(255, ("a" + ("" + "")))}
                fmt.Println(v25)
                _ = fi39
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            v3 = v3
            _ = fi41
        }
        var v42 int = 1
        _ = v42
    }
    if (((float32(-2.25) * 1.5) * 1e10) != 3.14159) {
        if true {
            for fi43 := 0; fi43 < 9; fi43++ {
                v44 := GenS3{F0: genFn0(1, ("a" + "a"))}
                _ = fi43
                _ = v44
            }
        } else {
            if false {
                v45 := GenS5{F0: -2147483648}
                var v46 uint64 = 4294967295
                _ = v45
                _ = v46
            } else {
                fmt.Println(v3)
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        switch "hello" {
        case "foo bar":
            for fi47 := 0; fi47 < 10; fi47++ {
                if false {
                    if (0.0 != 1e10) {
                        v48 := GenS5{F0: 0}
                        v49 := GenS6{F0: -1, F1: 128, F2: 42}
                        _ = len(v1)
                        _ = v48
                        _ = v49
                    } else {
                        fmt.Println(v4, fi47, v4)
                        var v50 uint32 = v3
                        _ = v50
                    }
                    var v51 float32 = ((0.1 * -2.25) * 3.14159)
                    if ((v4 % ((v4) | 1)) != int16(fi47)) {
                        v52 := GenS7{F0: true}
                        for fi53 := 0; fi53 < 15; fi53++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = len(v2)
                            if ((4294967295 == 376) || true) {
                                if (true || false) {
                                    fmt.Println(v3, v4)
                                    var v54 int = (fi47 % ((fi53) | 1))
                                    _ = v54
                                }
                                var v55 int16 = (v4 >> ((2) & 63))
                                v3 = v3
                                _ = v55
                            } else {
                                if ((3.14159 / 1.0) <= 3.14159) {
                                    var v56 bool = (32767 >= fi53)
                                    switch "a" {
                                    case "":
                                        if ((351 != -1) || (fi47 == (fi47 >> (((v25 % (((v25 >> (((v25 & v25)) & 63))) | 1))) & 63)))) {
                                            v56 = ((-(fi53)) != (fi53 % (((fi53 >> ((v25) & 63))) | 1)))
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v3)
                                            }
                                        } else {
                                            var v57 int8 = 42
                                            var v58 int = (^(fi47))
                                            _ = v57
                                            _ = v58
                                        }
                                    case "foo bar":
                                        v3 = 413
                                    case "\n\t":
                                        switch 32767 {
                                        case -1:
                                            var v59 string = ("hello" + "hello")
                                            _ = v59
                                        }
                                    default:
                                        var v60 uint16 = 42
                                        _ = v60
                                    }
                                    _ = v56
                                }
                            }
                            _ = fi53
                        }
                        v4 += (v4 >> (((^(v25))) & 63))
                        _ = v52
                    }
                    _ = v51
                } else {
                    v2 = append(v2, "hello")
                    for fi61 := 0; fi61 < 17; fi61++ {
                        switch fi47 {
                        case -9223372036854775808:
                            var v62 int32 = 400
                            _ = v62
                        case 4294967295:
                            if false {
                                if false {
                                    if ((-2147483648 <= 722) || (-333 <= -921)) {
                                        if (90 == 128) {
                                            v63 := GenS8{F0: (^(v4))}
                                            v4 -= 486
                                            if true {
                                                v3 = v3
                                                if false {
                                                    var v64 uint = 18446744073709551615
                                                    v25 -= v64
                                                    v3 = 2147483648
                                                    _ = v64
                                                }
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    var v65 bool = true
                                                    _ = v65
                                                case "\n\t":
                                                    fmt.Println(v4)
                                                case "a":
                                                    var v66 int = -339
                                                    _ = v66
                                                }
                                                var v67 int8 = 127
                                                fmt.Println(v4)
                                                var v68 uint16 = 2
                                                _ = v67
                                                _ = v68
                                            }
                                            v1[v3] = (true || true)
                                            _ = v63
                                        }
                                        v3 -= 144
                                    } else {
                                        if (false && false) {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        var v69 float32 = -2.25
                                        var v70 uint = 684
                                        if len(v2) > 0 {
                                            v2[0] = ("hello" + "\n\t")
                                        }
                                        var v71 int64 = 0
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                }
                                v72 := v1[1]
                                _ = v72
                            }
                        default:
                            if ((v25 * v25) < (v25 / (((v25 - 306)) | 1))) {
                                var v73 int64 = 205
                                v74 := GenS9{F0: 75}
                                _ = v73
                                _ = v74
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    v4 = v4
                                }
                            }
                        }
                        v4 = v4
                        var v75 int8 = 127
                        if (fi47 >= (fi61 >> ((0) & 63))) {
                            var v76 int = fi61
                            _ = v76
                        } else {
                            if false {
                                _ = len(v1)
                                v1[v3] = (42 == 100)
                            } else {
                                var v77 uint16 = 42
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v78 := GenS10{F0: -1, F1: 2}
                                var v79 int = (fi47 / ((32768) | 1))
                                _ = v77
                                _ = v78
                                _ = v79
                            }
                        }
                        _ = fi61
                        _ = v75
                    }
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "foo bar")
                    }
                }
                _ = fi47
            }
        case "\n\t":
            v2 = append(v2, "foo bar")
        case "hello":
            fmt.Println(v3)
        default:
            var v80 int = 996
            _ = v80
        }
    }
    var v81 int16 = v4
    for fi82 := 0; fi82 < 17; fi82++ {
        var v83 int32 = -1
        _ = fi82
        _ = v83
    }
    fmt.Println(v81, v25)
    if true {
        if false {
            var v84 uint8 = uint8(v25)
            _ = v84
        }
    }
    fmt.Println(v3)
    fmt.Println(v4, v3, v81)
    for fi85 := 0; fi85 < 7; fi85++ {
        v2 = append(v2, ("hello" + ""))
        switch fi85 {
        case -15:
            var v86 int64 = 994
            _ = v86
        case 95:
            var v87 int = 712
            _ = v87
        default:
            var v88 float32 = 1e10
            _ = v88
        }
        _ = fi85
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v81
}
