package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint64; F2 bool }
type GenS9 struct { F0 uint16; F1 int; F2 int64 }
type GenS10 struct { F0 uint16; F1 uint16; F2 int }
func genFn0() int8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 bool = true
        if v5 {
            v6 := GenS1{F0: ("" + "")}
            _ = v6
        } else {
            if true {
                switch (("foo bar" + "\n\t") + ("hello" + ("a" + "a"))) {
                case "\n\t":
                    switch "" {
                    case "\n\t":
                        var v7 float32 = 1.5
                        _ = v7
                    case "hello":
                        v5 = v5
                    }
                }
            }
            for fi8 := 0; fi8 < 13; fi8++ {
                if ((-116 <= 62) || ((242 == 863) && false)) {
                    v5 = (9223372036854775807 == 42)
                } else {
                    var v9 uint16 = 0
                    switch fi4 {
                    case -728:
                        v5 = v5
                    case -2147483648:
                        var v10 string = ("\n\t" + "foo bar")
                        _ = v10
                    case 128:
                        v11 := GenS1{F0: ""}
                        _ = v11
                    }
                    if (!(1e10 > 0.1)) {
                        var v12 uint32 = 79
                        _ = v12
                    }
                    _ = v9
                }
                v13 := []int{}
                _ = fi8
                _ = v13
            }
            switch "\n\t" {
            case "":
                v14 := []int{}
                _ = v14
            default:
                v15 := GenS1{F0: ("\n\t" + "hello")}
                _ = v15
            }
        }
        _ = fi4
        _ = v5
    }
    v16 := []int{4294967295, 9223372036854775807}
    switch (("hello" + ("a" + "\n\t")) + "hello") {
    case "":
        var v17 float32 = float32(3.14159)
        _ = v17
    }
    var v18 int = 726
    for fi19 := 0; fi19 < 18; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            switch "\n\t" {
            case "a":
                switch "a" {
                case "\n\t":
                    var v21 uint64 = 890
                    _ = v21
                case "a":
                    fmt.Println(fi20, fi20, fi20)
                case "foo bar":
                    fmt.Println(fi20, v18, fi20)
                }
            case "hello":
                var v22 int16 = 0
                _ = v22
            case "foo bar":
                v16 = append(v16, fi19)
            }
            v18 += v18
            if len(v16) > 0 {
                v16[0] = -690
            }
            _ = fi20
        }
        _ = fi19
    }
    _ = v16
    _ = v18
    return 127
}

