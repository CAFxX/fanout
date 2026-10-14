package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16; F2 int16 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 string; F1 string; F2 int64 }
type GenS9 struct { F0 int64; F1 int32; F2 int }
func genFn0() uint {
    var v5 uint8 = 54
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint64 = 317
        _ = fi6
        _ = v7
    }
    v5 = 128
    var v8 string = ("a" + (("hello" + "\n\t") + ""))
    _ = v5
    _ = v8
    return 255
}
func genFn1(p0 uint8) uint64 {
    var v9 float64 = -2.25
    p0 = p0
    switch ("a" + "hello") {
    case "":
        p0 += 216
    case "a":
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := []int{(fi11 - (fi10 + fi10)), fi10, fi11}
                v13 := GenS1{F0: (!(0 >= -450)), F1: 42}
                if (((1e10 * 0.0) + (3.14159 * -2.25)) > 0.1) {
                    _ = len(v12)
                    switch (fi10 << ((genFn0()) & 63)) {
                    case -9223372036854775808:
                        fmt.Println(fi10)
                    case -21:
                        fmt.Println(v9, p0)
                    case 9223372036854775807:
                        p0 = (p0 << ((18446744073709551615) & 63))
                    }
                }
                fmt.Println(fi10)
                v9 += (v9 * (v9 + v9))
                v14 := GenS1{F0: (0 > 42), F1: -1}
                if (-2147483648 < 0) {
                    v15 := GenS1{F0: true, F1: -48}
                    v16 := GenS2{F0: 283}
                    p0 += (^(p0))
                    _ = v15
                    _ = v16
                }
                v9 = 1.5
                v9 = (-2.25 - (0.1 / 0.1))
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v17 int16 = -370
            for fi18 := 0; fi18 < 1; fi18++ {
                if ((905 <= 321) || true) {
                    v19 := GenS3{F0: (fi18 % (((fi18 & (fi10 * fi18))) | 1)), F1: 540, F2: v17}
                    _ = v19
                } else {
                    switch (fi18 % (((^(fi18))) | 1)) {
                    case -308:
                        var v20 string = ("\n\t" + "hello")
                        _ = v20
                    case -1:
                        v21 := GenS4{F0: (fi10 + fi10), F1: 4294967295}
                        _ = v21
                    case -998:
                        switch "a" {
                        case "a":
                            if ((("foo bar" + "a") != "hello") || (!true)) {
                                switch fi10 {
                                case 127:
                                    v22 := []int{32768, (fi18 + (fi10 / ((127) | 1))), (fi10 % ((127) | 1))}
                                    _ = v22
                                case 132:
                                    var v23 float32 = 1.5
                                    _ = v23
                                case -2147483648:
                                    if ((v9 / v9) != ((1e10 + v9) / 3.14159)) {
                                        var v24 uint64 = 1
                                        fmt.Println(fi18)
                                        _ = v24
                                    }
                                }
                            } else {
                                var v25 uint64 = 0
                                fmt.Println(v17, v25, fi18)
                                _ = v25
                            }
                        default:
                            v26 := GenS3{F0: 432, F1: (v17 % ((0) | 1)), F2: -32768}
                            _ = v26
                        }
                    default:
                        if (((v17 / ((v17) | 1)) < (-(v17))) && ((((-2.25 <= v9) && true) || (!(42 >= -646))) && (!(true && (!(!(true || (182 > 44)))))))) {
                            v17 += 926
                            v27 := []int{-741}
                            _ = v27
                        }
                    }
                    v17 = (v17 >> ((65535) & 63))
                }
                v9 = 0.1
                fmt.Println(v9, fi18)
                v9 += v9
                _ = fi18
            }
            _ = fi10
            _ = v17
        }
    default:
        v28 := GenS4{F0: -73, F1: 4294967295}
        _ = v28
    }
    _ = p0
    _ = v9
    return 122
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 42
    var v3 string = ""
    v4 := make(map[string]bool)
    switch v3 {
    case "hello":
        if false {
            if true {
                fmt.Println(v3)
                if true {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v30 := GenS5{F0: -591, F1: 477}
                        _ = fi29
                        _ = v30
                    }
                    switch 4294967295 {
                    case 127:
                        var v31 float64 = 0.1
                        _ = v31
                    case -643:
                        if false {
                            v3 = v3
                            v2 += v2
                        } else {
                            v32 := GenS2{F0: 2}
                            for fi33 := 0; fi33 < 12; fi33++ {
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    v2 = v2
                                    fmt.Println(fi34, fi34)
                                    if (!false) {
                                        if (false && true) {
                                            switch (v3 + (v3 + v3)) {
                                            case "a":
                                                switch "hello" {
                                                case "hello":
                                                    var v35 float64 = float64(1.5)
                                                    _ = v35
                                                case "a":
                                                    fmt.Println(fi34)
                                                case "foo bar":
                                                    if len(v1) > 0 {
                                                        v1[0] = 0
                                                    }
                                                }
                                            }
                                            if true {
                                                delete(v4, "")
                                                v3 = ((("a" + v3) + (("foo bar" + "foo bar") + v3)) + (v3 + "foo bar"))
                                                v2 -= 9223372036854775807
                                                var v36 int8 = 50
                                                v37 := GenS6{F0: 4294967295, F1: 881, F2: 0}
                                                delete(v4, ("\n\t" + v3))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v36
                                                _ = v37
                                            }
                                            if (true || ((v2 - (v2 << ((255) & 63))) == (v2 - (v2 >> ((255) & 63))))) {
                                                fmt.Println(v2)
                                                var v38 uint8 = 255
                                                _ = len(v4)
                                                _ = v38
                                            }
                                        }
                                    }
                                    v39 := GenS7{F0: ((v3 + (v3 + ("" + v3))) + v3), F1: "foo bar"}
                                    _ = fi34
                                    _ = v39
                                }
                                _ = fi33
                            }
                            _ = v32
                        }
                    case 0:
                        v2 = (v2 & 299)
                    }
                    fmt.Println(v3, v2)
                }
            } else {
                _ = len(v4)
                var v40 int64 = (v2 & (v2 << ((4294967295) & 63)))
                _ = v40
            }
        }
    }
    var v41 uint = genFn0()
    var v42 float64 = (3.14159 - (3.14159 * 1e10))
    v41 += (v41 % (((v41 * genFn0())) | 1))
    v42 = v42
    if (false || false) {
        var v43 uint16 = 65535
        for fi44 := 0; fi44 < 3; fi44++ {
            v2 = v2
            _ = fi44
        }
        _ = v43
    } else {
        v45 := GenS8{F0: "foo bar", F1: "\n\t", F2: v2}
        var v46 int = -443
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 string = "hello"
        fmt.Println(v2, v48)
        var v49 int8 = 0
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        var v51 uint8 = 42
        var v52 string = v3
        if true {
            v42 += (0.0 * (v42 - v42))
            if len(v1) > 0 {
                v1[0] = 2
            }
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    if (551 > 2) {
        v2 += (v2 & v2)
        var v53 int8 = 127
        _ = v53
    }
    v3 = "foo bar"
    if len(v1) > 0 {
        v1[0] = 65535
    }
    _ = len(v4)
    v54 := GenS9{F0: 0, F1: 42, F2: 872}
    var v55 int8 = 111
    fmt.Println(v42, v3)
    var v56 uint16 = 65535
    var v57 uint = v41
    for fi58 := 0; fi58 < 4; fi58++ {
        var v59 string = "a"
        switch fi58 {
        case -394:
            var v60 int8 = (v55 * v55)
            _ = v60
        case 32768:
            var v61 float64 = (0.1 - (v42 / v42))
            _ = v61
        default:
            v2 -= v2
        }
        _ = fi58
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
