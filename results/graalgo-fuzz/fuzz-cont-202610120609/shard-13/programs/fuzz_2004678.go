package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int; F2 uint }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int16; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 bool; F2 int8 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 bool, p1 string, p2 int) int8 {
    var v5 uint64 = 4294967295
    switch "hello" {
    case "hello":
        if p0 {
            var v6 float64 = -2.25
            var v7 string = (((p1 + p1) + p1) + p1)
            _ = v6
            _ = v7
        }
    case "foo bar":
        var v8 uint64 = 0
        _ = v8
    case "\n\t":
        var v9 uint64 = 2
        _ = v9
    default:
        var v10 float64 = 0.1
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint64 = v5
        if (-374 > -166) {
            v13 := GenS1{F0: 376}
            p2 += fi11
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                p1 = p1
                switch (^(p2)) {
                case 32768:
                    fmt.Println(p2)
                case 127:
                    if p0 {
                        var v15 int16 = 1
                        fmt.Println(fi14, p0, fi11)
                        if (1 > (^(v12))) {
                            fmt.Println(p1, p2)
                            var v16 uint16 = 32767
                            v17 := []int{(fi14 >> ((42) & 63)), (fi14 << ((255) & 63))}
                            p0 = p0
                            if ((262 < 0) && false) {
                                var v18 int64 = 889
                                v16 = 0
                                v5 -= (v12 * v12)
                                _ = v18
                            } else {
                                if p0 {
                                    v16 -= v16
                                    var v19 int32 = 742
                                    _ = v19
                                }
                            }
                            v20 := GenS2{F0: (fi11 - fi14), F1: -9223372036854775808, F2: 395}
                            p2 += (p2 << ((1) & 63))
                            _ = v16
                            _ = v17
                            _ = v20
                        } else {
                            v21 := GenS3{F0: 1, F1: 1, F2: fi11}
                            _ = v21
                        }
                        _ = v15
                    } else {
                        var v22 string = ("" + p1)
                        v23 := GenS2{F0: fi14, F1: (p2 - fi14), F2: 4294967295}
                        var v24 uint = 2
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    switch (fi14 & p2) {
                    case 2147483647:
                        v25 := []int{p2, 32767, -1, p2}
                        var v26 uint32 = 429
                        switch "\n\t" {
                        case "":
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                        case "\n\t":
                            if (v26 != 1) {
                                fmt.Println(p1, fi11)
                                if (!p0) {
                                    var v27 int64 = -9223372036854775808
                                    _ = v27
                                }
                            }
                        default:
                            v12 = (v12 / ((4294967295) | 1))
                        }
                        if ((v12 | (v5 % ((4294967295) | 1))) >= (v5 | (v12 / ((18446744073709551615) | 1)))) {
                            v5 -= 18446744073709551615
                            v12 = uint64(v26)
                            _ = len(v25)
                            if (454 < 42) {
                                if (128 > fi11) {
                                    if len(v25) > 0 {
                                        v25[0] = 9223372036854775807
                                    }
                                    switch p1 {
                                    case "\n\t":
                                        if (((0.1 / float32(1.5)) - -2.25) <= 0.1) {
                                            if len(v25) > 0 {
                                                v25[0] = (fi14 ^ fi11)
                                            }
                                        } else {
                                            v5 = (v12 ^ (v5 / (((^(v12))) | 1)))
                                            fmt.Println(p1, fi14, v26)
                                            v26 = v26
                                        }
                                    case "":
                                        var v28 uint16 = 255
                                        _ = v28
                                    case "a":
                                        fmt.Println(fi11)
                                    default:
                                        fmt.Println(v12, p2, fi11)
                                    }
                                    var v29 bool = (42 >= 127)
                                    if false {
                                        fmt.Println(v12)
                                        p2 = p2
                                    } else {
                                        v30 := GenS1{F0: 740}
                                        fmt.Println(fi14, v29, p0)
                                        var v31 int32 = -2147483648
                                        if p0 {
                                            switch (p1 + p1) {
                                            case "hello":
                                                v29 = ((v31 + (^(v31))) <= v31)
                                            case "foo bar":
                                                var v32 uint64 = v12
                                                _ = v32
                                            case "a":
                                                var v33 uint8 = 197
                                                _ = v33
                                            default:
                                                p0 = true
                                            }
                                            var v34 float64 = ((1e10 - (0.0 * (3.14159 / 3.14159))) * ((0.0 * (((1e10 / ((3.14159 / float64(float32(1e10))) - 1.5)) - 1e10) / 1.5)) + 1e10))
                                            var v35 int64 = -9223372036854775808
                                            _ = v34
                                            _ = v35
                                        } else {
                                            v36 := GenS1{F0: 255}
                                            switch p1 {
                                            case "a":
                                                if len(v25) > 0 {
                                                    v25[0] = (^(p2))
                                                }
                                            case "hello":
                                                v37 := GenS3{F0: -549, F1: v5, F2: (fi14 / ((-162) | 1))}
                                                _ = v37
                                            case "\n\t":
                                                p0 = (((!(433 >= v31)) || (!p0)) || (0 < 449))
                                            default:
                                                switch (fi14 | fi11) {
                                                case 476:
                                                    var v38 float64 = (1.5 + (((0.0 * (1.5 + 0.0)) + 1.5) * 1e10))
                                                    _ = v38
                                                }
                                            }
                                            fmt.Println(fi14, v26)
                                            var v39 int64 = 42
                                            v40 := GenS3{F0: -32768, F1: (v5 / (((^(v12))) | 1)), F2: (fi11 ^ (fi11 | fi14))}
                                            _ = v36
                                            _ = v39
                                            _ = v40
                                        }
                                        var v41 string = (("" + "a") + (p1 + ((p1 + p1) + "hello")))
                                        _ = v30
                                        _ = v31
                                        _ = v41
                                    }
                                    _ = v29
                                }
                                if (p0 || (1.5 >= 0.0)) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v42 := GenS4{F0: 255, F1: 42, F2: 1}
                                        _ = v42
                                    }
                                    p0 = p0
                                }
                                v12 = (^(v5))
                            }
                            v26 -= v26
                            v12 = v5
                            v43 := GenS3{F0: 32767, F1: v12, F2: fi14}
                            _ = v43
                        } else {
                            var v44 int = fi14
                            fmt.Println(fi11, fi11, v26)
                            _ = v44
                        }
                        v45 := GenS4{F0: 224, F1: 833, F2: -2147483648}
                        _ = v25
                        _ = v26
                        _ = v45
                    case -1:
                        if (((!p0) && p0) && p0) {
                            v46 := []int{(fi14 << ((26) & 63)), 4294967295, fi11}
                            if len(v46) >= 2 {
                                v46 = v46[1:len(v46)]
                            }
                            _ = v46
                        }
                        if p0 {
                            p1 = "\n\t"
                        } else {
                            fmt.Println(v12)
                            var v47 float64 = 1.5
                            _ = v47
                        }
                        fmt.Println(fi14)
                    case 9223372036854775807:
                        v48 := GenS5{F0: p1}
                        _ = v48
                    }
                }
                v49 := []int{1, 315, 872}
                var v50 uint32 = 2147483648
                _ = fi14
                _ = v49
                _ = v50
            }
            switch fi11 {
            case -811:
                switch (((("a" + "\n\t") + ("\n\t" + "")) + "hello") + "\n\t") {
                case "\n\t":
                    v51 := GenS6{F0: p1, F1: p0, F2: 35}
                    _ = v51
                case "a":
                    p2 += (fi11 ^ 42)
                case "hello":
                    v5 += v12
                default:
                    switch 9223372036854775807 {
                    case -101:
                        switch (p2 * (fi11 + (fi11 - 637))) {
                        case 2147483647:
                            fmt.Println(fi11, fi11)
                        }
                    case 496:
                        var v52 bool = ("foo bar" >= p1)
                        _ = v52
                    case -2147483648:
                        var v53 int = (fi11 * (fi11 | 0))
                        _ = v53
                    default:
                        for fi54 := 0; fi54 < 1; fi54++ {
                            v55 := GenS6{F0: (p1 + p1), F1: (42 < (^(v5))), F2: 42}
                            p0 = false
                            var v56 int64 = 372
                            v56 = -656
                            fmt.Println(v56, fi11)
                            v57 := GenS7{F0: 1, F1: p1}
                            _ = fi54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                    }
                }
            default:
                p2 = (fi11 + (p2 ^ (p2 << ((65535) & 63))))
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1(p0 uint32) uint {
    p0 = 2
    var v58 uint32 = 4294967295
    _ = p0
    _ = v58
    return 0
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"", "foo bar", "a"}
    v3 := []string{}
    v4 := make(map[string]bool)
    var v59 uint16 = 2
    delete(v4, "\n\t")
    v1 = (v1 ^ v1)
    var v60 int32 = (v1 % ((42) | 1))
    var v61 uint8 = 174
    v62 := v4["a"]
    var v63 float64 = float64(-2.25)
    v3 = append(v3, (((("hello" + (("" + (("a" + ("foo bar" + ("\n\t" + ("foo bar" + "\n\t")))) + ("a" + "\n\t"))) + ("foo bar" + "a"))) + "\n\t") + "\n\t") + ("hello" + ("foo bar" + ""))))
    if false {
        for fi64 := 0; fi64 < 19; fi64++ {
            var v65 uint32 = 153
            v1 -= v60
            v3 = append(v3, "\n\t")
            _ = fi64
            _ = v65
        }
    }
    if true {
        switch ("" + ("foo bar" + (("foo bar" + "") + ""))) {
        case "hello":
            v66 := GenS8{F0: v61, F1: uint(v59), F2: v60}
            _ = v66
        case "a":
            if v62 {
                for fi67 := 0; fi67 < 3; fi67++ {
                    v68 := v4[("a" + ("" + "\n\t"))]
                    _ = fi67
                    _ = v68
                }
                _ = len(v3)
                if ((v59 * v59) != v59) {
                    v69 := GenS6{F0: "foo bar", F1: (42 != 32767), F2: -102}
                    _ = v69
                }
            } else {
                for fi70 := 0; fi70 < 20; fi70++ {
                    v1 = -2147483648
                    _ = fi70
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                if (!v62) {
                    fmt.Println(v62, v60, v60)
                    if v62 {
                        for fi71 := 0; fi71 < 7; fi71++ {
                            v1 = (v60 / (((v60 | (v1 << ((65535) & 63)))) | 1))
                            var v72 bool = (v62 || (!(!v62)))
                            switch (fi71 << ((4294967295) & 63)) {
                            case 127:
                                delete(v4, "hello")
                            case -205:
                                var v73 string = "a"
                                _ = v73
                            default:
                                v74 := GenS3{F0: -427, F1: 18446744073709551615, F2: (fi71 ^ (fi71 * -237))}
                                _ = v74
                            }
                            _ = fi71
                            _ = v72
                        }
                        v59 += 883
                        var v75 int64 = 0
                        _ = v75
                    }
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
            }
        case "":
            if v62 {
                for fi76 := 0; fi76 < 16; fi76++ {
                    var v77 uint64 = 76
                    v78 := v4["a"]
                    if v62 {
                        fmt.Println(v63, v1)
                    }
                    _ = fi76
                    _ = v77
                    _ = v78
                }
            }
        default:
            if v62 {
                v59 = (v59 & 32767)
                var v79 uint32 = 2147483647
                var v80 string = ""
                _ = v79
                _ = v80
            }
        }
    }
    fmt.Println(v59, v63, v61)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
