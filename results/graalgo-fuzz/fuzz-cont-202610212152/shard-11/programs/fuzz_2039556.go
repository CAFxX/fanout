package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 uint8; F2 bool }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 string; F1 uint64; F2 int }
type GenS6 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64) uint8 {
    v5 := GenS1{F0: -1}
    switch ("a" + ("\n\t" + "foo bar")) {
    case "foo bar":
        v6 := []int{-379, -1, 128, 1}
        _ = v6
    case "a":
        if true {
            fmt.Println(p0)
            v7 := []int{128, 0, 9223372036854775807, 9223372036854775807}
            var v8 bool = false
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint = 42
                v11 := []int{fi9}
                fmt.Println(p0)
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                p0 -= (p0 << ((0) & 63))
                _ = fi12
            }
        }
    default:
        p0 += -9223372036854775808
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 float64 = 0.0
    var v4 uint32 = 136
    fmt.Println(v3, v4)
    v13 := GenS1{F0: 42}
    if (1 <= 32768) {
        v14 := GenS2{F0: (v4 + v4), F1: 2, F2: -757}
        var v15 uint64 = 6
        _ = v14
        _ = v15
    } else {
        v4 += 0
    }
    delete(v2, ("hello" + "a"))
    v16 := GenS3{F0: false, F1: 1, F2: true}
    switch "\n\t" {
    case "hello":
        var v17 string = "a"
        _ = v17
    case "foo bar":
        v3 -= ((v3 * (v3 * 0.0)) - v3)
    case "":
        for fi18 := 0; fi18 < 9; fi18++ {
            v3 = (((v3 * (float64(0.1) - float64(3.14159))) * v3) / 3.14159)
            _ = fi18
        }
    }
    if (-1 != 127) {
        if true {
            v19 := []int{32768, -903}
            _ = v19
        }
        switch "foo bar" {
        case "\n\t":
            v3 = v3
        case "":
            for fi20 := 0; fi20 < 10; fi20++ {
                for fi21 := 0; fi21 < 14; fi21++ {
                    if false {
                        var v22 uint16 = 1
                        _ = v22
                    } else {
                        switch (("" + ((("" + "\n\t") + "a") + "hello")) + "") {
                        case "foo bar":
                            if (288 >= 0) {
                                var v23 uint32 = 663
                                fmt.Println(fi20, v23, v4)
                                var v24 int16 = -824
                                _ = v23
                                _ = v24
                            } else {
                                var v25 float32 = -2.25
                                if false {
                                    if ("a" != ("a" + ("\n\t" + ("foo bar" + ("\n\t" + ""))))) {
                                        var v26 int16 = 42
                                        _ = v26
                                    }
                                    if true {
                                        v27 := GenS4{F0: 127, F1: 2}
                                        v3 = (0.0 / float64(v25))
                                        _ = v27
                                    }
                                    var v28 uint64 = 2
                                    _ = v28
                                }
                                _ = v25
                            }
                        }
                    }
                    fmt.Println(fi21, v4)
                    _ = fi21
                }
                v29 := []int{int(v4), (fi20 * 2147483647), (fi20 ^ -475), (-(fi20))}
                var v30 float32 = ((1.5 + float32(float64(3.14159))) * float32(v3))
                _ = fi20
                _ = v29
                _ = v30
            }
        case "foo bar":
            if true {
                v4 = (^(v4))
                switch ("a" + ("foo bar" + "hello")) {
                case "foo bar":
                    v4 -= (^(v4))
                case "":
                    v3 = (float64(((((0.1 / 0.1) - 0.1) / 3.14159) + 1.5)) - 3.14159)
                case "hello":
                    v3 = (v3 * 3.14159)
                default:
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v32 := []int{fi31, (fi31 + (fi31 + 0)), (^(fi31)), (fi31 & (fi31 << ((729) & 63)))}
                        _ = fi31
                        _ = v32
                    }
                }
                v4 = 431
            } else {
                fmt.Println(v3)
                v33 := []int{}
                _ = v33
            }
        default:
            v34 := GenS5{F0: "foo bar", F1: 1, F2: 1}
            _ = v34
        }
        var v35 int8 = 100
        _ = v35
    }
    v3 = (v3 * v3)
    switch "foo bar" {
    case "a":
        v36 := GenS2{F0: (v4 - (v4 + v4)), F1: 556, F2: 758}
        _ = v36
    case "":
        switch 4294967295 {
        case -9223372036854775808:
            for fi37 := 0; fi37 < 7; fi37++ {
                if true {
                    var v38 float64 = float64(3.14159)
                    v39 := GenS3{F0: true, F1: 28, F2: (true || true)}
                    v40 := GenS2{F0: (v4 >> ((411) & 63)), F1: 42, F2: 0}
                    _ = v38
                    _ = v39
                    _ = v40
                } else {
                    _ = len(v2)
                }
                for fi41 := 0; fi41 < 14; fi41++ {
                    var v42 int8 = -128
                    _ = fi41
                    _ = v42
                }
                _ = fi37
            }
        case 962:
            var v43 int16 = 1000
            _ = v43
        case -196:
            if (true || (v3 != 1.5)) {
                for fi44 := 0; fi44 < 11; fi44++ {
                    fmt.Println(fi44, v4)
                    switch ((("a" + (("hello" + "\n\t") + "hello")) + "") + "\n\t") {
                    case "\n\t":
                        switch fi44 {
                        case 127:
                            fmt.Println(v4, v4)
                        default:
                            v45 := []int{(fi44 % ((fi44) | 1)), (fi44 | (^(fi44))), 4294967295, (fi44 * fi44)}
                            _ = v45
                        }
                    case "foo bar":
                        var v46 float64 = v3
                        _ = v46
                    default:
                        if false {
                            v47 := v2["\n\t"]
                            for fi48 := 0; fi48 < 15; fi48++ {
                                v49 := []int{-879}
                                _ = fi48
                                _ = v49
                            }
                            _ = v47
                        }
                    }
                    v4 = 65535
                    _ = fi44
                }
                v50 := GenS4{F0: 84, F1: genFn0(42)}
                fmt.Println(v4)
                _ = v50
            }
        default:
            var v51 uint = 4294967295
            _ = v51
        }
    default:
        v4 = v4
    }
    var v52 uint8 = 42
    if false {
        switch ("hello" + ((("" + "foo bar") + "\n\t") + ("a" + "foo bar"))) {
        case "a":
            var v53 uint16 = 2
            _ = v53
        }
        v52 -= (v52 >> ((18446744073709551615) & 63))
        fmt.Println(v3, v4)
    } else {
        var v54 uint8 = genFn0(-1)
        v55 := []int{355, 32767, -922, 128}
        var v56 uint16 = 2
        _ = v54
        _ = v55
        _ = v56
    }
    fmt.Println(v4, v4)
    if true {
        var v57 uint64 = 0
        if false {
            var v58 uint64 = 18446744073709551615
            v2["a"] = "foo bar"
            _ = v58
        }
        fmt.Println(v52)
        _ = v57
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        var v60 uint = 18446744073709551615
        if true {
            switch 32767 {
            case 4294967295:
                if false {
                    v61 := []int{fi59, -2147483648, (fi59 | fi59)}
                    v60 = v60
                    _ = v61
                } else {
                    fmt.Println(v60, fi59)
                    v62 := GenS6{F0: 329, F1: int16(fi59)}
                    _ = v62
                }
            default:
                v60 -= v60
            }
            fmt.Println(v60, v4)
            for fi63 := 0; fi63 < 12; fi63++ {
                if false {
                    v64 := []int{128, fi63, fi63, -1}
                    v52 = v52
                    _ = v64
                }
                _ = fi63
            }
        } else {
            var v65 uint = (v60 / ((v60) | 1))
            _ = v65
        }
        _ = fi59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v52
}
