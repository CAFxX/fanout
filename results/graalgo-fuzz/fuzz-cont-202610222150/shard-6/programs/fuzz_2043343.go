package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int; F2 bool }
type GenS6 struct { F0 uint16; F1 uint16 }
type GenS7 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"\n\t", "hello", "hello", "\n\t"}
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 string = "a"
    v3[32767] = 222
    var v6 int8 = 1
    _ = len(v4)
    fmt.Println(v1, v5, v5)
    if false {
        v5 = v5
        for fi7 := 0; fi7 < 13; fi7++ {
            v6 = (v6 ^ (v6 - v6))
            switch 258 {
            case 32768:
                if ((fi7 & (fi7 >> ((4294967295) & 63))) > fi7) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v5 = v5
                } else {
                    if (!(0 <= -1)) {
                        switch "a" {
                        case "foo bar":
                            for fi8 := 0; fi8 < 8; fi8++ {
                                v2 = append(v2, "hello")
                                var v9 uint = 614
                                _ = fi8
                                _ = v9
                            }
                        case "\n\t":
                            fmt.Println(v5, v6)
                        }
                        if ((!true) && (false || false)) {
                            v10 := GenS1{F0: fi7}
                            if false {
                                v11 := GenS1{F0: 4294967295}
                                _ = v11
                            }
                            _ = v10
                        }
                        fmt.Println(v1, v5, v5)
                    }
                }
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(fi7)
                var v13 uint16 = 42
                v5 = ("hello" + v5)
                _ = fi12
                _ = v13
            }
            _ = fi7
        }
        switch 32767 {
        case -9223372036854775808:
            v1 -= v1
        case 42:
            if true {
                var v14 uint32 = 566
                v15 := GenS1{F0: -9223372036854775808}
                switch v5 {
                case "hello":
                    var v16 float64 = ((0.1 + 0.0) * 0.1)
                    _ = v16
                default:
                    switch ((((v5 + ("\n\t" + v5)) + v5) + v5) + v5) {
                    case "foo bar":
                        for fi17 := 0; fi17 < 14; fi17++ {
                            v14 += (v14 >> ((653) & 63))
                            if true {
                                if (v6 == v6) {
                                    for fi18 := 0; fi18 < 10; fi18++ {
                                        var v19 uint32 = 384
                                        v20 := GenS2{F0: 0, F1: ((((1 == 248) || true) || false) && (1 < v1)), F2: v14}
                                        _ = fi18
                                        _ = v19
                                        _ = v20
                                    }
                                } else {
                                    var v21 uint = 18446744073709551615
                                    switch v5 {
                                    case "\n\t":
                                        var v22 int8 = (v6 * v6)
                                        _ = v22
                                    case "a":
                                        fmt.Println(fi17, v5, fi17)
                                    case "hello":
                                        var v23 float64 = (0.1 + 0.1)
                                        _ = v23
                                    }
                                    _ = v21
                                }
                                v6 += (-(v6))
                                if (1e10 > ((0.1 - (1.5 * float32(3.14159))) / (1e10 + 0.1))) {
                                    for fi24 := 0; fi24 < 16; fi24++ {
                                        v4[fi24] = (fi17 << ((846) & 63))
                                        _ = fi24
                                    }
                                    v25 := GenS3{F0: v6, F1: 575}
                                    _ = len(v3)
                                    _ = v25
                                }
                            }
                            _ = fi17
                        }
                    case "a":
                        var v26 uint16 = 65535
                        _ = v26
                    default:
                        switch -499 {
                        case 471:
                            v1 -= (v1 * (v1 / ((v1) | 1)))
                        }
                    }
                }
                _ = v14
                _ = v15
            } else {
                if (2 == 888) {
                    switch (v5 + ("hello" + v5)) {
                    case "":
                        for fi27 := 0; fi27 < 3; fi27++ {
                            if (4294967295 < 1) {
                                var v28 uint8 = 87
                                _ = v28
                            }
                            var v29 string = v5
                            var v30 uint16 = 143
                            _ = fi27
                            _ = v29
                            _ = v30
                        }
                    case "hello":
                        for fi31 := 0; fi31 < 13; fi31++ {
                            switch v5 {
                            case "\n\t":
                                v1 += 9223372036854775807
                            case "hello":
                                var v32 string = "hello"
                                _ = v32
                            default:
                                var v33 float32 = 1e10
                                _ = v33
                            }
                            _ = fi31
                        }
                    }
                    var v34 string = "a"
                    _ = v34
                }
                var v35 int8 = v6
                _ = v35
            }
        case 459:
            delete(v3, -472)
        }
    }
    switch v5 {
    case "hello":
        switch 1 {
        case 128:
            v36 := GenS4{F0: 42}
            _ = v36
        case 1:
            v5 = "a"
        case -1:
            v1 = (^(v1))
        }
    case "\n\t":
        var v37 int64 = (v1 >> ((0) & 63))
        _ = v37
    case "a":
        var v38 float32 = ((((1e10 + 1e10) + 1e10) + 1.5) / 3.14159)
        _ = v38
    default:
        v39 := GenS5{F0: v6, F1: 2147483647, F2: true}
        _ = v39
    }
    v40 := v4[719]
    var v41 bool = (!((v6 ^ -1) >= 107))
    var v42 int = 69
    switch ("a" + v5) {
    case "foo bar":
        fmt.Println(v6)
    }
    if ((v1 >= (v1 * v1)) || v41) {
        v5 = v5
        v43 := GenS4{F0: v42}
        _ = v43
    } else {
        if v41 {
            v44 := GenS6{F0: 65535, F1: 396}
            _ = v44
        } else {
            for fi45 := 0; fi45 < 5; fi45++ {
                for fi46 := 0; fi46 < 20; fi46++ {
                    v47 := v3[(-(v42))]
                    if v41 {
                        if (v41 || false) {
                            switch (-(v42)) {
                            case 128:
                                var v48 uint32 = (v47 ^ 778)
                                _ = v48
                            default:
                                _ = len(v2)
                            }
                            var v49 int = -586
                            v47 -= 501
                            _ = v49
                        }
                    } else {
                        v1 -= (v1 | v1)
                        fmt.Println(v41, v5)
                    }
                    v50 := GenS7{F0: v41, F1: 918}
                    _ = fi46
                    _ = v47
                    _ = v50
                }
                v41 = true
                v1 += v1
                _ = fi45
            }
            v6 -= -43
        }
        _ = len(v2)
        v42 = (^(v40))
    }
    if (v41 && v41) {
        v3[(v40 * v40)] = 42
        if (v41 || (768 >= 6)) {
            if (2 > 463) {
                v51 := GenS3{F0: (^(v6)), F1: 65535}
                if v41 {
                    var v52 float32 = 3.14159
                    v4[1] = 315
                    switch "a" {
                    case "hello":
                        fmt.Println(v5)
                    case "a":
                        v41 = (!(!v41))
                    case "":
                        var v53 int16 = -910
                        _ = v53
                    default:
                        var v54 uint8 = 0
                        _ = v54
                    }
                    _ = v52
                }
                _ = v51
            }
        }
        for fi55 := 0; fi55 < 9; fi55++ {
            v5 = (v5 + "\n\t")
            var v56 float32 = 1e10
            v57 := v4[-2147483648]
            _ = fi55
            _ = v56
            _ = v57
        }
    } else {
        if ((v41 && v41) && v41) {
            if (false || (59 >= 1)) {
                var v58 bool = v41
                _ = v58
            }
            _ = len(v3)
        }
        v59 := GenS2{F0: 0, F1: v41, F2: 1}
        _ = v59
    }
    var v60 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v60
}
