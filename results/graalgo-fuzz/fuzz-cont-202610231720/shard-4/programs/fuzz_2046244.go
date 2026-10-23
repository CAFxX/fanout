package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
func genFn0(p0 uint8, p1 uint, p2 uint64) string {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 953}
        v8 := GenS2{F0: p2, F1: 0}
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v11 uint = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return "\n\t"
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 898
    var v3 bool = (42 > 1)
    var v4 float32 = (((1.5 + (-2.25 / (0.1 * 3.14159))) * (float32(-2.25) - 0.1)) + (1e10 / ((1.5 + float32(0.0)) + 1.5)))
    v5 := []uint32{}
    v5 = append(v5, 0)
    v12 := GenS2{F0: 42, F1: 42}
    if len(v5) > 0 {
        v5[0] = 2
    }
    fmt.Println(v2)
    v13 := GenS1{F0: 0}
    _ = len(v5)
    if (-566 != 2147483647) {
        var v14 string = ("\n\t" + ("" + "foo bar"))
        if v3 {
            for fi15 := 0; fi15 < 3; fi15++ {
                if v3 {
                    var v16 int = fi15
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 float64 = (float64(v4) * ((0.1 * 0.0) * 3.14159))
                        var v19 float32 = v4
                        if v3 {
                            v18 += (v18 / v18)
                            switch v14 {
                            case "foo bar":
                                v1 = 679
                            case "hello":
                                if false {
                                    if len(v5) > 0 {
                                        v5[0] = 732
                                    }
                                    var v20 uint64 = 0
                                    v21 := GenS1{F0: 32767}
                                    v1 = (^(v1))
                                    v19 -= ((v19 - float32(v18)) + v4)
                                    _ = v20
                                    _ = v21
                                }
                            case "\n\t":
                                var v22 string = ("hello" + ("\n\t" + v14))
                                _ = v22
                            default:
                                _ = len(v5)
                            }
                            v3 = (42 > 127)
                        } else {
                            v23 := GenS2{F0: 827, F1: 100}
                            if v3 {
                                fmt.Println(v14)
                                if v3 {
                                    v24 := GenS2{F0: 42, F1: 119}
                                    v18 += float64(v4)
                                    v25 := GenS2{F0: 0, F1: -128}
                                    switch v16 {
                                    case -1:
                                        fmt.Println(v16, v2, v2)
                                    case 0:
                                        switch v16 {
                                        case 4294967295:
                                            fmt.Println(fi15)
                                            v19 = (3.14159 + (float32(float64((float32((v18 - v18)) / v4))) * v4))
                                            var v26 int64 = 1
                                            var v27 uint64 = 1
                                            _ = v26
                                            _ = v27
                                        case -2147483648:
                                            var v28 int64 = -843
                                            _ = v28
                                        default:
                                            _ = len(v5)
                                            var v29 int64 = 42
                                            _ = v29
                                        }
                                    case -9223372036854775808:
                                        var v30 int16 = 1000
                                        _ = v30
                                    default:
                                        _ = len(v5)
                                    }
                                    if v3 {
                                        var v31 int = fi17
                                        if false {
                                            v2 = (v2 >> ((v2) & 63))
                                            v32 := GenS1{F0: 1}
                                            var v33 uint8 = 42
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 float64 = (((v18 - 3.14159) / 0.1) / 1.0)
                                            if false {
                                                _ = len(v5)
                                            }
                                            _ = v34
                                        }
                                        fmt.Println(v19, v14, fi17)
                                        v2 = v2
                                        _ = v31
                                    }
                                    var v35 uint = (v2 + v2)
                                    v36 := GenS1{F0: -862}
                                    _ = v24
                                    _ = v25
                                    _ = v35
                                    _ = v36
                                }
                                _ = len(v5)
                                v37 := GenS2{F0: 975, F1: 100}
                                _ = v37
                            }
                            _ = v23
                        }
                        v38 := GenS1{F0: 32767}
                        _ = fi17
                        _ = v18
                        _ = v19
                        _ = v38
                    }
                    _ = v16
                }
                _ = fi15
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            fmt.Println(v2, v4)
        }
        switch v14 {
        case "":
            v39 := GenS2{F0: 930, F1: -46}
            _ = v39
        case "foo bar":
            v40 := GenS2{F0: 2, F1: 100}
            _ = v40
        case "hello":
            for fi41 := 0; fi41 < 4; fi41++ {
                for fi42 := 0; fi42 < 5; fi42++ {
                    if v3 {
                        if (-1 != -486) {
                            v1 = (^(v1))
                            v43 := GenS3{F0: 4294967295, F1: 127, F2: 128}
                            var v44 int8 = 0
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(v14, v14)
                        }
                    }
                    _ = fi42
                }
                v4 += v4
                v45 := GenS1{F0: 42}
                _ = fi41
                _ = v45
            }
        }
        _ = v14
    }
    if (0.0 < (((-2.25 * 1e10) + 0.1) - 0.0)) {
        if (!((v2 * v2) < (v2 / (((v2 - 1)) | 1)))) {
            v4 = (v4 + (float32(float64(v4)) - ((float32(float64(0.1)) * v4) - (3.14159 * v4))))
            v1 = v1
        }
    }
    var v46 int64 = 42
    v1 = (v1 & (^(v1)))
    var v47 uint = v2
    var v48 uint64 = 0
    switch "a" {
    case "foo bar":
        if true {
            v47 = v2
            fmt.Println(v3, v1, v2)
        } else {
            switch "\n\t" {
            case "\n\t":
                var v49 int64 = v46
                _ = v49
            case "foo bar":
                if (((42 >= (v1 / (((^(v1))) | 1))) && v3) && true) {
                    if true {
                        v47 += (v2 & 255)
                        v2 -= (v47 ^ (v2 << (((v47 << ((255) & 63))) & 63)))
                        v50 := GenS3{F0: (^(v48)), F1: 100, F2: 110}
                        _ = v50
                    }
                }
            }
            switch "hello" {
            case "a":
                var v51 uint32 = 65535
                _ = v51
            case "":
                if v3 {
                    v46 -= (v46 << ((v47) & 63))
                    v2 -= v47
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v52 := GenS2{F0: v48, F1: 127}
                    _ = v52
                }
            case "\n\t":
                switch "a" {
                case "\n\t":
                    v53 := GenS4{F0: 444, F1: 17, F2: 0}
                    _ = v53
                default:
                    var v54 float32 = (v4 * float32(((0.0 / 1e10) / 3.14159)))
                    _ = v54
                }
            default:
                switch "" {
                case "foo bar":
                    v5 = append(v5, uint32(v47))
                case "a":
                    switch -918 {
                    case 128:
                        fmt.Println(v3, v1)
                    case -426:
                        if v3 {
                            v3 = ((("a" + ("foo bar" + ("a" + "\n\t"))) + ("\n\t" + "foo bar")) == (("a" + ("foo bar" + (("a" + "hello") + "hello"))) + ("\n\t" + "\n\t")))
                            var v55 uint16 = 32767
                            _ = v55
                        } else {
                            v2 = (^(v47))
                            v2 = v2
                            var v56 int8 = 59
                            _ = v56
                        }
                    case -2147483648:
                        v1 += v1
                    }
                }
            }
            fmt.Println(v47, v1, v4)
        }
    case "\n\t":
        var v57 uint = 2
        _ = v57
    case "":
        if (0 == 127) {
            v46 -= v46
            v4 += ((0.0 / v4) / 0.1)
        } else {
            v3 = (0 <= (^(v1)))
            for fi58 := 0; fi58 < 3; fi58++ {
                v59 := GenS1{F0: 741}
                var v60 uint8 = 255
                _ = fi58
                _ = v59
                _ = v60
            }
        }
    }
    v46 = (v46 / (((v46 * (v46 & -9223372036854775808))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
}
