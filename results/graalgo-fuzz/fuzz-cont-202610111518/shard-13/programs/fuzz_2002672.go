package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int; F2 uint16 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int32 }
type GenS7 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS8 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0(p0 int32, p1 int8) uint64 {
    var v5 float32 = float32(((((0.0 + (0.0 / 1e10)) + 3.14159) + 0.1) / 3.14159))
    var v6 uint16 = 1
    var v7 int32 = p0
    fmt.Println(v7, p0, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int16) int32 {
    var v8 float64 = 1e10
    switch 42 {
    case 127:
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v8, p0, p0)
            _ = fi9
        }
    case -1:
        if (!false) {
            var v10 string = ("foo bar" + "hello")
            if false {
                v10 = v10
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch v10 {
                    case "foo bar":
                        if false {
                            v8 += v8
                            var v12 uint32 = 2147483648
                            _ = v12
                        } else {
                            p0 = 42
                        }
                    case "":
                        if (!(false || (!(32768 < fi11)))) {
                            v10 = "\n\t"
                            switch v10 {
                            case "\n\t":
                                if true {
                                    switch fi11 {
                                    case -479:
                                        var v13 bool = false
                                        _ = v13
                                    }
                                }
                            case "foo bar":
                                var v14 int64 = 191
                                _ = v14
                            default:
                                if (!(16 == 128)) {
                                    if true {
                                        if false {
                                            var v15 float64 = (v8 + (v8 - v8))
                                            for fi16 := 0; fi16 < 8; fi16++ {
                                                v15 -= float64(1e10)
                                                _ = fi16
                                            }
                                            var v17 int64 = 42
                                            _ = v15
                                            _ = v17
                                        } else {
                                            v18 := []int{fi11, 42, fi11}
                                            for fi19 := 0; fi19 < 13; fi19++ {
                                                if false {
                                                    fmt.Println(p0, fi11)
                                                    if false {
                                                        var v20 uint64 = 2
                                                        v21 := GenS1{F0: 1, F1: 4294967295}
                                                        var v22 int64 = 1
                                                        _ = v20
                                                        _ = v21
                                                        _ = v22
                                                    } else {
                                                        p0 -= (p0 % (((-(p0))) | 1))
                                                        v23 := GenS2{F0: 4294967295}
                                                        _ = v23
                                                    }
                                                    var v24 bool = (true && (1e10 <= 3.14159))
                                                    _ = v24
                                                } else {
                                                    if (-1 > 1) {
                                                        v25 := GenS3{F0: 1, F1: fi19, F2: 255}
                                                        fmt.Println(v10)
                                                        _ = v25
                                                    }
                                                    if (2147483647 == 2147483647) {
                                                        fmt.Println(p0)
                                                        if len(v18) > 0 {
                                                            v18[0] = fi11
                                                        }
                                                        fmt.Println(v8, fi11)
                                                    } else {
                                                        v26 := GenS4{F0: -2147483648, F1: v10}
                                                        p0 -= p0
                                                        p0 = 0
                                                        v27 := GenS2{F0: 4294967295}
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                                switch fi19 {
                                                case -712:
                                                    var v28 uint64 = 686
                                                    _ = v28
                                                }
                                                _ = fi19
                                            }
                                            if (true && (2 <= 2147483647)) {
                                                var v29 uint16 = 255
                                                _ = v29
                                            }
                                            _ = v18
                                        }
                                    }
                                    v30 := []int{(fi11 | (fi11 >> ((4294967295) & 63))), 0, (fi11 - (fi11 & fi11)), fi11}
                                    fmt.Println(v8)
                                    _ = v30
                                }
                            }
                        }
                    }
                    _ = fi11
                }
            } else {
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 uint16 = 1
                    _ = fi31
                    _ = v32
                }
                v10 = v10
            }
            _ = v10
        } else {
            var v33 int32 = 2147483647
            if ((!(255 >= 65535)) && (("a" + ((((("\n\t" + "hello") + "a") + "\n\t") + "a") + (((("" + ("foo bar" + "")) + ("foo bar" + "")) + "foo bar") + "hello"))) > "a")) {
                p0 += p0
                v34 := []int{42, 9223372036854775807}
                _ = v34
            }
            _ = v33
        }
    case 128:
        v35 := GenS1{F0: 1, F1: 2147483648}
        _ = v35
    default:
        switch 127 {
        case 981:
            v8 += (v8 / (float64((1.5 - 3.14159)) - (v8 * -2.25)))
        case 42:
            var v36 bool = false
            _ = v36
        }
    }
    fmt.Println(p0, v8)
    v37 := []int{-1, 128}
    _ = p0
    _ = v37
    _ = v8
    return 982
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 0
    var v3 float64 = v1
    v4 := []int{-239}
    if len(v4) > 0 {
        v4[0] = -334
    }
    var v38 uint64 = genFn0(0, 42)
    v39 := GenS3{F0: 42, F1: 1, F2: 1}
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 int32 = 0
        var v42 int64 = -355
        var v43 uint32 = uint32(v38)
        _ = fi40
        _ = v41
        _ = v42
        _ = v43
    }
    fmt.Println(v1, v3)
    fmt.Println(v2)
    v1 = (v3 * v3)
    v44 := GenS5{F0: 735}
    if (-58 <= 42) {
        switch -778 {
        case 42:
            switch "a" {
            case "a":
                v1 = 3.14159
            default:
                switch 2147483647 {
                case -941:
                    var v45 int16 = -1
                    _ = v45
                case -1:
                    v38 = v38
                }
            }
        case -1:
            if (2147483647 != 42) {
                var v46 uint = 2
                for fi47 := 0; fi47 < 9; fi47++ {
                    v2 += v2
                    _ = len(v4)
                    _ = fi47
                }
                for fi48 := 0; fi48 < 15; fi48++ {
                    v3 = 0.1
                    _ = fi48
                }
                _ = v46
            } else {
                v2 -= (v2 - v2)
            }
        case 656:
            var v49 int64 = -1
            _ = v49
        default:
            fmt.Println(v2, v1)
        }
        v50 := GenS3{F0: -1, F1: 4294967295, F2: 65535}
        _ = v50
    }
    v51 := GenS6{F0: 1, F1: 0}
    if len(v4) > 0 {
        v4[0] = 1
    }
    if true {
        for fi52 := 0; fi52 < 15; fi52++ {
            for fi53 := 0; fi53 < 11; fi53++ {
                var v54 float32 = 0.0
                v2 = v2
                if (917 < -32768) {
                    v2 = 128
                    v55 := GenS2{F0: genFn0(-976, 127)}
                    v56 := GenS7{F0: 65535, F1: 1000, F2: 673}
                    _ = v55
                    _ = v56
                } else {
                    v57 := GenS5{F0: 42}
                    v38 = 42
                    _ = v57
                }
                _ = fi53
                _ = v54
            }
            var v58 int16 = 1000
            fmt.Println(v3, v2)
            _ = fi52
            _ = v58
        }
    }
    v2 += v2
    if false {
        v59 := GenS8{F0: 2, F1: 441, F2: uint32(v38)}
        _ = v59
    } else {
        if len(v4) > 0 {
            v4[0] = -256
        }
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        for fi61 := 0; fi61 < 3; fi61++ {
            var v62 float32 = (((-2.25 * 0.1) * 1.5) + ((float32((float64(-2.25) * (v1 + (((v1 / ((1.5 / (1.5 + v3)) + float64(1.5))) + v3) * ((float64(-2.25) - 0.0) / (float64(3.14159) * v3)))))) * 1.5) + 0.0))
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    switch "\n\t" {
    case "":
        v2 = (v2 % (((v2 * (v2 & v2))) | 1))
    case "hello":
        var v63 string = (("foo bar" + "a") + ("hello" + "foo bar"))
        _ = v63
    case "a":
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    v64 := GenS7{F0: 2, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v51
    _ = v64
}
