package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 string; F1 uint32; F2 int }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint8 = 128
    var v2 int32 = 1
    v3 := []int{546, 32767}
    v2 = 42
    for fi4 := 0; fi4 < 5; fi4++ {
        switch (fi4 * fi4) {
        case 127:
            _ = len(v3)
        default:
            var v5 int = fi4
            _ = v5
        }
        v1 -= v1
        switch (^(fi4)) {
        case 42:
            if true {
                var v6 float32 = 0.1
                _ = v6
            }
        case -131:
            v1 = 143
        case 310:
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint64 = 0
                v9 := GenS1{F0: 325, F1: 81}
                switch (fi4 ^ (-(fi4))) {
                case -279:
                    var v10 uint8 = (v1 + 128)
                    _ = v10
                case 368:
                    v11 := GenS1{F0: 2147483648, F1: 42}
                    _ = v11
                default:
                    if (9223372036854775807 <= (^(fi7))) {
                        var v12 float64 = (((1.5 / ((3.14159 + float64(-2.25)) + -2.25)) * 1.5) - 3.14159)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v13 string = ("a" + "")
                        _ = v12
                        _ = v13
                    } else {
                        if true {
                            fmt.Println(v2)
                            switch 2147483647 {
                            case -2147483648:
                                _ = len(v3)
                            case 665:
                                var v14 uint64 = 42
                                _ = v14
                            case 0:
                                var v15 int16 = -648
                                _ = v15
                            }
                        } else {
                            v2 -= (^(v2))
                        }
                    }
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        _ = fi4
    }
    var v16 int8 = 0
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS2{F0: 54, F1: (v2 / (((-(v2))) | 1)), F2: 881}
        if (!((^(v1)) != v1)) {
            switch ((("foo bar" + "") + "a") + (("a" + "") + ("hello" + (((("hello" + ("a" + "a")) + (("foo bar" + ("\n\t" + ("a" + "\n\t"))) + "a")) + "\n\t") + "hello")))) {
            case "\n\t":
                switch fi17 {
                case 92:
                    _ = len(v3)
                case 594:
                    var v19 float64 = -2.25
                    _ = v19
                default:
                    switch fi17 {
                    case 4294967295:
                        if (!(!(((877 < 0) && (fi17 < 127)) && true))) {
                            v16 = (v16 | v16)
                        }
                    case -2147483648:
                        switch fi17 {
                        case 0:
                            fmt.Println(v2)
                        case 1:
                            var v20 string = "hello"
                            _ = v20
                        }
                    default:
                        switch ("foo bar" + ("a" + ("\n\t" + ""))) {
                        case "foo bar":
                            switch ("hello" + "") {
                            case "":
                                v2 += -1
                            default:
                                var v21 int = fi17
                                _ = v21
                            }
                        case "hello":
                            fmt.Println(fi17)
                        }
                    }
                }
            case "":
                fmt.Println(fi17, v1, v16)
            case "hello":
                v2 = v2
            }
        } else {
            switch "a" {
            case "":
                v22 := GenS3{F0: (v16 / (((v16 | v16)) | 1)), F1: true}
                _ = v22
            case "\n\t":
                v16 += (v16 * v16)
            case "a":
                v23 := GenS4{F0: 4294967295, F1: 4294967295}
                _ = v23
            }
            v1 += v1
            v24 := GenS5{F0: ("hello" + ("\n\t" + "")), F1: 173, F2: (-(fi17))}
            _ = v24
        }
        _ = fi17
        _ = v18
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = (v1 | 128)
    for fi25 := 0; fi25 < 14; fi25++ {
        if len(v3) > 0 {
            v3[0] = (fi25 ^ (fi25 * fi25))
        }
        _ = fi25
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v16 -= (v16 ^ int8(v2))
        v2 += (^(v2))
        _ = fi26
    }
    fmt.Println(v1, v2)
    v2 -= v2
    if false {
        if true {
            v1 -= v1
            var v27 int8 = -128
            _ = v27
        } else {
            switch "hello" {
            case "\n\t":
                for fi28 := 0; fi28 < 11; fi28++ {
                    if true {
                        var v29 int64 = 42
                        if false {
                            v30 := GenS6{F0: 735, F1: false}
                            _ = v30
                        } else {
                            v31 := GenS2{F0: 42, F1: int32(v16), F2: -9223372036854775808}
                            for fi32 := 0; fi32 < 17; fi32++ {
                                if false {
                                    fmt.Println(v29, v2)
                                    var v33 uint16 = 0
                                    _ = v33
                                } else {
                                    var v34 uint64 = 42
                                    fmt.Println(v2, fi28)
                                    _ = v34
                                }
                                if len(v3) > 0 {
                                    v3[0] = (fi32 - -367)
                                }
                                v16 = v16
                                _ = fi32
                            }
                            var v35 int = int(v16)
                            _ = v31
                            _ = v35
                        }
                        _ = v29
                    }
                    v36 := GenS6{F0: 4294967295, F1: ((1.5 / ((float32((float64((-2.25 * -2.25)) / 1.5)) * 0.1) * (1.5 + ((float32((1.5 - 1e10)) - 0.1) * 3.14159)))) < (float32(3.14159) - 0.0))}
                    _ = fi28
                    _ = v36
                }
            case "a":
                for fi37 := 0; fi37 < 10; fi37++ {
                    if (true || (true || (!false))) {
                        for fi38 := 0; fi38 < 2; fi38++ {
                            var v39 uint8 = (v1 + v1)
                            if (-238 >= 792) {
                                if false {
                                    switch fi37 {
                                    case 32768:
                                        v40 := GenS2{F0: 65535, F1: (v2 >> ((582) & 63)), F2: 0}
                                        _ = v40
                                    case 246:
                                        var v41 bool = false
                                        _ = v41
                                    default:
                                        fmt.Println(fi37, v2)
                                        v39 -= v1
                                        fmt.Println(v1, v1)
                                    }
                                    v3 = append(v3, (fi37 >> ((4294967295) & 63)))
                                    switch 4294967295 {
                                    case 725:
                                        switch (-(fi37)) {
                                        case 127:
                                            var v42 uint32 = 1
                                            _ = v42
                                        case 42:
                                            var v43 uint16 = 718
                                            _ = v43
                                        case 9223372036854775807:
                                            if len(v3) > 0 {
                                                v3[0] = (fi38 >> ((0) & 63))
                                            }
                                        }
                                    case 2147483647:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    default:
                                        v44 := GenS5{F0: ("" + ("\n\t" + ("hello" + ("\n\t" + ("" + ("" + ("\n\t" + ("hello" + ("foo bar" + ("foo bar" + "foo bar")))))))))), F1: 762, F2: fi37}
                                        _ = v44
                                    }
                                }
                                switch (fi37 / (((fi38 - fi37)) | 1)) {
                                case 128:
                                    var v45 uint8 = (v1 >> ((2) & 63))
                                    _ = v45
                                case 4294967295:
                                    if len(v3) > 0 {
                                        v3[0] = (fi37 / ((fi38) | 1))
                                    }
                                default:
                                    v2 = v2
                                    switch (("" + "foo bar") + "a") {
                                    case "foo bar":
                                        v16 = v16
                                    }
                                    if false {
                                        if false {
                                            v46 := GenS7{F0: int(v16), F1: 858}
                                            v39 = v39
                                            _ = v46
                                        } else {
                                            if ((0 != 0) && (2 == uint(v39))) {
                                                var v47 uint32 = 4294967295
                                                _ = v47
                                            }
                                            v48 := GenS1{F0: 611, F1: 651}
                                            _ = v48
                                        }
                                        fmt.Println(fi38, v2, fi37)
                                    }
                                }
                            }
                            _ = len(v3)
                            v39 = 127
                            _ = len(v3)
                            _ = fi38
                            _ = v39
                        }
                        v16 += -99
                        v1 = (v1 << ((992) & 63))
                    }
                    switch fi37 {
                    case -994:
                        var v49 uint64 = 0
                        _ = v49
                    case 9223372036854775807:
                        v16 -= 73
                    }
                    _ = fi37
                }
            case "":
                var v50 int16 = -32768
                _ = v50
            }
            var v51 uint = 255
            _ = v51
        }
    }
    var v52 int32 = (^(v2))
    var v53 bool = true
    fmt.Println(v16, v52)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v52
    _ = v53
}
