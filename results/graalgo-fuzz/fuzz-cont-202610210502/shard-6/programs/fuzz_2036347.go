package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 uint; F2 int16 }
type GenS5 struct { F0 uint32; F1 string; F2 int16 }
type GenS6 struct { F0 int32; F1 uint; F2 string }
type GenS7 struct { F0 bool; F1 int16 }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 2
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 19; fi4++ {
        v3["a"] = (((("" + "a") + "\n\t") + "foo bar") + (("" + "a") + "hello"))
        _ = fi4
    }
    for fi5 := 0; fi5 < 1; fi5++ {
        v2 = 2
        v6 := GenS1{F0: 363, F1: 1, F2: 1}
        v2 -= v2
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 2147483648
    for fi8 := 0; fi8 < 4; fi8++ {
        v2 = (v2 << ((18446744073709551615) & 63))
        _ = fi8
    }
    v9 := GenS1{F0: 42, F1: 660, F2: 4294967295}
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int64 = 0
        _ = fi10
        _ = v11
    }
    switch "foo bar" {
    case "a":
        v7 = 42
    case "":
        for fi12 := 0; fi12 < 13; fi12++ {
            v1 -= (v1 ^ 1)
            _ = len(v3)
            fmt.Println(v7, fi12)
            _ = fi12
        }
    }
    v13 := GenS2{F0: 0}
    if (false && false) {
        v14 := GenS3{F0: (v7 >> ((42) & 63))}
        _ = v14
    } else {
        var v15 uint8 = (v2 - (v2 * v2))
        var v16 bool = false
        _ = v15
        _ = v16
    }
    v2 = v2
    if true {
        var v17 uint64 = 1
        switch ("\n\t" + "foo bar") {
        case "hello":
            if false {
                v18 := GenS2{F0: -163}
                var v19 int64 = v1
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 2; fi20++ {
                    switch "" {
                    case "foo bar":
                        if true {
                            for fi21 := 0; fi21 < 3; fi21++ {
                                var v22 uint64 = (v17 / (((^(v17))) | 1))
                                var v23 int64 = (v1 & v1)
                                v24 := []int{(fi20 >> ((255) & 63)), fi21, (fi21 ^ -279)}
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            v25 := v3["hello"]
                            fmt.Println(v1)
                            _ = v25
                        }
                    case "hello":
                        if ("a" < (("foo bar" + "\n\t") + "")) {
                            if (973 == 255) {
                                for fi26 := 0; fi26 < 11; fi26++ {
                                    switch "\n\t" {
                                    case "":
                                        var v27 int64 = v1
                                        _ = v27
                                    }
                                    v7 -= 42
                                    _ = fi26
                                }
                            } else {
                                fmt.Println(v1)
                                var v28 uint32 = (v7 << ((65535) & 63))
                                _ = v28
                            }
                        } else {
                            switch fi20 {
                            case 721:
                                v2 = v2
                            case 127:
                                v1 -= v1
                            case -2147483648:
                                if false {
                                    var v29 uint = 764
                                    if (-1 > 1) {
                                        v7 = v7
                                        v30 := GenS1{F0: uint64(v2), F1: (v29 % ((0) | 1)), F2: v17}
                                        fmt.Println(v1)
                                        _ = v30
                                    } else {
                                        var v31 int = fi20
                                        v32 := GenS4{F0: -2147483648, F1: (^(v29)), F2: -837}
                                        fmt.Println(v17, v1)
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = v29
                                } else {
                                    var v33 int = (fi20 % ((2147483647) | 1))
                                    for fi34 := 0; fi34 < 14; fi34++ {
                                        if false {
                                            fmt.Println(v17)
                                            v33 = fi20
                                            var v35 int = fi34
                                            switch (-(fi34)) {
                                            case 42:
                                                v3["hello"] = "hello"
                                            default:
                                                var v36 int8 = 42
                                                _ = v36
                                            }
                                            _ = v35
                                        } else {
                                            if false {
                                                if false {
                                                    var v37 int = (fi34 & (fi34 | (fi34 % ((fi34) | 1))))
                                                    if true {
                                                        fmt.Println(v2, v37)
                                                        var v38 uint16 = 65535
                                                        v37 = 42
                                                        _ = v38
                                                    }
                                                    if ((float32(0.1) <= 0.1) && false) {
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(v2)
                                                        case "\n\t":
                                                            if (false && false) {
                                                                if false {
                                                                    fmt.Println(v37, fi20, v2)
                                                                    v33 = v37
                                                                } else {
                                                                    if true {
                                                                        var v39 int16 = -407
                                                                        _ = v39
                                                                    } else {
                                                                        v40 := GenS1{F0: (v17 + 4294967295), F1: 4294967295, F2: 0}
                                                                        v17 = v17
                                                                        _ = v40
                                                                    }
                                                                    v41 := []int{128, v33, (v33 + fi20), (v33 - (fi34 & (v33 + (fi34 - (v37 << ((359) & 63))))))}
                                                                    _ = v41
                                                                }
                                                                switch v33 {
                                                                case 71:
                                                                    switch (^(v37)) {
                                                                    case 32767:
                                                                        var v42 uint16 = 240
                                                                        switch v33 {
                                                                        case -448:
                                                                            fmt.Println(v37)
                                                                        case 1:
                                                                            fmt.Println(v42)
                                                                            if true {
                                                                                v3["foo bar"] = (("a" + "foo bar") + ("a" + "foo bar"))
                                                                                v43 := []int{v37}
                                                                                _ = v43
                                                                            }
                                                                        case 32767:
                                                                            v17 = (v17 << ((0) & 63))
                                                                        default:
                                                                            v33 = 4294967295
                                                                        }
                                                                        _ = v42
                                                                    }
                                                                case 780:
                                                                    if (0 == v37) {
                                                                        switch (^(fi20)) {
                                                                        case 9223372036854775807:
                                                                            var v44 int = 32767
                                                                            _ = v44
                                                                        case 0:
                                                                            fmt.Println(v37, v1)
                                                                        case 1:
                                                                            var v45 int32 = 42
                                                                            _ = v45
                                                                        }
                                                                        switch (^(fi34)) {
                                                                        case -143:
                                                                            v17 -= (v17 + (v17 | 361))
                                                                        default:
                                                                            v37 = v33
                                                                        }
                                                                    } else {
                                                                        v37 += fi34
                                                                        v17 = (^(v17))
                                                                        v46 := v3[("foo bar" + "hello")]
                                                                        var v47 int = fi20
                                                                        v46 = ("foo bar" + v46)
                                                                        switch (fi34 / ((1) | 1)) {
                                                                        case 127:
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                v17 -= v17
                                                                            default:
                                                                                var v48 uint32 = 479
                                                                                _ = v48
                                                                            }
                                                                            v46 = v46
                                                                            v49 := GenS3{F0: v7}
                                                                            _ = v49
                                                                        default:
                                                                            switch v47 {
                                                                            case 42:
                                                                                v46 = v46
                                                                            default:
                                                                                switch fi34 {
                                                                                case -514:
                                                                                    var v50 int64 = v1
                                                                                    _ = v50
                                                                                default:
                                                                                    var v51 uint64 = v17
                                                                                    _ = v51
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            if (v17 < (v17 * v17)) {
                                                                if (81 <= v33) {
                                                                    fmt.Println(v37, v1)
                                                                    if true {
                                                                        if false {
                                                                            if true {
                                                                                switch int(v17) {
                                                                                case 533:
                                                                                    v33 -= (v33 | (v33 % ((v33) | 1)))
                                                                                case -973:
                                                                                    v1 += v1
                                                                                case -1:
                                                                                    v52 := GenS5{F0: uint32(v2), F1: "\n\t", F2: 32767}
                                                                                    _ = v52
                                                                                default:
                                                                                    v1 += v1
                                                                                    fmt.Println(fi34, v1, v17)
                                                                                    v53 := []int{1, 127}
                                                                                    _ = v53
                                                                                }
                                                                                v7 += v7
                                                                            } else {
                                                                                switch 447 {
                                                                                case -531:
                                                                                    var v54 uint8 = 128
                                                                                    _ = v54
                                                                                case 127:
                                                                                    v37 = (^(v37))
                                                                                    if false {
                                                                                        v55 := GenS1{F0: (v17 | 18446744073709551615), F1: 18446744073709551615, F2: (v17 % (((^(v17))) | 1))}
                                                                                        _ = v55
                                                                                    } else {
                                                                                        v56 := GenS6{F0: 653, F1: 385, F2: "\n\t"}
                                                                                        var v57 int8 = 1
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    v2 -= 127
                                                                                default:
                                                                                    switch (("hello" + "a") + ("" + "")) {
                                                                                    case "\n\t":
                                                                                        if ((0.0 / (1.5 / 1.5)) == -2.25) {
                                                                                            if true {
                                                                                                var v58 uint64 = v17
                                                                                                switch ("a" + "a") {
                                                                                                case "\n\t":
                                                                                                    switch "\n\t" {
                                                                                                    case "hello":
                                                                                                        v37 -= (fi20 - 0)
                                                                                                    }
                                                                                                case "a":
                                                                                                    v17 = (v58 + v17)
                                                                                                default:
                                                                                                    if (-1 == -9) {
                                                                                                        v58 -= 548
                                                                                                    } else {
                                                                                                        var v59 int16 = 0
                                                                                                        if false {
                                                                                                            var v60 float64 = ((-2.25 / (((float64(3.14159) / float64(3.14159)) - (1.5 / 1.0)) + (((1e10 + float64(-2.25)) + (0.0 * float64(((3.14159 / 1.0) * float32(1e10))))) * (3.14159 - (1e10 / (1.5 - 0.1)))))) / -2.25)
                                                                                                            v61 := GenS5{F0: 2147483647, F1: ("foo bar" + ""), F2: int16(fi20)}
                                                                                                            switch (v37 << ((563) & 63)) {
                                                                                                            case -331:
                                                                                                                switch ("a" + ("hello" + (("" + ("\n\t" + (("hello" + ("foo bar" + "\n\t")) + "hello"))) + ("foo bar" + "foo bar")))) {
                                                                                                                case "":
                                                                                                                    v17 += (v58 ^ v58)
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v17, v58, v2)
                                                                                                                case "a":
                                                                                                                    v58 += v17
                                                                                                                    v62 := GenS4{F0: 0, F1: 255, F2: 1}
                                                                                                                    _ = v62
                                                                                                                }
                                                                                                            default:
                                                                                                                v63 := []int{-2147483648, v37, v37}
                                                                                                                _ = v63
                                                                                                            }
                                                                                                            var v64 float64 = v60
                                                                                                            _ = v60
                                                                                                            _ = v61
                                                                                                            _ = v64
                                                                                                        } else {
                                                                                                            fmt.Println(fi34, v2)
                                                                                                            if true {
                                                                                                                fmt.Println(fi20)
                                                                                                            } else {
                                                                                                                fmt.Println(v2, v2)
                                                                                                            }
                                                                                                        }
                                                                                                        v7 += v7
                                                                                                        var v65 float64 = 3.14159
                                                                                                        fmt.Println(v33)
                                                                                                        _ = v59
                                                                                                        _ = v65
                                                                                                    }
                                                                                                }
                                                                                                _ = v58
                                                                                            }
                                                                                        } else {
                                                                                            var v66 uint8 = (v2 & 77)
                                                                                            v7 = (v7 | v7)
                                                                                            _ = v66
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v33 += v37
                                                                            }
                                                                            var v67 uint16 = 42
                                                                            var v68 bool = false
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                        var v69 int64 = v1
                                                                        _ = v69
                                                                    } else {
                                                                        v2 -= v2
                                                                        var v70 int64 = v1
                                                                        v71 := GenS4{F0: 42, F1: 18446744073709551615, F2: -32768}
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    v37 = (-(fi20))
                                                                } else {
                                                                    var v72 int8 = int8(v1)
                                                                    var v73 string = "foo bar"
                                                                    var v74 bool = (!((0.1 * 0.0) != 1e10))
                                                                    _ = v72
                                                                    _ = v73
                                                                    _ = v74
                                                                }
                                                                switch fi20 {
                                                                case 9223372036854775807:
                                                                    delete(v3, (("\n\t" + "foo bar") + ""))
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                            } else {
                                                                v7 = (v7 << ((255) & 63))
                                                                v75 := []int{}
                                                                delete(v3, "")
                                                                _ = v75
                                                            }
                                                        }
                                                        v76 := GenS7{F0: true, F1: -153}
                                                        _ = v76
                                                    } else {
                                                        v77 := GenS6{F0: 0, F1: 42, F2: "hello"}
                                                        _ = v77
                                                    }
                                                    _ = v37
                                                } else {
                                                    fmt.Println(v17)
                                                }
                                                if true {
                                                    fmt.Println(fi34, v33, fi20)
                                                    v78 := []int{}
                                                    _ = v78
                                                }
                                            } else {
                                                if (!true) {
                                                    switch (v33 ^ (^(fi20))) {
                                                    case -9223372036854775808:
                                                        fmt.Println(fi20, fi34, v2)
                                                    case 935:
                                                        var v79 string = ("hello" + (("hello" + "\n\t") + ("\n\t" + ("" + "foo bar"))))
                                                        _ = v79
                                                    }
                                                    v80 := GenS7{F0: (v33 >= fi20), F1: -32768}
                                                    _ = v80
                                                } else {
                                                    var v81 uint = 218
                                                    _ = v81
                                                }
                                                fmt.Println(v7)
                                            }
                                        }
                                        _ = fi34
                                    }
                                    _ = v33
                                }
                            }
                            fmt.Println(v1, v1, v7)
                        }
                    default:
                        for fi82 := 0; fi82 < 9; fi82++ {
                            var v83 int8 = 1
                            var v84 bool = true
                            v7 = v7
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                    }
                    for fi85 := 0; fi85 < 4; fi85++ {
                        v1 = 1
                        _ = fi85
                    }
                    _ = fi20
                }
                v86 := GenS4{F0: -480, F1: 2, F2: -18}
                _ = v86
            }
        case "":
            var v87 bool = (42 <= 1)
            _ = v87
        default:
            for fi88 := 0; fi88 < 3; fi88++ {
                v1 += (^(v1))
                v89 := []int{}
                v2 = v2
                _ = fi88
                _ = v89
            }
        }
        _ = v17
    }
    fmt.Println(v1, v1, v7)
    var v90 int8 = -128
    if (!true) {
        if true {
            v91 := []int{int(v1), 42, 783, 1}
            for fi92 := 0; fi92 < 10; fi92++ {
                v1 += (^(v1))
                var v93 uint8 = v2
                v90 -= -1
                _ = fi92
                _ = v93
            }
            _ = v91
        } else {
            for fi94 := 0; fi94 < 3; fi94++ {
                v2 = v2
                _ = fi94
            }
            var v95 uint64 = 1
            fmt.Println(v7)
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v9
    _ = v90
}
