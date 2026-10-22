package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 uint; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 255
    v3 := []int{2147483647, -549, 4294967295, 42}
    v4 := []string{"\n\t"}
    var v5 int8 = 1
    v2 = 42
    v5 = (v5 >> ((uint(v2)) & 63))
    if (!(42 != 9223372036854775807)) {
        v6 := GenS1{F0: (v5 + 42), F1: 2}
        _ = v6
    }
    v7 := GenS1{F0: (v5 << ((2) & 63)), F1: 1}
    switch "hello" {
    case "\n\t":
        v3 = append(v3, 32767)
    }
    switch 42 {
    case 4294967295:
        var v8 uint16 = 0
        _ = v8
    }
    if ((v2 >> ((695) & 63)) != v2) {
        fmt.Println(v2)
        fmt.Println(v2, v2)
        if true {
            v5 = -128
            if (!(true || true)) {
                switch "\n\t" {
                case "\n\t":
                    for fi9 := 0; fi9 < 8; fi9++ {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            v11 := GenS2{F0: 42, F1: 1, F2: 2}
                            v12 := GenS1{F0: (v5 ^ (v5 % ((-12) | 1))), F1: 351}
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                        switch (-(fi9)) {
                        case -925:
                            for fi13 := 0; fi13 < 11; fi13++ {
                                var v14 uint16 = 729
                                var v15 uint32 = 548
                                _ = fi13
                                _ = v14
                                _ = v15
                            }
                        case -1:
                            v3 = append(v3, -1)
                        case 633:
                            fmt.Println(v2, v2)
                        }
                        if (1 >= -1) {
                            v2 -= (v2 - (v2 >> ((900) & 63)))
                            v16 := GenS3{F0: (-(v5))}
                            _ = v16
                        } else {
                            if (255 == 97) {
                                switch ("a" + (("\n\t" + "hello") + (("a" + ("foo bar" + "hello")) + "foo bar"))) {
                                case "foo bar":
                                    v2 -= (v2 ^ (v2 | (v2 * v2)))
                                case "hello":
                                    fmt.Println(fi9, v2)
                                default:
                                    if (false && (160 != v2)) {
                                        v5 -= v5
                                        var v17 uint16 = 255
                                        _ = v17
                                    } else {
                                        var v18 float32 = 1e10
                                        for fi19 := 0; fi19 < 11; fi19++ {
                                            v20 := GenS4{F0: 678, F1: 65535, F2: -14}
                                            _ = fi19
                                            _ = v20
                                        }
                                        var v21 int16 = 970
                                        _ = v18
                                        _ = v21
                                    }
                                }
                                var v22 int = (fi9 - 42)
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    v24 := GenS5{F0: 538}
                                    v25 := GenS6{F0: 4294967295}
                                    fmt.Println(v2, v2)
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                switch "" {
                                case "a":
                                    var v26 uint = 811
                                    _ = v26
                                case "foo bar":
                                    var v27 float32 = ((1.5 / 1.0) + 3.14159)
                                    _ = v27
                                case "\n\t":
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v3) > 0 {
                                            v3[0] = fi9
                                        }
                                    case "\n\t":
                                        v2 = v2
                                    case "a":
                                        v28 := GenS2{F0: 575, F1: 0, F2: 501}
                                        _ = v28
                                    default:
                                        var v29 float32 = -2.25
                                        _ = v29
                                    }
                                default:
                                    v2 += v2
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            delete(v1, ("foo bar" + "hello"))
                        }
                        _ = fi9
                    }
                case "":
                    if (2 >= 18446744073709551615) {
                        fmt.Println(v2, v2)
                    }
                case "hello":
                    v2 += (v2 / ((v2) | 1))
                }
                var v30 uint8 = 53
                if (2147483647 <= 42) {
                    for fi31 := 0; fi31 < 14; fi31++ {
                        var v32 int16 = int16(v30)
                        fmt.Println(v30)
                        var v33 bool = (-9223372036854775808 <= 42)
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    v30 = 2
                    v34 := GenS7{F0: 42, F1: 4294967295}
                    _ = v34
                }
                _ = v30
            }
            if true {
                for fi35 := 0; fi35 < 16; fi35++ {
                    fmt.Println(v5)
                    for fi36 := 0; fi36 < 12; fi36++ {
                        v4 = append(v4, "hello")
                        v5 += -128
                        _ = fi36
                    }
                    _ = fi35
                }
            }
        } else {
            var v37 int = 1
            for fi38 := 0; fi38 < 20; fi38++ {
                v5 = v5
                if false {
                    var v39 int16 = -1
                    var v40 int = (fi38 / ((681) | 1))
                    _ = v39
                    _ = v40
                }
                _ = fi38
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v37
        }
    }
    var v41 int8 = (v5 % ((42) | 1))
    var v42 int16 = -882
    var v43 float64 = 1.5
    v42 -= v42
    for fi44 := 0; fi44 < 12; fi44++ {
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "\n\t":
            v2 -= (v2 / ((61) | 1))
        case "hello":
            v4 = append(v4, "")
        case "":
            var v45 int32 = -474
            _ = v45
        }
        _ = fi44
    }
    v46 := GenS8{F0: 2147483647}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v46
    _ = v5
    _ = v7
}
