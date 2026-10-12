package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int; F1 uint32; F2 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint64 = 1
    var v3 uint16 = uint16(v2)
    for fi4 := 0; fi4 < 18; fi4++ {
        v2 += v2
        fmt.Println(v1, fi4)
        v3 -= v3
        _ = fi4
    }
    fmt.Println(v2)
    v2 = 694
    fmt.Println(v2, v3, v1)
    if true {
        v2 = v2
    } else {
        v2 = v2
        if true {
            v1 = (v1 << ((510) & 63))
        }
        switch -2147483648 {
        case 0:
            var v5 int32 = -2147483648
            _ = v5
        }
    }
    var v6 int32 = -440
    var v7 float64 = float64((float32((-2.25 + 0.0)) * 1.5))
    if (9223372036854775807 != -2147483648) {
        v8 := GenS1{F0: 9223372036854775807, F1: 128, F2: 0}
        if ((!false) || true) {
            if true {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := GenS1{F0: -9223372036854775808, F1: 0, F2: 255}
                    switch -517 {
                    case 27:
                        var v11 uint32 = 190
                        _ = v11
                    case -1:
                        var v12 uint8 = 97
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
                for fi13 := 0; fi13 < 6; fi13++ {
                    v7 += 0.0
                    _ = fi13
                }
                var v14 int8 = v1
                _ = v14
            } else {
                v7 -= (1.5 + float64((((3.14159 * (float32(float64(-2.25)) / float32(float64((0.1 / -2.25))))) + (3.14159 - (1.5 - (0.0 - 0.0)))) + (float32(((v7 * v7) + (-2.25 + float64(((0.0 / 1.0) * 0.1))))) * -2.25))))
                v15 := GenS2{F0: 961}
                _ = v15
            }
        }
        _ = v8
    }
    switch 42 {
    case 32768:
        var v16 uint8 = 128
        _ = v16
    case 1:
        switch ("foo bar" + "a") {
        case "foo bar":
            var v17 float64 = (v7 / v7)
            _ = v17
        case "a":
            v3 = (v3 >> ((487) & 63))
        case "\n\t":
            switch "hello" {
            case "a":
                fmt.Println(v1, v1, v1)
            default:
                var v18 uint32 = 2
                _ = v18
            }
        default:
            if true {
                fmt.Println(v7, v6, v7)
                if (0 < 2) {
                    fmt.Println(v6, v6)
                    for fi19 := 0; fi19 < 6; fi19++ {
                        v20 := GenS3{F0: 244, F1: 65535, F2: 1}
                        for fi21 := 0; fi21 < 7; fi21++ {
                            fmt.Println(fi21)
                            v1 += (v1 >> ((65535) & 63))
                            var v22 uint16 = 65535
                            if true {
                                v2 = 2
                                v3 = 65535
                                switch "hello" {
                                case "\n\t":
                                    v23 := GenS4{F0: 398, F1: 241, F2: v1}
                                    _ = v23
                                case "hello":
                                    if true {
                                        v24 := []int{(fi19 >> ((0) & 63)), (fi19 >> ((255) & 63))}
                                        switch fi19 {
                                        case 932:
                                            var v25 uint32 = 1
                                            _ = v25
                                        case 9223372036854775807:
                                            if (0.1 >= v7) {
                                                v26 := GenS5{F0: (^(v3)), F1: (^(v22))}
                                                v22 -= 2
                                                if len(v24) > 0 {
                                                    v24[0] = (-(fi21))
                                                }
                                                _ = v26
                                            }
                                        default:
                                            switch "" {
                                            case "hello":
                                                v3 += v22
                                            case "":
                                                v27 := GenS5{F0: (^(v22)), F1: (v3 + (^(v3)))}
                                                _ = v27
                                            case "a":
                                                var v28 uint = 0
                                                _ = v28
                                            default:
                                                var v29 float32 = float32(v7)
                                                _ = v29
                                            }
                                        }
                                        v30 := GenS6{F0: (-(fi21)), F1: 1, F2: 127}
                                        _ = v24
                                        _ = v30
                                    }
                                default:
                                    if ("" == (((("a" + ("a" + "\n\t")) + ("foo bar" + "\n\t")) + "") + "\n\t")) {
                                        v31 := []int{fi19, fi21}
                                        var v32 uint = 874
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            } else {
                                var v33 uint8 = 42
                                var v34 uint8 = uint8(v3)
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(fi19)
                            _ = fi21
                            _ = v22
                        }
                        v7 = (float64((0.1 - ((float32((0.0 - (v7 - v7))) + ((1.5 - 3.14159) / 1.5)) - (1e10 - 1e10)))) - v7)
                        _ = fi19
                        _ = v20
                    }
                }
                v1 -= v1
            } else {
                var v35 bool = true
                fmt.Println(v35)
                v35 = (!(2 < 29))
                _ = v35
            }
        }
    case -678:
        switch 1 {
        case 403:
            v36 := GenS6{F0: 2147483647, F1: 451, F2: 137}
            _ = v36
        case 32768:
            v37 := []int{-9223372036854775808}
            _ = v37
        }
    }
    if (!false) {
        v38 := []int{-9223372036854775808, int(v2), 9223372036854775807, 2147483647}
        _ = len(v38)
        _ = v38
    } else {
        if (false && true) {
            var v39 int32 = (v6 - (^(v6)))
            if true {
                if ((!true) && ((v2 << ((255) & 63)) > v2)) {
                    fmt.Println(v7, v6, v2)
                }
            }
            if (v3 < (^(v3))) {
                v2 += (v2 << ((0) & 63))
                var v40 string = "foo bar"
                _ = v40
            } else {
                fmt.Println(v1)
                if (697 < 1) {
                    v6 -= 2147483647
                    for fi41 := 0; fi41 < 1; fi41++ {
                        v42 := []int{int(v39), (fi41 - -698), (fi41 << ((1) & 63))}
                        v43 := GenS6{F0: (fi41 + fi41), F1: 528, F2: 0}
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 9; fi44++ {
                        v6 -= (v39 >> ((42) & 63))
                        v1 -= (^(v1))
                        v45 := []int{(fi44 / (((fi44 ^ 32767)) | 1)), fi44}
                        _ = fi44
                        _ = v45
                    }
                } else {
                    v46 := GenS1{F0: 9223372036854775807, F1: 2, F2: 2}
                    var v47 float32 = ((1e10 - 3.14159) + 1.5)
                    var v48 string = "hello"
                    _ = v46
                    _ = v47
                    _ = v48
                }
                v1 = -66
            }
            _ = v39
        } else {
            v7 += 0.0
            v49 := []int{}
            _ = v49
        }
        if ((v3 >> ((65535) & 63)) < 394) {
            fmt.Println(v6, v1, v7)
        }
        var v50 int64 = 42
        _ = v50
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
