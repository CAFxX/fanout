package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 1.5
    v3 := []string{"\n\t", ("" + "\n\t"), "hello"}
    v4 := make(map[uint32]int)
    v2 = 1e10
    if len(v3) > 0 {
        v3[0] = (("" + "a") + "\n\t")
    }
    var v5 float64 = -2.25
    switch 309 {
    case 543:
        v6 := GenS1{F0: "\n\t"}
        _ = v6
    }
    var v7 int8 = 1
    fmt.Println(v7)
    v7 += (v7 >> ((255) & 63))
    var v8 uint16 = 65535
    if (false || true) {
        switch "hello" {
        case "hello":
            v2 = v2
        case "a":
            var v9 float64 = v5
            _ = v9
        case "\n\t":
            var v10 float32 = ((((v2 * (v2 + ((v2 - v2) / (v2 / v2)))) * v2) - 1.5) + v2)
            _ = v10
        }
        fmt.Println(v8, v5)
        var v11 float32 = 1e10
        _ = v11
    } else {
        _ = len(v3)
        if (492 != 2147483647) {
            var v12 bool = false
            _ = v12
        } else {
            for fi13 := 0; fi13 < 15; fi13++ {
                if false {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 int = fi14
                        _ = fi14
                        _ = v15
                    }
                    if true {
                        switch ("hello" + (((("foo bar" + "a") + ("" + ("hello" + "hello"))) + "hello") + "")) {
                        case "foo bar":
                            for fi16 := 0; fi16 < 8; fi16++ {
                                switch (-(fi16)) {
                                case -9223372036854775808:
                                    if true {
                                        fmt.Println(v5)
                                        if ("" != "") {
                                            v5 = 1.5
                                            v1[0] = (fi13 - fi13)
                                        }
                                        v8 = (^(v8))
                                    } else {
                                        _ = len(v3)
                                        v7 = (v7 >> ((0) & 63))
                                    }
                                    if true {
                                        if (0 <= (v8 - v8)) {
                                            var v17 float64 = (((v5 + ((0.0 - ((v5 * v5) * v5)) * ((v5 - (0.1 + (v5 * 0.1))) - 0.0))) / v5) * 3.14159)
                                            var v18 int8 = (v7 >> ((255) & 63))
                                            _ = v17
                                            _ = v18
                                        } else {
                                            switch (("" + ("foo bar" + (("" + "a") + "a"))) + "\n\t") {
                                            case "a":
                                                v5 += (3.14159 - (v5 * v5))
                                            default:
                                                switch (fi16 ^ fi16) {
                                                case -1:
                                                    if (42 <= -163) {
                                                        v19 := GenS2{F0: false}
                                                        v7 -= v7
                                                        v7 += 100
                                                        _ = v19
                                                    } else {
                                                        var v20 string = "hello"
                                                        if false {
                                                            var v21 int64 = -9223372036854775808
                                                            fmt.Println(v5, fi16)
                                                            v7 = v7
                                                            _ = v21
                                                        } else {
                                                            var v22 uint16 = v8
                                                            var v23 uint64 = 1
                                                            _ = v22
                                                            _ = v23
                                                        }
                                                        v2 = 0.1
                                                        _ = v20
                                                    }
                                                case 32767:
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                case -9223372036854775808:
                                                    if len(v3) > 0 {
                                                        v3[0] = ""
                                                    }
                                                default:
                                                    v24 := GenS2{F0: false}
                                                    _ = v24
                                                }
                                            }
                                            fmt.Println(fi16)
                                        }
                                        v2 = (v2 * 0.1)
                                    } else {
                                        var v25 int8 = v7
                                        _ = v25
                                    }
                                    var v26 string = ""
                                    _ = v26
                                case 32768:
                                    if (55 < 32) {
                                        var v27 float32 = v2
                                        _ = v27
                                    }
                                    var v28 int = 1
                                    v29 := GenS1{F0: "hello"}
                                    _ = v28
                                    _ = v29
                                default:
                                    var v30 bool = (true && (!false))
                                    _ = v30
                                }
                                if true {
                                    var v31 uint64 = 42
                                    v7 += v7
                                    var v32 int64 = 0
                                    v31 += 2
                                    _ = v31
                                    _ = v32
                                }
                                _ = fi16
                            }
                        case "a":
                            var v33 uint = 2
                            _ = v33
                        }
                        if (v8 >= v8) {
                            switch "foo bar" {
                            case "\n\t":
                                v3 = append(v3, ("" + "hello"))
                            case "foo bar":
                                switch (fi13 * fi13) {
                                case 1:
                                    for fi34 := 0; fi34 < 4; fi34++ {
                                        var v35 string = "foo bar"
                                        delete(v1, 42)
                                        delete(v1, 0)
                                        var v36 float32 = v2
                                        var v37 uint32 = 65535
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                }
                            }
                            v8 = v8
                        } else {
                            for fi38 := 0; fi38 < 11; fi38++ {
                                var v39 uint32 = 42
                                if true {
                                    if (-88 <= v7) {
                                        var v40 uint32 = v39
                                        var v41 int64 = 1
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if false {
                                            v42 := GenS3{F0: 0}
                                            fmt.Println(v39)
                                            _ = v42
                                        } else {
                                            if ((!false) || true) {
                                                fmt.Println(fi38)
                                                v39 -= v39
                                                var v43 uint8 = 0
                                                _ = v43
                                            } else {
                                                v3 = append(v3, "")
                                            }
                                            var v44 uint = 533
                                            switch (^(fi38)) {
                                            case -2147483648:
                                                _ = len(v3)
                                            case -9223372036854775808:
                                                fmt.Println(v2, fi38, fi38)
                                            }
                                            var v45 bool = (!(!true))
                                            switch "" {
                                            case "a":
                                                fmt.Println(v2, fi38, fi13)
                                            case "foo bar":
                                                _ = len(v3)
                                            }
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                    _ = len(v3)
                                    v39 = v39
                                } else {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        if true {
                                            fmt.Println(v8)
                                            v8 += (v8 ^ (v8 & v8))
                                            v7 += v7
                                            var v46 float64 = 1e10
                                            var v47 uint32 = (^(v39))
                                            _ = v46
                                            _ = v47
                                        } else {
                                            var v48 int16 = 0
                                            if (v7 == int8(v8)) {
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                            }
                                            _ = v48
                                        }
                                    case "\n\t":
                                        switch (-(fi38)) {
                                        case 127:
                                            v7 -= (v7 << ((18446744073709551615) & 63))
                                        case -2147483648:
                                            v49 := GenS3{F0: 4294967295}
                                            _ = v49
                                        default:
                                            var v50 uint8 = 19
                                            _ = v50
                                        }
                                    default:
                                        var v51 float64 = ((v5 - -2.25) - float64((float32(v5) * v2)))
                                        _ = v51
                                    }
                                    v52 := GenS2{F0: false}
                                    _ = v52
                                }
                                v7 = (^(v7))
                                switch "" {
                                case "a":
                                    v53 := GenS4{F0: v7, F1: 2, F2: 128}
                                    _ = v53
                                case "foo bar":
                                    _ = len(v4)
                                case "":
                                    v39 += v39
                                default:
                                    v54 := GenS3{F0: uint(v39)}
                                    _ = v54
                                }
                                v7 -= v7
                                v5 -= v5
                                _ = fi38
                                _ = v39
                            }
                            v55 := GenS3{F0: 166}
                            var v56 int16 = 995
                            _ = v55
                            _ = v56
                        }
                    } else {
                        var v57 int64 = 42
                        switch "foo bar" {
                        case "hello":
                            switch "foo bar" {
                            case "\n\t":
                                v2 += (v2 * ((v2 / (v2 / (v2 - ((v2 + (1.5 + v2)) / float32(0.1))))) - v2))
                            case "hello":
                                for fi58 := 0; fi58 < 9; fi58++ {
                                    var v59 int32 = 0
                                    var v60 float32 = 0.1
                                    switch fi58 {
                                    case 128:
                                        if (false || false) {
                                            switch "a" {
                                            case "":
                                                v60 = v60
                                            case "a":
                                                var v61 int8 = v7
                                                _ = v61
                                            }
                                        }
                                    default:
                                        v62 := v4[65535]
                                        _ = v62
                                    }
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                }
                            case "a":
                                v7 = v7
                            default:
                                var v63 uint8 = 127
                                _ = v63
                            }
                        }
                        fmt.Println(v7, v8, v57)
                        _ = v57
                    }
                    switch ("" + (("" + "\n\t") + "hello")) {
                    case "hello":
                        for fi64 := 0; fi64 < 17; fi64++ {
                            var v65 int = fi64
                            if false {
                                delete(v4, 4294967295)
                                var v66 float64 = float64((v2 - ((1e10 + v2) * 3.14159)))
                                _ = v66
                            }
                            _ = fi64
                            _ = v65
                        }
                    case "a":
                        var v67 int32 = 1
                        _ = v67
                    }
                } else {
                    for fi68 := 0; fi68 < 19; fi68++ {
                        v7 = v7
                        v69 := GenS4{F0: v7, F1: 255, F2: 42}
                        _ = fi68
                        _ = v69
                    }
                }
                switch (fi13 + fi13) {
                case -732:
                    switch int(v8) {
                    case 9223372036854775807:
                        for fi70 := 0; fi70 < 18; fi70++ {
                            if (false && (true || true)) {
                                fmt.Println(v5, v2)
                                var v71 uint64 = 0
                                _ = v71
                            } else {
                                switch -923 {
                                case -798:
                                    v5 -= (1e10 - ((v5 / 1.5) * -2.25))
                                case 9223372036854775807:
                                    var v72 int32 = 1
                                    _ = v72
                                }
                                var v73 int32 = -2147483648
                                _ = v73
                            }
                            v74 := GenS1{F0: "hello"}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v5 = (v5 - (1.5 * (((v5 * float64(float32(v5))) - float64(1e10)) / -2.25)))
                            _ = fi70
                            _ = v74
                        }
                    case -900:
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            for fi75 := 0; fi75 < 4; fi75++ {
                                _ = len(v3)
                                _ = fi75
                            }
                        default:
                            fmt.Println(v2, fi13, v2)
                        }
                    case 32767:
                        var v76 string = "foo bar"
                        _ = v76
                    default:
                        var v77 float32 = 3.14159
                        _ = v77
                    }
                case 128:
                    v8 -= 416
                default:
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi13, v8, v2)
                    case "foo bar":
                        for fi78 := 0; fi78 < 8; fi78++ {
                            _ = len(v1)
                            v79 := v4[2147483648]
                            v80 := GenS3{F0: 0}
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    default:
                        var v81 int64 = 378
                        _ = v81
                    }
                }
                _ = fi13
            }
            for fi82 := 0; fi82 < 19; fi82++ {
                var v83 float64 = v5
                _ = fi82
                _ = v83
            }
            v2 += 0.1
        }
    }
    fmt.Println(v7)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        for fi84 := 0; fi84 < 4; fi84++ {
            _ = len(v4)
            for fi85 := 0; fi85 < 9; fi85++ {
                v2 -= 3.14159
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                var v86 uint32 = uint32(v8)
                _ = fi85
                _ = v86
            }
            _ = fi84
        }
        v7 = 87
        v2 = v2
    }
    var v87 int32 = 863
    if (true && (true || false)) {
        v7 = (v7 - -1)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v87
}
