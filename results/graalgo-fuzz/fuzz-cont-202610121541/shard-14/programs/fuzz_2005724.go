package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
func genFn0(p0 uint8, p1 uint) uint32 {
    if false {
        v5 := GenS1{F0: 0, F1: 0, F2: 472}
        if (354 >= 1) {
            var v6 int32 = 42
            v7 := []int{-811, 749, -9223372036854775808}
            if true {
                v6 = 562
            }
            _ = v6
            _ = v7
        } else {
            var v8 int64 = 9223372036854775807
            _ = v8
        }
        _ = v5
    }
    v9 := GenS1{F0: 42, F1: 0, F2: 1}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    var v1 uint64 = 57
    var v2 uint = 255
    var v3 float32 = -2.25
    var v4 int32 = 2147483647
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v1, v4, v2)
        switch "hello" {
        case "foo bar":
            v3 = (3.14159 - v3)
        case "\n\t":
            var v11 uint16 = 42
            _ = v11
        }
        switch "foo bar" {
        case "hello":
            if (v3 < (v3 + v3)) {
                var v12 int32 = (v4 - (v4 & 1))
                _ = v12
            } else {
                if (!true) {
                    if true {
                        var v13 int8 = -52
                        for fi14 := 0; fi14 < 4; fi14++ {
                            var v15 int16 = 1
                            var v16 int = (fi10 - fi14)
                            v13 += (v13 * v13)
                            v2 -= v2
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint32 = 533
                        _ = v13
                        _ = v17
                    }
                    switch 128 {
                    case 909:
                        var v18 string = "hello"
                        _ = v18
                    case 997:
                        if true {
                            switch fi10 {
                            case 128:
                                v19 := GenS1{F0: v1, F1: v1, F2: 0}
                                _ = v19
                            default:
                                switch ("hello" + "\n\t") {
                                case "foo bar":
                                    for fi20 := 0; fi20 < 7; fi20++ {
                                        v3 -= ((v3 / ((float32((3.14159 / float64(v3))) + (float32(3.14159) / v3)) * float32(float64(1.5)))) + (3.14159 + ((v3 * (v3 - v3)) - v3)))
                                        v21 := GenS2{F0: fi20, F1: -89, F2: genFn0(239, v2)}
                                        v2 = (v2 | v2)
                                        _ = fi20
                                        _ = v21
                                    }
                                default:
                                    fmt.Println(v4)
                                }
                            }
                            var v22 bool = false
                            v23 := []int{2147483647}
                            _ = v22
                            _ = v23
                        }
                    case 127:
                        if (114 >= 255) {
                            if ((v4 > (v4 * (v4 * v4))) || true) {
                                v1 = v1
                                v4 -= 42
                            }
                            if (577 > 111) {
                                if (-622 <= (fi10 / ((1) | 1))) {
                                    if (float64(v3) > (0.0 + (-2.25 - float64(v3)))) {
                                        v24 := []int{int(v2)}
                                        if len(v24) > 0 {
                                            v24[0] = fi10
                                        }
                                        _ = v24
                                    } else {
                                        v25 := GenS1{F0: (v1 / (((v1 | 32)) | 1)), F1: v1, F2: 255}
                                        _ = v25
                                    }
                                }
                            }
                        } else {
                            for fi26 := 0; fi26 < 11; fi26++ {
                                fmt.Println(v1, fi10, fi26)
                                _ = fi26
                            }
                            for fi27 := 0; fi27 < 12; fi27++ {
                                fmt.Println(v1, fi27, v4)
                                var v28 float32 = v3
                                _ = fi27
                                _ = v28
                            }
                        }
                    }
                    var v29 int = fi10
                    _ = v29
                } else {
                    v3 -= 1.5
                    if true {
                        v30 := []int{fi10, (-(fi10)), fi10}
                        switch (-(fi10)) {
                        case -436:
                            if len(v30) > 0 {
                                v30[0] = (fi10 ^ 884)
                            }
                        case 4294967295:
                            if (!true) {
                                var v31 int64 = 1
                                if (false && true) {
                                    v2 += v2
                                } else {
                                    v32 := GenS3{F0: 255, F1: 901, F2: (v4 % (((^(v4))) | 1))}
                                    if len(v30) > 0 {
                                        v30[0] = (fi10 / ((9223372036854775807) | 1))
                                    }
                                    _ = v32
                                }
                                _ = v31
                            }
                        case -559:
                            if (-2.25 > v3) {
                                switch -2147483648 {
                                case -824:
                                    v3 -= v3
                                case 9223372036854775807:
                                    var v33 bool = false
                                    _ = v33
                                case -494:
                                    v4 = 42
                                }
                            } else {
                                fmt.Println(v3, fi10, v2)
                            }
                        }
                        var v34 bool = false
                        _ = v30
                        _ = v34
                    }
                    v35 := GenS3{F0: 2, F1: 255, F2: v4}
                    _ = v35
                }
            }
        case "":
            v3 -= float32(-2.25)
        case "a":
            v36 := GenS1{F0: v1, F1: (v1 % (((v1 >> ((1) & 63))) | 1)), F2: 567}
            _ = v36
        }
        _ = fi10
    }
    v37 := GenS3{F0: 2, F1: 265, F2: v4}
    fmt.Println(v2, v4)
    for fi38 := 0; fi38 < 20; fi38++ {
        for fi39 := 0; fi39 < 15; fi39++ {
            v4 += int32(fi39)
            _ = fi39
        }
        var v40 float32 = v3
        _ = fi38
        _ = v40
    }
    v4 = 1
    for fi41 := 0; fi41 < 3; fi41++ {
        v2 -= 865
        switch -410 {
        case 4294967295:
            switch "hello" {
            case "foo bar":
                var v42 int = 423
                _ = v42
            case "a":
                switch fi41 {
                case -54:
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v44 := GenS4{F0: (fi43 % ((1) | 1)), F1: 42, F2: ("hello" + ((("foo bar" + "\n\t") + "hello") + ""))}
                        _ = fi43
                        _ = v44
                    }
                case 127:
                    var v45 uint64 = (v1 % ((v1) | 1))
                    _ = v45
                }
            case "":
                if (("" + ("foo bar" + "hello")) != "") {
                    switch 321 {
                    case -128:
                        v46 := GenS2{F0: fi41, F1: -1, F2: 42}
                        _ = v46
                    }
                    v47 := []int{fi41}
                    _ = v47
                } else {
                    v48 := GenS1{F0: 290, F1: 4294967295, F2: 65535}
                    v4 = 768
                    _ = v48
                }
            default:
                for fi49 := 0; fi49 < 1; fi49++ {
                    v50 := []int{fi41}
                    _ = fi49
                    _ = v50
                }
            }
        case 32767:
            for fi51 := 0; fi51 < 4; fi51++ {
                v52 := []int{(-(fi51)), -218}
                var v53 float32 = v3
                _ = fi51
                _ = v52
                _ = v53
            }
        }
        _ = fi41
    }
    var v54 uint16 = 1
    var v55 uint16 = 1
    var v56 uint32 = 2147483648
    if false {
        var v57 float64 = 0.1
        _ = v57
    } else {
        for fi58 := 0; fi58 < 7; fi58++ {
            fmt.Println(v3, v3, v55)
            v59 := GenS1{F0: v1, F1: v1, F2: v54}
            _ = fi58
            _ = v59
        }
        if false {
            v54 = (v55 / (((^(v55))) | 1))
            v60 := GenS4{F0: 265, F1: 2, F2: "hello"}
            var v61 float32 = (v3 + ((v3 * float32(((3.14159 + -2.25) / 1.0))) / ((v3 - 0.1) / (0.1 / (float32(float64(v3)) - float32(float64((v3 * float32(float64(0.1))))))))))
            _ = v60
            _ = v61
        } else {
            v2 = v2
        }
    }
    var v62 uint16 = (^(v54))
    for fi63 := 0; fi63 < 6; fi63++ {
        var v64 int64 = int64(v1)
        _ = fi63
        _ = v64
    }
    v1 = uint64(v55)
    fmt.Println(v56, v2)
    v65 := GenS2{F0: 4294967295, F1: -89, F2: (v56 >> (((v2 % (((v2 * uint(v62))) | 1))) & 63))}
    fmt.Println(v55)
    var v66 int16 = -642
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v62:", v62)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v54
    _ = v55
    _ = v56
    _ = v62
    _ = v65
    _ = v66
}
