package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint32, p1 string, p2 int32) uint8 {
    p1 = p1
    if (true || (((-9223372036854775808 <= 635) && ((p2 << ((255) & 63)) > p2)) && (127 > 42))) {
        var v6 float32 = float32(1.5)
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: (p0 % ((2147483647) | 1))}
            _ = fi7
            _ = v8
        }
        fmt.Println(v6)
        _ = v6
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint8) int16 {
    v9 := GenS2{F0: 255}
    var v10 bool = false
    _ = p0
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 int = 32767
    v2 := []string{"hello"}
    v3 := []uint32{42}
    var v4 string = (("foo bar" + "\n\t") + "a")
    var v5 float64 = 1e10
    v4 = v4
    fmt.Println(v5, v5, v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        _ = len(v3)
        _ = fi11
    }
    v12 := GenS3{F0: -43, F1: "a"}
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    var v13 uint = 18446744073709551615
    for fi14 := 0; fi14 < 6; fi14++ {
        _ = len(v2)
        v13 += (^(v13))
        var v15 uint = v13
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1.5
    if (255 <= 42) {
        v17 := GenS2{F0: 255}
        _ = v17
    }
    v1 = (-(v1))
    for fi18 := 0; fi18 < 18; fi18++ {
        v13 += (v13 ^ (^(v13)))
        _ = len(v2)
        _ = fi18
    }
    fmt.Println(v5)
    for fi19 := 0; fi19 < 6; fi19++ {
        if true {
            switch v4 {
            case "hello":
                switch v4 {
                case "a":
                    fmt.Println(v5)
                case "hello":
                    if true {
                        v4 = v4
                        var v20 uint32 = 0
                        for fi21 := 0; fi21 < 2; fi21++ {
                            var v22 float64 = (v5 / -2.25)
                            _ = fi21
                            _ = v22
                        }
                        _ = v20
                    }
                case "":
                    v13 += v13
                default:
                    v3 = append(v3, 2147483647)
                }
            default:
                switch v4 {
                case "\n\t":
                    fmt.Println(v16, fi19, v16)
                case "":
                    if false {
                        if (false && (42 < 2147483647)) {
                            _ = len(v2)
                        } else {
                            var v23 float64 = v5
                            if true {
                                for fi24 := 0; fi24 < 9; fi24++ {
                                    switch "foo bar" {
                                    case "hello":
                                        if false {
                                            switch fi19 {
                                            case 426:
                                                var v25 string = (v4 + v4)
                                                _ = v25
                                            case 128:
                                                switch "hello" {
                                                case "a":
                                                    v26 := GenS2{F0: 18446744073709551615}
                                                    _ = v26
                                                case "":
                                                    fmt.Println(v4, v16)
                                                case "hello":
                                                    var v27 int32 = -363
                                                    _ = v27
                                                }
                                            default:
                                                if true {
                                                    var v28 uint8 = 127
                                                    if (true && ((false || true) && (2 < 4294967295))) {
                                                        var v29 int8 = -128
                                                        if (-88 != 9223372036854775807) {
                                                            if false {
                                                                v16 -= 0.1
                                                                if true {
                                                                    var v30 bool = (4294967295 <= 655)
                                                                    v1 = (fi19 << ((v13) & 63))
                                                                    _ = v30
                                                                } else {
                                                                    v28 = v28
                                                                    switch fi24 {
                                                                    case 127:
                                                                        switch (v1 / ((v1) | 1)) {
                                                                        case 489:
                                                                            v16 = 0.0
                                                                        }
                                                                    case 1:
                                                                        if (true && false) {
                                                                            if (-466 >= -32768) {
                                                                                var v31 uint16 = 255
                                                                                var v32 int32 = 0
                                                                                if ((v31 >> ((v13) & 63)) == v31) {
                                                                                    _ = len(v2)
                                                                                    v29 = 100
                                                                                }
                                                                                _ = v31
                                                                                _ = v32
                                                                            }
                                                                            fmt.Println(v28)
                                                                        }
                                                                    default:
                                                                        var v33 int16 = -1
                                                                        _ = v33
                                                                    }
                                                                }
                                                                v4 = (v4 + ((v4 + v4) + v4))
                                                            }
                                                        }
                                                        _ = v29
                                                    } else {
                                                        switch (("\n\t" + ((((v4 + (((("a" + v4) + ("a" + "hello")) + v4) + (v4 + ((v4 + v4) + "")))) + v4) + "") + "a")) + v4) {
                                                        case "foo bar":
                                                            switch v4 {
                                                            case "foo bar":
                                                                var v34 float64 = 0.1
                                                                _ = v34
                                                            case "a":
                                                                v13 = (^(v13))
                                                            case "":
                                                                var v35 uint16 = 2
                                                                _ = v35
                                                            }
                                                        default:
                                                            var v36 int16 = 721
                                                            _ = v36
                                                        }
                                                        if true {
                                                            v37 := GenS1{F0: 2147483647}
                                                            v1 += (fi19 & fi19)
                                                            _ = v37
                                                        }
                                                        fmt.Println(fi19)
                                                    }
                                                    _ = v28
                                                } else {
                                                    fmt.Println(v4, v1, v13)
                                                }
                                            }
                                            _ = len(v3)
                                        }
                                    }
                                    v38 := GenS2{F0: v13}
                                    switch "\n\t" {
                                    case "foo bar":
                                        v5 = v23
                                    case "\n\t":
                                        fmt.Println(v4)
                                    case "hello":
                                        if true {
                                            if (154 < 913) {
                                                switch ((v4 + (v4 + ("a" + "hello"))) + v4) {
                                                case "\n\t":
                                                    if ((-(fi19)) <= (v1 << ((v13) & 63))) {
                                                        var v39 string = (v4 + "")
                                                        v16 += v16
                                                        _ = v39
                                                    } else {
                                                        switch ("foo bar" + ((v4 + v4) + v4)) {
                                                        case "\n\t":
                                                            v1 = fi24
                                                        }
                                                        fmt.Println(v1, v16)
                                                        var v40 float32 = 0.1
                                                        _ = v40
                                                    }
                                                case "a":
                                                    fmt.Println(v13, v16, fi24)
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                if true {
                                                    var v41 uint = v13
                                                    switch (v1 | v1) {
                                                    case 9223372036854775807:
                                                        v1 = 4294967295
                                                    default:
                                                        if true {
                                                            var v42 uint16 = 255
                                                            fmt.Println(fi24, fi24, v5)
                                                            if false {
                                                                if (true && true) {
                                                                    var v43 uint16 = v42
                                                                    switch (v1 ^ 128) {
                                                                    case 2147483647:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = v4
                                                                        }
                                                                    }
                                                                    _ = v43
                                                                }
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                switch "foo bar" {
                                                                case "":
                                                                    switch int(v42) {
                                                                    case 32768:
                                                                        var v44 uint = v13
                                                                        _ = v44
                                                                    case -1:
                                                                        if (2147483647 == 65535) {
                                                                            var v45 int = (fi24 & fi24)
                                                                            _ = v45
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(v16)
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        v46 := GenS4{F0: 0, F1: (v1 ^ (-(fi24)))}
                                                                        _ = v46
                                                                    }
                                                                case "a":
                                                                    var v47 int = v1
                                                                    _ = v47
                                                                case "hello":
                                                                    switch 0 {
                                                                    case -9223372036854775808:
                                                                        var v48 bool = false
                                                                        _ = v48
                                                                    case 127:
                                                                        if (78 <= 255) {
                                                                            var v49 uint16 = (v42 & v42)
                                                                            v50 := GenS3{F0: -128, F1: v4}
                                                                            _ = v49
                                                                            _ = v50
                                                                        } else {
                                                                            v51 := GenS3{F0: 127, F1: v4}
                                                                            _ = v51
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                v52 := GenS2{F0: (^(v13))}
                                                                v53 := GenS5{F0: 0}
                                                                _ = len(v3)
                                                                _ = v52
                                                                _ = v53
                                                            }
                                                            var v54 float64 = (((v23 + v23) / v23) / 3.14159)
                                                            _ = v42
                                                            _ = v54
                                                        } else {
                                                            v13 -= v13
                                                            if false {
                                                                if (false || (!true)) {
                                                                    v41 += 0
                                                                } else {
                                                                    v13 = 1
                                                                    if (!true) {
                                                                        var v55 int32 = 2147483647
                                                                        v3 = append(v3, 65535)
                                                                        if (42 < 128) {
                                                                            var v56 uint16 = 32767
                                                                            _ = v56
                                                                        }
                                                                        v1 -= (^(v1))
                                                                        _ = v55
                                                                    } else {
                                                                        v3 = append(v3, 0)
                                                                    }
                                                                    v41 = v13
                                                                }
                                                                v16 -= v16
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if ((!false) || true) {
                                                        v5 = v23
                                                        if (-1 < 83) {
                                                            fmt.Println(v5)
                                                        }
                                                    }
                                                    _ = v41
                                                } else {
                                                    v57 := GenS5{F0: 2}
                                                    v1 -= v1
                                                    v58 := GenS6{F0: true}
                                                    _ = v57
                                                    _ = v58
                                                }
                                            } else {
                                                var v59 int64 = -407
                                                fmt.Println(fi24, v4)
                                                v60 := GenS4{F0: 0, F1: fi24}
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                    }
                                    _ = fi24
                                    _ = v38
                                }
                                v61 := GenS2{F0: (v13 * v13)}
                                var v62 float32 = (-2.25 - ((((v16 / float32(v5)) + (float32(v23) * v16)) * ((float32((1e10 / v5)) - (((0.0 * v16) - ((v16 - v16) - (0.0 / (v16 / 1e10)))) / v16)) - v16)) / (v16 * (v16 * 1e10))))
                                _ = v61
                                _ = v62
                            }
                            if (((v23 / ((v23 - (v23 - 1.5)) * float64(v16))) - v23) > v23) {
                                v63 := GenS4{F0: 0, F1: (v1 % ((v1) | 1))}
                                var v64 float32 = float32(((v23 - 0.1) * v5))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v63
                                _ = v64
                            } else {
                                v5 -= v23
                            }
                            _ = v23
                        }
                    } else {
                        var v65 float32 = float32(v5)
                        _ = v65
                    }
                case "a":
                    fmt.Println(v13, v1)
                }
            }
            for fi66 := 0; fi66 < 8; fi66++ {
                v1 = (fi66 - 32767)
                v67 := GenS5{F0: 2}
                _ = fi66
                _ = v67
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
