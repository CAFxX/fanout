package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16; F1 uint16; F2 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint }
func genFn0(p0 string, p1 int) uint8 {
    p1 -= (p1 % ((32768) | 1))
    v4 := GenS1{F0: -1}
    if false {
        v5 := GenS1{F0: 42}
        v6 := []int{}
        for fi7 := 0; fi7 < 8; fi7++ {
            p0 = p0
            _ = fi7
        }
        _ = v5
        _ = v6
    }
    v8 := []int{p1, 32768, 65, (p1 - (-(p1)))}
    switch p1 {
    case 127:
        var v9 int8 = -68
        _ = v9
    case 0:
        fmt.Println(p1, p0)
    default:
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int64 = -376
            if len(v8) > 0 {
                v8[0] = fi10
            }
            fmt.Println(v11, v11, p0)
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 255
}
func genFn1(p0 int32, p1 int32) int8 {
    if (true && (!(false && false))) {
        for fi12 := 0; fi12 < 8; fi12++ {
            if false {
                if true {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v14 := []int{fi13, 450}
                        switch ("a" + "foo bar") {
                        case "":
                            p1 = p0
                        default:
                            var v15 int64 = 630
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                    v16 := []int{(fi12 - fi12), 4294967295, (fi12 << ((149) & 63))}
                    _ = v16
                }
            } else {
                var v17 bool = (((((1.5 + ((1e10 + (3.14159 * (1e10 - 0.1))) + 0.0)) / (0.1 - (float32((((((-2.25 / 1.0) - (-2.25 + (0.1 / float64(0.1)))) / 1e10) + 0.1) + (0.1 * -2.25))) - 3.14159))) - float32(0.0)) >= 0.0) || (425 > (fi12 << ((0) & 63))))
                _ = v17
            }
            if (!true) {
                if true {
                    fmt.Println(fi12, p0, p0)
                    p1 = (p0 - -1)
                }
                var v18 int64 = -796
                _ = v18
            } else {
                v19 := []int{42}
                p1 += (-(p1))
                var v20 float64 = -2.25
                _ = v19
                _ = v20
            }
            _ = fi12
        }
    } else {
        for fi21 := 0; fi21 < 18; fi21++ {
            switch "foo bar" {
            case "foo bar":
                if (((481 <= 255) || false) || (!true)) {
                    var v22 float64 = (3.14159 - (0.1 - 0.1))
                    _ = v22
                }
            case "a":
                switch 32767 {
                case 127:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v24 := GenS2{F0: p1, F1: 0}
                        p1 += (p0 << ((792) & 63))
                        p0 = p0
                        _ = fi23
                        _ = v24
                    }
                case 2147483647:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v26 := GenS2{F0: (^(p0)), F1: 32767}
                        if true {
                            switch ("a" + "") {
                            case "\n\t":
                                var v27 uint = 42
                                _ = v27
                            case "a":
                                p0 += p1
                                v28 := GenS1{F0: -13}
                                var v29 uint = 0
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v30 uint8 = 0
                                        _ = v30
                                    case "":
                                        var v31 uint32 = 65535
                                        _ = v31
                                    default:
                                        p1 += (p1 & (p0 & (^(p1))))
                                        v32 := GenS1{F0: int8(fi21)}
                                        v33 := []int{fi21, -183}
                                        _ = v32
                                        _ = v33
                                    }
                                    switch (fi21 & 4294967295) {
                                    case 540:
                                        if ((float32(((float64(0.0) / 1.5) + -2.25)) >= 1e10) || false) {
                                            var v34 int16 = 0
                                            if true {
                                                v35 := GenS3{F0: v34}
                                                _ = v35
                                            }
                                            if true {
                                                fmt.Println(fi21, fi21)
                                                if (1 != 1) {
                                                    v36 := []int{}
                                                    if (!(!(1.5 >= -2.25))) {
                                                        var v37 bool = (true || true)
                                                        v38 := GenS3{F0: (v34 ^ v34)}
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        v34 -= 1000
                                                    }
                                                    _ = v36
                                                } else {
                                                    v39 := []int{42, fi25, (fi21 ^ (-(fi25))), 2147483647}
                                                    v29 -= 69
                                                    _ = v39
                                                }
                                                if (1 > 1) {
                                                    var v40 uint32 = 4294967295
                                                    v41 := GenS2{F0: (p1 - p0), F1: (v34 - v34)}
                                                    _ = v40
                                                    _ = v41
                                                }
                                            } else {
                                                var v42 int64 = 0
                                                _ = v42
                                            }
                                            _ = v34
                                        }
                                    }
                                    switch (^(fi21)) {
                                    case -1:
                                        p0 = -725
                                    case 42:
                                        p0 = (p0 >> ((506) & 63))
                                    default:
                                        v43 := GenS3{F0: -327}
                                        _ = v43
                                    }
                                } else {
                                    fmt.Println(p1, fi25, fi25)
                                }
                                fmt.Println(p0)
                                _ = v28
                                _ = v29
                            default:
                                v44 := []int{}
                                _ = v44
                            }
                            var v45 uint16 = 0
                            _ = v45
                        }
                        _ = fi25
                        _ = v26
                    }
                }
            default:
                fmt.Println(p1, fi21)
            }
            p0 -= p0
            _ = fi21
        }
    }
    v46 := GenS4{F0: 1, F1: int(p1), F2: 578}
    _ = p0
    _ = p1
    _ = v46
    return 0
}

