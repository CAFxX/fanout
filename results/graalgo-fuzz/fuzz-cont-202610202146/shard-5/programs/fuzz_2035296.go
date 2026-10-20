package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint; F1 int32; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 197
    var v3 uint16 = (^(v2))
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 309
    fmt.Println(v3, v5, v2)
    fmt.Println(v3, v4, v3)
    if (369 <= 1) {
        switch ("\n\t" + "hello") {
        case "":
            var v6 int = -9223372036854775808
            _ = v6
        case "a":
            v7 := GenS1{F0: ("a" + ""), F1: 127, F2: 242}
            _ = v7
        default:
            for fi8 := 0; fi8 < 18; fi8++ {
                v3 += v3
                _ = fi8
            }
        }
        if (-1 <= 9) {
            var v9 uint = 42
            var v10 int8 = -1
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS1{F0: "hello", F1: -1, F2: 19}
                _ = fi11
                _ = v12
            }
            fmt.Println(v2, v4)
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 uint32 = 65535
                if false {
                    fmt.Println(v3, v2, v2)
                } else {
                    v4 -= v4
                    v15 := GenS1{F0: "\n\t", F1: 42, F2: 255}
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
        }
        switch -909 {
        case 32768:
            var v16 float64 = ((float64(1.5) + float64((float32(((0.0 - 0.1) * 3.14159)) / 1e10))) + (((float64(float32(3.14159)) / 1.5) - 1e10) * 1e10))
            _ = v16
        case 9223372036854775807:
            var v17 float32 = float32(float64((0.0 / (0.1 + 1.5))))
            _ = v17
        }
    } else {
        delete(v1, ("" + "\n\t"))
    }
    var v18 uint = 4294967295
    var v19 int32 = 0
    v20 := []int{-2147483648, 32768, 9223372036854775807}
    var v21 int16 = -1
    if false {
        if false {
            switch (("" + "hello") + "") {
            case "\n\t":
                if (1e10 != (0.0 * (0.1 / -2.25))) {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        v23 := GenS1{F0: "hello", F1: -128, F2: 128}
                        _ = fi22
                        _ = v23
                    }
                }
            case "":
                v19 = v19
            case "hello":
                for fi24 := 0; fi24 < 4; fi24++ {
                    v3 = v2
                    switch "\n\t" {
                    case "a":
                        v25 := GenS1{F0: "", F1: 42, F2: 2}
                        _ = v25
                    case "hello":
                        switch 0 {
                        case 9223372036854775807:
                            if (v2 > (v3 & 32767)) {
                                for fi26 := 0; fi26 < 3; fi26++ {
                                    if (v3 >= (v2 << ((18446744073709551615) & 63))) {
                                        switch "foo bar" {
                                        case "":
                                            _ = len(v20)
                                        default:
                                            if true {
                                                v1[""] = (fi24 - (fi26 >> (((v18 << ((491) & 63))) & 63)))
                                                var v27 uint = 1
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                                _ = v27
                                            } else {
                                                v28 := GenS2{F0: "hello"}
                                                _ = v28
                                            }
                                        }
                                        fmt.Println(v3, fi26, v19)
                                        var v29 int32 = (v19 >> (((v18 + (^(v18)))) & 63))
                                        _ = v29
                                    }
                                    var v30 string = ""
                                    v30 = "hello"
                                    _ = fi26
                                    _ = v30
                                }
                            }
                        case 783:
                            v19 = 0
                        case -272:
                            for fi31 := 0; fi31 < 11; fi31++ {
                                var v32 float32 = (0.0 / 0.1)
                                v33 := GenS1{F0: (("hello" + "hello") + "\n\t"), F1: 0, F2: 255}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    case "\n\t":
                        for fi34 := 0; fi34 < 7; fi34++ {
                            if false {
                                v35 := GenS1{F0: ("foo bar" + "a"), F1: 62, F2: 127}
                                v21 -= (v21 | (-(v21)))
                                _ = v35
                            }
                            switch (fi34 | int(v18)) {
                            case 42:
                                v2 = (v2 * (v2 ^ v3))
                            case -748:
                                switch (fi34 / (((fi34 - fi24)) | 1)) {
                                case 999:
                                    v36 := GenS1{F0: "hello", F1: 127, F2: 128}
                                    _ = v36
                                default:
                                    if false {
                                        v4 += v4
                                        v19 = (v19 ^ -869)
                                        v37 := GenS1{F0: "a", F1: 1, F2: uint8(v5)}
                                        _ = v37
                                    } else {
                                        v2 -= 178
                                    }
                                }
                            case 4294967295:
                                v38 := GenS3{F0: 100}
                                fmt.Println(v4, v18)
                                var v39 uint16 = v3
                                var v40 float64 = 1e10
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            _ = fi34
                        }
                    default:
                        var v41 uint32 = uint32(v2)
                        _ = v41
                    }
                    _ = fi24
                }
            default:
                var v42 int = -750
                _ = v42
            }
            if (v19 >= 42) {
                _ = len(v1)
                switch "foo bar" {
                case "foo bar":
                    v43 := GenS4{F0: 0}
                    _ = v43
                case "hello":
                    v18 = (v18 >> ((v18) & 63))
                }
            } else {
                v44 := GenS4{F0: 255}
                if (true || (v3 < v3)) {
                    if (42 <= -1) {
                        switch "foo bar" {
                        case "hello":
                            _ = len(v1)
                        case "\n\t":
                            var v45 uint32 = 0
                            _ = v45
                        case "foo bar":
                            v4 = (-(v4))
                        default:
                            if (true && false) {
                                var v46 string = "hello"
                                for fi47 := 0; fi47 < 8; fi47++ {
                                    v48 := v1[v46]
                                    v49 := GenS2{F0: "hello"}
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                                for fi50 := 0; fi50 < 13; fi50++ {
                                    var v51 int32 = v19
                                    _ = fi50
                                    _ = v51
                                }
                                _ = v46
                            }
                        }
                    }
                } else {
                    v52 := v1["\n\t"]
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v52, v21)
                    case "hello":
                        for fi53 := 0; fi53 < 14; fi53++ {
                            var v54 string = (((("" + (("hello" + "") + (("" + "\n\t") + "a"))) + "") + (("foo bar" + "") + (("foo bar" + ("\n\t" + "foo bar")) + ("" + "\n\t")))) + "\n\t")
                            _ = fi53
                            _ = v54
                        }
                    case "a":
                        switch "" {
                        case "a":
                            switch int(v21) {
                            case 9223372036854775807:
                                var v55 int32 = v19
                                _ = v55
                            case 4294967295:
                                var v56 int64 = (^(v4))
                                _ = v56
                            case 4:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            }
                        default:
                            v57 := v1[("foo bar" + "\n\t")]
                            _ = v57
                        }
                    default:
                        _ = len(v20)
                    }
                    switch v52 {
                    case 42:
                        if false {
                            for fi58 := 0; fi58 < 16; fi58++ {
                                for fi59 := 0; fi59 < 3; fi59++ {
                                    v20 = append(v20, v52)
                                    _ = fi59
                                }
                                _ = fi58
                            }
                            var v60 int8 = 1
                            var v61 int32 = v19
                            _ = v60
                            _ = v61
                        }
                    case -923:
                        var v62 uint16 = (v2 / ((561) | 1))
                        _ = v62
                    case 127:
                        _ = len(v20)
                    }
                    _ = v52
                }
                _ = v44
            }
            for fi63 := 0; fi63 < 15; fi63++ {
                v64 := GenS5{F0: (v19 << ((18446744073709551615) & 63)), F1: v5}
                switch "" {
                case "foo bar":
                    switch int(v2) {
                    case 961:
                        fmt.Println(v3, v4)
                    case 675:
                        for fi65 := 0; fi65 < 13; fi65++ {
                            if (true && (!false)) {
                                v5 = 0
                                var v66 float64 = ((((((1e10 / 1.0) + (float64(0.0) + (0.0 / (1e10 - (1.5 * ((0.0 - ((3.14159 - 0.0) / (float64(3.14159) * -2.25))) / 1.5)))))) + (0.0 - 0.0)) - 0.1) * (1.5 * 0.0)) - 0.1)
                                _ = v66
                            } else {
                                v3 = v2
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v21, v4)
                                case "":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v67 := GenS6{F0: v18, F1: (-(v19)), F2: -49}
                                        _ = v67
                                    }
                                case "hello":
                                    v68 := GenS7{F0: int16(v19)}
                                    switch fi63 {
                                    case -89:
                                        var v69 uint64 = 568
                                        _ = v69
                                    default:
                                        v2 -= (v3 | (v3 >> ((1) & 63)))
                                    }
                                    _ = v68
                                default:
                                    if true {
                                        fmt.Println(v3, fi63)
                                        v70 := GenS8{F0: 127}
                                        _ = v70
                                    } else {
                                        fmt.Println(v4)
                                        v20 = append(v20, (^(fi63)))
                                        if false {
                                            var v71 string = ("hello" + "foo bar")
                                            var v72 uint32 = (v5 + v5)
                                            _ = v71
                                            _ = v72
                                        }
                                        var v73 int32 = int32(v4)
                                        fmt.Println(fi65, v3, fi65)
                                        v74 := GenS9{F0: v21}
                                        v75 := GenS2{F0: "a"}
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                }
                            }
                            if (!(1 > 1)) {
                                var v76 uint64 = 2
                                fmt.Println(v18, v5, v21)
                                var v77 string = "\n\t"
                                _ = v76
                                _ = v77
                            }
                            switch "" {
                            case "":
                                if false {
                                    v19 = v19
                                } else {
                                    v78 := GenS3{F0: -49}
                                    _ = v78
                                }
                            case "foo bar":
                                if (4294967295 <= 4294967295) {
                                    var v79 uint64 = 382
                                    v18 = 0
                                    _ = v79
                                } else {
                                    v18 = 957
                                }
                            case "a":
                                v80 := GenS6{F0: v18, F1: (-(v19)), F2: 127}
                                _ = v80
                            default:
                                _ = len(v1)
                            }
                            _ = fi65
                        }
                    default:
                        var v81 uint = (v18 * 1)
                        _ = v81
                    }
                case "":
                    fmt.Println(v19, v3, fi63)
                default:
                    v82 := GenS2{F0: "a"}
                    _ = v82
                }
                fmt.Println(fi63)
                _ = fi63
                _ = v64
            }
        }
        var v83 int8 = 127
        _ = v83
    } else {
        v84 := GenS3{F0: 4}
        _ = v84
    }
    if (-12 == -66) {
        v85 := GenS10{F0: v4}
        fmt.Println(v5, v21)
        v21 = (v21 - (v21 / ((-304) | 1)))
        _ = v85
    } else {
        v86 := GenS4{F0: 127}
        var v87 int64 = (v4 >> ((v18) & 63))
        v88 := GenS8{F0: 42}
        _ = v86
        _ = v87
        _ = v88
    }
    switch -591 {
    case 4294967295:
        if (v5 <= 1) {
            var v89 int = 254
            for fi90 := 0; fi90 < 12; fi90++ {
                for fi91 := 0; fi91 < 15; fi91++ {
                    v92 := GenS11{F0: fi91, F1: v5}
                    _ = fi91
                    _ = v92
                }
                fmt.Println(v18, fi90, fi90)
                _ = fi90
            }
            _ = v89
        } else {
            if ((v4 << (((v18 ^ 255)) & 63)) >= -1) {
                v3 -= v3
                var v93 bool = (42 <= uint16(v18))
                _ = v93
            }
            for fi94 := 0; fi94 < 19; fi94++ {
                var v95 float64 = 0.1
                fmt.Println(v5, v18)
                v2 = (v3 ^ v2)
                _ = fi94
                _ = v95
            }
            var v96 int32 = 42
            _ = v96
        }
    case 2147483647:
        v2 += v2
    case -424:
        var v97 int64 = (v4 ^ v4)
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