func main() {
    var v1 int8 = 115
    v2 := []uint32{2147483648, 1, 2147483647, 65535}
    var v3 uint8 = 128
    switch "hello" {
    case "hello":
        v3 = 194
    case "a":
        v3 += 2
    }
    v3 = (v3 / (((v3 / (((v3 + (v3 ^ 127))) | 1))) | 1))
    _ = len(v2)
    switch ("" + "\n\t") {
    case "a":
        switch "a" {
        case "foo bar":
            switch "" {
            case "foo bar":
                if (42 > 42) {
                    fmt.Println(v3)
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                } else {
                    fmt.Println(v1, v1)
                    v23 := GenS2{F0: true}
                    _ = v23
                }
            case "hello":
                var v24 int = 0
                _ = v24
            }
        case "\n\t":
            if true {
                switch "\n\t" {
                case "a":
                    v2 = append(v2, 42)
                case "\n\t":
                    for fi25 := 0; fi25 < 3; fi25++ {
                        var v26 int = (-(fi25))
                        var v27 bool = false
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case "foo bar":
                    if ((v3 & 163) >= v3) {
                        v28 := GenS3{F0: 255, F1: uint16(v3)}
                        var v29 float32 = (((1.5 * (float32(0.0) - (float32(float64(1e10)) / ((0.1 + 1.5) + (-2.25 / float32(-2.25)))))) / 3.14159) - (0.1 + 0.1))
                        _ = v28
                        _ = v29
                    } else {
                        if true {
                            fmt.Println(v3, v3)
                            var v30 float32 = 3.14159
                            _ = v30
                        }
                        v31 := GenS2{F0: false}
                        var v32 int8 = -128
                        _ = v31
                        _ = v32
                    }
                }
                v1 += (-(v1))
            } else {
                _ = len(v2)
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                var v33 bool = false
                _ = v33
            }
        }
    }
    fmt.Println(v1, v3)
    v3 -= (v3 - (v3 / ((v3) | 1)))
    var v34 float64 = ((float64(-2.25) * (1.5 / 1.5)) - float64(-2.25))
    fmt.Println(v1, v3)
    if (((false || true) || true) || false) {
        v35 := GenS4{F0: -32768}
        _ = v35
    }
    switch -613 {
    case -773:
        _ = len(v2)
    case 572:
        for fi36 := 0; fi36 < 18; fi36++ {
            for fi37 := 0; fi37 < 15; fi37++ {
                var v38 uint8 = v3
                if false {
                    v34 += 3.14159
                    var v39 int8 = genFn0()
                    _ = v39
                } else {
                    if (0 >= 42) {
                        switch (-(fi36)) {
                        case 437:
                            var v40 int8 = v1
                            _ = v40
                        }
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi37
                _ = v38
            }
            switch "a" {
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            case "":
                if false {
                    switch fi36 {
                    case 196:
                        v41 := GenS4{F0: -746}
                        _ = v41
                    case 32767:
                        switch fi36 {
                        case 127:
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                        case 131:
                            v34 = v34
                        }
                    default:
                        switch fi36 {
                        case 2147483647:
                            v3 += 127
                        case 9223372036854775807:
                            v2 = append(v2, 2147483648)
                        case 32767:
                            fmt.Println(v34, v1)
                        default:
                            _ = len(v2)
                        }
                    }
                    if true {
                        v2 = append(v2, 42)
                    } else {
                        v34 = v34
                    }
                    v3 -= v3
                }
            case "\n\t":
                v3 = (v3 % ((v3) | 1))
            default:
                for fi42 := 0; fi42 < 16; fi42++ {
                    if false {
                        v43 := GenS5{F0: (("" + "") + ""), F1: 255, F2: -178}
                        if true {
                            v34 -= 0.1
                            if (!false) {
                                switch (fi36 | fi42) {
                                case -2147483648:
                                    if (0.0 >= float32(0.1)) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v44 := GenS5{F0: "a", F1: 1, F2: -34}
                                        _ = v44
                                    } else {
                                        var v45 int8 = genFn0()
                                        switch "a" {
                                        case "hello":
                                            if false {
                                                v34 = ((v34 - v34) + (1e10 - v34))
                                            }
                                        case "foo bar":
                                            fmt.Println(fi36, v34)
                                        }
                                        _ = v45
                                    }
                                case 0:
                                    var v46 uint32 = 4294967295
                                    _ = v46
                                default:
                                    var v47 int32 = 371
                                    _ = v47
                                }
                                switch "" {
                                case "\n\t":
                                    if (!(false || true)) {
                                        var v48 float32 = (1.5 * (-2.25 + (1.5 / (1e10 + 1e10))))
                                        fmt.Println(v1)
                                        _ = v48
                                    }
                                case "hello":
                                    v1 += v1
                                case "":
                                    if false {
                                        var v49 uint16 = 88
                                        var v50 string = "foo bar"
                                        v34 -= float64(3.14159)
                                        _ = v49
                                        _ = v50
                                    } else {
                                        if false {
                                            v51 := GenS6{F0: 2}
                                            v3 += (v3 + 199)
                                            _ = v51
                                        } else {
                                            fmt.Println(v34, v1)
                                            switch (fi36 | fi42) {
                                            case 4294967295:
                                                var v52 uint8 = 255
                                                _ = v52
                                            case 0:
                                                v34 = (0.1 - (v34 + ((v34 + v34) + v34)))
                                            }
                                            v1 = (v1 | (v1 % ((2) | 1)))
                                        }
                                    }
                                }
                            } else {
                                var v53 uint16 = 957
                                if true {
                                    fmt.Println(v3, fi42)
                                    v54 := GenS3{F0: v53, F1: v53}
                                    if true {
                                        if false {
                                            switch 146 {
                                            case 128:
                                                var v55 uint16 = (^(v53))
                                                _ = v55
                                            default:
                                                v56 := GenS1{F0: "hello"}
                                                _ = v56
                                            }
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = 0
                                            }
                                            fmt.Println(v3, v34)
                                            if true {
                                                switch "\n\t" {
                                                case "":
                                                    v57 := GenS5{F0: "", F1: 65535, F2: -161}
                                                    _ = v57
                                                default:
                                                    v3 = (v3 / (((^(v3))) | 1))
                                                }
                                                v3 += 187
                                                switch "foo bar" {
                                                case "hello":
                                                    v3 = v3
                                                case "a":
                                                    fmt.Println(v3, v1)
                                                    switch (fi36 + fi36) {
                                                    case 2147483647:
                                                        v53 += 62
                                                    default:
                                                        if true {
                                                            var v58 string = "hello"
                                                            var v59 int16 = 0
                                                            if true {
                                                                var v60 int64 = 9223372036854775807
                                                                _ = v60
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    }
                                                    v1 -= genFn0()
                                                }
                                                _ = len(v2)
                                                fmt.Println(v1, v34, v34)
                                            } else {
                                                if (!true) {
                                                    if true {
                                                        if len(v2) > 0 {
                                                            v2[0] = 4294967295
                                                        }
                                                        switch "hello" {
                                                        case "a":
                                                            if false {
                                                                var v61 uint64 = 2
                                                                _ = v61
                                                            }
                                                            v62 := GenS6{F0: 42}
                                                            if (230 > -1) {
                                                                switch "" {
                                                                case "":
                                                                    v63 := GenS7{F0: uint(v53)}
                                                                    _ = v63
                                                                case "hello":
                                                                    switch fi36 {
                                                                    case 9223372036854775807:
                                                                        _ = len(v2)
                                                                    case 2147483647:
                                                                        var v64 uint16 = 32767
                                                                        _ = v64
                                                                    }
                                                                }
                                                                fmt.Println(v53, v34)
                                                            }
                                                            _ = v62
                                                        }
                                                        v65 := GenS8{F0: 0, F1: 2, F2: (true || true)}
                                                        fmt.Println(v53)
                                                        _ = v65
                                                    } else {
                                                        switch fi36 {
                                                        case 4294967295:
                                                            v53 = (v53 * (v53 * v53))
                                                        }
                                                    }
                                                    fmt.Println(v1, fi42)
                                                    v34 += float64(0.0)
                                                }
                                            }
                                        }
                                    } else {
                                        _ = len(v2)
                                        fmt.Println(v3, v53)
                                        v66 := GenS7{F0: 0}
                                        v67 := GenS9{F0: (v53 >> ((2) & 63)), F1: 945, F2: 9223372036854775807}
                                        _ = v66
                                        _ = v67
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v68 := GenS5{F0: "", F1: 0, F2: 0}
                                        _ = v68
                                    case "\n\t":
                                        fmt.Println(fi42)
                                        v69 := GenS4{F0: 218}
                                        v1 += int8(v53)
                                        _ = v69
                                    default:
                                        var v70 uint16 = 1
                                        _ = v70
                                    }
                                    _ = v54
                                } else {
                                    v71 := GenS10{F0: (v53 & v53), F1: 394, F2: -9223372036854775808}
                                    _ = v71
                                }
                                _ = v53
                            }
                        }
                        _ = v43
                    }
                    _ = fi42
                }
            }
            _ = fi36
        }
    case 2147483647:
        v34 += 1e10
    default:
        v2 = append(v2, 2147483647)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
}
