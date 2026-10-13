package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int32; F1 uint8; F2 string }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int64; F1 string }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int8; F1 uint32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int32; F1 uint }
type GenS13 struct { F0 int32; F1 int8; F2 string }
func genFn0(p0 int16) int64 {
    var v5 bool = (0.1 < 0.0)
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := GenS1{F0: 1}
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS2{F0: (("\n\t" + ("a" + "")) + "foo bar")}
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    if (0 >= 1) {
        p0 = p0
        for fi10 := 0; fi10 < 4; fi10++ {
            fmt.Println(v5, fi10)
            switch "foo bar" {
            case "":
                var v11 int = 32768
                _ = v11
            case "a":
                var v12 uint16 = 255
                _ = v12
            case "\n\t":
                if v5 {
                    v5 = (true && (!v5))
                    var v13 int64 = 42
                    v14 := []int{127, (^(fi10))}
                    _ = v13
                    _ = v14
                }
            }
            v5 = (-1000 >= 471)
            _ = fi10
        }
        if v5 {
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := GenS2{F0: ""}
                    v18 := GenS2{F0: (("hello" + "") + "foo bar")}
                    v19 := []int{(fi16 + -9223372036854775808), (fi15 + 32768)}
                    v20 := GenS2{F0: (("" + ("" + "hello")) + "")}
                    switch "a" {
                    case "a":
                        var v21 float64 = 0.0
                        _ = v21
                    case "hello":
                        v22 := GenS3{F0: 0}
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi15
            }
            v23 := GenS4{F0: 42}
            v5 = (("a" + "\n\t") <= ("" + "hello"))
            _ = v23
        }
    } else {
        v5 = (2147483647 <= -9223372036854775808)
        v24 := []int{}
        var v25 bool = v5
        _ = v24
        _ = v25
    }
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int16 = -1
    var v2 uint = 1
    v3 := []uint32{2, 602}
    var v4 uint8 = 127
    fmt.Println(v1)
    switch -981 {
    case 32767:
        switch 42 {
        case 0:
            v26 := GenS3{F0: 65535}
            _ = v26
        default:
            var v27 string = "a"
            _ = v27
        }
    }
    var v28 uint32 = 2
    if (int64(v28) >= -1) {
        var v29 uint64 = 227
        var v30 int32 = 1
        _ = v29
        _ = v30
    }
    switch "a" {
    case "\n\t":
        for fi31 := 0; fi31 < 17; fi31++ {
            if len(v3) > 0 {
                v3[0] = v28
            }
            _ = fi31
        }
    case "a":
        v28 = (v28 * v28)
    case "foo bar":
        if (!(-58 >= -128)) {
            var v32 int = 9223372036854775807
            for fi33 := 0; fi33 < 17; fi33++ {
                for fi34 := 0; fi34 < 18; fi34++ {
                    v28 += 4294967295
                    fmt.Println(fi34)
                    v35 := GenS1{F0: 100}
                    v3 = append(v3, 2)
                    v3 = append(v3, v28)
                    _ = fi34
                    _ = v35
                }
                if false {
                    var v36 uint8 = v4
                    _ = v36
                }
                var v37 float64 = float64(0.0)
                _ = fi33
                _ = v37
            }
            _ = v32
        } else {
            v4 += 128
            v2 = (v2 & (v2 ^ v2))
        }
    }
    switch 32768 {
    case 32768:
        v38 := GenS5{F0: 18446744073709551615, F1: 1}
        _ = v38
    case 0:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if true {
        v4 -= (^(v4))
        v39 := GenS6{F0: 0, F1: (v4 / (((v4 | v4)) | 1)), F2: (("a" + ("hello" + ("" + "\n\t"))) + "hello")}
        _ = v39
    } else {
        if false {
            var v40 float32 = -2.25
            _ = v40
        } else {
            var v41 int64 = 0
            _ = v41
        }
        v42 := GenS6{F0: -869, F1: (v4 ^ v4), F2: "a"}
        if (((1.5 / 1.0) + (3.14159 * ((3.14159 * -2.25) + 0.1))) > -2.25) {
            v28 = uint32(v4)
        }
        _ = v42
    }
    if false {
        if len(v3) > 0 {
            v3[0] = (v28 << ((v2) & 63))
        }
        switch 128 {
        case -2147483648:
            v4 = (v4 / ((uint8(v28)) | 1))
        case 32767:
            for fi43 := 0; fi43 < 3; fi43++ {
                v2 = v2
                v3 = append(v3, (v28 / (((v28 >> ((255) & 63))) | 1)))
                for fi44 := 0; fi44 < 16; fi44++ {
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v1 = v1
                        }
                    } else {
                        var v45 bool = (!false)
                        var v46 uint32 = 65535
                        _ = v45
                        _ = v46
                    }
                    _ = fi44
                }
                _ = fi43
            }
        default:
            v2 -= v2
        }
    }
    v1 += (v1 & (v1 ^ (v1 - v1)))
    var v47 uint8 = (v4 + 42)
    v28 -= (^(v28))
    for fi48 := 0; fi48 < 20; fi48++ {
        switch ("a" + ("\n\t" + ("foo bar" + "foo bar"))) {
        case "a":
            if true {
                if false {
                    switch "a" {
                    case "":
                        fmt.Println(v28, v47, v47)
                    default:
                        v49 := GenS7{F0: v28, F1: (v4 / (((v4 << ((v2) & 63))) | 1))}
                        _ = v49
                    }
                    v3 = append(v3, (v28 << ((v2) & 63)))
                    v50 := GenS8{F0: 273, F1: "a"}
                    _ = v50
                } else {
                    if false {
                        if (!(true || true)) {
                            var v51 uint8 = v47
                            v51 = v4
                            _ = v51
                        } else {
                            var v52 uint = 770
                            v2 = (v2 | 65535)
                            for fi53 := 0; fi53 < 15; fi53++ {
                                v54 := GenS9{F0: 1, F1: (v52 + v52)}
                                _ = fi53
                                _ = v54
                            }
                            _ = v52
                        }
                        v1 += (v1 - v1)
                    }
                }
                if false {
                    v55 := GenS10{F0: -1, F1: 437}
                    if (841 <= 65535) {
                        for fi56 := 0; fi56 < 18; fi56++ {
                            v3 = append(v3, (v28 ^ (^(v28))))
                            _ = fi56
                        }
                        switch (fi48 * (fi48 * -9223372036854775808)) {
                        case -483:
                            var v57 float64 = ((1e10 + (-2.25 + 0.1)) * 0.1)
                            _ = v57
                        case -9223372036854775808:
                            var v58 uint = (^(v2))
                            _ = v58
                        }
                    }
                    var v59 float32 = (float32(3.14159) * 3.14159)
                    _ = v55
                    _ = v59
                }
                if len(v3) > 0 {
                    v3[0] = uint32(v4)
                }
            } else {
                var v60 float32 = ((0.1 - float32(((((0.1 / 1.5) * 0.0) / 1e10) * ((float64(-2.25) + 1.5) / 3.14159)))) + (1.5 - (float32(float64(0.0)) / 0.1)))
                for fi61 := 0; fi61 < 5; fi61++ {
                    v60 = -2.25
                    v62 := GenS11{F0: false}
                    _ = fi61
                    _ = v62
                }
                v2 = 277
                _ = v60
            }
        default:
            if true {
                for fi63 := 0; fi63 < 18; fi63++ {
                    v28 -= v28
                    v64 := GenS12{F0: 2147483647, F1: 680}
                    v28 -= (v28 + (^(v28)))
                    _ = fi63
                    _ = v64
                }
                v47 += (v4 | (v47 + v4))
                _ = len(v3)
            }
        }
        switch "foo bar" {
        case "hello":
            v65 := GenS13{F0: 2147483647, F1: 0, F2: ""}
            _ = v65
        case "":
            v4 += (v47 + (v47 & (v4 * 128)))
        default:
            v47 -= 0
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v47
}
