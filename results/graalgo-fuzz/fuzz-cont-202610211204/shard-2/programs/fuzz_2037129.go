package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16) uint {
    fmt.Println(p1, p1)
    p1 -= p1
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: 0, F1: 127}
    if (0.0 == 0.0) {
        delete(v2, 42)
        v5 := GenS1{F0: -9223372036854775808, F1: 42}
        _ = v5
    } else {
        v6 := GenS1{F0: -1, F1: -128}
        v7 := GenS1{F0: 9223372036854775807, F1: 1}
        var v8 uint8 = 42
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 float32 = (3.14159 - float32(1e10))
    v9 += -2.25
    if (1 <= -583) {
        if true {
            switch 939 {
            case 9223372036854775807:
                if (42 >= 32767) {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 uint64 = 2
                        if true {
                            if (false || true) {
                                v9 = float32(1.5)
                                v11 = (v11 * 357)
                                v11 = 4294967295
                            }
                        } else {
                            var v12 bool = (true || ("hello" != ""))
                            _ = v12
                        }
                        v9 = v9
                        _ = fi10
                        _ = v11
                    }
                }
            case 42:
                v9 -= 0.0
            }
        } else {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int = fi13
                v15 := GenS1{F0: 0, F1: 0}
                v16 := GenS1{F0: -477, F1: 100}
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 uint8 = 42
        _ = v17
    } else {
        var v18 string = (("a" + "") + "")
        var v19 uint64 = 1
        for fi20 := 0; fi20 < 12; fi20++ {
            for fi21 := 0; fi21 < 11; fi21++ {
                if (0.0 < float64(v9)) {
                    v18 = v18
                    v1 = append(v1, fi21)
                    switch v18 {
                    case "\n\t":
                        if false {
                            v22 := v3[fi20]
                            fmt.Println(fi20)
                            var v23 int8 = -1
                            _ = v22
                            _ = v23
                        }
                    case "a":
                        v24 := GenS2{F0: 4294967295, F1: 65535, F2: 18446744073709551615}
                        _ = v24
                    case "hello":
                        if true {
                            v9 = v9
                            if (127 > fi20) {
                                if (v18 == v18) {
                                    var v25 uint8 = 128
                                    fmt.Println(fi20, v9, v18)
                                    v26 := GenS2{F0: 0, F1: 876, F2: uint(v25)}
                                    _ = v25
                                    _ = v26
                                }
                            } else {
                                v9 -= (float32((3.14159 - (0.1 - (0.0 / 1e10)))) - 3.14159)
                                var v27 string = ((("" + v18) + v18) + v18)
                                var v28 uint32 = 1
                                _ = v27
                                _ = v28
                            }
                            switch fi21 {
                            case -218:
                                v18 = v18
                            case 1:
                                switch "hello" {
                                case "":
                                    var v29 uint = 18446744073709551615
                                    _ = v29
                                default:
                                    fmt.Println(v19, fi20)
                                }
                            default:
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    v30 := GenS3{F0: 1, F1: 0}
                                    _ = v30
                                }
                            }
                        } else {
                            if ((!false) || false) {
                                v1 = append(v1, fi21)
                                v31 := GenS3{F0: 9223372036854775807, F1: uint32(v19)}
                                var v32 uint16 = 243
                                _ = v31
                                _ = v32
                            }
                            if len(v1) > 0 {
                                v1[0] = (^(fi20))
                            }
                            v9 = 0.0
                        }
                    default:
                        switch "" {
                        case "":
                            var v33 int64 = 9223372036854775807
                            _ = v33
                        default:
                            if (((((float32(1.5) * v9) - (v9 - 1.5)) - (v9 - (1e10 * 0.1))) > v9) || false) {
                                v34 := GenS3{F0: -801, F1: 106}
                                _ = v34
                            }
                        }
                    }
                }
                switch -2147483648 {
                case 0:
                    _ = len(v2)
                case -635:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -2147483648:
                    var v35 int = (fi21 ^ 211)
                    _ = v35
                default:
                    if false {
                        if (!true) {
                            v36 := v3[9223372036854775807]
                            var v37 int8 = 100
                            _ = v36
                            _ = v37
                        }
                    }
                    switch -1 {
                    case 127:
                        delete(v3, (fi21 % (((fi21 - fi20)) | 1)))
                    case 0:
                        v9 -= v9
                    case 606:
                        v9 = 1e10
                    default:
                        v38 := GenS2{F0: 2147483647, F1: 2147483648, F2: 543}
                        _ = v38
                    }
                }
                _ = fi21
            }
            var v39 uint64 = v19
            v40 := GenS3{F0: -9223372036854775808, F1: 155}
            _ = fi20
            _ = v39
            _ = v40
        }
        _ = v18
        _ = v19
    }
    var v41 string = ""
    for fi42 := 0; fi42 < 18; fi42++ {
        v3[fi42] = v41
        fmt.Println(fi42, v9, v41)
        _ = fi42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        fmt.Println(v9, v41)
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "\n\t":
            var v44 float32 = (v9 / (0.1 + v9))
            _ = v44
        default:
            v45 := GenS4{F0: 1}
            _ = v45
        }
        if true {
            v9 -= v9
        }
        _ = fi43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v9
}
