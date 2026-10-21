package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 bool; F1 int8; F2 int16 }
func genFn0() int64 {
    var v4 uint8 = 2
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v4)
        _ = fi5
    }
    var v6 float64 = ((1.5 - -2.25) - 3.14159)
    _ = v4
    _ = v6
    return 0
}

func main() {
    var v1 bool = true
    var v2 int64 = 0
    v3 := make(map[uint32]bool)
    fmt.Println(v2, v1)
    v7 := []int{32768}
    switch "hello" {
    case "hello":
        delete(v3, 4294967295)
    default:
        v1 = v1
    }
    v2 += 0
    v2 -= genFn0()
    if ((false || true) && v1) {
        var v8 int16 = 0
        for fi9 := 0; fi9 < 18; fi9++ {
            if ((-2.25 == ((float32(3.14159) / 1.0) / 1e10)) || (65535 <= 42)) {
                v8 = v8
                switch ("hello" + ("" + "")) {
                case "":
                    v2 -= v2
                case "foo bar":
                    v10 := GenS1{F0: fi9, F1: 114, F2: 1}
                    _ = v10
                }
                var v11 int32 = 2147483647
                _ = v11
            } else {
                v12 := GenS2{F0: 65535, F1: -392}
                v13 := GenS2{F0: 2147483648, F1: 547}
                _ = v12
                _ = v13
            }
            switch 127 {
            case -9223372036854775808:
                switch fi9 {
                case -9223372036854775808:
                    v8 -= v8
                case -1:
                    var v14 float64 = 1.5
                    _ = v14
                case -468:
                    fmt.Println(v2, v8)
                }
            }
            fmt.Println(v2, v2)
            _ = fi9
        }
        fmt.Println(v1)
        _ = v8
    } else {
        fmt.Println(v1, v2)
        fmt.Println(v1)
        fmt.Println(v1, v2)
    }
    var v15 uint = 674
    switch ("a" + (("" + "") + ("\n\t" + "foo bar"))) {
    case "hello":
        if ((0 < -786) || (true && (96 < 42))) {
            switch 453 {
            case 42:
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 int64 = -1
                    if true {
                        _ = len(v7)
                    } else {
                        v1 = false
                        if false {
                            switch ("foo bar" + ("hello" + ("" + ("foo bar" + "\n\t")))) {
                            case "foo bar":
                                var v18 int8 = 100
                                _ = v18
                            case "hello":
                                v2 -= (-(v17))
                            }
                            for fi19 := 0; fi19 < 19; fi19++ {
                                v20 := GenS2{F0: 268, F1: -2147483648}
                                fmt.Println(v15)
                                var v21 string = ("hello" + (("hello" + ("" + "")) + "hello"))
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            v22 := GenS3{F0: 542}
                            _ = v22
                        }
                    }
                    _ = fi16
                    _ = v17
                }
            case 1:
                var v23 string = ((("\n\t" + "") + ("foo bar" + "\n\t")) + "hello")
                _ = v23
            case 32767:
                var v24 bool = (128 <= 128)
                _ = v24
            default:
                var v25 uint64 = 18446744073709551615
                _ = v25
            }
            v2 = 9223372036854775807
            var v26 float64 = 1.5
            _ = v26
        } else {
            if v1 {
                v27 := GenS4{F0: 127, F1: 95, F2: 1}
                if v1 {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v1 = ((((((^(v2)) > v2) || true) || v1) && v1) && (0.0 == 0.0))
                        for fi29 := 0; fi29 < 20; fi29++ {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v2)
                            case "foo bar":
                                fmt.Println(v15, fi28)
                            default:
                                if (true || v1) {
                                    var v30 uint64 = 42
                                    if v1 {
                                        v7 = append(v7, 0)
                                        v15 = v15
                                    }
                                    var v31 uint8 = 255
                                    _ = v30
                                    _ = v31
                                } else {
                                    v32 := GenS4{F0: uint8(v15), F1: 2, F2: 42}
                                    fmt.Println(v1, fi29)
                                    var v33 float32 = 0.0
                                    _ = v32
                                    _ = v33
                                }
                            }
                            if false {
                                fmt.Println(v2, v2, fi29)
                                v34 := GenS5{F0: (v2 + 1)}
                                fmt.Println(fi28)
                                _ = v34
                            } else {
                                fmt.Println(v15, v2, fi29)
                            }
                            _ = fi29
                        }
                        _ = fi28
                    }
                    var v35 uint32 = 53
                    v36 := GenS6{F0: ("" + ""), F1: v2}
                    _ = v35
                    _ = v36
                } else {
                    var v37 int16 = -1
                    var v38 string = ""
                    v37 = 42
                    _ = v37
                    _ = v38
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v27
            }
            v39 := GenS7{F0: true, F1: -71, F2: -32768}
            v7 = append(v7, 202)
            _ = v39
        }
    }
    if true {
        switch -1 {
        case 42:
            _ = len(v3)
        case 338:
            v15 += 106
        case 1:
            switch ("a" + ("\n\t" + "hello")) {
            case "\n\t":
                fmt.Println(v1, v15, v2)
            default:
                v15 -= (v15 + 856)
            }
        default:
            fmt.Println(v15, v1)
        }
        v2 += 0
        var v40 uint32 = 4294967295
        _ = v40
    } else {
        v41 := v3[65535]
        _ = v41
    }
    if len(v7) > 0 {
        v7[0] = -2147483648
    }
    var v42 int = 187
    for fi43 := 0; fi43 < 13; fi43++ {
        v44 := v3[274]
        var v45 uint16 = uint16(v15)
        _ = fi43
        _ = v44
        _ = v45
    }
    delete(v3, 63)
    var v46 uint16 = 0
    var v47 bool = ((v15 >> ((4294967295) & 63)) == (v15 | (v15 ^ 823)))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v42
    _ = v46
    _ = v47
    _ = v7
}
