package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 string; F2 uint }
type GenS8 struct { F0 int8; F1 uint; F2 uint32 }
type GenS9 struct { F0 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -1
    var v3 int8 = v2
    v4 := GenS1{F0: 32767}
    switch -2147483648 {
    case -530:
        var v5 uint8 = 255
        _ = v5
    }
    var v6 float32 = float32(0.0)
    for fi7 := 0; fi7 < 9; fi7++ {
        v2 += (v3 << ((65535) & 63))
        for fi8 := 0; fi8 < 16; fi8++ {
            v1[fi7] = fi7
            v9 := []int{32768}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            if (!(float64(3.14159) < 3.14159)) {
                _ = len(v1)
                switch fi10 {
                case 2147483647:
                    switch -2147483648 {
                    case -126:
                        if (!(369 <= 1)) {
                            switch ("foo bar" + (("a" + ("\n\t" + "\n\t")) + "a")) {
                            case "a":
                                if true {
                                    fmt.Println(v3)
                                    var v11 int64 = 0
                                    _ = v11
                                }
                            }
                            switch (fi7 / (((fi10 >> ((2) & 63))) | 1)) {
                            case 128:
                                v6 -= 1.5
                            case 4294967295:
                                fmt.Println(v6, v2)
                            case 714:
                                v12 := GenS1{F0: 32767}
                                _ = v12
                            default:
                                if (!(2147483647 == 572)) {
                                    v2 -= 1
                                    v3 += v2
                                    _ = len(v1)
                                    switch (fi10 << ((4294967295) & 63)) {
                                    case -710:
                                        var v13 uint16 = 364
                                        _ = v13
                                    }
                                    fmt.Println(v3, fi7)
                                }
                            }
                        } else {
                            if false {
                                fmt.Println(v6, v3, v2)
                                v2 = -128
                                v3 = (v2 << ((507) & 63))
                            } else {
                                v14 := GenS2{F0: 32767, F1: false}
                                switch ("" + "a") {
                                case "\n\t":
                                    var v15 int64 = 0
                                    _ = v15
                                }
                                var v16 int32 = -2147483648
                                v16 = v16
                                var v17 int64 = 0
                                _ = v14
                                _ = v16
                                _ = v17
                            }
                            v18 := GenS2{F0: 612, F1: (740 != 65535)}
                            if (-2147483648 != -2147483648) {
                                fmt.Println(v6, fi7, v6)
                            }
                            _ = v18
                        }
                    case 4294967295:
                        var v19 uint32 = 2
                        _ = v19
                    case -9223372036854775808:
                        switch "\n\t" {
                        case "a":
                            v3 -= (v2 >> ((65535) & 63))
                        case "\n\t":
                            var v20 float32 = float32(-2.25)
                            _ = v20
                        case "foo bar":
                            var v21 int8 = v2
                            _ = v21
                        default:
                            if true {
                                if true {
                                    fmt.Println(v6, v6)
                                } else {
                                    v3 -= v2
                                }
                            } else {
                                var v22 float64 = 0.1
                                _ = v22
                            }
                        }
                    default:
                        v23 := []int{(fi7 << ((894) & 63)), (fi7 << ((2) & 63)), 291}
                        _ = v23
                    }
                case -9223372036854775808:
                    if false {
                        if true {
                            var v24 int = -175
                            _ = v24
                        } else {
                            if (1.5 <= float64(0.0)) {
                                fmt.Println(v3, fi7, v3)
                                switch fi10 {
                                case 2147483647:
                                    switch ("hello" + ((("foo bar" + "\n\t") + ((("foo bar" + ((("\n\t" + ("foo bar" + "foo bar")) + (("\n\t" + "") + ("" + "\n\t"))) + "hello")) + "") + "hello")) + "foo bar")) {
                                    case "foo bar":
                                        v25 := GenS2{F0: 2, F1: true}
                                        _ = v25
                                    default:
                                        v26 := []int{fi7, fi10, fi7}
                                        _ = v26
                                    }
                                case -2147483648:
                                    switch (fi10 % ((2147483647) | 1)) {
                                    case 127:
                                        v2 = 1
                                    case -2147483648:
                                        v2 -= 127
                                    case -669:
                                        var v27 int32 = 1
                                        _ = v27
                                    default:
                                        if true {
                                            if (true && true) {
                                                var v28 int64 = 9223372036854775807
                                                _ = v28
                                            }
                                            v6 -= v6
                                            v29 := []int{(fi10 / ((fi7) | 1)), fi7, 127, (^(fi7))}
                                            _ = v29
                                        } else {
                                            if (((0 >= 0) && false) && (42 != 767)) {
                                                v2 = (v3 ^ (v3 & v3))
                                                var v30 bool = true
                                                v31 := GenS3{F0: 1, F1: fi10}
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v6 = 1.5
                                            }
                                            if (true && false) {
                                                v32 := GenS4{F0: fi10, F1: (false || false)}
                                                v33 := GenS5{F0: (("hello" + (("foo bar" + "a") + "")) + ("a" + "hello"))}
                                                _ = v32
                                                _ = v33
                                            } else {
                                                if (!(-93 >= -559)) {
                                                    var v34 uint8 = 0
                                                    delete(v1, (fi10 | fi7))
                                                    _ = v34
                                                }
                                            }
                                        }
                                    }
                                case 4294967295:
                                    v6 -= float32(float64((v6 / v6)))
                                }
                                fmt.Println(v3, fi7)
                            }
                        }
                        delete(v1, (fi10 & (fi10 & fi10)))
                    }
                }
            } else {
                fmt.Println(fi10, v3)
                var v35 uint32 = 0
                if (0 >= 2) {
                    v36 := []int{(fi10 << ((2) & 63))}
                    v6 -= v6
                    var v37 float64 = -2.25
                    _ = v36
                    _ = v37
                } else {
                    var v38 uint64 = 18446744073709551615
                    if (true && true) {
                        switch "" {
                        case "\n\t":
                            v3 = v3
                            switch fi10 {
                            case -9223372036854775808:
                                v35 -= (^(v35))
                            case -951:
                                if (!true) {
                                    v38 += 478
                                    v39 := GenS5{F0: ("a" + "hello")}
                                    fmt.Println(v2, v35, fi7)
                                    _ = v39
                                }
                            case 0:
                                var v40 uint = 42
                                _ = v40
                            default:
                                switch "a" {
                                case "a":
                                    v3 -= (v2 & (v2 / ((-99) | 1)))
                                case "":
                                    var v41 int = (-(fi7))
                                    _ = v41
                                case "hello":
                                    var v42 float32 = 1e10
                                    _ = v42
                                default:
                                    fmt.Println(v2, v6, v3)
                                }
                                switch "" {
                                case "":
                                    v43 := []int{int(v35), (fi7 / ((fi10) | 1)), (fi10 << ((975) & 63)), int(v3)}
                                    _ = v43
                                case "foo bar":
                                    var v44 uint32 = (v35 + (v35 % ((v35) | 1)))
                                    _ = v44
                                case "hello":
                                    switch (fi7 * (fi7 + fi7)) {
                                    case -851:
                                        var v45 string = "a"
                                        _ = v45
                                    }
                                default:
                                    switch ("foo bar" + "\n\t") {
                                    case "a":
                                        v46 := GenS6{F0: 255}
                                        _ = v46
                                    default:
                                        var v47 bool = true
                                        switch fi7 {
                                        case -1:
                                            fmt.Println(v38, v35)
                                        case -986:
                                            if (v2 > (v2 >> ((827) & 63))) {
                                                var v48 uint64 = (v38 - (v38 << ((973) & 63)))
                                                v6 += v6
                                                v38 = v38
                                                _ = v48
                                            } else {
                                                v49 := GenS2{F0: 32767, F1: ((1e10 >= 0.1) || v47)}
                                                v38 -= v38
                                                v50 := GenS7{F0: 65535, F1: "hello", F2: 42}
                                                _ = v49
                                                _ = v50
                                            }
                                        default:
                                            var v51 uint32 = v35
                                            _ = v51
                                        }
                                        var v52 uint32 = v35
                                        _ = v47
                                        _ = v52
                                    }
                                }
                                v53 := GenS4{F0: 1, F1: true}
                                _ = v53
                            }
                        case "hello":
                            switch (fi10 * (-(fi7))) {
                            case 9223372036854775807:
                                v54 := GenS7{F0: 42, F1: "", F2: 679}
                                _ = v54
                            case 0:
                                v38 -= (v38 >> ((4294967295) & 63))
                            case -1:
                                v6 = v6
                            default:
                                fmt.Println(v38, fi7, v6)
                            }
                        case "a":
                            if ((597 != 2) && false) {
                                var v55 uint8 = 128
                                _ = v55
                            }
                        }
                    } else {
                        v56 := GenS8{F0: (v3 - -1), F1: 853, F2: (v35 & v35)}
                        fmt.Println(fi10)
                        fmt.Println(v35)
                        _ = v56
                    }
                    _ = v38
                }
                _ = v35
            }
            if false {
                if true {
                    var v57 uint8 = 127
                    _ = v57
                }
            } else {
                var v58 int64 = -9223372036854775808
                v59 := []int{fi7}
                var v60 int = fi7
                _ = v58
                _ = v59
                _ = v60
            }
            v3 += (v3 >> ((0) & 63))
            _ = fi10
        }
        _ = fi7
    }
    switch "hello" {
    case "foo bar":
        if true {
            var v61 uint32 = 223
            delete(v1, 32768)
            v62 := GenS9{F0: 92}
            _ = v61
            _ = v62
        } else {
            var v63 uint32 = 0
            _ = v63
        }
    case "a":
        if true {
            v64 := GenS6{F0: 0}
            var v65 int64 = -1
            _ = v64
            _ = v65
        } else {
            for fi66 := 0; fi66 < 10; fi66++ {
                var v67 uint32 = 461
                _ = fi66
                _ = v67
            }
            for fi68 := 0; fi68 < 16; fi68++ {
                var v69 float32 = ((float32(3.14159) + v6) / ((1.5 - v6) * 3.14159))
                fmt.Println(fi68)
                v3 = (-(v3))
                _ = fi68
                _ = v69
            }
        }
    }
    if ((v2 % ((v3) | 1)) >= v2) {
        fmt.Println(v3, v6)
    }
    v2 = v3
    v70 := []int{}
    v71 := GenS7{F0: 255, F1: (("" + "a") + ("a" + "\n\t")), F2: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v70:", len(v70), v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v70
    _ = v71
}
