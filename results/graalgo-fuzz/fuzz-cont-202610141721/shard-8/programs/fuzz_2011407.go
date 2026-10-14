package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 uint16, p1 uint64, p2 string) uint32 {
    if true {
        p0 = p0
        fmt.Println(p1, p0, p0)
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := GenS1{F0: p2}
            _ = fi7
            _ = v8
        }
    } else {
        v9 := GenS1{F0: p2}
        var v10 uint16 = p0
        v11 := GenS1{F0: ""}
        _ = v10
        _ = v11
        _ = v9
    }
    fmt.Println(p1, p0, p0)
    var v12 int = 348
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 263
}
func genFn1(p0 uint16, p1 uint, p2 bool) bool {
    var v13 int = 127
    fmt.Println(v13, p0)
    if false {
        var v14 float32 = (3.14159 / 3.14159)
        _ = v14
    }
    if p2 {
        var v15 uint32 = 0
        _ = v15
    }
    fmt.Println(v13)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return ((1 <= uint32(p1)) && p2)
}

func main() {
    var v1 float32 = 3.14159
    var v2 float64 = 1.5
    v3 := []string{"hello", ((("hello" + (("" + "a") + ("hello" + ("hello" + "\n\t")))) + ("a" + "a")) + ("\n\t" + "a"))}
    var v4 uint32 = 522
    var v5 int8 = 1
    v6 := []uint32{v4, v4, (v4 * v4)}
    var v16 uint32 = 1
    var v17 float32 = (v1 - 0.0)
    v17 += 1e10
    v4 -= v16
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 uint = 18446744073709551615
        switch fi18 {
        case 127:
            v1 -= (float32((((v2 - float64(v1)) + (v2 + (float64((v1 + 1e10)) + 0.0))) + 0.1)) * (v1 * v17))
        case 136:
            switch "a" {
            case "foo bar":
                var v20 int8 = v5
                _ = v20
            case "a":
                var v21 uint32 = 180
                _ = v21
            case "\n\t":
                var v22 uint8 = 255
                _ = v22
            default:
                v23 := GenS2{F0: 0, F1: v5, F2: uint8(v19)}
                _ = v23
            }
        case 32767:
            if true {
                var v24 int32 = 2147483647
                _ = v24
            } else {
                v5 = -128
            }
        }
        v25 := GenS2{F0: 127, F1: -57, F2: 2}
        _ = fi18
        _ = v19
        _ = v25
    }
    v26 := GenS1{F0: "\n\t"}
    if true {
        var v27 int64 = 1
        if false {
            if false {
                if (uint8(v4) == 24) {
                    if false {
                        v28 := GenS3{F0: v27}
                        v5 -= 1
                        _ = v28
                    }
                } else {
                    var v29 int = 9223372036854775807
                    v27 = v27
                    _ = v29
                }
            } else {
                var v30 int = 260
                v16 += v16
                _ = v30
            }
            v5 = (v5 % (((v5 << ((480) & 63))) | 1))
            if len(v6) > 0 {
                v6[0] = v4
            }
        }
        switch -9223372036854775808 {
        case 667:
            var v31 string = "hello"
            _ = v31
        case 0:
            for fi32 := 0; fi32 < 15; fi32++ {
                if (42 != 255) {
                    switch ((("hello" + "foo bar") + ("" + "hello")) + ((("hello" + "foo bar") + "a") + ("\n\t" + ""))) {
                    case "foo bar":
                        for fi33 := 0; fi33 < 11; fi33++ {
                            var v34 float32 = v17
                            v4 -= v16
                            var v35 float64 = v2
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    case "hello":
                        if false {
                            if len(v6) > 0 {
                                v6[0] = v4
                            }
                            for fi36 := 0; fi36 < 10; fi36++ {
                                switch 32768 {
                                case 4294967295:
                                    if (!false) {
                                        if ((-(fi32)) >= -9223372036854775808) {
                                            v16 = 592
                                            var v37 float64 = 1e10
                                            _ = v37
                                        } else {
                                            if (42 > -1000) {
                                                var v38 int = (fi32 * (fi32 - fi32))
                                                v3 = append(v3, "hello")
                                                _ = v38
                                            }
                                            switch fi36 {
                                            case -9223372036854775808:
                                                fmt.Println(v1, fi32)
                                            case 32767:
                                                switch "foo bar" {
                                                case "hello":
                                                    v16 = v16
                                                case "\n\t":
                                                    v39 := GenS2{F0: 93, F1: (v5 * (v5 >> ((911) & 63))), F2: 1}
                                                    var v40 int32 = 509
                                                    _ = v39
                                                    _ = v40
                                                }
                                            case 9223372036854775807:
                                                v4 += v4
                                            default:
                                                fmt.Println(v4)
                                            }
                                            v41 := GenS3{F0: v27}
                                            var v42 string = ""
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            _ = v41
                                            _ = v42
                                        }
                                        var v43 int32 = 0
                                        v44 := GenS4{F0: 1}
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if true {
                                            v27 = v27
                                            v45 := GenS4{F0: 373}
                                            v2 = (v2 / v2)
                                            _ = v45
                                        } else {
                                            switch ("a" + "") {
                                            case "\n\t":
                                                v5 = v5
                                            case "a":
                                                fmt.Println(v27, v5)
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if ((18446744073709551615 > 2) || (18446744073709551615 < 42)) {
                                                var v46 uint = 255
                                                v47 := GenS1{F0: ("\n\t" + ((("a" + ("hello" + "")) + ((("\n\t" + "foo bar") + "\n\t") + "")) + ""))}
                                                _ = v46
                                                _ = v47
                                            }
                                        }
                                        var v48 float64 = ((1e10 - v2) + (v2 * float64(((v1 / ((v1 / (v1 * v1)) + (v17 - -2.25))) / v1))))
                                        _ = v48
                                    }
                                }
                                fmt.Println(fi32, v16, v27)
                                _ = fi36
                            }
                        } else {
                            v27 += v27
                        }
                    case "\n\t":
                        fmt.Println(v17, v2)
                    }
                } else {
                    var v49 uint = 699
                    switch "hello" {
                    case "foo bar":
                        var v50 uint16 = 32767
                        _ = v50
                    case "a":
                        v49 = v49
                    default:
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    }
                    _ = v49
                }
                switch fi32 {
                case -621:
                    if (!(!false)) {
                        var v51 uint = 18446744073709551615
                        for fi52 := 0; fi52 < 3; fi52++ {
                            var v53 int = fi52
                            _ = fi52
                            _ = v53
                        }
                        _ = v51
                    } else {
                        v54 := GenS5{F0: 180, F1: 2147483647, F2: 1}
                        v4 -= (v16 + v16)
                        for fi55 := 0; fi55 < 10; fi55++ {
                            v16 += (v16 & v16)
                            _ = fi55
                        }
                        _ = v54
                    }
                default:
                    if (false || (!(true && true))) {
                        for fi56 := 0; fi56 < 16; fi56++ {
                            switch "" {
                            case "a":
                                switch "a" {
                                case "hello":
                                    v5 += v5
                                }
                            }
                            var v57 float32 = v17
                            if false {
                                _ = len(v3)
                            } else {
                                v16 = v16
                                var v58 bool = (-312 < 42)
                                if (v17 == 0.1) {
                                    var v59 uint = 1
                                    _ = v59
                                } else {
                                    if (-2147483648 <= fi32) {
                                        v60 := GenS4{F0: 42}
                                        v27 = 9223372036854775807
                                        v61 := GenS2{F0: 255, F1: (v5 & (^(v5))), F2: 1}
                                        _ = v60
                                        _ = v61
                                    } else {
                                        var v62 bool = v58
                                        _ = v62
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = (v16 ^ 2147483648)
                                    }
                                    v63 := GenS1{F0: ("hello" + ("hello" + "hello"))}
                                    _ = v63
                                }
                                _ = v58
                            }
                            _ = fi56
                            _ = v57
                        }
                    }
                }
                _ = fi32
            }
        default:
            _ = len(v6)
        }
        _ = v27
    }
    var v64 int8 = v5
    var v65 float64 = -2.25
    var v66 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v66
}
