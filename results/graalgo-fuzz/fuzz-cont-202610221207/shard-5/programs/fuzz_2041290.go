package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint8; F1 uint; F2 string }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int8; F1 uint16 }
type GenS11 struct { F0 uint16; F1 int8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int16; F1 bool; F2 int8 }

func main() {
    v1 := []string{"hello", "\n\t", "a", ("" + "a")}
    var v2 uint16 = 32767
    v3 := []string{"\n\t", ("foo bar" + ("foo bar" + "")), "\n\t"}
    v2 = v2
    v4 := GenS1{F0: v2}
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(fi6)
            var v7 uint64 = 297
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    switch -2147483648 {
    case -843:
        var v8 uint16 = 0
        _ = v8
    case 32767:
        var v9 float64 = float64(-2.25)
        _ = v9
    default:
        if true {
            var v10 bool = ("" <= (("a" + (((("\n\t" + ("" + "\n\t")) + "a") + (("a" + (("" + "hello") + "a")) + "")) + "hello")) + ""))
            switch 578 {
            case 0:
                for fi11 := 0; fi11 < 16; fi11++ {
                    if v10 {
                        v2 += (v2 - 32767)
                    }
                    if (((1 != 1) && (1 == 245)) && false) {
                        v2 -= (^(v2))
                        switch fi11 {
                        case 2147483647:
                            for fi12 := 0; fi12 < 2; fi12++ {
                                v1 = append(v1, ((("" + "") + "") + (("a" + "hello") + "")))
                                if len(v1) > 0 {
                                    v1[0] = "foo bar"
                                }
                                _ = fi12
                            }
                        }
                    } else {
                        switch "hello" {
                        case "hello":
                            v2 = (v2 - 32767)
                        case "foo bar":
                            if false {
                                for fi13 := 0; fi13 < 15; fi13++ {
                                    _ = len(v1)
                                    v10 = (18446744073709551615 == 65535)
                                    fmt.Println(v2)
                                    _ = fi13
                                }
                                v14 := GenS2{F0: 333}
                                v10 = v10
                                _ = v14
                            }
                        case "":
                            var v15 uint32 = 190
                            _ = v15
                        }
                        _ = len(v3)
                        for fi16 := 0; fi16 < 13; fi16++ {
                            var v17 int32 = 300
                            v17 = 42
                            if v10 {
                                if (423 >= 495) {
                                    var v18 int64 = 0
                                    _ = v18
                                } else {
                                    v19 := GenS3{F0: 127, F1: (!(!(v2 < v2)))}
                                    var v20 uint = 1
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint64 = 388
                                v17 -= v17
                                _ = v21
                            } else {
                                if v10 {
                                    v2 = (v2 << ((873) & 63))
                                }
                                _ = len(v3)
                                if (fi16 != fi11) {
                                    switch fi11 {
                                    case -2147483648:
                                        switch fi11 {
                                        case 4294967295:
                                            if v10 {
                                                var v22 uint = 1
                                                var v23 uint16 = (v2 >> ((v22) & 63))
                                                var v24 uint8 = 116
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                            } else {
                                                var v25 uint32 = 2147483647
                                                _ = v25
                                            }
                                        case 32767:
                                            switch 32768 {
                                            case 223:
                                                switch (fi11 - int(v17)) {
                                                case 1:
                                                    if false {
                                                        v17 += v17
                                                        v26 := GenS4{F0: 65535, F1: 42}
                                                        _ = v26
                                                    } else {
                                                        if (-2.25 != (1e10 * 1.5)) {
                                                            fmt.Println(v17, v10, v17)
                                                            v17 = (v17 ^ 651)
                                                        }
                                                    }
                                                case -52:
                                                    fmt.Println(fi11, fi11, v17)
                                                case -9223372036854775808:
                                                    if ((-(fi11)) <= (-(fi16))) {
                                                        v10 = ((360 != v2) || ((fi11 >> ((42) & 63)) <= (fi16 / ((128) | 1))))
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = "a"
                                                        }
                                                        _ = len(v3)
                                                    }
                                                }
                                            case 96:
                                                if v10 {
                                                    switch (fi16 / ((fi16) | 1)) {
                                                    case 876:
                                                        switch fi16 {
                                                        case 128:
                                                            var v27 int64 = 1
                                                            _ = v27
                                                        case 32767:
                                                            v10 = ((1.5 + 0.0) <= 3.14159)
                                                        }
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v10 = (2 == 697)
                                                    case -963:
                                                        switch "foo bar" {
                                                        case "":
                                                            v2 -= v2
                                                        }
                                                    }
                                                    var v28 float32 = 1e10
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    _ = v28
                                                }
                                            case 32767:
                                                v29 := GenS5{F0: 0}
                                                _ = v29
                                            default:
                                                if v10 {
                                                    var v30 uint16 = (v2 / ((60) | 1))
                                                    if v10 {
                                                        v31 := GenS6{F0: -9223372036854775808, F1: 470, F2: 42}
                                                        var v32 string = "\n\t"
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    _ = v30
                                                } else {
                                                    var v33 uint16 = v2
                                                    _ = v33
                                                }
                                            }
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v34 int16 = 743
                                                _ = v34
                                            case "hello":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v35 string = ("hello" + "foo bar")
                                                    _ = v35
                                                default:
                                                    v17 += (v17 & (v17 << ((446) & 63)))
                                                    v17 = v17
                                                    switch (-(fi11)) {
                                                    case 127:
                                                        v10 = ((1.5 / 0.1) < ((0.0 + float64(float32(0.0))) / (1.5 - 0.1)))
                                                    case 128:
                                                        v1 = append(v1, "")
                                                    }
                                                    var v36 string = "hello"
                                                    _ = v36
                                                }
                                            case "foo bar":
                                                if v10 {
                                                    fmt.Println(v2, fi11, fi11)
                                                    fmt.Println(v10, fi11)
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = "hello"
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v37 uint32 = 158
                                                    _ = v37
                                                }
                                            }
                                        }
                                    default:
                                        v17 = v17
                                        if (v10 && true) {
                                            var v38 int16 = 1
                                            var v39 float64 = 0.1
                                            v10 = false
                                            var v40 int = 23
                                            fmt.Println(v10, v38)
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        }
                                    }
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                    for fi41 := 0; fi41 < 17; fi41++ {
                        v10 = v10
                        var v42 bool = (!((42 <= -1) || (true || true)))
                        _ = fi41
                        _ = v42
                    }
                    _ = fi11
                }
            case -1:
                switch "foo bar" {
                case "foo bar":
                    for fi43 := 0; fi43 < 20; fi43++ {
                        switch (fi43 << ((1) & 63)) {
                        case -2147483648:
                            v44 := GenS5{F0: 607}
                            _ = v44
                        }
                        _ = fi43
                    }
                case "hello":
                    v10 = (42 > 0)
                case "\n\t":
                    v10 = (-1 == -1)
                default:
                    for fi45 := 0; fi45 < 6; fi45++ {
                        v10 = (128 != fi45)
                        v46 := GenS4{F0: 464, F1: 42}
                        _ = fi45
                        _ = v46
                    }
                }
            }
            v2 -= 0
            _ = v10
        }
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        v2 = v2
        v2 += (v2 - (v2 % ((v2) | 1)))
        fmt.Println(fi47)
        _ = fi47
    }
    v48 := GenS7{F0: 2, F1: 42, F2: ((("\n\t" + "hello") + "foo bar") + "foo bar")}
    v2 -= v2
    var v49 uint64 = 898
    if true {
        if (!false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi50 := 0; fi50 < 12; fi50++ {
                switch (^(fi50)) {
                case -1:
                    var v51 int16 = -678
                    _ = v51
                }
                v1 = append(v1, ("a" + ""))
                _ = fi50
            }
        } else {
            v52 := GenS4{F0: 42, F1: v2}
            v3 = append(v3, "")
            _ = v52
        }
        if true {
            v49 = (v49 % ((v49) | 1))
            for fi53 := 0; fi53 < 7; fi53++ {
                v2 -= (^(v2))
                for fi54 := 0; fi54 < 4; fi54++ {
                    v55 := GenS2{F0: (v49 << ((0) & 63))}
                    fmt.Println(v2, fi54, v49)
                    var v56 string = "\n\t"
                    _ = fi54
                    _ = v55
                    _ = v56
                }
                var v57 int16 = 1
                _ = fi53
                _ = v57
            }
            v58 := GenS2{F0: (^(v49))}
            _ = v58
        } else {
            var v59 float32 = -2.25
            _ = v59
        }
    } else {
        v60 := GenS1{F0: v2}
        var v61 uint8 = 127
        fmt.Println(v61)
        _ = v60
        _ = v61
    }
    for fi62 := 0; fi62 < 20; fi62++ {
        if true {
            var v63 string = ("hello" + "")
            _ = v63
        }
        _ = fi62
    }
    fmt.Println(v49, v2)
    if false {
        v64 := GenS8{F0: -2147483648, F1: 2}
        v65 := GenS4{F0: 1, F1: 255}
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = v64
        _ = v65
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 128 {
        case 32768:
            for fi66 := 0; fi66 < 13; fi66++ {
                if (2 != 4294967295) {
                    switch fi66 {
                    case -607:
                        v3 = append(v3, "foo bar")
                    }
                } else {
                    _ = len(v3)
                }
                _ = fi66
            }
        case 42:
            for fi67 := 0; fi67 < 2; fi67++ {
                if false {
                    var v68 bool = (((479 > 2) && false) || false)
                    v2 -= v2
                    _ = v68
                } else {
                    fmt.Println(v2)
                    v49 -= v49
                    for fi69 := 0; fi69 < 11; fi69++ {
                        switch (fi67 ^ (fi69 ^ -553)) {
                        case 4294967295:
                            if (1 <= 1000) {
                                if true {
                                    fmt.Println(fi69, v49, v49)
                                    if true {
                                        if false {
                                            switch ("a" + "") {
                                            case "hello":
                                                if (int32(v2) > -2147483648) {
                                                    v2 = (^(v2))
                                                    v70 := GenS9{F0: (113 > 830), F1: (fi67 % ((fi67) | 1))}
                                                    _ = v70
                                                }
                                            }
                                            var v71 bool = ((-1 != 0) && true)
                                            var v72 uint64 = 42
                                            _ = v71
                                            _ = v72
                                        } else {
                                            switch (^(fi67)) {
                                            case 9223372036854775807:
                                                v49 = 1
                                            case 32768:
                                                var v73 uint32 = 65535
                                                _ = v73
                                            default:
                                                if (9223372036854775807 < -405) {
                                                    fmt.Println(v2)
                                                    var v74 string = ""
                                                    switch (fi67 + (fi67 * (fi69 >> ((965) & 63)))) {
                                                    case 42:
                                                        v49 += 0
                                                    case 270:
                                                        fmt.Println(fi69)
                                                    case 4294967295:
                                                        switch 144 {
                                                        case 2147483647:
                                                            fmt.Println(fi69, v74, fi67)
                                                        case 32768:
                                                            var v75 float64 = float64(0.1)
                                                            _ = v75
                                                        case -808:
                                                            v76 := GenS10{F0: 1, F1: v2}
                                                            _ = v76
                                                        default:
                                                            var v77 string = ""
                                                            _ = v77
                                                        }
                                                    default:
                                                        var v78 int = (fi69 + fi69)
                                                        _ = v78
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v79 int8 = -1
                                                    v2 += v2
                                                    _ = v79
                                                }
                                            }
                                            v2 -= (v2 * 0)
                                            v49 -= (^(v49))
                                        }
                                    } else {
                                        if ((!((^(v2)) >= v2)) && (9223372036854775807 == -1)) {
                                            if len(v1) > 0 {
                                                v1[0] = (("a" + ("hello" + ("hello" + "\n\t"))) + "\n\t")
                                            }
                                            var v80 int32 = 42
                                            _ = v80
                                        } else {
                                            if true {
                                                var v81 uint16 = (^(v2))
                                                _ = v81
                                            }
                                        }
                                        if ((1 > 63) || (true && false)) {
                                            if (-1 <= 0) {
                                                v82 := GenS11{F0: v2, F1: 1}
                                                _ = v82
                                            }
                                        }
                                        v83 := GenS2{F0: v49}
                                        _ = v83
                                    }
                                }
                            } else {
                                v49 = (v49 & v49)
                                var v84 int = (fi69 * 42)
                                fmt.Println(v49, v2, fi67)
                                _ = v84
                            }
                        }
                        var v85 int = fi69
                        switch fi69 {
                        case -204:
                            v3 = append(v3, "\n\t")
                        case 0:
                            var v86 int16 = 32767
                            _ = v86
                        default:
                            v87 := GenS12{F0: 0}
                            _ = v87
                        }
                        if (!(!(!("a" == "foo bar")))) {
                            switch "hello" {
                            case "\n\t":
                                v88 := GenS13{F0: 32767, F1: false, F2: -1}
                                _ = v88
                            }
                            fmt.Println(v2, fi67, fi67)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v2 = uint16(v49)
                        }
                        _ = fi69
                        _ = v85
                    }
                }
                _ = fi67
            }
        default:
            var v89 uint64 = v49
            _ = v89
        }
    }
    v1 = append(v1, "foo bar")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
