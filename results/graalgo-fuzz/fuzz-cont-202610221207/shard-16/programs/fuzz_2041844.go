package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int64; F2 string }
type GenS8 struct { F0 int16; F1 uint16; F2 int32 }
func genFn0(p0 uint64, p1 int) uint16 {
    fmt.Println(p1)
    p0 = p0
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(fi5, p0, p1)
        _ = fi5
    }
    v6 := GenS1{F0: 127, F1: 279}
    var v7 float32 = -2.25
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = 0
    var v3 uint32 = 0
    var v4 int = 9223372036854775807
    v8 := GenS2{F0: 42, F1: 65535, F2: -128}
    v9 := GenS3{F0: v2, F1: 128, F2: (("" + "hello") + (("\n\t" + "a") + ("" + "\n\t")))}
    if true {
        v10 := GenS4{F0: 0, F1: 0}
        var v11 int16 = 133
        fmt.Println(v4, v4)
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            if (!(42 > 127)) {
                v13 := []int{(v4 * v4), -9223372036854775808, -1, (^(fi12))}
                _ = v13
            } else {
                fmt.Println(v3, fi12)
                fmt.Println(v1)
            }
            switch "foo bar" {
            case "a":
                v1 -= v1
            case "\n\t":
                for fi14 := 0; fi14 < 11; fi14++ {
                    if (!true) {
                        v1 = 0.0
                    } else {
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            v2 -= (v2 / ((-1) | 1))
                        case "foo bar":
                            var v15 uint64 = 18446744073709551615
                            var v16 float32 = (((1.5 / (((v1 / (v1 * float32(-2.25))) * v1) + v1)) + 3.14159) / -2.25)
                            if ("a" < "hello") {
                                fmt.Println(v16)
                                v1 = v16
                            } else {
                                fmt.Println(v1, v1)
                                var v17 float64 = 1e10
                                if true {
                                    if true {
                                        fmt.Println(v1)
                                    }
                                    var v18 string = ("" + "\n\t")
                                    if true {
                                        var v19 uint64 = (v15 >> ((65535) & 63))
                                        switch int(v3) {
                                        case 0:
                                            var v20 uint32 = 153
                                            _ = v20
                                        case -1:
                                            fmt.Println(v3, v19)
                                        default:
                                            v15 += v15
                                        }
                                        v16 += (v1 - v16)
                                        _ = v19
                                    }
                                    var v21 float32 = float32(float64((v1 * ((v16 + v16) - (1e10 + v16)))))
                                    _ = v18
                                    _ = v21
                                }
                                if true {
                                    switch ("foo bar" + "foo bar") {
                                    case "foo bar":
                                        v4 = (fi14 - fi14)
                                    case "hello":
                                        v22 := GenS2{F0: 314, F1: uint(v3), F2: -128}
                                        _ = v22
                                    case "":
                                        var v23 uint32 = (v3 + (v3 % (((v3 / (((v3 - v3)) | 1))) | 1)))
                                        _ = v23
                                    }
                                    v24 := GenS2{F0: -532, F1: 2, F2: 101}
                                    _ = v24
                                }
                                _ = v17
                            }
                            switch ("a" + "") {
                            case "":
                                if (v1 <= (0.0 + v1)) {
                                    v25 := []int{-2147483648, 32767}
                                    v2 = (v2 / ((v2) | 1))
                                    _ = v25
                                } else {
                                    v26 := []int{fi14}
                                    _ = len(v26)
                                    _ = len(v26)
                                    v4 += fi12
                                    var v27 int32 = 231
                                    var v28 float64 = 3.14159
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                var v29 int32 = -2147483648
                                _ = v29
                            case "a":
                                switch 4294967295 {
                                case 32767:
                                    if true {
                                        v30 := GenS1{F0: 250, F1: (v15 % (((v15 / ((v15) | 1))) | 1))}
                                        v31 := GenS2{F0: int32(v15), F1: 2, F2: -1}
                                        _ = v30
                                        _ = v31
                                    }
                                case 128:
                                    v32 := []int{128, fi14}
                                    _ = v32
                                case 448:
                                    fmt.Println(v16, v1, v1)
                                    v3 -= v3
                                    var v33 float64 = 1e10
                                    _ = v33
                                default:
                                    v3 += v3
                                    v34 := GenS5{F0: (v3 + (v3 - (v3 >> ((2) & 63))))}
                                    if false {
                                        v2 = (v2 % (((v2 >> ((4294967295) & 63))) | 1))
                                        v35 := []int{(fi12 << ((65535) & 63))}
                                        _ = v35
                                    }
                                    _ = v34
                                }
                            case "hello":
                                v36 := []int{1, (-(fi14)), fi12}
                                fmt.Println(v2)
                                _ = v36
                            }
                            _ = v15
                            _ = v16
                        }
                        switch "hello" {
                        case "\n\t":
                            if false {
                                v1 -= v1
                                v37 := GenS6{F0: 42, F1: v3}
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v4, v1, v1)
                                case "foo bar":
                                    if false {
                                        v4 = (fi14 & v4)
                                        if (true && false) {
                                            v38 := GenS7{F0: 127, F1: v2, F2: "a"}
                                            v1 = v1
                                            var v39 int16 = 1
                                            var v40 uint = 0
                                            var v41 int32 = int32(v3)
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        } else {
                                            if (false || false) {
                                                v42 := []int{(fi12 >> ((65535) & 63)), fi12, -9223372036854775808, fi12}
                                                fmt.Println(fi12, fi12)
                                                _ = v42
                                            } else {
                                                switch 32768 {
                                                case -973:
                                                    var v43 uint32 = (v3 * (^(v3)))
                                                    _ = v43
                                                }
                                                v3 = (v3 * v3)
                                            }
                                            if true {
                                                if false {
                                                    v2 -= v2
                                                    var v44 bool = (128 >= 0)
                                                    _ = v44
                                                }
                                            }
                                            v45 := []int{fi12, (v4 << ((65535) & 63)), (fi12 ^ int(v2)), fi12}
                                            _ = v45
                                        }
                                    }
                                }
                                _ = v37
                            } else {
                                var v46 uint8 = 1
                                _ = v46
                            }
                        case "foo bar":
                            if false {
                                v47 := []int{(^(fi12)), fi12}
                                v48 := GenS4{F0: 1, F1: 58}
                                v49 := GenS4{F0: 2, F1: 127}
                                _ = v47
                                _ = v48
                                _ = v49
                            } else {
                                v2 -= v2
                            }
                        }
                    }
                    _ = fi14
                }
            case "foo bar":
                v50 := GenS4{F0: 494, F1: 39}
                _ = v50
            default:
                fmt.Println(v4, v4, v3)
            }
            _ = fi12
        }
        v51 := GenS8{F0: 1, F1: 65535, F2: 0}
        _ = v51
    }
    var v52 bool = true
    v53 := GenS2{F0: 0, F1: 769, F2: 11}
    fmt.Println(v2)
    fmt.Println(v3, v1)
    fmt.Println(v3, v3)
    if (-428 > -151) {
        if (true || (127 == 1)) {
            for fi54 := 0; fi54 < 20; fi54++ {
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 int16 = 474
                    if true {
                        fmt.Println(v1)
                    }
                    _ = fi55
                    _ = v56
                }
                _ = fi54
            }
            if true {
                v57 := []int{v4, (-(v4)), v4, v4}
                if v52 {
                    for fi58 := 0; fi58 < 10; fi58++ {
                        for fi59 := 0; fi59 < 17; fi59++ {
                            var v60 uint16 = 628
                            var v61 string = ("" + "hello")
                            var v62 int8 = -128
                            _ = fi59
                            _ = v60
                            _ = v61
                            _ = v62
                        }
                        v63 := GenS2{F0: -2147483648, F1: 411, F2: -1}
                        switch "foo bar" {
                        case "":
                            var v64 float64 = (float64(0.0) - 1e10)
                            _ = v64
                        }
                        _ = fi58
                        _ = v63
                    }
                }
                for fi65 := 0; fi65 < 8; fi65++ {
                    v4 = 128
                    _ = fi65
                }
                _ = v57
            } else {
                var v66 uint32 = 126
                _ = v66
            }
        }
        if (false || v52) {
            v4 -= -1
        } else {
            v52 = v52
        }
    } else {
        v67 := []int{v4, (v4 + v4), 32768, (v4 << ((199) & 63))}
        if len(v67) > 0 {
            v67[0] = v4
        }
        _ = v67
    }
    v68 := GenS1{F0: 223, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v68
    _ = v8
    _ = v9
}
