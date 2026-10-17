package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint8; F1 int64; F2 bool }
type GenS5 struct { F0 int; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS7 struct { F0 int; F1 uint; F2 string }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string; F1 uint32 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint8, p1 int16) int {
    p0 -= p0
    fmt.Println(p1)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{fi6, 2147483647, fi6, (fi6 % ((-472) | 1))}
        var v8 uint = 65535
        var v9 string = ""
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    fmt.Println(p1)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 int32 = 2147483647
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: p1}
            v14 := []int{(fi12 / ((2147483647) | 1)), (^(fi12)), (fi10 + fi12)}
            var v15 string = "foo bar"
            fmt.Println(fi12)
            fmt.Println(v11, fi12, v15)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint32, p1 uint8) uint64 {
    switch int(p1) {
    case 42:
        switch -79 {
        case 128:
            var v16 uint = 163
            _ = v16
        case -711:
            v17 := GenS2{F0: "foo bar"}
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            if false {
                var v20 int8 = -1
                switch (^(fi19)) {
                case -960:
                    switch (("\n\t" + ("\n\t" + "\n\t")) + "foo bar") {
                    case "foo bar":
                        if true {
                            v20 += (^(v20))
                            v21 := []int{(fi18 + fi18)}
                            _ = v21
                        } else {
                            switch fi18 {
                            case 9223372036854775807:
                                v22 := []int{}
                                _ = v22
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v20 = v20
                            }
                        }
                    case "a":
                        if true {
                            v23 := GenS3{F0: (false || ((0 <= 2) || ((true && ((fi19 / ((fi19) | 1)) == -2147483648)) || (-32768 == 1000)))), F1: ((p0 == (p0 + (p0 / ((65535) | 1)))) || ((false && (false && true)) && (!((-(fi19)) <= genFn0(p1, -32768)))))}
                            if false {
                                if true {
                                    var v24 string = (("hello" + "") + "hello")
                                    v25 := []int{(-(fi19)), fi19}
                                    _ = v24
                                    _ = v25
                                }
                                v26 := GenS4{F0: (p1 - p1), F1: 9223372036854775807, F2: (2 <= 4294967295)}
                                _ = v26
                            } else {
                                var v27 uint64 = 1
                                fmt.Println(fi19, fi18, v27)
                                _ = v27
                            }
                            var v28 int8 = v20
                            _ = v23
                            _ = v28
                        }
                    }
                case 32767:
                    v20 = (v20 | (v20 / (((v20 / ((v20) | 1))) | 1)))
                case -516:
                    p0 = p0
                default:
                    if ("foo bar" == "a") {
                        v20 = v20
                    } else {
                        v29 := []int{genFn0(255, -1), fi19, fi19}
                        _ = v29
                    }
                }
                _ = v20
            }
            _ = fi19
        }
        v30 := []int{(fi18 << ((35) & 63)), (^(fi18)), fi18, genFn0(127, 1000)}
        _ = fi18
        _ = v30
    }
    for fi31 := 0; fi31 < 19; fi31++ {
        fmt.Println(p1, fi31)
        _ = fi31
    }
    if false {
        p0 += p0
        if (!((0.1 - 0.1) == (1e10 / 3.14159))) {
            p0 += (^(p0))
            var v32 string = "a"
            var v33 int = 4294967295
            _ = v32
            _ = v33
        }
        var v34 float64 = 3.14159
        _ = v34
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := []uint32{1, 947, 4294967295}
    var v5 uint8 = 2
    fmt.Println(v5)
    var v35 uint = 4294967295
    if (true || false) {
        var v36 uint32 = 556
        var v37 uint16 = 1
        switch "a" {
        case "a":
            switch "hello" {
            case "\n\t":
                switch "a" {
                case "foo bar":
                    var v38 int16 = 32767
                    _ = v38
                case "\n\t":
                    var v39 int64 = -459
                    _ = v39
                }
            case "hello":
                if false {
                    if len(v4) > 0 {
                        v4[0] = v36
                    }
                }
            }
        case "":
            if true {
                var v40 uint8 = (^(v5))
                _ = v40
            } else {
                switch "" {
                case "a":
                    v35 += (v35 << (((v35 / ((2) | 1))) & 63))
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    for fi42 := 0; fi42 < 10; fi42++ {
                        var v43 int16 = -1
                        if false {
                            v43 = v43
                            var v44 float64 = -2.25
                            switch (^(fi41)) {
                            case 128:
                                switch "" {
                                case "":
                                    var v45 uint32 = v36
                                    _ = v45
                                case "hello":
                                    var v46 uint = v35
                                    _ = v46
                                case "\n\t":
                                    switch fi41 {
                                    case -621:
                                        v4 = append(v4, uint32(v35))
                                    }
                                default:
                                    var v47 float32 = 3.14159
                                    _ = v47
                                }
                            default:
                                var v48 int8 = (v1 >> ((2) & 63))
                                _ = v48
                            }
                            _ = v44
                        } else {
                            var v49 uint8 = 216
                            var v50 float64 = 0.1
                            _ = v49
                            _ = v50
                        }
                        if (4294967295 > genFn1((v36 + (v36 & (v36 - v36))), (^(v5)))) {
                            switch (("hello" + "\n\t") + "") {
                            case "hello":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            default:
                                v51 := GenS5{F0: (^(fi41)), F1: v37, F2: 516}
                                _ = v51
                            }
                            if true {
                                var v52 int8 = -68
                                _ = len(v2)
                                var v53 uint16 = v37
                                _ = v52
                                _ = v53
                            } else {
                                v54 := GenS6{F0: genFn1(v36, (^(v5))), F1: (v36 & v36), F2: v1}
                                _ = v54
                            }
                            v55 := GenS7{F0: 42, F1: 551, F2: "\n\t"}
                            _ = v55
                        }
                        _ = fi42
                        _ = v43
                    }
                    if ((fi41 > 643) || (false && (fi41 != (fi41 ^ (fi41 * (fi41 ^ -850)))))) {
                        fmt.Println(v36)
                        var v56 int8 = (v1 / ((1) | 1))
                        var v57 int16 = 0
                        _ = v56
                        _ = v57
                    } else {
                        switch (fi41 + genFn0(42, int16(v37))) {
                        case -1:
                            if ((0 < -32768) && false) {
                                var v58 uint8 = v5
                                if (!false) {
                                    fmt.Println(v36, v36)
                                    v59 := GenS5{F0: fi41, F1: 598, F2: uint64(v35)}
                                    _ = v59
                                } else {
                                    if true {
                                        for fi60 := 0; fi60 < 3; fi60++ {
                                            var v61 int = fi41
                                            v62 := v2[4294967295]
                                            _ = fi60
                                            _ = v61
                                            _ = v62
                                        }
                                        if true {
                                            var v63 int64 = 1
                                            for fi64 := 0; fi64 < 13; fi64++ {
                                                v5 -= v5
                                                v65 := v3[fi64]
                                                v1 -= (v1 + int8(v65))
                                                _ = fi64
                                                _ = v65
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v66 := GenS4{F0: v58, F1: (v63 * 42), F2: (0.1 == (1e10 * -2.25))}
                                                _ = v66
                                            default:
                                                for fi67 := 0; fi67 < 9; fi67++ {
                                                    switch ("\n\t" + "") {
                                                    case "a":
                                                        if (!true) {
                                                            var v68 uint8 = v58
                                                            _ = v68
                                                        }
                                                    case "hello":
                                                        var v69 uint = (v35 % (((v35 & v35)) | 1))
                                                        _ = v69
                                                    }
                                                    _ = fi67
                                                }
                                            }
                                            _ = v63
                                        }
                                        for fi70 := 0; fi70 < 10; fi70++ {
                                            var v71 uint16 = v37
                                            if false {
                                                var v72 bool = (0.1 < (0.1 - 0.1))
                                                var v73 int16 = 42
                                                if v72 {
                                                    v74 := GenS4{F0: v58, F1: -450, F2: true}
                                                    _ = v74
                                                } else {
                                                    var v75 int64 = -9223372036854775808
                                                    var v76 int16 = (v73 - v73)
                                                    _ = v75
                                                    _ = v76
                                                }
                                                _ = v72
                                                _ = v73
                                            } else {
                                                v77 := GenS8{F0: v35}
                                                _ = v77
                                            }
                                            _ = fi70
                                            _ = v71
                                        }
                                    }
                                    var v78 uint8 = (^(v58))
                                    _ = v78
                                }
                                var v79 uint64 = 4294967295
                                _ = v58
                                _ = v79
                            } else {
                                v80 := GenS4{F0: (v5 / (((v5 | 171)) | 1)), F1: 21, F2: (-1 == -768)}
                                var v81 float64 = 0.1
                                _ = v80
                                _ = v81
                            }
                        case 42:
                            fmt.Println(v5, v36)
                        default:
                            v1 -= v1
                        }
                        v82 := GenS9{F0: "hello", F1: v36}
                        for fi83 := 0; fi83 < 1; fi83++ {
                            var v84 float64 = (1.5 + float64(0.0))
                            fmt.Println(v84, v37, fi41)
                            v84 -= float64((float32(v84) / 1.0))
                            _ = fi83
                            _ = v84
                        }
                        _ = v82
                    }
                    var v85 uint32 = v36
                    _ = fi41
                    _ = v85
                }
            }
        }
        _ = v36
        _ = v37
    } else {
        var v86 int = 42
        if ((false && true) && (v86 == (v86 ^ genFn0(2, 42)))) {
            v1 = v1
            v87 := v3[(v86 >> (((v35 + 65535)) & 63))]
            if (((((0.1 - ((0.0 + (0.0 * (((1e10 / 1e10) - 3.14159) + 1.5))) / 1e10)) * float64(1e10)) * 0.1) < 3.14159) && false) {
                var v88 uint = uint(v5)
                _ = v88
            }
            _ = v87
        }
        _ = v86
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "hello":
        v35 += v35
    case "":
        v1 = v1
    default:
        if true {
            v1 = v1
            for fi89 := 0; fi89 < 4; fi89++ {
                if ((float64((1e10 + 1.5)) - float64((((((float32(float64((1e10 * 1.5))) / 1.5) / float32(float64(1e10))) + ((float32(((-2.25 / float64(1e10)) * float64((0.0 - (0.1 - ((0.0 * 3.14159) + 0.1)))))) - 1.5) / 0.1)) + 0.0) * 0.0))) == -2.25) {
                    fmt.Println(v5, fi89, v5)
                    v5 += (^(v5))
                }
                fmt.Println(v35, v1)
                _ = fi89
            }
        } else {
            delete(v2, 1)
        }
    }
    v5 -= 61
    fmt.Println(v5, v1, v5)
    v90 := GenS8{F0: v35}
    for fi91 := 0; fi91 < 10; fi91++ {
        for fi92 := 0; fi92 < 10; fi92++ {
            fmt.Println(v1)
            switch genFn0(v5, 42) {
            case 404:
                var v93 uint64 = 18446744073709551615
                v94 := GenS10{F0: 81}
                _ = v93
                _ = v94
            case -9223372036854775808:
                v95 := GenS10{F0: uint64(v35)}
                _ = v95
            default:
                v96 := GenS3{F0: (true && false), F1: ((!("hello" == "hello")) || (42 != 2))}
                _ = v96
            }
            v4 = append(v4, 260)
            _ = fi92
        }
        v5 = v5
        var v97 uint16 = uint16(v5)
        _ = fi91
        _ = v97
    }
    v35 += (^(v35))
    for fi98 := 0; fi98 < 13; fi98++ {
        if len(v4) > 0 {
            v4[0] = 2
        }
        switch fi98 {
        case 9223372036854775807:
            delete(v3, fi98)
        default:
            var v99 string = ("" + (("a" + "\n\t") + ("foo bar" + "a")))
            _ = v99
        }
        var v100 float64 = 1e10
        _ = fi98
        _ = v100
    }
    v1 = 0
    if ((42 == 2) || (2 < 21)) {
        switch "a" {
        case "\n\t":
            v35 -= (v35 * 65535)
        case "":
            delete(v2, 2147483648)
        default:
            if (!true) {
                v1 += -1
                v1 = v1
                var v101 int = genFn0(1, -32768)
                _ = v101
            } else {
                if ((v1 / ((-128) | 1)) > 42) {
                    v102 := GenS4{F0: 236, F1: 42, F2: true}
                    _ = v102
                }
            }
        }
    }
    var v103 uint16 = 0
    delete(v2, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v90
}
