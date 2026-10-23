package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint8; F1 bool }
func genFn0(p0 uint32, p1 uint64, p2 int) uint64 {
    var v7 uint = 232
    v8 := []int{p2, (p2 * (p2 * p2))}
    v8 = append(v8, (p2 * p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 << (((^(v7))) & 63))
}
func genFn1() uint32 {
    var v9 float32 = 0.1
    v10 := GenS1{F0: genFn0(65535, 4294967295, -2147483648), F1: true}
    _ = v10
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint16 = 712
    v2 := []int{9223372036854775807}
    v3 := []int{128, 42, 128, -560}
    var v4 int8 = 0
    var v5 int64 = -394
    v6 := []string{"\n\t"}
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    fmt.Println(v5, v4)
    v6 = append(v6, "foo bar")
    v13 := GenS1{F0: 42, F1: true}
    fmt.Println(v1)
    var v14 uint = 2
    _ = len(v3)
    v4 += v4
    if false {
        switch 32767 {
        case 32767:
            switch (("a" + "foo bar") + ("" + "")) {
            case "\n\t":
                v15 := GenS2{F0: (v5 & (v5 | 177)), F1: (388 != v14)}
                _ = v15
            case "a":
                var v16 float32 = float32(-2.25)
                _ = v16
            case "hello":
                var v17 uint64 = 2
                _ = v17
            }
        case -863:
            v4 = v4
        case 657:
            for fi18 := 0; fi18 < 9; fi18++ {
                if (false || ((0 == 1000) && true)) {
                    var v19 string = "foo bar"
                    if (42 == 2) {
                        v20 := GenS2{F0: (v5 / ((v5) | 1)), F1: (-1 <= 42)}
                        v14 += 2
                        _ = v20
                    } else {
                        v21 := GenS3{F0: v14}
                        v1 = v1
                        _ = v21
                    }
                    _ = v19
                }
                if (false && (!(v1 == (v1 + 255)))) {
                    switch ("\n\t" + "foo bar") {
                    case "\n\t":
                        switch (("\n\t" + "hello") + "hello") {
                        case "foo bar":
                            v5 -= 1
                        case "hello":
                            switch 489 {
                            case -188:
                                v14 = (^(v14))
                            case 0:
                                fmt.Println(v14)
                            default:
                                v22 := GenS4{F0: v1, F1: 83}
                                _ = v22
                            }
                        default:
                            if (genFn1() < 886) {
                                v1 -= (v1 >> (((v14 - 255)) & 63))
                                v2 = append(v2, 128)
                                var v23 uint64 = 18446744073709551615
                                _ = v23
                            }
                        }
                    case "foo bar":
                        var v24 float64 = 0.1
                        _ = v24
                    case "a":
                        if ("foo bar" > "a") {
                            fmt.Println(v14, fi18, v1)
                            if true {
                                v1 -= (v1 ^ uint16(v14))
                                v14 += 0
                                for fi25 := 0; fi25 < 9; fi25++ {
                                    var v26 int16 = 697
                                    v14 = 1
                                    v3 = append(v3, (fi18 << (((v14 >> (((v14 << (((^(v14))) & 63))) & 63))) & 63)))
                                    _ = fi25
                                    _ = v26
                                }
                            } else {
                                var v27 float32 = ((1e10 - 3.14159) * 1e10)
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    var v29 bool = true
                                    v6 = append(v6, "\n\t")
                                    v30 := GenS2{F0: (v5 << ((18446744073709551615) & 63)), F1: (42 <= (^(v1)))}
                                    switch 117 {
                                    case 2147483647:
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            var v31 int32 = int32(v5)
                                            var v32 uint8 = 42
                                            var v33 float32 = 0.1
                                            var v34 uint16 = v1
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        default:
                                            v35 := GenS2{F0: -1, F1: (65535 == 42)}
                                            _ = v35
                                        }
                                    case -51:
                                        v2 = append(v2, fi28)
                                    case -9223372036854775808:
                                        if (255 >= 255) {
                                            v4 = (-(v4))
                                            switch "\n\t" {
                                            case "foo bar":
                                                switch "a" {
                                                case "foo bar":
                                                    switch ("hello" + "hello") {
                                                    case "":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            var v36 string = ""
                                                            v5 += v5
                                                            var v37 int64 = (v5 * int64(v1))
                                                            _ = v36
                                                            _ = v37
                                                        case "hello":
                                                            v27 += (v27 - v27)
                                                        case "\n\t":
                                                            v38 := GenS5{F0: "a", F1: 239}
                                                            _ = v38
                                                        }
                                                    }
                                                    switch (^(fi18)) {
                                                    case -738:
                                                        var v39 int64 = int64(v4)
                                                        _ = v39
                                                    case 9223372036854775807:
                                                        var v40 uint64 = 1
                                                        _ = v40
                                                    }
                                                default:
                                                    fmt.Println(v29, v29)
                                                }
                                            case "hello":
                                                v29 = ((((^(v14)) >= 42) && (18446744073709551615 != 159)) && v29)
                                            case "":
                                                v29 = true
                                            default:
                                                v27 += float32(0.1)
                                            }
                                            v29 = v29
                                        }
                                    default:
                                        var v41 uint16 = v1
                                        _ = v41
                                    }
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                var v42 int8 = -19
                                _ = v27
                                _ = v42
                            }
                        }
                    default:
                        v14 = v14
                    }
                    var v43 uint16 = v1
                    _ = v43
                } else {
                    fmt.Println(v5)
                }
                for fi44 := 0; fi44 < 9; fi44++ {
                    if ((true || false) && ((!((v14 << (((v14 + v14)) & 63)) < (v14 - (v14 ^ 179)))) || false)) {
                        var v45 float64 = ((-2.25 * 3.14159) - float64(0.1))
                        var v46 uint8 = 127
                        _ = v45
                        _ = v46
                    }
                    if false {
                        v47 := GenS6{F0: (fi44 + (fi44 << ((18446744073709551615) & 63)))}
                        var v48 uint32 = uint32(v1)
                        _ = v47
                        _ = v48
                    }
                    _ = fi44
                }
                _ = fi18
            }
        default:
            fmt.Println(v5, v1, v14)
        }
        switch ("a" + "hello") {
        case "foo bar":
            fmt.Println(v14, v14, v4)
        default:
            v49 := GenS5{F0: "hello", F1: uint64(v1)}
            _ = v49
        }
        if true {
            for fi50 := 0; fi50 < 18; fi50++ {
                var v51 float64 = 0.1
                fmt.Println(v51, v14)
                var v52 int64 = (-(v5))
                _ = fi50
                _ = v51
                _ = v52
            }
        } else {
            var v53 bool = ((^(v14)) != 65535)
            fmt.Println(v53, v5)
            _ = v53
        }
    } else {
        if (true && (false || false)) {
            v54 := GenS7{F0: 559, F1: 0}
            _ = v54
        }
        if false {
            for fi55 := 0; fi55 < 12; fi55++ {
                var v56 bool = (1e10 <= -2.25)
                v1 = (v1 << ((v14) & 63))
                _ = fi55
                _ = v56
            }
            fmt.Println(v1, v4)
        }
    }
    var v57 uint = uint(v1)
    if ("\n\t" >= ("" + (("a" + "foo bar") + "hello"))) {
        _ = len(v2)
        _ = len(v3)
    }
    var v58 uint32 = 42
    var v59 float64 = 0.1
    v1 -= (v1 % ((2) | 1))
    var v60 float64 = (v59 - -2.25)
    if false {
        var v61 uint32 = genFn1()
        switch (("hello" + "\n\t") + (("foo bar" + "") + "")) {
        case "\n\t":
            _ = len(v6)
        default:
            for fi62 := 0; fi62 < 14; fi62++ {
                var v63 uint8 = 2
                for fi64 := 0; fi64 < 5; fi64++ {
                    v65 := GenS8{F0: 128, F1: false}
                    v4 -= (v4 - (v4 % (((v4 ^ 127)) | 1)))
                    _ = fi64
                    _ = v65
                }
                v1 += v1
                _ = fi62
                _ = v63
            }
        }
        for fi66 := 0; fi66 < 6; fi66++ {
            _ = len(v3)
            _ = fi66
        }
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
