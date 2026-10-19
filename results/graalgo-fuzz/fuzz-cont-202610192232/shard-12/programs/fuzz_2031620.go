package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 int16) int16 {
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := GenS1{F0: "a", F1: 42}
        p1 = p1
        _ = fi4
        _ = v5
    }
    if true {
        v6 := GenS1{F0: "foo bar", F1: 0}
        fmt.Println(p1, p2)
        _ = v6
    }
    var v7 uint = 180
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 % ((p2) | 1))
}
func genFn1(p0 int32) string {
    p0 += p0
    if false {
        fmt.Println(p0)
    }
    p0 += 579
    _ = p0
    return "a"
}

func main() {
    v1 := []string{"a", "foo bar"}
    var v2 string = "\n\t"
    var v3 int64 = 42
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            if (fi8 < (fi9 & 9223372036854775807)) {
                switch v2 {
                case "hello":
                    _ = len(v1)
                case "":
                    if true {
                        v3 = (v3 >> ((300) & 63))
                        if (255 == 65535) {
                            var v10 uint = 59
                            _ = v10
                        } else {
                            if ((1.5 * (0.0 / (-2.25 - (((0.0 * 0.0) + (0.0 + -2.25)) * 3.14159)))) >= (-2.25 / 1.0)) {
                                switch v2 {
                                case "hello":
                                    v3 += v3
                                case "a":
                                    v11 := GenS2{F0: ("a" + "a"), F1: (!(!true)), F2: -62}
                                    _ = v11
                                default:
                                    v12 := GenS1{F0: "", F1: (v3 >> ((0) & 63))}
                                    _ = v12
                                }
                                v13 := GenS1{F0: v2, F1: (^(v3))}
                                if len(v1) > 0 {
                                    v1[0] = v2
                                }
                                switch ("a" + (v2 + v2)) {
                                case "foo bar":
                                    fmt.Println(v3)
                                case "\n\t":
                                    fmt.Println(v2, fi8)
                                case "a":
                                    var v14 float64 = 0.0
                                    _ = v14
                                }
                                _ = v13
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        if (true && true) {
                            _ = len(v1)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        if false {
                            v15 := GenS2{F0: "", F1: (true || true), F2: -111}
                            if true {
                                var v16 uint8 = 42
                                var v17 int64 = v3
                                v1 = append(v1, v2)
                                _ = v16
                                _ = v17
                            } else {
                                if false {
                                    v18 := GenS3{F0: 2147483647}
                                    _ = v18
                                }
                                if (true && (42 <= 32767)) {
                                    var v19 uint8 = 2
                                    v20 := GenS4{F0: 2147483648, F1: int16(fi8)}
                                    if (((1.5 * (0.0 - 1.5)) * -2.25) > -2.25) {
                                        fmt.Println(v19)
                                        v2 = "foo bar"
                                        if (v19 != 42) {
                                            switch v2 {
                                            case "":
                                                v21 := GenS1{F0: v2, F1: v3}
                                                _ = v21
                                            }
                                        }
                                    } else {
                                        if true {
                                            v1 = append(v1, (((v2 + (v2 + v2)) + ("a" + (v2 + (v2 + v2)))) + v2))
                                            var v22 int = 2147483647
                                            if (v19 < 2) {
                                                var v23 uint16 = 65535
                                                var v24 int8 = -108
                                                _ = v23
                                                _ = v24
                                            }
                                            _ = v22
                                        } else {
                                            v25 := GenS4{F0: 437, F1: genFn0(191, 639, 0)}
                                            fmt.Println(fi9)
                                            switch "foo bar" {
                                            case "hello":
                                                switch ("a" + v2) {
                                                case "":
                                                    _ = len(v1)
                                                case "a":
                                                    v19 += v19
                                                    var v26 int32 = -1
                                                    _ = v26
                                                case "hello":
                                                    v3 += v3
                                                default:
                                                    fmt.Println(v19)
                                                }
                                            }
                                            switch v2 {
                                            case "\n\t":
                                                switch fi8 {
                                                case 1:
                                                    v3 += (-(v3))
                                                case -2147483648:
                                                    var v27 string = "\n\t"
                                                    _ = v27
                                                case 32767:
                                                    switch fi8 {
                                                    case 0:
                                                        if (v2 >= "foo bar") {
                                                            var v28 int32 = 2147483647
                                                            v19 -= (v19 << ((65535) & 63))
                                                            var v29 bool = (-2.25 <= 0.0)
                                                            _ = v28
                                                            _ = v29
                                                        } else {
                                                            fmt.Println(fi8)
                                                        }
                                                    case 4294967295:
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v19)
                                                        case "foo bar":
                                                            v30 := GenS3{F0: 65535}
                                                            _ = v30
                                                        case "hello":
                                                            switch (fi9 | (fi9 | fi8)) {
                                                            case 9223372036854775807:
                                                                if false {
                                                                    if true {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v31 uint64 = 18446744073709551615
                                                                            _ = v31
                                                                        }
                                                                        v32 := GenS5{F0: 42}
                                                                        v19 += (v19 + v19)
                                                                        _ = v32
                                                                    }
                                                                    v1 = append(v1, "foo bar")
                                                                    v2 = v2
                                                                }
                                                            }
                                                        default:
                                                            v3 = (v3 + v3)
                                                        }
                                                    default:
                                                        v2 = (v2 + (((v2 + v2) + (v2 + ((((v2 + v2) + v2) + ("" + "foo bar")) + v2))) + ("\n\t" + v2)))
                                                    }
                                                }
                                            case "":
                                                fmt.Println(v3, v2, v19)
                                            default:
                                                fmt.Println(fi9)
                                            }
                                            if (2147483647 >= -693) {
                                                v2 = "a"
                                                switch "foo bar" {
                                                case "hello":
                                                    _ = len(v1)
                                                case "":
                                                    v33 := GenS2{F0: v2, F1: (2 < 255), F2: 100}
                                                    _ = v33
                                                case "foo bar":
                                                    v34 := GenS1{F0: "\n\t", F1: 9223372036854775807}
                                                    _ = v34
                                                }
                                                v35 := GenS6{F0: -1}
                                                _ = v35
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = v2
                                            }
                                            fmt.Println(v2)
                                            _ = v25
                                        }
                                        var v36 float32 = 0.1
                                        v37 := GenS7{F0: v19, F1: 997, F2: v3}
                                        v19 += (v19 | v19)
                                        v36 = 1.5
                                        v2 = v2
                                        var v38 int64 = (v3 / ((v3) | 1))
                                        switch (fi9 / ((32767) | 1)) {
                                        case -1:
                                            switch 32768 {
                                            case 0:
                                                v2 = v2
                                            case -205:
                                                switch 32767 {
                                                case 346:
                                                    if true {
                                                        v36 = (v36 - 1.5)
                                                        v39 := GenS8{F0: 128}
                                                        _ = v39
                                                    }
                                                case 261:
                                                    v40 := GenS9{F0: 42}
                                                    _ = v40
                                                default:
                                                    var v41 int64 = (v38 / ((1) | 1))
                                                    _ = v41
                                                }
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = (v2 + (v2 + ""))
                                                }
                                            }
                                        case -510:
                                            v42 := GenS8{F0: (v19 - (v19 << ((42) & 63)))}
                                            _ = v42
                                        case 4294967295:
                                            var v43 uint8 = v19
                                            _ = v43
                                        }
                                        v2 = (v2 + ((v2 + "hello") + v2))
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    v19 -= 234
                                    _ = v19
                                    _ = v20
                                }
                            }
                            _ = len(v1)
                            if ((!(!(!(1 != 232)))) || (42 == 2)) {
                                var v44 int16 = genFn0(1, 38, genFn0(127, 743, genFn0(1, 1, 1)))
                                _ = v44
                            } else {
                                fmt.Println(fi9, fi8)
                                if (false && (!((-1 != -1) || (65535 > 2147483647)))) {
                                    switch (fi8 >> ((4294967295) & 63)) {
                                    case 2147483647:
                                        var v45 int16 = 32767
                                        _ = v45
                                    case 577:
                                        v3 -= v3
                                    case 847:
                                        v1 = append(v1, (v2 + (v2 + "\n\t")))
                                    default:
                                        fmt.Println(v3, fi9)
                                    }
                                    if false {
                                        fmt.Println(fi9)
                                    } else {
                                        var v46 uint32 = 2147483647
                                        _ = v46
                                    }
                                }
                                var v47 uint8 = 1
                                _ = v47
                            }
                            v48 := GenS10{F0: 2, F1: -1}
                            v3 = (v3 | 1)
                            _ = v15
                            _ = v48
                        }
                    }
                default:
                    var v49 float32 = 1e10
                    _ = v49
                }
                fmt.Println(v3, fi9)
                v3 -= 42
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = fi9
        }
        _ = fi8
    }
    v50 := GenS1{F0: v2, F1: (v3 % ((v3) | 1))}
    v3 -= (v3 << ((255) & 63))
    fmt.Println(v2, v3)
    fmt.Println(v3, v3)
    v3 = (v3 ^ v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi51 := 0; fi51 < 3; fi51++ {
        var v52 float64 = 0.0
        _ = fi51
        _ = v52
    }
    v2 = v2
    v3 -= 42
    v53 := GenS8{F0: 0}
    if len(v1) > 0 {
        v1[0] = v2
    }
    v54 := GenS5{F0: 905}
    for fi55 := 0; fi55 < 14; fi55++ {
        v56 := GenS11{F0: 0}
        v3 = v3
        for fi57 := 0; fi57 < 12; fi57++ {
            fmt.Println(v2, v3)
            _ = fi57
        }
        _ = fi55
        _ = v56
    }
    v58 := GenS9{F0: 868}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v53
    _ = v54
    _ = v58
}
