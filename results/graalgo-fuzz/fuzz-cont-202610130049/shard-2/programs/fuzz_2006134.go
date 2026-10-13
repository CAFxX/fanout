package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 string; F1 int32 }

func main() {
    var v1 string = ((((("a" + (("hello" + "hello") + "")) + ("\n\t" + "foo bar")) + ((("foo bar" + ("foo bar" + "hello")) + "hello") + "foo bar")) + "") + "\n\t")
    var v2 uint16 = 2
    var v3 float64 = (((1e10 + float64(3.14159)) / 0.1) / 3.14159)
    v4 := make(map[int]uint32)
    v5 := []string{v1, v1, v1, "a"}
    var v6 int8 = -69
    if true {
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(v2)
            for fi8 := 0; fi8 < 1; fi8++ {
                if false {
                    var v9 uint32 = 2147483648
                    _ = v9
                } else {
                    v10 := GenS1{F0: fi7}
                    switch v1 {
                    case "hello":
                        fmt.Println(v3, v3)
                    case "":
                        fmt.Println(v1, v2, fi7)
                    case "\n\t":
                        var v11 uint64 = 142
                        _ = v11
                    }
                    var v12 bool = ((!false) || (((((1.5 + (1.5 + 3.14159)) * 0.0) / 1e10) / 1e10) >= 3.14159))
                    _ = v10
                    _ = v12
                }
                v13 := GenS1{F0: 0}
                if true {
                    if true {
                        if (uint8(v2) != 130) {
                            switch (fi8 >> ((4294967295) & 63)) {
                            case 32768:
                                if ((222 > 533) && (255 <= v2)) {
                                    v3 += -2.25
                                    v14 := GenS1{F0: 0}
                                    var v15 int = (fi7 << ((1) & 63))
                                    switch 42 {
                                    case -33:
                                        var v16 int8 = v6
                                        _ = v16
                                    case 1:
                                        if true {
                                            v6 -= 100
                                            var v17 string = ("foo bar" + v1)
                                            _ = v17
                                        } else {
                                            var v18 int64 = 9223372036854775807
                                            var v19 string = v1
                                            _ = v18
                                            _ = v19
                                        }
                                    default:
                                        if len(v5) > 0 {
                                            v5[0] = ((v1 + v1) + "a")
                                        }
                                        if true {
                                            var v20 string = v1
                                            var v21 uint8 = 255
                                            var v22 int = fi8
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    _ = v14
                                    _ = v15
                                }
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                var v23 uint = 4294967295
                                _ = v23
                            default:
                                v1 = ("" + "")
                            }
                            switch fi7 {
                            case -372:
                                var v24 uint = 255
                                _ = v24
                            case -1:
                                switch (fi7 % ((860) | 1)) {
                                case 32768:
                                    switch (fi8 ^ (fi8 + (fi8 - (-(fi8))))) {
                                    case 129:
                                        v25 := v4[671]
                                        _ = v25
                                    case 0:
                                        var v26 int64 = 0
                                        _ = v26
                                    case 32768:
                                        v27 := GenS1{F0: fi7}
                                        _ = v27
                                    default:
                                        v28 := GenS1{F0: fi8}
                                        _ = v28
                                    }
                                    var v29 string = "hello"
                                    fmt.Println(fi8, v3)
                                    v6 = (^(v6))
                                    switch fi7 {
                                    case 372:
                                        v30 := GenS2{F0: -559}
                                        _ = v30
                                    default:
                                        switch v29 {
                                        case "":
                                            if (false || (42 != 255)) {
                                                var v31 int32 = 1
                                                v6 -= (v6 << ((255) & 63))
                                                if (42 != -184) {
                                                    var v32 int32 = (v31 % (((v31 - (v31 & v31))) | 1))
                                                    var v33 bool = true
                                                    v31 -= int32(v6)
                                                    fmt.Println(v29, v3)
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v31
                                            } else {
                                                v34 := GenS2{F0: 1}
                                                _ = v34
                                            }
                                        default:
                                            if (3.14159 != 1.5) {
                                                _ = len(v5)
                                                var v35 int32 = -2147483648
                                                switch (fi8 % ((127) | 1)) {
                                                case -508:
                                                    v5 = append(v5, "a")
                                                case 32767:
                                                    var v36 float64 = v3
                                                    _ = v36
                                                case 0:
                                                    var v37 float64 = (v3 + ((v3 - 1e10) - v3))
                                                    _ = v37
                                                }
                                                v38 := GenS3{F0: 1}
                                                _ = v35
                                                _ = v38
                                            }
                                        }
                                    }
                                    var v39 uint8 = 254
                                    _ = v29
                                    _ = v39
                                case 0:
                                    var v40 int = fi8
                                    _ = v40
                                default:
                                    var v41 uint16 = v2
                                    var v42 float32 = 0.1
                                    if (!(!((v41 << ((1) & 63)) >= v41))) {
                                        v6 = (^(v6))
                                        _ = len(v5)
                                        v42 = ((1e10 + v42) - (0.1 - v42))
                                    }
                                    _ = v41
                                    _ = v42
                                }
                            }
                        }
                        if (0 <= -1) {
                            var v43 float64 = (1e10 - v3)
                            _ = v43
                        } else {
                            var v44 int8 = v6
                            fmt.Println(fi7, v1)
                            _ = v44
                        }
                        var v45 uint32 = uint32(v2)
                        _ = v45
                    }
                    v2 -= v2
                    v6 -= -128
                }
                _ = fi8
                _ = v13
            }
            v46 := GenS2{F0: -9223372036854775808}
            _ = fi7
            _ = v46
        }
        var v47 int32 = -1
        v1 = "\n\t"
        _ = v47
    } else {
        for fi48 := 0; fi48 < 15; fi48++ {
            v1 = "a"
            _ = fi48
        }
    }
    _ = len(v5)
    for fi49 := 0; fi49 < 1; fi49++ {
        v6 -= v6
        if false {
            fmt.Println(v1, v2)
            delete(v4, (fi49 / (((fi49 & fi49)) | 1)))
            switch fi49 {
            case 32768:
                v50 := GenS2{F0: -9223372036854775808}
                _ = v50
            }
        } else {
            v51 := GenS4{F0: 767, F1: (fi49 / (((fi49 * (fi49 / ((-481) | 1)))) | 1))}
            var v52 int16 = 32767
            var v53 string = v1
            _ = v51
            _ = v52
            _ = v53
        }
        _ = fi49
    }
    v4[242] = 2147483648
    v6 += (v6 * 127)
    v2 = v2
    v54 := GenS5{F0: 831, F1: 163}
    if (255 == uint(v2)) {
        v5 = append(v5, v1)
    }
    v55 := GenS6{F0: (v1 + v1), F1: 2147483647}
    var v56 uint = 255
    var v57 int16 = -1
    if true {
        var v58 uint64 = 4294967295
        v56 += v56
        _ = v58
    }
    var v59 bool = (-677 <= 91)
    var v60 int32 = 2147483647
    for fi61 := 0; fi61 < 13; fi61++ {
        var v62 int8 = (-(v6))
        if v59 {
            var v63 int32 = (v60 / (((v60 + (^(v60)))) | 1))
            v64 := GenS3{F0: int16(v62)}
            _ = v63
            _ = v64
        }
        _ = fi61
        _ = v62
    }
    var v65 uint = v56
    if len(v5) > 0 {
        v5[0] = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v59
    _ = v6
    _ = v60
    _ = v65
}
