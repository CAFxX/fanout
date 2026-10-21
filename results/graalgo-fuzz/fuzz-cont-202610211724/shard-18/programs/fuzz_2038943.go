package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 int32 }
type GenS11 struct { F0 uint32; F1 bool }

func main() {
    var v1 int = 0
    v2 := []int{(v1 / ((v1) | 1)), v1, (v1 | v1)}
    var v3 int = v1
    var v4 int = (v1 - (v1 % ((1) | 1)))
    v5 := GenS1{F0: 2147483648}
    for fi6 := 0; fi6 < 20; fi6++ {
        if (0 >= 1) {
            for fi7 := 0; fi7 < 10; fi7++ {
                if (!true) {
                    fmt.Println(fi6)
                    fmt.Println(v1, v1, fi7)
                }
                _ = fi7
            }
            var v8 int16 = 32767
            _ = v8
        } else {
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 float32 = 1.5
                _ = fi9
                _ = v10
            }
        }
        _ = fi6
    }
    v1 = 0
    if true {
        var v11 int64 = 0
        for fi12 := 0; fi12 < 15; fi12++ {
            v2 = append(v2, fi12)
            var v13 uint64 = 1
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        if false {
            switch "" {
            case "a":
                var v14 int64 = -9223372036854775808
                _ = v14
            default:
                var v15 int64 = 42
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 string = ("a" + ("" + "hello"))
            switch (v3 - v1) {
            case -535:
                if (v3 == (v3 + 408)) {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v2 = append(v2, v3)
                        if true {
                            if (false && (false || true)) {
                                var v19 float64 = (0.1 + 1e10)
                                _ = v19
                            }
                            v2 = append(v2, v1)
                        } else {
                            if ((true || (82 == 18446744073709551615)) || true) {
                                var v20 float32 = 3.14159
                                if (360 >= 368) {
                                    v21 := GenS1{F0: 2147483648}
                                    _ = v21
                                }
                                _ = v20
                            }
                            v17 = "hello"
                            v22 := GenS1{F0: 4294967295}
                            v1 = v4
                            _ = v22
                        }
                        if true {
                            fmt.Println(v4, v17)
                            v3 = v4
                            v17 = (v17 + "")
                        }
                        _ = fi18
                    }
                    for fi23 := 0; fi23 < 7; fi23++ {
                        if (true || ((false || (!((v1 >> ((4294967295) & 63)) <= fi16))) || true)) {
                            v24 := GenS2{F0: 854, F1: (((("" + "hello") + "") + v17) + v17), F2: fi16}
                            _ = v24
                        } else {
                            if true {
                                var v25 uint8 = 196
                                v26 := GenS2{F0: 0, F1: ((v17 + v17) + v17), F2: fi16}
                                if false {
                                    switch (-(v3)) {
                                    case -2147483648:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case -1:
                                        if (0 == -2147483648) {
                                            fmt.Println(v4, v17, v4)
                                            v1 += 127
                                            v25 += v25
                                        } else {
                                            v27 := GenS3{F0: 364, F1: 1, F2: true}
                                            switch (v4 + fi16) {
                                            case 42:
                                                var v28 uint64 = 42
                                                _ = v28
                                            case 1:
                                                v1 = (v4 - v3)
                                            case 2147483647:
                                                var v29 uint16 = 2
                                                _ = v29
                                            }
                                            fmt.Println(v1, fi23, v17)
                                            _ = v27
                                        }
                                    }
                                    v30 := GenS1{F0: 42}
                                    _ = v30
                                }
                                _ = v25
                                _ = v26
                            }
                        }
                        _ = fi23
                    }
                } else {
                    if false {
                        var v31 float64 = 1.5
                        _ = v31
                    } else {
                        fmt.Println(v1)
                    }
                    var v32 uint8 = 0
                    _ = v32
                }
            case 42:
                _ = len(v2)
            case 420:
                var v33 int64 = 678
                _ = v33
            }
            _ = fi16
            _ = v17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v2 = append(v2, v4)
    for fi34 := 0; fi34 < 12; fi34++ {
        for fi35 := 0; fi35 < 4; fi35++ {
            switch (v3 << ((658) & 63)) {
            case 1:
                v36 := GenS4{F0: 42}
                _ = v36
            case -234:
                var v37 string = ""
                _ = v37
            }
            fmt.Println(v3, v3)
            _ = fi35
        }
        if true {
            switch (v1 % (((fi34 & 0)) | 1)) {
            case 127:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -2147483648:
                var v38 int8 = 100
                _ = v38
            }
        } else {
            v3 -= (v3 >> ((18446744073709551615) & 63))
            switch "\n\t" {
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "":
                var v39 int32 = 0
                _ = v39
            }
        }
        _ = fi34
    }
    var v40 uint16 = 32767
    v4 = v4
    v1 -= (v1 >> ((370) & 63))
    for fi41 := 0; fi41 < 2; fi41++ {
        switch "hello" {
        case "foo bar":
            var v42 float64 = (-2.25 / ((-2.25 * float64(float32(float64(1.5)))) / 1e10))
            _ = v42
        case "hello":
            for fi43 := 0; fi43 < 9; fi43++ {
                if (0 <= 0) {
                    v40 += (v40 | 0)
                    if len(v2) > 0 {
                        v2[0] = (v3 % ((v4) | 1))
                    }
                    var v44 uint = 255
                    _ = v44
                } else {
                    var v45 float64 = 0.1
                    v4 += fi41
                    if (611 >= uint64(v40)) {
                        var v46 int16 = -207
                        v47 := GenS2{F0: 42, F1: "a", F2: (v4 % ((fi41) | 1))}
                        v48 := GenS5{F0: (v40 * v40)}
                        _ = v46
                        _ = v47
                        _ = v48
                    } else {
                        fmt.Println(v4)
                        v3 = fi41
                        var v49 bool = (!false)
                        _ = v49
                    }
                    _ = v45
                }
                fmt.Println(v3, fi41, fi43)
                _ = len(v2)
                _ = fi43
            }
        default:
            v50 := GenS6{F0: ("a" + "hello"), F1: false}
            _ = v50
        }
        for fi51 := 0; fi51 < 17; fi51++ {
            fmt.Println(fi41)
            if (true && true) {
                v4 = (fi41 / ((127) | 1))
            }
            if (!(!(!false))) {
                fmt.Println(v3, v3)
                v1 -= (v1 | (-(v4)))
                var v52 uint64 = 2
                _ = v52
            }
            var v53 int64 = 42
            v54 := GenS7{F0: (v53 ^ (v53 & v53))}
            switch fi41 {
            case 128:
                fmt.Println(v40)
            }
            _ = fi51
            _ = v53
            _ = v54
        }
        var v55 int16 = -1
        _ = fi41
        _ = v55
    }
    v56 := GenS8{F0: 255, F1: 9223372036854775807}
    if (128 == 127) {
        switch 9223372036854775807 {
        case 0:
            var v57 int8 = 96
            _ = v57
        default:
            if true {
                switch "\n\t" {
                case "":
                    if true {
                        v4 -= v4
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v58 int32 = 42
                        _ = v58
                    } else {
                        v59 := GenS7{F0: 42}
                        if true {
                            switch v1 {
                            case 2147483647:
                                for fi60 := 0; fi60 < 13; fi60++ {
                                    for fi61 := 0; fi61 < 19; fi61++ {
                                        if true {
                                            switch ("" + "a") {
                                            case "a":
                                                v40 -= (v40 % ((v40) | 1))
                                            }
                                        } else {
                                            var v62 int32 = 2147483647
                                            switch -73 {
                                            case 42:
                                                fmt.Println(v1)
                                            case 127:
                                                if len(v2) > 0 {
                                                    v2[0] = (v1 - 32767)
                                                }
                                            case 1:
                                                if (-158 > 9223372036854775807) {
                                                    v4 += (^(v1))
                                                    var v63 int8 = -1
                                                    v64 := GenS9{F0: 42}
                                                    _ = v63
                                                    _ = v64
                                                }
                                            }
                                            v1 = v4
                                            fmt.Println(v40, v3, v1)
                                            fmt.Println(v40, v62, fi60)
                                            _ = v62
                                        }
                                        _ = fi61
                                    }
                                    v3 = -440
                                    _ = fi60
                                }
                            case -692:
                                var v65 bool = ((!((v3 >> ((4294967295) & 63)) != (v3 - (^(v3))))) || true)
                                _ = v65
                            default:
                                var v66 int32 = 1
                                _ = v66
                            }
                            v3 = (v3 + (-(v3)))
                        }
                        _ = v59
                    }
                default:
                    switch ("hello" + ("hello" + "\n\t")) {
                    case "":
                        _ = len(v2)
                    case "a":
                        v1 = v4
                    case "\n\t":
                        v67 := GenS4{F0: -69}
                        _ = v67
                    }
                }
                var v68 float32 = (0.0 / 0.1)
                for fi69 := 0; fi69 < 2; fi69++ {
                    var v70 int8 = 1
                    var v71 string = (("\n\t" + "hello") + "hello")
                    _ = fi69
                    _ = v70
                    _ = v71
                }
                _ = v68
            }
        }
        if len(v2) > 0 {
            v2[0] = (v4 - v1)
        }
    } else {
        if false {
            _ = len(v2)
            var v72 uint32 = uint32(v40)
            _ = v72
        } else {
            if (!false) {
                fmt.Println(v4, v40, v3)
            }
        }
        switch v4 {
        case -2147483648:
            v40 = v40
        case 1:
            for fi73 := 0; fi73 < 9; fi73++ {
                v1 -= fi73
                fmt.Println(v1, v1)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi73
            }
        case 492:
            for fi74 := 0; fi74 < 3; fi74++ {
                v75 := GenS3{F0: v40, F1: 1, F2: ("" >= "a")}
                _ = fi74
                _ = v75
            }
        default:
            for fi76 := 0; fi76 < 5; fi76++ {
                if false {
                    v77 := GenS7{F0: 954}
                    for fi78 := 0; fi78 < 16; fi78++ {
                        if true {
                            fmt.Println(v3, fi76, fi78)
                            v4 = (fi76 * 4294967295)
                            v4 = fi76
                        } else {
                            _ = len(v2)
                            if (((!true) && (!true)) && true) {
                                v79 := GenS4{F0: 100}
                                v80 := GenS9{F0: 9223372036854775807}
                                _ = v79
                                _ = v80
                            }
                        }
                        _ = fi78
                    }
                    v3 += fi76
                    _ = v77
                }
                v81 := GenS10{F0: 244, F1: 1}
                _ = fi76
                _ = v81
            }
        }
        fmt.Println(v3)
    }
    v82 := GenS11{F0: 4294967295, F1: (true || true)}
    v1 -= -976
    var v83 uint16 = (v40 & v40)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v56
    _ = v82
    _ = v83
}
