package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint32 }
func genFn0() int64 {
    var v6 uint64 = 1
    if (450 < 1) {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float64 = -2.25
            var v9 uint64 = (v6 >> ((2) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{128}
        v6 -= (v6 << ((2) & 63))
        _ = v10
    } else {
        v6 = v6
        if false {
            var v11 int = 32768
            v11 = v11
            if (-1 > -2147483648) {
                v11 = -1
            }
            _ = v11
        } else {
            v6 = (v6 / ((v6) | 1))
            v6 = (v6 >> ((255) & 63))
            v6 = (v6 - v6)
        }
    }
    var v12 int64 = 0
    _ = v12
    _ = v6
    return int64(v6)
}
func genFn1(p0 int16, p1 int16) int8 {
    fmt.Println(p0, p0)
    p1 += p1
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := GenS1{F0: -868, F1: -43}
        fmt.Println(fi13)
        _ = fi13
        _ = v14
    }
    p0 = -32768
    v15 := GenS2{F0: 1, F1: 2}
    _ = p0
    _ = p1
    _ = v15
    return 0
}

func main() {
    v1 := []string{""}
    var v2 uint32 = 65535
    var v3 float32 = 1.5
    var v4 int8 = 127
    v5 := make(map[uint32]string)
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint64 = 797
        for fi18 := 0; fi18 < 6; fi18++ {
            switch "hello" {
            case "hello":
                fmt.Println(v4, v17, v2)
            }
            fmt.Println(v17)
            if false {
                fmt.Println(fi16, v3)
                var v19 int = 32768
                _ = v19
            }
            switch (fi18 ^ (-(fi16))) {
            case 290:
                var v20 int = (fi16 - fi18)
                _ = v20
            case 32767:
                if false {
                    _ = len(v5)
                    v1 = append(v1, "hello")
                    if (!(0 > 1)) {
                        v21 := GenS2{F0: 255, F1: 2}
                        v22 := GenS3{F0: "\n\t", F1: (v2 & v2), F2: -564}
                        if false {
                            var v23 string = "foo bar"
                            v4 += v4
                            _ = v23
                        } else {
                            fmt.Println(v2, v3)
                            v24 := GenS2{F0: 127, F1: v17}
                            _ = v24
                        }
                        _ = v21
                        _ = v22
                    } else {
                        v3 = (1e10 + v3)
                        fmt.Println(v3)
                        switch ("a" + "\n\t") {
                        case "\n\t":
                            if ((!true) || ("\n\t" > "a")) {
                                fmt.Println(v17, fi16)
                            } else {
                                _ = len(v5)
                                switch (fi16 ^ fi16) {
                                case 9223372036854775807:
                                    if false {
                                        switch fi16 {
                                        case -2147483648:
                                            switch 1 {
                                            case 42:
                                                switch "" {
                                                case "foo bar":
                                                    v4 = 127
                                                case "\n\t":
                                                    switch fi18 {
                                                    case 700:
                                                        var v25 bool = ("hello" < "a")
                                                        _ = v25
                                                    case -2147483648:
                                                        var v26 string = "a"
                                                        _ = v26
                                                    }
                                                }
                                            default:
                                                v27 := GenS1{F0: 0, F1: (v4 & v4)}
                                                _ = v27
                                            }
                                        case -431:
                                            v28 := GenS2{F0: 165, F1: (v17 / ((v17) | 1))}
                                            _ = v28
                                        case 32768:
                                            var v29 int16 = -922
                                            _ = v29
                                        }
                                        v4 += int8(fi18)
                                        _ = len(v5)
                                    }
                                default:
                                    fmt.Println(v17)
                                }
                            }
                        case "":
                            var v30 float64 = 1e10
                            _ = v30
                        case "a":
                            if false {
                                v17 = v17
                            }
                        default:
                            switch fi18 {
                            case 4294967295:
                                v3 -= (0.1 * 0.1)
                            case 1:
                                if true {
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(fi16, v2, fi16)
                                        var v31 float64 = 3.14159
                                        _ = v31
                                    default:
                                        if false {
                                            var v32 uint64 = v17
                                            _ = v32
                                        } else {
                                            if (!(421 > v2)) {
                                                var v33 int32 = 0
                                                fmt.Println(fi18, v4, fi18)
                                                _ = v33
                                            } else {
                                                fmt.Println(fi18, fi16, v4)
                                                if (true || false) {
                                                    v3 -= 1.5
                                                    fmt.Println(v4, v4)
                                                }
                                                var v34 uint32 = v2
                                                _ = v34
                                            }
                                            _ = len(v1)
                                        }
                                    }
                                    v35 := GenS3{F0: ("foo bar" + "\n\t"), F1: v2, F2: -315}
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    _ = v35
                                }
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    var v36 float32 = float32(-2.25)
                                    _ = v36
                                case "":
                                    if (!true) {
                                        fmt.Println(v2, v4, v2)
                                        if (!true) {
                                            v4 += v4
                                        }
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            if (305 >= 42) {
                                                v3 = v3
                                                var v37 uint8 = 42
                                                v38 := GenS4{F0: 42}
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 uint16 = 42
                                                if true {
                                                    v17 = v17
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                v5[(v2 << ((1) & 63))] = ("a" + (((("foo bar" + "\n\t") + ("a" + "\n\t")) + ("foo bar" + "foo bar")) + "\n\t"))
                                                _ = v39
                                            }
                                        case "a":
                                            v2 += (v2 >> ((255) & 63))
                                        default:
                                            v2 = v2
                                            _ = len(v1)
                                            switch "" {
                                            case "":
                                                var v40 int = 264
                                                _ = v40
                                            case "foo bar":
                                                var v41 uint = 182
                                                _ = v41
                                            case "hello":
                                                v42 := GenS5{F0: 0, F1: 1000}
                                                _ = v42
                                            }
                                        }
                                        switch fi16 {
                                        case 32768:
                                            v4 += v4
                                        default:
                                            v43 := GenS5{F0: -9223372036854775808, F1: 1}
                                            _ = v43
                                        }
                                        _ = len(v5)
                                    }
                                case "\n\t":
                                    v4 = (v4 >> ((42) & 63))
                                }
                            }
                        }
                        v17 = 4294967295
                        var v44 uint8 = 127
                        v1 = append(v1, "\n\t")
                        if (0 < (v2 & (v2 ^ (v2 ^ 42)))) {
                            switch (-(fi18)) {
                            case 4294967295:
                                switch "a" {
                                case "foo bar":
                                    v45 := GenS4{F0: 42}
                                    _ = v45
                                }
                                var v46 int64 = -9223372036854775808
                                _ = v46
                            case 32768:
                                v47 := GenS6{F0: 4294967295}
                                _ = v47
                            default:
                                delete(v5, v2)
                            }
                            v4 -= -128
                            if false {
                                var v48 bool = false
                                v44 += 77
                                _ = v48
                            }
                            var v49 int16 = -32768
                            var v50 int32 = -1
                            _ = v49
                            _ = v50
                        }
                        _ = v44
                    }
                } else {
                    if (127 > (v4 & (v4 | v4))) {
                        v51 := GenS3{F0: "hello", F1: v2, F2: 2147483647}
                        fmt.Println(v2)
                        _ = v51
                    }
                }
            case 32768:
                if true {
                    var v52 uint16 = 42
                    v2 = 2147483647
                    v1 = append(v1, "a")
                    _ = v52
                } else {
                    switch fi16 {
                    case -1:
                        v3 -= float32(0.1)
                    default:
                        var v53 int = (fi16 << ((127) & 63))
                        _ = v53
                    }
                    var v54 float64 = (float64((0.0 / v3)) + 0.0)
                    _ = v54
                }
            }
            v4 = (^(v4))
            var v55 string = "a"
            fmt.Println(v55)
            _ = fi18
            _ = v55
        }
        _ = fi16
        _ = v17
    }
    v3 = ((v3 - v3) / (v3 + v3))
    v56 := GenS7{F0: -2147483648, F1: 409}
    for fi57 := 0; fi57 < 3; fi57++ {
        v3 = 0.0
        _ = fi57
    }
    v58 := v5[(v2 + (v2 / ((4294967295) | 1)))]
    var v59 int8 = -128
    v60 := v5[v2]
    switch 1 {
    case -2147483648:
        var v61 uint16 = 1
        _ = v61
    case 4294967295:
        v62 := GenS7{F0: 0, F1: 951}
        _ = v62
    case 128:
        fmt.Println(v3, v3, v4)
    }
    switch v60 {
    case "":
        switch 0 {
        case 2147483647:
            var v63 int64 = -9223372036854775808
            _ = v63
        case -2147483648:
            fmt.Println(v58, v59)
        case 423:
            for fi64 := 0; fi64 < 14; fi64++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v59 -= int8(fi64)
                switch v60 {
                case "":
                    var v65 uint16 = 0
                    _ = v65
                }
                _ = fi64
            }
        }
    default:
        if (true && (!false)) {
            if true {
                var v66 int32 = -2147483648
                if (-2.25 < (0.0 + 1e10)) {
                    fmt.Println(v2, v60)
                    _ = len(v1)
                }
                v67 := GenS8{F0: v2}
                _ = v66
                _ = v67
            }
            v59 -= 30
            var v68 float32 = float32(3.14159)
            _ = v68
        } else {
            fmt.Println(v60)
            for fi69 := 0; fi69 < 3; fi69++ {
                if true {
                    if false {
                        var v70 uint = 0
                        var v71 int64 = -51
                        _ = v70
                        _ = v71
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                for fi72 := 0; fi72 < 9; fi72++ {
                    if true {
                        var v73 uint = 4294967295
                        if (1 < 835) {
                            fmt.Println(v3)
                            var v74 int64 = 728
                            v75 := GenS6{F0: (v73 + 2)}
                            var v76 int32 = 1
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                        _ = v73
                    }
                    v5[(v2 | (v2 * 247))] = ((v60 + "foo bar") + ("" + (v60 + "a")))
                    _ = fi72
                }
                _ = fi69
            }
            switch 368 {
            case 0:
                for fi77 := 0; fi77 < 1; fi77++ {
                    var v78 int8 = -25
                    fmt.Println(v78)
                    _ = fi77
                    _ = v78
                }
            case -2147483648:
                for fi79 := 0; fi79 < 7; fi79++ {
                    switch v60 {
                    case "":
                        if false {
                            v4 = v59
                        } else {
                            v60 = ((v58 + v58) + "hello")
                            v58 = v60
                            fmt.Println(fi79, v59)
                        }
                    case "a":
                        var v80 float64 = 0.0
                        _ = v80
                    }
                    v4 -= genFn1(-457, -1)
                    _ = fi79
                }
            }
        }
    }
    v3 = 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v59
    _ = v60
}
