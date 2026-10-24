package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int; F1 uint64; F2 uint64 }
type GenS8 struct { F0 int32; F1 string }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 int16; F1 int }
type GenS11 struct { F0 string; F1 int16; F2 uint16 }

func main() {
    var v1 int16 = -970
    var v2 uint8 = 43
    var v3 int16 = (^(v1))
    v4 := make(map[string]uint32)
    if (!false) {
        for fi5 := 0; fi5 < 5; fi5++ {
            v2 -= v2
            v3 = 407
            v6 := GenS1{F0: (-(v3)), F1: 4294967295}
            _ = fi5
            _ = v6
        }
    }
    if false {
        v7 := []int{int(v2), 32768, 32768}
        var v8 uint32 = 0
        _ = v7
        _ = v8
    }
    fmt.Println(v3, v1, v2)
    if ((!true) && (true || (!false))) {
        v2 = v2
        v9 := GenS2{F0: 256, F1: 333}
        _ = v9
    } else {
        var v10 uint32 = 25
        var v11 int16 = (v1 * -605)
        _ = v10
        _ = v11
    }
    if false {
        if (0 > int(v3)) {
            var v12 int32 = 1
            _ = v12
        } else {
            var v13 int16 = v3
            _ = v13
        }
        if true {
            v14 := GenS3{F0: -1, F1: 18446744073709551615}
            v3 = (v3 | v3)
            if false {
                fmt.Println(v3, v2, v3)
                if false {
                    v15 := GenS1{F0: v1, F1: 4294967295}
                    v16 := GenS3{F0: 9223372036854775807, F1: 979}
                    v17 := []int{9223372036854775807, -9223372036854775808, 32767}
                    _ = v15
                    _ = v16
                    _ = v17
                }
            } else {
                switch "hello" {
                case "":
                    v18 := GenS4{F0: 4294967295, F1: 560}
                    _ = v18
                case "foo bar":
                    if true {
                        v3 = 826
                        var v19 uint32 = 1
                        _ = v19
                    }
                case "hello":
                    v20 := []int{42, 706, 4294967295, -2147483648}
                    _ = v20
                default:
                    var v21 uint16 = 0
                    _ = v21
                }
            }
            _ = v14
        } else {
            var v22 uint = 2
            _ = v22
        }
        v23 := GenS5{F0: 127, F1: 255}
        _ = v23
    } else {
        if ((!true) && true) {
            for fi24 := 0; fi24 < 9; fi24++ {
                switch "a" {
                case "\n\t":
                    switch "" {
                    case "a":
                        switch "\n\t" {
                        case "a":
                            var v25 uint64 = 314
                            _ = v25
                        case "":
                            v26 := GenS5{F0: (fi24 & fi24), F1: 979}
                            _ = v26
                        case "\n\t":
                            v27 := GenS4{F0: 1, F1: -830}
                            _ = v27
                        }
                    case "\n\t":
                        switch ((("" + "\n\t") + ("foo bar" + "hello")) + "a") {
                        case "foo bar":
                            switch "foo bar" {
                            case "a":
                                switch "\n\t" {
                                case "":
                                    if (4294967295 <= 4294967295) {
                                        if false {
                                            delete(v4, (("" + (("" + "a") + "foo bar")) + "foo bar"))
                                            switch ("hello" + "a") {
                                            case "":
                                                var v28 string = ""
                                                _ = v28
                                            case "a":
                                                var v29 uint8 = v2
                                                _ = v29
                                            case "foo bar":
                                                v1 += -32768
                                            }
                                        }
                                        v2 = 42
                                    }
                                case "\n\t":
                                    v30 := v4["hello"]
                                    _ = v30
                                default:
                                    if true {
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            if (1 != 4294967295) {
                                                v32 := GenS4{F0: 42, F1: 1}
                                                v33 := []int{fi24, (fi24 << ((4294967295) & 63)), (fi24 & -1)}
                                                var v34 float32 = float32(float64(0.0))
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            } else {
                                                var v35 string = ""
                                                var v36 int16 = v3
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = fi31
                                        }
                                    }
                                }
                            case "":
                                var v37 int64 = 1
                                _ = v37
                            case "hello":
                                v38 := GenS6{F0: 255, F1: "a"}
                                _ = v38
                            }
                        }
                    default:
                        for fi39 := 0; fi39 < 10; fi39++ {
                            v2 += (v2 - v2)
                            _ = fi39
                        }
                    }
                }
                var v40 string = "hello"
                _ = fi24
                _ = v40
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                v42 := GenS5{F0: 4294967295, F1: 757}
                _ = fi41
                _ = v42
            }
            v1 = (v1 << ((1) & 63))
        }
        delete(v4, "hello")
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        v2 = 0
        fmt.Println(fi43, fi43)
        _ = fi43
    }
    var v44 uint = 279
    if true {
        v45 := GenS4{F0: 2, F1: 42}
        v46 := []int{-663, -336, 9223372036854775807}
        switch -1 {
        case -149:
            if ((!false) && (true && false)) {
                v47 := GenS2{F0: uint(v2), F1: 2147483648}
                v2 = (^(v2))
                v1 += v1
                _ = v47
            } else {
                v1 = (v3 / ((-32768) | 1))
                v48 := GenS1{F0: v1, F1: (v44 / ((v44) | 1))}
                for fi49 := 0; fi49 < 2; fi49++ {
                    var v50 uint64 = 42
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        case 1:
            if false {
                for fi51 := 0; fi51 < 14; fi51++ {
                    fmt.Println(v3, v44, v3)
                    _ = fi51
                }
                switch "a" {
                case "a":
                    v1 = (v1 & (-(v1)))
                case "hello":
                    v2 += 127
                }
            }
        }
        _ = v45
        _ = v46
    }
    if false {
        var v52 float32 = ((3.14159 * 1.5) - float32(1e10))
        _ = len(v4)
        _ = v52
    } else {
        var v53 int = 9223372036854775807
        var v54 int16 = (v1 >> ((0) & 63))
        _ = v53
        _ = v54
    }
    v55 := GenS1{F0: (v3 + v3), F1: 1}
    var v56 int = 127
    v2 = v2
    if (!(false && false)) {
        v3 = (-(v3))
        switch ("" + "\n\t") {
        case "foo bar":
            if false {
                var v57 uint = (v44 * v44)
                v1 = v1
                _ = v57
            } else {
                var v58 bool = (!false)
                _ = len(v4)
                _ = v58
            }
        case "a":
            for fi59 := 0; fi59 < 19; fi59++ {
                switch (fi59 << ((v44) & 63)) {
                case -596:
                    var v60 int32 = 42
                    _ = v60
                case -657:
                    if true {
                        var v61 uint32 = 2147483647
                        _ = v61
                    }
                case 857:
                    switch v56 {
                    case -273:
                        for fi62 := 0; fi62 < 8; fi62++ {
                            v63 := []int{847, (-(v56)), (fi59 % ((fi62) | 1)), (fi62 >> ((2) & 63))}
                            v64 := GenS7{F0: (fi59 - fi59), F1: 0, F2: 42}
                            var v65 int64 = 732
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                    case 88:
                        for fi66 := 0; fi66 < 1; fi66++ {
                            fmt.Println(fi59)
                            _ = fi66
                        }
                    case -1:
                        var v67 uint32 = 2
                        _ = v67
                    default:
                        switch "\n\t" {
                        case "\n\t":
                            v2 -= v2
                        case "":
                            v68 := v4[""]
                            _ = v68
                        case "foo bar":
                            v4["\n\t"] = 856
                        }
                    }
                default:
                    switch fi59 {
                    case 42:
                        v69 := GenS8{F0: 42, F1: (("hello" + "a") + (("hello" + "hello") + "\n\t"))}
                        _ = v69
                    case 128:
                        if false {
                            v70 := []int{v56, fi59, (-(fi59)), -9223372036854775808}
                            var v71 float64 = (1.5 * float64(-2.25))
                            _ = v70
                            _ = v71
                        } else {
                            if ((!false) || (true || false)) {
                                var v72 string = "a"
                                var v73 uint = 147
                                _ = v72
                                _ = v73
                            } else {
                                for fi74 := 0; fi74 < 17; fi74++ {
                                    switch (fi74 * (fi74 * 128)) {
                                    case 42:
                                        if false {
                                            delete(v4, "")
                                            v75 := GenS9{F0: v2, F1: 0}
                                            _ = v75
                                        } else {
                                            if (false || true) {
                                                v44 -= 680
                                                v1 -= (-(v3))
                                            } else {
                                                if true {
                                                    var v76 float64 = 0.1
                                                    var v77 uint = v44
                                                    var v78 uint8 = 2
                                                    fmt.Println(v76, fi59, v77)
                                                    fmt.Println(v56)
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            }
                                            v4["a"] = 958
                                            v44 += (v44 % (((^(v44))) | 1))
                                        }
                                    case 263:
                                        fmt.Println(v3, v56, v2)
                                    default:
                                        v44 += (v44 ^ (v44 >> ((v44) & 63)))
                                    }
                                    fmt.Println(v2)
                                    var v79 int = 32767
                                    if false {
                                        fmt.Println(fi74)
                                        v80 := GenS1{F0: (v3 / ((42) | 1)), F1: 513}
                                        fmt.Println(v44, v1)
                                        _ = v80
                                    } else {
                                        v81 := []int{(^(v56)), (fi59 ^ (fi59 * 16))}
                                        v79 += (fi59 | (v79 >> ((v44) & 63)))
                                        _ = len(v4)
                                        if false {
                                            v82 := GenS4{F0: 771, F1: -9223372036854775808}
                                            switch ("a" + ("hello" + "foo bar")) {
                                            case "":
                                                fmt.Println(v79, v79)
                                            case "\n\t":
                                                var v83 int32 = 1
                                                _ = v83
                                            case "hello":
                                                if false {
                                                    v56 = (fi59 << (((^(v44))) & 63))
                                                    if (!((32767 != 0) || false)) {
                                                        if false {
                                                            v79 = fi74
                                                            var v84 uint64 = 956
                                                            fmt.Println(fi59, fi74)
                                                            v44 = 2
                                                            if (382 < 1) {
                                                                fmt.Println(v56, v3, v84)
                                                                v44 -= v44
                                                            }
                                                            v44 += v44
                                                            v85 := GenS1{F0: v3, F1: v44}
                                                            v86 := GenS3{F0: 9223372036854775807, F1: (v44 + (v44 + 1))}
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                    }
                                                }
                                            default:
                                                fmt.Println(v2, v79, v44)
                                            }
                                            _ = v82
                                        } else {
                                            v87 := GenS6{F0: 42, F1: "hello"}
                                            _ = v87
                                        }
                                        if len(v81) >= 2 {
                                            v81 = v81[1:len(v81)]
                                        }
                                        switch "\n\t" {
                                        case "":
                                            var v88 uint32 = 65535
                                            _ = v88
                                        }
                                        v1 = (v3 + (v3 + (v3 << ((222) & 63))))
                                        _ = v81
                                    }
                                    if true {
                                        var v89 uint = v44
                                        _ = v89
                                    }
                                    _ = fi74
                                    _ = v79
                                }
                                var v90 uint16 = 255
                                if (-1 >= -2147483648) {
                                    fmt.Println(v90, v56, fi59)
                                    v91 := GenS10{F0: v1, F1: 42}
                                    _ = v91
                                }
                                _ = v90
                            }
                        }
                    case 127:
                        var v92 int16 = -965
                        _ = v92
                    default:
                        if false {
                            var v93 uint32 = 42
                            v93 = v93
                            var v94 string = "a"
                            _ = v93
                            _ = v94
                        }
                    }
                }
                fmt.Println(v2, v44, v3)
                if false {
                    var v95 int16 = v1
                    if true {
                        v56 = fi59
                        if true {
                            v56 = (^(fi59))
                            if (true || ((2 != 8) || true)) {
                                v44 = (v44 | (v44 - v44))
                                var v96 float32 = 1.5
                                fmt.Println(v96, v3, v96)
                                _ = v96
                            }
                        }
                    }
                    _ = v95
                } else {
                    v56 = fi59
                    var v97 int64 = 843
                    fmt.Println(v1)
                    _ = v97
                }
                _ = fi59
            }
        case "\n\t":
            for fi98 := 0; fi98 < 2; fi98++ {
                fmt.Println(v2)
                v99 := GenS7{F0: 2147483647, F1: 51, F2: 18446744073709551615}
                v4[("\n\t" + ("a" + "foo bar"))] = 675
                _ = fi98
                _ = v99
            }
        }
    } else {
        if (!(1.5 <= (1.5 * -2.25))) {
            v44 = 4294967295
            v3 = 0
            if true {
                for fi100 := 0; fi100 < 6; fi100++ {
                    var v101 uint32 = 42
                    switch "" {
                    case "foo bar":
                        delete(v4, "\n\t")
                    case "\n\t":
                        v102 := []int{(fi100 ^ (-(fi100))), fi100, fi100, (-(fi100))}
                        _ = v102
                    case "a":
                        v101 += (^(v101))
                    }
                    _ = fi100
                    _ = v101
                }
                v103 := []int{}
                v104 := GenS6{F0: 2, F1: "hello"}
                _ = v103
                _ = v104
            }
        }
        var v105 float64 = 1e10
        if (v105 >= v105) {
            v106 := GenS11{F0: (("hello" + ("a" + "a")) + (("\n\t" + ("" + ("\n\t" + ("a" + (("hello" + "") + (("hello" + ((("foo bar" + "foo bar") + ("\n\t" + ("hello" + ("hello" + ("\n\t" + ("" + "foo bar")))))) + "\n\t")) + "a")))))) + "\n\t")), F1: 0, F2: 776}
            _ = v106
        } else {
            var v107 int = v56
            _ = v107
        }
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v55
    _ = v56
}
