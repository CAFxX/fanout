package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 string; F2 int64 }
type GenS9 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 int64, p1 uint16) int32 {
    v5 := GenS1{F0: 128, F1: p1, F2: 51}
    if true {
        if true {
            var v6 uint16 = p1
            var v7 int64 = (p0 / ((p0) | 1))
            if (1 < 467) {
                v8 := []int{-582, -187}
                if len(v8) > 0 {
                    v8[0] = 127
                }
                _ = v8
            }
            _ = v6
            _ = v7
        } else {
            p0 -= p0
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint32 = uint32(p1)
                var v11 bool = true
                switch "foo bar" {
                case "":
                    var v12 uint16 = 32767
                    _ = v12
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        fmt.Println(p1, p1)
    } else {
        switch "\n\t" {
        case "foo bar":
            if (!true) {
                if (0 < -1) {
                    p0 = (p0 & p0)
                    var v13 int8 = -128
                    v14 := GenS1{F0: 127, F1: p1, F2: v13}
                    _ = v13
                    _ = v14
                } else {
                    if ((499 > 2147483647) && false) {
                        p1 -= 155
                    } else {
                        fmt.Println(p0, p0)
                        v15 := GenS2{F0: 42}
                        v16 := []int{-719, -2147483648, 32767}
                        _ = v15
                        _ = v16
                    }
                }
            }
        case "\n\t":
            v17 := []int{128, -9223372036854775808, 903, -2147483648}
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483647
}
func genFn1() uint8 {
    var v18 string = ""
    var v19 uint = 42
    for fi20 := 0; fi20 < 18; fi20++ {
        switch v18 {
        case "hello":
            v19 = (v19 - (v19 / ((v19) | 1)))
        case "a":
            v21 := []int{42, 42, 32767, fi20}
            _ = v21
        case "":
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 int8 = -128
                v18 = v18
                _ = fi22
                _ = v23
            }
        }
        switch v18 {
        case "hello":
            v24 := []int{(fi20 % ((fi20) | 1))}
            _ = v24
        case "a":
            fmt.Println(fi20)
        case "foo bar":
            if (2 == 146) {
                switch v18 {
                case "":
                    v25 := GenS1{F0: 0, F1: 1, F2: 12}
                    _ = v25
                case "\n\t":
                    v26 := GenS1{F0: 0, F1: 0, F2: -37}
                    _ = v26
                }
            }
        default:
            v18 = v18
        }
        if true {
            v27 := GenS2{F0: 680}
            for fi28 := 0; fi28 < 13; fi28++ {
                v29 := GenS3{F0: 839, F1: 65535}
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        _ = fi20
    }
    var v30 uint = v19
    _ = v18
    _ = v19
    _ = v30
    return 77
}

func main() {
    var v1 uint16 = 2
    var v2 string = ("hello" + ("\n\t" + ("" + "a")))
    var v3 string = v2
    v4 := []int{}
    fmt.Println(v1, v1)
    var v31 uint16 = (^(v1))
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 float32 = 1.5
        for fi34 := 0; fi34 < 6; fi34++ {
            fmt.Println(v31)
            _ = len(v4)
            _ = fi34
        }
        _ = fi32
        _ = v33
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (0.0 == float64((1.5 * float32((float64(-2.25) + ((0.1 * (float64((3.14159 - (((-2.25 - 1e10) / 1.5) - ((float32((0.0 / float64(3.14159))) / 1.5) - (3.14159 - 3.14159))))) * (1e10 / 1.0))) - 3.14159)))))) {
        switch (((v2 + v2) + "a") + (v2 + v3)) {
        case "\n\t":
            if false {
                switch "" {
                case "foo bar":
                    v35 := GenS1{F0: 119, F1: 0, F2: 0}
                    _ = v35
                case "":
                    fmt.Println(v1)
                case "hello":
                    switch (v2 + "") {
                    case "hello":
                        v36 := GenS4{F0: 18446744073709551615, F1: -853, F2: 93}
                        _ = v36
                    case "foo bar":
                        var v37 uint = 1
                        _ = v37
                    case "":
                        v38 := GenS1{F0: 9, F1: 42, F2: 1}
                        _ = v38
                    }
                }
                for fi39 := 0; fi39 < 4; fi39++ {
                    var v40 int64 = 0
                    v31 = (^(v31))
                    _ = fi39
                    _ = v40
                }
            } else {
                var v41 bool = false
                var v42 int32 = -2147483648
                var v43 int8 = -128
                _ = v41
                _ = v42
                _ = v43
            }
        case "":
            _ = len(v4)
        default:
            v2 = v3
        }
        v3 = (v3 + v2)
        for fi44 := 0; fi44 < 19; fi44++ {
            if true {
                var v45 int8 = 42
                var v46 float32 = 3.14159
                _ = v45
                _ = v46
            }
            v47 := GenS5{F0: 129, F1: 18446744073709551615}
            _ = fi44
            _ = v47
        }
    } else {
        v3 = v2
        fmt.Println(v2, v2, v1)
    }
    if true {
        v48 := GenS3{F0: 9223372036854775807, F1: 65535}
        if true {
            for fi49 := 0; fi49 < 19; fi49++ {
                v50 := GenS4{F0: 472, F1: 1, F2: 1}
                v1 -= (v1 / ((v1) | 1))
                _ = fi49
                _ = v50
            }
            var v51 string = v3
            _ = v51
        }
        _ = v48
    }
    if true {
        v52 := GenS4{F0: 255, F1: 1, F2: 1}
        _ = v52
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        fmt.Println(fi53, v31, v2)
        var v54 int64 = 1
        if ((1e10 - 0.0) != (3.14159 + 1.5)) {
            var v55 int16 = -1
            if ((true && ((v55 + (v55 ^ v55)) <= v55)) || false) {
                v56 := GenS2{F0: -1}
                _ = v56
            } else {
                v3 = (("a" + v2) + v2)
                v57 := GenS3{F0: (v54 - (v54 % (((v54 / ((-1) | 1))) | 1))), F1: 684}
                _ = v57
            }
            _ = v55
        } else {
            for fi58 := 0; fi58 < 15; fi58++ {
                var v59 int = (^(fi53))
                switch (v3 + "foo bar") {
                case "a":
                    v54 = 1
                }
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v60 int32 = 913
                    var v61 float32 = 0.1
                    _ = v60
                    _ = v61
                } else {
                    var v62 bool = (false && (198 > v1))
                    _ = v62
                }
                _ = fi58
                _ = v59
            }
            v54 += 82
            var v63 int = (fi53 & fi53)
            _ = v63
        }
        _ = fi53
        _ = v54
    }
    _ = len(v4)
    if (!((uint32(v1) == 522) && (false && (226 > 220)))) {
        for fi64 := 0; fi64 < 1; fi64++ {
            var v65 uint64 = 65
            _ = fi64
            _ = v65
        }
        if true {
            v2 = "a"
            v1 = v31
        } else {
            v1 = 42
        }
        fmt.Println(v3)
    }
    fmt.Println(v31, v2, v2)
    v4 = append(v4, 9223372036854775807)
    if (((!(false && (false && (true && false)))) && false) || (-2147483648 >= 2147483647)) {
        for fi66 := 0; fi66 < 8; fi66++ {
            v1 = 2
            v67 := GenS6{F0: v1}
            _ = fi66
            _ = v67
        }
        v4 = append(v4, 32768)
        if (!(false && true)) {
            for fi68 := 0; fi68 < 8; fi68++ {
                for fi69 := 0; fi69 < 13; fi69++ {
                    var v70 float64 = (-2.25 * float64((((1.5 - 1e10) - 3.14159) + ((-2.25 / -2.25) - (1e10 * float32(1.5))))))
                    switch (-(fi69)) {
                    case -2147483648:
                        if (false || false) {
                            v71 := GenS7{F0: 127}
                            _ = v71
                        }
                    }
                    if (false || (!false)) {
                        fmt.Println(v31, v70, fi69)
                    }
                    _ = fi69
                    _ = v70
                }
                var v72 int64 = 42
                _ = fi68
                _ = v72
            }
            var v73 uint16 = v1
            _ = v73
        } else {
            v74 := GenS7{F0: 2}
            _ = v74
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v75 := GenS8{F0: (42 > v31), F1: v3, F2: 42}
    if (42 >= -857) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v76 := GenS9{F0: 456, F1: 237, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v75
    _ = v76
}
