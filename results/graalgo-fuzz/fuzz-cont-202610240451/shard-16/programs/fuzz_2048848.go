package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 string; F1 int16; F2 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8; F1 int; F2 int64 }
type GenS8 struct { F0 int; F1 int16; F2 uint16 }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 int; F1 uint64; F2 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{1, 65535}
    var v3 uint32 = 798
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint64 = 4294967295
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint16 = 32767
            _ = fi6
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v8 := v1["foo bar"]
        _ = v8
    default:
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
    }
    v3 = v3
    if (127 == -42) {
        _ = len(v1)
        fmt.Println(v3)
        if ((true && false) && (567 == 2147483647)) {
            v9 := GenS1{F0: -2147483648, F1: 1}
            switch -1 {
            case -1:
                v1[("\n\t" + "a")] = -945
            default:
                switch 183 {
                case 128:
                    v3 -= (v3 / ((v3) | 1))
                case 127:
                    v10 := v1["\n\t"]
                    _ = v10
                case 790:
                    for fi11 := 0; fi11 < 18; fi11++ {
                        switch (^(fi11)) {
                        case 372:
                            if (42 >= 1) {
                                fmt.Println(fi11)
                                v2 = append(v2, 1)
                            } else {
                                v1[("" + "")] = fi11
                                var v12 int16 = 32767
                                _ = v12
                            }
                        case 127:
                            v3 = v3
                        }
                        if false {
                            if (42 == 502) {
                                var v13 int32 = 536
                                v3 += (v3 - v3)
                                v14 := v1[("foo bar" + (("hello" + "hello") + ""))]
                                _ = v13
                                _ = v14
                            } else {
                                var v15 int = 1
                                v3 = (v3 << ((1) & 63))
                                _ = v15
                            }
                        } else {
                            var v16 int16 = 42
                            v1["hello"] = fi11
                            var v17 int16 = (v16 << ((2) & 63))
                            _ = v16
                            _ = v17
                        }
                        _ = fi11
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                }
            }
            _ = v9
        } else {
            if true {
                if true {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        v19 := GenS2{F0: "hello", F1: 42, F2: 0}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    var v20 int = 2147483647
                    var v21 bool = (584 < 4294967295)
                    for fi22 := 0; fi22 < 7; fi22++ {
                        if true {
                            var v23 uint16 = 93
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 2}
                            var v25 int64 = 1
                            _ = v24
                            _ = v25
                        }
                        _ = fi22
                    }
                    _ = v20
                    _ = v21
                }
                switch "\n\t" {
                case "":
                    var v26 int = 2147483647
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 2; fi27++ {
                    v28 := v1["\n\t"]
                    _ = fi27
                    _ = v28
                }
                switch 382 {
                case 2147483647:
                    v3 += 648
                }
            }
            v3 = v3
        }
    } else {
        _ = len(v1)
        v3 += 1
        if false {
            switch 0 {
            case 232:
                var v29 int64 = 1
                _ = v29
            }
        }
    }
    v3 = (v3 % (((^(v3))) | 1))
    v30 := GenS1{F0: -9223372036854775808, F1: 42}
    var v31 uint8 = 0
    v32 := GenS3{F0: v31}
    if ((716 != 1) && (-1 < -9223372036854775808)) {
        if ("\n\t" == "hello") {
            for fi33 := 0; fi33 < 9; fi33++ {
                v31 = (v31 - 0)
                for fi34 := 0; fi34 < 5; fi34++ {
                    if (!(1 >= 127)) {
                        v3 = (v3 << ((4294967295) & 63))
                        var v35 uint64 = 1
                        var v36 float32 = 0.0
                        _ = v35
                        _ = v36
                    } else {
                        var v37 int32 = 75
                        _ = v37
                    }
                    if true {
                        v31 += (v31 - (^(v31)))
                        v3 += 2147483648
                        v31 -= v31
                    } else {
                        v31 = v31
                    }
                    _ = fi34
                }
                _ = fi33
            }
            switch 127 {
            case 0:
                if len(v2) > 0 {
                    v2[0] = v3
                }
            case 1:
                v1["\n\t"] = -8
            default:
                for fi38 := 0; fi38 < 7; fi38++ {
                    var v39 int8 = -1
                    if (1.5 > (3.14159 + (1e10 / (-2.25 - 1e10)))) {
                        switch ("" + "hello") {
                        case "a":
                            var v40 int64 = 0
                            _ = v40
                        case "foo bar":
                            fmt.Println(fi38, v3)
                        default:
                            if ((false && true) && (273 <= 1)) {
                                delete(v1, "a")
                                if ((true && (!false)) && (!false)) {
                                    v2 = append(v2, v3)
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        fmt.Println(v39, fi38, v31)
                                        var v42 float32 = 0.0
                                        _ = fi41
                                        _ = v42
                                    }
                                    var v43 uint8 = (^(v31))
                                    _ = v43
                                } else {
                                    v1[""] = fi38
                                }
                                fmt.Println(v31, fi38, fi38)
                            }
                        }
                        for fi44 := 0; fi44 < 18; fi44++ {
                            if false {
                                switch fi44 {
                                case 982:
                                    v3 = v3
                                case -2147483648:
                                    if false {
                                        v45 := GenS1{F0: -2147483648, F1: 42}
                                        v39 += 0
                                        switch "\n\t" {
                                        case "":
                                            if true {
                                                v39 = -1
                                            } else {
                                                v46 := GenS3{F0: (v31 ^ (v31 * 1))}
                                                v47 := GenS4{F0: 2147483647, F1: 0}
                                                _ = v46
                                                _ = v47
                                            }
                                            v1[("a" + "a")] = (fi38 ^ fi44)
                                            v39 -= (v39 & v39)
                                        case "\n\t":
                                            v1["hello"] = (fi44 & (fi44 >> ((966) & 63)))
                                        case "a":
                                            if false {
                                                v48 := GenS5{F0: v3, F1: int8(fi44)}
                                                var v49 bool = (-161 > 388)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                var v50 uint32 = 2
                                                var v51 string = ("foo bar" + (((("" + "\n\t") + "") + "hello") + ("" + "hello")))
                                                _ = v50
                                                _ = v51
                                            }
                                        default:
                                            if false {
                                                var v52 int16 = -815
                                                _ = v52
                                            }
                                        }
                                        if true {
                                            if ((1e10 * ((1e10 / 3.14159) * (0.1 + 0.0))) != 3.14159) {
                                                switch "a" {
                                                case "hello":
                                                    if false {
                                                        v39 = v39
                                                        var v53 uint16 = 1
                                                        v54 := GenS1{F0: (fi44 & fi44), F1: (v53 | (v53 << ((uint(v3)) & 63)))}
                                                        _ = v53
                                                        _ = v54
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v3)
                                                case "foo bar":
                                                    var v55 float32 = ((((3.14159 - 1e10) / 1.5) + (0.0 + (3.14159 - 3.14159))) - 1e10)
                                                    _ = v55
                                                default:
                                                    v56 := GenS6{F0: "\n\t"}
                                                    _ = v56
                                                }
                                            }
                                            delete(v1, "a")
                                            if true {
                                                var v57 uint64 = 443
                                                _ = v57
                                            } else {
                                                v39 = -57
                                                if (!true) {
                                                    var v58 int64 = 0
                                                    _ = v58
                                                } else {
                                                    switch -1 {
                                                    case 2147483647:
                                                        _ = len(v1)
                                                    case -110:
                                                        if true {
                                                            if true {
                                                                switch (("foo bar" + "\n\t") + ("hello" + "foo bar")) {
                                                                case "\n\t":
                                                                    var v59 uint64 = 42
                                                                    _ = v59
                                                                case "hello":
                                                                    fmt.Println(fi44, fi38, fi38)
                                                                default:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v3 + v3)
                                                                    }
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            } else {
                                                                switch fi38 {
                                                                case -9223372036854775808:
                                                                    switch -2147483648 {
                                                                    case 127:
                                                                        var v60 float64 = 0.0
                                                                        switch fi38 {
                                                                        case -256:
                                                                            var v61 float32 = -2.25
                                                                            _ = v61
                                                                        case 37:
                                                                            var v62 string = ""
                                                                            _ = v62
                                                                        case -3:
                                                                            var v63 uint64 = 18446744073709551615
                                                                            _ = v63
                                                                        }
                                                                        _ = v60
                                                                    case 397:
                                                                        var v64 uint32 = v3
                                                                        _ = v64
                                                                    case 231:
                                                                        var v65 int16 = 42
                                                                        _ = v65
                                                                    default:
                                                                        v66 := GenS7{F0: (^(v31)), F1: 1, F2: -679}
                                                                        _ = v66
                                                                    }
                                                                case 42:
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (v3 + v3)
                                                                    }
                                                                default:
                                                                    v67 := GenS8{F0: fi44, F1: 1000, F2: 42}
                                                                    _ = v67
                                                                }
                                                                var v68 int64 = 9223372036854775807
                                                                var v69 int16 = -888
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                            v3 = (v3 + 2)
                                                        }
                                                    default:
                                                        if ((!(!((32767 > -311) || true))) || false) {
                                                            var v70 float64 = ((float64(0.0) + 1.5) * ((0.0 + 0.0) + ((1.5 - 1e10) + float64((1.5 + (float32(-2.25) / -2.25))))))
                                                            _ = v70
                                                        } else {
                                                            var v71 int32 = 2147483647
                                                            var v72 bool = (true || (!false))
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    }
                                                    v3 -= v3
                                                    v73 := GenS1{F0: 205, F1: 2}
                                                    _ = v73
                                                }
                                            }
                                        }
                                        v74 := GenS8{F0: 1, F1: 1, F2: 42}
                                        _ = v45
                                        _ = v74
                                    } else {
                                        switch (fi44 - fi38) {
                                        case 202:
                                            v31 = (v31 | 2)
                                            v39 -= (^(v39))
                                        case -721:
                                            var v75 uint16 = 2
                                            _ = v75
                                        case 32767:
                                            v39 = 62
                                        default:
                                            fmt.Println(fi44)
                                        }
                                        var v76 int32 = -1
                                        _ = v76
                                    }
                                case 338:
                                    v77 := GenS9{F0: 42, F1: v31}
                                    _ = v77
                                }
                                if false {
                                    v39 -= v39
                                    var v78 int16 = 1
                                    _ = v78
                                } else {
                                    v39 -= (v39 ^ (v39 / (((v39 * 100)) | 1)))
                                    var v79 uint64 = 42
                                    v31 -= (v31 + (v31 << ((42) & 63)))
                                    _ = v79
                                }
                                if false {
                                    v80 := GenS10{F0: fi44, F1: 39, F2: false}
                                    if false {
                                        v3 = v3
                                    }
                                    fmt.Println(v39)
                                    _ = v80
                                }
                            } else {
                                switch "foo bar" {
                                case "":
                                    var v81 int64 = -297
                                    v82 := GenS8{F0: int(v81), F1: 1, F2: 0}
                                    v83 := GenS2{F0: "a", F1: 42, F2: 9223372036854775807}
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                default:
                                    switch "a" {
                                    case "hello":
                                        var v84 uint32 = v3
                                        _ = v84
                                    default:
                                        v85 := v1["a"]
                                        _ = v85
                                    }
                                }
                            }
                            if false {
                                _ = len(v1)
                                v3 = v3
                                var v86 uint = 255
                                _ = v86
                            } else {
                                v31 += (v31 * v31)
                                if true {
                                    fmt.Println(fi44)
                                    switch "a" {
                                    case "a":
                                        v2 = append(v2, (v3 >> ((2) & 63)))
                                    }
                                    v87 := GenS2{F0: (("hello" + "") + ("" + "\n\t")), F1: 1000, F2: 0}
                                    _ = v87
                                } else {
                                    v88 := GenS11{F0: 1}
                                    if true {
                                        v3 = (v3 - 705)
                                        var v89 bool = (!false)
                                        v90 := GenS7{F0: 217, F1: fi38, F2: 9223372036854775807}
                                        v89 = (0 > -794)
                                        if v89 {
                                            var v91 float64 = 1e10
                                            _ = v91
                                        }
                                        _ = v89
                                        _ = v90
                                    } else {
                                        v3 = 2
                                        var v92 uint = 65535
                                        _ = v92
                                    }
                                    _ = v88
                                }
                            }
                            _ = fi44
                        }
                        v31 -= v31
                    }
                    _ = fi38
                    _ = v39
                }
            }
            switch 1 {
            case 128:
                if false {
                    fmt.Println(v3)
                }
            default:
                for fi93 := 0; fi93 < 5; fi93++ {
                    if true {
                        if false {
                            v1[(((("hello" + ("a" + "hello")) + ("a" + "a")) + "\n\t") + (("\n\t" + ("hello" + (("hello" + ((("foo bar" + ((("\n\t" + (((((("a" + "a") + "foo bar") + (("" + "\n\t") + ("a" + ""))) + "foo bar") + "") + "")) + (("" + "a") + ("hello" + "\n\t"))) + ("" + ("" + (("\n\t" + (("foo bar" + "\n\t") + "")) + ("foo bar" + "\n\t")))))) + "") + "foo bar")) + "hello"))) + "hello"))] = (fi93 | (^(fi93)))
                            v94 := GenS7{F0: v31, F1: (fi93 ^ (fi93 | fi93)), F2: 9223372036854775807}
                            fmt.Println(fi93)
                            _ = v94
                        }
                    }
                    v95 := GenS9{F0: 571, F1: (v31 + v31)}
                    _ = fi93
                    _ = v95
                }
            }
        }
        fmt.Println(v31, v3)
    }
    v96 := GenS12{F0: 100, F1: 4294967295, F2: 9223372036854775807}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v97 uint32 = (v3 >> ((42) & 63))
    var v98 int8 = 100
    v99 := GenS1{F0: 127, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v96
    _ = v97
    _ = v98
    _ = v99
}
