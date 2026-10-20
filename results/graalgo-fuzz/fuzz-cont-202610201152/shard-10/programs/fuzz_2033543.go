package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 string; F2 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool }

func main() {
    v1 := []int{-9223372036854775808}
    var v2 int32 = 1
    var v3 uint16 = 42
    var v4 uint64 = 1
    var v5 string = "a"
    fmt.Println(v3, v5)
    fmt.Println(v5, v4, v4)
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 = (v3 - (v3 + v3))
        _ = fi6
    }
    var v7 int = -672
    switch (v5 + v5) {
    case "":
        v8 := GenS1{F0: v2}
        _ = v8
    case "\n\t":
        if (0 > 2) {
            v4 = (v4 - v4)
            v1 = append(v1, v7)
            v4 -= v4
        } else {
            v9 := GenS1{F0: (v2 >> ((2) & 63))}
            switch "foo bar" {
            case "\n\t":
                var v10 uint32 = 65535
                _ = v10
            default:
                if true {
                    v11 := GenS1{F0: (^(v2))}
                    var v12 bool = false
                    _ = v11
                    _ = v12
                } else {
                    v3 = (v3 & (v3 & v3))
                    if (!(!(1 == 0))) {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            var v14 uint8 = 222
                            _ = fi13
                            _ = v14
                        }
                        var v15 int = int(v4)
                        fmt.Println(v4, v4)
                        _ = v15
                    }
                    var v16 bool = (-2.25 != 1.5)
                    _ = v16
                }
            }
            _ = v9
        }
    default:
        if (32767 >= (v3 % ((2) | 1))) {
            if true {
                switch (v5 + (((v5 + (("foo bar" + v5) + (((v5 + "") + "hello") + ("" + v5)))) + (v5 + v5)) + v5)) {
                case "foo bar":
                    for fi17 := 0; fi17 < 12; fi17++ {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            var v19 float32 = (float32(0.1) / float32(1.5))
                            v7 = (fi17 / ((-9223372036854775808) | 1))
                            _ = fi18
                            _ = v19
                        }
                        switch (v7 >> ((997) & 63)) {
                        case 0:
                            for fi20 := 0; fi20 < 16; fi20++ {
                                v4 -= v4
                                v5 = ("\n\t" + ("foo bar" + v5))
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (127 >= fi20) {
                                        var v21 uint16 = v3
                                        v22 := GenS1{F0: (v2 * (v2 + 2147483647))}
                                        fmt.Println(fi20)
                                        _ = v21
                                        _ = v22
                                    }
                                } else {
                                    v7 = fi20
                                    v2 -= (v2 | -2147483648)
                                }
                                _ = fi20
                            }
                        case 32768:
                            if (!((-97 <= 21) && (v7 >= v7))) {
                                for fi23 := 0; fi23 < 16; fi23++ {
                                    var v24 float32 = 3.14159
                                    _ = fi23
                                    _ = v24
                                }
                                var v25 uint16 = v3
                                _ = v25
                            }
                        }
                        v2 += v2
                        _ = fi17
                    }
                case "\n\t":
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v4 -= v4
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v27 := GenS1{F0: v2}
                        _ = fi26
                        _ = v27
                    }
                case "hello":
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if ((!true) || (false || (!(false || false)))) {
                            v29 := GenS2{F0: (fi28 % ((4294967295) | 1)), F1: 62, F2: 0}
                            if false {
                                var v30 uint = 18446744073709551615
                                switch (v5 + "hello") {
                                case "foo bar":
                                    fmt.Println(v5, v7)
                                case "\n\t":
                                    var v31 int = (v7 ^ 32768)
                                    _ = v31
                                default:
                                    _ = len(v1)
                                }
                                _ = v30
                            } else {
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    v7 = v7
                                    var v33 int32 = (v2 ^ (v2 * (^(v2))))
                                    switch v5 {
                                    case "hello":
                                        v2 -= v33
                                    case "":
                                        var v34 string = v5
                                        _ = v34
                                    }
                                    switch v5 {
                                    case "foo bar":
                                        fmt.Println(v4, v2, v4)
                                    default:
                                        switch (fi32 - (v7 / ((fi32) | 1))) {
                                        case 32768:
                                            switch v5 {
                                            case "a":
                                                v35 := GenS2{F0: 1, F1: 1, F2: -32768}
                                                _ = v35
                                            case "hello":
                                                switch v5 {
                                                case "hello":
                                                    v36 := GenS2{F0: v7, F1: 255, F2: -32768}
                                                    _ = v36
                                                case "a":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "":
                                                    var v37 uint64 = (v4 * v4)
                                                    _ = v37
                                                default:
                                                    switch fi32 {
                                                    case 32767:
                                                        v38 := GenS1{F0: v2}
                                                        _ = v38
                                                    default:
                                                        fmt.Println(v33, v33)
                                                        var v39 uint16 = (v3 + (v3 << ((0) & 63)))
                                                        _ = v39
                                                    }
                                                }
                                            case "foo bar":
                                                switch fi28 {
                                                case 127:
                                                    if (2147483647 != 688) {
                                                        fmt.Println(v2, fi32)
                                                        fmt.Println(v3, v5, v2)
                                                        v7 = (fi28 | (v7 / (((fi28 & 2147483647)) | 1)))
                                                        v5 = v5
                                                    } else {
                                                        switch v7 {
                                                        case -300:
                                                            v40 := GenS3{F0: (fi28 >> ((0) & 63)), F1: v5, F2: 2}
                                                            _ = v40
                                                        case 992:
                                                            if ((false || (!(false && (false || false)))) && ((0.1 * -2.25) != -2.25)) {
                                                                if true {
                                                                    var v41 int64 = 0
                                                                    fmt.Println(v5, v5)
                                                                    fmt.Println(v41, v7, v3)
                                                                    _ = len(v1)
                                                                    _ = v41
                                                                } else {
                                                                    var v42 int64 = 1
                                                                    _ = v42
                                                                }
                                                                fmt.Println(v3, v7)
                                                                if (-443 <= 1) {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        fmt.Println(fi32, fi32)
                                                                    default:
                                                                        var v43 uint8 = 42
                                                                        _ = v43
                                                                    }
                                                                    var v44 bool = false
                                                                    _ = v44
                                                                } else {
                                                                    switch ((("" + "foo bar") + (((v5 + ((v5 + v5) + v5)) + "a") + "\n\t")) + ((("" + v5) + ((v5 + "a") + (("" + (v5 + "")) + v5))) + "hello")) {
                                                                    case "\n\t":
                                                                        _ = len(v1)
                                                                        v4 = v4
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2147483647
                                                                        }
                                                                    case "a":
                                                                        var v45 uint8 = 2
                                                                        _ = v45
                                                                    default:
                                                                        v46 := GenS1{F0: (v2 >> ((255) & 63))}
                                                                        _ = v46
                                                                    }
                                                                }
                                                                var v47 float32 = 1e10
                                                                v2 = v33
                                                                _ = v47
                                                            } else {
                                                                var v48 int = (fi28 | fi32)
                                                                v7 = fi32
                                                                if true {
                                                                    v7 -= -829
                                                                } else {
                                                                    v49 := GenS4{F0: (^(fi32))}
                                                                    if false {
                                                                        if true {
                                                                            var v50 bool = (false || true)
                                                                            fmt.Println(fi32, v4)
                                                                            _ = v50
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = -938
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = fi28
                                                                        }
                                                                        var v51 float32 = (0.0 - (3.14159 / (-2.25 * 1.5)))
                                                                        if (98 < 42) {
                                                                            v4 += (v4 ^ (v4 & 57))
                                                                        } else {
                                                                            fmt.Println(v4, v2, v7)
                                                                            v52 := GenS1{F0: (v2 + 1)}
                                                                            _ = v52
                                                                        }
                                                                        _ = v51
                                                                    } else {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v53 int = v48
                                                                        fmt.Println(v3)
                                                                        _ = v53
                                                                    }
                                                                    switch (fi32 * 0) {
                                                                    case -2147483648:
                                                                        v5 = v5
                                                                    }
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            }
                                                        default:
                                                            switch -553 {
                                                            case 127:
                                                                v33 = v33
                                                            case 2147483647:
                                                                var v54 int8 = -40
                                                                _ = v54
                                                            case 104:
                                                                v55 := GenS2{F0: (fi32 % (((fi28 ^ 32767)) | 1)), F1: 1, F2: -292}
                                                                _ = v55
                                                            default:
                                                                v4 = (v4 + 765)
                                                            }
                                                        }
                                                        v56 := GenS4{F0: -2147483648}
                                                        v5 = "hello"
                                                        _ = v56
                                                    }
                                                default:
                                                    if false {
                                                        v4 = 79
                                                        var v57 uint = 255
                                                        switch v5 {
                                                        case "foo bar":
                                                            fmt.Println(v33, v3, v57)
                                                        default:
                                                            v58 := GenS5{F0: false}
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    } else {
                                                        v1 = append(v1, 42)
                                                        switch fi28 {
                                                        case -208:
                                                            var v59 float32 = 1e10
                                                            _ = v59
                                                        default:
                                                            v5 = "hello"
                                                        }
                                                    }
                                                }
                                            }
                                        case -531:
                                            var v60 int8 = 118
                                            _ = v60
                                        case -2147483648:
                                            var v61 int64 = 0
                                            _ = v61
                                        default:
                                            if (243 <= 42) {
                                                v62 := GenS2{F0: (fi28 & fi28), F1: 255, F2: -1}
                                                _ = v62
                                            } else {
                                                var v63 int8 = -1
                                                _ = v63
                                            }
                                            v64 := GenS3{F0: v7, F1: "\n\t", F2: 932}
                                            switch 4294967295 {
                                            case 42:
                                                v65 := GenS3{F0: fi32, F1: v5, F2: 333}
                                                _ = v65
                                            case -9223372036854775808:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            _ = v64
                                        }
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                if (1 == 149) {
                                    v7 = (fi28 | v7)
                                    v2 = v2
                                    var v66 bool = true
                                    _ = v66
                                }
                            }
                            var v67 int16 = -32768
                            _ = v29
                            _ = v67
                        } else {
                            v4 = (v4 / (((v4 << ((835) & 63))) | 1))
                            if true {
                                var v68 string = "a"
                                v4 = 2
                                _ = v68
                            }
                            if false {
                                v5 = v5
                                v3 = v3
                            } else {
                                v7 -= fi28
                                var v69 uint32 = 475
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    v69 = 1
                                    v69 = 42
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if false {
                                        var v71 uint16 = (v3 / (((v3 | 492)) | 1))
                                        fmt.Println(fi28, v5)
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                        _ = v71
                                    }
                                    _ = fi70
                                }
                                _ = v69
                            }
                        }
                        if len(v1) > 0 {
                            v1[0] = fi28
                        }
                        var v72 uint16 = (v3 / ((v3) | 1))
                        _ = fi28
                        _ = v72
                    }
                default:
                    var v73 int16 = 0
                    _ = v73
                }
            }
            fmt.Println(v4, v4)
            fmt.Println(v5, v7, v2)
        }
    }
    var v74 string = (v5 + "a")
    var v75 uint16 = v3
    switch v5 {
    case "foo bar":
        for fi76 := 0; fi76 < 16; fi76++ {
            for fi77 := 0; fi77 < 6; fi77++ {
                if (!true) {
                    v4 = (v4 ^ (v4 % ((42) | 1)))
                    v5 = "foo bar"
                } else {
                    var v78 float64 = (0.0 * 0.1)
                    if (0 > -490) {
                        v5 = (((((("\n\t" + ((("foo bar" + v5) + "") + ((v74 + "a") + (v74 + v5)))) + v74) + "") + v74) + "hello") + ((v5 + (v74 + (v5 + (("" + (v74 + "a")) + v74)))) + ""))
                    } else {
                        switch v74 {
                        case "hello":
                            fmt.Println(v78, fi76)
                        }
                    }
                    _ = v78
                }
                _ = fi77
            }
            var v79 uint8 = 1
            switch v7 {
            case -9223372036854775808:
                fmt.Println(fi76)
            case 42:
                v74 = ("" + "foo bar")
            default:
                if (0.0 < (float64(float32(((-2.25 * 0.0) + (0.1 - 0.1)))) + -2.25)) {
                    var v80 int32 = v2
                    _ = v80
                }
            }
            _ = fi76
            _ = v79
        }
    case "a":
        v4 -= v4
    case "\n\t":
        var v81 uint8 = 138
        _ = v81
    default:
        _ = len(v1)
    }
    var v82 int = v7
    _ = len(v1)
    v75 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v74
    _ = v75
    _ = v82
}
