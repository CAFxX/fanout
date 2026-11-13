package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint16; F1 int; F2 uint8 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint16 = 42
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 2
    var v4 int32 = 1
    var v5 int32 = -1
    var v6 int64 = (-(v2))
    fmt.Println(v2, v4, v4)
    if (14 >= 0) {
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: (fi7 & fi7), F1: v1, F2: (v3 & (v3 * 18446744073709551615))}
            var v9 bool = false
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS2{F0: 255, F1: 2147483647, F2: 1}
    var v11 int32 = v5
    v4 = v11
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint64 = 786
        v2 = (v2 | (v6 / ((int64(v13)) | 1)))
        _ = fi12
        _ = v13
    }
    v14 := []int{770, int(v6), 2147483647, 536}
    if false {
        var v15 uint = 65535
        var v16 float32 = 1.5
        _ = v15
        _ = v16
    }
    if true {
        var v17 uint8 = 2
        if (42 != 2) {
            var v18 float64 = (-2.25 / (-2.25 / 1.5))
            var v19 uint8 = v17
            _ = v18
            _ = v19
        } else {
            switch 32767 {
            case 128:
                var v20 int64 = (v6 & v6)
                _ = v20
            case 929:
                var v21 int32 = 0
                _ = v21
            }
            var v22 int16 = 1000
            if true {
                v22 -= -1
            } else {
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 uint8 = v17
                    if true {
                        for fi25 := 0; fi25 < 12; fi25++ {
                            var v26 float32 = -2.25
                            if true {
                                var v27 int32 = v11
                                _ = v27
                            } else {
                                var v28 int16 = v22
                                _ = v28
                            }
                            _ = fi25
                            _ = v26
                        }
                        v6 = 0
                    }
                    _ = fi23
                    _ = v24
                }
            }
            _ = v22
        }
        _ = v17
    }
    if (v6 != v2) {
        v2 = -116
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 bool = (int8(v4) > 127)
            var v31 float64 = 1e10
            switch "\n\t" {
            case "a":
                fmt.Println(fi29, v5)
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        for fi32 := 0; fi32 < 9; fi32++ {
            var v33 uint = 2
            var v34 uint64 = 2
            fmt.Println(v11, v1, fi32)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 int16 = -248
    var v36 bool = (((("hello" + "hello") + ((("" + (("foo bar" + ("a" + "hello")) + "foo bar")) + ((((("\n\t" + "foo bar") + "") + "\n\t") + "") + (("foo bar" + ("\n\t" + "\n\t")) + "foo bar"))) + "\n\t")) + "\n\t") > "")
    if v36 {
        fmt.Println(v2, v6, v4)
        if (true && false) {
            fmt.Println(v5)
            fmt.Println(v11, v35)
        } else {
            fmt.Println(v11, v4)
            var v37 int8 = -64
            _ = v37
        }
    } else {
        if ((155 <= 894) && (!(1 < 249))) {
            var v38 uint = 18446744073709551615
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 float64 = ((3.14159 * 0.0) / -2.25)
                if true {
                    v3 = (v3 / ((42) | 1))
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v6, v6)
                    case "hello":
                        if v36 {
                            v1 -= (v1 << ((v38) & 63))
                            var v41 uint64 = (v3 * (v3 & v3))
                            v36 = v36
                            _ = v41
                        }
                    case "a":
                        fmt.Println(v38)
                    default:
                        if v36 {
                            v42 := GenS3{F0: 1, F1: 513, F2: -1}
                            for fi43 := 0; fi43 < 2; fi43++ {
                                var v44 uint8 = 218
                                var v45 bool = v36
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                            fmt.Println(v6, v38)
                            _ = v42
                        } else {
                            if true {
                                _ = len(v14)
                            } else {
                                if v36 {
                                    v1 += (v1 / ((v1) | 1))
                                    switch (fi39 << ((v38) & 63)) {
                                    case 42:
                                        for fi46 := 0; fi46 < 10; fi46++ {
                                            v36 = ((fi39 + 4294967295) != fi46)
                                            if len(v14) > 0 {
                                                v14[0] = (fi39 << ((v38) & 63))
                                            }
                                            v5 -= (v11 | 0)
                                            _ = fi46
                                        }
                                    case 870:
                                        for fi47 := 0; fi47 < 2; fi47++ {
                                            v14 = append(v14, (fi39 % (((^(fi39))) | 1)))
                                            var v48 uint16 = v1
                                            if (v36 && ((v4 >> ((0) & 63)) != 0)) {
                                                fmt.Println(v38, v11, v48)
                                                v48 += 65535
                                            } else {
                                                if (!v36) {
                                                    if true {
                                                        if (!((1.5 + -2.25) < 0.1)) {
                                                            var v49 int32 = v11
                                                            _ = v49
                                                        }
                                                        var v50 int8 = -128
                                                        v51 := GenS3{F0: (v3 ^ 410), F1: (v1 - v48), F2: -128}
                                                        v52 := GenS1{F0: (fi47 << (((v38 & 65535)) & 63)), F1: v1, F2: v3}
                                                        v11 -= v5
                                                        v1 += v48
                                                        v35 += (v35 * (v35 / (((v35 ^ -32768)) | 1)))
                                                        _ = v50
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    if (v36 || v36) {
                                                        fmt.Println(v1, v40, v48)
                                                    }
                                                    fmt.Println(v4, v38)
                                                }
                                                v53 := GenS4{F0: v48, F1: (fi47 | 0), F2: 70}
                                                _ = v53
                                            }
                                            _ = fi47
                                            _ = v48
                                        }
                                    default:
                                        for fi54 := 0; fi54 < 17; fi54++ {
                                            var v55 uint32 = 2
                                            _ = fi54
                                            _ = v55
                                        }
                                    }
                                } else {
                                    fmt.Println(v35)
                                }
                                _ = len(v14)
                            }
                            var v56 uint64 = (v3 * v3)
                            for fi57 := 0; fi57 < 5; fi57++ {
                                v11 += v5
                                var v58 string = ""
                                _ = fi57
                                _ = v58
                            }
                            _ = v56
                        }
                    }
                    fmt.Println(v11)
                }
                _ = fi39
                _ = v40
            }
            _ = v38
        }
    }
    if (!(42 < (v1 ^ (^(v1))))) {
        if true {
            switch -1 {
            case 32768:
                var v59 int32 = v11
                _ = v59
            }
            var v60 uint32 = 65535
            _ = v60
        }
        for fi61 := 0; fi61 < 19; fi61++ {
            if true {
                for fi62 := 0; fi62 < 13; fi62++ {
                    v11 = v4
                    switch "\n\t" {
                    case "":
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    default:
                        switch ("" + ((("\n\t" + ("hello" + ("foo bar" + ("" + "hello")))) + "hello") + "hello")) {
                        case "":
                            var v63 string = (("foo bar" + "a") + "")
                            _ = v63
                        }
                    }
                    var v64 uint8 = 167
                    _ = fi62
                    _ = v64
                }
                if v36 {
                    v6 -= v6
                    fmt.Println(v6)
                } else {
                    v11 = v5
                    if v36 {
                        fmt.Println(v3, v35)
                        v5 = -811
                        var v65 int = (fi61 << ((42) & 63))
                        _ = v65
                    } else {
                        v35 = 1
                    }
                    var v66 uint64 = v3
                    _ = v66
                }
                if true {
                    if v36 {
                        if v36 {
                            fmt.Println(v36, v4, v11)
                            for fi67 := 0; fi67 < 4; fi67++ {
                                v36 = true
                                var v68 float64 = 0.1
                                v4 += 42
                                _ = fi67
                                _ = v68
                            }
                        }
                        v69 := GenS5{F0: 42, F1: -1}
                        v36 = false
                        _ = v69
                    }
                } else {
                    if (!(((-2.25 / (3.14159 / 0.1)) + ((0.1 / -2.25) - -2.25)) > 3.14159)) {
                        switch fi61 {
                        case 9223372036854775807:
                            v5 = int32(v3)
                        case -244:
                            if (float64(1e10) >= ((float64((3.14159 / 1e10)) - -2.25) - 1e10)) {
                                if (746 > (v3 & (^(v3)))) {
                                    if false {
                                        switch "foo bar" {
                                        case "a":
                                            if true {
                                                var v70 float64 = float64((0.0 - -2.25))
                                                if v36 {
                                                    switch "" {
                                                    case "hello":
                                                        for fi71 := 0; fi71 < 9; fi71++ {
                                                            var v72 int16 = 1000
                                                            var v73 int = (fi61 * (-(fi71)))
                                                            _ = fi71
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    }
                                                    var v74 bool = v36
                                                    _ = v74
                                                }
                                                _ = v70
                                            } else {
                                                if (v1 >= (^(v1))) {
                                                    var v75 int = fi61
                                                    var v76 int32 = v4
                                                    var v77 int16 = v35
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                } else {
                                                    v3 += (v3 * (^(v3)))
                                                }
                                            }
                                        case "":
                                            v4 -= v5
                                        default:
                                            for fi78 := 0; fi78 < 14; fi78++ {
                                                v4 += v4
                                                fmt.Println(v6, v3, v6)
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                                _ = fi78
                                            }
                                        }
                                        switch fi61 {
                                        case 127:
                                            v4 = -1
                                        case 785:
                                            var v79 uint16 = v1
                                            _ = v79
                                        default:
                                            for fi80 := 0; fi80 < 8; fi80++ {
                                                var v81 uint32 = 2147483647
                                                if (v35 >= -1) {
                                                    v4 += -2147483648
                                                    var v82 uint = 4294967295
                                                    fmt.Println(v2, v82, fi61)
                                                    var v83 uint = 255
                                                    v84 := GenS6{F0: 1}
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    var v85 string = ""
                                                    _ = v85
                                                }
                                                _ = fi80
                                                _ = v81
                                            }
                                        }
                                        for fi86 := 0; fi86 < 4; fi86++ {
                                            v5 = 2147483647
                                            v87 := GenS7{F0: v36, F1: int64(v11)}
                                            switch ((("a" + "") + ("\n\t" + (("a" + "\n\t") + (("" + "a") + ("foo bar" + "a"))))) + "") {
                                            case "a":
                                                v88 := GenS6{F0: 2147483648}
                                                _ = v88
                                            default:
                                                var v89 uint32 = 265
                                                _ = v89
                                            }
                                            _ = fi86
                                            _ = v87
                                        }
                                    }
                                    var v90 string = ("a" + "hello")
                                    _ = v90
                                }
                            } else {
                                v91 := GenS6{F0: 424}
                                _ = v91
                            }
                        default:
                            var v92 uint64 = 18446744073709551615
                            _ = v92
                        }
                    } else {
                        v6 = (v6 % ((v2) | 1))
                    }
                }
            } else {
                for fi93 := 0; fi93 < 15; fi93++ {
                    _ = len(v14)
                    if v36 {
                        var v94 uint = 324
                        switch fi61 {
                        case 4294967295:
                            if true {
                                v36 = true
                                if v36 {
                                    _ = len(v14)
                                    v3 += 18446744073709551615
                                    switch fi61 {
                                    case 32768:
                                        v2 += int64(v1)
                                    case -9223372036854775808:
                                        v5 -= (v5 & (v5 | -2147483648))
                                    }
                                }
                                v4 -= v11
                            } else {
                                fmt.Println(v35, v3)
                            }
                        case -854:
                            v95 := GenS8{F0: -9223372036854775808, F1: (^(v3))}
                            _ = v95
                        }
                        var v96 uint8 = 95
                        _ = v94
                        _ = v96
                    }
                    _ = fi93
                }
                var v97 int = fi61
                _ = v97
            }
            _ = fi61
        }
        fmt.Println(v3, v2, v1)
    }
    if (uint8(v1) <= 128) {
        v1 += (v1 - 42)
    }
    fmt.Println(v4, v3)
    v1 += (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
