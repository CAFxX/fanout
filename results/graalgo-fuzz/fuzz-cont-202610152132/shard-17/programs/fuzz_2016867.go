package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 int; F1 int16; F2 string }
type GenS6 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS8 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 uint64) int32 {
    v5 := []int{-1}
    var v6 uint32 = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 & 0)
}
func genFn1(p0 int8, p1 uint32, p2 int32) uint32 {
    var v7 uint8 = 2
    var v8 uint8 = (v7 % ((v7) | 1))
    if (42 > 56) {
        if ((p1 != 342) || (!true)) {
            switch int(p1) {
            case 4294967295:
                v9 := GenS1{F0: (p1 * uint32(v8)), F1: v7, F2: 1}
                _ = v9
            case 127:
                v10 := []int{9223372036854775807}
                _ = v10
            case -2147483648:
                p1 -= p1
            default:
                fmt.Println(v8, p1, p2)
            }
            fmt.Println(v7)
        }
    } else {
        var v11 uint64 = 615
        p0 = 42
        _ = v11
    }
    if false {
        fmt.Println(p2)
    } else {
        var v12 float32 = (1.5 + 0.0)
        var v13 int16 = 42
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 | 67)
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    v4 := make(map[string]bool)
    var v14 uint64 = 878
    if (42 == 18446744073709551615) {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint64 = (v14 | (^(v14)))
            _ = fi15
            _ = v16
        }
        v17 := []int{-237, 9223372036854775807, 42}
        _ = v17
    }
    if (2147483648 > 4294967295) {
        v14 -= v14
        switch ((("\n\t" + "") + "a") + (("\n\t" + ((("a" + "a") + "foo bar") + ("a" + "a"))) + "")) {
        case "hello":
            v18 := GenS2{F0: 42}
            _ = v18
        case "":
            v1 = v1
        }
        var v19 float32 = 1e10
        _ = v19
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 13; fi21++ {
            v1 += v1
            v22 := GenS2{F0: (v14 << ((408) & 63))}
            switch -357 {
            case 9223372036854775807:
                switch ("foo bar" + (("foo bar" + "foo bar") + ("" + (("hello" + "") + "foo bar")))) {
                case "":
                    v14 = (v14 % ((v14) | 1))
                default:
                    if ((0 < 65535) && (1 <= 186)) {
                        var v23 int16 = 206
                        v24 := GenS2{F0: (v14 & 1)}
                        v1 = 109
                        _ = v23
                        _ = v24
                    } else {
                        var v25 float64 = 0.1
                        if (!false) {
                            switch ("hello" + ("" + (("a" + "\n\t") + ""))) {
                            case "foo bar":
                                switch "" {
                                case "":
                                    if (!(float64(1.5) < 1e10)) {
                                        v1 = int8(fi21)
                                        if ((0.0 / 1.5) != 0.0) {
                                            if (v14 <= v14) {
                                                v26 := []int{-1, 42, fi21}
                                                var v27 uint16 = 2
                                                _ = v26
                                                _ = v27
                                            }
                                        } else {
                                            v28 := GenS3{F0: 0}
                                            _ = v28
                                        }
                                    } else {
                                        var v29 int32 = genFn0(-2147483648, (^(v14)))
                                        _ = v29
                                    }
                                case "\n\t":
                                    v30 := GenS4{F0: -1, F1: (false || ((false && (false || true)) || (127 <= 255))), F2: v1}
                                    _ = v30
                                default:
                                    v31 := GenS5{F0: fi21, F1: -32768, F2: "a"}
                                    _ = v31
                                }
                            case "\n\t":
                                var v32 uint = 255
                                _ = v32
                            default:
                                var v33 uint32 = 42
                                v3 += 0.1
                                var v34 uint = 18446744073709551615
                                _ = v33
                                _ = v34
                            }
                            if true {
                                var v35 float32 = 3.14159
                                v35 = (v35 + float32(((float64(v35) - ((0.0 * v25) + v3)) * v25)))
                                v14 -= 327
                                _ = v35
                            }
                        }
                        v25 = float64(3.14159)
                        _ = v25
                    }
                }
            case -662:
                v36 := []int{}
                v3 += float64(((3.14159 * (1e10 + (1e10 + 0.1))) - 0.0))
                var v37 uint32 = 65535
                _ = v36
                _ = v37
            case -1:
                if (!false) {
                    v14 += v14
                }
            }
            var v38 int32 = 306
            var v39 uint32 = 295
            var v40 uint = 0
            fmt.Println(fi20)
            _ = fi21
            _ = v22
            _ = v38
            _ = v39
            _ = v40
        }
        switch (fi20 & (fi20 % ((fi20) | 1))) {
        case 738:
            v41 := GenS4{F0: 1, F1: (255 <= 933), F2: (v1 - 38)}
            _ = v41
        case 266:
            if true {
                v42 := []int{}
                _ = v42
            } else {
                fmt.Println(fi20, v1)
            }
        case 4294967295:
            v14 -= v14
        default:
            v1 = (v1 << ((4294967295) & 63))
        }
        _ = fi20
    }
    for fi43 := 0; fi43 < 12; fi43++ {
        for fi44 := 0; fi44 < 17; fi44++ {
            v45 := GenS5{F0: (-(fi44)), F1: -32768, F2: ("foo bar" + "foo bar")}
            v46 := GenS6{F0: uint32(v14), F1: 411, F2: 42}
            var v47 float32 = 0.0
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        var v48 int8 = v1
        _ = fi43
        _ = v48
    }
    switch "hello" {
    case "hello":
        var v49 string = (((("\n\t" + ("foo bar" + "\n\t")) + "hello") + "hello") + "a")
        _ = v49
    case "\n\t":
        v4["a"] = (false || (255 != 0))
    case "foo bar":
        if (-2147483648 > 1) {
            fmt.Println(v1)
        } else {
            v1 -= -86
            if false {
                var v50 int64 = 535
                _ = v50
            } else {
                v51 := v4[("hello" + "foo bar")]
                v52 := GenS1{F0: 688, F1: 3, F2: -412}
                var v53 int32 = 171
                _ = v51
                _ = v52
                _ = v53
            }
        }
    default:
        var v54 string = ("\n\t" + "a")
        _ = v54
    }
    var v55 int8 = 100
    v56 := GenS7{F0: 42, F1: genFn0(89, v14), F2: 255}
    v57 := GenS2{F0: v14}
    v58 := []int{0, 392}
    var v59 uint8 = 128
    v60 := v2[9223372036854775807]
    for fi61 := 0; fi61 < 3; fi61++ {
        var v62 uint = 3
        var v63 uint32 = 2147483647
        fmt.Println(v62, fi61)
        _ = fi61
        _ = v62
        _ = v63
    }
    switch ((v60 + v60) + (("" + v60) + "hello")) {
    case "\n\t":
        if true {
            var v64 int = -1
            v65 := GenS8{F0: true, F1: 42}
            _ = v64
            _ = v65
        } else {
            v66 := v2[32767]
            if len(v58) > 0 {
                v58[0] = -2147483648
            }
            switch 1 {
            case 32767:
                if (!false) {
                    v14 += 0
                    v58 = append(v58, 4294967295)
                    for fi67 := 0; fi67 < 14; fi67++ {
                        switch v66 {
                        case "a":
                            if len(v58) >= 2 {
                                v58 = v58[1:len(v58)]
                            }
                        case "hello":
                            fmt.Println(v1)
                        case "":
                            for fi68 := 0; fi68 < 9; fi68++ {
                                v14 -= (v14 / (((v14 * v14)) | 1))
                                _ = fi68
                            }
                        }
                        v2[(fi67 & (fi67 / (((fi67 - (fi67 ^ (fi67 + fi67)))) | 1)))] = ((v66 + ((v66 + (v60 + "foo bar")) + v66)) + ((v60 + (v60 + ("" + "foo bar"))) + ((v66 + "foo bar") + v66)))
                        v55 = v55
                        _ = fi67
                    }
                } else {
                    var v69 uint32 = genFn1((v55 << ((550) & 63)), genFn1(v55, 65535, 1), -1)
                    _ = v69
                }
            }
            _ = v66
        }
    case "a":
        if true {
            var v70 int16 = 1000
            _ = v70
        } else {
            if (false && false) {
                var v71 uint32 = 1
                _ = v71
            } else {
                var v72 bool = false
                for fi73 := 0; fi73 < 12; fi73++ {
                    v74 := GenS5{F0: (fi73 - -2147483648), F1: 696, F2: v60}
                    _ = fi73
                    _ = v74
                }
                _ = v72
            }
        }
    case "hello":
        v4["foo bar"] = (v3 >= 1.5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
