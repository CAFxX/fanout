package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 int64; F1 int64; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS8 struct { F0 int64; F1 int; F2 int }
type GenS9 struct { F0 int; F1 uint8; F2 uint64 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS11 struct { F0 bool; F1 int64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint; F1 int32 }

func main() {
    var v1 int16 = 32767
    var v2 int32 = 1
    var v3 uint32 = 2147483647
    v4 := make(map[string]uint32)
    v5 := []uint32{v3, (v3 * (v3 - v3)), v3}
    v6 := GenS1{F0: 0}
    switch -979 {
    case -558:
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int8 = -53
            _ = fi7
            _ = v8
        }
    case 0:
        var v9 int16 = v1
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 -= v1
        v11 := v4["a"]
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 1
    if true {
        switch "foo bar" {
        case "":
            fmt.Println(v2, v12, v2)
        default:
            switch "foo bar" {
            case "foo bar":
                if (v3 < v3) {
                    v1 = v1
                    switch 42 {
                    case 127:
                        if (42 > 442) {
                            var v13 int8 = -1
                            v14 := GenS1{F0: 1}
                            _ = v13
                            _ = v14
                        }
                    case 0:
                        fmt.Println(v12, v2)
                    case 128:
                        for fi15 := 0; fi15 < 2; fi15++ {
                            if false {
                                if false {
                                    v1 += (v1 >> ((65535) & 63))
                                    v16 := GenS1{F0: 1}
                                    _ = v16
                                } else {
                                    fmt.Println(v2)
                                }
                                if false {
                                    fmt.Println(v2, v1)
                                    switch fi15 {
                                    case -9223372036854775808:
                                        fmt.Println(v12, v1)
                                    case 128:
                                        var v17 float32 = ((1e10 / 1.0) + (float32((1.5 + (1e10 * float64(-2.25)))) - (3.14159 - (0.0 + 0.0))))
                                        _ = v17
                                    default:
                                        var v18 int64 = -9223372036854775808
                                        _ = v18
                                    }
                                } else {
                                    v1 = (v1 - (v1 >> ((42) & 63)))
                                    for fi19 := 0; fi19 < 9; fi19++ {
                                        v20 := GenS1{F0: -883}
                                        _ = fi19
                                        _ = v20
                                    }
                                    for fi21 := 0; fi21 < 20; fi21++ {
                                        v5 = append(v5, v3)
                                        _ = fi21
                                    }
                                }
                            }
                            v22 := GenS1{F0: 9223372036854775807}
                            for fi23 := 0; fi23 < 11; fi23++ {
                                fmt.Println(v2, fi23)
                                var v24 int16 = v1
                                _ = fi23
                                _ = v24
                            }
                            _ = fi15
                            _ = v22
                        }
                    }
                    var v25 bool = true
                    _ = v25
                } else {
                    switch 9223372036854775807 {
                    case 32768:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            fmt.Println(v3, fi26)
                            _ = fi26
                        }
                    default:
                        v5 = append(v5, v3)
                    }
                    var v27 uint = 644
                    _ = v27
                }
            }
        }
        var v28 float32 = 1.5
        var v29 uint16 = (v12 ^ v12)
        _ = v28
        _ = v29
    } else {
        var v30 uint = 0
        _ = v30
    }
    v2 = -640
    for fi31 := 0; fi31 < 1; fi31++ {
        if ("foo bar" == "hello") {
            var v32 int64 = 1
            _ = v32
        }
        fmt.Println(v12, v1, v12)
        _ = fi31
    }
    fmt.Println(v2)
    if false {
        var v33 bool = true
        fmt.Println(v33, v33)
        if false {
            v12 = (v12 ^ (v12 >> ((1) & 63)))
        } else {
            var v34 int32 = v2
            for fi35 := 0; fi35 < 8; fi35++ {
                delete(v4, "foo bar")
                var v36 int8 = 100
                v1 -= (v1 + (v1 | (v1 ^ int16(v12))))
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        _ = v33
    }
    for fi37 := 0; fi37 < 18; fi37++ {
        switch ("foo bar" + "hello") {
        case "hello":
            v1 -= v1
        }
        var v38 float32 = 1e10
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: (271 != (v2 - (v2 * (v2 - (v2 >> ((42) & 63)))))), F1: ("foo bar" + "\n\t")}
    switch "\n\t" {
    case "foo bar":
        if (v2 <= v2) {
            v3 -= 0
            if ((18446744073709551615 >= 4294967295) && (false && (!true))) {
                v40 := GenS1{F0: 9223372036854775807}
                v3 = 4294967295
                if false {
                    v41 := GenS3{F0: -109, F1: 4294967295}
                    v42 := GenS4{F0: 42, F1: 115, F2: (v1 % (((v1 | v1)) | 1))}
                    _ = v41
                    _ = v42
                } else {
                    var v43 string = ""
                    _ = v43
                }
                _ = v40
            }
        }
    case "":
        var v44 uint32 = 0
        _ = v44
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        v1 = v1
    case -59:
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            fmt.Println(v2, v1, v1)
        case "a":
            var v45 uint64 = 1
            _ = v45
        default:
            if true {
                var v46 float32 = 1.5
                v47 := GenS5{F0: 83}
                if false {
                    if true {
                        v2 = v2
                    } else {
                        v48 := GenS4{F0: 1, F1: 9223372036854775807, F2: v1}
                        v49 := GenS5{F0: 255}
                        if true {
                            for fi50 := 0; fi50 < 18; fi50++ {
                                v51 := GenS6{F0: 42, F1: fi50}
                                var v52 uint32 = 2
                                if (2147483647 <= v3) {
                                    fmt.Println(fi50, fi50, v3)
                                } else {
                                    switch fi50 {
                                    case 2147483647:
                                        v53 := GenS7{F0: uint8(v12), F1: 1000, F2: 60}
                                        _ = v53
                                    }
                                }
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            var v54 int32 = v2
                            if false {
                                if (255 < 42) {
                                    switch 9223372036854775807 {
                                    case 32768:
                                        var v55 string = "hello"
                                        _ = v55
                                    default:
                                        var v56 int = -283
                                        _ = v56
                                    }
                                    v57 := GenS8{F0: 615, F1: 880, F2: -931}
                                    _ = v57
                                } else {
                                    _ = len(v4)
                                }
                                v46 += v46
                                if true {
                                    fmt.Println(v46)
                                    v58 := v4[(("" + "hello") + "a")]
                                    _ = len(v4)
                                    _ = v58
                                }
                            }
                            _ = v54
                        }
                        _ = v48
                        _ = v49
                    }
                    if (!((!true) && (float64(v46) > 1e10))) {
                        v59 := GenS9{F0: 891, F1: 127, F2: 253}
                        if (166 < 201) {
                            fmt.Println(v12)
                        }
                        _ = v59
                    }
                } else {
                    if (v2 < v2) {
                        v2 = v2
                        for fi60 := 0; fi60 < 11; fi60++ {
                            fmt.Println(v46, fi60)
                            _ = fi60
                        }
                        if true {
                            v3 += (v3 * 741)
                            switch ("" + "foo bar") {
                            case "\n\t":
                                v46 += v46
                            case "foo bar":
                                var v61 float32 = (0.1 - v46)
                                _ = v61
                            }
                            var v62 int64 = 442
                            _ = v62
                        } else {
                            var v63 int32 = v2
                            _ = v63
                        }
                    }
                }
                _ = v46
                _ = v47
            }
        }
    case -714:
        fmt.Println(v12)
    default:
        switch ("a" + "\n\t") {
        case "\n\t":
            var v64 bool = (32767 <= 638)
            _ = v64
        case "hello":
            switch "foo bar" {
            case "\n\t":
                v1 = (^(v1))
            case "a":
                if (true || (true && (570 < 65535))) {
                    for fi65 := 0; fi65 < 4; fi65++ {
                        if true {
                            var v66 uint8 = 42
                            v4["hello"] = v3
                            switch 127 {
                            case 127:
                                var v67 float32 = (float32((float64(-2.25) - ((3.14159 - (-2.25 + 1e10)) - float64(0.0)))) / float32(-2.25))
                                _ = v67
                            case 711:
                                for fi68 := 0; fi68 < 8; fi68++ {
                                    v12 -= 1
                                    _ = fi68
                                }
                            case -439:
                                v5 = append(v5, 4294967295)
                            }
                            _ = v66
                        }
                        for fi69 := 0; fi69 < 6; fi69++ {
                            v12 = (^(v12))
                            _ = fi69
                        }
                        for fi70 := 0; fi70 < 1; fi70++ {
                            var v71 uint8 = 127
                            v2 += (v2 | (^(v2)))
                            v72 := GenS10{F0: 0, F1: 2, F2: 127}
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        _ = fi65
                    }
                    for fi73 := 0; fi73 < 6; fi73++ {
                        var v74 int32 = -1
                        for fi75 := 0; fi75 < 10; fi75++ {
                            var v76 bool = true
                            v77 := GenS4{F0: 42, F1: 42, F2: v1}
                            v76 = v76
                            var v78 uint8 = 128
                            v5 = append(v5, v3)
                            _ = fi75
                            _ = v76
                            _ = v77
                            _ = v78
                        }
                        _ = fi73
                        _ = v74
                    }
                }
            case "":
                v12 = v12
            default:
                var v79 uint16 = (^(v12))
                _ = v79
            }
        case "a":
            for fi80 := 0; fi80 < 19; fi80++ {
                v5 = append(v5, v3)
                v2 = 639
                fmt.Println(fi80)
                _ = fi80
            }
        default:
            fmt.Println(v2, v12, v12)
        }
    }
    for fi81 := 0; fi81 < 5; fi81++ {
        if ((v1 % (((v1 << ((255) & 63))) | 1)) == 1) {
            var v82 float32 = 1e10
            v4[("a" + "a")] = v3
            _ = v82
        }
        switch fi81 {
        case -1:
            v1 -= (v1 / ((v1) | 1))
        case 4294967295:
            switch ((("hello" + "\n\t") + ("a" + "a")) + "hello") {
            case "\n\t":
                fmt.Println(fi81, v3)
            case "hello":
                var v83 uint = 4294967295
                _ = v83
            default:
                _ = len(v4)
            }
        default:
            v84 := v4["\n\t"]
            _ = v84
        }
        _ = fi81
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        switch int(v1) {
        case 42:
            switch (fi85 ^ int(v1)) {
            case -15:
                var v86 int64 = -9223372036854775808
                _ = v86
            case 128:
                for fi87 := 0; fi87 < 13; fi87++ {
                    v12 -= uint16(v3)
                    v4[""] = (v3 << ((uint(v3)) & 63))
                    if false {
                        if (!(!true)) {
                            var v88 int = (fi85 - 2147483647)
                            if ((1 >= 237) || false) {
                                switch ("" + ("a" + ("\n\t" + ("\n\t" + "hello")))) {
                                case "":
                                    var v89 int16 = 1000
                                    _ = v89
                                case "foo bar":
                                    v88 += fi87
                                }
                                fmt.Println(fi85, v88, v2)
                            } else {
                                fmt.Println(fi85, v1, fi87)
                            }
                            var v90 bool = ((3.14159 <= float64((-2.25 + 0.1))) && (42 == 4294967295))
                            v88 = 128
                            _ = v88
                            _ = v90
                        }
                    } else {
                        if true {
                            v91 := GenS11{F0: false, F1: -9223372036854775808}
                            _ = v91
                        } else {
                            v12 = 255
                        }
                        if (-128 <= -1) {
                            v2 = v2
                            v92 := v4["a"]
                            var v93 uint64 = 0
                            _ = v92
                            _ = v93
                        }
                    }
                    _ = fi87
                }
            case 594:
                var v94 uint = 732
                _ = v94
            default:
                if ((42 > 127) || (fi85 > fi85)) {
                    var v95 bool = (2 == 0)
                    if v95 {
                        fmt.Println(v95, v95, v3)
                    } else {
                        if (!false) {
                            v2 -= (v2 + -2147483648)
                            var v96 uint64 = 762
                            _ = v96
                        }
                    }
                    _ = v95
                }
            }
        case 4294967295:
            v97 := GenS12{F0: 0}
            _ = v97
        default:
            _ = len(v5)
        }
        if (1.5 > (3.14159 * (1.5 / 1.0))) {
            fmt.Println(v12, v2, v3)
            v98 := GenS10{F0: 223, F1: (v3 % ((v3) | 1)), F2: 0}
            for fi99 := 0; fi99 < 16; fi99++ {
                var v100 uint8 = 1
                v101 := GenS13{F0: 4294967295, F1: 2147483647}
                var v102 int16 = v1
                v103 := v4["foo bar"]
                _ = fi99
                _ = v100
                _ = v101
                _ = v102
                _ = v103
            }
            _ = v98
        } else {
            fmt.Println(v12, fi85)
            var v104 bool = true
            var v105 int16 = v1
            _ = v104
            _ = v105
        }
        _ = fi85
    }
    switch "foo bar" {
    case "\n\t":
        v12 -= v12
    case "":
        v106 := GenS5{F0: 42}
        _ = v106
    }
    v12 = (v12 | 887)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