func main() {
    var v1 uint64 = 779
    v2 := []uint32{}
    v3 := make(map[string]int)
    v1 -= v1
    var v47 uint64 = (^(v1))
    v1 += (v47 % ((v47) | 1))
    if (!true) {
        v47 = (^(v1))
        if (0 != 65535) {
            var v48 uint16 = 970
            _ = v48
        } else {
            v49 := GenS5{F0: 205}
            _ = v49
        }
    } else {
        for fi50 := 0; fi50 < 2; fi50++ {
            var v51 uint64 = v1
            for fi52 := 0; fi52 < 3; fi52++ {
                fmt.Println(v47, fi50)
                fmt.Println(v1, v51, fi52)
                var v53 uint = 18446744073709551615
                switch "\n\t" {
                case "":
                    if true {
                        if false {
                            var v54 string = "foo bar"
                            _ = v54
                        } else {
                            fmt.Println(v51)
                            if (!(float32(1e10) != 3.14159)) {
                                v53 = v53
                                v53 -= (v53 - (v53 << (((^(v53))) & 63)))
                                if (genFn0("foo bar", (^(fi50))) > 42) {
                                    v55 := GenS2{F0: 2147483647, F1: 1}
                                    _ = v55
                                } else {
                                    v56 := GenS2{F0: -1, F1: 298}
                                    v1 += v51
                                    _ = v56
                                }
                            } else {
                                switch ("hello" + ((("hello" + ("hello" + "hello")) + "a") + (("foo bar" + "foo bar") + "\n\t"))) {
                                case "\n\t":
                                    v51 = (v51 ^ v51)
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        if true {
                                            v47 = v1
                                        } else {
                                            if true {
                                                switch (-(fi52)) {
                                                case 127:
                                                    v51 = v47
                                                }
                                            }
                                            var v57 uint64 = 1
                                            if false {
                                                switch "foo bar" {
                                                case "a":
                                                    var v58 uint32 = 562
                                                    _ = v58
                                                }
                                                if (-9 == 1) {
                                                    fmt.Println(fi52)
                                                    var v59 int64 = -1
                                                    fmt.Println(fi52, v59)
                                                    _ = v59
                                                }
                                            } else {
                                                if (!(false && true)) {
                                                    var v60 int8 = genFn1(42, 42)
                                                    _ = v60
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        v51 += (v1 >> ((v53) & 63))
                                                    }
                                                }
                                                var v61 bool = (("hello" + "") < "a")
                                                v3["a"] = 510
                                                switch -443 {
                                                case 2147483647:
                                                    var v62 int16 = -727
                                                    _ = v62
                                                }
                                                _ = v61
                                            }
                                            _ = v57
                                        }
                                    case "a":
                                        if ("foo bar" > "\n\t") {
                                            var v63 uint16 = 277
                                            _ = v63
                                        }
                                    default:
                                        var v64 int32 = 311
                                        _ = v64
                                    }
                                }
                            }
                            var v65 uint64 = (^(v51))
                            _ = v65
                        }
                        var v66 bool = (true || false)
                        var v67 float64 = 0.1
                        switch "\n\t" {
                        case "":
                            var v68 int8 = -36
                            _ = v68
                        default:
                            v47 = (v1 << ((v53) & 63))
                        }
                        _ = v66
                        _ = v67
                    }
                case "foo bar":
                    v69 := GenS2{F0: -396, F1: 1}
                    _ = v69
                }
                v53 = 42
                _ = fi52
                _ = v53
            }
            var v70 uint64 = 42
            _ = fi50
            _ = v51
            _ = v70
        }
        if (false && true) {
            for fi71 := 0; fi71 < 15; fi71++ {
                for fi72 := 0; fi72 < 20; fi72++ {
                    var v73 float64 = float64(float32(1e10))
                    v74 := GenS5{F0: 0}
                    var v75 int = 1
                    _ = fi72
                    _ = v73
                    _ = v74
                    _ = v75
                }
                fmt.Println(v47, v1)
                var v76 float64 = 0.1
                _ = fi71
                _ = v76
            }
            if len(v2) > 0 {
                v2[0] = 624
            }
            for fi77 := 0; fi77 < 2; fi77++ {
                var v78 uint8 = 181
                for fi79 := 0; fi79 < 15; fi79++ {
                    var v80 float32 = 3.14159
                    _ = fi79
                    _ = v80
                }
                v78 = 127
                _ = fi77
                _ = v78
            }
        } else {
            var v81 int64 = -829
            _ = v81
        }
    }
    v1 += (v47 / ((354) | 1))
    fmt.Println(v1, v47)
    for fi82 := 0; fi82 < 5; fi82++ {
        v3["foo bar"] = fi82
        switch "foo bar" {
        case "":
            v83 := GenS4{F0: -2147483648, F1: fi82, F2: v47}
            _ = v83
        case "\n\t":
            if false {
                for fi84 := 0; fi84 < 10; fi84++ {
                    fmt.Println(v1, v47)
                    fmt.Println(fi84, fi84, v47)
                    var v85 int32 = 912
                    _ = fi84
                    _ = v85
                }
                v86 := GenS3{F0: -39}
                for fi87 := 0; fi87 < 7; fi87++ {
                    if false {
                        if false {
                            v1 = (v1 - (v1 | v47))
                        }
                    } else {
                        var v88 float64 = -2.25
                        _ = v88
                    }
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v47)
                    }
                    _ = fi87
                }
                _ = v86
            }
        case "foo bar":
            v47 = v47
        }
        _ = fi82
    }
    switch -446 {
    case -1:
        _ = len(v2)
    case 32767:
        fmt.Println(v47)
    case 9223372036854775807:
        if true {
            if (!true) {
                for fi89 := 0; fi89 < 13; fi89++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v90 := GenS1{F0: 1}
                    _ = fi89
                    _ = v90
                }
                switch "" {
                case "":
                    fmt.Println(v47, v47)
                default:
                    if (!true) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
            } else {
                for fi91 := 0; fi91 < 15; fi91++ {
                    switch (fi91 / (((fi91 << ((638) & 63))) | 1)) {
                    case 862:
                        v47 = (v47 ^ v47)
                    case 42:
                        if (569 <= 2147483647) {
                            v1 -= (v1 << ((18446744073709551615) & 63))
                            v92 := GenS6{F0: 2147483647}
                            _ = v92
                        }
                    case 2147483647:
                        var v93 uint8 = 128
                        _ = v93
                    }
                    v2 = append(v2, 263)
                    _ = fi91
                }
            }
            for fi94 := 0; fi94 < 14; fi94++ {
                v95 := GenS7{F0: 2, F1: 42, F2: (false && ((1e10 - 3.14159) > (0.1 - 0.0)))}
                _ = fi94
                _ = v95
            }
            for fi96 := 0; fi96 < 8; fi96++ {
                v1 = v1
                switch "foo bar" {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = 579
                    }
                case "\n\t":
                    v47 = v47
                }
                var v97 int = (^(fi96))
                _ = fi96
                _ = v97
            }
        }
    }
    var v98 int32 = 1
    for fi99 := 0; fi99 < 6; fi99++ {
        v1 = v1
        if true {
            for fi100 := 0; fi100 < 13; fi100++ {
                v47 -= (v47 % ((996) | 1))
                _ = fi100
            }
        } else {
            fmt.Println(v98, v1, fi99)
        }
        _ = fi99
    }
    for fi101 := 0; fi101 < 5; fi101++ {
        v1 -= v1
        v102 := GenS5{F0: 42}
        fmt.Println(fi101)
        _ = fi101
        _ = v102
    }
    v1 += 4294967295
    switch -433 {
    case 405:
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    case 4294967295:
        var v103 int8 = -1
        _ = v103
    default:
        _ = len(v2)
    }
    if false {
        var v104 int16 = -530
        var v105 string = "\n\t"
        _ = v104
        _ = v105
    } else {
        v98 += 504
        if (("foo bar" + "\n\t") < ("hello" + "\n\t")) {
            if ("hello" == "foo bar") {
                _ = len(v3)
                v106 := GenS6{F0: v98}
                _ = v106
            }
            var v107 float32 = 3.14159
            _ = v107
        } else {
            fmt.Println(v47, v1)
        }
        v108 := GenS8{F0: 830}
        _ = v108
    }
    for fi109 := 0; fi109 < 6; fi109++ {
        v110 := GenS9{F0: 587}
        v47 -= (v47 ^ v1)
        var v111 int64 = 42
        _ = fi109
        _ = v110
        _ = v111
    }
    fmt.Println(v1, v47)
    v1 += 329
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v47:", v47)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v98
}
