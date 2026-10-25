package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int32; F1 int32; F2 int8 }
type GenS10 struct { F0 int32; F1 int64; F2 int16 }
type GenS11 struct { F0 string; F1 int8; F2 uint64 }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 string = ("" + "\n\t")
    v2 := make(map[string]uint32)
    v3 := []string{"hello", v1, (v1 + "a")}
    var v4 float32 = (float32(3.14159) / 1e10)
    v5 := []int{128, -158, 32767}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint = 255
        var v8 int64 = 1
        for fi9 := 0; fi9 < 11; fi9++ {
            fmt.Println(fi9, v7, v8)
            v10 := GenS1{F0: (fi6 - 9223372036854775807), F1: 383, F2: 42}
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v2)
    fmt.Println(v1, v4)
    v5 = append(v5, 1)
    v2[v1] = 643
    v1 = ("foo bar" + "foo bar")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v11 int = 128
    switch ((v1 + v1) + v1) {
    case "foo bar":
        if (true && false) {
            switch 32768 {
            case 711:
                v12 := GenS1{F0: (^(v11)), F1: 1, F2: 2147483647}
                _ = v12
            case -955:
                var v13 uint32 = 2147483648
                _ = v13
            case 4294967295:
                var v14 float64 = 1e10
                _ = v14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: (^(fi15)), F1: 255, F2: 2147483647}
                    v18 := v2[(v1 + v1)]
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        } else {
            _ = len(v3)
            if true {
                if (0.0 >= ((1e10 - ((v4 - 0.1) - v4)) / ((1.5 - float32(0.0)) - v4))) {
                    switch ((v1 + "\n\t") + ("hello" + v1)) {
                    case "\n\t":
                        for fi19 := 0; fi19 < 14; fi19++ {
                            var v20 bool = true
                            var v21 bool = (!(42 > 9223372036854775807))
                            v11 = -1
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 int16 = 32767
                    fmt.Println(v11, v22)
                    _ = v22
                }
                v23 := v2[("\n\t" + v1)]
                _ = v23
            } else {
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 uint8 = 1
                    if false {
                        if false {
                            switch v1 {
                            case "foo bar":
                                if false {
                                    var v26 int = fi24
                                    _ = v26
                                } else {
                                    var v27 uint8 = (^(v25))
                                    _ = v27
                                }
                            case "\n\t":
                                var v28 int = 42
                                _ = v28
                            case "hello":
                                var v29 float32 = v4
                                _ = v29
                            default:
                                switch v1 {
                                case "\n\t":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    switch -1 {
                                    case 9223372036854775807:
                                        for fi30 := 0; fi30 < 18; fi30++ {
                                            v31 := GenS2{F0: "", F1: -551, F2: 1}
                                            v25 = v25
                                            _ = fi30
                                            _ = v31
                                        }
                                    case -2147483648:
                                        for fi32 := 0; fi32 < 3; fi32++ {
                                            var v33 float32 = v4
                                            v4 = v33
                                            v4 -= (v33 / -2.25)
                                            var v34 uint = 65535
                                            v25 = 2
                                            _ = fi32
                                            _ = v33
                                            _ = v34
                                        }
                                    case 798:
                                        v4 = 0.1
                                    }
                                }
                            }
                            fmt.Println(fi24, v4)
                        }
                    }
                    var v35 uint16 = 982
                    _ = fi24
                    _ = v25
                    _ = v35
                }
                fmt.Println(v1, v4)
                var v36 float32 = float32(3.14159)
                _ = v36
            }
        }
    case "":
        var v37 int8 = 127
        _ = v37
    default:
        if (v4 >= ((((1e10 * v4) / 1.0) + (0.1 * v4)) + (0.1 / 1.0))) {
            v38 := v2["\n\t"]
            if len(v3) > 0 {
                v3[0] = v1
            }
            v1 = v1
            _ = v38
        } else {
            switch (v11 << ((18446744073709551615) & 63)) {
            case 503:
                if (true && ((!true) || ("\n\t" > "a"))) {
                    if (!(!false)) {
                        v2[(v1 + ((v1 + v1) + v1))] = 2147483648
                        var v39 string = (v1 + v1)
                        _ = v39
                    } else {
                        v4 = (float32(3.14159) + (0.0 + v4))
                        if (true || true) {
                            var v40 int64 = 0
                            for fi41 := 0; fi41 < 12; fi41++ {
                                switch (-(fi41)) {
                                case 42:
                                    for fi42 := 0; fi42 < 14; fi42++ {
                                        var v43 bool = false
                                        _ = len(v2)
                                        var v44 uint = 42
                                        if true {
                                            if (-460 < -946) {
                                                v45 := GenS3{F0: 150}
                                                _ = v45
                                            }
                                            if v43 {
                                                v40 -= v40
                                                var v46 uint8 = 207
                                                _ = v46
                                            }
                                            v5 = append(v5, -9223372036854775808)
                                            if (((42 <= 2) || v43) || (1.5 < (0.0 + (float64((v4 / 1.0)) * ((3.14159 * (0.0 / 0.1)) / 1.0))))) {
                                                var v47 int32 = 2147483647
                                                _ = v47
                                            }
                                        } else {
                                            if (!v43) {
                                                v48 := GenS2{F0: (v1 + v1), F1: v40, F2: 42}
                                                _ = v48
                                            }
                                            v49 := GenS4{F0: 212}
                                            _ = len(v2)
                                            delete(v2, (((("\n\t" + ((v1 + "\n\t") + "a")) + ("foo bar" + v1)) + ((((v1 + "foo bar") + "\n\t") + v1) + v1)) + "a"))
                                            v43 = (((v44 - v44) == 4294967295) && v43)
                                            _ = v49
                                        }
                                        if ((-94 > 174) || v43) {
                                            _ = len(v3)
                                            var v50 int32 = -747
                                            _ = v50
                                        } else {
                                            v43 = ((32767 > -32768) || v43)
                                            if v43 {
                                                if ((v44 ^ 65535) <= (^(v44))) {
                                                    var v51 uint = v44
                                                    _ = v51
                                                } else {
                                                    switch (fi42 & (fi41 << (((v44 >> ((42) & 63))) & 63))) {
                                                    case -443:
                                                        if false {
                                                            v44 = v44
                                                        }
                                                    case 128:
                                                        v43 = (v40 != 0)
                                                    case -963:
                                                        switch v1 {
                                                        case "foo bar":
                                                            v40 -= v40
                                                        }
                                                    }
                                                    _ = len(v5)
                                                    fmt.Println(v43)
                                                }
                                            } else {
                                                v40 -= (v40 & 42)
                                                var v52 float32 = ((v4 - (((1.5 - ((v4 - ((v4 - (v4 + (1e10 - v4))) / (v4 + v4))) * (((v4 * 1e10) + v4) - v4))) + v4) / v4)) + (((v4 * v4) * (v4 - 0.0)) / v4))
                                                _ = v52
                                            }
                                            var v53 float64 = 0.1
                                            _ = v53
                                        }
                                        var v54 float32 = v4
                                        _ = fi42
                                        _ = v43
                                        _ = v44
                                        _ = v54
                                    }
                                case 2147483647:
                                    for fi55 := 0; fi55 < 6; fi55++ {
                                        fmt.Println(v4, v11)
                                        v1 = ("" + v1)
                                        v1 = ((v1 + "a") + "\n\t")
                                        _ = fi55
                                    }
                                case 1:
                                    v56 := GenS1{F0: (fi41 | fi41), F1: 255, F2: 1}
                                    _ = v56
                                default:
                                    var v57 uint16 = 65535
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 12; fi58++ {
                                    v59 := GenS4{F0: 2147483647}
                                    if ((0.0 - float64((((v4 * -2.25) * ((float32(1e10) / 1.0) - 0.0)) + float32(float64(float32(0.1)))))) < 0.1) {
                                        v1 = "\n\t"
                                    } else {
                                        var v60 int = (-(fi41))
                                        var v61 uint32 = 2
                                        var v62 int16 = -32768
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                                _ = fi41
                            }
                            var v63 int8 = 42
                            _ = v40
                            _ = v63
                        } else {
                            if (65535 == 42) {
                                fmt.Println(v4, v1)
                            } else {
                                v5 = append(v5, (v11 * -1))
                                for fi64 := 0; fi64 < 17; fi64++ {
                                    v65 := GenS5{F0: 2, F1: -9223372036854775808}
                                    _ = fi64
                                    _ = v65
                                }
                                if (!(((v1 + "\n\t") + v1) <= "hello")) {
                                    _ = len(v2)
                                    var v66 int8 = 1
                                    _ = v66
                                } else {
                                    var v67 int = v11
                                    fmt.Println(v4)
                                    v1 = (((v1 + v1) + (v1 + (v1 + (v1 + ("a" + v1))))) + ("foo bar" + ("a" + "")))
                                    _ = v67
                                }
                            }
                        }
                    }
                    v11 = 0
                    if false {
                        v11 = 128
                        v4 -= float32(1e10)
                        v11 += (v11 << ((1) & 63))
                    } else {
                        if (42 < 481) {
                            v4 += v4
                            var v68 int64 = -1
                            for fi69 := 0; fi69 < 5; fi69++ {
                                var v70 int = fi69
                                var v71 int32 = 208
                                _ = fi69
                                _ = v70
                                _ = v71
                            }
                            _ = v68
                        }
                        switch v11 {
                        case 0:
                            for fi72 := 0; fi72 < 10; fi72++ {
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    v74 := v2["a"]
                                    var v75 uint64 = 18446744073709551615
                                    fmt.Println(fi73)
                                    switch 2147483647 {
                                    case 127:
                                        switch fi72 {
                                        case -15:
                                            v1 = "a"
                                        }
                                    case -1:
                                        v1 = v1
                                    }
                                    _ = fi73
                                    _ = v74
                                    _ = v75
                                }
                                _ = fi72
                            }
                        }
                    }
                }
            case 128:
                v4 -= 0.0
            default:
                switch v11 {
                case 127:
                    v76 := GenS1{F0: -1, F1: 62, F2: 2147483647}
                    _ = v76
                case -356:
                    switch 32768 {
                    case 0:
                        if false {
                            v4 = 0.1
                        } else {
                            for fi77 := 0; fi77 < 1; fi77++ {
                                if true {
                                    if true {
                                        fmt.Println(v4, fi77)
                                    } else {
                                        v78 := GenS6{F0: 1, F1: 2147483648}
                                        var v79 uint64 = 0
                                        _ = v78
                                        _ = v79
                                    }
                                    v80 := GenS7{F0: 255}
                                    _ = v80
                                }
                                for fi81 := 0; fi81 < 13; fi81++ {
                                    v1 = "a"
                                    _ = fi81
                                }
                                delete(v2, ("a" + v1))
                                _ = fi77
                            }
                            for fi82 := 0; fi82 < 14; fi82++ {
                                v83 := GenS8{F0: 42}
                                switch (fi82 * 1) {
                                case -9223372036854775808:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    for fi84 := 0; fi84 < 20; fi84++ {
                                        _ = len(v5)
                                        if (!(128 <= 60)) {
                                            fmt.Println(v11)
                                        } else {
                                            if ((9223372036854775807 == -652) && (889 >= 0)) {
                                                v11 = fi84
                                            }
                                            var v85 int16 = 191
                                            switch (v11 - v11) {
                                            case 264:
                                                v86 := GenS9{F0: 586, F1: -429, F2: -1}
                                                _ = v86
                                            case 884:
                                                if (-1 >= 42) {
                                                    var v87 uint = 294
                                                    _ = v87
                                                }
                                            case 405:
                                                if true {
                                                    v1 = (v1 + "hello")
                                                }
                                            }
                                            v11 = fi82
                                            if (2 > 2147483648) {
                                                var v88 uint32 = 803
                                                _ = v88
                                            }
                                            if true {
                                                v1 = v1
                                            }
                                            if ((v1 + (v1 + v1)) > "hello") {
                                                v4 = 1.5
                                                var v89 int64 = -626
                                                v1 = ""
                                                _ = v89
                                            } else {
                                                var v90 string = (((("hello" + "\n\t") + v1) + v1) + "\n\t")
                                                _ = v90
                                            }
                                            _ = v85
                                        }
                                        var v91 uint64 = 18446744073709551615
                                        _ = fi84
                                        _ = v91
                                    }
                                }
                                v92 := GenS10{F0: -1, F1: 9223372036854775807, F2: 411}
                                _ = fi82
                                _ = v83
                                _ = v92
                            }
                            var v93 float64 = (1.5 / 3.14159)
                            _ = v93
                        }
                    case 32768:
                        for fi94 := 0; fi94 < 3; fi94++ {
                            v3 = append(v3, ("\n\t" + (v1 + v1)))
                            _ = fi94
                        }
                    }
                case 56:
                    var v95 uint8 = 62
                    _ = v95
                default:
                    v96 := GenS11{F0: "", F1: int8(v11), F2: 603}
                    _ = v96
                }
            }
            if true {
                var v97 float64 = (float64(v4) - (0.1 / -2.25))
                _ = v97
            }
        }
    }
    var v98 float64 = (0.1 * (3.14159 / 1.0))
    if (0 < 68) {
        for fi99 := 0; fi99 < 20; fi99++ {
            switch fi99 {
            case -9223372036854775808:
                v4 -= v4
            case 2147483647:
                for fi100 := 0; fi100 < 8; fi100++ {
                    var v101 uint8 = 255
                    _ = fi100
                    _ = v101
                }
            default:
                v1 = v1
            }
            var v102 uint16 = 255
            var v103 float32 = float32((0.0 / v98))
            _ = fi99
            _ = v102
            _ = v103
        }
        var v104 bool = ((0 >= 415) || true)
        _ = v104
    } else {
        for fi105 := 0; fi105 < 15; fi105++ {
            if true {
                var v106 int64 = -858
                var v107 uint8 = 1
                v108 := GenS12{F0: 65535}
                _ = v106
                _ = v107
                _ = v108
            } else {
                v2[v1] = 2
            }
            _ = fi105
        }
        v2[v1] = 222
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
}
