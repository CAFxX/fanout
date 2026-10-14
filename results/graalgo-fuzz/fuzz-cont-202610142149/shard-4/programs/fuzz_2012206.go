package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32; F2 uint16 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }

func main() {
    var v1 int16 = 42
    var v2 int8 = 100
    var v3 int16 = v1
    v4 := make(map[uint32]uint32)
    v5 := []int{-2147483648, 0}
    if false {
        v1 = (v1 + (v1 | (v1 << ((1) & 63))))
        for fi6 := 0; fi6 < 2; fi6++ {
            if (65535 == 253) {
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                v1 = v1
                v5 = append(v5, fi6)
            }
            _ = fi6
        }
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 float32 = 1e10
            v2 = 42
            v8 += v8
            _ = fi7
            _ = v8
        }
    } else {
        var v9 uint = 1
        fmt.Println(v3, v1)
        _ = v9
    }
    if (993 != -665) {
        var v10 int64 = -1
        v3 = -32768
        _ = v10
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        if true {
            switch 9223372036854775807 {
            case 32767:
                v2 = (v2 + -7)
            case 2147483647:
                var v12 uint64 = 4294967295
                _ = v12
            default:
                v4[2147483647] = 2147483647
            }
            v1 = (v1 - -32768)
            v13 := GenS1{F0: "\n\t"}
            _ = v13
        }
        if true {
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 float32 = ((0.1 + (0.1 / 1.0)) + float32(-2.25))
                var v16 int = (fi14 - (fi14 ^ (fi11 << ((18446744073709551615) & 63))))
                var v17 uint32 = 1
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            var v18 uint = 18446744073709551615
            if false {
                v3 = (v1 | v3)
            } else {
                _ = len(v4)
                for fi19 := 0; fi19 < 12; fi19++ {
                    if (false || ((!false) && true)) {
                        var v20 uint16 = 144
                        v21 := GenS2{F0: 42, F1: 0, F2: 255}
                        var v22 uint16 = (^(v20))
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 float32 = (3.14159 + -2.25)
                    _ = fi19
                    _ = v23
                }
                var v24 int8 = v2
                _ = v24
            }
            if true {
                var v25 uint64 = 1
                var v26 bool = (!true)
                _ = v25
                _ = v26
            } else {
                if true {
                    if len(v5) > 0 {
                        v5[0] = fi11
                    }
                } else {
                    v3 = (v1 | (v3 | v3))
                }
                for fi27 := 0; fi27 < 5; fi27++ {
                    switch ("\n\t" + (("foo bar" + "\n\t") + (("hello" + ("" + ("" + "a"))) + "hello"))) {
                    case "\n\t":
                        v3 -= (v1 / ((v3) | 1))
                    case "foo bar":
                        if (!(1 != -491)) {
                            v18 -= v18
                            if (2147483647 != 562) {
                                v1 = 42
                            }
                            v3 = (v1 + (v3 & (-(v1))))
                        } else {
                            v2 = v2
                        }
                    }
                    v2 -= v2
                    switch fi27 {
                    case -375:
                        var v28 uint8 = 127
                        if (((0.1 / 0.1) <= float64(-2.25)) || (!true)) {
                            var v29 uint64 = 2
                            if (!false) {
                                var v30 string = "hello"
                                _ = v30
                            } else {
                                var v31 int = 602
                                v29 -= v29
                                v31 -= v31
                                _ = v31
                            }
                            fmt.Println(fi11, v29)
                            _ = v29
                        }
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            if false {
                                if (((((1e10 / 1e10) - (float64(((0.1 / float32(1e10)) / (1.5 - ((1e10 / 1.0) + 1.5)))) * (((-2.25 * (3.14159 - 3.14159)) * (float64(((1e10 - 1.5) * 3.14159)) + (3.14159 - 0.0))) + float64((0.1 + -2.25))))) * 0.0) != (3.14159 - 3.14159)) || (-9223372036854775808 > -1)) {
                                    if false {
                                        var v32 int = (fi27 >> ((v18) & 63))
                                        _ = len(v4)
                                        v33 := GenS2{F0: -1, F1: -2147483648, F2: 884}
                                        _ = v32
                                        _ = v33
                                    } else {
                                        fmt.Println(v3, v28)
                                    }
                                }
                                var v34 int = 688
                                v34 -= (fi27 - 128)
                                _ = v34
                            }
                        case "a":
                            v4[2147483648] = 271
                        default:
                            var v35 uint = v18
                            _ = v35
                        }
                        _ = v28
                    case -1:
                        v2 += int8(v18)
                    default:
                        v2 += 1
                    }
                    if true {
                        v36 := GenS2{F0: -2147483648, F1: 0, F2: 0}
                        _ = v36
                    } else {
                        fmt.Println(fi27, v1, fi27)
                        if false {
                            v37 := GenS1{F0: "a"}
                            v18 = v18
                            _ = v37
                        } else {
                            v1 += (v3 ^ (-(v3)))
                        }
                    }
                    _ = fi27
                }
                fmt.Println(v3, v1, v18)
            }
            _ = v18
        }
        _ = fi11
    }
    for fi38 := 0; fi38 < 6; fi38++ {
        if true {
            var v39 int64 = 485
            _ = v39
        }
        v40 := GenS2{F0: 42, F1: -1, F2: 255}
        _ = fi38
        _ = v40
    }
    v3 -= v3
    v41 := GenS3{F0: (v2 - 100), F1: v2, F2: 2}
    var v42 uint = 968
    v1 = (v3 / ((v1) | 1))
    v43 := GenS3{F0: int8(v1), F1: v2, F2: 316}
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 int8 = v2
        v2 = (v2 - v2)
        fmt.Println(fi44, v1, v2)
        _ = fi44
        _ = v45
    }
    v1 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
