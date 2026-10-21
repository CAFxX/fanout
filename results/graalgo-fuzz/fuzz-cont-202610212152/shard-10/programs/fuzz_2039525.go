package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint64 }
func genFn0() uint16 {
    var v5 int16 = 32767
    v6 := GenS1{F0: -32768}
    switch "foo bar" {
    case "hello":
        switch -1 {
        case 13:
            if ((false || (1 > 53)) && true) {
                switch 2147483647 {
                case -173:
                    for fi7 := 0; fi7 < 15; fi7++ {
                        if false {
                            fmt.Println(v5, v5)
                            fmt.Println(v5, fi7)
                        } else {
                            for fi8 := 0; fi8 < 17; fi8++ {
                                v9 := GenS1{F0: v5}
                                v5 += (-(v5))
                                v10 := GenS1{F0: v5}
                                _ = fi8
                                _ = v10
                                _ = v9
                            }
                            v11 := []int{fi7, fi7, 583}
                            for fi12 := 0; fi12 < 20; fi12++ {
                                switch "a" {
                                case "":
                                    if (42 == 1) {
                                        v11 = append(v11, fi12)
                                        var v13 int32 = 1
                                        _ = v13
                                    } else {
                                        switch fi12 {
                                        case 32768:
                                            fmt.Println(v5)
                                        case 9223372036854775807:
                                            if true {
                                                var v14 uint16 = 595
                                                _ = v14
                                            }
                                        case -221:
                                            if true {
                                                var v15 uint16 = 0
                                                var v16 float32 = (-2.25 / (3.14159 - float32(1e10)))
                                                v5 = (v5 & v5)
                                                _ = v15
                                                _ = v16
                                            } else {
                                                if (float32((1e10 * 1e10)) == ((float32(3.14159) + 1.5) - 0.0)) {
                                                    v17 := GenS1{F0: 1}
                                                    if false {
                                                        fmt.Println(fi7)
                                                        var v18 bool = true
                                                        _ = v18
                                                    }
                                                    _ = v17
                                                }
                                                v5 -= v5
                                                v5 = v5
                                                if (2 == 127) {
                                                    var v19 string = ""
                                                    var v20 uint = 777
                                                    var v21 int8 = -128
                                                    if len(v11) >= 2 {
                                                        v11 = v11[1:len(v11)]
                                                    }
                                                    _ = v19
                                                    _ = v20
                                                    _ = v21
                                                }
                                                v5 -= (^(v5))
                                            }
                                        default:
                                            switch fi7 {
                                            case 127:
                                                if true {
                                                    v5 = v5
                                                    v11 = append(v11, -1)
                                                    var v22 int64 = 42
                                                    _ = v22
                                                }
                                            default:
                                                fmt.Println(fi7, v5)
                                            }
                                        }
                                    }
                                case "a":
                                    if true {
                                        if ((false || true) || (int8(fi7) > -102)) {
                                            v23 := GenS1{F0: (v5 / ((1) | 1))}
                                            switch "a" {
                                            case "\n\t":
                                                fmt.Println(fi7, fi12)
                                            }
                                            switch "a" {
                                            case "hello":
                                                var v24 int = 276
                                                _ = v24
                                            case "\n\t":
                                                v5 = v5
                                            default:
                                                fmt.Println(v5)
                                            }
                                            _ = v23
                                        }
                                        v25 := GenS1{F0: v5}
                                        _ = v25
                                    }
                                case "hello":
                                    fmt.Println(v5, fi7, fi12)
                                }
                                if true {
                                    if (!(1e10 != 1.5)) {
                                        v11 = append(v11, (fi12 & (fi12 & fi12)))
                                    }
                                    _ = len(v11)
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                } else {
                                    v26 := GenS2{F0: (-(fi7)), F1: (-(fi7)), F2: 0}
                                    _ = v26
                                }
                                if false {
                                    fmt.Println(v5, fi7)
                                } else {
                                    _ = len(v11)
                                    v5 = v5
                                    v27 := GenS2{F0: fi12, F1: (^(fi7)), F2: 928}
                                    _ = v27
                                }
                                _ = fi12
                            }
                            _ = v11
                        }
                        _ = fi7
                    }
                case -9223372036854775808:
                    fmt.Println(v5)
                }
            } else {
                if true {
                    switch -279 {
                    case 1:
                        v28 := GenS1{F0: -32768}
                        _ = v28
                    }
                }
            }
        case 0:
            v5 = v5
        }
    case "\n\t":
        v5 += -1
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int64, p1 int64, p2 uint) uint64 {
    p0 = p0
    v29 := GenS2{F0: 0, F1: 42, F2: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return 2
}

func main() {
    var v1 uint8 = 1
    var v2 string = "\n\t"
    var v3 uint8 = (^(v1))
    var v4 uint16 = 2
    var v30 uint = 18446744073709551615
    if true {
        switch v2 {
        case "a":
            v3 = (v3 & (v3 / ((v1) | 1)))
        case "foo bar":
            switch 727 {
            case 128:
                v31 := []int{2147483647, -345}
                _ = v31
            }
        case "hello":
            v32 := GenS2{F0: -913, F1: 32767, F2: 55}
            _ = v32
        }
        v33 := []int{}
        _ = v33
    }
    v34 := []int{1, 1}
    v35 := GenS2{F0: 42, F1: 32767, F2: 487}
    v30 = 18446744073709551615
    v30 = (v30 & v30)
    var v36 int32 = 2147483647
    v37 := GenS2{F0: 1, F1: 42, F2: genFn1(761, 42, (v30 ^ v30))}
    if true {
        var v38 int32 = v36
        if true {
            if (!false) {
                switch -247 {
                case 495:
                    v34 = append(v34, 9223372036854775807)
                default:
                    v39 := GenS1{F0: -32768}
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 1; fi40++ {
                    v41 := GenS3{F0: v38, F1: (v2 + v2)}
                    for fi42 := 0; fi42 < 14; fi42++ {
                        _ = len(v34)
                        if len(v34) >= 2 {
                            v34 = v34[1:len(v34)]
                        }
                        v43 := GenS2{F0: fi40, F1: (fi40 + (fi42 * (fi42 >> (((v30 - v30)) & 63)))), F2: genFn1(0, 0, 332)}
                        _ = fi42
                        _ = v43
                    }
                    var v44 int8 = -1
                    _ = fi40
                    _ = v41
                    _ = v44
                }
            }
        }
        _ = v38
    }
    v1 -= 255
    switch v2 {
    case "\n\t":
        switch v2 {
        case "":
            v3 = v1
        case "foo bar":
            var v45 uint8 = 7
            _ = v45
        case "hello":
            switch v2 {
            case "a":
                if (!true) {
                    var v46 int16 = 1
                    v3 = v1
                    _ = v46
                }
            case "\n\t":
                if true {
                    v47 := GenS4{F0: 538}
                    v30 = v30
                    for fi48 := 0; fi48 < 9; fi48++ {
                        v1 = v3
                        var v49 int8 = -128
                        for fi50 := 0; fi50 < 16; fi50++ {
                            if len(v34) > 0 {
                                v34[0] = (fi50 + (fi48 / ((fi50) | 1)))
                            }
                            _ = fi50
                        }
                        _ = fi48
                        _ = v49
                    }
                    _ = v47
                } else {
                    fmt.Println(v1, v3, v36)
                }
            }
        default:
            for fi51 := 0; fi51 < 1; fi51++ {
                v3 -= v1
                v36 += 0
                _ = fi51
            }
        }
    }
    for fi52 := 0; fi52 < 12; fi52++ {
        v30 -= (v30 % (((v30 / ((255) | 1))) | 1))
        fmt.Println(fi52)
        _ = fi52
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        var v54 uint8 = (v1 >> ((v30) & 63))
        v55 := GenS2{F0: (^(fi53)), F1: 128, F2: 42}
        for fi56 := 0; fi56 < 9; fi56++ {
            var v57 uint = (v30 + v30)
            _ = fi56
            _ = v57
        }
        _ = fi53
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
