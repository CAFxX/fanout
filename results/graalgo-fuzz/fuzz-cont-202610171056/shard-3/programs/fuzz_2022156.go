package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int; F2 int }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 uint64 = 1
    var v2 int64 = 438
    var v3 string = (("foo bar" + ("a" + ("a" + ("\n\t" + "hello")))) + "")
    v4 := make(map[uint32]bool)
    if true {
        v5 := []int{-705, 639}
        v6 := GenS1{F0: -128, F1: 242, F2: (((v3 + v3) + ("foo bar" + (v3 + v3))) > v3)}
        var v7 uint8 = 1
        _ = v5
        _ = v6
        _ = v7
    } else {
        var v8 float64 = (1.5 - 1e10)
        v3 = v3
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        if (-929 > 32767) {
            v3 = "a"
        } else {
            var v12 string = v3
            v13 := GenS1{F0: -128, F1: 1, F2: (!(false || (!(true || (!false)))))}
            var v14 uint64 = (v1 * 42)
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 string = "foo bar"
    fmt.Println(v2, v15, v1)
    if true {
        for fi16 := 0; fi16 < 3; fi16++ {
            if false {
                var v17 int = (-(fi16))
                _ = v17
            }
            if false {
                v18 := GenS1{F0: -31, F1: 42, F2: false}
                if (false && (1 != 0)) {
                    var v19 float64 = 0.1
                    _ = v19
                } else {
                    delete(v4, 21)
                }
                var v20 int = -2147483648
                _ = v18
                _ = v20
            }
            if false {
                if false {
                    if false {
                        if (!(!((-(fi16)) <= fi16))) {
                            for fi21 := 0; fi21 < 13; fi21++ {
                                if ((1 <= 0) && false) {
                                    v22 := GenS2{F0: 124}
                                    if ((fi21 != (-(fi16))) && (!false)) {
                                        if false {
                                            switch (v15 + v3) {
                                            case "a":
                                                if (!true) {
                                                    var v23 uint64 = 0
                                                    _ = v23
                                                }
                                            }
                                            v24 := []int{int(v2)}
                                            _ = v24
                                        } else {
                                            if true {
                                                if (!false) {
                                                    if false {
                                                        v1 = 0
                                                        v3 = (v3 + "hello")
                                                        v15 = (v15 + ((v3 + "foo bar") + ((v15 + "a") + ("foo bar" + "\n\t"))))
                                                    }
                                                    v1 = (v1 & 1)
                                                } else {
                                                    switch fi21 {
                                                    case -808:
                                                        v25 := GenS3{F0: v3, F1: v15}
                                                        _ = v25
                                                    case 42:
                                                        v26 := GenS3{F0: v15, F1: ""}
                                                        _ = v26
                                                    }
                                                    var v27 int16 = -321
                                                    v2 = (v2 ^ v2)
                                                    _ = v27
                                                }
                                            } else {
                                                var v28 int64 = (-(v2))
                                                fmt.Println(v1, v28, v28)
                                                v28 = 259
                                                var v29 int8 = 42
                                                v15 = (v15 + "")
                                                v30 := []int{fi16, (fi21 << ((4294967295) & 63))}
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            v2 = (v2 & 705)
                                            switch -300 {
                                            case 32768:
                                                var v31 uint16 = 1
                                                _ = v31
                                            }
                                            switch fi21 {
                                            case -2147483648:
                                                if (int64(fi21) == (v2 << ((0) & 63))) {
                                                    if (false || false) {
                                                        v32 := GenS1{F0: 0, F1: 255, F2: ((fi16 | (fi16 >> ((2) & 63))) != (fi21 ^ 320))}
                                                        if (!(0 < -840)) {
                                                            v33 := GenS4{F0: 1, F1: 1, F2: 722}
                                                            _ = v33
                                                        }
                                                        _ = v32
                                                    }
                                                    v34 := GenS2{F0: 71}
                                                    _ = v34
                                                } else {
                                                    var v35 int64 = (-(v2))
                                                    v3 = v3
                                                    _ = v35
                                                }
                                            case 317:
                                                if false {
                                                    v1 -= (v1 + v1)
                                                }
                                            }
                                        }
                                        var v36 int = 32768
                                        _ = v36
                                    }
                                    switch ((((v3 + v3) + (v3 + v15)) + "") + "") {
                                    case "a":
                                        var v37 string = v15
                                        _ = v37
                                    default:
                                        var v38 uint = 15
                                        switch 0 {
                                        case 9223372036854775807:
                                            var v39 uint64 = v1
                                            _ = v39
                                        default:
                                            v15 = (v3 + (("" + ("\n\t" + "foo bar")) + v3))
                                            v40 := []int{}
                                            _ = v40
                                        }
                                        var v41 int16 = 110
                                        _ = v38
                                        _ = v41
                                    }
                                    _ = v22
                                }
                                var v42 string = v15
                                v43 := GenS5{F0: -128, F1: 42}
                                _ = fi21
                                _ = v42
                                _ = v43
                            }
                            switch (v15 + "\n\t") {
                            case "\n\t":
                                switch "hello" {
                                case "foo bar":
                                    v44 := []int{fi16, 2147483647}
                                    _ = v44
                                }
                            case "a":
                                if (("\n\t" <= v15) || true) {
                                    for fi45 := 0; fi45 < 5; fi45++ {
                                        switch v15 {
                                        case "foo bar":
                                            switch fi16 {
                                            case -1:
                                                switch v15 {
                                                case "a":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    var v46 float64 = 0.0
                                                    _ = v46
                                                case "":
                                                    v47 := []int{}
                                                    _ = v47
                                                }
                                            default:
                                                if (0 < 42) {
                                                    if true {
                                                        v4[414] = (3.14159 == (3.14159 - 0.1))
                                                        fmt.Println(fi45, v2, v15)
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            if true {
                                                                if false {
                                                                    fmt.Println(v3, fi45)
                                                                    v48 := GenS6{F0: 0, F1: fi16, F2: fi16}
                                                                    _ = v48
                                                                }
                                                                var v49 int = 4294967295
                                                                _ = v49
                                                            } else {
                                                                switch (fi45 * fi45) {
                                                                case 32767:
                                                                    v50 := v4[65535]
                                                                    _ = v50
                                                                case 877:
                                                                    if true {
                                                                        var v51 int16 = 613
                                                                        v52 := []int{fi45}
                                                                        v1 = v1
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                default:
                                                                    var v53 int = (fi45 | 256)
                                                                    _ = v53
                                                                }
                                                                fmt.Println(v1, fi16, v1)
                                                                var v54 uint32 = 320
                                                                _ = v54
                                                            }
                                                            v55 := []int{}
                                                            var v56 uint = 778
                                                            _ = v55
                                                            _ = v56
                                                        case "foo bar":
                                                            var v57 int16 = 32767
                                                            _ = v57
                                                        }
                                                    }
                                                    v4[42] = ((v15 + v3) < v15)
                                                } else {
                                                    v4[2147483648] = ((0.1 / 1.5) >= 3.14159)
                                                    v3 = "hello"
                                                    fmt.Println(fi45, v3, v2)
                                                }
                                            }
                                        case "hello":
                                            fmt.Println(fi16)
                                        case "\n\t":
                                            fmt.Println(fi16, v15)
                                        default:
                                            var v58 uint32 = 42
                                            _ = v58
                                        }
                                        _ = fi45
                                    }
                                    fmt.Println(v2, v2)
                                }
                            case "foo bar":
                                var v59 int = (fi16 + (fi16 << ((1) & 63)))
                                _ = v59
                            }
                            v15 = v3
                        }
                        v3 = "\n\t"
                        v15 = v15
                    }
                    var v60 int8 = 127
                    _ = v60
                } else {
                    var v61 int16 = -1
                    _ = v61
                }
                for fi62 := 0; fi62 < 14; fi62++ {
                    fmt.Println(v3, v1)
                    var v63 uint32 = 534
                    if (true && true) {
                        var v64 float32 = (-2.25 - 0.0)
                        switch fi62 {
                        case -691:
                            switch v15 {
                            case "\n\t":
                                if (-170 < 0) {
                                    v64 += v64
                                }
                            }
                        default:
                            v63 = (v63 ^ v63)
                        }
                        switch (fi62 + 2147483647) {
                        case -126:
                            var v65 int8 = 42
                            _ = v65
                        default:
                            switch "a" {
                            case "":
                                _ = len(v4)
                            case "foo bar":
                                v66 := []int{}
                                _ = v66
                            }
                        }
                        _ = v64
                    }
                    v4[(v63 & (^(v63)))] = (1 > 18446744073709551615)
                    fmt.Println(v15)
                    _ = fi62
                    _ = v63
                }
                v3 = "\n\t"
            } else {
                var v67 float32 = 1.5
                switch ((v15 + ((v3 + v15) + v3)) + ((v3 + "hello") + "a")) {
                case "foo bar":
                    switch "\n\t" {
                    case "hello":
                        var v68 bool = (1e10 < ((v67 / v67) + (1.5 * float32(0.0))))
                        _ = v68
                    default:
                        for fi69 := 0; fi69 < 1; fi69++ {
                            var v70 float32 = (v67 - 3.14159)
                            v71 := GenS4{F0: (v2 + v2), F1: 65535, F2: (v1 * 18446744073709551615)}
                            _ = fi69
                            _ = v70
                            _ = v71
                        }
                    }
                case "a":
                    if (320 == 2) {
                        if (1 != 255) {
                            var v72 string = v15
                            var v73 float64 = 3.14159
                            var v74 int64 = (v2 / ((v2) | 1))
                            _ = v72
                            _ = v73
                            _ = v74
                        }
                        var v75 int32 = 42
                        _ = v75
                    } else {
                        var v76 float32 = v67
                        _ = v76
                    }
                default:
                    var v77 uint = 65535
                    _ = v77
                }
                _ = v67
            }
            _ = fi16
        }
        if false {
            fmt.Println(v1, v2, v2)
            fmt.Println(v2)
            var v78 int = 259
            _ = v78
        }
    } else {
        fmt.Println(v15, v2)
        switch 32768 {
        case 32767:
            v1 = v1
        case -306:
            v15 = ("\n\t" + "\n\t")
        case 720:
            switch 0 {
            case 2147483647:
                for fi79 := 0; fi79 < 16; fi79++ {
                    var v80 int64 = v2
                    _ = len(v4)
                    v81 := v4[2147483648]
                    _ = fi79
                    _ = v80
                    _ = v81
                }
            default:
                v3 = ((("" + v15) + "foo bar") + ("" + "hello"))
            }
        }
        for fi82 := 0; fi82 < 16; fi82++ {
            v83 := GenS7{F0: 127}
            if ((100 != 0) || (-1 != 1000)) {
                v84 := GenS5{F0: -1, F1: 2}
                _ = v84
            } else {
                fmt.Println(v3, v1)
                var v85 int = fi82
                switch (v85 ^ (fi82 - 128)) {
                case -2147483648:
                    var v86 bool = (!false)
                    _ = v86
                case 32767:
                    if false {
                        for fi87 := 0; fi87 < 8; fi87++ {
                            fmt.Println(v3)
                            v88 := []int{}
                            if (float32(0.0) >= 0.0) {
                                _ = len(v88)
                                v89 := GenS5{F0: 42, F1: 548}
                                if true {
                                    if (!true) {
                                        var v90 string = v3
                                        if (-512 > -909) {
                                            if len(v88) > 0 {
                                                v88[0] = (-(fi82))
                                            }
                                            v90 = ((v3 + v3) + "a")
                                            delete(v4, 772)
                                        } else {
                                            v3 = ""
                                        }
                                        _ = v90
                                    } else {
                                        fmt.Println(fi82, v2)
                                        if true {
                                            v91 := GenS6{F0: 0, F1: v85, F2: (v85 * v85)}
                                            _ = v91
                                        }
                                    }
                                }
                                _ = v89
                            }
                            v92 := GenS1{F0: 0, F1: 0, F2: false}
                            fmt.Println(fi82, fi87)
                            _ = fi87
                            _ = v88
                            _ = v92
                        }
                    }
                }
                _ = v85
            }
            _ = fi82
            _ = v83
        }
    }
    v93 := GenS2{F0: 2}
    v3 = (v15 + ("hello" + v3))
    for fi94 := 0; fi94 < 13; fi94++ {
        fmt.Println(v3, v3)
        _ = fi94
    }
    for fi95 := 0; fi95 < 17; fi95++ {
        var v96 float32 = -2.25
        v15 = (v3 + ((v3 + ("foo bar" + v15)) + v15))
        _ = fi95
        _ = v96
    }
    v2 -= (v2 << ((510) & 63))
    var v97 uint8 = 255
    v3 = v3
    for fi98 := 0; fi98 < 11; fi98++ {
        delete(v4, 65535)
        for fi99 := 0; fi99 < 17; fi99++ {
            _ = len(v4)
            switch v15 {
            case "foo bar":
                fmt.Println(v2)
            }
            _ = fi99
        }
        _ = fi98
    }
    fmt.Println(v3, v3, v3)
    v2 = v2
    for fi100 := 0; fi100 < 11; fi100++ {
        if ((-(fi100)) > 515) {
            for fi101 := 0; fi101 < 11; fi101++ {
                var v102 bool = (4294967295 == 787)
                _ = fi101
                _ = v102
            }
            v97 = (^(v97))
        } else {
            if true {
                if false {
                    switch (fi100 << ((530) & 63)) {
                    case 4294967295:
                        var v103 int32 = 0
                        _ = v103
                    default:
                        v2 += -1
                    }
                }
            }
            fmt.Println(v97, v3)
        }
        v3 = v3
        _ = fi100
    }
    fmt.Println(v97, v97, v3)
    v104 := GenS4{F0: v2, F1: 988, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v97:", v97)
    _ = v1
    _ = v104
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v93
    _ = v97
}
