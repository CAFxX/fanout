package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint32; F1 int64; F2 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 4294967295
    var v3 int16 = 32767
    var v4 uint = 1
    var v5 bool = true
    v3 = v3
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(v4, v5, fi6)
        var v7 int64 = -1
        fmt.Println(v3, v5)
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v2 = (v2 << (((^(v4))) & 63))
    var v9 uint16 = 9
    var v10 uint64 = 18446744073709551615
    var v11 int64 = 42
    if false {
        fmt.Println(v10)
    } else {
        _ = len(v1)
        if false {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint16 = 387
                v14 := v1[(v2 / ((v2) | 1))]
                for fi15 := 0; fi15 < 14; fi15++ {
                    if v5 {
                        v3 = -425
                    } else {
                        switch ((("\n\t" + "foo bar") + ("a" + ("" + "a"))) + "foo bar") {
                        case "foo bar":
                            v16 := GenS1{F0: v4, F1: v3, F2: (fi15 / ((fi15) | 1))}
                            _ = v16
                        case "":
                            fmt.Println(v5, v10)
                        default:
                            switch "hello" {
                            case "foo bar":
                                var v17 uint16 = v13
                                _ = v17
                            case "hello":
                                v9 = (v9 / (((v9 >> (((v4 | v4)) & 63))) | 1))
                            }
                        }
                        v18 := GenS1{F0: v4, F1: v3, F2: fi12}
                        switch ("a" + (("" + "hello") + (("\n\t" + "a") + ("\n\t" + ("\n\t" + "a"))))) {
                        case "\n\t":
                            v1[0] = 65535
                        }
                        _ = v18
                    }
                    _ = fi15
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            if true {
                var v19 int16 = -32768
                _ = v19
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        switch ((((("\n\t" + ((("hello" + ("a" + "hello")) + "\n\t") + "foo bar")) + "foo bar") + "foo bar") + ("hello" + "a")) + "hello") {
        case "\n\t":
            v20 := GenS2{F0: 0, F1: (^(v3))}
            _ = v20
        default:
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                fmt.Println(v10, v2, v9)
            case "":
                delete(v1, v2)
            }
        }
    case "hello":
        switch 0 {
        case 42:
            v21 := GenS1{F0: (v4 >> (((v4 * v4)) & 63)), F1: v3, F2: int(v2)}
            _ = v21
        case 514:
            v11 = v11
        }
    }
    v11 = v11
    if false {
        v11 = (v11 >> ((255) & 63))
        switch ("a" + (("a" + (("a" + "\n\t") + ((("foo bar" + "hello") + ("hello" + ((("\n\t" + (("" + "\n\t") + "hello")) + "foo bar") + ("hello" + "a")))) + (("\n\t" + (((((("" + ("" + ("foo bar" + ("\n\t" + ((("\n\t" + "foo bar") + "a") + "hello"))))) + ("foo bar" + "\n\t")) + (("" + "a") + ("foo bar" + (("" + "hello") + (("foo bar" + "") + ("foo bar" + "hello")))))) + (("foo bar" + "hello") + "a")) + (("a" + "\n\t") + "\n\t")) + ("hello" + "\n\t"))) + ("a" + "foo bar"))))) + "")) {
        case "foo bar":
            switch (("" + "a") + "a") {
            case "a":
                v22 := GenS3{F0: v9, F1: 588}
                _ = v22
            }
        case "":
            if (!(v5 || v5)) {
                v23 := GenS4{F0: "foo bar", F1: v11}
                _ = v23
            } else {
                var v24 int32 = 2147483647
                if (v2 == v2) {
                    v24 += v24
                    _ = len(v1)
                }
                fmt.Println(v9)
                _ = v24
            }
        default:
            if v5 {
                v25 := v1[v2]
                v1[(^(v25))] = v25
                v26 := GenS5{F0: (v9 >> (((v4 + (v4 / ((2) | 1)))) & 63)), F1: 1}
                _ = v25
                _ = v26
            } else {
                switch "\n\t" {
                case "a":
                    var v27 int64 = v11
                    _ = v27
                }
                v4 += (v4 - (v4 | uint(v10)))
            }
        }
        switch -9223372036854775808 {
        case -29:
            if ((129 == 255) || v5) {
                var v28 bool = v5
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 float64 = float64(1.5)
                    _ = fi29
                    _ = v30
                }
                v11 += (^(v11))
                _ = v28
            } else {
                if true {
                    for fi31 := 0; fi31 < 7; fi31++ {
                        if (68 <= 2) {
                            var v32 int64 = v11
                            if (!v5) {
                                var v33 int = fi31
                                _ = v33
                            }
                            _ = v32
                        }
                        _ = fi31
                    }
                } else {
                    if false {
                        switch 42 {
                        case 2147483647:
                            _ = len(v1)
                        case -974:
                            var v34 int16 = v3
                            _ = v34
                        case 128:
                            for fi35 := 0; fi35 < 5; fi35++ {
                                fmt.Println(v11)
                                _ = fi35
                            }
                        }
                        v36 := GenS4{F0: (("" + ("" + "a")) + "hello"), F1: v11}
                        _ = v36
                    } else {
                        v10 = 4294967295
                        if (v5 && v5) {
                            v37 := []int{}
                            fmt.Println(v5, v2, v5)
                            if v5 {
                                var v38 float64 = (0.0 + 3.14159)
                                for fi39 := 0; fi39 < 6; fi39++ {
                                    for fi40 := 0; fi40 < 8; fi40++ {
                                        var v41 float32 = (1.5 + (0.1 + ((1.5 + float32((1e10 / ((v38 - (v38 + -2.25)) * ((v38 * -2.25) + 3.14159))))) - 3.14159)))
                                        _ = fi40
                                        _ = v41
                                    }
                                    var v42 bool = (2147483647 <= 1)
                                    _ = fi39
                                    _ = v42
                                }
                                v43 := GenS3{F0: (v9 & v9), F1: 42}
                                _ = v38
                                _ = v43
                            }
                            _ = v37
                        } else {
                            _ = len(v1)
                            switch -18 {
                            case -796:
                                var v44 string = ""
                                _ = v44
                            case 9223372036854775807:
                                for fi45 := 0; fi45 < 8; fi45++ {
                                    if true {
                                        var v46 int32 = -644
                                        _ = len(v1)
                                        var v47 int8 = 0
                                        _ = v46
                                        _ = v47
                                    }
                                    var v48 int64 = v11
                                    v49 := GenS6{F0: 37, F1: (v48 - 587), F2: v48}
                                    _ = fi45
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                        fmt.Println(v5)
                    }
                    switch 32768 {
                    case -146:
                        if v5 {
                            v50 := []int{}
                            _ = v50
                        }
                    default:
                        if (v5 && v5) {
                            switch ("a" + "") {
                            case "":
                                for fi51 := 0; fi51 < 5; fi51++ {
                                    var v52 int8 = int8(v2)
                                    switch (fi51 - (fi51 ^ (-(fi51)))) {
                                    case 4294967295:
                                        if true {
                                            var v53 uint32 = 65535
                                            v4 -= (^(v4))
                                            _ = v53
                                        }
                                    default:
                                        if v5 {
                                            if (true && v5) {
                                                if (v5 && true) {
                                                    if v5 {
                                                        if (!v5) {
                                                            var v54 uint32 = (^(v2))
                                                            _ = v54
                                                        } else {
                                                            v52 -= 127
                                                            var v55 int8 = 46
                                                            for fi56 := 0; fi56 < 17; fi56++ {
                                                                var v57 int64 = v11
                                                                v10 = v10
                                                                fmt.Println(v5, v5, v9)
                                                                _ = fi56
                                                                _ = v57
                                                            }
                                                            _ = v55
                                                        }
                                                        var v58 int16 = v3
                                                        _ = v58
                                                    }
                                                    v59 := GenS3{F0: 255, F1: (fi51 >> ((uint(v2)) & 63))}
                                                    _ = v59
                                                } else {
                                                    v3 -= 857
                                                    var v60 int32 = -554
                                                    _ = v60
                                                }
                                                fmt.Println(v10, v10, v2)
                                            }
                                            v3 = v3
                                            v2 += 1
                                        } else {
                                            v61 := []int{-1}
                                            _ = v61
                                        }
                                    }
                                    _ = fi51
                                    _ = v52
                                }
                            case "foo bar":
                                if (1.5 == ((0.1 / 3.14159) + 3.14159)) {
                                    var v62 uint16 = (^(v9))
                                    fmt.Println(v3, v2)
                                    var v63 uint64 = (v10 / ((18446744073709551615) | 1))
                                    _ = v62
                                    _ = v63
                                } else {
                                    fmt.Println(v5, v11)
                                    fmt.Println(v2)
                                    fmt.Println(v2)
                                }
                            case "a":
                                fmt.Println(v2, v2)
                            default:
                                var v64 uint8 = 116
                                _ = v64
                            }
                            v11 += v11
                        } else {
                            var v65 bool = (-128 < -128)
                            v9 = (v9 - (v9 & v9))
                            var v66 bool = v5
                            _ = v65
                            _ = v66
                        }
                    }
                    var v67 float64 = 1e10
                    _ = v67
                }
                if ((true || false) && v5) {
                    if v5 {
                        var v68 int = 42
                        var v69 float64 = 3.14159
                        for fi70 := 0; fi70 < 7; fi70++ {
                            var v71 int8 = 100
                            var v72 int16 = v3
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        _ = v68
                        _ = v69
                    }
                    v3 = v3
                    v73 := []int{4294967295, 2147483647, 127, int(v3)}
                    _ = v73
                } else {
                    v4 += (v4 | (v4 - (v4 ^ v4)))
                    for fi74 := 0; fi74 < 18; fi74++ {
                        v11 += v11
                        v5 = ((v9 | v9) != 429)
                        _ = fi74
                    }
                }
            }
        case -9223372036854775808:
            for fi75 := 0; fi75 < 11; fi75++ {
                v10 += (v10 - 0)
                v76 := []int{fi75, fi75, fi75, int(v11)}
                v77 := GenS3{F0: uint16(v10), F1: -300}
                _ = fi75
                _ = v76
                _ = v77
            }
        case 2147483647:
            if false {
                var v78 int16 = 42
                _ = v78
            }
        }
    } else {
        for fi79 := 0; fi79 < 12; fi79++ {
            var v80 int32 = 2147483647
            var v81 int8 = 100
            _ = fi79
            _ = v80
            _ = v81
        }
        v10 = v10
    }
    for fi82 := 0; fi82 < 6; fi82++ {
        switch fi82 {
        case 42:
            v83 := []int{fi82}
            _ = v83
        }
        _ = fi82
    }
    var v84 uint64 = (v10 * v10)
    for fi85 := 0; fi85 < 9; fi85++ {
        var v86 int8 = 0
        var v87 float64 = 0.1
        _ = fi85
        _ = v86
        _ = v87
    }
    var v88 string = ""
    v89 := GenS1{F0: 1, F1: 42, F2: 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v88:", v88)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v84
    _ = v88
    _ = v89
    _ = v9
}
