package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int }
func genFn0(p0 uint8, p1 int32, p2 int32) int {
    v6 := []int{933, -949}
    var v7 float32 = 1e10
    if (-2.25 <= 0.1) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v7 = 3.14159
    var v8 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -2147483648
}

func main() {
    var v1 uint64 = 2
    var v2 uint16 = 283
    var v3 string = (("\n\t" + "") + "\n\t")
    var v4 int = int(v1)
    v5 := make(map[string]uint32)
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            switch v3 {
            case "":
                switch (^(fi10)) {
                case 125:
                    v11 := []int{genFn0(222, -2147483648, 0), (fi10 / ((32768) | 1))}
                    _ = v11
                case 0:
                    fmt.Println(fi9, fi10)
                default:
                    var v12 int16 = int16(v2)
                    _ = v12
                }
            case "a":
                if ((v3 + (((("a" + (v3 + "a")) + v3) + "hello") + v3)) > v3) {
                    if (442 < 454) {
                        v13 := []int{}
                        fmt.Println(fi9, fi10)
                        if (-12 > 100) {
                            var v14 uint16 = (v2 & v2)
                            _ = v14
                        } else {
                            fmt.Println(fi10)
                            v2 += v2
                            if (1000 != -32768) {
                                v2 = (v2 % ((0) | 1))
                                if (!true) {
                                    v3 = v3
                                    var v15 float64 = 3.14159
                                    _ = v15
                                }
                                var v16 string = "a"
                                _ = v16
                            } else {
                                v17 := GenS1{F0: 42, F1: 250}
                                v18 := GenS1{F0: 1, F1: 0}
                                fmt.Println(v3, fi10, v3)
                                _ = v17
                                _ = v18
                            }
                        }
                        _ = v13
                    } else {
                        if (!(711 != v2)) {
                            v1 += 4294967295
                            v5[v3] = 254
                        }
                        v19 := GenS1{F0: -1, F1: 443}
                        _ = v19
                    }
                    v2 += uint16(v1)
                }
                fmt.Println(v1, v3)
                if ((0.0 / 0.1) != 1e10) {
                    v4 = 9223372036854775807
                } else {
                    v3 = (v3 + "a")
                }
                v2 += v2
            case "\n\t":
                v20 := GenS1{F0: 0, F1: 2147483648}
                _ = v20
            default:
                v2 = 65535
            }
            if (!true) {
                v4 = (-(v4))
                switch (v3 + (v3 + ((v3 + v3) + "a"))) {
                case "\n\t":
                    v5["foo bar"] = 751
                case "":
                    var v21 float32 = (1e10 * 1.5)
                    _ = v21
                case "hello":
                    var v22 int16 = 103
                    _ = v22
                default:
                    if true {
                        fmt.Println(fi10, v3)
                        if true {
                            switch "\n\t" {
                            case "a":
                                switch "hello" {
                                case "foo bar":
                                    v4 = v4
                                case "a":
                                    if false {
                                        var v23 int64 = 1
                                        v4 = genFn0(255, 0, 0)
                                        if true {
                                            v24 := []int{}
                                            _ = v24
                                        }
                                        _ = v23
                                    }
                                default:
                                    if (-1 <= -1) {
                                        var v25 float32 = 1e10
                                        _ = v25
                                    }
                                }
                            case "":
                                v26 := GenS1{F0: -14, F1: 65535}
                                _ = v26
                            case "hello":
                                v27 := []int{(fi10 + (fi9 - (fi9 * (fi10 ^ genFn0(128, 42, -2147483648))))), fi9, v4}
                                _ = v27
                            }
                            v1 = 18446744073709551615
                            delete(v5, v3)
                            v1 = v1
                        } else {
                            fmt.Println(v4, v2, v2)
                            var v28 uint16 = 1
                            _ = v28
                        }
                        v29 := GenS2{F0: -1, F1: v2, F2: v2}
                        v30 := GenS3{F0: 183, F1: 42}
                        _ = v29
                        _ = v30
                    } else {
                        switch (v4 << ((1) & 63)) {
                        case 9223372036854775807:
                            if false {
                                var v31 int = 987
                                var v32 int = v31
                                v1 = 0
                                _ = v31
                                _ = v32
                            } else {
                                v33 := []int{(fi9 >> ((0) & 63)), 32768}
                                v2 -= (^(v2))
                                v1 = v1
                                _ = v33
                            }
                        case -200:
                            fmt.Println(fi9, v3, v3)
                        case 1:
                            var v34 uint8 = 0
                            _ = v34
                        default:
                            v35 := GenS1{F0: 94, F1: 0}
                            _ = v35
                        }
                        fmt.Println(v2)
                        var v36 float64 = 0.1
                        _ = v36
                    }
                }
            } else {
                v2 = (v2 & (^(v2)))
                delete(v5, "a")
            }
            delete(v5, (v3 + ""))
            _ = fi10
        }
        v2 += v2
        _ = fi9
    }
    v37 := []int{(v4 ^ v4), v4, 0}
    fmt.Println(v3)
    _ = len(v37)
    if ((v4 - (v4 | (v4 + -1))) >= 988) {
        fmt.Println(v3, v3, v4)
        switch genFn0(128, 0, 1) {
        case 128:
            v2 = 0
        case 911:
            var v38 int = genFn0(1, -2147483648, 620)
            _ = v38
        default:
            var v39 uint8 = 128
            _ = v39
        }
        if len(v37) >= 2 {
            v37 = v37[1:len(v37)]
        }
    } else {
        v40 := GenS4{F0: 1}
        var v41 int = v4
        if (529 > 4294967295) {
            switch "" {
            case "":
                v41 += (v41 >> ((44) & 63))
            }
        }
        _ = v40
        _ = v41
    }
    fmt.Println(v1, v2, v2)
    v3 = v3
    if (1 <= 303) {
        v2 = v2
        v42 := GenS5{F0: 13, F1: -912}
        _ = v42
    }
    var v43 int = (v4 << ((4294967295) & 63))
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            fmt.Println(v1, v43)
            _ = fi45
        }
        v2 += 2
        if false {
            var v46 float32 = (0.0 - 1e10)
            v2 = (v2 << ((263) & 63))
            v37 = append(v37, (^(v4)))
            _ = v46
        } else {
            v47 := GenS1{F0: 1, F1: 65535}
            var v48 float32 = (1e10 - float32(3.14159))
            _ = v47
            _ = v48
        }
        _ = fi44
    }
    var v49 uint32 = 65535
    v3 = (("a" + (v3 + v3)) + v3)
    v2 += (v2 / ((32767) | 1))
    fmt.Println(v43, v1)
    fmt.Println(v4, v3, v3)
    v50 := v5[(((v3 + "\n\t") + v3) + v3)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v43
    _ = v49
    _ = v5
    _ = v50
}
