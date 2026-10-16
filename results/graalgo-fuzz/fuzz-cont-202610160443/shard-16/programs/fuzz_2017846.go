package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int16; F1 string; F2 uint }
type GenS5 struct { F0 int32; F1 uint32; F2 bool }
type GenS6 struct { F0 uint; F1 int64; F2 uint16 }
type GenS7 struct { F0 int; F1 uint32; F2 int }
type GenS8 struct { F0 uint16; F1 int; F2 bool }
type GenS9 struct { F0 int64; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint8; F1 bool }
type GenS11 struct { F0 uint }
func genFn0(p0 int16) bool {
    var v4 uint64 = 42
    switch -713 {
    case 9223372036854775807:
        for fi5 := 0; fi5 < 11; fi5++ {
            if (!false) {
                v4 -= 42
            }
            if false {
                v6 := []int{}
                _ = v6
            }
            var v7 uint32 = uint32(v4)
            _ = fi5
            _ = v7
        }
    case 1:
        v8 := GenS1{F0: 829, F1: (p0 - 1)}
        _ = v8
    default:
        switch "foo bar" {
        case "hello":
            v9 := []int{42, 32767, 32768, 32767}
            _ = v9
        default:
            v10 := []int{42}
            _ = v10
        }
    }
    if (-2147483648 < 1) {
        if false {
            v11 := GenS1{F0: 1000, F1: (p0 / ((p0) | 1))}
            if (-2.25 == 1.5) {
                v4 = 1
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 float64 = float64(1e10)
                    v14 := GenS2{F0: 544, F1: 42}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 uint16 = 65535
                _ = v15
            } else {
                v4 = v4
                var v16 uint32 = 1
                _ = v16
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v19 := GenS2{F0: 42, F1: 32767}
                    _ = fi18
                    _ = v19
                }
                v20 := []int{fi17, 1}
                _ = fi17
                _ = v20
            }
            _ = v11
        }
        fmt.Println(v4, p0)
    } else {
        p0 -= (p0 % (((p0 ^ p0)) | 1))
    }
    v21 := []int{32768, 128}
    _ = p0
    _ = v21
    _ = v4
    return false
}
func genFn1(p0 uint16, p1 int32, p2 uint16) uint64 {
    fmt.Println(p2)
    var v22 uint16 = p0
    switch (("" + "\n\t") + "foo bar") {
    case "hello":
        if ((1e10 * (1.5 / 1.0)) >= -2.25) {
            fmt.Println(p1, p1, p0)
        }
    case "a":
        var v23 int32 = (p1 - 1)
        _ = v23
    default:
        var v24 uint16 = (^(p2))
        _ = v24
    }
    if true {
        v22 -= (v22 + 255)
        switch 1 {
        case 609:
            fmt.Println(p1, p1, v22)
        case 2147483647:
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 uint = 0
                v26 -= uint(v22)
                _ = fi25
                _ = v26
            }
        default:
            switch -924 {
            case 32768:
                if false {
                    if true {
                        v27 := []int{127, -667, 128}
                        _ = len(v27)
                        _ = v27
                    } else {
                        v28 := GenS2{F0: 2, F1: p0}
                        if true {
                            v29 := GenS3{F0: 42, F1: (^(v22))}
                            var v30 uint = 294
                            v31 := GenS1{F0: 0, F1: -947}
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 12; fi32++ {
                                switch fi32 {
                                case -325:
                                    switch "a" {
                                    case "\n\t":
                                        switch ("" + "\n\t") {
                                        case "foo bar":
                                            switch (("a" + "foo bar") + "hello") {
                                            case "a":
                                                if true {
                                                    p1 += (p1 * (p1 & (p1 * (p1 * -1))))
                                                    if true {
                                                        p0 -= v22
                                                    } else {
                                                        p1 += 42
                                                    }
                                                }
                                            default:
                                                switch (^(fi32)) {
                                                case 42:
                                                    v22 -= (^(p0))
                                                }
                                            }
                                        default:
                                            fmt.Println(v22)
                                        }
                                    case "":
                                        v33 := GenS1{F0: 991, F1: 0}
                                        _ = v33
                                    case "a":
                                        fmt.Println(fi32)
                                    default:
                                        for fi34 := 0; fi34 < 13; fi34++ {
                                            v22 += p0
                                            v35 := GenS4{F0: 1000, F1: (("hello" + "\n\t") + "hello"), F2: 182}
                                            _ = fi34
                                            _ = v35
                                        }
                                    }
                                }
                                p2 = p0
                                v36 := []int{fi32, fi32, (fi32 << ((42) & 63))}
                                _ = fi32
                                _ = v36
                            }
                        }
                        _ = v28
                    }
                    switch -527 {
                    case 2147483647:
                        switch 112 {
                        case 42:
                            switch "" {
                            case "\n\t":
                                v37 := GenS2{F0: 4294967295, F1: p2}
                                _ = v37
                            case "":
                                p0 = (v22 - v22)
                            case "foo bar":
                                fmt.Println(p1)
                            }
                        }
                    }
                }
            case -2147483648:
                p2 -= 65535
            default:
                p0 -= (^(v22))
            }
        }
        var v38 int64 = 0
        _ = v38
    } else {
        if true {
            for fi39 := 0; fi39 < 9; fi39++ {
                v22 = p0
                var v40 uint = 950
                _ = fi39
                _ = v40
            }
            fmt.Println(p0)
            p1 = p1
        }
        p2 -= v22
        if (3.14159 > (((0.1 * ((float32(1e10) - 3.14159) * 1.5)) + 0.0) * 0.0)) {
            v41 := []int{-2147483648}
            if true {
                switch "a" {
                case "foo bar":
                    p2 -= p2
                case "":
                    v42 := GenS5{F0: 2147483647, F1: 58, F2: true}
                    _ = v42
                case "hello":
                    var v43 uint8 = 42
                    _ = v43
                }
            } else {
                p2 -= 255
                v44 := GenS5{F0: p1, F1: 134, F2: true}
                if len(v41) >= 2 {
                    v41 = v41[1:len(v41)]
                }
                _ = v44
            }
            if len(v41) > 0 {
                v41[0] = 9223372036854775807
            }
            _ = v41
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return 4294967295
}

func main() {
    v1 := []string{("a" + ("" + "")), "foo bar", "foo bar"}
    var v2 int16 = 32767
    var v3 int64 = 1
    var v45 bool = ((!false) || ((v3 | v3) != (v3 >> ((668) & 63))))
    switch ((("foo bar" + "foo bar") + ("\n\t" + "hello")) + "foo bar") {
    case "\n\t":
        v1 = append(v1, "\n\t")
    case "foo bar":
        v45 = (true || true)
    case "":
        for fi46 := 0; fi46 < 15; fi46++ {
            v45 = false
            if v45 {
                switch "a" {
                case "":
                    for fi47 := 0; fi47 < 9; fi47++ {
                        if (false || true) {
                            v48 := GenS6{F0: 4294967295, F1: 868, F2: 65535}
                            switch fi46 {
                            case 0:
                                var v49 float64 = 1e10
                                _ = v49
                            case 9223372036854775807:
                                if (1 < genFn1(255, 921, 32767)) {
                                    v50 := GenS7{F0: 32768, F1: 2147483648, F2: (fi47 + fi47)}
                                    v1 = append(v1, "hello")
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v50
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v51 uint64 = 545
                                    _ = v51
                                }
                            default:
                                if (!(1 > 1)) {
                                    fmt.Println(v3, fi47, v3)
                                    if v45 {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if (4294967295 >= 420) {
                                                var v52 int = 4294967295
                                                var v53 string = ("\n\t" + "")
                                                v54 := GenS8{F0: 65535, F1: (v52 & (fi46 | fi46)), F2: true}
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                        case "a":
                                            v1 = append(v1, "hello")
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            var v55 int8 = -1
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v45, fi46, fi47)
                                            v3 = -9223372036854775808
                                            v56 := GenS9{F0: (v3 * v3), F1: 536, F2: -98}
                                            _ = v55
                                            _ = v56
                                        case "a":
                                            var v57 uint16 = 0
                                            _ = v57
                                        }
                                        v45 = ((fi46 ^ 128) < (-(fi47)))
                                        v3 = (-(v3))
                                        v45 = true
                                    }
                                }
                            }
                            _ = len(v1)
                            _ = v48
                        } else {
                            switch ((("hello" + (((("" + ("hello" + ("" + (("hello" + "foo bar") + (("a" + "hello") + "hello"))))) + "a") + "\n\t") + "")) + "hello") + "") {
                            case "a":
                                fmt.Println(fi46, v3, v45)
                            case "foo bar":
                                switch ("hello" + "") {
                                case "":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "hello":
                                    var v58 int16 = (v2 * (v2 - (v2 - (v2 >> ((2) & 63)))))
                                    _ = v58
                                case "a":
                                    var v59 int = (fi46 ^ fi46)
                                    _ = v59
                                }
                            }
                            var v60 uint64 = 1
                            _ = v60
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi47
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = ("hello" + ("foo bar" + "a"))
                    }
                }
            }
            v61 := GenS7{F0: int(v3), F1: 65535, F2: (fi46 & 32768)}
            _ = fi46
            _ = v61
        }
    }
    v45 = (54 != 255)
    var v62 int64 = (v3 / (((v3 >> ((65535) & 63))) | 1))
    v63 := GenS8{F0: 2, F1: 164, F2: (1 >= 42)}
    v62 -= v3
    switch "a" {
    case "a":
        v45 = true
    case "hello":
        v2 = int16(v3)
    default:
        if ((((-1 == -1) && v45) && v45) || (895 > 2)) {
            v45 = v45
            _ = len(v1)
            for fi64 := 0; fi64 < 11; fi64++ {
                v45 = true
                _ = fi64
            }
        } else {
            var v65 float32 = float32(((float64(0.0) * float64(((-2.25 - 0.1) + -2.25))) / float64(1.5)))
            var v66 uint32 = 4294967295
            _ = v65
            _ = v66
        }
    }
    v67 := GenS5{F0: 2147483647, F1: 65535, F2: (-1 <= 1)}
    var v68 bool = v45
    fmt.Println(v68, v2)
    switch "foo bar" {
    case "foo bar":
        var v69 uint64 = 849
        _ = v69
    case "":
        if (2 <= 0) {
            switch -1 {
            case 127:
                for fi70 := 0; fi70 < 11; fi70++ {
                    var v71 int16 = (v2 << ((18446744073709551615) & 63))
                    v68 = true
                    var v72 int64 = v3
                    _ = fi70
                    _ = v71
                    _ = v72
                }
            default:
                v73 := GenS3{F0: v3, F1: 952}
                _ = v73
            }
            if (0 <= 482) {
                var v74 uint64 = genFn1(2, -2147483648, 1)
                _ = v74
            }
        }
    case "\n\t":
        _ = len(v1)
    default:
        v62 = v62
    }
    var v75 bool = (!(317 != 0))
    var v76 int8 = 108
    v77 := GenS10{F0: 255, F1: v75}
    if v75 {
        for fi78 := 0; fi78 < 8; fi78++ {
            for fi79 := 0; fi79 < 10; fi79++ {
                v80 := GenS11{F0: 828}
                v62 = (v62 * v3)
                _ = fi79
                _ = v80
            }
            _ = fi78
        }
    }
    var v81 float64 = float64((0.1 - 1e10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v62:", v62)
    fmt.Println("v68:", v68)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v62
    _ = v63
    _ = v67
    _ = v68
    _ = v75
    _ = v76
    _ = v77
    _ = v81
}
