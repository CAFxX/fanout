package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int16, p1 uint32, p2 uint32) uint8 {
    var v5 string = "foo bar"
    v6 := []int{}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v5 = v5
            v5 = v5
            v5 = v5
            _ = fi8
        }
        var v9 uint = 615
        switch fi7 {
        case 32768:
            if false {
                if len(v6) > 0 {
                    v6[0] = fi7
                }
                var v10 int32 = 0
                _ = v10
            } else {
                for fi11 := 0; fi11 < 18; fi11++ {
                    var v12 int16 = -32768
                    _ = fi11
                    _ = v12
                }
                v13 := GenS1{F0: (!(-2.25 >= (1e10 - 1e10))), F1: int16(fi7), F2: 2}
                switch (fi7 << (((v9 ^ v9)) & 63)) {
                case 127:
                    fmt.Println(p2, fi7)
                case 31:
                    p0 -= (p0 + (p0 * 707))
                }
                _ = v13
            }
        case 781:
            switch fi7 {
            case -989:
                var v14 string = v5
                _ = v14
            case -514:
                if (!(!false)) {
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 float32 = (1.5 * (0.1 - 0.1))
                        _ = fi15
                        _ = v16
                    }
                    switch "\n\t" {
                    case "hello":
                        v17 := GenS1{F0: (false && (-2.25 >= 0.0)), F1: 32767, F2: 2147483648}
                        _ = v17
                    case "":
                        var v18 int = fi7
                        _ = v18
                    case "\n\t":
                        v19 := GenS2{F0: (fi7 << ((v9) & 63))}
                        _ = v19
                    }
                }
            }
        default:
            var v20 uint16 = 2
            _ = v20
        }
        _ = fi7
        _ = v9
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v6 = append(v6, fi21)
        v5 = "\n\t"
        _ = fi21
    }
    v6 = append(v6, -652)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 135
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    var v3 uint64 = 42
    var v4 uint64 = (v3 ^ v3)
    fmt.Println(v2, v3)
    v22 := []int{}
    for fi23 := 0; fi23 < 8; fi23++ {
        v4 += v3
        v24 := GenS1{F0: (!true), F1: (v2 | 70), F2: 2}
        for fi25 := 0; fi25 < 6; fi25++ {
            v26 := GenS3{F0: -1, F1: 339}
            var v27 int8 = 127
            switch (^(fi23)) {
            case -1:
                v28 := GenS3{F0: (fi25 / ((128) | 1)), F1: 758}
                _ = v28
            case -2147483648:
                if ((true || (!(4294967295 == 2))) || (!false)) {
                    v3 -= (v4 | v3)
                    fmt.Println(v2)
                } else {
                    fmt.Println(fi25, v3)
                    var v29 int64 = -765
                    v27 -= (v27 ^ -1)
                    var v30 int8 = (-(v27))
                    switch ("a" + "\n\t") {
                    case "":
                        switch "a" {
                        case "hello":
                            var v31 uint8 = 1
                            _ = v31
                        case "a":
                            switch "hello" {
                            case "":
                                if false {
                                    switch -1 {
                                    case 0:
                                        fmt.Println(v2, v30)
                                    case 42:
                                        var v32 uint16 = 946
                                        _ = v32
                                    case 128:
                                        v33 := GenS4{F0: v3, F1: v3}
                                        _ = v33
                                    default:
                                        v30 = (v27 << ((345) & 63))
                                    }
                                }
                            case "hello":
                                if true {
                                    v34 := GenS4{F0: v3, F1: v4}
                                    v35 := GenS3{F0: (^(fi25)), F1: 217}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v36 := GenS1{F0: (true && (42 > 127)), F1: v2, F2: 755}
                                    v37 := GenS2{F0: 2147483647}
                                    _ = v36
                                    _ = v37
                                }
                            default:
                                fmt.Println(fi23)
                            }
                        case "":
                            if ((true && (!(-802 >= 0))) && (true || true)) {
                                v4 = (v4 % ((v3) | 1))
                            }
                        default:
                            var v38 string = ("" + ((("\n\t" + "foo bar") + "a") + "hello"))
                            _ = v38
                        }
                    default:
                        v27 = v30
                    }
                    _ = v29
                    _ = v30
                }
            case 0:
                v2 -= (v2 * (-(v2)))
            default:
                var v39 int8 = v27
                _ = v39
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    if true {
        var v40 int32 = 772
        _ = v40
    } else {
        var v41 float64 = ((0.1 + (float64(0.1) - -2.25)) + -2.25)
        _ = len(v22)
        _ = v41
    }
    switch 1 {
    case 540:
        v42 := v1[-1]
        _ = v42
    case 4294967295:
        v3 = (v4 / (((^(v4))) | 1))
    default:
        v3 += (v4 * (v4 & (v4 * 0)))
    }
    if (v2 > (v2 & (^(v2)))) {
        for fi43 := 0; fi43 < 13; fi43++ {
            fmt.Println(v2, v4)
            v44 := GenS1{F0: ((true || true) && (567 < 325)), F1: (v2 | v2), F2: 2147483648}
            v4 = 0
            _ = fi43
            _ = v44
        }
    } else {
        for fi45 := 0; fi45 < 8; fi45++ {
            for fi46 := 0; fi46 < 8; fi46++ {
                switch (fi46 - fi46) {
                case 42:
                    v47 := GenS2{F0: fi45}
                    v4 = 1
                    switch (fi45 / ((2147483647) | 1)) {
                    case -2147483648:
                        var v48 float32 = (1e10 / -2.25)
                        _ = v48
                    case -35:
                        if false {
                            if true {
                                switch ("hello" + "hello") {
                                case "a":
                                    var v49 int32 = 2147483647
                                    _ = v49
                                case "\n\t":
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                case "hello":
                                    var v50 bool = (0 != 42)
                                    _ = v50
                                }
                            }
                        }
                    case 1:
                        var v51 int64 = 9223372036854775807
                        _ = v51
                    default:
                        switch (("a" + "foo bar") + (("hello" + (("foo bar" + (("foo bar" + ("foo bar" + "a")) + (("" + "foo bar") + "hello"))) + "a")) + "")) {
                        case "\n\t":
                            v4 = (v3 / (((v3 & (v4 % ((871) | 1)))) | 1))
                        case "a":
                            var v52 uint = 491
                            _ = v52
                        }
                    }
                    switch ((("\n\t" + "hello") + "hello") + ("" + (((("foo bar" + "a") + "hello") + "") + ("\n\t" + ("" + "a"))))) {
                    case "":
                        var v53 uint64 = 4294967295
                        _ = v53
                    case "foo bar":
                        var v54 uint32 = 0
                        _ = v54
                    case "a":
                        if (590 != 640) {
                            if ((((("" + ("" + (("a" + "\n\t") + "foo bar"))) + "") + ("hello" + ("\n\t" + (("hello" + ("" + "\n\t")) + ("" + ""))))) + "") >= ("" + "hello")) {
                                var v55 int64 = 1
                                v56 := GenS3{F0: 32768, F1: 42}
                                _ = v55
                                _ = v56
                            } else {
                                switch ((("foo bar" + "foo bar") + "\n\t") + "hello") {
                                case "hello":
                                    fmt.Println(v2, v2)
                                    v2 += 1
                                case "a":
                                    var v57 float32 = (0.1 - 0.0)
                                    _ = v57
                                case "\n\t":
                                    var v58 uint32 = 360
                                    _ = v58
                                default:
                                    var v59 float64 = 0.0
                                    _ = v59
                                }
                            }
                        } else {
                            var v60 uint32 = 42
                            _ = v60
                        }
                    default:
                        switch "hello" {
                        case "\n\t":
                            var v61 uint8 = 2
                            _ = v61
                        case "foo bar":
                            switch ((("" + "") + ("hello" + "foo bar")) + "foo bar") {
                            case "hello":
                                fmt.Println(fi45, v3)
                            case "a":
                                var v62 float64 = float64(((float32((1.5 / (0.1 * 1e10))) - float32(0.1)) - float32(1e10)))
                                _ = v62
                            case "\n\t":
                                if (351 < 306) {
                                    var v63 uint64 = (v4 - v4)
                                    _ = v63
                                } else {
                                    v3 -= (v4 / (((v3 - v4)) | 1))
                                    fmt.Println(fi46, v3)
                                }
                            }
                        }
                    }
                    _ = v47
                default:
                    fmt.Println(fi46)
                }
                v64 := GenS5{F0: v2}
                _ = fi46
                _ = v64
            }
            _ = fi45
        }
        var v65 float64 = (0.0 * 0.1)
        _ = v65
    }
    for fi66 := 0; fi66 < 2; fi66++ {
        if (!(!(100 <= 21))) {
            if (0 != 1) {
                for fi67 := 0; fi67 < 18; fi67++ {
                    fmt.Println(v3)
                    v68 := v1[fi67]
                    v2 = 1
                    _ = fi67
                    _ = v68
                }
            } else {
                fmt.Println(v2, v4)
                v4 -= v4
            }
            switch fi66 {
            case -703:
                v2 -= -323
            }
        } else {
            v69 := GenS4{F0: v4, F1: v4}
            if ((fi66 >= (-(fi66))) || true) {
                if false {
                    var v70 uint8 = 177
                    _ = v70
                } else {
                    if (4294967295 == 65535) {
                        var v71 float32 = float32(0.1)
                        _ = v71
                    }
                }
                delete(v1, fi66)
            }
            v72 := GenS2{F0: (fi66 + fi66)}
            _ = v69
            _ = v72
        }
        v73 := GenS2{F0: (fi66 << ((4294967295) & 63))}
        if (!false) {
            v74 := GenS4{F0: 42, F1: v3}
            var v75 string = "a"
            v2 = -1
            _ = v74
            _ = v75
        }
        _ = fi66
        _ = v73
    }
    v3 = (v4 & (^(v4)))
    for fi76 := 0; fi76 < 12; fi76++ {
        if (18446744073709551615 < (^(v3))) {
            for fi77 := 0; fi77 < 1; fi77++ {
                if ((!(!(!true))) || ((fi77 - -447) > -1)) {
                    v78 := GenS3{F0: (fi76 + (fi77 << ((65535) & 63))), F1: 255}
                    _ = v78
                }
                v2 = (v2 | (^(v2)))
                var v79 uint8 = 1
                _ = fi77
                _ = v79
            }
            v3 += v3
        }
        _ = fi76
    }
    fmt.Println(v2)
    var v80 uint64 = (v3 / ((359) | 1))
    fmt.Println(v2, v2)
    fmt.Println(v3)
    v3 = 137
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v80
}
