package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int; F2 int }
type GenS3 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0() int32 {
    var v6 int16 = 1
    switch -23 {
    case -616:
        v6 = v6
    case 32767:
        switch "" {
        case "a":
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 13; fi8++ {
                    v6 -= v6
                    var v9 float64 = -2.25
                    _ = fi8
                    _ = v9
                }
                v6 = v6
                _ = fi7
            }
        case "foo bar":
            if false {
                fmt.Println(v6)
            } else {
                fmt.Println(v6)
            }
        }
    case 42:
        switch "\n\t" {
        case "":
            v6 -= v6
        case "hello":
            if (-2.25 < (float64((float32(0.1) + 1e10)) - 3.14159)) {
                v6 += 690
            } else {
                var v10 int32 = -864
                var v11 uint8 = 61
                _ = v10
                _ = v11
            }
        }
    default:
        var v12 uint = 0
        _ = v12
    }
    v6 += (v6 | (v6 % ((v6) | 1)))
    if false {
        var v13 bool = (470 <= 2147483647)
        v13 = v13
        v13 = (!(1.5 == (3.14159 * 1e10)))
        _ = v13
    } else {
        v14 := []int{32768, 127, 42, -19}
        _ = v14
    }
    _ = v6
    return 42
}
func genFn1(p0 int8, p1 int, p2 int64) uint32 {
    fmt.Println(p2)
    switch 127 {
    case 1:
        for fi15 := 0; fi15 < 6; fi15++ {
            p2 = (p2 & 1)
            p2 += -9223372036854775808
            _ = fi15
        }
    case 128:
        p0 += p0
    default:
        p0 = p0
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int = 9223372036854775807
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    return 403
}

func main() {
    var v1 int64 = 873
    var v2 float32 = 1.5
    v3 := make(map[int]int)
    v4 := make(map[string]bool)
    var v5 uint32 = 2147483648
    v18 := []int{1, 128, 0, 1}
    if (1 >= 1) {
        v2 = 0.1
    }
    var v19 float32 = 1e10
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint8 = 2
        if ("" <= "a") {
            var v22 uint32 = v5
            v23 := GenS1{F0: 0}
            _ = v22
            _ = v23
        } else {
            v21 += (^(v21))
            v5 = genFn1(0, 0, (v1 << ((976) & 63)))
        }
        _ = fi20
        _ = v21
    }
    switch 1 {
    case -2147483648:
        switch "foo bar" {
        case "\n\t":
            for fi24 := 0; fi24 < 4; fi24++ {
                if false {
                    if ((false || (12 <= 42)) || (false || (v1 > 9223372036854775807))) {
                        if true {
                            if true {
                                if (42 >= 4294967295) {
                                    for fi25 := 0; fi25 < 2; fi25++ {
                                        v26 := GenS1{F0: 42}
                                        _ = fi25
                                        _ = v26
                                    }
                                } else {
                                    switch "" {
                                    case "a":
                                        var v27 uint64 = 18446744073709551615
                                        _ = v27
                                    case "foo bar":
                                        v19 = v2
                                    case "hello":
                                        var v28 uint8 = 194
                                        _ = v28
                                    }
                                    if (false && false) {
                                        switch 0 {
                                        case 9223372036854775807:
                                            fmt.Println(v19, v5)
                                        case 416:
                                            v29 := GenS1{F0: 580}
                                            _ = v29
                                        case 4294967295:
                                            v30 := GenS2{F0: fi24, F1: (fi24 % ((33) | 1)), F2: int(v1)}
                                            _ = v30
                                        }
                                        var v31 uint32 = 769
                                        v32 := GenS1{F0: 0}
                                        _ = v31
                                        _ = v32
                                    }
                                    v19 += v2
                                }
                                v33 := GenS2{F0: (fi24 << ((4294967295) & 63)), F1: fi24, F2: (^(fi24))}
                                v18 = append(v18, (fi24 >> ((631) & 63)))
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    delete(v3, (fi24 ^ fi34))
                                    if len(v18) > 0 {
                                        v18[0] = fi34
                                    }
                                    if true {
                                        fmt.Println(fi24)
                                        if false {
                                            var v35 int32 = 42
                                            _ = v35
                                        }
                                        var v36 uint8 = 42
                                        _ = v36
                                    }
                                    _ = fi34
                                }
                            }
                        } else {
                            v19 += float32((-2.25 + (1e10 + 1.5)))
                        }
                        var v37 float64 = (0.0 - (1e10 + (-2.25 - 1.5)))
                        _ = v37
                    }
                    v1 = -9223372036854775808
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                } else {
                    v38 := GenS2{F0: 42, F1: fi24, F2: fi24}
                    _ = v38
                }
                var v39 float64 = -2.25
                _ = fi24
                _ = v39
            }
        }
    case -780:
        var v40 float32 = v19
        _ = v40
    }
    var v41 uint8 = 255
    _ = len(v4)
    fmt.Println(v19, v19, v41)
    for fi42 := 0; fi42 < 8; fi42++ {
        var v43 int8 = -1
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v45 := GenS1{F0: 553}
        _ = fi44
        _ = v45
    }
    if false {
        if (42 != 1) {
            switch int(v1) {
            case -254:
                var v46 float64 = 0.0
                _ = v46
            case 0:
                v2 += 0.0
            }
            v47 := GenS2{F0: 770, F1: -1, F2: -798}
            if true {
                var v48 int64 = 0
                _ = v48
            } else {
                for fi49 := 0; fi49 < 15; fi49++ {
                    v1 = 42
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            for fi50 := 0; fi50 < 17; fi50++ {
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    v51 := v3[(^(fi49))]
                                    _ = v51
                                case "foo bar":
                                    v52 := GenS2{F0: (fi50 / ((fi49) | 1)), F1: (fi49 / ((fi50) | 1)), F2: fi50}
                                    _ = v52
                                case "":
                                    _ = len(v4)
                                }
                                var v53 uint32 = v5
                                _ = fi50
                                _ = v53
                            }
                            for fi54 := 0; fi54 < 1; fi54++ {
                                v2 += float32(0.0)
                                v19 -= v2
                                if true {
                                    fmt.Println(fi49, v19, v19)
                                    if len(v18) > 0 {
                                        v18[0] = (fi49 - fi54)
                                    }
                                } else {
                                    v1 -= 0
                                    v3[42] = fi54
                                    delete(v3, 42)
                                }
                                switch (^(fi54)) {
                                case -736:
                                    if true {
                                        v3[1] = (-(fi49))
                                    }
                                case -2147483648:
                                    var v55 float32 = v19
                                    _ = v55
                                case 32767:
                                    var v56 float64 = 3.14159
                                    _ = v56
                                default:
                                    switch ("\n\t" + ("" + "")) {
                                    case "":
                                        var v57 uint32 = v5
                                        _ = v57
                                    }
                                    v2 = float32(0.0)
                                }
                                _ = fi54
                            }
                            v58 := GenS1{F0: 983}
                            _ = v58
                        } else {
                            fmt.Println(v2, v2, v5)
                        }
                    case "a":
                        fmt.Println(v5, v2, v2)
                    case "":
                        v59 := GenS2{F0: (fi49 * 32768), F1: -388, F2: 189}
                        _ = v59
                    default:
                        if true {
                            switch ("hello" + ("hello" + (((("\n\t" + "a") + "a") + "hello") + "\n\t"))) {
                            case "\n\t":
                                v19 += v2
                            case "a":
                                v1 -= v1
                            case "":
                                v1 -= v1
                            }
                            switch ("foo bar" + "a") {
                            case "hello":
                                var v60 int16 = -985
                                _ = v60
                            }
                        }
                    }
                    switch fi49 {
                    case 830:
                        v61 := GenS1{F0: 255}
                        _ = v61
                    case 4294967295:
                        switch fi49 {
                        case 585:
                            for fi62 := 0; fi62 < 7; fi62++ {
                                if ((uint64(v5) != 42) || true) {
                                    if true {
                                        _ = len(v3)
                                        if true {
                                            v41 = (^(v41))
                                            delete(v4, "")
                                            var v63 uint32 = v5
                                            _ = v63
                                        } else {
                                            v64 := GenS1{F0: 1}
                                            _ = v64
                                        }
                                        var v65 int32 = 0
                                        _ = v65
                                    } else {
                                        v5 = (^(v5))
                                        switch fi49 {
                                        case -2147483648:
                                            var v66 int64 = v1
                                            _ = v66
                                        case -9223372036854775808:
                                            var v67 float32 = 0.1
                                            _ = v67
                                        case 32768:
                                            v19 = 1e10
                                        }
                                    }
                                } else {
                                    v19 = 1.5
                                }
                                _ = fi62
                            }
                        }
                    default:
                        v41 += (^(v41))
                    }
                    _ = fi49
                }
                for fi68 := 0; fi68 < 17; fi68++ {
                    var v69 uint = 18446744073709551615
                    var v70 string = "\n\t"
                    if ((v5 & v5) == genFn1(1, (^(fi68)), v1)) {
                        var v71 uint8 = (v41 >> ((99) & 63))
                        _ = v71
                    }
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
            _ = v47
        }
        v2 = (float32((float64(v19) / 1.0)) * v2)
        if (true && (2 > 1)) {
            v1 -= (v1 - v1)
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                if (true && false) {
                    fmt.Println(v2)
                    v2 -= (1.5 * ((float32(float64((v2 - 0.0))) - v2) + 0.1))
                } else {
                    if len(v18) > 0 {
                        v18[0] = 128
                    }
                    for fi72 := 0; fi72 < 2; fi72++ {
                        if false {
                            fmt.Println(v2, v1, v41)
                            fmt.Println(v1)
                        } else {
                            v41 = v41
                            var v73 bool = (fi72 > (fi72 - fi72))
                            if ((fi72 << ((255) & 63)) != (fi72 | fi72)) {
                                var v74 uint64 = 1
                                _ = v74
                            }
                            _ = v73
                        }
                        var v75 int32 = 2147483647
                        v5 = 2147483647
                        _ = fi72
                        _ = v75
                    }
                }
            default:
                v19 -= 1e10
            }
            if (true || false) {
                v76 := v3[17]
                fmt.Println(v2, v41)
                var v77 int = (v76 & (v76 * v76))
                _ = v76
                _ = v77
            }
            switch "hello" {
            case "\n\t":
                var v78 int = 42
                _ = v78
            case "a":
                v5 = (v5 % ((v5) | 1))
            case "hello":
                var v79 int8 = -125
                _ = v79
            }
        }
    } else {
        if ((!(!((!false) || (!(!(false && false)))))) && false) {
            var v80 uint8 = 35
            if true {
                switch -198 {
                case -275:
                    v81 := GenS3{F0: 100, F1: 128, F2: 2}
                    _ = v81
                default:
                    switch 989 {
                    case -9223372036854775808:
                        var v82 uint = 255
                        _ = v82
                    case 32768:
                        var v83 uint = 422
                        _ = v83
                    case -2147483648:
                        v84 := v4[""]
                        _ = v84
                    }
                }
                switch 9223372036854775807 {
                case 1:
                    v85 := GenS4{F0: "a"}
                    _ = v85
                }
            }
            _ = v80
        } else {
            fmt.Println(v5, v2, v1)
            var v86 uint32 = v5
            var v87 int16 = -32768
            _ = v86
            _ = v87
        }
        v5 = genFn1(42, 526, v1)
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
    }
    if true {
        for fi88 := 0; fi88 < 15; fi88++ {
            var v89 uint64 = 2
            _ = fi88
            _ = v89
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
