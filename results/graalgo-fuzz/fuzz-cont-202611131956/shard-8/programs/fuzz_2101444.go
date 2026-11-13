package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 string; F2 uint64 }
type GenS6 struct { F0 int16; F1 int64; F2 bool }
type GenS7 struct { F0 uint8 }
func genFn0() uint {
    v4 := []int{32767, 127}
    for fi5 := 0; fi5 < 6; fi5++ {
        v4 = append(v4, fi5)
        fmt.Println(fi5)
        v6 := GenS1{F0: 2147483647, F1: -32768}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint64 = 1
        _ = fi7
        _ = v8
    }
    if (0 == 1000) {
        if ((127 < -128) && true) {
            var v9 int = 4294967295
            var v10 int16 = -929
            _ = v10
            _ = v9
        }
    }
    _ = v4
    return 2
}
func genFn1() int32 {
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 uint16 = 1
            switch "hello" {
            case "foo bar":
                switch 720 {
                case 746:
                    v14 := []int{(fi11 ^ fi12), 127}
                    _ = v14
                case -9223372036854775808:
                    if (float32(3.14159) <= 1.5) {
                        switch "foo bar" {
                        case "foo bar":
                            if (65535 == 42) {
                                if false {
                                    var v15 uint64 = 1
                                    var v16 bool = true
                                    _ = v15
                                    _ = v16
                                } else {
                                    fmt.Println(fi12)
                                    var v17 int16 = 91
                                    _ = v17
                                }
                                v18 := GenS1{F0: 2147483648, F1: -93}
                                _ = v18
                            }
                        case "a":
                            v13 = (v13 * v13)
                        case "hello":
                            v19 := GenS1{F0: 1, F1: -211}
                            _ = v19
                        default:
                            var v20 uint32 = 2147483647
                            _ = v20
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            if (((("a" + ("" + ("hello" + "hello"))) + "") + "") != ("\n\t" + ("" + "a"))) {
                                v13 = v13
                                v13 = (v13 & v13)
                                var v21 uint8 = 112
                                fmt.Println(v13, fi11)
                                if false {
                                    var v22 uint8 = v21
                                    var v23 int32 = 0
                                    _ = v22
                                    _ = v23
                                } else {
                                    v13 = v13
                                }
                                fmt.Println(v21, v13)
                                var v24 float32 = (0.1 + -2.25)
                                _ = v21
                                _ = v24
                            } else {
                                v25 := []int{fi12}
                                switch 606 {
                                case 32767:
                                    var v26 uint16 = v13
                                    _ = v26
                                case 226:
                                    if (false && (!(false || (!true)))) {
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                    }
                                case 127:
                                    switch "hello" {
                                    case "hello":
                                        v13 += (v13 & (^(v13)))
                                    case "a":
                                        if (9223372036854775807 >= 610) {
                                            if true {
                                                var v27 string = ("" + ((("a" + ("" + "hello")) + "\n\t") + ("a" + ((("hello" + "a") + "") + (("hello" + "") + "hello")))))
                                                v13 = v13
                                                if false {
                                                    v28 := GenS1{F0: 42, F1: 42}
                                                    var v29 uint16 = v13
                                                    v29 = v29
                                                    fmt.Println(v27, v27)
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    v30 := GenS2{F0: 731}
                                                    switch fi12 {
                                                    case 47:
                                                        fmt.Println(fi12, fi11, v13)
                                                    case 42:
                                                        if (!true) {
                                                            v27 = v27
                                                            if len(v25) > 0 {
                                                                v25[0] = -2147483648
                                                            }
                                                            if (0.0 < 3.14159) {
                                                                v27 = ((v27 + "a") + v27)
                                                                fmt.Println(fi11, v13)
                                                                switch "" {
                                                                case "\n\t":
                                                                    switch fi11 {
                                                                    case 0:
                                                                        v27 = v27
                                                                    case -2147483648:
                                                                        if (((2147483647 == 65) && (!true)) && ((^(v13)) == (v13 << ((42) & 63)))) {
                                                                            v31 := GenS1{F0: 256, F1: 1000}
                                                                            var v32 float32 = 1e10
                                                                            v33 := GenS2{F0: 2147483647}
                                                                            v27 = ""
                                                                            _ = v31
                                                                            _ = v32
                                                                            _ = v33
                                                                        }
                                                                    case -672:
                                                                        _ = len(v25)
                                                                        switch (fi11 - fi12) {
                                                                        case -2147483648:
                                                                            if (1 <= 42) {
                                                                                _ = len(v25)
                                                                                v27 = v27
                                                                            } else {
                                                                                v27 = "hello"
                                                                            }
                                                                        case 2147483647:
                                                                            var v34 uint = 4294967295
                                                                            v35 := GenS1{F0: 2147483648, F1: 1}
                                                                            _ = v34
                                                                            _ = v35
                                                                        case -1:
                                                                            var v36 float64 = 0.1
                                                                            fmt.Println(v13, v27)
                                                                            _ = v36
                                                                        default:
                                                                            if (!(0.0 == 3.14159)) {
                                                                                if (-2147483648 < 1) {
                                                                                    v37 := GenS3{F0: -66, F1: 302}
                                                                                    v25 = append(v25, fi12)
                                                                                    _ = v37
                                                                                } else {
                                                                                    fmt.Println(fi12)
                                                                                    var v38 uint8 = 1
                                                                                    _ = v38
                                                                                }
                                                                                switch 32768 {
                                                                                case -808:
                                                                                    var v39 int64 = -9223372036854775808
                                                                                    _ = v39
                                                                                case -398:
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v40 uint64 = 686
                                                                                        _ = v40
                                                                                    }
                                                                                    var v41 float32 = (3.14159 / 0.1)
                                                                                    var v42 uint64 = 1
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                case 32767:
                                                                                    var v43 int8 = -74
                                                                                    _ = v43
                                                                                default:
                                                                                    fmt.Println(fi11)
                                                                                }
                                                                                fmt.Println(v13)
                                                                                v44 := GenS2{F0: 4294967295}
                                                                                v13 = (^(v13))
                                                                                v45 := GenS4{F0: 480}
                                                                                _ = v44
                                                                                _ = v45
                                                                            } else {
                                                                                v27 = "foo bar"
                                                                                var v46 uint64 = 592
                                                                                _ = v46
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    v13 = 630
                                                                    var v47 int32 = -195
                                                                    if ((261 == (^(fi12))) || (!true)) {
                                                                        v13 -= v13
                                                                        var v48 string = "a"
                                                                        switch int(v47) {
                                                                        case 2147483647:
                                                                            _ = len(v25)
                                                                        case -803:
                                                                            fmt.Println(v47, fi11, v13)
                                                                        case -1:
                                                                            if len(v25) > 0 {
                                                                                v25[0] = (fi12 >> ((65535) & 63))
                                                                            }
                                                                        default:
                                                                            var v49 int16 = -32768
                                                                            fmt.Println(v48)
                                                                            var v50 int = -9223372036854775808
                                                                            _ = v49
                                                                            _ = v50
                                                                        }
                                                                        _ = v48
                                                                    } else {
                                                                        v27 = v27
                                                                        switch "a" {
                                                                        case "hello":
                                                                            v13 -= (v13 ^ v13)
                                                                        case "\n\t":
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                        case "foo bar":
                                                                            if false {
                                                                                v51 := GenS3{F0: -128, F1: -1}
                                                                                fmt.Println(v27)
                                                                                v25 = append(v25, 42)
                                                                                var v52 uint16 = v13
                                                                                v52 = 0
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                        }
                                                                        var v53 float32 = 1.5
                                                                        _ = v53
                                                                    }
                                                                    _ = v47
                                                                }
                                                            }
                                                            var v54 uint32 = 0
                                                            var v55 string = v27
                                                            if true {
                                                                switch fi11 {
                                                                case 204:
                                                                    v56 := GenS4{F0: 42}
                                                                    _ = v56
                                                                default:
                                                                    if true {
                                                                        switch -1 {
                                                                        case 32768:
                                                                            v55 = "a"
                                                                        }
                                                                        var v57 float64 = -2.25
                                                                        _ = v57
                                                                    }
                                                                }
                                                                if false {
                                                                    v25 = append(v25, 1)
                                                                    v54 += v54
                                                                    v27 = "\n\t"
                                                                }
                                                                switch v55 {
                                                                case "foo bar":
                                                                    switch ((v27 + (v27 + v55)) + (v55 + v27)) {
                                                                    case "\n\t":
                                                                        v58 := GenS2{F0: 0}
                                                                        _ = v58
                                                                    case "":
                                                                        v13 = v13
                                                                    }
                                                                default:
                                                                    switch v55 {
                                                                    case "a":
                                                                        var v59 uint64 = 18446744073709551615
                                                                        _ = v59
                                                                    case "\n\t":
                                                                        switch (fi12 & fi12) {
                                                                        case -101:
                                                                            if true {
                                                                                fmt.Println(fi11, v54, v54)
                                                                                if len(v25) > 0 {
                                                                                    v25[0] = (fi12 / (((fi12 - fi12)) | 1))
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v60 float64 = (1e10 + 1e10)
                                                                    v54 = (v54 | v54)
                                                                    _ = v60
                                                                }
                                                                v55 = v55
                                                                fmt.Println(fi12, fi11, v13)
                                                            }
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    default:
                                                        var v61 int64 = 1
                                                        _ = v61
                                                    }
                                                    _ = v30
                                                }
                                                _ = v27
                                            }
                                            v13 = (v13 - v13)
                                        } else {
                                            v62 := GenS4{F0: 0}
                                            _ = v62
                                        }
                                    }
                                }
                                _ = v25
                            }
                        case "hello":
                            v13 += v13
                        }
                        fmt.Println(fi12)
                        v63 := []int{fi12}
                        v13 = v13
                        _ = v63
                    }
                case 0:
                    var v64 int32 = -1
                    _ = v64
                }
            case "a":
                v13 += 42
            case "hello":
                switch ("foo bar" + "a") {
                case "hello":
                    switch 32767 {
                    case -2147483648:
                        v65 := GenS5{F0: 42, F1: "a", F2: 1}
                        _ = v65
                    case 4294967295:
                        var v66 int = (^(fi12))
                        _ = v66
                    }
                case "a":
                    var v67 int64 = 1
                    _ = v67
                case "foo bar":
                    var v68 string = "\n\t"
                    _ = v68
                }
            default:
                v13 -= (^(v13))
            }
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v69 uint8 = 127
    v69 = (v69 / ((v69) | 1))
    v69 -= v69
    _ = v69
    return 2147483647
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ("hello" + ("foo bar" + (("a" + ((("hello" + "") + "hello") + ("a" + "hello"))) + (("hello" + ("hello" + "\n\t")) + (("a" + "") + "")))))
    v3 := make(map[uint32]uint32)
    var v70 float32 = 3.14159
    if true {
        var v71 uint8 = 128
        _ = v71
    }
    var v72 int64 = 42
    v72 -= v72
    switch (v2 + v2) {
    case "":
        var v73 uint32 = 4294967295
        _ = v73
    case "foo bar":
        var v74 uint8 = 255
        _ = v74
    case "a":
        fmt.Println(v70, v72, v72)
    }
    if false {
        if (-497 == 99) {
            v75 := GenS4{F0: 6}
            _ = v75
        } else {
            var v76 int = -2147483648
            if true {
                var v77 string = ((("\n\t" + v2) + ("foo bar" + "hello")) + v2)
                _ = v77
            }
            for fi78 := 0; fi78 < 1; fi78++ {
                for fi79 := 0; fi79 < 11; fi79++ {
                    if false {
                        v72 = v72
                        v70 = 1.5
                        fmt.Println(v76, fi78, v76)
                    } else {
                        fmt.Println(v2, fi79, v70)
                    }
                    _ = fi79
                }
                _ = fi78
            }
            _ = v76
        }
        v2 = v2
        var v80 string = (("\n\t" + v2) + v2)
        _ = v80
    }
    var v81 float64 = (0.1 / 1e10)
    v72 = (v72 ^ v72)
    v82 := GenS4{F0: 728}
    v83 := GenS6{F0: 355, F1: 585, F2: false}
    switch "foo bar" {
    case "hello":
        v84 := GenS7{F0: 0}
        _ = v84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v70
    _ = v72
    _ = v81
    _ = v82
    _ = v83
}
