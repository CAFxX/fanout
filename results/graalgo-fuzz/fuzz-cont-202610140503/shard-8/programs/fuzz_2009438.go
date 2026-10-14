package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint32; F2 int }
func genFn0(p0 bool, p1 uint, p2 uint) string {
    fmt.Println(p2)
    if p0 {
        v6 := []int{0}
        p2 -= p2
        fmt.Println(p1, p0, p2)
        _ = v6
    } else {
        fmt.Println(p2, p2)
        fmt.Println(p0, p1, p1)
    }
    v7 := GenS1{F0: 1, F1: (p2 > (p2 >> (((p2 + p1)) & 63)))}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ((("a" + "foo bar") + "foo bar") + "\n\t")
}
func genFn1(p0 uint32) uint16 {
    var v8 uint8 = 0
    v9 := GenS1{F0: -1, F1: (!false)}
    var v10 int = 4294967295
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 string = ""
    var v2 uint32 = 895
    var v3 uint32 = (v2 | 65535)
    var v4 bool = true
    var v5 uint = 984
    v3 = v3
    v1 = "foo bar"
    v11 := []int{1, -9223372036854775808}
    for fi12 := 0; fi12 < 2; fi12++ {
        if v4 {
            v13 := GenS2{F0: 18446744073709551615, F1: 1}
            fmt.Println(v4, v5, fi12)
            _ = len(v11)
            _ = v13
        } else {
            var v14 int16 = 0
            var v15 int16 = (v14 % (((v14 ^ v14)) | 1))
            _ = v14
            _ = v15
        }
        var v16 int64 = -572
        var v17 uint = v5
        _ = fi12
        _ = v16
        _ = v17
    }
    if false {
        v5 -= v5
        switch (((v1 + ("" + v1)) + ((("hello" + v1) + v1) + ("a" + v1))) + (((v1 + (("\n\t" + v1) + "hello")) + "a") + v1)) {
        case "":
            if v4 {
                var v18 int32 = -1
                v18 -= v18
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS3{F0: 84}
                    if (v18 <= 418) {
                        switch v1 {
                        case "":
                            var v21 uint = (v5 % (((v5 & 65535)) | 1))
                            _ = v21
                        case "hello":
                            for fi22 := 0; fi22 < 17; fi22++ {
                                fmt.Println(v1, v1)
                                v11 = append(v11, -506)
                                fmt.Println(v18)
                                _ = fi22
                            }
                        default:
                            fmt.Println(v2, fi19)
                        }
                        for fi23 := 0; fi23 < 10; fi23++ {
                            var v24 uint16 = 737
                            var v25 int32 = v18
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := GenS3{F0: 148}
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        var v28 uint32 = v3
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v4 = false
                        switch v1 {
                        case "":
                            switch fi29 {
                            case 439:
                                var v30 float64 = 1.5
                                _ = v30
                            case 22:
                                if false {
                                    switch ((v1 + (v1 + ("foo bar" + v1))) + "\n\t") {
                                    case "hello":
                                        if v4 {
                                            var v31 uint16 = genFn1(v2)
                                            _ = v31
                                        }
                                    case "\n\t":
                                        v5 = 788
                                    default:
                                        fmt.Println(v2)
                                    }
                                    if v4 {
                                        var v32 int = fi19
                                        _ = v32
                                    }
                                } else {
                                    v3 -= 662
                                    v3 = v2
                                    fmt.Println(v2, v1, v2)
                                }
                            case 9223372036854775807:
                                v18 = v18
                            }
                        case "a":
                            v2 -= (v2 | v2)
                        case "\n\t":
                            var v33 float32 = 3.14159
                            _ = v33
                        default:
                            if (!(68 > 2)) {
                                v34 := GenS3{F0: 128}
                                switch v1 {
                                case "a":
                                    fmt.Println(v5, v4, fi19)
                                case "hello":
                                    v1 = "foo bar"
                                case "foo bar":
                                    v5 -= (v5 + 383)
                                default:
                                    v5 = 18446744073709551615
                                }
                                _ = v34
                            } else {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                v35 := GenS1{F0: 1, F1: v4}
                                _ = v35
                            }
                        }
                        v4 = (((((("foo bar" + v1) + v1) + ("" + v1)) + v1) < "") || false)
                        fmt.Println(v1, v5)
                        _ = fi29
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                switch v1 {
                case "a":
                    v36 := GenS1{F0: int16(v3), F1: (v5 != (v5 % ((v5) | 1)))}
                    _ = v36
                default:
                    switch (("" + "") + v1) {
                    case "":
                        v2 -= (^(v3))
                    }
                }
                for fi37 := 0; fi37 < 13; fi37++ {
                    v38 := GenS4{F0: -2147483648, F1: 639, F2: 100}
                    v2 += v2
                    _ = fi37
                    _ = v38
                }
            }
        case "a":
            var v39 uint32 = v2
            _ = v39
        default:
            var v40 int = 42
            _ = v40
        }
    }
    if ((1.5 >= (((0.1 + (3.14159 / (float64((0.1 * 3.14159)) + ((1.5 * 1.5) + 0.1)))) - (float64((1.5 / 0.1)) - 1e10)) / 1.5)) && (((v1 + (("foo bar" + (v1 + v1)) + v1)) < v1) || (0 >= 827))) {
        switch 128 {
        case 1:
            var v41 int64 = -9223372036854775808
            _ = v41
        }
        if (278 != uint64(v3)) {
            if (2 <= 1) {
                var v42 uint = v5
                _ = v42
            } else {
                v11 = append(v11, 1)
                var v43 int32 = -2147483648
                v44 := GenS5{F0: uint32(v5), F1: v4, F2: 127}
                _ = v43
                _ = v44
            }
            fmt.Println(v2, v4)
            var v45 float32 = 1.5
            _ = v45
        }
        if (104 >= 83) {
            switch 127 {
            case 1:
                fmt.Println(v2, v2, v2)
            case -836:
                var v46 uint32 = 4294967295
                _ = v46
            case 12:
                fmt.Println(v3)
            }
            v47 := GenS3{F0: 1}
            _ = v47
        } else {
            var v48 bool = (1e10 > float32(float64(1e10)))
            var v49 int16 = -1
            _ = v48
            _ = v49
        }
    }
    var v50 float32 = 3.14159
    v2 -= 42
    v4 = false
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v2)
    case "":
        _ = len(v11)
    case "a":
        v51 := GenS4{F0: -2147483648, F1: 32767, F2: -1}
        _ = v51
    default:
        _ = len(v11)
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if (("foo bar" + "foo bar") == (v1 + (v1 + ("\n\t" + v1)))) {
        switch 2147483647 {
        case 1:
            if v4 {
                if v4 {
                    if (-9223372036854775808 <= 1) {
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v53 := GenS2{F0: 2, F1: 18446744073709551615}
                            v2 = uint32(v5)
                            v54 := GenS4{F0: 2147483647, F1: 32767, F2: -128}
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                        fmt.Println(v1, v5)
                        for fi55 := 0; fi55 < 16; fi55++ {
                            if (v3 <= v3) {
                                v1 = v1
                                for fi56 := 0; fi56 < 3; fi56++ {
                                    var v57 uint8 = 2
                                    var v58 uint16 = 32767
                                    if false {
                                        if len(v11) > 0 {
                                            v11[0] = fi55
                                        }
                                        fmt.Println(fi56)
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        switch ("a" + v1) {
                                        case "foo bar":
                                            v59 := GenS2{F0: 4294967295, F1: 0}
                                            _ = v59
                                        case "\n\t":
                                            var v60 uint16 = v58
                                            _ = v60
                                        default:
                                            v61 := GenS4{F0: -1, F1: v58, F2: 42}
                                            _ = v61
                                        }
                                    }
                                    v62 := GenS6{F0: (v5 * (v5 & (v5 * v5))), F1: uint8(v5), F2: v57}
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                    _ = v62
                                }
                                var v63 float64 = (1.5 - ((((float64(v50) / 0.1) - (3.14159 - ((1.5 * ((0.0 - 0.0) / float64((3.14159 / float32(((float64(v50) + ((0.0 - 1e10) - -2.25)) / 3.14159)))))) / 3.14159))) * 1e10) - ((1.5 - float64(v50)) - -2.25)))
                                _ = v63
                            }
                            _ = fi55
                        }
                    }
                    v64 := GenS6{F0: (v5 & uint(v3)), F1: 56, F2: 0}
                    _ = v64
                } else {
                    fmt.Println(v2)
                    v65 := GenS7{F0: -139, F1: v1}
                    if len(v11) > 0 {
                        v11[0] = 32767
                    }
                    _ = v65
                }
                v3 = (^(v2))
            }
        case 42:
            for fi66 := 0; fi66 < 2; fi66++ {
                for fi67 := 0; fi67 < 16; fi67++ {
                    if v4 {
                        var v68 int16 = -604
                        _ = v68
                    }
                    _ = fi67
                }
                fmt.Println(v5, v5)
                v1 = v1
                _ = fi66
            }
        default:
            v3 += 252
        }
        v1 = (((v1 + v1) + v1) + "")
    }
    fmt.Println(v1, v4)
    v69 := GenS8{F0: v2, F1: 65535, F2: 100}
    switch v1 {
    case "\n\t":
        var v70 uint64 = 18446744073709551615
        _ = v70
    case "hello":
        var v71 float32 = v50
        _ = v71
    case "":
        if v4 {
            v3 += 2147483648
        }
    }
    v72 := GenS9{F0: 2, F1: (v2 * 2), F2: 811}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v69
    _ = v72
}
