package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 uint; F1 uint; F2 int }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint16; F1 uint32; F2 int16 }
func genFn0(p0 uint32, p1 uint8) int {
    var v6 string = "hello"
    v7 := GenS1{F0: v6}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1() int32 {
    var v8 uint16 = 33
    var v9 int8 = 0
    v10 := []int{genFn0(65535, 127), 85, 1}
    v11 := GenS1{F0: "\n\t"}
    v12 := GenS2{F0: 173}
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return int32(v9)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -868
    var v3 uint16 = 68
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    v3 -= v3
    v13 := []int{(v2 << ((2) & 63)), (v2 * genFn0(4294967295, 91)), v2}
    if false {
        v14 := v4[1]
        v14 += (v14 * 2147483647)
        var v15 int64 = 0
        _ = v14
        _ = v15
    }
    switch (v2 << ((733) & 63)) {
    case 2147483647:
        if true {
            fmt.Println(v2)
        } else {
            if (genFn1() < 2147483647) {
                switch v2 {
                case 32768:
                    for fi16 := 0; fi16 < 3; fi16++ {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            v3 -= (v3 - v3)
                            switch ("" + "\n\t") {
                            case "hello":
                                v18 := GenS1{F0: "a"}
                                _ = v18
                            case "a":
                                if len(v13) > 0 {
                                    v13[0] = fi16
                                }
                            case "foo bar":
                                v19 := GenS3{F0: 42, F1: fi16, F2: v3}
                                _ = v19
                            default:
                                var v20 int32 = -2147483648
                                var v21 uint8 = 42
                                if true {
                                    var v22 float32 = 1.5
                                    _ = v22
                                }
                                _ = v20
                                _ = v21
                            }
                            var v23 uint8 = 121
                            _ = fi17
                            _ = v23
                        }
                        _ = fi16
                    }
                default:
                    v24 := GenS4{F0: 2147483647, F1: -40}
                    _ = v24
                }
                v13 = append(v13, v2)
            } else {
                if false {
                    switch v2 {
                    case -2147483648:
                        for fi25 := 0; fi25 < 12; fi25++ {
                            if ((v2 + -85) != -9223372036854775808) {
                                var v26 int = 128
                                if true {
                                    var v27 int = 32767
                                    v3 = v3
                                    for fi28 := 0; fi28 < 3; fi28++ {
                                        fmt.Println(v26)
                                        v2 = v26
                                        _ = fi28
                                    }
                                    _ = v27
                                }
                                _ = v26
                            }
                            for fi29 := 0; fi29 < 6; fi29++ {
                                v2 = (fi29 - (-(fi29)))
                                var v30 float64 = 1e10
                                if false {
                                    v31 := GenS5{F0: 2, F1: 738}
                                    v3 = 1
                                    var v32 uint16 = v3
                                    _ = v31
                                    _ = v32
                                }
                                _ = fi29
                                _ = v30
                            }
                            _ = fi25
                        }
                    case 4294967295:
                        for fi33 := 0; fi33 < 15; fi33++ {
                            v3 -= 65535
                            if true {
                                if (!false) {
                                    _ = len(v4)
                                    var v34 uint16 = v3
                                    _ = v34
                                } else {
                                    var v35 int = genFn0(1, 42)
                                    _ = v35
                                }
                            }
                            _ = fi33
                        }
                    case -9223372036854775808:
                        for fi36 := 0; fi36 < 7; fi36++ {
                            var v37 string = "foo bar"
                            for fi38 := 0; fi38 < 5; fi38++ {
                                switch (((v37 + (((v37 + v37) + v37) + v37)) + (v37 + v37)) + v37) {
                                case "foo bar":
                                    if true {
                                        v13 = append(v13, genFn0(1, 135))
                                        var v39 int16 = -1
                                        _ = v39
                                    } else {
                                        var v40 uint8 = 0
                                        _ = v40
                                    }
                                case "\n\t":
                                    v41 := GenS6{F0: -1, F1: (v3 << ((188) & 63))}
                                    _ = v41
                                }
                                v37 = v37
                                _ = fi38
                            }
                            _ = fi36
                            _ = v37
                        }
                    }
                }
            }
            switch (v2 ^ 4294967295) {
            case 0:
                var v42 bool = (127 == 225)
                _ = v42
            default:
                var v43 int32 = 383
                _ = v43
            }
        }
    case 32767:
        fmt.Println(v2, v3)
    case -114:
        var v44 uint8 = 244
        _ = v44
    }
    v2 += 0
    switch "foo bar" {
    case "a":
        v13 = append(v13, (v2 % (((^(v2))) | 1)))
    case "\n\t":
        if true {
            v45 := v5["\n\t"]
            v2 -= (v2 - -2147483648)
            fmt.Println(v2, v2, v2)
            _ = v45
        } else {
            v2 -= v2
            var v46 int32 = 1
            _ = v46
        }
    case "hello":
        v2 = genFn0(826, 0)
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        if (65535 != 2147483647) {
            var v48 uint8 = 157
            _ = v48
        } else {
            v2 += 0
            for fi49 := 0; fi49 < 9; fi49++ {
                switch (("foo bar" + "a") + ("\n\t" + "\n\t")) {
                case "foo bar":
                    v2 = (v2 - (fi47 << ((2) & 63)))
                case "a":
                    var v50 float64 = 1e10
                    _ = v50
                case "hello":
                    fmt.Println(v2)
                    var v51 float32 = 1.5
                    v2 -= (fi47 * (v2 >> ((42) & 63)))
                    _ = v51
                default:
                    if ((0 > 582) && false) {
                        v2 -= v2
                    }
                }
                var v52 uint8 = 2
                _ = fi49
                _ = v52
            }
        }
        _ = fi47
    }
    var v53 uint8 = 2
    v1[(-(v2))] = false
    var v54 int16 = 1
    v54 += v54
    if false {
        if true {
            var v55 int64 = 1
            v53 = v53
            v55 += (v55 ^ -9223372036854775808)
            _ = v55
        } else {
            switch (v2 | genFn0(0, v53)) {
            case 1:
                var v56 uint8 = 0
                _ = v56
            default:
                switch (v2 % ((v2) | 1)) {
                case -2147483648:
                    v57 := GenS2{F0: v53}
                    _ = v57
                case -91:
                    switch "a" {
                    case "a":
                        if false {
                            v54 = (v54 | v54)
                            var v58 int8 = 0
                            if (genFn0(42, (v53 / ((uint8(v3)) | 1))) > (v2 * v2)) {
                                if (v58 < v58) {
                                    switch (("" + "") + "") {
                                    case "":
                                        if false {
                                            switch (v2 * genFn0(318, uint8(v3))) {
                                            case 1:
                                                var v59 uint32 = 930
                                                _ = v59
                                            case 4294967295:
                                                if true {
                                                    delete(v1, 800)
                                                    var v60 uint32 = 0
                                                    _ = v60
                                                } else {
                                                    switch (^(v2)) {
                                                    case -1:
                                                        var v61 uint32 = 2147483647
                                                        _ = v61
                                                    default:
                                                        var v62 int64 = 1
                                                        _ = v62
                                                    }
                                                    if ((v58 / ((v58) | 1)) < (^(v58))) {
                                                        v58 += -1
                                                        if (float64(1.5) != 0.1) {
                                                            if len(v13) > 0 {
                                                                v13[0] = v2
                                                            }
                                                            fmt.Println(v58, v54, v3)
                                                            for fi63 := 0; fi63 < 2; fi63++ {
                                                                if false {
                                                                    v53 = 234
                                                                    v58 = (v58 << ((0) & 63))
                                                                    for fi64 := 0; fi64 < 20; fi64++ {
                                                                        v65 := GenS7{F0: 9223372036854775807}
                                                                        v66 := GenS5{F0: 255, F1: 4294967295}
                                                                        v67 := v1[genFn0(2147483647, (v53 & (^(v53))))]
                                                                        _ = fi64
                                                                        _ = v65
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                } else {
                                                                    switch ("a" + "foo bar") {
                                                                    case "foo bar":
                                                                        v68 := GenS4{F0: -99, F1: (v58 ^ v58)}
                                                                        _ = v68
                                                                    case "":
                                                                        for fi69 := 0; fi69 < 15; fi69++ {
                                                                            if false {
                                                                                var v70 int = (v2 - (fi63 | v2))
                                                                                _ = v70
                                                                            }
                                                                            _ = fi69
                                                                        }
                                                                    default:
                                                                        v3 += (^(v3))
                                                                    }
                                                                    if false {
                                                                        var v71 int = 42
                                                                        var v72 int16 = (v54 >> ((42) & 63))
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v72 = v72
                                                                        }
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                }
                                                                fmt.Println(v3, v2)
                                                                switch fi63 {
                                                                case 32767:
                                                                    var v73 int = v2
                                                                    _ = v73
                                                                case 766:
                                                                    if len(v13) > 0 {
                                                                        v13[0] = fi63
                                                                    }
                                                                default:
                                                                    fmt.Println(fi63)
                                                                }
                                                                _ = fi63
                                                            }
                                                        }
                                                        switch (v2 << ((255) & 63)) {
                                                        case 9223372036854775807:
                                                            var v74 uint64 = 1
                                                            _ = v74
                                                        case -118:
                                                            v2 = v2
                                                        }
                                                    }
                                                }
                                            case 145:
                                                v75 := GenS8{F0: false, F1: 42}
                                                _ = v75
                                            }
                                            v76 := v5[("foo bar" + "hello")]
                                            v54 -= v54
                                            _ = v76
                                        } else {
                                            fmt.Println(v53)
                                            fmt.Println(v54)
                                            if (true && (!(v54 >= v54))) {
                                                v53 = v53
                                                if (true || (127 <= 163)) {
                                                    switch genFn0(751, v53) {
                                                    case -2147483648:
                                                        for fi77 := 0; fi77 < 5; fi77++ {
                                                            if false {
                                                                v3 += v3
                                                                fmt.Println(v54, v2, v2)
                                                            } else {
                                                                _ = len(v5)
                                                                if false {
                                                                    delete(v4, int(v3))
                                                                } else {
                                                                    v78 := GenS9{F0: 836, F1: 65535, F2: fi77}
                                                                    _ = v78
                                                                }
                                                            }
                                                            v79 := GenS10{F0: 65535}
                                                            if false {
                                                                if (!false) {
                                                                    var v80 int16 = (-(v54))
                                                                    v2 = fi77
                                                                    _ = v80
                                                                }
                                                                var v81 float64 = 1e10
                                                                _ = v81
                                                            }
                                                            _ = fi77
                                                            _ = v79
                                                        }
                                                    case 127:
                                                        fmt.Println(v2, v58)
                                                    case -124:
                                                        v82 := GenS8{F0: true, F1: 2}
                                                        _ = v82
                                                    }
                                                } else {
                                                    for fi83 := 0; fi83 < 16; fi83++ {
                                                        var v84 float32 = ((3.14159 + 3.14159) - 3.14159)
                                                        _ = fi83
                                                        _ = v84
                                                    }
                                                    v85 := GenS11{F0: 255, F1: 205, F2: -1}
                                                    _ = v85
                                                }
                                            } else {
                                                v86 := GenS3{F0: 0, F1: int(v58), F2: v3}
                                                switch 32767 {
                                                case -1:
                                                    var v87 uint8 = 2
                                                    _ = v87
                                                }
                                                _ = v86
                                            }
                                        }
                                    }
                                }
                                v53 -= v53
                            }
                            _ = v58
                        }
                    case "\n\t":
                        switch "a" {
                        case "hello":
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                var v88 uint64 = 42
                                _ = v88
                            }
                        }
                    case "hello":
                        if (127 > (v53 | v53)) {
                            var v89 string = "a"
                            v90 := v5[(v89 + (v89 + v89))]
                            switch "" {
                            case "\n\t":
                                if len(v13) > 0 {
                                    v13[0] = 1
                                }
                            case "hello":
                                v91 := GenS7{F0: 1}
                                _ = v91
                            default:
                                v92 := GenS10{F0: 0}
                                _ = v92
                            }
                            _ = v89
                            _ = v90
                        } else {
                            fmt.Println(v53, v2)
                            v2 = (v2 & v2)
                        }
                    default:
                        if len(v13) > 0 {
                            v13[0] = (v2 >> ((0) & 63))
                        }
                    }
                case 755:
                    switch "" {
                    case "hello":
                        for fi93 := 0; fi93 < 16; fi93++ {
                            v53 += (v53 << ((18446744073709551615) & 63))
                            v2 = fi93
                            _ = fi93
                        }
                    }
                }
            }
            for fi94 := 0; fi94 < 7; fi94++ {
                _ = len(v13)
                var v95 uint = 0
                if (66 < 42) {
                    var v96 uint = v95
                    _ = v96
                } else {
                    v2 = (^(fi94))
                    var v97 float32 = (-2.25 - 1.5)
                    switch -2147483648 {
                    case -2147483648:
                        v2 -= int(v53)
                    case 722:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case -316:
                        for fi98 := 0; fi98 < 17; fi98++ {
                            if true {
                                var v99 int = genFn0(2, 89)
                                _ = len(v13)
                                _ = v99
                            }
                            v97 -= (-2.25 * (v97 / 3.14159))
                            var v100 int16 = int16(v3)
                            _ = fi98
                            _ = v100
                        }
                    default:
                        var v101 bool = ((977 > 42) || false)
                        _ = v101
                    }
                    _ = v97
                }
                _ = fi94
                _ = v95
            }
            v102 := GenS10{F0: 65535}
            _ = v102
        }
        v103 := GenS4{F0: 0, F1: -1}
        _ = v103
    } else {
        fmt.Println(v2)
        if (v3 <= v3) {
            var v104 float64 = 3.14159
            _ = v104
        }
        v53 -= v53
    }
    var v105 string = "foo bar"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v105
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
}
