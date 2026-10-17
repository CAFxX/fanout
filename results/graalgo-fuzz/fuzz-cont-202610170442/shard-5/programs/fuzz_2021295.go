package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    var v1 int8 = 127
    var v2 string = (((((("a" + "a") + "") + "a") + "") + ("a" + ("\n\t" + "hello"))) + "foo bar")
    var v3 uint32 = 65535
    var v4 int16 = -1
    var v5 int32 = 1
    var v6 int16 = -808
    v7 := []int{218, 32767, -1, 9223372036854775807}
    fmt.Println(v6)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (1.5 < 3.14159) {
            switch fi8 {
            case -1:
                if true {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        switch fi9 {
                        case 381:
                            if (fi9 >= (fi8 % (((^(fi8))) | 1))) {
                                v7 = append(v7, (^(fi9)))
                            } else {
                                if (false && true) {
                                    if (!true) {
                                        switch "a" {
                                        case "":
                                            fmt.Println(v1)
                                        }
                                    }
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    v4 = -1
                                    if ((v5 + -791) == v5) {
                                        v4 = -445
                                        v10 := GenS1{F0: "\n\t", F1: (^(v3)), F2: 0}
                                        _ = v10
                                    }
                                }
                                v4 = (v6 & v6)
                            }
                        default:
                            v1 = -85
                        }
                        _ = fi9
                    }
                    for fi11 := 0; fi11 < 14; fi11++ {
                        var v12 bool = false
                        v2 = v2
                        _ = fi11
                        _ = v12
                    }
                }
            case 775:
                if (((4294967295 == 4294967295) || (!true)) && (false || (true || true))) {
                    var v13 uint8 = 42
                    v14 := GenS2{F0: v1}
                    _ = v13
                    _ = v14
                }
            default:
                v15 := GenS2{F0: v1}
                _ = v15
            }
            var v16 float32 = (0.1 + ((1.5 + float32(1.5)) / ((1.5 / 0.1) * -2.25)))
            v2 = v2
            _ = v16
        }
        var v17 float32 = 0.0
        if ((!(v6 > (-(v4)))) && (v3 == v3)) {
            if (((fi8 > 0) && (-1 == 1)) && (699 < 32767)) {
                var v18 int16 = (^(v4))
                v5 = v5
                _ = v18
            } else {
                v2 = v2
                var v19 float64 = 3.14159
                _ = v19
            }
            var v20 uint = 2
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v20
        } else {
            v21 := GenS1{F0: v2, F1: v3, F2: v4}
            _ = v21
        }
        _ = fi8
        _ = v17
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            switch 9223372036854775807 {
            case 4294967295:
                fmt.Println(fi22, v4)
                var v24 uint8 = 127
                var v25 int64 = -9223372036854775808
                _ = v24
                _ = v25
            }
            _ = fi23
        }
        var v26 uint16 = 32767
        _ = fi22
        _ = v26
    }
    fmt.Println(v4, v6)
    v1 -= (v1 ^ 127)
    var v27 string = (v2 + v2)
    switch -918 {
    case -9223372036854775808:
        v6 = (v6 + (v4 >> ((0) & 63)))
    default:
        var v28 uint32 = 265
        _ = v28
    }
    if (float64(0.0) > ((0.1 / 1.0) / 1.5)) {
        v29 := GenS1{F0: v27, F1: (v3 - 346), F2: v4}
        var v30 uint8 = 230
        v31 := GenS1{F0: "\n\t", F1: v3, F2: v4}
        _ = v29
        _ = v30
        _ = v31
    } else {
        v4 -= v4
        v32 := GenS1{F0: (v2 + v2), F1: v3, F2: (v4 + (v4 >> ((255) & 63)))}
        v33 := GenS3{F0: (v5 % ((v5) | 1)), F1: -149, F2: 9223372036854775807}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        v7 = append(v7, (fi34 | fi34))
        var v35 int16 = 42
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
