package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 42
    v3 := []int{}
    var v4 float32 = -2.25
    var v5 int64 = -9223372036854775808
    v4 += v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 9223372036854775807 {
    case 42:
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: 0}
            _ = fi6
            _ = v7
        }
    }
    var v8 int64 = (^(v2))
    fmt.Println(v5, v2, v2)
    if len(v3) > 0 {
        v3[0] = 32768
    }
    var v9 bool = false
    v3 = append(v3, 594)
    for fi10 := 0; fi10 < 7; fi10++ {
        v3 = append(v3, (^(fi10)))
        if v9 {
            v9 = (0 < 1)
        } else {
            var v11 int = (fi10 >> ((42) & 63))
            _ = v11
        }
        _ = fi10
    }
    var v12 uint8 = 2
    switch -1 {
    case 9223372036854775807:
        fmt.Println(v8, v4)
    case 32767:
        v13 := GenS2{F0: 32767, F1: -128, F2: 487}
        _ = v13
    case 4294967295:
        if v9 {
            switch "a" {
            case "":
                fmt.Println(v4, v12, v5)
            case "a":
                delete(v1, 1)
            }
            fmt.Println(v4, v9)
            fmt.Println(v12, v12)
        } else {
            switch 1 {
            case 32768:
                if (!false) {
                    v5 = (v2 & 570)
                    v12 -= 115
                } else {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v5, v9)
                    case "foo bar":
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 int = 42
                            _ = fi14
                            _ = v15
                        }
                    case "a":
                        if (-54 == -10) {
                            delete(v1, 888)
                            fmt.Println(v5, v4, v8)
                            var v16 int32 = 671
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                var v18 uint8 = 0
                                for fi19 := 0; fi19 < 20; fi19++ {
                                    v1[1] = 370
                                    _ = fi19
                                }
                                _ = fi17
                                _ = v18
                            }
                            if len(v3) > 0 {
                                v3[0] = -9223372036854775808
                            }
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v2)
                            case "":
                                if (-1 != -1) {
                                    _ = len(v3)
                                } else {
                                    if v9 {
                                        for fi20 := 0; fi20 < 11; fi20++ {
                                            var v21 float64 = ((-2.25 * 1.5) / -2.25)
                                            var v22 int = fi20
                                            for fi23 := 0; fi23 < 3; fi23++ {
                                                v3 = append(v3, (^(v22)))
                                                fmt.Println(fi23, v2)
                                                v24 := GenS3{F0: 0}
                                                _ = fi23
                                                _ = v24
                                            }
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        switch 32768 {
                                        case -2147483648:
                                            if ((-128 > 127) && (true && v9)) {
                                                v2 += v5
                                                v12 = (v12 * v12)
                                                switch -908 {
                                                case 4294967295:
                                                    v25 := v1[42]
                                                    _ = v25
                                                case 9223372036854775807:
                                                    var v26 int16 = 1000
                                                    _ = v26
                                                case -799:
                                                    var v27 bool = false
                                                    _ = v27
                                                default:
                                                    v5 -= 42
                                                }
                                            } else {
                                                v28 := GenS2{F0: 91, F1: 125, F2: 2147483647}
                                                _ = v28
                                            }
                                        default:
                                            v2 = v8
                                        }
                                        if (float32(0.0) > (float32(-2.25) - (v4 + 3.14159))) {
                                            var v29 float64 = 0.1
                                            if v9 {
                                                var v30 int = 323
                                                var v31 uint = 2
                                                fmt.Println(v4, v4, v31)
                                                _ = v30
                                                _ = v31
                                            }
                                            _ = v29
                                        } else {
                                            _ = len(v1)
                                            delete(v1, 0)
                                        }
                                    }
                                    v32 := GenS4{F0: true, F1: -32768, F2: -1}
                                    _ = v32
                                }
                            case "hello":
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    var v34 int8 = 72
                                    for fi35 := 0; fi35 < 14; fi35++ {
                                        _ = len(v1)
                                        v36 := GenS1{F0: -244}
                                        var v37 uint32 = 2
                                        var v38 bool = v9
                                        var v39 float32 = 3.14159
                                        v40 := GenS1{F0: v5}
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                    v9 = ((v34 / (((v34 << ((18446744073709551615) & 63))) | 1)) > (^(v34)))
                                    _ = fi33
                                    _ = v34
                                }
                            default:
                                _ = len(v3)
                            }
                        }
                    default:
                        v9 = ((v4 <= 0.1) && v9)
                    }
                    v12 = (v12 * (^(v12)))
                }
            case 707:
                v4 = v4
            }
        }
    }
    v12 -= v12
    if (-181 > 1) {
        if len(v3) > 0 {
            v3[0] = 32768
        }
    } else {
        v5 += v2
        var v41 int8 = 82
        _ = v41
    }
    v42 := GenS3{F0: 2147483647}
    var v43 float32 = (v4 * v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
