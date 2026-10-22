package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool; F1 bool; F2 bool }
type GenS8 struct { F0 string; F1 int8; F2 uint32 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int }

func main() {
    v1 := []uint32{1, 4294967295, 65535, 2147483647}
    v2 := []uint32{950, 251, 557}
    v3 := []int{}
    var v4 float32 = ((1e10 - 0.1) - (3.14159 + 1.5))
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    if false {
        if (((!(!true)) || true) || false) {
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
    }
    v5[2147483647] = true
    for fi7 := 0; fi7 < 16; fi7++ {
        switch 42 {
        case 128:
            var v8 string = "a"
            _ = v8
        case 28:
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(fi9)
                var v10 float32 = ((v4 - v4) - (1.5 - v4))
                v10 = 0.1
                _ = fi9
                _ = v10
            }
        case 936:
            v4 -= float32(1e10)
        }
        v4 = v4
        _ = fi7
    }
    var v11 float64 = (((3.14159 + -2.25) * 0.0) / 1e10)
    var v12 uint64 = 18446744073709551615
    if (true || ((("foo bar" + "") + "hello") >= "")) {
        fmt.Println(v12, v12)
        switch "" {
        case "hello":
            v12 = (v12 % ((v12) | 1))
        case "foo bar":
            var v13 string = ("hello" + "foo bar")
            _ = v13
        case "":
            switch ("a" + "a") {
            case "hello":
                v4 -= v4
            default:
                var v14 int32 = 869
                _ = v14
            }
        default:
            v15 := v5[2]
            _ = v15
        }
    }
    if (!true) {
        v11 = v11
        v16 := GenS1{F0: (("\n\t" + (((((("" + ("" + "\n\t")) + "foo bar") + "") + "a") + ("hello" + "foo bar")) + "hello")) > (("hello" + "foo bar") + "\n\t")), F1: -2147483648, F2: -2147483648}
        _ = v16
    }
    switch (("foo bar" + "hello") + "a") {
    case "foo bar":
        if (!(42 <= 127)) {
            for fi17 := 0; fi17 < 10; fi17++ {
                v11 = v11
                v4 = float32(((1e10 - (0.1 - v11)) * v11))
                fmt.Println(fi17, v4)
                _ = fi17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := GenS2{F0: 455, F1: false, F2: 2}
                v20 := GenS2{F0: 0, F1: (float32(v11) >= ((float32(-2.25) - 1.5) + v4)), F2: 1}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            var v21 int8 = 127
            if false {
                v22 := GenS2{F0: 911, F1: (18446744073709551615 >= 650), F2: 2}
                for fi23 := 0; fi23 < 1; fi23++ {
                    v24 := GenS3{F0: (-(v21)), F1: 42}
                    switch fi23 {
                    case 0:
                        if (-2.25 > v4) {
                            switch (("\n\t" + (("a" + "\n\t") + (("\n\t" + ((("hello" + ((("\n\t" + "hello") + "foo bar") + ("a" + ("foo bar" + ("a" + "\n\t"))))) + ("hello" + "\n\t")) + "foo bar")) + "hello"))) + "hello") {
                            case "":
                                if (!true) {
                                    v25 := v6["\n\t"]
                                    var v26 uint64 = (^(v12))
                                    for fi27 := 0; fi27 < 15; fi27++ {
                                        var v28 bool = (534 >= 0)
                                        _ = fi27
                                        _ = v28
                                    }
                                    _ = v25
                                    _ = v26
                                } else {
                                    for fi29 := 0; fi29 < 11; fi29++ {
                                        var v30 int32 = 375
                                        switch ("" + "\n\t") {
                                        case "a":
                                            var v31 uint32 = 0
                                            _ = v31
                                        case "hello":
                                            var v32 uint8 = 42
                                            _ = v32
                                        case "":
                                            _ = len(v1)
                                        default:
                                            v12 = 1
                                        }
                                        v33 := GenS4{F0: ""}
                                        var v34 uint32 = 1
                                        if ((42 == int16(v30)) || true) {
                                            var v35 int64 = -1
                                            switch "" {
                                            case "\n\t":
                                                var v36 float32 = 1e10
                                                v30 = -2147483648
                                                _ = v36
                                            case "foo bar":
                                                v34 = v34
                                            }
                                            v12 = (^(v12))
                                            _ = v35
                                        }
                                        _ = fi29
                                        _ = v30
                                        _ = v33
                                        _ = v34
                                    }
                                }
                            default:
                                v37 := GenS5{F0: 597}
                                _ = v37
                            }
                            v12 -= (v12 ^ (^(v12)))
                        }
                    case 679:
                        if false {
                            if ((!(!true)) && false) {
                                var v38 float64 = v11
                                v39 := GenS4{F0: ("foo bar" + ("\n\t" + ("a" + "\n\t")))}
                                _ = v38
                                _ = v39
                            }
                            var v40 uint64 = 675
                            fmt.Println(v12, v12)
                            _ = v40
                        }
                    case -429:
                        v41 := GenS6{F0: (v12 + (v12 / ((v12) | 1)))}
                        _ = v41
                    default:
                        var v42 uint64 = v12
                        _ = v42
                    }
                    if (false && (520 >= 204)) {
                        switch (fi23 * (fi23 << ((236) & 63))) {
                        case 42:
                            var v43 bool = ("hello" <= "foo bar")
                            _ = v43
                        case 9223372036854775807:
                            switch "foo bar" {
                            case "foo bar":
                                if (v4 == (v4 - float32(v11))) {
                                    v4 += v4
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                _ = v22
            } else {
                v6[("\n\t" + ("foo bar" + "a"))] = false
            }
            if ((v12 / ((v12) | 1)) != (^(v12))) {
                switch "foo bar" {
                case "foo bar":
                    if true {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v44 := GenS7{F0: false, F1: (!(!false)), F2: (true || true)}
                        v11 += ((float64(float32(v11)) - ((v11 + v11) / (v11 + float64(float32((v11 - v11)))))) + 1.5)
                        _ = v44
                    } else {
                        v11 -= 1e10
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                        for fi45 := 0; fi45 < 9; fi45++ {
                            v12 += v12
                            for fi46 := 0; fi46 < 7; fi46++ {
                                switch fi46 {
                                case 127:
                                    var v47 float64 = ((float64(v4) - 0.0) + v11)
                                    _ = v47
                                default:
                                    var v48 uint64 = (v12 - (v12 % (((^(v12))) | 1)))
                                    _ = v48
                                }
                                v4 += 0.0
                                var v49 uint16 = 255
                                _ = fi46
                                _ = v49
                            }
                            _ = fi45
                        }
                    }
                case "hello":
                    if (0 > 255) {
                        var v50 string = ""
                        for fi51 := 0; fi51 < 17; fi51++ {
                            fmt.Println(v11, v11, v4)
                            if false {
                                v52 := v5[uint32(v12)]
                                _ = v52
                            }
                            _ = fi51
                        }
                        switch 346 {
                        case 32768:
                            if (0 != 1) {
                                v11 = 0.1
                                switch v50 {
                                case "":
                                    fmt.Println(v50, v50, v50)
                                }
                                v50 = "foo bar"
                            } else {
                                v50 = ((((v50 + ("hello" + (v50 + ("\n\t" + "")))) + "\n\t") + ("foo bar" + v50)) + ("a" + v50))
                            }
                        }
                        _ = v50
                    } else {
                        switch 42 {
                        case 9223372036854775807:
                            var v53 float32 = -2.25
                            _ = v53
                        case 32768:
                            delete(v5, 785)
                        default:
                            for fi54 := 0; fi54 < 5; fi54++ {
                                var v55 uint16 = 872
                                _ = fi54
                                _ = v55
                            }
                        }
                    }
                default:
                    v12 = v12
                }
                var v56 int = -1
                _ = v56
            } else {
                v1 = append(v1, 928)
                if (!(-1 > 32767)) {
                    v11 -= v11
                    var v57 int16 = -32768
                    _ = v57
                } else {
                    for fi58 := 0; fi58 < 14; fi58++ {
                        var v59 int16 = 0
                        _ = fi58
                        _ = v59
                    }
                    fmt.Println(v21, v21, v12)
                    v60 := GenS8{F0: ((("a" + ("\n\t" + "\n\t")) + ("foo bar" + "a")) + ("foo bar" + "a")), F1: v21, F2: 4294967295}
                    _ = v60
                }
            }
            _ = v21
        }
    }
    v61 := GenS9{F0: 1, F1: 2}
    if true {
        var v62 uint16 = 2
        _ = v62
    }
    if (!(2 == 157)) {
        v11 = 0.0
        fmt.Println(v12, v12, v11)
        var v63 bool = ((42 <= -53) && false)
        _ = v63
    } else {
        v64 := GenS10{F0: 32767}
        var v65 float64 = v11
        delete(v6, "")
        _ = v64
        _ = v65
    }
    switch "hello" {
    case "hello":
        v12 = (v12 - (v12 | 777))
    case "foo bar":
        for fi66 := 0; fi66 < 19; fi66++ {
            v11 -= (v11 - v11)
            _ = fi66
        }
    default:
        v4 = (1e10 * 0.1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
}
