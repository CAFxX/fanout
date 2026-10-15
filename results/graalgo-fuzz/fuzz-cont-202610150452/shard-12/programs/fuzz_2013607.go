package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint64; F1 int; F2 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 string; F1 uint64 }
type GenS14 struct { F0 int16; F1 int8; F2 uint }

func main() {
    v1 := []int{9223372036854775807, 32768, -9223372036854775808}
    var v2 string = "hello"
    var v3 uint16 = 65535
    var v4 uint16 = v3
    var v5 float64 = 3.14159
    switch "hello" {
    case "\n\t":
        v5 = 3.14159
    case "foo bar":
        fmt.Println(v5)
    default:
        v6 := GenS1{F0: (v2 + v2), F1: 42, F2: 144}
        _ = v6
    }
    switch "\n\t" {
    case "\n\t":
        v2 = (v2 + v2)
    case "foo bar":
        fmt.Println(v4, v5, v5)
    case "a":
        v1 = append(v1, 880)
    }
    v1 = append(v1, 42)
    v5 += (v5 / v5)
    v2 = (((v2 + v2) + v2) + v2)
    v3 = 32767
    var v7 string = ""
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint16 = v3
        v4 = 2
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 4294967295, F1: 1, F2: 949}
    var v11 uint16 = 860
    fmt.Println(v2)
    switch -906 {
    case 907:
        for fi12 := 0; fi12 < 12; fi12++ {
            if ((!((v2 + v7) < (v2 + "\n\t"))) || false) {
                v5 -= v5
            }
            var v13 float32 = 1.5
            switch v7 {
            case "hello":
                if true {
                    if true {
                        v14 := GenS2{F0: 4294967295, F1: 128, F2: 1}
                        if false {
                            v15 := GenS2{F0: 2147483647, F1: 42, F2: 792}
                            v5 += v5
                            _ = v15
                        }
                        var v16 int = 4294967295
                        _ = v14
                        _ = v16
                    }
                    if (1 < 128) {
                        v7 = v7
                    } else {
                        var v17 uint32 = 1
                        _ = v17
                    }
                }
            case "\n\t":
                if true {
                    var v18 uint32 = 0
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v20 := GenS3{F0: (fi12 + -901)}
                        switch fi12 {
                        case 452:
                            v13 -= v13
                        }
                        switch 9223372036854775807 {
                        case 1:
                            fmt.Println(v4, v11, v3)
                        default:
                            v13 = float32((-2.25 - v5))
                        }
                        _ = fi19
                        _ = v20
                    }
                    _ = len(v1)
                    _ = v18
                } else {
                    v21 := GenS4{F0: 255, F1: true}
                    switch ((v7 + "a") + v2) {
                    case "foo bar":
                        if true {
                            switch fi12 {
                            case 115:
                                if len(v1) > 0 {
                                    v1[0] = -1
                                }
                            case 9223372036854775807:
                                var v22 int16 = -211
                                _ = v22
                            default:
                                fmt.Println(fi12)
                            }
                        }
                    case "a":
                        var v23 int16 = 792
                        _ = v23
                    case "":
                        var v24 uint8 = 84
                        _ = v24
                    default:
                        switch (v7 + v7) {
                        case "foo bar":
                            for fi25 := 0; fi25 < 17; fi25++ {
                                switch v2 {
                                case "":
                                    var v26 float64 = (((v5 * 0.0) - v5) * float64(-2.25))
                                    _ = v26
                                case "a":
                                    v11 = (v3 + (v3 ^ 754))
                                case "\n\t":
                                    if len(v1) > 0 {
                                        v1[0] = (-(fi25))
                                    }
                                }
                                v11 -= (^(v3))
                                switch (fi12 % (((fi12 + fi25)) | 1)) {
                                case 42:
                                    switch (fi12 << ((576) & 63)) {
                                    case 350:
                                        if (2147483647 == 1) {
                                            var v27 int = 32768
                                            v28 := GenS5{F0: 2}
                                            var v29 int16 = -145
                                            v30 := GenS1{F0: (v7 + "hello"), F1: 103, F2: 42}
                                            _ = len(v1)
                                            fmt.Println(fi25, v5)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                    case -512:
                                        var v31 int16 = -13
                                        _ = v31
                                    }
                                case 1:
                                    if (1 < 6) {
                                        v1 = append(v1, fi25)
                                    } else {
                                        switch (v7 + v2) {
                                        case "foo bar":
                                            v32 := GenS5{F0: 4294967295}
                                            _ = v32
                                        case "":
                                            v1 = append(v1, (fi25 | 32768))
                                        default:
                                            var v33 bool = (42 == 0)
                                            _ = v33
                                        }
                                        v34 := GenS1{F0: ("hello" + (v2 + ((v7 + v7) + v2))), F1: -9223372036854775808, F2: 4294967295}
                                        var v35 bool = false
                                        _ = v34
                                        _ = v35
                                    }
                                case 105:
                                    switch (v2 + "\n\t") {
                                    case "a":
                                        if (4294967295 < 42) {
                                            var v36 uint8 = 1
                                            if true {
                                                v1 = append(v1, (fi12 - (fi12 << ((4294967295) & 63))))
                                                switch v2 {
                                                case "":
                                                    v4 = 1
                                                case "foo bar":
                                                    if false {
                                                        v5 = v5
                                                    } else {
                                                        var v37 uint = 645
                                                        v38 := GenS6{F0: (fi12 * 2147483647), F1: (((v7 + v2) + v7) + (((v2 + "a") + "") + (v7 + v2)))}
                                                        if true {
                                                            fmt.Println(fi25, v5, v37)
                                                        }
                                                        var v39 float32 = v13
                                                        if false {
                                                            v40 := GenS7{F0: 0}
                                                            if (!((v7 <= ((v7 + "hello") + v2)) && (!(1000 >= int16(v37))))) {
                                                                v41 := GenS5{F0: v37}
                                                                var v42 int8 = 1
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                            _ = v40
                                                        }
                                                        v2 = ("\n\t" + "a")
                                                        _ = v37
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                }
                                            }
                                            fmt.Println(v11, v13, fi25)
                                            _ = v36
                                        } else {
                                            v4 = (v11 & v4)
                                            v13 = ((v13 / 1.0) * (((v13 - 0.0) / float32(1.5)) / (3.14159 - 1.5)))
                                            switch "a" {
                                            case "":
                                                var v43 float32 = (-2.25 - v13)
                                                _ = v43
                                            case "hello":
                                                v44 := GenS5{F0: 5}
                                                _ = v44
                                            case "foo bar":
                                                if false {
                                                    fmt.Println(fi12)
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v3, fi25, v7)
                                    var v45 float64 = 1.5
                                    var v46 float64 = ((v45 + 3.14159) - (1.5 * (v45 - (float64(((0.0 + v13) - (float32((v5 / ((3.14159 * 3.14159) - v5))) * v13))) + v45))))
                                    var v47 float64 = 3.14159
                                    v11 -= (v11 % (((v4 + (v4 ^ 32767))) | 1))
                                    if (false || false) {
                                        var v48 uint = 536
                                        v5 = (0.0 + v47)
                                        switch v7 {
                                        case "":
                                            v49 := GenS8{F0: 127, F1: 1}
                                            _ = v49
                                        case "foo bar":
                                            v3 = v11
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        default:
                                            fmt.Println(v13, v11, v2)
                                        }
                                        _ = v48
                                    } else {
                                        v2 = (v2 + v2)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v50 := GenS6{F0: fi12, F1: v2}
                                        var v51 int16 = 843
                                        switch "foo bar" {
                                        case "":
                                            if (!false) {
                                                var v52 uint = 4294967295
                                                v53 := GenS1{F0: v7, F1: 0, F2: 42}
                                                _ = v52
                                                _ = v53
                                            } else {
                                                v51 = 42
                                            }
                                        case "foo bar":
                                            var v54 uint = 42
                                            _ = v54
                                        }
                                        _ = v50
                                        _ = v51
                                    }
                                    if (!false) {
                                        v55 := GenS2{F0: 779, F1: 127, F2: -1}
                                        v13 -= v13
                                        if true {
                                            if (-9223372036854775808 >= (^(fi25))) {
                                                var v56 bool = (!(v3 <= v11))
                                                if (!true) {
                                                    v56 = (0.0 != v13)
                                                } else {
                                                    var v57 int = fi12
                                                    v58 := GenS1{F0: "a", F1: -1, F2: 495}
                                                    _ = v57
                                                    _ = v58
                                                }
                                                var v59 float64 = (((float64(v13) + v5) - v46) - (((float64(float32(0.1)) + v45) / (3.14159 * 1.5)) * v5))
                                                _ = v56
                                                _ = v59
                                            } else {
                                                if (v7 > "") {
                                                    v60 := GenS9{F0: 202}
                                                    if true {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        var v61 bool = (!true)
                                                        _ = v61
                                                    }
                                                    var v62 int16 = 922
                                                    _ = v60
                                                    _ = v62
                                                } else {
                                                    var v63 uint8 = 91
                                                    _ = v63
                                                }
                                                var v64 int16 = 1
                                                v65 := GenS1{F0: (("foo bar" + v2) + "\n\t"), F1: -1, F2: 4294967295}
                                                _ = v64
                                                _ = v65
                                            }
                                            v66 := GenS10{F0: uint64(v3)}
                                            v67 := GenS5{F0: 0}
                                            _ = v66
                                            _ = v67
                                        }
                                        _ = v55
                                    } else {
                                        switch ("hello" + ((("" + (v2 + (("\n\t" + "") + v7))) + ("foo bar" + v7)) + v2)) {
                                        case "foo bar":
                                            if (!(2147483648 != 42)) {
                                                v68 := GenS1{F0: (((v2 + ("foo bar" + (v2 + v2))) + ((v2 + v2) + v7)) + v7), F1: int64(fi12), F2: 2147483647}
                                                if (true && false) {
                                                    var v69 uint8 = 251
                                                    var v70 uint16 = v3
                                                    v71 := GenS6{F0: (fi25 % ((fi12) | 1)), F1: v2}
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                } else {
                                                    var v72 float64 = v47
                                                    _ = v72
                                                }
                                                if true {
                                                    fmt.Println(v46, v45, v45)
                                                    v2 = v7
                                                }
                                                _ = v68
                                            } else {
                                                v1 = append(v1, fi12)
                                                var v73 int32 = 922
                                                _ = v73
                                            }
                                        case "\n\t":
                                            v3 += v3
                                        case "":
                                            if (42 != 9223372036854775807) {
                                                var v74 int = fi25
                                                _ = v74
                                            }
                                        }
                                        switch ((("foo bar" + ((v7 + (v7 + v2)) + "foo bar")) + ((v2 + v7) + (v2 + ((v2 + v2) + ("foo bar" + v2))))) + v7) {
                                        case "\n\t":
                                            if (676 <= 1) {
                                                v7 = v2
                                                var v75 uint8 = 0
                                                _ = v75
                                            }
                                        case "a":
                                            if false {
                                                v76 := GenS11{F0: 799, F1: (fi25 * (fi25 * 980)), F2: v11}
                                                fmt.Println(v2, v47, v3)
                                                _ = v76
                                            }
                                        case "hello":
                                            v77 := GenS3{F0: fi12}
                                            _ = v77
                                        default:
                                            v78 := GenS12{F0: 42}
                                            _ = v78
                                        }
                                        fmt.Println(v47, fi25)
                                    }
                                    v79 := GenS13{F0: v7, F1: 1}
                                    if true {
                                        var v80 uint8 = 146
                                        var v81 int8 = 117
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v80
                                        _ = v81
                                    }
                                    var v82 uint = 1
                                    var v83 int = 32767
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                    _ = v79
                                    _ = v82
                                    _ = v83
                                }
                                v2 = (v2 + v7)
                                v13 += (v13 * v13)
                                _ = fi25
                            }
                        case "a":
                            v2 = v7
                        default:
                            if false {
                                v4 += 42
                                v7 = (v7 + (("a" + (v2 + v7)) + v7))
                                if len(v1) > 0 {
                                    v1[0] = 9223372036854775807
                                }
                            } else {
                                v5 = 0.1
                                var v84 uint32 = 986
                                _ = len(v1)
                                _ = v84
                            }
                        }
                    }
                    if true {
                        fmt.Println(v7, v5)
                        for fi85 := 0; fi85 < 15; fi85++ {
                            switch "hello" {
                            case "a":
                                v11 -= v4
                            case "":
                                var v86 int16 = -1
                                _ = v86
                            default:
                                if len(v1) > 0 {
                                    v1[0] = fi12
                                }
                            }
                            if (true || true) {
                                var v87 uint16 = (v3 >> ((42) & 63))
                                var v88 uint = 42
                                _ = v87
                                _ = v88
                            }
                            _ = fi85
                        }
                        v4 -= v11
                    }
                    _ = v21
                }
            default:
                v13 += 1e10
            }
            _ = fi12
            _ = v13
        }
    case 2147483647:
        for fi89 := 0; fi89 < 17; fi89++ {
            var v90 int16 = int16(fi89)
            _ = fi89
            _ = v90
        }
    default:
        for fi91 := 0; fi91 < 6; fi91++ {
            if (v5 != (v5 * v5)) {
                switch (v7 + v2) {
                case "hello":
                    var v92 uint16 = v4
                    _ = v92
                case "":
                    switch v2 {
                    case "":
                        var v93 int32 = -514
                        _ = v93
                    case "a":
                        for fi94 := 0; fi94 < 20; fi94++ {
                            var v95 uint8 = 137
                            var v96 uint64 = 700
                            var v97 uint32 = 2147483647
                            _ = fi94
                            _ = v95
                            _ = v96
                            _ = v97
                        }
                    case "\n\t":
                        for fi98 := 0; fi98 < 15; fi98++ {
                            v99 := GenS10{F0: 18446744073709551615}
                            v3 = (^(v4))
                            _ = fi98
                            _ = v99
                        }
                    default:
                        v4 = v3
                    }
                default:
                    v4 = v4
                }
                for fi100 := 0; fi100 < 15; fi100++ {
                    var v101 uint16 = (v11 & (v11 / ((v11) | 1)))
                    _ = len(v1)
                    _ = fi100
                    _ = v101
                }
                var v102 float64 = ((v5 * ((v5 * ((1e10 + ((-2.25 - v5) / v5)) + (1e10 / 1e10))) + v5)) - float64(1e10))
                _ = v102
            }
            _ = fi91
        }
    }
    for fi103 := 0; fi103 < 20; fi103++ {
        if (true && true) {
            v11 += 0
            v4 = v4
            v3 = (v4 | 42)
        }
        for fi104 := 0; fi104 < 19; fi104++ {
            var v105 uint = 255
            _ = fi104
            _ = v105
        }
        v2 = v7
        _ = fi103
    }
    v106 := GenS14{F0: 42, F1: 42, F2: 2}
    _ = len(v1)
    v107 := GenS3{F0: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v106
    _ = v107
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
