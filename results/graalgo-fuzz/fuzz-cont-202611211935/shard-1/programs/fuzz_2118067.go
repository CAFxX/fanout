package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }

func main() {
    var v1 int8 = 100
    var v2 int16 = -957
    v3 := make(map[int]uint32)
    v4 := make(map[int]bool)
    v5 := GenS1{F0: "a", F1: 2147483647, F2: 721}
    switch 127 {
    case -9223372036854775808:
        v2 += v2
    default:
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 float32 = 0.1
                v9 := []int{fi7, -1}
                var v10 int64 = 0
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi6
        }
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 string = "hello"
            fmt.Println(v2, fi12, v2)
            fmt.Println(v2, fi11)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v14 := []int{-343, 1, 32767}
    var v15 int64 = 829
    if (2 <= 786) {
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                fmt.Println(fi16, v1)
                fmt.Println(v2, fi16, v1)
                _ = fi17
            }
            v1 = (^(v1))
            v14 = append(v14, (^(fi16)))
            _ = fi16
        }
    } else {
        v1 = int8(v2)
    }
    if len(v14) > 0 {
        v14[0] = 233
    }
    switch 128 {
    case 1:
        switch -9223372036854775808 {
        case 32767:
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint16 = 0
                delete(v4, fi18)
                _ = fi18
                _ = v19
            }
        default:
            v1 += (v1 / (((v1 ^ v1)) | 1))
        }
    default:
        for fi20 := 0; fi20 < 4; fi20++ {
            for fi21 := 0; fi21 < 19; fi21++ {
                if true {
                    fmt.Println(v1, v1)
                }
                var v22 uint8 = 128
                switch fi20 {
                case 128:
                    if len(v14) > 0 {
                        v14[0] = 32767
                    }
                case -1:
                    fmt.Println(v2, fi21, fi21)
                default:
                    var v23 uint64 = 1
                    _ = v23
                }
                _ = fi21
                _ = v22
            }
            if true {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                v15 = (-(v15))
            } else {
                if false {
                    if false {
                        switch "hello" {
                        case "":
                            v2 = 42
                        case "hello":
                            if false {
                                v2 += v2
                            }
                        case "a":
                            fmt.Println(v15, fi20, v15)
                        default:
                            v24 := GenS2{F0: "a", F1: 0, F2: 255}
                            _ = v24
                        }
                    } else {
                        switch -1 {
                        case 84:
                            for fi25 := 0; fi25 < 7; fi25++ {
                                if true {
                                    var v26 int16 = (^(v2))
                                    _ = v26
                                }
                                if (-1 <= 42) {
                                    v1 += (v1 & (^(v1)))
                                    v15 -= v15
                                }
                                switch "a" {
                                case "a":
                                    fmt.Println(fi20, v15)
                                default:
                                    switch (fi25 >> ((745) & 63)) {
                                    case 4294967295:
                                        switch "" {
                                        case "hello":
                                            v27 := GenS1{F0: ("" + "\n\t"), F1: 2147483647, F2: 65535}
                                            _ = v27
                                        case "\n\t":
                                            v1 = 100
                                        case "a":
                                            if true {
                                                if (42 > 42) {
                                                    if ((!(127 == 127)) && false) {
                                                        var v28 int16 = v2
                                                        var v29 uint8 = 46
                                                        var v30 uint64 = 18446744073709551615
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                    switch (-(fi25)) {
                                                    case 288:
                                                        _ = len(v14)
                                                    case 0:
                                                        _ = len(v3)
                                                        v31 := GenS1{F0: ("\n\t" + ""), F1: 2147483647, F2: 65535}
                                                        _ = v31
                                                    case -2147483648:
                                                        fmt.Println(v15, v1)
                                                    default:
                                                        if ((float32(1.5) / 3.14159) != ((3.14159 + 3.14159) * (3.14159 - 0.1))) {
                                                            fmt.Println(fi25, v1, fi25)
                                                            v32 := GenS2{F0: ("" + (("" + "foo bar") + "")), F1: -32768, F2: 65535}
                                                            _ = v32
                                                        } else {
                                                            if len(v14) >= 2 {
                                                                v14 = v14[1:len(v14)]
                                                            }
                                                        }
                                                    }
                                                }
                                                v2 = 1000
                                                var v33 string = ("" + "\n\t")
                                                _ = v33
                                            } else {
                                                fmt.Println(v2)
                                            }
                                        default:
                                            v14 = append(v14, (^(fi25)))
                                        }
                                    default:
                                        fmt.Println(fi20, fi25)
                                    }
                                }
                                _ = fi25
                            }
                        case 0:
                            for fi34 := 0; fi34 < 13; fi34++ {
                                if len(v14) > 0 {
                                    v14[0] = 32768
                                }
                                var v35 uint8 = 128
                                v2 += v2
                                _ = fi34
                                _ = v35
                            }
                        default:
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        }
                        v2 = v2
                        fmt.Println(v15, fi20)
                    }
                } else {
                    var v36 int64 = (v15 - (^(v15)))
                    for fi37 := 0; fi37 < 19; fi37++ {
                        if false {
                            _ = len(v4)
                            var v38 int32 = -2147483648
                            _ = v38
                        } else {
                            var v39 int = fi37
                            if len(v14) > 0 {
                                v14[0] = v39
                            }
                            switch "" {
                            case "":
                                v2 += v2
                            case "a":
                                delete(v4, (fi37 | (fi37 / ((fi20) | 1))))
                            }
                            _ = v39
                        }
                        _ = fi37
                    }
                    if true {
                        var v40 int = 4294967295
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 19; fi41++ {
                            v2 -= 1000
                            if (1e10 >= 1e10) {
                                delete(v4, int(v2))
                                if len(v14) > 0 {
                                    v14[0] = fi41
                                }
                                v15 = v15
                            } else {
                                v42 := GenS3{F0: 65535}
                                var v43 int32 = 0
                                var v44 float64 = float64((3.14159 + (3.14159 + 3.14159)))
                                _ = v42
                                _ = v43
                                _ = v44
                            }
                            _ = fi41
                        }
                        var v45 string = ("a" + ((("foo bar" + "\n\t") + "foo bar") + "a"))
                        switch -1 {
                        case -9:
                            fmt.Println(v15, v2, v1)
                        case 943:
                            var v46 int8 = v1
                            _ = v46
                        case -786:
                            switch "" {
                            case "\n\t":
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    fmt.Println(v15, v2)
                                    if false {
                                        switch "a" {
                                        case "a":
                                            switch v45 {
                                            case "a":
                                                if false {
                                                    v36 = (v15 >> ((65535) & 63))
                                                    v14 = append(v14, (fi20 & 9223372036854775807))
                                                }
                                            case "hello":
                                                if (42 > 2) {
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    fmt.Println(v36, v15)
                                                    if len(v14) > 0 {
                                                        v14[0] = (fi20 & (fi20 % (((fi20 + (fi20 % ((10) | 1)))) | 1)))
                                                    }
                                                }
                                            case "\n\t":
                                                v48 := v3[(fi20 + fi20)]
                                                _ = v48
                                            default:
                                                if (false && true) {
                                                    if (-1 <= 2147483647) {
                                                        v2 = (v2 - v2)
                                                        v1 += -128
                                                        var v49 int64 = (v36 & (v15 * v36))
                                                        var v50 float64 = 3.14159
                                                        fmt.Println(v15, v1)
                                                        if true {
                                                            v1 -= (v1 / ((v1) | 1))
                                                            switch v45 {
                                                            case "a":
                                                                var v51 int32 = 42
                                                                _ = v51
                                                            case "\n\t":
                                                                v52 := GenS4{F0: 51, F1: 2147483647}
                                                                _ = v52
                                                            case "foo bar":
                                                                v50 += v50
                                                            default:
                                                                v50 = v50
                                                            }
                                                        } else {
                                                            v50 = (v50 * float64((0.1 - 3.14159)))
                                                            fmt.Println(v45)
                                                            if false {
                                                                fmt.Println(v36, v45, fi47)
                                                            }
                                                        }
                                                        _ = v49
                                                        _ = v50
                                                    } else {
                                                        var v53 int32 = -2147483648
                                                        v54 := GenS1{F0: ("foo bar" + ((v45 + "a") + "foo bar")), F1: 1, F2: 255}
                                                        v2 += (v2 >> ((1) & 63))
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                }
                                            }
                                        case "hello":
                                            var v55 float64 = -2.25
                                            _ = v55
                                        case "":
                                            v56 := GenS5{F0: 252}
                                            _ = v56
                                        default:
                                            v1 = v1
                                        }
                                        var v57 string = (("" + v45) + "hello")
                                        if false {
                                            v58 := v3[(fi47 * fi20)]
                                            if false {
                                                var v59 int16 = v2
                                                if false {
                                                    v2 = v2
                                                } else {
                                                    v57 = v45
                                                    var v60 int16 = 141
                                                    _ = v60
                                                }
                                                v45 = (v57 + v57)
                                                _ = v59
                                            }
                                            v61 := GenS1{F0: v45, F1: 1, F2: 42}
                                            _ = v58
                                            _ = v61
                                        } else {
                                            v62 := GenS5{F0: 18446744073709551615}
                                            _ = v62
                                        }
                                        fmt.Println(fi47, v36, v1)
                                        fmt.Println(v45)
                                        var v63 uint64 = 18446744073709551615
                                        var v64 uint32 = 757
                                        _ = v57
                                        _ = v63
                                        _ = v64
                                    } else {
                                        switch v45 {
                                        case "":
                                            v45 = ("foo bar" + v45)
                                            var v65 int32 = 42
                                            _ = v65
                                        case "foo bar":
                                            v36 += (v36 ^ (-(v15)))
                                        }
                                    }
                                    _ = fi47
                                }
                            case "":
                                if (162 == 1) {
                                    v66 := GenS6{F0: 56}
                                    v15 = -205
                                    fmt.Println(v1, v1)
                                    _ = v66
                                } else {
                                    if false {
                                        if (false || (!(768 < 18446744073709551615))) {
                                            v67 := GenS6{F0: 4294967295}
                                            _ = v67
                                        }
                                        fmt.Println(v15, v2, v2)
                                    }
                                    if len(v14) > 0 {
                                        v14[0] = fi20
                                    }
                                }
                            case "foo bar":
                                v45 = (("foo bar" + "a") + "hello")
                            default:
                                v2 -= v2
                            }
                        default:
                            if (-1 == 719) {
                                fmt.Println(fi20, v15)
                                var v68 uint8 = 127
                                _ = len(v3)
                                _ = v68
                            }
                        }
                        _ = v45
                    }
                    _ = v36
                }
                _ = len(v3)
                if true {
                    var v69 int = 1
                    _ = v69
                } else {
                    v2 -= int16(v1)
                    for fi70 := 0; fi70 < 12; fi70++ {
                        v71 := GenS7{F0: 628}
                        var v72 uint16 = 255
                        switch (fi70 | fi70) {
                        case 865:
                            v1 += 0
                        case -1:
                            _ = len(v14)
                        default:
                            v3[fi20] = 876
                        }
                        v73 := GenS1{F0: "foo bar", F1: 0, F2: 890}
                        _ = fi70
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                    _ = len(v3)
                }
            }
            for fi74 := 0; fi74 < 3; fi74++ {
                v2 += v2
                _ = fi74
            }
            _ = fi20
        }
    }
    if true {
        v75 := v3[1]
        if len(v14) > 0 {
            v14[0] = 1
        }
        if (("" + ("\n\t" + "\n\t")) != "") {
            v75 = (v75 >> ((1) & 63))
            switch 1 {
            case 1:
                v1 = v1
            case -884:
                for fi76 := 0; fi76 < 2; fi76++ {
                    v2 -= (v2 - 128)
                    _ = fi76
                }
            case 32767:
                var v77 float32 = -2.25
                _ = v77
            default:
                if true {
                    v78 := GenS1{F0: "", F1: 968, F2: 255}
                    _ = v78
                } else {
                    v1 -= 72
                    var v79 uint64 = 42
                    _ = v79
                }
            }
        }
        _ = v75
    } else {
        v80 := GenS4{F0: 217, F1: 42}
        _ = v80
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
