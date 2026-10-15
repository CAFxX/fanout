package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint; F2 int16 }

func main() {
    var v1 bool = (638 <= -478)
    var v2 uint8 = 127
    v3 := make(map[int]uint32)
    var v4 float64 = 3.14159
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 1}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "foo bar" {
        case "hello":
            var v8 uint32 = 423
            _ = v8
        default:
            if (v1 && false) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    switch (fi7 + (fi9 << ((0) & 63))) {
                    case 1:
                        switch (^(fi9)) {
                        case 32768:
                            _ = len(v3)
                        case 1:
                            var v10 bool = (!true)
                            _ = v10
                        }
                    case 32768:
                        switch (fi9 * (fi9 >> ((129) & 63))) {
                        case 42:
                            v1 = (v1 || (!false))
                        default:
                            var v11 uint32 = 1
                            _ = v11
                        }
                    default:
                        if (!v1) {
                            v12 := []int{(fi9 | 249), fi9, -314}
                            _ = v12
                        }
                    }
                    v13 := GenS1{F0: 403}
                    v3[fi7] = 201
                    _ = fi9
                    _ = v13
                }
                var v14 uint16 = 1
                _ = v14
            } else {
                switch "" {
                case "hello":
                    v15 := GenS1{F0: uint(v2)}
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 7; fi16++ {
                        switch "" {
                        case "":
                            var v17 int8 = -128
                            _ = v17
                        default:
                            fmt.Println(v2, v2, fi16)
                        }
                        var v18 bool = false
                        var v19 int16 = 0
                        _ = fi16
                        _ = v18
                        _ = v19
                    }
                case "a":
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 int16 = 42
                        _ = fi20
                        _ = v21
                    }
                }
            }
        }
        v4 = 3.14159
        v2 -= (v2 & 2)
        _ = fi7
    }
    v2 = v2
    v22 := GenS1{F0: 255}
    var v23 int16 = -430
    _ = len(v3)
    var v24 uint8 = (v2 * 127)
    switch "" {
    case "foo bar":
        if ("" != ("\n\t" + "foo bar")) {
            v25 := GenS2{F0: (v23 & 32767), F1: 1, F2: "a"}
            switch -9223372036854775808 {
            case 2147483647:
                if false {
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v4 = v4
                        switch "" {
                        case "hello":
                            var v27 int8 = -128
                            _ = v27
                        case "a":
                            if (-1 >= 61) {
                                v28 := GenS2{F0: v23, F1: 0, F2: "\n\t"}
                                var v29 int8 = -105
                                _ = v28
                                _ = v29
                            }
                        }
                        for fi30 := 0; fi30 < 18; fi30++ {
                            switch (fi26 - -271) {
                            case 127:
                                var v31 uint = 0
                                _ = v31
                            case -332:
                                v32 := []int{(fi30 & -228), 2147483647, -2147483648}
                                _ = v32
                            }
                            _ = fi30
                        }
                        _ = fi26
                    }
                    if v1 {
                        for fi33 := 0; fi33 < 8; fi33++ {
                            for fi34 := 0; fi34 < 16; fi34++ {
                                v23 += v23
                                _ = fi34
                            }
                            if (2147483647 < fi33) {
                                var v35 bool = true
                                switch "\n\t" {
                                case "a":
                                    if (100 >= -128) {
                                        v24 += 193
                                        switch "\n\t" {
                                        case "":
                                            _ = len(v3)
                                        case "\n\t":
                                            v24 -= (v2 & (v24 + v24))
                                        case "foo bar":
                                            if ((0 == 65535) && (fi33 != int(v23))) {
                                                fmt.Println(fi33, v4)
                                                v4 = v4
                                                var v36 float64 = 1.5
                                                _ = v36
                                            }
                                        default:
                                            var v37 uint16 = 32767
                                            _ = v37
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(fi33)
                                case "\n\t":
                                    v38 := []int{fi33, fi33, fi33}
                                    _ = v38
                                default:
                                    switch "hello" {
                                    case "":
                                        for fi39 := 0; fi39 < 15; fi39++ {
                                            var v40 uint = 18446744073709551615
                                            _ = fi39
                                            _ = v40
                                        }
                                    case "hello":
                                        for fi41 := 0; fi41 < 13; fi41++ {
                                            if (1e10 >= (0.1 + (0.0 + (1e10 / -2.25)))) {
                                                if true {
                                                    if v35 {
                                                        v4 -= 0.1
                                                        switch ("\n\t" + "\n\t") {
                                                        case "\n\t":
                                                            v42 := []int{}
                                                            _ = v42
                                                        case "":
                                                            if (v4 >= ((v4 * v4) * v4)) {
                                                                delete(v3, 9223372036854775807)
                                                            } else {
                                                                v2 += v24
                                                                v23 += (v23 + 32767)
                                                                _ = len(v5)
                                                            }
                                                        }
                                                        v23 += (v23 * (v23 % ((v23) | 1)))
                                                    }
                                                    v43 := GenS2{F0: (v23 * v23), F1: 0, F2: "foo bar"}
                                                    _ = v43
                                                } else {
                                                    fmt.Println(v2, v4)
                                                }
                                            }
                                            var v44 float64 = v4
                                            var v45 float64 = v4
                                            _ = fi41
                                            _ = v44
                                            _ = v45
                                        }
                                    case "foo bar":
                                        fmt.Println(v35, v2, fi33)
                                    default:
                                        if v35 {
                                            for fi46 := 0; fi46 < 16; fi46++ {
                                                if (18446744073709551615 < 192) {
                                                    var v47 float32 = 3.14159
                                                    var v48 int = fi46
                                                    _ = v47
                                                    _ = v48
                                                }
                                                var v49 int16 = v23
                                                v24 += 2
                                                _ = fi46
                                                _ = v49
                                            }
                                            v50 := []int{(fi33 >> ((42) & 63))}
                                            if len(v50) >= 2 {
                                                v50 = v50[1:len(v50)]
                                            }
                                            _ = v50
                                        } else {
                                            for fi51 := 0; fi51 < 3; fi51++ {
                                                var v52 uint8 = 0
                                                v53 := GenS3{F0: "", F1: 100}
                                                _ = fi51
                                                _ = v52
                                                _ = v53
                                            }
                                            var v54 float64 = 3.14159
                                            _ = v54
                                        }
                                    }
                                }
                                _ = v35
                            }
                            if false {
                                v55 := []int{(-(fi33)), fi33, -9223372036854775808, fi33}
                                var v56 uint = 1
                                _ = v55
                                _ = v56
                            }
                            _ = fi33
                        }
                    }
                }
            case 32767:
                if v1 {
                    fmt.Println(v24, v4, v4)
                }
            case 804:
                for fi57 := 0; fi57 < 14; fi57++ {
                    v4 += (float64(0.1) + v4)
                    var v58 string = "foo bar"
                    _ = fi57
                    _ = v58
                }
            default:
                v24 -= v2
            }
            _ = v25
        }
    case "a":
        for fi59 := 0; fi59 < 17; fi59++ {
            var v60 uint32 = 4294967295
            v1 = false
            _ = fi59
            _ = v60
        }
    }
    v1 = false
    if (32768 <= 128) {
        fmt.Println(v23, v24, v1)
        v61 := GenS4{F0: ""}
        _ = v61
    } else {
        v62 := GenS5{F0: 0, F1: 920, F2: (v23 - (v23 - v23))}
        _ = v62
    }
    var v63 int8 = -128
    if false {
        v1 = v1
    } else {
        v64 := []int{}
        v24 = (^(v2))
        for fi65 := 0; fi65 < 9; fi65++ {
            v66 := GenS5{F0: (fi65 + fi65), F1: 42, F2: v23}
            _ = fi65
            _ = v66
        }
        _ = v64
    }
    var v67 bool = v1
    v24 = (^(v2))
    var v68 uint8 = v24
    fmt.Println(v2, v63, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v67
    _ = v68
}
