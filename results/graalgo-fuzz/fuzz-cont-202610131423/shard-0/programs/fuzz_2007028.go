package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int32; F2 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint }

func main() {
    v1 := []uint32{358, 4294967295, 42}
    v2 := []uint32{738, 2, 42}
    v3 := make(map[string]uint32)
    v4 := []uint32{42}
    var v5 uint8 = 0
    var v6 float32 = (3.14159 + 0.0)
    if true {
        v7 := GenS1{F0: 32767}
        fmt.Println(v6, v6)
        if (18446744073709551615 < 155) {
            var v8 uint16 = 855
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 int32 = 42
                if false {
                    if ((!((true && (!false)) || true)) && true) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi11 := 0; fi11 < 10; fi11++ {
                            fmt.Println(v6, v5)
                            _ = fi11
                        }
                        for fi12 := 0; fi12 < 8; fi12++ {
                            var v13 int64 = 9223372036854775807
                            var v14 float32 = v6
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                    }
                    if (((-1 < v10) || (!false)) && (!(!((0.0 + (3.14159 - 3.14159)) >= 3.14159)))) {
                        var v15 uint = 18446744073709551615
                        _ = v15
                    }
                    if (!false) {
                        v10 = 42
                        if (false || true) {
                            delete(v3, ((("foo bar" + "foo bar") + ("foo bar" + "foo bar")) + "a"))
                        }
                        if false {
                            v8 = v8
                        } else {
                            var v16 float32 = v6
                            var v17 uint = 2
                            _ = v16
                            _ = v17
                        }
                    } else {
                        var v18 int = fi9
                        _ = v18
                    }
                }
                _ = fi9
                _ = v10
            }
            v5 += v5
            _ = v8
        } else {
            v6 = float32(0.0)
            v5 = v5
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 uint = 961
                v20 = (v20 - v20)
                _ = fi19
                _ = v20
            }
        }
        _ = v7
    } else {
        v5 -= v5
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS1{F0: -9223372036854775808}
            var v23 uint16 = 255
            delete(v3, "hello")
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 float32 = (((0.1 * (1.5 / v6)) / v6) * v6)
        _ = v24
    }
    var v25 int32 = 261
    var v26 string = "foo bar"
    if (float64(v6) >= (-2.25 / 1e10)) {
        var v27 int64 = 9223372036854775807
        v1 = append(v1, 4294967295)
        _ = v27
    }
    v26 = v26
    v26 = v26
    fmt.Println(v5)
    for fi28 := 0; fi28 < 13; fi28++ {
        fmt.Println(v26)
        if (42 <= 1) {
            switch fi28 {
            case 128:
                for fi29 := 0; fi29 < 7; fi29++ {
                    _ = len(v3)
                    v30 := GenS2{F0: (!true), F1: v25, F2: 962}
                    _ = fi29
                    _ = v30
                }
            }
            v25 = v25
            var v31 int = (^(fi28))
            _ = v31
        } else {
            switch v26 {
            case "":
                v32 := v3[""]
                _ = v32
            default:
                v26 = (v26 + "\n\t")
            }
            switch fi28 {
            case 578:
                v5 = (^(v5))
            case 2147483647:
                v25 += (-(v25))
            default:
                v5 += v5
            }
            if false {
                if true {
                    fmt.Println(v6, v5, v5)
                    var v33 float64 = 1e10
                    _ = v33
                }
                if false {
                    v6 = 0.0
                } else {
                    var v34 float64 = 0.1
                    if false {
                        var v35 string = v26
                        switch ("a" + "foo bar") {
                        case "hello":
                            v34 = (v34 + (v34 * v34))
                        case "\n\t":
                            for fi36 := 0; fi36 < 1; fi36++ {
                                v5 -= v5
                                v26 = ((v35 + (("" + ("hello" + ((("\n\t" + "a") + v35) + v35))) + v26)) + ((v26 + v35) + v26))
                                v35 = ("hello" + "\n\t")
                                _ = fi36
                            }
                        case "a":
                            if true {
                                var v37 uint = 4294967295
                                if true {
                                    for fi38 := 0; fi38 < 4; fi38++ {
                                        v35 = (v26 + (v35 + "a"))
                                        fmt.Println(v37, v37)
                                        var v39 bool = false
                                        _ = fi38
                                        _ = v39
                                    }
                                    switch fi28 {
                                    case -540:
                                        v3[v35] = 2147483648
                                    }
                                    switch v26 {
                                    case "\n\t":
                                        for fi40 := 0; fi40 < 11; fi40++ {
                                            v26 = v35
                                            switch (v35 + v26) {
                                            case "\n\t":
                                                fmt.Println(v34, v34)
                                            }
                                            fmt.Println(v35, v37, v6)
                                            _ = fi40
                                        }
                                    case "hello":
                                        var v41 uint64 = 784
                                        _ = v41
                                    case "":
                                        var v42 float64 = (float64((-2.25 / 0.1)) * ((-2.25 / ((3.14159 * (((v34 - -2.25) / float64(3.14159)) + v34)) + (float64(v6) * v34))) + v34))
                                        _ = v42
                                    default:
                                        if false {
                                            v5 -= 127
                                            var v43 int32 = v25
                                            _ = v43
                                        }
                                    }
                                }
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    v45 := GenS2{F0: (!true), F1: v25, F2: 2147483647}
                                    _ = fi44
                                    _ = v45
                                }
                                _ = v37
                            }
                        default:
                            v46 := GenS2{F0: (1 < 1), F1: 794, F2: 2}
                            _ = v46
                        }
                        _ = v35
                    } else {
                        v5 = v5
                    }
                    v26 = (v26 + "hello")
                    _ = v34
                }
            } else {
                fmt.Println(v6, v5, v25)
                var v47 int32 = 42
                v6 -= (v6 * 1e10)
                _ = v47
            }
        }
        switch v26 {
        case "\n\t":
            for fi48 := 0; fi48 < 2; fi48++ {
                switch (fi48 / ((4294967295) | 1)) {
                case 2147483647:
                    v49 := GenS3{F0: "\n\t", F1: false, F2: -1}
                    _ = v49
                case 42:
                    v50 := GenS3{F0: (v26 + (v26 + (v26 + "a"))), F1: true, F2: 32767}
                    _ = v50
                case 1:
                    var v51 float32 = float32(float64(v6))
                    _ = v51
                default:
                    if false {
                        switch "hello" {
                        case "":
                            fmt.Println(v26, fi28, v5)
                        case "foo bar":
                            fmt.Println(v5)
                        case "hello":
                            v5 += 156
                        }
                        if false {
                            v52 := GenS4{F0: (((-2.25 * (1e10 + v6)) + (((v6 - v6) * v6) * v6)) > v6)}
                            v53 := GenS5{F0: fi28}
                            if true {
                                fmt.Println(v25, v25, v6)
                            } else {
                                v25 += v25
                            }
                            v5 -= 1
                            v54 := v3[v26]
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    } else {
                        var v55 int8 = 42
                        _ = v55
                    }
                }
                _ = fi48
            }
        }
        _ = fi28
    }
    var v56 int32 = v25
    if true {
        for fi57 := 0; fi57 < 3; fi57++ {
            var v58 uint32 = 896
            if (true && false) {
                v58 = (v58 | v58)
            }
            _ = fi57
            _ = v58
        }
        fmt.Println(v6, v26, v6)
        _ = len(v3)
    }
    if true {
        v59 := GenS5{F0: -2147483648}
        _ = v59
    }
    v5 -= 7
    _ = len(v3)
    var v60 uint64 = 0
    if (!(-75 > 1)) {
        if (-203 < 9223372036854775807) {
            switch v26 {
            case "\n\t":
                for fi61 := 0; fi61 < 10; fi61++ {
                    var v62 uint = 65535
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v60 -= (v60 + (v60 << (((^(v62))) & 63)))
                    _ = fi61
                    _ = v62
                }
            case "":
                fmt.Println(v25, v6, v26)
            default:
                v63 := GenS3{F0: "", F1: true, F2: -344}
                _ = v63
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        } else {
            v5 = 147
            var v64 int16 = 42
            _ = v64
        }
    }
    var v65 uint16 = 255
    for fi66 := 0; fi66 < 17; fi66++ {
        v67 := GenS6{F0: v65, F1: 2, F2: 1}
        var v68 uint16 = 65535
        _ = fi66
        _ = v67
        _ = v68
    }
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v60
    _ = v65
}
