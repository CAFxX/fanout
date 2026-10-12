package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 string; F2 bool }
type GenS7 struct { F0 int32; F1 string; F2 int8 }
type GenS8 struct { F0 bool; F1 int16; F2 uint }
type GenS9 struct { F0 int8; F1 int8 }
type GenS10 struct { F0 uint; F1 int; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string; F1 uint64 }
func genFn0() int32 {
    v4 := GenS1{F0: 42}
    if (0 > 1) {
        var v5 int16 = 32767
        _ = v5
    }
    var v6 uint32 = 74
    _ = v4
    _ = v6
    return -71
}
func genFn1(p0 int32, p1 uint32, p2 uint) int8 {
    v7 := []int{42, -9223372036854775808, 9223372036854775807, int(p1)}
    _ = len(v7)
    v8 := GenS1{F0: 746}
    for fi9 := 0; fi9 < 6; fi9++ {
        p1 -= (p1 * (p1 | (p1 - (p1 - (p1 >> (((p2 ^ (^(p2)))) & 63))))))
        p2 = 42
        _ = len(v7)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 uint = 756
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]bool)
    fmt.Println(v1, v1)
    if false {
        v10 := []int{}
        var v11 int = 4294967295
        for fi12 := 0; fi12 < 1; fi12++ {
            switch 4294967295 {
            case 638:
                switch ((("foo bar" + "") + "") + "") {
                case "":
                    fmt.Println(fi12, fi12, fi12)
                case "a":
                    for fi13 := 0; fi13 < 15; fi13++ {
                        var v14 string = "hello"
                        _ = fi13
                        _ = v14
                    }
                }
            case 780:
                if false {
                    if (-2.25 != -2.25) {
                        fmt.Println(v11)
                    }
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                } else {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v10 = append(v10, (-(fi15)))
                        switch "hello" {
                        case "hello":
                            v11 = v11
                            fmt.Println(v11)
                            if true {
                                v16 := GenS1{F0: 2}
                                v17 := GenS2{F0: genFn0()}
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(v1, v1, v11)
                            }
                            if false {
                                v2 += (v2 - -9223372036854775808)
                                _ = len(v10)
                            }
                            var v18 int64 = v2
                            _ = v18
                        case "\n\t":
                            var v19 int64 = v2
                            _ = v19
                        }
                        if (42 <= -1) {
                            var v20 uint32 = 2147483647
                            v21 := v3[("\n\t" + "a")]
                            _ = v20
                            _ = v21
                        } else {
                            switch "a" {
                            case "hello":
                                var v22 int32 = 42
                                _ = v22
                            case "":
                                delete(v3, "")
                            case "foo bar":
                                if true {
                                    v23 := GenS2{F0: int32(v1)}
                                    _ = v23
                                }
                            default:
                                var v24 uint64 = 1
                                _ = v24
                            }
                            if false {
                                _ = len(v3)
                                switch "\n\t" {
                                case "\n\t":
                                    switch fi12 {
                                    case -9223372036854775808:
                                        var v25 int8 = 100
                                        _ = v25
                                    case 127:
                                        switch ("\n\t" + (("" + "foo bar") + ((("\n\t" + "a") + "\n\t") + "hello"))) {
                                        case "hello":
                                            var v26 string = "a"
                                            _ = v26
                                        default:
                                            v27 := GenS3{F0: uint16(v1), F1: 2}
                                            _ = v27
                                        }
                                    default:
                                        v2 = v2
                                    }
                                    if (false && false) {
                                        _ = len(v3)
                                    } else {
                                        v1 = v1
                                        if true {
                                            v2 -= v2
                                            switch fi12 {
                                            case 1:
                                                var v28 int32 = genFn0()
                                                _ = v28
                                            case 357:
                                                v10 = append(v10, v11)
                                            }
                                            switch fi15 {
                                            case 32768:
                                                var v29 string = "\n\t"
                                                _ = v29
                                            case -9223372036854775808:
                                                if ((0.0 - ((0.0 - 1e10) * 1e10)) > (1e10 - -2.25)) {
                                                    v30 := GenS4{F0: int64(v1)}
                                                    v11 += (fi12 - 42)
                                                    v2 -= v2
                                                    _ = v30
                                                } else {
                                                    if (42 >= 127) {
                                                        _ = len(v3)
                                                        switch fi15 {
                                                        case -1:
                                                            if ("\n\t" <= "foo bar") {
                                                                if len(v10) > 0 {
                                                                    v10[0] = (v11 % (((v11 << ((v1) & 63))) | 1))
                                                                }
                                                                v31 := GenS4{F0: -1}
                                                                _ = v31
                                                            } else {
                                                                v32 := GenS2{F0: 1}
                                                                _ = v32
                                                            }
                                                        case 32768:
                                                            fmt.Println(v1, v1)
                                                        case 394:
                                                            fmt.Println(v11, fi12)
                                                        default:
                                                            if (!(((0 > 505) && (2 == (v1 >> ((v1) & 63)))) && (42 < 633))) {
                                                                v33 := GenS4{F0: (v2 ^ (v2 >> ((924) & 63)))}
                                                                var v34 int = (fi15 / (((^(v11))) | 1))
                                                                _ = v33
                                                                _ = v34
                                                            } else {
                                                                switch "a" {
                                                                case "a":
                                                                    v35 := v3[("hello" + "")]
                                                                    _ = v35
                                                                case "\n\t":
                                                                    if false {
                                                                        fmt.Println(fi15, v1)
                                                                    }
                                                                default:
                                                                    v36 := GenS5{F0: ("\n\t" + "\n\t")}
                                                                    _ = v36
                                                                }
                                                                v11 = fi12
                                                                var v37 float32 = (0.1 - 1.5)
                                                                _ = v37
                                                            }
                                                        }
                                                    }
                                                    v1 = (v1 + v1)
                                                }
                                            default:
                                                v1 = (v1 / ((v1) | 1))
                                            }
                                        }
                                    }
                                    v11 = v11
                                    v10 = append(v10, (v11 << ((v1) & 63)))
                                }
                                v11 = 32768
                            }
                        }
                        _ = fi15
                    }
                    if ((0 <= 676) || false) {
                        v10 = append(v10, (fi12 << ((v1) & 63)))
                        var v38 uint8 = 14
                        _ = v38
                    } else {
                        v39 := v3["\n\t"]
                        v2 = (v2 | (v2 | v2))
                        v40 := GenS6{F0: (v39 && (90 != 1)), F1: "hello", F2: v39}
                        _ = v39
                        _ = v40
                    }
                    _ = len(v3)
                }
            default:
                v2 += v2
            }
            var v41 int = (fi12 - (v11 % ((int(v1)) | 1)))
            if (-381 > v2) {
                var v42 uint8 = 0
                fmt.Println(v11, v1, v41)
                _ = v42
            }
            _ = fi12
            _ = v41
        }
        _ = v10
        _ = v11
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        var v44 uint16 = 748
        switch ((("foo bar" + ("\n\t" + (("hello" + "\n\t") + (("foo bar" + "hello") + "foo bar")))) + ("hello" + "foo bar")) + (("" + (("hello" + "\n\t") + "foo bar")) + "a")) {
        case "a":
            if true {
                v45 := GenS7{F0: 2147483647, F1: "foo bar", F2: 100}
                for fi46 := 0; fi46 < 15; fi46++ {
                    v47 := GenS1{F0: 4294967295}
                    var v48 int8 = 0
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = v45
            }
        case "foo bar":
            if (fi43 >= (fi43 / (((fi43 << (((^(v1))) & 63))) | 1))) {
                v49 := GenS4{F0: (^(v2))}
                _ = v49
            } else {
                switch ("\n\t" + "hello") {
                case "\n\t":
                    v50 := GenS8{F0: true, F1: 0, F2: (v1 >> ((439) & 63))}
                    _ = v50
                case "":
                    if (!(!false)) {
                        if false {
                            var v51 float32 = 3.14159
                            switch (("\n\t" + "") + "foo bar") {
                            case "hello":
                                _ = len(v3)
                            case "":
                                if true {
                                    var v52 uint = v1
                                    for fi53 := 0; fi53 < 20; fi53++ {
                                        v51 -= float32(0.1)
                                        _ = fi53
                                    }
                                    switch "hello" {
                                    case "a":
                                        v54 := GenS7{F0: -2147483648, F1: "foo bar", F2: 42}
                                        _ = v54
                                    default:
                                        var v55 int16 = 1
                                        _ = v55
                                    }
                                    _ = v52
                                }
                            }
                            _ = v51
                        }
                    } else {
                        switch "foo bar" {
                        case "a":
                            if ((857 >= 0) && ((v2 / ((-1) | 1)) != (v2 | (v2 - 1)))) {
                                var v56 float64 = ((((1e10 / 1.5) - (3.14159 - (1e10 + (3.14159 + float64((3.14159 * (-2.25 - ((1.5 - (0.0 + 1e10)) / 1.0)))))))) / 1.0) * float64((float32(3.14159) - float32((1e10 * (0.1 - 0.1))))))
                                v57 := GenS9{F0: 42, F1: -63}
                                var v58 uint32 = 2
                                _ = v56
                                _ = v57
                                _ = v58
                            } else {
                                fmt.Println(v1)
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    v60 := GenS10{F0: 2, F1: fi59, F2: 2}
                                    if false {
                                        if false {
                                            switch fi59 {
                                            case 165:
                                                v61 := []int{fi59}
                                                _ = v61
                                            case 880:
                                                v62 := GenS8{F0: false, F1: 32767, F2: 62}
                                                _ = v62
                                            case 4294967295:
                                                v3[""] = false
                                            default:
                                                var v63 int64 = v2
                                                _ = v63
                                            }
                                            var v64 int = (fi59 + (-(fi43)))
                                            _ = v64
                                        } else {
                                            v65 := GenS4{F0: (^(v2))}
                                            v44 = 42
                                            _ = v65
                                        }
                                        v66 := []int{}
                                        _ = v66
                                    }
                                    _ = fi59
                                    _ = v60
                                }
                                if true {
                                    for fi67 := 0; fi67 < 1; fi67++ {
                                        if (0.0 != 1.5) {
                                            switch (fi43 | fi67) {
                                            case 1:
                                                if (4294967295 != 43) {
                                                    v68 := GenS4{F0: v2}
                                                    v44 = v44
                                                    _ = v68
                                                } else {
                                                    if (false || false) {
                                                        var v69 int64 = v2
                                                        fmt.Println(v44)
                                                        _ = v69
                                                    } else {
                                                        var v70 float32 = 1e10
                                                        var v71 int32 = -722
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                            }
                                        } else {
                                            v72 := v3["a"]
                                            v73 := GenS1{F0: 478}
                                            _ = v72
                                            _ = v73
                                        }
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(fi67, fi67)
                                        }
                                        v2 = v2
                                        var v74 uint = v1
                                        _ = fi67
                                        _ = v74
                                    }
                                }
                            }
                        default:
                            var v75 float32 = (0.1 + (0.0 * 1e10))
                            _ = v75
                        }
                        v76 := GenS3{F0: 65535, F1: 0}
                        v2 = (-(v2))
                        _ = v76
                    }
                case "hello":
                    if false {
                        var v77 int32 = 42
                        _ = v77
                    } else {
                        fmt.Println(fi43)
                    }
                }
            }
        default:
            var v78 int32 = 42
            _ = v78
        }
        _ = fi43
        _ = v44
    }
    fmt.Println(v2)
    var v79 uint16 = 65535
    var v80 string = "hello"
    if (false || (1 != -32768)) {
        for fi81 := 0; fi81 < 7; fi81++ {
            v82 := []int{(fi81 >> ((255) & 63)), fi81, fi81, fi81}
            var v83 uint32 = 4294967295
            v2 += 0
            _ = fi81
            _ = v82
            _ = v83
        }
        v84 := []int{0, -9223372036854775808}
        _ = v84
    } else {
        for fi85 := 0; fi85 < 10; fi85++ {
            for fi86 := 0; fi86 < 7; fi86++ {
                v87 := []int{(^(fi86))}
                fmt.Println(fi86)
                var v88 bool = true
                _ = fi86
                _ = v87
                _ = v88
            }
            v89 := []int{(fi85 ^ fi85), (fi85 + (^(fi85)))}
            _ = fi85
            _ = v89
        }
        v79 += (v79 >> (((v1 + v1)) & 63))
    }
    v1 = v1
    for fi90 := 0; fi90 < 8; fi90++ {
        var v91 uint8 = 128
        _ = fi90
        _ = v91
    }
    if true {
        v2 = (v2 << ((v1) & 63))
        var v92 int64 = v2
        for fi93 := 0; fi93 < 7; fi93++ {
            fmt.Println(v80, v80)
            fmt.Println(fi93, v79, v1)
            var v94 int = 0
            _ = fi93
            _ = v94
        }
        _ = v92
    } else {
        var v95 float64 = ((((((3.14159 / ((1e10 / 1e10) + 1.5)) - float64(1e10)) + (1.5 * float64(0.0))) - ((3.14159 - 0.0) + 3.14159)) * (0.1 * float64(((1e10 * float32(((1.5 + (-2.25 + float64(1e10))) / ((((float64(0.1) - float64(1.5)) - (-2.25 * 1e10)) / 1.5) / -2.25)))) + 3.14159)))) / ((1.5 + 1e10) - 1e10))
        v96 := GenS1{F0: 4294967295}
        _ = v95
        _ = v96
    }
    v80 = ((v80 + "hello") + ((v80 + v80) + v80))
    v80 = ("\n\t" + v80)
    v1 += v1
    if false {
        var v97 uint = (^(v1))
        v97 = 4294967295
        for fi98 := 0; fi98 < 11; fi98++ {
            v2 = v2
            _ = fi98
        }
        _ = v97
    } else {
        v3[("a" + v80)] = true
        _ = len(v3)
    }
    fmt.Println(v1, v2, v80)
    for fi99 := 0; fi99 < 14; fi99++ {
        fmt.Println(v2)
        if false {
            v100 := GenS11{F0: 68}
            if false {
                v80 = (("hello" + v80) + v80)
            }
            fmt.Println(v1, v1)
            _ = v100
        } else {
            for fi101 := 0; fi101 < 13; fi101++ {
                v79 = (v79 | (v79 * 843))
                v2 -= v2
                v3[v80] = (fi99 >= (fi101 * (fi101 + fi99)))
                var v102 int16 = 1000
                if ("foo bar" > (v80 + v80)) {
                    switch fi101 {
                    case 569:
                        switch fi99 {
                        case -2147483648:
                            switch 128 {
                            case 9223372036854775807:
                                var v103 string = "\n\t"
                                _ = v103
                            case 42:
                                switch (v80 + v80) {
                                case "foo bar":
                                    v104 := GenS12{F0: "hello", F1: 18446744073709551615}
                                    _ = v104
                                case "\n\t":
                                    v105 := GenS3{F0: (v79 + 2), F1: 102}
                                    _ = v105
                                }
                            case 183:
                                var v106 uint32 = 65535
                                _ = v106
                            }
                        }
                    }
                    v107 := []int{}
                    _ = v107
                } else {
                    v79 = (v79 / ((1) | 1))
                }
                _ = fi101
                _ = v102
            }
            fmt.Println(v2, fi99)
        }
        _ = fi99
    }
    fmt.Println(v80, v79, v2)
    var v108 uint64 = 993
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v108
    _ = v2
    _ = v3
    _ = v79
    _ = v80
}
