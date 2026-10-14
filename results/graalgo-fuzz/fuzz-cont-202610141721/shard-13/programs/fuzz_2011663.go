package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 int32; F2 bool }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]bool)
    v3 := []int{-80, 0, 32767, 128}
    v4 := []string{(("hello" + "a") + "a"), "\n\t", "hello", ""}
    v5 := make(map[string]int)
    var v6 int32 = -1
    for fi7 := 0; fi7 < 8; fi7++ {
        v6 = v6
        var v8 int16 = 32767
        _ = fi7
        _ = v8
    }
    fmt.Println(v6)
    if false {
        if true {
            var v9 uint32 = 2147483647
            _ = v9
        } else {
            v6 += -512
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 float32 = float32(0.0)
                _ = fi10
                _ = v11
            }
        }
        switch (((("\n\t" + "hello") + (("foo bar" + ("a" + "hello")) + (("hello" + (("foo bar" + ("" + (("a" + (("foo bar" + "\n\t") + "a")) + "foo bar"))) + "a")) + "hello"))) + "\n\t") + "foo bar") {
        case "foo bar":
            fmt.Println(v6)
        default:
            v6 -= 0
        }
        if (228 <= 71) {
            v6 -= (-(v6))
            v6 -= 2147483647
        }
    }
    v6 -= 344
    delete(v1, 0)
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v6)
        var v13 int32 = (v6 + (v6 & v6))
        if (!(4294967295 > 253)) {
            v13 = v6
            switch fi12 {
            case -1:
                v6 = v13
            case -9223372036854775808:
                fmt.Println(v6, v13, fi12)
            default:
                fmt.Println(v13)
            }
        } else {
            fmt.Println(fi12, v6)
            var v14 uint32 = 2147483647
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    v4 = append(v4, "a")
    var v15 bool = true
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 int8 = 0
        fmt.Println(v17)
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: 101, F2: v6}
    var v19 int32 = (v6 * -648)
    switch -2147483648 {
    case 32768:
        if (("foo bar" + ("\n\t" + ("" + ((("" + "hello") + ("a" + "hello")) + "")))) != (("a" + "foo bar") + ("a" + ("a" + ("foo bar" + "hello"))))) {
            v1[0] = -1
            v19 = 1
        }
    case 127:
        switch "\n\t" {
        case "":
            v6 = 660
        default:
            if true {
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 float32 = -2.25
                    var v22 float32 = (((float32(-2.25) + ((((((3.14159 / ((v21 * (v21 - 1e10)) - v21)) * v21) * (v21 - v21)) - v21) / 1.0) - float32((1.5 / 1.0)))) / 0.1) / v21)
                    var v23 int32 = -477
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                delete(v5, "")
                _ = len(v1)
            } else {
                var v24 int8 = 0
                if true {
                    var v25 int16 = -596
                    fmt.Println(v19, v15, v6)
                    v15 = (!(false || true))
                    _ = v25
                }
                v19 = (v6 << ((4294967295) & 63))
                _ = v24
            }
        }
    }
    if (v15 || v15) {
        v26 := GenS1{F0: int8(v6), F1: 128, F2: v6}
        fmt.Println(v15, v6)
        _ = v26
    } else {
        if v15 {
            v27 := GenS1{F0: 42, F1: 234, F2: 1}
            if (-128 < -1) {
                var v28 uint32 = 4294967295
                if v15 {
                    switch "a" {
                    case "":
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v6 += (v19 % (((v6 ^ v19)) | 1))
                            _ = fi29
                        }
                    default:
                        if v15 {
                            var v30 int32 = v6
                            _ = v30
                        } else {
                            fmt.Println(v6, v19)
                            var v31 float64 = 3.14159
                            _ = v31
                        }
                    }
                } else {
                    for fi32 := 0; fi32 < 15; fi32++ {
                        v33 := GenS1{F0: 0, F1: 127, F2: (v6 / (((v6 / (((^(v6))) | 1))) | 1))}
                        v19 = v19
                        _ = fi32
                        _ = v33
                    }
                    var v34 float64 = 3.14159
                    var v35 int8 = 0
                    _ = v34
                    _ = v35
                }
                v19 += v19
                _ = v28
            } else {
                var v36 float64 = 0.1
                var v37 int32 = v6
                _ = v36
                _ = v37
            }
            _ = v27
        }
        v19 += (v6 ^ 2147483647)
        if v15 {
            v38 := GenS1{F0: 0, F1: 2, F2: (^(v19))}
            if (0 != 1) {
                var v39 bool = ((v19 + 1) != v6)
                var v40 int = 1
                _ = v39
                _ = v40
            } else {
                switch -2147483648 {
                case -1:
                    if (941 == 0) {
                        if true {
                            v15 = false
                            v6 = 0
                            v41 := GenS2{F0: 560, F1: 2147483647, F2: v15}
                            _ = v41
                        } else {
                            v19 -= (v19 - v19)
                            delete(v1, 1)
                            switch 42 {
                            case 32767:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        }
                        v42 := GenS2{F0: -338, F1: (v19 << ((437) & 63)), F2: (!(1 > 0))}
                        for fi43 := 0; fi43 < 12; fi43++ {
                            var v44 uint16 = 718
                            v45 := GenS2{F0: (^(fi43)), F1: v19, F2: (2 <= 950)}
                            var v46 int32 = v6
                            _ = fi43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        _ = v42
                    }
                case 127:
                    var v47 string = ""
                    _ = v47
                case 0:
                    if false {
                        v48 := GenS2{F0: -517, F1: v19, F2: v15}
                        _ = v48
                    } else {
                        v2[2147483648] = false
                        if (true || (100 > 0)) {
                            v15 = (-2.25 <= 0.0)
                            v19 = v19
                        }
                        switch "hello" {
                        case "":
                            for fi49 := 0; fi49 < 18; fi49++ {
                                v15 = v15
                                for fi50 := 0; fi50 < 20; fi50++ {
                                    fmt.Println(fi49, v19, v15)
                                    v15 = v15
                                    _ = fi50
                                }
                                var v51 float32 = ((float32(float64(3.14159)) + (0.1 * (0.0 / 1.0))) - 0.0)
                                _ = fi49
                                _ = v51
                            }
                        case "\n\t":
                            v52 := GenS3{F0: 1, F1: 120, F2: 2147483647}
                            _ = v52
                        case "hello":
                            for fi53 := 0; fi53 < 7; fi53++ {
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    var v55 uint = 42
                                    _ = fi54
                                    _ = v55
                                }
                                var v56 uint16 = 1
                                v57 := GenS1{F0: 127, F1: 255, F2: (v6 - v6)}
                                _ = fi53
                                _ = v56
                                _ = v57
                            }
                        default:
                            var v58 int16 = 506
                            _ = v58
                        }
                    }
                default:
                    var v59 string = "\n\t"
                    _ = v59
                }
                v15 = v15
            }
            _ = v38
        }
    }
    v15 = (0.1 == 3.14159)
    var v60 float32 = -2.25
    fmt.Println(v15, v15)
    for fi61 := 0; fi61 < 13; fi61++ {
        switch fi61 {
        case 42:
            v62 := GenS4{F0: "foo bar"}
            _ = v62
        }
        var v63 string = (("hello" + "hello") + ("foo bar" + (("a" + ("\n\t" + "a")) + "hello")))
        if (1.5 < -2.25) {
            fmt.Println(v6, v60)
            switch fi61 {
            case 1:
                fmt.Println(fi61, v63)
            case 127:
                for fi64 := 0; fi64 < 17; fi64++ {
                    v60 = v60
                    _ = fi64
                }
            default:
                v65 := GenS1{F0: -1, F1: 217, F2: (v19 | (v6 & (v19 >> ((0) & 63))))}
                _ = v65
            }
        }
        _ = fi61
        _ = v63
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
