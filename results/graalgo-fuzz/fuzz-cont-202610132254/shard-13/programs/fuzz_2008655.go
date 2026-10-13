package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint8; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint32, p1 bool) uint8 {
    if ((922 != 18446744073709551615) && p1) {
        v6 := []int{9223372036854775807, 32768}
        p1 = true
        if p1 {
            if len(v6) > 0 {
                v6[0] = 128
            }
            var v7 uint64 = 42
            p0 += p0
            _ = v7
        }
        _ = v6
    } else {
        p0 -= 2
        v8 := GenS1{F0: 749, F1: 2, F2: 555}
        if (((3.14159 - 0.0) > 1.5) && (!(624 < 2))) {
            if p1 {
                p1 = true
            } else {
                switch "hello" {
                case "a":
                    for fi9 := 0; fi9 < 19; fi9++ {
                        fmt.Println(p1, fi9, fi9)
                        for fi10 := 0; fi10 < 15; fi10++ {
                            var v11 int8 = 42
                            v12 := GenS1{F0: 666, F1: 940, F2: 42}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 3; fi13++ {
                            if (!(((fi9 >> ((709) & 63)) <= fi9) && (p1 || p1))) {
                                fmt.Println(fi13, fi9, p1)
                                v14 := GenS2{F0: 42, F1: 255, F2: ("" + ("hello" + "\n\t"))}
                                p1 = p1
                                v15 := []int{(fi13 ^ fi9), 42}
                                _ = v14
                                _ = v15
                            } else {
                                if p1 {
                                    fmt.Println(p1, p1, p1)
                                }
                            }
                            if false {
                                if (!(0.1 > (0.0 / (0.0 - float32(((1.5 - (1e10 / 1e10)) + (-2.25 - 3.14159))))))) {
                                    fmt.Println(fi13, p1)
                                } else {
                                    var v16 bool = p1
                                    v17 := GenS1{F0: 0, F1: 1, F2: 226}
                                    _ = v16
                                    _ = v17
                                }
                            }
                            _ = fi13
                        }
                        _ = fi9
                    }
                default:
                    for fi18 := 0; fi18 < 6; fi18++ {
                        switch ("hello" + "") {
                        case "a":
                            for fi19 := 0; fi19 < 7; fi19++ {
                                p1 = (100 <= 42)
                                var v20 uint = 18446744073709551615
                                _ = fi19
                                _ = v20
                            }
                        case "hello":
                            for fi21 := 0; fi21 < 15; fi21++ {
                                p0 = p0
                                v22 := GenS3{F0: false}
                                _ = fi21
                                _ = v22
                            }
                        default:
                            if (-98 > 99) {
                                v23 := []int{(fi18 - fi18), (fi18 + (fi18 >> ((42) & 63))), fi18, fi18}
                                p1 = (0 < 1)
                                _ = v23
                            } else {
                                var v24 uint8 = 255
                                v25 := GenS2{F0: uint64(p0), F1: (v24 + v24), F2: (("hello" + ("hello" + (((("hello" + "hello") + ("foo bar" + "\n\t")) + "hello") + ((("a" + "a") + "\n\t") + "a")))) + "a")}
                                v26 := GenS4{F0: 4294967295, F1: 452}
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        }
                        var v27 int16 = 1
                        var v28 uint = 42
                        _ = fi18
                        _ = v27
                        _ = v28
                    }
                }
            }
            switch "a" {
            case "":
                for fi29 := 0; fi29 < 18; fi29++ {
                    fmt.Println(p0, p0)
                    v30 := GenS4{F0: 18446744073709551615, F1: 1}
                    v31 := []int{873, (fi29 << ((4294967295) & 63)), fi29}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case "hello":
                var v32 int = 1
                _ = v32
            default:
                v33 := []int{-9223372036854775808, 217, int(p0)}
                _ = v33
            }
            p1 = true
        }
        _ = v8
    }
    var v34 bool = false
    if (2 > 37) {
        p1 = p1
        switch "foo bar" {
        case "\n\t":
            var v35 uint = 42
            _ = v35
        case "a":
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 int64 = 9223372036854775807
                _ = fi36
                _ = v37
            }
        }
    }
    _ = p0
    _ = p1
    _ = v34
    return 91
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    var v3 float32 = (1e10 + 1e10)
    var v4 string = ("foo bar" + "foo bar")
    v5 := []uint32{65535, 2147483648}
    if false {
        v4 = ((v4 + (((v4 + v4) + v4) + (v4 + "\n\t"))) + v4)
        _ = len(v5)
    } else {
        v3 -= v3
    }
    v5 = append(v5, 42)
    if (!(!false)) {
        v38 := GenS5{F0: 432}
        v2 -= (v2 + (v2 + (-(v2))))
        _ = v38
    } else {
        var v39 int8 = int8(v2)
        _ = v39
    }
    switch "a" {
    case "a":
        fmt.Println(v2, v3, v2)
    case "\n\t":
        var v40 float64 = 1.5
        _ = v40
    }
    v2 = (v2 & -1)
    var v41 int = v2
    v2 = (v41 % (((v2 * v41)) | 1))
    v42 := GenS6{F0: 0}
    var v43 int64 = -562
    if ((float64((v3 * float32(1e10))) + ((float64(1.5) * 3.14159) + 1.5)) < (0.0 * -2.25)) {
        fmt.Println(v43, v4)
        if true {
            for fi44 := 0; fi44 < 8; fi44++ {
                v2 = (v41 / ((9223372036854775807) | 1))
                v45 := GenS2{F0: 2, F1: 128, F2: "\n\t"}
                if true {
                    if (1 <= 87) {
                        var v46 uint16 = 65535
                        var v47 int32 = -2147483648
                        switch (v41 ^ fi44) {
                        case 1:
                            for fi48 := 0; fi48 < 15; fi48++ {
                                v49 := GenS2{F0: 0, F1: 244, F2: (v4 + ("foo bar" + v4))}
                                v47 = v47
                                if true {
                                    switch (fi44 + fi44) {
                                    case -2147483648:
                                        fmt.Println(v43)
                                    case -785:
                                        v50 := GenS3{F0: false}
                                        _ = v50
                                    case 1:
                                        v4 = v4
                                    default:
                                        _ = len(v5)
                                    }
                                    v51 := GenS7{F0: (v43 - (v43 + -540))}
                                    _ = v51
                                }
                                var v52 int8 = -9
                                if true {
                                    var v53 uint8 = 127
                                    v3 = v3
                                    var v54 int = fi48
                                    v4 = "a"
                                    _ = v53
                                    _ = v54
                                }
                                _ = fi48
                                _ = v49
                                _ = v52
                            }
                        case -2147483648:
                            if len(v5) > 0 {
                                v5[0] = 145
                            }
                        case 32768:
                            _ = len(v1)
                        }
                        _ = v46
                        _ = v47
                    } else {
                        if (4294967295 <= 2) {
                            var v55 string = v4
                            switch 508 {
                            case 32767:
                                v4 = (v4 + "hello")
                            case -21:
                                var v56 uint32 = 65535
                                _ = v56
                            default:
                                v57 := GenS8{F0: 0, F1: 36}
                                _ = v57
                            }
                            _ = v55
                        } else {
                            var v58 uint16 = 2
                            for fi59 := 0; fi59 < 2; fi59++ {
                                var v60 uint8 = 7
                                switch v41 {
                                case 0:
                                    v61 := GenS9{F0: v60, F1: -2147483648, F2: 18446744073709551615}
                                    _ = v61
                                default:
                                    _ = len(v1)
                                }
                                v62 := GenS3{F0: true}
                                _ = fi59
                                _ = v60
                                _ = v62
                            }
                            var v63 bool = true
                            _ = v58
                            _ = v63
                        }
                    }
                    for fi64 := 0; fi64 < 18; fi64++ {
                        switch "foo bar" {
                        case "a":
                            v43 -= -1
                        }
                        switch ((v4 + "foo bar") + ("\n\t" + "a")) {
                        case "a":
                            switch (^(v41)) {
                            case 0:
                                v65 := GenS3{F0: true}
                                _ = v65
                            case 2147483647:
                                if true {
                                    fmt.Println(v3, v4)
                                    v1[(-(v2))] = true
                                } else {
                                    v66 := v1[(^(v41))]
                                    _ = v66
                                }
                            }
                        case "foo bar":
                            v67 := GenS5{F0: 42}
                            _ = v67
                        case "\n\t":
                            if (273 < 2) {
                                _ = len(v1)
                                if (518 == 2) {
                                    fmt.Println(v3)
                                    if false {
                                        var v68 uint16 = 0
                                        v5 = append(v5, 4294967295)
                                        if (-1 > -755) {
                                            v69 := GenS1{F0: 18446744073709551615, F1: 703, F2: 255}
                                            v3 += (-2.25 - v3)
                                            _ = v69
                                        }
                                        _ = v68
                                    } else {
                                        switch (fi64 * fi44) {
                                        case -1:
                                            v1[(v2 + (-(v2)))] = false
                                        default:
                                            v2 = (v2 / ((v2) | 1))
                                        }
                                        v2 = fi44
                                    }
                                }
                                var v70 float32 = v3
                                _ = v70
                            }
                        default:
                            v3 = float32(0.1)
                        }
                        _ = fi64
                    }
                    v71 := GenS10{F0: v43}
                    _ = v71
                } else {
                    var v72 int8 = 0
                    for fi73 := 0; fi73 < 20; fi73++ {
                        v74 := v1[fi73]
                        fmt.Println(v2)
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                }
                _ = fi44
                _ = v45
            }
            var v75 int32 = 54
            _ = v75
        }
        v5 = append(v5, 2)
    } else {
        if len(v5) > 0 {
            v5[0] = 42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
