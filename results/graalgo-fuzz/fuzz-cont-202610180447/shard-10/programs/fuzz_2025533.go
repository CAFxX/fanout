package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int8; F2 string }
type GenS6 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint, p1 int) uint16 {
    p0 -= p0
    var v7 uint64 = 978
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 177
    var v3 int8 = 42
    v4 := make(map[uint32]int)
    v5 := []string{("a" + "\n\t"), "\n\t"}
    var v6 int64 = v2
    if false {
        v3 = (v3 & 42)
        v5 = append(v5, ("a" + ("foo bar" + "")))
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                for fi9 := 0; fi9 < 17; fi9++ {
                    if false {
                        switch "a" {
                        case "hello":
                            if true {
                                v3 = 100
                            }
                        }
                    }
                    if (917 > 1) {
                        if (0 > 18446744073709551615) {
                            var v10 int64 = int64(fi9)
                            _ = v10
                        }
                        var v11 uint64 = 0
                        v12 := GenS1{F0: 42, F1: (v1 % (((v1 + 1)) | 1))}
                        _ = v11
                        _ = v12
                    }
                    v2 = v6
                    v13 := GenS1{F0: 0, F1: genFn0(0, (fi9 + fi9))}
                    _ = fi9
                    _ = v13
                }
            } else {
                v6 = (v6 + v2)
            }
            v6 = (v6 ^ v6)
            switch "" {
            case "a":
                var v14 string = "a"
                _ = v14
            case "":
                fmt.Println(v3)
            default:
                if true {
                    var v15 int64 = 0
                    v15 -= -1
                    _ = v15
                }
            }
            _ = fi8
        }
    } else {
        v16 := GenS2{F0: 127, F1: 1}
        _ = v16
    }
    _ = len(v5)
    v1 -= v1
    fmt.Println(v1, v3, v2)
    v17 := GenS1{F0: 42, F1: (v1 ^ (^(v1)))}
    var v18 int = 0
    if (!true) {
        v19 := GenS1{F0: -94, F1: v1}
        if false {
            v1 -= v1
            v1 = v1
            v20 := v4[236]
            _ = v20
        }
        _ = v19
    }
    var v21 float64 = ((((0.1 / 1.0) - 0.0) - -2.25) - -2.25)
    var v22 uint = 2
    switch v18 {
    case 128:
        v6 = (-(v2))
    case 4294967295:
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                switch (fi24 << (((v22 >> ((v22) & 63))) & 63)) {
                case 127:
                    var v25 int64 = (v6 & -1)
                    _ = v25
                case 32767:
                    v26 := GenS2{F0: v18, F1: 176}
                    _ = v26
                case 128:
                    var v27 uint16 = (v1 / (((v1 << (((v22 >> ((v22) & 63))) & 63))) | 1))
                    _ = v27
                default:
                    switch "" {
                    case "foo bar":
                        v28 := GenS3{F0: 2, F1: 2, F2: 4294967295}
                        _ = v28
                    default:
                        v18 += fi23
                    }
                }
                _ = fi24
            }
            _ = fi23
        }
    }
    var v29 uint16 = genFn0(v22, v18)
    v18 += v18
    if (127 != v3) {
        var v30 uint32 = 387
        _ = v30
    }
    switch v18 {
    case 32768:
        switch (("a" + (("" + "\n\t") + "\n\t")) + "a") {
        case "":
            if (!false) {
                switch "a" {
                case "a":
                    v5 = append(v5, "\n\t")
                }
                if (true && false) {
                    if true {
                        v31 := GenS3{F0: 2, F1: 1, F2: 42}
                        fmt.Println(v3)
                        _ = v31
                    } else {
                        var v32 uint16 = v29
                        _ = v32
                    }
                }
                for fi33 := 0; fi33 < 4; fi33++ {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        var v35 bool = ((^(v22)) != 4294967295)
                        v36 := GenS2{F0: v18, F1: -2147483648}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    var v37 float32 = (-2.25 / ((1e10 - ((1e10 + 3.14159) - -2.25)) - (0.0 + 3.14159)))
                    v22 = (^(v22))
                    _ = fi33
                    _ = v37
                }
            } else {
                if (!(int32(v3) >= -2147483648)) {
                    if true {
                        switch 736 {
                        case 9223372036854775807:
                            v38 := v4[2147483647]
                            _ = v38
                        case 798:
                            v3 = v3
                        }
                        var v39 uint8 = 42
                        var v40 int32 = 2147483647
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS2{F0: (v18 ^ -1), F1: -1}
                        _ = v41
                    }
                    _ = len(v4)
                } else {
                    for fi42 := 0; fi42 < 9; fi42++ {
                        _ = len(v4)
                        _ = fi42
                    }
                    v3 = v3
                    v43 := GenS1{F0: 2147483647, F1: (v1 | (v1 % (((v1 - (v1 * 0))) | 1)))}
                    _ = v43
                }
                if true {
                    var v44 uint64 = 42
                    if (("hello" != "\n\t") || ((((0.1 - -2.25) + (-2.25 * float32((((float64(0.1) + (3.14159 + -2.25)) / 1.5) * v21)))) - 3.14159) > 1e10)) {
                        var v45 uint = (v22 + (^(v22)))
                        fmt.Println(v21, v18, v44)
                        _ = v45
                    }
                    _ = v44
                } else {
                    switch "a" {
                    case "a":
                        for fi46 := 0; fi46 < 14; fi46++ {
                            var v47 int16 = 32767
                            var v48 uint16 = (v29 % ((983) | 1))
                            _ = fi46
                            _ = v47
                            _ = v48
                        }
                    case "foo bar":
                        if true {
                            var v49 uint8 = 1
                            fmt.Println(v49, v1, v1)
                            v50 := GenS3{F0: 42, F1: 127, F2: 531}
                            _ = v49
                            _ = v50
                        } else {
                            var v51 bool = ((("hello" + "foo bar") + "foo bar") >= "a")
                            _ = v51
                        }
                    }
                }
            }
        case "hello":
            var v52 uint8 = 255
            _ = v52
        case "a":
            v3 = (v3 * v3)
        default:
            for fi53 := 0; fi53 < 18; fi53++ {
                for fi54 := 0; fi54 < 14; fi54++ {
                    v2 = 390
                    if (float64(0.1) == v21) {
                        v55 := GenS4{F0: (fi53 >> (((v22 - v22)) & 63))}
                        v56 := GenS1{F0: -1, F1: v29}
                        _ = v55
                        _ = v56
                    }
                    _ = fi54
                }
                v57 := v4[65535]
                fmt.Println(fi53, v18)
                _ = fi53
                _ = v57
            }
        }
    case 223:
        v58 := GenS5{F0: "hello", F1: (v3 | (v3 * (v3 / ((v3) | 1)))), F2: ("\n\t" + "a")}
        _ = v58
    }
    if true {
        for fi59 := 0; fi59 < 20; fi59++ {
            v60 := GenS6{F0: 42, F1: -32768}
            _ = fi59
            _ = v60
        }
    }
    fmt.Println(v22, v2)
    var v61 uint8 = 14
    var v62 int = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
