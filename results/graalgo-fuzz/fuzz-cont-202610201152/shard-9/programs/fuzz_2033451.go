package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS5 struct { F0 int16; F1 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int16; F2 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int; F1 uint8; F2 string }

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 670
    var v3 uint8 = 42
    var v4 uint = 2
    var v5 int8 = 17
    delete(v1, "\n\t")
    v2 += -9223372036854775808
    if false {
        v3 -= v3
        v6 := []int{9223372036854775807}
        _ = v6
    }
    var v7 int32 = 134
    fmt.Println(v5, v5, v4)
    if (v5 < (-(v5))) {
        v8 := v1["hello"]
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{}
            if true {
                if false {
                    v1["\n\t"] = (v8 + "")
                    if (true || ((!true) || false)) {
                        v11 := GenS1{F0: 4294967295, F1: int64(v7)}
                        if true {
                            v12 := GenS1{F0: (^(fi9)), F1: v2}
                            _ = v12
                        }
                        _ = v11
                    }
                    switch "\n\t" {
                    case "hello":
                        v3 += v3
                    case "foo bar":
                        switch v8 {
                        case "":
                            for fi13 := 0; fi13 < 16; fi13++ {
                                if false {
                                    v14 := GenS2{F0: 45}
                                    switch v8 {
                                    case "foo bar":
                                        fmt.Println(fi13)
                                    case "a":
                                        fmt.Println(v8, fi13, v5)
                                    default:
                                        if false {
                                            v15 := GenS3{F0: (v8 + (v8 + "")), F1: true}
                                            v16 := GenS3{F0: "hello", F1: ((v3 | v3) <= (v3 << ((v4) & 63)))}
                                            if (2 < 142) {
                                                v7 = -2147483648
                                                switch fi9 {
                                                case 128:
                                                    v3 -= (v3 | 130)
                                                case 32768:
                                                    if ((true || (0 >= 42)) && (147 != uint64(v4))) {
                                                        var v17 uint32 = 4294967295
                                                        _ = v17
                                                    } else {
                                                        switch (-(fi9)) {
                                                        case 4294967295:
                                                            var v18 int32 = -1
                                                            _ = v18
                                                        case 9223372036854775807:
                                                            var v19 bool = (2 <= 1)
                                                            _ = v19
                                                        case 1:
                                                            fmt.Println(v5, v5)
                                                            if false {
                                                                switch (v8 + (((((v8 + ((("a" + (v8 + v8)) + (("" + ("\n\t" + (v8 + (v8 + (v8 + "a"))))) + "a")) + (v8 + "a"))) + v8) + ((v8 + ("" + (("hello" + (v8 + ("a" + ("foo bar" + v8)))) + (v8 + "")))) + (("\n\t" + "") + v8))) + v8) + v8)) {
                                                                case "\n\t":
                                                                    v4 += v4
                                                                default:
                                                                    if ((v8 + ((v8 + v8) + v8)) < v8) {
                                                                        v20 := GenS2{F0: 0}
                                                                        _ = v20
                                                                    } else {
                                                                        if (false && false) {
                                                                            v2 = (v2 >> ((v4) & 63))
                                                                            _ = len(v10)
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                var v21 int64 = (v2 * (v2 >> (((v4 | 380)) & 63)))
                                                                                _ = v21
                                                                            case "\n\t":
                                                                                if (true && true) {
                                                                                    v3 -= 2
                                                                                    v22 := GenS2{F0: uint16(v4)}
                                                                                    _ = v22
                                                                                }
                                                                            case "a":
                                                                                v2 -= 9223372036854775807
                                                                            }
                                                                        }
                                                                        switch (fi13 & -542) {
                                                                        case -2147483648:
                                                                            fmt.Println(v2, fi9)
                                                                        case 127:
                                                                            delete(v1, (((v8 + "\n\t") + v8) + (v8 + v8)))
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    v4 += v4
                                                                }
                                                                switch 55 {
                                                                case 32768:
                                                                    v2 += (v2 / ((v2) | 1))
                                                                default:
                                                                    _ = len(v1)
                                                                }
                                                            } else {
                                                                fmt.Println(v2, fi13)
                                                                var v23 uint = v4
                                                                var v24 int8 = (v5 - v5)
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                        default:
                                                            v25 := GenS2{F0: 673}
                                                            _ = v25
                                                        }
                                                        v8 = (v8 + "")
                                                        if ((((0.1 / 1e10) < 0.0) && (!false)) || false) {
                                                            if ((v5 * v5) != v5) {
                                                                var v26 int = fi13
                                                                switch (fi13 * (v26 / ((42) | 1))) {
                                                                case 1:
                                                                    if false {
                                                                        v1[""] = v8
                                                                        v2 -= (v2 & -849)
                                                                        var v27 float32 = 1e10
                                                                        if len(v10) > 0 {
                                                                            v10[0] = (fi9 % ((v26) | 1))
                                                                        }
                                                                        var v28 int16 = int16(v3)
                                                                        if true {
                                                                            var v29 uint32 = 2147483648
                                                                            _ = v29
                                                                        }
                                                                        if (42 > 4294967295) {
                                                                            v3 = 44
                                                                            v5 = -1
                                                                            v28 = v28
                                                                            var v30 int32 = v7
                                                                            fmt.Println(v26)
                                                                            _ = v30
                                                                        } else {
                                                                            switch fi9 {
                                                                            case 215:
                                                                                v31 := GenS3{F0: v8, F1: (0 > 1)}
                                                                                _ = v31
                                                                            case 2147483647:
                                                                                v27 -= v27
                                                                            }
                                                                            var v32 uint32 = 42
                                                                            v26 -= (fi9 - int(v7))
                                                                            _ = v32
                                                                        }
                                                                        _ = v27
                                                                        _ = v28
                                                                    } else {
                                                                        v5 = (-(v5))
                                                                        var v33 int16 = -865
                                                                        _ = v33
                                                                    }
                                                                default:
                                                                    if true {
                                                                        switch (fi13 | 165) {
                                                                        case 9223372036854775807:
                                                                            var v34 uint64 = 488
                                                                            _ = v34
                                                                        case 1:
                                                                            delete(v1, v8)
                                                                        case -1:
                                                                            fmt.Println(v26)
                                                                        default:
                                                                            v35 := GenS4{F0: 255, F1: true, F2: 0}
                                                                            _ = v35
                                                                        }
                                                                        v36 := GenS5{F0: 1, F1: false}
                                                                        v37 := v1[v8]
                                                                        fmt.Println(v4, fi13, v37)
                                                                        _ = v36
                                                                        _ = v37
                                                                    } else {
                                                                        var v38 string = v8
                                                                        _ = v38
                                                                    }
                                                                }
                                                                var v39 string = ""
                                                                _ = v26
                                                                _ = v39
                                                            } else {
                                                                if (42 > (v7 ^ (v7 + (v7 * 42)))) {
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v40 int64 = (v2 - -9223372036854775808)
                                                                        _ = v40
                                                                    case "a":
                                                                        if true {
                                                                            var v41 string = v8
                                                                            switch v41 {
                                                                            case "hello":
                                                                                if false {
                                                                                    var v42 int = (fi13 / ((651) | 1))
                                                                                    v43 := GenS2{F0: 255}
                                                                                    v44 := GenS2{F0: 1}
                                                                                    v2 = (v2 - v2)
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                    _ = v44
                                                                                } else {
                                                                                    var v45 bool = (0 > 1)
                                                                                    _ = v45
                                                                                }
                                                                            case "foo bar":
                                                                                v3 += (v3 - 255)
                                                                            default:
                                                                                v3 -= (^(v3))
                                                                            }
                                                                            fmt.Println(v5)
                                                                            v1[("hello" + v41)] = ("\n\t" + v41)
                                                                            _ = v41
                                                                        } else {
                                                                            var v46 float32 = 0.0
                                                                            _ = v46
                                                                        }
                                                                    case "foo bar":
                                                                        v5 = v5
                                                                    default:
                                                                        if true {
                                                                            v47 := GenS6{F0: v8}
                                                                            fmt.Println(v3, v2, v3)
                                                                            _ = len(v10)
                                                                            _ = v47
                                                                        }
                                                                    }
                                                                    var v48 uint = (^(v4))
                                                                    _ = v48
                                                                }
                                                                v3 += v3
                                                                fmt.Println(fi13)
                                                            }
                                                            v49 := GenS7{F0: v7}
                                                            v50 := GenS5{F0: 42, F1: ((-213 < (v2 << (((v4 + v4)) & 63))) || (false && ((true || true) || true)))}
                                                            _ = v49
                                                            _ = v50
                                                        } else {
                                                            v5 = v5
                                                            if true {
                                                                var v51 int64 = (v2 / ((-171) | 1))
                                                                v52 := GenS7{F0: int32(v4)}
                                                                v5 -= v5
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                var v53 int = -1
                                                                _ = v53
                                                            }
                                                            fmt.Println(v2, v2)
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    if true {
                                                        v3 = v3
                                                        var v54 float32 = float32(0.1)
                                                        _ = v54
                                                    }
                                                default:
                                                    var v55 int8 = (v5 & (v5 + (v5 >> (((^(v4))) & 63))))
                                                    _ = v55
                                                }
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    if (!false) {
                                                        if (float64(3.14159) == 0.0) {
                                                            var v56 uint16 = 2
                                                            var v57 string = v8
                                                            if len(v10) > 0 {
                                                                v10[0] = (fi9 - fi13)
                                                            }
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    } else {
                                                        v5 += v5
                                                        if (v5 == -113) {
                                                            fmt.Println(fi9)
                                                            var v58 uint64 = uint64(v4)
                                                            v5 += (v5 | 0)
                                                            _ = v58
                                                        } else {
                                                            v59 := GenS1{F0: fi13, F1: (v2 ^ 1)}
                                                            v7 = v7
                                                            switch fi13 {
                                                            case 32767:
                                                                v5 -= v5
                                                            default:
                                                                v5 -= int8(v7)
                                                            }
                                                            _ = v59
                                                        }
                                                        v7 = 1
                                                        var v60 uint64 = 42
                                                        _ = v60
                                                    }
                                                case 0:
                                                    fmt.Println(fi13, fi9, v7)
                                                }
                                            }
                                            _ = v15
                                            _ = v16
                                        } else {
                                            v2 = (v2 - v2)
                                            if len(v10) >= 2 {
                                                v10 = v10[1:len(v10)]
                                            }
                                        }
                                    }
                                    _ = v14
                                }
                                switch (fi13 >> (((v4 & 1)) & 63)) {
                                case 413:
                                    v8 = v8
                                default:
                                    fmt.Println(v3)
                                }
                                _ = fi13
                            }
                        case "\n\t":
                            var v61 uint16 = 0
                            _ = v61
                        default:
                            if (2 != 2147483647) {
                                v62 := v1[v8]
                                v2 -= v2
                                _ = v62
                            } else {
                                v2 = (v2 << ((4294967295) & 63))
                                v63 := v1["foo bar"]
                                _ = v63
                            }
                        }
                    case "":
                        delete(v1, v8)
                    }
                }
            } else {
                for fi64 := 0; fi64 < 13; fi64++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v8 = v8
                    switch "hello" {
                    case "a":
                        v65 := GenS1{F0: fi64, F1: -121}
                        _ = v65
                    case "":
                        fmt.Println(fi9, v5)
                    }
                    v66 := GenS8{F0: (v8 + v8), F1: -63, F2: "hello"}
                    _ = fi64
                    _ = v66
                }
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    var v67 int32 = (-(v7))
    v2 = 1
    if false {
        var v68 int16 = int16(v2)
        if false {
            v3 -= v3
            v69 := GenS9{F0: 4294967295}
            _ = v69
        } else {
            fmt.Println(v3, v7)
            for fi70 := 0; fi70 < 15; fi70++ {
                for fi71 := 0; fi71 < 10; fi71++ {
                    if ((!true) || (!(!false))) {
                        switch "" {
                        case "foo bar":
                            var v72 bool = ((-(v68)) <= 340)
                            _ = v72
                        case "a":
                            var v73 string = ("" + "a")
                            _ = v73
                        case "":
                            v74 := GenS10{F0: (fi70 - fi70), F1: (^(v3)), F2: "hello"}
                            _ = v74
                        }
                        v67 -= -1
                        v2 = v2
                    }
                    var v75 uint16 = 32767
                    _ = fi71
                    _ = v75
                }
                _ = fi70
            }
        }
        v76 := []int{127, 32768}
        _ = v68
        _ = v76
    } else {
        for fi77 := 0; fi77 < 18; fi77++ {
            v78 := GenS9{F0: 709}
            v7 = (v7 / ((v67) | 1))
            _ = fi77
            _ = v78
        }
    }
    switch "foo bar" {
    case "":
        for fi79 := 0; fi79 < 8; fi79++ {
            v80 := GenS9{F0: 1}
            v3 = v3
            _ = fi79
            _ = v80
        }
    case "\n\t":
        if true {
            var v81 string = "foo bar"
            switch 4294967295 {
            case 325:
                var v82 int64 = v2
                _ = v82
            case -9223372036854775808:
                if false {
                    var v83 float32 = 0.0
                    v84 := []int{9223372036854775807, 0}
                    _ = v83
                    _ = v84
                } else {
                    if true {
                        fmt.Println(v5, v5)
                        fmt.Println(v4)
                        fmt.Println(v81, v7)
                    }
                }
            }
            _ = v81
        }
    }
    var v85 string = "a"
    v86 := []int{-9223372036854775808, -1, -2147483648, -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v7
    _ = v85
    _ = v86
}
