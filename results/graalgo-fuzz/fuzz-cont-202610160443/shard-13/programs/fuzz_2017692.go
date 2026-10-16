package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := []int{2147483647, 450, 32768, 698}
    v2 := make(map[int]uint32)
    var v3 uint = 255
    v4 := make(map[int]string)
    var v5 int8 = -128
    var v6 int16 = 486
    var v7 int = 42
    v3 += v3
    for fi8 := 0; fi8 < 12; fi8++ {
        v3 = v3
        var v9 int16 = v6
        var v10 int = (fi8 - (v7 << (((v3 << ((v3) & 63))) & 63)))
        _ = fi8
        _ = v10
        _ = v9
    }
    switch v7 {
    case 127:
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                switch "a" {
                case "a":
                    v13 := GenS1{F0: "a", F1: fi11, F2: -317}
                    _ = v13
                default:
                    var v14 bool = true
                    _ = v14
                }
                v6 -= v6
                _ = fi12
            }
            v3 -= v3
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int = fi15
                fmt.Println(v16)
                _ = fi15
                _ = v16
            }
            _ = fi11
        }
    case 42:
        var v17 int32 = -2147483648
        _ = v17
    case 32768:
        if len(v1) > 0 {
            v1[0] = 0
        }
    }
    v18 := GenS2{F0: (v3 + (^(v3))), F1: -918}
    if false {
        for fi19 := 0; fi19 < 18; fi19++ {
            v20 := GenS2{F0: 757, F1: -9223372036854775808}
            _ = fi19
            _ = v20
        }
        var v21 bool = false
        _ = v21
    } else {
        v22 := GenS3{F0: (v7 % ((v7) | 1)), F1: 48}
        var v23 uint16 = 0
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v5 = (v5 | v5)
        _ = fi24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 4; fi26++ {
            var v27 bool = (!((v6 ^ 1000) <= v6))
            v6 = -1
            _ = fi26
            _ = v27
        }
        v28 := GenS2{F0: v3, F1: 9223372036854775807}
        fmt.Println(v7, v5)
        _ = fi25
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 5; fi29++ {
            v3 += v3
            _ = fi29
        }
    } else {
        var v30 uint = (v3 << ((v3) & 63))
        for fi31 := 0; fi31 < 17; fi31++ {
            v3 += v3
            for fi32 := 0; fi32 < 20; fi32++ {
                if true {
                    var v33 int8 = 100
                    _ = v33
                } else {
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v34 := GenS3{F0: v7, F1: 42}
                        _ = v34
                    default:
                        var v35 uint16 = 255
                        _ = v35
                    }
                    v6 = (v6 - -1)
                }
                v30 -= v30
                switch v7 {
                case -2147483648:
                    fmt.Println(v6, fi31)
                    if (5 != 42) {
                        var v36 float32 = (-2.25 + 1.5)
                        v36 = (v36 + ((v36 + v36) - v36))
                        if ("\n\t" < (("\n\t" + "hello") + "")) {
                            if (("\n\t" + "") >= ("hello" + ("hello" + ("" + ("\n\t" + (("foo bar" + "hello") + ("a" + "hello"))))))) {
                                switch 32768 {
                                case 127:
                                    var v37 int = (fi32 | (v7 % ((824) | 1)))
                                    v5 += (v5 * (v5 << ((v30) & 63)))
                                    v38 := GenS4{F0: 1, F1: -1, F2: 2147483648}
                                    _ = v37
                                    _ = v38
                                case 32767:
                                    v7 = 730
                                }
                                if ((v3 - v3) != 255) {
                                    fmt.Println(fi32, v5, v3)
                                    switch ("" + "foo bar") {
                                    case "foo bar":
                                        var v39 uint64 = 1
                                        _ = v39
                                    case "\n\t":
                                        if (true && (!(42 == -942))) {
                                            v40 := GenS5{F0: 255, F1: 65535}
                                            v2[fi32] = 118
                                            _ = v40
                                        } else {
                                            v41 := GenS6{F0: (^(v3)), F1: 0}
                                            _ = v41
                                        }
                                        v1 = append(v1, v7)
                                        switch ("" + "hello") {
                                        case "\n\t":
                                            v4[1] = ((((("foo bar" + "hello") + ("a" + (((("hello" + (("" + ("hello" + ("foo bar" + ("" + "foo bar")))) + (("\n\t" + "hello") + (("foo bar" + (("hello" + "") + (("foo bar" + "foo bar") + "hello"))) + "")))) + "") + "") + (("foo bar" + ("a" + ("" + ""))) + "foo bar")))) + (("foo bar" + "a") + "hello")) + ("\n\t" + ("foo bar" + ""))) + "foo bar")
                                        case "":
                                            if (128 != 19) {
                                                if true {
                                                    switch "" {
                                                    case "":
                                                        if true {
                                                            var v42 uint16 = 0
                                                            _ = v42
                                                        } else {
                                                            v43 := GenS3{F0: int(v6), F1: 9}
                                                            var v44 string = (("a" + "") + "\n\t")
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    case "hello":
                                                        v30 = 2
                                                    case "foo bar":
                                                        fmt.Println(v30, v6)
                                                        _ = len(v1)
                                                        var v45 uint = 255
                                                        _ = v45
                                                    default:
                                                        fmt.Println(v36)
                                                    }
                                                    delete(v2, 127)
                                                }
                                                if true {
                                                    _ = len(v4)
                                                }
                                            } else {
                                                v46 := GenS4{F0: 127, F1: -1, F2: 49}
                                                v30 += 0
                                                if (0.1 < 0.0) {
                                                    v47 := GenS5{F0: 98, F1: 0}
                                                    _ = v47
                                                }
                                                _ = v46
                                            }
                                        default:
                                            if true {
                                                v36 = v36
                                                if ((false && false) || (((v36 / 1.0) - (v36 * ((0.1 * (v36 - v36)) + (((v36 - v36) - (1.5 / v36)) * (float32(-2.25) + (1e10 * (v36 * (float32(1.5) + v36)))))))) >= v36)) {
                                                    var v48 int = 962
                                                    v3 += 42
                                                    _ = v48
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = (^(fi31))
                                                    }
                                                    v30 += 0
                                                }
                                            }
                                        }
                                    case "":
                                        switch ("" + ("\n\t" + "a")) {
                                        case "":
                                            v49 := GenS3{F0: fi31, F1: 2}
                                            v36 -= (1.5 / v36)
                                            _ = v49
                                        }
                                    default:
                                        v50 := GenS3{F0: int(v3), F1: 24}
                                        _ = v50
                                    }
                                    var v51 float32 = (float32(-2.25) * ((float32(float64(float32(1e10))) / v36) - ((v36 / float32(((3.14159 + 0.0) + 3.14159))) / ((v36 * ((((-2.25 - -2.25) + ((float32(-2.25) / (v36 - ((((v36 + v36) - (-2.25 + (float32(3.14159) * (v36 - v36)))) - v36) - v36))) / float32(((3.14159 + 1.5) - ((0.0 * 3.14159) * 3.14159))))) * 1e10) + v36)) - (v36 + (0.0 * 1.5))))))
                                    _ = v51
                                }
                                v3 = (v3 ^ (v30 | (v30 % ((v30) | 1))))
                            }
                            v1 = append(v1, v7)
                            var v52 uint = v30
                            _ = v52
                        } else {
                            _ = len(v4)
                            fmt.Println(fi32)
                            var v53 int = v7
                            _ = v53
                        }
                        var v54 uint64 = 1
                        _ = v36
                        _ = v54
                    }
                }
                var v55 float32 = 3.14159
                if (true && true) {
                    v55 += 0.1
                }
                _ = fi32
                _ = v55
            }
            _ = fi31
        }
        switch v7 {
        case -1:
            v56 := v2[v7]
            _ = v56
        case 901:
            v5 += (v5 >> (((v3 << ((v30) & 63))) & 63))
        case 127:
            v30 = (v30 / ((42) | 1))
        default:
            if false {
                if ((v6 <= (v6 / ((-1) | 1))) && (!(0 <= 42))) {
                    var v57 string = "foo bar"
                    _ = v57
                }
                var v58 int8 = v5
                _ = v58
            } else {
                fmt.Println(v5)
            }
        }
        _ = v30
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        v4[fi59] = ("hello" + ("foo bar" + "foo bar"))
        _ = fi59
    }
    if true {
        var v60 uint64 = 495
        v61 := GenS6{F0: (v3 - (v3 * (^(v3)))), F1: 1}
        if ((v3 < 1) || (!((!(!(true && true))) && false))) {
            if true {
                for fi62 := 0; fi62 < 5; fi62++ {
                    v63 := GenS5{F0: 127, F1: 987}
                    switch ((("hello" + "\n\t") + "\n\t") + ((("" + ("\n\t" + "a")) + "\n\t") + "hello")) {
                    case "a":
                        fmt.Println(v6, v6)
                    }
                    switch ("" + "foo bar") {
                    case "hello":
                        for fi64 := 0; fi64 < 16; fi64++ {
                            var v65 int = (fi64 ^ (v7 >> ((v3) & 63)))
                            if false {
                                v66 := GenS1{F0: ("a" + "hello"), F1: 128, F2: fi62}
                                v6 += (v6 % ((-862) | 1))
                                if false {
                                    v3 += v3
                                    v1 = append(v1, v7)
                                    fmt.Println(v6, v6, v60)
                                } else {
                                    fmt.Println(v6, fi62, fi64)
                                    switch "foo bar" {
                                    case "\n\t":
                                        if ((((0.0 / (float64(3.14159) / 1.5)) / (1e10 * ((((-2.25 + ((-2.25 + (0.1 - 0.1)) / 0.1)) + float64((0.0 - 3.14159))) - (((((1.5 - (0.1 * 0.1)) / 3.14159) * (((3.14159 / 1.0) + 1.5) - ((1.5 - (0.1 * -2.25)) / 0.1))) - ((1e10 / 1.0) * (3.14159 - -2.25))) * -2.25)) * 1e10))) * (-2.25 - float64(1e10))) != 0.0) {
                                            var v67 uint = 255
                                            var v68 int64 = -788
                                            _ = v67
                                            _ = v68
                                        }
                                    default:
                                        var v69 uint32 = 2147483648
                                        var v70 int64 = 1
                                        var v71 uint64 = (v60 ^ (v60 * (v60 / ((v60) | 1))))
                                        var v72 int16 = (v6 + 1000)
                                        if true {
                                            switch (((("\n\t" + ("" + ("\n\t" + (("hello" + "foo bar") + ("a" + "foo bar"))))) + "foo bar") + "\n\t") + "hello") {
                                            case "hello":
                                                var v73 int = fi62
                                                _ = v73
                                            case "":
                                                fmt.Println(v7, v65, v3)
                                            case "foo bar":
                                                v74 := GenS7{F0: (v5 ^ v5)}
                                                _ = v74
                                            }
                                            v75 := GenS3{F0: 2147483647, F1: 255}
                                            _ = v75
                                        } else {
                                            var v76 float64 = 3.14159
                                            switch (fi64 & (v7 % (((v7 & int(v69))) | 1))) {
                                            case 226:
                                                var v77 uint32 = 2147483648
                                                _ = v77
                                            default:
                                                fmt.Println(v71)
                                                v6 = -705
                                                var v78 float64 = ((v76 / 1e10) + float64(-2.25))
                                                _ = v78
                                            }
                                            if (((v6 ^ (-(v6))) > (-(v72))) || (2 < v69)) {
                                                delete(v2, v7)
                                                var v79 float64 = v76
                                                fmt.Println(v5, v7, v71)
                                                _ = v79
                                            }
                                            _ = v76
                                        }
                                        v60 = v71
                                        v80 := v2[(fi64 + (fi64 * fi64))]
                                        var v81 bool = true
                                        var v82 bool = (v81 || v81)
                                        var v83 uint32 = 1
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    }
                                    _ = len(v2)
                                }
                                _ = v66
                            } else {
                                var v84 string = ("foo bar" + "a")
                                _ = v84
                            }
                            switch -9223372036854775808 {
                            case 2147483647:
                                v65 = (v7 % ((fi64) | 1))
                            case 0:
                                v85 := GenS6{F0: v3, F1: 0}
                                _ = v85
                            }
                            var v86 bool = true
                            _ = fi64
                            _ = v65
                            _ = v86
                        }
                    }
                    _ = fi62
                    _ = v63
                }
                if false {
                    v3 -= 4294967295
                    var v87 int32 = -2147483648
                    v5 = (v5 + -1)
                    _ = v87
                } else {
                    v5 += v5
                }
                var v88 string = "a"
                _ = v88
            }
            fmt.Println(v5, v6, v7)
        }
        _ = v60
        _ = v61
    } else {
        if (2147483647 < -689) {
            fmt.Println(v5, v5, v3)
            var v89 bool = false
            _ = v89
        } else {
            v90 := GenS1{F0: ("\n\t" + ""), F1: v7, F2: v7}
            var v91 uint8 = 1
            switch "a" {
            case "a":
                v2[v7] = 42
            case "foo bar":
                v92 := v2[(v7 << ((65535) & 63))]
                _ = v92
            default:
                v2[v7] = uint32(v91)
            }
            _ = v90
            _ = v91
        }
        for fi93 := 0; fi93 < 7; fi93++ {
            switch "" {
            case "a":
                if (!false) {
                    for fi94 := 0; fi94 < 11; fi94++ {
                        fmt.Println(fi94, fi94, fi93)
                        v1 = append(v1, (fi93 % (((^(v7))) | 1)))
                        _ = fi94
                    }
                    v7 -= fi93
                }
            case "":
                switch (fi93 << ((v3) & 63)) {
                case 2147483647:
                    v6 -= (-(v6))
                default:
                    v1 = append(v1, -257)
                }
            case "hello":
                for fi95 := 0; fi95 < 15; fi95++ {
                    var v96 int8 = v5
                    v96 += v5
                    fmt.Println(fi93, v3, v3)
                    if false {
                        v96 = v96
                        v3 = 1
                        var v97 float32 = 0.0
                        _ = v97
                    } else {
                        var v98 int8 = 100
                        fmt.Println(v98)
                        fmt.Println(v5, v3, v7)
                        v3 = 559
                        _ = v98
                    }
                    _ = fi95
                    _ = v96
                }
            }
            _ = fi93
        }
    }
    v99 := GenS8{F0: 1}
    v100 := v2[32768]
    switch "a" {
    case "hello":
        v101 := GenS7{F0: v5}
        _ = v101
    case "a":
        v102 := GenS6{F0: v3, F1: 837}
        _ = v102
    default:
        v3 -= (v3 >> (((v3 - 141)) & 63))
    }
    var v103 string = ""
    var v104 uint8 = 93
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v103:", v103)
    fmt.Println("v104:", v104)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v100
    _ = v103
    _ = v104
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v99
}
