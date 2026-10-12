package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 int64 }
func genFn0() uint64 {
    var v5 bool = false
    if v5 {
        fmt.Println(v5)
        v6 := GenS1{F0: 854, F1: 9223372036854775807, F2: 100}
        _ = v6
    }
    switch "" {
    case "\n\t":
        var v7 uint32 = 0
        _ = v7
    case "":
        v5 = v5
    case "foo bar":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint64 = 897
            _ = fi8
            _ = v9
        }
    }
    _ = v5
    return 42
}

func main() {
    var v1 bool = (!(!(1 < 1)))
    var v2 float64 = -2.25
    var v3 int = 1
    var v4 float64 = -2.25
    var v10 string = ""
    for fi11 := 0; fi11 < 5; fi11++ {
        v10 = v10
        var v12 string = ("a" + v10)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 float32 = 1.5
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: (-500 != -792), F1: 128, F2: 0}
    for fi16 := 0; fi16 < 3; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            v10 = "\n\t"
            v18 := []int{fi16, fi16}
            v19 := GenS3{F0: 1, F1: -591}
            _ = fi17
            _ = v18
            _ = v19
        }
        switch 1 {
        case 32767:
            switch fi16 {
            case 9223372036854775807:
                v2 += float64(0.0)
            case 220:
                if v1 {
                    v20 := []int{fi16}
                    v4 = ((v2 / v2) - (v2 / float64(3.14159)))
                    _ = v20
                } else {
                    var v21 string = v10
                    _ = v21
                }
            case 4294967295:
                v22 := []int{(fi16 ^ fi16)}
                _ = v22
            }
        case -9223372036854775808:
            switch v10 {
            case "foo bar":
                v23 := []int{v3, v3, (fi16 | (fi16 + fi16))}
                _ = v23
            case "a":
                v4 += float64(-2.25)
            default:
                v10 = "a"
            }
        default:
            var v24 uint64 = 865
            _ = v24
        }
        var v25 int64 = 0
        _ = fi16
        _ = v25
    }
    if (255 != 0) {
        v3 -= (^(v3))
        fmt.Println(v4, v10)
        for fi26 := 0; fi26 < 16; fi26++ {
            v27 := GenS4{F0: 4294967295, F1: 127}
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS3{F0: 255, F1: 835}
    if (42 == 42) {
        v29 := []int{v3, v3}
        var v30 uint32 = 715
        if len(v29) > 0 {
            v29[0] = (v3 & (v3 % (((v3 << ((2) & 63))) | 1)))
        }
        _ = v29
        _ = v30
    }
    v4 -= v4
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 string = ((v10 + v10) + ((("" + "a") + "hello") + "foo bar"))
        _ = fi31
        _ = v32
    }
    if (v1 || v1) {
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 int = v3
            _ = fi33
            _ = v34
        }
        v35 := []int{v3, v3}
        _ = v35
    } else {
        switch "a" {
        case "a":
            for fi36 := 0; fi36 < 17; fi36++ {
                for fi37 := 0; fi37 < 7; fi37++ {
                    var v38 int32 = 42
                    var v39 bool = ((v1 || (1 <= 2)) && false)
                    v40 := GenS1{F0: 1, F1: 1, F2: int8(v38)}
                    v41 := []int{(fi37 | (fi37 | (fi37 - fi37))), 1, v3}
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                fmt.Println(v2)
                for fi42 := 0; fi42 < 20; fi42++ {
                    v3 -= -2147483648
                    fmt.Println(v3, fi42, v2)
                    v43 := GenS3{F0: 18446744073709551615, F1: -230}
                    _ = fi42
                    _ = v43
                }
                _ = fi36
            }
        case "hello":
            for fi44 := 0; fi44 < 18; fi44++ {
                fmt.Println(v10, v2, v1)
                v45 := GenS5{F0: 0}
                var v46 uint16 = 944
                _ = fi44
                _ = v45
                _ = v46
            }
        }
        v10 = "\n\t"
    }
    v47 := []int{-690}
    if (v10 > (v10 + ("a" + v10))) {
        var v48 int64 = 42
        var v49 float32 = 1.5
        _ = v48
        _ = v49
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            v47 = append(v47, (v3 * (fi50 | -2147483648)))
            if false {
                v51 := GenS4{F0: 0, F1: 255}
                v2 = v4
                _ = v51
            } else {
                v52 := GenS2{F0: (!(-2.25 != 1.5)), F1: 127, F2: -128}
                var v53 uint = 4294967295
                v53 += 2
                _ = v52
                _ = v53
            }
            _ = fi50
        }
        v10 = ("hello" + v10)
        for fi54 := 0; fi54 < 12; fi54++ {
            for fi55 := 0; fi55 < 2; fi55++ {
                v56 := GenS6{F0: 859}
                _ = fi55
                _ = v56
            }
            if v1 {
                v1 = false
                fmt.Println(v10, v4)
            }
            var v57 uint32 = 4294967295
            _ = fi54
            _ = v57
        }
    }
    v58 := GenS7{F0: 42, F1: 1}
    if (v1 && (v1 && (false && (2147483648 < 42)))) {
        for fi59 := 0; fi59 < 17; fi59++ {
            if (((0.0 + v2) + v4) < 1e10) {
                if false {
                    var v60 int = (-(v3))
                    v47 = append(v47, 373)
                    _ = v60
                } else {
                    v47 = append(v47, 2147483647)
                    var v61 uint64 = 0
                    var v62 int16 = -1
                    _ = v61
                    _ = v62
                }
                var v63 string = v10
                _ = v63
            }
            fmt.Println(fi59, v2)
            _ = fi59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v47
    _ = v58
}
