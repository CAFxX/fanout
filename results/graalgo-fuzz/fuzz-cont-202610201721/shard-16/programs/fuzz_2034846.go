package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint; F1 int64; F2 uint }
type GenS9 struct { F0 int8 }

func main() {
    v1 := []uint32{556, 1, 824, 42}
    v2 := make(map[uint32]string)
    var v3 string = ("hello" + "a")
    var v4 int8 = 107
    v5 := []string{v3, ("hello" + ((v3 + ("a" + v3)) + ((v3 + v3) + ""))), "foo bar"}
    v6 := GenS1{F0: 255}
    v7 := GenS1{F0: 0}
    v4 -= (v4 * v4)
    for fi8 := 0; fi8 < 2; fi8++ {
        if true {
            var v9 int16 = 300
            switch fi8 {
            case -1:
                switch "foo bar" {
                case "a":
                    fmt.Println(fi8, fi8, v4)
                case "\n\t":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        switch (^(fi10)) {
                        case -1:
                            v11 := GenS1{F0: 42}
                            _ = v11
                        }
                        _ = fi10
                    }
                case "":
                    v12 := GenS2{F0: (fi8 + 661)}
                    _ = v12
                default:
                    var v13 uint8 = 255
                    _ = v13
                }
            case 32767:
                v3 = v3
            }
            _ = v9
        }
        v4 += (v4 / ((42) | 1))
        _ = fi8
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v3 = v3
        fmt.Println(v3, v3)
        _ = fi14
    }
    if (false && false) {
        switch 0 {
        case -964:
            v15 := GenS2{F0: 355}
            _ = v15
        case 17:
            v2[2147483648] = ("hello" + "")
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float32 = 1.5
            if false {
                v18 := GenS3{F0: (true || (((-64 == -2147483648) || true) && false))}
                _ = v18
            }
            if false {
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                _ = len(v1)
            }
            _ = fi16
            _ = v17
        }
        v4 -= -128
    }
    v19 := GenS4{F0: 2147483647, F1: 42}
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := GenS2{F0: fi20}
        switch fi20 {
        case -358:
            for fi22 := 0; fi22 < 18; fi22++ {
                if false {
                    v23 := GenS5{F0: (v4 + (v4 | (v4 / ((-17) | 1)))), F1: 65535, F2: -834}
                    var v24 uint16 = 42
                    v24 = (v24 ^ (v24 + 818))
                    v24 += v24
                    if false {
                        var v25 float32 = float32(0.1)
                        _ = v25
                    }
                    _ = v23
                    _ = v24
                } else {
                    if true {
                        if false {
                            switch "" {
                            case "\n\t":
                                var v26 float32 = -2.25
                                _ = v26
                            case "hello":
                                v3 = v3
                            default:
                                var v27 int = (fi20 ^ (fi22 << ((255) & 63)))
                                _ = v27
                            }
                            var v28 uint = 18446744073709551615
                            var v29 uint64 = 132
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS5{F0: (v4 / ((v4) | 1)), F1: 255, F2: -1}
                            _ = v30
                        }
                        v4 -= (v4 * v4)
                    } else {
                        var v31 uint32 = 408
                        fmt.Println(v3, v4)
                        var v32 int16 = 32767
                        var v33 uint16 = 624
                        switch v3 {
                        case "":
                            fmt.Println(v3, fi20)
                        case "\n\t":
                            v33 = v33
                        default:
                            if (42 < 868) {
                                if true {
                                    switch "a" {
                                    case "hello":
                                        if false {
                                            v32 = (v32 & v32)
                                        } else {
                                            var v34 uint = 33
                                            v35 := GenS3{F0: false}
                                            _ = v34
                                            _ = v35
                                        }
                                    case "foo bar":
                                        v36 := GenS6{F0: 7}
                                        _ = v36
                                    }
                                }
                                var v37 float64 = ((1.5 / 3.14159) - (1e10 * 0.0))
                                v38 := GenS6{F0: 18446744073709551615}
                                _ = v37
                                _ = v38
                            }
                        }
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v39 := GenS7{F0: 958}
                    _ = v39
                }
                v40 := GenS8{F0: 2, F1: 9223372036854775807, F2: 18446744073709551615}
                v41 := GenS9{F0: v4}
                _ = fi22
                _ = v40
                _ = v41
            }
        case 42:
            if false {
                switch fi20 {
                case -9223372036854775808:
                    var v42 uint = 1
                    _ = v42
                case -1:
                    var v43 float32 = float32((float64((1e10 + 3.14159)) / 1.0))
                    _ = v43
                default:
                    switch (-(fi20)) {
                    case 32768:
                        for fi44 := 0; fi44 < 20; fi44++ {
                            if true {
                                var v45 int = (fi20 & fi20)
                                _ = v45
                            } else {
                                v4 -= v4
                            }
                            _ = fi44
                        }
                    }
                }
                var v46 uint = 0
                _ = v46
            }
        case 1:
            var v47 uint = 65535
            _ = v47
        }
        _ = fi20
        _ = v21
    }
    if true {
        switch -463 {
        case 352:
            v48 := GenS7{F0: 21}
            _ = v48
        case 32767:
            fmt.Println(v3, v3)
        }
    } else {
        var v49 uint8 = 255
        _ = v49
    }
    if (187 != 1) {
        if (4294967295 == -737) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            delete(v2, 564)
            v50 := v2[2147483647]
            _ = v50
        } else {
            v3 = "foo bar"
            if false {
                if len(v5) > 0 {
                    v5[0] = "a"
                }
                if true {
                    var v51 int32 = 42
                    _ = v51
                }
            }
        }
    } else {
        v3 = ((v3 + "a") + "hello")
    }
    var v52 int32 = -29
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v7
}
