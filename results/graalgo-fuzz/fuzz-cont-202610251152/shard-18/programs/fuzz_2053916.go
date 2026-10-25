package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int; F1 int32; F2 uint64 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{65535, 4294967295, 967}
    var v3 uint64 = 2
    var v4 int64 = 1
    var v5 uint16 = 920
    var v6 float64 = (0.0 - (0.1 / (0.1 - 0.0)))
    var v7 int = -9223372036854775808
    fmt.Println(v3)
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            v4 -= 191
            v7 = (v7 + fi8)
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            if (((1.5 * -2.25) + 3.14159) < (((0.0 - 0.0) * 0.0) / (0.0 / (float32((v6 - 1.5)) / (-2.25 + -2.25))))) {
                var v10 string = (("a" + "foo bar") + (("a" + ("a" + "\n\t")) + "a"))
                v11 := GenS1{F0: 42, F1: v5, F2: (v5 / ((2) | 1))}
                v12 := v1[0]
                if false {
                    v7 = -500
                    var v13 uint = 970
                    _ = v13
                } else {
                    v7 -= 956
                    _ = len(v2)
                    var v14 int32 = 2147483647
                    _ = v14
                }
                _ = v10
                _ = v11
                _ = v12
            } else {
                fmt.Println(v3, v6)
                v15 := GenS1{F0: -32768, F1: v5, F2: v5}
                v16 := GenS1{F0: 32767, F1: (^(v5)), F2: (v5 & (v5 & v5))}
                _ = v15
                _ = v16
            }
            var v17 uint16 = (^(v5))
            _ = fi9
            _ = v17
        }
        _ = fi8
    }
    if true {
        v1[v7] = -9223372036854775808
        v4 -= v4
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            switch int(v3) {
            case 128:
                _ = len(v1)
            case -9223372036854775808:
                var v19 uint32 = 42
                _ = v19
            default:
                v20 := GenS2{F0: (-(fi18)), F1: ("a" + (("" + "a") + "a"))}
                _ = v20
            }
            if (130 > 68) {
                if true {
                    switch "foo bar" {
                    case "hello":
                        for fi21 := 0; fi21 < 5; fi21++ {
                            fmt.Println(v3, fi18)
                            if (152 <= uint8(v5)) {
                                var v22 float32 = -2.25
                                v3 -= v3
                                v23 := GenS3{F0: (!((-2.25 * v22) > 1.5)), F1: 42, F2: 174}
                                _ = v22
                                _ = v23
                            }
                            fmt.Println(v4)
                            var v24 uint = 18446744073709551615
                            v3 -= (v3 - uint64(v5))
                            fmt.Println(fi18)
                            _ = fi21
                            _ = v24
                        }
                    case "":
                        var v25 uint16 = (v5 | (^(v5)))
                        _ = v25
                    }
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        switch int(v4) {
                        case 128:
                            var v26 uint64 = v3
                            _ = v26
                        case -214:
                            fmt.Println(v3, v3, v6)
                        case -996:
                            for fi27 := 0; fi27 < 18; fi27++ {
                                var v28 string = ("hello" + ("\n\t" + ("a" + "a")))
                                v7 += -9223372036854775808
                                _ = fi27
                                _ = v28
                            }
                        default:
                            _ = len(v2)
                        }
                    default:
                        v29 := GenS4{F0: v5}
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 12; fi30++ {
                    switch ("" + "foo bar") {
                    case "a":
                        if false {
                            _ = len(v1)
                            switch (("" + "a") + "foo bar") {
                            case "foo bar":
                                switch (("" + ("a" + "foo bar")) + "hello") {
                                case "hello":
                                    if false {
                                        v7 -= (fi30 ^ 1)
                                        v31 := GenS2{F0: fi18, F1: "hello"}
                                        if false {
                                            switch ("\n\t" + "\n\t") {
                                            case "":
                                                v3 += (v3 % (((^(v3))) | 1))
                                            case "hello":
                                                if false {
                                                    var v32 string = ""
                                                    fmt.Println(v3)
                                                    _ = v32
                                                }
                                            }
                                            v33 := GenS2{F0: (fi30 << ((1) & 63)), F1: "a"}
                                            v4 += (v4 << ((2) & 63))
                                            var v34 uint64 = uint64(v5)
                                            _ = v33
                                            _ = v34
                                        }
                                        _ = v31
                                    } else {
                                        var v35 uint16 = 34
                                        v36 := GenS5{F0: "", F1: -1}
                                        switch (v7 - fi18) {
                                        case -191:
                                            var v37 int64 = 9223372036854775807
                                            _ = v37
                                        case 32768:
                                            switch ("" + ("hello" + "foo bar")) {
                                            case "foo bar":
                                                _ = len(v2)
                                            case "a":
                                                v38 := GenS4{F0: 255}
                                                _ = v38
                                            case "\n\t":
                                                v3 = 425
                                            }
                                        case 101:
                                            v39 := GenS6{F0: fi18, F1: 2147483647, F2: (v3 / ((18446744073709551615) | 1))}
                                            _ = v39
                                        }
                                        delete(v1, (fi18 << ((42) & 63)))
                                        _ = v35
                                        _ = v36
                                    }
                                case "\n\t":
                                    if ((true || (0 < 128)) || (v7 != v7)) {
                                        if len(v2) > 0 {
                                            v2[0] = uint32(v3)
                                        }
                                        v6 += 1.5
                                    }
                                }
                            case "a":
                                v5 += 735
                            default:
                                v40 := v1[fi18]
                                v41 := v1[(fi18 ^ fi30)]
                                _ = v40
                                _ = v41
                            }
                            var v42 uint32 = 0
                            _ = v42
                        }
                    default:
                        v5 += 255
                    }
                    _ = fi30
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v43 int64 = (v4 ^ int64(v7))
            _ = fi18
            _ = v43
        }
        if (133 >= 0) {
            if false {
                v3 = (v3 + (v3 & 1))
            } else {
                if true {
                    for fi44 := 0; fi44 < 1; fi44++ {
                        fmt.Println(v5, v3, v3)
                        _ = fi44
                    }
                    v4 += -212
                } else {
                    v45 := GenS7{F0: v7, F1: ("a" + ("hello" + (("" + "") + "foo bar")))}
                    var v46 float64 = (float64(float32(v6)) + ((float64(-2.25) * (((float64(-2.25) * (v6 + v6)) + (1e10 + float64(0.0))) + -2.25)) + (float64(3.14159) + 0.1)))
                    v3 -= (v3 ^ v3)
                    _ = v45
                    _ = v46
                }
                if len(v2) > 0 {
                    v2[0] = 30
                }
            }
            switch (v7 % ((-2147483648) | 1)) {
            case 2147483647:
                fmt.Println(v4, v5, v3)
            case 32768:
                switch "a" {
                case "\n\t":
                    fmt.Println(v5)
                }
            case -1:
                fmt.Println(v7)
            }
        } else {
            v47 := GenS5{F0: "", F1: 1000}
            fmt.Println(v6, v6, v4)
            v7 -= (v7 << ((2) & 63))
            _ = v47
        }
    }
    var v48 int64 = (v4 % (((v4 * v4)) | 1))
    v49 := GenS8{F0: -32768, F1: v5}
    for fi50 := 0; fi50 < 10; fi50++ {
        var v51 uint = 4294967295
        var v52 int32 = -983
        if (((("hello" + (("" + "hello") + (("" + "") + "\n\t"))) + "\n\t") + "foo bar") <= (((((("hello" + "hello") + "\n\t") + "") + "hello") + (("hello" + "hello") + ((("\n\t" + "a") + ("hello" + ("foo bar" + (("a" + "") + "foo bar")))) + ("foo bar" + ("a" + "a"))))) + "foo bar")) {
            for fi53 := 0; fi53 < 6; fi53++ {
                if ("\n\t" < "") {
                    var v54 int32 = (v52 % (((v52 + v52)) | 1))
                    switch (fi50 / ((int(v4)) | 1)) {
                    case 4294967295:
                        v51 = v51
                    }
                    _ = v54
                }
                fmt.Println(fi50, fi50)
                var v55 int16 = 991
                _ = fi53
                _ = v55
            }
            for fi56 := 0; fi56 < 15; fi56++ {
                switch "" {
                case "":
                    var v57 string = "foo bar"
                    _ = v57
                case "hello":
                    v52 -= v52
                default:
                    var v58 int8 = 0
                    _ = v58
                }
                _ = fi56
            }
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    var v59 float32 = (-2.25 * 1.5)
    v60 := v1[(-(v7))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v7
}
