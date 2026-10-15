package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int }
type GenS3 struct { F0 int; F1 int; F2 uint }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint32; F1 int; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 int }
func genFn0(p0 int64, p1 int64) int32 {
    var v4 uint16 = 1
    p0 -= 1
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: v4, F1: (p0 & -1)}
        v7 := GenS1{F0: v4, F1: p0}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{127, 32767, int(p1), 267}
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 115
}
func genFn1(p0 int8) int8 {
    p0 += p0
    switch 922 {
    case 4294967295:
        if ((-1 <= 1) && true) {
            v9 := GenS2{F0: 155, F1: 42, F2: -522}
            switch 128 {
            case 128:
                var v10 string = "hello"
                _ = v10
            case 527:
                var v11 int64 = 779
                _ = v11
            case 1:
                switch "foo bar" {
                case "":
                    p0 = p0
                case "foo bar":
                    var v12 int8 = -114
                    _ = v12
                }
            }
            _ = v9
        } else {
            var v13 uint16 = 989
            v14 := GenS3{F0: -9223372036854775808, F1: 127, F2: 65535}
            v15 := []int{127, 4294967295}
            _ = v13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    return p0
}

func main() {
    v1 := []string{"hello", "", ("hello" + "\n\t")}
    var v2 uint64 = 1
    var v3 string = ((("" + (("" + "a") + "hello")) + "") + "")
    if false {
        var v16 int8 = -1
        fmt.Println(v2, v2, v3)
        _ = v16
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            var v19 bool = (false || (0.0 > 3.14159))
            _ = fi18
            _ = v19
        }
        if (1000 == -565) {
            fmt.Println(v2, fi17, v2)
        } else {
            if ((42 > -460) || false) {
                v20 := GenS3{F0: int(v2), F1: (fi17 - (^(fi17))), F2: 65535}
                for fi21 := 0; fi21 < 14; fi21++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v22 := GenS4{F0: 1, F1: 1}
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(fi17, fi17)
            if false {
                v24 := GenS3{F0: -2147483648, F1: (fi17 >> ((2) & 63)), F2: 255}
                v3 = (v3 + (v3 + (v3 + "\n\t")))
                _ = v24
            } else {
                v2 = (v2 * v2)
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            _ = fi23
        }
        _ = fi17
    }
    switch v3 {
    case "foo bar":
        _ = len(v1)
    case "":
        fmt.Println(v2, v2)
    case "\n\t":
        switch ((("foo bar" + (v3 + (("" + v3) + ("a" + (v3 + v3))))) + v3) + v3) {
        case "a":
            v26 := GenS5{F0: 2, F1: 4294967295, F2: -532}
            _ = v26
        case "foo bar":
            v3 = v3
        case "\n\t":
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 uint8 = 255
                _ = fi27
                _ = v28
            }
        }
    default:
        v3 = v3
    }
    switch 32767 {
    case 771:
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 float32 = 3.14159
            _ = fi29
            _ = v30
        }
    case -913:
        v2 -= (^(v2))
    }
    var v31 int32 = genFn0(9223372036854775807, 0)
    switch 917 {
    case 312:
        if ((!(-9223372036854775808 > 1)) || false) {
            v32 := GenS6{F0: v31}
            for fi33 := 0; fi33 < 19; fi33++ {
                v31 = (v31 << ((65535) & 63))
                var v34 bool = (10 <= 0)
                for fi35 := 0; fi35 < 7; fi35++ {
                    fmt.Println(v31, v34)
                    if v34 {
                        var v36 uint8 = 1
                        switch (v3 + "foo bar") {
                        case "hello":
                            var v37 int32 = (v31 << ((uint(v36)) & 63))
                            _ = v37
                        case "\n\t":
                            if v34 {
                                v2 = (v2 % (((v2 ^ v2)) | 1))
                            } else {
                                v1 = append(v1, v3)
                            }
                        case "a":
                            v34 = v34
                        }
                        _ = v36
                    } else {
                        if len(v1) > 0 {
                            v1[0] = (v3 + (((("hello" + v3) + (v3 + "\n\t")) + v3) + v3))
                        }
                        v31 -= (v31 >> ((631) & 63))
                        var v38 bool = v34
                        _ = v38
                    }
                    if true {
                        if true {
                            v39 := GenS7{F0: fi33}
                            _ = v39
                        }
                        var v40 int = (fi35 | fi33)
                        var v41 float32 = float32(3.14159)
                        _ = v40
                        _ = v41
                    }
                    _ = fi35
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
    case 127:
        var v42 float32 = 3.14159
        _ = v42
    case 4294967295:
        var v43 int = 706
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 int = fi44
        switch v3 {
        case "":
            for fi46 := 0; fi46 < 10; fi46++ {
                switch v3 {
                case "hello":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "a":
                    var v47 int64 = 0
                    switch ((v3 + v3) + "foo bar") {
                    case "":
                        if false {
                            switch -9223372036854775808 {
                            case -810:
                                v48 := GenS1{F0: 255, F1: v47}
                                _ = v48
                            case -1:
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                                fmt.Println(v2)
                                v49 := GenS8{F0: v31}
                                _ = v49
                            default:
                                fmt.Println(v2)
                            }
                            v45 += (fi44 + fi46)
                        }
                    case "foo bar":
                        v47 -= (v47 - 1)
                    }
                    var v50 uint = 265
                    var v51 uint16 = 65535
                    _ = v47
                    _ = v50
                    _ = v51
                case "\n\t":
                    v45 -= (fi44 % (((^(fi46))) | 1))
                }
                _ = fi46
            }
        default:
            for fi52 := 0; fi52 < 14; fi52++ {
                v53 := GenS8{F0: (v31 | v31)}
                _ = fi52
                _ = v53
            }
        }
        v54 := GenS9{F0: 75, F1: 18446744073709551615}
        _ = fi44
        _ = v45
        _ = v54
    }
    if (-128 == genFn1(genFn1(-1))) {
        var v55 int32 = v31
        _ = v55
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        for fi57 := 0; fi57 < 14; fi57++ {
            v58 := GenS3{F0: fi57, F1: -1, F2: 326}
            _ = fi57
            _ = v58
        }
        v2 = v2
        switch fi56 {
        case 32768:
            fmt.Println(v3, fi56)
        case 4294967295:
            fmt.Println(v3, v3, v2)
        default:
            var v59 int8 = 100
            _ = v59
        }
        _ = fi56
    }
    v60 := GenS9{F0: 0, F1: 255}
    switch v3 {
    case "\n\t":
        var v61 float64 = 3.14159
        _ = v61
    }
    switch ((v3 + (("\n\t" + v3) + v3)) + v3) {
    case "foo bar":
        _ = len(v1)
    case "\n\t":
        v62 := GenS10{F0: 128}
        _ = v62
    default:
        for fi63 := 0; fi63 < 14; fi63++ {
            if (!(false || (32767 >= -1))) {
                var v64 int = fi63
                fmt.Println(v3, fi63)
                _ = v64
            } else {
                fmt.Println(fi63, v31)
                var v65 bool = (255 == 633)
                v66 := GenS3{F0: 32767, F1: fi63, F2: 539}
                _ = v65
                _ = v66
            }
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v67 := GenS10{F0: -1}
                fmt.Println(fi63, v2, v31)
                _ = v67
            }
            _ = len(v1)
            _ = fi63
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v60
}
