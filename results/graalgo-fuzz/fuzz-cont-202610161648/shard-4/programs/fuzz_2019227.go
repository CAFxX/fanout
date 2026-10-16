package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 int64; F1 uint8 }
type GenS9 struct { F0 uint8; F1 uint64; F2 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = 0.0
    var v2 uint8 = 128
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := GenS1{F0: -2147483648, F1: 9223372036854775807, F2: -78}
    fmt.Println(v1)
    v6 := v4["hello"]
    var v7 int64 = 9223372036854775807
    var v8 int16 = 42
    v9 := v4[(("hello" + "") + "a")]
    var v10 uint16 = 797
    switch 32767 {
    case 9223372036854775807:
        if v6 {
            var v11 float64 = 0.0
            _ = v11
        }
    case 0:
        switch 2147483647 {
        case 32767:
            switch 998 {
            case -9223372036854775808:
                var v12 int16 = (v8 + v8)
                _ = v12
            default:
                var v13 bool = ((false || (1 != -53)) && v9)
                _ = v13
            }
        case 981:
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := GenS1{F0: 1, F1: (fi14 << ((721) & 63)), F2: 127}
                _ = fi14
                _ = v15
            }
        }
    default:
        var v16 float64 = float64((v1 - float32(float64(0.0))))
        _ = v16
    }
    v2 = v2
    v4["a"] = (v7 != v7)
    var v17 float64 = float64(v1)
    switch (("hello" + "a") + "foo bar") {
    case "hello":
        fmt.Println(v2, v17, v9)
    case "":
        if (!false) {
            _ = len(v4)
            var v18 int8 = 127
            v8 = -1
            _ = v18
        } else {
            if true {
                var v19 int32 = 2147483647
                switch "foo bar" {
                case "a":
                    if (!(42 != 42)) {
                        fmt.Println(v2, v10, v19)
                        var v20 int16 = (v8 + (-(v8)))
                        _ = v20
                    } else {
                        v21 := GenS1{F0: (v19 / (((v19 << ((435) & 63))) | 1)), F1: 426, F2: 42}
                        switch 1 {
                        case 9223372036854775807:
                            if false {
                                for fi22 := 0; fi22 < 1; fi22++ {
                                    var v23 float32 = (float32((v17 / 1.5)) * v1)
                                    var v24 bool = (!true)
                                    if (float64(v1) >= float64(3.14159)) {
                                        v19 = -2147483648
                                        var v25 int8 = 60
                                        var v26 uint64 = uint64(v10)
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = fi22
                                    _ = v23
                                    _ = v24
                                }
                                v7 += v7
                            }
                        case 32767:
                            v27 := GenS2{F0: 0}
                            _ = v27
                        case 42:
                            var v28 uint8 = 148
                            _ = v28
                        default:
                            v9 = ((0.0 * ((v17 / ((v17 / v17) * v17)) - v17)) != (-2.25 + ((v17 / v17) * 0.1)))
                        }
                        _ = v21
                    }
                case "hello":
                    switch int(v10) {
                    case -416:
                        if v6 {
                            v6 = true
                        } else {
                            fmt.Println(v1, v10)
                            var v29 float32 = (((1.5 - (v1 * 0.0)) * v1) + (0.1 + ((1e10 * 3.14159) + (0.0 * v1))))
                            _ = v29
                        }
                    default:
                        var v30 uint16 = (^(v10))
                        _ = v30
                    }
                case "\n\t":
                    if false {
                        switch 127 {
                        case 394:
                            if v6 {
                                v9 = (!v6)
                                switch 32767 {
                                case 321:
                                    fmt.Println(v6, v9)
                                case 2147483647:
                                    var v31 uint16 = 2
                                    _ = v31
                                case 143:
                                    var v32 uint = 618
                                    _ = v32
                                default:
                                    var v33 float32 = float32(-2.25)
                                    _ = v33
                                }
                                v10 = 899
                            } else {
                                var v34 string = ("\n\t" + "foo bar")
                                var v35 float32 = float32(0.0)
                                fmt.Println(v8)
                                _ = v34
                                _ = v35
                            }
                        case 32768:
                            switch -451 {
                            case -1:
                                switch 42 {
                                case 529:
                                    if v6 {
                                        var v36 string = ("" + ("" + ("hello" + "foo bar")))
                                        _ = len(v3)
                                        switch 127 {
                                        case 42:
                                            v17 = float64(1.5)
                                        }
                                        _ = v36
                                    }
                                case 313:
                                    var v37 int16 = (v8 >> ((2) & 63))
                                    _ = v37
                                default:
                                    for fi38 := 0; fi38 < 6; fi38++ {
                                        v1 -= v1
                                        _ = fi38
                                    }
                                }
                            default:
                                if v6 {
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        var v40 uint32 = 4294967295
                                        _ = fi39
                                        _ = v40
                                    }
                                    for fi41 := 0; fi41 < 14; fi41++ {
                                        var v42 int64 = (^(v7))
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            if (2147483647 <= 1) {
                                                v4["foo bar"] = false
                                                v2 -= v2
                                                v1 -= 1e10
                                                if len(v3) > 0 {
                                                    v3[0] = "foo bar"
                                                }
                                            }
                                            switch "\n\t" {
                                            case "":
                                                v8 = -32768
                                            case "hello":
                                                fmt.Println(v42)
                                            default:
                                                var v44 uint32 = 0
                                                _ = v44
                                            }
                                            if true {
                                                v45 := GenS2{F0: (v10 + 0)}
                                                v46 := GenS3{F0: v10, F1: v19}
                                                if (-119 != 1) {
                                                    if v9 {
                                                        if v9 {
                                                            fmt.Println(fi43)
                                                            v10 -= v10
                                                        } else {
                                                            v2 += (^(v2))
                                                            fmt.Println(v19, v7)
                                                            fmt.Println(v7, v1, v9)
                                                        }
                                                        switch (("a" + "hello") + (("" + ("foo bar" + "a")) + (((("a" + "foo bar") + "") + ("foo bar" + ("foo bar" + "a"))) + ("a" + "")))) {
                                                        case "\n\t":
                                                            var v47 float32 = (((0.0 / v1) / 1.5) * float32(v17))
                                                            _ = v47
                                                        }
                                                    } else {
                                                        var v48 int16 = -32768
                                                        if (624 > 2) {
                                                            var v49 int = 128
                                                            switch "" {
                                                            case "foo bar":
                                                                var v50 uint16 = (v10 | v10)
                                                                _ = v50
                                                            case "hello":
                                                                var v51 uint8 = 255
                                                                _ = v51
                                                            default:
                                                                switch (((("foo bar" + "\n\t") + ("hello" + "foo bar")) + "a") + ("foo bar" + "")) {
                                                                case "":
                                                                    v19 += 453
                                                                default:
                                                                    if ((v48 | (v8 / (((-(v48))) | 1))) == v8) {
                                                                        v10 -= 788
                                                                    }
                                                                }
                                                                var v52 uint64 = 1
                                                                v53 := v4["hello"]
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            _ = v49
                                                        }
                                                        if v6 {
                                                            v7 += 648
                                                            if (fi43 < fi41) {
                                                                v54 := GenS4{F0: (v7 >> ((18446744073709551615) & 63))}
                                                                if (!(v6 && v9)) {
                                                                    if true {
                                                                        fmt.Println(v1, fi41)
                                                                        v9 = (v10 < uint16(v2))
                                                                    } else {
                                                                        if v9 {
                                                                            var v55 int64 = -9223372036854775808
                                                                            _ = v55
                                                                        } else {
                                                                            var v56 int8 = 1
                                                                            v57 := GenS3{F0: (v10 << ((4294967295) & 63)), F1: (v19 + (v19 - 553))}
                                                                            if (!true) {
                                                                                v58 := GenS1{F0: (v19 & v19), F1: fi41, F2: 42}
                                                                                _ = len(v4)
                                                                                _ = v58
                                                                            }
                                                                            v1 = v1
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                        if (2 <= 0) {
                                                                            if ((42 >= 2) && ((-(v8)) == 1000)) {
                                                                                v59 := GenS1{F0: v19, F1: (^(fi41)), F2: -8}
                                                                                _ = v59
                                                                            } else {
                                                                                v17 = 3.14159
                                                                            }
                                                                            switch ((("a" + (("\n\t" + ("" + "foo bar")) + "")) + "a") + ("hello" + "foo bar")) {
                                                                            case "":
                                                                                v7 = int64(v2)
                                                                            case "foo bar":
                                                                                v48 -= v8
                                                                            case "hello":
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            default:
                                                                                v60 := GenS5{F0: 293}
                                                                                _ = v60
                                                                            }
                                                                        }
                                                                        var v61 string = ("\n\t" + "\n\t")
                                                                        _ = v61
                                                                    }
                                                                    var v62 uint16 = v10
                                                                    _ = v62
                                                                }
                                                                _ = v54
                                                            } else {
                                                                if v6 {
                                                                    v63 := GenS6{F0: (v42 & (v42 ^ v42)), F1: v42}
                                                                    var v64 uint64 = 88
                                                                    _ = v63
                                                                    _ = v64
                                                                } else {
                                                                    var v65 int64 = 946
                                                                    _ = v65
                                                                }
                                                                v66 := GenS7{F0: int(v19), F1: v2}
                                                                _ = v66
                                                            }
                                                            var v67 uint64 = 1
                                                            if (v6 || false) {
                                                                var v68 string = "hello"
                                                                if (!v6) {
                                                                    v7 = v7
                                                                    fmt.Println(v7)
                                                                    v17 = v17
                                                                } else {
                                                                    if (fi41 < fi43) {
                                                                        v69 := GenS8{F0: v42, F1: (v2 ^ v2)}
                                                                        var v70 string = v68
                                                                        if v6 {
                                                                            if ((v2 % ((42) | 1)) >= (v2 & v2)) {
                                                                                v71 := GenS9{F0: v2, F1: (v67 ^ uint64(v2)), F2: fi41}
                                                                                fmt.Println(v42, v10)
                                                                                v48 += v48
                                                                                _ = v71
                                                                            } else {
                                                                                v48 = int16(fi43)
                                                                            }
                                                                            v48 = (v8 * (v8 / ((int16(v7)) | 1)))
                                                                            var v72 int32 = (v19 + (v19 & v19))
                                                                            _ = v72
                                                                        } else {
                                                                            v68 = ((v70 + "") + v70)
                                                                            if (((3.14159 * v1) + -2.25) > v1) {
                                                                                switch v70 {
                                                                                case "foo bar":
                                                                                    var v73 uint8 = v2
                                                                                    _ = v73
                                                                                case "\n\t":
                                                                                    v67 += 0
                                                                                case "":
                                                                                    fmt.Println(v1, v42, v9)
                                                                                default:
                                                                                    v4[(v68 + "hello")] = v6
                                                                                    switch ("\n\t" + (v70 + ((v70 + v68) + "hello"))) {
                                                                                    case "":
                                                                                        var v74 string = v68
                                                                                        _ = v74
                                                                                    default:
                                                                                        switch (((("hello" + (v68 + (v70 + v68))) + "\n\t") + v68) + ("hello" + v70)) {
                                                                                        case "a":
                                                                                            delete(v4, "\n\t")
                                                                                        }
                                                                                    }
                                                                                }
                                                                                var v75 bool = (float64(v1) <= (3.14159 + v17))
                                                                                _ = v75
                                                                            } else {
                                                                                _ = len(v3)
                                                                                fmt.Println(v42, v9)
                                                                                v76 := GenS2{F0: v10}
                                                                                _ = v76
                                                                            }
                                                                        }
                                                                        _ = v69
                                                                        _ = v70
                                                                    } else {
                                                                        v77 := GenS5{F0: v67}
                                                                        v9 = true
                                                                        if (v6 && (655 > 0)) {
                                                                            fmt.Println(fi41, v10)
                                                                            var v78 int8 = 83
                                                                            switch 264 {
                                                                            case 42:
                                                                                delete(v4, (("a" + (v68 + "hello")) + ((v68 + v68) + (v68 + "a"))))
                                                                            case 373:
                                                                                delete(v4, (v68 + v68))
                                                                            case -334:
                                                                                v79 := GenS6{F0: 0, F1: -891}
                                                                                _ = v79
                                                                            default:
                                                                                v80 := GenS10{F0: (v48 >> ((18446744073709551615) & 63))}
                                                                                _ = v80
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                        _ = v77
                                                                    }
                                                                    if (!v9) {
                                                                        v81 := v4["hello"]
                                                                        _ = v81
                                                                    } else {
                                                                        v7 -= (-(v7))
                                                                        v82 := GenS11{F0: 4294967295, F1: v10}
                                                                        _ = v82
                                                                    }
                                                                }
                                                                fmt.Println(v17, v68, v1)
                                                                _ = v68
                                                            }
                                                            _ = v67
                                                        } else {
                                                            var v83 float64 = v17
                                                            v9 = v6
                                                            var v84 float32 = (float32(float64((v1 / 3.14159))) / (v1 / (((1e10 - 3.14159) * 0.0) + 3.14159)))
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        _ = v48
                                                    }
                                                    if true {
                                                        v85 := GenS12{F0: v42}
                                                        v7 += 9223372036854775807
                                                        _ = v85
                                                    }
                                                    v86 := GenS13{F0: (v42 & v42), F1: uint64(v10)}
                                                    _ = v86
                                                }
                                                _ = v45
                                                _ = v46
                                            } else {
                                                v2 += v2
                                            }
                                            _ = fi43
                                        }
                                        v2 = v2
                                        _ = fi41
                                        _ = v42
                                    }
                                }
                            }
                        case -995:
                            v2 -= 127
                        }
                        _ = len(v3)
                    }
                }
                var v87 uint8 = v2
                _ = v19
                _ = v87
            }
            v6 = true
        }
    }
    _ = len(v4)
    for fi88 := 0; fi88 < 13; fi88++ {
        var v89 uint16 = 593
        for fi90 := 0; fi90 < 20; fi90++ {
            _ = len(v4)
            v17 += 0.1
            _ = fi90
        }
        _ = fi88
        _ = v89
    }
    switch -837 {
    case 4294967295:
        for fi91 := 0; fi91 < 20; fi91++ {
            for fi92 := 0; fi92 < 5; fi92++ {
                v93 := GenS13{F0: (^(v7)), F1: 1}
                v6 = (v9 && v6)
                _ = len(v4)
                _ = fi92
                _ = v93
            }
            _ = fi91
        }
    default:
        for fi94 := 0; fi94 < 14; fi94++ {
            if (!v9) {
                fmt.Println(v17)
                if len(v3) > 0 {
                    v3[0] = ((("\n\t" + "foo bar") + "\n\t") + "a")
                }
                _ = len(v3)
            }
            fmt.Println(v2, v2, v9)
            _ = fi94
        }
    }
    for fi95 := 0; fi95 < 14; fi95++ {
        v96 := GenS5{F0: 2}
        switch (-(fi95)) {
        case -9223372036854775808:
            v4["\n\t"] = v6
        case 32767:
            fmt.Println(v17, v6)
        }
        _ = fi95
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
