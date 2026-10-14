package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int16; F2 int8 }
type GenS3 struct { F0 int16; F1 int64; F2 int32 }
type GenS4 struct { F0 int8; F1 int64; F2 int64 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 uint8; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1e10
    var v3 int32 = 42
    if (true && (-2147483648 <= 636)) {
        v4 := v1[2147483648]
        if (true && false) {
            if (!v4) {
                if v4 {
                    var v5 int64 = 175
                    _ = v5
                }
            }
            var v6 string = "\n\t"
            switch v6 {
            case "a":
                var v7 int = 4294967295
                _ = v7
            }
            _ = v6
        } else {
            var v8 int32 = (v3 % (((v3 & v3)) | 1))
            _ = v8
        }
        _ = v4
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint16 = 32767
            fmt.Println(v3, v3)
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 765, F1: 0, F2: 2}
        _ = v11
    }
    _ = len(v1)
    delete(v1, 277)
    if false {
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v2)
        case "":
            if (255 == 1) {
                fmt.Println(v3, v2)
            } else {
                delete(v1, 0)
                for fi12 := 0; fi12 < 13; fi12++ {
                    switch "hello" {
                    case "":
                        delete(v1, 65535)
                    }
                    v2 -= v2
                    v13 := GenS2{F0: 346, F1: -443, F2: -77}
                    _ = fi12
                    _ = v13
                }
                var v14 uint = 4294967295
                _ = v14
            }
        default:
            v2 += 0.1
        }
    }
    if true {
        var v15 float32 = 3.14159
        _ = v15
    } else {
        fmt.Println(v3)
    }
    v16 := []int{}
    for fi17 := 0; fi17 < 3; fi17++ {
        if len(v16) > 0 {
            v16[0] = (fi17 >> ((18446744073709551615) & 63))
        }
        _ = fi17
    }
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint64 = 4294967295
        switch (fi19 & int(v20)) {
        case 646:
            var v21 float64 = 0.0
            _ = v21
        case 0:
            if false {
                v22 := GenS1{F0: (-(v3)), F1: (fi19 | (fi19 * 4294967295)), F2: 32767}
                var v23 int8 = 42
                _ = v22
                _ = v23
            } else {
                for fi24 := 0; fi24 < 5; fi24++ {
                    if true {
                        if true {
                            v3 = (-(v3))
                            v20 += v20
                            switch "" {
                            case "foo bar":
                                v25 := GenS3{F0: 32767, F1: 42, F2: v3}
                                _ = v25
                            case "\n\t":
                                if (!(1 >= -128)) {
                                    var v26 float32 = 0.1
                                    v2 += v2
                                    fmt.Println(v20, fi19, fi24)
                                    _ = v26
                                }
                            case "a":
                                var v27 int8 = 0
                                _ = v27
                            }
                        } else {
                            v20 += v20
                        }
                        var v28 uint8 = (v18 * v18)
                        var v29 string = "\n\t"
                        var v30 int32 = -2147483648
                        fmt.Println(v2)
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                    _ = len(v16)
                    _ = fi24
                }
                v2 = (v2 + v2)
                switch "foo bar" {
                case "foo bar":
                    for fi31 := 0; fi31 < 2; fi31++ {
                        var v32 bool = true
                        if v32 {
                            v20 = (v20 / (((^(v20))) | 1))
                        } else {
                            if false {
                                v33 := GenS4{F0: 12, F1: 1, F2: 1}
                                if true {
                                    var v34 float32 = (float32((float64(float32(v2)) / -2.25)) - 3.14159)
                                    switch fi31 {
                                    case -454:
                                        if (true && true) {
                                            v34 += 1.5
                                            var v35 uint8 = uint8(v20)
                                            fmt.Println(v18)
                                            var v36 int32 = v3
                                            _ = v35
                                            _ = v36
                                        } else {
                                            var v37 int32 = v3
                                            v38 := GenS1{F0: v3, F1: fi19, F2: 65535}
                                            v3 = (v37 << ((53) & 63))
                                            _ = v37
                                            _ = v38
                                        }
                                    case -9223372036854775808:
                                        var v39 bool = ((v18 >> ((712) & 63)) >= (^(v18)))
                                        _ = v39
                                    case 0:
                                        v18 = v18
                                    }
                                    _ = v34
                                } else {
                                    var v40 uint8 = v18
                                    if ((^(fi31)) <= (fi19 % ((fi19) | 1))) {
                                        if v32 {
                                            v40 = (v40 + v40)
                                            var v41 int16 = 194
                                            _ = v41
                                        }
                                    }
                                    _ = v40
                                }
                                v42 := GenS5{F0: (v32 || ((42 < 65535) || v32)), F1: uint32(v18)}
                                _ = v33
                                _ = v42
                            } else {
                                if true {
                                    var v43 uint = 18446744073709551615
                                    if (!((753 < uint16(v43)) && v32)) {
                                        var v44 float64 = v2
                                        _ = v44
                                    }
                                    _ = len(v16)
                                    _ = v43
                                }
                            }
                            var v45 uint8 = 255
                            var v46 int8 = -35
                            _ = v45
                            _ = v46
                        }
                        _ = fi31
                        _ = v32
                    }
                default:
                    v47 := GenS1{F0: v3, F1: fi19, F2: 206}
                    _ = v47
                }
            }
        }
        _ = fi19
        _ = v20
    }
    for fi48 := 0; fi48 < 6; fi48++ {
        fmt.Println(v3, v18)
        fmt.Println(fi48)
        if (-2.25 != 1.5) {
            if true {
                if true {
                    v49 := GenS2{F0: 2, F1: 1000, F2: -1}
                    v18 -= 0
                    for fi50 := 0; fi50 < 14; fi50++ {
                        v51 := GenS6{F0: -729, F1: -1}
                        v52 := GenS7{F0: false}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                }
                if true {
                    switch fi48 {
                    case 42:
                        if false {
                            v53 := GenS6{F0: (fi48 | (fi48 - fi48)), F1: -465}
                            if (true && (!true)) {
                                v18 -= (v18 % ((v18) | 1))
                            } else {
                                for fi54 := 0; fi54 < 16; fi54++ {
                                    v2 = v2
                                    switch (((("a" + "\n\t") + ("\n\t" + "a")) + "\n\t") + "a") {
                                    case "a":
                                        if false {
                                            if (!false) {
                                                var v55 int = fi54
                                                fmt.Println(v2, v3, v18)
                                                _ = len(v16)
                                                if true {
                                                    if (((((float64(1e10) / (1.5 / (3.14159 / v2))) - -2.25) * v2) - v2) != ((v2 + 1.5) - v2)) {
                                                        v56 := GenS1{F0: (v3 ^ v3), F1: -9223372036854775808, F2: 255}
                                                        _ = v56
                                                    } else {
                                                        v57 := GenS8{F0: 4294967295, F1: 1}
                                                        _ = v57
                                                    }
                                                } else {
                                                    v18 += (v18 << ((18446744073709551615) & 63))
                                                }
                                                var v58 bool = true
                                                v58 = v58
                                                if len(v16) >= 2 {
                                                    v16 = v16[1:len(v16)]
                                                }
                                                var v59 int64 = 860
                                                _ = v55
                                                _ = v58
                                                _ = v59
                                            }
                                        } else {
                                            switch 864 {
                                            case 508:
                                                v60 := GenS9{F0: v18, F1: "foo bar"}
                                                _ = v60
                                            }
                                        }
                                    }
                                    _ = fi54
                                }
                                delete(v1, 42)
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    var v62 float32 = 0.1
                                    v63 := GenS9{F0: 239, F1: "foo bar"}
                                    _ = fi61
                                    _ = v62
                                    _ = v63
                                }
                            }
                            v64 := GenS6{F0: 32767, F1: 0}
                            _ = v53
                            _ = v64
                        }
                    case 306:
                        var v65 uint32 = 4294967295
                        _ = v65
                    case 32768:
                        switch (fi48 / ((0) | 1)) {
                        case -898:
                            for fi66 := 0; fi66 < 13; fi66++ {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    var v67 uint16 = 0
                                    _ = v67
                                case "hello":
                                    v68 := v1[42]
                                    _ = v68
                                default:
                                    var v69 int8 = 100
                                    _ = v69
                                }
                                var v70 string = "foo bar"
                                _ = len(v16)
                                var v71 int8 = 4
                                _ = fi66
                                _ = v70
                                _ = v71
                            }
                        case -1:
                            fmt.Println(fi48, v18, v3)
                        case 1:
                            var v72 int8 = 100
                            _ = v72
                        }
                    }
                    for fi73 := 0; fi73 < 10; fi73++ {
                        var v74 float32 = -2.25
                        switch fi48 {
                        case -2147483648:
                            switch (fi73 + fi48) {
                            case 2147483647:
                                v75 := GenS7{F0: true}
                                _ = v75
                            case 128:
                                v74 = v74
                            case 364:
                                if (4294967295 <= (fi73 % ((fi48) | 1))) {
                                    v3 = (-(v3))
                                    v74 = (float32(v2) + v74)
                                    var v76 float64 = float64(v74)
                                    _ = v76
                                }
                                _ = len(v16)
                                v77 := GenS9{F0: (v18 & 255), F1: "\n\t"}
                                v18 = 255
                                _ = v77
                            default:
                                v74 += v74
                            }
                        case 80:
                            var v78 int32 = (v3 / (((v3 << ((460) & 63))) | 1))
                            _ = v78
                        }
                        _ = fi73
                        _ = v74
                    }
                    var v79 int16 = 1000
                    _ = v79
                }
            }
        }
        _ = fi48
    }
    v3 = 0
    for fi80 := 0; fi80 < 4; fi80++ {
        for fi81 := 0; fi81 < 8; fi81++ {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = fi81
        }
        var v82 uint16 = uint16(v18)
        _ = fi80
        _ = v82
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v83 float64 = (v2 / 1e10)
    for fi84 := 0; fi84 < 2; fi84++ {
        v3 = (v3 - (^(v3)))
        fmt.Println(v83, v18, v3)
        _ = fi84
    }
    v3 += (v3 << ((4294967295) & 63))
    fmt.Println(v83)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v83
}
