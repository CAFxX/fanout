package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int64; F1 uint64 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int, p2 int64) bool {
    p0 -= (^(p0))
    v5 := GenS1{F0: false}
    switch "foo bar" {
    case "hello":
        var v6 uint32 = 4294967295
        _ = v6
    case "\n\t":
        if false {
            switch (p1 % ((p1) | 1)) {
            case 188:
                switch (^(p1)) {
                case 957:
                    fmt.Println(p0, p2)
                default:
                    switch p1 {
                    case -9223372036854775808:
                        v7 := []int{(p1 << ((42) & 63)), (p1 * 2147483647)}
                        _ = v7
                    case 9223372036854775807:
                        p1 = p1
                    case 128:
                        var v8 uint16 = p0
                        _ = v8
                    }
                }
            case 127:
                var v9 uint = 1
                _ = v9
            case 947:
                if ((255 > 2) || (true || (713 > -34))) {
                    if ((0 < 18446744073709551615) || true) {
                        v10 := []int{4294967295}
                        v11 := GenS1{F0: false}
                        _ = v10
                        _ = v11
                    }
                    switch ("hello" + "foo bar") {
                    case "":
                        if false {
                            for fi12 := 0; fi12 < 4; fi12++ {
                                switch fi12 {
                                case 4294967295:
                                    v13 := []int{fi12, fi12}
                                    _ = v13
                                }
                                _ = fi12
                            }
                        }
                    default:
                        p2 = (p2 >> ((18446744073709551615) & 63))
                    }
                    fmt.Println(p2)
                } else {
                    v14 := []int{}
                    _ = v14
                }
            }
        } else {
            p1 += 32767
            if (815 == 2147483647) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    switch (fi15 % (((fi15 + p1)) | 1)) {
                    case 0:
                        for fi16 := 0; fi16 < 20; fi16++ {
                            fmt.Println(p2, fi16, fi16)
                            _ = fi16
                        }
                    case 977:
                        fmt.Println(p0, p2, fi15)
                    default:
                        for fi17 := 0; fi17 < 3; fi17++ {
                            p0 = p0
                            _ = fi17
                        }
                    }
                    var v18 int32 = 0
                    _ = fi15
                    _ = v18
                }
                p0 += p0
            }
        }
    }
    v19 := []int{498}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v5
    return false
}
func genFn1(p0 bool) uint64 {
    v20 := GenS1{F0: (0 > -2147483648)}
    if true {
        if false {
            fmt.Println(p0)
        }
    } else {
        switch ("hello" + ("foo bar" + ("hello" + "hello"))) {
        case "foo bar":
            p0 = (!((0 > -377) && (false && ("a" <= ("" + ("" + "a"))))))
        case "\n\t":
            v21 := []int{-235, -385, 0, 641}
            _ = v21
        case "a":
            if (1e10 > 3.14159) {
                fmt.Println(p0)
            } else {
                if p0 {
                    p0 = (0 <= 0)
                } else {
                    var v22 uint32 = 65535
                    var v23 uint64 = 2
                    _ = v22
                    _ = v23
                }
                switch ("" + ("a" + "a")) {
                case "":
                    fmt.Println(p0)
                case "foo bar":
                    p0 = true
                default:
                    var v24 int64 = 9223372036854775807
                    _ = v24
                }
                if false {
                    v25 := GenS1{F0: (("\n\t" + "foo bar") == "hello")}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        if (true && (false && p0)) {
                            v27 := GenS2{F0: 110, F1: "", F2: ((("hello" + "hello") + ("hello" + "\n\t")) + ("hello" + ""))}
                            switch fi26 {
                            case 32768:
                                v28 := GenS1{F0: p0}
                                _ = v28
                            case 983:
                                for fi29 := 0; fi29 < 15; fi29++ {
                                    switch fi26 {
                                    case -106:
                                        v30 := GenS2{F0: 42, F1: "foo bar", F2: ""}
                                        var v31 int8 = 1
                                        var v32 int = 127
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                    case 78:
                                        var v33 int = fi29
                                        _ = v33
                                    case 128:
                                        var v34 int32 = -1
                                        p0 = (!(221 < 127))
                                        _ = v34
                                    }
                                    var v35 int16 = 1000
                                    v35 = v35
                                    _ = fi29
                                    _ = v35
                                }
                            case 32767:
                                v36 := []int{(fi26 * (fi26 >> ((18446744073709551615) & 63))), fi26, fi26}
                                _ = v36
                            }
                            _ = v27
                        }
                        v37 := GenS1{F0: (fi26 == (fi26 * fi26))}
                        for fi38 := 0; fi38 < 9; fi38++ {
                            v39 := GenS1{F0: p0}
                            v40 := GenS2{F0: 127, F1: "foo bar", F2: "\n\t"}
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        _ = fi26
                        _ = v37
                    }
                    var v41 int64 = 42
                    _ = v41
                }
            }
        default:
            v42 := GenS3{F0: 42}
            _ = v42
        }
    }
    p0 = (2 > 4294967295)
    _ = p0
    _ = v20
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 338
    v2 := make(map[string]bool)
    var v3 string = "\n\t"
    var v4 uint32 = 740
    v4 -= (v4 | 0)
    fmt.Println(v3)
    for fi43 := 0; fi43 < 17; fi43++ {
        for fi44 := 0; fi44 < 14; fi44++ {
            v45 := []int{(^(fi44))}
            _ = fi44
            _ = v45
        }
        _ = fi43
    }
    var v46 float64 = 0.1
    if true {
        v4 += v4
    }
    if false {
        switch 2147483647 {
        case 32767:
            var v47 uint16 = 602
            _ = v47
        }
        v48 := []int{42}
        _ = v48
    } else {
        if false {
            for fi49 := 0; fi49 < 20; fi49++ {
                v3 = v3
                _ = fi49
            }
            _ = len(v2)
            v50 := []int{-1}
            _ = v50
        }
    }
    switch 32768 {
    case -9223372036854775808:
        for fi51 := 0; fi51 < 10; fi51++ {
            switch 32767 {
            case 32767:
                switch fi51 {
                case 32768:
                    switch (v3 + (v3 + v3)) {
                    case "foo bar":
                        v1 = v1
                    case "":
                        if false {
                            var v52 int32 = 2147483647
                            v46 = v46
                            _ = v52
                        }
                    }
                case -1:
                    for fi53 := 0; fi53 < 12; fi53++ {
                        v1 = 42
                        v1 -= (^(v1))
                        v54 := GenS4{F0: v3, F1: false, F2: (v1 % ((0) | 1))}
                        _ = fi53
                        _ = v54
                    }
                case -9223372036854775808:
                    v1 = v1
                default:
                    if ((v1 | v1) != (v1 & (v1 + v1))) {
                        switch "foo bar" {
                        case "foo bar":
                            v4 -= uint32(v1)
                        case "a":
                            var v55 float64 = 0.1
                            _ = v55
                        }
                        switch 32768 {
                        case 75:
                            if true {
                                for fi56 := 0; fi56 < 4; fi56++ {
                                    fmt.Println(v1)
                                    switch (fi51 / (((fi51 * -185)) | 1)) {
                                    case 42:
                                        if (-1 > -675) {
                                            v4 = (^(v4))
                                            var v57 bool = true
                                            _ = v57
                                        }
                                    case 161:
                                        var v58 float64 = 1.5
                                        _ = v58
                                    }
                                    var v59 string = (v3 + (("\n\t" + ((v3 + v3) + (v3 + v3))) + v3))
                                    _ = fi56
                                    _ = v59
                                }
                                var v60 uint8 = 42
                                fmt.Println(v60, v46)
                                _ = v60
                            }
                        }
                    } else {
                        fmt.Println(v46, v4)
                    }
                }
            case -49:
                fmt.Println(v46, v3)
            case 2147483647:
                if (65535 == 42) {
                    if true {
                        switch "\n\t" {
                        case "a":
                            if (!((^(v4)) >= (v4 / ((532) | 1)))) {
                                switch fi51 {
                                case -713:
                                    v1 -= (v1 * (v1 - v1))
                                case -415:
                                    switch -2147483648 {
                                    case 128:
                                        v61 := GenS5{F0: 0}
                                        _ = v61
                                    case -90:
                                        v3 = (v3 + v3)
                                    }
                                default:
                                    var v62 uint16 = (v1 << ((2) & 63))
                                    _ = v62
                                }
                            }
                        default:
                            v3 = v3
                        }
                        v63 := GenS6{F0: 0, F1: 58}
                        var v64 uint64 = 456
                        _ = v63
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 3; fi65++ {
                        v46 = -2.25
                        v66 := []int{}
                        _ = fi65
                        _ = v66
                    }
                }
            default:
                v46 -= (((0.1 * 1e10) / (3.14159 + v46)) - ((v46 - float64(1e10)) * v46))
            }
            _ = fi51
        }
    case 42:
        v46 = (float64(float32(v46)) + (v46 + v46))
    case 32768:
        v4 -= v4
    default:
        v1 = 0
    }
    if false {
        v67 := GenS5{F0: 787}
        for fi68 := 0; fi68 < 19; fi68++ {
            v69 := GenS2{F0: 42, F1: (v3 + (((v3 + "foo bar") + (((v3 + "hello") + "foo bar") + ((v3 + (((v3 + v3) + v3) + v3)) + v3))) + "foo bar")), F2: (("\n\t" + (v3 + v3)) + (v3 + "hello"))}
            _ = fi68
            _ = v69
        }
        if (!(18446744073709551615 > genFn1(true))) {
            for fi70 := 0; fi70 < 5; fi70++ {
                if (v3 != v3) {
                    v71 := []int{}
                    fmt.Println(v4, v4, v1)
                    _ = v71
                } else {
                    v72 := GenS7{F0: 0, F1: 116}
                    v4 -= v4
                    _ = v72
                }
                _ = fi70
            }
        }
        _ = v67
    }
    fmt.Println(v3, v3)
    switch v3 {
    case "":
        for fi73 := 0; fi73 < 20; fi73++ {
            v4 += (^(v4))
            v3 = "foo bar"
            _ = fi73
        }
    case "a":
        if ((-398 > 42) && (!false)) {
            for fi74 := 0; fi74 < 9; fi74++ {
                for fi75 := 0; fi75 < 9; fi75++ {
                    var v76 bool = true
                    v77 := GenS8{F0: genFn1((((fi74 >> ((42) & 63)) != (^(fi75))) && (!false))), F1: (v4 ^ (v4 * v4))}
                    v78 := GenS9{F0: (v1 + v1)}
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                }
                _ = fi74
            }
            var v79 string = v3
            if true {
                v1 = v1
                v80 := []int{127, -9223372036854775808, -618, 0}
                _ = v80
            } else {
                v4 -= v4
                v1 = 255
            }
            _ = v79
        }
    case "\n\t":
        v81 := []int{}
        _ = v81
    }
    fmt.Println(v3)
    switch 42 {
    case 127:
        if (228 > 1) {
            switch v3 {
            case "hello":
                for fi82 := 0; fi82 < 13; fi82++ {
                    var v83 string = "\n\t"
                    _ = fi82
                    _ = v83
                }
            }
        }
    default:
        v2[(v3 + "a")] = false
    }
    if false {
        var v84 int16 = 1
        if false {
            var v85 string = ("" + v3)
            v86 := GenS8{F0: 0, F1: v4}
            _ = v85
            _ = v86
        }
        _ = v84
    }
    var v87 string = ("\n\t" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v87
}
