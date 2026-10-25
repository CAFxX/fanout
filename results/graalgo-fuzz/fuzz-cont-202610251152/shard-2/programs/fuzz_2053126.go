package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 uint8; F1 int8 }
func genFn0() uint32 {
    var v7 uint8 = 1
    v8 := []int{2147483647, 9223372036854775807}
    v7 = (^(v7))
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        v7 = 87
        _ = fi9
    }
    _ = v7
    _ = v8
    return 747
}
func genFn1(p0 int64, p1 uint32, p2 uint8) int8 {
    for fi10 := 0; fi10 < 19; fi10++ {
        if (32767 >= 65535) {
            p0 = p0
        } else {
            var v11 uint8 = (p2 & p2)
            p0 = (-(p0))
            _ = v11
        }
        fmt.Println(p0)
        var v12 int32 = int32(fi10)
        _ = fi10
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            p0 = 42
            _ = fi14
        }
        var v15 int64 = (p0 ^ 42)
        _ = fi13
        _ = v15
    }
    v16 := []int{9223372036854775807, 4294967295, -286, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 float32 = ((((3.14159 + float32(-2.25)) + (1e10 + (((((1e10 + ((1e10 - 0.0) + 1.5)) + float32((3.14159 / -2.25))) / 1e10) + 0.0) + -2.25))) - 3.14159) * 1.5)
    var v4 float64 = float64(v3)
    var v5 uint64 = 354
    v6 := make(map[string]string)
    var v17 int16 = 1
    switch "hello" {
    case "\n\t":
        if ((-207 == 512) && true) {
            fmt.Println(v17, v4, v5)
            switch 414 {
            case -756:
                var v18 uint32 = 916
                _ = v18
            case 42:
                v19 := v1[42]
                _ = v19
            case -115:
                var v20 int64 = 82
                _ = v20
            }
        }
    case "a":
        delete(v6, ("\n\t" + ""))
    case "hello":
        v21 := GenS1{F0: 32768, F1: 2, F2: (^(v5))}
        _ = v21
    default:
        v5 -= (^(v5))
    }
    if false {
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(v5, fi22)
            var v23 uint16 = 42
            _ = fi22
            _ = v23
        }
        v24 := GenS2{F0: 1, F1: 6, F2: -75}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 13; fi25++ {
            v4 += float64(((v3 * 3.14159) * v3))
            fmt.Println(fi25, v17)
            _ = fi25
        }
    }
    if (18446744073709551615 != 1) {
        if (true && true) {
            if true {
                if false {
                    v17 = (-(v17))
                    var v26 string = (("foo bar" + "a") + "foo bar")
                    _ = v26
                } else {
                    v5 += (v5 << ((896) & 63))
                }
                fmt.Println(v17, v17)
                if (true || (v17 > (v17 - (^(v17))))) {
                    for fi27 := 0; fi27 < 20; fi27++ {
                        if (false && (-2147483648 >= 454)) {
                            switch "foo bar" {
                            case "":
                                var v28 uint = 69
                                _ = v28
                            case "hello":
                                var v29 float32 = v3
                                _ = v29
                            case "\n\t":
                                if (42 <= 4294967295) {
                                    fmt.Println(v17, v17, v4)
                                    var v30 bool = (1 == 42)
                                    if (v30 && true) {
                                        v31 := GenS3{F0: fi27}
                                        var v32 uint8 = 202
                                        v33 := []int{fi27, fi27, fi27}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    }
                                    _ = v30
                                }
                            }
                        } else {
                            var v34 uint = 65535
                            v17 = 1000
                            _ = v34
                        }
                        v5 = v5
                        switch ("a" + "foo bar") {
                        case "a":
                            switch fi27 {
                            case 0:
                                if (v5 > (v5 ^ v5)) {
                                    v35 := GenS4{F0: -532}
                                    for fi36 := 0; fi36 < 15; fi36++ {
                                        if (false && false) {
                                            var v37 int32 = -163
                                            v6[("\n\t" + ("\n\t" + "a"))] = "foo bar"
                                            var v38 int32 = v37
                                            var v39 uint = 2
                                            var v40 float32 = 1.5
                                            v41 := GenS1{F0: fi36, F1: 4294967295, F2: (v5 ^ (v5 % ((2) | 1)))}
                                            v42 := GenS5{F0: 9223372036854775807, F1: 0}
                                            var v43 uint32 = 2
                                            v44 := []int{32768, 4294967295, fi36}
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                        v17 = v17
                                        if true {
                                            v3 -= (0.0 + (1e10 + v3))
                                            v45 := []int{fi36, 128, 323, (fi27 - (^(fi27)))}
                                            _ = v45
                                        }
                                        _ = fi36
                                    }
                                    _ = v35
                                }
                            }
                        case "foo bar":
                            for fi46 := 0; fi46 < 16; fi46++ {
                                switch (fi27 % (((fi46 / ((fi27) | 1))) | 1)) {
                                case -2147483648:
                                    fmt.Println(v5, fi46)
                                }
                                v4 += 0.1
                                _ = fi46
                            }
                        case "hello":
                            var v47 float64 = 3.14159
                            _ = v47
                        }
                        _ = fi27
                    }
                    for fi48 := 0; fi48 < 19; fi48++ {
                        if (0 > 18446744073709551615) {
                            var v49 int = (fi48 >> ((4294967295) & 63))
                            if (!(float64(v3) >= (v4 / 1.0))) {
                                switch "" {
                                case "foo bar":
                                    var v50 int16 = (v17 | -1)
                                    _ = v50
                                case "hello":
                                    v17 -= (v17 / (((v17 - (v17 << ((65535) & 63)))) | 1))
                                case "a":
                                    v5 = 4294967295
                                default:
                                    var v51 uint64 = 0
                                    _ = v51
                                }
                                var v52 int64 = 0
                                _ = v52
                            }
                            _ = v49
                        } else {
                            v17 = -32768
                            fmt.Println(fi48, fi48, v5)
                            var v53 int64 = 1
                            _ = v53
                        }
                        fmt.Println(v17, v17)
                        var v54 uint = 413
                        _ = fi48
                        _ = v54
                    }
                }
            } else {
                var v55 uint16 = 912
                v4 = v4
                _ = v55
            }
        }
        var v56 uint64 = (v5 * v5)
        switch ("" + "") {
        case "foo bar":
            var v57 bool = (!true)
            _ = v57
        }
        _ = v56
    }
    var v58 uint8 = 2
    v59 := GenS6{F0: -1, F1: 4294967295}
    v4 += (3.14159 * 3.14159)
    if false {
        if (422 < 1) {
            if true {
                v17 = v17
                if (-383 <= 4294967295) {
                    v60 := GenS7{F0: 127, F1: (v17 | 1)}
                    v61 := []int{}
                    _ = v60
                    _ = v61
                }
            }
        }
    }
    v3 += (v3 - 3.14159)
    for fi62 := 0; fi62 < 1; fi62++ {
        if (9223372036854775807 == 952) {
            v63 := GenS4{F0: -985}
            v58 -= 35
            var v64 float64 = ((3.14159 / v4) * -2.25)
            _ = v63
            _ = v64
        }
        if (2 >= 0) {
            var v65 uint16 = 917
            for fi66 := 0; fi66 < 4; fi66++ {
                var v67 int64 = 174
                v17 = (-(v17))
                v3 = float32(v4)
                _ = fi66
                _ = v67
            }
            switch 2147483647 {
            case 4294967295:
                fmt.Println(fi62)
            case -9223372036854775808:
                for fi68 := 0; fi68 < 8; fi68++ {
                    var v69 float32 = -2.25
                    v70 := GenS8{F0: (v58 | v58), F1: -128}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
            }
            _ = v65
        }
        _ = fi62
    }
    fmt.Println(v5, v3)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
