package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32 }

func main() {
    var v1 uint8 = 29
    var v2 uint64 = 4294967295
    var v3 uint8 = v1
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 4; fi5++ {
        fmt.Println(v3)
        fmt.Println(fi5, v3, v3)
        if false {
            var v6 uint32 = 2147483647
            _ = v6
        } else {
            v7 := GenS1{F0: (v2 % (((v2 - 1)) | 1))}
            _ = v7
        }
        _ = fi5
    }
    if ((1.5 + (1.5 * float64((((((-2.25 - 1.5) * ((1e10 + 1e10) * (1e10 * ((-2.25 + float32(1e10)) / 0.1)))) / 0.1) + (1.5 + (1.5 * -2.25))) / float32(1e10))))) <= (0.1 / (float64(1.5) * 0.1))) {
        v8 := GenS1{F0: (v2 * 4294967295)}
        v1 = v3
        _ = v8
    } else {
        if true {
            fmt.Println(v3, v1)
        } else {
            v2 += uint64(v1)
        }
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                v2 = 1
                switch fi9 {
                case 32767:
                    switch "hello" {
                    case "hello":
                        delete(v4, (("hello" + ("a" + "hello")) + "hello"))
                    case "a":
                        v10 := GenS1{F0: 418}
                        _ = v10
                    case "":
                        if (true || false) {
                            v3 = v1
                            v11 := GenS1{F0: 229}
                            _ = v11
                        } else {
                            var v12 uint16 = 435
                            var v13 int16 = -1
                            v14 := GenS2{F0: fi9, F1: (v3 / ((v1) | 1))}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    default:
                        var v15 uint16 = 447
                        _ = v15
                    }
                case -9223372036854775808:
                    v16 := []int{fi9, (^(fi9)), 32767}
                    _ = v16
                }
                v17 := GenS3{F0: -2147483648, F1: 128}
                _ = v17
            } else {
                v4["hello"] = (2 == 65535)
            }
            _ = len(v4)
            _ = fi9
        }
        var v18 int64 = 0
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            switch "a" {
            case "\n\t":
                if (1 > fi20) {
                    if (!true) {
                        fmt.Println(fi20)
                    }
                }
            case "a":
                switch fi19 {
                case 32768:
                    v3 = v1
                case 0:
                    v21 := GenS2{F0: fi20, F1: 255}
                    _ = v21
                default:
                    switch (^(fi20)) {
                    case -727:
                        fmt.Println(v2)
                        v22 := []int{(fi20 % ((-574) | 1)), (fi20 / (((fi19 * (fi20 * (fi20 | 377)))) | 1))}
                        delete(v4, ("a" + "foo bar"))
                        _ = v22
                    case -97:
                        var v23 uint32 = 0
                        _ = v23
                    default:
                        v24 := GenS3{F0: 42, F1: (v3 / (((^(v3))) | 1))}
                        _ = v24
                    }
                }
            case "":
                v2 = (v2 & 2)
            default:
                var v25 uint8 = v3
                _ = v25
            }
            var v26 float64 = float64(((0.0 + (1e10 * (-2.25 * 1e10))) * (3.14159 - float32(0.1))))
            switch "\n\t" {
            case "\n\t":
                v27 := GenS4{F0: v2, F1: "\n\t"}
                _ = v27
            case "foo bar":
                v28 := []int{fi19, fi19, fi19}
                _ = v28
            case "a":
                v29 := GenS2{F0: fi19, F1: (v1 * v3)}
                _ = v29
            }
            _ = fi20
            _ = v26
        }
        if ((-32768 <= int16(v2)) || (-2.25 < 1e10)) {
            v30 := v4[("hello" + "a")]
            _ = v30
        } else {
            v2 = 18446744073709551615
            var v31 int8 = -128
            var v32 int32 = -1
            _ = v31
            _ = v32
        }
        _ = fi19
    }
    fmt.Println(v1, v2)
    v33 := GenS4{F0: 0, F1: ""}
    v3 -= (v3 & (v1 * v1))
    var v34 float32 = 0.0
    v1 += v1
    v4[("hello" + "\n\t")] = (((0.1 - float64(1.5)) > (0.0 + float64(3.14159))) || true)
    v35 := GenS5{F0: 837, F1: 0, F2: 245}
    v3 = (v1 / (((v1 & 255)) | 1))
    v36 := GenS5{F0: 2, F1: 1, F2: 0}
    if false {
        v37 := v4[""]
        v3 = 152
        v38 := GenS6{F0: 42}
        _ = v37
        _ = v38
    } else {
        var v39 int = -597
        if ((("foo bar" >= ("a" + "")) && false) && (-32768 <= 0)) {
            for fi40 := 0; fi40 < 4; fi40++ {
                if (!((v34 * (float32(1.5) + v34)) <= float32(-2.25))) {
                    v4["foo bar"] = true
                    var v41 uint16 = 474
                    _ = v41
                } else {
                    fmt.Println(v39)
                }
                if ((1 >= -729) && false) {
                    v42 := v4["\n\t"]
                    _ = v42
                }
                v39 = 9223372036854775807
                _ = fi40
            }
            for fi43 := 0; fi43 < 13; fi43++ {
                v44 := []int{fi43, fi43, v39}
                v45 := GenS7{F0: 808, F1: -128}
                if (793 > 0) {
                    var v46 uint8 = v3
                    var v47 float32 = 1.5
                    fmt.Println(v46, fi43, v3)
                    _ = v46
                    _ = v47
                }
                _ = fi43
                _ = v44
                _ = v45
            }
            for fi48 := 0; fi48 < 15; fi48++ {
                switch fi48 {
                case -1:
                    v49 := []int{(fi48 >> ((255) & 63)), v39}
                    _ = v49
                default:
                    v50 := []int{128, (v39 / ((9223372036854775807) | 1)), -897}
                    _ = v50
                }
                _ = fi48
            }
        }
        if true {
            if (0 <= 65535) {
                if true {
                    var v51 float32 = 1e10
                    var v52 int32 = 0
                    _ = v51
                    _ = v52
                }
                fmt.Println(v1)
                if false {
                    var v53 uint = 701
                    _ = v53
                } else {
                    if true {
                        switch 1 {
                        case 32768:
                            var v54 float32 = (1e10 * float32((float64(0.1) - ((float64(((v34 - v34) - (3.14159 * (float32(0.0) - 0.0)))) + 1e10) - 0.1))))
                            _ = v54
                        case -1:
                            var v55 uint32 = 42
                            _ = v55
                        case -729:
                            var v56 uint64 = 2
                            _ = v56
                        default:
                            v1 += (v3 & (v3 ^ (^(v1))))
                        }
                        v4[("foo bar" + "foo bar")] = false
                    } else {
                        switch "a" {
                        case "foo bar":
                            for fi57 := 0; fi57 < 3; fi57++ {
                                for fi58 := 0; fi58 < 2; fi58++ {
                                    v59 := GenS7{F0: 0, F1: -64}
                                    var v60 uint32 = 4294967295
                                    var v61 string = ("" + "hello")
                                    v62 := GenS8{F0: v3}
                                    _ = fi58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                }
                                if false {
                                    v63 := []int{v39}
                                    v64 := GenS5{F0: 58, F1: -2147483648, F2: 65535}
                                    v39 = (^(fi57))
                                    _ = v63
                                    _ = v64
                                }
                                if (!false) {
                                    v65 := GenS7{F0: 769, F1: 1}
                                    for fi66 := 0; fi66 < 12; fi66++ {
                                        v34 -= (float32(1e10) - v34)
                                        switch 0 {
                                        case 2147483647:
                                            _ = len(v4)
                                        default:
                                            var v67 float32 = (1.5 + (v34 - 0.1))
                                            _ = v67
                                        }
                                        var v68 int8 = -26
                                        _ = fi66
                                        _ = v68
                                    }
                                    _ = v65
                                }
                                _ = fi57
                            }
                        }
                    }
                    fmt.Println(v1, v39)
                }
            }
        } else {
            for fi69 := 0; fi69 < 10; fi69++ {
                for fi70 := 0; fi70 < 20; fi70++ {
                    v3 -= v3
                    v71 := []int{-2147483648}
                    _ = fi70
                    _ = v71
                }
                var v72 int = v39
                for fi73 := 0; fi73 < 15; fi73++ {
                    if (2 == 1) {
                        v72 -= -1
                    } else {
                        switch "\n\t" {
                        case "foo bar":
                            if (1 < 65535) {
                                v74 := GenS9{F0: 9}
                                v34 -= v34
                                _ = v74
                            } else {
                                _ = len(v4)
                            }
                        }
                        fmt.Println(v3, v34, v72)
                        delete(v4, "")
                    }
                    _ = fi73
                }
                _ = fi69
                _ = v72
            }
            v75 := []int{(v39 ^ 916), (v39 / ((-409) | 1))}
            _ = v75
        }
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
