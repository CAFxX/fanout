package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int32; F1 int64; F2 bool }
type GenS9 struct { F0 bool; F1 uint; F2 int16 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    v4 := []int{-9223372036854775808}
    var v5 float32 = ((1e10 / ((1e10 - float32((1.5 + (0.1 + (1.5 - 1.5))))) - 1.5)) - (((1e10 / 0.1) - 3.14159) / 1.0))
    v6 := GenS1{F0: 0}
    if true {
        fmt.Println(v5)
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            v4 = append(v4, (fi7 ^ (fi7 >> ((255) & 63))))
            v5 -= (float32(-2.25) * -2.25)
            _ = fi7
        }
        fmt.Println(v5)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v5 = float32((0.0 - ((0.1 * -2.25) / 1.0)))
    v8 := GenS1{F0: -120}
    v5 -= v5
    v5 -= float32((0.0 * 1.5))
    v5 -= 1.5
    if false {
        var v9 uint32 = 65535
        v9 -= 0
        var v10 float64 = float64(v5)
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 14; fi12++ {
                if true {
                    if (-1 == 2147483647) {
                        v5 = ((1e10 * (1e10 * (v5 * v5))) + (v5 / (v5 / (v5 / 1e10))))
                        v13 := GenS1{F0: 42}
                        _ = v13
                    } else {
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v5, fi12)
                            fmt.Println(fi12, fi11, v5)
                            fmt.Println(fi12)
                        default:
                            switch 4294967295 {
                            case 2147483647:
                                switch 0 {
                                case 42:
                                    fmt.Println(v5)
                                case 127:
                                    switch ("hello" + ("a" + ((("\n\t" + "a") + "foo bar") + (("foo bar" + "\n\t") + "hello")))) {
                                    case "foo bar":
                                        fmt.Println(fi12)
                                    default:
                                        var v14 int = fi11
                                        _ = v14
                                    }
                                case 32767:
                                    v15 := GenS1{F0: -128}
                                    _ = v15
                                default:
                                    v5 = v5
                                }
                            default:
                                if false {
                                    v2[0] = fi11
                                    v5 -= 1e10
                                    v5 = ((0.0 / v5) / float32(0.1))
                                } else {
                                    var v16 int32 = -1
                                    v17 := v2[2147483648]
                                    switch fi12 {
                                    case 42:
                                        if false {
                                            if true {
                                                delete(v2, 2147483647)
                                                _ = len(v3)
                                            } else {
                                                switch "" {
                                                case "\n\t":
                                                    if (false && true) {
                                                        v4 = append(v4, fi12)
                                                        v18 := GenS2{F0: "", F1: 2, F2: -49}
                                                        _ = v18
                                                    }
                                                case "a":
                                                    if true {
                                                        _ = len(v3)
                                                        v19 := GenS3{F0: 0, F1: 1}
                                                        v20 := GenS4{F0: 255}
                                                        var v21 uint16 = 32767
                                                        var v22 bool = false
                                                        _ = v19
                                                        _ = v20
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                case "hello":
                                                    var v23 uint = 65535
                                                    _ = v23
                                                default:
                                                    if ("" < "hello") {
                                                        switch 32768 {
                                                        case 128:
                                                            if (32767 != -32768) {
                                                                fmt.Println(fi12, v5, v16)
                                                            } else {
                                                                fmt.Println(fi12, v17)
                                                                switch (fi11 / ((fi11) | 1)) {
                                                                case 42:
                                                                    var v24 int64 = 9223372036854775807
                                                                    _ = v24
                                                                case 32768:
                                                                    v4 = append(v4, 128)
                                                                case -2147483648:
                                                                    var v25 uint32 = 1
                                                                    _ = v25
                                                                }
                                                            }
                                                        case 914:
                                                            v5 = v5
                                                        case 42:
                                                            v26 := GenS5{F0: 1, F1: 57, F2: -852}
                                                            var v27 int8 = 37
                                                            fmt.Println(v17, v17)
                                                            _ = v26
                                                            _ = v27
                                                        default:
                                                            v28 := GenS2{F0: (("foo bar" + "hello") + ""), F1: 337, F2: 86}
                                                            _ = v28
                                                        }
                                                        v29 := GenS2{F0: ((("a" + ("" + "\n\t")) + "a") + "\n\t"), F1: 18446744073709551615, F2: 39}
                                                        v17 = (fi11 ^ (v17 & fi11))
                                                        switch "hello" {
                                                        case "hello":
                                                            if len(v4) > 0 {
                                                                v4[0] = fi12
                                                            }
                                                        default:
                                                            v30 := GenS6{F0: 273}
                                                            _ = v30
                                                        }
                                                        _ = v29
                                                    }
                                                }
                                                fmt.Println(v17, fi11)
                                                var v31 uint32 = 2147483648
                                                _ = v31
                                            }
                                        }
                                    default:
                                        var v32 uint = 18446744073709551615
                                        _ = v32
                                    }
                                    if (195 > (v17 / ((4294967295) | 1))) {
                                        var v33 int = (fi12 / ((fi11) | 1))
                                        var v34 int8 = -1
                                        var v35 int64 = 1
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = v16
                                    _ = v17
                                }
                            }
                        }
                        v5 = (((0.0 * (3.14159 / v5)) + 1.5) - float32((1e10 / ((0.0 / 1.0) - (1.5 / 1.0)))))
                    }
                    var v36 int16 = 42
                    _ = v36
                } else {
                    var v37 int8 = -128
                    v37 -= (-(v37))
                    if true {
                        if true {
                            v3["foo bar"] = ("a" + "hello")
                        }
                        if false {
                            var v38 uint = 730
                            _ = v38
                        }
                        var v39 uint32 = 2
                        switch "" {
                        case "\n\t":
                            v37 = (v37 ^ v37)
                            v4 = append(v4, (fi12 << ((4294967295) & 63)))
                            v37 -= -1
                        case "foo bar":
                            v39 = (v39 + v39)
                        default:
                            var v40 uint64 = 977
                            _ = v40
                        }
                        if false {
                            v5 = v5
                            v41 := GenS1{F0: v37}
                            _ = v41
                        }
                        _ = v39
                    } else {
                        v5 -= ((v5 * ((0.0 / v5) - 3.14159)) - (v5 + v5))
                        switch "\n\t" {
                        case "":
                            v42 := GenS7{F0: 1, F1: -1}
                            _ = v42
                        case "hello":
                            var v43 uint8 = 0
                            _ = v43
                        default:
                            if (!false) {
                                v37 -= v37
                                var v44 int32 = 181
                                _ = v44
                            }
                        }
                        v37 = (v37 % (((v37 / (((v37 & (v37 + v37))) | 1))) | 1))
                    }
                    _ = v37
                }
                fmt.Println(fi12, fi12)
                _ = fi12
            }
            v45 := GenS3{F0: -128, F1: 4294967295}
            var v46 int32 = 1
            _ = fi11
            _ = v45
            _ = v46
        }
        v5 = 3.14159
        if false {
            v5 += v5
            fmt.Println(v5)
        } else {
            v4 = append(v4, 2147483647)
            var v47 int32 = -2147483648
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        for fi49 := 0; fi49 < 19; fi49++ {
            var v50 int64 = -613
            _ = fi49
            _ = v50
        }
        if (true && ("hello" != ("hello" + ((("foo bar" + "foo bar") + "\n\t") + "foo bar")))) {
            delete(v3, ("a" + "hello"))
            if (32767 > 65535) {
                v5 += float32((0.0 + 1.5))
                var v51 bool = (!(!false))
                _ = v51
            }
            if false {
                if (758 > 42) {
                    var v52 int = -1
                    switch "" {
                    case "foo bar":
                        for fi53 := 0; fi53 < 10; fi53++ {
                            fmt.Println(fi48, fi48, v52)
                            v54 := GenS8{F0: 42, F1: -1, F2: true}
                            switch "\n\t" {
                            case "":
                                if (((3.14159 * 1.5) / 3.14159) > 1.5) {
                                    var v55 uint32 = 336
                                    _ = v55
                                }
                            default:
                                var v56 float64 = 3.14159
                                _ = v56
                            }
                            v52 = 42
                            fmt.Println(v52, fi53)
                            _ = fi53
                            _ = v54
                        }
                    case "hello":
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    if false {
                                        v5 = v5
                                    }
                                } else {
                                    v57 := GenS9{F0: (1e10 > (3.14159 + -2.25)), F1: 0, F2: 1}
                                    _ = v57
                                }
                            case "":
                                v1[-540] = (!false)
                            case "hello":
                                v52 = (v52 | fi48)
                            default:
                                var v58 int = 0
                                _ = v58
                            }
                            v52 += 0
                        } else {
                            v4 = append(v4, -1)
                        }
                    case "\n\t":
                        v5 -= v5
                    }
                    _ = v52
                } else {
                    if (1 != 42) {
                        v59 := GenS10{F0: 32767, F1: 0}
                        v5 = (v5 * (1e10 - 0.0))
                        _ = v59
                    }
                    _ = len(v4)
                    v60 := GenS6{F0: 1}
                    _ = v60
                }
                if (-1 > -1) {
                    delete(v1, fi48)
                }
                if (18446744073709551615 >= 255) {
                    v61 := GenS11{F0: -7, F1: fi48}
                    _ = v61
                } else {
                    var v62 uint16 = 2
                    fmt.Println(fi48)
                    v5 = v5
                    _ = v62
                }
            } else {
                v63 := GenS10{F0: 42, F1: -123}
                for fi64 := 0; fi64 < 19; fi64++ {
                    _ = len(v1)
                    _ = fi64
                }
                _ = v63
            }
        } else {
            if (792 > -9223372036854775808) {
                for fi65 := 0; fi65 < 16; fi65++ {
                    v5 = v5
                    var v66 uint16 = 1
                    _ = fi65
                    _ = v66
                }
                if len(v4) > 0 {
                    v4[0] = (fi48 * (fi48 & fi48))
                }
                if ((-407 == 1) || true) {
                    if false {
                        var v67 int32 = 1
                        v68 := GenS6{F0: 0}
                        _ = v67
                        _ = v68
                    }
                } else {
                    for fi69 := 0; fi69 < 18; fi69++ {
                        if (255 != 4294967295) {
                            v5 -= (float32(1.5) * (-2.25 / 1.0))
                        } else {
                            fmt.Println(fi48)
                            v5 = v5
                            var v70 uint8 = 127
                            _ = v70
                        }
                        v71 := GenS12{F0: 127, F1: 65535}
                        _ = fi69
                        _ = v71
                    }
                    switch "\n\t" {
                    case "foo bar":
                        switch -1 {
                        case 127:
                            var v72 uint = 2
                            _ = v72
                        case 9223372036854775807:
                            v73 := GenS7{F0: 65535, F1: 1}
                            _ = v73
                        case -9223372036854775808:
                            var v74 int16 = 1000
                            _ = v74
                        default:
                            var v75 uint32 = 65535
                            _ = v75
                        }
                    case "":
                        v5 += float32((((1e10 - (-2.25 + 1e10)) / ((((0.0 * 1.5) - 1e10) + -2.25) + (3.14159 - -2.25))) - float64(1.5)))
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v2[0] = -972
                }
            } else {
                v5 = v5
                fmt.Println(v5, v5)
            }
        }
        var v76 string = "\n\t"
        _ = fi48
        _ = v76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
