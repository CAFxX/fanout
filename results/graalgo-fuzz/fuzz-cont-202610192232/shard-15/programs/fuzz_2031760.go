package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64; F1 int16; F2 uint }

func main() {
    var v1 bool = true
    var v2 int8 = 0
    var v3 int64 = 0
    v3 -= v3
    switch 42 {
    case -1:
        v4 := GenS1{F0: 42, F1: v1}
        _ = v4
    }
    var v5 uint32 = 2
    var v6 int8 = v2
    v5 = (v5 * v5)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            v2 = -32
            fmt.Println(v6)
            _ = fi8
        }
        _ = fi7
    }
    if v1 {
        switch -1 {
        case 1:
            var v9 int16 = 1
            _ = v9
        default:
            v10 := GenS1{F0: (v5 + v5), F1: (((18446744073709551615 <= 0) && true) && false)}
            _ = v10
        }
        if (127 != 1) {
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 2147483647, F1: (1 <= 2147483647)}
                v13 := GenS1{F0: (^(v5)), F1: false}
                if (!v1) {
                    v14 := []int{128}
                    _ = v14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v1, v5)
        }
    }
    var v15 int8 = v2
    switch -9223372036854775808 {
    case 4294967295:
        if (0 != 1) {
            v16 := GenS2{F0: -32768, F1: (v5 | v5)}
            if false {
                var v17 float64 = (1e10 / 1e10)
                if (v15 <= (v6 / (((v15 * 0)) | 1))) {
                    v1 = false
                    var v18 float64 = v17
                    _ = v18
                }
                _ = v17
            }
            _ = v16
        } else {
            var v19 string = "hello"
            v2 -= v2
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(v15, v15, v19)
                v6 = int8(v3)
                switch (((v19 + "hello") + ((v19 + "hello") + v19)) + ("\n\t" + v19)) {
                case "a":
                    v21 := []int{fi20, 2147483647, fi20, -521}
                    _ = v21
                }
                _ = fi20
            }
            _ = v19
        }
    case -550:
        for fi22 := 0; fi22 < 1; fi22++ {
            switch fi22 {
            case 42:
                if false {
                    v23 := GenS2{F0: -32768, F1: v5}
                    _ = v23
                } else {
                    v24 := []int{32768, fi22, int(v5)}
                    _ = v24
                }
            case 32767:
                v6 = v6
            case 128:
                for fi25 := 0; fi25 < 10; fi25++ {
                    v26 := GenS2{F0: -32768, F1: (^(v5))}
                    fmt.Println(v5, v6)
                    switch 128 {
                    case 2147483647:
                        v27 := []int{fi25, fi25, fi22, fi25}
                        _ = v27
                    case -364:
                        switch fi25 {
                        case 4294967295:
                            v6 = v15
                        default:
                            v5 = 2147483647
                        }
                    default:
                        v28 := []int{fi25, (fi22 << ((548) & 63))}
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
            }
            if (v1 && false) {
                v29 := GenS1{F0: v5, F1: ((!true) && false)}
                var v30 uint32 = 221
                switch 32768 {
                case 0:
                    switch ((("foo bar" + "") + ("foo bar" + "\n\t")) + "foo bar") {
                    case "":
                        v31 := GenS3{F0: 42}
                        _ = v31
                    case "a":
                        var v32 bool = true
                        _ = v32
                    case "hello":
                        v5 += 42
                    }
                case 128:
                    fmt.Println(v2)
                }
                _ = v29
                _ = v30
            }
            var v33 uint = 381
            _ = fi22
            _ = v33
        }
    case 439:
        var v34 int64 = 9223372036854775807
        _ = v34
    }
    v1 = true
    fmt.Println(v6, v1)
    var v35 int8 = v2
    v36 := GenS4{F0: 1}
    v15 += 0
    switch ("hello" + "hello") {
    case "hello":
        fmt.Println(v3)
    case "foo bar":
        v37 := []int{-87, 4294967295, 1, int(v5)}
        _ = v37
    case "":
        v6 = (v15 % (((v6 - (v15 - v35))) | 1))
    }
    fmt.Println(v35, v5, v35)
    if true {
        switch 127 {
        case -597:
            v38 := []int{2147483647, 32768}
            _ = v38
        case -151:
            var v39 uint8 = 2
            _ = v39
        case 9223372036854775807:
            for fi40 := 0; fi40 < 12; fi40++ {
                v5 = v5
                v35 = -128
                for fi41 := 0; fi41 < 16; fi41++ {
                    v42 := []int{}
                    fmt.Println(v3)
                    _ = fi41
                    _ = v42
                }
                _ = fi40
            }
        }
    } else {
        for fi43 := 0; fi43 < 5; fi43++ {
            v44 := []int{fi43, 9223372036854775807}
            if ((v5 << ((398) & 63)) < 1) {
                var v45 int32 = -2147483648
                _ = v45
            } else {
                switch fi43 {
                case 0:
                    if v1 {
                        if true {
                            var v46 string = "\n\t"
                            _ = len(v44)
                            for fi47 := 0; fi47 < 4; fi47++ {
                                if (534 >= 4294967295) {
                                    v5 -= (v5 << ((793) & 63))
                                    fmt.Println(fi43)
                                    if ((v5 << ((255) & 63)) != (v5 ^ (^(v5)))) {
                                        var v48 bool = v1
                                        if (!true) {
                                            switch fi43 {
                                            case -1:
                                                var v49 uint8 = 2
                                                _ = v49
                                            }
                                        }
                                        fmt.Println(v6, fi47)
                                        v2 = (v2 * v6)
                                        if (v15 >= v15) {
                                            v50 := GenS5{F0: 4294967295, F1: -1}
                                            var v51 int16 = -571
                                            _ = len(v44)
                                            _ = v50
                                            _ = v51
                                        }
                                        var v52 string = ("hello" + "")
                                        _ = v48
                                        _ = v52
                                    }
                                }
                                if false {
                                    v53 := GenS2{F0: -1, F1: v5}
                                    var v54 uint8 = 0
                                    v55 := GenS5{F0: (v5 - v5), F1: -32768}
                                    var v56 uint64 = 4294967295
                                    switch fi47 {
                                    case -2147483648:
                                        if true {
                                            if (v1 || true) {
                                                switch (^(fi47)) {
                                                case 9223372036854775807:
                                                    var v57 int32 = -832
                                                    _ = v57
                                                case 42:
                                                    var v58 string = ""
                                                    _ = v58
                                                default:
                                                    var v59 string = (v46 + (v46 + "\n\t"))
                                                    v60 := GenS4{F0: v5}
                                                    _ = v59
                                                    _ = v60
                                                }
                                                if v1 {
                                                    v61 := GenS3{F0: v2}
                                                    _ = v61
                                                } else {
                                                    switch v46 {
                                                    case "a":
                                                        switch 127 {
                                                        case 128:
                                                            if len(v44) > 0 {
                                                                v44[0] = 1
                                                            }
                                                        case 43:
                                                            if v1 {
                                                                switch (("hello" + v46) + "") {
                                                                case "":
                                                                    v62 := GenS6{F0: uint32(v56)}
                                                                    _ = v62
                                                                }
                                                                var v63 uint = 666
                                                                _ = v63
                                                            }
                                                        default:
                                                            if (70 != v54) {
                                                                if (fi43 != 908) {
                                                                    var v64 uint = 822
                                                                    fmt.Println(fi47, v6)
                                                                    var v65 float32 = 0.1
                                                                    _ = v64
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        v66 := GenS7{F0: 2}
                                                        _ = v66
                                                    case "foo bar":
                                                        switch (fi47 << ((816) & 63)) {
                                                        case 42:
                                                            v44 = append(v44, fi43)
                                                        case 693:
                                                            v67 := GenS1{F0: v5, F1: true}
                                                            _ = v67
                                                        }
                                                    }
                                                    if (v1 && (!(true || v1))) {
                                                        v68 := GenS8{F0: -94}
                                                        _ = v68
                                                    } else {
                                                        v44 = append(v44, fi43)
                                                        switch fi43 {
                                                        case 42:
                                                            if ((v2 / ((v6) | 1)) >= (-(v35))) {
                                                                fmt.Println(v35, v56, v2)
                                                                fmt.Println(v5, v56)
                                                            } else {
                                                                v15 = (v2 & (-(v15)))
                                                                _ = len(v44)
                                                                if len(v44) >= 2 {
                                                                    v44 = v44[1:len(v44)]
                                                                }
                                                            }
                                                        case 0:
                                                            var v69 int8 = v2
                                                            _ = v69
                                                        case 127:
                                                            v70 := GenS1{F0: (^(v5)), F1: false}
                                                            _ = v70
                                                        }
                                                    }
                                                }
                                            } else {
                                                switch -9223372036854775808 {
                                                case 9223372036854775807:
                                                    v54 += v54
                                                default:
                                                    v15 = (v35 << ((18446744073709551615) & 63))
                                                }
                                            }
                                            v1 = true
                                        } else {
                                            v15 -= v35
                                            if v1 {
                                                v71 := GenS9{F0: -924}
                                                _ = v71
                                            } else {
                                                v72 := GenS3{F0: 0}
                                                fmt.Println(v56)
                                                switch "a" {
                                                case "a":
                                                    if v1 {
                                                        var v73 uint16 = 32767
                                                        _ = len(v44)
                                                        _ = v73
                                                    } else {
                                                        v56 += v56
                                                        var v74 int = 127
                                                        switch "" {
                                                        case "\n\t":
                                                            if len(v44) >= 2 {
                                                                v44 = v44[1:len(v44)]
                                                            }
                                                        case "a":
                                                            v56 = (v56 << ((1) & 63))
                                                        }
                                                        _ = v74
                                                    }
                                                }
                                                _ = v72
                                            }
                                        }
                                    }
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                                var v75 int64 = (v3 - v3)
                                _ = fi47
                                _ = v75
                            }
                            _ = v46
                        }
                        var v76 float32 = float32(1e10)
                        var v77 int8 = v6
                        _ = v76
                        _ = v77
                    } else {
                        fmt.Println(fi43, v15)
                    }
                }
                var v78 int = fi43
                _ = v78
            }
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
            _ = fi43
            _ = v44
        }
        switch ((("\n\t" + "a") + "hello") + "\n\t") {
        case "\n\t":
            for fi79 := 0; fi79 < 7; fi79++ {
                if false {
                    for fi80 := 0; fi80 < 9; fi80++ {
                        v81 := []int{}
                        switch 42 {
                        case 4294967295:
                            v15 = 0
                        case -9223372036854775808:
                            if false {
                                v82 := GenS5{F0: v5, F1: 32767}
                                switch (fi79 ^ 2147483647) {
                                case 128:
                                    if v1 {
                                        var v83 bool = (!false)
                                        v1 = v83
                                        if v83 {
                                            if len(v81) > 0 {
                                                v81[0] = fi79
                                            }
                                        } else {
                                            v84 := GenS8{F0: -2147483648}
                                            var v85 float64 = 1e10
                                            _ = v84
                                            _ = v85
                                        }
                                        _ = v83
                                    }
                                case -316:
                                    fmt.Println(v5, v6)
                                }
                                _ = v82
                            } else {
                                if v1 {
                                    var v86 int8 = (v6 << ((1) & 63))
                                    _ = len(v81)
                                    switch -9223372036854775808 {
                                    case -372:
                                        if (0.0 <= 1.5) {
                                            v87 := GenS10{F0: v3}
                                            if len(v81) > 0 {
                                                v81[0] = fi79
                                            }
                                            v88 := GenS9{F0: int32(v3)}
                                            if (!(1 > 4294967295)) {
                                                v1 = (true || (v1 || false))
                                                var v89 int8 = (v86 & v35)
                                                var v90 uint8 = 0
                                                _ = v89
                                                _ = v90
                                            }
                                            _ = v87
                                            _ = v88
                                        }
                                    default:
                                        v6 -= (v15 * v86)
                                    }
                                    _ = v86
                                } else {
                                    var v91 string = "foo bar"
                                    fmt.Println(v6)
                                    var v92 bool = ((0.1 * (0.1 * (1e10 / (1.5 * ((3.14159 * 0.0) - 3.14159))))) > (1.5 + (1e10 + 3.14159)))
                                    _ = v91
                                    _ = v92
                                }
                            }
                        case -274:
                            switch fi79 {
                            case -1:
                                switch fi79 {
                                case 4294967295:
                                    v81 = append(v81, fi80)
                                case -378:
                                    if false {
                                        v35 = (v35 * v6)
                                        switch "hello" {
                                        case "":
                                            switch ("foo bar" + "foo bar") {
                                            case "":
                                                v93 := GenS2{F0: -32768, F1: (^(v5))}
                                                _ = v93
                                            }
                                            v2 = (v35 + 127)
                                        }
                                    }
                                }
                            case -2147483648:
                                v3 -= (v3 << ((18446744073709551615) & 63))
                            case 32767:
                                v94 := GenS7{F0: 427}
                                _ = v94
                            }
                        }
                        _ = fi80
                        _ = v81
                    }
                } else {
                    for fi95 := 0; fi95 < 15; fi95++ {
                        var v96 uint8 = 42
                        _ = fi95
                        _ = v96
                    }
                    var v97 uint16 = 0
                    var v98 string = "foo bar"
                    _ = v97
                    _ = v98
                }
                if ((fi79 | 9223372036854775807) == (fi79 % (((fi79 - (fi79 / ((fi79) | 1)))) | 1))) {
                    switch (fi79 >> ((32) & 63)) {
                    case 1:
                        switch "" {
                        case "hello":
                            var v99 bool = (123 <= 65535)
                            _ = v99
                        case "\n\t":
                            v100 := GenS11{F0: v3, F1: -32768, F2: 65535}
                            _ = v100
                        case "a":
                            var v101 int8 = int8(fi79)
                            _ = v101
                        }
                    }
                    var v102 uint64 = 719
                    _ = v102
                }
                _ = fi79
            }
        default:
            switch 596 {
            case -9223372036854775808:
                var v103 int32 = 42
                _ = v103
            case 923:
                v104 := []int{}
                _ = v104
            default:
                v3 = (v3 >> ((2) & 63))
            }
        }
        var v105 float32 = (3.14159 + 3.14159)
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v5
    _ = v6
}
