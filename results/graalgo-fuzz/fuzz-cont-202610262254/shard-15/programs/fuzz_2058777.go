package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int32 }
type GenS5 struct { F0 int8; F1 uint8; F2 int }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int64; F1 uint16; F2 int16 }
type GenS8 struct { F0 int }
func genFn0(p0 uint32, p1 string, p2 int64) int16 {
    var v4 uint32 = (p0 * (p0 / (((p0 | (^(p0)))) | 1)))
    v5 := GenS1{F0: v4, F1: 1, F2: (^(p2))}
    fmt.Println(p0, p2, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int16, p1 uint8) int8 {
    var v6 float32 = 0.1
    fmt.Println(v6, p0, v6)
    var v7 float32 = v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}

func main() {
    v1 := []int{-762}
    v2 := []int{}
    v3 := []string{"a", "foo bar"}
    var v8 uint = 2
    if len(v2) > 0 {
        v2[0] = 127
    }
    v8 = 6
    v8 -= 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS1{F0: 2147483648, F1: 42, F2: 1}
    switch "foo bar" {
    case "hello":
        for fi10 := 0; fi10 < 18; fi10++ {
            v8 = (v8 * v8)
            for fi11 := 0; fi11 < 12; fi11++ {
                v8 = v8
                if false {
                    var v12 uint = v8
                    v8 -= v12
                    _ = v12
                } else {
                    v13 := GenS1{F0: 42, F1: -32768, F2: -834}
                    switch "foo bar" {
                    case "hello":
                        v14 := GenS2{F0: 42, F1: 0, F2: -1}
                        _ = v14
                    }
                    var v15 int8 = -57
                    _ = len(v2)
                    _ = v13
                    _ = v15
                }
                var v16 uint32 = 2
                _ = fi11
                _ = v16
            }
            if false {
                if false {
                    fmt.Println(v8, v8)
                    var v17 bool = ("foo bar" < "foo bar")
                    _ = v17
                }
                if true {
                    v18 := GenS1{F0: 65535, F1: -745, F2: -1}
                    if true {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            var v20 string = ((("" + "hello") + "\n\t") + ("foo bar" + ("foo bar" + ("a" + "hello"))))
                            if true {
                                v21 := GenS3{F0: 4294967295, F1: 32767}
                                if (false || false) {
                                    if (fi19 != fi19) {
                                        v22 := GenS3{F0: 347, F1: -574}
                                        switch 32767 {
                                        case 2147483647:
                                            v23 := GenS4{F0: 170, F1: -2147483648, F2: 2147483647}
                                            _ = v23
                                        case -102:
                                            var v24 uint16 = 255
                                            _ = v24
                                        default:
                                            v25 := GenS1{F0: 1, F1: 1, F2: 769}
                                            _ = v25
                                        }
                                        var v26 int32 = 0
                                        _ = v22
                                        _ = v26
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if (!((304 >= 1) || true)) {
                                        switch (^(fi19)) {
                                        case 32768:
                                            v8 = (v8 / (((v8 >> ((227) & 63))) | 1))
                                        default:
                                            v27 := GenS5{F0: 0, F1: 0, F2: fi10}
                                            _ = v27
                                        }
                                    } else {
                                        v28 := GenS6{F0: 42, F1: (fi10 - (^(fi10)))}
                                        var v29 uint8 = uint8(v8)
                                        _ = v28
                                        _ = v29
                                    }
                                } else {
                                    if (!true) {
                                        v8 += (v8 - v8)
                                    }
                                    v8 += (v8 + (v8 + (v8 / (((v8 | (v8 * (^(v8))))) | 1))))
                                    switch fi19 {
                                    case -2147483648:
                                        v8 -= 931
                                    default:
                                        v20 = (v20 + "")
                                    }
                                }
                                _ = v21
                            }
                            fmt.Println(v20, fi19, fi10)
                            _ = fi19
                            _ = v20
                        }
                    } else {
                        v30 := GenS7{F0: 42, F1: 32767, F2: 42}
                        for fi31 := 0; fi31 < 19; fi31++ {
                            var v32 uint32 = 562
                            var v33 float32 = (1e10 * 0.0)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        _ = v30
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi10 | fi10)
                    }
                    _ = v18
                } else {
                    if (false && false) {
                        var v34 uint16 = 633
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v8 = (^(v8))
                            _ = fi35
                        }
                        v34 += 582
                        _ = v34
                    }
                    v36 := GenS4{F0: 32767, F1: 2147483647, F2: 42}
                    if false {
                        var v37 float32 = (-2.25 * 3.14159)
                        var v38 uint16 = 0
                        if true {
                            var v39 uint16 = (v38 ^ (v38 * v38))
                            var v40 int8 = 9
                            _ = v39
                            _ = v40
                        }
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
                v1 = append(v1, fi10)
            } else {
                var v41 string = ("\n\t" + "foo bar")
                _ = v41
            }
            _ = fi10
        }
    case "":
        fmt.Println(v8)
    default:
        fmt.Println(v8)
    }
    fmt.Println(v8)
    v3 = append(v3, "")
    switch -726 {
    case 991:
        var v42 uint64 = 18446744073709551615
        _ = v42
    case 861:
        if false {
            switch "a" {
            case "":
                var v43 uint = (v8 ^ (v8 + (v8 * 255)))
                _ = v43
            case "\n\t":
                var v44 int = 128
                _ = v44
            default:
                v45 := GenS2{F0: 0, F1: -2147483648, F2: -1}
                _ = v45
            }
            for fi46 := 0; fi46 < 3; fi46++ {
                if (42 >= 18446744073709551615) {
                    v8 -= v8
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if ((0.1 > ((float64(0.1) + (((1e10 - 1e10) / (((3.14159 + 1.5) * (0.0 / 1.5)) - 3.14159)) - (-2.25 - 1e10))) + 1.5)) || (-225 == 2147483647)) {
                        v8 -= (v8 * v8)
                        if (!(((-2.25 * 0.1) + (1.5 - 3.14159)) <= ((((float32((1e10 + (1e10 + (((float64(float32((((((0.0 - 0.1) + float64(0.1)) + -2.25) + ((1e10 - 0.1) * ((0.0 * ((1.5 + 1.5) * 1e10)) * 0.1))) + (float64((0.0 * (1e10 - 1.5))) + (1e10 + 1e10))))) - -2.25) + 1.5) / 1e10)))) * -2.25) - float32((float64(0.0) * (-2.25 * (1e10 / 1.5))))) - (0.0 + 3.14159)) - 0.1))) {
                            fmt.Println(fi46)
                        } else {
                            switch "\n\t" {
                            case "a":
                                switch "foo bar" {
                                case "\n\t":
                                    for fi47 := 0; fi47 < 3; fi47++ {
                                        var v48 int32 = 1
                                        var v49 uint = 18446744073709551615
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                case "a":
                                    for fi50 := 0; fi50 < 7; fi50++ {
                                        _ = len(v3)
                                        _ = fi50
                                    }
                                }
                            default:
                                var v51 uint8 = 0
                                _ = v51
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            var v52 int8 = 23
                            _ = v52
                        case "\n\t":
                            var v53 int8 = 100
                            _ = v53
                        case "hello":
                            var v54 uint32 = 2147483647
                            _ = v54
                        }
                    } else {
                        for fi55 := 0; fi55 < 5; fi55++ {
                            v1 = append(v1, fi46)
                            switch "\n\t" {
                            case "a":
                                v8 = (v8 ^ v8)
                            default:
                                if len(v2) > 0 {
                                    v2[0] = fi55
                                }
                            }
                            var v56 int = (fi55 / ((fi55) | 1))
                            _ = fi55
                            _ = v56
                        }
                        v8 = v8
                        var v57 uint8 = 0
                        _ = v57
                    }
                    var v58 int16 = 1000
                    var v59 uint16 = 424
                    _ = v58
                    _ = v59
                }
                v60 := GenS8{F0: fi46}
                var v61 int16 = genFn0(2147483648, "foo bar", 42)
                _ = fi46
                _ = v60
                _ = v61
            }
        }
    case 4294967295:
        v8 = (v8 * 42)
    }
    var v62 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v8
    _ = v9
}
