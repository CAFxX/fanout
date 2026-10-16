package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 int; F2 uint }
type GenS5 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int64; F1 int; F2 string }
type GenS9 struct { F0 int8; F1 uint16; F2 int32 }
type GenS10 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 539
    v3 := []string{("hello" + (("a" + "\n\t") + "")), "hello"}
    v4 := make(map[int]uint32)
    v2 += (v2 << ((357) & 63))
    fmt.Println(v2)
    v2 = 42
    v3 = append(v3, "a")
    if (!((("foo bar" + "foo bar") == ("" + "foo bar")) || (v2 == v2))) {
        v5 := v4[32768]
        var v6 float64 = 0.0
        var v7 float64 = (v6 - v6)
        _ = v5
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 479, F1: 259, F2: 255}
        var v9 int8 = 8
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: (0.1 != ((-2.25 * float64((1e10 / 3.14159))) * float64(1e10)))}
    v4[128] = 42
    v11 := GenS1{F0: 2, F1: 42, F2: 86}
    switch "hello" {
    case "foo bar":
        for fi12 := 0; fi12 < 10; fi12++ {
            if true {
                for fi13 := 0; fi13 < 13; fi13++ {
                    switch fi13 {
                    case 128:
                        if true {
                            switch 32768 {
                            case -386:
                                v2 = v2
                            case 128:
                                if true {
                                    v2 = v2
                                } else {
                                    if true {
                                        v2 -= (v2 << ((18446744073709551615) & 63))
                                    } else {
                                        if ((1 == 809) && false) {
                                            v14 := GenS2{F0: (-1 < 1)}
                                            switch (-(fi13)) {
                                            case 42:
                                                if ((2147483647 >= 2147483647) && true) {
                                                    var v15 uint = 753
                                                    var v16 int = fi12
                                                    _ = v15
                                                    _ = v16
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = "foo bar"
                                                    }
                                                }
                                            case 791:
                                                v17 := GenS3{F0: (fi12 < 128), F1: (-(fi13)), F2: 260}
                                                _ = v17
                                            case 128:
                                                switch (fi12 & fi12) {
                                                case 42:
                                                    var v18 uint32 = 2
                                                    if ((false || ((true && true) || true)) && false) {
                                                        switch (^(fi12)) {
                                                        case 128:
                                                            var v19 int16 = 42
                                                            _ = v19
                                                        default:
                                                            var v20 uint64 = 174
                                                            _ = v20
                                                        }
                                                    }
                                                    if false {
                                                        v2 += v2
                                                        v21 := GenS2{F0: ((fi12 * (^(fi12))) <= fi13)}
                                                        _ = v21
                                                    } else {
                                                        v22 := GenS4{F0: 0, F1: fi13, F2: 955}
                                                        v23 := GenS3{F0: ((36 >= 204) || (!(fi12 <= 32768))), F1: (fi13 % ((fi12) | 1)), F2: 1}
                                                        if false {
                                                            v18 = v18
                                                            if true {
                                                                if true {
                                                                    v24 := GenS3{F0: false, F1: fi13, F2: uint16(v2)}
                                                                    _ = v24
                                                                } else {
                                                                    v2 = (^(v2))
                                                                    v18 -= (^(v18))
                                                                }
                                                                if ((!((!false) && false)) && (2 < 35)) {
                                                                    var v25 bool = false
                                                                    _ = v25
                                                                } else {
                                                                    switch fi12 {
                                                                    case 1:
                                                                        _ = len(v3)
                                                                    case 2147483647:
                                                                        v2 += (v2 * v2)
                                                                        v2 += 4294967295
                                                                    }
                                                                    v26 := GenS2{F0: false}
                                                                    if true {
                                                                        var v27 string = ("\n\t" + ("hello" + "\n\t"))
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        if true {
                                                                            v27 = "hello"
                                                                        } else {
                                                                            v27 = v27
                                                                            v18 = (v18 + (v18 >> ((371) & 63)))
                                                                        }
                                                                        _ = v27
                                                                    } else {
                                                                        fmt.Println(v2, v18)
                                                                        v18 = 2147483647
                                                                        var v28 float64 = 0.0
                                                                        delete(v4, fi12)
                                                                        _ = v28
                                                                    }
                                                                    _ = v26
                                                                }
                                                                var v29 uint8 = 255
                                                                _ = v29
                                                            } else {
                                                                var v30 int32 = -2147483648
                                                                _ = v30
                                                            }
                                                            v2 += v2
                                                            v4[fi12] = (v18 | v18)
                                                        }
                                                        _ = v22
                                                        _ = v23
                                                    }
                                                    _ = v18
                                                case 32768:
                                                    switch "a" {
                                                    case "hello":
                                                        v4[fi13] = 65535
                                                    }
                                                case -506:
                                                    var v31 int32 = -2147483648
                                                    _ = v31
                                                }
                                            default:
                                                fmt.Println(fi12, fi13, fi12)
                                            }
                                            if true {
                                                fmt.Println(fi12, fi13, fi12)
                                                v32 := GenS3{F0: true, F1: fi12, F2: 848}
                                                var v33 uint = 902
                                                v2 -= v2
                                                var v34 int = 2147483647
                                                v2 -= 724
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            } else {
                                                if (("hello" + ("\n\t" + (("hello" + ("\n\t" + (("foo bar" + "\n\t") + ("foo bar" + "foo bar")))) + "hello"))) == "hello") {
                                                    var v35 uint16 = 574
                                                    _ = v35
                                                }
                                                if (-543 > -1) {
                                                    v36 := v4[(fi13 >> ((707) & 63))]
                                                    v2 = 4294967295
                                                    v2 = v2
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    fmt.Println(fi12, v2)
                                                    _ = v36
                                                }
                                                var v37 uint32 = 4294967295
                                                _ = v37
                                            }
                                            _ = v14
                                        } else {
                                            if ("" <= "foo bar") {
                                                var v38 uint = 42
                                                fmt.Println(v38)
                                                var v39 int8 = 100
                                                _ = v38
                                                _ = v39
                                            }
                                            var v40 int64 = 0
                                            if true {
                                                var v41 string = "a"
                                                var v42 uint64 = v2
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v2 = v2
                                                fmt.Println(v2, v2, v40)
                                                var v43 int16 = -32768
                                                _ = v43
                                            }
                                            _ = v40
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "hello")
                                        }
                                    }
                                    v2 = (v2 | v2)
                                }
                            case 737:
                                v2 = (v2 ^ (v2 >> ((119) & 63)))
                            }
                            var v44 uint = 832
                            var v45 string = "hello"
                            _ = v44
                            _ = v45
                        } else {
                            v1[32768] = 2147483647
                            switch -963 {
                            case 4294967295:
                                var v46 int64 = -1
                                v46 -= 42
                                _ = v46
                            case 32768:
                                var v47 int = (fi12 ^ fi12)
                                _ = v47
                            default:
                                var v48 int32 = 324
                                _ = v48
                            }
                        }
                    case 441:
                        if (255 > 1) {
                            v1[fi12] = 4294967295
                            var v49 int = (fi12 << ((2) & 63))
                            _ = v49
                        } else {
                            if true {
                                if (!false) {
                                    if false {
                                        switch (fi13 ^ 171) {
                                        case -9223372036854775808:
                                            v50 := GenS5{F0: 42, F1: 160, F2: 2}
                                            _ = v50
                                        case 9223372036854775807:
                                            if (!(((float32(((((1.5 / 0.1) / (1e10 / 1e10)) * (1.5 * -2.25)) - (0.0 - -2.25))) / -2.25) * 1e10) >= (3.14159 * 1.5))) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if true {
                                                    v51 := GenS1{F0: 4294967295, F1: 540, F2: 1}
                                                    if true {
                                                        v52 := GenS2{F0: true}
                                                        var v53 int32 = 42
                                                        _ = v52
                                                        _ = v53
                                                    } else {
                                                        var v54 uint8 = 128
                                                        if false {
                                                            var v55 string = ("hello" + "")
                                                            _ = v55
                                                        }
                                                        _ = v54
                                                    }
                                                    var v56 uint = 0
                                                    _ = v51
                                                    _ = v56
                                                }
                                            }
                                        case -408:
                                            if false {
                                                v2 = (v2 << ((141) & 63))
                                            } else {
                                                var v57 float64 = (1.5 * ((1e10 - 1.5) * 3.14159))
                                                delete(v4, fi13)
                                                var v58 uint = 478
                                                _ = v57
                                                _ = v58
                                            }
                                        default:
                                            if false {
                                                switch fi13 {
                                                case 9223372036854775807:
                                                    v59 := GenS6{F0: 0, F1: 1}
                                                    _ = v59
                                                }
                                                v60 := GenS7{F0: (^(fi13)), F1: 255}
                                                _ = v60
                                            } else {
                                                v61 := GenS6{F0: 1, F1: -1}
                                                _ = v61
                                            }
                                        }
                                        if ((true || (-1 <= 0)) && true) {
                                            _ = len(v3)
                                            var v62 uint64 = (v2 >> ((42) & 63))
                                            switch "foo bar" {
                                            case "hello":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if (true || false) {
                                                    switch "\n\t" {
                                                    case "":
                                                        v62 -= 4294967295
                                                    case "\n\t":
                                                        switch (fi13 - (fi13 % ((817) | 1))) {
                                                        case -1:
                                                            v63 := GenS3{F0: (4294967295 != 395), F1: 127, F2: uint16(v62)}
                                                            _ = v63
                                                        case 1:
                                                            fmt.Println(v2, v2)
                                                        case 691:
                                                            if len(v3) > 0 {
                                                                v3[0] = "a"
                                                            }
                                                            var v64 bool = true
                                                            fmt.Println(v62, v62)
                                                            var v65 int = 0
                                                            _ = v64
                                                            _ = v65
                                                        default:
                                                            if true {
                                                                fmt.Println(v2, v62)
                                                                v62 = v2
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v2 = (v2 & v62)
                                                    }
                                                    var v66 uint = 42
                                                    var v67 uint64 = uint64(v66)
                                                    _ = v66
                                                    _ = v67
                                                }
                                                v62 += 1
                                                var v68 int32 = 1
                                                _ = v68
                                            case "":
                                                var v69 int64 = 964
                                                _ = v69
                                            }
                                            if false {
                                                if (true && (!(878 != 554))) {
                                                    switch fi12 {
                                                    case 128:
                                                        var v70 uint = 36
                                                        _ = v70
                                                    case -2147483648:
                                                        v71 := v1[(fi13 & fi12)]
                                                        _ = v71
                                                    case 127:
                                                        var v72 int = (fi13 & (^(fi13)))
                                                        _ = v72
                                                    }
                                                    if true {
                                                        fmt.Println(v2, fi12)
                                                        var v73 uint = 0
                                                        _ = v73
                                                    } else {
                                                        var v74 int16 = -685
                                                        if (40 > 127) {
                                                            switch "a" {
                                                            case "a":
                                                                var v75 float32 = (0.1 + 3.14159)
                                                                _ = v75
                                                            case "":
                                                                v76 := v1[(fi12 + (fi12 / ((32767) | 1)))]
                                                                _ = v76
                                                            }
                                                        } else {
                                                            v74 -= v74
                                                        }
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            fmt.Println(v62, v62)
                                            _ = v62
                                        }
                                        var v77 int = 320
                                        v78 := GenS2{F0: false}
                                        _ = v77
                                        _ = v78
                                    }
                                } else {
                                    fmt.Println(fi12)
                                    if true {
                                        v2 -= v2
                                        v2 += v2
                                        fmt.Println(v2, fi12, fi12)
                                        var v79 float32 = (0.0 + 0.0)
                                        _ = v79
                                    }
                                }
                            }
                            fmt.Println(v2)
                            fmt.Println(fi12, v2, fi13)
                        }
                    case -9223372036854775808:
                        if true {
                            var v80 uint8 = 2
                            switch "foo bar" {
                            case "":
                                var v81 int8 = 127
                                _ = v81
                            case "hello":
                                switch "a" {
                                case "\n\t":
                                    _ = len(v3)
                                case "foo bar":
                                    if (float64((-2.25 - -2.25)) == 0.0) {
                                        if true {
                                            var v82 uint8 = v80
                                            v83 := GenS6{F0: 65535, F1: -736}
                                            fmt.Println(fi13, fi13)
                                            v84 := v4[(fi12 * fi12)]
                                            v85 := GenS8{F0: -350, F1: fi13, F2: "\n\t"}
                                            _ = v82
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                        }
                                        var v86 uint64 = (v2 >> ((18446744073709551615) & 63))
                                        switch 128 {
                                        case -295:
                                            v86 = (v2 + v2)
                                        case 32768:
                                            v87 := GenS2{F0: (-1 == 42)}
                                            _ = v87
                                        case 166:
                                            v80 -= (v80 | (v80 * v80))
                                        default:
                                            v88 := GenS3{F0: (0 != v80), F1: -2147483648, F2: 0}
                                            _ = v88
                                        }
                                        _ = v86
                                    }
                                }
                            case "a":
                                var v89 int16 = 1000
                                _ = v89
                            default:
                                if (32767 == 0) {
                                    v2 -= 1
                                }
                            }
                            _ = v80
                        } else {
                            switch ((("a" + "foo bar") + "\n\t") + "") {
                            case "hello":
                                switch fi12 {
                                case 128:
                                    v2 -= v2
                                case 623:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 9223372036854775807:
                                    fmt.Println(fi13, fi13, fi12)
                                default:
                                    v90 := GenS1{F0: 2, F1: 198, F2: 2}
                                    _ = v90
                                }
                            case "a":
                                fmt.Println(fi12, fi12)
                            case "foo bar":
                                v91 := GenS8{F0: 0, F1: fi12, F2: ("" + "foo bar")}
                                _ = v91
                            }
                            fmt.Println(fi12)
                        }
                        v92 := GenS9{F0: -128, F1: 32767, F2: 0}
                        _ = v92
                    }
                    var v93 int8 = -128
                    v2 = 2
                    _ = fi13
                    _ = v93
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v94 := GenS7{F0: 42, F1: 1}
                _ = v94
            }
            v3 = append(v3, "hello")
            _ = fi12
        }
    case "hello":
        if false {
            var v95 uint32 = 65535
            var v96 string = "\n\t"
            _ = v95
            _ = v96
        }
    }
    v97 := GenS10{F0: -1}
    var v98 uint32 = 508
    if (128 != -1) {
        for fi99 := 0; fi99 < 18; fi99++ {
            fmt.Println(v2, fi99, fi99)
            _ = fi99
        }
    }
    var v100 int8 = -128
    v98 = v98
    v101 := GenS10{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v98:", v98)
    _ = v1
    _ = v10
    _ = v100
    _ = v101
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v97
    _ = v98
}
