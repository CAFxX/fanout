package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int32; F1 uint; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 string; F1 uint8; F2 uint16 }
type GenS10 struct { F0 int8; F1 int; F2 string }
type GenS11 struct { F0 int16; F1 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0(p0 bool) int16 {
    var v7 bool = true
    v7 = (509 >= 2147483648)
    if p0 {
        p0 = false
        v7 = (!(0.1 == 1e10))
        p0 = v7
    }
    v8 := GenS1{F0: 117, F1: -1}
    _ = p0
    _ = v7
    _ = v8
    return 1000
}
func genFn1() int32 {
    v9 := []int{865, 863}
    var v10 int8 = 100
    switch 32767 {
    case 32767:
        var v11 uint64 = 2
        _ = v11
    case 0:
        if (680 < 1) {
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v10, v10)
                v10 += v10
                _ = fi12
            }
            if (!true) {
                if (true || true) {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            v10 += (-(v10))
                            var v15 int16 = 397
                            var v16 int8 = (v10 + v10)
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        v10 += v10
                        switch fi13 {
                        case 4294967295:
                            if len(v9) > 0 {
                                v9[0] = 42
                            }
                        case 32767:
                            v17 := GenS2{F0: 498, F1: 270, F2: v10}
                            _ = v17
                        case -1:
                            var v18 int32 = -1
                            _ = v18
                        default:
                            switch (fi13 << ((255) & 63)) {
                            case -434:
                                if false {
                                    switch "" {
                                    case "\n\t":
                                        if true {
                                            v19 := GenS2{F0: 42, F1: 103, F2: (v10 - v10)}
                                            _ = v19
                                        } else {
                                            _ = len(v9)
                                        }
                                    default:
                                        v20 := GenS3{F0: 224, F1: 42}
                                        _ = v20
                                    }
                                } else {
                                    for fi21 := 0; fi21 < 1; fi21++ {
                                        var v22 bool = (100 <= 132)
                                        fmt.Println(v22)
                                        v22 = false
                                        v22 = (v22 || ((fi13 | (fi21 - 634)) >= 2147483647))
                                        _ = fi21
                                        _ = v22
                                    }
                                }
                            case -1:
                                var v23 int16 = 456
                                _ = v23
                            case -2147483648:
                                v24 := GenS4{F0: ("hello" + "foo bar"), F1: ("foo bar" + "hello")}
                                _ = v24
                            default:
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    if true {
                                        if false {
                                            if len(v9) > 0 {
                                                v9[0] = (fi13 << ((574) & 63))
                                            }
                                            if true {
                                                var v26 float64 = 3.14159
                                                _ = v26
                                            } else {
                                                if (491 != 42) {
                                                    v10 = (v10 % (((v10 & -1)) | 1))
                                                    if false {
                                                        if true {
                                                            fmt.Println(v10, fi13, fi25)
                                                            v27 := GenS5{F0: -443, F1: 18446744073709551615, F2: 1}
                                                            _ = v27
                                                        } else {
                                                            if (-9223372036854775808 == -9223372036854775808) {
                                                                v10 += (v10 % ((v10) | 1))
                                                                if (1 == -553) {
                                                                    _ = len(v9)
                                                                    var v28 uint = 4294967295
                                                                    fmt.Println(v28, v28, fi25)
                                                                    if (!(-2.25 != 1e10)) {
                                                                        v29 := GenS6{F0: 508}
                                                                        if ("a" > "foo bar") {
                                                                            switch (fi25 + (fi13 - (fi13 * (fi13 & fi13)))) {
                                                                            case 127:
                                                                                var v30 float32 = (1e10 + -2.25)
                                                                                _ = v30
                                                                            case 277:
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if true {
                                                                                        v31 := GenS7{F0: fi13, F1: 42}
                                                                                        if (1 == 65535) {
                                                                                            var v32 int32 = 0
                                                                                            if len(v9) >= 2 {
                                                                                                v9 = v9[1:len(v9)]
                                                                                            }
                                                                                            var v33 uint = v28
                                                                                            _ = v32
                                                                                            _ = v33
                                                                                        }
                                                                                        _ = v31
                                                                                    }
                                                                                case "foo bar":
                                                                                    if len(v9) > 0 {
                                                                                        v9[0] = fi25
                                                                                    }
                                                                                }
                                                                            }
                                                                            v28 -= v28
                                                                        } else {
                                                                            _ = len(v9)
                                                                            var v34 int64 = 9223372036854775807
                                                                            v10 -= v10
                                                                            var v35 float64 = 1.5
                                                                            v10 += (v10 ^ -1)
                                                                            _ = v34
                                                                            _ = v35
                                                                        }
                                                                        if true {
                                                                            v36 := GenS4{F0: ("hello" + "foo bar"), F1: "foo bar"}
                                                                            _ = v36
                                                                        }
                                                                        if true {
                                                                            v37 := GenS8{F0: 128, F1: 32768}
                                                                            var v38 uint32 = 65535
                                                                            v39 := GenS9{F0: ("" + "\n\t"), F1: 1, F2: 1}
                                                                            _ = v37
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        var v40 uint8 = 42
                                                                        _ = v29
                                                                        _ = v40
                                                                    }
                                                                    v41 := GenS5{F0: -481, F1: (v28 & 65535), F2: 42}
                                                                    _ = v28
                                                                    _ = v41
                                                                } else {
                                                                    v42 := GenS1{F0: 42, F1: v10}
                                                                    _ = v42
                                                                }
                                                            } else {
                                                                var v43 float32 = float32(1.5)
                                                                _ = len(v9)
                                                                v43 = (1.5 / (((float32((0.0 / 3.14159)) + 3.14159) + (0.1 - float32(((3.14159 + (3.14159 - 0.1)) * (1.5 - -2.25))))) * v43))
                                                                _ = v43
                                                            }
                                                            if false {
                                                                var v44 float32 = ((((((-2.25 / 3.14159) * (1.5 / 0.1)) / 3.14159) / (float32(-2.25) / 0.1)) / float32(1.5)) / 1e10)
                                                                _ = v44
                                                            } else {
                                                                if false {
                                                                    var v45 int = fi25
                                                                    _ = v45
                                                                } else {
                                                                    var v46 int16 = -1
                                                                    _ = v46
                                                                }
                                                            }
                                                        }
                                                        v10 += (v10 | -1)
                                                        if true {
                                                            if (704 > 333) {
                                                                v10 = (v10 >> ((255) & 63))
                                                            }
                                                        } else {
                                                            v10 = v10
                                                        }
                                                        fmt.Println(fi13, fi25, fi13)
                                                    } else {
                                                        var v47 float32 = (-2.25 + 0.1)
                                                        if (false && false) {
                                                            var v48 uint8 = 0
                                                            v48 -= v48
                                                            _ = v48
                                                        } else {
                                                            _ = len(v9)
                                                            var v49 bool = false
                                                            var v50 float32 = (float32(((3.14159 - 3.14159) - 1e10)) + (float32(0.0) / float32((1.5 - 0.1))))
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        _ = v47
                                                    }
                                                    if true {
                                                        fmt.Println(fi25, v10)
                                                    }
                                                    if true {
                                                        var v51 bool = (2 > 4294967295)
                                                        var v52 int = fi13
                                                        switch (fi13 >> ((65535) & 63)) {
                                                        case 1:
                                                            if v51 {
                                                                fmt.Println(v51, fi13)
                                                            }
                                                        default:
                                                            v51 = v51
                                                        }
                                                        v53 := GenS5{F0: -830, F1: 65535, F2: 0}
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                }
                                                _ = len(v9)
                                            }
                                            switch "hello" {
                                            case "":
                                                v10 -= -116
                                            case "foo bar":
                                                if false {
                                                    fmt.Println(v10, v10, v10)
                                                    v9 = append(v9, fi13)
                                                    if ((v10 ^ v10) <= (v10 - (v10 - (^(v10))))) {
                                                        v9 = append(v9, fi13)
                                                    } else {
                                                        var v54 string = "foo bar"
                                                        v54 = ((("hello" + "a") + v54) + v54)
                                                        _ = v54
                                                    }
                                                } else {
                                                    v10 = 1
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                }
                                            }
                                        }
                                        v10 -= -1
                                    } else {
                                        if (!false) {
                                            if false {
                                                if (!(387 >= 18446744073709551615)) {
                                                    v10 += (v10 ^ (v10 << ((1) & 63)))
                                                    switch fi25 {
                                                    case -462:
                                                        fmt.Println(fi13)
                                                    case -1:
                                                        if false {
                                                            v55 := GenS8{F0: 1, F1: (fi13 << ((18446744073709551615) & 63))}
                                                            _ = v55
                                                        }
                                                    }
                                                    switch ("" + "\n\t") {
                                                    case "a":
                                                        fmt.Println(fi25, fi13, fi13)
                                                    }
                                                    var v56 uint8 = 2
                                                    _ = v56
                                                } else {
                                                    var v57 float32 = 1.5
                                                    v10 -= (-(v10))
                                                    _ = v57
                                                }
                                            } else {
                                                v10 = (v10 / (((v10 & (v10 % ((-37) | 1)))) | 1))
                                                v10 -= -1
                                                fmt.Println(fi25, fi25, fi25)
                                            }
                                            v10 -= 127
                                        } else {
                                            _ = len(v9)
                                            var v58 uint = 650
                                            if len(v9) > 0 {
                                                v9[0] = int(v58)
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            v58 += 255
                                            v59 := GenS8{F0: 2, F1: (fi25 | fi13)}
                                            var v60 uint16 = 255
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        if (0 == 199) {
                                            var v61 int16 = 1000
                                            _ = v61
                                        } else {
                                            fmt.Println(fi13)
                                            fmt.Println(fi13, v10, v10)
                                            v62 := GenS5{F0: 175, F1: 4294967295, F2: 4294967295}
                                            if true {
                                                var v63 uint = 734
                                                if false {
                                                    var v64 int16 = -32768
                                                    _ = v64
                                                }
                                                _ = v63
                                            } else {
                                                switch (fi13 << ((65535) & 63)) {
                                                case -990:
                                                    v65 := GenS10{F0: (v10 / ((1) | 1)), F1: (fi25 - fi13), F2: "\n\t"}
                                                    _ = v65
                                                case -957:
                                                    v10 += (-(v10))
                                                case 930:
                                                    v10 += 0
                                                }
                                            }
                                            _ = v62
                                        }
                                    }
                                    v66 := GenS5{F0: 1, F1: 2, F2: 339}
                                    var v67 float32 = 0.0
                                    _ = len(v9)
                                    var v68 int = fi25
                                    if len(v9) > 0 {
                                        v9[0] = fi25
                                    }
                                    _ = fi25
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                }
                            }
                        }
                        _ = fi13
                    }
                } else {
                    if len(v9) > 0 {
                        v9[0] = 0
                    }
                }
                if (32767 >= 0) {
                    v9 = append(v9, -146)
                }
            }
            v69 := GenS11{F0: 0, F1: 413}
            _ = v69
        } else {
            for fi70 := 0; fi70 < 4; fi70++ {
                for fi71 := 0; fi71 < 20; fi71++ {
                    var v72 uint16 = 0
                    _ = fi71
                    _ = v72
                }
                _ = fi70
            }
            switch (((("hello" + "a") + "a") + "") + "hello") {
            case "a":
                fmt.Println(v10)
            case "":
                fmt.Println(v10)
            case "foo bar":
                v73 := GenS12{F0: 1}
                _ = v73
            }
            if (0 != 1) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            } else {
                fmt.Println(v10)
                var v74 uint64 = 42
                _ = v74
            }
        }
    default:
        fmt.Println(v10)
    }
    v10 -= -1
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 uint = 255
    var v2 bool = true
    v3 := []string{"foo bar"}
    v4 := make(map[int]bool)
    var v5 float64 = (-2.25 * 3.14159)
    var v6 uint = v1
    v75 := GenS13{F0: 919, F1: 42, F2: 42}
    v3 = append(v3, "a")
    v76 := GenS4{F0: ((("a" + "") + ("" + "\n\t")) + "hello"), F1: ("a" + "")}
    var v77 int64 = -556
    if v2 {
        var v78 string = ("hello" + "hello")
        fmt.Println(v1, v77, v5)
        var v79 float64 = v5
        _ = v78
        _ = v79
    }
    if (genFn0((!v2)) >= 42) {
        var v80 float32 = ((((1e10 - (0.1 * 0.1)) - (3.14159 + 1e10)) - 0.0) / (((float32(v5) - (1.5 / ((0.1 / 1.0) * 0.1))) - 0.0) - (3.14159 * (-2.25 - 3.14159))))
        if ((2147483647 > -296) && v2) {
            var v81 float64 = v5
            _ = v81
        }
        fmt.Println(v5, v1, v80)
        _ = v80
    } else {
        v82 := v4[-2147483648]
        _ = v82
    }
    delete(v4, 127)
    if v2 {
        for fi83 := 0; fi83 < 8; fi83++ {
            if (false || (v1 != (v1 ^ (v6 - v1)))) {
                switch "a" {
                case "hello":
                    switch ("a" + "foo bar") {
                    case "foo bar":
                        var v84 uint64 = 18446744073709551615
                        _ = v84
                    case "":
                        fmt.Println(fi83)
                    }
                }
            }
            _ = fi83
        }
        _ = len(v3)
    }
    v1 -= (v6 % ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v77
}
