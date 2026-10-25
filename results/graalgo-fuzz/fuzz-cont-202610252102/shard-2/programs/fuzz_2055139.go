package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 int64; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 int8; F1 uint64 }

func main() {
    v1 := []string{"\n\t", "", (((("\n\t" + "a") + "hello") + (("" + ("a" + "")) + ("\n\t" + ("\n\t" + ("hello" + ""))))) + "hello")}
    var v2 float64 = float64((1.5 * 3.14159))
    v3 := []int{9223372036854775807, -1, 127, 4294967295}
    var v4 float64 = v2
    var v5 uint8 = 128
    var v6 string = "foo bar"
    fmt.Println(v5, v5, v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        v6 = "a"
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if ((fi9 >> ((2) & 63)) != fi9) {
            if (1 >= 4294967295) {
                switch (v6 + v6) {
                case "a":
                    var v10 string = ((("a" + "hello") + v6) + v6)
                    _ = v10
                }
                v4 = v4
            } else {
                switch v6 {
                case "hello":
                    if ((-1 <= 1) || (false || false)) {
                        v11 := GenS1{F0: -322}
                        var v12 int = (fi9 & 32768)
                        v3 = append(v3, fi9)
                        _ = v11
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            if true {
                                v14 := GenS1{F0: -1}
                                v6 = "a"
                                var v15 uint8 = (v5 % ((v5) | 1))
                                if false {
                                    switch ((("a" + v6) + v6) + "\n\t") {
                                    case "a":
                                        v3 = append(v3, 9223372036854775807)
                                    case "foo bar":
                                        v6 = v6
                                    }
                                    var v16 int16 = 1
                                    if len(v1) > 0 {
                                        v1[0] = ((v6 + (v6 + "a")) + v6)
                                    }
                                    _ = v16
                                }
                                v17 := GenS2{F0: 2147483648, F1: "foo bar", F2: 581}
                                _ = v14
                                _ = v15
                                _ = v17
                            } else {
                                var v18 int16 = 32767
                                v18 += 42
                                _ = v18
                            }
                            var v19 int16 = -32768
                            var v20 string = v6
                            _ = fi13
                            _ = v19
                            _ = v20
                        }
                        v21 := GenS1{F0: 1}
                        v3 = append(v3, (fi9 + (fi9 << ((403) & 63))))
                        _ = v21
                    }
                }
                switch "a" {
                case "a":
                    v5 = 128
                }
            }
            v22 := GenS1{F0: 42}
            _ = v22
        } else {
            if (4294967295 != 65535) {
                v4 += (-2.25 / float64(3.14159))
            } else {
                v23 := GenS2{F0: 65535, F1: v6, F2: 0}
                for fi24 := 0; fi24 < 17; fi24++ {
                    if len(v1) > 0 {
                        v1[0] = (v6 + "")
                    }
                    var v25 bool = false
                    _ = fi24
                    _ = v25
                }
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 uint64 = 700
                    if (!(v27 <= (v27 | v27))) {
                        fmt.Println(v27, v4, v4)
                        fmt.Println(v5)
                        var v28 bool = (!(-117 != -128))
                        _ = v28
                    } else {
                        v29 := GenS2{F0: 65535, F1: v6, F2: 0}
                        var v30 bool = false
                        v31 := GenS3{F0: 42, F1: -32768}
                        var v32 int = fi26
                        _ = v29
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    v1 = append(v1, v6)
                    _ = fi26
                    _ = v27
                }
                _ = v23
            }
        }
        _ = fi9
    }
    if true {
        v4 += ((v2 / 0.1) / v2)
        for fi33 := 0; fi33 < 12; fi33++ {
            var v34 uint32 = 2147483648
            if false {
                v34 = v34
            } else {
                var v35 uint64 = 2
                v36 := GenS3{F0: 2147483647, F1: -1}
                _ = v35
                _ = v36
            }
            _ = fi33
            _ = v34
        }
        var v37 int32 = -1
        _ = v37
    }
    _ = len(v3)
    if true {
        if false {
            switch 2147483647 {
            case -9223372036854775808:
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
            }
            v6 = "a"
            switch 54 {
            case -9223372036854775808:
                if false {
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v5 += v5
                        if (false || true) {
                            for fi39 := 0; fi39 < 1; fi39++ {
                                var v40 uint64 = 2
                                var v41 uint32 = 42
                                v40 -= (v40 >> ((65535) & 63))
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            v4 -= 3.14159
                            if true {
                                fmt.Println(v4, v6, v2)
                                if true {
                                    switch v6 {
                                    case "hello":
                                        _ = len(v1)
                                    case "":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "a":
                                        fmt.Println(v2, v6)
                                    }
                                    v42 := GenS1{F0: 732}
                                    _ = v42
                                } else {
                                    for fi43 := 0; fi43 < 15; fi43++ {
                                        if (false && ((1 < 789) && (!(false || false)))) {
                                            var v44 uint8 = v5
                                            switch v6 {
                                            case "\n\t":
                                                var v45 int8 = 0
                                                _ = v45
                                            case "foo bar":
                                                var v46 bool = true
                                                _ = v46
                                            }
                                            var v47 string = (v6 + v6)
                                            _ = v44
                                            _ = v47
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v48 := GenS3{F0: -856, F1: 1000}
                                            if true {
                                                var v49 uint32 = 0
                                                var v50 uint8 = (v5 & v5)
                                                var v51 bool = (false && false)
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v52 := GenS4{F0: 32767, F1: 2147483647}
                                                var v53 uint = 255
                                                _ = v52
                                                _ = v53
                                            }
                                            switch fi38 {
                                            case 128:
                                                v5 = (v5 + v5)
                                            case 258:
                                                if true {
                                                    v54 := GenS5{F0: 42}
                                                    _ = v54
                                                }
                                            }
                                            _ = v48
                                        }
                                        if ((-2147483648 <= -689) || (-105 == -128)) {
                                            switch v6 {
                                            case "\n\t":
                                                v2 = 0.1
                                            default:
                                                v55 := GenS4{F0: 1, F1: -2147483648}
                                                _ = v55
                                            }
                                            v56 := GenS1{F0: 2147483647}
                                            if true {
                                                if true {
                                                    if (!(!(false && false))) {
                                                        v6 = (v6 + "hello")
                                                        if true {
                                                            var v57 uint64 = 42
                                                            _ = v57
                                                        } else {
                                                            v4 += ((-2.25 + ((((v2 - v2) / 1.5) / ((v4 - float64(0.0)) - float64((0.1 + (1.5 - 0.0))))) + 0.1)) + ((v4 - v2) + (v2 * v4)))
                                                        }
                                                        v58 := GenS6{F0: (!(42 < -1)), F1: 0, F2: 32767}
                                                        _ = v58
                                                    }
                                                    fmt.Println(v6)
                                                }
                                                var v59 int = fi38
                                                _ = v59
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    var v60 int8 = 1
                                                    if (!((("foo bar" + v6) + v6) > v6)) {
                                                        if (v6 != v6) {
                                                            v61 := GenS7{F0: -32768}
                                                            var v62 bool = false
                                                            fmt.Println(v4)
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                        var v63 uint16 = 2
                                                        var v64 int64 = 1
                                                        switch (v6 + "hello") {
                                                        case "a":
                                                            if len(v3) > 0 {
                                                                v3[0] = fi38
                                                            }
                                                        default:
                                                            fmt.Println(v6, v60)
                                                        }
                                                        _ = v63
                                                        _ = v64
                                                    }
                                                    _ = v60
                                                default:
                                                    var v65 uint32 = 2147483647
                                                    _ = v65
                                                }
                                                var v66 float32 = -2.25
                                                _ = v66
                                            }
                                            _ = v56
                                        } else {
                                            if (3.14159 >= (((((((1.5 + (((-2.25 + ((0.0 + 0.1) * ((-2.25 / 0.1) + 0.0))) / (3.14159 / ((0.0 - 3.14159) + (3.14159 * (1e10 - -2.25))))) / float32(v2))) - (3.14159 - (0.0 / (((float32((v4 + (v2 + 1.5))) * 1.5) / (((3.14159 / float32(1e10)) - (0.0 * 3.14159)) / -2.25)) - 1e10)))) * 3.14159) - (0.1 / -2.25)) / 1.0) / float32(v4)) + (((3.14159 * (1.5 - (1e10 * 0.1))) + (0.0 + ((1e10 + -2.25) / 3.14159))) - ((float32((v4 / v4)) / 1.5) - 0.1)))) {
                                                v6 = "hello"
                                                if false {
                                                    if (277 > 277) {
                                                        v5 += v5
                                                        var v67 uint32 = 0
                                                        _ = v67
                                                    } else {
                                                        v6 = "\n\t"
                                                    }
                                                    switch v6 {
                                                    case "":
                                                        v68 := GenS4{F0: -32768, F1: -2147483648}
                                                        _ = v68
                                                    }
                                                    v69 := GenS6{F0: true, F1: -401, F2: 776}
                                                    _ = v69
                                                } else {
                                                    switch (fi38 ^ (fi43 % ((fi43) | 1))) {
                                                    case 1:
                                                        var v70 uint8 = v5
                                                        _ = v70
                                                    case 42:
                                                        var v71 int16 = int16(v5)
                                                        _ = v71
                                                    case 127:
                                                        v5 = 255
                                                    }
                                                    v2 = 0.0
                                                    switch 128 {
                                                    case -1:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case 42:
                                                        v72 := GenS5{F0: 905}
                                                        _ = v72
                                                    }
                                                }
                                                var v73 float64 = (v2 / ((v4 - (v2 * 0.1)) + v4))
                                                if false {
                                                    v6 = v6
                                                    var v74 int16 = 0
                                                    var v75 float32 = float32(v73)
                                                    _ = v74
                                                    _ = v75
                                                } else {
                                                    var v76 int16 = -62
                                                    _ = v76
                                                }
                                                _ = v73
                                            } else {
                                                v6 = "a"
                                                var v77 int = (fi43 | fi43)
                                                fmt.Println(v5, v6)
                                                _ = v77
                                            }
                                            if false {
                                                v78 := GenS6{F0: false, F1: -9223372036854775808, F2: 42}
                                                _ = v78
                                            }
                                        }
                                        v6 = v6
                                        _ = fi43
                                    }
                                }
                            }
                        }
                        _ = fi38
                    }
                }
            }
        }
        v6 = v6
    } else {
        var v79 int8 = 52
        _ = v79
    }
    switch v6 {
    case "hello":
        if true {
            for fi80 := 0; fi80 < 11; fi80++ {
                if (9223372036854775807 >= 289) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = ("hello" + v6)
                    }
                }
                _ = fi80
            }
            if ((9223372036854775807 > 42) || true) {
                if false {
                    if len(v1) > 0 {
                        v1[0] = v6
                    }
                    for fi81 := 0; fi81 < 14; fi81++ {
                        for fi82 := 0; fi82 < 15; fi82++ {
                            _ = len(v3)
                            switch 0 {
                            case 0:
                                switch 0 {
                                case -358:
                                    v6 = ((("foo bar" + v6) + ((v6 + (v6 + ("a" + ("\n\t" + v6)))) + v6)) + v6)
                                default:
                                    var v83 float32 = -2.25
                                    _ = v83
                                }
                            case 9223372036854775807:
                                v84 := GenS4{F0: 1, F1: 0}
                                _ = v84
                            case -91:
                                var v85 string = v6
                                _ = v85
                            }
                            if (!((v5 - 255) < (v5 << ((1) & 63)))) {
                                v4 = (v2 * v4)
                                fmt.Println(v6)
                                v86 := GenS8{F0: 2147483648, F1: 656}
                                _ = v86
                            } else {
                                switch -1 {
                                case 128:
                                    if true {
                                        v5 = (v5 + v5)
                                        v5 = (v5 / (((^(v5))) | 1))
                                    }
                                }
                            }
                            _ = fi82
                        }
                        _ = fi81
                    }
                    for fi87 := 0; fi87 < 11; fi87++ {
                        switch fi87 {
                        case -2147483648:
                            for fi88 := 0; fi88 < 6; fi88++ {
                                var v89 bool = false
                                _ = fi88
                                _ = v89
                            }
                        case 357:
                            fmt.Println(v5, v5)
                        case 0:
                            switch fi87 {
                            case -843:
                                fmt.Println(v6, v5)
                            }
                        }
                        for fi90 := 0; fi90 < 20; fi90++ {
                            var v91 int8 = -128
                            _ = fi90
                            _ = v91
                        }
                        var v92 string = v6
                        _ = fi87
                        _ = v92
                    }
                } else {
                    for fi93 := 0; fi93 < 11; fi93++ {
                        var v94 int = -350
                        _ = fi93
                        _ = v94
                    }
                }
                fmt.Println(v4, v5, v6)
            }
        } else {
            fmt.Println(v6, v5, v2)
        }
    case "foo bar":
        var v95 uint8 = v5
        _ = v95
    case "a":
        v96 := GenS9{F0: -106, F1: 535}
        _ = v96
    default:
        fmt.Println(v5, v5)
    }
    v4 = 0.1
    if false {
        var v97 uint32 = 42
        _ = v97
    } else {
        var v98 int8 = -1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v98
    }
    _ = len(v3)
    v2 = float64(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
