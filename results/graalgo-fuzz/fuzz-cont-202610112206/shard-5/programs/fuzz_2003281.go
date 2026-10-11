package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 string; F1 uint16 }

func main() {
    var v1 int16 = -1
    var v2 uint16 = 42
    var v3 string = "foo bar"
    var v4 float32 = 0.1
    if false {
        var v5 int64 = 42
        v6 := []int{127, -2147483648}
        var v7 uint = 727
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{612}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int = (fi9 ^ -1)
        _ = fi9
        _ = v10
    }
    if (42 <= 0) {
        if (-1 != -834) {
            v1 -= -1
            var v11 int16 = -32768
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v3)
            case "a":
                if (v1 < (v11 | (v1 >> ((18446744073709551615) & 63)))) {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 uint8 = 128
                        v8 = append(v8, -9223372036854775808)
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 uint64 = 1
                            v4 += 1.5
                            _ = fi14
                            _ = v15
                        }
                        _ = fi12
                        _ = v13
                    }
                    var v16 float64 = float64(0.0)
                    switch v3 {
                    case "hello":
                        if len(v8) > 0 {
                            v8[0] = 964
                        }
                    }
                    _ = v16
                } else {
                    var v17 int16 = 1
                    _ = v17
                }
            case "hello":
                if false {
                    var v18 uint = 1
                    var v19 int8 = -36
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v4)
                    if true {
                        v1 += (v11 / ((v1) | 1))
                    } else {
                        if (42 != -407) {
                            for fi20 := 0; fi20 < 19; fi20++ {
                                v21 := GenS1{F0: 1, F1: -128}
                                if (150 != 0) {
                                    switch "" {
                                    case "a":
                                        var v22 int64 = -1
                                        _ = v22
                                    case "":
                                        var v23 uint8 = 215
                                        _ = v23
                                    case "\n\t":
                                        v1 = (v1 >> ((65535) & 63))
                                    }
                                    for fi24 := 0; fi24 < 4; fi24++ {
                                        v25 := GenS2{F0: (65535 >= 18446744073709551615), F1: (v2 >> ((0) & 63))}
                                        _ = fi24
                                        _ = v25
                                    }
                                    v3 = ("foo bar" + "")
                                }
                                var v26 uint8 = 42
                                _ = fi20
                                _ = v21
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 1; fi27++ {
                                v4 = ((-2.25 - v4) / v4)
                                _ = fi27
                            }
                            for fi28 := 0; fi28 < 6; fi28++ {
                                var v29 string = v3
                                fmt.Println(v2)
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    if true {
                                        v1 = v1
                                        v31 := GenS3{F0: 38, F1: 318, F2: 42}
                                        _ = v31
                                    } else {
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        if true {
                                            v32 := GenS4{F0: 733, F1: (v2 > (v2 % ((v2) | 1)))}
                                            switch fi30 {
                                            case -584:
                                                switch v3 {
                                                case "":
                                                    fmt.Println(v4, v29)
                                                default:
                                                    var v33 uint16 = v2
                                                    _ = v33
                                                }
                                            default:
                                                v34 := GenS5{F0: 2}
                                                _ = v34
                                            }
                                            _ = v32
                                        } else {
                                            switch "" {
                                            case "":
                                                v35 := GenS6{F0: 42, F1: 0}
                                                _ = v35
                                            case "hello":
                                                if true {
                                                    fmt.Println(v4)
                                                    v36 := GenS7{F0: (-2.25 != ((0.0 / 0.1) / ((0.0 / 3.14159) - float64((v4 / 1e10)))))}
                                                    _ = v36
                                                }
                                            default:
                                                v37 := GenS1{F0: -8, F1: 127}
                                                _ = v37
                                            }
                                            v1 -= -32768
                                            var v38 int32 = 0
                                            _ = v38
                                        }
                                        fmt.Println(v29)
                                    }
                                    _ = fi30
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                        fmt.Println(v2, v11)
                    }
                    if len(v8) > 0 {
                        v8[0] = 1
                    }
                }
            }
            _ = v11
        }
    }
    switch v3 {
    case "hello":
        var v39 int16 = v1
        _ = v39
    default:
        v40 := GenS8{F0: v3}
        _ = v40
    }
    v41 := GenS3{F0: uint8(v2), F1: 42, F2: 936}
    var v42 int32 = -860
    for fi43 := 0; fi43 < 15; fi43++ {
        if false {
            v44 := GenS9{F0: v42, F1: (^(v42))}
            v3 = (v3 + (v3 + (v3 + ((v3 + v3) + v3))))
            fmt.Println(v1, v3)
            _ = v44
        } else {
            var v45 uint16 = (v2 - 65535)
            for fi46 := 0; fi46 < 9; fi46++ {
                v3 = v3
                switch (v3 + v3) {
                case "":
                    if len(v8) > 0 {
                        v8[0] = 0
                    }
                case "hello":
                    var v47 int32 = (-(v42))
                    _ = v47
                case "a":
                    v3 = (v3 + "foo bar")
                default:
                    var v48 float32 = (v4 * v4)
                    _ = v48
                }
                v1 = v1
                _ = fi46
            }
            var v49 bool = false
            _ = v45
            _ = v49
        }
        for fi50 := 0; fi50 < 15; fi50++ {
            if true {
                v8 = append(v8, fi50)
                v42 = v42
                v2 += v2
            } else {
                v51 := GenS8{F0: v3}
                var v52 int16 = v1
                v53 := GenS10{F0: "hello", F1: v2}
                _ = v51
                _ = v52
                _ = v53
            }
            _ = fi50
        }
        _ = len(v8)
        _ = fi43
    }
    var v54 uint64 = 4294967295
    v42 -= (v42 ^ 2147483647)
    v8 = append(v8, 426)
    var v55 bool = (-2.25 < 1e10)
    switch (v3 + "a") {
    case "foo bar":
        switch 337 {
        case 504:
            fmt.Println(v4, v2, v1)
        case -2147483648:
            if (false && true) {
                var v56 uint32 = 42
                var v57 uint32 = (v56 + (v56 ^ (v56 & v56)))
                _ = v56
                _ = v57
            }
        case 9223372036854775807:
            _ = len(v8)
        default:
            switch ("foo bar" + "\n\t") {
            case "a":
                fmt.Println(v4, v1)
            case "\n\t":
                for fi58 := 0; fi58 < 14; fi58++ {
                    var v59 bool = true
                    _ = fi58
                    _ = v59
                }
            default:
                switch (v3 + (v3 + (v3 + v3))) {
                case "hello":
                    var v60 bool = v55
                    _ = v60
                case "\n\t":
                    v4 = 0.0
                }
            }
        }
    case "hello":
        v54 = 4294967295
    default:
        v2 = (v2 ^ v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v54
    _ = v55
    _ = v8
}
