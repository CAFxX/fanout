package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint16; F1 uint; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 bool = false
    var v2 int16 = -32768
    v3 := []string{"hello", "a", "foo bar", "foo bar"}
    var v4 uint64 = 0
    if v1 {
        v1 = v1
        fmt.Println(v1)
        v5 := GenS1{F0: 2, F1: 2, F2: 1}
        _ = v5
    } else {
        v2 -= 42
        if v1 {
            for fi6 := 0; fi6 < 8; fi6++ {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = fi6
            }
        } else {
            for fi7 := 0; fi7 < 18; fi7++ {
                v3 = append(v3, ("" + "hello"))
                var v8 uint16 = 255
                _ = fi7
                _ = v8
            }
            switch "foo bar" {
            case "\n\t":
                v9 := GenS1{F0: 42, F1: 65535, F2: 100}
                _ = v9
            case "hello":
                v4 += (v4 ^ v4)
            }
        }
    }
    v1 = true
    if true {
        var v10 int16 = (v2 >> ((0) & 63))
        var v11 float64 = 1.5
        var v12 float64 = float64(3.14159)
        _ = v10
        _ = v11
        _ = v12
    } else {
        fmt.Println(v4)
        v13 := GenS2{F0: 128}
        if false {
            v14 := GenS3{F0: 1, F1: ""}
            if false {
                var v15 float64 = (3.14159 - (0.1 / (float64(-2.25) * 0.1)))
                v3 = append(v3, "a")
                var v16 bool = false
                _ = v15
                _ = v16
            }
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "\n\t")
            }
            _ = v14
        } else {
            _ = len(v3)
        }
        _ = v13
    }
    if (!true) {
        var v17 int64 = 42
        v18 := GenS4{F0: 65535, F1: 712, F2: v17}
        for fi19 := 0; fi19 < 7; fi19++ {
            v17 += (v17 / (((v17 + (v17 | -9223372036854775808))) | 1))
            var v20 int16 = -265
            _ = len(v3)
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    }
    fmt.Println(v4, v2, v2)
    fmt.Println(v2, v4)
    fmt.Println(v2)
    var v21 uint = 18446744073709551615
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = (("foo bar" + (("foo bar" + "\n\t") + ("" + (("foo bar" + ("foo bar" + "\n\t")) + "")))) + "foo bar")
    }
    if (0 != 128) {
        var v22 uint32 = 2
        switch 42 {
        case 32767:
            v22 = 0
        case -464:
            switch 32767 {
            case 2147483647:
                var v23 int64 = int64(v4)
                _ = v23
            case 42:
                if v1 {
                    var v24 uint16 = 0
                    if ((v2 | 32767) < -1) {
                        v21 = (v21 & 42)
                        var v25 string = ""
                        _ = v25
                    }
                    if (v22 > (v22 * 422)) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if (1 > 1) {
                            v21 = 293
                            if v1 {
                                _ = len(v3)
                                v21 += v21
                            }
                            for fi26 := 0; fi26 < 1; fi26++ {
                                if v1 {
                                    for fi27 := 0; fi27 < 14; fi27++ {
                                        v3 = append(v3, ("" + ("" + "\n\t")))
                                        _ = fi27
                                    }
                                    if ("\n\t" <= "") {
                                        if v1 {
                                            for fi28 := 0; fi28 < 3; fi28++ {
                                                v29 := GenS2{F0: -1}
                                                _ = fi28
                                                _ = v29
                                            }
                                        } else {
                                            for fi30 := 0; fi30 < 14; fi30++ {
                                                fmt.Println(v22)
                                                if (42 > (v22 & v22)) {
                                                    v4 = (^(v4))
                                                    if v1 {
                                                        v2 = -777
                                                        v4 = v4
                                                    }
                                                    if (float64(1e10) <= (3.14159 / ((float64(1.5) - float64((1e10 + 3.14159))) * ((0.0 / (1.5 + float64((0.0 + 3.14159)))) - 3.14159)))) {
                                                        v4 += v4
                                                        if (v2 == v2) {
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                if ((v1 && true) && (436 < int64(v2))) {
                                                                    fmt.Println(v1, v21, v4)
                                                                    fmt.Println(v21)
                                                                    switch "a" {
                                                                    case "":
                                                                        fmt.Println(v21)
                                                                    }
                                                                    if (!(9223372036854775807 != 42)) {
                                                                        v31 := GenS2{F0: fi26}
                                                                        switch "a" {
                                                                        case "":
                                                                            v32 := GenS2{F0: fi30}
                                                                            _ = v32
                                                                        case "foo bar":
                                                                            _ = len(v3)
                                                                        }
                                                                        _ = v31
                                                                    }
                                                                    if (((false || v1) || v1) || v1) {
                                                                        v33 := GenS1{F0: 42, F1: (v24 + (v24 & 680)), F2: int8(v2)}
                                                                        switch (fi30 % (((^(fi30))) | 1)) {
                                                                        case -82:
                                                                            switch 42 {
                                                                            case -661:
                                                                                fmt.Println(v1, v22)
                                                                            case 42:
                                                                                v34 := GenS5{F0: (v22 & v22), F1: (v4 / ((v4) | 1))}
                                                                                _ = v34
                                                                            case 1:
                                                                                var v35 int64 = -1
                                                                                _ = v35
                                                                            }
                                                                        case 2147483647:
                                                                            switch (("foo bar" + "hello") + "\n\t") {
                                                                            case "foo bar":
                                                                                var v36 uint16 = 374
                                                                                _ = v36
                                                                            case "\n\t":
                                                                                var v37 int8 = 127
                                                                                _ = v37
                                                                            }
                                                                        case 9223372036854775807:
                                                                            fmt.Println(v22)
                                                                        }
                                                                        v38 := GenS4{F0: (v24 ^ (v24 + v24)), F1: (v21 % (((^(v21))) | 1)), F2: -9223372036854775808}
                                                                        _ = v33
                                                                        _ = v38
                                                                    }
                                                                } else {
                                                                    var v39 int16 = v2
                                                                    _ = v39
                                                                }
                                                            case -2147483648:
                                                                v21 = 18446744073709551615
                                                            }
                                                            v40 := GenS6{F0: v1, F1: 1}
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v40
                                                        }
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                } else {
                                                    var v41 int32 = -2147483648
                                                    var v42 float64 = (0.1 - (float64(((1e10 - 0.1) + ((3.14159 + -2.25) + 1e10))) + (float64(1.5) * 1e10)))
                                                    _ = v41
                                                    _ = v42
                                                }
                                                v43 := GenS7{F0: 0, F1: 2147483647}
                                                _ = fi30
                                                _ = v43
                                            }
                                            switch ("a" + ("a" + "\n\t")) {
                                            case "":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                        }
                                        var v44 uint8 = 127
                                        fmt.Println(v2, v22, fi26)
                                        _ = v44
                                    }
                                }
                                _ = fi26
                            }
                        } else {
                            for fi45 := 0; fi45 < 3; fi45++ {
                                var v46 int16 = v2
                                _ = fi45
                                _ = v46
                            }
                            v21 = (v21 << (((v21 >> ((v21) & 63))) & 63))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        switch "\n\t" {
                        case "":
                            var v47 uint32 = v22
                            _ = v47
                        case "foo bar":
                            if (!(!(true && true))) {
                                var v48 float32 = 3.14159
                                var v49 uint16 = (v24 * 175)
                                _ = len(v3)
                                _ = v48
                                _ = v49
                            }
                        default:
                            var v50 uint = v21
                            _ = v50
                        }
                    }
                    _ = v24
                }
            default:
                v21 += (v21 / ((v21) | 1))
            }
        default:
            fmt.Println(v4, v22, v21)
        }
        v4 = (v4 - v4)
        _ = v22
    } else {
        v3 = append(v3, ("a" + "\n\t"))
        var v51 uint8 = 94
        v1 = (18446744073709551615 == (v21 * (v21 ^ v21)))
        _ = v51
    }
    v52 := GenS8{F0: 928}
    v53 := GenS2{F0: 32768}
    if ((((((1e10 / (((1.5 / float32(float64(1.5))) + -2.25) / (3.14159 + (3.14159 * -2.25)))) / 1e10) / (1.5 / float32(1.5))) - 1e10) - float32(-2.25)) >= ((3.14159 - (0.1 / float32(((0.1 - float64(((((0.0 * (1.5 / -2.25)) - (0.0 + (1e10 * -2.25))) * (3.14159 * 0.0)) / 1e10))) + (-2.25 + 0.1))))) * (((float32(1e10) * float32((float64(0.1) + (0.1 / 1.5)))) - 0.1) + 0.1))) {
        var v54 float32 = 0.1
        v3 = append(v3, "\n\t")
        _ = len(v3)
        _ = v54
    }
    v55 := GenS3{F0: 4294967295, F1: ("" + ((("foo bar" + "") + "\n\t") + ((("foo bar" + "hello") + "foo bar") + "a")))}
    for fi56 := 0; fi56 < 6; fi56++ {
        var v57 float64 = (float64(1.5) * (-2.25 / 3.14159))
        switch "hello" {
        case "a":
            var v58 uint64 = (v4 * 42)
            _ = v58
        }
        v1 = true
        _ = fi56
        _ = v57
    }
    fmt.Println(v2)
    for fi59 := 0; fi59 < 8; fi59++ {
        var v60 uint8 = 42
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v55
}
