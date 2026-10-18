package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int64; F1 string; F2 string }
type GenS6 struct { F0 uint16; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"foo bar", ("a" + "hello"), "\n\t", ("hello" + "hello")}
    var v3 int8 = 42
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    v6 := []string{"foo bar", ("" + ("a" + "")), "hello"}
    switch 4294967295 {
    case 0:
        v7 := GenS1{F0: -337, F1: -766, F2: 584}
        _ = v7
    case 1:
        switch "hello" {
        case "\n\t":
            v8 := GenS1{F0: 0, F1: -310, F2: 18446744073709551615}
            _ = v8
        case "a":
            _ = len(v2)
        }
    default:
        v3 = 100
    }
    var v9 string = ("foo bar" + ("a" + ""))
    v1[""] = v9
    switch -1 {
    case -2147483648:
        for fi10 := 0; fi10 < 16; fi10++ {
            if (v9 <= "foo bar") {
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch 2147483647 {
                    case -721:
                        var v12 int32 = 2147483647
                        _ = v12
                    case 32768:
                        v3 = (v3 ^ (v3 * 1))
                    case -1:
                        fmt.Println(v9, fi11)
                    }
                    var v13 uint8 = 0
                    if true {
                        v9 = v9
                        var v14 int32 = -657
                        var v15 int8 = v3
                        _ = v14
                        _ = v15
                    }
                    if len(v6) > 0 {
                        v6[0] = v9
                    }
                    v13 -= 127
                    _ = fi11
                    _ = v13
                }
            }
            fmt.Println(v9, fi10, v3)
            _ = fi10
        }
    default:
        v16 := GenS1{F0: 870, F1: 42, F2: 57}
        _ = v16
    }
    v9 = ("\n\t" + v9)
    v9 = "a"
    fmt.Println(v9, v9)
    var v17 int = -867
    var v18 float64 = (float64((1.5 - 1.5)) / ((3.14159 * float64((1.5 + 1e10))) + ((((1.5 + (float64(3.14159) * 1e10)) - 0.0) + (-2.25 * (3.14159 - (0.0 + float64(0.0))))) - 0.1)))
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: -1, F1: 1, F2: 255}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (true || (v3 == (v3 << ((1) & 63)))) {
        if false {
            switch v17 {
            case 128:
                v3 -= 100
            case 127:
                v22 := GenS1{F0: int32(v3), F1: -795, F2: 18446744073709551615}
                _ = v22
            }
        }
        var v23 uint32 = 42
        var v24 string = ("\n\t" + v9)
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 uint = 65535
            if (true || (("\n\t" >= (v9 + "\n\t")) || true)) {
                v5[32768] = false
                var v27 int64 = 9223372036854775807
                _ = v27
            }
            _ = fi25
            _ = v26
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 15; fi29++ {
                if false {
                    v9 = v9
                    v18 = (1.5 - 0.0)
                } else {
                    v30 := GenS1{F0: -2147483648, F1: 42, F2: 247}
                    var v31 float32 = 3.14159
                    switch (v9 + "") {
                    case "\n\t":
                        var v32 int64 = -1
                        _ = v32
                    }
                    _ = v30
                    _ = v31
                }
                _ = fi29
            }
            var v33 string = "foo bar"
            _ = fi28
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 int64 = -9223372036854775808
        var v36 uint8 = 42
        v37 := GenS1{F0: -1, F1: 369, F2: 42}
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    if true {
        fmt.Println(v18, v9, v17)
        if (true || true) {
            v9 = v9
            if false {
                var v38 bool = false
                _ = v38
            }
        } else {
            if (!false) {
                fmt.Println(v3)
                _ = len(v5)
                var v39 bool = false
                _ = v39
            }
            fmt.Println(v18, v3, v18)
            var v40 uint64 = 18446744073709551615
            _ = v40
        }
        v41 := GenS1{F0: 0, F1: 32767, F2: 2}
        _ = v41
    } else {
        v9 = v9
        if false {
            if ((((3.14159 - 0.0) + (3.14159 - (1e10 + ((3.14159 + float32(0.0)) - 1e10)))) >= (1e10 * (1e10 * 0.1))) && false) {
                v42 := GenS1{F0: 1, F1: 938, F2: 18446744073709551615}
                _ = v42
            }
            switch "a" {
            case "a":
                if true {
                    fmt.Println(v18)
                    for fi43 := 0; fi43 < 9; fi43++ {
                        v1["hello"] = ("\n\t" + v9)
                        v3 -= 42
                        var v44 int8 = v3
                        _ = fi43
                        _ = v44
                    }
                }
            case "\n\t":
                for fi45 := 0; fi45 < 14; fi45++ {
                    for fi46 := 0; fi46 < 12; fi46++ {
                        switch v9 {
                        case "hello":
                            if (true || true) {
                                v18 = 1.5
                                _ = len(v4)
                            } else {
                                var v47 int8 = 1
                                if (1 >= 948) {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        var v48 string = v9
                                        _ = v48
                                    }
                                    switch (^(fi46)) {
                                    case 0:
                                        if len(v2) > 0 {
                                            v2[0] = v9
                                        }
                                    default:
                                        var v49 int16 = -1
                                        _ = v49
                                    }
                                    var v50 uint = 861
                                    if false {
                                        if true {
                                            var v51 int8 = (v3 * -89)
                                            var v52 uint8 = 186
                                            var v53 int8 = -1
                                            if true {
                                                if true {
                                                    if (v9 > v9) {
                                                        if (v53 >= (v53 ^ (v3 * (v3 * (v47 & (^(v53))))))) {
                                                            v52 -= 175
                                                            if (true || (!false)) {
                                                                v54 := GenS1{F0: -1, F1: 32767, F2: 552}
                                                                if true {
                                                                    switch fi46 {
                                                                    case 548:
                                                                        if (508 <= 2) {
                                                                            var v55 int8 = 108
                                                                            _ = v55
                                                                        }
                                                                    default:
                                                                        v51 -= 0
                                                                    }
                                                                    var v56 uint32 = 170
                                                                    switch fi46 {
                                                                    case 595:
                                                                        delete(v4, (v17 ^ fi45))
                                                                    case 145:
                                                                        v52 += 0
                                                                    }
                                                                    delete(v5, (v17 + (v17 * 62)))
                                                                    if false {
                                                                        v57 := GenS1{F0: 0, F1: -1, F2: (v50 & v50)}
                                                                        v47 += (v47 * (v47 * v51))
                                                                        fmt.Println(v56, fi46)
                                                                        if false {
                                                                            v58 := v5[(^(fi45))]
                                                                            v17 += 128
                                                                            if v58 {
                                                                                var v59 uint32 = v56
                                                                                _ = v59
                                                                            } else {
                                                                                var v60 uint16 = 65535
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = ((v9 + v9) + (v9 + v9))
                                                                                }
                                                                                _ = v60
                                                                            }
                                                                            _ = v58
                                                                        } else {
                                                                            if true {
                                                                                var v61 uint32 = 65535
                                                                                _ = len(v2)
                                                                                switch (^(fi45)) {
                                                                                case 206:
                                                                                    v51 += (-(v3))
                                                                                case 127:
                                                                                    fmt.Println(v9)
                                                                                }
                                                                                _ = v61
                                                                            }
                                                                            var v62 float32 = float32(float64((1.5 * 3.14159)))
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v63 := GenS1{F0: -2147483648, F1: int16(v52), F2: (v50 | v50)}
                                                                                _ = v63
                                                                            case "\n\t":
                                                                                fmt.Println(fi46, fi45)
                                                                                _ = len(v2)
                                                                            case "a":
                                                                                var v64 bool = true
                                                                                _ = v64
                                                                            default:
                                                                                var v65 uint8 = v52
                                                                                _ = v65
                                                                            }
                                                                            _ = v62
                                                                        }
                                                                        v2 = append(v2, v9)
                                                                        _ = v57
                                                                    }
                                                                    switch (v9 + "hello") {
                                                                    case "foo bar":
                                                                        if true {
                                                                            v53 = -1
                                                                        } else {
                                                                            v56 += (v56 / ((v56) | 1))
                                                                            var v66 int16 = 42
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    _ = v56
                                                                } else {
                                                                    if true {
                                                                        v9 = "a"
                                                                        v50 = v50
                                                                        v50 = (^(v50))
                                                                        var v67 int = -9223372036854775808
                                                                        _ = v67
                                                                    }
                                                                    fmt.Println(v17, v18)
                                                                    v3 += (v3 & v51)
                                                                }
                                                                _ = v54
                                                            }
                                                        } else {
                                                            switch v9 {
                                                            case "\n\t":
                                                                fmt.Println(v17)
                                                            case "hello":
                                                                fmt.Println(v18)
                                                            default:
                                                                var v68 int = (fi46 / (((^(fi46))) | 1))
                                                                v69 := GenS1{F0: -1, F1: 0, F2: (v50 | 65535)}
                                                                v70 := GenS2{F0: true}
                                                                delete(v1, v9)
                                                                if false {
                                                                    fmt.Println(v68, v50, fi45)
                                                                    if false {
                                                                        v71 := GenS3{F0: 49}
                                                                        if (!((v18 != 3.14159) || ((fi45 * 32767) <= fi45))) {
                                                                            var v72 uint = v50
                                                                            if (((!(false || (!true))) || (!false)) && true) {
                                                                                var v73 bool = true
                                                                                _ = v73
                                                                            }
                                                                            _ = v72
                                                                        } else {
                                                                            v74 := v1[(v9 + ((("foo bar" + "") + (v9 + v9)) + v9))]
                                                                            _ = v74
                                                                        }
                                                                        if true {
                                                                            var v75 int = v68
                                                                            _ = v75
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        v76 := GenS4{F0: -1, F1: "foo bar"}
                                                                        _ = v76
                                                                    }
                                                                    fmt.Println(v47, v9)
                                                                }
                                                                v51 = (-(v3))
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v77 int8 = 50
                                                        v18 = v18
                                                        _ = v77
                                                    } else {
                                                        switch "foo bar" {
                                                        case "a":
                                                            if ((-771 == -1) && false) {
                                                                fmt.Println(v50, fi45)
                                                                var v78 string = ""
                                                                _ = v78
                                                            }
                                                        }
                                                        v17 = (fi46 & 32768)
                                                        v18 = v18
                                                    }
                                                    if (400 == 9223372036854775807) {
                                                        v79 := GenS5{F0: -82, F1: (v9 + ("a" + ("a" + v9))), F2: v9}
                                                        _ = v79
                                                    } else {
                                                        var v80 int8 = (v3 / (((v51 & (v47 ^ (v3 & v51)))) | 1))
                                                        var v81 bool = (42 >= -1)
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                } else {
                                                    var v82 uint16 = 470
                                                    v2 = append(v2, ((v9 + "") + v9))
                                                    v9 = ""
                                                    switch "a" {
                                                    case "foo bar":
                                                        v9 = v9
                                                    default:
                                                        var v83 int32 = -1
                                                        _ = v83
                                                    }
                                                    _ = v82
                                                }
                                                fmt.Println(v9)
                                                _ = len(v1)
                                            }
                                            v84 := GenS4{F0: 2147483647, F1: v9}
                                            var v85 uint64 = uint64(v50)
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                            _ = v84
                                            _ = v85
                                        } else {
                                            v1[v9] = v9
                                        }
                                        if true {
                                            var v86 int = v17
                                            _ = v86
                                        } else {
                                            var v87 int8 = (v47 * (v47 | (^(v3))))
                                            var v88 int8 = v87
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                    _ = v50
                                }
                                _ = v47
                            }
                        case "\n\t":
                            v89 := GenS1{F0: 589, F1: 0, F2: 0}
                            _ = v89
                        }
                        var v90 int = (fi46 & fi46)
                        _ = fi46
                        _ = v90
                    }
                    switch fi45 {
                    case 186:
                        v91 := GenS4{F0: -513, F1: ("\n\t" + "\n\t")}
                        _ = v91
                    case 32768:
                        var v92 int64 = -9223372036854775808
                        _ = v92
                    }
                    _ = fi45
                }
            default:
                if len(v6) > 0 {
                    v6[0] = v9
                }
            }
            delete(v1, v9)
        }
    }
    if ((!false) && (14 < 2)) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (((((v9 + (v9 + ((v9 + v9) + v9))) + v9) + v9) + "hello") + v9) {
        case "":
            v18 = 3.14159
        case "hello":
            v93 := GenS5{F0: -158, F1: (v9 + "foo bar"), F2: "foo bar"}
            _ = v93
        }
    } else {
        _ = len(v5)
    }
    switch v9 {
    case "\n\t":
        var v94 bool = true
        _ = v94
    default:
        for fi95 := 0; fi95 < 1; fi95++ {
            var v96 int8 = (^(v3))
            var v97 float64 = 1.5
            switch (fi95 / ((v17) | 1)) {
            case 864:
                for fi98 := 0; fi98 < 10; fi98++ {
                    fmt.Println(fi98, v17)
                    v97 = (1e10 * v18)
                    if (2147483647 > 414) {
                        switch v9 {
                        case "foo bar":
                            v99 := GenS1{F0: 42, F1: 1, F2: 2}
                            _ = v99
                        case "hello":
                            var v100 uint16 = 65535
                            v101 := GenS3{F0: 564}
                            _ = v100
                            _ = v101
                        case "":
                            switch (fi95 ^ (fi95 % ((9223372036854775807) | 1))) {
                            case 42:
                                _ = len(v4)
                                v102 := GenS3{F0: 0}
                                v97 -= (1.5 + ((float64(0.1) + 3.14159) * ((v97 + (1e10 + v18)) / 3.14159)))
                                _ = v102
                            default:
                                var v103 int16 = 42
                                _ = v103
                            }
                        default:
                            if len(v2) > 0 {
                                v2[0] = (v9 + "a")
                            }
                        }
                        v104 := GenS6{F0: 42, F1: 0}
                        _ = v104
                    }
                    var v105 uint32 = 1
                    v105 -= 881
                    _ = fi98
                    _ = v105
                }
            }
            _ = fi95
            _ = v96
            _ = v97
        }
    }
    for fi106 := 0; fi106 < 6; fi106++ {
        v18 -= 0.0
        _ = fi106
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
