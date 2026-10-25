package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint }

func main() {
    var v1 int8 = 1
    var v2 uint8 = 61
    v3 := []uint32{12}
    var v4 int32 = -1
    v5 := GenS1{F0: 1, F1: 290, F2: 2}
    for fi6 := 0; fi6 < 7; fi6++ {
        v2 += 203
        v2 = (v2 % ((v2) | 1))
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v2)
            _ = len(v3)
            fmt.Println(fi7)
            _ = fi7
        }
        _ = fi6
    }
    var v8 int8 = 25
    v9 := GenS2{F0: 42}
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS3{F0: (fi10 | (fi10 - fi10))}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if true {
            v8 -= 1
            v8 += (-(v8))
        }
        _ = fi12
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v13 uint = 2
    v4 -= v4
    if (v2 >= 1) {
        fmt.Println(v2, v4, v2)
        for fi14 := 0; fi14 < 11; fi14++ {
            switch "foo bar" {
            case "":
                var v15 uint8 = (v2 >> (((v13 & (v13 & (v13 ^ (v13 >> ((1) & 63)))))) & 63))
                _ = v15
            case "foo bar":
                v16 := GenS1{F0: 313, F1: 2147483648, F2: 1}
                _ = v16
            default:
                for fi17 := 0; fi17 < 12; fi17++ {
                    if (-508 != 0) {
                        if (1.5 != (float64(1.5) - ((float64(1.5) - 0.1) + (((0.1 * (1.5 / ((float64(0.0) - -2.25) * -2.25))) + (float64(float32(((-2.25 * 1.5) + 1e10))) / 1.5)) / (1e10 + float64(float32((1e10 / 1.0)))))))) {
                            if (1 < -1) {
                                if true {
                                    v18 := GenS4{F0: (^(fi14))}
                                    var v19 int32 = -2147483648
                                    var v20 uint16 = 916
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    if true {
                                        v21 := GenS5{F0: v13, F1: (fi17 - (fi14 | 42)), F2: 255}
                                        if false {
                                            if len(v3) > 0 {
                                                v3[0] = 1
                                            }
                                            fmt.Println(v8, fi14, v2)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v22 int8 = -128
                                            if true {
                                                fmt.Println(fi14, v13)
                                                v13 += 345
                                            } else {
                                                v22 -= (v22 + (v22 | (v8 ^ v22)))
                                                v22 -= v8
                                                var v23 int16 = int16(v13)
                                                v24 := GenS2{F0: 9223372036854775807}
                                                if (!(255 > 65535)) {
                                                    fmt.Println(v13, v22)
                                                    v8 -= (-(v22))
                                                } else {
                                                    var v25 uint = v13
                                                    _ = v25
                                                }
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v22
                                        } else {
                                            fmt.Println(fi14)
                                            v2 -= 1
                                            switch -765 {
                                            case -1:
                                                var v26 int8 = (v1 * v8)
                                                _ = v26
                                            case -9:
                                                v1 -= 1
                                            case 4294967295:
                                                if false {
                                                    if (1 < -347) {
                                                        var v27 bool = (951 <= 597)
                                                        fmt.Println(v13, fi17)
                                                        _ = v27
                                                    }
                                                } else {
                                                    switch 32768 {
                                                    case 4294967295:
                                                        var v28 uint16 = 1
                                                        _ = v28
                                                    case 0:
                                                        fmt.Println(fi17, v1)
                                                    case 127:
                                                        if len(v3) > 0 {
                                                            v3[0] = 76
                                                        }
                                                    default:
                                                        v2 += (^(v2))
                                                    }
                                                    var v29 bool = true
                                                    _ = v29
                                                }
                                            default:
                                                fmt.Println(v4, v2, v8)
                                            }
                                        }
                                        var v30 string = ""
                                        _ = v21
                                        _ = v30
                                    }
                                    var v31 int = fi14
                                    _ = v31
                                }
                            } else {
                                fmt.Println(v2)
                            }
                            v13 -= v13
                        } else {
                            var v32 uint32 = 65535
                            v1 = (v1 << ((v13) & 63))
                            v13 -= (^(v13))
                            _ = v32
                        }
                        if (("foo bar" + "hello") != "") {
                            var v33 int64 = 42
                            v1 += (-(v8))
                            _ = v33
                        } else {
                            var v34 int32 = 2147483647
                            _ = v34
                        }
                    } else {
                        var v35 string = "foo bar"
                        var v36 uint8 = v2
                        _ = len(v3)
                        _ = v35
                        _ = v36
                    }
                    var v37 int64 = 1
                    _ = fi17
                    _ = v37
                }
            }
            _ = len(v3)
            _ = fi14
        }
    }
    fmt.Println(v13, v8)
    if (!(803 != 2147483647)) {
        if (32 > -1) {
            var v38 float32 = (3.14159 - 1.5)
            fmt.Println(v2, v13)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v38
        } else {
            var v39 uint8 = v2
            if true {
                var v40 int8 = (^(v1))
                fmt.Println(v4, v4, v2)
                _ = v40
            } else {
                v39 = v2
                _ = len(v3)
                switch "" {
                case "hello":
                    var v41 int8 = v1
                    _ = v41
                case "":
                    v42 := GenS3{F0: 2147483647}
                    _ = v42
                case "a":
                    v43 := GenS2{F0: 93}
                    _ = v43
                default:
                    v4 = v4
                }
            }
            var v44 int16 = -1
            _ = v39
            _ = v44
        }
        var v45 string = "a"
        var v46 uint64 = 18446744073709551615
        _ = v45
        _ = v46
    }
    v47 := GenS6{F0: (false && (!true)), F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v8
    _ = v9
}
