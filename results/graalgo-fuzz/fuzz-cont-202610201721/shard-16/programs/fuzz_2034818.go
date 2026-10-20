package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 string) int {
    switch ("foo bar" + p0) {
    case "hello":
        p0 = p0
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := []int{fi6, (fi6 / ((fi6) | 1)), 448, (fi6 * (fi6 * 42))}
            _ = fi6
            _ = v7
        }
    case "a":
        var v8 float32 = ((float32(3.14159) / 3.14159) - (1.5 - float32(3.14159)))
        _ = v8
    }
    p0 = "\n\t"
    v9 := GenS1{F0: 21, F1: -823}
    v10 := GenS1{F0: -127, F1: -2147483648}
    _ = p0
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 bool = true
    var v2 bool = (873 >= -9223372036854775808)
    v3 := make(map[string]int)
    var v4 uint32 = 42
    var v5 uint32 = v4
    if (!v1) {
        switch "" {
        case "hello":
            var v11 uint = 42
            _ = v11
        case "foo bar":
            v12 := v3["hello"]
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 4; fi13++ {
                v4 = v5
                for fi14 := 0; fi14 < 7; fi14++ {
                    v1 = v1
                    v1 = (65535 >= (v4 / ((v4) | 1)))
                    _ = fi14
                }
                _ = fi13
            }
        default:
            var v15 uint32 = 4294967295
            _ = v15
        }
        delete(v3, "\n\t")
    } else {
        if (v2 || (("foo bar" + "foo bar") != ("" + "foo bar"))) {
            switch "\n\t" {
            case "\n\t":
                v2 = (128 > uint8(v4))
            }
            switch "" {
            case "":
                switch "foo bar" {
                case "foo bar":
                    v16 := GenS2{F0: 1, F1: -2147483648, F2: 18446744073709551615}
                    _ = v16
                case "":
                    v3["a"] = 194
                case "hello":
                    var v17 int32 = 42
                    _ = v17
                default:
                    if ((377 > 212) || true) {
                        v3[""] = 202
                    } else {
                        switch "foo bar" {
                        case "hello":
                            if (v1 || ((v2 || true) || (581 <= 18446744073709551615))) {
                                v18 := []int{1, 9223372036854775807, -364, -9223372036854775808}
                                switch "a" {
                                case "a":
                                    var v19 bool = (v2 && (!(2 <= 4294967295)))
                                    _ = v19
                                case "foo bar":
                                    for fi20 := 0; fi20 < 6; fi20++ {
                                        switch (("hello" + ("foo bar" + "\n\t")) + "a") {
                                        case "":
                                            switch fi20 {
                                            case 918:
                                                v21 := GenS2{F0: 475, F1: 340, F2: 18446744073709551615}
                                                _ = v21
                                            case 96:
                                                for fi22 := 0; fi22 < 1; fi22++ {
                                                    v2 = (2 == 1)
                                                    var v23 float64 = float64(1.5)
                                                    if false {
                                                        switch "hello" {
                                                        case "a":
                                                            fmt.Println(fi22, fi22)
                                                            if ((-2.25 / -2.25) >= (-2.25 + 0.0)) {
                                                                v2 = (0 >= (v5 >> ((65535) & 63)))
                                                            } else {
                                                                if (!(v1 && (v4 < v4))) {
                                                                    var v24 uint16 = 1
                                                                    var v25 int64 = -1
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v26 uint = 0
                                                                        _ = v26
                                                                    default:
                                                                        var v27 uint8 = 128
                                                                        _ = v27
                                                                    }
                                                                    v1 = v1
                                                                    v23 -= float64(3.14159)
                                                                    _ = v24
                                                                    _ = v25
                                                                } else {
                                                                    v2 = (-61 <= 42)
                                                                    v18 = append(v18, fi20)
                                                                }
                                                                var v28 bool = false
                                                                _ = v28
                                                            }
                                                        case "":
                                                            delete(v3, "foo bar")
                                                        }
                                                        var v29 int8 = 89
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v30 := GenS3{F0: int32(v5), F1: 1}
                                                            _ = v30
                                                        case "\n\t":
                                                            v31 := GenS1{F0: (v29 * 100), F1: int32(fi22)}
                                                            _ = v31
                                                        case "a":
                                                            if false {
                                                                v23 -= v23
                                                            } else {
                                                                v32 := GenS2{F0: 425, F1: int32(fi20), F2: 4294967295}
                                                                _ = v32
                                                            }
                                                        default:
                                                            if ((v5 * 395) < 1) {
                                                                v33 := GenS4{F0: 552, F1: 283}
                                                                _ = v33
                                                            } else {
                                                                var v34 string = ("foo bar" + (("foo bar" + (("\n\t" + "") + "")) + "a"))
                                                                var v35 uint = 12
                                                                v36 := GenS1{F0: 0, F1: 2147483647}
                                                                _ = v34
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                        }
                                                        v2 = (3.14159 < 1.5)
                                                        _ = v29
                                                    }
                                                    _ = fi22
                                                    _ = v23
                                                }
                                            default:
                                                v37 := GenS3{F0: -70, F1: -1}
                                                _ = v37
                                            }
                                        }
                                        v1 = v2
                                        _ = fi20
                                    }
                                default:
                                    v2 = false
                                }
                                _ = v18
                            }
                        case "a":
                            var v38 bool = (-567 == 0)
                            _ = v38
                        case "":
                            v39 := GenS1{F0: -1, F1: 2147483647}
                            _ = v39
                        }
                        var v40 int16 = -760
                        if false {
                            if (v40 == (-(v40))) {
                                if true {
                                    var v41 int64 = 9223372036854775807
                                    var v42 uint16 = 42
                                    _ = v41
                                    _ = v42
                                }
                                v1 = false
                            }
                            v40 -= v40
                            for fi43 := 0; fi43 < 8; fi43++ {
                                v44 := GenS2{F0: 926, F1: -189, F2: 0}
                                v45 := []int{}
                                _ = fi43
                                _ = v44
                                _ = v45
                            }
                        } else {
                            v2 = false
                        }
                        _ = v40
                    }
                }
            default:
                var v46 uint32 = 65535
                _ = v46
            }
            v47 := []int{127, 42}
            _ = v47
        }
        v1 = (1 >= -2147483648)
    }
    var v48 uint16 = 65535
    v4 -= 2147483647
    switch -2147483648 {
    case -648:
        v49 := GenS5{F0: (("a" + "") + (("" + ("\n\t" + "hello")) + "a")), F1: int8(v48)}
        _ = v49
    case 9223372036854775807:
        _ = len(v3)
    case 226:
        if v2 {
            if v1 {
                _ = len(v3)
                v50 := GenS3{F0: -534, F1: 42}
                _ = v50
            }
            var v51 int16 = -895
            _ = v51
        }
    default:
        fmt.Println(v4, v5, v1)
    }
    v48 = (^(v48))
    v52 := []int{}
    v53 := GenS2{F0: v48, F1: -1, F2: 2}
    fmt.Println(v4)
    if false {
        v54 := GenS2{F0: (v48 + v48), F1: -2147483648, F2: 1}
        _ = v54
    }
    v52 = append(v52, 338)
    if len(v52) >= 2 {
        v52 = v52[1:len(v52)]
    }
    for fi55 := 0; fi55 < 11; fi55++ {
        for fi56 := 0; fi56 < 6; fi56++ {
            if (!(1 >= 0)) {
                v2 = false
                v5 -= v5
                v48 = (^(v48))
            }
            _ = fi56
        }
        fmt.Println(v2, fi55, fi55)
        for fi57 := 0; fi57 < 16; fi57++ {
            if (103 > 0) {
                delete(v3, "hello")
                var v58 uint32 = (v5 % ((v4) | 1))
                fmt.Println(v48, fi55)
                _ = v58
            } else {
                var v59 int64 = -1
                var v60 uint64 = 42
                v4 += (^(v4))
                v60 -= 1
                _ = v59
                _ = v60
            }
            v5 = v4
            v5 += v4
            fmt.Println(v4, v5)
            v4 -= (v4 >> ((65535) & 63))
            _ = fi57
        }
        _ = fi55
    }
    v5 = (v4 | (v5 / ((v4) | 1)))
    v5 = 285
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v52:", len(v52), v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v52
    _ = v53
}
