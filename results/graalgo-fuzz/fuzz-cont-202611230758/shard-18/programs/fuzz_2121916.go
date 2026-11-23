package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0(p0 int8, p1 uint64, p2 int) uint {
    p2 += 9223372036854775807
    if ((0.0 + 1.5) > ((0.1 / 1.0) - 0.1)) {
        p1 = p1
    }
    var v4 uint = 0
    p0 = int8(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 - 18446744073709551615)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 0
    var v3 float64 = (0.0 * 3.14159)
    fmt.Println(v2, v3)
    var v5 int8 = 34
    var v6 uint16 = 765
    fmt.Println(v2)
    fmt.Println(v5)
    v7 := GenS1{F0: "hello", F1: ((42 != 9223372036854775807) && ((!false) || true)), F2: -9223372036854775808}
    v8 := GenS1{F0: "hello", F1: true, F2: 610}
    switch "a" {
    case "hello":
        switch "a" {
        case "foo bar":
            v9 := GenS1{F0: "hello", F1: (!false), F2: -1}
            _ = v9
        default:
            switch ("\n\t" + "") {
            case "\n\t":
                var v10 float32 = (3.14159 / (0.1 / 3.14159))
                _ = v10
            case "foo bar":
                fmt.Println(v2)
            }
        }
    case "foo bar":
        if false {
            v2 -= 104
            var v11 bool = true
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{(fi12 & -198), 171, int(v6)}
            fmt.Println(v5)
            _ = len(v13)
            _ = fi12
            _ = v13
        }
    }
    if (-372 > 75) {
        switch 0 {
        case 2147483647:
            var v14 int = 32767
            _ = v14
        case 0:
            v15 := GenS1{F0: "hello", F1: false, F2: 1}
            _ = v15
        case 4294967295:
            v3 += (((v3 / v3) * (1e10 * float64((float32(v3) - 0.1)))) + v3)
        }
        v2 -= v2
    } else {
        v16 := GenS2{F0: 1, F1: int8(v2), F2: 52}
        var v17 int = 9223372036854775807
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        switch (("\n\t" + "") + "a") {
        case "\n\t":
            v19 := []int{(fi18 + fi18), (fi18 + 32767)}
            _ = v19
        case "":
            var v20 float32 = 0.0
            _ = v20
        case "a":
            switch int(v6) {
            case 4294967295:
                var v21 int64 = -9223372036854775808
                _ = v21
            default:
                var v22 uint64 = 4294967295
                _ = v22
            }
        default:
            switch fi18 {
            case 358:
                if ((-128 < 127) && (421 < -2147483648)) {
                    if true {
                        var v23 int32 = 1
                        _ = v23
                    }
                    if (42 <= int32(v2)) {
                        if true {
                            var v24 float32 = 1e10
                            v25 := GenS2{F0: 0, F1: v5, F2: 550}
                            _ = v24
                            _ = v25
                        }
                        v3 -= v3
                    }
                    switch ("a" + "foo bar") {
                    case "a":
                        switch "\n\t" {
                        case "":
                            var v26 uint64 = 4294967295
                            _ = v26
                        case "a":
                            var v27 int8 = v5
                            _ = v27
                        case "\n\t":
                            var v28 uint = 4294967295
                            _ = v28
                        }
                    case "\n\t":
                        if false {
                            v3 = 3.14159
                        } else {
                            v6 = v6
                            v29 := GenS3{F0: ((-2.25 * 1.5) != 1.5)}
                            v30 := GenS4{F0: "hello", F1: 4294967295, F2: 32767}
                            _ = v29
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 3; fi31++ {
                            switch ("a" + "\n\t") {
                            case "":
                                v3 -= (0.1 / ((3.14159 * (1e10 / 1.0)) * (v3 * (v3 * (v3 * (0.1 / (float64(1e10) - float64(((float32(v3) + 1.5) - (float32(((v3 / (v3 * 0.0)) + (float64(0.0) / v3))) + 1e10))))))))))
                            }
                            _ = fi31
                        }
                    }
                } else {
                    for fi32 := 0; fi32 < 11; fi32++ {
                        if (-273 <= 975) {
                            var v33 uint32 = 0
                            _ = v33
                        }
                        var v34 string = "foo bar"
                        _ = fi32
                        _ = v34
                    }
                    fmt.Println(v2)
                    for fi35 := 0; fi35 < 16; fi35++ {
                        fmt.Println(v3)
                        var v36 uint32 = 963
                        if ((fi18 & fi18) >= -1) {
                            v3 += 0.0
                        }
                        fmt.Println(fi18, v2)
                        _ = fi35
                        _ = v36
                    }
                }
            default:
                var v37 int = (fi18 >> ((4294967295) & 63))
                _ = v37
            }
        }
        _ = fi18
    }
    switch ("a" + "foo bar") {
    case "hello":
        v5 += 0
    default:
        fmt.Println(v6)
    }
    if false {
        v5 += 0
        switch "a" {
        case "":
            v6 = 734
        case "\n\t":
            switch "foo bar" {
            case "\n\t":
                var v38 int32 = 42
                _ = v38
            case "":
                for fi39 := 0; fi39 < 18; fi39++ {
                    var v40 int = (fi39 | fi39)
                    for fi41 := 0; fi41 < 20; fi41++ {
                        v42 := GenS5{F0: 100}
                        if true {
                            fmt.Println(fi39)
                        } else {
                            var v43 uint64 = 540
                            if true {
                                fmt.Println(v43)
                                v3 = v3
                            }
                            v6 -= v6
                            var v44 uint16 = (v6 / ((uint16(v2)) | 1))
                            _ = v43
                            _ = v44
                        }
                        v45 := GenS2{F0: 707, F1: (-(v5)), F2: 0}
                        _ = fi41
                        _ = v42
                        _ = v45
                    }
                    _ = fi39
                    _ = v40
                }
            case "a":
                fmt.Println(v6, v5, v6)
            default:
                switch 128 {
                case 171:
                    if (true && true) {
                        if (198 < 4294967295) {
                            var v46 float64 = float64(-2.25)
                            var v47 string = ""
                            _ = v46
                            _ = v47
                        } else {
                            v2 = (v2 << ((18446744073709551615) & 63))
                        }
                        v3 = v3
                    }
                case -1:
                    switch 139 {
                    case 4294967295:
                        if false {
                            v48 := []int{-1, -2147483648}
                            _ = v48
                        }
                    case 127:
                        var v49 int32 = 2147483647
                        _ = v49
                    case -578:
                        var v50 int8 = 127
                        _ = v50
                    }
                default:
                    if true {
                        if false {
                            var v51 uint64 = 28
                            v5 = (v5 % ((v5) | 1))
                            v52 := []int{-2147483648, -476}
                            _ = v51
                            _ = v52
                        }
                        if (111 >= v5) {
                            v2 = v2
                            v53 := GenS6{F0: (v2 | (^(v2))), F1: 60, F2: -9223372036854775808}
                            _ = v53
                        }
                        switch "a" {
                        case "\n\t":
                            for fi54 := 0; fi54 < 16; fi54++ {
                                if (("hello" <= "a") || true) {
                                    var v55 uint8 = v2
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            v56 := GenS7{F0: 1, F1: true, F2: -1}
                                            _ = v56
                                        case "foo bar":
                                            fmt.Println(v55, v5, v5)
                                        case "\n\t":
                                            fmt.Println(fi54, v2, v3)
                                        }
                                    case "hello":
                                        switch ("foo bar" + "a") {
                                        case "":
                                            fmt.Println(v6)
                                        default:
                                            v5 -= v5
                                        }
                                    case "":
                                        v2 = (v2 | (v55 ^ 42))
                                    }
                                    for fi57 := 0; fi57 < 10; fi57++ {
                                        fmt.Println(v55, v3, v3)
                                        _ = fi57
                                    }
                                    _ = v55
                                }
                                v58 := GenS7{F0: 261, F1: true, F2: -2147483648}
                                _ = fi54
                                _ = v58
                            }
                        case "":
                            fmt.Println(v6)
                        case "hello":
                            v59 := GenS6{F0: v2, F1: v6, F2: 42}
                            _ = v59
                        default:
                            for fi60 := 0; fi60 < 10; fi60++ {
                                v5 = (v5 >> ((uint(v2)) & 63))
                                var v61 float64 = (float64(-2.25) - float64((1.5 + -2.25)))
                                v6 = (v6 | (v6 ^ v6))
                                _ = fi60
                                _ = v61
                            }
                        }
                    }
                }
            }
        default:
            v62 := []int{391, 516, -700}
            _ = v62
        }
        v63 := []int{}
        _ = v63
    } else {
        v64 := GenS2{F0: 910, F1: (v5 >> ((1) & 63)), F2: 4294967295}
        if true {
            var v65 bool = ((v3 > (0.0 / ((-2.25 * v3) - 3.14159))) || true)
            if v65 {
                var v66 int8 = (v5 ^ v5)
                _ = v66
            }
            _ = v65
        } else {
            var v67 uint32 = 2
            _ = v67
        }
        _ = v64
    }
    v1[2147483648] = 4294967295
    if false {
        var v68 uint = 94
        var v69 int16 = 356
        var v70 uint8 = (v2 | v2)
        _ = v68
        _ = v69
        _ = v70
    } else {
        v71 := GenS2{F0: 65535, F1: v5, F2: 18446744073709551615}
        _ = v71
    }
    v72 := v1[482]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v72
    _ = v8
}
