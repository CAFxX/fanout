package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint; F1 int; F2 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 uint }
func genFn0(p0 int8) uint {
    for fi6 := 0; fi6 < 2; fi6++ {
        switch fi6 {
        case -1:
            switch "" {
            case "foo bar":
                if (true && true) {
                    v7 := GenS1{F0: -128, F1: 2}
                    _ = v7
                } else {
                    var v8 int8 = p0
                    for fi9 := 0; fi9 < 3; fi9++ {
                        switch 128 {
                        case 42:
                            if (fi9 <= (fi9 + fi9)) {
                                switch "a" {
                                case "a":
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v10 uint32 = 2
                                        _ = v10
                                    case "foo bar":
                                        switch fi9 {
                                        case 329:
                                            if true {
                                                switch ("hello" + "") {
                                                case "hello":
                                                    v11 := []int{(fi9 ^ (fi9 << ((42) & 63))), fi9}
                                                    _ = v11
                                                default:
                                                    var v12 float32 = 1e10
                                                    _ = v12
                                                }
                                            }
                                        case 882:
                                            v13 := []int{(^(fi6)), 247}
                                            _ = v13
                                        case -9223372036854775808:
                                            fmt.Println(fi6, v8)
                                        }
                                    case "hello":
                                        var v14 float64 = 1e10
                                        _ = v14
                                    }
                                case "hello":
                                    switch 0 {
                                    case -2147483648:
                                        if true {
                                            v8 += p0
                                            var v15 float32 = 1.5
                                            var v16 string = "hello"
                                            _ = v15
                                            _ = v16
                                        }
                                    case 42:
                                        var v17 int64 = 477
                                        if true {
                                            p0 = v8
                                        }
                                        _ = v17
                                    }
                                case "foo bar":
                                    if ((!(3.14159 < -2.25)) || true) {
                                        p0 -= -128
                                    } else {
                                        v18 := GenS1{F0: v8, F1: 75}
                                        switch ("hello" + (("\n\t" + "foo bar") + "")) {
                                        case "a":
                                            var v19 bool = false
                                            _ = v19
                                        case "hello":
                                            p0 = (p0 * (p0 ^ v8))
                                            v20 := []int{}
                                            _ = v20
                                        case "\n\t":
                                            v21 := []int{553, fi6, -124}
                                            _ = v21
                                        }
                                        _ = v18
                                    }
                                }
                            } else {
                                if (127 >= 2) {
                                    var v22 uint32 = 809
                                    _ = v22
                                } else {
                                    switch 32767 {
                                    case 128:
                                        fmt.Println(fi6, fi6)
                                    default:
                                        var v23 bool = false
                                        _ = v23
                                    }
                                    fmt.Println(p0, fi9, fi6)
                                }
                                v24 := []int{fi6}
                                fmt.Println(v8)
                                _ = v24
                            }
                        case -254:
                            if (true && false) {
                                v25 := GenS1{F0: p0, F1: 2}
                                v26 := []int{}
                                p0 += p0
                                _ = v25
                                _ = v26
                            }
                        default:
                            var v27 float64 = 3.14159
                            _ = v27
                        }
                        _ = fi9
                    }
                    _ = v8
                }
            case "a":
                for fi28 := 0; fi28 < 16; fi28++ {
                    if true {
                        v29 := []int{}
                        var v30 int32 = -2147483648
                        switch "a" {
                        case "foo bar":
                            v30 += v30
                        default:
                            p0 -= 42
                        }
                        _ = v29
                        _ = v30
                    }
                    if false {
                        p0 -= p0
                    }
                    _ = fi28
                }
            case "":
                for fi31 := 0; fi31 < 1; fi31++ {
                    p0 += p0
                    p0 = (p0 / (((p0 + p0)) | 1))
                    switch -84 {
                    case -328:
                        p0 = (p0 + p0)
                    }
                    if true {
                        if true {
                            v32 := GenS2{F0: (-(fi31)), F1: 128}
                            _ = v32
                        } else {
                            var v33 bool = (false && false)
                            _ = v33
                        }
                        var v34 string = ""
                        _ = v34
                    }
                    v35 := GenS3{F0: ("\n\t" + ("foo bar" + "a")), F1: 18446744073709551615, F2: 611}
                    _ = fi31
                    _ = v35
                }
            }
        case 128:
            p0 = p0
        default:
            v36 := []int{128, (fi6 % ((fi6) | 1)), fi6, fi6}
            _ = v36
        }
        _ = fi6
    }
    for fi37 := 0; fi37 < 20; fi37++ {
        p0 = p0
        _ = fi37
    }
    for fi38 := 0; fi38 < 20; fi38++ {
        if false {
            var v39 int64 = 1
            var v40 float32 = 0.1
            _ = v39
            _ = v40
        } else {
            v41 := GenS2{F0: 9223372036854775807, F1: 42}
            p0 -= p0
            switch "foo bar" {
            case "\n\t":
                v42 := []int{(fi38 ^ 4294967295), 32767}
                _ = v42
            }
            _ = v41
        }
        _ = fi38
    }
    _ = p0
    return 255
}
func genFn1(p0 uint32, p1 uint32) int16 {
    var v43 uint = 780
    switch "" {
    case "\n\t":
        switch ("foo bar" + "a") {
        case "a":
            var v44 uint = 1
            _ = v44
        case "foo bar":
            var v45 uint = (v43 >> ((uint(p0)) & 63))
            _ = v45
        case "":
            p0 = 0
        default:
            for fi46 := 0; fi46 < 2; fi46++ {
                p0 -= (p0 / (((p1 >> ((255) & 63))) | 1))
                fmt.Println(fi46, v43)
                fmt.Println(v43, p1)
                _ = fi46
            }
        }
    case "":
        if (1 == -2147483648) {
            p0 = (p0 + (^(p1)))
        } else {
            for fi47 := 0; fi47 < 2; fi47++ {
                var v48 uint16 = uint16(p0)
                p1 -= uint32(v43)
                v49 := []int{(fi47 >> (((v43 << (((v43 << ((1) & 63))) & 63))) & 63)), fi47, 32767}
                _ = fi47
                _ = v48
                _ = v49
            }
            v50 := []int{0, -213, -444, -509}
            if (true || false) {
                fmt.Println(v43, p0)
            } else {
                var v51 int8 = 42
                for fi52 := 0; fi52 < 1; fi52++ {
                    for fi53 := 0; fi53 < 18; fi53++ {
                        fmt.Println(fi52, fi53)
                        var v54 bool = false
                        _ = fi53
                        _ = v54
                    }
                    switch "foo bar" {
                    case "a":
                        v55 := GenS3{F0: "foo bar", F1: 904, F2: 0}
                        _ = v55
                    case "hello":
                        switch 42 {
                        case 128:
                            v56 := GenS4{F0: 1, F1: fi52, F2: 604}
                            _ = v56
                        case -1:
                            v50 = append(v50, fi52)
                        }
                    }
                    if ((v43 & v43) < genFn0((v51 << ((genFn0(100)) & 63)))) {
                        for fi57 := 0; fi57 < 18; fi57++ {
                            var v58 uint = (v43 % ((65535) | 1))
                            var v59 int32 = -421
                            v43 = (v58 * (^(v58)))
                            v60 := GenS1{F0: v51, F1: 128}
                            var v61 uint64 = 18446744073709551615
                            v62 := GenS5{F0: (v59 | 42)}
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = fi52
                }
                _ = v51
            }
            _ = v50
        }
    }
    _ = p0
    _ = p1
    _ = v43
    return 1000
}

func main() {
    var v1 int = 525
    v2 := make(map[string]string)
    v3 := []int{511, 32767}
    v4 := []int{-9223372036854775808}
    v5 := make(map[string]bool)
    for fi63 := 0; fi63 < 8; fi63++ {
        for fi64 := 0; fi64 < 13; fi64++ {
            switch ("\n\t" + "a") {
            case "foo bar":
                v1 += (fi63 ^ v1)
            case "a":
                if false {
                    switch fi64 {
                    case -750:
                        var v65 string = "hello"
                        _ = v65
                    case 1:
                        _ = len(v4)
                    case -889:
                        fmt.Println(v1, v1, fi64)
                    }
                    switch "a" {
                    case "hello":
                        v5[""] = (3.14159 >= (1e10 * 3.14159))
                    case "\n\t":
                        if (1000 == -723) {
                            _ = len(v3)
                        } else {
                            v1 += 32768
                            v1 -= fi63
                        }
                    case "foo bar":
                        var v66 int64 = 67
                        _ = v66
                    }
                    var v67 int16 = -785
                    _ = v67
                }
            default:
                _ = len(v3)
            }
            _ = fi64
        }
        v1 += (fi63 % ((32768) | 1))
        _ = fi63
    }
    for fi68 := 0; fi68 < 6; fi68++ {
        var v69 float32 = (0.0 + 1e10)
        _ = fi68
        _ = v69
    }
    v70 := GenS6{F0: 1, F1: 2}
    for fi71 := 0; fi71 < 11; fi71++ {
        v1 += 0
        var v72 string = (("hello" + ((("\n\t" + "\n\t") + "foo bar") + "\n\t")) + "\n\t")
        _ = fi71
        _ = v72
    }
    v73 := GenS2{F0: v1, F1: 255}
    if (false || (true && true)) {
        var v74 uint8 = 248
        for fi75 := 0; fi75 < 2; fi75++ {
            v74 = (v74 >> ((0) & 63))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi76 := 0; fi76 < 16; fi76++ {
                v2["hello"] = "foo bar"
                if len(v4) > 0 {
                    v4[0] = (fi75 & 2147483647)
                }
                _ = fi76
            }
            _ = fi75
        }
        _ = v74
    } else {
        fmt.Println(v1)
        if false {
            if len(v3) > 0 {
                v3[0] = v1
            }
        }
        var v77 float32 = (1e10 * float32(1.5))
        _ = v77
    }
    v1 = 1
    var v78 uint = 18446744073709551615
    var v79 int64 = 42
    for fi80 := 0; fi80 < 12; fi80++ {
        v81 := GenS7{F0: 4294967295}
        _ = fi80
        _ = v81
    }
    switch "" {
    case "a":
        fmt.Println(v78)
    case "":
        v1 += v1
    case "hello":
        switch (v1 / ((127) | 1)) {
        case -1:
            v78 = 1
        case 4294967295:
            switch "hello" {
            case "\n\t":
                var v82 int8 = 100
                _ = v82
            }
        case 1:
            v78 = v78
        default:
            v83 := GenS8{F0: v79}
            _ = v83
        }
    default:
        var v84 uint16 = 42
        _ = v84
    }
    for fi85 := 0; fi85 < 3; fi85++ {
        v86 := GenS9{F0: 775, F1: (v78 ^ v78)}
        var v87 int8 = -1
        if true {
            v88 := v5[(("hello" + "foo bar") + (("a" + ("" + "\n\t")) + "a"))]
            var v89 uint32 = uint32(v78)
            _ = v88
            _ = v89
        }
        _ = fi85
        _ = v86
        _ = v87
    }
    switch v1 {
    case 32768:
        switch "foo bar" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "":
            var v90 int8 = 1
            _ = v90
        case "a":
            _ = len(v5)
        default:
            v79 = int64(v78)
        }
    default:
        var v91 bool = (2147483648 < 32)
        _ = v91
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v73
    _ = v78
    _ = v79
}
