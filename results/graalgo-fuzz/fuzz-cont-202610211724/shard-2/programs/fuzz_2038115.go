package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "foo bar"
    v3 := make(map[int]bool)
    var v4 int32 = 1
    fmt.Println(v4)
    var v5 float64 = 1e10
    switch v2 {
    case "":
        var v6 float64 = (0.1 - v5)
        _ = v6
    case "\n\t":
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                switch (fi8 % ((fi7) | 1)) {
                case 516:
                    v9 := []int{fi7}
                    _ = v9
                case 128:
                    fmt.Println(fi8)
                case 194:
                    if ((int(v4) > fi8) || (false || (1.5 != 1.5))) {
                        switch (v2 + "a") {
                        case "a":
                            v2 = ("" + (v2 + ""))
                        case "":
                            v10 := []int{(fi7 << ((255) & 63)), fi8}
                            _ = v10
                        default:
                            v11 := GenS1{F0: (121 < -128), F1: 65535}
                            _ = v11
                        }
                        fmt.Println(fi8, v2)
                    }
                default:
                    switch "\n\t" {
                    case "\n\t":
                        switch v2 {
                        case "":
                            v2 = (("hello" + v2) + v2)
                        }
                    case "":
                        if false {
                            fmt.Println(fi8)
                            v2 = v2
                        } else {
                            v12 := GenS1{F0: false, F1: 0}
                            fmt.Println(v2, v5)
                            _ = v12
                        }
                    case "hello":
                        v13 := []int{}
                        _ = v13
                    }
                    _ = len(v3)
                    v14 := []int{(fi8 << ((4294967295) & 63))}
                    _ = v14
                }
                v15 := GenS1{F0: (!true), F1: 0}
                _ = fi8
                _ = v15
            }
            v2 = v2
            v16 := []int{fi7, fi7}
            _ = fi7
            _ = v16
        }
    }
    if true {
        v4 -= v4
        if true {
            var v17 int8 = 0
            v18 := []int{1, 1, 9223372036854775807, -989}
            if (true && false) {
                fmt.Println(v4)
                var v19 int = -1
                v4 = (v4 * v4)
                _ = v19
            }
            _ = v17
            _ = v18
        } else {
            v5 += -2.25
            if true {
                v4 = v4
            } else {
                v20 := GenS2{F0: -1, F1: (v2 + v2)}
                _ = v20
            }
        }
        fmt.Println(v4)
    }
    var v21 uint8 = 1
    var v22 float32 = ((((0.1 - -2.25) / (-2.25 + 1e10)) - 1e10) / (0.1 / float32(v5)))
    if (float64(v22) < (v5 + 1e10)) {
        if false {
            v23 := GenS2{F0: 0, F1: v2}
            if true {
                if (0 == 72) {
                    switch 572 {
                    case 4294967295:
                        delete(v3, 2147483647)
                    }
                }
            }
            _ = v23
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                v25 := v3[(fi24 - 128)]
                if (v21 < (v21 << ((2) & 63))) {
                    if v25 {
                        switch (fi24 & fi24) {
                        case -1:
                            fmt.Println(v25, v4)
                        case 127:
                            var v26 uint64 = 18446744073709551615
                            _ = v26
                        }
                        v27 := GenS1{F0: (!v25), F1: 65535}
                        if (-145 == 665) {
                            v25 = v25
                        } else {
                            if true {
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    var v29 uint8 = (v21 + v21)
                                    var v30 uint16 = 806
                                    fmt.Println(v25)
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v2 = (v2 + v2)
                                var v31 uint8 = 128
                                _ = v31
                            }
                            v32 := []int{}
                            _ = v32
                        }
                        _ = v27
                    }
                    fmt.Println(v5, v5)
                    var v33 float32 = (v22 / 0.1)
                    _ = v33
                } else {
                    var v34 int = (fi24 / ((128) | 1))
                    _ = v34
                }
                _ = fi24
                _ = v25
            }
            if false {
                for fi35 := 0; fi35 < 18; fi35++ {
                    switch v2 {
                    case "hello":
                        v36 := GenS2{F0: -714, F1: (((("" + v2) + v2) + ("" + v2)) + (v2 + (v2 + "")))}
                        _ = v36
                    case "a":
                        var v37 int8 = 42
                        _ = v37
                    case "":
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            var v38 bool = true
                            _ = v38
                        }
                    default:
                        v22 = 1e10
                    }
                    v39 := GenS2{F0: 0, F1: v2}
                    v40 := GenS1{F0: false, F1: 255}
                    _ = fi35
                    _ = v39
                    _ = v40
                }
                switch (("a" + v2) + ("\n\t" + "\n\t")) {
                case "foo bar":
                    for fi41 := 0; fi41 < 9; fi41++ {
                        for fi42 := 0; fi42 < 8; fi42++ {
                            fmt.Println(fi42, fi41)
                            v43 := v1[(v2 + (("foo bar" + v2) + v2))]
                            v44 := GenS3{F0: uint64(v43), F1: 822, F2: 255}
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        switch (^(fi41)) {
                        case 42:
                            var v45 int8 = 28
                            _ = v45
                        default:
                            var v46 uint64 = 2
                            _ = v46
                        }
                        _ = fi41
                    }
                default:
                    v2 = v2
                }
            }
            if true {
                if (v21 >= (v21 | v21)) {
                    var v47 float64 = float64(0.0)
                    switch v2 {
                    case "\n\t":
                        var v48 int = 4294967295
                        _ = v48
                    case "foo bar":
                        var v49 int16 = -1
                        _ = v49
                    default:
                        v50 := v3[-2147483648]
                        _ = v50
                    }
                    v51 := []int{128}
                    _ = v47
                    _ = v51
                }
                v21 = 127
            } else {
                fmt.Println(v22, v5, v22)
                v52 := []int{527}
                if (18446744073709551615 > 42) {
                    v4 += v4
                    v53 := GenS4{F0: 0, F1: 1, F2: 32767}
                    var v54 uint64 = 761
                    _ = v53
                    _ = v54
                } else {
                    var v55 int = 547
                    if (42 <= 343) {
                        fmt.Println(v5, v22)
                        v56 := GenS2{F0: -32768, F1: ("" + ("a" + ("hello" + "hello")))}
                        if false {
                            v22 += v22
                        } else {
                            v57 := GenS5{F0: (v21 - 1)}
                            var v58 bool = false
                            _ = v57
                            _ = v58
                        }
                        _ = v56
                    } else {
                        if true {
                            fmt.Println(v55)
                            fmt.Println(v2, v55)
                        } else {
                            v55 += v55
                            var v59 bool = false
                            v22 += v22
                            _ = v59
                        }
                    }
                    switch (v2 + v2) {
                    case "\n\t":
                        switch "foo bar" {
                        case "\n\t":
                            var v60 int16 = 0
                            _ = v60
                        case "foo bar":
                            v61 := GenS6{F0: 2}
                            _ = v61
                        }
                    }
                    _ = v55
                }
                _ = v52
            }
        }
    }
    v21 = v21
    var v62 int32 = (v4 / ((v4) | 1))
    for fi63 := 0; fi63 < 5; fi63++ {
        var v64 uint32 = 414
        _ = len(v3)
        _ = fi63
        _ = v64
    }
    var v65 int64 = 1
    if true {
        if (!false) {
            v66 := []int{0, 127, 42}
            if ("a" != "") {
                v22 = (3.14159 + v22)
            } else {
                if false {
                    if (441 == -709) {
                        fmt.Println(v5, v5)
                    }
                    var v67 int8 = -128
                    _ = v67
                }
            }
            _ = v66
        } else {
            fmt.Println(v2)
        }
    } else {
        var v68 uint8 = 0
        var v69 int8 = -1
        fmt.Println(v5, v21, v69)
        _ = v68
        _ = v69
    }
    v70 := GenS6{F0: 163}
    switch ((v2 + v2) + v2) {
    case "":
        v22 = 3.14159
    }
    switch 2147483647 {
    case 1:
        v22 -= v22
    case 128:
        var v71 uint = 4294967295
        _ = v71
    default:
        v72 := []int{32768, 42}
        _ = v72
    }
    v73 := GenS5{F0: (v21 | (^(v21)))}
    fmt.Println(v62)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v65
    _ = v70
    _ = v73
}
